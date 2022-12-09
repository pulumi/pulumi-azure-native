


package v20210101

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type WebAppSwiftVirtualNetworkConnection struct {
	pulumi.CustomResourceState

	Kind             pulumi.StringPtrOutput `pulumi:"kind"`
	Name             pulumi.StringOutput    `pulumi:"name"`
	SubnetResourceId pulumi.StringPtrOutput `pulumi:"subnetResourceId"`
	SwiftSupported   pulumi.BoolPtrOutput   `pulumi:"swiftSupported"`
	Type             pulumi.StringOutput    `pulumi:"type"`
}


func NewWebAppSwiftVirtualNetworkConnection(ctx *pulumi.Context,
	name string, args *WebAppSwiftVirtualNetworkConnectionArgs, opts ...pulumi.ResourceOption) (*WebAppSwiftVirtualNetworkConnection, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Name == nil {
		return nil, errors.New("invalid value for required argument 'Name'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:web:WebAppSwiftVirtualNetworkConnection"),
		},
		{
			Type: pulumi.String("azure-native:web/v20180201:WebAppSwiftVirtualNetworkConnection"),
		},
		{
			Type: pulumi.String("azure-native:web/v20181101:WebAppSwiftVirtualNetworkConnection"),
		},
		{
			Type: pulumi.String("azure-native:web/v20190801:WebAppSwiftVirtualNetworkConnection"),
		},
		{
			Type: pulumi.String("azure-native:web/v20200601:WebAppSwiftVirtualNetworkConnection"),
		},
		{
			Type: pulumi.String("azure-native:web/v20200901:WebAppSwiftVirtualNetworkConnection"),
		},
		{
			Type: pulumi.String("azure-native:web/v20201001:WebAppSwiftVirtualNetworkConnection"),
		},
		{
			Type: pulumi.String("azure-native:web/v20201201:WebAppSwiftVirtualNetworkConnection"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210115:WebAppSwiftVirtualNetworkConnection"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210201:WebAppSwiftVirtualNetworkConnection"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210301:WebAppSwiftVirtualNetworkConnection"),
		},
		{
			Type: pulumi.String("azure-native:web/v20220301:WebAppSwiftVirtualNetworkConnection"),
		},
	})
	opts = append(opts, aliases)
	var resource WebAppSwiftVirtualNetworkConnection
	err := ctx.RegisterResource("azure-native:web/v20210101:WebAppSwiftVirtualNetworkConnection", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetWebAppSwiftVirtualNetworkConnection(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WebAppSwiftVirtualNetworkConnectionState, opts ...pulumi.ResourceOption) (*WebAppSwiftVirtualNetworkConnection, error) {
	var resource WebAppSwiftVirtualNetworkConnection
	err := ctx.ReadResource("azure-native:web/v20210101:WebAppSwiftVirtualNetworkConnection", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type webAppSwiftVirtualNetworkConnectionState struct {
}

type WebAppSwiftVirtualNetworkConnectionState struct {
}

func (WebAppSwiftVirtualNetworkConnectionState) ElementType() reflect.Type {
	return reflect.TypeOf((*webAppSwiftVirtualNetworkConnectionState)(nil)).Elem()
}

type webAppSwiftVirtualNetworkConnectionArgs struct {
	Kind              *string `pulumi:"kind"`
	Name              string  `pulumi:"name"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
	SubnetResourceId  *string `pulumi:"subnetResourceId"`
	SwiftSupported    *bool   `pulumi:"swiftSupported"`
}


type WebAppSwiftVirtualNetworkConnectionArgs struct {
	Kind              pulumi.StringPtrInput
	Name              pulumi.StringInput
	ResourceGroupName pulumi.StringInput
	SubnetResourceId  pulumi.StringPtrInput
	SwiftSupported    pulumi.BoolPtrInput
}

func (WebAppSwiftVirtualNetworkConnectionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*webAppSwiftVirtualNetworkConnectionArgs)(nil)).Elem()
}

type WebAppSwiftVirtualNetworkConnectionInput interface {
	pulumi.Input

	ToWebAppSwiftVirtualNetworkConnectionOutput() WebAppSwiftVirtualNetworkConnectionOutput
	ToWebAppSwiftVirtualNetworkConnectionOutputWithContext(ctx context.Context) WebAppSwiftVirtualNetworkConnectionOutput
}

func (*WebAppSwiftVirtualNetworkConnection) ElementType() reflect.Type {
	return reflect.TypeOf((**WebAppSwiftVirtualNetworkConnection)(nil)).Elem()
}

func (i *WebAppSwiftVirtualNetworkConnection) ToWebAppSwiftVirtualNetworkConnectionOutput() WebAppSwiftVirtualNetworkConnectionOutput {
	return i.ToWebAppSwiftVirtualNetworkConnectionOutputWithContext(context.Background())
}

func (i *WebAppSwiftVirtualNetworkConnection) ToWebAppSwiftVirtualNetworkConnectionOutputWithContext(ctx context.Context) WebAppSwiftVirtualNetworkConnectionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebAppSwiftVirtualNetworkConnectionOutput)
}

type WebAppSwiftVirtualNetworkConnectionOutput struct{ *pulumi.OutputState }

func (WebAppSwiftVirtualNetworkConnectionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WebAppSwiftVirtualNetworkConnection)(nil)).Elem()
}

func (o WebAppSwiftVirtualNetworkConnectionOutput) ToWebAppSwiftVirtualNetworkConnectionOutput() WebAppSwiftVirtualNetworkConnectionOutput {
	return o
}

func (o WebAppSwiftVirtualNetworkConnectionOutput) ToWebAppSwiftVirtualNetworkConnectionOutputWithContext(ctx context.Context) WebAppSwiftVirtualNetworkConnectionOutput {
	return o
}

func (o WebAppSwiftVirtualNetworkConnectionOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebAppSwiftVirtualNetworkConnection) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o WebAppSwiftVirtualNetworkConnectionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *WebAppSwiftVirtualNetworkConnection) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o WebAppSwiftVirtualNetworkConnectionOutput) SubnetResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebAppSwiftVirtualNetworkConnection) pulumi.StringPtrOutput { return v.SubnetResourceId }).(pulumi.StringPtrOutput)
}

func (o WebAppSwiftVirtualNetworkConnectionOutput) SwiftSupported() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *WebAppSwiftVirtualNetworkConnection) pulumi.BoolPtrOutput { return v.SwiftSupported }).(pulumi.BoolPtrOutput)
}

func (o WebAppSwiftVirtualNetworkConnectionOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *WebAppSwiftVirtualNetworkConnection) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(WebAppSwiftVirtualNetworkConnectionOutput{})
}
