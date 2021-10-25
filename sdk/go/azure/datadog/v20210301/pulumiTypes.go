


package v20210301

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type DatadogApiKeyResponse struct {
	Created   *string `pulumi:"created"`
	CreatedBy *string `pulumi:"createdBy"`
	Key       string  `pulumi:"key"`
	Name      *string `pulumi:"name"`
}





type DatadogApiKeyResponseInput interface {
	pulumi.Input

	ToDatadogApiKeyResponseOutput() DatadogApiKeyResponseOutput
	ToDatadogApiKeyResponseOutputWithContext(context.Context) DatadogApiKeyResponseOutput
}

type DatadogApiKeyResponseArgs struct {
	Created   pulumi.StringPtrInput `pulumi:"created"`
	CreatedBy pulumi.StringPtrInput `pulumi:"createdBy"`
	Key       pulumi.StringInput    `pulumi:"key"`
	Name      pulumi.StringPtrInput `pulumi:"name"`
}

func (DatadogApiKeyResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DatadogApiKeyResponse)(nil)).Elem()
}

func (i DatadogApiKeyResponseArgs) ToDatadogApiKeyResponseOutput() DatadogApiKeyResponseOutput {
	return i.ToDatadogApiKeyResponseOutputWithContext(context.Background())
}

func (i DatadogApiKeyResponseArgs) ToDatadogApiKeyResponseOutputWithContext(ctx context.Context) DatadogApiKeyResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatadogApiKeyResponseOutput)
}





type DatadogApiKeyResponseArrayInput interface {
	pulumi.Input

	ToDatadogApiKeyResponseArrayOutput() DatadogApiKeyResponseArrayOutput
	ToDatadogApiKeyResponseArrayOutputWithContext(context.Context) DatadogApiKeyResponseArrayOutput
}

type DatadogApiKeyResponseArray []DatadogApiKeyResponseInput

func (DatadogApiKeyResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DatadogApiKeyResponse)(nil)).Elem()
}

func (i DatadogApiKeyResponseArray) ToDatadogApiKeyResponseArrayOutput() DatadogApiKeyResponseArrayOutput {
	return i.ToDatadogApiKeyResponseArrayOutputWithContext(context.Background())
}

func (i DatadogApiKeyResponseArray) ToDatadogApiKeyResponseArrayOutputWithContext(ctx context.Context) DatadogApiKeyResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatadogApiKeyResponseArrayOutput)
}

type DatadogApiKeyResponseOutput struct{ *pulumi.OutputState }

func (DatadogApiKeyResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DatadogApiKeyResponse)(nil)).Elem()
}

func (o DatadogApiKeyResponseOutput) ToDatadogApiKeyResponseOutput() DatadogApiKeyResponseOutput {
	return o
}

func (o DatadogApiKeyResponseOutput) ToDatadogApiKeyResponseOutputWithContext(ctx context.Context) DatadogApiKeyResponseOutput {
	return o
}

func (o DatadogApiKeyResponseOutput) Created() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DatadogApiKeyResponse) *string { return v.Created }).(pulumi.StringPtrOutput)
}

func (o DatadogApiKeyResponseOutput) CreatedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DatadogApiKeyResponse) *string { return v.CreatedBy }).(pulumi.StringPtrOutput)
}

func (o DatadogApiKeyResponseOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v DatadogApiKeyResponse) string { return v.Key }).(pulumi.StringOutput)
}

func (o DatadogApiKeyResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DatadogApiKeyResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

type DatadogApiKeyResponseArrayOutput struct{ *pulumi.OutputState }

func (DatadogApiKeyResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DatadogApiKeyResponse)(nil)).Elem()
}

func (o DatadogApiKeyResponseArrayOutput) ToDatadogApiKeyResponseArrayOutput() DatadogApiKeyResponseArrayOutput {
	return o
}

func (o DatadogApiKeyResponseArrayOutput) ToDatadogApiKeyResponseArrayOutputWithContext(ctx context.Context) DatadogApiKeyResponseArrayOutput {
	return o
}

func (o DatadogApiKeyResponseArrayOutput) Index(i pulumi.IntInput) DatadogApiKeyResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) DatadogApiKeyResponse {
		return vs[0].([]DatadogApiKeyResponse)[vs[1].(int)]
	}).(DatadogApiKeyResponseOutput)
}

type DatadogHostMetadataResponse struct {
	AgentVersion  *string                       `pulumi:"agentVersion"`
	InstallMethod *DatadogInstallMethodResponse `pulumi:"installMethod"`
	LogsAgent     *DatadogLogsAgentResponse     `pulumi:"logsAgent"`
}





type DatadogHostMetadataResponseInput interface {
	pulumi.Input

	ToDatadogHostMetadataResponseOutput() DatadogHostMetadataResponseOutput
	ToDatadogHostMetadataResponseOutputWithContext(context.Context) DatadogHostMetadataResponseOutput
}

type DatadogHostMetadataResponseArgs struct {
	AgentVersion  pulumi.StringPtrInput                `pulumi:"agentVersion"`
	InstallMethod DatadogInstallMethodResponsePtrInput `pulumi:"installMethod"`
	LogsAgent     DatadogLogsAgentResponsePtrInput     `pulumi:"logsAgent"`
}

func (DatadogHostMetadataResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DatadogHostMetadataResponse)(nil)).Elem()
}

func (i DatadogHostMetadataResponseArgs) ToDatadogHostMetadataResponseOutput() DatadogHostMetadataResponseOutput {
	return i.ToDatadogHostMetadataResponseOutputWithContext(context.Background())
}

func (i DatadogHostMetadataResponseArgs) ToDatadogHostMetadataResponseOutputWithContext(ctx context.Context) DatadogHostMetadataResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatadogHostMetadataResponseOutput)
}

func (i DatadogHostMetadataResponseArgs) ToDatadogHostMetadataResponsePtrOutput() DatadogHostMetadataResponsePtrOutput {
	return i.ToDatadogHostMetadataResponsePtrOutputWithContext(context.Background())
}

func (i DatadogHostMetadataResponseArgs) ToDatadogHostMetadataResponsePtrOutputWithContext(ctx context.Context) DatadogHostMetadataResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatadogHostMetadataResponseOutput).ToDatadogHostMetadataResponsePtrOutputWithContext(ctx)
}









type DatadogHostMetadataResponsePtrInput interface {
	pulumi.Input

	ToDatadogHostMetadataResponsePtrOutput() DatadogHostMetadataResponsePtrOutput
	ToDatadogHostMetadataResponsePtrOutputWithContext(context.Context) DatadogHostMetadataResponsePtrOutput
}

type datadogHostMetadataResponsePtrType DatadogHostMetadataResponseArgs

func DatadogHostMetadataResponsePtr(v *DatadogHostMetadataResponseArgs) DatadogHostMetadataResponsePtrInput {
	return (*datadogHostMetadataResponsePtrType)(v)
}

func (*datadogHostMetadataResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DatadogHostMetadataResponse)(nil)).Elem()
}

func (i *datadogHostMetadataResponsePtrType) ToDatadogHostMetadataResponsePtrOutput() DatadogHostMetadataResponsePtrOutput {
	return i.ToDatadogHostMetadataResponsePtrOutputWithContext(context.Background())
}

func (i *datadogHostMetadataResponsePtrType) ToDatadogHostMetadataResponsePtrOutputWithContext(ctx context.Context) DatadogHostMetadataResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatadogHostMetadataResponsePtrOutput)
}

type DatadogHostMetadataResponseOutput struct{ *pulumi.OutputState }

func (DatadogHostMetadataResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DatadogHostMetadataResponse)(nil)).Elem()
}

func (o DatadogHostMetadataResponseOutput) ToDatadogHostMetadataResponseOutput() DatadogHostMetadataResponseOutput {
	return o
}

func (o DatadogHostMetadataResponseOutput) ToDatadogHostMetadataResponseOutputWithContext(ctx context.Context) DatadogHostMetadataResponseOutput {
	return o
}

func (o DatadogHostMetadataResponseOutput) ToDatadogHostMetadataResponsePtrOutput() DatadogHostMetadataResponsePtrOutput {
	return o.ToDatadogHostMetadataResponsePtrOutputWithContext(context.Background())
}

func (o DatadogHostMetadataResponseOutput) ToDatadogHostMetadataResponsePtrOutputWithContext(ctx context.Context) DatadogHostMetadataResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DatadogHostMetadataResponse) *DatadogHostMetadataResponse {
		return &v
	}).(DatadogHostMetadataResponsePtrOutput)
}

func (o DatadogHostMetadataResponseOutput) AgentVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DatadogHostMetadataResponse) *string { return v.AgentVersion }).(pulumi.StringPtrOutput)
}

func (o DatadogHostMetadataResponseOutput) InstallMethod() DatadogInstallMethodResponsePtrOutput {
	return o.ApplyT(func(v DatadogHostMetadataResponse) *DatadogInstallMethodResponse { return v.InstallMethod }).(DatadogInstallMethodResponsePtrOutput)
}

func (o DatadogHostMetadataResponseOutput) LogsAgent() DatadogLogsAgentResponsePtrOutput {
	return o.ApplyT(func(v DatadogHostMetadataResponse) *DatadogLogsAgentResponse { return v.LogsAgent }).(DatadogLogsAgentResponsePtrOutput)
}

type DatadogHostMetadataResponsePtrOutput struct{ *pulumi.OutputState }

func (DatadogHostMetadataResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DatadogHostMetadataResponse)(nil)).Elem()
}

func (o DatadogHostMetadataResponsePtrOutput) ToDatadogHostMetadataResponsePtrOutput() DatadogHostMetadataResponsePtrOutput {
	return o
}

func (o DatadogHostMetadataResponsePtrOutput) ToDatadogHostMetadataResponsePtrOutputWithContext(ctx context.Context) DatadogHostMetadataResponsePtrOutput {
	return o
}

func (o DatadogHostMetadataResponsePtrOutput) Elem() DatadogHostMetadataResponseOutput {
	return o.ApplyT(func(v *DatadogHostMetadataResponse) DatadogHostMetadataResponse {
		if v != nil {
			return *v
		}
		var ret DatadogHostMetadataResponse
		return ret
	}).(DatadogHostMetadataResponseOutput)
}

func (o DatadogHostMetadataResponsePtrOutput) AgentVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DatadogHostMetadataResponse) *string {
		if v == nil {
			return nil
		}
		return v.AgentVersion
	}).(pulumi.StringPtrOutput)
}

func (o DatadogHostMetadataResponsePtrOutput) InstallMethod() DatadogInstallMethodResponsePtrOutput {
	return o.ApplyT(func(v *DatadogHostMetadataResponse) *DatadogInstallMethodResponse {
		if v == nil {
			return nil
		}
		return v.InstallMethod
	}).(DatadogInstallMethodResponsePtrOutput)
}

func (o DatadogHostMetadataResponsePtrOutput) LogsAgent() DatadogLogsAgentResponsePtrOutput {
	return o.ApplyT(func(v *DatadogHostMetadataResponse) *DatadogLogsAgentResponse {
		if v == nil {
			return nil
		}
		return v.LogsAgent
	}).(DatadogLogsAgentResponsePtrOutput)
}

type DatadogHostResponse struct {
	Aliases []string                     `pulumi:"aliases"`
	Apps    []string                     `pulumi:"apps"`
	Meta    *DatadogHostMetadataResponse `pulumi:"meta"`
	Name    *string                      `pulumi:"name"`
}





type DatadogHostResponseInput interface {
	pulumi.Input

	ToDatadogHostResponseOutput() DatadogHostResponseOutput
	ToDatadogHostResponseOutputWithContext(context.Context) DatadogHostResponseOutput
}

type DatadogHostResponseArgs struct {
	Aliases pulumi.StringArrayInput             `pulumi:"aliases"`
	Apps    pulumi.StringArrayInput             `pulumi:"apps"`
	Meta    DatadogHostMetadataResponsePtrInput `pulumi:"meta"`
	Name    pulumi.StringPtrInput               `pulumi:"name"`
}

func (DatadogHostResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DatadogHostResponse)(nil)).Elem()
}

func (i DatadogHostResponseArgs) ToDatadogHostResponseOutput() DatadogHostResponseOutput {
	return i.ToDatadogHostResponseOutputWithContext(context.Background())
}

func (i DatadogHostResponseArgs) ToDatadogHostResponseOutputWithContext(ctx context.Context) DatadogHostResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatadogHostResponseOutput)
}





type DatadogHostResponseArrayInput interface {
	pulumi.Input

	ToDatadogHostResponseArrayOutput() DatadogHostResponseArrayOutput
	ToDatadogHostResponseArrayOutputWithContext(context.Context) DatadogHostResponseArrayOutput
}

