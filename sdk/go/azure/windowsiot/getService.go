


package windowsiot

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupService(ctx *pulumi.Context, args *LookupServiceArgs, opts ...pulumi.InvokeOption) (*LookupServiceResult, error) {
	var rv LookupServiceResult
	err := ctx.Invoke("azure-native:windowsiot:getService", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupServiceArgs struct {
	DeviceName        string `pulumi:"deviceName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupServiceResult struct {
	AdminDomainName   *string           `pulumi:"adminDomainName"`
	BillingDomainName *string           `pulumi:"billingDomainName"`
	Etag              *string           `pulumi:"etag"`
	Id                string            `pulumi:"id"`
	Location          *string           `pulumi:"location"`
	Name              string            `pulumi:"name"`
	Notes             *string           `pulumi:"notes"`
	Quantity          *float64          `pulumi:"quantity"`
	StartDate         string            `pulumi:"startDate"`
	Tags              map[string]string `pulumi:"tags"`
	Type              string            `pulumi:"type"`
}

func LookupServiceOutput(ctx *pulumi.Context, args LookupServiceOutputArgs, opts ...pulumi.InvokeOption) LookupServiceResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupServiceResult, error) {
			args := v.(LookupServiceArgs)
			r, err := LookupService(ctx, &args, opts...)
			var s LookupServiceResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupServiceResultOutput)
}

type LookupServiceOutputArgs struct {
	DeviceName        pulumi.StringInput `pulumi:"deviceName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupServiceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupServiceArgs)(nil)).Elem()
}


type LookupServiceResultOutput struct{ *pulumi.OutputState }

func (LookupServiceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupServiceResult)(nil)).Elem()
}

func (o LookupServiceResultOutput) ToLookupServiceResultOutput() LookupServiceResultOutput {
	return o
}

func (o LookupServiceResultOutput) ToLookupServiceResultOutputWithContext(ctx context.Context) LookupServiceResultOutput {
	return o
}

func (o LookupServiceResultOutput) AdminDomainName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupServiceResult) *string { return v.AdminDomainName }).(pulumi.StringPtrOutput)
}

func (o LookupServiceResultOutput) BillingDomainName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupServiceResult) *string { return v.BillingDomainName }).(pulumi.StringPtrOutput)
}

func (o LookupServiceResultOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupServiceResult) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o LookupServiceResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServiceResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupServiceResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupServiceResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupServiceResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServiceResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupServiceResultOutput) Notes() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupServiceResult) *string { return v.Notes }).(pulumi.StringPtrOutput)
}

func (o LookupServiceResultOutput) Quantity() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v LookupServiceResult) *float64 { return v.Quantity }).(pulumi.Float64PtrOutput)
}

func (o LookupServiceResultOutput) StartDate() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServiceResult) string { return v.StartDate }).(pulumi.StringOutput)
}

func (o LookupServiceResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupServiceResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupServiceResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServiceResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupServiceResultOutput{})
}
