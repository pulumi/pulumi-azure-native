


package v20170601

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Manager struct {
	pulumi.CustomResourceState

	CisIntrinsicSettings ManagerIntrinsicSettingsResponsePtrOutput `pulumi:"cisIntrinsicSettings"`
	Etag                 pulumi.StringPtrOutput                    `pulumi:"etag"`
	Location             pulumi.StringOutput                       `pulumi:"location"`
	Name                 pulumi.StringOutput                       `pulumi:"name"`
	ProvisioningState    pulumi.StringPtrOutput                    `pulumi:"provisioningState"`
	Sku                  ManagerSkuResponsePtrOutput               `pulumi:"sku"`
	Tags                 pulumi.StringMapOutput                    `pulumi:"tags"`
	Type                 pulumi.StringOutput                       `pulumi:"type"`
}


func NewManager(ctx *pulumi.Context,
	name string, args *ManagerArgs, opts ...pulumi.ResourceOption) (*Manager, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:storsimple:Manager"),
		},
		{
			Type: pulumi.String("azure-native:storsimple/v20161001:Manager"),
		},
	})
	opts = append(opts, aliases)
	var resource Manager
	err := ctx.RegisterResource("azure-native:storsimple/v20170601:Manager", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetManager(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ManagerState, opts ...pulumi.ResourceOption) (*Manager, error) {
	var resource Manager
	err := ctx.ReadResource("azure-native:storsimple/v20170601:Manager", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type managerState struct {
}

type ManagerState struct {
}

func (ManagerState) ElementType() reflect.Type {
	return reflect.TypeOf((*managerState)(nil)).Elem()
}

type managerArgs struct {
	CisIntrinsicSettings *ManagerIntrinsicSettings `pulumi:"cisIntrinsicSettings"`
	Location             *string                   `pulumi:"location"`
	ManagerName          *string                   `pulumi:"managerName"`
	ProvisioningState    *string                   `pulumi:"provisioningState"`
	ResourceGroupName    string                    `pulumi:"resourceGroupName"`
	Sku                  *ManagerSku               `pulumi:"sku"`
	Tags                 map[string]string         `pulumi:"tags"`
}


type ManagerArgs struct {
	CisIntrinsicSettings ManagerIntrinsicSettingsPtrInput
	Location             pulumi.StringPtrInput
	ManagerName          pulumi.StringPtrInput
	ProvisioningState    pulumi.StringPtrInput
	ResourceGroupName    pulumi.StringInput
	Sku                  ManagerSkuPtrInput
	Tags                 pulumi.StringMapInput
}

func (ManagerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*managerArgs)(nil)).Elem()
}

type ManagerInput interface {
	pulumi.Input

	ToManagerOutput() ManagerOutput
	ToManagerOutputWithContext(ctx context.Context) ManagerOutput
}

func (*Manager) ElementType() reflect.Type {
	return reflect.TypeOf((**Manager)(nil)).Elem()
}

func (i *Manager) ToManagerOutput() ManagerOutput {
	return i.ToManagerOutputWithContext(context.Background())
}

func (i *Manager) ToManagerOutputWithContext(ctx context.Context) ManagerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagerOutput)
}

type ManagerOutput struct{ *pulumi.OutputState }

func (ManagerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Manager)(nil)).Elem()
}

func (o ManagerOutput) ToManagerOutput() ManagerOutput {
	return o
}

func (o ManagerOutput) ToManagerOutputWithContext(ctx context.Context) ManagerOutput {
	return o
}

func (o ManagerOutput) CisIntrinsicSettings() ManagerIntrinsicSettingsResponsePtrOutput {
	return o.ApplyT(func(v *Manager) ManagerIntrinsicSettingsResponsePtrOutput { return v.CisIntrinsicSettings }).(ManagerIntrinsicSettingsResponsePtrOutput)
}

func (o ManagerOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Manager) pulumi.StringPtrOutput { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o ManagerOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *Manager) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o ManagerOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Manager) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ManagerOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Manager) pulumi.StringPtrOutput { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

func (o ManagerOutput) Sku() ManagerSkuResponsePtrOutput {
	return o.ApplyT(func(v *Manager) ManagerSkuResponsePtrOutput { return v.Sku }).(ManagerSkuResponsePtrOutput)
}

func (o ManagerOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Manager) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o ManagerOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Manager) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ManagerOutput{})
}
