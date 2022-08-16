


package v20171001

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


type CapacityDetails struct {
	pulumi.CustomResourceState

	Administration    DedicatedCapacityAdministratorsResponsePtrOutput `pulumi:"administration"`
	FriendlyName      pulumi.StringOutput                              `pulumi:"friendlyName"`
	Location          pulumi.StringOutput                              `pulumi:"location"`
	Mode              pulumi.StringOutput                              `pulumi:"mode"`
	Name              pulumi.StringOutput                              `pulumi:"name"`
	ProvisioningState pulumi.StringOutput                              `pulumi:"provisioningState"`
	Sku               ResourceSkuResponseOutput                        `pulumi:"sku"`
	State             pulumi.StringOutput                              `pulumi:"state"`
	Tags              pulumi.StringMapOutput                           `pulumi:"tags"`
	TenantId          pulumi.StringOutput                              `pulumi:"tenantId"`
	Type              pulumi.StringOutput                              `pulumi:"type"`
}


func NewCapacityDetails(ctx *pulumi.Context,
	name string, args *CapacityDetailsArgs, opts ...pulumi.ResourceOption) (*CapacityDetails, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.Sku == nil {
		return nil, errors.New("invalid value for required argument 'Sku'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:powerbidedicated:CapacityDetails"),
		},
		{
			Type: pulumi.String("azure-native:powerbidedicated/v20210101:CapacityDetails"),
		},
	})
	opts = append(opts, aliases)
	var resource CapacityDetails
	err := ctx.RegisterResource("azure-native:powerbidedicated/v20171001:CapacityDetails", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetCapacityDetails(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CapacityDetailsState, opts ...pulumi.ResourceOption) (*CapacityDetails, error) {
	var resource CapacityDetails
	err := ctx.ReadResource("azure-native:powerbidedicated/v20171001:CapacityDetails", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type capacityDetailsState struct {
}

type CapacityDetailsState struct {
}

func (CapacityDetailsState) ElementType() reflect.Type {
	return reflect.TypeOf((*capacityDetailsState)(nil)).Elem()
}

type capacityDetailsArgs struct {
	Administration        *DedicatedCapacityAdministrators `pulumi:"administration"`
	DedicatedCapacityName *string                          `pulumi:"dedicatedCapacityName"`
	Location              *string                          `pulumi:"location"`
	ResourceGroupName     string                           `pulumi:"resourceGroupName"`
	Sku                   ResourceSku                      `pulumi:"sku"`
	Tags                  map[string]string                `pulumi:"tags"`
}


type CapacityDetailsArgs struct {
	Administration        DedicatedCapacityAdministratorsPtrInput
	DedicatedCapacityName pulumi.StringPtrInput
	Location              pulumi.StringPtrInput
	ResourceGroupName     pulumi.StringInput
	Sku                   ResourceSkuInput
	Tags                  pulumi.StringMapInput
}

func (CapacityDetailsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*capacityDetailsArgs)(nil)).Elem()
}

type CapacityDetailsInput interface {
	pulumi.Input

	ToCapacityDetailsOutput() CapacityDetailsOutput
	ToCapacityDetailsOutputWithContext(ctx context.Context) CapacityDetailsOutput
}

func (*CapacityDetails) ElementType() reflect.Type {
	return reflect.TypeOf((**CapacityDetails)(nil)).Elem()
}

func (i *CapacityDetails) ToCapacityDetailsOutput() CapacityDetailsOutput {
	return i.ToCapacityDetailsOutputWithContext(context.Background())
}

func (i *CapacityDetails) ToCapacityDetailsOutputWithContext(ctx context.Context) CapacityDetailsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CapacityDetailsOutput)
}

type CapacityDetailsOutput struct{ *pulumi.OutputState }

func (CapacityDetailsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CapacityDetails)(nil)).Elem()
}

func (o CapacityDetailsOutput) ToCapacityDetailsOutput() CapacityDetailsOutput {
	return o
}

func (o CapacityDetailsOutput) ToCapacityDetailsOutputWithContext(ctx context.Context) CapacityDetailsOutput {
	return o
}

func (o CapacityDetailsOutput) Administration() DedicatedCapacityAdministratorsResponsePtrOutput {
	return o.ApplyT(func(v *CapacityDetails) DedicatedCapacityAdministratorsResponsePtrOutput { return v.Administration }).(DedicatedCapacityAdministratorsResponsePtrOutput)
}

func (o CapacityDetailsOutput) FriendlyName() pulumi.StringOutput {
	return o.ApplyT(func(v *CapacityDetails) pulumi.StringOutput { return v.FriendlyName }).(pulumi.StringOutput)
}

func (o CapacityDetailsOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *CapacityDetails) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o CapacityDetailsOutput) Mode() pulumi.StringOutput {
	return o.ApplyT(func(v *CapacityDetails) pulumi.StringOutput { return v.Mode }).(pulumi.StringOutput)
}

func (o CapacityDetailsOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *CapacityDetails) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o CapacityDetailsOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *CapacityDetails) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o CapacityDetailsOutput) Sku() ResourceSkuResponseOutput {
	return o.ApplyT(func(v *CapacityDetails) ResourceSkuResponseOutput { return v.Sku }).(ResourceSkuResponseOutput)
}

func (o CapacityDetailsOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v *CapacityDetails) pulumi.StringOutput { return v.State }).(pulumi.StringOutput)
}

func (o CapacityDetailsOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *CapacityDetails) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o CapacityDetailsOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v *CapacityDetails) pulumi.StringOutput { return v.TenantId }).(pulumi.StringOutput)
}

func (o CapacityDetailsOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *CapacityDetails) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(CapacityDetailsOutput{})
}
