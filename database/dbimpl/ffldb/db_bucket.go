// Copyright (c) 2018-2020 The asimov developers
// Copyright (c) 2013-2017 The btcsuite developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.
package ffldb

import (
	"bytes"
	"fmt"
	"github.com/AsimovNetwork/asimov/database"
	"runtime"
)

// bucket is an internal type used to represent a collection of key/value pairs
// and implements the database.Bucket interface.
type bucket struct {
	tx *transaction
	id [4]byte
}

// Enforce bucket implements the database.Bucket interface.
var _ database.Bucket = (*bucket)(nil)

// bucketIndexKey returns the actual key to use for storing and retrieving a
// child bucket in the bucket index.  This is required because additional
// information is needed to distinguish nested buckets with the same name.
func bucketIndexKey(parentID [4]byte, key []byte) []byte {
	// The serialized bucket index key format is:
	//   <bucketindexprefix><parentbucketid><bucketname>
	indexKey := make([]byte, len(bucketIndexPrefix)+4+len(key))
	copy(indexKey, bucketIndexPrefix)
	copy(indexKey[len(bucketIndexPrefix):], parentID[:])
	copy(indexKey[len(bucketIndexPrefix)+4:], key)
	return indexKey
}

// bucketizedKey returns the actual key to use for storing and retrieving a key
// for the provided bucket ID.  This is required because bucketizing is handled
// through the use of a unique prefix per bucket.
func bucketizedKey(bucketID [4]byte, key []byte) []byte {
	// The serialized block index key format is:
	//   <bucketid><key>
	bKey := make([]byte, 4+len(key))
	copy(bKey, bucketID[:])
	copy(bKey[4:], key)
	return bKey
}

// Bucket retrieves a nested bucket with the given key.  Returns nil if
// the bucket does not exist.
//
// This function is part of the database.Bucket interface implementation.
func (b *bucket) Bucket(key []byte) database.Bucket {
	// Ensure transaction state is valid.
	if err := b.tx.checkClosed(); err != nil {
		return nil
	}

	// Attempt to fetch the ID for the child bucket.  The bucket does not
	// exist if the bucket index entry does not exist.
	childID := b.tx.fetchKey(bucketIndexKey(b.id, key))
	if childID == nil {
		return nil
	}

	childBucket := &bucket{tx: b.tx}
	copy(childBucket.id[:], childID)
	return childBucket
}

// CreateBucket creates and returns a new nested bucket with the given key.
//
// Returns the following errors as required by the interface contract:
//   - ErrBucketExists if the bucket already exists
//   - ErrBucketNameRequired if the key is empty
//   - ErrIncompatibleValue if the key is otherwise invalid for the particular
//     implementation
//   - ErrTxNotWritable if attempted against a read-only transaction
//   - ErrTxClosed if the transaction has already been closed
//
// This function is part of the database.Bucket interface implementation.
func (b *bucket) CreateBucket(key []byte) (database.Bucket, error) {
	// Ensure transaction state is valid.
	if err := b.tx.checkClosed(); err != nil {
		return nil, err
	}

	// Ensure the transaction is writable.
	if !b.tx.writable {
		str := "create bucket requires a writable database transaction"
		return nil, database.MakeError(database.ErrTxNotWritable, str, nil)
	}

	// Ensure a key was provided.
	if len(key) == 0 {
		str := "create bucket requires a key"
		return nil, database.MakeError(database.ErrBucketNameRequired, str, nil)
	}

	// Ensure bucket does not already exist.
	bidxKey := bucketIndexKey(b.id, key)
	if b.tx.hasKey(bidxKey) {
		str := "bucket already exists"
		return nil, database.MakeError(database.ErrBucketExists, str, nil)
	}

	// Find the appropriate next bucket ID to use for the new bucket.  In
	// the case of the special internal block index, keep the fixed ID.
	var childID [4]byte
	if b.id == metadataBucketID && bytes.Equal(key, blockIdxBucketName) {
		childID = blockIdxBucketID
	} else {
		var err error
		childID, err = b.tx.nextBucketID()
		if err != nil {
			return nil, err
		}
	}

	// Add the new bucket to the bucket index.
	if err := b.tx.putKey(bidxKey, childID[:]); err != nil {
		str := fmt.Sprintf("failed to create bucket with key %q", key)
		return nil, database.ConvertErr(str, err)
	}
	return &bucket{tx: b.tx, id: childID}, nil
}

