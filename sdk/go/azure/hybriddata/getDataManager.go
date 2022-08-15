


package hybriddata

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupDataManager(ctx *pulumi.Context, args *LookupDataManagerArgs, opts ...pulumi.InvokeOption) (*LookupDataManagerResult, error) {
	var rv LookupDataManagerResult
	err := ctx.Invoke("azure-native:hybriddata:getDataManager", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDataManagerArgs struct {
	DataManagerName   string `pulumi:"dataManagerName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupDataManagerResult struct {
	Etag     *string           `pulumi:"etag"`
	Id       string            `pulumi:"id"`
	Location string            `pulumi:"location"`
	Name     string            `pulumi:"name"`
	Sku      *SkuResponse      `pulumi:"sku"`
	Tags     map[string]string `pulumi:"tags"`
	Type     string            `pulumi:"type"`
}

func LookupDataManagerOutput(ctx *pulumi.Context, args LookupDataManagerOutputArgs, opts ...pulumi.InvokeOption) LookupDataManagerResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDataManagerResult, error) {
			args := v.(LookupDataManagerArgs)
			r, err := LookupDataManager(ctx, &args, opts...)
			var s LookupDataManagerResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupDataManagerResultOutput)
}

type LookupDataManagerOutputArgs struct {
	DataManagerName   pulumi.StringInput `pulumi:"dataManagerName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupDataManagerOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDataManagerArgs)(nil)).Elem()
}


type LookupDataManagerResultOutput struct{ *pulumi.OutputState }

func (LookupDataManagerResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDataManagerResult)(nil)).Elem()
}

func (o LookupDataManagerResultOutput) ToLookupDataManagerResultOutput() LookupDataManagerResultOutput {
	return o
}

func (o LookupDataManagerResultOutput) ToLookupDataManagerResultOutputWithContext(ctx context.Context) LookupDataManagerResultOutput {
	return o
}

func (o LookupDataManagerResultOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDataManagerResult) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o LookupDataManagerResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDataManagerResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupDataManagerResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDataManagerResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupDataManagerResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDataManagerResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupDataManagerResultOutput) Sku() SkuResponsePtrOutput {
	return o.ApplyT(func(v LookupDataManagerResult) *SkuResponse { return v.Sku }).(SkuResponsePtrOutput)
}

func (o LookupDataManagerResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupDataManagerResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupDataManagerResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDataManagerResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDataManagerResultOutput{})
}
