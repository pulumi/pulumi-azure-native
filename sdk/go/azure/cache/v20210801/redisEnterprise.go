


package v20210801

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
			Type: pulumi.String("azure-nextgen:cache/v20210801:RedisEnterprise"),
		},
		{
			Type: pulumi.String("azure-native:cache:RedisEnterprise"),
		},
		{
			Type: pulumi.String("azure-nextgen:cache:RedisEnterprise"),
		},
		{
			Type: pulumi.String("azure-native:cache/v20201001preview:RedisEnterprise"),
		},
		{
			Type: pulumi.String("azure-nextgen:cache/v20201001preview:RedisEnterprise"),
		},
		{
			Type: pulumi.String("azure-native:cache/v20210201preview:RedisEnterprise"),
		},
		{
			Type: pulumi.String("azure-nextgen:cache/v20210201preview:RedisEnterprise"),
		},
		{
			Type: pulumi.String("azure-native:cache/v20210301:RedisEnterprise"),
		},
		{
			Type: pulumi.String("azure-nextgen:cache/v20210301:RedisEnterprise"),
		},
	})
	opts = append(opts, aliases)
	var resource RedisEnterprise
	err := ctx.RegisterResource("azure-native:cache/v20210801:RedisEnterprise", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetRedisEnterprise(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RedisEnterpriseState, opts ...pulumi.ResourceOption) (*RedisEnterprise, error) {
	var resource RedisEnterprise
	err := ctx.ReadResource("azure-native:cache/v20210801:RedisEnterprise", name, id, state, &resource, opts...)
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
	return reflect.TypeOf((*RedisEnterprise)(nil))
}

func (i *RedisEnterprise) ToRedisEnterpriseOutput() RedisEnterpriseOutput {
	return i.ToRedisEnterpriseOutputWithContext(context.Background())
}

func (i *RedisEnterprise) ToRedisEnterpriseOutputWithContext(ctx context.Context) RedisEnterpriseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RedisEnterpriseOutput)
}

type RedisEnterpriseOutput struct{ *pulumi.OutputState }

func (RedisEnterpriseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RedisEnterprise)(nil))
}

func (o RedisEnterpriseOutput) ToRedisEnterpriseOutput() RedisEnterpriseOutput {
	return o
}

func (o RedisEnterpriseOutput) ToRedisEnterpriseOutputWithContext(ctx context.Context) RedisEnterpriseOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RedisEnterpriseInput)(nil)).Elem(), &RedisEnterprise{})
	pulumi.RegisterOutputType(RedisEnterpriseOutput{})
}
