


package v20181015

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Environment struct {
	pulumi.CustomResourceState

	ClaimedByUserName        pulumi.StringOutput                 `pulumi:"claimedByUserName"`
	ClaimedByUserObjectId    pulumi.StringOutput                 `pulumi:"claimedByUserObjectId"`
	ClaimedByUserPrincipalId pulumi.StringOutput                 `pulumi:"claimedByUserPrincipalId"`
	IsClaimed                pulumi.BoolOutput                   `pulumi:"isClaimed"`
	LastKnownPowerState      pulumi.StringOutput                 `pulumi:"lastKnownPowerState"`
	LatestOperationResult    LatestOperationResultResponseOutput `pulumi:"latestOperationResult"`
	Location                 pulumi.StringPtrOutput              `pulumi:"location"`
	Name                     pulumi.StringOutput                 `pulumi:"name"`
	NetworkInterface         NetworkInterfaceResponseOutput      `pulumi:"networkInterface"`
	PasswordLastReset        pulumi.StringOutput                 `pulumi:"passwordLastReset"`
	ProvisioningState        pulumi.StringPtrOutput              `pulumi:"provisioningState"`
	ResourceSets             ResourceSetResponsePtrOutput        `pulumi:"resourceSets"`
	Tags                     pulumi.StringMapOutput              `pulumi:"tags"`
	TotalUsage               pulumi.StringOutput                 `pulumi:"totalUsage"`
	Type                     pulumi.StringOutput                 `pulumi:"type"`
	UniqueIdentifier         pulumi.StringPtrOutput              `pulumi:"uniqueIdentifier"`
}


func NewEnvironment(ctx *pulumi.Context,
	name string, args *EnvironmentArgs, opts ...pulumi.ResourceOption) (*Environment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.EnvironmentSettingName == nil {
		return nil, errors.New("invalid value for required argument 'EnvironmentSettingName'")
	}
	if args.LabAccountName == nil {
		return nil, errors.New("invalid value for required argument 'LabAccountName'")
	}
	if args.LabName == nil {
		return nil, errors.New("invalid value for required argument 'LabName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:labservices:Environment"),
		},
	})
	opts = append(opts, aliases)
	var resource Environment
	err := ctx.RegisterResource("azure-native:labservices/v20181015:Environment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetEnvironment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EnvironmentState, opts ...pulumi.ResourceOption) (*Environment, error) {
	var resource Environment
	err := ctx.ReadResource("azure-native:labservices/v20181015:Environment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type environmentState struct {
}

type EnvironmentState struct {
}

func (EnvironmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*environmentState)(nil)).Elem()
}

type environmentArgs struct {
	EnvironmentName        *string           `pulumi:"environmentName"`
	EnvironmentSettingName string            `pulumi:"environmentSettingName"`
	LabAccountName         string            `pulumi:"labAccountName"`
	LabName                string            `pulumi:"labName"`
	Location               *string           `pulumi:"location"`
	ProvisioningState      *string           `pulumi:"provisioningState"`
	ResourceGroupName      string            `pulumi:"resourceGroupName"`
	ResourceSets           *ResourceSet      `pulumi:"resourceSets"`
	Tags                   map[string]string `pulumi:"tags"`
	UniqueIdentifier       *string           `pulumi:"uniqueIdentifier"`
}


type EnvironmentArgs struct {
	EnvironmentName        pulumi.StringPtrInput
	EnvironmentSettingName pulumi.StringInput
	LabAccountName         pulumi.StringInput
	LabName                pulumi.StringInput
	Location               pulumi.StringPtrInput
	ProvisioningState      pulumi.StringPtrInput
	ResourceGroupName      pulumi.StringInput
	ResourceSets           ResourceSetPtrInput
	Tags                   pulumi.StringMapInput
	UniqueIdentifier       pulumi.StringPtrInput
}

func (EnvironmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*environmentArgs)(nil)).Elem()
}

type EnvironmentInput interface {
	pulumi.Input

	ToEnvironmentOutput() EnvironmentOutput
	ToEnvironmentOutputWithContext(ctx context.Context) EnvironmentOutput
}

func (*Environment) ElementType() reflect.Type {
	return reflect.TypeOf((*Environment)(nil))
}

func (i *Environment) ToEnvironmentOutput() EnvironmentOutput {
	return i.ToEnvironmentOutputWithContext(context.Background())
}

func (i *Environment) ToEnvironmentOutputWithContext(ctx context.Context) EnvironmentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EnvironmentOutput)
}

type EnvironmentOutput struct{ *pulumi.OutputState }

func (EnvironmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Environment)(nil))
}

func (o EnvironmentOutput) ToEnvironmentOutput() EnvironmentOutput {
	return o
}

func (o EnvironmentOutput) ToEnvironmentOutputWithContext(ctx context.Context) EnvironmentOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(EnvironmentOutput{})
}
