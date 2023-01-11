


package v20210114preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ApplicationGroup struct {
	pulumi.CustomResourceState

	ApplicationGroupType pulumi.StringOutput                                          `pulumi:"applicationGroupType"`
	CloudPcResource      pulumi.BoolOutput                                            `pulumi:"cloudPcResource"`
	Description          pulumi.StringPtrOutput                                       `pulumi:"description"`
	Etag                 pulumi.StringOutput                                          `pulumi:"etag"`
	FriendlyName         pulumi.StringPtrOutput                                       `pulumi:"friendlyName"`
	HostPoolArmPath      pulumi.StringOutput                                          `pulumi:"hostPoolArmPath"`
	Identity             ResourceModelWithAllowedPropertySetResponseIdentityPtrOutput `pulumi:"identity"`
	Kind                 pulumi.StringPtrOutput                                       `pulumi:"kind"`
	Location             pulumi.StringPtrOutput                                       `pulumi:"location"`
	ManagedBy            pulumi.StringPtrOutput                                       `pulumi:"managedBy"`
	MigrationRequest     MigrationRequestPropertiesResponsePtrOutput                  `pulumi:"migrationRequest"`
	Name                 pulumi.StringOutput                                          `pulumi:"name"`
	ObjectId             pulumi.StringOutput                                          `pulumi:"objectId"`
	Plan                 ResourceModelWithAllowedPropertySetResponsePlanPtrOutput     `pulumi:"plan"`
	Sku                  ResourceModelWithAllowedPropertySetResponseSkuPtrOutput      `pulumi:"sku"`
	Tags                 pulumi.StringMapOutput                                       `pulumi:"tags"`
	Type                 pulumi.StringOutput                                          `pulumi:"type"`
	WorkspaceArmPath     pulumi.StringOutput                                          `pulumi:"workspaceArmPath"`
}


func NewApplicationGroup(ctx *pulumi.Context,
	name string, args *ApplicationGroupArgs, opts ...pulumi.ResourceOption) (*ApplicationGroup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ApplicationGroupType == nil {
		return nil, errors.New("invalid value for required argument 'ApplicationGroupType'")
	}
	if args.HostPoolArmPath == nil {
		return nil, errors.New("invalid value for required argument 'HostPoolArmPath'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:desktopvirtualization:ApplicationGroup"),
		},
		{
			Type: pulumi.String("azure-native:desktopvirtualization/v20190123preview:ApplicationGroup"),
		},
		{
			Type: pulumi.String("azure-native:desktopvirtualization/v20190924preview:ApplicationGroup"),
		},
		{
			Type: pulumi.String("azure-native:desktopvirtualization/v20191210preview:ApplicationGroup"),
		},
		{
			Type: pulumi.String("azure-native:desktopvirtualization/v20200921preview:ApplicationGroup"),
		},
		{
			Type: pulumi.String("azure-native:desktopvirtualization/v20201019preview:ApplicationGroup"),
		},
		{
			Type: pulumi.String("azure-native:desktopvirtualization/v20201102preview:ApplicationGroup"),
		},
		{
			Type: pulumi.String("azure-native:desktopvirtualization/v20201110preview:ApplicationGroup"),
		},
		{
			Type: pulumi.String("azure-native:desktopvirtualization/v20210201preview:ApplicationGroup"),
		},
		{
			Type: pulumi.String("azure-native:desktopvirtualization/v20210309preview:ApplicationGroup"),
		},
		{
			Type: pulumi.String("azure-native:desktopvirtualization/v20210401preview:ApplicationGroup"),
		},
		{
			Type: pulumi.String("azure-native:desktopvirtualization/v20210712:ApplicationGroup"),
		},
		{
			Type: pulumi.String("azure-native:desktopvirtualization/v20210903preview:ApplicationGroup"),
		},
		{
			Type: pulumi.String("azure-native:desktopvirtualization/v20220210preview:ApplicationGroup"),
		},
		{
			Type: pulumi.String("azure-native:desktopvirtualization/v20220401preview:ApplicationGroup"),
		},
		{
			Type: pulumi.String("azure-native:desktopvirtualization/v20220909:ApplicationGroup"),
		},
		{
			Type: pulumi.String("azure-native:desktopvirtualization/v20221014preview:ApplicationGroup"),
		},
	})
	opts = append(opts, aliases)
	var resource ApplicationGroup
	err := ctx.RegisterResource("azure-native:desktopvirtualization/v20210114preview:ApplicationGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetApplicationGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ApplicationGroupState, opts ...pulumi.ResourceOption) (*ApplicationGroup, error) {
	var resource ApplicationGroup
	err := ctx.ReadResource("azure-native:desktopvirtualization/v20210114preview:ApplicationGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type applicationGroupState struct {
}

type ApplicationGroupState struct {
}

func (ApplicationGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*applicationGroupState)(nil)).Elem()
}

type applicationGroupArgs struct {
	ApplicationGroupName *string                                      `pulumi:"applicationGroupName"`
	ApplicationGroupType string                                       `pulumi:"applicationGroupType"`
	Description          *string                                      `pulumi:"description"`
	FriendlyName         *string                                      `pulumi:"friendlyName"`
	HostPoolArmPath      string                                       `pulumi:"hostPoolArmPath"`
	Identity             *ResourceModelWithAllowedPropertySetIdentity `pulumi:"identity"`
	Kind                 *string                                      `pulumi:"kind"`
	Location             *string                                      `pulumi:"location"`
	ManagedBy            *string                                      `pulumi:"managedBy"`
	MigrationRequest     *MigrationRequestProperties                  `pulumi:"migrationRequest"`
	Plan                 *ResourceModelWithAllowedPropertySetPlan     `pulumi:"plan"`
	ResourceGroupName    string                                       `pulumi:"resourceGroupName"`
	Sku                  *ResourceModelWithAllowedPropertySetSku      `pulumi:"sku"`
	Tags                 map[string]string                            `pulumi:"tags"`
}


