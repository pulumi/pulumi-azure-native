


package v20180501

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupRecordSet(ctx *pulumi.Context, args *LookupRecordSetArgs, opts ...pulumi.InvokeOption) (*LookupRecordSetResult, error) {
	var rv LookupRecordSetResult
	err := ctx.Invoke("azure-native:network/v20180501:getRecordSet", args, &rv, opts...)
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
	ARecords          []ARecordResponse    `pulumi:"aRecords"`
	AaaaRecords       []AaaaRecordResponse `pulumi:"aaaaRecords"`
	CaaRecords        []CaaRecordResponse  `pulumi:"caaRecords"`
	CnameRecord       *CnameRecordResponse `pulumi:"cnameRecord"`
	Etag              *string              `pulumi:"etag"`
	Fqdn              string               `pulumi:"fqdn"`
	Id                string               `pulumi:"id"`
	Metadata          map[string]string    `pulumi:"metadata"`
	MxRecords         []MxRecordResponse   `pulumi:"mxRecords"`
	Name              string               `pulumi:"name"`
	NsRecords         []NsRecordResponse   `pulumi:"nsRecords"`
	ProvisioningState string               `pulumi:"provisioningState"`
	PtrRecords        []PtrRecordResponse  `pulumi:"ptrRecords"`
	SoaRecord         *SoaRecordResponse   `pulumi:"soaRecord"`
	SrvRecords        []SrvRecordResponse  `pulumi:"srvRecords"`
	TargetResource    *SubResourceResponse `pulumi:"targetResource"`
	Ttl               *float64             `pulumi:"ttl"`
	TxtRecords        []TxtRecordResponse  `pulumi:"txtRecords"`
	Type              string               `pulumi:"type"`
}
