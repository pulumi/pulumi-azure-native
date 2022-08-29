


package v20210301preview

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

type ComputePowerAction string

const (
	ComputePowerActionStart = ComputePowerAction("Start")
	ComputePowerActionStop  = ComputePowerAction("Stop")
)

type ComputeType string

const (
	ComputeTypeAKS               = ComputeType("AKS")
	ComputeTypeAmlCompute        = ComputeType("AmlCompute")
	ComputeTypeComputeInstance   = ComputeType("ComputeInstance")
	ComputeTypeDataFactory       = ComputeType("DataFactory")
	ComputeTypeVirtualMachine    = ComputeType("VirtualMachine")
	ComputeTypeHDInsight         = ComputeType("HDInsight")
	ComputeTypeDatabricks        = ComputeType("Databricks")
	ComputeTypeDataLakeAnalytics = ComputeType("DataLakeAnalytics")
	ComputeTypeSynapseSpark      = ComputeType("SynapseSpark")
)

type ContainerType string

const (
	ContainerTypeStorageInitializer = ContainerType("StorageInitializer")
	ContainerTypeInferenceServer    = ContainerType("InferenceServer")
)

type ContentsType string

const (
	ContentsTypeAzureBlob         = ContentsType("AzureBlob")
	ContentsTypeAzureDataLakeGen1 = ContentsType("AzureDataLakeGen1")
	ContentsTypeAzureDataLakeGen2 = ContentsType("AzureDataLakeGen2")
	ContentsTypeAzureFile         = ContentsType("AzureFile")
	ContentsTypeAzureMySql        = ContentsType("AzureMySql")
	ContentsTypeAzurePostgreSql   = ContentsType("AzurePostgreSql")
	ContentsTypeAzureSqlDatabase  = ContentsType("AzureSqlDatabase")
	ContentsTypeGlusterFs         = ContentsType("GlusterFs")
)

type CredentialsType string

const (
	CredentialsTypeAccountKey       = CredentialsType("AccountKey")
	CredentialsTypeCertificate      = CredentialsType("Certificate")
	CredentialsTypeNone             = CredentialsType("None")
	CredentialsTypeSas              = CredentialsType("Sas")
	CredentialsTypeServicePrincipal = CredentialsType("ServicePrincipal")
	CredentialsTypeSqlAdmin         = CredentialsType("SqlAdmin")
)

type DataBindingMode string

const (
	DataBindingModeMount          = DataBindingMode("Mount")
	DataBindingModeDownload       = DataBindingMode("Download")
	DataBindingModeUpload         = DataBindingMode("Upload")
	DataBindingModeReadOnlyMount  = DataBindingMode("ReadOnlyMount")
	DataBindingModeReadWriteMount = DataBindingMode("ReadWriteMount")
	DataBindingModeDirect         = DataBindingMode("Direct")
	DataBindingModeEvalMount      = DataBindingMode("EvalMount")
	DataBindingModeEvalDownload   = DataBindingMode("EvalDownload")
)

type DatasetType string

const (
	DatasetTypeSimple   = DatasetType("Simple")
	DatasetTypeDataflow = DatasetType("Dataflow")
)

type DaysOfWeek string

const (
	DaysOfWeekSunday    = DaysOfWeek("Sunday")
	DaysOfWeekMonday    = DaysOfWeek("Monday")
	DaysOfWeekTuesday   = DaysOfWeek("Tuesday")
	DaysOfWeekWednesday = DaysOfWeek("Wednesday")
	DaysOfWeekThursday  = DaysOfWeek("Thursday")
	DaysOfWeekFriday    = DaysOfWeek("Friday")
	DaysOfWeekSaturday  = DaysOfWeek("Saturday")
)

type DistributionType string

const (
	DistributionTypePyTorch    = DistributionType("PyTorch")
	DistributionTypeTensorFlow = DistributionType("TensorFlow")
	DistributionTypeMpi        = DistributionType("Mpi")
)

type DockerSpecificationType string

const (
	DockerSpecificationTypeBuild = DockerSpecificationType("Build")
	DockerSpecificationTypeImage = DockerSpecificationType("Image")
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
	EndpointComputeTypeK8S            = EndpointComputeType("K8S")
	EndpointComputeTypeAzureMLCompute = EndpointComputeType("AzureMLCompute")
)

type Goal string

const (
	GoalMinimize = Goal("Minimize")
	GoalMaximize = Goal("Maximize")
)

type IdentityConfigurationType string

const (
	IdentityConfigurationTypeManaged  = IdentityConfigurationType("Managed")
	IdentityConfigurationTypeAMLToken = IdentityConfigurationType("AMLToken")
)

type ImageAnnotationType string

