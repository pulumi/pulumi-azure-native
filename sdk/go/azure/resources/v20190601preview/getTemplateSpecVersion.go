


package v20190601preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupTemplateSpecVersion(ctx *pulumi.Context, args *LookupTemplateSpecVersionArgs, opts ...pulumi.InvokeOption) (*LookupTemplateSpecVersionResult, error) {
	var rv LookupTemplateSpecVersionResult
	err := ctx.Invoke("azure-native:resources/v20190601preview:getTemplateSpecVersion", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupTemplateSpecVersionArgs struct {
	ResourceGroupName   string `pulumi:"resourceGroupName"`
	TemplateSpecName    string `pulumi:"templateSpecName"`
	TemplateSpecVersion string `pulumi:"templateSpecVersion"`
}


type LookupTemplateSpecVersionResult struct {
	Artifacts   []TemplateSpecTemplateArtifactResponse `pulumi:"artifacts"`
	Description *string                                `pulumi:"description"`
	Id          string                                 `pulumi:"id"`
	Location    string                                 `pulumi:"location"`
	Name        string                                 `pulumi:"name"`
	SystemData  SystemDataResponse                     `pulumi:"systemData"`
	Tags        map[string]string                      `pulumi:"tags"`
	Template    interface{}                            `pulumi:"template"`
	Type        string                                 `pulumi:"type"`
}