// CreateBucketIfNotExists creates and returns a new nested bucket with the
// given key if it does not already exist.
//
// Returns the following errors as required by the interface contract:
//   - ErrBucketNameRequired if the key is empty
//   - ErrIncompatibleValue if the key is otherwise invalid for the particular
//     implementation
//   - ErrTxNotWritable if attempted against a read-only transaction
//   - ErrTxClosed if the transaction has already been closed
//
// This function is part of the database.Bucket interface implementation.
func (b *bucket) CreateBucketIfNotExists(key []byte) (database.Bucket, error) {
	// Ensure transaction state is valid.
	if err := b.tx.checkClosed(); err != nil {
		return nil, err
	}

	// Ensure the transaction is writable.
	if !b.tx.writable {
		str := "create bucket requires a writable database transaction"
		return nil, database.MakeError(database.ErrTxNotWritable, str, nil)
	}

	// Return existing bucket if it already exists, otherwise create it.
	if bucket := b.Bucket(key); bucket != nil {
		return bucket, nil
	}
	return b.CreateBucket(key)
}

// DeleteBucket removes a nested bucket with the given key.
//
// Returns the following errors as required by the interface contract:
//   - ErrBucketNotFound if the specified bucket does not exist
//   - ErrTxNotWritable if attempted against a read-only transaction
//   - ErrTxClosed if the transaction has already been closed
//
// This function is part of the database.Bucket interface implementation.
func (b *bucket) DeleteBucket(key []byte) error {
	// Ensure transaction state is valid.
	if err := b.tx.checkClosed(); err != nil {
		return err
	}

	// Ensure the transaction is writable.
	if !b.tx.writable {
		str := "delete bucket requires a writable database transaction"
		return database.MakeError(database.ErrTxNotWritable, str, nil)
	}

	// Attempt to fetch the ID for the child bucket.  The bucket does not
	// exist if the bucket index entry does not exist.  In the case of the
	// special internal block index, keep the fixed ID.
	bidxKey := bucketIndexKey(b.id, key)
	childID := b.tx.fetchKey(bidxKey)
	if childID == nil {
		str := fmt.Sprintf("bucket %q does not exist", key)
		return database.MakeError(database.ErrBucketNotFound, str, nil)
	}

	// Remove all nested buckets and their keys.
	childIDs := [][]byte{childID}
	for len(childIDs) > 0 {
		childID = childIDs[len(childIDs)-1]
		childIDs = childIDs[:len(childIDs)-1]

		// Delete all keys in the nested bucket.
		keyCursor := newCursor(b, childID, ctKeys)
		for ok := keyCursor.First(); ok; ok = keyCursor.Next() {
			b.tx.deleteKey(keyCursor.rawKey(), false)
		}
		cursorFinalizer(keyCursor)

		// Iterate through all nested buckets.
		bucketCursor := newCursor(b, childID, ctBuckets)
		for ok := bucketCursor.First(); ok; ok = bucketCursor.Next() {
			// Push the id of the nested bucket onto the stack for
			// the next iteration.
			childID := bucketCursor.rawValue()
			childIDs = append(childIDs, childID)

			// Remove the nested bucket from the bucket index.
			b.tx.deleteKey(bucketCursor.rawKey(), false)
		}
		cursorFinalizer(bucketCursor)
	}

	// Remove the nested bucket from the bucket index.  Any buckets nested
	// under it were already removed above.
	b.tx.deleteKey(bidxKey, true)
	return nil
}

// Cursor returns a new cursor, allowing for iteration over the bucket's
// key/value pairs and nested buckets in forward or backward order.
//
// You must seek to a position using the First, Last, or Seek functions before
// calling the Next, Prev, Key, or Value functions.  Failure to do so will
// result in the same return values as an exhausted cursor, which is false for
// the Prev and Next functions and nil for Key and Value functions.
//
// This function is part of the database.Bucket interface implementation.
func (b *bucket) Cursor() database.Cursor {
	// Ensure transaction state is valid.
	if err := b.tx.checkClosed(); err != nil {
		return &cursor{bucket: b}
	}

	// Create the cursor and setup a runtime finalizer to ensure the
	// iterators are released when the cursor is garbage collected.
	c := newCursor(b, b.id[:], ctFull)
	runtime.SetFinalizer(c, cursorFinalizer)
	return c
}

