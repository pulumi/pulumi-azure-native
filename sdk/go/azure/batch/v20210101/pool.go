


package v20210101

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Pool struct {
	pulumi.CustomResourceState

	AllocationState                 pulumi.StringOutput                            `pulumi:"allocationState"`
	AllocationStateTransitionTime   pulumi.StringOutput                            `pulumi:"allocationStateTransitionTime"`
	ApplicationLicenses             pulumi.StringArrayOutput                       `pulumi:"applicationLicenses"`
	ApplicationPackages             ApplicationPackageReferenceResponseArrayOutput `pulumi:"applicationPackages"`
	AutoScaleRun                    AutoScaleRunResponseOutput                     `pulumi:"autoScaleRun"`
	Certificates                    CertificateReferenceResponseArrayOutput        `pulumi:"certificates"`
	CreationTime                    pulumi.StringOutput                            `pulumi:"creationTime"`
	CurrentDedicatedNodes           pulumi.IntOutput                               `pulumi:"currentDedicatedNodes"`
	CurrentLowPriorityNodes         pulumi.IntOutput                               `pulumi:"currentLowPriorityNodes"`
	DeploymentConfiguration         DeploymentConfigurationResponsePtrOutput       `pulumi:"deploymentConfiguration"`
	DisplayName                     pulumi.StringPtrOutput                         `pulumi:"displayName"`
	Etag                            pulumi.StringOutput                            `pulumi:"etag"`
	Identity                        BatchPoolIdentityResponsePtrOutput             `pulumi:"identity"`
	InterNodeCommunication          pulumi.StringPtrOutput                         `pulumi:"interNodeCommunication"`
	LastModified                    pulumi.StringOutput                            `pulumi:"lastModified"`
	Metadata                        MetadataItemResponseArrayOutput                `pulumi:"metadata"`
	MountConfiguration              MountConfigurationResponseArrayOutput          `pulumi:"mountConfiguration"`
	Name                            pulumi.StringOutput                            `pulumi:"name"`
	NetworkConfiguration            NetworkConfigurationResponsePtrOutput          `pulumi:"networkConfiguration"`
	ProvisioningState               pulumi.StringOutput                            `pulumi:"provisioningState"`
	ProvisioningStateTransitionTime pulumi.StringOutput                            `pulumi:"provisioningStateTransitionTime"`
	ResizeOperationStatus           ResizeOperationStatusResponseOutput            `pulumi:"resizeOperationStatus"`
	ScaleSettings                   ScaleSettingsResponsePtrOutput                 `pulumi:"scaleSettings"`
	StartTask                       StartTaskResponsePtrOutput                     `pulumi:"startTask"`
	TaskSchedulingPolicy            TaskSchedulingPolicyResponsePtrOutput          `pulumi:"taskSchedulingPolicy"`
	TaskSlotsPerNode                pulumi.IntPtrOutput                            `pulumi:"taskSlotsPerNode"`
	Type                            pulumi.StringOutput                            `pulumi:"type"`
	UserAccounts                    UserAccountResponseArrayOutput                 `pulumi:"userAccounts"`
	VmSize                          pulumi.StringPtrOutput                         `pulumi:"vmSize"`
}


