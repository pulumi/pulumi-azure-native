


package v20210501preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AdminRuleCollection struct {
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


func NewAdminRuleCollection(ctx *pulumi.Context,
	name string, args *AdminRuleCollectionArgs, opts ...pulumi.ResourceOption) (*AdminRuleCollection, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AppliesToGroups == nil {
		return nil, errors.New("invalid value for required argument 'AppliesToGroups'")
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
			Type: pulumi.String("azure-native:network:AdminRuleCollection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210201preview:AdminRuleCollection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220101:AdminRuleCollection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220201preview:AdminRuleCollection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220401preview:AdminRuleCollection"),
		},
	})
	opts = append(opts, aliases)
	var resource AdminRuleCollection
	err := ctx.RegisterResource("azure-native:network/v20210501preview:AdminRuleCollection", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetAdminRuleCollection(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AdminRuleCollectionState, opts ...pulumi.ResourceOption) (*AdminRuleCollection, error) {
	var resource AdminRuleCollection
	err := ctx.ReadResource("azure-native:network/v20210501preview:AdminRuleCollection", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type adminRuleCollectionState struct {
}

type AdminRuleCollectionState struct {
}

func (AdminRuleCollectionState) ElementType() reflect.Type {
	return reflect.TypeOf((*adminRuleCollectionState)(nil)).Elem()
}

type adminRuleCollectionArgs struct {
	AppliesToGroups    []NetworkManagerSecurityGroupItem `pulumi:"appliesToGroups"`
	ConfigurationName  string                            `pulumi:"configurationName"`
	Description        *string                           `pulumi:"description"`
	DisplayName        *string                           `pulumi:"displayName"`
	NetworkManagerName string                            `pulumi:"networkManagerName"`
	ResourceGroupName  string                            `pulumi:"resourceGroupName"`
	RuleCollectionName *string                           `pulumi:"ruleCollectionName"`
}


type AdminRuleCollectionArgs struct {
	AppliesToGroups    NetworkManagerSecurityGroupItemArrayInput
	ConfigurationName  pulumi.StringInput
	Description        pulumi.StringPtrInput
	DisplayName        pulumi.StringPtrInput
	NetworkManagerName pulumi.StringInput
	ResourceGroupName  pulumi.StringInput
	RuleCollectionName pulumi.StringPtrInput
}

func (AdminRuleCollectionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*adminRuleCollectionArgs)(nil)).Elem()
}

type AdminRuleCollectionInput interface {
	pulumi.Input

	ToAdminRuleCollectionOutput() AdminRuleCollectionOutput
	ToAdminRuleCollectionOutputWithContext(ctx context.Context) AdminRuleCollectionOutput
}

func (*AdminRuleCollection) ElementType() reflect.Type {
	return reflect.TypeOf((**AdminRuleCollection)(nil)).Elem()
}

func (i *AdminRuleCollection) ToAdminRuleCollectionOutput() AdminRuleCollectionOutput {
	return i.ToAdminRuleCollectionOutputWithContext(context.Background())
}

func (i *AdminRuleCollection) ToAdminRuleCollectionOutputWithContext(ctx context.Context) AdminRuleCollectionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AdminRuleCollectionOutput)
}

type AdminRuleCollectionOutput struct{ *pulumi.OutputState }

func (AdminRuleCollectionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AdminRuleCollection)(nil)).Elem()
}

func (o AdminRuleCollectionOutput) ToAdminRuleCollectionOutput() AdminRuleCollectionOutput {
	return o
}

func (o AdminRuleCollectionOutput) ToAdminRuleCollectionOutputWithContext(ctx context.Context) AdminRuleCollectionOutput {
	return o
}

func (o AdminRuleCollectionOutput) AppliesToGroups() NetworkManagerSecurityGroupItemResponseArrayOutput {
	return o.ApplyT(func(v *AdminRuleCollection) NetworkManagerSecurityGroupItemResponseArrayOutput {
		return v.AppliesToGroups
	}).(NetworkManagerSecurityGroupItemResponseArrayOutput)
}

func (o AdminRuleCollectionOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AdminRuleCollection) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

func (o AdminRuleCollectionOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AdminRuleCollection) pulumi.StringPtrOutput { return v.DisplayName }).(pulumi.StringPtrOutput)
}

func (o AdminRuleCollectionOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *AdminRuleCollection) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

func (o AdminRuleCollectionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *AdminRuleCollection) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o AdminRuleCollectionOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *AdminRuleCollection) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o AdminRuleCollectionOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *AdminRuleCollection) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o AdminRuleCollectionOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *AdminRuleCollection) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(AdminRuleCollectionOutput{})
}
