


package v20210201preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type UserRuleCollection struct {
	pulumi.CustomResourceState

	AppliesToGroups   NetworkManagerSecurityGroupItemResponseArrayOutput `pulumi:"appliesToGroups"`
	Description       pulumi.StringPtrOutput                             `pulumi:"description"`
	DisplayName       pulumi.StringPtrOutput                             `pulumi:"displayName"`
	Etag              pulumi.StringOutput                                `pulumi:"etag"`
	Name              pulumi.StringOutput                                `pulumi:"name"`
	ProvisioningState pulumi.StringOutput                                `pulumi:"provisioningState"`
	SystemData        SystemDataResponseOutput                           `pulumi:"systemData"`
	Type              pulumi.StringOutput                                `pulumi:"type"`
}


func NewUserRuleCollection(ctx *pulumi.Context,
	name string, args *UserRuleCollectionArgs, opts ...pulumi.ResourceOption) (*UserRuleCollection, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ConfigurationName == nil {
		return nil, errors.New("invalid value for required argument 'ConfigurationName'")
	}
	if args.NetworkManagerName == nil {
		return nil, errors.New("invalid value for required argument 'NetworkManagerName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:network:UserRuleCollection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210501preview:UserRuleCollection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220201preview:UserRuleCollection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220401preview:UserRuleCollection"),
		},
	})
	opts = append(opts, aliases)
	var resource UserRuleCollection
	err := ctx.RegisterResource("azure-native:network/v20210201preview:UserRuleCollection", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetUserRuleCollection(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *UserRuleCollectionState, opts ...pulumi.ResourceOption) (*UserRuleCollection, error) {
	var resource UserRuleCollection
	err := ctx.ReadResource("azure-native:network/v20210201preview:UserRuleCollection", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type userRuleCollectionState struct {
}

type UserRuleCollectionState struct {
}

func (UserRuleCollectionState) ElementType() reflect.Type {
	return reflect.TypeOf((*userRuleCollectionState)(nil)).Elem()
}

type userRuleCollectionArgs struct {
	AppliesToGroups    []NetworkManagerSecurityGroupItem `pulumi:"appliesToGroups"`
	ConfigurationName  string                            `pulumi:"configurationName"`
	Description        *string                           `pulumi:"description"`
	DisplayName        *string                           `pulumi:"displayName"`
	NetworkManagerName string                            `pulumi:"networkManagerName"`
	ResourceGroupName  string                            `pulumi:"resourceGroupName"`
	RuleCollectionName *string                           `pulumi:"ruleCollectionName"`
}


type UserRuleCollectionArgs struct {
	AppliesToGroups    NetworkManagerSecurityGroupItemArrayInput
	ConfigurationName  pulumi.StringInput
	Description        pulumi.StringPtrInput
	DisplayName        pulumi.StringPtrInput
	NetworkManagerName pulumi.StringInput
	ResourceGroupName  pulumi.StringInput
	RuleCollectionName pulumi.StringPtrInput
}

func (UserRuleCollectionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*userRuleCollectionArgs)(nil)).Elem()
}

type UserRuleCollectionInput interface {
	pulumi.Input

	ToUserRuleCollectionOutput() UserRuleCollectionOutput
	ToUserRuleCollectionOutputWithContext(ctx context.Context) UserRuleCollectionOutput
}

func (*UserRuleCollection) ElementType() reflect.Type {
	return reflect.TypeOf((**UserRuleCollection)(nil)).Elem()
}

func (i *UserRuleCollection) ToUserRuleCollectionOutput() UserRuleCollectionOutput {
	return i.ToUserRuleCollectionOutputWithContext(context.Background())
}

func (i *UserRuleCollection) ToUserRuleCollectionOutputWithContext(ctx context.Context) UserRuleCollectionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserRuleCollectionOutput)
}

type UserRuleCollectionOutput struct{ *pulumi.OutputState }

func (UserRuleCollectionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**UserRuleCollection)(nil)).Elem()
}

func (o UserRuleCollectionOutput) ToUserRuleCollectionOutput() UserRuleCollectionOutput {
	return o
}

func (o UserRuleCollectionOutput) ToUserRuleCollectionOutputWithContext(ctx context.Context) UserRuleCollectionOutput {
	return o
}

func (o UserRuleCollectionOutput) AppliesToGroups() NetworkManagerSecurityGroupItemResponseArrayOutput {
	return o.ApplyT(func(v *UserRuleCollection) NetworkManagerSecurityGroupItemResponseArrayOutput {
		return v.AppliesToGroups
	}).(NetworkManagerSecurityGroupItemResponseArrayOutput)
}

func (o UserRuleCollectionOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UserRuleCollection) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

func (o UserRuleCollectionOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UserRuleCollection) pulumi.StringPtrOutput { return v.DisplayName }).(pulumi.StringPtrOutput)
}

func (o UserRuleCollectionOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *UserRuleCollection) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

func (o UserRuleCollectionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *UserRuleCollection) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o UserRuleCollectionOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *UserRuleCollection) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o UserRuleCollectionOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *UserRuleCollection) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o UserRuleCollectionOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *UserRuleCollection) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(UserRuleCollectionOutput{})
}
