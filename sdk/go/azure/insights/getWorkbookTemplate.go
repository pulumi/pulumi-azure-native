


package insights

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupWorkbookTemplate(ctx *pulumi.Context, args *LookupWorkbookTemplateArgs, opts ...pulumi.InvokeOption) (*LookupWorkbookTemplateResult, error) {
	var rv LookupWorkbookTemplateResult
	err := ctx.Invoke("azure-native:insights:getWorkbookTemplate", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupWorkbookTemplateArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ResourceName      string `pulumi:"resourceName"`
}


type LookupWorkbookTemplateResult struct {
	Author       *string                                               `pulumi:"author"`
	Galleries    []WorkbookTemplateGalleryResponse                     `pulumi:"galleries"`
	Id           string                                                `pulumi:"id"`
	Localized    map[string][]WorkbookTemplateLocalizedGalleryResponse `pulumi:"localized"`
	Location     string                                                `pulumi:"location"`
	Name         string                                                `pulumi:"name"`
	Priority     *int                                                  `pulumi:"priority"`
	Tags         map[string]string                                     `pulumi:"tags"`
	TemplateData interface{}                                           `pulumi:"templateData"`
	Type         string                                                `pulumi:"type"`
}
