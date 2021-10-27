


package v20160701

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type WCFRelayAuthorizationRule struct {
	pulumi.CustomResourceState

	Name   pulumi.StringOutput      `pulumi:"name"`
	Rights pulumi.StringArrayOutput `pulumi:"rights"`
	Type   pulumi.StringOutput      `pulumi:"type"`
}


func NewWCFRelayAuthorizationRule(ctx *pulumi.Context,
	name string, args *WCFRelayAuthorizationRuleArgs, opts ...pulumi.ResourceOption) (*WCFRelayAuthorizationRule, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.NamespaceName == nil {
		return nil, errors.New("invalid value for required argument 'NamespaceName'")
	}
	if args.RelayName == nil {
		return nil, errors.New("invalid value for required argument 'RelayName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.Rights == nil {
		return nil, errors.New("invalid value for required argument 'Rights'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:relay/v20160701:WCFRelayAuthorizationRule"),
		},
		{
			Type: pulumi.String("azure-native:relay:WCFRelayAuthorizationRule"),
		},
		{
			Type: pulumi.String("azure-nextgen:relay:WCFRelayAuthorizationRule"),
		},
		{
			Type: pulumi.String("azure-native:relay/v20170401:WCFRelayAuthorizationRule"),
		},
		{
			Type: pulumi.String("azure-nextgen:relay/v20170401:WCFRelayAuthorizationRule"),
		},
	})
	opts = append(opts, aliases)
	var resource WCFRelayAuthorizationRule
	err := ctx.RegisterResource("azure-native:relay/v20160701:WCFRelayAuthorizationRule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetWCFRelayAuthorizationRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WCFRelayAuthorizationRuleState, opts ...pulumi.ResourceOption) (*WCFRelayAuthorizationRule, error) {
	var resource WCFRelayAuthorizationRule
	err := ctx.ReadResource("azure-native:relay/v20160701:WCFRelayAuthorizationRule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type wcfrelayAuthorizationRuleState struct {
}

type WCFRelayAuthorizationRuleState struct {
}

func (WCFRelayAuthorizationRuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*wcfrelayAuthorizationRuleState)(nil)).Elem()
}

type wcfrelayAuthorizationRuleArgs struct {
	AuthorizationRuleName *string  `pulumi:"authorizationRuleName"`
	NamespaceName         string   `pulumi:"namespaceName"`
	RelayName             string   `pulumi:"relayName"`
	ResourceGroupName     string   `pulumi:"resourceGroupName"`
	Rights                []string `pulumi:"rights"`
}


type WCFRelayAuthorizationRuleArgs struct {
	AuthorizationRuleName pulumi.StringPtrInput
	NamespaceName         pulumi.StringInput
	RelayName             pulumi.StringInput
	ResourceGroupName     pulumi.StringInput
	Rights                pulumi.StringArrayInput
}

func (WCFRelayAuthorizationRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*wcfrelayAuthorizationRuleArgs)(nil)).Elem()
}

type WCFRelayAuthorizationRuleInput interface {
	pulumi.Input

	ToWCFRelayAuthorizationRuleOutput() WCFRelayAuthorizationRuleOutput
	ToWCFRelayAuthorizationRuleOutputWithContext(ctx context.Context) WCFRelayAuthorizationRuleOutput
}

func (*WCFRelayAuthorizationRule) ElementType() reflect.Type {
	return reflect.TypeOf((*WCFRelayAuthorizationRule)(nil))
}

func (i *WCFRelayAuthorizationRule) ToWCFRelayAuthorizationRuleOutput() WCFRelayAuthorizationRuleOutput {
	return i.ToWCFRelayAuthorizationRuleOutputWithContext(context.Background())
}

func (i *WCFRelayAuthorizationRule) ToWCFRelayAuthorizationRuleOutputWithContext(ctx context.Context) WCFRelayAuthorizationRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WCFRelayAuthorizationRuleOutput)
}

type WCFRelayAuthorizationRuleOutput struct{ *pulumi.OutputState }

func (WCFRelayAuthorizationRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WCFRelayAuthorizationRule)(nil))
}

func (o WCFRelayAuthorizationRuleOutput) ToWCFRelayAuthorizationRuleOutput() WCFRelayAuthorizationRuleOutput {
	return o
}

func (o WCFRelayAuthorizationRuleOutput) ToWCFRelayAuthorizationRuleOutputWithContext(ctx context.Context) WCFRelayAuthorizationRuleOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*WCFRelayAuthorizationRuleInput)(nil)).Elem(), &WCFRelayAuthorizationRule{})
	pulumi.RegisterOutputType(WCFRelayAuthorizationRuleOutput{})
}
