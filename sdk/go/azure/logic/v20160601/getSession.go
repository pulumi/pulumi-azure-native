


package v20160601

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupSession(ctx *pulumi.Context, args *LookupSessionArgs, opts ...pulumi.InvokeOption) (*LookupSessionResult, error) {
	var rv LookupSessionResult
	err := ctx.Invoke("azure-native:logic/v20160601:getSession", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSessionArgs struct {
	IntegrationAccountName string `pulumi:"integrationAccountName"`
	ResourceGroupName      string `pulumi:"resourceGroupName"`
	SessionName            string `pulumi:"sessionName"`
}


type LookupSessionResult struct {
	ChangedTime string            `pulumi:"changedTime"`
	Content     interface{}       `pulumi:"content"`
	CreatedTime string            `pulumi:"createdTime"`
	Id          string            `pulumi:"id"`
	Location    *string           `pulumi:"location"`
	Name        string            `pulumi:"name"`
	Tags        map[string]string `pulumi:"tags"`
	Type        string            `pulumi:"type"`
}
