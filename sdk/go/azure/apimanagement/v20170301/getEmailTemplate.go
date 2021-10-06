


package v20170301

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupEmailTemplate(ctx *pulumi.Context, args *LookupEmailTemplateArgs, opts ...pulumi.InvokeOption) (*LookupEmailTemplateResult, error) {
	var rv LookupEmailTemplateResult
	err := ctx.Invoke("azure-native:apimanagement/v20170301:getEmailTemplate", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupEmailTemplateArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ServiceName       string `pulumi:"serviceName"`
	TemplateName      string `pulumi:"templateName"`
}


type LookupEmailTemplateResult struct {
	Body        string                                              `pulumi:"body"`
	Description *string                                             `pulumi:"description"`
	Id          string                                              `pulumi:"id"`
	IsDefault   bool                                                `pulumi:"isDefault"`
	Name        string                                              `pulumi:"name"`
	Parameters  []EmailTemplateParametersContractPropertiesResponse `pulumi:"parameters"`
	Subject     string                                              `pulumi:"subject"`
	Title       *string                                             `pulumi:"title"`
	Type        string                                              `pulumi:"type"`
}
