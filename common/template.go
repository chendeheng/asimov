// Copyright (c) 2018-2020. The asimov developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.
package common

import (
	"errors"
)

const (
	// A standard template contract must contains
	// two functions, getTemplateInfo and initTemplate
	// Following two fields are name for these functions.
	GetTemplateInfoFunc = "getTemplateInfo"

	InitTemplateFunc = "initTemplate"

	// empty value of an address type, formatted with string
	EmptyAddressValue = "0x000000000000000000000000000000000000000000"

	// This is the standard template's abi.
	TemplateABI = "[{\"constant\": true, \"inputs\": [], \"name\": \"getTemplateInfo\", \"outputs\": [{\"name\": \"\", \"type\": \"uint16\"}, {\"name\": \"\", \"type\": \"string\"} ], \"payable\": false, \"stateMutability\": \"view\", \"type\": \"function\"}, {\"constant\": false, \"inputs\": [{\"name\": \"_category\", \"type\": \"uint16\"}, {\"name\": \"_templateName\", \"type\": \"string\"} ], \"name\": \"initTemplate\", \"outputs\": [], \"payable\": false, \"stateMutability\": \"nonpayable\", \"type\": \"function\"} ]"
)

var (
	// callCode for `getTemplateInfo` function.
	GetTemplateInfoCallCode = Hex2Bytes("2fb97c1d")

	// callCode for `canTransfer` function.
	CanTransferFuncByte = Hex2Bytes("fc588476")

	// callCode for `isRestrictedAsset` function.
	IsRestrictedAssetByte = Hex2Bytes("f0d94a13")

	// callCode for `getOrganizationAddressById` function.
	GetOrganizationAddressByIdByte = Hex2Bytes("c0bb5900")
)

// PackCanTransferInput returns a byte slice according to given parameters
func PackCanTransferInput(transferAddress []byte, assetIndex uint32) []byte {
	var input []byte
	input = append(input, CanTransferFuncByte...)
	input = append(input, LeftPadBytes(transferAddress, 32)...)
	input = append(input, LeftPadBytes([]byte{byte(assetIndex)}, 32)...)
	return input
}

// PackIsRestrictedAssetInput returns a byte slice according to given parameters
func PackIsRestrictedAssetInput(organizationId uint32, assetIndex uint32) []byte {
	var input []byte
	input = append(input, IsRestrictedAssetByte...)
	input = append(input, LeftPadBytes([]byte{byte(organizationId)}, 32)...)
	input = append(input, LeftPadBytes([]byte{byte(assetIndex)}, 32)...)
	return input
}

// PackGetOrganizationAddressByIdInput returns a byte slice according to given parameters
func PackGetOrganizationAddressByIdInput(organizationId uint32, assetIndex uint32) []byte {
	var input []byte
	input = append(input, GetOrganizationAddressByIdByte...)
	input = append(input, LeftPadBytes([]byte{byte(organizationId)}, 32)...)
	input = append(input, LeftPadBytes([]byte{byte(assetIndex)}, 32)...)
	return input
}

// UnPackBoolResult returns bool value by unpacking given byte slice
func UnPackBoolResult(ret []byte) (bool, error) {
	if len(ret) != 32 {
		return false, errors.New("invalid length of ret of canTransfer func")
	}
	var support bool
	switch ret[31] {
	case 0:
		support = false
	case 1:
		support = true
	default:
		support = false
	}
	return support, nil
}

// UnPackBoolResult returns bool value by unpacking given byte slice
func UnPackIsRestrictedAssetResult(ret []byte) (bool, bool, error) {
	if len(ret) != 64 {
		return false, false, errors.New("invalid length of ret of isRestrictedAsset func")
	}
	var existed, support bool
	if ret[31] == 1 {
		existed = true
	} else {
		existed = false
	}

	if ret[63] == 1 {
		support = true
	} else {
		support = false
	}
	return existed, support, nil
}