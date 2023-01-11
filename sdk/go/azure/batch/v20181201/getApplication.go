


package v20181201

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


func LookupApplication(ctx *pulumi.Context, args *LookupApplicationArgs, opts ...pulumi.InvokeOption) (*LookupApplicationResult, error) {
	var rv LookupApplicationResult
	err := ctx.Invoke("azure-native:batch/v20181201:getApplication", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupApplicationArgs struct {
	AccountName       string `pulumi:"accountName"`
	ApplicationName   string `pulumi:"applicationName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupApplicationResult struct {
	AllowUpdates   *bool   `pulumi:"allowUpdates"`
	DefaultVersion *string `pulumi:"defaultVersion"`
	DisplayName    *string `pulumi:"displayName"`
	Etag           string  `pulumi:"etag"`
	Id             string  `pulumi:"id"`
	Name           string  `pulumi:"name"`
	Type           string  `pulumi:"type"`
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
	ApplicationName   pulumi.StringInput `pulumi:"applicationName"`
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

func (o LookupApplicationResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApplicationResult) string { return v.Etag }).(pulumi.StringOutput)
}

func (o LookupApplicationResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApplicationResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupApplicationResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApplicationResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupApplicationResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApplicationResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupApplicationResultOutput{})
}
