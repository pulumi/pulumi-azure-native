


package v20200901

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
			Type: pulumi.String("azure-nextgen:batch/v20200901:Pool"),
		},
		{
			Type: pulumi.String("azure-native:batch:Pool"),
		},
		{
			Type: pulumi.String("azure-nextgen:batch:Pool"),
		},
		{
			Type: pulumi.String("azure-native:batch/v20170901:Pool"),
		},
		{
			Type: pulumi.String("azure-nextgen:batch/v20170901:Pool"),
		},
		{
			Type: pulumi.String("azure-native:batch/v20181201:Pool"),
		},
		{
			Type: pulumi.String("azure-nextgen:batch/v20181201:Pool"),
		},
		{
			Type: pulumi.String("azure-native:batch/v20190401:Pool"),
		},
		{
			Type: pulumi.String("azure-nextgen:batch/v20190401:Pool"),
		},
		{
			Type: pulumi.String("azure-native:batch/v20190801:Pool"),
		},
		{
			Type: pulumi.String("azure-nextgen:batch/v20190801:Pool"),
		},
		{
			Type: pulumi.String("azure-native:batch/v20200301:Pool"),
		},
		{
			Type: pulumi.String("azure-nextgen:batch/v20200301:Pool"),
		},
		{
			Type: pulumi.String("azure-native:batch/v20200501:Pool"),
		},
		{
			Type: pulumi.String("azure-nextgen:batch/v20200501:Pool"),
		},
		{
			Type: pulumi.String("azure-native:batch/v20210101:Pool"),
		},
		{
			Type: pulumi.String("azure-nextgen:batch/v20210101:Pool"),
		},
		{
			Type: pulumi.String("azure-native:batch/v20210601:Pool"),
		},
		{
			Type: pulumi.String("azure-nextgen:batch/v20210601:Pool"),
		},
	})
	opts = append(opts, aliases)
	var resource Pool
	err := ctx.RegisterResource("azure-native:batch/v20200901:Pool", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetPool(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PoolState, opts ...pulumi.ResourceOption) (*Pool, error) {
	var resource Pool
	err := ctx.ReadResource("azure-native:batch/v20200901:Pool", name, id, state, &resource, opts...)
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
	return reflect.TypeOf((*Pool)(nil))
}

func (i *Pool) ToPoolOutput() PoolOutput {
	return i.ToPoolOutputWithContext(context.Background())
}

func (i *Pool) ToPoolOutputWithContext(ctx context.Context) PoolOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PoolOutput)
}

type PoolOutput struct{ *pulumi.OutputState }

func (PoolOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Pool)(nil))
}

func (o PoolOutput) ToPoolOutput() PoolOutput {
	return o
}

func (o PoolOutput) ToPoolOutputWithContext(ctx context.Context) PoolOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(PoolOutput{})
}
