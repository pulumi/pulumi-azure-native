


package confidentialledger

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupLedger(ctx *pulumi.Context, args *LookupLedgerArgs, opts ...pulumi.InvokeOption) (*LookupLedgerResult, error) {
	var rv LookupLedgerResult
	err := ctx.Invoke("azure-native:confidentialledger:getLedger", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupLedgerArgs struct {
	LedgerName        string `pulumi:"ledgerName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupLedgerResult struct {
	Id         string                   `pulumi:"id"`
	Location   *string                  `pulumi:"location"`
	Name       string                   `pulumi:"name"`
	Properties LedgerPropertiesResponse `pulumi:"properties"`
	SystemData SystemDataResponse       `pulumi:"systemData"`
	Tags       map[string]string        `pulumi:"tags"`
	Type       string                   `pulumi:"type"`
}
