


package v20191212

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


func LookupGremlinResourceGremlinGraph(ctx *pulumi.Context, args *LookupGremlinResourceGremlinGraphArgs, opts ...pulumi.InvokeOption) (*LookupGremlinResourceGremlinGraphResult, error) {
	var rv LookupGremlinResourceGremlinGraphResult
	err := ctx.Invoke("azure-native:documentdb/v20191212:getGremlinResourceGremlinGraph", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupGremlinResourceGremlinGraphArgs struct {
	AccountName       string `pulumi:"accountName"`
	DatabaseName      string `pulumi:"databaseName"`
	GraphName         string `pulumi:"graphName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupGremlinResourceGremlinGraphResult struct {
	Id       string                                     `pulumi:"id"`
	Location *string                                    `pulumi:"location"`
	Name     string                                     `pulumi:"name"`
	Resource *GremlinGraphGetPropertiesResponseResource `pulumi:"resource"`
	Tags     map[string]string                          `pulumi:"tags"`
	Type     string                                     `pulumi:"type"`
}


func (val *LookupGremlinResourceGremlinGraphResult) Defaults() *LookupGremlinResourceGremlinGraphResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.Resource = tmp.Resource.Defaults()

	return &tmp
}

func LookupGremlinResourceGremlinGraphOutput(ctx *pulumi.Context, args LookupGremlinResourceGremlinGraphOutputArgs, opts ...pulumi.InvokeOption) LookupGremlinResourceGremlinGraphResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupGremlinResourceGremlinGraphResult, error) {
			args := v.(LookupGremlinResourceGremlinGraphArgs)
			r, err := LookupGremlinResourceGremlinGraph(ctx, &args, opts...)
			var s LookupGremlinResourceGremlinGraphResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupGremlinResourceGremlinGraphResultOutput)
}

type LookupGremlinResourceGremlinGraphOutputArgs struct {
	AccountName       pulumi.StringInput `pulumi:"accountName"`
	DatabaseName      pulumi.StringInput `pulumi:"databaseName"`
	GraphName         pulumi.StringInput `pulumi:"graphName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupGremlinResourceGremlinGraphOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupGremlinResourceGremlinGraphArgs)(nil)).Elem()
}


type LookupGremlinResourceGremlinGraphResultOutput struct{ *pulumi.OutputState }

func (LookupGremlinResourceGremlinGraphResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupGremlinResourceGremlinGraphResult)(nil)).Elem()
}

func (o LookupGremlinResourceGremlinGraphResultOutput) ToLookupGremlinResourceGremlinGraphResultOutput() LookupGremlinResourceGremlinGraphResultOutput {
	return o
}

func (o LookupGremlinResourceGremlinGraphResultOutput) ToLookupGremlinResourceGremlinGraphResultOutputWithContext(ctx context.Context) LookupGremlinResourceGremlinGraphResultOutput {
	return o
}

func (o LookupGremlinResourceGremlinGraphResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGremlinResourceGremlinGraphResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupGremlinResourceGremlinGraphResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupGremlinResourceGremlinGraphResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupGremlinResourceGremlinGraphResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGremlinResourceGremlinGraphResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupGremlinResourceGremlinGraphResultOutput) Resource() GremlinGraphGetPropertiesResponseResourcePtrOutput {
	return o.ApplyT(func(v LookupGremlinResourceGremlinGraphResult) *GremlinGraphGetPropertiesResponseResource {
		return v.Resource
	}).(GremlinGraphGetPropertiesResponseResourcePtrOutput)
}

func (o LookupGremlinResourceGremlinGraphResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupGremlinResourceGremlinGraphResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupGremlinResourceGremlinGraphResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGremlinResourceGremlinGraphResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupGremlinResourceGremlinGraphResultOutput{})
}
