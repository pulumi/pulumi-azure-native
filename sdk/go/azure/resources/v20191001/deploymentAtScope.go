


package v20191001

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type DeploymentAtScope struct {
	pulumi.CustomResourceState

	Location   pulumi.StringPtrOutput                     `pulumi:"location"`
	Name       pulumi.StringOutput                        `pulumi:"name"`
	Properties DeploymentPropertiesExtendedResponseOutput `pulumi:"properties"`
	Tags       pulumi.StringMapOutput                     `pulumi:"tags"`
	Type       pulumi.StringOutput                        `pulumi:"type"`
}


func NewDeploymentAtScope(ctx *pulumi.Context,
	name string, args *DeploymentAtScopeArgs, opts ...pulumi.ResourceOption) (*DeploymentAtScope, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Properties == nil {
		return nil, errors.New("invalid value for required argument 'Properties'")
	}
	if args.Scope == nil {
		return nil, errors.New("invalid value for required argument 'Scope'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:resources:DeploymentAtScope"),
		},
		{
			Type: pulumi.String("azure-native:resources/v20190701:DeploymentAtScope"),
		},
		{
			Type: pulumi.String("azure-native:resources/v20190801:DeploymentAtScope"),
		},
		{
			Type: pulumi.String("azure-native:resources/v20200601:DeploymentAtScope"),
		},
		{
			Type: pulumi.String("azure-native:resources/v20200801:DeploymentAtScope"),
		},
		{
			Type: pulumi.String("azure-native:resources/v20201001:DeploymentAtScope"),
		},
		{
			Type: pulumi.String("azure-native:resources/v20210101:DeploymentAtScope"),
		},
		{
			Type: pulumi.String("azure-native:resources/v20210401:DeploymentAtScope"),
		},
	})
	opts = append(opts, aliases)
	var resource DeploymentAtScope
	err := ctx.RegisterResource("azure-native:resources/v20191001:DeploymentAtScope", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetDeploymentAtScope(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DeploymentAtScopeState, opts ...pulumi.ResourceOption) (*DeploymentAtScope, error) {
	var resource DeploymentAtScope
	err := ctx.ReadResource("azure-native:resources/v20191001:DeploymentAtScope", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type deploymentAtScopeState struct {
}

type DeploymentAtScopeState struct {
}

func (DeploymentAtScopeState) ElementType() reflect.Type {
	return reflect.TypeOf((*deploymentAtScopeState)(nil)).Elem()
}

type deploymentAtScopeArgs struct {
	DeploymentName *string              `pulumi:"deploymentName"`
	Location       *string              `pulumi:"location"`
	Properties     DeploymentProperties `pulumi:"properties"`
	Scope          string               `pulumi:"scope"`
	Tags           map[string]string    `pulumi:"tags"`
}


type DeploymentAtScopeArgs struct {
	DeploymentName pulumi.StringPtrInput
	Location       pulumi.StringPtrInput
	Properties     DeploymentPropertiesInput
	Scope          pulumi.StringInput
	Tags           pulumi.StringMapInput
}

func (DeploymentAtScopeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*deploymentAtScopeArgs)(nil)).Elem()
}

type DeploymentAtScopeInput interface {
	pulumi.Input

	ToDeploymentAtScopeOutput() DeploymentAtScopeOutput
	ToDeploymentAtScopeOutputWithContext(ctx context.Context) DeploymentAtScopeOutput
}

func (*DeploymentAtScope) ElementType() reflect.Type {
	return reflect.TypeOf((**DeploymentAtScope)(nil)).Elem()
}

func (i *DeploymentAtScope) ToDeploymentAtScopeOutput() DeploymentAtScopeOutput {
	return i.ToDeploymentAtScopeOutputWithContext(context.Background())
}

func (i *DeploymentAtScope) ToDeploymentAtScopeOutputWithContext(ctx context.Context) DeploymentAtScopeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeploymentAtScopeOutput)
}

type DeploymentAtScopeOutput struct{ *pulumi.OutputState }

func (DeploymentAtScopeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DeploymentAtScope)(nil)).Elem()
}

func (o DeploymentAtScopeOutput) ToDeploymentAtScopeOutput() DeploymentAtScopeOutput {
	return o
}

func (o DeploymentAtScopeOutput) ToDeploymentAtScopeOutputWithContext(ctx context.Context) DeploymentAtScopeOutput {
	return o
}

func (o DeploymentAtScopeOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DeploymentAtScope) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

func (o DeploymentAtScopeOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *DeploymentAtScope) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o DeploymentAtScopeOutput) Properties() DeploymentPropertiesExtendedResponseOutput {
	return o.ApplyT(func(v *DeploymentAtScope) DeploymentPropertiesExtendedResponseOutput { return v.Properties }).(DeploymentPropertiesExtendedResponseOutput)
}

func (o DeploymentAtScopeOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *DeploymentAtScope) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o DeploymentAtScopeOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *DeploymentAtScope) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(DeploymentAtScopeOutput{})
}
