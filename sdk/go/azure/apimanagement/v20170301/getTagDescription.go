


package v20170301

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupTagDescription(ctx *pulumi.Context, args *LookupTagDescriptionArgs, opts ...pulumi.InvokeOption) (*LookupTagDescriptionResult, error) {
	var rv LookupTagDescriptionResult
	err := ctx.Invoke("azure-native:apimanagement/v20170301:getTagDescription", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupTagDescriptionArgs struct {
	ApiId             string `pulumi:"apiId"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ServiceName       string `pulumi:"serviceName"`
	TagId             string `pulumi:"tagId"`
}


type LookupTagDescriptionResult struct {
	Description             *string `pulumi:"description"`
	DisplayName             *string `pulumi:"displayName"`
	ExternalDocsDescription *string `pulumi:"externalDocsDescription"`
	ExternalDocsUrl         *string `pulumi:"externalDocsUrl"`
	Id                      string  `pulumi:"id"`
	Name                    string  `pulumi:"name"`
	Type                    string  `pulumi:"type"`
}
