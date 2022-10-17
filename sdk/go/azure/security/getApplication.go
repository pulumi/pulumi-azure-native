


package security

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupApplication(ctx *pulumi.Context, args *LookupApplicationArgs, opts ...pulumi.InvokeOption) (*LookupApplicationResult, error) {
	var rv LookupApplicationResult
	err := ctx.Invoke("azure-native:security:getApplication", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupApplicationArgs struct {
	ApplicationId string `pulumi:"applicationId"`
}


type LookupApplicationResult struct {
	Description        *string `pulumi:"description"`
	DisplayName        *string `pulumi:"displayName"`
	Id                 string  `pulumi:"id"`
	Name               string  `pulumi:"name"`
	SourceResourceType string  `pulumi:"sourceResourceType"`
	Type               string  `pulumi:"type"`
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
	ApplicationId pulumi.StringInput `pulumi:"applicationId"`
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

func (o LookupApplicationResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupApplicationResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupApplicationResultOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupApplicationResult) *string { return v.DisplayName }).(pulumi.StringPtrOutput)
}

func (o LookupApplicationResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApplicationResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupApplicationResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApplicationResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupApplicationResultOutput) SourceResourceType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApplicationResult) string { return v.SourceResourceType }).(pulumi.StringOutput)
}

func (o LookupApplicationResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApplicationResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupApplicationResultOutput{})
}
