//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armtestbase

const (
	moduleName    = "armtestbase"
	moduleVersion = "v0.5.0"
)

// Action - The action of the command.
type Action string

const (
	ActionClose     Action = "Close"
	ActionCustom    Action = "Custom"
	ActionInstall   Action = "Install"
	ActionLaunch    Action = "Launch"
	ActionUninstall Action = "Uninstall"
)

// PossibleActionValues returns the possible values for the Action const type.
func PossibleActionValues() []Action {
	return []Action{
		ActionClose,
		ActionCustom,
		ActionInstall,
		ActionLaunch,
		ActionUninstall,
	}
}

type AnalysisResultName string

const (
	AnalysisResultNameCPURegression     AnalysisResultName = "cpuRegression"
	AnalysisResultNameCPUUtilization    AnalysisResultName = "cpuUtilization"
	AnalysisResultNameMemoryRegression  AnalysisResultName = "memoryRegression"
	AnalysisResultNameMemoryUtilization AnalysisResultName = "memoryUtilization"
	AnalysisResultNameReliability       AnalysisResultName = "reliability"
	AnalysisResultNameScriptExecution   AnalysisResultName = "scriptExecution"
	AnalysisResultNameTestAnalysis      AnalysisResultName = "testAnalysis"
)

// PossibleAnalysisResultNameValues returns the possible values for the AnalysisResultName const type.
func PossibleAnalysisResultNameValues() []AnalysisResultName {
	return []AnalysisResultName{
		AnalysisResultNameCPURegression,
		AnalysisResultNameCPUUtilization,
		AnalysisResultNameMemoryRegression,
		AnalysisResultNameMemoryUtilization,
		AnalysisResultNameReliability,
		AnalysisResultNameScriptExecution,
		AnalysisResultNameTestAnalysis,
	}
}

// AnalysisResultType - Type of the Analysis Result.
type AnalysisResultType string

const (
	AnalysisResultTypeCPURegression     AnalysisResultType = "CPURegression"
	AnalysisResultTypeCPUUtilization    AnalysisResultType = "CPUUtilization"
	AnalysisResultTypeMemoryRegression  AnalysisResultType = "MemoryRegression"
	AnalysisResultTypeMemoryUtilization AnalysisResultType = "MemoryUtilization"
	AnalysisResultTypeReliability       AnalysisResultType = "Reliability"
	AnalysisResultTypeScriptExecution   AnalysisResultType = "ScriptExecution"
	AnalysisResultTypeTestAnalysis      AnalysisResultType = "TestAnalysis"
)

// PossibleAnalysisResultTypeValues returns the possible values for the AnalysisResultType const type.
func PossibleAnalysisResultTypeValues() []AnalysisResultType {
	return []AnalysisResultType{
		AnalysisResultTypeCPURegression,
		AnalysisResultTypeCPUUtilization,
		AnalysisResultTypeMemoryRegression,
		AnalysisResultTypeMemoryUtilization,
		AnalysisResultTypeReliability,
		AnalysisResultTypeScriptExecution,
		AnalysisResultTypeTestAnalysis,
	}
}

// AnalysisStatus - The analysis status.
type AnalysisStatus string

const (
	AnalysisStatusAvailable    AnalysisStatus = "Available"
	AnalysisStatusCompleted    AnalysisStatus = "Completed"
	AnalysisStatusFailed       AnalysisStatus = "Failed"
	AnalysisStatusInProgress   AnalysisStatus = "InProgress"
	AnalysisStatusNone         AnalysisStatus = "None"
	AnalysisStatusNotAvailable AnalysisStatus = "NotAvailable"
	AnalysisStatusSucceeded    AnalysisStatus = "Succeeded"
)

// PossibleAnalysisStatusValues returns the possible values for the AnalysisStatus const type.
func PossibleAnalysisStatusValues() []AnalysisStatus {
	return []AnalysisStatus{
		AnalysisStatusAvailable,
		AnalysisStatusCompleted,
		AnalysisStatusFailed,
		AnalysisStatusInProgress,
		AnalysisStatusNone,
		AnalysisStatusNotAvailable,
		AnalysisStatusSucceeded,
	}
}

// Category - The category of the failure.
type Category string

const (
	CategoryInfrastructure Category = "Infrastructure"
	CategoryNone           Category = "None"
	CategoryOSUpdate       Category = "OSUpdate"
	CategoryPackage        Category = "Package"
	CategoryUnidentified   Category = "Unidentified"
)

// PossibleCategoryValues returns the possible values for the Category const type.
func PossibleCategoryValues() []Category {
	return []Category{
		CategoryInfrastructure,
		CategoryNone,
		CategoryOSUpdate,
		CategoryPackage,
		CategoryUnidentified,
	}
}

// ContentType - The type of command content.
type ContentType string