// ForEach invokes the passed function with every key/value pair in the bucket.
// This does not include nested buckets or the key/value pairs within those
// nested buckets.
//
// WARNING: It is not safe to mutate data while iterating with this method.
// Doing so may cause the underlying cursor to be invalidated and return
// unexpected keys and/or values.
//
// Returns the following errors as required by the interface contract:
//   - ErrTxClosed if the transaction has already been closed
//
// NOTE: The values returned by this function are only valid during a
// transaction.  Attempting to access them after a transaction has ended will
// likely result in an access violation.
//
// This function is part of the database.Bucket interface implementation.
func (b *bucket) ForEach(fn func(k, v []byte) error) error {
	// Ensure transaction state is valid.
	if err := b.tx.checkClosed(); err != nil {
		return err
	}

	// Invoke the callback for each cursor item.  Return the error returned
	// from the callback when it is non-nil.
	c := newCursor(b, b.id[:], ctKeys)
	defer cursorFinalizer(c)
	for ok := c.First(); ok; ok = c.Next() {
		err := fn(c.Key(), c.Value())
		if err != nil {
			return err
		}
	}

	return nil
}

// ForEachBucket invokes the passed function with the key of every nested bucket
// in the current bucket.  This does not include any nested buckets within those
// nested buckets.
//
// WARNING: It is not safe to mutate data while iterating with this method.
// Doing so may cause the underlying cursor to be invalidated and return
// unexpected keys.
//
// Returns the following errors as required by the interface contract:
//   - ErrTxClosed if the transaction has already been closed
//
// NOTE: The values returned by this function are only valid during a
// transaction.  Attempting to access them after a transaction has ended will
// likely result in an access violation.
//
// This function is part of the database.Bucket interface implementation.
func (b *bucket) ForEachBucket(fn func(k []byte) error) error {
	// Ensure transaction state is valid.
	if err := b.tx.checkClosed(); err != nil {
		return err
	}

	// Invoke the callback for each cursor item.  Return the error returned
	// from the callback when it is non-nil.
	c := newCursor(b, b.id[:], ctBuckets)
	defer cursorFinalizer(c)
	for ok := c.First(); ok; ok = c.Next() {
		err := fn(c.Key())
		if err != nil {
			return err
		}
	}

	return nil
}

// Writable returns whether or not the bucket is writable.
//
// This function is part of the database.Bucket interface implementation.
func (b *bucket) Writable() bool {
	return b.tx.writable
}

// Put saves the specified key/value pair to the bucket.  Keys that do not
// already exist are added and keys that already exist are overwritten.
//
// Returns the following errors as required by the interface contract:
//   - ErrKeyRequired if the key is empty
//   - ErrIncompatibleValue if the key is the same as an existing bucket
//   - ErrTxNotWritable if attempted against a read-only transaction
//   - ErrTxClosed if the transaction has already been closed
//
// This function is part of the database.Bucket interface implementation.
func (b *bucket) Put(key, value []byte) error {
	// Ensure transaction state is valid.
	if err := b.tx.checkClosed(); err != nil {
		return err
	}

	// Ensure the transaction is writable.
	if !b.tx.writable {
		str := "setting a key requires a writable database transaction"
		return database.MakeError(database.ErrTxNotWritable, str, nil)
	}

	// Ensure a key was provided.
	if len(key) == 0 {
		str := "put requires a key"
		return database.MakeError(database.ErrKeyRequired, str, nil)
	}

	return b.tx.putKey(bucketizedKey(b.id, key), value)
}

// Get returns the value for the given key.  Returns nil if the key does not
// exist in this bucket.  An empty slice is returned for keys that exist but
// have no value assigned.
//
// NOTE: The value returned by this function is only valid during a transaction.
// Attempting to access it after a transaction has ended results in undefined
// behavior.  Additionally, the value must NOT be modified by the caller.
//
// This function is part of the database.Bucket interface implementation.
func (b *bucket) Get(key []byte) []byte {
	// Ensure transaction state is valid.
	if err := b.tx.checkClosed(); err != nil {
		return nil
	}

	// Nothing to return if there is no key.
	if len(key) == 0 {
		return nil
	}

	return b.tx.fetchKey(bucketizedKey(b.id, key))
}

// Delete removes the specified key from the bucket.  Deleting a key that does
// not exist does not return an error.
//
// Returns the following errors as required by the interface contract:
//   - ErrKeyRequired if the key is empty
//   - ErrIncompatibleValue if the key is the same as an existing bucket
//   - ErrTxNotWritable if attempted against a read-only transaction
//   - ErrTxClosed if the transaction has already been closed
//
// This function is part of the database.Bucket interface implementation.
func (b *bucket) Delete(key []byte) error {
	// Ensure transaction state is valid.
	if err := b.tx.checkClosed(); err != nil {
		return err
	}

	// Ensure the transaction is writable.
	if !b.tx.writable {
		str := "deleting a value requires a writable database transaction"
		return database.MakeError(database.ErrTxNotWritable, str, nil)
	}

	// Nothing to do if there is no key.
	if len(key) == 0 {
		return nil
	}

	b.tx.deleteKey(bucketizedKey(b.id, key), true)
	return nil
}

