


package v20191201

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type BastionHost struct {
	pulumi.CustomResourceState

	DnsName           pulumi.StringPtrOutput                        `pulumi:"dnsName"`
	Etag              pulumi.StringOutput                           `pulumi:"etag"`
	IpConfigurations  BastionHostIPConfigurationResponseArrayOutput `pulumi:"ipConfigurations"`
	Location          pulumi.StringPtrOutput                        `pulumi:"location"`
	Name              pulumi.StringOutput                           `pulumi:"name"`
	ProvisioningState pulumi.StringOutput                           `pulumi:"provisioningState"`
	Tags              pulumi.StringMapOutput                        `pulumi:"tags"`
	Type              pulumi.StringOutput                           `pulumi:"type"`
}


func NewBastionHost(ctx *pulumi.Context,
	name string, args *BastionHostArgs, opts ...pulumi.ResourceOption) (*BastionHost, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:network:BastionHost"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190401:BastionHost"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190601:BastionHost"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190701:BastionHost"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190801:BastionHost"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190901:BastionHost"),
		},
		{
			Type: pulumi.String("azure-native:network/v20191101:BastionHost"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200301:BastionHost"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200401:BastionHost"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200501:BastionHost"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200601:BastionHost"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200701:BastionHost"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200801:BastionHost"),
		},
		{
			Type: pulumi.String("azure-native:network/v20201101:BastionHost"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210201:BastionHost"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210301:BastionHost"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210501:BastionHost"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210801:BastionHost"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220101:BastionHost"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220501:BastionHost"),
		},
	})
	opts = append(opts, aliases)
	var resource BastionHost
	err := ctx.RegisterResource("azure-native:network/v20191201:BastionHost", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetBastionHost(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BastionHostState, opts ...pulumi.ResourceOption) (*BastionHost, error) {
	var resource BastionHost
	err := ctx.ReadResource("azure-native:network/v20191201:BastionHost", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type bastionHostState struct {
}

type BastionHostState struct {
}

func (BastionHostState) ElementType() reflect.Type {
	return reflect.TypeOf((*bastionHostState)(nil)).Elem()
}

type bastionHostArgs struct {
	BastionHostName   *string                      `pulumi:"bastionHostName"`
	DnsName           *string                      `pulumi:"dnsName"`
	Id                *string                      `pulumi:"id"`
	IpConfigurations  []BastionHostIPConfiguration `pulumi:"ipConfigurations"`
	Location          *string                      `pulumi:"location"`
	ResourceGroupName string                       `pulumi:"resourceGroupName"`
	Tags              map[string]string            `pulumi:"tags"`
}


type BastionHostArgs struct {
	BastionHostName   pulumi.StringPtrInput
	DnsName           pulumi.StringPtrInput
	Id                pulumi.StringPtrInput
	IpConfigurations  BastionHostIPConfigurationArrayInput
	Location          pulumi.StringPtrInput
	ResourceGroupName pulumi.StringInput
	Tags              pulumi.StringMapInput
}

func (BastionHostArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*bastionHostArgs)(nil)).Elem()
}

type BastionHostInput interface {
	pulumi.Input

	ToBastionHostOutput() BastionHostOutput
	ToBastionHostOutputWithContext(ctx context.Context) BastionHostOutput
}

func (*BastionHost) ElementType() reflect.Type {
	return reflect.TypeOf((**BastionHost)(nil)).Elem()
}

func (i *BastionHost) ToBastionHostOutput() BastionHostOutput {
	return i.ToBastionHostOutputWithContext(context.Background())
}

func (i *BastionHost) ToBastionHostOutputWithContext(ctx context.Context) BastionHostOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BastionHostOutput)
}

type BastionHostOutput struct{ *pulumi.OutputState }

func (BastionHostOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BastionHost)(nil)).Elem()
}

func (o BastionHostOutput) ToBastionHostOutput() BastionHostOutput {
	return o
}

func (o BastionHostOutput) ToBastionHostOutputWithContext(ctx context.Context) BastionHostOutput {
	return o
}

func (o BastionHostOutput) DnsName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BastionHost) pulumi.StringPtrOutput { return v.DnsName }).(pulumi.StringPtrOutput)
}

func (o BastionHostOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *BastionHost) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

func (o BastionHostOutput) IpConfigurations() BastionHostIPConfigurationResponseArrayOutput {
	return o.ApplyT(func(v *BastionHost) BastionHostIPConfigurationResponseArrayOutput { return v.IpConfigurations }).(BastionHostIPConfigurationResponseArrayOutput)
}

func (o BastionHostOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BastionHost) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

func (o BastionHostOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *BastionHost) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o BastionHostOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *BastionHost) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o BastionHostOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *BastionHost) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o BastionHostOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *BastionHost) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(BastionHostOutput{})
}
