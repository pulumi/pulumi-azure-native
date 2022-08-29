


package v20220501

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ApplicationSharingPolicy string

const (
	ApplicationSharingPolicyPersonal = ApplicationSharingPolicy("Personal")
	ApplicationSharingPolicyShared   = ApplicationSharingPolicy("Shared")
)

type BatchLoggingLevel string

const (
	BatchLoggingLevelInfo    = BatchLoggingLevel("Info")
	BatchLoggingLevelWarning = BatchLoggingLevel("Warning")
	BatchLoggingLevelDebug   = BatchLoggingLevel("Debug")
)

type BatchOutputAction string

const (
	BatchOutputActionSummaryOnly = BatchOutputAction("SummaryOnly")
	BatchOutputActionAppendRow   = BatchOutputAction("AppendRow")
)

type ClusterPurpose string

const (
	ClusterPurposeFastProd  = ClusterPurpose("FastProd")
	ClusterPurposeDenseProd = ClusterPurpose("DenseProd")
	ClusterPurposeDevTest   = ClusterPurpose("DevTest")
)

type ComputeInstanceAuthorizationType string

const (
	ComputeInstanceAuthorizationTypePersonal = ComputeInstanceAuthorizationType("personal")
)

type ComputeType string

const (
	ComputeTypeAKS               = ComputeType("AKS")
	ComputeTypeKubernetes        = ComputeType("Kubernetes")
	ComputeTypeAmlCompute        = ComputeType("AmlCompute")
	ComputeTypeComputeInstance   = ComputeType("ComputeInstance")
	ComputeTypeDataFactory       = ComputeType("DataFactory")
	ComputeTypeVirtualMachine    = ComputeType("VirtualMachine")
	ComputeTypeHDInsight         = ComputeType("HDInsight")
	ComputeTypeDatabricks        = ComputeType("Databricks")
	ComputeTypeDataLakeAnalytics = ComputeType("DataLakeAnalytics")
	ComputeTypeSynapseSpark      = ComputeType("SynapseSpark")
)

type ConnectionAuthType string

const (
	ConnectionAuthTypePAT              = ConnectionAuthType("PAT")
	ConnectionAuthTypeManagedIdentity  = ConnectionAuthType("ManagedIdentity")
	ConnectionAuthTypeUsernamePassword = ConnectionAuthType("UsernamePassword")
	ConnectionAuthTypeNone             = ConnectionAuthType("None")
	ConnectionAuthTypeSAS              = ConnectionAuthType("SAS")
)

type ConnectionCategory string

const (
	ConnectionCategoryPythonFeed        = ConnectionCategory("PythonFeed")
	ConnectionCategoryContainerRegistry = ConnectionCategory("ContainerRegistry")
	ConnectionCategoryGit               = ConnectionCategory("Git")
)

type ContainerType string

const (
	ContainerTypeStorageInitializer = ContainerType("StorageInitializer")
	ContainerTypeInferenceServer    = ContainerType("InferenceServer")
)

type CredentialsType string

const (
	CredentialsTypeAccountKey       = CredentialsType("AccountKey")
	CredentialsTypeCertificate      = CredentialsType("Certificate")
	CredentialsTypeNone             = CredentialsType("None")
	CredentialsTypeSas              = CredentialsType("Sas")
	CredentialsTypeServicePrincipal = CredentialsType("ServicePrincipal")
)

type DataType string

const (
	DataType_Uri_file   = DataType("uri_file")
	DataType_Uri_folder = DataType("uri_folder")
	DataTypeMltable     = DataType("mltable")
)

type DatastoreType string

const (
	DatastoreTypeAzureBlob         = DatastoreType("AzureBlob")
	DatastoreTypeAzureDataLakeGen1 = DatastoreType("AzureDataLakeGen1")
	DatastoreTypeAzureDataLakeGen2 = DatastoreType("AzureDataLakeGen2")
	DatastoreTypeAzureFile         = DatastoreType("AzureFile")
)

type DistributionType string

const (
	DistributionTypePyTorch    = DistributionType("PyTorch")
	DistributionTypeTensorFlow = DistributionType("TensorFlow")
	DistributionTypeMpi        = DistributionType("Mpi")
)

type EarlyTerminationPolicyType string

