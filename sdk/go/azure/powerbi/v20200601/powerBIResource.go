


package v20200601

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type PowerBIResource struct {
	pulumi.CustomResourceState

	Location                   pulumi.StringPtrOutput                       `pulumi:"location"`
	Name                       pulumi.StringOutput                          `pulumi:"name"`
	PrivateEndpointConnections PrivateEndpointConnectionResponseArrayOutput `pulumi:"privateEndpointConnections"`
	SystemData                 SystemDataResponseOutput                     `pulumi:"systemData"`
	Tags                       pulumi.StringMapOutput                       `pulumi:"tags"`
	TenantId                   pulumi.StringPtrOutput                       `pulumi:"tenantId"`
	Type                       pulumi.StringOutput                          `pulumi:"type"`
}


func NewPowerBIResource(ctx *pulumi.Context,
	name string, args *PowerBIResourceArgs, opts ...pulumi.ResourceOption) (*PowerBIResource, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:powerbi:PowerBIResource"),
		},
	})
	opts = append(opts, aliases)
	var resource PowerBIResource
	err := ctx.RegisterResource("azure-native:powerbi/v20200601:PowerBIResource", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetPowerBIResource(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PowerBIResourceState, opts ...pulumi.ResourceOption) (*PowerBIResource, error) {
	var resource PowerBIResource
	err := ctx.ReadResource("azure-native:powerbi/v20200601:PowerBIResource", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type powerBIResourceState struct {
}

type PowerBIResourceState struct {
}

func (PowerBIResourceState) ElementType() reflect.Type {
	return reflect.TypeOf((*powerBIResourceState)(nil)).Elem()
}

type powerBIResourceArgs struct {
	AzureResourceName          *string                         `pulumi:"azureResourceName"`
	Location                   *string                         `pulumi:"location"`
	PrivateEndpointConnections []PrivateEndpointConnectionType `pulumi:"privateEndpointConnections"`
	ResourceGroupName          string                          `pulumi:"resourceGroupName"`
	Tags                       map[string]string               `pulumi:"tags"`
	TenantId                   *string                         `pulumi:"tenantId"`
}


type PowerBIResourceArgs struct {
	AzureResourceName          pulumi.StringPtrInput
	Location                   pulumi.StringPtrInput
	PrivateEndpointConnections PrivateEndpointConnectionTypeArrayInput
	ResourceGroupName          pulumi.StringInput
	Tags                       pulumi.StringMapInput
	TenantId                   pulumi.StringPtrInput
}

func (PowerBIResourceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*powerBIResourceArgs)(nil)).Elem()
}

type PowerBIResourceInput interface {
	pulumi.Input

	ToPowerBIResourceOutput() PowerBIResourceOutput
	ToPowerBIResourceOutputWithContext(ctx context.Context) PowerBIResourceOutput
}

func (*PowerBIResource) ElementType() reflect.Type {
	return reflect.TypeOf((**PowerBIResource)(nil)).Elem()
}

func (i *PowerBIResource) ToPowerBIResourceOutput() PowerBIResourceOutput {
	return i.ToPowerBIResourceOutputWithContext(context.Background())
}

func (i *PowerBIResource) ToPowerBIResourceOutputWithContext(ctx context.Context) PowerBIResourceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PowerBIResourceOutput)
}

type PowerBIResourceOutput struct{ *pulumi.OutputState }

func (PowerBIResourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PowerBIResource)(nil)).Elem()
}

func (o PowerBIResourceOutput) ToPowerBIResourceOutput() PowerBIResourceOutput {
	return o
}

func (o PowerBIResourceOutput) ToPowerBIResourceOutputWithContext(ctx context.Context) PowerBIResourceOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(PowerBIResourceOutput{})
}