func NewPool(ctx *pulumi.Context,
	name string, args *PoolArgs, opts ...pulumi.ResourceOption) (*Pool, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccountName == nil {
		return nil, errors.New("invalid value for required argument 'AccountName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:batch:Pool"),
		},
		{
			Type: pulumi.String("azure-native:batch/v20170901:Pool"),
		},
		{
			Type: pulumi.String("azure-native:batch/v20181201:Pool"),
		},
		{
			Type: pulumi.String("azure-native:batch/v20190401:Pool"),
		},
		{
			Type: pulumi.String("azure-native:batch/v20190801:Pool"),
		},
		{
			Type: pulumi.String("azure-native:batch/v20200301:Pool"),
		},
		{
			Type: pulumi.String("azure-native:batch/v20200501:Pool"),
		},
		{
			Type: pulumi.String("azure-native:batch/v20200901:Pool"),
		},
		{
			Type: pulumi.String("azure-native:batch/v20210601:Pool"),
		},
		{
			Type: pulumi.String("azure-native:batch/v20220101:Pool"),
		},
		{
			Type: pulumi.String("azure-native:batch/v20220601:Pool"),
		},
		{
			Type: pulumi.String("azure-native:batch/v20221001:Pool"),
		},
	})
	opts = append(opts, aliases)
	var resource Pool
	err := ctx.RegisterResource("azure-native:batch/v20210101:Pool", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetPool(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PoolState, opts ...pulumi.ResourceOption) (*Pool, error) {
	var resource Pool
	err := ctx.ReadResource("azure-native:batch/v20210101:Pool", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type poolState struct {
}

type PoolState struct {
}

func (PoolState) ElementType() reflect.Type {
	return reflect.TypeOf((*poolState)(nil)).Elem()
}

type poolArgs struct {
	AccountName             string                        `pulumi:"accountName"`
	ApplicationLicenses     []string                      `pulumi:"applicationLicenses"`
	ApplicationPackages     []ApplicationPackageReference `pulumi:"applicationPackages"`
	Certificates            []CertificateReference        `pulumi:"certificates"`
	DeploymentConfiguration *DeploymentConfiguration      `pulumi:"deploymentConfiguration"`
	DisplayName             *string                       `pulumi:"displayName"`
	Identity                *BatchPoolIdentity            `pulumi:"identity"`
	InterNodeCommunication  *InterNodeCommunicationState  `pulumi:"interNodeCommunication"`
	Metadata                []MetadataItem                `pulumi:"metadata"`
	MountConfiguration      []MountConfiguration          `pulumi:"mountConfiguration"`
	NetworkConfiguration    *NetworkConfiguration         `pulumi:"networkConfiguration"`
	PoolName                *string                       `pulumi:"poolName"`
	ResourceGroupName       string                        `pulumi:"resourceGroupName"`
	ScaleSettings           *ScaleSettings                `pulumi:"scaleSettings"`
	StartTask               *StartTask                    `pulumi:"startTask"`
	TaskSchedulingPolicy    *TaskSchedulingPolicy         `pulumi:"taskSchedulingPolicy"`
	TaskSlotsPerNode        *int                          `pulumi:"taskSlotsPerNode"`
	UserAccounts            []UserAccount                 `pulumi:"userAccounts"`
	VmSize                  *string                       `pulumi:"vmSize"`
}


type PoolArgs struct {
	AccountName             pulumi.StringInput
	ApplicationLicenses     pulumi.StringArrayInput
	ApplicationPackages     ApplicationPackageReferenceArrayInput
	Certificates            CertificateReferenceArrayInput
	DeploymentConfiguration DeploymentConfigurationPtrInput
	DisplayName             pulumi.StringPtrInput
	Identity                BatchPoolIdentityPtrInput
	InterNodeCommunication  InterNodeCommunicationStatePtrInput
	Metadata                MetadataItemArrayInput
	MountConfiguration      MountConfigurationArrayInput
	NetworkConfiguration    NetworkConfigurationPtrInput
	PoolName                pulumi.StringPtrInput
	ResourceGroupName       pulumi.StringInput
	ScaleSettings           ScaleSettingsPtrInput
	StartTask               StartTaskPtrInput
	TaskSchedulingPolicy    TaskSchedulingPolicyPtrInput
	TaskSlotsPerNode        pulumi.IntPtrInput
	UserAccounts            UserAccountArrayInput
	VmSize                  pulumi.StringPtrInput
}

func (PoolArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*poolArgs)(nil)).Elem()
}

type PoolInput interface {
	pulumi.Input

	ToPoolOutput() PoolOutput
	ToPoolOutputWithContext(ctx context.Context) PoolOutput
}

func (*Pool) ElementType() reflect.Type {
	return reflect.TypeOf((**Pool)(nil)).Elem()
}

func (i *Pool) ToPoolOutput() PoolOutput {
	return i.ToPoolOutputWithContext(context.Background())
}

func (i *Pool) ToPoolOutputWithContext(ctx context.Context) PoolOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PoolOutput)
}

type PoolOutput struct{ *pulumi.OutputState }

func (PoolOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Pool)(nil)).Elem()
}

func (o PoolOutput) ToPoolOutput() PoolOutput {
	return o
}

func (o PoolOutput) ToPoolOutputWithContext(ctx context.Context) PoolOutput {
	return o
}

func (o PoolOutput) AllocationState() pulumi.StringOutput {
	return o.ApplyT(func(v *Pool) pulumi.StringOutput { return v.AllocationState }).(pulumi.StringOutput)
}

func (o PoolOutput) AllocationStateTransitionTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Pool) pulumi.StringOutput { return v.AllocationStateTransitionTime }).(pulumi.StringOutput)
}

func (o PoolOutput) ApplicationLicenses() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Pool) pulumi.StringArrayOutput { return v.ApplicationLicenses }).(pulumi.StringArrayOutput)
}

func (o PoolOutput) ApplicationPackages() ApplicationPackageReferenceResponseArrayOutput {
	return o.ApplyT(func(v *Pool) ApplicationPackageReferenceResponseArrayOutput { return v.ApplicationPackages }).(ApplicationPackageReferenceResponseArrayOutput)
}

func (o PoolOutput) AutoScaleRun() AutoScaleRunResponseOutput {
	return o.ApplyT(func(v *Pool) AutoScaleRunResponseOutput { return v.AutoScaleRun }).(AutoScaleRunResponseOutput)
}

