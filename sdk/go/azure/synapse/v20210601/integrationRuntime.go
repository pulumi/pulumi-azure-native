


package v20210601

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type IntegrationRuntime struct {
	pulumi.CustomResourceState

	Etag       pulumi.StringOutput `pulumi:"etag"`
	Name       pulumi.StringOutput `pulumi:"name"`
	Properties pulumi.AnyOutput    `pulumi:"properties"`
	Type       pulumi.StringOutput `pulumi:"type"`
}


func NewIntegrationRuntime(ctx *pulumi.Context,
	name string, args *IntegrationRuntimeArgs, opts ...pulumi.ResourceOption) (*IntegrationRuntime, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Properties == nil {
		return nil, errors.New("invalid value for required argument 'Properties'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.WorkspaceName == nil {
		return nil, errors.New("invalid value for required argument 'WorkspaceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:synapse:IntegrationRuntime"),
		},
		{
			Type: pulumi.String("azure-native:synapse/v20190601preview:IntegrationRuntime"),
		},
		{
			Type: pulumi.String("azure-native:synapse/v20201201:IntegrationRuntime"),
		},
		{
			Type: pulumi.String("azure-native:synapse/v20210301:IntegrationRuntime"),
		},
		{
			Type: pulumi.String("azure-native:synapse/v20210401preview:IntegrationRuntime"),
		},
		{
			Type: pulumi.String("azure-native:synapse/v20210501:IntegrationRuntime"),
		},
		{
			Type: pulumi.String("azure-native:synapse/v20210601preview:IntegrationRuntime"),
		},
	})
	opts = append(opts, aliases)
	var resource IntegrationRuntime
	err := ctx.RegisterResource("azure-native:synapse/v20210601:IntegrationRuntime", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetIntegrationRuntime(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IntegrationRuntimeState, opts ...pulumi.ResourceOption) (*IntegrationRuntime, error) {
	var resource IntegrationRuntime
	err := ctx.ReadResource("azure-native:synapse/v20210601:IntegrationRuntime", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type integrationRuntimeState struct {
}

type IntegrationRuntimeState struct {
}

func (IntegrationRuntimeState) ElementType() reflect.Type {
	return reflect.TypeOf((*integrationRuntimeState)(nil)).Elem()
}

type integrationRuntimeArgs struct {
	IntegrationRuntimeName *string     `pulumi:"integrationRuntimeName"`
	Properties             interface{} `pulumi:"properties"`
	ResourceGroupName      string      `pulumi:"resourceGroupName"`
	WorkspaceName          string      `pulumi:"workspaceName"`
}


type IntegrationRuntimeArgs struct {
	IntegrationRuntimeName pulumi.StringPtrInput
	Properties             pulumi.Input
	ResourceGroupName      pulumi.StringInput
	WorkspaceName          pulumi.StringInput
}

func (IntegrationRuntimeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*integrationRuntimeArgs)(nil)).Elem()
}

type IntegrationRuntimeInput interface {
	pulumi.Input

	ToIntegrationRuntimeOutput() IntegrationRuntimeOutput
	ToIntegrationRuntimeOutputWithContext(ctx context.Context) IntegrationRuntimeOutput
}

func (*IntegrationRuntime) ElementType() reflect.Type {
	return reflect.TypeOf((*IntegrationRuntime)(nil))
}

func (i *IntegrationRuntime) ToIntegrationRuntimeOutput() IntegrationRuntimeOutput {
	return i.ToIntegrationRuntimeOutputWithContext(context.Background())
}

func (i *IntegrationRuntime) ToIntegrationRuntimeOutputWithContext(ctx context.Context) IntegrationRuntimeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IntegrationRuntimeOutput)
}

type IntegrationRuntimeOutput struct{ *pulumi.OutputState }

func (IntegrationRuntimeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IntegrationRuntime)(nil))
}

func (o IntegrationRuntimeOutput) ToIntegrationRuntimeOutput() IntegrationRuntimeOutput {
	return o
}

func (o IntegrationRuntimeOutput) ToIntegrationRuntimeOutputWithContext(ctx context.Context) IntegrationRuntimeOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(IntegrationRuntimeOutput{})
}
