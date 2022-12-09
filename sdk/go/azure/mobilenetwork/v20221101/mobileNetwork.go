


package v20221101

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type MobileNetwork struct {
	pulumi.CustomResourceState

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
			Type: pulumi.String("azure-native:mobilenetwork:MobileNetwork"),
		},
		{
			Type: pulumi.String("azure-native:mobilenetwork/v20220301preview:MobileNetwork"),
		},
		{
			Type: pulumi.String("azure-native:mobilenetwork/v20220401preview:MobileNetwork"),
		},
	})
	opts = append(opts, aliases)
	var resource MobileNetwork
	err := ctx.RegisterResource("azure-native:mobilenetwork/v20221101:MobileNetwork", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetMobileNetwork(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MobileNetworkState, opts ...pulumi.ResourceOption) (*MobileNetwork, error) {
	var resource MobileNetwork
	err := ctx.ReadResource("azure-native:mobilenetwork/v20221101:MobileNetwork", name, id, state, &resource, opts...)
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
	Location                          *string           `pulumi:"location"`
	MobileNetworkName                 *string           `pulumi:"mobileNetworkName"`
	PublicLandMobileNetworkIdentifier PlmnId            `pulumi:"publicLandMobileNetworkIdentifier"`
	ResourceGroupName                 string            `pulumi:"resourceGroupName"`
	Tags                              map[string]string `pulumi:"tags"`
}


type MobileNetworkArgs struct {
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
