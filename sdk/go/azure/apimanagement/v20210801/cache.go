


package v20210801

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Cache struct {
	pulumi.CustomResourceState

	ConnectionString pulumi.StringOutput    `pulumi:"connectionString"`
	Description      pulumi.StringPtrOutput `pulumi:"description"`
	Name             pulumi.StringOutput    `pulumi:"name"`
	ResourceId       pulumi.StringPtrOutput `pulumi:"resourceId"`
	Type             pulumi.StringOutput    `pulumi:"type"`
	UseFromLocation  pulumi.StringOutput    `pulumi:"useFromLocation"`
}


func NewCache(ctx *pulumi.Context,
	name string, args *CacheArgs, opts ...pulumi.ResourceOption) (*Cache, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ConnectionString == nil {
		return nil, errors.New("invalid value for required argument 'ConnectionString'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ServiceName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceName'")
	}
	if args.UseFromLocation == nil {
		return nil, errors.New("invalid value for required argument 'UseFromLocation'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:apimanagement:Cache"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20180601preview:Cache"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20190101:Cache"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20191201:Cache"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20191201preview:Cache"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20200601preview:Cache"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20201201:Cache"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20210101preview:Cache"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20210401preview:Cache"),
		},
	})
	opts = append(opts, aliases)
	var resource Cache
	err := ctx.RegisterResource("azure-native:apimanagement/v20210801:Cache", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetCache(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CacheState, opts ...pulumi.ResourceOption) (*Cache, error) {
	var resource Cache
	err := ctx.ReadResource("azure-native:apimanagement/v20210801:Cache", name, id, state, &resource, opts...)
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
	CacheId           *string `pulumi:"cacheId"`
	ConnectionString  string  `pulumi:"connectionString"`
	Description       *string `pulumi:"description"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
	ResourceId        *string `pulumi:"resourceId"`
	ServiceName       string  `pulumi:"serviceName"`
	UseFromLocation   string  `pulumi:"useFromLocation"`
}


type CacheArgs struct {
	CacheId           pulumi.StringPtrInput
	ConnectionString  pulumi.StringInput
	Description       pulumi.StringPtrInput
	ResourceGroupName pulumi.StringInput
	ResourceId        pulumi.StringPtrInput
	ServiceName       pulumi.StringInput
	UseFromLocation   pulumi.StringInput
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
	return reflect.TypeOf((*Cache)(nil))
}

func (i *Cache) ToCacheOutput() CacheOutput {
	return i.ToCacheOutputWithContext(context.Background())
}

func (i *Cache) ToCacheOutputWithContext(ctx context.Context) CacheOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CacheOutput)
}

type CacheOutput struct{ *pulumi.OutputState }

func (CacheOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Cache)(nil))
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
