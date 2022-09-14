


package v20211001preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type DnsRecordResponse struct {
	Name  string `pulumi:"name"`
	Ttl   int    `pulumi:"ttl"`
	Type  string `pulumi:"type"`
	Value string `pulumi:"value"`
}

type DnsRecordResponseOutput struct{ *pulumi.OutputState }

func (DnsRecordResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DnsRecordResponse)(nil)).Elem()
}

func (o DnsRecordResponseOutput) ToDnsRecordResponseOutput() DnsRecordResponseOutput {
	return o
}

func (o DnsRecordResponseOutput) ToDnsRecordResponseOutputWithContext(ctx context.Context) DnsRecordResponseOutput {
	return o
}

func (o DnsRecordResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v DnsRecordResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o DnsRecordResponseOutput) Ttl() pulumi.IntOutput {
	return o.ApplyT(func(v DnsRecordResponse) int { return v.Ttl }).(pulumi.IntOutput)
}

func (o DnsRecordResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v DnsRecordResponse) string { return v.Type }).(pulumi.StringOutput)
}

func (o DnsRecordResponseOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v DnsRecordResponse) string { return v.Value }).(pulumi.StringOutput)
}

type DnsRecordResponsePtrOutput struct{ *pulumi.OutputState }

func (DnsRecordResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DnsRecordResponse)(nil)).Elem()
}

func (o DnsRecordResponsePtrOutput) ToDnsRecordResponsePtrOutput() DnsRecordResponsePtrOutput {
	return o
}

func (o DnsRecordResponsePtrOutput) ToDnsRecordResponsePtrOutputWithContext(ctx context.Context) DnsRecordResponsePtrOutput {
	return o
}

func (o DnsRecordResponsePtrOutput) Elem() DnsRecordResponseOutput {
	return o.ApplyT(func(v *DnsRecordResponse) DnsRecordResponse {
		if v != nil {
			return *v
		}
		var ret DnsRecordResponse
		return ret
	}).(DnsRecordResponseOutput)
}

func (o DnsRecordResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DnsRecordResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

func (o DnsRecordResponsePtrOutput) Ttl() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *DnsRecordResponse) *int {
		if v == nil {
			return nil
		}
		return &v.Ttl
	}).(pulumi.IntPtrOutput)
}

func (o DnsRecordResponsePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DnsRecordResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Type
	}).(pulumi.StringPtrOutput)
}

func (o DnsRecordResponsePtrOutput) Value() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DnsRecordResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Value
	}).(pulumi.StringPtrOutput)
}

type DomainPropertiesResponseVerificationRecords struct {
	DKIM   *DnsRecordResponse `pulumi:"dKIM"`
	DKIM2  *DnsRecordResponse `pulumi:"dKIM2"`
	DMARC  *DnsRecordResponse `pulumi:"dMARC"`
	Domain *DnsRecordResponse `pulumi:"domain"`
	SPF    *DnsRecordResponse `pulumi:"sPF"`
}

type DomainPropertiesResponseVerificationRecordsOutput struct{ *pulumi.OutputState }

func (DomainPropertiesResponseVerificationRecordsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DomainPropertiesResponseVerificationRecords)(nil)).Elem()
}

func (o DomainPropertiesResponseVerificationRecordsOutput) ToDomainPropertiesResponseVerificationRecordsOutput() DomainPropertiesResponseVerificationRecordsOutput {
	return o
}

func (o DomainPropertiesResponseVerificationRecordsOutput) ToDomainPropertiesResponseVerificationRecordsOutputWithContext(ctx context.Context) DomainPropertiesResponseVerificationRecordsOutput {
	return o
}

func (o DomainPropertiesResponseVerificationRecordsOutput) DKIM() DnsRecordResponsePtrOutput {
	return o.ApplyT(func(v DomainPropertiesResponseVerificationRecords) *DnsRecordResponse { return v.DKIM }).(DnsRecordResponsePtrOutput)
}

func (o DomainPropertiesResponseVerificationRecordsOutput) DKIM2() DnsRecordResponsePtrOutput {
	return o.ApplyT(func(v DomainPropertiesResponseVerificationRecords) *DnsRecordResponse { return v.DKIM2 }).(DnsRecordResponsePtrOutput)
}

