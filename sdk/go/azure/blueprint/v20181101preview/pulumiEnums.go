


package v20181101preview

type ArtifactKind string

const (
	ArtifactKindTemplate         = ArtifactKind("template")
	ArtifactKindRoleAssignment   = ArtifactKind("roleAssignment")
	ArtifactKindPolicyAssignment = ArtifactKind("policyAssignment")
)

type AssignmentLockMode string

const (
	AssignmentLockModeNone                    = AssignmentLockMode("None")
	AssignmentLockModeAllResourcesReadOnly    = AssignmentLockMode("AllResourcesReadOnly")
	AssignmentLockModeAllResourcesDoNotDelete = AssignmentLockMode("AllResourcesDoNotDelete")
)

type BlueprintTargetScope string

const (
	// The blueprint targets a subscription during blueprint assignment.
	BlueprintTargetScopeSubscription = BlueprintTargetScope("subscription")
	// The blueprint targets a management group during blueprint assignment. This is reserved for future use.
	BlueprintTargetScopeManagementGroup = BlueprintTargetScope("managementGroup")
)

type ManagedServiceIdentityType string

const (
	ManagedServiceIdentityTypeNone           = ManagedServiceIdentityType("None")
	ManagedServiceIdentityTypeSystemAssigned = ManagedServiceIdentityType("SystemAssigned")
	ManagedServiceIdentityTypeUserAssigned   = ManagedServiceIdentityType("UserAssigned")
)

type TemplateParameterType string

const (
	TemplateParameterTypeString       = TemplateParameterType("string")
	TemplateParameterTypeArray        = TemplateParameterType("array")
	TemplateParameterTypeBool         = TemplateParameterType("bool")
	TemplateParameterTypeInt          = TemplateParameterType("int")
	TemplateParameterTypeObject       = TemplateParameterType("object")
	TemplateParameterTypeSecureObject = TemplateParameterType("secureObject")
	TemplateParameterTypeSecureString = TemplateParameterType("secureString")
)

func init() {
}
