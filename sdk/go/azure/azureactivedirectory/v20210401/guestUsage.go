


package v20210401

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type GuestUsage struct {
	pulumi.CustomResourceState

	Location   pulumi.StringPtrOutput   `pulumi:"location"`
	Name       pulumi.StringOutput      `pulumi:"name"`
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	Tags       pulumi.StringMapOutput   `pulumi:"tags"`
	TenantId   pulumi.StringPtrOutput   `pulumi:"tenantId"`
	Type       pulumi.StringOutput      `pulumi:"type"`
}


func NewGuestUsage(ctx *pulumi.Context,
	name string, args *GuestUsageArgs, opts ...pulumi.ResourceOption) (*GuestUsage, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:azureactivedirectory:GuestUsage"),
		},
		{
			Type: pulumi.String("azure-native:azureactivedirectory/v20200501preview:GuestUsage"),
		},
	})
	opts = append(opts, aliases)
	var resource GuestUsage
	err := ctx.RegisterResource("azure-native:azureactivedirectory/v20210401:GuestUsage", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetGuestUsage(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GuestUsageState, opts ...pulumi.ResourceOption) (*GuestUsage, error) {
	var resource GuestUsage
	err := ctx.ReadResource("azure-native:azureactivedirectory/v20210401:GuestUsage", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type guestUsageState struct {
}

type GuestUsageState struct {
}

func (GuestUsageState) ElementType() reflect.Type {
	return reflect.TypeOf((*guestUsageState)(nil)).Elem()
}

type guestUsageArgs struct {
	Location          *string           `pulumi:"location"`
	ResourceGroupName string            `pulumi:"resourceGroupName"`
	ResourceName      *string           `pulumi:"resourceName"`
	Tags              map[string]string `pulumi:"tags"`
	TenantId          *string           `pulumi:"tenantId"`
}


type GuestUsageArgs struct {
	Location          pulumi.StringPtrInput
	ResourceGroupName pulumi.StringInput
	ResourceName      pulumi.StringPtrInput
	Tags              pulumi.StringMapInput
	TenantId          pulumi.StringPtrInput
}

func (GuestUsageArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*guestUsageArgs)(nil)).Elem()
}

type GuestUsageInput interface {
	pulumi.Input

	ToGuestUsageOutput() GuestUsageOutput
	ToGuestUsageOutputWithContext(ctx context.Context) GuestUsageOutput
}

func (*GuestUsage) ElementType() reflect.Type {
	return reflect.TypeOf((**GuestUsage)(nil)).Elem()
}

func (i *GuestUsage) ToGuestUsageOutput() GuestUsageOutput {
	return i.ToGuestUsageOutputWithContext(context.Background())
}

func (i *GuestUsage) ToGuestUsageOutputWithContext(ctx context.Context) GuestUsageOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GuestUsageOutput)
}

type GuestUsageOutput struct{ *pulumi.OutputState }

func (GuestUsageOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GuestUsage)(nil)).Elem()
}

func (o GuestUsageOutput) ToGuestUsageOutput() GuestUsageOutput {
	return o
}

func (o GuestUsageOutput) ToGuestUsageOutputWithContext(ctx context.Context) GuestUsageOutput {
	return o
}

func (o GuestUsageOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GuestUsage) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

func (o GuestUsageOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *GuestUsage) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o GuestUsageOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *GuestUsage) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o GuestUsageOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *GuestUsage) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o GuestUsageOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GuestUsage) pulumi.StringPtrOutput { return v.TenantId }).(pulumi.StringPtrOutput)
}

func (o GuestUsageOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *GuestUsage) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GuestUsageOutput{})
}
