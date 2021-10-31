


package v20210301preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupTemplateSpecVersion(ctx *pulumi.Context, args *LookupTemplateSpecVersionArgs, opts ...pulumi.InvokeOption) (*LookupTemplateSpecVersionResult, error) {
	var rv LookupTemplateSpecVersionResult
	err := ctx.Invoke("azure-native:resources/v20210301preview:getTemplateSpecVersion", args, &rv, opts...)
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
	Description      *string                          `pulumi:"description"`
	Id               string                           `pulumi:"id"`
	LinkedTemplates  []LinkedTemplateArtifactResponse `pulumi:"linkedTemplates"`
	Location         string                           `pulumi:"location"`
	MainTemplate     interface{}                      `pulumi:"mainTemplate"`
	Metadata         interface{}                      `pulumi:"metadata"`
	Name             string                           `pulumi:"name"`
	SystemData       SystemDataResponse               `pulumi:"systemData"`
	Tags             map[string]string                `pulumi:"tags"`
	Type             string                           `pulumi:"type"`
	UiFormDefinition interface{}                      `pulumi:"uiFormDefinition"`
}