type DatadogHostResponseArray []DatadogHostResponseInput

func (DatadogHostResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DatadogHostResponse)(nil)).Elem()
}

func (i DatadogHostResponseArray) ToDatadogHostResponseArrayOutput() DatadogHostResponseArrayOutput {
	return i.ToDatadogHostResponseArrayOutputWithContext(context.Background())
}

func (i DatadogHostResponseArray) ToDatadogHostResponseArrayOutputWithContext(ctx context.Context) DatadogHostResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatadogHostResponseArrayOutput)
}

type DatadogHostResponseOutput struct{ *pulumi.OutputState }

func (DatadogHostResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DatadogHostResponse)(nil)).Elem()
}

func (o DatadogHostResponseOutput) ToDatadogHostResponseOutput() DatadogHostResponseOutput {
	return o
}

func (o DatadogHostResponseOutput) ToDatadogHostResponseOutputWithContext(ctx context.Context) DatadogHostResponseOutput {
	return o
}

func (o DatadogHostResponseOutput) Aliases() pulumi.StringArrayOutput {
	return o.ApplyT(func(v DatadogHostResponse) []string { return v.Aliases }).(pulumi.StringArrayOutput)
}

func (o DatadogHostResponseOutput) Apps() pulumi.StringArrayOutput {
	return o.ApplyT(func(v DatadogHostResponse) []string { return v.Apps }).(pulumi.StringArrayOutput)
}

func (o DatadogHostResponseOutput) Meta() DatadogHostMetadataResponsePtrOutput {
	return o.ApplyT(func(v DatadogHostResponse) *DatadogHostMetadataResponse { return v.Meta }).(DatadogHostMetadataResponsePtrOutput)
}

func (o DatadogHostResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DatadogHostResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

type DatadogHostResponseArrayOutput struct{ *pulumi.OutputState }

func (DatadogHostResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DatadogHostResponse)(nil)).Elem()
}

func (o DatadogHostResponseArrayOutput) ToDatadogHostResponseArrayOutput() DatadogHostResponseArrayOutput {
	return o
}

func (o DatadogHostResponseArrayOutput) ToDatadogHostResponseArrayOutputWithContext(ctx context.Context) DatadogHostResponseArrayOutput {
	return o
}

func (o DatadogHostResponseArrayOutput) Index(i pulumi.IntInput) DatadogHostResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) DatadogHostResponse {
		return vs[0].([]DatadogHostResponse)[vs[1].(int)]
	}).(DatadogHostResponseOutput)
}

type DatadogInstallMethodResponse struct {
	InstallerVersion *string `pulumi:"installerVersion"`
	Tool             *string `pulumi:"tool"`
	ToolVersion      *string `pulumi:"toolVersion"`
}





type DatadogInstallMethodResponseInput interface {
	pulumi.Input

	ToDatadogInstallMethodResponseOutput() DatadogInstallMethodResponseOutput
	ToDatadogInstallMethodResponseOutputWithContext(context.Context) DatadogInstallMethodResponseOutput
}

type DatadogInstallMethodResponseArgs struct {
	InstallerVersion pulumi.StringPtrInput `pulumi:"installerVersion"`
	Tool             pulumi.StringPtrInput `pulumi:"tool"`
	ToolVersion      pulumi.StringPtrInput `pulumi:"toolVersion"`
}

func (DatadogInstallMethodResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DatadogInstallMethodResponse)(nil)).Elem()
}

func (i DatadogInstallMethodResponseArgs) ToDatadogInstallMethodResponseOutput() DatadogInstallMethodResponseOutput {
	return i.ToDatadogInstallMethodResponseOutputWithContext(context.Background())
}

func (i DatadogInstallMethodResponseArgs) ToDatadogInstallMethodResponseOutputWithContext(ctx context.Context) DatadogInstallMethodResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatadogInstallMethodResponseOutput)
}

func (i DatadogInstallMethodResponseArgs) ToDatadogInstallMethodResponsePtrOutput() DatadogInstallMethodResponsePtrOutput {
	return i.ToDatadogInstallMethodResponsePtrOutputWithContext(context.Background())
}

func (i DatadogInstallMethodResponseArgs) ToDatadogInstallMethodResponsePtrOutputWithContext(ctx context.Context) DatadogInstallMethodResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatadogInstallMethodResponseOutput).ToDatadogInstallMethodResponsePtrOutputWithContext(ctx)
}









type DatadogInstallMethodResponsePtrInput interface {
	pulumi.Input

	ToDatadogInstallMethodResponsePtrOutput() DatadogInstallMethodResponsePtrOutput
	ToDatadogInstallMethodResponsePtrOutputWithContext(context.Context) DatadogInstallMethodResponsePtrOutput
}

type datadogInstallMethodResponsePtrType DatadogInstallMethodResponseArgs

func DatadogInstallMethodResponsePtr(v *DatadogInstallMethodResponseArgs) DatadogInstallMethodResponsePtrInput {
	return (*datadogInstallMethodResponsePtrType)(v)
}

func (*datadogInstallMethodResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DatadogInstallMethodResponse)(nil)).Elem()
}

func (i *datadogInstallMethodResponsePtrType) ToDatadogInstallMethodResponsePtrOutput() DatadogInstallMethodResponsePtrOutput {
	return i.ToDatadogInstallMethodResponsePtrOutputWithContext(context.Background())
}

func (i *datadogInstallMethodResponsePtrType) ToDatadogInstallMethodResponsePtrOutputWithContext(ctx context.Context) DatadogInstallMethodResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatadogInstallMethodResponsePtrOutput)
}

type DatadogInstallMethodResponseOutput struct{ *pulumi.OutputState }

func (DatadogInstallMethodResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DatadogInstallMethodResponse)(nil)).Elem()
}

func (o DatadogInstallMethodResponseOutput) ToDatadogInstallMethodResponseOutput() DatadogInstallMethodResponseOutput {
	return o
}

func (o DatadogInstallMethodResponseOutput) ToDatadogInstallMethodResponseOutputWithContext(ctx context.Context) DatadogInstallMethodResponseOutput {
	return o
}

func (o DatadogInstallMethodResponseOutput) ToDatadogInstallMethodResponsePtrOutput() DatadogInstallMethodResponsePtrOutput {
	return o.ToDatadogInstallMethodResponsePtrOutputWithContext(context.Background())
}

func (o DatadogInstallMethodResponseOutput) ToDatadogInstallMethodResponsePtrOutputWithContext(ctx context.Context) DatadogInstallMethodResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DatadogInstallMethodResponse) *DatadogInstallMethodResponse {
		return &v
	}).(DatadogInstallMethodResponsePtrOutput)
}

func (o DatadogInstallMethodResponseOutput) InstallerVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DatadogInstallMethodResponse) *string { return v.InstallerVersion }).(pulumi.StringPtrOutput)
}

func (o DatadogInstallMethodResponseOutput) Tool() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DatadogInstallMethodResponse) *string { return v.Tool }).(pulumi.StringPtrOutput)
}

func (o DatadogInstallMethodResponseOutput) ToolVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DatadogInstallMethodResponse) *string { return v.ToolVersion }).(pulumi.StringPtrOutput)
}

type DatadogInstallMethodResponsePtrOutput struct{ *pulumi.OutputState }

func (DatadogInstallMethodResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DatadogInstallMethodResponse)(nil)).Elem()
}

func (o DatadogInstallMethodResponsePtrOutput) ToDatadogInstallMethodResponsePtrOutput() DatadogInstallMethodResponsePtrOutput {
	return o
}

func (o DatadogInstallMethodResponsePtrOutput) ToDatadogInstallMethodResponsePtrOutputWithContext(ctx context.Context) DatadogInstallMethodResponsePtrOutput {
	return o
}

func (o DatadogInstallMethodResponsePtrOutput) Elem() DatadogInstallMethodResponseOutput {
	return o.ApplyT(func(v *DatadogInstallMethodResponse) DatadogInstallMethodResponse {
		if v != nil {
			return *v
		}
		var ret DatadogInstallMethodResponse
		return ret
	}).(DatadogInstallMethodResponseOutput)
}

func (o DatadogInstallMethodResponsePtrOutput) InstallerVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DatadogInstallMethodResponse) *string {
		if v == nil {
			return nil
		}
		return v.InstallerVersion
	}).(pulumi.StringPtrOutput)
}

func (o DatadogInstallMethodResponsePtrOutput) Tool() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DatadogInstallMethodResponse) *string {
		if v == nil {
			return nil
		}
		return v.Tool
	}).(pulumi.StringPtrOutput)
}

func (o DatadogInstallMethodResponsePtrOutput) ToolVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DatadogInstallMethodResponse) *string {
		if v == nil {
			return nil
		}
		return v.ToolVersion
	}).(pulumi.StringPtrOutput)
}

type DatadogLogsAgentResponse struct {
	Transport *string `pulumi:"transport"`
}





type DatadogLogsAgentResponseInput interface {
	pulumi.Input

	ToDatadogLogsAgentResponseOutput() DatadogLogsAgentResponseOutput
	ToDatadogLogsAgentResponseOutputWithContext(context.Context) DatadogLogsAgentResponseOutput
}

type DatadogLogsAgentResponseArgs struct {
	Transport pulumi.StringPtrInput `pulumi:"transport"`
}

func (DatadogLogsAgentResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DatadogLogsAgentResponse)(nil)).Elem()
}

func (i DatadogLogsAgentResponseArgs) ToDatadogLogsAgentResponseOutput() DatadogLogsAgentResponseOutput {
	return i.ToDatadogLogsAgentResponseOutputWithContext(context.Background())
}

func (i DatadogLogsAgentResponseArgs) ToDatadogLogsAgentResponseOutputWithContext(ctx context.Context) DatadogLogsAgentResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatadogLogsAgentResponseOutput)
}

func (i DatadogLogsAgentResponseArgs) ToDatadogLogsAgentResponsePtrOutput() DatadogLogsAgentResponsePtrOutput {
	return i.ToDatadogLogsAgentResponsePtrOutputWithContext(context.Background())
}

func (i DatadogLogsAgentResponseArgs) ToDatadogLogsAgentResponsePtrOutputWithContext(ctx context.Context) DatadogLogsAgentResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatadogLogsAgentResponseOutput).ToDatadogLogsAgentResponsePtrOutputWithContext(ctx)
}









type DatadogLogsAgentResponsePtrInput interface {
	pulumi.Input

	ToDatadogLogsAgentResponsePtrOutput() DatadogLogsAgentResponsePtrOutput
	ToDatadogLogsAgentResponsePtrOutputWithContext(context.Context) DatadogLogsAgentResponsePtrOutput
}

type datadogLogsAgentResponsePtrType DatadogLogsAgentResponseArgs

func DatadogLogsAgentResponsePtr(v *DatadogLogsAgentResponseArgs) DatadogLogsAgentResponsePtrInput {
	return (*datadogLogsAgentResponsePtrType)(v)
}

func (*datadogLogsAgentResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DatadogLogsAgentResponse)(nil)).Elem()
}

func (i *datadogLogsAgentResponsePtrType) ToDatadogLogsAgentResponsePtrOutput() DatadogLogsAgentResponsePtrOutput {
	return i.ToDatadogLogsAgentResponsePtrOutputWithContext(context.Background())
}

func (i *datadogLogsAgentResponsePtrType) ToDatadogLogsAgentResponsePtrOutputWithContext(ctx context.Context) DatadogLogsAgentResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatadogLogsAgentResponsePtrOutput)
}

type DatadogLogsAgentResponseOutput struct{ *pulumi.OutputState }

func (DatadogLogsAgentResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DatadogLogsAgentResponse)(nil)).Elem()
}

func (o DatadogLogsAgentResponseOutput) ToDatadogLogsAgentResponseOutput() DatadogLogsAgentResponseOutput {
	return o
}

func (o DatadogLogsAgentResponseOutput) ToDatadogLogsAgentResponseOutputWithContext(ctx context.Context) DatadogLogsAgentResponseOutput {
	return o
}

func (o DatadogLogsAgentResponseOutput) ToDatadogLogsAgentResponsePtrOutput() DatadogLogsAgentResponsePtrOutput {
	return o.ToDatadogLogsAgentResponsePtrOutputWithContext(context.Background())
}

func (o DatadogLogsAgentResponseOutput) ToDatadogLogsAgentResponsePtrOutputWithContext(ctx context.Context) DatadogLogsAgentResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DatadogLogsAgentResponse) *DatadogLogsAgentResponse {
		return &v
	}).(DatadogLogsAgentResponsePtrOutput)
}

func (o DatadogLogsAgentResponseOutput) Transport() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DatadogLogsAgentResponse) *string { return v.Transport }).(pulumi.StringPtrOutput)
}

