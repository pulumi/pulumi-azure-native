


package v20171111preview

type ArtifactKind string

const (
	ArtifactKindTemplate         = ArtifactKind("template")
	ArtifactKindRoleAssignment   = ArtifactKind("roleAssignment")
	ArtifactKindPolicyAssignment = ArtifactKind("policyAssignment")
)

type AssignmentLockMode string

const (
	AssignmentLockModeNone         = AssignmentLockMode("None")
	AssignmentLockModeAllResources = AssignmentLockMode("AllResources")
)

type BlueprintTargetScope string

const (
	BlueprintTargetScopeSubscription    = BlueprintTargetScope("subscription")
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
