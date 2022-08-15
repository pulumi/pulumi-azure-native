


package v20170201

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


type RedisLinkedServer struct {
	pulumi.CustomResourceState

	LinkedRedisCacheId       pulumi.StringOutput `pulumi:"linkedRedisCacheId"`
	LinkedRedisCacheLocation pulumi.StringOutput `pulumi:"linkedRedisCacheLocation"`
	Name                     pulumi.StringOutput `pulumi:"name"`
	ProvisioningState        pulumi.StringOutput `pulumi:"provisioningState"`
	ServerRole               pulumi.StringOutput `pulumi:"serverRole"`
	Type                     pulumi.StringOutput `pulumi:"type"`
}


func NewRedisLinkedServer(ctx *pulumi.Context,
	name string, args *RedisLinkedServerArgs, opts ...pulumi.ResourceOption) (*RedisLinkedServer, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.LinkedRedisCacheId == nil {
		return nil, errors.New("invalid value for required argument 'LinkedRedisCacheId'")
	}
	if args.LinkedRedisCacheLocation == nil {
		return nil, errors.New("invalid value for required argument 'LinkedRedisCacheLocation'")
	}
	if args.Name == nil {
		return nil, errors.New("invalid value for required argument 'Name'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ServerRole == nil {
		return nil, errors.New("invalid value for required argument 'ServerRole'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:cache:RedisLinkedServer"),
		},
		{
			Type: pulumi.String("azure-native:cache/v20171001:RedisLinkedServer"),
		},
		{
			Type: pulumi.String("azure-native:cache/v20180301:RedisLinkedServer"),
		},
		{
			Type: pulumi.String("azure-native:cache/v20190701:RedisLinkedServer"),
		},
		{
			Type: pulumi.String("azure-native:cache/v20200601:RedisLinkedServer"),
		},
		{
			Type: pulumi.String("azure-native:cache/v20201201:RedisLinkedServer"),
		},
		{
			Type: pulumi.String("azure-native:cache/v20210601:RedisLinkedServer"),
		},
		{
			Type: pulumi.String("azure-native:cache/v20220501:RedisLinkedServer"),
		},
	})
	opts = append(opts, aliases)
	var resource RedisLinkedServer
	err := ctx.RegisterResource("azure-native:cache/v20170201:RedisLinkedServer", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetRedisLinkedServer(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RedisLinkedServerState, opts ...pulumi.ResourceOption) (*RedisLinkedServer, error) {
	var resource RedisLinkedServer
	err := ctx.ReadResource("azure-native:cache/v20170201:RedisLinkedServer", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type redisLinkedServerState struct {
}

type RedisLinkedServerState struct {
}

func (RedisLinkedServerState) ElementType() reflect.Type {
	return reflect.TypeOf((*redisLinkedServerState)(nil)).Elem()
}

type redisLinkedServerArgs struct {
	LinkedRedisCacheId       string          `pulumi:"linkedRedisCacheId"`
	LinkedRedisCacheLocation string          `pulumi:"linkedRedisCacheLocation"`
	LinkedServerName         *string         `pulumi:"linkedServerName"`
	Name                     string          `pulumi:"name"`
	ResourceGroupName        string          `pulumi:"resourceGroupName"`
	ServerRole               ReplicationRole `pulumi:"serverRole"`
}


type RedisLinkedServerArgs struct {
	LinkedRedisCacheId       pulumi.StringInput
	LinkedRedisCacheLocation pulumi.StringInput
	LinkedServerName         pulumi.StringPtrInput
	Name                     pulumi.StringInput
	ResourceGroupName        pulumi.StringInput
	ServerRole               ReplicationRoleInput
}

func (RedisLinkedServerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*redisLinkedServerArgs)(nil)).Elem()
}

type RedisLinkedServerInput interface {
	pulumi.Input

	ToRedisLinkedServerOutput() RedisLinkedServerOutput
	ToRedisLinkedServerOutputWithContext(ctx context.Context) RedisLinkedServerOutput
}

func (*RedisLinkedServer) ElementType() reflect.Type {
	return reflect.TypeOf((**RedisLinkedServer)(nil)).Elem()
}

func (i *RedisLinkedServer) ToRedisLinkedServerOutput() RedisLinkedServerOutput {
	return i.ToRedisLinkedServerOutputWithContext(context.Background())
}

func (i *RedisLinkedServer) ToRedisLinkedServerOutputWithContext(ctx context.Context) RedisLinkedServerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RedisLinkedServerOutput)
}

type RedisLinkedServerOutput struct{ *pulumi.OutputState }

func (RedisLinkedServerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RedisLinkedServer)(nil)).Elem()
}

func (o RedisLinkedServerOutput) ToRedisLinkedServerOutput() RedisLinkedServerOutput {
	return o
}

func (o RedisLinkedServerOutput) ToRedisLinkedServerOutputWithContext(ctx context.Context) RedisLinkedServerOutput {
	return o
}

func (o RedisLinkedServerOutput) LinkedRedisCacheId() pulumi.StringOutput {
	return o.ApplyT(func(v *RedisLinkedServer) pulumi.StringOutput { return v.LinkedRedisCacheId }).(pulumi.StringOutput)
}

func (o RedisLinkedServerOutput) LinkedRedisCacheLocation() pulumi.StringOutput {
	return o.ApplyT(func(v *RedisLinkedServer) pulumi.StringOutput { return v.LinkedRedisCacheLocation }).(pulumi.StringOutput)
}

func (o RedisLinkedServerOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *RedisLinkedServer) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o RedisLinkedServerOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *RedisLinkedServer) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o RedisLinkedServerOutput) ServerRole() pulumi.StringOutput {
	return o.ApplyT(func(v *RedisLinkedServer) pulumi.StringOutput { return v.ServerRole }).(pulumi.StringOutput)
}

func (o RedisLinkedServerOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *RedisLinkedServer) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(RedisLinkedServerOutput{})
}
