


package v20190201

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type InterfaceEndpoint struct {
	pulumi.CustomResourceState

	EndpointService   EndpointServiceResponsePtrOutput    `pulumi:"endpointService"`
	Etag              pulumi.StringPtrOutput              `pulumi:"etag"`
	Fqdn              pulumi.StringPtrOutput              `pulumi:"fqdn"`
	Location          pulumi.StringPtrOutput              `pulumi:"location"`
	Name              pulumi.StringOutput                 `pulumi:"name"`
	NetworkInterfaces NetworkInterfaceResponseArrayOutput `pulumi:"networkInterfaces"`
	Owner             pulumi.StringOutput                 `pulumi:"owner"`
	ProvisioningState pulumi.StringOutput                 `pulumi:"provisioningState"`
	Subnet            SubnetResponsePtrOutput             `pulumi:"subnet"`
	Tags              pulumi.StringMapOutput              `pulumi:"tags"`
	Type              pulumi.StringOutput                 `pulumi:"type"`
}


func NewInterfaceEndpoint(ctx *pulumi.Context,
	name string, args *InterfaceEndpointArgs, opts ...pulumi.ResourceOption) (*InterfaceEndpoint, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:network:InterfaceEndpoint"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180801:InterfaceEndpoint"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181001:InterfaceEndpoint"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181101:InterfaceEndpoint"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181201:InterfaceEndpoint"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190401:InterfaceEndpoint"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190601:InterfaceEndpoint"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190701:InterfaceEndpoint"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190801:InterfaceEndpoint"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190901:InterfaceEndpoint"),
		},
		{
			Type: pulumi.String("azure-native:network/v20191101:InterfaceEndpoint"),
		},
		{
			Type: pulumi.String("azure-native:network/v20191201:InterfaceEndpoint"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200301:InterfaceEndpoint"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200401:InterfaceEndpoint"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200501:InterfaceEndpoint"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200601:InterfaceEndpoint"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200701:InterfaceEndpoint"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200801:InterfaceEndpoint"),
		},
		{
			Type: pulumi.String("azure-native:network/v20201101:InterfaceEndpoint"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210201:InterfaceEndpoint"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210301:InterfaceEndpoint"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210501:InterfaceEndpoint"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210801:InterfaceEndpoint"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220101:InterfaceEndpoint"),
		},
	})
	opts = append(opts, aliases)
	var resource InterfaceEndpoint
	err := ctx.RegisterResource("azure-native:network/v20190201:InterfaceEndpoint", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetInterfaceEndpoint(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *InterfaceEndpointState, opts ...pulumi.ResourceOption) (*InterfaceEndpoint, error) {
	var resource InterfaceEndpoint
	err := ctx.ReadResource("azure-native:network/v20190201:InterfaceEndpoint", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type interfaceEndpointState struct {
}

type InterfaceEndpointState struct {
}

func (InterfaceEndpointState) ElementType() reflect.Type {
	return reflect.TypeOf((*interfaceEndpointState)(nil)).Elem()
}

type interfaceEndpointArgs struct {
	EndpointService       *EndpointService  `pulumi:"endpointService"`
	Fqdn                  *string           `pulumi:"fqdn"`
	Id                    *string           `pulumi:"id"`
	InterfaceEndpointName *string           `pulumi:"interfaceEndpointName"`
	Location              *string           `pulumi:"location"`
	ResourceGroupName     string            `pulumi:"resourceGroupName"`
	Subnet                *SubnetType       `pulumi:"subnet"`
	Tags                  map[string]string `pulumi:"tags"`
}


type InterfaceEndpointArgs struct {
	EndpointService       EndpointServicePtrInput
	Fqdn                  pulumi.StringPtrInput
	Id                    pulumi.StringPtrInput
	InterfaceEndpointName pulumi.StringPtrInput
	Location              pulumi.StringPtrInput
	ResourceGroupName     pulumi.StringInput
	Subnet                SubnetTypePtrInput
	Tags                  pulumi.StringMapInput
}

func (InterfaceEndpointArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*interfaceEndpointArgs)(nil)).Elem()
}

type InterfaceEndpointInput interface {
	pulumi.Input

	ToInterfaceEndpointOutput() InterfaceEndpointOutput
	ToInterfaceEndpointOutputWithContext(ctx context.Context) InterfaceEndpointOutput
}

func (*InterfaceEndpoint) ElementType() reflect.Type {
	return reflect.TypeOf((**InterfaceEndpoint)(nil)).Elem()
}

func (i *InterfaceEndpoint) ToInterfaceEndpointOutput() InterfaceEndpointOutput {
	return i.ToInterfaceEndpointOutputWithContext(context.Background())
}

func (i *InterfaceEndpoint) ToInterfaceEndpointOutputWithContext(ctx context.Context) InterfaceEndpointOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InterfaceEndpointOutput)
}

type InterfaceEndpointOutput struct{ *pulumi.OutputState }

func (InterfaceEndpointOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**InterfaceEndpoint)(nil)).Elem()
}

func (o InterfaceEndpointOutput) ToInterfaceEndpointOutput() InterfaceEndpointOutput {
	return o
}

func (o InterfaceEndpointOutput) ToInterfaceEndpointOutputWithContext(ctx context.Context) InterfaceEndpointOutput {
	return o
}

func (o InterfaceEndpointOutput) EndpointService() EndpointServiceResponsePtrOutput {
	return o.ApplyT(func(v *InterfaceEndpoint) EndpointServiceResponsePtrOutput { return v.EndpointService }).(EndpointServiceResponsePtrOutput)
}

func (o InterfaceEndpointOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *InterfaceEndpoint) pulumi.StringPtrOutput { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o InterfaceEndpointOutput) Fqdn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *InterfaceEndpoint) pulumi.StringPtrOutput { return v.Fqdn }).(pulumi.StringPtrOutput)
}

func (o InterfaceEndpointOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *InterfaceEndpoint) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

func (o InterfaceEndpointOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *InterfaceEndpoint) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o InterfaceEndpointOutput) NetworkInterfaces() NetworkInterfaceResponseArrayOutput {
	return o.ApplyT(func(v *InterfaceEndpoint) NetworkInterfaceResponseArrayOutput { return v.NetworkInterfaces }).(NetworkInterfaceResponseArrayOutput)
}

func (o InterfaceEndpointOutput) Owner() pulumi.StringOutput {
	return o.ApplyT(func(v *InterfaceEndpoint) pulumi.StringOutput { return v.Owner }).(pulumi.StringOutput)
}

func (o InterfaceEndpointOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *InterfaceEndpoint) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o InterfaceEndpointOutput) Subnet() SubnetResponsePtrOutput {
	return o.ApplyT(func(v *InterfaceEndpoint) SubnetResponsePtrOutput { return v.Subnet }).(SubnetResponsePtrOutput)
}

func (o InterfaceEndpointOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *InterfaceEndpoint) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o InterfaceEndpointOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *InterfaceEndpoint) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(InterfaceEndpointOutput{})
}
