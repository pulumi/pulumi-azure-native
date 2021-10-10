


package v20210501

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupTemplateSpec(ctx *pulumi.Context, args *LookupTemplateSpecArgs, opts ...pulumi.InvokeOption) (*LookupTemplateSpecResult, error) {
	var rv LookupTemplateSpecResult
	err := ctx.Invoke("azure-native:resources/v20210501:getTemplateSpec", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupTemplateSpecArgs struct {
	Expand            *string `pulumi:"expand"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
	TemplateSpecName  string  `pulumi:"templateSpecName"`
}


type LookupTemplateSpecResult struct {
	Description *string                                    `pulumi:"description"`
	DisplayName *string                                    `pulumi:"displayName"`
	Id          string                                     `pulumi:"id"`
	Location    string                                     `pulumi:"location"`
	Metadata    interface{}                                `pulumi:"metadata"`
	Name        string                                     `pulumi:"name"`
	SystemData  SystemDataResponse                         `pulumi:"systemData"`
	Tags        map[string]string                          `pulumi:"tags"`
	Type        string                                     `pulumi:"type"`
	Versions    map[string]TemplateSpecVersionInfoResponse `pulumi:"versions"`
}
