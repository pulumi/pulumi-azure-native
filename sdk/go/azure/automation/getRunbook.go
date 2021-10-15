


package automation

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupRunbook(ctx *pulumi.Context, args *LookupRunbookArgs, opts ...pulumi.InvokeOption) (*LookupRunbookResult, error) {
	var rv LookupRunbookResult
	err := ctx.Invoke("azure-native:automation:getRunbook", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupRunbookArgs struct {
	AutomationAccountName string `pulumi:"automationAccountName"`
	ResourceGroupName     string `pulumi:"resourceGroupName"`
	RunbookName           string `pulumi:"runbookName"`
}


type LookupRunbookResult struct {
	CreationTime       *string                             `pulumi:"creationTime"`
	Description        *string                             `pulumi:"description"`
	Draft              *RunbookDraftResponse               `pulumi:"draft"`
	Etag               *string                             `pulumi:"etag"`
	Id                 string                              `pulumi:"id"`
	JobCount           *int                                `pulumi:"jobCount"`
	LastModifiedBy     *string                             `pulumi:"lastModifiedBy"`
	LastModifiedTime   *string                             `pulumi:"lastModifiedTime"`
	Location           *string                             `pulumi:"location"`
	LogActivityTrace   *int                                `pulumi:"logActivityTrace"`
	LogProgress        *bool                               `pulumi:"logProgress"`
	LogVerbose         *bool                               `pulumi:"logVerbose"`
	Name               string                              `pulumi:"name"`
	OutputTypes        []string                            `pulumi:"outputTypes"`
	Parameters         map[string]RunbookParameterResponse `pulumi:"parameters"`
	ProvisioningState  *string                             `pulumi:"provisioningState"`
	PublishContentLink *ContentLinkResponse                `pulumi:"publishContentLink"`
	RunbookType        *string                             `pulumi:"runbookType"`
	State              *string                             `pulumi:"state"`
	Tags               map[string]string                   `pulumi:"tags"`
	Type               string                              `pulumi:"type"`
}