func (o DomainPropertiesResponseVerificationRecordsOutput) DMARC() DnsRecordResponsePtrOutput {
	return o.ApplyT(func(v DomainPropertiesResponseVerificationRecords) *DnsRecordResponse { return v.DMARC }).(DnsRecordResponsePtrOutput)
}

func (o DomainPropertiesResponseVerificationRecordsOutput) Domain() DnsRecordResponsePtrOutput {
	return o.ApplyT(func(v DomainPropertiesResponseVerificationRecords) *DnsRecordResponse { return v.Domain }).(DnsRecordResponsePtrOutput)
}

func (o DomainPropertiesResponseVerificationRecordsOutput) SPF() DnsRecordResponsePtrOutput {
	return o.ApplyT(func(v DomainPropertiesResponseVerificationRecords) *DnsRecordResponse { return v.SPF }).(DnsRecordResponsePtrOutput)
}

type DomainPropertiesResponseVerificationStates struct {
	DKIM   *VerificationStatusRecordResponse `pulumi:"dKIM"`
	DKIM2  *VerificationStatusRecordResponse `pulumi:"dKIM2"`
	DMARC  *VerificationStatusRecordResponse `pulumi:"dMARC"`
	Domain *VerificationStatusRecordResponse `pulumi:"domain"`
	SPF    *VerificationStatusRecordResponse `pulumi:"sPF"`
}

type DomainPropertiesResponseVerificationStatesOutput struct{ *pulumi.OutputState }

func (DomainPropertiesResponseVerificationStatesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DomainPropertiesResponseVerificationStates)(nil)).Elem()
}

func (o DomainPropertiesResponseVerificationStatesOutput) ToDomainPropertiesResponseVerificationStatesOutput() DomainPropertiesResponseVerificationStatesOutput {
	return o
}

func (o DomainPropertiesResponseVerificationStatesOutput) ToDomainPropertiesResponseVerificationStatesOutputWithContext(ctx context.Context) DomainPropertiesResponseVerificationStatesOutput {
	return o
}

func (o DomainPropertiesResponseVerificationStatesOutput) DKIM() VerificationStatusRecordResponsePtrOutput {
	return o.ApplyT(func(v DomainPropertiesResponseVerificationStates) *VerificationStatusRecordResponse { return v.DKIM }).(VerificationStatusRecordResponsePtrOutput)
}

func (o DomainPropertiesResponseVerificationStatesOutput) DKIM2() VerificationStatusRecordResponsePtrOutput {
	return o.ApplyT(func(v DomainPropertiesResponseVerificationStates) *VerificationStatusRecordResponse { return v.DKIM2 }).(VerificationStatusRecordResponsePtrOutput)
}

func (o DomainPropertiesResponseVerificationStatesOutput) DMARC() VerificationStatusRecordResponsePtrOutput {
	return o.ApplyT(func(v DomainPropertiesResponseVerificationStates) *VerificationStatusRecordResponse { return v.DMARC }).(VerificationStatusRecordResponsePtrOutput)
}

func (o DomainPropertiesResponseVerificationStatesOutput) Domain() VerificationStatusRecordResponsePtrOutput {
	return o.ApplyT(func(v DomainPropertiesResponseVerificationStates) *VerificationStatusRecordResponse { return v.Domain }).(VerificationStatusRecordResponsePtrOutput)
}

func (o DomainPropertiesResponseVerificationStatesOutput) SPF() VerificationStatusRecordResponsePtrOutput {
	return o.ApplyT(func(v DomainPropertiesResponseVerificationStates) *VerificationStatusRecordResponse { return v.SPF }).(VerificationStatusRecordResponsePtrOutput)
}

type SystemDataResponse struct {
	CreatedAt          *string `pulumi:"createdAt"`
	CreatedBy          *string `pulumi:"createdBy"`
	CreatedByType      *string `pulumi:"createdByType"`
	LastModifiedAt     *string `pulumi:"lastModifiedAt"`
	LastModifiedBy     *string `pulumi:"lastModifiedBy"`
	LastModifiedByType *string `pulumi:"lastModifiedByType"`
}

