


package v20200701

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type DscpConfiguration struct {
	pulumi.CustomResourceState

	AssociatedNetworkInterfaces NetworkInterfaceResponseArrayOutput `pulumi:"associatedNetworkInterfaces"`
	DestinationIpRanges         QosIpRangeResponseArrayOutput       `pulumi:"destinationIpRanges"`
	DestinationPortRanges       QosPortRangeResponseArrayOutput     `pulumi:"destinationPortRanges"`
	Etag                        pulumi.StringOutput                 `pulumi:"etag"`
	Location                    pulumi.StringPtrOutput              `pulumi:"location"`
	Markings                    pulumi.IntArrayOutput               `pulumi:"markings"`
	Name                        pulumi.StringOutput                 `pulumi:"name"`
	Protocol                    pulumi.StringPtrOutput              `pulumi:"protocol"`
	ProvisioningState           pulumi.StringOutput                 `pulumi:"provisioningState"`
	QosCollectionId             pulumi.StringOutput                 `pulumi:"qosCollectionId"`
	ResourceGuid                pulumi.StringOutput                 `pulumi:"resourceGuid"`
	SourceIpRanges              QosIpRangeResponseArrayOutput       `pulumi:"sourceIpRanges"`
	SourcePortRanges            QosPortRangeResponseArrayOutput     `pulumi:"sourcePortRanges"`
	Tags                        pulumi.StringMapOutput              `pulumi:"tags"`
	Type                        pulumi.StringOutput                 `pulumi:"type"`
}


func NewDscpConfiguration(ctx *pulumi.Context,
	name string, args *DscpConfigurationArgs, opts ...pulumi.ResourceOption) (*DscpConfiguration, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:network:DscpConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200601:DscpConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200801:DscpConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:network/v20201101:DscpConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210201:DscpConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210301:DscpConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210501:DscpConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210801:DscpConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220101:DscpConfiguration"),
		},
	})
	opts = append(opts, aliases)
	var resource DscpConfiguration
	err := ctx.RegisterResource("azure-native:network/v20200701:DscpConfiguration", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetDscpConfiguration(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DscpConfigurationState, opts ...pulumi.ResourceOption) (*DscpConfiguration, error) {
	var resource DscpConfiguration
	err := ctx.ReadResource("azure-native:network/v20200701:DscpConfiguration", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type dscpConfigurationState struct {
}

type DscpConfigurationState struct {
}

func (DscpConfigurationState) ElementType() reflect.Type {
	return reflect.TypeOf((*dscpConfigurationState)(nil)).Elem()
}

type dscpConfigurationArgs struct {
	DestinationIpRanges   []QosIpRange      `pulumi:"destinationIpRanges"`
	DestinationPortRanges []QosPortRange    `pulumi:"destinationPortRanges"`
	DscpConfigurationName *string           `pulumi:"dscpConfigurationName"`
	Id                    *string           `pulumi:"id"`
	Location              *string           `pulumi:"location"`
	Markings              []int             `pulumi:"markings"`
	Protocol              *string           `pulumi:"protocol"`
	ResourceGroupName     string            `pulumi:"resourceGroupName"`
	SourceIpRanges        []QosIpRange      `pulumi:"sourceIpRanges"`
	SourcePortRanges      []QosPortRange    `pulumi:"sourcePortRanges"`
	Tags                  map[string]string `pulumi:"tags"`
}


type DscpConfigurationArgs struct {
	DestinationIpRanges   QosIpRangeArrayInput
	DestinationPortRanges QosPortRangeArrayInput
	DscpConfigurationName pulumi.StringPtrInput
	Id                    pulumi.StringPtrInput
	Location              pulumi.StringPtrInput
	Markings              pulumi.IntArrayInput
	Protocol              pulumi.StringPtrInput
	ResourceGroupName     pulumi.StringInput
	SourceIpRanges        QosIpRangeArrayInput
	SourcePortRanges      QosPortRangeArrayInput
	Tags                  pulumi.StringMapInput
}

func (DscpConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dscpConfigurationArgs)(nil)).Elem()
}

type DscpConfigurationInput interface {
	pulumi.Input

	ToDscpConfigurationOutput() DscpConfigurationOutput
	ToDscpConfigurationOutputWithContext(ctx context.Context) DscpConfigurationOutput
}

func (*DscpConfiguration) ElementType() reflect.Type {
	return reflect.TypeOf((**DscpConfiguration)(nil)).Elem()
}

func (i *DscpConfiguration) ToDscpConfigurationOutput() DscpConfigurationOutput {
	return i.ToDscpConfigurationOutputWithContext(context.Background())
}

func (i *DscpConfiguration) ToDscpConfigurationOutputWithContext(ctx context.Context) DscpConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DscpConfigurationOutput)
}