const (
	EarlyTerminationPolicyTypeBandit              = EarlyTerminationPolicyType("Bandit")
	EarlyTerminationPolicyTypeMedianStopping      = EarlyTerminationPolicyType("MedianStopping")
	EarlyTerminationPolicyTypeTruncationSelection = EarlyTerminationPolicyType("TruncationSelection")
)

type EncryptionStatus string

const (
	EncryptionStatusEnabled  = EncryptionStatus("Enabled")
	EncryptionStatusDisabled = EncryptionStatus("Disabled")
)

type EndpointAuthMode string

const (
	EndpointAuthModeAMLToken = EndpointAuthMode("AMLToken")
	EndpointAuthModeKey      = EndpointAuthMode("Key")
	EndpointAuthModeAADToken = EndpointAuthMode("AADToken")
)

type EndpointComputeType string

const (
	EndpointComputeTypeManaged        = EndpointComputeType("Managed")
	EndpointComputeTypeKubernetes     = EndpointComputeType("Kubernetes")
	EndpointComputeTypeAzureMLCompute = EndpointComputeType("AzureMLCompute")
)

type Goal string

const (
	GoalMinimize = Goal("Minimize")
	GoalMaximize = Goal("Maximize")
)

type IdentityConfigurationType string

const (
	IdentityConfigurationTypeManaged      = IdentityConfigurationType("Managed")
	IdentityConfigurationTypeAMLToken     = IdentityConfigurationType("AMLToken")
	IdentityConfigurationTypeUserIdentity = IdentityConfigurationType("UserIdentity")
)

type InputDeliveryMode string

const (
	InputDeliveryModeReadOnlyMount  = InputDeliveryMode("ReadOnlyMount")
	InputDeliveryModeReadWriteMount = InputDeliveryMode("ReadWriteMount")
	InputDeliveryModeDownload       = InputDeliveryMode("Download")
	InputDeliveryModeDirect         = InputDeliveryMode("Direct")
	InputDeliveryModeEvalMount      = InputDeliveryMode("EvalMount")
	InputDeliveryModeEvalDownload   = InputDeliveryMode("EvalDownload")
)

type JobInputType string

const (
	JobInputTypeLiteral       = JobInputType("literal")
	JobInputType_Uri_file     = JobInputType("uri_file")
	JobInputType_Uri_folder   = JobInputType("uri_folder")
	JobInputTypeMltable       = JobInputType("mltable")
	JobInputType_Custom_model = JobInputType("custom_model")
	JobInputType_Mlflow_model = JobInputType("mlflow_model")
	JobInputType_Triton_model = JobInputType("triton_model")
)

type JobLimitsType string

const (
	JobLimitsTypeCommand = JobLimitsType("Command")
	JobLimitsTypeSweep   = JobLimitsType("Sweep")
)

type JobOutputType string

const (
	JobOutputType_Uri_file     = JobOutputType("uri_file")
	JobOutputType_Uri_folder   = JobOutputType("uri_folder")
	JobOutputTypeMltable       = JobOutputType("mltable")
	JobOutputType_Custom_model = JobOutputType("custom_model")
	JobOutputType_Mlflow_model = JobOutputType("mlflow_model")
	JobOutputType_Triton_model = JobOutputType("triton_model")
)

type JobType string

const (
	JobTypeCommand  = JobType("Command")
	JobTypeSweep    = JobType("Sweep")
	JobTypePipeline = JobType("Pipeline")
)

type LoadBalancerType string

const (
	LoadBalancerTypePublicIp             = LoadBalancerType("PublicIp")
	LoadBalancerTypeInternalLoadBalancer = LoadBalancerType("InternalLoadBalancer")
)

type ManagedServiceIdentityType string

const (
	ManagedServiceIdentityTypeNone                         = ManagedServiceIdentityType("None")
	ManagedServiceIdentityTypeSystemAssigned               = ManagedServiceIdentityType("SystemAssigned")
	ManagedServiceIdentityTypeUserAssigned                 = ManagedServiceIdentityType("UserAssigned")
	ManagedServiceIdentityType_SystemAssigned_UserAssigned = ManagedServiceIdentityType("SystemAssigned,UserAssigned")
)

type OperatingSystemType string

