


package devcenter

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type DevCenter struct {
	pulumi.CustomResourceState

	Identity          ManagedServiceIdentityResponsePtrOutput `pulumi:"identity"`
	Location          pulumi.StringOutput                     `pulumi:"location"`
	Name              pulumi.StringOutput                     `pulumi:"name"`
	ProvisioningState pulumi.StringOutput                     `pulumi:"provisioningState"`
	SystemData        SystemDataResponseOutput                `pulumi:"systemData"`
	Tags              pulumi.StringMapOutput                  `pulumi:"tags"`
	Type              pulumi.StringOutput                     `pulumi:"type"`
}


func NewDevCenter(ctx *pulumi.Context,
	name string, args *DevCenterArgs, opts ...pulumi.ResourceOption) (*DevCenter, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:devcenter/v20220801preview:DevCenter"),
		},
		{
			Type: pulumi.String("azure-native:devcenter/v20220901preview:DevCenter"),
		},
	})
	opts = append(opts, aliases)
	var resource DevCenter
	err := ctx.RegisterResource("azure-native:devcenter:DevCenter", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetDevCenter(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DevCenterState, opts ...pulumi.ResourceOption) (*DevCenter, error) {
	var resource DevCenter
	err := ctx.ReadResource("azure-native:devcenter:DevCenter", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type devCenterState struct {
}

type DevCenterState struct {
}

func (DevCenterState) ElementType() reflect.Type {
	return reflect.TypeOf((*devCenterState)(nil)).Elem()
}

type devCenterArgs struct {
	DevCenterName     *string                 `pulumi:"devCenterName"`
	Identity          *ManagedServiceIdentity `pulumi:"identity"`
	Location          *string                 `pulumi:"location"`
	ResourceGroupName string                  `pulumi:"resourceGroupName"`
	Tags              map[string]string       `pulumi:"tags"`
}


type DevCenterArgs struct {
	DevCenterName     pulumi.StringPtrInput
	Identity          ManagedServiceIdentityPtrInput
	Location          pulumi.StringPtrInput
	ResourceGroupName pulumi.StringInput
	Tags              pulumi.StringMapInput
}

func (DevCenterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*devCenterArgs)(nil)).Elem()
}

type DevCenterInput interface {
	pulumi.Input

	ToDevCenterOutput() DevCenterOutput
	ToDevCenterOutputWithContext(ctx context.Context) DevCenterOutput
}

func (*DevCenter) ElementType() reflect.Type {
	return reflect.TypeOf((**DevCenter)(nil)).Elem()
}

func (i *DevCenter) ToDevCenterOutput() DevCenterOutput {
	return i.ToDevCenterOutputWithContext(context.Background())
}

func (i *DevCenter) ToDevCenterOutputWithContext(ctx context.Context) DevCenterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DevCenterOutput)
}

type DevCenterOutput struct{ *pulumi.OutputState }

func (DevCenterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DevCenter)(nil)).Elem()
}

func (o DevCenterOutput) ToDevCenterOutput() DevCenterOutput {
	return o
}

func (o DevCenterOutput) ToDevCenterOutputWithContext(ctx context.Context) DevCenterOutput {
	return o
}

func (o DevCenterOutput) Identity() ManagedServiceIdentityResponsePtrOutput {
	return o.ApplyT(func(v *DevCenter) ManagedServiceIdentityResponsePtrOutput { return v.Identity }).(ManagedServiceIdentityResponsePtrOutput)
}

func (o DevCenterOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *DevCenter) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o DevCenterOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *DevCenter) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o DevCenterOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *DevCenter) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o DevCenterOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *DevCenter) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o DevCenterOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *DevCenter) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o DevCenterOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *DevCenter) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(DevCenterOutput{})
}
