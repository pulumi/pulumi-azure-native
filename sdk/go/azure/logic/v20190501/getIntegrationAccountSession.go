


package v20190501

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupIntegrationAccountSession(ctx *pulumi.Context, args *LookupIntegrationAccountSessionArgs, opts ...pulumi.InvokeOption) (*LookupIntegrationAccountSessionResult, error) {
	var rv LookupIntegrationAccountSessionResult
	err := ctx.Invoke("azure-native:logic/v20190501:getIntegrationAccountSession", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupIntegrationAccountSessionArgs struct {
	IntegrationAccountName string `pulumi:"integrationAccountName"`
	ResourceGroupName      string `pulumi:"resourceGroupName"`
	SessionName            string `pulumi:"sessionName"`
}


type LookupIntegrationAccountSessionResult struct {
	ChangedTime string            `pulumi:"changedTime"`
	Content     interface{}       `pulumi:"content"`
	CreatedTime string            `pulumi:"createdTime"`
	Id          string            `pulumi:"id"`
	Location    *string           `pulumi:"location"`
	Name        string            `pulumi:"name"`
	Tags        map[string]string `pulumi:"tags"`
	Type        string            `pulumi:"type"`
}
