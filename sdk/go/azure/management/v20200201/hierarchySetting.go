


package v20200201

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type HierarchySetting struct {
	pulumi.CustomResourceState

	DefaultManagementGroup               pulumi.StringPtrOutput `pulumi:"defaultManagementGroup"`
	Name                                 pulumi.StringOutput    `pulumi:"name"`
	RequireAuthorizationForGroupCreation pulumi.BoolPtrOutput   `pulumi:"requireAuthorizationForGroupCreation"`
	TenantId                             pulumi.StringPtrOutput `pulumi:"tenantId"`
	Type                                 pulumi.StringOutput    `pulumi:"type"`
}


func NewHierarchySetting(ctx *pulumi.Context,
	name string, args *HierarchySettingArgs, opts ...pulumi.ResourceOption) (*HierarchySetting, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.GroupId == nil {
		return nil, errors.New("invalid value for required argument 'GroupId'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:management/v20200201:HierarchySetting"),
		},
		{
			Type: pulumi.String("azure-native:management:HierarchySetting"),
		},
		{
			Type: pulumi.String("azure-nextgen:management:HierarchySetting"),
		},
		{
			Type: pulumi.String("azure-native:management/v20200501:HierarchySetting"),
		},
		{
			Type: pulumi.String("azure-nextgen:management/v20200501:HierarchySetting"),
		},
		{
			Type: pulumi.String("azure-native:management/v20201001:HierarchySetting"),
		},
		{
			Type: pulumi.String("azure-nextgen:management/v20201001:HierarchySetting"),
		},
		{
			Type: pulumi.String("azure-native:management/v20210401:HierarchySetting"),
		},
		{
			Type: pulumi.String("azure-nextgen:management/v20210401:HierarchySetting"),
		},
	})
	opts = append(opts, aliases)
	var resource HierarchySetting
	err := ctx.RegisterResource("azure-native:management/v20200201:HierarchySetting", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetHierarchySetting(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *HierarchySettingState, opts ...pulumi.ResourceOption) (*HierarchySetting, error) {
	var resource HierarchySetting
	err := ctx.ReadResource("azure-native:management/v20200201:HierarchySetting", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type hierarchySettingState struct {
}

type HierarchySettingState struct {
}

func (HierarchySettingState) ElementType() reflect.Type {
	return reflect.TypeOf((*hierarchySettingState)(nil)).Elem()
}

type hierarchySettingArgs struct {
	DefaultManagementGroup               *string `pulumi:"defaultManagementGroup"`
	GroupId                              string  `pulumi:"groupId"`
	RequireAuthorizationForGroupCreation *bool   `pulumi:"requireAuthorizationForGroupCreation"`
}


type HierarchySettingArgs struct {
	DefaultManagementGroup               pulumi.StringPtrInput
	GroupId                              pulumi.StringInput
	RequireAuthorizationForGroupCreation pulumi.BoolPtrInput
}

func (HierarchySettingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*hierarchySettingArgs)(nil)).Elem()
}

type HierarchySettingInput interface {
	pulumi.Input

	ToHierarchySettingOutput() HierarchySettingOutput
	ToHierarchySettingOutputWithContext(ctx context.Context) HierarchySettingOutput
}

func (*HierarchySetting) ElementType() reflect.Type {
	return reflect.TypeOf((*HierarchySetting)(nil))
}

func (i *HierarchySetting) ToHierarchySettingOutput() HierarchySettingOutput {
	return i.ToHierarchySettingOutputWithContext(context.Background())
}

func (i *HierarchySetting) ToHierarchySettingOutputWithContext(ctx context.Context) HierarchySettingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HierarchySettingOutput)
}

type HierarchySettingOutput struct{ *pulumi.OutputState }

func (HierarchySettingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*HierarchySetting)(nil))
}

func (o HierarchySettingOutput) ToHierarchySettingOutput() HierarchySettingOutput {
	return o
}

func (o HierarchySettingOutput) ToHierarchySettingOutputWithContext(ctx context.Context) HierarchySettingOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*HierarchySettingInput)(nil)).Elem(), &HierarchySetting{})
	pulumi.RegisterOutputType(HierarchySettingOutput{})
}