const (
	ContentTypeFile   ContentType = "File"
	ContentTypeInline ContentType = "Inline"
	ContentTypePath   ContentType = "Path"
)

// PossibleContentTypeValues returns the possible values for the ContentType const type.
func PossibleContentTypeValues() []ContentType {
	return []ContentType{
		ContentTypeFile,
		ContentTypeInline,
		ContentTypePath,
	}
}

// CreatedByType - The type of identity that created the resource.
type CreatedByType string

const (
	CreatedByTypeApplication     CreatedByType = "Application"
	CreatedByTypeKey             CreatedByType = "Key"
	CreatedByTypeManagedIdentity CreatedByType = "ManagedIdentity"
	CreatedByTypeUser            CreatedByType = "User"
)

// PossibleCreatedByTypeValues returns the possible values for the CreatedByType const type.
func PossibleCreatedByTypeValues() []CreatedByType {
	return []CreatedByType{
		CreatedByTypeApplication,
		CreatedByTypeKey,
		CreatedByTypeManagedIdentity,
		CreatedByTypeUser,
	}
}

// ExecutionStatus - The execution status of a test.
type ExecutionStatus string

const (
	ExecutionStatusCompleted   ExecutionStatus = "Completed"
	ExecutionStatusFailed      ExecutionStatus = "Failed"
	ExecutionStatusInProgress  ExecutionStatus = "InProgress"
	ExecutionStatusIncomplete  ExecutionStatus = "Incomplete"
	ExecutionStatusNone        ExecutionStatus = "None"
	ExecutionStatusNotExecuted ExecutionStatus = "NotExecuted"
	ExecutionStatusProcessing  ExecutionStatus = "Processing"
	ExecutionStatusSucceeded   ExecutionStatus = "Succeeded"
)

// PossibleExecutionStatusValues returns the possible values for the ExecutionStatus const type.
func PossibleExecutionStatusValues() []ExecutionStatus {
	return []ExecutionStatus{
		ExecutionStatusCompleted,
		ExecutionStatusFailed,
		ExecutionStatusInProgress,
		ExecutionStatusIncomplete,
		ExecutionStatusNone,
		ExecutionStatusNotExecuted,
		ExecutionStatusProcessing,
		ExecutionStatusSucceeded,
	}
}

// Grade - The grade of a test.
type Grade string

const (
	GradeFail         Grade = "Fail"
	GradeNone         Grade = "None"
	GradeNotAvailable Grade = "NotAvailable"
	GradePass         Grade = "Pass"
)

// PossibleGradeValues returns the possible values for the Grade const type.
func PossibleGradeValues() []Grade {
	return []Grade{
		GradeFail,
		GradeNone,
		GradeNotAvailable,
		GradePass,
	}
}

type OsUpdateType string

const (
	OsUpdateTypeFeatureUpdate  OsUpdateType = "FeatureUpdate"
	OsUpdateTypeSecurityUpdate OsUpdateType = "SecurityUpdate"
)

// PossibleOsUpdateTypeValues returns the possible values for the OsUpdateType const type.
func PossibleOsUpdateTypeValues() []OsUpdateType {
	return []OsUpdateType{
		OsUpdateTypeFeatureUpdate,
		OsUpdateTypeSecurityUpdate,
	}
}

// PackageStatus - The status of the package.
type PackageStatus string

const (
	PackageStatusDeleted                   PackageStatus = "Deleted"
	PackageStatusError                     PackageStatus = "Error"
	PackageStatusPreValidationCheckPass    PackageStatus = "PreValidationCheckPass"
	PackageStatusReady                     PackageStatus = "Ready"
	PackageStatusRegistered                PackageStatus = "Registered"
	PackageStatusUnknown                   PackageStatus = "Unknown"
	PackageStatusValidatingPackage         PackageStatus = "ValidatingPackage"
	PackageStatusValidationLongerThanUsual PackageStatus = "ValidationLongerThanUsual"
	PackageStatusVerifyingPackage          PackageStatus = "VerifyingPackage"
)

// PossiblePackageStatusValues returns the possible values for the PackageStatus const type.
func PossiblePackageStatusValues() []PackageStatus {
	return []PackageStatus{
		PackageStatusDeleted,
		PackageStatusError,
		PackageStatusPreValidationCheckPass,
		PackageStatusReady,
		PackageStatusRegistered,
		PackageStatusUnknown,
		PackageStatusValidatingPackage,
		PackageStatusValidationLongerThanUsual,
		PackageStatusVerifyingPackage,
	}
}

// ProvisioningState - ARM provisioning state, see https://github.com/Azure/azure-resource-manager-rpc/blob/master/v1.0/Addendum.md#provisioningstate-property
type ProvisioningState string

