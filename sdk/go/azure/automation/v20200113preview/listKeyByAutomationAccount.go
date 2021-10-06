


package v20200113preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListKeyByAutomationAccount(ctx *pulumi.Context, args *ListKeyByAutomationAccountArgs, opts ...pulumi.InvokeOption) (*ListKeyByAutomationAccountResult, error) {
	var rv ListKeyByAutomationAccountResult
	err := ctx.Invoke("azure-native:automation/v20200113preview:listKeyByAutomationAccount", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListKeyByAutomationAccountArgs struct {
	AutomationAccountName string `pulumi:"automationAccountName"`
	ResourceGroupName     string `pulumi:"resourceGroupName"`
}

type ListKeyByAutomationAccountResult struct {
	Keys []KeyResponse `pulumi:"keys"`
}
