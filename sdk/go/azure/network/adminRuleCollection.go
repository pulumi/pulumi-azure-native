


package network

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
			Type: pulumi.String("azure-native:network/v20210201preview:AdminRuleCollection"),
		},
	})
	opts = append(opts, aliases)
	var resource AdminRuleCollection
	err := ctx.RegisterResource("azure-native:network:AdminRuleCollection", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetAdminRuleCollection(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AdminRuleCollectionState, opts ...pulumi.ResourceOption) (*AdminRuleCollection, error) {
	var resource AdminRuleCollection
	err := ctx.ReadResource("azure-native:network:AdminRuleCollection", name, id, state, &resource, opts...)
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
	return reflect.TypeOf((*AdminRuleCollection)(nil))
}

func (i *AdminRuleCollection) ToAdminRuleCollectionOutput() AdminRuleCollectionOutput {
	return i.ToAdminRuleCollectionOutputWithContext(context.Background())
}

func (i *AdminRuleCollection) ToAdminRuleCollectionOutputWithContext(ctx context.Context) AdminRuleCollectionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AdminRuleCollectionOutput)
}

type AdminRuleCollectionOutput struct{ *pulumi.OutputState }

func (AdminRuleCollectionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AdminRuleCollection)(nil))
}

func (o AdminRuleCollectionOutput) ToAdminRuleCollectionOutput() AdminRuleCollectionOutput {
	return o
}

func (o AdminRuleCollectionOutput) ToAdminRuleCollectionOutputWithContext(ctx context.Context) AdminRuleCollectionOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(AdminRuleCollectionOutput{})
}