const (
	OperatingSystemTypeLinux   = OperatingSystemType("Linux")
	OperatingSystemTypeWindows = OperatingSystemType("Windows")
)

type OsType string

const (
	OsTypeLinux   = OsType("Linux")
	OsTypeWindows = OsType("Windows")
)

type OutputDeliveryMode string

const (
	OutputDeliveryModeReadWriteMount = OutputDeliveryMode("ReadWriteMount")
	OutputDeliveryModeUpload         = OutputDeliveryMode("Upload")
)

type PrivateEndpointServiceConnectionStatus string

const (
	PrivateEndpointServiceConnectionStatusPending      = PrivateEndpointServiceConnectionStatus("Pending")
	PrivateEndpointServiceConnectionStatusApproved     = PrivateEndpointServiceConnectionStatus("Approved")
	PrivateEndpointServiceConnectionStatusRejected     = PrivateEndpointServiceConnectionStatus("Rejected")
	PrivateEndpointServiceConnectionStatusDisconnected = PrivateEndpointServiceConnectionStatus("Disconnected")
	PrivateEndpointServiceConnectionStatusTimeout      = PrivateEndpointServiceConnectionStatus("Timeout")
)

type PublicNetworkAccess string

const (
	PublicNetworkAccessEnabled  = PublicNetworkAccess("Enabled")
	PublicNetworkAccessDisabled = PublicNetworkAccess("Disabled")
)

type RandomSamplingAlgorithmRule string

const (
	RandomSamplingAlgorithmRuleRandom = RandomSamplingAlgorithmRule("Random")
	RandomSamplingAlgorithmRuleSobol  = RandomSamplingAlgorithmRule("Sobol")
)

type ReferenceType string

const (
	ReferenceTypeId         = ReferenceType("Id")
	ReferenceTypeDataPath   = ReferenceType("DataPath")
	ReferenceTypeOutputPath = ReferenceType("OutputPath")
)

type RemoteLoginPortPublicAccess string

const (
	RemoteLoginPortPublicAccessEnabled      = RemoteLoginPortPublicAccess("Enabled")
	RemoteLoginPortPublicAccessDisabled     = RemoteLoginPortPublicAccess("Disabled")
	RemoteLoginPortPublicAccessNotSpecified = RemoteLoginPortPublicAccess("NotSpecified")
)

type SamplingAlgorithmType string

const (
	SamplingAlgorithmTypeGrid     = SamplingAlgorithmType("Grid")
	SamplingAlgorithmTypeRandom   = SamplingAlgorithmType("Random")
	SamplingAlgorithmTypeBayesian = SamplingAlgorithmType("Bayesian")
)

type ScaleType string

const (
	ScaleTypeDefault           = ScaleType("Default")
	ScaleTypeTargetUtilization = ScaleType("TargetUtilization")
)

type SecretsType string

const (
	SecretsTypeAccountKey       = SecretsType("AccountKey")
	SecretsTypeCertificate      = SecretsType("Certificate")
	SecretsTypeSas              = SecretsType("Sas")
	SecretsTypeServicePrincipal = SecretsType("ServicePrincipal")
)

type ServiceDataAccessAuthIdentity string

const (
	// Do not use any identity for service data access.
	ServiceDataAccessAuthIdentityNone = ServiceDataAccessAuthIdentity("None")
	// Use the system assigned managed identity of the Workspace to authenticate service data access.
	ServiceDataAccessAuthIdentityWorkspaceSystemAssignedIdentity = ServiceDataAccessAuthIdentity("WorkspaceSystemAssignedIdentity")
	// Use the user assigned managed identity of the Workspace to authenticate service data access.
	ServiceDataAccessAuthIdentityWorkspaceUserAssignedIdentity = ServiceDataAccessAuthIdentity("WorkspaceUserAssignedIdentity")
)

type SkuTier string

const (
	SkuTierFree     = SkuTier("Free")
	SkuTierBasic    = SkuTier("Basic")
	SkuTierStandard = SkuTier("Standard")
	SkuTierPremium  = SkuTier("Premium")
)

func (SkuTier) ElementType() reflect.Type {
	return reflect.TypeOf((*SkuTier)(nil)).Elem()
}

func (e SkuTier) ToSkuTierOutput() SkuTierOutput {
	return pulumi.ToOutput(e).(SkuTierOutput)
}

