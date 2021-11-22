


package v20210801

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupWorkbook(ctx *pulumi.Context, args *LookupWorkbookArgs, opts ...pulumi.InvokeOption) (*LookupWorkbookResult, error) {
	var rv LookupWorkbookResult
	err := ctx.Invoke("azure-native:insights/v20210801:getWorkbook", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupWorkbookArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ResourceName      string `pulumi:"resourceName"`
}


type LookupWorkbookResult struct {
	Category       string                            `pulumi:"category"`
	Description    *string                           `pulumi:"description"`
	DisplayName    string                            `pulumi:"displayName"`
	Etag           map[string]string                 `pulumi:"etag"`
	Id             string                            `pulumi:"id"`
	Identity       *WorkbookResourceResponseIdentity `pulumi:"identity"`
	Kind           *string                           `pulumi:"kind"`
	Location       string                            `pulumi:"location"`
	Name           string                            `pulumi:"name"`
	Revision       string                            `pulumi:"revision"`
	SerializedData string                            `pulumi:"serializedData"`
	SourceId       *string                           `pulumi:"sourceId"`
	StorageUri     *string                           `pulumi:"storageUri"`
	SystemData     SystemDataResponse                `pulumi:"systemData"`
	Tags           map[string]string                 `pulumi:"tags"`
	TimeModified   string                            `pulumi:"timeModified"`
	Type           string                            `pulumi:"type"`
	UserId         string                            `pulumi:"userId"`
	Version        *string                           `pulumi:"version"`
}
