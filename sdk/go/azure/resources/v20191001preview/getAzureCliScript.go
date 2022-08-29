


package v20191001preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupAzureCliScript(ctx *pulumi.Context, args *LookupAzureCliScriptArgs, opts ...pulumi.InvokeOption) (*LookupAzureCliScriptResult, error) {
	var rv LookupAzureCliScriptResult
	err := ctx.Invoke("azure-native:resources/v20191001preview:getAzureCliScript", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupAzureCliScriptArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ScriptName        string `pulumi:"scriptName"`
}


type LookupAzureCliScriptResult struct {
	Arguments              *string                              `pulumi:"arguments"`
	AzCliVersion           string                               `pulumi:"azCliVersion"`
	CleanupPreference      *string                              `pulumi:"cleanupPreference"`
	ContainerSettings      *ContainerConfigurationResponse      `pulumi:"containerSettings"`
	EnvironmentVariables   []EnvironmentVariableResponse        `pulumi:"environmentVariables"`
	ForceUpdateTag         *string                              `pulumi:"forceUpdateTag"`
	Id                     string                               `pulumi:"id"`
	Identity               ManagedServiceIdentityResponse       `pulumi:"identity"`
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


func (val *LookupAzureCliScriptResult) Defaults() *LookupAzureCliScriptResult {
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

func LookupAzureCliScriptOutput(ctx *pulumi.Context, args LookupAzureCliScriptOutputArgs, opts ...pulumi.InvokeOption) LookupAzureCliScriptResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupAzureCliScriptResult, error) {
			args := v.(LookupAzureCliScriptArgs)
			r, err := LookupAzureCliScript(ctx, &args, opts...)
			var s LookupAzureCliScriptResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupAzureCliScriptResultOutput)
}

type LookupAzureCliScriptOutputArgs struct {
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	ScriptName        pulumi.StringInput `pulumi:"scriptName"`
}

func (LookupAzureCliScriptOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAzureCliScriptArgs)(nil)).Elem()
}


type LookupAzureCliScriptResultOutput struct{ *pulumi.OutputState }

func (LookupAzureCliScriptResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAzureCliScriptResult)(nil)).Elem()
}

func (o LookupAzureCliScriptResultOutput) ToLookupAzureCliScriptResultOutput() LookupAzureCliScriptResultOutput {
	return o
}

func (o LookupAzureCliScriptResultOutput) ToLookupAzureCliScriptResultOutputWithContext(ctx context.Context) LookupAzureCliScriptResultOutput {
	return o
}

func (o LookupAzureCliScriptResultOutput) Arguments() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAzureCliScriptResult) *string { return v.Arguments }).(pulumi.StringPtrOutput)
}

func (o LookupAzureCliScriptResultOutput) AzCliVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAzureCliScriptResult) string { return v.AzCliVersion }).(pulumi.StringOutput)
}

func (o LookupAzureCliScriptResultOutput) CleanupPreference() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAzureCliScriptResult) *string { return v.CleanupPreference }).(pulumi.StringPtrOutput)
}

func (o LookupAzureCliScriptResultOutput) ContainerSettings() ContainerConfigurationResponsePtrOutput {
	return o.ApplyT(func(v LookupAzureCliScriptResult) *ContainerConfigurationResponse { return v.ContainerSettings }).(ContainerConfigurationResponsePtrOutput)
}

func (o LookupAzureCliScriptResultOutput) EnvironmentVariables() EnvironmentVariableResponseArrayOutput {
	return o.ApplyT(func(v LookupAzureCliScriptResult) []EnvironmentVariableResponse { return v.EnvironmentVariables }).(EnvironmentVariableResponseArrayOutput)
}

func (o LookupAzureCliScriptResultOutput) ForceUpdateTag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAzureCliScriptResult) *string { return v.ForceUpdateTag }).(pulumi.StringPtrOutput)
}

func (o LookupAzureCliScriptResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAzureCliScriptResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupAzureCliScriptResultOutput) Identity() ManagedServiceIdentityResponseOutput {
	return o.ApplyT(func(v LookupAzureCliScriptResult) ManagedServiceIdentityResponse { return v.Identity }).(ManagedServiceIdentityResponseOutput)
}

func (o LookupAzureCliScriptResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAzureCliScriptResult) string { return v.Kind }).(pulumi.StringOutput)
}

func (o LookupAzureCliScriptResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAzureCliScriptResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupAzureCliScriptResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAzureCliScriptResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupAzureCliScriptResultOutput) Outputs() pulumi.MapOutput {
	return o.ApplyT(func(v LookupAzureCliScriptResult) map[string]interface{} { return v.Outputs }).(pulumi.MapOutput)
}

func (o LookupAzureCliScriptResultOutput) PrimaryScriptUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAzureCliScriptResult) *string { return v.PrimaryScriptUri }).(pulumi.StringPtrOutput)
}

func (o LookupAzureCliScriptResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAzureCliScriptResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupAzureCliScriptResultOutput) RetentionInterval() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAzureCliScriptResult) string { return v.RetentionInterval }).(pulumi.StringOutput)
}

func (o LookupAzureCliScriptResultOutput) ScriptContent() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAzureCliScriptResult) *string { return v.ScriptContent }).(pulumi.StringPtrOutput)
}

func (o LookupAzureCliScriptResultOutput) Status() ScriptStatusResponseOutput {
	return o.ApplyT(func(v LookupAzureCliScriptResult) ScriptStatusResponse { return v.Status }).(ScriptStatusResponseOutput)
}

func (o LookupAzureCliScriptResultOutput) StorageAccountSettings() StorageAccountConfigurationResponsePtrOutput {
	return o.ApplyT(func(v LookupAzureCliScriptResult) *StorageAccountConfigurationResponse {
		return v.StorageAccountSettings
	}).(StorageAccountConfigurationResponsePtrOutput)
}

func (o LookupAzureCliScriptResultOutput) SupportingScriptUris() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupAzureCliScriptResult) []string { return v.SupportingScriptUris }).(pulumi.StringArrayOutput)
}

func (o LookupAzureCliScriptResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupAzureCliScriptResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupAzureCliScriptResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupAzureCliScriptResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupAzureCliScriptResultOutput) Timeout() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAzureCliScriptResult) *string { return v.Timeout }).(pulumi.StringPtrOutput)
}

func (o LookupAzureCliScriptResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAzureCliScriptResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupAzureCliScriptResultOutput{})
}
