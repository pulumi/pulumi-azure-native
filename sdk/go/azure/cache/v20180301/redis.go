


package v20180301

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


type Redis struct {
	pulumi.CustomResourceState

	AccessKeys         RedisAccessKeysResponseOutput        `pulumi:"accessKeys"`
	EnableNonSslPort   pulumi.BoolPtrOutput                 `pulumi:"enableNonSslPort"`
	HostName           pulumi.StringOutput                  `pulumi:"hostName"`
	LinkedServers      RedisLinkedServerResponseArrayOutput `pulumi:"linkedServers"`
	Location           pulumi.StringOutput                  `pulumi:"location"`
	MinimumTlsVersion  pulumi.StringPtrOutput               `pulumi:"minimumTlsVersion"`
	Name               pulumi.StringOutput                  `pulumi:"name"`
	Port               pulumi.IntOutput                     `pulumi:"port"`
	ProvisioningState  pulumi.StringOutput                  `pulumi:"provisioningState"`
	RedisConfiguration pulumi.StringMapOutput               `pulumi:"redisConfiguration"`
	RedisVersion       pulumi.StringOutput                  `pulumi:"redisVersion"`
	ShardCount         pulumi.IntPtrOutput                  `pulumi:"shardCount"`
	Sku                SkuResponseOutput                    `pulumi:"sku"`
	SslPort            pulumi.IntOutput                     `pulumi:"sslPort"`
	StaticIP           pulumi.StringPtrOutput               `pulumi:"staticIP"`
	SubnetId           pulumi.StringPtrOutput               `pulumi:"subnetId"`
	Tags               pulumi.StringMapOutput               `pulumi:"tags"`
	TenantSettings     pulumi.StringMapOutput               `pulumi:"tenantSettings"`
	Type               pulumi.StringOutput                  `pulumi:"type"`
	Zones              pulumi.StringArrayOutput             `pulumi:"zones"`
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
			Type: pulumi.String("azure-native:cache:Redis"),
		},
		{
			Type: pulumi.String("azure-native:cache/v20150801:Redis"),
		},
		{
			Type: pulumi.String("azure-native:cache/v20160401:Redis"),
		},
		{
			Type: pulumi.String("azure-native:cache/v20170201:Redis"),
		},
		{
			Type: pulumi.String("azure-native:cache/v20171001:Redis"),
		},
		{
			Type: pulumi.String("azure-native:cache/v20190701:Redis"),
		},
		{
			Type: pulumi.String("azure-native:cache/v20200601:Redis"),
		},
		{
			Type: pulumi.String("azure-native:cache/v20201201:Redis"),
		},
		{
			Type: pulumi.String("azure-native:cache/v20210601:Redis"),
		},
		{
			Type: pulumi.String("azure-native:cache/v20220501:Redis"),
		},
		{
			Type: pulumi.String("azure-native:cache/v20220601:Redis"),
		},
	})
	opts = append(opts, aliases)
	var resource Redis
	err := ctx.RegisterResource("azure-native:cache/v20180301:Redis", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetRedis(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RedisState, opts ...pulumi.ResourceOption) (*Redis, error) {
	var resource Redis
	err := ctx.ReadResource("azure-native:cache/v20180301:Redis", name, id, state, &resource, opts...)
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
	MinimumTlsVersion  *string           `pulumi:"minimumTlsVersion"`
	Name               *string           `pulumi:"name"`
	RedisConfiguration map[string]string `pulumi:"redisConfiguration"`
	ResourceGroupName  string            `pulumi:"resourceGroupName"`
	ShardCount         *int              `pulumi:"shardCount"`
	Sku                Sku               `pulumi:"sku"`
	StaticIP           *string           `pulumi:"staticIP"`
	SubnetId           *string           `pulumi:"subnetId"`
	Tags               map[string]string `pulumi:"tags"`
	TenantSettings     map[string]string `pulumi:"tenantSettings"`
	Zones              []string          `pulumi:"zones"`
}


type RedisArgs struct {
	EnableNonSslPort   pulumi.BoolPtrInput
	Location           pulumi.StringPtrInput
	MinimumTlsVersion  pulumi.StringPtrInput
	Name               pulumi.StringPtrInput
	RedisConfiguration pulumi.StringMapInput
	ResourceGroupName  pulumi.StringInput
	ShardCount         pulumi.IntPtrInput
	Sku                SkuInput
	StaticIP           pulumi.StringPtrInput
	SubnetId           pulumi.StringPtrInput
	Tags               pulumi.StringMapInput
	TenantSettings     pulumi.StringMapInput
	Zones              pulumi.StringArrayInput
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
	return reflect.TypeOf((**Redis)(nil)).Elem()
}

func (i *Redis) ToRedisOutput() RedisOutput {
	return i.ToRedisOutputWithContext(context.Background())
}

func (i *Redis) ToRedisOutputWithContext(ctx context.Context) RedisOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RedisOutput)
}

