


package v20210101

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupStaticSiteUserProvidedFunctionAppForStaticSite(ctx *pulumi.Context, args *LookupStaticSiteUserProvidedFunctionAppForStaticSiteArgs, opts ...pulumi.InvokeOption) (*LookupStaticSiteUserProvidedFunctionAppForStaticSiteResult, error) {
	var rv LookupStaticSiteUserProvidedFunctionAppForStaticSiteResult
	err := ctx.Invoke("azure-native:web/v20210101:getStaticSiteUserProvidedFunctionAppForStaticSite", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupStaticSiteUserProvidedFunctionAppForStaticSiteArgs struct {
	FunctionAppName   string `pulumi:"functionAppName"`
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupStaticSiteUserProvidedFunctionAppForStaticSiteResult struct {
	CreatedOn             string  `pulumi:"createdOn"`
	FunctionAppRegion     *string `pulumi:"functionAppRegion"`
	FunctionAppResourceId *string `pulumi:"functionAppResourceId"`
	Id                    string  `pulumi:"id"`
	Kind                  *string `pulumi:"kind"`
	Name                  string  `pulumi:"name"`
	Type                  string  `pulumi:"type"`
}

func LookupStaticSiteUserProvidedFunctionAppForStaticSiteOutput(ctx *pulumi.Context, args LookupStaticSiteUserProvidedFunctionAppForStaticSiteOutputArgs, opts ...pulumi.InvokeOption) LookupStaticSiteUserProvidedFunctionAppForStaticSiteResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupStaticSiteUserProvidedFunctionAppForStaticSiteResult, error) {
			args := v.(LookupStaticSiteUserProvidedFunctionAppForStaticSiteArgs)
			r, err := LookupStaticSiteUserProvidedFunctionAppForStaticSite(ctx, &args, opts...)
			var s LookupStaticSiteUserProvidedFunctionAppForStaticSiteResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupStaticSiteUserProvidedFunctionAppForStaticSiteResultOutput)
}

type LookupStaticSiteUserProvidedFunctionAppForStaticSiteOutputArgs struct {
	FunctionAppName   pulumi.StringInput `pulumi:"functionAppName"`
	Name              pulumi.StringInput `pulumi:"name"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupStaticSiteUserProvidedFunctionAppForStaticSiteOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupStaticSiteUserProvidedFunctionAppForStaticSiteArgs)(nil)).Elem()
}


type LookupStaticSiteUserProvidedFunctionAppForStaticSiteResultOutput struct{ *pulumi.OutputState }

func (LookupStaticSiteUserProvidedFunctionAppForStaticSiteResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupStaticSiteUserProvidedFunctionAppForStaticSiteResult)(nil)).Elem()
}

func (o LookupStaticSiteUserProvidedFunctionAppForStaticSiteResultOutput) ToLookupStaticSiteUserProvidedFunctionAppForStaticSiteResultOutput() LookupStaticSiteUserProvidedFunctionAppForStaticSiteResultOutput {
	return o
}

func (o LookupStaticSiteUserProvidedFunctionAppForStaticSiteResultOutput) ToLookupStaticSiteUserProvidedFunctionAppForStaticSiteResultOutputWithContext(ctx context.Context) LookupStaticSiteUserProvidedFunctionAppForStaticSiteResultOutput {
	return o
}

func (o LookupStaticSiteUserProvidedFunctionAppForStaticSiteResultOutput) CreatedOn() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStaticSiteUserProvidedFunctionAppForStaticSiteResult) string { return v.CreatedOn }).(pulumi.StringOutput)
}

func (o LookupStaticSiteUserProvidedFunctionAppForStaticSiteResultOutput) FunctionAppRegion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupStaticSiteUserProvidedFunctionAppForStaticSiteResult) *string { return v.FunctionAppRegion }).(pulumi.StringPtrOutput)
}

func (o LookupStaticSiteUserProvidedFunctionAppForStaticSiteResultOutput) FunctionAppResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupStaticSiteUserProvidedFunctionAppForStaticSiteResult) *string {
		return v.FunctionAppResourceId
	}).(pulumi.StringPtrOutput)
}

func (o LookupStaticSiteUserProvidedFunctionAppForStaticSiteResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStaticSiteUserProvidedFunctionAppForStaticSiteResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupStaticSiteUserProvidedFunctionAppForStaticSiteResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupStaticSiteUserProvidedFunctionAppForStaticSiteResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o LookupStaticSiteUserProvidedFunctionAppForStaticSiteResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStaticSiteUserProvidedFunctionAppForStaticSiteResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupStaticSiteUserProvidedFunctionAppForStaticSiteResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStaticSiteUserProvidedFunctionAppForStaticSiteResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupStaticSiteUserProvidedFunctionAppForStaticSiteResultOutput{})
}
