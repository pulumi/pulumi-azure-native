


package v20180201

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupPartner(ctx *pulumi.Context, args *LookupPartnerArgs, opts ...pulumi.InvokeOption) (*LookupPartnerResult, error) {
	var rv LookupPartnerResult
	err := ctx.Invoke("azure-native:managementpartner/v20180201:getPartner", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupPartnerArgs struct {
	PartnerId string `pulumi:"partnerId"`
}


type LookupPartnerResult struct {
	CreatedTime *string `pulumi:"createdTime"`
	Etag        *int    `pulumi:"etag"`
	Id          string  `pulumi:"id"`
	Name        string  `pulumi:"name"`
	ObjectId    *string `pulumi:"objectId"`
	PartnerId   *string `pulumi:"partnerId"`
	PartnerName *string `pulumi:"partnerName"`
	TenantId    *string `pulumi:"tenantId"`
	Type        string  `pulumi:"type"`
	UpdatedTime *string `pulumi:"updatedTime"`
	Version     *int    `pulumi:"version"`
}

func LookupPartnerOutput(ctx *pulumi.Context, args LookupPartnerOutputArgs, opts ...pulumi.InvokeOption) LookupPartnerResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupPartnerResult, error) {
			args := v.(LookupPartnerArgs)
			r, err := LookupPartner(ctx, &args, opts...)
			var s LookupPartnerResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupPartnerResultOutput)
}

type LookupPartnerOutputArgs struct {
	PartnerId pulumi.StringInput `pulumi:"partnerId"`
}

func (LookupPartnerOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPartnerArgs)(nil)).Elem()
}


type LookupPartnerResultOutput struct{ *pulumi.OutputState }

func (LookupPartnerResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPartnerResult)(nil)).Elem()
}

func (o LookupPartnerResultOutput) ToLookupPartnerResultOutput() LookupPartnerResultOutput {
	return o
}

func (o LookupPartnerResultOutput) ToLookupPartnerResultOutputWithContext(ctx context.Context) LookupPartnerResultOutput {
	return o
}

func (o LookupPartnerResultOutput) CreatedTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPartnerResult) *string { return v.CreatedTime }).(pulumi.StringPtrOutput)
}

func (o LookupPartnerResultOutput) Etag() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupPartnerResult) *int { return v.Etag }).(pulumi.IntPtrOutput)
}

func (o LookupPartnerResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPartnerResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupPartnerResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPartnerResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupPartnerResultOutput) ObjectId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPartnerResult) *string { return v.ObjectId }).(pulumi.StringPtrOutput)
}

func (o LookupPartnerResultOutput) PartnerId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPartnerResult) *string { return v.PartnerId }).(pulumi.StringPtrOutput)
}

func (o LookupPartnerResultOutput) PartnerName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPartnerResult) *string { return v.PartnerName }).(pulumi.StringPtrOutput)
}

func (o LookupPartnerResultOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPartnerResult) *string { return v.TenantId }).(pulumi.StringPtrOutput)
}

func (o LookupPartnerResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPartnerResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupPartnerResultOutput) UpdatedTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPartnerResult) *string { return v.UpdatedTime }).(pulumi.StringPtrOutput)
}

func (o LookupPartnerResultOutput) Version() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupPartnerResult) *int { return v.Version }).(pulumi.IntPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupPartnerResultOutput{})
}