func (o PoolOutput) Certificates() CertificateReferenceResponseArrayOutput {
	return o.ApplyT(func(v *Pool) CertificateReferenceResponseArrayOutput { return v.Certificates }).(CertificateReferenceResponseArrayOutput)
}

func (o PoolOutput) CreationTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Pool) pulumi.StringOutput { return v.CreationTime }).(pulumi.StringOutput)
}

func (o PoolOutput) CurrentDedicatedNodes() pulumi.IntOutput {
	return o.ApplyT(func(v *Pool) pulumi.IntOutput { return v.CurrentDedicatedNodes }).(pulumi.IntOutput)
}

func (o PoolOutput) CurrentLowPriorityNodes() pulumi.IntOutput {
	return o.ApplyT(func(v *Pool) pulumi.IntOutput { return v.CurrentLowPriorityNodes }).(pulumi.IntOutput)
}

func (o PoolOutput) DeploymentConfiguration() DeploymentConfigurationResponsePtrOutput {
	return o.ApplyT(func(v *Pool) DeploymentConfigurationResponsePtrOutput { return v.DeploymentConfiguration }).(DeploymentConfigurationResponsePtrOutput)
}

func (o PoolOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Pool) pulumi.StringPtrOutput { return v.DisplayName }).(pulumi.StringPtrOutput)
}

func (o PoolOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *Pool) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

func (o PoolOutput) Identity() BatchPoolIdentityResponsePtrOutput {
	return o.ApplyT(func(v *Pool) BatchPoolIdentityResponsePtrOutput { return v.Identity }).(BatchPoolIdentityResponsePtrOutput)
}

func (o PoolOutput) InterNodeCommunication() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Pool) pulumi.StringPtrOutput { return v.InterNodeCommunication }).(pulumi.StringPtrOutput)
}

func (o PoolOutput) LastModified() pulumi.StringOutput {
	return o.ApplyT(func(v *Pool) pulumi.StringOutput { return v.LastModified }).(pulumi.StringOutput)
}

func (o PoolOutput) Metadata() MetadataItemResponseArrayOutput {
	return o.ApplyT(func(v *Pool) MetadataItemResponseArrayOutput { return v.Metadata }).(MetadataItemResponseArrayOutput)
}

func (o PoolOutput) MountConfiguration() MountConfigurationResponseArrayOutput {
	return o.ApplyT(func(v *Pool) MountConfigurationResponseArrayOutput { return v.MountConfiguration }).(MountConfigurationResponseArrayOutput)
}

func (o PoolOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Pool) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o PoolOutput) NetworkConfiguration() NetworkConfigurationResponsePtrOutput {
	return o.ApplyT(func(v *Pool) NetworkConfigurationResponsePtrOutput { return v.NetworkConfiguration }).(NetworkConfigurationResponsePtrOutput)
}

func (o PoolOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *Pool) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o PoolOutput) ProvisioningStateTransitionTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Pool) pulumi.StringOutput { return v.ProvisioningStateTransitionTime }).(pulumi.StringOutput)
}

func (o PoolOutput) ResizeOperationStatus() ResizeOperationStatusResponseOutput {
	return o.ApplyT(func(v *Pool) ResizeOperationStatusResponseOutput { return v.ResizeOperationStatus }).(ResizeOperationStatusResponseOutput)
}

func (o PoolOutput) ScaleSettings() ScaleSettingsResponsePtrOutput {
	return o.ApplyT(func(v *Pool) ScaleSettingsResponsePtrOutput { return v.ScaleSettings }).(ScaleSettingsResponsePtrOutput)
}

func (o PoolOutput) StartTask() StartTaskResponsePtrOutput {
	return o.ApplyT(func(v *Pool) StartTaskResponsePtrOutput { return v.StartTask }).(StartTaskResponsePtrOutput)
}

func (o PoolOutput) TaskSchedulingPolicy() TaskSchedulingPolicyResponsePtrOutput {
	return o.ApplyT(func(v *Pool) TaskSchedulingPolicyResponsePtrOutput { return v.TaskSchedulingPolicy }).(TaskSchedulingPolicyResponsePtrOutput)
}

func (o PoolOutput) TaskSlotsPerNode() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Pool) pulumi.IntPtrOutput { return v.TaskSlotsPerNode }).(pulumi.IntPtrOutput)
}

func (o PoolOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Pool) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o PoolOutput) UserAccounts() UserAccountResponseArrayOutput {
	return o.ApplyT(func(v *Pool) UserAccountResponseArrayOutput { return v.UserAccounts }).(UserAccountResponseArrayOutput)
}

func (o PoolOutput) VmSize() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Pool) pulumi.StringPtrOutput { return v.VmSize }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(PoolOutput{})
}
