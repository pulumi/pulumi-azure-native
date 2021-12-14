


package v20191101

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Cache struct {
	pulumi.CustomResourceState

	CacheSizeGB       pulumi.IntPtrOutput                 `pulumi:"cacheSizeGB"`
	Health            CacheHealthResponseOutput           `pulumi:"health"`
	Location          pulumi.StringPtrOutput              `pulumi:"location"`
	MountAddresses    pulumi.StringArrayOutput            `pulumi:"mountAddresses"`
	Name              pulumi.StringOutput                 `pulumi:"name"`
	ProvisioningState pulumi.StringPtrOutput              `pulumi:"provisioningState"`
	Sku               CacheResponseSkuPtrOutput           `pulumi:"sku"`
	Subnet            pulumi.StringPtrOutput              `pulumi:"subnet"`
	Tags              pulumi.AnyOutput                    `pulumi:"tags"`
	Type              pulumi.StringOutput                 `pulumi:"type"`
	UpgradeStatus     CacheUpgradeStatusResponsePtrOutput `pulumi:"upgradeStatus"`
}


func NewCache(ctx *pulumi.Context,
	name string, args *CacheArgs, opts ...pulumi.ResourceOption) (*Cache, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:storagecache:Cache"),
		},
		{
			Type: pulumi.String("azure-native:storagecache/v20190801preview:Cache"),
		},
		{
			Type: pulumi.String("azure-native:storagecache/v20200301:Cache"),
		},
		{
			Type: pulumi.String("azure-native:storagecache/v20201001:Cache"),
		},
		{
			Type: pulumi.String("azure-native:storagecache/v20210301:Cache"),
		},
		{
			Type: pulumi.String("azure-native:storagecache/v20210501:Cache"),
		},
		{
			Type: pulumi.String("azure-native:storagecache/v20210901:Cache"),
		},
	})
	opts = append(opts, aliases)
	var resource Cache
	err := ctx.RegisterResource("azure-native:storagecache/v20191101:Cache", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetCache(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CacheState, opts ...pulumi.ResourceOption) (*Cache, error) {
	var resource Cache
	err := ctx.ReadResource("azure-native:storagecache/v20191101:Cache", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type cacheState struct {
}

type CacheState struct {
}

func (CacheState) ElementType() reflect.Type {
	return reflect.TypeOf((*cacheState)(nil)).Elem()
}

type cacheArgs struct {
	CacheName         *string     `pulumi:"cacheName"`
	CacheSizeGB       *int        `pulumi:"cacheSizeGB"`
	Location          *string     `pulumi:"location"`
	ProvisioningState *string     `pulumi:"provisioningState"`
	ResourceGroupName string      `pulumi:"resourceGroupName"`
	Sku               *CacheSku   `pulumi:"sku"`
	Subnet            *string     `pulumi:"subnet"`
	Tags              interface{} `pulumi:"tags"`
}


type CacheArgs struct {
	CacheName         pulumi.StringPtrInput
	CacheSizeGB       pulumi.IntPtrInput
	Location          pulumi.StringPtrInput
	ProvisioningState pulumi.StringPtrInput
	ResourceGroupName pulumi.StringInput
	Sku               CacheSkuPtrInput
	Subnet            pulumi.StringPtrInput
	Tags              pulumi.Input
}

func (CacheArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*cacheArgs)(nil)).Elem()
}

type CacheInput interface {
	pulumi.Input

	ToCacheOutput() CacheOutput
	ToCacheOutputWithContext(ctx context.Context) CacheOutput
}

func (*Cache) ElementType() reflect.Type {
	return reflect.TypeOf((**Cache)(nil)).Elem()
}

func (i *Cache) ToCacheOutput() CacheOutput {
	return i.ToCacheOutputWithContext(context.Background())
}

func (i *Cache) ToCacheOutputWithContext(ctx context.Context) CacheOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CacheOutput)
}

type CacheOutput struct{ *pulumi.OutputState }

func (CacheOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Cache)(nil)).Elem()
}

func (o CacheOutput) ToCacheOutput() CacheOutput {
	return o
}

func (o CacheOutput) ToCacheOutputWithContext(ctx context.Context) CacheOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(CacheOutput{})
}
