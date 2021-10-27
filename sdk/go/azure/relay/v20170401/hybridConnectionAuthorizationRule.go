


package v20170401

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type HybridConnectionAuthorizationRule struct {
	pulumi.CustomResourceState

	Name   pulumi.StringOutput      `pulumi:"name"`
	Rights pulumi.StringArrayOutput `pulumi:"rights"`
	Type   pulumi.StringOutput      `pulumi:"type"`
}


func NewHybridConnectionAuthorizationRule(ctx *pulumi.Context,
	name string, args *HybridConnectionAuthorizationRuleArgs, opts ...pulumi.ResourceOption) (*HybridConnectionAuthorizationRule, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.HybridConnectionName == nil {
		return nil, errors.New("invalid value for required argument 'HybridConnectionName'")
	}
	if args.NamespaceName == nil {
		return nil, errors.New("invalid value for required argument 'NamespaceName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.Rights == nil {
		return nil, errors.New("invalid value for required argument 'Rights'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:relay/v20170401:HybridConnectionAuthorizationRule"),
		},
		{
			Type: pulumi.String("azure-native:relay:HybridConnectionAuthorizationRule"),
		},
		{
			Type: pulumi.String("azure-nextgen:relay:HybridConnectionAuthorizationRule"),
		},
		{
			Type: pulumi.String("azure-native:relay/v20160701:HybridConnectionAuthorizationRule"),
		},
		{
			Type: pulumi.String("azure-nextgen:relay/v20160701:HybridConnectionAuthorizationRule"),
		},
	})
	opts = append(opts, aliases)
	var resource HybridConnectionAuthorizationRule
	err := ctx.RegisterResource("azure-native:relay/v20170401:HybridConnectionAuthorizationRule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetHybridConnectionAuthorizationRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *HybridConnectionAuthorizationRuleState, opts ...pulumi.ResourceOption) (*HybridConnectionAuthorizationRule, error) {
	var resource HybridConnectionAuthorizationRule
	err := ctx.ReadResource("azure-native:relay/v20170401:HybridConnectionAuthorizationRule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type hybridConnectionAuthorizationRuleState struct {
}

type HybridConnectionAuthorizationRuleState struct {
}

func (HybridConnectionAuthorizationRuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*hybridConnectionAuthorizationRuleState)(nil)).Elem()
}

type hybridConnectionAuthorizationRuleArgs struct {
	AuthorizationRuleName *string        `pulumi:"authorizationRuleName"`
	HybridConnectionName  string         `pulumi:"hybridConnectionName"`
	NamespaceName         string         `pulumi:"namespaceName"`
	ResourceGroupName     string         `pulumi:"resourceGroupName"`
	Rights                []AccessRights `pulumi:"rights"`
}


type HybridConnectionAuthorizationRuleArgs struct {
	AuthorizationRuleName pulumi.StringPtrInput
	HybridConnectionName  pulumi.StringInput
	NamespaceName         pulumi.StringInput
	ResourceGroupName     pulumi.StringInput
	Rights                AccessRightsArrayInput
}

func (HybridConnectionAuthorizationRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*hybridConnectionAuthorizationRuleArgs)(nil)).Elem()
}

type HybridConnectionAuthorizationRuleInput interface {
	pulumi.Input

	ToHybridConnectionAuthorizationRuleOutput() HybridConnectionAuthorizationRuleOutput
	ToHybridConnectionAuthorizationRuleOutputWithContext(ctx context.Context) HybridConnectionAuthorizationRuleOutput
}

func (*HybridConnectionAuthorizationRule) ElementType() reflect.Type {
	return reflect.TypeOf((*HybridConnectionAuthorizationRule)(nil))
}

func (i *HybridConnectionAuthorizationRule) ToHybridConnectionAuthorizationRuleOutput() HybridConnectionAuthorizationRuleOutput {
	return i.ToHybridConnectionAuthorizationRuleOutputWithContext(context.Background())
}

func (i *HybridConnectionAuthorizationRule) ToHybridConnectionAuthorizationRuleOutputWithContext(ctx context.Context) HybridConnectionAuthorizationRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HybridConnectionAuthorizationRuleOutput)
}

type HybridConnectionAuthorizationRuleOutput struct{ *pulumi.OutputState }

func (HybridConnectionAuthorizationRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*HybridConnectionAuthorizationRule)(nil))
}

func (o HybridConnectionAuthorizationRuleOutput) ToHybridConnectionAuthorizationRuleOutput() HybridConnectionAuthorizationRuleOutput {
	return o
}

func (o HybridConnectionAuthorizationRuleOutput) ToHybridConnectionAuthorizationRuleOutputWithContext(ctx context.Context) HybridConnectionAuthorizationRuleOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*HybridConnectionAuthorizationRuleInput)(nil)).Elem(), &HybridConnectionAuthorizationRule{})
	pulumi.RegisterOutputType(HybridConnectionAuthorizationRuleOutput{})
}
