


package v20191001

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type DeploymentAtTenantScope struct {
	pulumi.CustomResourceState

	Location   pulumi.StringPtrOutput                     `pulumi:"location"`
	Name       pulumi.StringOutput                        `pulumi:"name"`
	Properties DeploymentPropertiesExtendedResponseOutput `pulumi:"properties"`
	Tags       pulumi.StringMapOutput                     `pulumi:"tags"`
	Type       pulumi.StringOutput                        `pulumi:"type"`
}


func NewDeploymentAtTenantScope(ctx *pulumi.Context,
	name string, args *DeploymentAtTenantScopeArgs, opts ...pulumi.ResourceOption) (*DeploymentAtTenantScope, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Properties == nil {
		return nil, errors.New("invalid value for required argument 'Properties'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:resources/v20191001:DeploymentAtTenantScope"),
		},
		{
			Type: pulumi.String("azure-native:resources:DeploymentAtTenantScope"),
		},
		{
			Type: pulumi.String("azure-nextgen:resources:DeploymentAtTenantScope"),
		},
		{
			Type: pulumi.String("azure-native:resources/v20190701:DeploymentAtTenantScope"),
		},
		{
			Type: pulumi.String("azure-nextgen:resources/v20190701:DeploymentAtTenantScope"),
		},
		{
			Type: pulumi.String("azure-native:resources/v20190801:DeploymentAtTenantScope"),
		},
		{
			Type: pulumi.String("azure-nextgen:resources/v20190801:DeploymentAtTenantScope"),
		},
		{
			Type: pulumi.String("azure-native:resources/v20200601:DeploymentAtTenantScope"),
		},
		{
			Type: pulumi.String("azure-nextgen:resources/v20200601:DeploymentAtTenantScope"),
		},
		{
			Type: pulumi.String("azure-native:resources/v20200801:DeploymentAtTenantScope"),
		},
		{
			Type: pulumi.String("azure-nextgen:resources/v20200801:DeploymentAtTenantScope"),
		},
		{
			Type: pulumi.String("azure-native:resources/v20201001:DeploymentAtTenantScope"),
		},
		{
			Type: pulumi.String("azure-nextgen:resources/v20201001:DeploymentAtTenantScope"),
		},
		{
			Type: pulumi.String("azure-native:resources/v20210101:DeploymentAtTenantScope"),
		},
		{
			Type: pulumi.String("azure-nextgen:resources/v20210101:DeploymentAtTenantScope"),
		},
		{
			Type: pulumi.String("azure-native:resources/v20210401:DeploymentAtTenantScope"),
		},
		{
			Type: pulumi.String("azure-nextgen:resources/v20210401:DeploymentAtTenantScope"),
		},
	})
	opts = append(opts, aliases)
	var resource DeploymentAtTenantScope
	err := ctx.RegisterResource("azure-native:resources/v20191001:DeploymentAtTenantScope", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetDeploymentAtTenantScope(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DeploymentAtTenantScopeState, opts ...pulumi.ResourceOption) (*DeploymentAtTenantScope, error) {
	var resource DeploymentAtTenantScope
	err := ctx.ReadResource("azure-native:resources/v20191001:DeploymentAtTenantScope", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type deploymentAtTenantScopeState struct {
}

type DeploymentAtTenantScopeState struct {
}

func (DeploymentAtTenantScopeState) ElementType() reflect.Type {
	return reflect.TypeOf((*deploymentAtTenantScopeState)(nil)).Elem()
}

type deploymentAtTenantScopeArgs struct {
	DeploymentName *string              `pulumi:"deploymentName"`
	Location       *string              `pulumi:"location"`
	Properties     DeploymentProperties `pulumi:"properties"`
	Tags           map[string]string    `pulumi:"tags"`
}


type DeploymentAtTenantScopeArgs struct {
	DeploymentName pulumi.StringPtrInput
	Location       pulumi.StringPtrInput
	Properties     DeploymentPropertiesInput
	Tags           pulumi.StringMapInput
}

func (DeploymentAtTenantScopeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*deploymentAtTenantScopeArgs)(nil)).Elem()
}

type DeploymentAtTenantScopeInput interface {
	pulumi.Input

	ToDeploymentAtTenantScopeOutput() DeploymentAtTenantScopeOutput
	ToDeploymentAtTenantScopeOutputWithContext(ctx context.Context) DeploymentAtTenantScopeOutput
}

func (*DeploymentAtTenantScope) ElementType() reflect.Type {
	return reflect.TypeOf((*DeploymentAtTenantScope)(nil))
}

func (i *DeploymentAtTenantScope) ToDeploymentAtTenantScopeOutput() DeploymentAtTenantScopeOutput {
	return i.ToDeploymentAtTenantScopeOutputWithContext(context.Background())
}

func (i *DeploymentAtTenantScope) ToDeploymentAtTenantScopeOutputWithContext(ctx context.Context) DeploymentAtTenantScopeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeploymentAtTenantScopeOutput)
}

type DeploymentAtTenantScopeOutput struct{ *pulumi.OutputState }

func (DeploymentAtTenantScopeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DeploymentAtTenantScope)(nil))
}

func (o DeploymentAtTenantScopeOutput) ToDeploymentAtTenantScopeOutput() DeploymentAtTenantScopeOutput {
	return o
}

func (o DeploymentAtTenantScopeOutput) ToDeploymentAtTenantScopeOutputWithContext(ctx context.Context) DeploymentAtTenantScopeOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DeploymentAtTenantScopeInput)(nil)).Elem(), &DeploymentAtTenantScope{})
	pulumi.RegisterOutputType(DeploymentAtTenantScopeOutput{})
}