const (
	ImageAnnotationTypeClassification       = ImageAnnotationType("Classification")
	ImageAnnotationTypeBoundingBox          = ImageAnnotationType("BoundingBox")
	ImageAnnotationTypeInstanceSegmentation = ImageAnnotationType("InstanceSegmentation")
)

type JobType string

const (
	JobTypeCommand  = JobType("Command")
	JobTypeSweep    = JobType("Sweep")
	JobTypeLabeling = JobType("Labeling")
)

type LoadBalancerType string

const (
	LoadBalancerTypePublicIp             = LoadBalancerType("PublicIp")
	LoadBalancerTypeInternalLoadBalancer = LoadBalancerType("InternalLoadBalancer")
)

type MediaType string

const (
	MediaTypeImage = MediaType("Image")
	MediaTypeText  = MediaType("Text")
)

type OperatingSystemType string

const (
	OperatingSystemTypeLinux   = OperatingSystemType("Linux")
	OperatingSystemTypeWindows = OperatingSystemType("Windows")
)

type OriginType string

const (
	OriginTypeSynapse = OriginType("Synapse")
)

type OsType string

const (
	OsTypeLinux   = OsType("Linux")
	OsTypeWindows = OsType("Windows")
)

type PrivateEndpointServiceConnectionStatus string

const (
	PrivateEndpointServiceConnectionStatusPending      = PrivateEndpointServiceConnectionStatus("Pending")
	PrivateEndpointServiceConnectionStatusApproved     = PrivateEndpointServiceConnectionStatus("Approved")
	PrivateEndpointServiceConnectionStatusRejected     = PrivateEndpointServiceConnectionStatus("Rejected")
	PrivateEndpointServiceConnectionStatusDisconnected = PrivateEndpointServiceConnectionStatus("Disconnected")
	PrivateEndpointServiceConnectionStatusTimeout      = PrivateEndpointServiceConnectionStatus("Timeout")
)

type RecurrenceFrequency string

const (
	RecurrenceFrequencyNotSpecified = RecurrenceFrequency("NotSpecified")
	RecurrenceFrequencySecond       = RecurrenceFrequency("Second")
	RecurrenceFrequencyMinute       = RecurrenceFrequency("Minute")
	RecurrenceFrequencyHour         = RecurrenceFrequency("Hour")
	RecurrenceFrequencyDay          = RecurrenceFrequency("Day")
	RecurrenceFrequencyWeek         = RecurrenceFrequency("Week")
	RecurrenceFrequencyMonth        = RecurrenceFrequency("Month")
	RecurrenceFrequencyYear         = RecurrenceFrequency("Year")
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

type ResourceIdentityAssignment string

const (
	ResourceIdentityAssignmentSystemAssigned               = ResourceIdentityAssignment("SystemAssigned")
	ResourceIdentityAssignmentUserAssigned                 = ResourceIdentityAssignment("UserAssigned")
	ResourceIdentityAssignment_SystemAssigned_UserAssigned = ResourceIdentityAssignment("SystemAssigned,UserAssigned")
	ResourceIdentityAssignmentNone                         = ResourceIdentityAssignment("None")
)

type ResourceIdentityType string

const (
	ResourceIdentityTypeSystemAssigned               = ResourceIdentityType("SystemAssigned")
	ResourceIdentityType_SystemAssigned_UserAssigned = ResourceIdentityType("SystemAssigned,UserAssigned")
	ResourceIdentityTypeUserAssigned                 = ResourceIdentityType("UserAssigned")
	ResourceIdentityTypeNone                         = ResourceIdentityType("None")
)

func (ResourceIdentityType) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceIdentityType)(nil)).Elem()
}

func (e ResourceIdentityType) ToResourceIdentityTypeOutput() ResourceIdentityTypeOutput {
	return pulumi.ToOutput(e).(ResourceIdentityTypeOutput)
}

func (e ResourceIdentityType) ToResourceIdentityTypeOutputWithContext(ctx context.Context) ResourceIdentityTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ResourceIdentityTypeOutput)
}

func (e ResourceIdentityType) ToResourceIdentityTypePtrOutput() ResourceIdentityTypePtrOutput {
	return e.ToResourceIdentityTypePtrOutputWithContext(context.Background())
}

func (e ResourceIdentityType) ToResourceIdentityTypePtrOutputWithContext(ctx context.Context) ResourceIdentityTypePtrOutput {
	return ResourceIdentityType(e).ToResourceIdentityTypeOutputWithContext(ctx).ToResourceIdentityTypePtrOutputWithContext(ctx)
}

func (e ResourceIdentityType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ResourceIdentityType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ResourceIdentityType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ResourceIdentityType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ResourceIdentityTypeOutput struct{ *pulumi.OutputState }

func (ResourceIdentityTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceIdentityType)(nil)).Elem()
}

func (o ResourceIdentityTypeOutput) ToResourceIdentityTypeOutput() ResourceIdentityTypeOutput {
	return o
}

