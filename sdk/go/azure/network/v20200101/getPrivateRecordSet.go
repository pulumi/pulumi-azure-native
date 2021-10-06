


package v20200101

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupPrivateRecordSet(ctx *pulumi.Context, args *LookupPrivateRecordSetArgs, opts ...pulumi.InvokeOption) (*LookupPrivateRecordSetResult, error) {
	var rv LookupPrivateRecordSetResult
	err := ctx.Invoke("azure-native:network/v20200101:getPrivateRecordSet", args, &rv, opts...)
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
