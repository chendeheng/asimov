// Code generated by github.com/AsimovNetwork/asimov/cmd/genesis/write_contract. DO NOT EDIT.

package common

// ContractCode identifies a kind of genesis contract.
type ContractCode string

const (
	ConsensusPOA ContractCode = "0x63000000000000000000000000000000000000006c"
	ConsensusSatoshiPlus ContractCode = "0x63000000000000000000000000000000000000006a"
	GenesisOrganization ContractCode = "0x630000000000000000000000000000000000000064"
	RegistryCenter ContractCode = "0x630000000000000000000000000000000000000065"
	TemplateWarehouse ContractCode = "0x630000000000000000000000000000000000000067"
	ValidatorCommittee ContractCode = "0x63000000000000000000000000000000000000006b"
)

// Map of ContractCode values back to their constant names for pretty printing.
var ContractCodeStrings = map[ContractCode]string{
	ConsensusPOA : "ConsensusPOA",
	ConsensusSatoshiPlus : "ConsensusSatoshiPlus",
	GenesisOrganization : "GenesisOrganization",
	RegistryCenter : "RegistryCenter",
	TemplateWarehouse : "TemplateWarehouse",
	ValidatorCommittee : "ValidatorCommittee",
}

// String returns the ErrorCode as a human-readable name.
func (c ContractCode) String() string {
	if s := ContractCodeStrings[c]; s != "" {
		return s
	}
	return "Unknown ContractCode " + string(c)
}


// Contract ConsensusPOA Definition

func ContractConsensusPOA_BatchRemoveAdminsFunction() (string) {
	return "batchRemoveAdmins"
}

func ContractConsensusPOA_GetAddressRolesMapFunction() (string) {
	return "getAddressRolesMap"
}

func ContractConsensusPOA_InitFunction() (string) {
	return "init"
}

func ContractConsensusPOA_ConfigureFunctionRoleSuperFunction() (string) {
	return "configureFunctionRoleSuper"
}

func ContractConsensusPOA_ConfigureFunctionAddressAdvancedFunction() (string) {
	return "configureFunctionAddressAdvanced"
}

func ContractConsensusPOA_ConfigureFunctionAddressSuperFunction() (string) {
	return "configureFunctionAddressSuper"
}

func ContractConsensusPOA_BatchRemoveValidatorsFunction() (string) {
	return "batchRemoveValidators"
}

func ContractConsensusPOA_ConfigureFunctionRoleFunction() (string) {
	return "configureFunctionRole"
}

func ContractConsensusPOA_ConfigureAddressRoleFunction() (string) {
	return "configureAddressRole"
}

func ContractConsensusPOA_UpdateValidatorsBlockInfoFunction() (string) {
	return "updateValidatorsBlockInfo"
}

func ContractConsensusPOA_GetSignupValidatorsFunction() (string) {
	return "getSignupValidators"
}

func ContractConsensusPOA_CanPerformFunction() (string) {
	return "canPerform"
}

func ContractConsensusPOA_ConfigureFunctionAddressFunction() (string) {
	return "configureFunctionAddress"
}

func ContractConsensusPOA_ConfigureFunctionRoleAdvancedFunction() (string) {
	return "configureFunctionRoleAdvanced"
}

func ContractConsensusPOA_GetAdminsAndValidatorsFunction() (string) {
	return "getAdminsAndValidators"
}

func ContractConsensusPOA_BatchAddValidatorsFunction() (string) {
	return "batchAddValidators"
}

func ContractConsensusPOA_BatchAddAdminsFunction() (string) {
	return "batchAddAdmins"
}

func ContractConsensusPOA_TransferAssetFunction() (string) {
	return "transferAsset"
}

// Contract ConsensusSatoshiPlus Definition

func ContractConsensusSatoshiPlus_UpdateValidatorsBlockInfoFunction() (string) {
	return "updateValidatorsBlockInfo"
}

func ContractConsensusSatoshiPlus_GetSignupValidatorsFunction() (string) {
	return "getSignupValidators"
}

func ContractConsensusSatoshiPlus_SignupFunction() (string) {
	return "signup"
}

func ContractConsensusSatoshiPlus_GetValidatorBlockInfoFunction() (string) {
	return "getValidatorBlockInfo"
}

func ContractConsensusSatoshiPlus_InitFunction() (string) {
	return "init"
}

// Contract GenesisOrganization Definition

func ContractGenesisOrganization_StartProposalFunction() (string) {
	return "startProposal"
}

func ContractGenesisOrganization_ExistedFunction() (string) {
	return "existed"
}

func ContractGenesisOrganization_InitFunction() (string) {
	return "init"
}

func ContractGenesisOrganization_VoteFunction() (string) {
	return "vote"
}

// Contract RegistryCenter Definition

func ContractRegistryCenter_DeactivateTemplateFunction() (string) {
	return "deactivateTemplate"
}

func ContractRegistryCenter_RemoveAssetRestrictionFunction() (string) {
	return "removeAssetRestriction"
}

func ContractRegistryCenter_RegisterOrganizationFunction() (string) {
	return "registerOrganization"
}

func ContractRegistryCenter_UpdateOrganizationStatusFunction() (string) {
	return "updateOrganizationStatus"
}

