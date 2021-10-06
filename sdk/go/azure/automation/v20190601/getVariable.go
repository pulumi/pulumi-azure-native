


package v20190601

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupVariable(ctx *pulumi.Context, args *LookupVariableArgs, opts ...pulumi.InvokeOption) (*LookupVariableResult, error) {
	var rv LookupVariableResult
	err := ctx.Invoke("azure-native:automation/v20190601:getVariable", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupVariableArgs struct {
	AutomationAccountName string `pulumi:"automationAccountName"`
	ResourceGroupName     string `pulumi:"resourceGroupName"`
	VariableName          string `pulumi:"variableName"`
}


type LookupVariableResult struct {
	CreationTime     *string `pulumi:"creationTime"`
	Description      *string `pulumi:"description"`
	Id               string  `pulumi:"id"`
	IsEncrypted      *bool   `pulumi:"isEncrypted"`
	LastModifiedTime *string `pulumi:"lastModifiedTime"`
	Name             string  `pulumi:"name"`
	Type             string  `pulumi:"type"`
	Value            *string `pulumi:"value"`
}
