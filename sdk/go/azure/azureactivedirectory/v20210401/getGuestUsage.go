


package v20210401

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupGuestUsage(ctx *pulumi.Context, args *LookupGuestUsageArgs, opts ...pulumi.InvokeOption) (*LookupGuestUsageResult, error) {
	var rv LookupGuestUsageResult
	err := ctx.Invoke("azure-native:azureactivedirectory/v20210401:getGuestUsage", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupGuestUsageArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ResourceName      string `pulumi:"resourceName"`
}


type LookupGuestUsageResult struct {
	Id         string             `pulumi:"id"`
	Location   *string            `pulumi:"location"`
	Name       string             `pulumi:"name"`
	SystemData SystemDataResponse `pulumi:"systemData"`
	Tags       map[string]string  `pulumi:"tags"`
	TenantId   *string            `pulumi:"tenantId"`
	Type       string             `pulumi:"type"`
}

func LookupGuestUsageOutput(ctx *pulumi.Context, args LookupGuestUsageOutputArgs, opts ...pulumi.InvokeOption) LookupGuestUsageResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupGuestUsageResult, error) {
			args := v.(LookupGuestUsageArgs)
			r, err := LookupGuestUsage(ctx, &args, opts...)
			var s LookupGuestUsageResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupGuestUsageResultOutput)
}

type LookupGuestUsageOutputArgs struct {
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	ResourceName      pulumi.StringInput `pulumi:"resourceName"`
}

func (LookupGuestUsageOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupGuestUsageArgs)(nil)).Elem()
}


type LookupGuestUsageResultOutput struct{ *pulumi.OutputState }

func (LookupGuestUsageResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupGuestUsageResult)(nil)).Elem()
}

func (o LookupGuestUsageResultOutput) ToLookupGuestUsageResultOutput() LookupGuestUsageResultOutput {
	return o
}

func (o LookupGuestUsageResultOutput) ToLookupGuestUsageResultOutputWithContext(ctx context.Context) LookupGuestUsageResultOutput {
	return o
}

func (o LookupGuestUsageResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGuestUsageResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupGuestUsageResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupGuestUsageResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupGuestUsageResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGuestUsageResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupGuestUsageResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupGuestUsageResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupGuestUsageResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupGuestUsageResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupGuestUsageResultOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupGuestUsageResult) *string { return v.TenantId }).(pulumi.StringPtrOutput)
}

func (o LookupGuestUsageResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGuestUsageResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupGuestUsageResultOutput{})
}
