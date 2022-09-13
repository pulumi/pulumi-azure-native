


package v20220301preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type MobileNetwork struct {
	pulumi.CustomResourceState

	CreatedAt                         pulumi.StringPtrOutput   `pulumi:"createdAt"`
	CreatedBy                         pulumi.StringPtrOutput   `pulumi:"createdBy"`
	CreatedByType                     pulumi.StringPtrOutput   `pulumi:"createdByType"`
	LastModifiedAt                    pulumi.StringPtrOutput   `pulumi:"lastModifiedAt"`
	LastModifiedBy                    pulumi.StringPtrOutput   `pulumi:"lastModifiedBy"`
	LastModifiedByType                pulumi.StringPtrOutput   `pulumi:"lastModifiedByType"`
	Location                          pulumi.StringOutput      `pulumi:"location"`
	Name                              pulumi.StringOutput      `pulumi:"name"`
	ProvisioningState                 pulumi.StringOutput      `pulumi:"provisioningState"`
	PublicLandMobileNetworkIdentifier PlmnIdResponseOutput     `pulumi:"publicLandMobileNetworkIdentifier"`
	ServiceKey                        pulumi.StringOutput      `pulumi:"serviceKey"`
	SystemData                        SystemDataResponseOutput `pulumi:"systemData"`
	Tags                              pulumi.StringMapOutput   `pulumi:"tags"`
	Type                              pulumi.StringOutput      `pulumi:"type"`
}


func NewMobileNetwork(ctx *pulumi.Context,
	name string, args *MobileNetworkArgs, opts ...pulumi.ResourceOption) (*MobileNetwork, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.PublicLandMobileNetworkIdentifier == nil {
		return nil, errors.New("invalid value for required argument 'PublicLandMobileNetworkIdentifier'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:mobilenetwork/v20220401preview:MobileNetwork"),
		},
	})
	opts = append(opts, aliases)
	var resource MobileNetwork
	err := ctx.RegisterResource("azure-native:mobilenetwork/v20220301preview:MobileNetwork", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetMobileNetwork(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MobileNetworkState, opts ...pulumi.ResourceOption) (*MobileNetwork, error) {
	var resource MobileNetwork
	err := ctx.ReadResource("azure-native:mobilenetwork/v20220301preview:MobileNetwork", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type mobileNetworkState struct {
}

type MobileNetworkState struct {
}

func (MobileNetworkState) ElementType() reflect.Type {
	return reflect.TypeOf((*mobileNetworkState)(nil)).Elem()
}

type mobileNetworkArgs struct {
	CreatedAt                         *string           `pulumi:"createdAt"`
	CreatedBy                         *string           `pulumi:"createdBy"`
	CreatedByType                     *string           `pulumi:"createdByType"`
	LastModifiedAt                    *string           `pulumi:"lastModifiedAt"`
	LastModifiedBy                    *string           `pulumi:"lastModifiedBy"`
	LastModifiedByType                *string           `pulumi:"lastModifiedByType"`
	Location                          *string           `pulumi:"location"`
	MobileNetworkName                 *string           `pulumi:"mobileNetworkName"`
	PublicLandMobileNetworkIdentifier PlmnId            `pulumi:"publicLandMobileNetworkIdentifier"`
	ResourceGroupName                 string            `pulumi:"resourceGroupName"`
	Tags                              map[string]string `pulumi:"tags"`
}


type MobileNetworkArgs struct {
	CreatedAt                         pulumi.StringPtrInput
	CreatedBy                         pulumi.StringPtrInput
	CreatedByType                     pulumi.StringPtrInput
	LastModifiedAt                    pulumi.StringPtrInput
	LastModifiedBy                    pulumi.StringPtrInput
	LastModifiedByType                pulumi.StringPtrInput
	Location                          pulumi.StringPtrInput
	MobileNetworkName                 pulumi.StringPtrInput
	PublicLandMobileNetworkIdentifier PlmnIdInput
	ResourceGroupName                 pulumi.StringInput
	Tags                              pulumi.StringMapInput
}

func (MobileNetworkArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*mobileNetworkArgs)(nil)).Elem()
}

type MobileNetworkInput interface {
	pulumi.Input

	ToMobileNetworkOutput() MobileNetworkOutput
	ToMobileNetworkOutputWithContext(ctx context.Context) MobileNetworkOutput
}

func (*MobileNetwork) ElementType() reflect.Type {
	return reflect.TypeOf((**MobileNetwork)(nil)).Elem()
}

func (i *MobileNetwork) ToMobileNetworkOutput() MobileNetworkOutput {
	return i.ToMobileNetworkOutputWithContext(context.Background())
}

func (i *MobileNetwork) ToMobileNetworkOutputWithContext(ctx context.Context) MobileNetworkOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MobileNetworkOutput)
}

type MobileNetworkOutput struct{ *pulumi.OutputState }

func (MobileNetworkOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MobileNetwork)(nil)).Elem()
}

func (o MobileNetworkOutput) ToMobileNetworkOutput() MobileNetworkOutput {
	return o
}

func (o MobileNetworkOutput) ToMobileNetworkOutputWithContext(ctx context.Context) MobileNetworkOutput {
	return o
}

func (o MobileNetworkOutput) CreatedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MobileNetwork) pulumi.StringPtrOutput { return v.CreatedAt }).(pulumi.StringPtrOutput)
}

func (o MobileNetworkOutput) CreatedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MobileNetwork) pulumi.StringPtrOutput { return v.CreatedBy }).(pulumi.StringPtrOutput)
}

func (o MobileNetworkOutput) CreatedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MobileNetwork) pulumi.StringPtrOutput { return v.CreatedByType }).(pulumi.StringPtrOutput)
}

func (o MobileNetworkOutput) LastModifiedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MobileNetwork) pulumi.StringPtrOutput { return v.LastModifiedAt }).(pulumi.StringPtrOutput)
}

func (o MobileNetworkOutput) LastModifiedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MobileNetwork) pulumi.StringPtrOutput { return v.LastModifiedBy }).(pulumi.StringPtrOutput)
}

func (o MobileNetworkOutput) LastModifiedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MobileNetwork) pulumi.StringPtrOutput { return v.LastModifiedByType }).(pulumi.StringPtrOutput)
}

func (o MobileNetworkOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *MobileNetwork) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o MobileNetworkOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *MobileNetwork) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o MobileNetworkOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *MobileNetwork) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o MobileNetworkOutput) PublicLandMobileNetworkIdentifier() PlmnIdResponseOutput {
	return o.ApplyT(func(v *MobileNetwork) PlmnIdResponseOutput { return v.PublicLandMobileNetworkIdentifier }).(PlmnIdResponseOutput)
}

func (o MobileNetworkOutput) ServiceKey() pulumi.StringOutput {
	return o.ApplyT(func(v *MobileNetwork) pulumi.StringOutput { return v.ServiceKey }).(pulumi.StringOutput)
}

func (o MobileNetworkOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *MobileNetwork) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o MobileNetworkOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *MobileNetwork) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o MobileNetworkOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *MobileNetwork) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(MobileNetworkOutput{})
}
