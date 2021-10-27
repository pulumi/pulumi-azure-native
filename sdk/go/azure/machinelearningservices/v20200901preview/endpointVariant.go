


package v20200901preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type EndpointVariant struct {
	pulumi.CustomResourceState

	Identity   IdentityResponsePtrOutput `pulumi:"identity"`
	Location   pulumi.StringPtrOutput    `pulumi:"location"`
	Name       pulumi.StringOutput       `pulumi:"name"`
	Properties pulumi.AnyOutput          `pulumi:"properties"`
	Sku        SkuResponsePtrOutput      `pulumi:"sku"`
	Tags       pulumi.StringMapOutput    `pulumi:"tags"`
	Type       pulumi.StringOutput       `pulumi:"type"`
}


func NewEndpointVariant(ctx *pulumi.Context,
	name string, args *EndpointVariantArgs, opts ...pulumi.ResourceOption) (*EndpointVariant, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ComputeType == nil {
		return nil, errors.New("invalid value for required argument 'ComputeType'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.WorkspaceName == nil {
		return nil, errors.New("invalid value for required argument 'WorkspaceName'")
	}
	args.ComputeType = pulumi.String("Custom")
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:machinelearningservices/v20200901preview:EndpointVariant"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices:EndpointVariant"),
		},
		{
			Type: pulumi.String("azure-nextgen:machinelearningservices:EndpointVariant"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20200501preview:EndpointVariant"),
		},
		{
			Type: pulumi.String("azure-nextgen:machinelearningservices/v20200501preview:EndpointVariant"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20200515preview:EndpointVariant"),
		},
		{
			Type: pulumi.String("azure-nextgen:machinelearningservices/v20200515preview:EndpointVariant"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20210101:EndpointVariant"),
		},
		{
			Type: pulumi.String("azure-nextgen:machinelearningservices/v20210101:EndpointVariant"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20210401:EndpointVariant"),
		},
		{
			Type: pulumi.String("azure-nextgen:machinelearningservices/v20210401:EndpointVariant"),
		},
	})
	opts = append(opts, aliases)
	var resource EndpointVariant
	err := ctx.RegisterResource("azure-native:machinelearningservices/v20200901preview:EndpointVariant", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetEndpointVariant(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EndpointVariantState, opts ...pulumi.ResourceOption) (*EndpointVariant, error) {
	var resource EndpointVariant
	err := ctx.ReadResource("azure-native:machinelearningservices/v20200901preview:EndpointVariant", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type endpointVariantState struct {
}

type EndpointVariantState struct {
}

func (EndpointVariantState) ElementType() reflect.Type {
	return reflect.TypeOf((*endpointVariantState)(nil)).Elem()
}

type endpointVariantArgs struct {
	ComputeType             string                                       `pulumi:"computeType"`
	Description             *string                                      `pulumi:"description"`
	EnvironmentImageRequest *CreateServiceRequestEnvironmentImageRequest `pulumi:"environmentImageRequest"`
	IsDefault               *bool                                        `pulumi:"isDefault"`
	Keys                    *CreateServiceRequestKeys                    `pulumi:"keys"`
	KvTags                  map[string]string                            `pulumi:"kvTags"`
	Location                *string                                      `pulumi:"location"`
	Properties              map[string]string                            `pulumi:"properties"`
	ResourceGroupName       string                                       `pulumi:"resourceGroupName"`
	ServiceName             *string                                      `pulumi:"serviceName"`
	TrafficPercentile       *float64                                     `pulumi:"trafficPercentile"`
	Type                    *string                                      `pulumi:"type"`
	WorkspaceName           string                                       `pulumi:"workspaceName"`
}


type EndpointVariantArgs struct {
	ComputeType             pulumi.StringInput
	Description             pulumi.StringPtrInput
	EnvironmentImageRequest CreateServiceRequestEnvironmentImageRequestPtrInput
	IsDefault               pulumi.BoolPtrInput
	Keys                    CreateServiceRequestKeysPtrInput
	KvTags                  pulumi.StringMapInput
	Location                pulumi.StringPtrInput
	Properties              pulumi.StringMapInput
	ResourceGroupName       pulumi.StringInput
	ServiceName             pulumi.StringPtrInput
	TrafficPercentile       pulumi.Float64PtrInput
	Type                    pulumi.StringPtrInput
	WorkspaceName           pulumi.StringInput
}

func (EndpointVariantArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*endpointVariantArgs)(nil)).Elem()
}

type EndpointVariantInput interface {
	pulumi.Input

	ToEndpointVariantOutput() EndpointVariantOutput
	ToEndpointVariantOutputWithContext(ctx context.Context) EndpointVariantOutput
}

func (*EndpointVariant) ElementType() reflect.Type {
	return reflect.TypeOf((*EndpointVariant)(nil))
}

func (i *EndpointVariant) ToEndpointVariantOutput() EndpointVariantOutput {
	return i.ToEndpointVariantOutputWithContext(context.Background())
}

func (i *EndpointVariant) ToEndpointVariantOutputWithContext(ctx context.Context) EndpointVariantOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EndpointVariantOutput)
}

type EndpointVariantOutput struct{ *pulumi.OutputState }

func (EndpointVariantOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EndpointVariant)(nil))
}

func (o EndpointVariantOutput) ToEndpointVariantOutput() EndpointVariantOutput {
	return o
}

func (o EndpointVariantOutput) ToEndpointVariantOutputWithContext(ctx context.Context) EndpointVariantOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*EndpointVariantInput)(nil)).Elem(), &EndpointVariant{})
	pulumi.RegisterOutputType(EndpointVariantOutput{})
}