type ApplicationGroupArgs struct {
	ApplicationGroupName pulumi.StringPtrInput
	ApplicationGroupType pulumi.StringInput
	Description          pulumi.StringPtrInput
	FriendlyName         pulumi.StringPtrInput
	HostPoolArmPath      pulumi.StringInput
	Identity             ResourceModelWithAllowedPropertySetIdentityPtrInput
	Kind                 pulumi.StringPtrInput
	Location             pulumi.StringPtrInput
	ManagedBy            pulumi.StringPtrInput
	MigrationRequest     MigrationRequestPropertiesPtrInput
	Plan                 ResourceModelWithAllowedPropertySetPlanPtrInput
	ResourceGroupName    pulumi.StringInput
	Sku                  ResourceModelWithAllowedPropertySetSkuPtrInput
	Tags                 pulumi.StringMapInput
}

func (ApplicationGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*applicationGroupArgs)(nil)).Elem()
}

type ApplicationGroupInput interface {
	pulumi.Input

	ToApplicationGroupOutput() ApplicationGroupOutput
	ToApplicationGroupOutputWithContext(ctx context.Context) ApplicationGroupOutput
}

func (*ApplicationGroup) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationGroup)(nil)).Elem()
}

func (i *ApplicationGroup) ToApplicationGroupOutput() ApplicationGroupOutput {
	return i.ToApplicationGroupOutputWithContext(context.Background())
}

func (i *ApplicationGroup) ToApplicationGroupOutputWithContext(ctx context.Context) ApplicationGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationGroupOutput)
}

type ApplicationGroupOutput struct{ *pulumi.OutputState }

func (ApplicationGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationGroup)(nil)).Elem()
}

func (o ApplicationGroupOutput) ToApplicationGroupOutput() ApplicationGroupOutput {
	return o
}

func (o ApplicationGroupOutput) ToApplicationGroupOutputWithContext(ctx context.Context) ApplicationGroupOutput {
	return o
}

func (o ApplicationGroupOutput) ApplicationGroupType() pulumi.StringOutput {
	return o.ApplyT(func(v *ApplicationGroup) pulumi.StringOutput { return v.ApplicationGroupType }).(pulumi.StringOutput)
}

func (o ApplicationGroupOutput) CloudPcResource() pulumi.BoolOutput {
	return o.ApplyT(func(v *ApplicationGroup) pulumi.BoolOutput { return v.CloudPcResource }).(pulumi.BoolOutput)
}

func (o ApplicationGroupOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApplicationGroup) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

func (o ApplicationGroupOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *ApplicationGroup) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

func (o ApplicationGroupOutput) FriendlyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApplicationGroup) pulumi.StringPtrOutput { return v.FriendlyName }).(pulumi.StringPtrOutput)
}

func (o ApplicationGroupOutput) HostPoolArmPath() pulumi.StringOutput {
	return o.ApplyT(func(v *ApplicationGroup) pulumi.StringOutput { return v.HostPoolArmPath }).(pulumi.StringOutput)
}

func (o ApplicationGroupOutput) Identity() ResourceModelWithAllowedPropertySetResponseIdentityPtrOutput {
	return o.ApplyT(func(v *ApplicationGroup) ResourceModelWithAllowedPropertySetResponseIdentityPtrOutput {
		return v.Identity
	}).(ResourceModelWithAllowedPropertySetResponseIdentityPtrOutput)
}

func (o ApplicationGroupOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApplicationGroup) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o ApplicationGroupOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApplicationGroup) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

func (o ApplicationGroupOutput) ManagedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApplicationGroup) pulumi.StringPtrOutput { return v.ManagedBy }).(pulumi.StringPtrOutput)
}

func (o ApplicationGroupOutput) MigrationRequest() MigrationRequestPropertiesResponsePtrOutput {
	return o.ApplyT(func(v *ApplicationGroup) MigrationRequestPropertiesResponsePtrOutput { return v.MigrationRequest }).(MigrationRequestPropertiesResponsePtrOutput)
}

func (o ApplicationGroupOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ApplicationGroup) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ApplicationGroupOutput) ObjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *ApplicationGroup) pulumi.StringOutput { return v.ObjectId }).(pulumi.StringOutput)
}

func (o ApplicationGroupOutput) Plan() ResourceModelWithAllowedPropertySetResponsePlanPtrOutput {
	return o.ApplyT(func(v *ApplicationGroup) ResourceModelWithAllowedPropertySetResponsePlanPtrOutput { return v.Plan }).(ResourceModelWithAllowedPropertySetResponsePlanPtrOutput)
}

func (o ApplicationGroupOutput) Sku() ResourceModelWithAllowedPropertySetResponseSkuPtrOutput {
	return o.ApplyT(func(v *ApplicationGroup) ResourceModelWithAllowedPropertySetResponseSkuPtrOutput { return v.Sku }).(ResourceModelWithAllowedPropertySetResponseSkuPtrOutput)
}

func (o ApplicationGroupOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ApplicationGroup) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o ApplicationGroupOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ApplicationGroup) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o ApplicationGroupOutput) WorkspaceArmPath() pulumi.StringOutput {
	return o.ApplyT(func(v *ApplicationGroup) pulumi.StringOutput { return v.WorkspaceArmPath }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ApplicationGroupOutput{})
}
