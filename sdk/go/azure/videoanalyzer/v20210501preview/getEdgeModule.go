


package v20210501preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupEdgeModule(ctx *pulumi.Context, args *LookupEdgeModuleArgs, opts ...pulumi.InvokeOption) (*LookupEdgeModuleResult, error) {
	var rv LookupEdgeModuleResult
	err := ctx.Invoke("azure-native:videoanalyzer/v20210501preview:getEdgeModule", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupEdgeModuleArgs struct {
	AccountName       string `pulumi:"accountName"`
	EdgeModuleName    string `pulumi:"edgeModuleName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupEdgeModuleResult struct {
	EdgeModuleId string             `pulumi:"edgeModuleId"`
	Id           string             `pulumi:"id"`
	Name         string             `pulumi:"name"`
	SystemData   SystemDataResponse `pulumi:"systemData"`
	Type         string             `pulumi:"type"`
}

func LookupEdgeModuleOutput(ctx *pulumi.Context, args LookupEdgeModuleOutputArgs, opts ...pulumi.InvokeOption) LookupEdgeModuleResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupEdgeModuleResult, error) {
			args := v.(LookupEdgeModuleArgs)
			r, err := LookupEdgeModule(ctx, &args, opts...)
			var s LookupEdgeModuleResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupEdgeModuleResultOutput)
}

type LookupEdgeModuleOutputArgs struct {
	AccountName       pulumi.StringInput `pulumi:"accountName"`
	EdgeModuleName    pulumi.StringInput `pulumi:"edgeModuleName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupEdgeModuleOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupEdgeModuleArgs)(nil)).Elem()
}


type LookupEdgeModuleResultOutput struct{ *pulumi.OutputState }

func (LookupEdgeModuleResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupEdgeModuleResult)(nil)).Elem()
}

func (o LookupEdgeModuleResultOutput) ToLookupEdgeModuleResultOutput() LookupEdgeModuleResultOutput {
	return o
}

func (o LookupEdgeModuleResultOutput) ToLookupEdgeModuleResultOutputWithContext(ctx context.Context) LookupEdgeModuleResultOutput {
	return o
}

func (o LookupEdgeModuleResultOutput) EdgeModuleId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEdgeModuleResult) string { return v.EdgeModuleId }).(pulumi.StringOutput)
}

func (o LookupEdgeModuleResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEdgeModuleResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupEdgeModuleResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEdgeModuleResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupEdgeModuleResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupEdgeModuleResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupEdgeModuleResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEdgeModuleResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupEdgeModuleResultOutput{})
}
