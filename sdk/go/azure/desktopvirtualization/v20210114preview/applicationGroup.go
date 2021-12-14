


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

func init() {
	pulumi.RegisterOutputType(ApplicationGroupOutput{})
}
