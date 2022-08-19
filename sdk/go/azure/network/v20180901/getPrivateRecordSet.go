


package v20180901

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupPrivateRecordSet(ctx *pulumi.Context, args *LookupPrivateRecordSetArgs, opts ...pulumi.InvokeOption) (*LookupPrivateRecordSetResult, error) {
	var rv LookupPrivateRecordSetResult
	err := ctx.Invoke("azure-native:network/v20180901:getPrivateRecordSet", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupPrivateRecordSetArgs struct {
	PrivateZoneName       string `pulumi:"privateZoneName"`
	RecordType            string `pulumi:"recordType"`
	RelativeRecordSetName string `pulumi:"relativeRecordSetName"`
	ResourceGroupName     string `pulumi:"resourceGroupName"`
}


type LookupPrivateRecordSetResult struct {
	ARecords         []ARecordResponse    `pulumi:"aRecords"`
	AaaaRecords      []AaaaRecordResponse `pulumi:"aaaaRecords"`
	CnameRecord      *CnameRecordResponse `pulumi:"cnameRecord"`
	Etag             *string              `pulumi:"etag"`
	Fqdn             string               `pulumi:"fqdn"`
	Id               string               `pulumi:"id"`
	IsAutoRegistered bool                 `pulumi:"isAutoRegistered"`
	Metadata         map[string]string    `pulumi:"metadata"`
	MxRecords        []MxRecordResponse   `pulumi:"mxRecords"`
	Name             string               `pulumi:"name"`
	PtrRecords       []PtrRecordResponse  `pulumi:"ptrRecords"`
	SoaRecord        *SoaRecordResponse   `pulumi:"soaRecord"`
	SrvRecords       []SrvRecordResponse  `pulumi:"srvRecords"`
	Ttl              *float64             `pulumi:"ttl"`
	TxtRecords       []TxtRecordResponse  `pulumi:"txtRecords"`
	Type             string               `pulumi:"type"`
}

func LookupPrivateRecordSetOutput(ctx *pulumi.Context, args LookupPrivateRecordSetOutputArgs, opts ...pulumi.InvokeOption) LookupPrivateRecordSetResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupPrivateRecordSetResult, error) {
			args := v.(LookupPrivateRecordSetArgs)
			r, err := LookupPrivateRecordSet(ctx, &args, opts...)
			var s LookupPrivateRecordSetResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupPrivateRecordSetResultOutput)
}

type LookupPrivateRecordSetOutputArgs struct {
	PrivateZoneName       pulumi.StringInput `pulumi:"privateZoneName"`
	RecordType            pulumi.StringInput `pulumi:"recordType"`
	RelativeRecordSetName pulumi.StringInput `pulumi:"relativeRecordSetName"`
	ResourceGroupName     pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupPrivateRecordSetOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPrivateRecordSetArgs)(nil)).Elem()
}


type LookupPrivateRecordSetResultOutput struct{ *pulumi.OutputState }

func (LookupPrivateRecordSetResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPrivateRecordSetResult)(nil)).Elem()
}

func (o LookupPrivateRecordSetResultOutput) ToLookupPrivateRecordSetResultOutput() LookupPrivateRecordSetResultOutput {
	return o
}

func (o LookupPrivateRecordSetResultOutput) ToLookupPrivateRecordSetResultOutputWithContext(ctx context.Context) LookupPrivateRecordSetResultOutput {
	return o
}

func (o LookupPrivateRecordSetResultOutput) ARecords() ARecordResponseArrayOutput {
	return o.ApplyT(func(v LookupPrivateRecordSetResult) []ARecordResponse { return v.ARecords }).(ARecordResponseArrayOutput)
}

func (o LookupPrivateRecordSetResultOutput) AaaaRecords() AaaaRecordResponseArrayOutput {
	return o.ApplyT(func(v LookupPrivateRecordSetResult) []AaaaRecordResponse { return v.AaaaRecords }).(AaaaRecordResponseArrayOutput)
}

func (o LookupPrivateRecordSetResultOutput) CnameRecord() CnameRecordResponsePtrOutput {
	return o.ApplyT(func(v LookupPrivateRecordSetResult) *CnameRecordResponse { return v.CnameRecord }).(CnameRecordResponsePtrOutput)
}

func (o LookupPrivateRecordSetResultOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPrivateRecordSetResult) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o LookupPrivateRecordSetResultOutput) Fqdn() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateRecordSetResult) string { return v.Fqdn }).(pulumi.StringOutput)
}

func (o LookupPrivateRecordSetResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateRecordSetResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupPrivateRecordSetResultOutput) IsAutoRegistered() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupPrivateRecordSetResult) bool { return v.IsAutoRegistered }).(pulumi.BoolOutput)
}

func (o LookupPrivateRecordSetResultOutput) Metadata() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupPrivateRecordSetResult) map[string]string { return v.Metadata }).(pulumi.StringMapOutput)
}

func (o LookupPrivateRecordSetResultOutput) MxRecords() MxRecordResponseArrayOutput {
	return o.ApplyT(func(v LookupPrivateRecordSetResult) []MxRecordResponse { return v.MxRecords }).(MxRecordResponseArrayOutput)
}

func (o LookupPrivateRecordSetResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateRecordSetResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupPrivateRecordSetResultOutput) PtrRecords() PtrRecordResponseArrayOutput {
	return o.ApplyT(func(v LookupPrivateRecordSetResult) []PtrRecordResponse { return v.PtrRecords }).(PtrRecordResponseArrayOutput)
}

func (o LookupPrivateRecordSetResultOutput) SoaRecord() SoaRecordResponsePtrOutput {
	return o.ApplyT(func(v LookupPrivateRecordSetResult) *SoaRecordResponse { return v.SoaRecord }).(SoaRecordResponsePtrOutput)
}

func (o LookupPrivateRecordSetResultOutput) SrvRecords() SrvRecordResponseArrayOutput {
	return o.ApplyT(func(v LookupPrivateRecordSetResult) []SrvRecordResponse { return v.SrvRecords }).(SrvRecordResponseArrayOutput)
}

func (o LookupPrivateRecordSetResultOutput) Ttl() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v LookupPrivateRecordSetResult) *float64 { return v.Ttl }).(pulumi.Float64PtrOutput)
}

func (o LookupPrivateRecordSetResultOutput) TxtRecords() TxtRecordResponseArrayOutput {
	return o.ApplyT(func(v LookupPrivateRecordSetResult) []TxtRecordResponse { return v.TxtRecords }).(TxtRecordResponseArrayOutput)
}

func (o LookupPrivateRecordSetResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateRecordSetResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupPrivateRecordSetResultOutput{})
}
