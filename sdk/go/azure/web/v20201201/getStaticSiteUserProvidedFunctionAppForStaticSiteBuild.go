


package v20201201

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupStaticSiteUserProvidedFunctionAppForStaticSiteBuild(ctx *pulumi.Context, args *LookupStaticSiteUserProvidedFunctionAppForStaticSiteBuildArgs, opts ...pulumi.InvokeOption) (*LookupStaticSiteUserProvidedFunctionAppForStaticSiteBuildResult, error) {
	var rv LookupStaticSiteUserProvidedFunctionAppForStaticSiteBuildResult
	err := ctx.Invoke("azure-native:web/v20201201:getStaticSiteUserProvidedFunctionAppForStaticSiteBuild", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupStaticSiteUserProvidedFunctionAppForStaticSiteBuildArgs struct {
	EnvironmentName   string `pulumi:"environmentName"`
	FunctionAppName   string `pulumi:"functionAppName"`
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupStaticSiteUserProvidedFunctionAppForStaticSiteBuildResult struct {
	CreatedOn             string  `pulumi:"createdOn"`
	FunctionAppRegion     *string `pulumi:"functionAppRegion"`
	FunctionAppResourceId *string `pulumi:"functionAppResourceId"`
	Id                    string  `pulumi:"id"`
	Kind                  *string `pulumi:"kind"`
	Name                  string  `pulumi:"name"`
	Type                  string  `pulumi:"type"`
}

func LookupStaticSiteUserProvidedFunctionAppForStaticSiteBuildOutput(ctx *pulumi.Context, args LookupStaticSiteUserProvidedFunctionAppForStaticSiteBuildOutputArgs, opts ...pulumi.InvokeOption) LookupStaticSiteUserProvidedFunctionAppForStaticSiteBuildResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupStaticSiteUserProvidedFunctionAppForStaticSiteBuildResult, error) {
			args := v.(LookupStaticSiteUserProvidedFunctionAppForStaticSiteBuildArgs)
			r, err := LookupStaticSiteUserProvidedFunctionAppForStaticSiteBuild(ctx, &args, opts...)
			var s LookupStaticSiteUserProvidedFunctionAppForStaticSiteBuildResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupStaticSiteUserProvidedFunctionAppForStaticSiteBuildResultOutput)
}

type LookupStaticSiteUserProvidedFunctionAppForStaticSiteBuildOutputArgs struct {
	EnvironmentName   pulumi.StringInput `pulumi:"environmentName"`
	FunctionAppName   pulumi.StringInput `pulumi:"functionAppName"`
	Name              pulumi.StringInput `pulumi:"name"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupStaticSiteUserProvidedFunctionAppForStaticSiteBuildOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupStaticSiteUserProvidedFunctionAppForStaticSiteBuildArgs)(nil)).Elem()
}


type LookupStaticSiteUserProvidedFunctionAppForStaticSiteBuildResultOutput struct{ *pulumi.OutputState }

func (LookupStaticSiteUserProvidedFunctionAppForStaticSiteBuildResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupStaticSiteUserProvidedFunctionAppForStaticSiteBuildResult)(nil)).Elem()
}

func (o LookupStaticSiteUserProvidedFunctionAppForStaticSiteBuildResultOutput) ToLookupStaticSiteUserProvidedFunctionAppForStaticSiteBuildResultOutput() LookupStaticSiteUserProvidedFunctionAppForStaticSiteBuildResultOutput {
	return o
}

func (o LookupStaticSiteUserProvidedFunctionAppForStaticSiteBuildResultOutput) ToLookupStaticSiteUserProvidedFunctionAppForStaticSiteBuildResultOutputWithContext(ctx context.Context) LookupStaticSiteUserProvidedFunctionAppForStaticSiteBuildResultOutput {
	return o
}

func (o LookupStaticSiteUserProvidedFunctionAppForStaticSiteBuildResultOutput) CreatedOn() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStaticSiteUserProvidedFunctionAppForStaticSiteBuildResult) string { return v.CreatedOn }).(pulumi.StringOutput)
}

func (o LookupStaticSiteUserProvidedFunctionAppForStaticSiteBuildResultOutput) FunctionAppRegion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupStaticSiteUserProvidedFunctionAppForStaticSiteBuildResult) *string {
		return v.FunctionAppRegion
	}).(pulumi.StringPtrOutput)
}

func (o LookupStaticSiteUserProvidedFunctionAppForStaticSiteBuildResultOutput) FunctionAppResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupStaticSiteUserProvidedFunctionAppForStaticSiteBuildResult) *string {
		return v.FunctionAppResourceId
	}).(pulumi.StringPtrOutput)
}

func (o LookupStaticSiteUserProvidedFunctionAppForStaticSiteBuildResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStaticSiteUserProvidedFunctionAppForStaticSiteBuildResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupStaticSiteUserProvidedFunctionAppForStaticSiteBuildResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupStaticSiteUserProvidedFunctionAppForStaticSiteBuildResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o LookupStaticSiteUserProvidedFunctionAppForStaticSiteBuildResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStaticSiteUserProvidedFunctionAppForStaticSiteBuildResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupStaticSiteUserProvidedFunctionAppForStaticSiteBuildResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStaticSiteUserProvidedFunctionAppForStaticSiteBuildResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupStaticSiteUserProvidedFunctionAppForStaticSiteBuildResultOutput{})
}
