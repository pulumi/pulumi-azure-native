


package v20170501

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


func LookupApplication(ctx *pulumi.Context, args *LookupApplicationArgs, opts ...pulumi.InvokeOption) (*LookupApplicationResult, error) {
	var rv LookupApplicationResult
	err := ctx.Invoke("azure-native:batch/v20170501:getApplication", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupApplicationArgs struct {
	AccountName       string `pulumi:"accountName"`
	ApplicationId     string `pulumi:"applicationId"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupApplicationResult struct {
	AllowUpdates   *bool                        `pulumi:"allowUpdates"`
	DefaultVersion *string                      `pulumi:"defaultVersion"`
	DisplayName    *string                      `pulumi:"displayName"`
	Id             *string                      `pulumi:"id"`
	Packages       []ApplicationPackageResponse `pulumi:"packages"`
}

func LookupApplicationOutput(ctx *pulumi.Context, args LookupApplicationOutputArgs, opts ...pulumi.InvokeOption) LookupApplicationResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupApplicationResult, error) {
			args := v.(LookupApplicationArgs)
			r, err := LookupApplication(ctx, &args, opts...)
			var s LookupApplicationResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupApplicationResultOutput)
}

type LookupApplicationOutputArgs struct {
	AccountName       pulumi.StringInput `pulumi:"accountName"`
	ApplicationId     pulumi.StringInput `pulumi:"applicationId"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupApplicationOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupApplicationArgs)(nil)).Elem()
}


type LookupApplicationResultOutput struct{ *pulumi.OutputState }

func (LookupApplicationResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupApplicationResult)(nil)).Elem()
}

func (o LookupApplicationResultOutput) ToLookupApplicationResultOutput() LookupApplicationResultOutput {
	return o
}

func (o LookupApplicationResultOutput) ToLookupApplicationResultOutputWithContext(ctx context.Context) LookupApplicationResultOutput {
	return o
}

func (o LookupApplicationResultOutput) AllowUpdates() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupApplicationResult) *bool { return v.AllowUpdates }).(pulumi.BoolPtrOutput)
}

func (o LookupApplicationResultOutput) DefaultVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupApplicationResult) *string { return v.DefaultVersion }).(pulumi.StringPtrOutput)
}

func (o LookupApplicationResultOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupApplicationResult) *string { return v.DisplayName }).(pulumi.StringPtrOutput)
}

func (o LookupApplicationResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupApplicationResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupApplicationResultOutput) Packages() ApplicationPackageResponseArrayOutput {
	return o.ApplyT(func(v LookupApplicationResult) []ApplicationPackageResponse { return v.Packages }).(ApplicationPackageResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupApplicationResultOutput{})
}
