


package v20181203

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupEnterpriseKnowledgeGraph(ctx *pulumi.Context, args *LookupEnterpriseKnowledgeGraphArgs, opts ...pulumi.InvokeOption) (*LookupEnterpriseKnowledgeGraphResult, error) {
	var rv LookupEnterpriseKnowledgeGraphResult
	err := ctx.Invoke("azure-native:enterpriseknowledgegraph/v20181203:getEnterpriseKnowledgeGraph", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupEnterpriseKnowledgeGraphArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ResourceName      string `pulumi:"resourceName"`
}


type LookupEnterpriseKnowledgeGraphResult struct {
	Id         string                                     `pulumi:"id"`
	Location   *string                                    `pulumi:"location"`
	Name       string                                     `pulumi:"name"`
	Properties EnterpriseKnowledgeGraphPropertiesResponse `pulumi:"properties"`
	Sku        *SkuResponse                               `pulumi:"sku"`
	Tags       map[string]string                          `pulumi:"tags"`
	Type       string                                     `pulumi:"type"`
}

func LookupEnterpriseKnowledgeGraphOutput(ctx *pulumi.Context, args LookupEnterpriseKnowledgeGraphOutputArgs, opts ...pulumi.InvokeOption) LookupEnterpriseKnowledgeGraphResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupEnterpriseKnowledgeGraphResult, error) {
			args := v.(LookupEnterpriseKnowledgeGraphArgs)
			r, err := LookupEnterpriseKnowledgeGraph(ctx, &args, opts...)
			var s LookupEnterpriseKnowledgeGraphResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupEnterpriseKnowledgeGraphResultOutput)
}

type LookupEnterpriseKnowledgeGraphOutputArgs struct {
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	ResourceName      pulumi.StringInput `pulumi:"resourceName"`
}

func (LookupEnterpriseKnowledgeGraphOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupEnterpriseKnowledgeGraphArgs)(nil)).Elem()
}


type LookupEnterpriseKnowledgeGraphResultOutput struct{ *pulumi.OutputState }

func (LookupEnterpriseKnowledgeGraphResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupEnterpriseKnowledgeGraphResult)(nil)).Elem()
}

func (o LookupEnterpriseKnowledgeGraphResultOutput) ToLookupEnterpriseKnowledgeGraphResultOutput() LookupEnterpriseKnowledgeGraphResultOutput {
	return o
}

func (o LookupEnterpriseKnowledgeGraphResultOutput) ToLookupEnterpriseKnowledgeGraphResultOutputWithContext(ctx context.Context) LookupEnterpriseKnowledgeGraphResultOutput {
	return o
}

func (o LookupEnterpriseKnowledgeGraphResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEnterpriseKnowledgeGraphResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupEnterpriseKnowledgeGraphResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupEnterpriseKnowledgeGraphResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupEnterpriseKnowledgeGraphResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEnterpriseKnowledgeGraphResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupEnterpriseKnowledgeGraphResultOutput) Properties() EnterpriseKnowledgeGraphPropertiesResponseOutput {
	return o.ApplyT(func(v LookupEnterpriseKnowledgeGraphResult) EnterpriseKnowledgeGraphPropertiesResponse {
		return v.Properties
	}).(EnterpriseKnowledgeGraphPropertiesResponseOutput)
}

func (o LookupEnterpriseKnowledgeGraphResultOutput) Sku() SkuResponsePtrOutput {
	return o.ApplyT(func(v LookupEnterpriseKnowledgeGraphResult) *SkuResponse { return v.Sku }).(SkuResponsePtrOutput)
}

func (o LookupEnterpriseKnowledgeGraphResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupEnterpriseKnowledgeGraphResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupEnterpriseKnowledgeGraphResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEnterpriseKnowledgeGraphResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupEnterpriseKnowledgeGraphResultOutput{})
}