type DatadogLogsAgentResponsePtrOutput struct{ *pulumi.OutputState }

func (DatadogLogsAgentResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DatadogLogsAgentResponse)(nil)).Elem()
}

func (o DatadogLogsAgentResponsePtrOutput) ToDatadogLogsAgentResponsePtrOutput() DatadogLogsAgentResponsePtrOutput {
	return o
}

func (o DatadogLogsAgentResponsePtrOutput) ToDatadogLogsAgentResponsePtrOutputWithContext(ctx context.Context) DatadogLogsAgentResponsePtrOutput {
	return o
}

func (o DatadogLogsAgentResponsePtrOutput) Elem() DatadogLogsAgentResponseOutput {
	return o.ApplyT(func(v *DatadogLogsAgentResponse) DatadogLogsAgentResponse {
		if v != nil {
			return *v
		}
		var ret DatadogLogsAgentResponse
		return ret
	}).(DatadogLogsAgentResponseOutput)
}

func (o DatadogLogsAgentResponsePtrOutput) Transport() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DatadogLogsAgentResponse) *string {
		if v == nil {
			return nil
		}
		return v.Transport
	}).(pulumi.StringPtrOutput)
}

type DatadogOrganizationProperties struct {
	ApiKey          *string `pulumi:"apiKey"`
	ApplicationKey  *string `pulumi:"applicationKey"`
	EnterpriseAppId *string `pulumi:"enterpriseAppId"`
	LinkingAuthCode *string `pulumi:"linkingAuthCode"`
	LinkingClientId *string `pulumi:"linkingClientId"`
	RedirectUri     *string `pulumi:"redirectUri"`
}





type DatadogOrganizationPropertiesInput interface {
	pulumi.Input

	ToDatadogOrganizationPropertiesOutput() DatadogOrganizationPropertiesOutput
	ToDatadogOrganizationPropertiesOutputWithContext(context.Context) DatadogOrganizationPropertiesOutput
}

type DatadogOrganizationPropertiesArgs struct {
	ApiKey          pulumi.StringPtrInput `pulumi:"apiKey"`
	ApplicationKey  pulumi.StringPtrInput `pulumi:"applicationKey"`
	EnterpriseAppId pulumi.StringPtrInput `pulumi:"enterpriseAppId"`
	LinkingAuthCode pulumi.StringPtrInput `pulumi:"linkingAuthCode"`
	LinkingClientId pulumi.StringPtrInput `pulumi:"linkingClientId"`
	RedirectUri     pulumi.StringPtrInput `pulumi:"redirectUri"`
}

func (DatadogOrganizationPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DatadogOrganizationProperties)(nil)).Elem()
}

func (i DatadogOrganizationPropertiesArgs) ToDatadogOrganizationPropertiesOutput() DatadogOrganizationPropertiesOutput {
	return i.ToDatadogOrganizationPropertiesOutputWithContext(context.Background())
}

func (i DatadogOrganizationPropertiesArgs) ToDatadogOrganizationPropertiesOutputWithContext(ctx context.Context) DatadogOrganizationPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatadogOrganizationPropertiesOutput)
}

func (i DatadogOrganizationPropertiesArgs) ToDatadogOrganizationPropertiesPtrOutput() DatadogOrganizationPropertiesPtrOutput {
	return i.ToDatadogOrganizationPropertiesPtrOutputWithContext(context.Background())
}

func (i DatadogOrganizationPropertiesArgs) ToDatadogOrganizationPropertiesPtrOutputWithContext(ctx context.Context) DatadogOrganizationPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatadogOrganizationPropertiesOutput).ToDatadogOrganizationPropertiesPtrOutputWithContext(ctx)
}









type DatadogOrganizationPropertiesPtrInput interface {
	pulumi.Input

	ToDatadogOrganizationPropertiesPtrOutput() DatadogOrganizationPropertiesPtrOutput
	ToDatadogOrganizationPropertiesPtrOutputWithContext(context.Context) DatadogOrganizationPropertiesPtrOutput
}

type datadogOrganizationPropertiesPtrType DatadogOrganizationPropertiesArgs

func DatadogOrganizationPropertiesPtr(v *DatadogOrganizationPropertiesArgs) DatadogOrganizationPropertiesPtrInput {
	return (*datadogOrganizationPropertiesPtrType)(v)
}

func (*datadogOrganizationPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DatadogOrganizationProperties)(nil)).Elem()
}

func (i *datadogOrganizationPropertiesPtrType) ToDatadogOrganizationPropertiesPtrOutput() DatadogOrganizationPropertiesPtrOutput {
	return i.ToDatadogOrganizationPropertiesPtrOutputWithContext(context.Background())
}

func (i *datadogOrganizationPropertiesPtrType) ToDatadogOrganizationPropertiesPtrOutputWithContext(ctx context.Context) DatadogOrganizationPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatadogOrganizationPropertiesPtrOutput)
}

type DatadogOrganizationPropertiesOutput struct{ *pulumi.OutputState }

func (DatadogOrganizationPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DatadogOrganizationProperties)(nil)).Elem()
}

func (o DatadogOrganizationPropertiesOutput) ToDatadogOrganizationPropertiesOutput() DatadogOrganizationPropertiesOutput {
	return o
}

func (o DatadogOrganizationPropertiesOutput) ToDatadogOrganizationPropertiesOutputWithContext(ctx context.Context) DatadogOrganizationPropertiesOutput {
	return o
}

func (o DatadogOrganizationPropertiesOutput) ToDatadogOrganizationPropertiesPtrOutput() DatadogOrganizationPropertiesPtrOutput {
	return o.ToDatadogOrganizationPropertiesPtrOutputWithContext(context.Background())
}

func (o DatadogOrganizationPropertiesOutput) ToDatadogOrganizationPropertiesPtrOutputWithContext(ctx context.Context) DatadogOrganizationPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DatadogOrganizationProperties) *DatadogOrganizationProperties {
		return &v
	}).(DatadogOrganizationPropertiesPtrOutput)
}

func (o DatadogOrganizationPropertiesOutput) ApiKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DatadogOrganizationProperties) *string { return v.ApiKey }).(pulumi.StringPtrOutput)
}

func (o DatadogOrganizationPropertiesOutput) ApplicationKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DatadogOrganizationProperties) *string { return v.ApplicationKey }).(pulumi.StringPtrOutput)
}

func (o DatadogOrganizationPropertiesOutput) EnterpriseAppId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DatadogOrganizationProperties) *string { return v.EnterpriseAppId }).(pulumi.StringPtrOutput)
}

func (o DatadogOrganizationPropertiesOutput) LinkingAuthCode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DatadogOrganizationProperties) *string { return v.LinkingAuthCode }).(pulumi.StringPtrOutput)
}

func (o DatadogOrganizationPropertiesOutput) LinkingClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DatadogOrganizationProperties) *string { return v.LinkingClientId }).(pulumi.StringPtrOutput)
}

func (o DatadogOrganizationPropertiesOutput) RedirectUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DatadogOrganizationProperties) *string { return v.RedirectUri }).(pulumi.StringPtrOutput)
}

type DatadogOrganizationPropertiesPtrOutput struct{ *pulumi.OutputState }

func (DatadogOrganizationPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DatadogOrganizationProperties)(nil)).Elem()
}

func (o DatadogOrganizationPropertiesPtrOutput) ToDatadogOrganizationPropertiesPtrOutput() DatadogOrganizationPropertiesPtrOutput {
	return o
}

func (o DatadogOrganizationPropertiesPtrOutput) ToDatadogOrganizationPropertiesPtrOutputWithContext(ctx context.Context) DatadogOrganizationPropertiesPtrOutput {
	return o
}

func (o DatadogOrganizationPropertiesPtrOutput) Elem() DatadogOrganizationPropertiesOutput {
	return o.ApplyT(func(v *DatadogOrganizationProperties) DatadogOrganizationProperties {
		if v != nil {
			return *v
		}
		var ret DatadogOrganizationProperties
		return ret
	}).(DatadogOrganizationPropertiesOutput)
}

func (o DatadogOrganizationPropertiesPtrOutput) ApiKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DatadogOrganizationProperties) *string {
		if v == nil {
			return nil
		}
		return v.ApiKey
	}).(pulumi.StringPtrOutput)
}

func (o DatadogOrganizationPropertiesPtrOutput) ApplicationKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DatadogOrganizationProperties) *string {
		if v == nil {
			return nil
		}
		return v.ApplicationKey
	}).(pulumi.StringPtrOutput)
}

func (o DatadogOrganizationPropertiesPtrOutput) EnterpriseAppId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DatadogOrganizationProperties) *string {
		if v == nil {
			return nil
		}
		return v.EnterpriseAppId
	}).(pulumi.StringPtrOutput)
}

func (o DatadogOrganizationPropertiesPtrOutput) LinkingAuthCode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DatadogOrganizationProperties) *string {
		if v == nil {
			return nil
		}
		return v.LinkingAuthCode
	}).(pulumi.StringPtrOutput)
}

func (o DatadogOrganizationPropertiesPtrOutput) LinkingClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DatadogOrganizationProperties) *string {
		if v == nil {
			return nil
		}
		return v.LinkingClientId
	}).(pulumi.StringPtrOutput)
}

func (o DatadogOrganizationPropertiesPtrOutput) RedirectUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DatadogOrganizationProperties) *string {
		if v == nil {
			return nil
		}
		return v.RedirectUri
	}).(pulumi.StringPtrOutput)
}

type DatadogOrganizationPropertiesResponse struct {
	Id   string `pulumi:"id"`
	Name string `pulumi:"name"`
}





type DatadogOrganizationPropertiesResponseInput interface {
	pulumi.Input

	ToDatadogOrganizationPropertiesResponseOutput() DatadogOrganizationPropertiesResponseOutput
	ToDatadogOrganizationPropertiesResponseOutputWithContext(context.Context) DatadogOrganizationPropertiesResponseOutput
}

type DatadogOrganizationPropertiesResponseArgs struct {
	Id   pulumi.StringInput `pulumi:"id"`
	Name pulumi.StringInput `pulumi:"name"`
}

func (DatadogOrganizationPropertiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DatadogOrganizationPropertiesResponse)(nil)).Elem()
}

func (i DatadogOrganizationPropertiesResponseArgs) ToDatadogOrganizationPropertiesResponseOutput() DatadogOrganizationPropertiesResponseOutput {
	return i.ToDatadogOrganizationPropertiesResponseOutputWithContext(context.Background())
}

func (i DatadogOrganizationPropertiesResponseArgs) ToDatadogOrganizationPropertiesResponseOutputWithContext(ctx context.Context) DatadogOrganizationPropertiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatadogOrganizationPropertiesResponseOutput)
}

func (i DatadogOrganizationPropertiesResponseArgs) ToDatadogOrganizationPropertiesResponsePtrOutput() DatadogOrganizationPropertiesResponsePtrOutput {
	return i.ToDatadogOrganizationPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i DatadogOrganizationPropertiesResponseArgs) ToDatadogOrganizationPropertiesResponsePtrOutputWithContext(ctx context.Context) DatadogOrganizationPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatadogOrganizationPropertiesResponseOutput).ToDatadogOrganizationPropertiesResponsePtrOutputWithContext(ctx)
}









type DatadogOrganizationPropertiesResponsePtrInput interface {
	pulumi.Input

	ToDatadogOrganizationPropertiesResponsePtrOutput() DatadogOrganizationPropertiesResponsePtrOutput
	ToDatadogOrganizationPropertiesResponsePtrOutputWithContext(context.Context) DatadogOrganizationPropertiesResponsePtrOutput
}

type datadogOrganizationPropertiesResponsePtrType DatadogOrganizationPropertiesResponseArgs

func DatadogOrganizationPropertiesResponsePtr(v *DatadogOrganizationPropertiesResponseArgs) DatadogOrganizationPropertiesResponsePtrInput {
	return (*datadogOrganizationPropertiesResponsePtrType)(v)
}

func (*datadogOrganizationPropertiesResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DatadogOrganizationPropertiesResponse)(nil)).Elem()
}

func (i *datadogOrganizationPropertiesResponsePtrType) ToDatadogOrganizationPropertiesResponsePtrOutput() DatadogOrganizationPropertiesResponsePtrOutput {
	return i.ToDatadogOrganizationPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i *datadogOrganizationPropertiesResponsePtrType) ToDatadogOrganizationPropertiesResponsePtrOutputWithContext(ctx context.Context) DatadogOrganizationPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatadogOrganizationPropertiesResponsePtrOutput)
}

type DatadogOrganizationPropertiesResponseOutput struct{ *pulumi.OutputState }

func (DatadogOrganizationPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DatadogOrganizationPropertiesResponse)(nil)).Elem()
}

