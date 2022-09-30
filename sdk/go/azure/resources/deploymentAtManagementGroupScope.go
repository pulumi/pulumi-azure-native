


package resources

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type DeploymentAtManagementGroupScope struct {
	pulumi.CustomResourceState

	Location   pulumi.StringPtrOutput                     `pulumi:"location"`
	Name       pulumi.StringOutput                        `pulumi:"name"`
	Properties DeploymentPropertiesExtendedResponseOutput `pulumi:"properties"`
	Tags       pulumi.StringMapOutput                     `pulumi:"tags"`
	Type       pulumi.StringOutput                        `pulumi:"type"`
}


func NewDeploymentAtManagementGroupScope(ctx *pulumi.Context,
	name string, args *DeploymentAtManagementGroupScopeArgs, opts ...pulumi.ResourceOption) (*DeploymentAtManagementGroupScope, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.GroupId == nil {
		return nil, errors.New("invalid value for required argument 'GroupId'")
	}
	if args.Properties == nil {
		return nil, errors.New("invalid value for required argument 'Properties'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:resources/v20190501:DeploymentAtManagementGroupScope"),
		},
		{
			Type: pulumi.String("azure-native:resources/v20190510:DeploymentAtManagementGroupScope"),
		},
		{
			Type: pulumi.String("azure-native:resources/v20190701:DeploymentAtManagementGroupScope"),
		},
		{
			Type: pulumi.String("azure-native:resources/v20190801:DeploymentAtManagementGroupScope"),
		},
		{
			Type: pulumi.String("azure-native:resources/v20191001:DeploymentAtManagementGroupScope"),
		},
		{
			Type: pulumi.String("azure-native:resources/v20200601:DeploymentAtManagementGroupScope"),
		},
		{
			Type: pulumi.String("azure-native:resources/v20200801:DeploymentAtManagementGroupScope"),
		},
		{
			Type: pulumi.String("azure-native:resources/v20201001:DeploymentAtManagementGroupScope"),
		},
		{
			Type: pulumi.String("azure-native:resources/v20210101:DeploymentAtManagementGroupScope"),
		},
		{
			Type: pulumi.String("azure-native:resources/v20210401:DeploymentAtManagementGroupScope"),
		},
	})
	opts = append(opts, aliases)
	var resource DeploymentAtManagementGroupScope
	err := ctx.RegisterResource("azure-native:resources:DeploymentAtManagementGroupScope", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetDeploymentAtManagementGroupScope(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DeploymentAtManagementGroupScopeState, opts ...pulumi.ResourceOption) (*DeploymentAtManagementGroupScope, error) {
	var resource DeploymentAtManagementGroupScope
	err := ctx.ReadResource("azure-native:resources:DeploymentAtManagementGroupScope", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type deploymentAtManagementGroupScopeState struct {
}

type DeploymentAtManagementGroupScopeState struct {
}

func (DeploymentAtManagementGroupScopeState) ElementType() reflect.Type {
	return reflect.TypeOf((*deploymentAtManagementGroupScopeState)(nil)).Elem()
}

type deploymentAtManagementGroupScopeArgs struct {
	DeploymentName *string              `pulumi:"deploymentName"`
	GroupId        string               `pulumi:"groupId"`
	Location       *string              `pulumi:"location"`
	Properties     DeploymentProperties `pulumi:"properties"`
	Tags           map[string]string    `pulumi:"tags"`
}


type DeploymentAtManagementGroupScopeArgs struct {
	DeploymentName pulumi.StringPtrInput
	GroupId        pulumi.StringInput
	Location       pulumi.StringPtrInput
	Properties     DeploymentPropertiesInput
	Tags           pulumi.StringMapInput
}

func (DeploymentAtManagementGroupScopeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*deploymentAtManagementGroupScopeArgs)(nil)).Elem()
}

type DeploymentAtManagementGroupScopeInput interface {
	pulumi.Input

	ToDeploymentAtManagementGroupScopeOutput() DeploymentAtManagementGroupScopeOutput
	ToDeploymentAtManagementGroupScopeOutputWithContext(ctx context.Context) DeploymentAtManagementGroupScopeOutput
}

func (*DeploymentAtManagementGroupScope) ElementType() reflect.Type {
	return reflect.TypeOf((**DeploymentAtManagementGroupScope)(nil)).Elem()
}

func (i *DeploymentAtManagementGroupScope) ToDeploymentAtManagementGroupScopeOutput() DeploymentAtManagementGroupScopeOutput {
	return i.ToDeploymentAtManagementGroupScopeOutputWithContext(context.Background())
}

func (i *DeploymentAtManagementGroupScope) ToDeploymentAtManagementGroupScopeOutputWithContext(ctx context.Context) DeploymentAtManagementGroupScopeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeploymentAtManagementGroupScopeOutput)
}

type DeploymentAtManagementGroupScopeOutput struct{ *pulumi.OutputState }

func (DeploymentAtManagementGroupScopeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DeploymentAtManagementGroupScope)(nil)).Elem()
}

func (o DeploymentAtManagementGroupScopeOutput) ToDeploymentAtManagementGroupScopeOutput() DeploymentAtManagementGroupScopeOutput {
	return o
}

func (o DeploymentAtManagementGroupScopeOutput) ToDeploymentAtManagementGroupScopeOutputWithContext(ctx context.Context) DeploymentAtManagementGroupScopeOutput {
	return o
}

func (o DeploymentAtManagementGroupScopeOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DeploymentAtManagementGroupScope) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

func (o DeploymentAtManagementGroupScopeOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *DeploymentAtManagementGroupScope) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o DeploymentAtManagementGroupScopeOutput) Properties() DeploymentPropertiesExtendedResponseOutput {
	return o.ApplyT(func(v *DeploymentAtManagementGroupScope) DeploymentPropertiesExtendedResponseOutput {
		return v.Properties
	}).(DeploymentPropertiesExtendedResponseOutput)
}

func (o DeploymentAtManagementGroupScopeOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *DeploymentAtManagementGroupScope) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o DeploymentAtManagementGroupScopeOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *DeploymentAtManagementGroupScope) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(DeploymentAtManagementGroupScopeOutput{})
}
