


package v20210827

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupScript(ctx *pulumi.Context, args *LookupScriptArgs, opts ...pulumi.InvokeOption) (*LookupScriptResult, error) {
	var rv LookupScriptResult
	err := ctx.Invoke("azure-native:kusto/v20210827:getScript", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupScriptArgs struct {
	ClusterName       string `pulumi:"clusterName"`
	DatabaseName      string `pulumi:"databaseName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ScriptName        string `pulumi:"scriptName"`
}


type LookupScriptResult struct {
	ContinueOnErrors  *bool              `pulumi:"continueOnErrors"`
	ForceUpdateTag    *string            `pulumi:"forceUpdateTag"`
	Id                string             `pulumi:"id"`
	Name              string             `pulumi:"name"`
	ProvisioningState string             `pulumi:"provisioningState"`
	ScriptUrl         string             `pulumi:"scriptUrl"`
	SystemData        SystemDataResponse `pulumi:"systemData"`
	Type              string             `pulumi:"type"`
}


func (val *LookupScriptResult) Defaults() *LookupScriptResult {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.ContinueOnErrors) {
		continueOnErrors_ := false
		tmp.ContinueOnErrors = &continueOnErrors_
	}
	return &tmp
}
