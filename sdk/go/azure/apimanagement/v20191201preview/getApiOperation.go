


package v20191201preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupApiOperation(ctx *pulumi.Context, args *LookupApiOperationArgs, opts ...pulumi.InvokeOption) (*LookupApiOperationResult, error) {
	var rv LookupApiOperationResult
	err := ctx.Invoke("azure-native:apimanagement/v20191201preview:getApiOperation", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupApiOperationArgs struct {
	ApiId             string `pulumi:"apiId"`
	OperationId       string `pulumi:"operationId"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ServiceName       string `pulumi:"serviceName"`
}


type LookupApiOperationResult struct {
	Description        *string                     `pulumi:"description"`
	DisplayName        string                      `pulumi:"displayName"`
	Id                 string                      `pulumi:"id"`
	Method             string                      `pulumi:"method"`
	Name               string                      `pulumi:"name"`
	Policies           *string                     `pulumi:"policies"`
	Request            *RequestContractResponse    `pulumi:"request"`
	Responses          []ResponseContractResponse  `pulumi:"responses"`
	TemplateParameters []ParameterContractResponse `pulumi:"templateParameters"`
	Type               string                      `pulumi:"type"`
	UrlTemplate        string                      `pulumi:"urlTemplate"`
}

func LookupApiOperationOutput(ctx *pulumi.Context, args LookupApiOperationOutputArgs, opts ...pulumi.InvokeOption) LookupApiOperationResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupApiOperationResult, error) {
			args := v.(LookupApiOperationArgs)
			r, err := LookupApiOperation(ctx, &args, opts...)
			var s LookupApiOperationResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupApiOperationResultOutput)
}

type LookupApiOperationOutputArgs struct {
	ApiId             pulumi.StringInput `pulumi:"apiId"`
	OperationId       pulumi.StringInput `pulumi:"operationId"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	ServiceName       pulumi.StringInput `pulumi:"serviceName"`
}

func (LookupApiOperationOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupApiOperationArgs)(nil)).Elem()
}


type LookupApiOperationResultOutput struct{ *pulumi.OutputState }

func (LookupApiOperationResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupApiOperationResult)(nil)).Elem()
}

func (o LookupApiOperationResultOutput) ToLookupApiOperationResultOutput() LookupApiOperationResultOutput {
	return o
}

func (o LookupApiOperationResultOutput) ToLookupApiOperationResultOutputWithContext(ctx context.Context) LookupApiOperationResultOutput {
	return o
}

func (o LookupApiOperationResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupApiOperationResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupApiOperationResultOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApiOperationResult) string { return v.DisplayName }).(pulumi.StringOutput)
}

func (o LookupApiOperationResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApiOperationResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupApiOperationResultOutput) Method() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApiOperationResult) string { return v.Method }).(pulumi.StringOutput)
}

func (o LookupApiOperationResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApiOperationResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupApiOperationResultOutput) Policies() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupApiOperationResult) *string { return v.Policies }).(pulumi.StringPtrOutput)
}

func (o LookupApiOperationResultOutput) Request() RequestContractResponsePtrOutput {
	return o.ApplyT(func(v LookupApiOperationResult) *RequestContractResponse { return v.Request }).(RequestContractResponsePtrOutput)
}

func (o LookupApiOperationResultOutput) Responses() ResponseContractResponseArrayOutput {
	return o.ApplyT(func(v LookupApiOperationResult) []ResponseContractResponse { return v.Responses }).(ResponseContractResponseArrayOutput)
}

func (o LookupApiOperationResultOutput) TemplateParameters() ParameterContractResponseArrayOutput {
	return o.ApplyT(func(v LookupApiOperationResult) []ParameterContractResponse { return v.TemplateParameters }).(ParameterContractResponseArrayOutput)
}

func (o LookupApiOperationResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApiOperationResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupApiOperationResultOutput) UrlTemplate() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApiOperationResult) string { return v.UrlTemplate }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupApiOperationResultOutput{})
}