const (
	ProvisioningStateCancelled ProvisioningState = "Cancelled"
	ProvisioningStateCreating  ProvisioningState = "Creating"
	ProvisioningStateDeleting  ProvisioningState = "Deleting"
	ProvisioningStateFailed    ProvisioningState = "Failed"
	ProvisioningStateSucceeded ProvisioningState = "Succeeded"
	ProvisioningStateUpdating  ProvisioningState = "Updating"
)

// PossibleProvisioningStateValues returns the possible values for the ProvisioningState const type.
func PossibleProvisioningStateValues() []ProvisioningState {
	return []ProvisioningState{
		ProvisioningStateCancelled,
		ProvisioningStateCreating,
		ProvisioningStateDeleting,
		ProvisioningStateFailed,
		ProvisioningStateSucceeded,
		ProvisioningStateUpdating,
	}
}

// Reason - The reason for unavailability of a name. Required if nameAvailable == false.
type Reason string

const (
	ReasonAlreadyExists Reason = "AlreadyExists"
	ReasonInvalid       Reason = "Invalid"
)

// PossibleReasonValues returns the possible values for the Reason const type.
func PossibleReasonValues() []Reason {
	return []Reason{
		ReasonAlreadyExists,
		ReasonInvalid,
	}
}

// TestAnalysisStatus - The status of the analysis.
type TestAnalysisStatus string

const (
	TestAnalysisStatusAnalyzing TestAnalysisStatus = "Analyzing"
	TestAnalysisStatusCompleted TestAnalysisStatus = "Completed"
	TestAnalysisStatusFailed    TestAnalysisStatus = "Failed"
	TestAnalysisStatusNone      TestAnalysisStatus = "None"
)

// PossibleTestAnalysisStatusValues returns the possible values for the TestAnalysisStatus const type.
func PossibleTestAnalysisStatusValues() []TestAnalysisStatus {
	return []TestAnalysisStatus{
		TestAnalysisStatusAnalyzing,
		TestAnalysisStatusCompleted,
		TestAnalysisStatusFailed,
		TestAnalysisStatusNone,
	}
}

// TestStatus - The status of a test.
type TestStatus string

const (
	TestStatusCompleted               TestStatus = "Completed"
	TestStatusDataProcessing          TestStatus = "DataProcessing"
	TestStatusInfrastructureFailure   TestStatus = "InfrastructureFailure"
	TestStatusNone                    TestStatus = "None"
	TestStatusTestAndUpdateFailure    TestStatus = "TestAndUpdateFailure"
	TestStatusTestExecutionInProgress TestStatus = "TestExecutionInProgress"
	TestStatusTestFailure             TestStatus = "TestFailure"
	TestStatusUpdateFailure           TestStatus = "UpdateFailure"
)

// PossibleTestStatusValues returns the possible values for the TestStatus const type.
func PossibleTestStatusValues() []TestStatus {
	return []TestStatus{
		TestStatusCompleted,
		TestStatusDataProcessing,
		TestStatusInfrastructureFailure,
		TestStatusNone,
		TestStatusTestAndUpdateFailure,
		TestStatusTestExecutionInProgress,
		TestStatusTestFailure,
		TestStatusUpdateFailure,
	}
}

// TestType - The test type.
type TestType string

const (
	TestTypeFunctionalTest TestType = "FunctionalTest"
	TestTypeOutOfBoxTest   TestType = "OutOfBoxTest"
)

// PossibleTestTypeValues returns the possible values for the TestType const type.
func PossibleTestTypeValues() []TestType {
	return []TestType{
		TestTypeFunctionalTest,
		TestTypeOutOfBoxTest,
	}
}

// Tier - The tier of this particular SKU.
type Tier string

const (
	TierStandard Tier = "Standard"
)

// PossibleTierValues returns the possible values for the Tier const type.
func PossibleTierValues() []Tier {
	return []Tier{
		TierStandard,
	}
}

// Type - The type of this release (OS update).
type Type string

const (
	TypeFeatureUpdate  Type = "FeatureUpdate"
	TypeSecurityUpdate Type = "SecurityUpdate"
)

// PossibleTypeValues returns the possible values for the Type const type.
func PossibleTypeValues() []Type {
	return []Type{
		TypeFeatureUpdate,
		TypeSecurityUpdate,
	}
}

// ValidationRunStatus - The status of the validation run of the package.
type ValidationRunStatus string

const (
	ValidationRunStatusFailed  ValidationRunStatus = "Failed"
	ValidationRunStatusPassed  ValidationRunStatus = "Passed"
	ValidationRunStatusPending ValidationRunStatus = "Pending"
	ValidationRunStatusUnknown ValidationRunStatus = "Unknown"
)

// PossibleValidationRunStatusValues returns the possible values for the ValidationRunStatus const type.
func PossibleValidationRunStatusValues() []ValidationRunStatus {
	return []ValidationRunStatus{
		ValidationRunStatusFailed,
		ValidationRunStatusPassed,
		ValidationRunStatusPending,
		ValidationRunStatusUnknown,
	}
}