func (o ResourceIdentityTypeOutput) ToResourceIdentityTypeOutputWithContext(ctx context.Context) ResourceIdentityTypeOutput {
	return o
}

func (o ResourceIdentityTypeOutput) ToResourceIdentityTypePtrOutput() ResourceIdentityTypePtrOutput {
	return o.ToResourceIdentityTypePtrOutputWithContext(context.Background())
}

func (o ResourceIdentityTypeOutput) ToResourceIdentityTypePtrOutputWithContext(ctx context.Context) ResourceIdentityTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ResourceIdentityType) *ResourceIdentityType {
		return &v
	}).(ResourceIdentityTypePtrOutput)
}

func (o ResourceIdentityTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ResourceIdentityTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ResourceIdentityType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ResourceIdentityTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ResourceIdentityTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ResourceIdentityType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ResourceIdentityTypePtrOutput struct{ *pulumi.OutputState }

func (ResourceIdentityTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceIdentityType)(nil)).Elem()
}

func (o ResourceIdentityTypePtrOutput) ToResourceIdentityTypePtrOutput() ResourceIdentityTypePtrOutput {
	return o
}

func (o ResourceIdentityTypePtrOutput) ToResourceIdentityTypePtrOutputWithContext(ctx context.Context) ResourceIdentityTypePtrOutput {
	return o
}

func (o ResourceIdentityTypePtrOutput) Elem() ResourceIdentityTypeOutput {
	return o.ApplyT(func(v *ResourceIdentityType) ResourceIdentityType {
		if v != nil {
			return *v
		}
		var ret ResourceIdentityType
		return ret
	}).(ResourceIdentityTypeOutput)
}

func (o ResourceIdentityTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ResourceIdentityTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ResourceIdentityType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type ResourceIdentityTypeInput interface {
	pulumi.Input

	ToResourceIdentityTypeOutput() ResourceIdentityTypeOutput
	ToResourceIdentityTypeOutputWithContext(context.Context) ResourceIdentityTypeOutput
}

var resourceIdentityTypePtrType = reflect.TypeOf((**ResourceIdentityType)(nil)).Elem()

type ResourceIdentityTypePtrInput interface {
	pulumi.Input

	ToResourceIdentityTypePtrOutput() ResourceIdentityTypePtrOutput
	ToResourceIdentityTypePtrOutputWithContext(context.Context) ResourceIdentityTypePtrOutput
}

type resourceIdentityTypePtr string

func ResourceIdentityTypePtr(v string) ResourceIdentityTypePtrInput {
	return (*resourceIdentityTypePtr)(&v)
}

func (*resourceIdentityTypePtr) ElementType() reflect.Type {
	return resourceIdentityTypePtrType
}

func (in *resourceIdentityTypePtr) ToResourceIdentityTypePtrOutput() ResourceIdentityTypePtrOutput {
	return pulumi.ToOutput(in).(ResourceIdentityTypePtrOutput)
}

func (in *resourceIdentityTypePtr) ToResourceIdentityTypePtrOutputWithContext(ctx context.Context) ResourceIdentityTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ResourceIdentityTypePtrOutput)
}

type SamplingAlgorithm string

const (
	SamplingAlgorithmGrid     = SamplingAlgorithm("Grid")
	SamplingAlgorithmRandom   = SamplingAlgorithm("Random")
	SamplingAlgorithmBayesian = SamplingAlgorithm("Bayesian")
)

type ScaleType string

const (
	ScaleTypeAuto   = ScaleType("Auto")
	ScaleTypeManual = ScaleType("Manual")
)

type ScheduleStatus string

const (
	ScheduleStatusEnabled  = ScheduleStatus("Enabled")
	ScheduleStatusDisabled = ScheduleStatus("Disabled")
)

type SecretsType string

const (
	SecretsTypeAccountKey       = SecretsType("AccountKey")
	SecretsTypeCertificate      = SecretsType("Certificate")
	SecretsTypeNone             = SecretsType("None")
	SecretsTypeSas              = SecretsType("Sas")
	SecretsTypeServicePrincipal = SecretsType("ServicePrincipal")
	SecretsTypeSqlAdmin         = SecretsType("SqlAdmin")
)

type SshPublicAccess string

const (
	SshPublicAccessEnabled  = SshPublicAccess("Enabled")
	SshPublicAccessDisabled = SshPublicAccess("Disabled")
)

type TextAnnotationType string

const (
	TextAnnotationTypeClassification = TextAnnotationType("Classification")
)

type TriggerType string

const (
	TriggerTypeRecurrence = TriggerType("Recurrence")
	TriggerTypeCron       = TriggerType("Cron")
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
	pulumi.RegisterOutputType(ResourceIdentityTypeOutput{})
	pulumi.RegisterOutputType(ResourceIdentityTypePtrOutput{})
}
