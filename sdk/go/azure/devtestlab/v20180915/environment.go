


package v20180915

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Environment struct {
	pulumi.CustomResourceState

	ArmTemplateDisplayName pulumi.StringPtrOutput                           `pulumi:"armTemplateDisplayName"`
	CreatedByUser          pulumi.StringOutput                              `pulumi:"createdByUser"`
	DeploymentProperties   EnvironmentDeploymentPropertiesResponsePtrOutput `pulumi:"deploymentProperties"`
	Location               pulumi.StringPtrOutput                           `pulumi:"location"`
	Name                   pulumi.StringOutput                              `pulumi:"name"`
	ProvisioningState      pulumi.StringOutput                              `pulumi:"provisioningState"`
	ResourceGroupId        pulumi.StringOutput                              `pulumi:"resourceGroupId"`
	Tags                   pulumi.StringMapOutput                           `pulumi:"tags"`
	Type                   pulumi.StringOutput                              `pulumi:"type"`
	UniqueIdentifier       pulumi.StringOutput                              `pulumi:"uniqueIdentifier"`
}


func NewEnvironment(ctx *pulumi.Context,
	name string, args *EnvironmentArgs, opts ...pulumi.ResourceOption) (*Environment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.LabName == nil {
		return nil, errors.New("invalid value for required argument 'LabName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.UserName == nil {
		return nil, errors.New("invalid value for required argument 'UserName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:devtestlab/v20180915:Environment"),
		},
		{
			Type: pulumi.String("azure-native:devtestlab:Environment"),
		},
		{
			Type: pulumi.String("azure-nextgen:devtestlab:Environment"),
		},
		{
			Type: pulumi.String("azure-native:devtestlab/v20160515:Environment"),
		},
		{
			Type: pulumi.String("azure-nextgen:devtestlab/v20160515:Environment"),
		},
	})
	opts = append(opts, aliases)
	var resource Environment
	err := ctx.RegisterResource("azure-native:devtestlab/v20180915:Environment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetEnvironment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EnvironmentState, opts ...pulumi.ResourceOption) (*Environment, error) {
	var resource Environment
	err := ctx.ReadResource("azure-native:devtestlab/v20180915:Environment", name, id, state, &resource, opts...)
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
	ArmTemplateDisplayName *string                          `pulumi:"armTemplateDisplayName"`
	DeploymentProperties   *EnvironmentDeploymentProperties `pulumi:"deploymentProperties"`
	LabName                string                           `pulumi:"labName"`
	Location               *string                          `pulumi:"location"`
	Name                   *string                          `pulumi:"name"`
	ResourceGroupName      string                           `pulumi:"resourceGroupName"`
	Tags                   map[string]string                `pulumi:"tags"`
	UserName               string                           `pulumi:"userName"`
}


type EnvironmentArgs struct {
	ArmTemplateDisplayName pulumi.StringPtrInput
	DeploymentProperties   EnvironmentDeploymentPropertiesPtrInput
	LabName                pulumi.StringInput
	Location               pulumi.StringPtrInput
	Name                   pulumi.StringPtrInput
	ResourceGroupName      pulumi.StringInput
	Tags                   pulumi.StringMapInput
	UserName               pulumi.StringInput
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
