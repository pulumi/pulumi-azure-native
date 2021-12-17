


package v20150501

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupWorkbook(ctx *pulumi.Context, args *LookupWorkbookArgs, opts ...pulumi.InvokeOption) (*LookupWorkbookResult, error) {
	var rv LookupWorkbookResult
	err := ctx.Invoke("azure-native:insights/v20150501:getWorkbook", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupWorkbookArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ResourceName      string `pulumi:"resourceName"`
}


type LookupWorkbookResult struct {
	Category         string            `pulumi:"category"`
	Id               string            `pulumi:"id"`
	Kind             *string           `pulumi:"kind"`
	Location         *string           `pulumi:"location"`
	Name             string            `pulumi:"name"`
	SerializedData   string            `pulumi:"serializedData"`
	SharedTypeKind   string            `pulumi:"sharedTypeKind"`
	SourceResourceId *string           `pulumi:"sourceResourceId"`
	Tags             map[string]string `pulumi:"tags"`
	TimeModified     string            `pulumi:"timeModified"`
	Type             string            `pulumi:"type"`
	UserId           string            `pulumi:"userId"`
	Version          *string           `pulumi:"version"`
	WorkbookId       string            `pulumi:"workbookId"`
}


func (val *LookupWorkbookResult) Defaults() *LookupWorkbookResult {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.SharedTypeKind) {
		tmp.SharedTypeKind = "shared"
	}
	return &tmp
}
