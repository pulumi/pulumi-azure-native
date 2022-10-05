


package v20180101

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupApiVersionSet(ctx *pulumi.Context, args *LookupApiVersionSetArgs, opts ...pulumi.InvokeOption) (*LookupApiVersionSetResult, error) {
	var rv LookupApiVersionSetResult
	err := ctx.Invoke("azure-native:apimanagement/v20180101:getApiVersionSet", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupApiVersionSetArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ServiceName       string `pulumi:"serviceName"`
	VersionSetId      string `pulumi:"versionSetId"`
}


type LookupApiVersionSetResult struct {
	Description       *string `pulumi:"description"`
	DisplayName       string  `pulumi:"displayName"`
	Id                string  `pulumi:"id"`
	Name              string  `pulumi:"name"`
	Type              string  `pulumi:"type"`
	VersionHeaderName *string `pulumi:"versionHeaderName"`
	VersionQueryName  *string `pulumi:"versionQueryName"`
	VersioningScheme  string  `pulumi:"versioningScheme"`
}

func LookupApiVersionSetOutput(ctx *pulumi.Context, args LookupApiVersionSetOutputArgs, opts ...pulumi.InvokeOption) LookupApiVersionSetResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupApiVersionSetResult, error) {
			args := v.(LookupApiVersionSetArgs)
			r, err := LookupApiVersionSet(ctx, &args, opts...)
			var s LookupApiVersionSetResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupApiVersionSetResultOutput)
}

type LookupApiVersionSetOutputArgs struct {
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	ServiceName       pulumi.StringInput `pulumi:"serviceName"`
	VersionSetId      pulumi.StringInput `pulumi:"versionSetId"`
}

func (LookupApiVersionSetOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupApiVersionSetArgs)(nil)).Elem()
}


type LookupApiVersionSetResultOutput struct{ *pulumi.OutputState }

func (LookupApiVersionSetResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupApiVersionSetResult)(nil)).Elem()
}

func (o LookupApiVersionSetResultOutput) ToLookupApiVersionSetResultOutput() LookupApiVersionSetResultOutput {
	return o
}

func (o LookupApiVersionSetResultOutput) ToLookupApiVersionSetResultOutputWithContext(ctx context.Context) LookupApiVersionSetResultOutput {
	return o
}

func (o LookupApiVersionSetResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupApiVersionSetResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupApiVersionSetResultOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApiVersionSetResult) string { return v.DisplayName }).(pulumi.StringOutput)
}

func (o LookupApiVersionSetResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApiVersionSetResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupApiVersionSetResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApiVersionSetResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupApiVersionSetResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApiVersionSetResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupApiVersionSetResultOutput) VersionHeaderName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupApiVersionSetResult) *string { return v.VersionHeaderName }).(pulumi.StringPtrOutput)
}

func (o LookupApiVersionSetResultOutput) VersionQueryName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupApiVersionSetResult) *string { return v.VersionQueryName }).(pulumi.StringPtrOutput)
}

func (o LookupApiVersionSetResultOutput) VersioningScheme() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApiVersionSetResult) string { return v.VersioningScheme }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupApiVersionSetResultOutput{})
}
