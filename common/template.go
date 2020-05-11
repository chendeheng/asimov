// Copyright (c) 2018-2020. The asimov developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.
package common

import (
	"encoding/binary"
	"errors"
	"github.com/bytom/common"
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

	// callCode for `getApprovedTemplatesCount` function.
	GetApprovedTemplatesCountByte = Hex2Bytes("ff2e9b4d")

	// callCode for `getSubmittedTemplatesCount` function.
	GetSubmittedTemplatesCountByte = Hex2Bytes("2c47086a")

	// callCode for `getApprovedTemplate` function.
	GetApprovedTemplateByte = Hex2Bytes("214b38d9")

	// callCode for `getSubmittedTemplate` function.
	GetSubmittedTemplateByte = Hex2Bytes("c17168bc")

	// callCode for `getTemplate` function.
	GetTemplateByte = Hex2Bytes("e9e89524")
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

// PackGetTemplateCountInput returns a byte slice according to given parameters
func PackGetTemplateCountInput(funcName string, category uint16) []byte {
	var input []byte
	if ContractTemplateWarehouse_GetApprovedTemplatesCountFunction() == funcName {
		input = append(input, GetApprovedTemplatesCountByte...)
	} else if ContractTemplateWarehouse_GetSubmittedTemplatesCountFunction() == funcName {
		input = append(input, GetSubmittedTemplatesCountByte...)
	}
	input = append(input, LeftPadBytes([]byte{byte(category)}, 32)...)
	return input
}

// PackGetTemplateDetailInput returns a byte slice according to given parameters
func PackGetTemplateDetailInput(funcName string, category uint16, index int64) []byte {
	var input []byte
	if ContractTemplateWarehouse_GetApprovedTemplateFunction() == funcName {
		input = append(input, GetApprovedTemplateByte...)
	} else if ContractTemplateWarehouse_GetSubmittedTemplateFunction() == funcName {
		input = append(input, GetSubmittedTemplateByte...)
	}
	input = append(input, LeftPadBytes([]byte{byte(category)}, 32)...)
	input = append(input, LeftPadBytes([]byte{byte(index)}, 32)...)
	return input
}

// PackGetTemplateInput returns a byte slice according to given parameters
func PackGetTemplateInput(category uint16, name string) []byte {
	var input []byte
	input = append(input, GetTemplateByte...)
	input = append(input, LeftPadBytes([]byte{byte(category)}, 32)...)
	input = append(input, LeftPadBytes([]byte{byte(64)}, 32)...)
	input = append(input, LeftPadBytes([]byte{byte(len(name))}, 32)...)
	input = append(input, common.RightPadBytes([]byte(name), 32)...)
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

// UnPackGetTemplatesCountResult returns template number by unpacking given byte slice
func UnPackGetTemplatesCountResult(ret []byte) (int64, error) {
	if len(ret) != 32 {
		return 0, errors.New("invalid length of ret of getOrganizationAddressesByAssets func")
	}

	return int64(ret[31]), nil
}

// UnPackGetTemplateDetailResult returns template detail by unpacking given byte slice
func UnPackGetTemplateDetailResult(ret []byte) (string, []byte, int64, uint8, uint8, uint8, uint8, error) {
	name := string(ret[256:])
	key := ret[32:64][0:32]
	createTime := int64(binary.BigEndian.Uint64(ret[64:96][len(ret[64:96])-8:]))
	approveCount := ret[96:128][len(ret[96:128])-1]
	rejectCount := ret[128:160][len(ret[128:160])-1]
	reviewers := ret[160:192][len(ret[160:192])-1]
	status := ret[192:224][len(ret[192:224])-1]

	return name, key, createTime, approveCount, rejectCount, reviewers, status, nil
}

