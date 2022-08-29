


package v20220101

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type RedisEnterprise struct {
	pulumi.CustomResourceState

	HostName                   pulumi.StringOutput                          `pulumi:"hostName"`
	Location                   pulumi.StringOutput                          `pulumi:"location"`
	MinimumTlsVersion          pulumi.StringPtrOutput                       `pulumi:"minimumTlsVersion"`
	Name                       pulumi.StringOutput                          `pulumi:"name"`
	PrivateEndpointConnections PrivateEndpointConnectionResponseArrayOutput `pulumi:"privateEndpointConnections"`
	ProvisioningState          pulumi.StringOutput                          `pulumi:"provisioningState"`
	RedisVersion               pulumi.StringOutput                          `pulumi:"redisVersion"`
	ResourceState              pulumi.StringOutput                          `pulumi:"resourceState"`
	Sku                        EnterpriseSkuResponseOutput                  `pulumi:"sku"`
	Tags                       pulumi.StringMapOutput                       `pulumi:"tags"`
	Type                       pulumi.StringOutput                          `pulumi:"type"`
	Zones                      pulumi.StringArrayOutput                     `pulumi:"zones"`
}


func NewRedisEnterprise(ctx *pulumi.Context,
	name string, args *RedisEnterpriseArgs, opts ...pulumi.ResourceOption) (*RedisEnterprise, error) {
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
			Type: pulumi.String("azure-native:cache:RedisEnterprise"),
		},
		{
			Type: pulumi.String("azure-native:cache/v20201001preview:RedisEnterprise"),
		},
		{
			Type: pulumi.String("azure-native:cache/v20210201preview:RedisEnterprise"),
		},
		{
			Type: pulumi.String("azure-native:cache/v20210301:RedisEnterprise"),
		},
		{
			Type: pulumi.String("azure-native:cache/v20210801:RedisEnterprise"),
		},
	})
	opts = append(opts, aliases)
	var resource RedisEnterprise
	err := ctx.RegisterResource("azure-native:cache/v20220101:RedisEnterprise", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetRedisEnterprise(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RedisEnterpriseState, opts ...pulumi.ResourceOption) (*RedisEnterprise, error) {
	var resource RedisEnterprise
	err := ctx.ReadResource("azure-native:cache/v20220101:RedisEnterprise", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type redisEnterpriseState struct {
}

type RedisEnterpriseState struct {
}

func (RedisEnterpriseState) ElementType() reflect.Type {
	return reflect.TypeOf((*redisEnterpriseState)(nil)).Elem()
}

type redisEnterpriseArgs struct {
	ClusterName       *string           `pulumi:"clusterName"`
	Location          *string           `pulumi:"location"`
	MinimumTlsVersion *string           `pulumi:"minimumTlsVersion"`
	ResourceGroupName string            `pulumi:"resourceGroupName"`
	Sku               EnterpriseSku     `pulumi:"sku"`
	Tags              map[string]string `pulumi:"tags"`
	Zones             []string          `pulumi:"zones"`
}


type RedisEnterpriseArgs struct {
	ClusterName       pulumi.StringPtrInput
	Location          pulumi.StringPtrInput
	MinimumTlsVersion pulumi.StringPtrInput
	ResourceGroupName pulumi.StringInput
	Sku               EnterpriseSkuInput
	Tags              pulumi.StringMapInput
	Zones             pulumi.StringArrayInput
}

func (RedisEnterpriseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*redisEnterpriseArgs)(nil)).Elem()
}

type RedisEnterpriseInput interface {
	pulumi.Input

	ToRedisEnterpriseOutput() RedisEnterpriseOutput
	ToRedisEnterpriseOutputWithContext(ctx context.Context) RedisEnterpriseOutput
}

func (*RedisEnterprise) ElementType() reflect.Type {
	return reflect.TypeOf((**RedisEnterprise)(nil)).Elem()
}

func (i *RedisEnterprise) ToRedisEnterpriseOutput() RedisEnterpriseOutput {
	return i.ToRedisEnterpriseOutputWithContext(context.Background())
}

func (i *RedisEnterprise) ToRedisEnterpriseOutputWithContext(ctx context.Context) RedisEnterpriseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RedisEnterpriseOutput)
}

type RedisEnterpriseOutput struct{ *pulumi.OutputState }

func (RedisEnterpriseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RedisEnterprise)(nil)).Elem()
}

func (o RedisEnterpriseOutput) ToRedisEnterpriseOutput() RedisEnterpriseOutput {
	return o
}

func (o RedisEnterpriseOutput) ToRedisEnterpriseOutputWithContext(ctx context.Context) RedisEnterpriseOutput {
	return o
}

func (o RedisEnterpriseOutput) HostName() pulumi.StringOutput {
	return o.ApplyT(func(v *RedisEnterprise) pulumi.StringOutput { return v.HostName }).(pulumi.StringOutput)
}

func (o RedisEnterpriseOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *RedisEnterprise) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o RedisEnterpriseOutput) MinimumTlsVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RedisEnterprise) pulumi.StringPtrOutput { return v.MinimumTlsVersion }).(pulumi.StringPtrOutput)
}

func (o RedisEnterpriseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *RedisEnterprise) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o RedisEnterpriseOutput) PrivateEndpointConnections() PrivateEndpointConnectionResponseArrayOutput {
	return o.ApplyT(func(v *RedisEnterprise) PrivateEndpointConnectionResponseArrayOutput {
		return v.PrivateEndpointConnections
	}).(PrivateEndpointConnectionResponseArrayOutput)
}

func (o RedisEnterpriseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *RedisEnterprise) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o RedisEnterpriseOutput) RedisVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *RedisEnterprise) pulumi.StringOutput { return v.RedisVersion }).(pulumi.StringOutput)
}

func (o RedisEnterpriseOutput) ResourceState() pulumi.StringOutput {
	return o.ApplyT(func(v *RedisEnterprise) pulumi.StringOutput { return v.ResourceState }).(pulumi.StringOutput)
}

func (o RedisEnterpriseOutput) Sku() EnterpriseSkuResponseOutput {
	return o.ApplyT(func(v *RedisEnterprise) EnterpriseSkuResponseOutput { return v.Sku }).(EnterpriseSkuResponseOutput)
}

func (o RedisEnterpriseOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *RedisEnterprise) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o RedisEnterpriseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *RedisEnterprise) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o RedisEnterpriseOutput) Zones() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *RedisEnterprise) pulumi.StringArrayOutput { return v.Zones }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(RedisEnterpriseOutput{})
}
