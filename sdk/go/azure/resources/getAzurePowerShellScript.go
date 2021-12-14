


package resources

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupAzurePowerShellScript(ctx *pulumi.Context, args *LookupAzurePowerShellScriptArgs, opts ...pulumi.InvokeOption) (*LookupAzurePowerShellScriptResult, error) {
	var rv LookupAzurePowerShellScriptResult
	err := ctx.Invoke("azure-native:resources:getAzurePowerShellScript", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupAzurePowerShellScriptArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ScriptName        string `pulumi:"scriptName"`
}


type LookupAzurePowerShellScriptResult struct {
	Arguments              *string                              `pulumi:"arguments"`
	AzPowerShellVersion    string                               `pulumi:"azPowerShellVersion"`
	CleanupPreference      *string                              `pulumi:"cleanupPreference"`
	ContainerSettings      *ContainerConfigurationResponse      `pulumi:"containerSettings"`
	EnvironmentVariables   []EnvironmentVariableResponse        `pulumi:"environmentVariables"`
	ForceUpdateTag         *string                              `pulumi:"forceUpdateTag"`
	Id                     string                               `pulumi:"id"`
	Identity               *ManagedServiceIdentityResponse      `pulumi:"identity"`
	Kind                   string                               `pulumi:"kind"`
	Location               string                               `pulumi:"location"`
	Name                   string                               `pulumi:"name"`
	Outputs                map[string]interface{}               `pulumi:"outputs"`
	PrimaryScriptUri       *string                              `pulumi:"primaryScriptUri"`
	ProvisioningState      string                               `pulumi:"provisioningState"`
	RetentionInterval      string                               `pulumi:"retentionInterval"`
	ScriptContent          *string                              `pulumi:"scriptContent"`
	Status                 ScriptStatusResponse                 `pulumi:"status"`
	StorageAccountSettings *StorageAccountConfigurationResponse `pulumi:"storageAccountSettings"`
	SupportingScriptUris   []string                             `pulumi:"supportingScriptUris"`
	SystemData             SystemDataResponse                   `pulumi:"systemData"`
	Tags                   map[string]string                    `pulumi:"tags"`
	Timeout                *string                              `pulumi:"timeout"`
	Type                   string                               `pulumi:"type"`
}


func (val *LookupAzurePowerShellScriptResult) Defaults() *LookupAzurePowerShellScriptResult {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.CleanupPreference) {
		cleanupPreference_ := "Always"
		tmp.CleanupPreference = &cleanupPreference_
	}
	if isZero(tmp.Timeout) {
		timeout_ := "P1D"
		tmp.Timeout = &timeout_
	}
	return &tmp
}
