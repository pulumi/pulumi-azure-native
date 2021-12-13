


package v20200601

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupPool(ctx *pulumi.Context, args *LookupPoolArgs, opts ...pulumi.InvokeOption) (*LookupPoolResult, error) {
	var rv LookupPoolResult
	err := ctx.Invoke("azure-native:netapp/v20200601:getPool", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupPoolArgs struct {
	AccountName       string `pulumi:"accountName"`
	PoolName          string `pulumi:"poolName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupPoolResult struct {
	Id                      string            `pulumi:"id"`
	Location                string            `pulumi:"location"`
	Name                    string            `pulumi:"name"`
	PoolId                  string            `pulumi:"poolId"`
	ProvisioningState       string            `pulumi:"provisioningState"`
	QosType                 *string           `pulumi:"qosType"`
	ServiceLevel            string            `pulumi:"serviceLevel"`
	Size                    float64           `pulumi:"size"`
	Tags                    map[string]string `pulumi:"tags"`
	TotalThroughputMibps    float64           `pulumi:"totalThroughputMibps"`
	Type                    string            `pulumi:"type"`
	UtilizedThroughputMibps float64           `pulumi:"utilizedThroughputMibps"`
}


func (val *LookupPoolResult) Defaults() *LookupPoolResult {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.QosType) {
		qosType_ := "Auto"
		tmp.QosType = &qosType_
	}
	if isZero(tmp.ServiceLevel) {
		tmp.ServiceLevel = "Premium"
	}
	return &tmp
}
