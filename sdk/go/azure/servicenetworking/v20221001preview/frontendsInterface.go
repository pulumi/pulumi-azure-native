


package v20221001preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type FrontendsInterface struct {
	pulumi.CustomResourceState

	IpAddressVersion  pulumi.StringPtrOutput                       `pulumi:"ipAddressVersion"`
	Location          pulumi.StringOutput                          `pulumi:"location"`
	Mode              pulumi.StringPtrOutput                       `pulumi:"mode"`
	Name              pulumi.StringOutput                          `pulumi:"name"`
	ProvisioningState pulumi.StringOutput                          `pulumi:"provisioningState"`
	PublicIPAddress   FrontendPropertiesIPAddressResponsePtrOutput `pulumi:"publicIPAddress"`
	SystemData        SystemDataResponseOutput                     `pulumi:"systemData"`
	Tags              pulumi.StringMapOutput                       `pulumi:"tags"`
	Type              pulumi.StringOutput                          `pulumi:"type"`
}


func NewFrontendsInterface(ctx *pulumi.Context,
	name string, args *FrontendsInterfaceArgs, opts ...pulumi.ResourceOption) (*FrontendsInterface, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.TrafficControllerName == nil {
		return nil, errors.New("invalid value for required argument 'TrafficControllerName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:servicenetworking:FrontendsInterface"),
		},
	})
	opts = append(opts, aliases)
	var resource FrontendsInterface
	err := ctx.RegisterResource("azure-native:servicenetworking/v20221001preview:FrontendsInterface", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetFrontendsInterface(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FrontendsInterfaceState, opts ...pulumi.ResourceOption) (*FrontendsInterface, error) {
	var resource FrontendsInterface
	err := ctx.ReadResource("azure-native:servicenetworking/v20221001preview:FrontendsInterface", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type frontendsInterfaceState struct {
}

type FrontendsInterfaceState struct {
}

func (FrontendsInterfaceState) ElementType() reflect.Type {
	return reflect.TypeOf((*frontendsInterfaceState)(nil)).Elem()
}

type frontendsInterfaceArgs struct {
	FrontendName          *string                      `pulumi:"frontendName"`
	IpAddressVersion      *FrontendIPAddressVersion    `pulumi:"ipAddressVersion"`
	Location              *string                      `pulumi:"location"`
	Mode                  *FrontendMode                `pulumi:"mode"`
	PublicIPAddress       *FrontendPropertiesIPAddress `pulumi:"publicIPAddress"`
	ResourceGroupName     string                       `pulumi:"resourceGroupName"`
	Tags                  map[string]string            `pulumi:"tags"`
	TrafficControllerName string                       `pulumi:"trafficControllerName"`
}


type FrontendsInterfaceArgs struct {
	FrontendName          pulumi.StringPtrInput
	IpAddressVersion      FrontendIPAddressVersionPtrInput
	Location              pulumi.StringPtrInput
	Mode                  FrontendModePtrInput
	PublicIPAddress       FrontendPropertiesIPAddressPtrInput
	ResourceGroupName     pulumi.StringInput
	Tags                  pulumi.StringMapInput
	TrafficControllerName pulumi.StringInput
}

func (FrontendsInterfaceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*frontendsInterfaceArgs)(nil)).Elem()
}

type FrontendsInterfaceInput interface {
	pulumi.Input

	ToFrontendsInterfaceOutput() FrontendsInterfaceOutput
	ToFrontendsInterfaceOutputWithContext(ctx context.Context) FrontendsInterfaceOutput
}

func (*FrontendsInterface) ElementType() reflect.Type {
	return reflect.TypeOf((**FrontendsInterface)(nil)).Elem()
}

func (i *FrontendsInterface) ToFrontendsInterfaceOutput() FrontendsInterfaceOutput {
	return i.ToFrontendsInterfaceOutputWithContext(context.Background())
}

func (i *FrontendsInterface) ToFrontendsInterfaceOutputWithContext(ctx context.Context) FrontendsInterfaceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FrontendsInterfaceOutput)
}

type FrontendsInterfaceOutput struct{ *pulumi.OutputState }

func (FrontendsInterfaceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FrontendsInterface)(nil)).Elem()
}

func (o FrontendsInterfaceOutput) ToFrontendsInterfaceOutput() FrontendsInterfaceOutput {
	return o
}

func (o FrontendsInterfaceOutput) ToFrontendsInterfaceOutputWithContext(ctx context.Context) FrontendsInterfaceOutput {
	return o
}

func (o FrontendsInterfaceOutput) IpAddressVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FrontendsInterface) pulumi.StringPtrOutput { return v.IpAddressVersion }).(pulumi.StringPtrOutput)
}

func (o FrontendsInterfaceOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *FrontendsInterface) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o FrontendsInterfaceOutput) Mode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FrontendsInterface) pulumi.StringPtrOutput { return v.Mode }).(pulumi.StringPtrOutput)
}

func (o FrontendsInterfaceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *FrontendsInterface) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o FrontendsInterfaceOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *FrontendsInterface) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o FrontendsInterfaceOutput) PublicIPAddress() FrontendPropertiesIPAddressResponsePtrOutput {
	return o.ApplyT(func(v *FrontendsInterface) FrontendPropertiesIPAddressResponsePtrOutput { return v.PublicIPAddress }).(FrontendPropertiesIPAddressResponsePtrOutput)
}

func (o FrontendsInterfaceOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *FrontendsInterface) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o FrontendsInterfaceOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *FrontendsInterface) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o FrontendsInterfaceOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *FrontendsInterface) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(FrontendsInterfaceOutput{})
}
