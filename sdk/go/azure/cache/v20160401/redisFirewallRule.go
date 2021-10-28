


package v20160401

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type RedisFirewallRule struct {
	pulumi.CustomResourceState

	EndIP   pulumi.StringOutput `pulumi:"endIP"`
	Name    pulumi.StringOutput `pulumi:"name"`
	StartIP pulumi.StringOutput `pulumi:"startIP"`
	Type    pulumi.StringOutput `pulumi:"type"`
}


func NewRedisFirewallRule(ctx *pulumi.Context,
	name string, args *RedisFirewallRuleArgs, opts ...pulumi.ResourceOption) (*RedisFirewallRule, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.CacheName == nil {
		return nil, errors.New("invalid value for required argument 'CacheName'")
	}
	if args.EndIP == nil {
		return nil, errors.New("invalid value for required argument 'EndIP'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.StartIP == nil {
		return nil, errors.New("invalid value for required argument 'StartIP'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:cache/v20160401:RedisFirewallRule"),
		},
		{
			Type: pulumi.String("azure-native:cache:RedisFirewallRule"),
		},
		{
			Type: pulumi.String("azure-nextgen:cache:RedisFirewallRule"),
		},
		{
			Type: pulumi.String("azure-native:cache/v20170201:RedisFirewallRule"),
		},
		{
			Type: pulumi.String("azure-nextgen:cache/v20170201:RedisFirewallRule"),
		},
		{
			Type: pulumi.String("azure-native:cache/v20171001:RedisFirewallRule"),
		},
		{
			Type: pulumi.String("azure-nextgen:cache/v20171001:RedisFirewallRule"),
		},
		{
			Type: pulumi.String("azure-native:cache/v20180301:RedisFirewallRule"),
		},
		{
			Type: pulumi.String("azure-nextgen:cache/v20180301:RedisFirewallRule"),
		},
		{
			Type: pulumi.String("azure-native:cache/v20190701:RedisFirewallRule"),
		},
		{
			Type: pulumi.String("azure-nextgen:cache/v20190701:RedisFirewallRule"),
		},
		{
			Type: pulumi.String("azure-native:cache/v20200601:RedisFirewallRule"),
		},
		{
			Type: pulumi.String("azure-nextgen:cache/v20200601:RedisFirewallRule"),
		},
		{
			Type: pulumi.String("azure-native:cache/v20201201:RedisFirewallRule"),
		},
		{
			Type: pulumi.String("azure-nextgen:cache/v20201201:RedisFirewallRule"),
		},
	})
	opts = append(opts, aliases)
	var resource RedisFirewallRule
	err := ctx.RegisterResource("azure-native:cache/v20160401:RedisFirewallRule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetRedisFirewallRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RedisFirewallRuleState, opts ...pulumi.ResourceOption) (*RedisFirewallRule, error) {
	var resource RedisFirewallRule
	err := ctx.ReadResource("azure-native:cache/v20160401:RedisFirewallRule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type redisFirewallRuleState struct {
}

type RedisFirewallRuleState struct {
}

func (RedisFirewallRuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*redisFirewallRuleState)(nil)).Elem()
}

type redisFirewallRuleArgs struct {
	CacheName         string  `pulumi:"cacheName"`
	EndIP             string  `pulumi:"endIP"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
	RuleName          *string `pulumi:"ruleName"`
	StartIP           string  `pulumi:"startIP"`
}


type RedisFirewallRuleArgs struct {
	CacheName         pulumi.StringInput
	EndIP             pulumi.StringInput
	ResourceGroupName pulumi.StringInput
	RuleName          pulumi.StringPtrInput
	StartIP           pulumi.StringInput
}

func (RedisFirewallRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*redisFirewallRuleArgs)(nil)).Elem()
}

type RedisFirewallRuleInput interface {
	pulumi.Input

	ToRedisFirewallRuleOutput() RedisFirewallRuleOutput
	ToRedisFirewallRuleOutputWithContext(ctx context.Context) RedisFirewallRuleOutput
}

func (*RedisFirewallRule) ElementType() reflect.Type {
	return reflect.TypeOf((*RedisFirewallRule)(nil))
}

func (i *RedisFirewallRule) ToRedisFirewallRuleOutput() RedisFirewallRuleOutput {
	return i.ToRedisFirewallRuleOutputWithContext(context.Background())
}

func (i *RedisFirewallRule) ToRedisFirewallRuleOutputWithContext(ctx context.Context) RedisFirewallRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RedisFirewallRuleOutput)
}

type RedisFirewallRuleOutput struct{ *pulumi.OutputState }

func (RedisFirewallRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RedisFirewallRule)(nil))
}

func (o RedisFirewallRuleOutput) ToRedisFirewallRuleOutput() RedisFirewallRuleOutput {
	return o
}

func (o RedisFirewallRuleOutput) ToRedisFirewallRuleOutputWithContext(ctx context.Context) RedisFirewallRuleOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(RedisFirewallRuleOutput{})
}
