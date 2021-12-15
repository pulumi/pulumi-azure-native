


package v20200801preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupFailoverGroup(ctx *pulumi.Context, args *LookupFailoverGroupArgs, opts ...pulumi.InvokeOption) (*LookupFailoverGroupResult, error) {
	var rv LookupFailoverGroupResult
	err := ctx.Invoke("azure-native:sql/v20200801preview:getFailoverGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupFailoverGroupArgs struct {
	FailoverGroupName string `pulumi:"failoverGroupName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ServerName        string `pulumi:"serverName"`
}


type LookupFailoverGroupResult struct {
	Databases         []string                               `pulumi:"databases"`
	Id                string                                 `pulumi:"id"`
	Location          string                                 `pulumi:"location"`
	Name              string                                 `pulumi:"name"`
	PartnerServers    []PartnerInfoResponse                  `pulumi:"partnerServers"`
	ReadOnlyEndpoint  *FailoverGroupReadOnlyEndpointResponse `pulumi:"readOnlyEndpoint"`
	ReadWriteEndpoint FailoverGroupReadWriteEndpointResponse `pulumi:"readWriteEndpoint"`
	ReplicationRole   string                                 `pulumi:"replicationRole"`
	ReplicationState  string                                 `pulumi:"replicationState"`
	Tags              map[string]string                      `pulumi:"tags"`
	Type              string                                 `pulumi:"type"`
}