func (o DatadogOrganizationPropertiesResponseOutput) ToDatadogOrganizationPropertiesResponseOutput() DatadogOrganizationPropertiesResponseOutput {
	return o
}

func (o DatadogOrganizationPropertiesResponseOutput) ToDatadogOrganizationPropertiesResponseOutputWithContext(ctx context.Context) DatadogOrganizationPropertiesResponseOutput {
	return o
}

func (o DatadogOrganizationPropertiesResponseOutput) ToDatadogOrganizationPropertiesResponsePtrOutput() DatadogOrganizationPropertiesResponsePtrOutput {
	return o.ToDatadogOrganizationPropertiesResponsePtrOutputWithContext(context.Background())
}

func (o DatadogOrganizationPropertiesResponseOutput) ToDatadogOrganizationPropertiesResponsePtrOutputWithContext(ctx context.Context) DatadogOrganizationPropertiesResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DatadogOrganizationPropertiesResponse) *DatadogOrganizationPropertiesResponse {
		return &v
	}).(DatadogOrganizationPropertiesResponsePtrOutput)
}

func (o DatadogOrganizationPropertiesResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v DatadogOrganizationPropertiesResponse) string { return v.Id }).(pulumi.StringOutput)
}

func (o DatadogOrganizationPropertiesResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v DatadogOrganizationPropertiesResponse) string { return v.Name }).(pulumi.StringOutput)
}

type DatadogOrganizationPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (DatadogOrganizationPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DatadogOrganizationPropertiesResponse)(nil)).Elem()
}

func (o DatadogOrganizationPropertiesResponsePtrOutput) ToDatadogOrganizationPropertiesResponsePtrOutput() DatadogOrganizationPropertiesResponsePtrOutput {
	return o
}

func (o DatadogOrganizationPropertiesResponsePtrOutput) ToDatadogOrganizationPropertiesResponsePtrOutputWithContext(ctx context.Context) DatadogOrganizationPropertiesResponsePtrOutput {
	return o
}

func (o DatadogOrganizationPropertiesResponsePtrOutput) Elem() DatadogOrganizationPropertiesResponseOutput {
	return o.ApplyT(func(v *DatadogOrganizationPropertiesResponse) DatadogOrganizationPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret DatadogOrganizationPropertiesResponse
		return ret
	}).(DatadogOrganizationPropertiesResponseOutput)
}

func (o DatadogOrganizationPropertiesResponsePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DatadogOrganizationPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Id
	}).(pulumi.StringPtrOutput)
}

func (o DatadogOrganizationPropertiesResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DatadogOrganizationPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

type IdentityProperties struct {
	Type *string `pulumi:"type"`
}





type IdentityPropertiesInput interface {
	pulumi.Input

	ToIdentityPropertiesOutput() IdentityPropertiesOutput
	ToIdentityPropertiesOutputWithContext(context.Context) IdentityPropertiesOutput
}

type IdentityPropertiesArgs struct {
	Type pulumi.StringPtrInput `pulumi:"type"`
}

func (IdentityPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*IdentityProperties)(nil)).Elem()
}

func (i IdentityPropertiesArgs) ToIdentityPropertiesOutput() IdentityPropertiesOutput {
	return i.ToIdentityPropertiesOutputWithContext(context.Background())
}

func (i IdentityPropertiesArgs) ToIdentityPropertiesOutputWithContext(ctx context.Context) IdentityPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentityPropertiesOutput)
}

func (i IdentityPropertiesArgs) ToIdentityPropertiesPtrOutput() IdentityPropertiesPtrOutput {
	return i.ToIdentityPropertiesPtrOutputWithContext(context.Background())
}

func (i IdentityPropertiesArgs) ToIdentityPropertiesPtrOutputWithContext(ctx context.Context) IdentityPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentityPropertiesOutput).ToIdentityPropertiesPtrOutputWithContext(ctx)
}









type IdentityPropertiesPtrInput interface {
	pulumi.Input

	ToIdentityPropertiesPtrOutput() IdentityPropertiesPtrOutput
	ToIdentityPropertiesPtrOutputWithContext(context.Context) IdentityPropertiesPtrOutput
}

type identityPropertiesPtrType IdentityPropertiesArgs

func IdentityPropertiesPtr(v *IdentityPropertiesArgs) IdentityPropertiesPtrInput {
	return (*identityPropertiesPtrType)(v)
}

func (*identityPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**IdentityProperties)(nil)).Elem()
}

func (i *identityPropertiesPtrType) ToIdentityPropertiesPtrOutput() IdentityPropertiesPtrOutput {
	return i.ToIdentityPropertiesPtrOutputWithContext(context.Background())
}

func (i *identityPropertiesPtrType) ToIdentityPropertiesPtrOutputWithContext(ctx context.Context) IdentityPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentityPropertiesPtrOutput)
}

type IdentityPropertiesOutput struct{ *pulumi.OutputState }

func (IdentityPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IdentityProperties)(nil)).Elem()
}

func (o IdentityPropertiesOutput) ToIdentityPropertiesOutput() IdentityPropertiesOutput {
	return o
}

func (o IdentityPropertiesOutput) ToIdentityPropertiesOutputWithContext(ctx context.Context) IdentityPropertiesOutput {
	return o
}

func (o IdentityPropertiesOutput) ToIdentityPropertiesPtrOutput() IdentityPropertiesPtrOutput {
	return o.ToIdentityPropertiesPtrOutputWithContext(context.Background())
}

func (o IdentityPropertiesOutput) ToIdentityPropertiesPtrOutputWithContext(ctx context.Context) IdentityPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v IdentityProperties) *IdentityProperties {
		return &v
	}).(IdentityPropertiesPtrOutput)
}

func (o IdentityPropertiesOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IdentityProperties) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type IdentityPropertiesPtrOutput struct{ *pulumi.OutputState }

func (IdentityPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IdentityProperties)(nil)).Elem()
}

func (o IdentityPropertiesPtrOutput) ToIdentityPropertiesPtrOutput() IdentityPropertiesPtrOutput {
	return o
}

func (o IdentityPropertiesPtrOutput) ToIdentityPropertiesPtrOutputWithContext(ctx context.Context) IdentityPropertiesPtrOutput {
	return o
}

func (o IdentityPropertiesPtrOutput) Elem() IdentityPropertiesOutput {
	return o.ApplyT(func(v *IdentityProperties) IdentityProperties {
		if v != nil {
			return *v
		}
		var ret IdentityProperties
		return ret
	}).(IdentityPropertiesOutput)
}

func (o IdentityPropertiesPtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IdentityProperties) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

type IdentityPropertiesResponse struct {
	PrincipalId string  `pulumi:"principalId"`
	TenantId    string  `pulumi:"tenantId"`
	Type        *string `pulumi:"type"`
}





type IdentityPropertiesResponseInput interface {
	pulumi.Input

	ToIdentityPropertiesResponseOutput() IdentityPropertiesResponseOutput
	ToIdentityPropertiesResponseOutputWithContext(context.Context) IdentityPropertiesResponseOutput
}

type IdentityPropertiesResponseArgs struct {
	PrincipalId pulumi.StringInput    `pulumi:"principalId"`
	TenantId    pulumi.StringInput    `pulumi:"tenantId"`
	Type        pulumi.StringPtrInput `pulumi:"type"`
}

func (IdentityPropertiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*IdentityPropertiesResponse)(nil)).Elem()
}

func (i IdentityPropertiesResponseArgs) ToIdentityPropertiesResponseOutput() IdentityPropertiesResponseOutput {
	return i.ToIdentityPropertiesResponseOutputWithContext(context.Background())
}

func (i IdentityPropertiesResponseArgs) ToIdentityPropertiesResponseOutputWithContext(ctx context.Context) IdentityPropertiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentityPropertiesResponseOutput)
}

func (i IdentityPropertiesResponseArgs) ToIdentityPropertiesResponsePtrOutput() IdentityPropertiesResponsePtrOutput {
	return i.ToIdentityPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i IdentityPropertiesResponseArgs) ToIdentityPropertiesResponsePtrOutputWithContext(ctx context.Context) IdentityPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentityPropertiesResponseOutput).ToIdentityPropertiesResponsePtrOutputWithContext(ctx)
}









type IdentityPropertiesResponsePtrInput interface {
	pulumi.Input

	ToIdentityPropertiesResponsePtrOutput() IdentityPropertiesResponsePtrOutput
	ToIdentityPropertiesResponsePtrOutputWithContext(context.Context) IdentityPropertiesResponsePtrOutput
}

type identityPropertiesResponsePtrType IdentityPropertiesResponseArgs

func IdentityPropertiesResponsePtr(v *IdentityPropertiesResponseArgs) IdentityPropertiesResponsePtrInput {
	return (*identityPropertiesResponsePtrType)(v)
}

func (*identityPropertiesResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**IdentityPropertiesResponse)(nil)).Elem()
}

func (i *identityPropertiesResponsePtrType) ToIdentityPropertiesResponsePtrOutput() IdentityPropertiesResponsePtrOutput {
	return i.ToIdentityPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i *identityPropertiesResponsePtrType) ToIdentityPropertiesResponsePtrOutputWithContext(ctx context.Context) IdentityPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentityPropertiesResponsePtrOutput)
}

type IdentityPropertiesResponseOutput struct{ *pulumi.OutputState }

func (IdentityPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IdentityPropertiesResponse)(nil)).Elem()
}

func (o IdentityPropertiesResponseOutput) ToIdentityPropertiesResponseOutput() IdentityPropertiesResponseOutput {
	return o
}

func (o IdentityPropertiesResponseOutput) ToIdentityPropertiesResponseOutputWithContext(ctx context.Context) IdentityPropertiesResponseOutput {
	return o
}

func (o IdentityPropertiesResponseOutput) ToIdentityPropertiesResponsePtrOutput() IdentityPropertiesResponsePtrOutput {
	return o.ToIdentityPropertiesResponsePtrOutputWithContext(context.Background())
}

func (o IdentityPropertiesResponseOutput) ToIdentityPropertiesResponsePtrOutputWithContext(ctx context.Context) IdentityPropertiesResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v IdentityPropertiesResponse) *IdentityPropertiesResponse {
		return &v
	}).(IdentityPropertiesResponsePtrOutput)
}

func (o IdentityPropertiesResponseOutput) PrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v IdentityPropertiesResponse) string { return v.PrincipalId }).(pulumi.StringOutput)
}

func (o IdentityPropertiesResponseOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v IdentityPropertiesResponse) string { return v.TenantId }).(pulumi.StringOutput)
}

func (o IdentityPropertiesResponseOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IdentityPropertiesResponse) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type IdentityPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (IdentityPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IdentityPropertiesResponse)(nil)).Elem()
}

func (o IdentityPropertiesResponsePtrOutput) ToIdentityPropertiesResponsePtrOutput() IdentityPropertiesResponsePtrOutput {
	return o
}

func (o IdentityPropertiesResponsePtrOutput) ToIdentityPropertiesResponsePtrOutputWithContext(ctx context.Context) IdentityPropertiesResponsePtrOutput {
	return o
}

func (o IdentityPropertiesResponsePtrOutput) Elem() IdentityPropertiesResponseOutput {
	return o.ApplyT(func(v *IdentityPropertiesResponse) IdentityPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret IdentityPropertiesResponse
		return ret
	}).(IdentityPropertiesResponseOutput)
}

func (o IdentityPropertiesResponsePtrOutput) PrincipalId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IdentityPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.PrincipalId
	}).(pulumi.StringPtrOutput)
}

func (o IdentityPropertiesResponsePtrOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IdentityPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.TenantId
	}).(pulumi.StringPtrOutput)
}

func (o IdentityPropertiesResponsePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IdentityPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

type LinkedResourceResponse struct {
	Id *string `pulumi:"id"`
}





type LinkedResourceResponseInput interface {
	pulumi.Input

	ToLinkedResourceResponseOutput() LinkedResourceResponseOutput
	ToLinkedResourceResponseOutputWithContext(context.Context) LinkedResourceResponseOutput
}

type LinkedResourceResponseArgs struct {
	Id pulumi.StringPtrInput `pulumi:"id"`
}

func (LinkedResourceResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LinkedResourceResponse)(nil)).Elem()
}

func (i LinkedResourceResponseArgs) ToLinkedResourceResponseOutput() LinkedResourceResponseOutput {
	return i.ToLinkedResourceResponseOutputWithContext(context.Background())
}

func (i LinkedResourceResponseArgs) ToLinkedResourceResponseOutputWithContext(ctx context.Context) LinkedResourceResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LinkedResourceResponseOutput)
}





type LinkedResourceResponseArrayInput interface {
	pulumi.Input

	ToLinkedResourceResponseArrayOutput() LinkedResourceResponseArrayOutput
	ToLinkedResourceResponseArrayOutputWithContext(context.Context) LinkedResourceResponseArrayOutput
}