type SystemDataResponseOutput struct{ *pulumi.OutputState }

func (SystemDataResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SystemDataResponse)(nil)).Elem()
}

func (o SystemDataResponseOutput) ToSystemDataResponseOutput() SystemDataResponseOutput {
	return o
}

func (o SystemDataResponseOutput) ToSystemDataResponseOutputWithContext(ctx context.Context) SystemDataResponseOutput {
	return o
}

func (o SystemDataResponseOutput) CreatedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.CreatedAt }).(pulumi.StringPtrOutput)
}

func (o SystemDataResponseOutput) CreatedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.CreatedBy }).(pulumi.StringPtrOutput)
}

func (o SystemDataResponseOutput) CreatedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.CreatedByType }).(pulumi.StringPtrOutput)
}

func (o SystemDataResponseOutput) LastModifiedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.LastModifiedAt }).(pulumi.StringPtrOutput)
}

func (o SystemDataResponseOutput) LastModifiedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.LastModifiedBy }).(pulumi.StringPtrOutput)
}

func (o SystemDataResponseOutput) LastModifiedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.LastModifiedByType }).(pulumi.StringPtrOutput)
}

type VerificationStatusRecordResponse struct {
	ErrorCode string `pulumi:"errorCode"`
	Status    string `pulumi:"status"`
}

type VerificationStatusRecordResponseOutput struct{ *pulumi.OutputState }

func (VerificationStatusRecordResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VerificationStatusRecordResponse)(nil)).Elem()
}

func (o VerificationStatusRecordResponseOutput) ToVerificationStatusRecordResponseOutput() VerificationStatusRecordResponseOutput {
	return o
}

func (o VerificationStatusRecordResponseOutput) ToVerificationStatusRecordResponseOutputWithContext(ctx context.Context) VerificationStatusRecordResponseOutput {
	return o
}

func (o VerificationStatusRecordResponseOutput) ErrorCode() pulumi.StringOutput {
	return o.ApplyT(func(v VerificationStatusRecordResponse) string { return v.ErrorCode }).(pulumi.StringOutput)
}

func (o VerificationStatusRecordResponseOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v VerificationStatusRecordResponse) string { return v.Status }).(pulumi.StringOutput)
}

type VerificationStatusRecordResponsePtrOutput struct{ *pulumi.OutputState }

func (VerificationStatusRecordResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VerificationStatusRecordResponse)(nil)).Elem()
}

func (o VerificationStatusRecordResponsePtrOutput) ToVerificationStatusRecordResponsePtrOutput() VerificationStatusRecordResponsePtrOutput {
	return o
}

func (o VerificationStatusRecordResponsePtrOutput) ToVerificationStatusRecordResponsePtrOutputWithContext(ctx context.Context) VerificationStatusRecordResponsePtrOutput {
	return o
}

func (o VerificationStatusRecordResponsePtrOutput) Elem() VerificationStatusRecordResponseOutput {
	return o.ApplyT(func(v *VerificationStatusRecordResponse) VerificationStatusRecordResponse {
		if v != nil {
			return *v
		}
		var ret VerificationStatusRecordResponse
		return ret
	}).(VerificationStatusRecordResponseOutput)
}

func (o VerificationStatusRecordResponsePtrOutput) ErrorCode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VerificationStatusRecordResponse) *string {
		if v == nil {
			return nil
		}
		return &v.ErrorCode
	}).(pulumi.StringPtrOutput)
}

func (o VerificationStatusRecordResponsePtrOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VerificationStatusRecordResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Status
	}).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(DnsRecordResponseOutput{})
	pulumi.RegisterOutputType(DnsRecordResponsePtrOutput{})
	pulumi.RegisterOutputType(DomainPropertiesResponseVerificationRecordsOutput{})
	pulumi.RegisterOutputType(DomainPropertiesResponseVerificationStatesOutput{})
	pulumi.RegisterOutputType(SystemDataResponseOutput{})
	pulumi.RegisterOutputType(VerificationStatusRecordResponseOutput{})
	pulumi.RegisterOutputType(VerificationStatusRecordResponsePtrOutput{})
}