func ContractRegistryCenter_GetOrganizationAddressesByAssetsFunction() (string) {
	return "getOrganizationAddressesByAssets"
}

func ContractRegistryCenter_BurnAssetFunction() (string) {
	return "burnAsset"
}

func ContractRegistryCenter_CreateFunction() (string) {
	return "create"
}

func ContractRegistryCenter_HasRegisteredFunction() (string) {
	return "hasRegistered"
}

func ContractRegistryCenter_RenameOrganizationFunction() (string) {
	return "renameOrganization"
}

func ContractRegistryCenter_MintAssetFunction() (string) {
	return "mintAsset"
}

func ContractRegistryCenter_GetOrganizationAddressByIdFunction() (string) {
	return "getOrganizationAddressById"
}

func ContractRegistryCenter_CloseOrganizationFunction() (string) {
	return "closeOrganization"
}

func ContractRegistryCenter_NewAssetInfoFunction() (string) {
	return "newAssetInfo"
}

func ContractRegistryCenter_FindAssetFunction() (string) {
	return "findAsset"
}

func ContractRegistryCenter_GetOrganizationIdFunction() (string) {
	return "getOrganizationId"
}

func ContractRegistryCenter_InitFunction() (string) {
	return "init"
}

func ContractRegistryCenter_GetAssetInfoByAssetIdFunction() (string) {
	return "getAssetInfoByAssetId"
}

func ContractRegistryCenter_IsRestrictedAssetFunction() (string) {
	return "isRestrictedAsset"
}

// Contract TemplateWarehouse Definition

func ContractTemplateWarehouse_EnableFunction() (string) {
	return "enable"
}

func ContractTemplateWarehouse_GetApprovedTemplateFunction() (string) {
	return "getApprovedTemplate"
}

func ContractTemplateWarehouse_GetSubmittedTemplatesCountFunction() (string) {
	return "getSubmittedTemplatesCount"
}

func ContractTemplateWarehouse_NameExistFunction() (string) {
	return "nameExist"
}

func ContractTemplateWarehouse_THRESHOLDFunction() (string) {
	return "THRESHOLD"
}

func ContractTemplateWarehouse_CreateFunction() (string) {
	return "create"
}

func ContractTemplateWarehouse_GetSubmittedTemplateFunction() (string) {
	return "getSubmittedTemplate"
}

func ContractTemplateWarehouse_InitFunction() (string) {
	return "init"
}

func ContractTemplateWarehouse_DisableFunction() (string) {
	return "disable"
}

func ContractTemplateWarehouse_GetTemplateFunction() (string) {
	return "getTemplate"
}

func ContractTemplateWarehouse_ALL_APPROVERFunction() (string) {
	return "ALL_APPROVER"
}

func ContractTemplateWarehouse_GetApprovedTemplatesCountFunction() (string) {
	return "getApprovedTemplatesCount"
}

// Contract ValidatorCommittee Definition

func ContractValidatorCommittee_GetAddressRolesMapFunction() (string) {
	return "getAddressRolesMap"
}

func ContractValidatorCommittee_GetAssetFeeListFunction() (string) {
	return "getAssetFeeList"
}

func ContractValidatorCommittee_TestGetRoundParamsFunction() (string) {
	return "testGetRoundParams"
}

func ContractValidatorCommittee_TestGetProposalDetailFunction() (string) {
	return "testGetProposalDetail"
}

func ContractValidatorCommittee_TestUpdateValidatorBlocksFunction() (string) {
	return "testUpdateValidatorBlocks"
}

func ContractValidatorCommittee_ConfigureFunctionRoleSuperFunction() (string) {
	return "configureFunctionRoleSuper"
}

func ContractValidatorCommittee_ConfigureFunctionAddressAdvancedFunction() (string) {
	return "configureFunctionAddressAdvanced"
}

func ContractValidatorCommittee_ConfigureFunctionAddressSuperFunction() (string) {
	return "configureFunctionAddressSuper"
}

func ContractValidatorCommittee_ConfigureFunctionRoleFunction() (string) {
	return "configureFunctionRole"
}

func ContractValidatorCommittee_ConfigureAddressRoleFunction() (string) {
	return "configureAddressRole"
}

func ContractValidatorCommittee_StartNewRoundFunction() (string) {
	return "startNewRound"
}

func ContractValidatorCommittee_GetRoundLastUpdateHeightFunction() (string) {
	return "getRoundLastUpdateHeight"
}

func ContractValidatorCommittee_TestUpdateRoundLengthFunction() (string) {
	return "testUpdateRoundLength"
}

func ContractValidatorCommittee_SignupFunction() (string) {
	return "signup"
}

func ContractValidatorCommittee_CanPerformFunction() (string) {
	return "canPerform"
}

func ContractValidatorCommittee_ConfigureFunctionAddressFunction() (string) {
	return "configureFunctionAddress"
}

func ContractValidatorCommittee_StartProposalFunction() (string) {
	return "startProposal"
}

func ContractValidatorCommittee_ConfigureFunctionRoleAdvancedFunction() (string) {
	return "configureFunctionRoleAdvanced"
}

func ContractValidatorCommittee_InitFunction() (string) {
	return "init"
}

func ContractValidatorCommittee_VoteFunction() (string) {
	return "vote"
}