type LinkedResourceResponseArray []LinkedResourceResponseInput

func (LinkedResourceResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]LinkedResourceResponse)(nil)).Elem()
}

func (i LinkedResourceResponseArray) ToLinkedResourceResponseArrayOutput() LinkedResourceResponseArrayOutput {
	return i.ToLinkedResourceResponseArrayOutputWithContext(context.Background())
}

func (i LinkedResourceResponseArray) ToLinkedResourceResponseArrayOutputWithContext(ctx context.Context) LinkedResourceResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LinkedResourceResponseArrayOutput)
}

type LinkedResourceResponseOutput struct{ *pulumi.OutputState }

func (LinkedResourceResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LinkedResourceResponse)(nil)).Elem()
}

func (o LinkedResourceResponseOutput) ToLinkedResourceResponseOutput() LinkedResourceResponseOutput {
	return o
}

func (o LinkedResourceResponseOutput) ToLinkedResourceResponseOutputWithContext(ctx context.Context) LinkedResourceResponseOutput {
	return o
}

func (o LinkedResourceResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LinkedResourceResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

type LinkedResourceResponseArrayOutput struct{ *pulumi.OutputState }

func (LinkedResourceResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]LinkedResourceResponse)(nil)).Elem()
}

func (o LinkedResourceResponseArrayOutput) ToLinkedResourceResponseArrayOutput() LinkedResourceResponseArrayOutput {
	return o
}

func (o LinkedResourceResponseArrayOutput) ToLinkedResourceResponseArrayOutputWithContext(ctx context.Context) LinkedResourceResponseArrayOutput {
	return o
}

func (o LinkedResourceResponseArrayOutput) Index(i pulumi.IntInput) LinkedResourceResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) LinkedResourceResponse {
		return vs[0].([]LinkedResourceResponse)[vs[1].(int)]
	}).(LinkedResourceResponseOutput)
}

type MonitorProperties struct {
	DatadogOrganizationProperties *DatadogOrganizationProperties `pulumi:"datadogOrganizationProperties"`
	MonitoringStatus              *string                        `pulumi:"monitoringStatus"`
	UserInfo                      *UserInfo                      `pulumi:"userInfo"`
}





type MonitorPropertiesInput interface {
	pulumi.Input

	ToMonitorPropertiesOutput() MonitorPropertiesOutput
	ToMonitorPropertiesOutputWithContext(context.Context) MonitorPropertiesOutput
}

type MonitorPropertiesArgs struct {
	DatadogOrganizationProperties DatadogOrganizationPropertiesPtrInput `pulumi:"datadogOrganizationProperties"`
	MonitoringStatus              pulumi.StringPtrInput                 `pulumi:"monitoringStatus"`
	UserInfo                      UserInfoPtrInput                      `pulumi:"userInfo"`
}

func (MonitorPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MonitorProperties)(nil)).Elem()
}

func (i MonitorPropertiesArgs) ToMonitorPropertiesOutput() MonitorPropertiesOutput {
	return i.ToMonitorPropertiesOutputWithContext(context.Background())
}

func (i MonitorPropertiesArgs) ToMonitorPropertiesOutputWithContext(ctx context.Context) MonitorPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MonitorPropertiesOutput)
}

func (i MonitorPropertiesArgs) ToMonitorPropertiesPtrOutput() MonitorPropertiesPtrOutput {
	return i.ToMonitorPropertiesPtrOutputWithContext(context.Background())
}

func (i MonitorPropertiesArgs) ToMonitorPropertiesPtrOutputWithContext(ctx context.Context) MonitorPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MonitorPropertiesOutput).ToMonitorPropertiesPtrOutputWithContext(ctx)
}









type MonitorPropertiesPtrInput interface {
	pulumi.Input

	ToMonitorPropertiesPtrOutput() MonitorPropertiesPtrOutput
	ToMonitorPropertiesPtrOutputWithContext(context.Context) MonitorPropertiesPtrOutput
}

type monitorPropertiesPtrType MonitorPropertiesArgs

func MonitorPropertiesPtr(v *MonitorPropertiesArgs) MonitorPropertiesPtrInput {
	return (*monitorPropertiesPtrType)(v)
}

func (*monitorPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**MonitorProperties)(nil)).Elem()
}

func (i *monitorPropertiesPtrType) ToMonitorPropertiesPtrOutput() MonitorPropertiesPtrOutput {
	return i.ToMonitorPropertiesPtrOutputWithContext(context.Background())
}

func (i *monitorPropertiesPtrType) ToMonitorPropertiesPtrOutputWithContext(ctx context.Context) MonitorPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MonitorPropertiesPtrOutput)
}

type MonitorPropertiesOutput struct{ *pulumi.OutputState }

func (MonitorPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MonitorProperties)(nil)).Elem()
}

func (o MonitorPropertiesOutput) ToMonitorPropertiesOutput() MonitorPropertiesOutput {
	return o
}

func (o MonitorPropertiesOutput) ToMonitorPropertiesOutputWithContext(ctx context.Context) MonitorPropertiesOutput {
	return o
}

func (o MonitorPropertiesOutput) ToMonitorPropertiesPtrOutput() MonitorPropertiesPtrOutput {
	return o.ToMonitorPropertiesPtrOutputWithContext(context.Background())
}

func (o MonitorPropertiesOutput) ToMonitorPropertiesPtrOutputWithContext(ctx context.Context) MonitorPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v MonitorProperties) *MonitorProperties {
		return &v
	}).(MonitorPropertiesPtrOutput)
}

func (o MonitorPropertiesOutput) DatadogOrganizationProperties() DatadogOrganizationPropertiesPtrOutput {
	return o.ApplyT(func(v MonitorProperties) *DatadogOrganizationProperties { return v.DatadogOrganizationProperties }).(DatadogOrganizationPropertiesPtrOutput)
}

func (o MonitorPropertiesOutput) MonitoringStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MonitorProperties) *string { return v.MonitoringStatus }).(pulumi.StringPtrOutput)
}

func (o MonitorPropertiesOutput) UserInfo() UserInfoPtrOutput {
	return o.ApplyT(func(v MonitorProperties) *UserInfo { return v.UserInfo }).(UserInfoPtrOutput)
}

type MonitorPropertiesPtrOutput struct{ *pulumi.OutputState }

func (MonitorPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MonitorProperties)(nil)).Elem()
}

func (o MonitorPropertiesPtrOutput) ToMonitorPropertiesPtrOutput() MonitorPropertiesPtrOutput {
	return o
}

func (o MonitorPropertiesPtrOutput) ToMonitorPropertiesPtrOutputWithContext(ctx context.Context) MonitorPropertiesPtrOutput {
	return o
}

func (o MonitorPropertiesPtrOutput) Elem() MonitorPropertiesOutput {
	return o.ApplyT(func(v *MonitorProperties) MonitorProperties {
		if v != nil {
			return *v
		}
		var ret MonitorProperties
		return ret
	}).(MonitorPropertiesOutput)
}

func (o MonitorPropertiesPtrOutput) DatadogOrganizationProperties() DatadogOrganizationPropertiesPtrOutput {
	return o.ApplyT(func(v *MonitorProperties) *DatadogOrganizationProperties {
		if v == nil {
			return nil
		}
		return v.DatadogOrganizationProperties
	}).(DatadogOrganizationPropertiesPtrOutput)
}

func (o MonitorPropertiesPtrOutput) MonitoringStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MonitorProperties) *string {
		if v == nil {
			return nil
		}
		return v.MonitoringStatus
	}).(pulumi.StringPtrOutput)
}

func (o MonitorPropertiesPtrOutput) UserInfo() UserInfoPtrOutput {
	return o.ApplyT(func(v *MonitorProperties) *UserInfo {
		if v == nil {
			return nil
		}
		return v.UserInfo
	}).(UserInfoPtrOutput)
}

type MonitorPropertiesResponse struct {
	DatadogOrganizationProperties *DatadogOrganizationPropertiesResponse `pulumi:"datadogOrganizationProperties"`
	LiftrResourceCategory         string                                 `pulumi:"liftrResourceCategory"`
	LiftrResourcePreference       int                                    `pulumi:"liftrResourcePreference"`
	MarketplaceSubscriptionStatus string                                 `pulumi:"marketplaceSubscriptionStatus"`
	MonitoringStatus              *string                                `pulumi:"monitoringStatus"`
	ProvisioningState             string                                 `pulumi:"provisioningState"`
	UserInfo                      *UserInfoResponse                      `pulumi:"userInfo"`
}





type MonitorPropertiesResponseInput interface {
	pulumi.Input

	ToMonitorPropertiesResponseOutput() MonitorPropertiesResponseOutput
	ToMonitorPropertiesResponseOutputWithContext(context.Context) MonitorPropertiesResponseOutput
}

type MonitorPropertiesResponseArgs struct {
	DatadogOrganizationProperties DatadogOrganizationPropertiesResponsePtrInput `pulumi:"datadogOrganizationProperties"`
	LiftrResourceCategory         pulumi.StringInput                            `pulumi:"liftrResourceCategory"`
	LiftrResourcePreference       pulumi.IntInput                               `pulumi:"liftrResourcePreference"`
	MarketplaceSubscriptionStatus pulumi.StringInput                            `pulumi:"marketplaceSubscriptionStatus"`
	MonitoringStatus              pulumi.StringPtrInput                         `pulumi:"monitoringStatus"`
	ProvisioningState             pulumi.StringInput                            `pulumi:"provisioningState"`
	UserInfo                      UserInfoResponsePtrInput                      `pulumi:"userInfo"`
}

func (MonitorPropertiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MonitorPropertiesResponse)(nil)).Elem()
}

func (i MonitorPropertiesResponseArgs) ToMonitorPropertiesResponseOutput() MonitorPropertiesResponseOutput {
	return i.ToMonitorPropertiesResponseOutputWithContext(context.Background())
}

func (i MonitorPropertiesResponseArgs) ToMonitorPropertiesResponseOutputWithContext(ctx context.Context) MonitorPropertiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MonitorPropertiesResponseOutput)
}

func (i MonitorPropertiesResponseArgs) ToMonitorPropertiesResponsePtrOutput() MonitorPropertiesResponsePtrOutput {
	return i.ToMonitorPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i MonitorPropertiesResponseArgs) ToMonitorPropertiesResponsePtrOutputWithContext(ctx context.Context) MonitorPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MonitorPropertiesResponseOutput).ToMonitorPropertiesResponsePtrOutputWithContext(ctx)
}









type MonitorPropertiesResponsePtrInput interface {
	pulumi.Input

	ToMonitorPropertiesResponsePtrOutput() MonitorPropertiesResponsePtrOutput
	ToMonitorPropertiesResponsePtrOutputWithContext(context.Context) MonitorPropertiesResponsePtrOutput
}

type monitorPropertiesResponsePtrType MonitorPropertiesResponseArgs

func MonitorPropertiesResponsePtr(v *MonitorPropertiesResponseArgs) MonitorPropertiesResponsePtrInput {
	return (*monitorPropertiesResponsePtrType)(v)
}

func (*monitorPropertiesResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**MonitorPropertiesResponse)(nil)).Elem()
}

func (i *monitorPropertiesResponsePtrType) ToMonitorPropertiesResponsePtrOutput() MonitorPropertiesResponsePtrOutput {
	return i.ToMonitorPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i *monitorPropertiesResponsePtrType) ToMonitorPropertiesResponsePtrOutputWithContext(ctx context.Context) MonitorPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MonitorPropertiesResponsePtrOutput)
}

type MonitorPropertiesResponseOutput struct{ *pulumi.OutputState }

func (MonitorPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MonitorPropertiesResponse)(nil)).Elem()
}

func (o MonitorPropertiesResponseOutput) ToMonitorPropertiesResponseOutput() MonitorPropertiesResponseOutput {
	return o
}

func (o MonitorPropertiesResponseOutput) ToMonitorPropertiesResponseOutputWithContext(ctx context.Context) MonitorPropertiesResponseOutput {
	return o
}

func (o MonitorPropertiesResponseOutput) ToMonitorPropertiesResponsePtrOutput() MonitorPropertiesResponsePtrOutput {
	return o.ToMonitorPropertiesResponsePtrOutputWithContext(context.Background())
}

func (o MonitorPropertiesResponseOutput) ToMonitorPropertiesResponsePtrOutputWithContext(ctx context.Context) MonitorPropertiesResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v MonitorPropertiesResponse) *MonitorPropertiesResponse {
		return &v
	}).(MonitorPropertiesResponsePtrOutput)
}

