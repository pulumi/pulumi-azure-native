


package v20180301preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


func LookupRecordSet(ctx *pulumi.Context, args *LookupRecordSetArgs, opts ...pulumi.InvokeOption) (*LookupRecordSetResult, error) {
	var rv LookupRecordSetResult
	err := ctx.Invoke("azure-native:network/v20180301preview:getRecordSet", args, &rv, opts...)
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
	ARecords    []ARecordResponse    `pulumi:"aRecords"`
	AaaaRecords []AaaaRecordResponse `pulumi:"aaaaRecords"`
	CaaRecords  []CaaRecordResponse  `pulumi:"caaRecords"`
	CnameRecord *CnameRecordResponse `pulumi:"cnameRecord"`
	Etag        *string              `pulumi:"etag"`
	Fqdn        string               `pulumi:"fqdn"`
	Id          string               `pulumi:"id"`
	Metadata    map[string]string    `pulumi:"metadata"`
	MxRecords   []MxRecordResponse   `pulumi:"mxRecords"`
	Name        string               `pulumi:"name"`
	NsRecords   []NsRecordResponse   `pulumi:"nsRecords"`
	PtrRecords  []PtrRecordResponse  `pulumi:"ptrRecords"`
	SoaRecord   *SoaRecordResponse   `pulumi:"soaRecord"`
	SrvRecords  []SrvRecordResponse  `pulumi:"srvRecords"`
	Ttl         *float64             `pulumi:"ttl"`
	TxtRecords  []TxtRecordResponse  `pulumi:"txtRecords"`
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

func (o LookupRecordSetResultOutput) ARecords() ARecordResponseArrayOutput {
	return o.ApplyT(func(v LookupRecordSetResult) []ARecordResponse { return v.ARecords }).(ARecordResponseArrayOutput)
}

func (o LookupRecordSetResultOutput) AaaaRecords() AaaaRecordResponseArrayOutput {
	return o.ApplyT(func(v LookupRecordSetResult) []AaaaRecordResponse { return v.AaaaRecords }).(AaaaRecordResponseArrayOutput)
}

func (o LookupRecordSetResultOutput) CaaRecords() CaaRecordResponseArrayOutput {
	return o.ApplyT(func(v LookupRecordSetResult) []CaaRecordResponse { return v.CaaRecords }).(CaaRecordResponseArrayOutput)
}

func (o LookupRecordSetResultOutput) CnameRecord() CnameRecordResponsePtrOutput {
	return o.ApplyT(func(v LookupRecordSetResult) *CnameRecordResponse { return v.CnameRecord }).(CnameRecordResponsePtrOutput)
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

func (o LookupRecordSetResultOutput) Metadata() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupRecordSetResult) map[string]string { return v.Metadata }).(pulumi.StringMapOutput)
}

func (o LookupRecordSetResultOutput) MxRecords() MxRecordResponseArrayOutput {
	return o.ApplyT(func(v LookupRecordSetResult) []MxRecordResponse { return v.MxRecords }).(MxRecordResponseArrayOutput)
}

func (o LookupRecordSetResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRecordSetResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupRecordSetResultOutput) NsRecords() NsRecordResponseArrayOutput {
	return o.ApplyT(func(v LookupRecordSetResult) []NsRecordResponse { return v.NsRecords }).(NsRecordResponseArrayOutput)
}

func (o LookupRecordSetResultOutput) PtrRecords() PtrRecordResponseArrayOutput {
	return o.ApplyT(func(v LookupRecordSetResult) []PtrRecordResponse { return v.PtrRecords }).(PtrRecordResponseArrayOutput)
}

func (o LookupRecordSetResultOutput) SoaRecord() SoaRecordResponsePtrOutput {
	return o.ApplyT(func(v LookupRecordSetResult) *SoaRecordResponse { return v.SoaRecord }).(SoaRecordResponsePtrOutput)
}

func (o LookupRecordSetResultOutput) SrvRecords() SrvRecordResponseArrayOutput {
	return o.ApplyT(func(v LookupRecordSetResult) []SrvRecordResponse { return v.SrvRecords }).(SrvRecordResponseArrayOutput)
}

func (o LookupRecordSetResultOutput) Ttl() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v LookupRecordSetResult) *float64 { return v.Ttl }).(pulumi.Float64PtrOutput)
}

func (o LookupRecordSetResultOutput) TxtRecords() TxtRecordResponseArrayOutput {
	return o.ApplyT(func(v LookupRecordSetResult) []TxtRecordResponse { return v.TxtRecords }).(TxtRecordResponseArrayOutput)
}

func (o LookupRecordSetResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRecordSetResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupRecordSetResultOutput{})
}