type DscpConfigurationOutput struct{ *pulumi.OutputState }

func (DscpConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DscpConfiguration)(nil)).Elem()
}

func (o DscpConfigurationOutput) ToDscpConfigurationOutput() DscpConfigurationOutput {
	return o
}

func (o DscpConfigurationOutput) ToDscpConfigurationOutputWithContext(ctx context.Context) DscpConfigurationOutput {
	return o
}

func (o DscpConfigurationOutput) AssociatedNetworkInterfaces() NetworkInterfaceResponseArrayOutput {
	return o.ApplyT(func(v *DscpConfiguration) NetworkInterfaceResponseArrayOutput { return v.AssociatedNetworkInterfaces }).(NetworkInterfaceResponseArrayOutput)
}

func (o DscpConfigurationOutput) DestinationIpRanges() QosIpRangeResponseArrayOutput {
	return o.ApplyT(func(v *DscpConfiguration) QosIpRangeResponseArrayOutput { return v.DestinationIpRanges }).(QosIpRangeResponseArrayOutput)
}

func (o DscpConfigurationOutput) DestinationPortRanges() QosPortRangeResponseArrayOutput {
	return o.ApplyT(func(v *DscpConfiguration) QosPortRangeResponseArrayOutput { return v.DestinationPortRanges }).(QosPortRangeResponseArrayOutput)
}

func (o DscpConfigurationOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *DscpConfiguration) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

func (o DscpConfigurationOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DscpConfiguration) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

func (o DscpConfigurationOutput) Markings() pulumi.IntArrayOutput {
	return o.ApplyT(func(v *DscpConfiguration) pulumi.IntArrayOutput { return v.Markings }).(pulumi.IntArrayOutput)
}

func (o DscpConfigurationOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *DscpConfiguration) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o DscpConfigurationOutput) Protocol() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DscpConfiguration) pulumi.StringPtrOutput { return v.Protocol }).(pulumi.StringPtrOutput)
}

func (o DscpConfigurationOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *DscpConfiguration) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o DscpConfigurationOutput) QosCollectionId() pulumi.StringOutput {
	return o.ApplyT(func(v *DscpConfiguration) pulumi.StringOutput { return v.QosCollectionId }).(pulumi.StringOutput)
}

func (o DscpConfigurationOutput) ResourceGuid() pulumi.StringOutput {
	return o.ApplyT(func(v *DscpConfiguration) pulumi.StringOutput { return v.ResourceGuid }).(pulumi.StringOutput)
}

func (o DscpConfigurationOutput) SourceIpRanges() QosIpRangeResponseArrayOutput {
	return o.ApplyT(func(v *DscpConfiguration) QosIpRangeResponseArrayOutput { return v.SourceIpRanges }).(QosIpRangeResponseArrayOutput)
}

func (o DscpConfigurationOutput) SourcePortRanges() QosPortRangeResponseArrayOutput {
	return o.ApplyT(func(v *DscpConfiguration) QosPortRangeResponseArrayOutput { return v.SourcePortRanges }).(QosPortRangeResponseArrayOutput)
}

func (o DscpConfigurationOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *DscpConfiguration) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o DscpConfigurationOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *DscpConfiguration) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(DscpConfigurationOutput{})
}