func (o MonitorPropertiesResponseOutput) DatadogOrganizationProperties() DatadogOrganizationPropertiesResponsePtrOutput {
	return o.ApplyT(func(v MonitorPropertiesResponse) *DatadogOrganizationPropertiesResponse {
		return v.DatadogOrganizationProperties
	}).(DatadogOrganizationPropertiesResponsePtrOutput)
}

func (o MonitorPropertiesResponseOutput) LiftrResourceCategory() pulumi.StringOutput {
	return o.ApplyT(func(v MonitorPropertiesResponse) string { return v.LiftrResourceCategory }).(pulumi.StringOutput)
}

func (o MonitorPropertiesResponseOutput) LiftrResourcePreference() pulumi.IntOutput {
	return o.ApplyT(func(v MonitorPropertiesResponse) int { return v.LiftrResourcePreference }).(pulumi.IntOutput)
}

func (o MonitorPropertiesResponseOutput) MarketplaceSubscriptionStatus() pulumi.StringOutput {
	return o.ApplyT(func(v MonitorPropertiesResponse) string { return v.MarketplaceSubscriptionStatus }).(pulumi.StringOutput)
}

func (o MonitorPropertiesResponseOutput) MonitoringStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MonitorPropertiesResponse) *string { return v.MonitoringStatus }).(pulumi.StringPtrOutput)
}

func (o MonitorPropertiesResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v MonitorPropertiesResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o MonitorPropertiesResponseOutput) UserInfo() UserInfoResponsePtrOutput {
	return o.ApplyT(func(v MonitorPropertiesResponse) *UserInfoResponse { return v.UserInfo }).(UserInfoResponsePtrOutput)
}

type MonitorPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (MonitorPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MonitorPropertiesResponse)(nil)).Elem()
}

func (o MonitorPropertiesResponsePtrOutput) ToMonitorPropertiesResponsePtrOutput() MonitorPropertiesResponsePtrOutput {
	return o
}

func (o MonitorPropertiesResponsePtrOutput) ToMonitorPropertiesResponsePtrOutputWithContext(ctx context.Context) MonitorPropertiesResponsePtrOutput {
	return o
}

func (o MonitorPropertiesResponsePtrOutput) Elem() MonitorPropertiesResponseOutput {
	return o.ApplyT(func(v *MonitorPropertiesResponse) MonitorPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret MonitorPropertiesResponse
		return ret
	}).(MonitorPropertiesResponseOutput)
}

func (o MonitorPropertiesResponsePtrOutput) DatadogOrganizationProperties() DatadogOrganizationPropertiesResponsePtrOutput {
	return o.ApplyT(func(v *MonitorPropertiesResponse) *DatadogOrganizationPropertiesResponse {
		if v == nil {
			return nil
		}
		return v.DatadogOrganizationProperties
	}).(DatadogOrganizationPropertiesResponsePtrOutput)
}

func (o MonitorPropertiesResponsePtrOutput) LiftrResourceCategory() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MonitorPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.LiftrResourceCategory
	}).(pulumi.StringPtrOutput)
}

func (o MonitorPropertiesResponsePtrOutput) LiftrResourcePreference() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *MonitorPropertiesResponse) *int {
		if v == nil {
			return nil
		}
		return &v.LiftrResourcePreference
	}).(pulumi.IntPtrOutput)
}

func (o MonitorPropertiesResponsePtrOutput) MarketplaceSubscriptionStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MonitorPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.MarketplaceSubscriptionStatus
	}).(pulumi.StringPtrOutput)
}

func (o MonitorPropertiesResponsePtrOutput) MonitoringStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MonitorPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.MonitoringStatus
	}).(pulumi.StringPtrOutput)
}

func (o MonitorPropertiesResponsePtrOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MonitorPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.ProvisioningState
	}).(pulumi.StringPtrOutput)
}

func (o MonitorPropertiesResponsePtrOutput) UserInfo() UserInfoResponsePtrOutput {
	return o.ApplyT(func(v *MonitorPropertiesResponse) *UserInfoResponse {
		if v == nil {
			return nil
		}
		return v.UserInfo
	}).(UserInfoResponsePtrOutput)
}

type MonitoredResourceResponse struct {
	Id                     *string `pulumi:"id"`
	ReasonForLogsStatus    *string `pulumi:"reasonForLogsStatus"`
	ReasonForMetricsStatus *string `pulumi:"reasonForMetricsStatus"`
	SendingLogs            *bool   `pulumi:"sendingLogs"`
	SendingMetrics         *bool   `pulumi:"sendingMetrics"`
}





type MonitoredResourceResponseInput interface {
	pulumi.Input

	ToMonitoredResourceResponseOutput() MonitoredResourceResponseOutput
	ToMonitoredResourceResponseOutputWithContext(context.Context) MonitoredResourceResponseOutput
}

type MonitoredResourceResponseArgs struct {
	Id                     pulumi.StringPtrInput `pulumi:"id"`
	ReasonForLogsStatus    pulumi.StringPtrInput `pulumi:"reasonForLogsStatus"`
	ReasonForMetricsStatus pulumi.StringPtrInput `pulumi:"reasonForMetricsStatus"`
	SendingLogs            pulumi.BoolPtrInput   `pulumi:"sendingLogs"`
	SendingMetrics         pulumi.BoolPtrInput   `pulumi:"sendingMetrics"`
}

func (MonitoredResourceResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MonitoredResourceResponse)(nil)).Elem()
}

func (i MonitoredResourceResponseArgs) ToMonitoredResourceResponseOutput() MonitoredResourceResponseOutput {
	return i.ToMonitoredResourceResponseOutputWithContext(context.Background())
}

func (i MonitoredResourceResponseArgs) ToMonitoredResourceResponseOutputWithContext(ctx context.Context) MonitoredResourceResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MonitoredResourceResponseOutput)
}





type MonitoredResourceResponseArrayInput interface {
	pulumi.Input

	ToMonitoredResourceResponseArrayOutput() MonitoredResourceResponseArrayOutput
	ToMonitoredResourceResponseArrayOutputWithContext(context.Context) MonitoredResourceResponseArrayOutput
}

type MonitoredResourceResponseArray []MonitoredResourceResponseInput

func (MonitoredResourceResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]MonitoredResourceResponse)(nil)).Elem()
}

func (i MonitoredResourceResponseArray) ToMonitoredResourceResponseArrayOutput() MonitoredResourceResponseArrayOutput {
	return i.ToMonitoredResourceResponseArrayOutputWithContext(context.Background())
}

func (i MonitoredResourceResponseArray) ToMonitoredResourceResponseArrayOutputWithContext(ctx context.Context) MonitoredResourceResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MonitoredResourceResponseArrayOutput)
}

type MonitoredResourceResponseOutput struct{ *pulumi.OutputState }

func (MonitoredResourceResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MonitoredResourceResponse)(nil)).Elem()
}

func (o MonitoredResourceResponseOutput) ToMonitoredResourceResponseOutput() MonitoredResourceResponseOutput {
	return o
}

func (o MonitoredResourceResponseOutput) ToMonitoredResourceResponseOutputWithContext(ctx context.Context) MonitoredResourceResponseOutput {
	return o
}

func (o MonitoredResourceResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MonitoredResourceResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o MonitoredResourceResponseOutput) ReasonForLogsStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MonitoredResourceResponse) *string { return v.ReasonForLogsStatus }).(pulumi.StringPtrOutput)
}

func (o MonitoredResourceResponseOutput) ReasonForMetricsStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MonitoredResourceResponse) *string { return v.ReasonForMetricsStatus }).(pulumi.StringPtrOutput)
}

func (o MonitoredResourceResponseOutput) SendingLogs() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v MonitoredResourceResponse) *bool { return v.SendingLogs }).(pulumi.BoolPtrOutput)
}

func (o MonitoredResourceResponseOutput) SendingMetrics() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v MonitoredResourceResponse) *bool { return v.SendingMetrics }).(pulumi.BoolPtrOutput)
}

type MonitoredResourceResponseArrayOutput struct{ *pulumi.OutputState }

func (MonitoredResourceResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]MonitoredResourceResponse)(nil)).Elem()
}

func (o MonitoredResourceResponseArrayOutput) ToMonitoredResourceResponseArrayOutput() MonitoredResourceResponseArrayOutput {
	return o
}

func (o MonitoredResourceResponseArrayOutput) ToMonitoredResourceResponseArrayOutputWithContext(ctx context.Context) MonitoredResourceResponseArrayOutput {
	return o
}

func (o MonitoredResourceResponseArrayOutput) Index(i pulumi.IntInput) MonitoredResourceResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) MonitoredResourceResponse {
		return vs[0].([]MonitoredResourceResponse)[vs[1].(int)]
	}).(MonitoredResourceResponseOutput)
}

type ResourceSku struct {
	Name string `pulumi:"name"`
}





type ResourceSkuInput interface {
	pulumi.Input

	ToResourceSkuOutput() ResourceSkuOutput
	ToResourceSkuOutputWithContext(context.Context) ResourceSkuOutput
}

type ResourceSkuArgs struct {
	Name pulumi.StringInput `pulumi:"name"`
}

func (ResourceSkuArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceSku)(nil)).Elem()
}

func (i ResourceSkuArgs) ToResourceSkuOutput() ResourceSkuOutput {
	return i.ToResourceSkuOutputWithContext(context.Background())
}

func (i ResourceSkuArgs) ToResourceSkuOutputWithContext(ctx context.Context) ResourceSkuOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceSkuOutput)
}

func (i ResourceSkuArgs) ToResourceSkuPtrOutput() ResourceSkuPtrOutput {
	return i.ToResourceSkuPtrOutputWithContext(context.Background())
}

func (i ResourceSkuArgs) ToResourceSkuPtrOutputWithContext(ctx context.Context) ResourceSkuPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceSkuOutput).ToResourceSkuPtrOutputWithContext(ctx)
}









type ResourceSkuPtrInput interface {
	pulumi.Input

	ToResourceSkuPtrOutput() ResourceSkuPtrOutput
	ToResourceSkuPtrOutputWithContext(context.Context) ResourceSkuPtrOutput
}

type resourceSkuPtrType ResourceSkuArgs

func ResourceSkuPtr(v *ResourceSkuArgs) ResourceSkuPtrInput {
	return (*resourceSkuPtrType)(v)
}

func (*resourceSkuPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceSku)(nil)).Elem()
}

func (i *resourceSkuPtrType) ToResourceSkuPtrOutput() ResourceSkuPtrOutput {
	return i.ToResourceSkuPtrOutputWithContext(context.Background())
}

func (i *resourceSkuPtrType) ToResourceSkuPtrOutputWithContext(ctx context.Context) ResourceSkuPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceSkuPtrOutput)
}

type ResourceSkuOutput struct{ *pulumi.OutputState }

func (ResourceSkuOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceSku)(nil)).Elem()
}

func (o ResourceSkuOutput) ToResourceSkuOutput() ResourceSkuOutput {
	return o
}

func (o ResourceSkuOutput) ToResourceSkuOutputWithContext(ctx context.Context) ResourceSkuOutput {
	return o
}

func (o ResourceSkuOutput) ToResourceSkuPtrOutput() ResourceSkuPtrOutput {
	return o.ToResourceSkuPtrOutputWithContext(context.Background())
}

func (o ResourceSkuOutput) ToResourceSkuPtrOutputWithContext(ctx context.Context) ResourceSkuPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ResourceSku) *ResourceSku {
		return &v
	}).(ResourceSkuPtrOutput)
}

func (o ResourceSkuOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ResourceSku) string { return v.Name }).(pulumi.StringOutput)
}

type ResourceSkuPtrOutput struct{ *pulumi.OutputState }

func (ResourceSkuPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceSku)(nil)).Elem()
}

func (o ResourceSkuPtrOutput) ToResourceSkuPtrOutput() ResourceSkuPtrOutput {
	return o
}

func (o ResourceSkuPtrOutput) ToResourceSkuPtrOutputWithContext(ctx context.Context) ResourceSkuPtrOutput {
	return o
}

func (o ResourceSkuPtrOutput) Elem() ResourceSkuOutput {
	return o.ApplyT(func(v *ResourceSku) ResourceSku {
		if v != nil {
			return *v
		}
		var ret ResourceSku
		return ret
	}).(ResourceSkuOutput)
}

func (o ResourceSkuPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ResourceSku) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

type ResourceSkuResponse struct {
	Name string `pulumi:"name"`
}





type ResourceSkuResponseInput interface {
	pulumi.Input

	ToResourceSkuResponseOutput() ResourceSkuResponseOutput
	ToResourceSkuResponseOutputWithContext(context.Context) ResourceSkuResponseOutput
}

type ResourceSkuResponseArgs struct {
	Name pulumi.StringInput `pulumi:"name"`
}

func (ResourceSkuResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceSkuResponse)(nil)).Elem()
}

