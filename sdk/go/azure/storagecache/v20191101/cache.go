


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
		{
			Type: pulumi.String("azure-native:storagecache/v20220101:Cache"),
		},
		{
			Type: pulumi.String("azure-native:storagecache/v20220501:Cache"),
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

func (o CacheOutput) CacheSizeGB() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Cache) pulumi.IntPtrOutput { return v.CacheSizeGB }).(pulumi.IntPtrOutput)
}

func (o CacheOutput) Health() CacheHealthResponseOutput {
	return o.ApplyT(func(v *Cache) CacheHealthResponseOutput { return v.Health }).(CacheHealthResponseOutput)
}

func (o CacheOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Cache) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

func (o CacheOutput) MountAddresses() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Cache) pulumi.StringArrayOutput { return v.MountAddresses }).(pulumi.StringArrayOutput)
}

func (o CacheOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Cache) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o CacheOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Cache) pulumi.StringPtrOutput { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

func (o CacheOutput) Sku() CacheResponseSkuPtrOutput {
	return o.ApplyT(func(v *Cache) CacheResponseSkuPtrOutput { return v.Sku }).(CacheResponseSkuPtrOutput)
}

func (o CacheOutput) Subnet() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Cache) pulumi.StringPtrOutput { return v.Subnet }).(pulumi.StringPtrOutput)
}

func (o CacheOutput) Tags() pulumi.AnyOutput {
	return o.ApplyT(func(v *Cache) pulumi.AnyOutput { return v.Tags }).(pulumi.AnyOutput)
}

func (o CacheOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Cache) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o CacheOutput) UpgradeStatus() CacheUpgradeStatusResponsePtrOutput {
	return o.ApplyT(func(v *Cache) CacheUpgradeStatusResponsePtrOutput { return v.UpgradeStatus }).(CacheUpgradeStatusResponsePtrOutput)
}

func init() {
	pulumi.RegisterOutputType(CacheOutput{})
}
