


package v20150504preview

import (
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