func (i ResourceSkuResponseArgs) ToResourceSkuResponseOutput() ResourceSkuResponseOutput {
	return i.ToResourceSkuResponseOutputWithContext(context.Background())
}

func (i ResourceSkuResponseArgs) ToResourceSkuResponseOutputWithContext(ctx context.Context) ResourceSkuResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceSkuResponseOutput)
}

func (i ResourceSkuResponseArgs) ToResourceSkuResponsePtrOutput() ResourceSkuResponsePtrOutput {
	return i.ToResourceSkuResponsePtrOutputWithContext(context.Background())
}

func (i ResourceSkuResponseArgs) ToResourceSkuResponsePtrOutputWithContext(ctx context.Context) ResourceSkuResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceSkuResponseOutput).ToResourceSkuResponsePtrOutputWithContext(ctx)
}









type ResourceSkuResponsePtrInput interface {
	pulumi.Input

	ToResourceSkuResponsePtrOutput() ResourceSkuResponsePtrOutput
	ToResourceSkuResponsePtrOutputWithContext(context.Context) ResourceSkuResponsePtrOutput
}

type resourceSkuResponsePtrType ResourceSkuResponseArgs

func ResourceSkuResponsePtr(v *ResourceSkuResponseArgs) ResourceSkuResponsePtrInput {
	return (*resourceSkuResponsePtrType)(v)
}

func (*resourceSkuResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceSkuResponse)(nil)).Elem()
}

func (i *resourceSkuResponsePtrType) ToResourceSkuResponsePtrOutput() ResourceSkuResponsePtrOutput {
	return i.ToResourceSkuResponsePtrOutputWithContext(context.Background())
}

func (i *resourceSkuResponsePtrType) ToResourceSkuResponsePtrOutputWithContext(ctx context.Context) ResourceSkuResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceSkuResponsePtrOutput)
}

type ResourceSkuResponseOutput struct{ *pulumi.OutputState }

func (ResourceSkuResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceSkuResponse)(nil)).Elem()
}

func (o ResourceSkuResponseOutput) ToResourceSkuResponseOutput() ResourceSkuResponseOutput {
	return o
}

func (o ResourceSkuResponseOutput) ToResourceSkuResponseOutputWithContext(ctx context.Context) ResourceSkuResponseOutput {
	return o
}

func (o ResourceSkuResponseOutput) ToResourceSkuResponsePtrOutput() ResourceSkuResponsePtrOutput {
	return o.ToResourceSkuResponsePtrOutputWithContext(context.Background())
}

func (o ResourceSkuResponseOutput) ToResourceSkuResponsePtrOutputWithContext(ctx context.Context) ResourceSkuResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ResourceSkuResponse) *ResourceSkuResponse {
		return &v
	}).(ResourceSkuResponsePtrOutput)
}

func (o ResourceSkuResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ResourceSkuResponse) string { return v.Name }).(pulumi.StringOutput)
}

type ResourceSkuResponsePtrOutput struct{ *pulumi.OutputState }

func (ResourceSkuResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceSkuResponse)(nil)).Elem()
}

func (o ResourceSkuResponsePtrOutput) ToResourceSkuResponsePtrOutput() ResourceSkuResponsePtrOutput {
	return o
}

func (o ResourceSkuResponsePtrOutput) ToResourceSkuResponsePtrOutputWithContext(ctx context.Context) ResourceSkuResponsePtrOutput {
	return o
}

func (o ResourceSkuResponsePtrOutput) Elem() ResourceSkuResponseOutput {
	return o.ApplyT(func(v *ResourceSkuResponse) ResourceSkuResponse {
		if v != nil {
			return *v
		}
		var ret ResourceSkuResponse
		return ret
	}).(ResourceSkuResponseOutput)
}

func (o ResourceSkuResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ResourceSkuResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

type SystemDataResponse struct {
	CreatedAt          *string `pulumi:"createdAt"`
	CreatedBy          *string `pulumi:"createdBy"`
	CreatedByType      *string `pulumi:"createdByType"`
	LastModifiedAt     *string `pulumi:"lastModifiedAt"`
	LastModifiedBy     *string `pulumi:"lastModifiedBy"`
	LastModifiedByType *string `pulumi:"lastModifiedByType"`
}





type SystemDataResponseInput interface {
	pulumi.Input

	ToSystemDataResponseOutput() SystemDataResponseOutput
	ToSystemDataResponseOutputWithContext(context.Context) SystemDataResponseOutput
}

type SystemDataResponseArgs struct {
	CreatedAt          pulumi.StringPtrInput `pulumi:"createdAt"`
	CreatedBy          pulumi.StringPtrInput `pulumi:"createdBy"`
	CreatedByType      pulumi.StringPtrInput `pulumi:"createdByType"`
	LastModifiedAt     pulumi.StringPtrInput `pulumi:"lastModifiedAt"`
	LastModifiedBy     pulumi.StringPtrInput `pulumi:"lastModifiedBy"`
	LastModifiedByType pulumi.StringPtrInput `pulumi:"lastModifiedByType"`
}

func (SystemDataResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SystemDataResponse)(nil)).Elem()
}

func (i SystemDataResponseArgs) ToSystemDataResponseOutput() SystemDataResponseOutput {
	return i.ToSystemDataResponseOutputWithContext(context.Background())
}

func (i SystemDataResponseArgs) ToSystemDataResponseOutputWithContext(ctx context.Context) SystemDataResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemDataResponseOutput)
}

func (i SystemDataResponseArgs) ToSystemDataResponsePtrOutput() SystemDataResponsePtrOutput {
	return i.ToSystemDataResponsePtrOutputWithContext(context.Background())
}

func (i SystemDataResponseArgs) ToSystemDataResponsePtrOutputWithContext(ctx context.Context) SystemDataResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemDataResponseOutput).ToSystemDataResponsePtrOutputWithContext(ctx)
}









type SystemDataResponsePtrInput interface {
	pulumi.Input

	ToSystemDataResponsePtrOutput() SystemDataResponsePtrOutput
	ToSystemDataResponsePtrOutputWithContext(context.Context) SystemDataResponsePtrOutput
}

type systemDataResponsePtrType SystemDataResponseArgs

func SystemDataResponsePtr(v *SystemDataResponseArgs) SystemDataResponsePtrInput {
	return (*systemDataResponsePtrType)(v)
}

func (*systemDataResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemDataResponse)(nil)).Elem()
}

func (i *systemDataResponsePtrType) ToSystemDataResponsePtrOutput() SystemDataResponsePtrOutput {
	return i.ToSystemDataResponsePtrOutputWithContext(context.Background())
}

func (i *systemDataResponsePtrType) ToSystemDataResponsePtrOutputWithContext(ctx context.Context) SystemDataResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemDataResponsePtrOutput)
}

type SystemDataResponseOutput struct{ *pulumi.OutputState }

func (SystemDataResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SystemDataResponse)(nil)).Elem()
}

func (o SystemDataResponseOutput) ToSystemDataResponseOutput() SystemDataResponseOutput {
	return o
}

func (o SystemDataResponseOutput) ToSystemDataResponseOutputWithContext(ctx context.Context) SystemDataResponseOutput {
	return o
}

func (o SystemDataResponseOutput) ToSystemDataResponsePtrOutput() SystemDataResponsePtrOutput {
	return o.ToSystemDataResponsePtrOutputWithContext(context.Background())
}

func (o SystemDataResponseOutput) ToSystemDataResponsePtrOutputWithContext(ctx context.Context) SystemDataResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SystemDataResponse) *SystemDataResponse {
		return &v
	}).(SystemDataResponsePtrOutput)
}

func (o SystemDataResponseOutput) CreatedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.CreatedAt }).(pulumi.StringPtrOutput)
}

func (o SystemDataResponseOutput) CreatedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.CreatedBy }).(pulumi.StringPtrOutput)
}

func (o SystemDataResponseOutput) CreatedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.CreatedByType }).(pulumi.StringPtrOutput)
}

func (o SystemDataResponseOutput) LastModifiedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.LastModifiedAt }).(pulumi.StringPtrOutput)
}

func (o SystemDataResponseOutput) LastModifiedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.LastModifiedBy }).(pulumi.StringPtrOutput)
}

func (o SystemDataResponseOutput) LastModifiedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.LastModifiedByType }).(pulumi.StringPtrOutput)
}

type SystemDataResponsePtrOutput struct{ *pulumi.OutputState }

func (SystemDataResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemDataResponse)(nil)).Elem()
}

func (o SystemDataResponsePtrOutput) ToSystemDataResponsePtrOutput() SystemDataResponsePtrOutput {
	return o
}

func (o SystemDataResponsePtrOutput) ToSystemDataResponsePtrOutputWithContext(ctx context.Context) SystemDataResponsePtrOutput {
	return o
}

func (o SystemDataResponsePtrOutput) Elem() SystemDataResponseOutput {
	return o.ApplyT(func(v *SystemDataResponse) SystemDataResponse {
		if v != nil {
			return *v
		}
		var ret SystemDataResponse
		return ret
	}).(SystemDataResponseOutput)
}

func (o SystemDataResponsePtrOutput) CreatedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemDataResponse) *string {
		if v == nil {
			return nil
		}
		return v.CreatedAt
	}).(pulumi.StringPtrOutput)
}

func (o SystemDataResponsePtrOutput) CreatedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemDataResponse) *string {
		if v == nil {
			return nil
		}
		return v.CreatedBy
	}).(pulumi.StringPtrOutput)
}

func (o SystemDataResponsePtrOutput) CreatedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemDataResponse) *string {
		if v == nil {
			return nil
		}
		return v.CreatedByType
	}).(pulumi.StringPtrOutput)
}

func (o SystemDataResponsePtrOutput) LastModifiedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemDataResponse) *string {
		if v == nil {
			return nil
		}
		return v.LastModifiedAt
	}).(pulumi.StringPtrOutput)
}

func (o SystemDataResponsePtrOutput) LastModifiedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemDataResponse) *string {
		if v == nil {
			return nil
		}
		return v.LastModifiedBy
	}).(pulumi.StringPtrOutput)
}

func (o SystemDataResponsePtrOutput) LastModifiedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemDataResponse) *string {
		if v == nil {
			return nil
		}
		return v.LastModifiedByType
	}).(pulumi.StringPtrOutput)
}

type UserInfo struct {
	EmailAddress *string `pulumi:"emailAddress"`
	Name         *string `pulumi:"name"`
	PhoneNumber  *string `pulumi:"phoneNumber"`
}





type UserInfoInput interface {
	pulumi.Input

	ToUserInfoOutput() UserInfoOutput
	ToUserInfoOutputWithContext(context.Context) UserInfoOutput
}

type UserInfoArgs struct {
	EmailAddress pulumi.StringPtrInput `pulumi:"emailAddress"`
	Name         pulumi.StringPtrInput `pulumi:"name"`
	PhoneNumber  pulumi.StringPtrInput `pulumi:"phoneNumber"`
}

func (UserInfoArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*UserInfo)(nil)).Elem()
}

func (i UserInfoArgs) ToUserInfoOutput() UserInfoOutput {
	return i.ToUserInfoOutputWithContext(context.Background())
}

func (i UserInfoArgs) ToUserInfoOutputWithContext(ctx context.Context) UserInfoOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserInfoOutput)
}

func (i UserInfoArgs) ToUserInfoPtrOutput() UserInfoPtrOutput {
	return i.ToUserInfoPtrOutputWithContext(context.Background())
}

func (i UserInfoArgs) ToUserInfoPtrOutputWithContext(ctx context.Context) UserInfoPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserInfoOutput).ToUserInfoPtrOutputWithContext(ctx)
}









type UserInfoPtrInput interface {
	pulumi.Input

	ToUserInfoPtrOutput() UserInfoPtrOutput
	ToUserInfoPtrOutputWithContext(context.Context) UserInfoPtrOutput
}

type userInfoPtrType UserInfoArgs

func UserInfoPtr(v *UserInfoArgs) UserInfoPtrInput {
	return (*userInfoPtrType)(v)
}

func (*userInfoPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**UserInfo)(nil)).Elem()
}

func (i *userInfoPtrType) ToUserInfoPtrOutput() UserInfoPtrOutput {
	return i.ToUserInfoPtrOutputWithContext(context.Background())
}

func (i *userInfoPtrType) ToUserInfoPtrOutputWithContext(ctx context.Context) UserInfoPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserInfoPtrOutput)
}

type UserInfoOutput struct{ *pulumi.OutputState }

func (UserInfoOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UserInfo)(nil)).Elem()
}

func (o UserInfoOutput) ToUserInfoOutput() UserInfoOutput {
	return o
}

func (o UserInfoOutput) ToUserInfoOutputWithContext(ctx context.Context) UserInfoOutput {
	return o
}

