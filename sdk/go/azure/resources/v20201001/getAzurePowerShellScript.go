


package v20201001

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupAzurePowerShellScript(ctx *pulumi.Context, args *LookupAzurePowerShellScriptArgs, opts ...pulumi.InvokeOption) (*LookupAzurePowerShellScriptResult, error) {
	var rv LookupAzurePowerShellScriptResult
	err := ctx.Invoke("azure-native:resources/v20201001:getAzurePowerShellScript", args, &rv, opts...)
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

func LookupAzurePowerShellScriptOutput(ctx *pulumi.Context, args LookupAzurePowerShellScriptOutputArgs, opts ...pulumi.InvokeOption) LookupAzurePowerShellScriptResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupAzurePowerShellScriptResult, error) {
			args := v.(LookupAzurePowerShellScriptArgs)
			r, err := LookupAzurePowerShellScript(ctx, &args, opts...)
			var s LookupAzurePowerShellScriptResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupAzurePowerShellScriptResultOutput)
}

type LookupAzurePowerShellScriptOutputArgs struct {
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	ScriptName        pulumi.StringInput `pulumi:"scriptName"`
}

func (LookupAzurePowerShellScriptOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAzurePowerShellScriptArgs)(nil)).Elem()
}


type LookupAzurePowerShellScriptResultOutput struct{ *pulumi.OutputState }

func (LookupAzurePowerShellScriptResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAzurePowerShellScriptResult)(nil)).Elem()
}

func (o LookupAzurePowerShellScriptResultOutput) ToLookupAzurePowerShellScriptResultOutput() LookupAzurePowerShellScriptResultOutput {
	return o
}

func (o LookupAzurePowerShellScriptResultOutput) ToLookupAzurePowerShellScriptResultOutputWithContext(ctx context.Context) LookupAzurePowerShellScriptResultOutput {
	return o
}

func (o LookupAzurePowerShellScriptResultOutput) Arguments() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAzurePowerShellScriptResult) *string { return v.Arguments }).(pulumi.StringPtrOutput)
}

func (o LookupAzurePowerShellScriptResultOutput) AzPowerShellVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAzurePowerShellScriptResult) string { return v.AzPowerShellVersion }).(pulumi.StringOutput)
}

func (o LookupAzurePowerShellScriptResultOutput) CleanupPreference() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAzurePowerShellScriptResult) *string { return v.CleanupPreference }).(pulumi.StringPtrOutput)
}

func (o LookupAzurePowerShellScriptResultOutput) ContainerSettings() ContainerConfigurationResponsePtrOutput {
	return o.ApplyT(func(v LookupAzurePowerShellScriptResult) *ContainerConfigurationResponse { return v.ContainerSettings }).(ContainerConfigurationResponsePtrOutput)
}

func (o LookupAzurePowerShellScriptResultOutput) EnvironmentVariables() EnvironmentVariableResponseArrayOutput {
	return o.ApplyT(func(v LookupAzurePowerShellScriptResult) []EnvironmentVariableResponse { return v.EnvironmentVariables }).(EnvironmentVariableResponseArrayOutput)
}

func (o LookupAzurePowerShellScriptResultOutput) ForceUpdateTag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAzurePowerShellScriptResult) *string { return v.ForceUpdateTag }).(pulumi.StringPtrOutput)
}

func (o LookupAzurePowerShellScriptResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAzurePowerShellScriptResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupAzurePowerShellScriptResultOutput) Identity() ManagedServiceIdentityResponsePtrOutput {
	return o.ApplyT(func(v LookupAzurePowerShellScriptResult) *ManagedServiceIdentityResponse { return v.Identity }).(ManagedServiceIdentityResponsePtrOutput)
}

func (o LookupAzurePowerShellScriptResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAzurePowerShellScriptResult) string { return v.Kind }).(pulumi.StringOutput)
}

func (o LookupAzurePowerShellScriptResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAzurePowerShellScriptResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupAzurePowerShellScriptResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAzurePowerShellScriptResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupAzurePowerShellScriptResultOutput) Outputs() pulumi.MapOutput {
	return o.ApplyT(func(v LookupAzurePowerShellScriptResult) map[string]interface{} { return v.Outputs }).(pulumi.MapOutput)
}

func (o LookupAzurePowerShellScriptResultOutput) PrimaryScriptUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAzurePowerShellScriptResult) *string { return v.PrimaryScriptUri }).(pulumi.StringPtrOutput)
}

func (o LookupAzurePowerShellScriptResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAzurePowerShellScriptResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupAzurePowerShellScriptResultOutput) RetentionInterval() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAzurePowerShellScriptResult) string { return v.RetentionInterval }).(pulumi.StringOutput)
}

func (o LookupAzurePowerShellScriptResultOutput) ScriptContent() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAzurePowerShellScriptResult) *string { return v.ScriptContent }).(pulumi.StringPtrOutput)
}

func (o LookupAzurePowerShellScriptResultOutput) Status() ScriptStatusResponseOutput {
	return o.ApplyT(func(v LookupAzurePowerShellScriptResult) ScriptStatusResponse { return v.Status }).(ScriptStatusResponseOutput)
}

func (o LookupAzurePowerShellScriptResultOutput) StorageAccountSettings() StorageAccountConfigurationResponsePtrOutput {
	return o.ApplyT(func(v LookupAzurePowerShellScriptResult) *StorageAccountConfigurationResponse {
		return v.StorageAccountSettings
	}).(StorageAccountConfigurationResponsePtrOutput)
}

func (o LookupAzurePowerShellScriptResultOutput) SupportingScriptUris() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupAzurePowerShellScriptResult) []string { return v.SupportingScriptUris }).(pulumi.StringArrayOutput)
}

func (o LookupAzurePowerShellScriptResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupAzurePowerShellScriptResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupAzurePowerShellScriptResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupAzurePowerShellScriptResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupAzurePowerShellScriptResultOutput) Timeout() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAzurePowerShellScriptResult) *string { return v.Timeout }).(pulumi.StringPtrOutput)
}

func (o LookupAzurePowerShellScriptResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAzurePowerShellScriptResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupAzurePowerShellScriptResultOutput{})
}
