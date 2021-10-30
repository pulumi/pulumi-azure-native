


package deviceupdate

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupAccount(ctx *pulumi.Context, args *LookupAccountArgs, opts ...pulumi.InvokeOption) (*LookupAccountResult, error) {
	var rv LookupAccountResult
	err := ctx.Invoke("azure-native:deviceupdate:getAccount", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupAccountArgs struct {
	AccountName       string `pulumi:"accountName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupAccountResult struct {
	HostName            string                          `pulumi:"hostName"`
	Id                  string                          `pulumi:"id"`
	Identity            *ManagedServiceIdentityResponse `pulumi:"identity"`
	Location            string                          `pulumi:"location"`
	Name                string                          `pulumi:"name"`
	ProvisioningState   string                          `pulumi:"provisioningState"`
	PublicNetworkAccess *string                         `pulumi:"publicNetworkAccess"`
	SystemData          SystemDataResponse              `pulumi:"systemData"`
	Tags                map[string]string               `pulumi:"tags"`
	Type                string                          `pulumi:"type"`
}
