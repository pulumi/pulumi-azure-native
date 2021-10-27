


package v20170201

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Redis struct {
	pulumi.CustomResourceState

	AccessKeys         RedisAccessKeysResponseOutput       `pulumi:"accessKeys"`
	EnableNonSslPort   pulumi.BoolPtrOutput                `pulumi:"enableNonSslPort"`
	HostName           pulumi.StringOutput                 `pulumi:"hostName"`
	LinkedServers      RedisLinkedServerListResponseOutput `pulumi:"linkedServers"`
	Location           pulumi.StringOutput                 `pulumi:"location"`
	Name               pulumi.StringOutput                 `pulumi:"name"`
	Port               pulumi.IntOutput                    `pulumi:"port"`
	ProvisioningState  pulumi.StringOutput                 `pulumi:"provisioningState"`
	RedisConfiguration pulumi.StringMapOutput              `pulumi:"redisConfiguration"`
	RedisVersion       pulumi.StringOutput                 `pulumi:"redisVersion"`
	ShardCount         pulumi.IntPtrOutput                 `pulumi:"shardCount"`
	Sku                SkuResponsePtrOutput                `pulumi:"sku"`
	SslPort            pulumi.IntOutput                    `pulumi:"sslPort"`
	StaticIP           pulumi.StringPtrOutput              `pulumi:"staticIP"`
	SubnetId           pulumi.StringPtrOutput              `pulumi:"subnetId"`
	Tags               pulumi.StringMapOutput              `pulumi:"tags"`
	TenantSettings     pulumi.StringMapOutput              `pulumi:"tenantSettings"`
	Type               pulumi.StringOutput                 `pulumi:"type"`
}


func NewRedis(ctx *pulumi.Context,
	name string, args *RedisArgs, opts ...pulumi.ResourceOption) (*Redis, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.Sku == nil {
		return nil, errors.New("invalid value for required argument 'Sku'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:cache/v20170201:Redis"),
		},
		{
			Type: pulumi.String("azure-native:cache:Redis"),
		},
		{
			Type: pulumi.String("azure-nextgen:cache:Redis"),
		},
		{
			Type: pulumi.String("azure-native:cache/v20150801:Redis"),
		},
		{
			Type: pulumi.String("azure-nextgen:cache/v20150801:Redis"),
		},
		{
			Type: pulumi.String("azure-native:cache/v20160401:Redis"),
		},
		{
			Type: pulumi.String("azure-nextgen:cache/v20160401:Redis"),
		},
		{
			Type: pulumi.String("azure-native:cache/v20171001:Redis"),
		},
		{
			Type: pulumi.String("azure-nextgen:cache/v20171001:Redis"),
		},
		{
			Type: pulumi.String("azure-native:cache/v20180301:Redis"),
		},
		{
			Type: pulumi.String("azure-nextgen:cache/v20180301:Redis"),
		},
		{
			Type: pulumi.String("azure-native:cache/v20190701:Redis"),
		},
		{
			Type: pulumi.String("azure-nextgen:cache/v20190701:Redis"),
		},
		{
			Type: pulumi.String("azure-native:cache/v20200601:Redis"),
		},
		{
			Type: pulumi.String("azure-nextgen:cache/v20200601:Redis"),
		},
		{
			Type: pulumi.String("azure-native:cache/v20201201:Redis"),
		},
		{
			Type: pulumi.String("azure-nextgen:cache/v20201201:Redis"),
		},
	})
	opts = append(opts, aliases)
	var resource Redis
	err := ctx.RegisterResource("azure-native:cache/v20170201:Redis", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetRedis(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RedisState, opts ...pulumi.ResourceOption) (*Redis, error) {
	var resource Redis
	err := ctx.ReadResource("azure-native:cache/v20170201:Redis", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type redisState struct {
}

type RedisState struct {
}

func (RedisState) ElementType() reflect.Type {
	return reflect.TypeOf((*redisState)(nil)).Elem()
}

type redisArgs struct {
	EnableNonSslPort   *bool             `pulumi:"enableNonSslPort"`
	Location           *string           `pulumi:"location"`
	Name               *string           `pulumi:"name"`
	RedisConfiguration map[string]string `pulumi:"redisConfiguration"`
	ResourceGroupName  string            `pulumi:"resourceGroupName"`
	ShardCount         *int              `pulumi:"shardCount"`
	Sku                Sku               `pulumi:"sku"`
	StaticIP           *string           `pulumi:"staticIP"`
	SubnetId           *string           `pulumi:"subnetId"`
	Tags               map[string]string `pulumi:"tags"`
	TenantSettings     map[string]string `pulumi:"tenantSettings"`
}


type RedisArgs struct {
	EnableNonSslPort   pulumi.BoolPtrInput
	Location           pulumi.StringPtrInput
	Name               pulumi.StringPtrInput
	RedisConfiguration pulumi.StringMapInput
	ResourceGroupName  pulumi.StringInput
	ShardCount         pulumi.IntPtrInput
	Sku                SkuInput
	StaticIP           pulumi.StringPtrInput
	SubnetId           pulumi.StringPtrInput
	Tags               pulumi.StringMapInput
	TenantSettings     pulumi.StringMapInput
}

func (RedisArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*redisArgs)(nil)).Elem()
}

type RedisInput interface {
	pulumi.Input

	ToRedisOutput() RedisOutput
	ToRedisOutputWithContext(ctx context.Context) RedisOutput
}

func (*Redis) ElementType() reflect.Type {
	return reflect.TypeOf((*Redis)(nil))
}

func (i *Redis) ToRedisOutput() RedisOutput {
	return i.ToRedisOutputWithContext(context.Background())
}

func (i *Redis) ToRedisOutputWithContext(ctx context.Context) RedisOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RedisOutput)
}

type RedisOutput struct{ *pulumi.OutputState }

func (RedisOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Redis)(nil))
}

func (o RedisOutput) ToRedisOutput() RedisOutput {
	return o
}

func (o RedisOutput) ToRedisOutputWithContext(ctx context.Context) RedisOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RedisInput)(nil)).Elem(), &Redis{})
	pulumi.RegisterOutputType(RedisOutput{})
}
