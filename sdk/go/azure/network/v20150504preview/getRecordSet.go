


package v20150504preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


func LookupRecordSet(ctx *pulumi.Context, args *LookupRecordSetArgs, opts ...pulumi.InvokeOption) (*LookupRecordSetResult, error) {
	var rv LookupRecordSetResult
	err := ctx.Invoke("azure-native:network/v20150504preview:getRecordSet", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupRecordSetArgs struct {
	RecordType            string `pulumi:"recordType"`
	RelativeRecordSetName string `pulumi:"relativeRecordSetName"`
	ResourceGroupName     string `pulumi:"resourceGroupName"`
	ZoneName              string `pulumi:"zoneName"`
}


type LookupRecordSetResult struct {
	AAAARecords []AaaaRecordResponse `pulumi:"aAAARecords"`
	ARecords    []ARecordResponse    `pulumi:"aRecords"`
	CNAMERecord *CnameRecordResponse `pulumi:"cNAMERecord"`
	Etag        *string              `pulumi:"etag"`
	Fqdn        string               `pulumi:"fqdn"`
	Id          string               `pulumi:"id"`
	MXRecords   []MxRecordResponse   `pulumi:"mXRecords"`
	NSRecords   []NsRecordResponse   `pulumi:"nSRecords"`
	Name        string               `pulumi:"name"`
	PTRRecords  []PtrRecordResponse  `pulumi:"pTRRecords"`
	SOARecord   *SoaRecordResponse   `pulumi:"sOARecord"`
	SRVRecords  []SrvRecordResponse  `pulumi:"sRVRecords"`
	TXTRecords  []TxtRecordResponse  `pulumi:"tXTRecords"`
	Ttl         *float64             `pulumi:"ttl"`
	Type        string               `pulumi:"type"`
}

func LookupRecordSetOutput(ctx *pulumi.Context, args LookupRecordSetOutputArgs, opts ...pulumi.InvokeOption) LookupRecordSetResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupRecordSetResult, error) {
			args := v.(LookupRecordSetArgs)
			r, err := LookupRecordSet(ctx, &args, opts...)
			var s LookupRecordSetResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupRecordSetResultOutput)
}

type LookupRecordSetOutputArgs struct {
	RecordType            pulumi.StringInput `pulumi:"recordType"`
	RelativeRecordSetName pulumi.StringInput `pulumi:"relativeRecordSetName"`
	ResourceGroupName     pulumi.StringInput `pulumi:"resourceGroupName"`
	ZoneName              pulumi.StringInput `pulumi:"zoneName"`
}

func (LookupRecordSetOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRecordSetArgs)(nil)).Elem()
}


type LookupRecordSetResultOutput struct{ *pulumi.OutputState }

func (LookupRecordSetResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRecordSetResult)(nil)).Elem()
}

func (o LookupRecordSetResultOutput) ToLookupRecordSetResultOutput() LookupRecordSetResultOutput {
	return o
}

func (o LookupRecordSetResultOutput) ToLookupRecordSetResultOutputWithContext(ctx context.Context) LookupRecordSetResultOutput {
	return o
}

func (o LookupRecordSetResultOutput) AAAARecords() AaaaRecordResponseArrayOutput {
	return o.ApplyT(func(v LookupRecordSetResult) []AaaaRecordResponse { return v.AAAARecords }).(AaaaRecordResponseArrayOutput)
}

func (o LookupRecordSetResultOutput) ARecords() ARecordResponseArrayOutput {
	return o.ApplyT(func(v LookupRecordSetResult) []ARecordResponse { return v.ARecords }).(ARecordResponseArrayOutput)
}

func (o LookupRecordSetResultOutput) CNAMERecord() CnameRecordResponsePtrOutput {
	return o.ApplyT(func(v LookupRecordSetResult) *CnameRecordResponse { return v.CNAMERecord }).(CnameRecordResponsePtrOutput)
}

func (o LookupRecordSetResultOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRecordSetResult) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o LookupRecordSetResultOutput) Fqdn() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRecordSetResult) string { return v.Fqdn }).(pulumi.StringOutput)
}

func (o LookupRecordSetResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRecordSetResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupRecordSetResultOutput) MXRecords() MxRecordResponseArrayOutput {
	return o.ApplyT(func(v LookupRecordSetResult) []MxRecordResponse { return v.MXRecords }).(MxRecordResponseArrayOutput)
}

func (o LookupRecordSetResultOutput) NSRecords() NsRecordResponseArrayOutput {
	return o.ApplyT(func(v LookupRecordSetResult) []NsRecordResponse { return v.NSRecords }).(NsRecordResponseArrayOutput)
}

func (o LookupRecordSetResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRecordSetResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupRecordSetResultOutput) PTRRecords() PtrRecordResponseArrayOutput {
	return o.ApplyT(func(v LookupRecordSetResult) []PtrRecordResponse { return v.PTRRecords }).(PtrRecordResponseArrayOutput)
}

func (o LookupRecordSetResultOutput) SOARecord() SoaRecordResponsePtrOutput {
	return o.ApplyT(func(v LookupRecordSetResult) *SoaRecordResponse { return v.SOARecord }).(SoaRecordResponsePtrOutput)
}

func (o LookupRecordSetResultOutput) SRVRecords() SrvRecordResponseArrayOutput {
	return o.ApplyT(func(v LookupRecordSetResult) []SrvRecordResponse { return v.SRVRecords }).(SrvRecordResponseArrayOutput)
}

func (o LookupRecordSetResultOutput) TXTRecords() TxtRecordResponseArrayOutput {
	return o.ApplyT(func(v LookupRecordSetResult) []TxtRecordResponse { return v.TXTRecords }).(TxtRecordResponseArrayOutput)
}

func (o LookupRecordSetResultOutput) Ttl() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v LookupRecordSetResult) *float64 { return v.Ttl }).(pulumi.Float64PtrOutput)
}

func (o LookupRecordSetResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRecordSetResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupRecordSetResultOutput{})
}
