


package v20210801

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupCache(ctx *pulumi.Context, args *LookupCacheArgs, opts ...pulumi.InvokeOption) (*LookupCacheResult, error) {
	var rv LookupCacheResult
	err := ctx.Invoke("azure-native:apimanagement/v20210801:getCache", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupCacheArgs struct {
	CacheId           string `pulumi:"cacheId"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ServiceName       string `pulumi:"serviceName"`
}


type LookupCacheResult struct {
	ConnectionString string  `pulumi:"connectionString"`
	Description      *string `pulumi:"description"`
	Id               string  `pulumi:"id"`
	Name             string  `pulumi:"name"`
	ResourceId       *string `pulumi:"resourceId"`
	Type             string  `pulumi:"type"`
	UseFromLocation  string  `pulumi:"useFromLocation"`
}

func LookupCacheOutput(ctx *pulumi.Context, args LookupCacheOutputArgs, opts ...pulumi.InvokeOption) LookupCacheResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupCacheResult, error) {
			args := v.(LookupCacheArgs)
			r, err := LookupCache(ctx, &args, opts...)
			var s LookupCacheResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupCacheResultOutput)
}

type LookupCacheOutputArgs struct {
	CacheId           pulumi.StringInput `pulumi:"cacheId"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	ServiceName       pulumi.StringInput `pulumi:"serviceName"`
}

func (LookupCacheOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCacheArgs)(nil)).Elem()
}


type LookupCacheResultOutput struct{ *pulumi.OutputState }

func (LookupCacheResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCacheResult)(nil)).Elem()
}

func (o LookupCacheResultOutput) ToLookupCacheResultOutput() LookupCacheResultOutput {
	return o
}

func (o LookupCacheResultOutput) ToLookupCacheResultOutputWithContext(ctx context.Context) LookupCacheResultOutput {
	return o
}

func (o LookupCacheResultOutput) ConnectionString() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCacheResult) string { return v.ConnectionString }).(pulumi.StringOutput)
}

func (o LookupCacheResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCacheResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupCacheResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCacheResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupCacheResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCacheResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupCacheResultOutput) ResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCacheResult) *string { return v.ResourceId }).(pulumi.StringPtrOutput)
}

func (o LookupCacheResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCacheResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupCacheResultOutput) UseFromLocation() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCacheResult) string { return v.UseFromLocation }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupCacheResultOutput{})
}