func (o UserInfoOutput) ToUserInfoPtrOutput() UserInfoPtrOutput {
	return o.ToUserInfoPtrOutputWithContext(context.Background())
}

func (o UserInfoOutput) ToUserInfoPtrOutputWithContext(ctx context.Context) UserInfoPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v UserInfo) *UserInfo {
		return &v
	}).(UserInfoPtrOutput)
}

func (o UserInfoOutput) EmailAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v UserInfo) *string { return v.EmailAddress }).(pulumi.StringPtrOutput)
}

func (o UserInfoOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v UserInfo) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o UserInfoOutput) PhoneNumber() pulumi.StringPtrOutput {
	return o.ApplyT(func(v UserInfo) *string { return v.PhoneNumber }).(pulumi.StringPtrOutput)
}

type UserInfoPtrOutput struct{ *pulumi.OutputState }

func (UserInfoPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**UserInfo)(nil)).Elem()
}

func (o UserInfoPtrOutput) ToUserInfoPtrOutput() UserInfoPtrOutput {
	return o
}

func (o UserInfoPtrOutput) ToUserInfoPtrOutputWithContext(ctx context.Context) UserInfoPtrOutput {
	return o
}

func (o UserInfoPtrOutput) Elem() UserInfoOutput {
	return o.ApplyT(func(v *UserInfo) UserInfo {
		if v != nil {
			return *v
		}
		var ret UserInfo
		return ret
	}).(UserInfoOutput)
}

func (o UserInfoPtrOutput) EmailAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UserInfo) *string {
		if v == nil {
			return nil
		}
		return v.EmailAddress
	}).(pulumi.StringPtrOutput)
}

func (o UserInfoPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UserInfo) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o UserInfoPtrOutput) PhoneNumber() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UserInfo) *string {
		if v == nil {
			return nil
		}
		return v.PhoneNumber
	}).(pulumi.StringPtrOutput)
}

type UserInfoResponse struct {
	EmailAddress *string `pulumi:"emailAddress"`
	Name         *string `pulumi:"name"`
	PhoneNumber  *string `pulumi:"phoneNumber"`
}





type UserInfoResponseInput interface {
	pulumi.Input

	ToUserInfoResponseOutput() UserInfoResponseOutput
	ToUserInfoResponseOutputWithContext(context.Context) UserInfoResponseOutput
}

type UserInfoResponseArgs struct {
	EmailAddress pulumi.StringPtrInput `pulumi:"emailAddress"`
	Name         pulumi.StringPtrInput `pulumi:"name"`
	PhoneNumber  pulumi.StringPtrInput `pulumi:"phoneNumber"`
}

func (UserInfoResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*UserInfoResponse)(nil)).Elem()
}

func (i UserInfoResponseArgs) ToUserInfoResponseOutput() UserInfoResponseOutput {
	return i.ToUserInfoResponseOutputWithContext(context.Background())
}

func (i UserInfoResponseArgs) ToUserInfoResponseOutputWithContext(ctx context.Context) UserInfoResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserInfoResponseOutput)
}

func (i UserInfoResponseArgs) ToUserInfoResponsePtrOutput() UserInfoResponsePtrOutput {
	return i.ToUserInfoResponsePtrOutputWithContext(context.Background())
}

func (i UserInfoResponseArgs) ToUserInfoResponsePtrOutputWithContext(ctx context.Context) UserInfoResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserInfoResponseOutput).ToUserInfoResponsePtrOutputWithContext(ctx)
}









type UserInfoResponsePtrInput interface {
	pulumi.Input

	ToUserInfoResponsePtrOutput() UserInfoResponsePtrOutput
	ToUserInfoResponsePtrOutputWithContext(context.Context) UserInfoResponsePtrOutput
}

type userInfoResponsePtrType UserInfoResponseArgs

func UserInfoResponsePtr(v *UserInfoResponseArgs) UserInfoResponsePtrInput {
	return (*userInfoResponsePtrType)(v)
}

func (*userInfoResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**UserInfoResponse)(nil)).Elem()
}

func (i *userInfoResponsePtrType) ToUserInfoResponsePtrOutput() UserInfoResponsePtrOutput {
	return i.ToUserInfoResponsePtrOutputWithContext(context.Background())
}

func (i *userInfoResponsePtrType) ToUserInfoResponsePtrOutputWithContext(ctx context.Context) UserInfoResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserInfoResponsePtrOutput)
}

type UserInfoResponseOutput struct{ *pulumi.OutputState }

func (UserInfoResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UserInfoResponse)(nil)).Elem()
}

func (o UserInfoResponseOutput) ToUserInfoResponseOutput() UserInfoResponseOutput {
	return o
}

func (o UserInfoResponseOutput) ToUserInfoResponseOutputWithContext(ctx context.Context) UserInfoResponseOutput {
	return o
}

func (o UserInfoResponseOutput) ToUserInfoResponsePtrOutput() UserInfoResponsePtrOutput {
	return o.ToUserInfoResponsePtrOutputWithContext(context.Background())
}

func (o UserInfoResponseOutput) ToUserInfoResponsePtrOutputWithContext(ctx context.Context) UserInfoResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v UserInfoResponse) *UserInfoResponse {
		return &v
	}).(UserInfoResponsePtrOutput)
}

func (o UserInfoResponseOutput) EmailAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v UserInfoResponse) *string { return v.EmailAddress }).(pulumi.StringPtrOutput)
}

func (o UserInfoResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v UserInfoResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o UserInfoResponseOutput) PhoneNumber() pulumi.StringPtrOutput {
	return o.ApplyT(func(v UserInfoResponse) *string { return v.PhoneNumber }).(pulumi.StringPtrOutput)
}

type UserInfoResponsePtrOutput struct{ *pulumi.OutputState }

func (UserInfoResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**UserInfoResponse)(nil)).Elem()
}

func (o UserInfoResponsePtrOutput) ToUserInfoResponsePtrOutput() UserInfoResponsePtrOutput {
	return o
}

func (o UserInfoResponsePtrOutput) ToUserInfoResponsePtrOutputWithContext(ctx context.Context) UserInfoResponsePtrOutput {
	return o
}

func (o UserInfoResponsePtrOutput) Elem() UserInfoResponseOutput {
	return o.ApplyT(func(v *UserInfoResponse) UserInfoResponse {
		if v != nil {
			return *v
		}
		var ret UserInfoResponse
		return ret
	}).(UserInfoResponseOutput)
}

func (o UserInfoResponsePtrOutput) EmailAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UserInfoResponse) *string {
		if v == nil {
			return nil
		}
		return v.EmailAddress
	}).(pulumi.StringPtrOutput)
}

func (o UserInfoResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UserInfoResponse) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o UserInfoResponsePtrOutput) PhoneNumber() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UserInfoResponse) *string {
		if v == nil {
			return nil
		}
		return v.PhoneNumber
	}).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DatadogApiKeyResponseInput)(nil)).Elem(), DatadogApiKeyResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*DatadogApiKeyResponseArrayInput)(nil)).Elem(), DatadogApiKeyResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DatadogHostMetadataResponseInput)(nil)).Elem(), DatadogHostMetadataResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*DatadogHostMetadataResponsePtrInput)(nil)).Elem(), DatadogHostMetadataResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*DatadogHostResponseInput)(nil)).Elem(), DatadogHostResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*DatadogHostResponseArrayInput)(nil)).Elem(), DatadogHostResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DatadogInstallMethodResponseInput)(nil)).Elem(), DatadogInstallMethodResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*DatadogInstallMethodResponsePtrInput)(nil)).Elem(), DatadogInstallMethodResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*DatadogLogsAgentResponseInput)(nil)).Elem(), DatadogLogsAgentResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*DatadogLogsAgentResponsePtrInput)(nil)).Elem(), DatadogLogsAgentResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*DatadogOrganizationPropertiesInput)(nil)).Elem(), DatadogOrganizationPropertiesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*DatadogOrganizationPropertiesPtrInput)(nil)).Elem(), DatadogOrganizationPropertiesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*DatadogOrganizationPropertiesResponseInput)(nil)).Elem(), DatadogOrganizationPropertiesResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*DatadogOrganizationPropertiesResponsePtrInput)(nil)).Elem(), DatadogOrganizationPropertiesResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*IdentityPropertiesInput)(nil)).Elem(), IdentityPropertiesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*IdentityPropertiesPtrInput)(nil)).Elem(), IdentityPropertiesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*IdentityPropertiesResponseInput)(nil)).Elem(), IdentityPropertiesResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*IdentityPropertiesResponsePtrInput)(nil)).Elem(), IdentityPropertiesResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*LinkedResourceResponseInput)(nil)).Elem(), LinkedResourceResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*LinkedResourceResponseArrayInput)(nil)).Elem(), LinkedResourceResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*MonitorPropertiesInput)(nil)).Elem(), MonitorPropertiesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*MonitorPropertiesPtrInput)(nil)).Elem(), MonitorPropertiesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*MonitorPropertiesResponseInput)(nil)).Elem(), MonitorPropertiesResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*MonitorPropertiesResponsePtrInput)(nil)).Elem(), MonitorPropertiesResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*MonitoredResourceResponseInput)(nil)).Elem(), MonitoredResourceResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*MonitoredResourceResponseArrayInput)(nil)).Elem(), MonitoredResourceResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ResourceSkuInput)(nil)).Elem(), ResourceSkuArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ResourceSkuPtrInput)(nil)).Elem(), ResourceSkuArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ResourceSkuResponseInput)(nil)).Elem(), ResourceSkuResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ResourceSkuResponsePtrInput)(nil)).Elem(), ResourceSkuResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemDataResponseInput)(nil)).Elem(), SystemDataResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemDataResponsePtrInput)(nil)).Elem(), SystemDataResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*UserInfoInput)(nil)).Elem(), UserInfoArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*UserInfoPtrInput)(nil)).Elem(), UserInfoArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*UserInfoResponseInput)(nil)).Elem(), UserInfoResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*UserInfoResponsePtrInput)(nil)).Elem(), UserInfoResponseArgs{})
	pulumi.RegisterOutputType(DatadogApiKeyResponseOutput{})
	pulumi.RegisterOutputType(DatadogApiKeyResponseArrayOutput{})
	pulumi.RegisterOutputType(DatadogHostMetadataResponseOutput{})
	pulumi.RegisterOutputType(DatadogHostMetadataResponsePtrOutput{})
	pulumi.RegisterOutputType(DatadogHostResponseOutput{})
	pulumi.RegisterOutputType(DatadogHostResponseArrayOutput{})
	pulumi.RegisterOutputType(DatadogInstallMethodResponseOutput{})
	pulumi.RegisterOutputType(DatadogInstallMethodResponsePtrOutput{})
	pulumi.RegisterOutputType(DatadogLogsAgentResponseOutput{})
	pulumi.RegisterOutputType(DatadogLogsAgentResponsePtrOutput{})
	pulumi.RegisterOutputType(DatadogOrganizationPropertiesOutput{})
	pulumi.RegisterOutputType(DatadogOrganizationPropertiesPtrOutput{})
	pulumi.RegisterOutputType(DatadogOrganizationPropertiesResponseOutput{})
	pulumi.RegisterOutputType(DatadogOrganizationPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(IdentityPropertiesOutput{})
	pulumi.RegisterOutputType(IdentityPropertiesPtrOutput{})
	pulumi.RegisterOutputType(IdentityPropertiesResponseOutput{})
	pulumi.RegisterOutputType(IdentityPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(LinkedResourceResponseOutput{})
	pulumi.RegisterOutputType(LinkedResourceResponseArrayOutput{})
	pulumi.RegisterOutputType(MonitorPropertiesOutput{})
	pulumi.RegisterOutputType(MonitorPropertiesPtrOutput{})
	pulumi.RegisterOutputType(MonitorPropertiesResponseOutput{})
	pulumi.RegisterOutputType(MonitorPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(MonitoredResourceResponseOutput{})
	pulumi.RegisterOutputType(MonitoredResourceResponseArrayOutput{})
	pulumi.RegisterOutputType(ResourceSkuOutput{})
	pulumi.RegisterOutputType(ResourceSkuPtrOutput{})
	pulumi.RegisterOutputType(ResourceSkuResponseOutput{})
	pulumi.RegisterOutputType(ResourceSkuResponsePtrOutput{})
	pulumi.RegisterOutputType(SystemDataResponseOutput{})
	pulumi.RegisterOutputType(SystemDataResponsePtrOutput{})
	pulumi.RegisterOutputType(UserInfoOutput{})
	pulumi.RegisterOutputType(UserInfoPtrOutput{})
	pulumi.RegisterOutputType(UserInfoResponseOutput{})
	pulumi.RegisterOutputType(UserInfoResponsePtrOutput{})
}