func (e SkuTier) ToSkuTierOutputWithContext(ctx context.Context) SkuTierOutput {
	return pulumi.ToOutputWithContext(ctx, e).(SkuTierOutput)
}

func (e SkuTier) ToSkuTierPtrOutput() SkuTierPtrOutput {
	return e.ToSkuTierPtrOutputWithContext(context.Background())
}

func (e SkuTier) ToSkuTierPtrOutputWithContext(ctx context.Context) SkuTierPtrOutput {
	return SkuTier(e).ToSkuTierOutputWithContext(ctx).ToSkuTierPtrOutputWithContext(ctx)
}

func (e SkuTier) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e SkuTier) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e SkuTier) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e SkuTier) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type SkuTierOutput struct{ *pulumi.OutputState }

func (SkuTierOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SkuTier)(nil)).Elem()
}

func (o SkuTierOutput) ToSkuTierOutput() SkuTierOutput {
	return o
}

func (o SkuTierOutput) ToSkuTierOutputWithContext(ctx context.Context) SkuTierOutput {
	return o
}

func (o SkuTierOutput) ToSkuTierPtrOutput() SkuTierPtrOutput {
	return o.ToSkuTierPtrOutputWithContext(context.Background())
}

func (o SkuTierOutput) ToSkuTierPtrOutputWithContext(ctx context.Context) SkuTierPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SkuTier) *SkuTier {
		return &v
	}).(SkuTierPtrOutput)
}

func (o SkuTierOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o SkuTierOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SkuTier) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o SkuTierOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SkuTierOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SkuTier) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type SkuTierPtrOutput struct{ *pulumi.OutputState }

func (SkuTierPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SkuTier)(nil)).Elem()
}

func (o SkuTierPtrOutput) ToSkuTierPtrOutput() SkuTierPtrOutput {
	return o
}

func (o SkuTierPtrOutput) ToSkuTierPtrOutputWithContext(ctx context.Context) SkuTierPtrOutput {
	return o
}

func (o SkuTierPtrOutput) Elem() SkuTierOutput {
	return o.ApplyT(func(v *SkuTier) SkuTier {
		if v != nil {
			return *v
		}
		var ret SkuTier
		return ret
	}).(SkuTierOutput)
}

func (o SkuTierPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SkuTierPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *SkuTier) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type SkuTierInput interface {
	pulumi.Input

	ToSkuTierOutput() SkuTierOutput
	ToSkuTierOutputWithContext(context.Context) SkuTierOutput
}

var skuTierPtrType = reflect.TypeOf((**SkuTier)(nil)).Elem()

type SkuTierPtrInput interface {
	pulumi.Input

	ToSkuTierPtrOutput() SkuTierPtrOutput
	ToSkuTierPtrOutputWithContext(context.Context) SkuTierPtrOutput
}

type skuTierPtr string

func SkuTierPtr(v string) SkuTierPtrInput {
	return (*skuTierPtr)(&v)
}

func (*skuTierPtr) ElementType() reflect.Type {
	return skuTierPtrType
}

func (in *skuTierPtr) ToSkuTierPtrOutput() SkuTierPtrOutput {
	return pulumi.ToOutput(in).(SkuTierPtrOutput)
}

func (in *skuTierPtr) ToSkuTierPtrOutputWithContext(ctx context.Context) SkuTierPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(SkuTierPtrOutput)
}

type SshPublicAccess string

const (
	SshPublicAccessEnabled  = SshPublicAccess("Enabled")
	SshPublicAccessDisabled = SshPublicAccess("Disabled")
)

type SslConfigStatus string

const (
	SslConfigStatusDisabled = SslConfigStatus("Disabled")
	SslConfigStatusEnabled  = SslConfigStatus("Enabled")
	SslConfigStatusAuto     = SslConfigStatus("Auto")
)

type ValueFormat string

const (
	ValueFormatJSON = ValueFormat("JSON")
)

type VmPriority string

const (
	VmPriorityDedicated   = VmPriority("Dedicated")
	VmPriorityLowPriority = VmPriority("LowPriority")
)

func init() {
	pulumi.RegisterOutputType(SkuTierOutput{})
	pulumi.RegisterOutputType(SkuTierPtrOutput{})
}
