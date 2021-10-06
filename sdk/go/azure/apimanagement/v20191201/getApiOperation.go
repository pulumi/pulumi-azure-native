


package v20191201

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupApiOperation(ctx *pulumi.Context, args *LookupApiOperationArgs, opts ...pulumi.InvokeOption) (*LookupApiOperationResult, error) {
	var rv LookupApiOperationResult
	err := ctx.Invoke("azure-native:apimanagement/v20191201:getApiOperation", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupApiOperationArgs struct {
	ApiId             string `pulumi:"apiId"`
	OperationId       string `pulumi:"operationId"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ServiceName       string `pulumi:"serviceName"`
}


type LookupApiOperationResult struct {
	Description        *string                     `pulumi:"description"`
	DisplayName        string                      `pulumi:"displayName"`
	Id                 string                      `pulumi:"id"`
	Method             string                      `pulumi:"method"`
	Name               string                      `pulumi:"name"`
	Policies           *string                     `pulumi:"policies"`
	Request            *RequestContractResponse    `pulumi:"request"`
	Responses          []ResponseContractResponse  `pulumi:"responses"`
	TemplateParameters []ParameterContractResponse `pulumi:"templateParameters"`
	Type               string                      `pulumi:"type"`
	UrlTemplate        string                      `pulumi:"urlTemplate"`
}
