


package v20201101

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type CustomIPPrefix struct {
	pulumi.CustomResourceState

	Cidr              pulumi.StringPtrOutput            `pulumi:"cidr"`
	CommissionedState pulumi.StringPtrOutput            `pulumi:"commissionedState"`
	Etag              pulumi.StringOutput               `pulumi:"etag"`
	ExtendedLocation  ExtendedLocationResponsePtrOutput `pulumi:"extendedLocation"`
	Location          pulumi.StringPtrOutput            `pulumi:"location"`
	Name              pulumi.StringOutput               `pulumi:"name"`
	ProvisioningState pulumi.StringOutput               `pulumi:"provisioningState"`
	PublicIpPrefixes  SubResourceResponseArrayOutput    `pulumi:"publicIpPrefixes"`
	ResourceGuid      pulumi.StringOutput               `pulumi:"resourceGuid"`
	Tags              pulumi.StringMapOutput            `pulumi:"tags"`
	Type              pulumi.StringOutput               `pulumi:"type"`
	Zones             pulumi.StringArrayOutput          `pulumi:"zones"`
}


func NewCustomIPPrefix(ctx *pulumi.Context,
	name string, args *CustomIPPrefixArgs, opts ...pulumi.ResourceOption) (*CustomIPPrefix, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:network:CustomIPPrefix"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200601:CustomIPPrefix"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200701:CustomIPPrefix"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200801:CustomIPPrefix"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210201:CustomIPPrefix"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210301:CustomIPPrefix"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210501:CustomIPPrefix"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210801:CustomIPPrefix"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220101:CustomIPPrefix"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220501:CustomIPPrefix"),
		},
	})
	opts = append(opts, aliases)
	var resource CustomIPPrefix
	err := ctx.RegisterResource("azure-native:network/v20201101:CustomIPPrefix", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetCustomIPPrefix(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CustomIPPrefixState, opts ...pulumi.ResourceOption) (*CustomIPPrefix, error) {
	var resource CustomIPPrefix
	err := ctx.ReadResource("azure-native:network/v20201101:CustomIPPrefix", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type customIPPrefixState struct {
}

type CustomIPPrefixState struct {
}

func (CustomIPPrefixState) ElementType() reflect.Type {
	return reflect.TypeOf((*customIPPrefixState)(nil)).Elem()
}

type customIPPrefixArgs struct {
	Cidr               *string           `pulumi:"cidr"`
	CommissionedState  *string           `pulumi:"commissionedState"`
	CustomIpPrefixName *string           `pulumi:"customIpPrefixName"`
	ExtendedLocation   *ExtendedLocation `pulumi:"extendedLocation"`
	Id                 *string           `pulumi:"id"`
	Location           *string           `pulumi:"location"`
	ResourceGroupName  string            `pulumi:"resourceGroupName"`
	Tags               map[string]string `pulumi:"tags"`
	Zones              []string          `pulumi:"zones"`
}


type CustomIPPrefixArgs struct {
	Cidr               pulumi.StringPtrInput
	CommissionedState  pulumi.StringPtrInput
	CustomIpPrefixName pulumi.StringPtrInput
	ExtendedLocation   ExtendedLocationPtrInput
	Id                 pulumi.StringPtrInput
	Location           pulumi.StringPtrInput
	ResourceGroupName  pulumi.StringInput
	Tags               pulumi.StringMapInput
	Zones              pulumi.StringArrayInput
}

func (CustomIPPrefixArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*customIPPrefixArgs)(nil)).Elem()
}

type CustomIPPrefixInput interface {
	pulumi.Input

	ToCustomIPPrefixOutput() CustomIPPrefixOutput
	ToCustomIPPrefixOutputWithContext(ctx context.Context) CustomIPPrefixOutput
}

func (*CustomIPPrefix) ElementType() reflect.Type {
	return reflect.TypeOf((**CustomIPPrefix)(nil)).Elem()
}

func (i *CustomIPPrefix) ToCustomIPPrefixOutput() CustomIPPrefixOutput {
	return i.ToCustomIPPrefixOutputWithContext(context.Background())
}

func (i *CustomIPPrefix) ToCustomIPPrefixOutputWithContext(ctx context.Context) CustomIPPrefixOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomIPPrefixOutput)
}

type CustomIPPrefixOutput struct{ *pulumi.OutputState }

func (CustomIPPrefixOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CustomIPPrefix)(nil)).Elem()
}

func (o CustomIPPrefixOutput) ToCustomIPPrefixOutput() CustomIPPrefixOutput {
	return o
}

func (o CustomIPPrefixOutput) ToCustomIPPrefixOutputWithContext(ctx context.Context) CustomIPPrefixOutput {
	return o
}

func (o CustomIPPrefixOutput) Cidr() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CustomIPPrefix) pulumi.StringPtrOutput { return v.Cidr }).(pulumi.StringPtrOutput)
}

func (o CustomIPPrefixOutput) CommissionedState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CustomIPPrefix) pulumi.StringPtrOutput { return v.CommissionedState }).(pulumi.StringPtrOutput)
}

func (o CustomIPPrefixOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *CustomIPPrefix) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

func (o CustomIPPrefixOutput) ExtendedLocation() ExtendedLocationResponsePtrOutput {
	return o.ApplyT(func(v *CustomIPPrefix) ExtendedLocationResponsePtrOutput { return v.ExtendedLocation }).(ExtendedLocationResponsePtrOutput)
}

func (o CustomIPPrefixOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CustomIPPrefix) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

func (o CustomIPPrefixOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *CustomIPPrefix) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o CustomIPPrefixOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *CustomIPPrefix) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o CustomIPPrefixOutput) PublicIpPrefixes() SubResourceResponseArrayOutput {
	return o.ApplyT(func(v *CustomIPPrefix) SubResourceResponseArrayOutput { return v.PublicIpPrefixes }).(SubResourceResponseArrayOutput)
}

func (o CustomIPPrefixOutput) ResourceGuid() pulumi.StringOutput {
	return o.ApplyT(func(v *CustomIPPrefix) pulumi.StringOutput { return v.ResourceGuid }).(pulumi.StringOutput)
}

func (o CustomIPPrefixOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *CustomIPPrefix) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o CustomIPPrefixOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *CustomIPPrefix) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o CustomIPPrefixOutput) Zones() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *CustomIPPrefix) pulumi.StringArrayOutput { return v.Zones }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(CustomIPPrefixOutput{})
}