type RedisOutput struct{ *pulumi.OutputState }

func (RedisOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Redis)(nil)).Elem()
}

func (o RedisOutput) ToRedisOutput() RedisOutput {
	return o
}

func (o RedisOutput) ToRedisOutputWithContext(ctx context.Context) RedisOutput {
	return o
}

func (o RedisOutput) AccessKeys() RedisAccessKeysResponseOutput {
	return o.ApplyT(func(v *Redis) RedisAccessKeysResponseOutput { return v.AccessKeys }).(RedisAccessKeysResponseOutput)
}

func (o RedisOutput) EnableNonSslPort() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Redis) pulumi.BoolPtrOutput { return v.EnableNonSslPort }).(pulumi.BoolPtrOutput)
}

func (o RedisOutput) HostName() pulumi.StringOutput {
	return o.ApplyT(func(v *Redis) pulumi.StringOutput { return v.HostName }).(pulumi.StringOutput)
}

func (o RedisOutput) LinkedServers() RedisLinkedServerResponseArrayOutput {
	return o.ApplyT(func(v *Redis) RedisLinkedServerResponseArrayOutput { return v.LinkedServers }).(RedisLinkedServerResponseArrayOutput)
}

func (o RedisOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *Redis) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o RedisOutput) MinimumTlsVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Redis) pulumi.StringPtrOutput { return v.MinimumTlsVersion }).(pulumi.StringPtrOutput)
}

func (o RedisOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Redis) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o RedisOutput) Port() pulumi.IntOutput {
	return o.ApplyT(func(v *Redis) pulumi.IntOutput { return v.Port }).(pulumi.IntOutput)
}

func (o RedisOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *Redis) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o RedisOutput) RedisConfiguration() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Redis) pulumi.StringMapOutput { return v.RedisConfiguration }).(pulumi.StringMapOutput)
}

func (o RedisOutput) RedisVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *Redis) pulumi.StringOutput { return v.RedisVersion }).(pulumi.StringOutput)
}

func (o RedisOutput) ShardCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Redis) pulumi.IntPtrOutput { return v.ShardCount }).(pulumi.IntPtrOutput)
}

func (o RedisOutput) Sku() SkuResponseOutput {
	return o.ApplyT(func(v *Redis) SkuResponseOutput { return v.Sku }).(SkuResponseOutput)
}

func (o RedisOutput) SslPort() pulumi.IntOutput {
	return o.ApplyT(func(v *Redis) pulumi.IntOutput { return v.SslPort }).(pulumi.IntOutput)
}

func (o RedisOutput) StaticIP() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Redis) pulumi.StringPtrOutput { return v.StaticIP }).(pulumi.StringPtrOutput)
}

func (o RedisOutput) SubnetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Redis) pulumi.StringPtrOutput { return v.SubnetId }).(pulumi.StringPtrOutput)
}

func (o RedisOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Redis) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o RedisOutput) TenantSettings() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Redis) pulumi.StringMapOutput { return v.TenantSettings }).(pulumi.StringMapOutput)
}

func (o RedisOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Redis) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o RedisOutput) Zones() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Redis) pulumi.StringArrayOutput { return v.Zones }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(RedisOutput{})
}
