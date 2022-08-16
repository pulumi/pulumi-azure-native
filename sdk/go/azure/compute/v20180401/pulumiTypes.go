


package v20180401

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AdditionalUnattendContent struct {
	ComponentName *ComponentNames `pulumi:"componentName"`
	Content       *string         `pulumi:"content"`
	PassName      *PassNames      `pulumi:"passName"`
	SettingName   *SettingNames   `pulumi:"settingName"`
}





type AdditionalUnattendContentInput interface {
	pulumi.Input

	ToAdditionalUnattendContentOutput() AdditionalUnattendContentOutput
	ToAdditionalUnattendContentOutputWithContext(context.Context) AdditionalUnattendContentOutput
}

type AdditionalUnattendContentArgs struct {
	ComponentName ComponentNamesPtrInput `pulumi:"componentName"`
	Content       pulumi.StringPtrInput  `pulumi:"content"`
	PassName      PassNamesPtrInput      `pulumi:"passName"`
	SettingName   SettingNamesPtrInput   `pulumi:"settingName"`
}

func (AdditionalUnattendContentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AdditionalUnattendContent)(nil)).Elem()
}

func (i AdditionalUnattendContentArgs) ToAdditionalUnattendContentOutput() AdditionalUnattendContentOutput {
	return i.ToAdditionalUnattendContentOutputWithContext(context.Background())
}

func (i AdditionalUnattendContentArgs) ToAdditionalUnattendContentOutputWithContext(ctx context.Context) AdditionalUnattendContentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AdditionalUnattendContentOutput)
}





type AdditionalUnattendContentArrayInput interface {
	pulumi.Input

	ToAdditionalUnattendContentArrayOutput() AdditionalUnattendContentArrayOutput
	ToAdditionalUnattendContentArrayOutputWithContext(context.Context) AdditionalUnattendContentArrayOutput
}

type AdditionalUnattendContentArray []AdditionalUnattendContentInput

func (AdditionalUnattendContentArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AdditionalUnattendContent)(nil)).Elem()
}

func (i AdditionalUnattendContentArray) ToAdditionalUnattendContentArrayOutput() AdditionalUnattendContentArrayOutput {
	return i.ToAdditionalUnattendContentArrayOutputWithContext(context.Background())
}

func (i AdditionalUnattendContentArray) ToAdditionalUnattendContentArrayOutputWithContext(ctx context.Context) AdditionalUnattendContentArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AdditionalUnattendContentArrayOutput)
}

type AdditionalUnattendContentOutput struct{ *pulumi.OutputState }

func (AdditionalUnattendContentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AdditionalUnattendContent)(nil)).Elem()
}

func (o AdditionalUnattendContentOutput) ToAdditionalUnattendContentOutput() AdditionalUnattendContentOutput {
	return o
}

func (o AdditionalUnattendContentOutput) ToAdditionalUnattendContentOutputWithContext(ctx context.Context) AdditionalUnattendContentOutput {
	return o
}

func (o AdditionalUnattendContentOutput) ComponentName() ComponentNamesPtrOutput {
	return o.ApplyT(func(v AdditionalUnattendContent) *ComponentNames { return v.ComponentName }).(ComponentNamesPtrOutput)
}

func (o AdditionalUnattendContentOutput) Content() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AdditionalUnattendContent) *string { return v.Content }).(pulumi.StringPtrOutput)
}

func (o AdditionalUnattendContentOutput) PassName() PassNamesPtrOutput {
	return o.ApplyT(func(v AdditionalUnattendContent) *PassNames { return v.PassName }).(PassNamesPtrOutput)
}

func (o AdditionalUnattendContentOutput) SettingName() SettingNamesPtrOutput {
	return o.ApplyT(func(v AdditionalUnattendContent) *SettingNames { return v.SettingName }).(SettingNamesPtrOutput)
}

type AdditionalUnattendContentArrayOutput struct{ *pulumi.OutputState }

func (AdditionalUnattendContentArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AdditionalUnattendContent)(nil)).Elem()
}

func (o AdditionalUnattendContentArrayOutput) ToAdditionalUnattendContentArrayOutput() AdditionalUnattendContentArrayOutput {
	return o
}

func (o AdditionalUnattendContentArrayOutput) ToAdditionalUnattendContentArrayOutputWithContext(ctx context.Context) AdditionalUnattendContentArrayOutput {
	return o
}

func (o AdditionalUnattendContentArrayOutput) Index(i pulumi.IntInput) AdditionalUnattendContentOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) AdditionalUnattendContent {
		return vs[0].([]AdditionalUnattendContent)[vs[1].(int)]
	}).(AdditionalUnattendContentOutput)
}

type AdditionalUnattendContentResponse struct {
	ComponentName *string `pulumi:"componentName"`
	Content       *string `pulumi:"content"`
	PassName      *string `pulumi:"passName"`
	SettingName   *string `pulumi:"settingName"`
}

type AdditionalUnattendContentResponseOutput struct{ *pulumi.OutputState }

func (AdditionalUnattendContentResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AdditionalUnattendContentResponse)(nil)).Elem()
}

func (o AdditionalUnattendContentResponseOutput) ToAdditionalUnattendContentResponseOutput() AdditionalUnattendContentResponseOutput {
	return o
}

func (o AdditionalUnattendContentResponseOutput) ToAdditionalUnattendContentResponseOutputWithContext(ctx context.Context) AdditionalUnattendContentResponseOutput {
	return o
}

func (o AdditionalUnattendContentResponseOutput) ComponentName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AdditionalUnattendContentResponse) *string { return v.ComponentName }).(pulumi.StringPtrOutput)
}

func (o AdditionalUnattendContentResponseOutput) Content() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AdditionalUnattendContentResponse) *string { return v.Content }).(pulumi.StringPtrOutput)
}

func (o AdditionalUnattendContentResponseOutput) PassName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AdditionalUnattendContentResponse) *string { return v.PassName }).(pulumi.StringPtrOutput)
}

func (o AdditionalUnattendContentResponseOutput) SettingName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AdditionalUnattendContentResponse) *string { return v.SettingName }).(pulumi.StringPtrOutput)
}

type AdditionalUnattendContentResponseArrayOutput struct{ *pulumi.OutputState }

func (AdditionalUnattendContentResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AdditionalUnattendContentResponse)(nil)).Elem()
}

func (o AdditionalUnattendContentResponseArrayOutput) ToAdditionalUnattendContentResponseArrayOutput() AdditionalUnattendContentResponseArrayOutput {
	return o
}

func (o AdditionalUnattendContentResponseArrayOutput) ToAdditionalUnattendContentResponseArrayOutputWithContext(ctx context.Context) AdditionalUnattendContentResponseArrayOutput {
	return o
}

func (o AdditionalUnattendContentResponseArrayOutput) Index(i pulumi.IntInput) AdditionalUnattendContentResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) AdditionalUnattendContentResponse {
		return vs[0].([]AdditionalUnattendContentResponse)[vs[1].(int)]
	}).(AdditionalUnattendContentResponseOutput)
}

type ApiEntityReference struct {
	Id *string `pulumi:"id"`
}





type ApiEntityReferenceInput interface {
	pulumi.Input

	ToApiEntityReferenceOutput() ApiEntityReferenceOutput
	ToApiEntityReferenceOutputWithContext(context.Context) ApiEntityReferenceOutput
}

type ApiEntityReferenceArgs struct {
	Id pulumi.StringPtrInput `pulumi:"id"`
}

func (ApiEntityReferenceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ApiEntityReference)(nil)).Elem()
}

func (i ApiEntityReferenceArgs) ToApiEntityReferenceOutput() ApiEntityReferenceOutput {
	return i.ToApiEntityReferenceOutputWithContext(context.Background())
}

func (i ApiEntityReferenceArgs) ToApiEntityReferenceOutputWithContext(ctx context.Context) ApiEntityReferenceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiEntityReferenceOutput)
}

func (i ApiEntityReferenceArgs) ToApiEntityReferencePtrOutput() ApiEntityReferencePtrOutput {
	return i.ToApiEntityReferencePtrOutputWithContext(context.Background())
}

func (i ApiEntityReferenceArgs) ToApiEntityReferencePtrOutputWithContext(ctx context.Context) ApiEntityReferencePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiEntityReferenceOutput).ToApiEntityReferencePtrOutputWithContext(ctx)
}









type ApiEntityReferencePtrInput interface {
	pulumi.Input

	ToApiEntityReferencePtrOutput() ApiEntityReferencePtrOutput
	ToApiEntityReferencePtrOutputWithContext(context.Context) ApiEntityReferencePtrOutput
}

type apiEntityReferencePtrType ApiEntityReferenceArgs

func ApiEntityReferencePtr(v *ApiEntityReferenceArgs) ApiEntityReferencePtrInput {
	return (*apiEntityReferencePtrType)(v)
}

func (*apiEntityReferencePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ApiEntityReference)(nil)).Elem()
}

func (i *apiEntityReferencePtrType) ToApiEntityReferencePtrOutput() ApiEntityReferencePtrOutput {
	return i.ToApiEntityReferencePtrOutputWithContext(context.Background())
}

func (i *apiEntityReferencePtrType) ToApiEntityReferencePtrOutputWithContext(ctx context.Context) ApiEntityReferencePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiEntityReferencePtrOutput)
}

type ApiEntityReferenceOutput struct{ *pulumi.OutputState }

func (ApiEntityReferenceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApiEntityReference)(nil)).Elem()
}

func (o ApiEntityReferenceOutput) ToApiEntityReferenceOutput() ApiEntityReferenceOutput {
	return o
}

func (o ApiEntityReferenceOutput) ToApiEntityReferenceOutputWithContext(ctx context.Context) ApiEntityReferenceOutput {
	return o
}

func (o ApiEntityReferenceOutput) ToApiEntityReferencePtrOutput() ApiEntityReferencePtrOutput {
	return o.ToApiEntityReferencePtrOutputWithContext(context.Background())
}

func (o ApiEntityReferenceOutput) ToApiEntityReferencePtrOutputWithContext(ctx context.Context) ApiEntityReferencePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ApiEntityReference) *ApiEntityReference {
		return &v
	}).(ApiEntityReferencePtrOutput)
}

func (o ApiEntityReferenceOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApiEntityReference) *string { return v.Id }).(pulumi.StringPtrOutput)
}

type ApiEntityReferencePtrOutput struct{ *pulumi.OutputState }

func (ApiEntityReferencePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApiEntityReference)(nil)).Elem()
}

func (o ApiEntityReferencePtrOutput) ToApiEntityReferencePtrOutput() ApiEntityReferencePtrOutput {
	return o
}

func (o ApiEntityReferencePtrOutput) ToApiEntityReferencePtrOutputWithContext(ctx context.Context) ApiEntityReferencePtrOutput {
	return o
}

func (o ApiEntityReferencePtrOutput) Elem() ApiEntityReferenceOutput {
	return o.ApplyT(func(v *ApiEntityReference) ApiEntityReference {
		if v != nil {
			return *v
		}
		var ret ApiEntityReference
		return ret
	}).(ApiEntityReferenceOutput)
}

func (o ApiEntityReferencePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApiEntityReference) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

type ApiEntityReferenceResponse struct {
	Id *string `pulumi:"id"`
}

type ApiEntityReferenceResponseOutput struct{ *pulumi.OutputState }

func (ApiEntityReferenceResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApiEntityReferenceResponse)(nil)).Elem()
}

func (o ApiEntityReferenceResponseOutput) ToApiEntityReferenceResponseOutput() ApiEntityReferenceResponseOutput {
	return o
}

func (o ApiEntityReferenceResponseOutput) ToApiEntityReferenceResponseOutputWithContext(ctx context.Context) ApiEntityReferenceResponseOutput {
	return o
}

func (o ApiEntityReferenceResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApiEntityReferenceResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

type ApiEntityReferenceResponsePtrOutput struct{ *pulumi.OutputState }

func (ApiEntityReferenceResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApiEntityReferenceResponse)(nil)).Elem()
}

func (o ApiEntityReferenceResponsePtrOutput) ToApiEntityReferenceResponsePtrOutput() ApiEntityReferenceResponsePtrOutput {
	return o
}

func (o ApiEntityReferenceResponsePtrOutput) ToApiEntityReferenceResponsePtrOutputWithContext(ctx context.Context) ApiEntityReferenceResponsePtrOutput {
	return o
}

func (o ApiEntityReferenceResponsePtrOutput) Elem() ApiEntityReferenceResponseOutput {
	return o.ApplyT(func(v *ApiEntityReferenceResponse) ApiEntityReferenceResponse {
		if v != nil {
			return *v
		}
		var ret ApiEntityReferenceResponse
		return ret
	}).(ApiEntityReferenceResponseOutput)
}

func (o ApiEntityReferenceResponsePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApiEntityReferenceResponse) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

type AutoOSUpgradePolicy struct {
	DisableAutoRollback *bool `pulumi:"disableAutoRollback"`
}





type AutoOSUpgradePolicyInput interface {
	pulumi.Input

	ToAutoOSUpgradePolicyOutput() AutoOSUpgradePolicyOutput
	ToAutoOSUpgradePolicyOutputWithContext(context.Context) AutoOSUpgradePolicyOutput
}

type AutoOSUpgradePolicyArgs struct {
	DisableAutoRollback pulumi.BoolPtrInput `pulumi:"disableAutoRollback"`
}

func (AutoOSUpgradePolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AutoOSUpgradePolicy)(nil)).Elem()
}

func (i AutoOSUpgradePolicyArgs) ToAutoOSUpgradePolicyOutput() AutoOSUpgradePolicyOutput {
	return i.ToAutoOSUpgradePolicyOutputWithContext(context.Background())
}

func (i AutoOSUpgradePolicyArgs) ToAutoOSUpgradePolicyOutputWithContext(ctx context.Context) AutoOSUpgradePolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutoOSUpgradePolicyOutput)
}

func (i AutoOSUpgradePolicyArgs) ToAutoOSUpgradePolicyPtrOutput() AutoOSUpgradePolicyPtrOutput {
	return i.ToAutoOSUpgradePolicyPtrOutputWithContext(context.Background())
}

func (i AutoOSUpgradePolicyArgs) ToAutoOSUpgradePolicyPtrOutputWithContext(ctx context.Context) AutoOSUpgradePolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutoOSUpgradePolicyOutput).ToAutoOSUpgradePolicyPtrOutputWithContext(ctx)
}









type AutoOSUpgradePolicyPtrInput interface {
	pulumi.Input

	ToAutoOSUpgradePolicyPtrOutput() AutoOSUpgradePolicyPtrOutput
	ToAutoOSUpgradePolicyPtrOutputWithContext(context.Context) AutoOSUpgradePolicyPtrOutput
}

type autoOSUpgradePolicyPtrType AutoOSUpgradePolicyArgs

func AutoOSUpgradePolicyPtr(v *AutoOSUpgradePolicyArgs) AutoOSUpgradePolicyPtrInput {
	return (*autoOSUpgradePolicyPtrType)(v)
}

func (*autoOSUpgradePolicyPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AutoOSUpgradePolicy)(nil)).Elem()
}

func (i *autoOSUpgradePolicyPtrType) ToAutoOSUpgradePolicyPtrOutput() AutoOSUpgradePolicyPtrOutput {
	return i.ToAutoOSUpgradePolicyPtrOutputWithContext(context.Background())
}

func (i *autoOSUpgradePolicyPtrType) ToAutoOSUpgradePolicyPtrOutputWithContext(ctx context.Context) AutoOSUpgradePolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutoOSUpgradePolicyPtrOutput)
}

type AutoOSUpgradePolicyOutput struct{ *pulumi.OutputState }

func (AutoOSUpgradePolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AutoOSUpgradePolicy)(nil)).Elem()
}

func (o AutoOSUpgradePolicyOutput) ToAutoOSUpgradePolicyOutput() AutoOSUpgradePolicyOutput {
	return o
}

func (o AutoOSUpgradePolicyOutput) ToAutoOSUpgradePolicyOutputWithContext(ctx context.Context) AutoOSUpgradePolicyOutput {
	return o
}

func (o AutoOSUpgradePolicyOutput) ToAutoOSUpgradePolicyPtrOutput() AutoOSUpgradePolicyPtrOutput {
	return o.ToAutoOSUpgradePolicyPtrOutputWithContext(context.Background())
}

func (o AutoOSUpgradePolicyOutput) ToAutoOSUpgradePolicyPtrOutputWithContext(ctx context.Context) AutoOSUpgradePolicyPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AutoOSUpgradePolicy) *AutoOSUpgradePolicy {
		return &v
	}).(AutoOSUpgradePolicyPtrOutput)
}

func (o AutoOSUpgradePolicyOutput) DisableAutoRollback() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v AutoOSUpgradePolicy) *bool { return v.DisableAutoRollback }).(pulumi.BoolPtrOutput)
}

type AutoOSUpgradePolicyPtrOutput struct{ *pulumi.OutputState }

func (AutoOSUpgradePolicyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AutoOSUpgradePolicy)(nil)).Elem()
}

func (o AutoOSUpgradePolicyPtrOutput) ToAutoOSUpgradePolicyPtrOutput() AutoOSUpgradePolicyPtrOutput {
	return o
}

func (o AutoOSUpgradePolicyPtrOutput) ToAutoOSUpgradePolicyPtrOutputWithContext(ctx context.Context) AutoOSUpgradePolicyPtrOutput {
	return o
}

func (o AutoOSUpgradePolicyPtrOutput) Elem() AutoOSUpgradePolicyOutput {
	return o.ApplyT(func(v *AutoOSUpgradePolicy) AutoOSUpgradePolicy {
		if v != nil {
			return *v
		}
		var ret AutoOSUpgradePolicy
		return ret
	}).(AutoOSUpgradePolicyOutput)
}

func (o AutoOSUpgradePolicyPtrOutput) DisableAutoRollback() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AutoOSUpgradePolicy) *bool {
		if v == nil {
			return nil
		}
		return v.DisableAutoRollback
	}).(pulumi.BoolPtrOutput)
}

type AutoOSUpgradePolicyResponse struct {
	DisableAutoRollback *bool `pulumi:"disableAutoRollback"`
}

type AutoOSUpgradePolicyResponseOutput struct{ *pulumi.OutputState }

func (AutoOSUpgradePolicyResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AutoOSUpgradePolicyResponse)(nil)).Elem()
}

func (o AutoOSUpgradePolicyResponseOutput) ToAutoOSUpgradePolicyResponseOutput() AutoOSUpgradePolicyResponseOutput {
	return o
}

func (o AutoOSUpgradePolicyResponseOutput) ToAutoOSUpgradePolicyResponseOutputWithContext(ctx context.Context) AutoOSUpgradePolicyResponseOutput {
	return o
}

func (o AutoOSUpgradePolicyResponseOutput) DisableAutoRollback() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v AutoOSUpgradePolicyResponse) *bool { return v.DisableAutoRollback }).(pulumi.BoolPtrOutput)
}

type AutoOSUpgradePolicyResponsePtrOutput struct{ *pulumi.OutputState }

func (AutoOSUpgradePolicyResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AutoOSUpgradePolicyResponse)(nil)).Elem()
}

func (o AutoOSUpgradePolicyResponsePtrOutput) ToAutoOSUpgradePolicyResponsePtrOutput() AutoOSUpgradePolicyResponsePtrOutput {
	return o
}

func (o AutoOSUpgradePolicyResponsePtrOutput) ToAutoOSUpgradePolicyResponsePtrOutputWithContext(ctx context.Context) AutoOSUpgradePolicyResponsePtrOutput {
	return o
}

func (o AutoOSUpgradePolicyResponsePtrOutput) Elem() AutoOSUpgradePolicyResponseOutput {
	return o.ApplyT(func(v *AutoOSUpgradePolicyResponse) AutoOSUpgradePolicyResponse {
		if v != nil {
			return *v
		}
		var ret AutoOSUpgradePolicyResponse
		return ret
	}).(AutoOSUpgradePolicyResponseOutput)
}

func (o AutoOSUpgradePolicyResponsePtrOutput) DisableAutoRollback() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AutoOSUpgradePolicyResponse) *bool {
		if v == nil {
			return nil
		}
		return v.DisableAutoRollback
	}).(pulumi.BoolPtrOutput)
}

type BootDiagnostics struct {
	Enabled    *bool   `pulumi:"enabled"`
	StorageUri *string `pulumi:"storageUri"`
}





type BootDiagnosticsInput interface {
	pulumi.Input

	ToBootDiagnosticsOutput() BootDiagnosticsOutput
	ToBootDiagnosticsOutputWithContext(context.Context) BootDiagnosticsOutput
}

type BootDiagnosticsArgs struct {
	Enabled    pulumi.BoolPtrInput   `pulumi:"enabled"`
	StorageUri pulumi.StringPtrInput `pulumi:"storageUri"`
}

func (BootDiagnosticsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*BootDiagnostics)(nil)).Elem()
}

func (i BootDiagnosticsArgs) ToBootDiagnosticsOutput() BootDiagnosticsOutput {
	return i.ToBootDiagnosticsOutputWithContext(context.Background())
}

func (i BootDiagnosticsArgs) ToBootDiagnosticsOutputWithContext(ctx context.Context) BootDiagnosticsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BootDiagnosticsOutput)
}

func (i BootDiagnosticsArgs) ToBootDiagnosticsPtrOutput() BootDiagnosticsPtrOutput {
	return i.ToBootDiagnosticsPtrOutputWithContext(context.Background())
}

func (i BootDiagnosticsArgs) ToBootDiagnosticsPtrOutputWithContext(ctx context.Context) BootDiagnosticsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BootDiagnosticsOutput).ToBootDiagnosticsPtrOutputWithContext(ctx)
}









type BootDiagnosticsPtrInput interface {
	pulumi.Input

	ToBootDiagnosticsPtrOutput() BootDiagnosticsPtrOutput
	ToBootDiagnosticsPtrOutputWithContext(context.Context) BootDiagnosticsPtrOutput
}

type bootDiagnosticsPtrType BootDiagnosticsArgs

func BootDiagnosticsPtr(v *BootDiagnosticsArgs) BootDiagnosticsPtrInput {
	return (*bootDiagnosticsPtrType)(v)
}

func (*bootDiagnosticsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**BootDiagnostics)(nil)).Elem()
}

func (i *bootDiagnosticsPtrType) ToBootDiagnosticsPtrOutput() BootDiagnosticsPtrOutput {
	return i.ToBootDiagnosticsPtrOutputWithContext(context.Background())
}

func (i *bootDiagnosticsPtrType) ToBootDiagnosticsPtrOutputWithContext(ctx context.Context) BootDiagnosticsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BootDiagnosticsPtrOutput)
}

type BootDiagnosticsOutput struct{ *pulumi.OutputState }

func (BootDiagnosticsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BootDiagnostics)(nil)).Elem()
}

func (o BootDiagnosticsOutput) ToBootDiagnosticsOutput() BootDiagnosticsOutput {
	return o
}

func (o BootDiagnosticsOutput) ToBootDiagnosticsOutputWithContext(ctx context.Context) BootDiagnosticsOutput {
	return o
}

func (o BootDiagnosticsOutput) ToBootDiagnosticsPtrOutput() BootDiagnosticsPtrOutput {
	return o.ToBootDiagnosticsPtrOutputWithContext(context.Background())
}

func (o BootDiagnosticsOutput) ToBootDiagnosticsPtrOutputWithContext(ctx context.Context) BootDiagnosticsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v BootDiagnostics) *BootDiagnostics {
		return &v
	}).(BootDiagnosticsPtrOutput)
}

func (o BootDiagnosticsOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v BootDiagnostics) *bool { return v.Enabled }).(pulumi.BoolPtrOutput)
}

func (o BootDiagnosticsOutput) StorageUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BootDiagnostics) *string { return v.StorageUri }).(pulumi.StringPtrOutput)
}

type BootDiagnosticsPtrOutput struct{ *pulumi.OutputState }

func (BootDiagnosticsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BootDiagnostics)(nil)).Elem()
}

func (o BootDiagnosticsPtrOutput) ToBootDiagnosticsPtrOutput() BootDiagnosticsPtrOutput {
	return o
}

func (o BootDiagnosticsPtrOutput) ToBootDiagnosticsPtrOutputWithContext(ctx context.Context) BootDiagnosticsPtrOutput {
	return o
}

func (o BootDiagnosticsPtrOutput) Elem() BootDiagnosticsOutput {
	return o.ApplyT(func(v *BootDiagnostics) BootDiagnostics {
		if v != nil {
			return *v
		}
		var ret BootDiagnostics
		return ret
	}).(BootDiagnosticsOutput)
}

func (o BootDiagnosticsPtrOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *BootDiagnostics) *bool {
		if v == nil {
			return nil
		}
		return v.Enabled
	}).(pulumi.BoolPtrOutput)
}

func (o BootDiagnosticsPtrOutput) StorageUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BootDiagnostics) *string {
		if v == nil {
			return nil
		}
		return v.StorageUri
	}).(pulumi.StringPtrOutput)
}

type BootDiagnosticsInstanceViewResponse struct {
	ConsoleScreenshotBlobUri string                     `pulumi:"consoleScreenshotBlobUri"`
	SerialConsoleLogBlobUri  string                     `pulumi:"serialConsoleLogBlobUri"`
	Status                   InstanceViewStatusResponse `pulumi:"status"`
}

type BootDiagnosticsInstanceViewResponseOutput struct{ *pulumi.OutputState }

func (BootDiagnosticsInstanceViewResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BootDiagnosticsInstanceViewResponse)(nil)).Elem()
}

func (o BootDiagnosticsInstanceViewResponseOutput) ToBootDiagnosticsInstanceViewResponseOutput() BootDiagnosticsInstanceViewResponseOutput {
	return o
}

func (o BootDiagnosticsInstanceViewResponseOutput) ToBootDiagnosticsInstanceViewResponseOutputWithContext(ctx context.Context) BootDiagnosticsInstanceViewResponseOutput {
	return o
}

func (o BootDiagnosticsInstanceViewResponseOutput) ConsoleScreenshotBlobUri() pulumi.StringOutput {
	return o.ApplyT(func(v BootDiagnosticsInstanceViewResponse) string { return v.ConsoleScreenshotBlobUri }).(pulumi.StringOutput)
}

func (o BootDiagnosticsInstanceViewResponseOutput) SerialConsoleLogBlobUri() pulumi.StringOutput {
	return o.ApplyT(func(v BootDiagnosticsInstanceViewResponse) string { return v.SerialConsoleLogBlobUri }).(pulumi.StringOutput)
}

func (o BootDiagnosticsInstanceViewResponseOutput) Status() InstanceViewStatusResponseOutput {
	return o.ApplyT(func(v BootDiagnosticsInstanceViewResponse) InstanceViewStatusResponse { return v.Status }).(InstanceViewStatusResponseOutput)
}

type BootDiagnosticsInstanceViewResponsePtrOutput struct{ *pulumi.OutputState }

func (BootDiagnosticsInstanceViewResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BootDiagnosticsInstanceViewResponse)(nil)).Elem()
}

func (o BootDiagnosticsInstanceViewResponsePtrOutput) ToBootDiagnosticsInstanceViewResponsePtrOutput() BootDiagnosticsInstanceViewResponsePtrOutput {
	return o
}

func (o BootDiagnosticsInstanceViewResponsePtrOutput) ToBootDiagnosticsInstanceViewResponsePtrOutputWithContext(ctx context.Context) BootDiagnosticsInstanceViewResponsePtrOutput {
	return o
}

func (o BootDiagnosticsInstanceViewResponsePtrOutput) Elem() BootDiagnosticsInstanceViewResponseOutput {
	return o.ApplyT(func(v *BootDiagnosticsInstanceViewResponse) BootDiagnosticsInstanceViewResponse {
		if v != nil {
			return *v
		}
		var ret BootDiagnosticsInstanceViewResponse
		return ret
	}).(BootDiagnosticsInstanceViewResponseOutput)
}

func (o BootDiagnosticsInstanceViewResponsePtrOutput) ConsoleScreenshotBlobUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BootDiagnosticsInstanceViewResponse) *string {
		if v == nil {
			return nil
		}
		return &v.ConsoleScreenshotBlobUri
	}).(pulumi.StringPtrOutput)
}

func (o BootDiagnosticsInstanceViewResponsePtrOutput) SerialConsoleLogBlobUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BootDiagnosticsInstanceViewResponse) *string {
		if v == nil {
			return nil
		}
		return &v.SerialConsoleLogBlobUri
	}).(pulumi.StringPtrOutput)
}

func (o BootDiagnosticsInstanceViewResponsePtrOutput) Status() InstanceViewStatusResponsePtrOutput {
	return o.ApplyT(func(v *BootDiagnosticsInstanceViewResponse) *InstanceViewStatusResponse {
		if v == nil {
			return nil
		}
		return &v.Status
	}).(InstanceViewStatusResponsePtrOutput)
}

type BootDiagnosticsResponse struct {
	Enabled    *bool   `pulumi:"enabled"`
	StorageUri *string `pulumi:"storageUri"`
}

type BootDiagnosticsResponseOutput struct{ *pulumi.OutputState }

func (BootDiagnosticsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BootDiagnosticsResponse)(nil)).Elem()
}

func (o BootDiagnosticsResponseOutput) ToBootDiagnosticsResponseOutput() BootDiagnosticsResponseOutput {
	return o
}

func (o BootDiagnosticsResponseOutput) ToBootDiagnosticsResponseOutputWithContext(ctx context.Context) BootDiagnosticsResponseOutput {
	return o
}

func (o BootDiagnosticsResponseOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v BootDiagnosticsResponse) *bool { return v.Enabled }).(pulumi.BoolPtrOutput)
}

func (o BootDiagnosticsResponseOutput) StorageUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BootDiagnosticsResponse) *string { return v.StorageUri }).(pulumi.StringPtrOutput)
}

type BootDiagnosticsResponsePtrOutput struct{ *pulumi.OutputState }

func (BootDiagnosticsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BootDiagnosticsResponse)(nil)).Elem()
}

func (o BootDiagnosticsResponsePtrOutput) ToBootDiagnosticsResponsePtrOutput() BootDiagnosticsResponsePtrOutput {
	return o
}

func (o BootDiagnosticsResponsePtrOutput) ToBootDiagnosticsResponsePtrOutputWithContext(ctx context.Context) BootDiagnosticsResponsePtrOutput {
	return o
}

func (o BootDiagnosticsResponsePtrOutput) Elem() BootDiagnosticsResponseOutput {
	return o.ApplyT(func(v *BootDiagnosticsResponse) BootDiagnosticsResponse {
		if v != nil {
			return *v
		}
		var ret BootDiagnosticsResponse
		return ret
	}).(BootDiagnosticsResponseOutput)
}

func (o BootDiagnosticsResponsePtrOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *BootDiagnosticsResponse) *bool {
		if v == nil {
			return nil
		}
		return v.Enabled
	}).(pulumi.BoolPtrOutput)
}

func (o BootDiagnosticsResponsePtrOutput) StorageUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BootDiagnosticsResponse) *string {
		if v == nil {
			return nil
		}
		return v.StorageUri
	}).(pulumi.StringPtrOutput)
}

type CreationData struct {
	CreateOption     string              `pulumi:"createOption"`
	ImageReference   *ImageDiskReference `pulumi:"imageReference"`
	SourceResourceId *string             `pulumi:"sourceResourceId"`
	SourceUri        *string             `pulumi:"sourceUri"`
	StorageAccountId *string             `pulumi:"storageAccountId"`
}





type CreationDataInput interface {
	pulumi.Input

	ToCreationDataOutput() CreationDataOutput
	ToCreationDataOutputWithContext(context.Context) CreationDataOutput
}

type CreationDataArgs struct {
	CreateOption     pulumi.StringInput         `pulumi:"createOption"`
	ImageReference   ImageDiskReferencePtrInput `pulumi:"imageReference"`
	SourceResourceId pulumi.StringPtrInput      `pulumi:"sourceResourceId"`
	SourceUri        pulumi.StringPtrInput      `pulumi:"sourceUri"`
	StorageAccountId pulumi.StringPtrInput      `pulumi:"storageAccountId"`
}

func (CreationDataArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CreationData)(nil)).Elem()
}

func (i CreationDataArgs) ToCreationDataOutput() CreationDataOutput {
	return i.ToCreationDataOutputWithContext(context.Background())
}

func (i CreationDataArgs) ToCreationDataOutputWithContext(ctx context.Context) CreationDataOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CreationDataOutput)
}

type CreationDataOutput struct{ *pulumi.OutputState }

func (CreationDataOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CreationData)(nil)).Elem()
}

func (o CreationDataOutput) ToCreationDataOutput() CreationDataOutput {
	return o
}

func (o CreationDataOutput) ToCreationDataOutputWithContext(ctx context.Context) CreationDataOutput {
	return o
}

func (o CreationDataOutput) CreateOption() pulumi.StringOutput {
	return o.ApplyT(func(v CreationData) string { return v.CreateOption }).(pulumi.StringOutput)
}

func (o CreationDataOutput) ImageReference() ImageDiskReferencePtrOutput {
	return o.ApplyT(func(v CreationData) *ImageDiskReference { return v.ImageReference }).(ImageDiskReferencePtrOutput)
}

func (o CreationDataOutput) SourceResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CreationData) *string { return v.SourceResourceId }).(pulumi.StringPtrOutput)
}

func (o CreationDataOutput) SourceUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CreationData) *string { return v.SourceUri }).(pulumi.StringPtrOutput)
}

func (o CreationDataOutput) StorageAccountId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CreationData) *string { return v.StorageAccountId }).(pulumi.StringPtrOutput)
}

type CreationDataResponse struct {
	CreateOption     string                      `pulumi:"createOption"`
	ImageReference   *ImageDiskReferenceResponse `pulumi:"imageReference"`
	SourceResourceId *string                     `pulumi:"sourceResourceId"`
	SourceUri        *string                     `pulumi:"sourceUri"`
	StorageAccountId *string                     `pulumi:"storageAccountId"`
}

type CreationDataResponseOutput struct{ *pulumi.OutputState }

func (CreationDataResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CreationDataResponse)(nil)).Elem()
}

func (o CreationDataResponseOutput) ToCreationDataResponseOutput() CreationDataResponseOutput {
	return o
}

func (o CreationDataResponseOutput) ToCreationDataResponseOutputWithContext(ctx context.Context) CreationDataResponseOutput {
	return o
}

func (o CreationDataResponseOutput) CreateOption() pulumi.StringOutput {
	return o.ApplyT(func(v CreationDataResponse) string { return v.CreateOption }).(pulumi.StringOutput)
}

func (o CreationDataResponseOutput) ImageReference() ImageDiskReferenceResponsePtrOutput {
	return o.ApplyT(func(v CreationDataResponse) *ImageDiskReferenceResponse { return v.ImageReference }).(ImageDiskReferenceResponsePtrOutput)
}

func (o CreationDataResponseOutput) SourceResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CreationDataResponse) *string { return v.SourceResourceId }).(pulumi.StringPtrOutput)
}

func (o CreationDataResponseOutput) SourceUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CreationDataResponse) *string { return v.SourceUri }).(pulumi.StringPtrOutput)
}

func (o CreationDataResponseOutput) StorageAccountId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CreationDataResponse) *string { return v.StorageAccountId }).(pulumi.StringPtrOutput)
}

type DataDisk struct {
	Caching                 *CachingTypes          `pulumi:"caching"`
	CreateOption            string                 `pulumi:"createOption"`
	DiskSizeGB              *int                   `pulumi:"diskSizeGB"`
	Image                   *VirtualHardDisk       `pulumi:"image"`
	Lun                     int                    `pulumi:"lun"`
	ManagedDisk             *ManagedDiskParameters `pulumi:"managedDisk"`
	Name                    *string                `pulumi:"name"`
	Vhd                     *VirtualHardDisk       `pulumi:"vhd"`
	WriteAcceleratorEnabled *bool                  `pulumi:"writeAcceleratorEnabled"`
}





type DataDiskInput interface {
	pulumi.Input

	ToDataDiskOutput() DataDiskOutput
	ToDataDiskOutputWithContext(context.Context) DataDiskOutput
}

type DataDiskArgs struct {
	Caching                 CachingTypesPtrInput          `pulumi:"caching"`
	CreateOption            pulumi.StringInput            `pulumi:"createOption"`
	DiskSizeGB              pulumi.IntPtrInput            `pulumi:"diskSizeGB"`
	Image                   VirtualHardDiskPtrInput       `pulumi:"image"`
	Lun                     pulumi.IntInput               `pulumi:"lun"`
	ManagedDisk             ManagedDiskParametersPtrInput `pulumi:"managedDisk"`
	Name                    pulumi.StringPtrInput         `pulumi:"name"`
	Vhd                     VirtualHardDiskPtrInput       `pulumi:"vhd"`
	WriteAcceleratorEnabled pulumi.BoolPtrInput           `pulumi:"writeAcceleratorEnabled"`
}

func (DataDiskArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DataDisk)(nil)).Elem()
}

func (i DataDiskArgs) ToDataDiskOutput() DataDiskOutput {
	return i.ToDataDiskOutputWithContext(context.Background())
}

func (i DataDiskArgs) ToDataDiskOutputWithContext(ctx context.Context) DataDiskOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataDiskOutput)
}





type DataDiskArrayInput interface {
	pulumi.Input

	ToDataDiskArrayOutput() DataDiskArrayOutput
	ToDataDiskArrayOutputWithContext(context.Context) DataDiskArrayOutput
}

type DataDiskArray []DataDiskInput

func (DataDiskArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DataDisk)(nil)).Elem()
}

func (i DataDiskArray) ToDataDiskArrayOutput() DataDiskArrayOutput {
	return i.ToDataDiskArrayOutputWithContext(context.Background())
}

func (i DataDiskArray) ToDataDiskArrayOutputWithContext(ctx context.Context) DataDiskArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataDiskArrayOutput)
}

type DataDiskOutput struct{ *pulumi.OutputState }

func (DataDiskOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DataDisk)(nil)).Elem()
}

func (o DataDiskOutput) ToDataDiskOutput() DataDiskOutput {
	return o
}

func (o DataDiskOutput) ToDataDiskOutputWithContext(ctx context.Context) DataDiskOutput {
	return o
}

func (o DataDiskOutput) Caching() CachingTypesPtrOutput {
	return o.ApplyT(func(v DataDisk) *CachingTypes { return v.Caching }).(CachingTypesPtrOutput)
}

func (o DataDiskOutput) CreateOption() pulumi.StringOutput {
	return o.ApplyT(func(v DataDisk) string { return v.CreateOption }).(pulumi.StringOutput)
}

func (o DataDiskOutput) DiskSizeGB() pulumi.IntPtrOutput {
	return o.ApplyT(func(v DataDisk) *int { return v.DiskSizeGB }).(pulumi.IntPtrOutput)
}

func (o DataDiskOutput) Image() VirtualHardDiskPtrOutput {
	return o.ApplyT(func(v DataDisk) *VirtualHardDisk { return v.Image }).(VirtualHardDiskPtrOutput)
}

func (o DataDiskOutput) Lun() pulumi.IntOutput {
	return o.ApplyT(func(v DataDisk) int { return v.Lun }).(pulumi.IntOutput)
}

func (o DataDiskOutput) ManagedDisk() ManagedDiskParametersPtrOutput {
	return o.ApplyT(func(v DataDisk) *ManagedDiskParameters { return v.ManagedDisk }).(ManagedDiskParametersPtrOutput)
}

func (o DataDiskOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DataDisk) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o DataDiskOutput) Vhd() VirtualHardDiskPtrOutput {
	return o.ApplyT(func(v DataDisk) *VirtualHardDisk { return v.Vhd }).(VirtualHardDiskPtrOutput)
}

func (o DataDiskOutput) WriteAcceleratorEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v DataDisk) *bool { return v.WriteAcceleratorEnabled }).(pulumi.BoolPtrOutput)
}

type DataDiskArrayOutput struct{ *pulumi.OutputState }

func (DataDiskArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DataDisk)(nil)).Elem()
}

func (o DataDiskArrayOutput) ToDataDiskArrayOutput() DataDiskArrayOutput {
	return o
}

func (o DataDiskArrayOutput) ToDataDiskArrayOutputWithContext(ctx context.Context) DataDiskArrayOutput {
	return o
}

func (o DataDiskArrayOutput) Index(i pulumi.IntInput) DataDiskOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) DataDisk {
		return vs[0].([]DataDisk)[vs[1].(int)]
	}).(DataDiskOutput)
}

type DataDiskResponse struct {
	Caching                 *string                        `pulumi:"caching"`
	CreateOption            string                         `pulumi:"createOption"`
	DiskSizeGB              *int                           `pulumi:"diskSizeGB"`
	Image                   *VirtualHardDiskResponse       `pulumi:"image"`
	Lun                     int                            `pulumi:"lun"`
	ManagedDisk             *ManagedDiskParametersResponse `pulumi:"managedDisk"`
	Name                    *string                        `pulumi:"name"`
	Vhd                     *VirtualHardDiskResponse       `pulumi:"vhd"`
	WriteAcceleratorEnabled *bool                          `pulumi:"writeAcceleratorEnabled"`
}

type DataDiskResponseOutput struct{ *pulumi.OutputState }

func (DataDiskResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DataDiskResponse)(nil)).Elem()
}

func (o DataDiskResponseOutput) ToDataDiskResponseOutput() DataDiskResponseOutput {
	return o
}

func (o DataDiskResponseOutput) ToDataDiskResponseOutputWithContext(ctx context.Context) DataDiskResponseOutput {
	return o
}

func (o DataDiskResponseOutput) Caching() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DataDiskResponse) *string { return v.Caching }).(pulumi.StringPtrOutput)
}

func (o DataDiskResponseOutput) CreateOption() pulumi.StringOutput {
	return o.ApplyT(func(v DataDiskResponse) string { return v.CreateOption }).(pulumi.StringOutput)
}

func (o DataDiskResponseOutput) DiskSizeGB() pulumi.IntPtrOutput {
	return o.ApplyT(func(v DataDiskResponse) *int { return v.DiskSizeGB }).(pulumi.IntPtrOutput)
}

func (o DataDiskResponseOutput) Image() VirtualHardDiskResponsePtrOutput {
	return o.ApplyT(func(v DataDiskResponse) *VirtualHardDiskResponse { return v.Image }).(VirtualHardDiskResponsePtrOutput)
}

func (o DataDiskResponseOutput) Lun() pulumi.IntOutput {
	return o.ApplyT(func(v DataDiskResponse) int { return v.Lun }).(pulumi.IntOutput)
}

func (o DataDiskResponseOutput) ManagedDisk() ManagedDiskParametersResponsePtrOutput {
	return o.ApplyT(func(v DataDiskResponse) *ManagedDiskParametersResponse { return v.ManagedDisk }).(ManagedDiskParametersResponsePtrOutput)
}

func (o DataDiskResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DataDiskResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o DataDiskResponseOutput) Vhd() VirtualHardDiskResponsePtrOutput {
	return o.ApplyT(func(v DataDiskResponse) *VirtualHardDiskResponse { return v.Vhd }).(VirtualHardDiskResponsePtrOutput)
}

func (o DataDiskResponseOutput) WriteAcceleratorEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v DataDiskResponse) *bool { return v.WriteAcceleratorEnabled }).(pulumi.BoolPtrOutput)
}

type DataDiskResponseArrayOutput struct{ *pulumi.OutputState }

func (DataDiskResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DataDiskResponse)(nil)).Elem()
}

func (o DataDiskResponseArrayOutput) ToDataDiskResponseArrayOutput() DataDiskResponseArrayOutput {
	return o
}

func (o DataDiskResponseArrayOutput) ToDataDiskResponseArrayOutputWithContext(ctx context.Context) DataDiskResponseArrayOutput {
	return o
}

func (o DataDiskResponseArrayOutput) Index(i pulumi.IntInput) DataDiskResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) DataDiskResponse {
		return vs[0].([]DataDiskResponse)[vs[1].(int)]
	}).(DataDiskResponseOutput)
}

type DiagnosticsProfile struct {
	BootDiagnostics *BootDiagnostics `pulumi:"bootDiagnostics"`
}





type DiagnosticsProfileInput interface {
	pulumi.Input

	ToDiagnosticsProfileOutput() DiagnosticsProfileOutput
	ToDiagnosticsProfileOutputWithContext(context.Context) DiagnosticsProfileOutput
}

type DiagnosticsProfileArgs struct {
	BootDiagnostics BootDiagnosticsPtrInput `pulumi:"bootDiagnostics"`
}

func (DiagnosticsProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DiagnosticsProfile)(nil)).Elem()
}

func (i DiagnosticsProfileArgs) ToDiagnosticsProfileOutput() DiagnosticsProfileOutput {
	return i.ToDiagnosticsProfileOutputWithContext(context.Background())
}

func (i DiagnosticsProfileArgs) ToDiagnosticsProfileOutputWithContext(ctx context.Context) DiagnosticsProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DiagnosticsProfileOutput)
}

func (i DiagnosticsProfileArgs) ToDiagnosticsProfilePtrOutput() DiagnosticsProfilePtrOutput {
	return i.ToDiagnosticsProfilePtrOutputWithContext(context.Background())
}

func (i DiagnosticsProfileArgs) ToDiagnosticsProfilePtrOutputWithContext(ctx context.Context) DiagnosticsProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DiagnosticsProfileOutput).ToDiagnosticsProfilePtrOutputWithContext(ctx)
}









type DiagnosticsProfilePtrInput interface {
	pulumi.Input

	ToDiagnosticsProfilePtrOutput() DiagnosticsProfilePtrOutput
	ToDiagnosticsProfilePtrOutputWithContext(context.Context) DiagnosticsProfilePtrOutput
}

type diagnosticsProfilePtrType DiagnosticsProfileArgs

func DiagnosticsProfilePtr(v *DiagnosticsProfileArgs) DiagnosticsProfilePtrInput {
	return (*diagnosticsProfilePtrType)(v)
}

func (*diagnosticsProfilePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DiagnosticsProfile)(nil)).Elem()
}

func (i *diagnosticsProfilePtrType) ToDiagnosticsProfilePtrOutput() DiagnosticsProfilePtrOutput {
	return i.ToDiagnosticsProfilePtrOutputWithContext(context.Background())
}

func (i *diagnosticsProfilePtrType) ToDiagnosticsProfilePtrOutputWithContext(ctx context.Context) DiagnosticsProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DiagnosticsProfilePtrOutput)
}

type DiagnosticsProfileOutput struct{ *pulumi.OutputState }

func (DiagnosticsProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DiagnosticsProfile)(nil)).Elem()
}

func (o DiagnosticsProfileOutput) ToDiagnosticsProfileOutput() DiagnosticsProfileOutput {
	return o
}

func (o DiagnosticsProfileOutput) ToDiagnosticsProfileOutputWithContext(ctx context.Context) DiagnosticsProfileOutput {
	return o
}

func (o DiagnosticsProfileOutput) ToDiagnosticsProfilePtrOutput() DiagnosticsProfilePtrOutput {
	return o.ToDiagnosticsProfilePtrOutputWithContext(context.Background())
}

func (o DiagnosticsProfileOutput) ToDiagnosticsProfilePtrOutputWithContext(ctx context.Context) DiagnosticsProfilePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DiagnosticsProfile) *DiagnosticsProfile {
		return &v
	}).(DiagnosticsProfilePtrOutput)
}

func (o DiagnosticsProfileOutput) BootDiagnostics() BootDiagnosticsPtrOutput {
	return o.ApplyT(func(v DiagnosticsProfile) *BootDiagnostics { return v.BootDiagnostics }).(BootDiagnosticsPtrOutput)
}

type DiagnosticsProfilePtrOutput struct{ *pulumi.OutputState }

func (DiagnosticsProfilePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DiagnosticsProfile)(nil)).Elem()
}

func (o DiagnosticsProfilePtrOutput) ToDiagnosticsProfilePtrOutput() DiagnosticsProfilePtrOutput {
	return o
}

func (o DiagnosticsProfilePtrOutput) ToDiagnosticsProfilePtrOutputWithContext(ctx context.Context) DiagnosticsProfilePtrOutput {
	return o
}

func (o DiagnosticsProfilePtrOutput) Elem() DiagnosticsProfileOutput {
	return o.ApplyT(func(v *DiagnosticsProfile) DiagnosticsProfile {
		if v != nil {
			return *v
		}
		var ret DiagnosticsProfile
		return ret
	}).(DiagnosticsProfileOutput)
}

func (o DiagnosticsProfilePtrOutput) BootDiagnostics() BootDiagnosticsPtrOutput {
	return o.ApplyT(func(v *DiagnosticsProfile) *BootDiagnostics {
		if v == nil {
			return nil
		}
		return v.BootDiagnostics
	}).(BootDiagnosticsPtrOutput)
}

type DiagnosticsProfileResponse struct {
	BootDiagnostics *BootDiagnosticsResponse `pulumi:"bootDiagnostics"`
}

type DiagnosticsProfileResponseOutput struct{ *pulumi.OutputState }

func (DiagnosticsProfileResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DiagnosticsProfileResponse)(nil)).Elem()
}

func (o DiagnosticsProfileResponseOutput) ToDiagnosticsProfileResponseOutput() DiagnosticsProfileResponseOutput {
	return o
}

func (o DiagnosticsProfileResponseOutput) ToDiagnosticsProfileResponseOutputWithContext(ctx context.Context) DiagnosticsProfileResponseOutput {
	return o
}

func (o DiagnosticsProfileResponseOutput) BootDiagnostics() BootDiagnosticsResponsePtrOutput {
	return o.ApplyT(func(v DiagnosticsProfileResponse) *BootDiagnosticsResponse { return v.BootDiagnostics }).(BootDiagnosticsResponsePtrOutput)
}

type DiagnosticsProfileResponsePtrOutput struct{ *pulumi.OutputState }

func (DiagnosticsProfileResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DiagnosticsProfileResponse)(nil)).Elem()
}

func (o DiagnosticsProfileResponsePtrOutput) ToDiagnosticsProfileResponsePtrOutput() DiagnosticsProfileResponsePtrOutput {
	return o
}

func (o DiagnosticsProfileResponsePtrOutput) ToDiagnosticsProfileResponsePtrOutputWithContext(ctx context.Context) DiagnosticsProfileResponsePtrOutput {
	return o
}

func (o DiagnosticsProfileResponsePtrOutput) Elem() DiagnosticsProfileResponseOutput {
	return o.ApplyT(func(v *DiagnosticsProfileResponse) DiagnosticsProfileResponse {
		if v != nil {
			return *v
		}
		var ret DiagnosticsProfileResponse
		return ret
	}).(DiagnosticsProfileResponseOutput)
}

func (o DiagnosticsProfileResponsePtrOutput) BootDiagnostics() BootDiagnosticsResponsePtrOutput {
	return o.ApplyT(func(v *DiagnosticsProfileResponse) *BootDiagnosticsResponse {
		if v == nil {
			return nil
		}
		return v.BootDiagnostics
	}).(BootDiagnosticsResponsePtrOutput)
}

type DiskEncryptionSettings struct {
	DiskEncryptionKey *KeyVaultSecretReference `pulumi:"diskEncryptionKey"`
	Enabled           *bool                    `pulumi:"enabled"`
	KeyEncryptionKey  *KeyVaultKeyReference    `pulumi:"keyEncryptionKey"`
}





type DiskEncryptionSettingsInput interface {
	pulumi.Input

	ToDiskEncryptionSettingsOutput() DiskEncryptionSettingsOutput
	ToDiskEncryptionSettingsOutputWithContext(context.Context) DiskEncryptionSettingsOutput
}

type DiskEncryptionSettingsArgs struct {
	DiskEncryptionKey KeyVaultSecretReferencePtrInput `pulumi:"diskEncryptionKey"`
	Enabled           pulumi.BoolPtrInput             `pulumi:"enabled"`
	KeyEncryptionKey  KeyVaultKeyReferencePtrInput    `pulumi:"keyEncryptionKey"`
}

func (DiskEncryptionSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DiskEncryptionSettings)(nil)).Elem()
}

func (i DiskEncryptionSettingsArgs) ToDiskEncryptionSettingsOutput() DiskEncryptionSettingsOutput {
	return i.ToDiskEncryptionSettingsOutputWithContext(context.Background())
}

func (i DiskEncryptionSettingsArgs) ToDiskEncryptionSettingsOutputWithContext(ctx context.Context) DiskEncryptionSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DiskEncryptionSettingsOutput)
}

func (i DiskEncryptionSettingsArgs) ToDiskEncryptionSettingsPtrOutput() DiskEncryptionSettingsPtrOutput {
	return i.ToDiskEncryptionSettingsPtrOutputWithContext(context.Background())
}

func (i DiskEncryptionSettingsArgs) ToDiskEncryptionSettingsPtrOutputWithContext(ctx context.Context) DiskEncryptionSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DiskEncryptionSettingsOutput).ToDiskEncryptionSettingsPtrOutputWithContext(ctx)
}









type DiskEncryptionSettingsPtrInput interface {
	pulumi.Input

	ToDiskEncryptionSettingsPtrOutput() DiskEncryptionSettingsPtrOutput
	ToDiskEncryptionSettingsPtrOutputWithContext(context.Context) DiskEncryptionSettingsPtrOutput
}

type diskEncryptionSettingsPtrType DiskEncryptionSettingsArgs

func DiskEncryptionSettingsPtr(v *DiskEncryptionSettingsArgs) DiskEncryptionSettingsPtrInput {
	return (*diskEncryptionSettingsPtrType)(v)
}

func (*diskEncryptionSettingsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DiskEncryptionSettings)(nil)).Elem()
}

func (i *diskEncryptionSettingsPtrType) ToDiskEncryptionSettingsPtrOutput() DiskEncryptionSettingsPtrOutput {
	return i.ToDiskEncryptionSettingsPtrOutputWithContext(context.Background())
}

func (i *diskEncryptionSettingsPtrType) ToDiskEncryptionSettingsPtrOutputWithContext(ctx context.Context) DiskEncryptionSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DiskEncryptionSettingsPtrOutput)
}

type DiskEncryptionSettingsOutput struct{ *pulumi.OutputState }

func (DiskEncryptionSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DiskEncryptionSettings)(nil)).Elem()
}

func (o DiskEncryptionSettingsOutput) ToDiskEncryptionSettingsOutput() DiskEncryptionSettingsOutput {
	return o
}

func (o DiskEncryptionSettingsOutput) ToDiskEncryptionSettingsOutputWithContext(ctx context.Context) DiskEncryptionSettingsOutput {
	return o
}

func (o DiskEncryptionSettingsOutput) ToDiskEncryptionSettingsPtrOutput() DiskEncryptionSettingsPtrOutput {
	return o.ToDiskEncryptionSettingsPtrOutputWithContext(context.Background())
}

func (o DiskEncryptionSettingsOutput) ToDiskEncryptionSettingsPtrOutputWithContext(ctx context.Context) DiskEncryptionSettingsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DiskEncryptionSettings) *DiskEncryptionSettings {
		return &v
	}).(DiskEncryptionSettingsPtrOutput)
}

func (o DiskEncryptionSettingsOutput) DiskEncryptionKey() KeyVaultSecretReferencePtrOutput {
	return o.ApplyT(func(v DiskEncryptionSettings) *KeyVaultSecretReference { return v.DiskEncryptionKey }).(KeyVaultSecretReferencePtrOutput)
}

func (o DiskEncryptionSettingsOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v DiskEncryptionSettings) *bool { return v.Enabled }).(pulumi.BoolPtrOutput)
}

func (o DiskEncryptionSettingsOutput) KeyEncryptionKey() KeyVaultKeyReferencePtrOutput {
	return o.ApplyT(func(v DiskEncryptionSettings) *KeyVaultKeyReference { return v.KeyEncryptionKey }).(KeyVaultKeyReferencePtrOutput)
}

type DiskEncryptionSettingsPtrOutput struct{ *pulumi.OutputState }

func (DiskEncryptionSettingsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DiskEncryptionSettings)(nil)).Elem()
}

func (o DiskEncryptionSettingsPtrOutput) ToDiskEncryptionSettingsPtrOutput() DiskEncryptionSettingsPtrOutput {
	return o
}

func (o DiskEncryptionSettingsPtrOutput) ToDiskEncryptionSettingsPtrOutputWithContext(ctx context.Context) DiskEncryptionSettingsPtrOutput {
	return o
}

func (o DiskEncryptionSettingsPtrOutput) Elem() DiskEncryptionSettingsOutput {
	return o.ApplyT(func(v *DiskEncryptionSettings) DiskEncryptionSettings {
		if v != nil {
			return *v
		}
		var ret DiskEncryptionSettings
		return ret
	}).(DiskEncryptionSettingsOutput)
}

func (o DiskEncryptionSettingsPtrOutput) DiskEncryptionKey() KeyVaultSecretReferencePtrOutput {
	return o.ApplyT(func(v *DiskEncryptionSettings) *KeyVaultSecretReference {
		if v == nil {
			return nil
		}
		return v.DiskEncryptionKey
	}).(KeyVaultSecretReferencePtrOutput)
}

func (o DiskEncryptionSettingsPtrOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *DiskEncryptionSettings) *bool {
		if v == nil {
			return nil
		}
		return v.Enabled
	}).(pulumi.BoolPtrOutput)
}

func (o DiskEncryptionSettingsPtrOutput) KeyEncryptionKey() KeyVaultKeyReferencePtrOutput {
	return o.ApplyT(func(v *DiskEncryptionSettings) *KeyVaultKeyReference {
		if v == nil {
			return nil
		}
		return v.KeyEncryptionKey
	}).(KeyVaultKeyReferencePtrOutput)
}

type DiskEncryptionSettingsResponse struct {
	DiskEncryptionKey *KeyVaultSecretReferenceResponse `pulumi:"diskEncryptionKey"`
	Enabled           *bool                            `pulumi:"enabled"`
	KeyEncryptionKey  *KeyVaultKeyReferenceResponse    `pulumi:"keyEncryptionKey"`
}

type DiskEncryptionSettingsResponseOutput struct{ *pulumi.OutputState }

func (DiskEncryptionSettingsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DiskEncryptionSettingsResponse)(nil)).Elem()
}

func (o DiskEncryptionSettingsResponseOutput) ToDiskEncryptionSettingsResponseOutput() DiskEncryptionSettingsResponseOutput {
	return o
}

func (o DiskEncryptionSettingsResponseOutput) ToDiskEncryptionSettingsResponseOutputWithContext(ctx context.Context) DiskEncryptionSettingsResponseOutput {
	return o
}

func (o DiskEncryptionSettingsResponseOutput) DiskEncryptionKey() KeyVaultSecretReferenceResponsePtrOutput {
	return o.ApplyT(func(v DiskEncryptionSettingsResponse) *KeyVaultSecretReferenceResponse { return v.DiskEncryptionKey }).(KeyVaultSecretReferenceResponsePtrOutput)
}

func (o DiskEncryptionSettingsResponseOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v DiskEncryptionSettingsResponse) *bool { return v.Enabled }).(pulumi.BoolPtrOutput)
}

func (o DiskEncryptionSettingsResponseOutput) KeyEncryptionKey() KeyVaultKeyReferenceResponsePtrOutput {
	return o.ApplyT(func(v DiskEncryptionSettingsResponse) *KeyVaultKeyReferenceResponse { return v.KeyEncryptionKey }).(KeyVaultKeyReferenceResponsePtrOutput)
}

type DiskEncryptionSettingsResponsePtrOutput struct{ *pulumi.OutputState }

func (DiskEncryptionSettingsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DiskEncryptionSettingsResponse)(nil)).Elem()
}

func (o DiskEncryptionSettingsResponsePtrOutput) ToDiskEncryptionSettingsResponsePtrOutput() DiskEncryptionSettingsResponsePtrOutput {
	return o
}

func (o DiskEncryptionSettingsResponsePtrOutput) ToDiskEncryptionSettingsResponsePtrOutputWithContext(ctx context.Context) DiskEncryptionSettingsResponsePtrOutput {
	return o
}

func (o DiskEncryptionSettingsResponsePtrOutput) Elem() DiskEncryptionSettingsResponseOutput {
	return o.ApplyT(func(v *DiskEncryptionSettingsResponse) DiskEncryptionSettingsResponse {
		if v != nil {
			return *v
		}
		var ret DiskEncryptionSettingsResponse
		return ret
	}).(DiskEncryptionSettingsResponseOutput)
}

func (o DiskEncryptionSettingsResponsePtrOutput) DiskEncryptionKey() KeyVaultSecretReferenceResponsePtrOutput {
	return o.ApplyT(func(v *DiskEncryptionSettingsResponse) *KeyVaultSecretReferenceResponse {
		if v == nil {
			return nil
		}
		return v.DiskEncryptionKey
	}).(KeyVaultSecretReferenceResponsePtrOutput)
}

func (o DiskEncryptionSettingsResponsePtrOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *DiskEncryptionSettingsResponse) *bool {
		if v == nil {
			return nil
		}
		return v.Enabled
	}).(pulumi.BoolPtrOutput)
}

func (o DiskEncryptionSettingsResponsePtrOutput) KeyEncryptionKey() KeyVaultKeyReferenceResponsePtrOutput {
	return o.ApplyT(func(v *DiskEncryptionSettingsResponse) *KeyVaultKeyReferenceResponse {
		if v == nil {
			return nil
		}
		return v.KeyEncryptionKey
	}).(KeyVaultKeyReferenceResponsePtrOutput)
}

type DiskEncryptionSettingsResponseArrayOutput struct{ *pulumi.OutputState }

func (DiskEncryptionSettingsResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DiskEncryptionSettingsResponse)(nil)).Elem()
}

func (o DiskEncryptionSettingsResponseArrayOutput) ToDiskEncryptionSettingsResponseArrayOutput() DiskEncryptionSettingsResponseArrayOutput {
	return o
}

func (o DiskEncryptionSettingsResponseArrayOutput) ToDiskEncryptionSettingsResponseArrayOutputWithContext(ctx context.Context) DiskEncryptionSettingsResponseArrayOutput {
	return o
}

func (o DiskEncryptionSettingsResponseArrayOutput) Index(i pulumi.IntInput) DiskEncryptionSettingsResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) DiskEncryptionSettingsResponse {
		return vs[0].([]DiskEncryptionSettingsResponse)[vs[1].(int)]
	}).(DiskEncryptionSettingsResponseOutput)
}

type DiskInstanceViewResponse struct {
	EncryptionSettings []DiskEncryptionSettingsResponse `pulumi:"encryptionSettings"`
	Name               *string                          `pulumi:"name"`
	Statuses           []InstanceViewStatusResponse     `pulumi:"statuses"`
}

type DiskInstanceViewResponseOutput struct{ *pulumi.OutputState }

func (DiskInstanceViewResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DiskInstanceViewResponse)(nil)).Elem()
}

func (o DiskInstanceViewResponseOutput) ToDiskInstanceViewResponseOutput() DiskInstanceViewResponseOutput {
	return o
}

func (o DiskInstanceViewResponseOutput) ToDiskInstanceViewResponseOutputWithContext(ctx context.Context) DiskInstanceViewResponseOutput {
	return o
}

func (o DiskInstanceViewResponseOutput) EncryptionSettings() DiskEncryptionSettingsResponseArrayOutput {
	return o.ApplyT(func(v DiskInstanceViewResponse) []DiskEncryptionSettingsResponse { return v.EncryptionSettings }).(DiskEncryptionSettingsResponseArrayOutput)
}

func (o DiskInstanceViewResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DiskInstanceViewResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o DiskInstanceViewResponseOutput) Statuses() InstanceViewStatusResponseArrayOutput {
	return o.ApplyT(func(v DiskInstanceViewResponse) []InstanceViewStatusResponse { return v.Statuses }).(InstanceViewStatusResponseArrayOutput)
}

type DiskInstanceViewResponseArrayOutput struct{ *pulumi.OutputState }

func (DiskInstanceViewResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DiskInstanceViewResponse)(nil)).Elem()
}

func (o DiskInstanceViewResponseArrayOutput) ToDiskInstanceViewResponseArrayOutput() DiskInstanceViewResponseArrayOutput {
	return o
}

func (o DiskInstanceViewResponseArrayOutput) ToDiskInstanceViewResponseArrayOutputWithContext(ctx context.Context) DiskInstanceViewResponseArrayOutput {
	return o
}

func (o DiskInstanceViewResponseArrayOutput) Index(i pulumi.IntInput) DiskInstanceViewResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) DiskInstanceViewResponse {
		return vs[0].([]DiskInstanceViewResponse)[vs[1].(int)]
	}).(DiskInstanceViewResponseOutput)
}

type DiskSku struct {
	Name *string `pulumi:"name"`
}





type DiskSkuInput interface {
	pulumi.Input

	ToDiskSkuOutput() DiskSkuOutput
	ToDiskSkuOutputWithContext(context.Context) DiskSkuOutput
}

type DiskSkuArgs struct {
	Name pulumi.StringPtrInput `pulumi:"name"`
}

func (DiskSkuArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DiskSku)(nil)).Elem()
}

func (i DiskSkuArgs) ToDiskSkuOutput() DiskSkuOutput {
	return i.ToDiskSkuOutputWithContext(context.Background())
}

func (i DiskSkuArgs) ToDiskSkuOutputWithContext(ctx context.Context) DiskSkuOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DiskSkuOutput)
}

func (i DiskSkuArgs) ToDiskSkuPtrOutput() DiskSkuPtrOutput {
	return i.ToDiskSkuPtrOutputWithContext(context.Background())
}

func (i DiskSkuArgs) ToDiskSkuPtrOutputWithContext(ctx context.Context) DiskSkuPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DiskSkuOutput).ToDiskSkuPtrOutputWithContext(ctx)
}









type DiskSkuPtrInput interface {
	pulumi.Input

	ToDiskSkuPtrOutput() DiskSkuPtrOutput
	ToDiskSkuPtrOutputWithContext(context.Context) DiskSkuPtrOutput
}

type diskSkuPtrType DiskSkuArgs

func DiskSkuPtr(v *DiskSkuArgs) DiskSkuPtrInput {
	return (*diskSkuPtrType)(v)
}

func (*diskSkuPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DiskSku)(nil)).Elem()
}

func (i *diskSkuPtrType) ToDiskSkuPtrOutput() DiskSkuPtrOutput {
	return i.ToDiskSkuPtrOutputWithContext(context.Background())
}

func (i *diskSkuPtrType) ToDiskSkuPtrOutputWithContext(ctx context.Context) DiskSkuPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DiskSkuPtrOutput)
}

type DiskSkuOutput struct{ *pulumi.OutputState }

func (DiskSkuOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DiskSku)(nil)).Elem()
}

func (o DiskSkuOutput) ToDiskSkuOutput() DiskSkuOutput {
	return o
}

func (o DiskSkuOutput) ToDiskSkuOutputWithContext(ctx context.Context) DiskSkuOutput {
	return o
}

func (o DiskSkuOutput) ToDiskSkuPtrOutput() DiskSkuPtrOutput {
	return o.ToDiskSkuPtrOutputWithContext(context.Background())
}

func (o DiskSkuOutput) ToDiskSkuPtrOutputWithContext(ctx context.Context) DiskSkuPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DiskSku) *DiskSku {
		return &v
	}).(DiskSkuPtrOutput)
}

func (o DiskSkuOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DiskSku) *string { return v.Name }).(pulumi.StringPtrOutput)
}

type DiskSkuPtrOutput struct{ *pulumi.OutputState }

func (DiskSkuPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DiskSku)(nil)).Elem()
}

func (o DiskSkuPtrOutput) ToDiskSkuPtrOutput() DiskSkuPtrOutput {
	return o
}

func (o DiskSkuPtrOutput) ToDiskSkuPtrOutputWithContext(ctx context.Context) DiskSkuPtrOutput {
	return o
}

func (o DiskSkuPtrOutput) Elem() DiskSkuOutput {
	return o.ApplyT(func(v *DiskSku) DiskSku {
		if v != nil {
			return *v
		}
		var ret DiskSku
		return ret
	}).(DiskSkuOutput)
}

func (o DiskSkuPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DiskSku) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

type DiskSkuResponse struct {
	Name *string `pulumi:"name"`
	Tier string  `pulumi:"tier"`
}


func (val *DiskSkuResponse) Defaults() *DiskSkuResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Tier) {
		tmp.Tier = "Standard"
	}
	return &tmp
}

type DiskSkuResponseOutput struct{ *pulumi.OutputState }

func (DiskSkuResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DiskSkuResponse)(nil)).Elem()
}

func (o DiskSkuResponseOutput) ToDiskSkuResponseOutput() DiskSkuResponseOutput {
	return o
}

func (o DiskSkuResponseOutput) ToDiskSkuResponseOutputWithContext(ctx context.Context) DiskSkuResponseOutput {
	return o
}

func (o DiskSkuResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DiskSkuResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o DiskSkuResponseOutput) Tier() pulumi.StringOutput {
	return o.ApplyT(func(v DiskSkuResponse) string { return v.Tier }).(pulumi.StringOutput)
}

type DiskSkuResponsePtrOutput struct{ *pulumi.OutputState }

func (DiskSkuResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DiskSkuResponse)(nil)).Elem()
}

func (o DiskSkuResponsePtrOutput) ToDiskSkuResponsePtrOutput() DiskSkuResponsePtrOutput {
	return o
}

func (o DiskSkuResponsePtrOutput) ToDiskSkuResponsePtrOutputWithContext(ctx context.Context) DiskSkuResponsePtrOutput {
	return o
}

func (o DiskSkuResponsePtrOutput) Elem() DiskSkuResponseOutput {
	return o.ApplyT(func(v *DiskSkuResponse) DiskSkuResponse {
		if v != nil {
			return *v
		}
		var ret DiskSkuResponse
		return ret
	}).(DiskSkuResponseOutput)
}

func (o DiskSkuResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DiskSkuResponse) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o DiskSkuResponsePtrOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DiskSkuResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Tier
	}).(pulumi.StringPtrOutput)
}

type EncryptionSettings struct {
	DiskEncryptionKey *KeyVaultAndSecretReference `pulumi:"diskEncryptionKey"`
	Enabled           *bool                       `pulumi:"enabled"`
	KeyEncryptionKey  *KeyVaultAndKeyReference    `pulumi:"keyEncryptionKey"`
}





type EncryptionSettingsInput interface {
	pulumi.Input

	ToEncryptionSettingsOutput() EncryptionSettingsOutput
	ToEncryptionSettingsOutputWithContext(context.Context) EncryptionSettingsOutput
}

type EncryptionSettingsArgs struct {
	DiskEncryptionKey KeyVaultAndSecretReferencePtrInput `pulumi:"diskEncryptionKey"`
	Enabled           pulumi.BoolPtrInput                `pulumi:"enabled"`
	KeyEncryptionKey  KeyVaultAndKeyReferencePtrInput    `pulumi:"keyEncryptionKey"`
}

func (EncryptionSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EncryptionSettings)(nil)).Elem()
}

func (i EncryptionSettingsArgs) ToEncryptionSettingsOutput() EncryptionSettingsOutput {
	return i.ToEncryptionSettingsOutputWithContext(context.Background())
}

func (i EncryptionSettingsArgs) ToEncryptionSettingsOutputWithContext(ctx context.Context) EncryptionSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EncryptionSettingsOutput)
}

func (i EncryptionSettingsArgs) ToEncryptionSettingsPtrOutput() EncryptionSettingsPtrOutput {
	return i.ToEncryptionSettingsPtrOutputWithContext(context.Background())
}

func (i EncryptionSettingsArgs) ToEncryptionSettingsPtrOutputWithContext(ctx context.Context) EncryptionSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EncryptionSettingsOutput).ToEncryptionSettingsPtrOutputWithContext(ctx)
}









type EncryptionSettingsPtrInput interface {
	pulumi.Input

	ToEncryptionSettingsPtrOutput() EncryptionSettingsPtrOutput
	ToEncryptionSettingsPtrOutputWithContext(context.Context) EncryptionSettingsPtrOutput
}

type encryptionSettingsPtrType EncryptionSettingsArgs

func EncryptionSettingsPtr(v *EncryptionSettingsArgs) EncryptionSettingsPtrInput {
	return (*encryptionSettingsPtrType)(v)
}

func (*encryptionSettingsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**EncryptionSettings)(nil)).Elem()
}

func (i *encryptionSettingsPtrType) ToEncryptionSettingsPtrOutput() EncryptionSettingsPtrOutput {
	return i.ToEncryptionSettingsPtrOutputWithContext(context.Background())
}

func (i *encryptionSettingsPtrType) ToEncryptionSettingsPtrOutputWithContext(ctx context.Context) EncryptionSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EncryptionSettingsPtrOutput)
}

type EncryptionSettingsOutput struct{ *pulumi.OutputState }

func (EncryptionSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EncryptionSettings)(nil)).Elem()
}

func (o EncryptionSettingsOutput) ToEncryptionSettingsOutput() EncryptionSettingsOutput {
	return o
}

func (o EncryptionSettingsOutput) ToEncryptionSettingsOutputWithContext(ctx context.Context) EncryptionSettingsOutput {
	return o
}

func (o EncryptionSettingsOutput) ToEncryptionSettingsPtrOutput() EncryptionSettingsPtrOutput {
	return o.ToEncryptionSettingsPtrOutputWithContext(context.Background())
}

func (o EncryptionSettingsOutput) ToEncryptionSettingsPtrOutputWithContext(ctx context.Context) EncryptionSettingsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v EncryptionSettings) *EncryptionSettings {
		return &v
	}).(EncryptionSettingsPtrOutput)
}

func (o EncryptionSettingsOutput) DiskEncryptionKey() KeyVaultAndSecretReferencePtrOutput {
	return o.ApplyT(func(v EncryptionSettings) *KeyVaultAndSecretReference { return v.DiskEncryptionKey }).(KeyVaultAndSecretReferencePtrOutput)
}

func (o EncryptionSettingsOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v EncryptionSettings) *bool { return v.Enabled }).(pulumi.BoolPtrOutput)
}

func (o EncryptionSettingsOutput) KeyEncryptionKey() KeyVaultAndKeyReferencePtrOutput {
	return o.ApplyT(func(v EncryptionSettings) *KeyVaultAndKeyReference { return v.KeyEncryptionKey }).(KeyVaultAndKeyReferencePtrOutput)
}

type EncryptionSettingsPtrOutput struct{ *pulumi.OutputState }

func (EncryptionSettingsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EncryptionSettings)(nil)).Elem()
}

func (o EncryptionSettingsPtrOutput) ToEncryptionSettingsPtrOutput() EncryptionSettingsPtrOutput {
	return o
}

func (o EncryptionSettingsPtrOutput) ToEncryptionSettingsPtrOutputWithContext(ctx context.Context) EncryptionSettingsPtrOutput {
	return o
}

func (o EncryptionSettingsPtrOutput) Elem() EncryptionSettingsOutput {
	return o.ApplyT(func(v *EncryptionSettings) EncryptionSettings {
		if v != nil {
			return *v
		}
		var ret EncryptionSettings
		return ret
	}).(EncryptionSettingsOutput)
}

func (o EncryptionSettingsPtrOutput) DiskEncryptionKey() KeyVaultAndSecretReferencePtrOutput {
	return o.ApplyT(func(v *EncryptionSettings) *KeyVaultAndSecretReference {
		if v == nil {
			return nil
		}
		return v.DiskEncryptionKey
	}).(KeyVaultAndSecretReferencePtrOutput)
}

func (o EncryptionSettingsPtrOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *EncryptionSettings) *bool {
		if v == nil {
			return nil
		}
		return v.Enabled
	}).(pulumi.BoolPtrOutput)
}

func (o EncryptionSettingsPtrOutput) KeyEncryptionKey() KeyVaultAndKeyReferencePtrOutput {
	return o.ApplyT(func(v *EncryptionSettings) *KeyVaultAndKeyReference {
		if v == nil {
			return nil
		}
		return v.KeyEncryptionKey
	}).(KeyVaultAndKeyReferencePtrOutput)
}

type EncryptionSettingsResponse struct {
	DiskEncryptionKey *KeyVaultAndSecretReferenceResponse `pulumi:"diskEncryptionKey"`
	Enabled           *bool                               `pulumi:"enabled"`
	KeyEncryptionKey  *KeyVaultAndKeyReferenceResponse    `pulumi:"keyEncryptionKey"`
}

type EncryptionSettingsResponseOutput struct{ *pulumi.OutputState }

func (EncryptionSettingsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EncryptionSettingsResponse)(nil)).Elem()
}

func (o EncryptionSettingsResponseOutput) ToEncryptionSettingsResponseOutput() EncryptionSettingsResponseOutput {
	return o
}

func (o EncryptionSettingsResponseOutput) ToEncryptionSettingsResponseOutputWithContext(ctx context.Context) EncryptionSettingsResponseOutput {
	return o
}

func (o EncryptionSettingsResponseOutput) DiskEncryptionKey() KeyVaultAndSecretReferenceResponsePtrOutput {
	return o.ApplyT(func(v EncryptionSettingsResponse) *KeyVaultAndSecretReferenceResponse { return v.DiskEncryptionKey }).(KeyVaultAndSecretReferenceResponsePtrOutput)
}

func (o EncryptionSettingsResponseOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v EncryptionSettingsResponse) *bool { return v.Enabled }).(pulumi.BoolPtrOutput)
}

func (o EncryptionSettingsResponseOutput) KeyEncryptionKey() KeyVaultAndKeyReferenceResponsePtrOutput {
	return o.ApplyT(func(v EncryptionSettingsResponse) *KeyVaultAndKeyReferenceResponse { return v.KeyEncryptionKey }).(KeyVaultAndKeyReferenceResponsePtrOutput)
}

type EncryptionSettingsResponsePtrOutput struct{ *pulumi.OutputState }

func (EncryptionSettingsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EncryptionSettingsResponse)(nil)).Elem()
}

func (o EncryptionSettingsResponsePtrOutput) ToEncryptionSettingsResponsePtrOutput() EncryptionSettingsResponsePtrOutput {
	return o
}

func (o EncryptionSettingsResponsePtrOutput) ToEncryptionSettingsResponsePtrOutputWithContext(ctx context.Context) EncryptionSettingsResponsePtrOutput {
	return o
}

func (o EncryptionSettingsResponsePtrOutput) Elem() EncryptionSettingsResponseOutput {
	return o.ApplyT(func(v *EncryptionSettingsResponse) EncryptionSettingsResponse {
		if v != nil {
			return *v
		}
		var ret EncryptionSettingsResponse
		return ret
	}).(EncryptionSettingsResponseOutput)
}

func (o EncryptionSettingsResponsePtrOutput) DiskEncryptionKey() KeyVaultAndSecretReferenceResponsePtrOutput {
	return o.ApplyT(func(v *EncryptionSettingsResponse) *KeyVaultAndSecretReferenceResponse {
		if v == nil {
			return nil
		}
		return v.DiskEncryptionKey
	}).(KeyVaultAndSecretReferenceResponsePtrOutput)
}

func (o EncryptionSettingsResponsePtrOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *EncryptionSettingsResponse) *bool {
		if v == nil {
			return nil
		}
		return v.Enabled
	}).(pulumi.BoolPtrOutput)
}

func (o EncryptionSettingsResponsePtrOutput) KeyEncryptionKey() KeyVaultAndKeyReferenceResponsePtrOutput {
	return o.ApplyT(func(v *EncryptionSettingsResponse) *KeyVaultAndKeyReferenceResponse {
		if v == nil {
			return nil
		}
		return v.KeyEncryptionKey
	}).(KeyVaultAndKeyReferenceResponsePtrOutput)
}

type HardwareProfile struct {
	VmSize *string `pulumi:"vmSize"`
}





type HardwareProfileInput interface {
	pulumi.Input

	ToHardwareProfileOutput() HardwareProfileOutput
	ToHardwareProfileOutputWithContext(context.Context) HardwareProfileOutput
}

type HardwareProfileArgs struct {
	VmSize pulumi.StringPtrInput `pulumi:"vmSize"`
}

func (HardwareProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*HardwareProfile)(nil)).Elem()
}

func (i HardwareProfileArgs) ToHardwareProfileOutput() HardwareProfileOutput {
	return i.ToHardwareProfileOutputWithContext(context.Background())
}

func (i HardwareProfileArgs) ToHardwareProfileOutputWithContext(ctx context.Context) HardwareProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HardwareProfileOutput)
}

func (i HardwareProfileArgs) ToHardwareProfilePtrOutput() HardwareProfilePtrOutput {
	return i.ToHardwareProfilePtrOutputWithContext(context.Background())
}

func (i HardwareProfileArgs) ToHardwareProfilePtrOutputWithContext(ctx context.Context) HardwareProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HardwareProfileOutput).ToHardwareProfilePtrOutputWithContext(ctx)
}









type HardwareProfilePtrInput interface {
	pulumi.Input

	ToHardwareProfilePtrOutput() HardwareProfilePtrOutput
	ToHardwareProfilePtrOutputWithContext(context.Context) HardwareProfilePtrOutput
}

type hardwareProfilePtrType HardwareProfileArgs

func HardwareProfilePtr(v *HardwareProfileArgs) HardwareProfilePtrInput {
	return (*hardwareProfilePtrType)(v)
}

func (*hardwareProfilePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**HardwareProfile)(nil)).Elem()
}

func (i *hardwareProfilePtrType) ToHardwareProfilePtrOutput() HardwareProfilePtrOutput {
	return i.ToHardwareProfilePtrOutputWithContext(context.Background())
}

func (i *hardwareProfilePtrType) ToHardwareProfilePtrOutputWithContext(ctx context.Context) HardwareProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HardwareProfilePtrOutput)
}

type HardwareProfileOutput struct{ *pulumi.OutputState }

func (HardwareProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*HardwareProfile)(nil)).Elem()
}

func (o HardwareProfileOutput) ToHardwareProfileOutput() HardwareProfileOutput {
	return o
}

func (o HardwareProfileOutput) ToHardwareProfileOutputWithContext(ctx context.Context) HardwareProfileOutput {
	return o
}

func (o HardwareProfileOutput) ToHardwareProfilePtrOutput() HardwareProfilePtrOutput {
	return o.ToHardwareProfilePtrOutputWithContext(context.Background())
}

func (o HardwareProfileOutput) ToHardwareProfilePtrOutputWithContext(ctx context.Context) HardwareProfilePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v HardwareProfile) *HardwareProfile {
		return &v
	}).(HardwareProfilePtrOutput)
}

func (o HardwareProfileOutput) VmSize() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HardwareProfile) *string { return v.VmSize }).(pulumi.StringPtrOutput)
}

type HardwareProfilePtrOutput struct{ *pulumi.OutputState }

func (HardwareProfilePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**HardwareProfile)(nil)).Elem()
}

func (o HardwareProfilePtrOutput) ToHardwareProfilePtrOutput() HardwareProfilePtrOutput {
	return o
}

func (o HardwareProfilePtrOutput) ToHardwareProfilePtrOutputWithContext(ctx context.Context) HardwareProfilePtrOutput {
	return o
}

func (o HardwareProfilePtrOutput) Elem() HardwareProfileOutput {
	return o.ApplyT(func(v *HardwareProfile) HardwareProfile {
		if v != nil {
			return *v
		}
		var ret HardwareProfile
		return ret
	}).(HardwareProfileOutput)
}

func (o HardwareProfilePtrOutput) VmSize() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HardwareProfile) *string {
		if v == nil {
			return nil
		}
		return v.VmSize
	}).(pulumi.StringPtrOutput)
}

type HardwareProfileResponse struct {
	VmSize *string `pulumi:"vmSize"`
}

type HardwareProfileResponseOutput struct{ *pulumi.OutputState }

func (HardwareProfileResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*HardwareProfileResponse)(nil)).Elem()
}

func (o HardwareProfileResponseOutput) ToHardwareProfileResponseOutput() HardwareProfileResponseOutput {
	return o
}

func (o HardwareProfileResponseOutput) ToHardwareProfileResponseOutputWithContext(ctx context.Context) HardwareProfileResponseOutput {
	return o
}

func (o HardwareProfileResponseOutput) VmSize() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HardwareProfileResponse) *string { return v.VmSize }).(pulumi.StringPtrOutput)
}

type HardwareProfileResponsePtrOutput struct{ *pulumi.OutputState }

func (HardwareProfileResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**HardwareProfileResponse)(nil)).Elem()
}

func (o HardwareProfileResponsePtrOutput) ToHardwareProfileResponsePtrOutput() HardwareProfileResponsePtrOutput {
	return o
}

func (o HardwareProfileResponsePtrOutput) ToHardwareProfileResponsePtrOutputWithContext(ctx context.Context) HardwareProfileResponsePtrOutput {
	return o
}

func (o HardwareProfileResponsePtrOutput) Elem() HardwareProfileResponseOutput {
	return o.ApplyT(func(v *HardwareProfileResponse) HardwareProfileResponse {
		if v != nil {
			return *v
		}
		var ret HardwareProfileResponse
		return ret
	}).(HardwareProfileResponseOutput)
}

func (o HardwareProfileResponsePtrOutput) VmSize() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HardwareProfileResponse) *string {
		if v == nil {
			return nil
		}
		return v.VmSize
	}).(pulumi.StringPtrOutput)
}

type ImageDataDisk struct {
	BlobUri            *string       `pulumi:"blobUri"`
	Caching            *CachingTypes `pulumi:"caching"`
	DiskSizeGB         *int          `pulumi:"diskSizeGB"`
	Lun                int           `pulumi:"lun"`
	ManagedDisk        *SubResource  `pulumi:"managedDisk"`
	Snapshot           *SubResource  `pulumi:"snapshot"`
	StorageAccountType *string       `pulumi:"storageAccountType"`
}





type ImageDataDiskInput interface {
	pulumi.Input

	ToImageDataDiskOutput() ImageDataDiskOutput
	ToImageDataDiskOutputWithContext(context.Context) ImageDataDiskOutput
}

type ImageDataDiskArgs struct {
	BlobUri            pulumi.StringPtrInput `pulumi:"blobUri"`
	Caching            CachingTypesPtrInput  `pulumi:"caching"`
	DiskSizeGB         pulumi.IntPtrInput    `pulumi:"diskSizeGB"`
	Lun                pulumi.IntInput       `pulumi:"lun"`
	ManagedDisk        SubResourcePtrInput   `pulumi:"managedDisk"`
	Snapshot           SubResourcePtrInput   `pulumi:"snapshot"`
	StorageAccountType pulumi.StringPtrInput `pulumi:"storageAccountType"`
}

func (ImageDataDiskArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ImageDataDisk)(nil)).Elem()
}

func (i ImageDataDiskArgs) ToImageDataDiskOutput() ImageDataDiskOutput {
	return i.ToImageDataDiskOutputWithContext(context.Background())
}

func (i ImageDataDiskArgs) ToImageDataDiskOutputWithContext(ctx context.Context) ImageDataDiskOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImageDataDiskOutput)
}





type ImageDataDiskArrayInput interface {
	pulumi.Input

	ToImageDataDiskArrayOutput() ImageDataDiskArrayOutput
	ToImageDataDiskArrayOutputWithContext(context.Context) ImageDataDiskArrayOutput
}

type ImageDataDiskArray []ImageDataDiskInput

func (ImageDataDiskArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ImageDataDisk)(nil)).Elem()
}

func (i ImageDataDiskArray) ToImageDataDiskArrayOutput() ImageDataDiskArrayOutput {
	return i.ToImageDataDiskArrayOutputWithContext(context.Background())
}

func (i ImageDataDiskArray) ToImageDataDiskArrayOutputWithContext(ctx context.Context) ImageDataDiskArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImageDataDiskArrayOutput)
}

type ImageDataDiskOutput struct{ *pulumi.OutputState }

func (ImageDataDiskOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ImageDataDisk)(nil)).Elem()
}

func (o ImageDataDiskOutput) ToImageDataDiskOutput() ImageDataDiskOutput {
	return o
}

func (o ImageDataDiskOutput) ToImageDataDiskOutputWithContext(ctx context.Context) ImageDataDiskOutput {
	return o
}

func (o ImageDataDiskOutput) BlobUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ImageDataDisk) *string { return v.BlobUri }).(pulumi.StringPtrOutput)
}

func (o ImageDataDiskOutput) Caching() CachingTypesPtrOutput {
	return o.ApplyT(func(v ImageDataDisk) *CachingTypes { return v.Caching }).(CachingTypesPtrOutput)
}

func (o ImageDataDiskOutput) DiskSizeGB() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ImageDataDisk) *int { return v.DiskSizeGB }).(pulumi.IntPtrOutput)
}

func (o ImageDataDiskOutput) Lun() pulumi.IntOutput {
	return o.ApplyT(func(v ImageDataDisk) int { return v.Lun }).(pulumi.IntOutput)
}

func (o ImageDataDiskOutput) ManagedDisk() SubResourcePtrOutput {
	return o.ApplyT(func(v ImageDataDisk) *SubResource { return v.ManagedDisk }).(SubResourcePtrOutput)
}

func (o ImageDataDiskOutput) Snapshot() SubResourcePtrOutput {
	return o.ApplyT(func(v ImageDataDisk) *SubResource { return v.Snapshot }).(SubResourcePtrOutput)
}

func (o ImageDataDiskOutput) StorageAccountType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ImageDataDisk) *string { return v.StorageAccountType }).(pulumi.StringPtrOutput)
}

type ImageDataDiskArrayOutput struct{ *pulumi.OutputState }

func (ImageDataDiskArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ImageDataDisk)(nil)).Elem()
}

func (o ImageDataDiskArrayOutput) ToImageDataDiskArrayOutput() ImageDataDiskArrayOutput {
	return o
}

func (o ImageDataDiskArrayOutput) ToImageDataDiskArrayOutputWithContext(ctx context.Context) ImageDataDiskArrayOutput {
	return o
}

func (o ImageDataDiskArrayOutput) Index(i pulumi.IntInput) ImageDataDiskOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ImageDataDisk {
		return vs[0].([]ImageDataDisk)[vs[1].(int)]
	}).(ImageDataDiskOutput)
}

type ImageDataDiskResponse struct {
	BlobUri            *string              `pulumi:"blobUri"`
	Caching            *string              `pulumi:"caching"`
	DiskSizeGB         *int                 `pulumi:"diskSizeGB"`
	Lun                int                  `pulumi:"lun"`
	ManagedDisk        *SubResourceResponse `pulumi:"managedDisk"`
	Snapshot           *SubResourceResponse `pulumi:"snapshot"`
	StorageAccountType *string              `pulumi:"storageAccountType"`
}

type ImageDataDiskResponseOutput struct{ *pulumi.OutputState }

func (ImageDataDiskResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ImageDataDiskResponse)(nil)).Elem()
}

func (o ImageDataDiskResponseOutput) ToImageDataDiskResponseOutput() ImageDataDiskResponseOutput {
	return o
}

func (o ImageDataDiskResponseOutput) ToImageDataDiskResponseOutputWithContext(ctx context.Context) ImageDataDiskResponseOutput {
	return o
}

func (o ImageDataDiskResponseOutput) BlobUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ImageDataDiskResponse) *string { return v.BlobUri }).(pulumi.StringPtrOutput)
}

func (o ImageDataDiskResponseOutput) Caching() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ImageDataDiskResponse) *string { return v.Caching }).(pulumi.StringPtrOutput)
}

func (o ImageDataDiskResponseOutput) DiskSizeGB() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ImageDataDiskResponse) *int { return v.DiskSizeGB }).(pulumi.IntPtrOutput)
}

func (o ImageDataDiskResponseOutput) Lun() pulumi.IntOutput {
	return o.ApplyT(func(v ImageDataDiskResponse) int { return v.Lun }).(pulumi.IntOutput)
}

func (o ImageDataDiskResponseOutput) ManagedDisk() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v ImageDataDiskResponse) *SubResourceResponse { return v.ManagedDisk }).(SubResourceResponsePtrOutput)
}

func (o ImageDataDiskResponseOutput) Snapshot() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v ImageDataDiskResponse) *SubResourceResponse { return v.Snapshot }).(SubResourceResponsePtrOutput)
}

func (o ImageDataDiskResponseOutput) StorageAccountType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ImageDataDiskResponse) *string { return v.StorageAccountType }).(pulumi.StringPtrOutput)
}

type ImageDataDiskResponseArrayOutput struct{ *pulumi.OutputState }

func (ImageDataDiskResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ImageDataDiskResponse)(nil)).Elem()
}

func (o ImageDataDiskResponseArrayOutput) ToImageDataDiskResponseArrayOutput() ImageDataDiskResponseArrayOutput {
	return o
}

func (o ImageDataDiskResponseArrayOutput) ToImageDataDiskResponseArrayOutputWithContext(ctx context.Context) ImageDataDiskResponseArrayOutput {
	return o
}

func (o ImageDataDiskResponseArrayOutput) Index(i pulumi.IntInput) ImageDataDiskResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ImageDataDiskResponse {
		return vs[0].([]ImageDataDiskResponse)[vs[1].(int)]
	}).(ImageDataDiskResponseOutput)
}

type ImageDiskReference struct {
	Id  string `pulumi:"id"`
	Lun *int   `pulumi:"lun"`
}





type ImageDiskReferenceInput interface {
	pulumi.Input

	ToImageDiskReferenceOutput() ImageDiskReferenceOutput
	ToImageDiskReferenceOutputWithContext(context.Context) ImageDiskReferenceOutput
}

type ImageDiskReferenceArgs struct {
	Id  pulumi.StringInput `pulumi:"id"`
	Lun pulumi.IntPtrInput `pulumi:"lun"`
}

func (ImageDiskReferenceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ImageDiskReference)(nil)).Elem()
}

func (i ImageDiskReferenceArgs) ToImageDiskReferenceOutput() ImageDiskReferenceOutput {
	return i.ToImageDiskReferenceOutputWithContext(context.Background())
}

func (i ImageDiskReferenceArgs) ToImageDiskReferenceOutputWithContext(ctx context.Context) ImageDiskReferenceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImageDiskReferenceOutput)
}

func (i ImageDiskReferenceArgs) ToImageDiskReferencePtrOutput() ImageDiskReferencePtrOutput {
	return i.ToImageDiskReferencePtrOutputWithContext(context.Background())
}

func (i ImageDiskReferenceArgs) ToImageDiskReferencePtrOutputWithContext(ctx context.Context) ImageDiskReferencePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImageDiskReferenceOutput).ToImageDiskReferencePtrOutputWithContext(ctx)
}









type ImageDiskReferencePtrInput interface {
	pulumi.Input

	ToImageDiskReferencePtrOutput() ImageDiskReferencePtrOutput
	ToImageDiskReferencePtrOutputWithContext(context.Context) ImageDiskReferencePtrOutput
}

type imageDiskReferencePtrType ImageDiskReferenceArgs

func ImageDiskReferencePtr(v *ImageDiskReferenceArgs) ImageDiskReferencePtrInput {
	return (*imageDiskReferencePtrType)(v)
}

func (*imageDiskReferencePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ImageDiskReference)(nil)).Elem()
}

func (i *imageDiskReferencePtrType) ToImageDiskReferencePtrOutput() ImageDiskReferencePtrOutput {
	return i.ToImageDiskReferencePtrOutputWithContext(context.Background())
}

func (i *imageDiskReferencePtrType) ToImageDiskReferencePtrOutputWithContext(ctx context.Context) ImageDiskReferencePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImageDiskReferencePtrOutput)
}

type ImageDiskReferenceOutput struct{ *pulumi.OutputState }

func (ImageDiskReferenceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ImageDiskReference)(nil)).Elem()
}

func (o ImageDiskReferenceOutput) ToImageDiskReferenceOutput() ImageDiskReferenceOutput {
	return o
}

func (o ImageDiskReferenceOutput) ToImageDiskReferenceOutputWithContext(ctx context.Context) ImageDiskReferenceOutput {
	return o
}

func (o ImageDiskReferenceOutput) ToImageDiskReferencePtrOutput() ImageDiskReferencePtrOutput {
	return o.ToImageDiskReferencePtrOutputWithContext(context.Background())
}

func (o ImageDiskReferenceOutput) ToImageDiskReferencePtrOutputWithContext(ctx context.Context) ImageDiskReferencePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ImageDiskReference) *ImageDiskReference {
		return &v
	}).(ImageDiskReferencePtrOutput)
}

func (o ImageDiskReferenceOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v ImageDiskReference) string { return v.Id }).(pulumi.StringOutput)
}

func (o ImageDiskReferenceOutput) Lun() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ImageDiskReference) *int { return v.Lun }).(pulumi.IntPtrOutput)
}

type ImageDiskReferencePtrOutput struct{ *pulumi.OutputState }

func (ImageDiskReferencePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ImageDiskReference)(nil)).Elem()
}

func (o ImageDiskReferencePtrOutput) ToImageDiskReferencePtrOutput() ImageDiskReferencePtrOutput {
	return o
}

func (o ImageDiskReferencePtrOutput) ToImageDiskReferencePtrOutputWithContext(ctx context.Context) ImageDiskReferencePtrOutput {
	return o
}

func (o ImageDiskReferencePtrOutput) Elem() ImageDiskReferenceOutput {
	return o.ApplyT(func(v *ImageDiskReference) ImageDiskReference {
		if v != nil {
			return *v
		}
		var ret ImageDiskReference
		return ret
	}).(ImageDiskReferenceOutput)
}

func (o ImageDiskReferencePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ImageDiskReference) *string {
		if v == nil {
			return nil
		}
		return &v.Id
	}).(pulumi.StringPtrOutput)
}

func (o ImageDiskReferencePtrOutput) Lun() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ImageDiskReference) *int {
		if v == nil {
			return nil
		}
		return v.Lun
	}).(pulumi.IntPtrOutput)
}

type ImageDiskReferenceResponse struct {
	Id  string `pulumi:"id"`
	Lun *int   `pulumi:"lun"`
}

type ImageDiskReferenceResponseOutput struct{ *pulumi.OutputState }

func (ImageDiskReferenceResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ImageDiskReferenceResponse)(nil)).Elem()
}

func (o ImageDiskReferenceResponseOutput) ToImageDiskReferenceResponseOutput() ImageDiskReferenceResponseOutput {
	return o
}

func (o ImageDiskReferenceResponseOutput) ToImageDiskReferenceResponseOutputWithContext(ctx context.Context) ImageDiskReferenceResponseOutput {
	return o
}

func (o ImageDiskReferenceResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v ImageDiskReferenceResponse) string { return v.Id }).(pulumi.StringOutput)
}

func (o ImageDiskReferenceResponseOutput) Lun() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ImageDiskReferenceResponse) *int { return v.Lun }).(pulumi.IntPtrOutput)
}

type ImageDiskReferenceResponsePtrOutput struct{ *pulumi.OutputState }

func (ImageDiskReferenceResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ImageDiskReferenceResponse)(nil)).Elem()
}

func (o ImageDiskReferenceResponsePtrOutput) ToImageDiskReferenceResponsePtrOutput() ImageDiskReferenceResponsePtrOutput {
	return o
}

func (o ImageDiskReferenceResponsePtrOutput) ToImageDiskReferenceResponsePtrOutputWithContext(ctx context.Context) ImageDiskReferenceResponsePtrOutput {
	return o
}

func (o ImageDiskReferenceResponsePtrOutput) Elem() ImageDiskReferenceResponseOutput {
	return o.ApplyT(func(v *ImageDiskReferenceResponse) ImageDiskReferenceResponse {
		if v != nil {
			return *v
		}
		var ret ImageDiskReferenceResponse
		return ret
	}).(ImageDiskReferenceResponseOutput)
}

func (o ImageDiskReferenceResponsePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ImageDiskReferenceResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Id
	}).(pulumi.StringPtrOutput)
}

func (o ImageDiskReferenceResponsePtrOutput) Lun() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ImageDiskReferenceResponse) *int {
		if v == nil {
			return nil
		}
		return v.Lun
	}).(pulumi.IntPtrOutput)
}

type ImageOSDisk struct {
	BlobUri            *string                   `pulumi:"blobUri"`
	Caching            *CachingTypes             `pulumi:"caching"`
	DiskSizeGB         *int                      `pulumi:"diskSizeGB"`
	ManagedDisk        *SubResource              `pulumi:"managedDisk"`
	OsState            OperatingSystemStateTypes `pulumi:"osState"`
	OsType             OperatingSystemTypes      `pulumi:"osType"`
	Snapshot           *SubResource              `pulumi:"snapshot"`
	StorageAccountType *string                   `pulumi:"storageAccountType"`
}





type ImageOSDiskInput interface {
	pulumi.Input

	ToImageOSDiskOutput() ImageOSDiskOutput
	ToImageOSDiskOutputWithContext(context.Context) ImageOSDiskOutput
}

type ImageOSDiskArgs struct {
	BlobUri            pulumi.StringPtrInput          `pulumi:"blobUri"`
	Caching            CachingTypesPtrInput           `pulumi:"caching"`
	DiskSizeGB         pulumi.IntPtrInput             `pulumi:"diskSizeGB"`
	ManagedDisk        SubResourcePtrInput            `pulumi:"managedDisk"`
	OsState            OperatingSystemStateTypesInput `pulumi:"osState"`
	OsType             OperatingSystemTypesInput      `pulumi:"osType"`
	Snapshot           SubResourcePtrInput            `pulumi:"snapshot"`
	StorageAccountType pulumi.StringPtrInput          `pulumi:"storageAccountType"`
}

func (ImageOSDiskArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ImageOSDisk)(nil)).Elem()
}

func (i ImageOSDiskArgs) ToImageOSDiskOutput() ImageOSDiskOutput {
	return i.ToImageOSDiskOutputWithContext(context.Background())
}

func (i ImageOSDiskArgs) ToImageOSDiskOutputWithContext(ctx context.Context) ImageOSDiskOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImageOSDiskOutput)
}

func (i ImageOSDiskArgs) ToImageOSDiskPtrOutput() ImageOSDiskPtrOutput {
	return i.ToImageOSDiskPtrOutputWithContext(context.Background())
}

func (i ImageOSDiskArgs) ToImageOSDiskPtrOutputWithContext(ctx context.Context) ImageOSDiskPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImageOSDiskOutput).ToImageOSDiskPtrOutputWithContext(ctx)
}









type ImageOSDiskPtrInput interface {
	pulumi.Input

	ToImageOSDiskPtrOutput() ImageOSDiskPtrOutput
	ToImageOSDiskPtrOutputWithContext(context.Context) ImageOSDiskPtrOutput
}

type imageOSDiskPtrType ImageOSDiskArgs

func ImageOSDiskPtr(v *ImageOSDiskArgs) ImageOSDiskPtrInput {
	return (*imageOSDiskPtrType)(v)
}

func (*imageOSDiskPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ImageOSDisk)(nil)).Elem()
}

func (i *imageOSDiskPtrType) ToImageOSDiskPtrOutput() ImageOSDiskPtrOutput {
	return i.ToImageOSDiskPtrOutputWithContext(context.Background())
}

func (i *imageOSDiskPtrType) ToImageOSDiskPtrOutputWithContext(ctx context.Context) ImageOSDiskPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImageOSDiskPtrOutput)
}

type ImageOSDiskOutput struct{ *pulumi.OutputState }

func (ImageOSDiskOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ImageOSDisk)(nil)).Elem()
}

func (o ImageOSDiskOutput) ToImageOSDiskOutput() ImageOSDiskOutput {
	return o
}

func (o ImageOSDiskOutput) ToImageOSDiskOutputWithContext(ctx context.Context) ImageOSDiskOutput {
	return o
}

func (o ImageOSDiskOutput) ToImageOSDiskPtrOutput() ImageOSDiskPtrOutput {
	return o.ToImageOSDiskPtrOutputWithContext(context.Background())
}

func (o ImageOSDiskOutput) ToImageOSDiskPtrOutputWithContext(ctx context.Context) ImageOSDiskPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ImageOSDisk) *ImageOSDisk {
		return &v
	}).(ImageOSDiskPtrOutput)
}

func (o ImageOSDiskOutput) BlobUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ImageOSDisk) *string { return v.BlobUri }).(pulumi.StringPtrOutput)
}

func (o ImageOSDiskOutput) Caching() CachingTypesPtrOutput {
	return o.ApplyT(func(v ImageOSDisk) *CachingTypes { return v.Caching }).(CachingTypesPtrOutput)
}

func (o ImageOSDiskOutput) DiskSizeGB() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ImageOSDisk) *int { return v.DiskSizeGB }).(pulumi.IntPtrOutput)
}

func (o ImageOSDiskOutput) ManagedDisk() SubResourcePtrOutput {
	return o.ApplyT(func(v ImageOSDisk) *SubResource { return v.ManagedDisk }).(SubResourcePtrOutput)
}

func (o ImageOSDiskOutput) OsState() OperatingSystemStateTypesOutput {
	return o.ApplyT(func(v ImageOSDisk) OperatingSystemStateTypes { return v.OsState }).(OperatingSystemStateTypesOutput)
}

func (o ImageOSDiskOutput) OsType() OperatingSystemTypesOutput {
	return o.ApplyT(func(v ImageOSDisk) OperatingSystemTypes { return v.OsType }).(OperatingSystemTypesOutput)
}

func (o ImageOSDiskOutput) Snapshot() SubResourcePtrOutput {
	return o.ApplyT(func(v ImageOSDisk) *SubResource { return v.Snapshot }).(SubResourcePtrOutput)
}

func (o ImageOSDiskOutput) StorageAccountType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ImageOSDisk) *string { return v.StorageAccountType }).(pulumi.StringPtrOutput)
}

type ImageOSDiskPtrOutput struct{ *pulumi.OutputState }

func (ImageOSDiskPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ImageOSDisk)(nil)).Elem()
}

func (o ImageOSDiskPtrOutput) ToImageOSDiskPtrOutput() ImageOSDiskPtrOutput {
	return o
}

func (o ImageOSDiskPtrOutput) ToImageOSDiskPtrOutputWithContext(ctx context.Context) ImageOSDiskPtrOutput {
	return o
}

func (o ImageOSDiskPtrOutput) Elem() ImageOSDiskOutput {
	return o.ApplyT(func(v *ImageOSDisk) ImageOSDisk {
		if v != nil {
			return *v
		}
		var ret ImageOSDisk
		return ret
	}).(ImageOSDiskOutput)
}

func (o ImageOSDiskPtrOutput) BlobUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ImageOSDisk) *string {
		if v == nil {
			return nil
		}
		return v.BlobUri
	}).(pulumi.StringPtrOutput)
}

func (o ImageOSDiskPtrOutput) Caching() CachingTypesPtrOutput {
	return o.ApplyT(func(v *ImageOSDisk) *CachingTypes {
		if v == nil {
			return nil
		}
		return v.Caching
	}).(CachingTypesPtrOutput)
}

func (o ImageOSDiskPtrOutput) DiskSizeGB() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ImageOSDisk) *int {
		if v == nil {
			return nil
		}
		return v.DiskSizeGB
	}).(pulumi.IntPtrOutput)
}

func (o ImageOSDiskPtrOutput) ManagedDisk() SubResourcePtrOutput {
	return o.ApplyT(func(v *ImageOSDisk) *SubResource {
		if v == nil {
			return nil
		}
		return v.ManagedDisk
	}).(SubResourcePtrOutput)
}

func (o ImageOSDiskPtrOutput) OsState() OperatingSystemStateTypesPtrOutput {
	return o.ApplyT(func(v *ImageOSDisk) *OperatingSystemStateTypes {
		if v == nil {
			return nil
		}
		return &v.OsState
	}).(OperatingSystemStateTypesPtrOutput)
}

func (o ImageOSDiskPtrOutput) OsType() OperatingSystemTypesPtrOutput {
	return o.ApplyT(func(v *ImageOSDisk) *OperatingSystemTypes {
		if v == nil {
			return nil
		}
		return &v.OsType
	}).(OperatingSystemTypesPtrOutput)
}

func (o ImageOSDiskPtrOutput) Snapshot() SubResourcePtrOutput {
	return o.ApplyT(func(v *ImageOSDisk) *SubResource {
		if v == nil {
			return nil
		}
		return v.Snapshot
	}).(SubResourcePtrOutput)
}

func (o ImageOSDiskPtrOutput) StorageAccountType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ImageOSDisk) *string {
		if v == nil {
			return nil
		}
		return v.StorageAccountType
	}).(pulumi.StringPtrOutput)
}

type ImageOSDiskResponse struct {
	BlobUri            *string              `pulumi:"blobUri"`
	Caching            *string              `pulumi:"caching"`
	DiskSizeGB         *int                 `pulumi:"diskSizeGB"`
	ManagedDisk        *SubResourceResponse `pulumi:"managedDisk"`
	OsState            string               `pulumi:"osState"`
	OsType             string               `pulumi:"osType"`
	Snapshot           *SubResourceResponse `pulumi:"snapshot"`
	StorageAccountType *string              `pulumi:"storageAccountType"`
}

type ImageOSDiskResponseOutput struct{ *pulumi.OutputState }

func (ImageOSDiskResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ImageOSDiskResponse)(nil)).Elem()
}

func (o ImageOSDiskResponseOutput) ToImageOSDiskResponseOutput() ImageOSDiskResponseOutput {
	return o
}

func (o ImageOSDiskResponseOutput) ToImageOSDiskResponseOutputWithContext(ctx context.Context) ImageOSDiskResponseOutput {
	return o
}

func (o ImageOSDiskResponseOutput) BlobUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ImageOSDiskResponse) *string { return v.BlobUri }).(pulumi.StringPtrOutput)
}

func (o ImageOSDiskResponseOutput) Caching() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ImageOSDiskResponse) *string { return v.Caching }).(pulumi.StringPtrOutput)
}

func (o ImageOSDiskResponseOutput) DiskSizeGB() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ImageOSDiskResponse) *int { return v.DiskSizeGB }).(pulumi.IntPtrOutput)
}

func (o ImageOSDiskResponseOutput) ManagedDisk() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v ImageOSDiskResponse) *SubResourceResponse { return v.ManagedDisk }).(SubResourceResponsePtrOutput)
}

func (o ImageOSDiskResponseOutput) OsState() pulumi.StringOutput {
	return o.ApplyT(func(v ImageOSDiskResponse) string { return v.OsState }).(pulumi.StringOutput)
}

func (o ImageOSDiskResponseOutput) OsType() pulumi.StringOutput {
	return o.ApplyT(func(v ImageOSDiskResponse) string { return v.OsType }).(pulumi.StringOutput)
}

func (o ImageOSDiskResponseOutput) Snapshot() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v ImageOSDiskResponse) *SubResourceResponse { return v.Snapshot }).(SubResourceResponsePtrOutput)
}

func (o ImageOSDiskResponseOutput) StorageAccountType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ImageOSDiskResponse) *string { return v.StorageAccountType }).(pulumi.StringPtrOutput)
}

type ImageOSDiskResponsePtrOutput struct{ *pulumi.OutputState }

func (ImageOSDiskResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ImageOSDiskResponse)(nil)).Elem()
}

func (o ImageOSDiskResponsePtrOutput) ToImageOSDiskResponsePtrOutput() ImageOSDiskResponsePtrOutput {
	return o
}

func (o ImageOSDiskResponsePtrOutput) ToImageOSDiskResponsePtrOutputWithContext(ctx context.Context) ImageOSDiskResponsePtrOutput {
	return o
}

func (o ImageOSDiskResponsePtrOutput) Elem() ImageOSDiskResponseOutput {
	return o.ApplyT(func(v *ImageOSDiskResponse) ImageOSDiskResponse {
		if v != nil {
			return *v
		}
		var ret ImageOSDiskResponse
		return ret
	}).(ImageOSDiskResponseOutput)
}

func (o ImageOSDiskResponsePtrOutput) BlobUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ImageOSDiskResponse) *string {
		if v == nil {
			return nil
		}
		return v.BlobUri
	}).(pulumi.StringPtrOutput)
}

func (o ImageOSDiskResponsePtrOutput) Caching() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ImageOSDiskResponse) *string {
		if v == nil {
			return nil
		}
		return v.Caching
	}).(pulumi.StringPtrOutput)
}

func (o ImageOSDiskResponsePtrOutput) DiskSizeGB() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ImageOSDiskResponse) *int {
		if v == nil {
			return nil
		}
		return v.DiskSizeGB
	}).(pulumi.IntPtrOutput)
}

func (o ImageOSDiskResponsePtrOutput) ManagedDisk() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v *ImageOSDiskResponse) *SubResourceResponse {
		if v == nil {
			return nil
		}
		return v.ManagedDisk
	}).(SubResourceResponsePtrOutput)
}

func (o ImageOSDiskResponsePtrOutput) OsState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ImageOSDiskResponse) *string {
		if v == nil {
			return nil
		}
		return &v.OsState
	}).(pulumi.StringPtrOutput)
}

func (o ImageOSDiskResponsePtrOutput) OsType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ImageOSDiskResponse) *string {
		if v == nil {
			return nil
		}
		return &v.OsType
	}).(pulumi.StringPtrOutput)
}

func (o ImageOSDiskResponsePtrOutput) Snapshot() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v *ImageOSDiskResponse) *SubResourceResponse {
		if v == nil {
			return nil
		}
		return v.Snapshot
	}).(SubResourceResponsePtrOutput)
}

func (o ImageOSDiskResponsePtrOutput) StorageAccountType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ImageOSDiskResponse) *string {
		if v == nil {
			return nil
		}
		return v.StorageAccountType
	}).(pulumi.StringPtrOutput)
}

type ImageReference struct {
	Id        *string `pulumi:"id"`
	Offer     *string `pulumi:"offer"`
	Publisher *string `pulumi:"publisher"`
	Sku       *string `pulumi:"sku"`
	Version   *string `pulumi:"version"`
}





type ImageReferenceInput interface {
	pulumi.Input

	ToImageReferenceOutput() ImageReferenceOutput
	ToImageReferenceOutputWithContext(context.Context) ImageReferenceOutput
}

type ImageReferenceArgs struct {
	Id        pulumi.StringPtrInput `pulumi:"id"`
	Offer     pulumi.StringPtrInput `pulumi:"offer"`
	Publisher pulumi.StringPtrInput `pulumi:"publisher"`
	Sku       pulumi.StringPtrInput `pulumi:"sku"`
	Version   pulumi.StringPtrInput `pulumi:"version"`
}

func (ImageReferenceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ImageReference)(nil)).Elem()
}

func (i ImageReferenceArgs) ToImageReferenceOutput() ImageReferenceOutput {
	return i.ToImageReferenceOutputWithContext(context.Background())
}

func (i ImageReferenceArgs) ToImageReferenceOutputWithContext(ctx context.Context) ImageReferenceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImageReferenceOutput)
}

func (i ImageReferenceArgs) ToImageReferencePtrOutput() ImageReferencePtrOutput {
	return i.ToImageReferencePtrOutputWithContext(context.Background())
}

func (i ImageReferenceArgs) ToImageReferencePtrOutputWithContext(ctx context.Context) ImageReferencePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImageReferenceOutput).ToImageReferencePtrOutputWithContext(ctx)
}









type ImageReferencePtrInput interface {
	pulumi.Input

	ToImageReferencePtrOutput() ImageReferencePtrOutput
	ToImageReferencePtrOutputWithContext(context.Context) ImageReferencePtrOutput
}

type imageReferencePtrType ImageReferenceArgs

func ImageReferencePtr(v *ImageReferenceArgs) ImageReferencePtrInput {
	return (*imageReferencePtrType)(v)
}

func (*imageReferencePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ImageReference)(nil)).Elem()
}

func (i *imageReferencePtrType) ToImageReferencePtrOutput() ImageReferencePtrOutput {
	return i.ToImageReferencePtrOutputWithContext(context.Background())
}

func (i *imageReferencePtrType) ToImageReferencePtrOutputWithContext(ctx context.Context) ImageReferencePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImageReferencePtrOutput)
}

type ImageReferenceOutput struct{ *pulumi.OutputState }

func (ImageReferenceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ImageReference)(nil)).Elem()
}

func (o ImageReferenceOutput) ToImageReferenceOutput() ImageReferenceOutput {
	return o
}

func (o ImageReferenceOutput) ToImageReferenceOutputWithContext(ctx context.Context) ImageReferenceOutput {
	return o
}

func (o ImageReferenceOutput) ToImageReferencePtrOutput() ImageReferencePtrOutput {
	return o.ToImageReferencePtrOutputWithContext(context.Background())
}

func (o ImageReferenceOutput) ToImageReferencePtrOutputWithContext(ctx context.Context) ImageReferencePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ImageReference) *ImageReference {
		return &v
	}).(ImageReferencePtrOutput)
}

func (o ImageReferenceOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ImageReference) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o ImageReferenceOutput) Offer() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ImageReference) *string { return v.Offer }).(pulumi.StringPtrOutput)
}

func (o ImageReferenceOutput) Publisher() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ImageReference) *string { return v.Publisher }).(pulumi.StringPtrOutput)
}

func (o ImageReferenceOutput) Sku() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ImageReference) *string { return v.Sku }).(pulumi.StringPtrOutput)
}

func (o ImageReferenceOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ImageReference) *string { return v.Version }).(pulumi.StringPtrOutput)
}

type ImageReferencePtrOutput struct{ *pulumi.OutputState }

func (ImageReferencePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ImageReference)(nil)).Elem()
}

func (o ImageReferencePtrOutput) ToImageReferencePtrOutput() ImageReferencePtrOutput {
	return o
}

func (o ImageReferencePtrOutput) ToImageReferencePtrOutputWithContext(ctx context.Context) ImageReferencePtrOutput {
	return o
}

func (o ImageReferencePtrOutput) Elem() ImageReferenceOutput {
	return o.ApplyT(func(v *ImageReference) ImageReference {
		if v != nil {
			return *v
		}
		var ret ImageReference
		return ret
	}).(ImageReferenceOutput)
}

func (o ImageReferencePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ImageReference) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

func (o ImageReferencePtrOutput) Offer() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ImageReference) *string {
		if v == nil {
			return nil
		}
		return v.Offer
	}).(pulumi.StringPtrOutput)
}

func (o ImageReferencePtrOutput) Publisher() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ImageReference) *string {
		if v == nil {
			return nil
		}
		return v.Publisher
	}).(pulumi.StringPtrOutput)
}

func (o ImageReferencePtrOutput) Sku() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ImageReference) *string {
		if v == nil {
			return nil
		}
		return v.Sku
	}).(pulumi.StringPtrOutput)
}

func (o ImageReferencePtrOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ImageReference) *string {
		if v == nil {
			return nil
		}
		return v.Version
	}).(pulumi.StringPtrOutput)
}

type ImageReferenceResponse struct {
	Id        *string `pulumi:"id"`
	Offer     *string `pulumi:"offer"`
	Publisher *string `pulumi:"publisher"`
	Sku       *string `pulumi:"sku"`
	Version   *string `pulumi:"version"`
}

type ImageReferenceResponseOutput struct{ *pulumi.OutputState }

func (ImageReferenceResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ImageReferenceResponse)(nil)).Elem()
}

func (o ImageReferenceResponseOutput) ToImageReferenceResponseOutput() ImageReferenceResponseOutput {
	return o
}

func (o ImageReferenceResponseOutput) ToImageReferenceResponseOutputWithContext(ctx context.Context) ImageReferenceResponseOutput {
	return o
}

func (o ImageReferenceResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ImageReferenceResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o ImageReferenceResponseOutput) Offer() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ImageReferenceResponse) *string { return v.Offer }).(pulumi.StringPtrOutput)
}

func (o ImageReferenceResponseOutput) Publisher() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ImageReferenceResponse) *string { return v.Publisher }).(pulumi.StringPtrOutput)
}

func (o ImageReferenceResponseOutput) Sku() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ImageReferenceResponse) *string { return v.Sku }).(pulumi.StringPtrOutput)
}

func (o ImageReferenceResponseOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ImageReferenceResponse) *string { return v.Version }).(pulumi.StringPtrOutput)
}

type ImageReferenceResponsePtrOutput struct{ *pulumi.OutputState }

func (ImageReferenceResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ImageReferenceResponse)(nil)).Elem()
}

func (o ImageReferenceResponsePtrOutput) ToImageReferenceResponsePtrOutput() ImageReferenceResponsePtrOutput {
	return o
}

func (o ImageReferenceResponsePtrOutput) ToImageReferenceResponsePtrOutputWithContext(ctx context.Context) ImageReferenceResponsePtrOutput {
	return o
}

func (o ImageReferenceResponsePtrOutput) Elem() ImageReferenceResponseOutput {
	return o.ApplyT(func(v *ImageReferenceResponse) ImageReferenceResponse {
		if v != nil {
			return *v
		}
		var ret ImageReferenceResponse
		return ret
	}).(ImageReferenceResponseOutput)
}

func (o ImageReferenceResponsePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ImageReferenceResponse) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

func (o ImageReferenceResponsePtrOutput) Offer() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ImageReferenceResponse) *string {
		if v == nil {
			return nil
		}
		return v.Offer
	}).(pulumi.StringPtrOutput)
}

func (o ImageReferenceResponsePtrOutput) Publisher() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ImageReferenceResponse) *string {
		if v == nil {
			return nil
		}
		return v.Publisher
	}).(pulumi.StringPtrOutput)
}

func (o ImageReferenceResponsePtrOutput) Sku() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ImageReferenceResponse) *string {
		if v == nil {
			return nil
		}
		return v.Sku
	}).(pulumi.StringPtrOutput)
}

func (o ImageReferenceResponsePtrOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ImageReferenceResponse) *string {
		if v == nil {
			return nil
		}
		return v.Version
	}).(pulumi.StringPtrOutput)
}

type ImageStorageProfile struct {
	DataDisks     []ImageDataDisk `pulumi:"dataDisks"`
	OsDisk        *ImageOSDisk    `pulumi:"osDisk"`
	ZoneResilient *bool           `pulumi:"zoneResilient"`
}





type ImageStorageProfileInput interface {
	pulumi.Input

	ToImageStorageProfileOutput() ImageStorageProfileOutput
	ToImageStorageProfileOutputWithContext(context.Context) ImageStorageProfileOutput
}

type ImageStorageProfileArgs struct {
	DataDisks     ImageDataDiskArrayInput `pulumi:"dataDisks"`
	OsDisk        ImageOSDiskPtrInput     `pulumi:"osDisk"`
	ZoneResilient pulumi.BoolPtrInput     `pulumi:"zoneResilient"`
}

func (ImageStorageProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ImageStorageProfile)(nil)).Elem()
}

func (i ImageStorageProfileArgs) ToImageStorageProfileOutput() ImageStorageProfileOutput {
	return i.ToImageStorageProfileOutputWithContext(context.Background())
}

func (i ImageStorageProfileArgs) ToImageStorageProfileOutputWithContext(ctx context.Context) ImageStorageProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImageStorageProfileOutput)
}

func (i ImageStorageProfileArgs) ToImageStorageProfilePtrOutput() ImageStorageProfilePtrOutput {
	return i.ToImageStorageProfilePtrOutputWithContext(context.Background())
}

func (i ImageStorageProfileArgs) ToImageStorageProfilePtrOutputWithContext(ctx context.Context) ImageStorageProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImageStorageProfileOutput).ToImageStorageProfilePtrOutputWithContext(ctx)
}









type ImageStorageProfilePtrInput interface {
	pulumi.Input

	ToImageStorageProfilePtrOutput() ImageStorageProfilePtrOutput
	ToImageStorageProfilePtrOutputWithContext(context.Context) ImageStorageProfilePtrOutput
}

type imageStorageProfilePtrType ImageStorageProfileArgs

func ImageStorageProfilePtr(v *ImageStorageProfileArgs) ImageStorageProfilePtrInput {
	return (*imageStorageProfilePtrType)(v)
}

func (*imageStorageProfilePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ImageStorageProfile)(nil)).Elem()
}

func (i *imageStorageProfilePtrType) ToImageStorageProfilePtrOutput() ImageStorageProfilePtrOutput {
	return i.ToImageStorageProfilePtrOutputWithContext(context.Background())
}

func (i *imageStorageProfilePtrType) ToImageStorageProfilePtrOutputWithContext(ctx context.Context) ImageStorageProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImageStorageProfilePtrOutput)
}

type ImageStorageProfileOutput struct{ *pulumi.OutputState }

func (ImageStorageProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ImageStorageProfile)(nil)).Elem()
}

func (o ImageStorageProfileOutput) ToImageStorageProfileOutput() ImageStorageProfileOutput {
	return o
}

func (o ImageStorageProfileOutput) ToImageStorageProfileOutputWithContext(ctx context.Context) ImageStorageProfileOutput {
	return o
}

func (o ImageStorageProfileOutput) ToImageStorageProfilePtrOutput() ImageStorageProfilePtrOutput {
	return o.ToImageStorageProfilePtrOutputWithContext(context.Background())
}

func (o ImageStorageProfileOutput) ToImageStorageProfilePtrOutputWithContext(ctx context.Context) ImageStorageProfilePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ImageStorageProfile) *ImageStorageProfile {
		return &v
	}).(ImageStorageProfilePtrOutput)
}

func (o ImageStorageProfileOutput) DataDisks() ImageDataDiskArrayOutput {
	return o.ApplyT(func(v ImageStorageProfile) []ImageDataDisk { return v.DataDisks }).(ImageDataDiskArrayOutput)
}

func (o ImageStorageProfileOutput) OsDisk() ImageOSDiskPtrOutput {
	return o.ApplyT(func(v ImageStorageProfile) *ImageOSDisk { return v.OsDisk }).(ImageOSDiskPtrOutput)
}

func (o ImageStorageProfileOutput) ZoneResilient() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ImageStorageProfile) *bool { return v.ZoneResilient }).(pulumi.BoolPtrOutput)
}

type ImageStorageProfilePtrOutput struct{ *pulumi.OutputState }

func (ImageStorageProfilePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ImageStorageProfile)(nil)).Elem()
}

func (o ImageStorageProfilePtrOutput) ToImageStorageProfilePtrOutput() ImageStorageProfilePtrOutput {
	return o
}

func (o ImageStorageProfilePtrOutput) ToImageStorageProfilePtrOutputWithContext(ctx context.Context) ImageStorageProfilePtrOutput {
	return o
}

func (o ImageStorageProfilePtrOutput) Elem() ImageStorageProfileOutput {
	return o.ApplyT(func(v *ImageStorageProfile) ImageStorageProfile {
		if v != nil {
			return *v
		}
		var ret ImageStorageProfile
		return ret
	}).(ImageStorageProfileOutput)
}

func (o ImageStorageProfilePtrOutput) DataDisks() ImageDataDiskArrayOutput {
	return o.ApplyT(func(v *ImageStorageProfile) []ImageDataDisk {
		if v == nil {
			return nil
		}
		return v.DataDisks
	}).(ImageDataDiskArrayOutput)
}

func (o ImageStorageProfilePtrOutput) OsDisk() ImageOSDiskPtrOutput {
	return o.ApplyT(func(v *ImageStorageProfile) *ImageOSDisk {
		if v == nil {
			return nil
		}
		return v.OsDisk
	}).(ImageOSDiskPtrOutput)
}

func (o ImageStorageProfilePtrOutput) ZoneResilient() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ImageStorageProfile) *bool {
		if v == nil {
			return nil
		}
		return v.ZoneResilient
	}).(pulumi.BoolPtrOutput)
}

type ImageStorageProfileResponse struct {
	DataDisks     []ImageDataDiskResponse `pulumi:"dataDisks"`
	OsDisk        *ImageOSDiskResponse    `pulumi:"osDisk"`
	ZoneResilient *bool                   `pulumi:"zoneResilient"`
}

type ImageStorageProfileResponseOutput struct{ *pulumi.OutputState }

func (ImageStorageProfileResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ImageStorageProfileResponse)(nil)).Elem()
}

func (o ImageStorageProfileResponseOutput) ToImageStorageProfileResponseOutput() ImageStorageProfileResponseOutput {
	return o
}

func (o ImageStorageProfileResponseOutput) ToImageStorageProfileResponseOutputWithContext(ctx context.Context) ImageStorageProfileResponseOutput {
	return o
}

func (o ImageStorageProfileResponseOutput) DataDisks() ImageDataDiskResponseArrayOutput {
	return o.ApplyT(func(v ImageStorageProfileResponse) []ImageDataDiskResponse { return v.DataDisks }).(ImageDataDiskResponseArrayOutput)
}

func (o ImageStorageProfileResponseOutput) OsDisk() ImageOSDiskResponsePtrOutput {
	return o.ApplyT(func(v ImageStorageProfileResponse) *ImageOSDiskResponse { return v.OsDisk }).(ImageOSDiskResponsePtrOutput)
}

func (o ImageStorageProfileResponseOutput) ZoneResilient() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ImageStorageProfileResponse) *bool { return v.ZoneResilient }).(pulumi.BoolPtrOutput)
}

type ImageStorageProfileResponsePtrOutput struct{ *pulumi.OutputState }

func (ImageStorageProfileResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ImageStorageProfileResponse)(nil)).Elem()
}

func (o ImageStorageProfileResponsePtrOutput) ToImageStorageProfileResponsePtrOutput() ImageStorageProfileResponsePtrOutput {
	return o
}

func (o ImageStorageProfileResponsePtrOutput) ToImageStorageProfileResponsePtrOutputWithContext(ctx context.Context) ImageStorageProfileResponsePtrOutput {
	return o
}

func (o ImageStorageProfileResponsePtrOutput) Elem() ImageStorageProfileResponseOutput {
	return o.ApplyT(func(v *ImageStorageProfileResponse) ImageStorageProfileResponse {
		if v != nil {
			return *v
		}
		var ret ImageStorageProfileResponse
		return ret
	}).(ImageStorageProfileResponseOutput)
}

func (o ImageStorageProfileResponsePtrOutput) DataDisks() ImageDataDiskResponseArrayOutput {
	return o.ApplyT(func(v *ImageStorageProfileResponse) []ImageDataDiskResponse {
		if v == nil {
			return nil
		}
		return v.DataDisks
	}).(ImageDataDiskResponseArrayOutput)
}

func (o ImageStorageProfileResponsePtrOutput) OsDisk() ImageOSDiskResponsePtrOutput {
	return o.ApplyT(func(v *ImageStorageProfileResponse) *ImageOSDiskResponse {
		if v == nil {
			return nil
		}
		return v.OsDisk
	}).(ImageOSDiskResponsePtrOutput)
}

func (o ImageStorageProfileResponsePtrOutput) ZoneResilient() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ImageStorageProfileResponse) *bool {
		if v == nil {
			return nil
		}
		return v.ZoneResilient
	}).(pulumi.BoolPtrOutput)
}

type InstanceViewStatus struct {
	Code          *string           `pulumi:"code"`
	DisplayStatus *string           `pulumi:"displayStatus"`
	Level         *StatusLevelTypes `pulumi:"level"`
	Message       *string           `pulumi:"message"`
	Time          *string           `pulumi:"time"`
}





type InstanceViewStatusInput interface {
	pulumi.Input

	ToInstanceViewStatusOutput() InstanceViewStatusOutput
	ToInstanceViewStatusOutputWithContext(context.Context) InstanceViewStatusOutput
}

type InstanceViewStatusArgs struct {
	Code          pulumi.StringPtrInput    `pulumi:"code"`
	DisplayStatus pulumi.StringPtrInput    `pulumi:"displayStatus"`
	Level         StatusLevelTypesPtrInput `pulumi:"level"`
	Message       pulumi.StringPtrInput    `pulumi:"message"`
	Time          pulumi.StringPtrInput    `pulumi:"time"`
}

func (InstanceViewStatusArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*InstanceViewStatus)(nil)).Elem()
}

func (i InstanceViewStatusArgs) ToInstanceViewStatusOutput() InstanceViewStatusOutput {
	return i.ToInstanceViewStatusOutputWithContext(context.Background())
}

func (i InstanceViewStatusArgs) ToInstanceViewStatusOutputWithContext(ctx context.Context) InstanceViewStatusOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceViewStatusOutput)
}





type InstanceViewStatusArrayInput interface {
	pulumi.Input

	ToInstanceViewStatusArrayOutput() InstanceViewStatusArrayOutput
	ToInstanceViewStatusArrayOutputWithContext(context.Context) InstanceViewStatusArrayOutput
}

type InstanceViewStatusArray []InstanceViewStatusInput

func (InstanceViewStatusArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]InstanceViewStatus)(nil)).Elem()
}

func (i InstanceViewStatusArray) ToInstanceViewStatusArrayOutput() InstanceViewStatusArrayOutput {
	return i.ToInstanceViewStatusArrayOutputWithContext(context.Background())
}

func (i InstanceViewStatusArray) ToInstanceViewStatusArrayOutputWithContext(ctx context.Context) InstanceViewStatusArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceViewStatusArrayOutput)
}

type InstanceViewStatusOutput struct{ *pulumi.OutputState }

func (InstanceViewStatusOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*InstanceViewStatus)(nil)).Elem()
}

func (o InstanceViewStatusOutput) ToInstanceViewStatusOutput() InstanceViewStatusOutput {
	return o
}

func (o InstanceViewStatusOutput) ToInstanceViewStatusOutputWithContext(ctx context.Context) InstanceViewStatusOutput {
	return o
}

func (o InstanceViewStatusOutput) Code() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InstanceViewStatus) *string { return v.Code }).(pulumi.StringPtrOutput)
}

func (o InstanceViewStatusOutput) DisplayStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InstanceViewStatus) *string { return v.DisplayStatus }).(pulumi.StringPtrOutput)
}

func (o InstanceViewStatusOutput) Level() StatusLevelTypesPtrOutput {
	return o.ApplyT(func(v InstanceViewStatus) *StatusLevelTypes { return v.Level }).(StatusLevelTypesPtrOutput)
}

func (o InstanceViewStatusOutput) Message() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InstanceViewStatus) *string { return v.Message }).(pulumi.StringPtrOutput)
}

func (o InstanceViewStatusOutput) Time() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InstanceViewStatus) *string { return v.Time }).(pulumi.StringPtrOutput)
}

type InstanceViewStatusArrayOutput struct{ *pulumi.OutputState }

func (InstanceViewStatusArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]InstanceViewStatus)(nil)).Elem()
}

func (o InstanceViewStatusArrayOutput) ToInstanceViewStatusArrayOutput() InstanceViewStatusArrayOutput {
	return o
}

func (o InstanceViewStatusArrayOutput) ToInstanceViewStatusArrayOutputWithContext(ctx context.Context) InstanceViewStatusArrayOutput {
	return o
}

func (o InstanceViewStatusArrayOutput) Index(i pulumi.IntInput) InstanceViewStatusOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) InstanceViewStatus {
		return vs[0].([]InstanceViewStatus)[vs[1].(int)]
	}).(InstanceViewStatusOutput)
}

type InstanceViewStatusResponse struct {
	Code          *string `pulumi:"code"`
	DisplayStatus *string `pulumi:"displayStatus"`
	Level         *string `pulumi:"level"`
	Message       *string `pulumi:"message"`
	Time          *string `pulumi:"time"`
}

type InstanceViewStatusResponseOutput struct{ *pulumi.OutputState }

func (InstanceViewStatusResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*InstanceViewStatusResponse)(nil)).Elem()
}

func (o InstanceViewStatusResponseOutput) ToInstanceViewStatusResponseOutput() InstanceViewStatusResponseOutput {
	return o
}

func (o InstanceViewStatusResponseOutput) ToInstanceViewStatusResponseOutputWithContext(ctx context.Context) InstanceViewStatusResponseOutput {
	return o
}

func (o InstanceViewStatusResponseOutput) Code() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InstanceViewStatusResponse) *string { return v.Code }).(pulumi.StringPtrOutput)
}

func (o InstanceViewStatusResponseOutput) DisplayStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InstanceViewStatusResponse) *string { return v.DisplayStatus }).(pulumi.StringPtrOutput)
}

func (o InstanceViewStatusResponseOutput) Level() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InstanceViewStatusResponse) *string { return v.Level }).(pulumi.StringPtrOutput)
}

func (o InstanceViewStatusResponseOutput) Message() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InstanceViewStatusResponse) *string { return v.Message }).(pulumi.StringPtrOutput)
}

func (o InstanceViewStatusResponseOutput) Time() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InstanceViewStatusResponse) *string { return v.Time }).(pulumi.StringPtrOutput)
}

type InstanceViewStatusResponsePtrOutput struct{ *pulumi.OutputState }

func (InstanceViewStatusResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**InstanceViewStatusResponse)(nil)).Elem()
}

func (o InstanceViewStatusResponsePtrOutput) ToInstanceViewStatusResponsePtrOutput() InstanceViewStatusResponsePtrOutput {
	return o
}

func (o InstanceViewStatusResponsePtrOutput) ToInstanceViewStatusResponsePtrOutputWithContext(ctx context.Context) InstanceViewStatusResponsePtrOutput {
	return o
}

func (o InstanceViewStatusResponsePtrOutput) Elem() InstanceViewStatusResponseOutput {
	return o.ApplyT(func(v *InstanceViewStatusResponse) InstanceViewStatusResponse {
		if v != nil {
			return *v
		}
		var ret InstanceViewStatusResponse
		return ret
	}).(InstanceViewStatusResponseOutput)
}

func (o InstanceViewStatusResponsePtrOutput) Code() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *InstanceViewStatusResponse) *string {
		if v == nil {
			return nil
		}
		return v.Code
	}).(pulumi.StringPtrOutput)
}

func (o InstanceViewStatusResponsePtrOutput) DisplayStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *InstanceViewStatusResponse) *string {
		if v == nil {
			return nil
		}
		return v.DisplayStatus
	}).(pulumi.StringPtrOutput)
}

func (o InstanceViewStatusResponsePtrOutput) Level() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *InstanceViewStatusResponse) *string {
		if v == nil {
			return nil
		}
		return v.Level
	}).(pulumi.StringPtrOutput)
}

func (o InstanceViewStatusResponsePtrOutput) Message() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *InstanceViewStatusResponse) *string {
		if v == nil {
			return nil
		}
		return v.Message
	}).(pulumi.StringPtrOutput)
}

func (o InstanceViewStatusResponsePtrOutput) Time() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *InstanceViewStatusResponse) *string {
		if v == nil {
			return nil
		}
		return v.Time
	}).(pulumi.StringPtrOutput)
}

type InstanceViewStatusResponseArrayOutput struct{ *pulumi.OutputState }

func (InstanceViewStatusResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]InstanceViewStatusResponse)(nil)).Elem()
}

func (o InstanceViewStatusResponseArrayOutput) ToInstanceViewStatusResponseArrayOutput() InstanceViewStatusResponseArrayOutput {
	return o
}

func (o InstanceViewStatusResponseArrayOutput) ToInstanceViewStatusResponseArrayOutputWithContext(ctx context.Context) InstanceViewStatusResponseArrayOutput {
	return o
}

func (o InstanceViewStatusResponseArrayOutput) Index(i pulumi.IntInput) InstanceViewStatusResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) InstanceViewStatusResponse {
		return vs[0].([]InstanceViewStatusResponse)[vs[1].(int)]
	}).(InstanceViewStatusResponseOutput)
}

type KeyVaultAndKeyReference struct {
	KeyUrl      string      `pulumi:"keyUrl"`
	SourceVault SourceVault `pulumi:"sourceVault"`
}





type KeyVaultAndKeyReferenceInput interface {
	pulumi.Input

	ToKeyVaultAndKeyReferenceOutput() KeyVaultAndKeyReferenceOutput
	ToKeyVaultAndKeyReferenceOutputWithContext(context.Context) KeyVaultAndKeyReferenceOutput
}

type KeyVaultAndKeyReferenceArgs struct {
	KeyUrl      pulumi.StringInput `pulumi:"keyUrl"`
	SourceVault SourceVaultInput   `pulumi:"sourceVault"`
}

func (KeyVaultAndKeyReferenceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*KeyVaultAndKeyReference)(nil)).Elem()
}

func (i KeyVaultAndKeyReferenceArgs) ToKeyVaultAndKeyReferenceOutput() KeyVaultAndKeyReferenceOutput {
	return i.ToKeyVaultAndKeyReferenceOutputWithContext(context.Background())
}

func (i KeyVaultAndKeyReferenceArgs) ToKeyVaultAndKeyReferenceOutputWithContext(ctx context.Context) KeyVaultAndKeyReferenceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyVaultAndKeyReferenceOutput)
}

func (i KeyVaultAndKeyReferenceArgs) ToKeyVaultAndKeyReferencePtrOutput() KeyVaultAndKeyReferencePtrOutput {
	return i.ToKeyVaultAndKeyReferencePtrOutputWithContext(context.Background())
}

func (i KeyVaultAndKeyReferenceArgs) ToKeyVaultAndKeyReferencePtrOutputWithContext(ctx context.Context) KeyVaultAndKeyReferencePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyVaultAndKeyReferenceOutput).ToKeyVaultAndKeyReferencePtrOutputWithContext(ctx)
}









type KeyVaultAndKeyReferencePtrInput interface {
	pulumi.Input

	ToKeyVaultAndKeyReferencePtrOutput() KeyVaultAndKeyReferencePtrOutput
	ToKeyVaultAndKeyReferencePtrOutputWithContext(context.Context) KeyVaultAndKeyReferencePtrOutput
}

type keyVaultAndKeyReferencePtrType KeyVaultAndKeyReferenceArgs

func KeyVaultAndKeyReferencePtr(v *KeyVaultAndKeyReferenceArgs) KeyVaultAndKeyReferencePtrInput {
	return (*keyVaultAndKeyReferencePtrType)(v)
}

func (*keyVaultAndKeyReferencePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**KeyVaultAndKeyReference)(nil)).Elem()
}

func (i *keyVaultAndKeyReferencePtrType) ToKeyVaultAndKeyReferencePtrOutput() KeyVaultAndKeyReferencePtrOutput {
	return i.ToKeyVaultAndKeyReferencePtrOutputWithContext(context.Background())
}

func (i *keyVaultAndKeyReferencePtrType) ToKeyVaultAndKeyReferencePtrOutputWithContext(ctx context.Context) KeyVaultAndKeyReferencePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyVaultAndKeyReferencePtrOutput)
}

type KeyVaultAndKeyReferenceOutput struct{ *pulumi.OutputState }

func (KeyVaultAndKeyReferenceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*KeyVaultAndKeyReference)(nil)).Elem()
}

func (o KeyVaultAndKeyReferenceOutput) ToKeyVaultAndKeyReferenceOutput() KeyVaultAndKeyReferenceOutput {
	return o
}

func (o KeyVaultAndKeyReferenceOutput) ToKeyVaultAndKeyReferenceOutputWithContext(ctx context.Context) KeyVaultAndKeyReferenceOutput {
	return o
}

func (o KeyVaultAndKeyReferenceOutput) ToKeyVaultAndKeyReferencePtrOutput() KeyVaultAndKeyReferencePtrOutput {
	return o.ToKeyVaultAndKeyReferencePtrOutputWithContext(context.Background())
}

func (o KeyVaultAndKeyReferenceOutput) ToKeyVaultAndKeyReferencePtrOutputWithContext(ctx context.Context) KeyVaultAndKeyReferencePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v KeyVaultAndKeyReference) *KeyVaultAndKeyReference {
		return &v
	}).(KeyVaultAndKeyReferencePtrOutput)
}

func (o KeyVaultAndKeyReferenceOutput) KeyUrl() pulumi.StringOutput {
	return o.ApplyT(func(v KeyVaultAndKeyReference) string { return v.KeyUrl }).(pulumi.StringOutput)
}

func (o KeyVaultAndKeyReferenceOutput) SourceVault() SourceVaultOutput {
	return o.ApplyT(func(v KeyVaultAndKeyReference) SourceVault { return v.SourceVault }).(SourceVaultOutput)
}

type KeyVaultAndKeyReferencePtrOutput struct{ *pulumi.OutputState }

func (KeyVaultAndKeyReferencePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**KeyVaultAndKeyReference)(nil)).Elem()
}

func (o KeyVaultAndKeyReferencePtrOutput) ToKeyVaultAndKeyReferencePtrOutput() KeyVaultAndKeyReferencePtrOutput {
	return o
}

func (o KeyVaultAndKeyReferencePtrOutput) ToKeyVaultAndKeyReferencePtrOutputWithContext(ctx context.Context) KeyVaultAndKeyReferencePtrOutput {
	return o
}

func (o KeyVaultAndKeyReferencePtrOutput) Elem() KeyVaultAndKeyReferenceOutput {
	return o.ApplyT(func(v *KeyVaultAndKeyReference) KeyVaultAndKeyReference {
		if v != nil {
			return *v
		}
		var ret KeyVaultAndKeyReference
		return ret
	}).(KeyVaultAndKeyReferenceOutput)
}

func (o KeyVaultAndKeyReferencePtrOutput) KeyUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KeyVaultAndKeyReference) *string {
		if v == nil {
			return nil
		}
		return &v.KeyUrl
	}).(pulumi.StringPtrOutput)
}

func (o KeyVaultAndKeyReferencePtrOutput) SourceVault() SourceVaultPtrOutput {
	return o.ApplyT(func(v *KeyVaultAndKeyReference) *SourceVault {
		if v == nil {
			return nil
		}
		return &v.SourceVault
	}).(SourceVaultPtrOutput)
}

type KeyVaultAndKeyReferenceResponse struct {
	KeyUrl      string              `pulumi:"keyUrl"`
	SourceVault SourceVaultResponse `pulumi:"sourceVault"`
}

type KeyVaultAndKeyReferenceResponseOutput struct{ *pulumi.OutputState }

func (KeyVaultAndKeyReferenceResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*KeyVaultAndKeyReferenceResponse)(nil)).Elem()
}

func (o KeyVaultAndKeyReferenceResponseOutput) ToKeyVaultAndKeyReferenceResponseOutput() KeyVaultAndKeyReferenceResponseOutput {
	return o
}

func (o KeyVaultAndKeyReferenceResponseOutput) ToKeyVaultAndKeyReferenceResponseOutputWithContext(ctx context.Context) KeyVaultAndKeyReferenceResponseOutput {
	return o
}

func (o KeyVaultAndKeyReferenceResponseOutput) KeyUrl() pulumi.StringOutput {
	return o.ApplyT(func(v KeyVaultAndKeyReferenceResponse) string { return v.KeyUrl }).(pulumi.StringOutput)
}

func (o KeyVaultAndKeyReferenceResponseOutput) SourceVault() SourceVaultResponseOutput {
	return o.ApplyT(func(v KeyVaultAndKeyReferenceResponse) SourceVaultResponse { return v.SourceVault }).(SourceVaultResponseOutput)
}

type KeyVaultAndKeyReferenceResponsePtrOutput struct{ *pulumi.OutputState }

func (KeyVaultAndKeyReferenceResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**KeyVaultAndKeyReferenceResponse)(nil)).Elem()
}

func (o KeyVaultAndKeyReferenceResponsePtrOutput) ToKeyVaultAndKeyReferenceResponsePtrOutput() KeyVaultAndKeyReferenceResponsePtrOutput {
	return o
}

func (o KeyVaultAndKeyReferenceResponsePtrOutput) ToKeyVaultAndKeyReferenceResponsePtrOutputWithContext(ctx context.Context) KeyVaultAndKeyReferenceResponsePtrOutput {
	return o
}

func (o KeyVaultAndKeyReferenceResponsePtrOutput) Elem() KeyVaultAndKeyReferenceResponseOutput {
	return o.ApplyT(func(v *KeyVaultAndKeyReferenceResponse) KeyVaultAndKeyReferenceResponse {
		if v != nil {
			return *v
		}
		var ret KeyVaultAndKeyReferenceResponse
		return ret
	}).(KeyVaultAndKeyReferenceResponseOutput)
}

func (o KeyVaultAndKeyReferenceResponsePtrOutput) KeyUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KeyVaultAndKeyReferenceResponse) *string {
		if v == nil {
			return nil
		}
		return &v.KeyUrl
	}).(pulumi.StringPtrOutput)
}

func (o KeyVaultAndKeyReferenceResponsePtrOutput) SourceVault() SourceVaultResponsePtrOutput {
	return o.ApplyT(func(v *KeyVaultAndKeyReferenceResponse) *SourceVaultResponse {
		if v == nil {
			return nil
		}
		return &v.SourceVault
	}).(SourceVaultResponsePtrOutput)
}

type KeyVaultAndSecretReference struct {
	SecretUrl   string      `pulumi:"secretUrl"`
	SourceVault SourceVault `pulumi:"sourceVault"`
}





type KeyVaultAndSecretReferenceInput interface {
	pulumi.Input

	ToKeyVaultAndSecretReferenceOutput() KeyVaultAndSecretReferenceOutput
	ToKeyVaultAndSecretReferenceOutputWithContext(context.Context) KeyVaultAndSecretReferenceOutput
}

type KeyVaultAndSecretReferenceArgs struct {
	SecretUrl   pulumi.StringInput `pulumi:"secretUrl"`
	SourceVault SourceVaultInput   `pulumi:"sourceVault"`
}

func (KeyVaultAndSecretReferenceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*KeyVaultAndSecretReference)(nil)).Elem()
}

func (i KeyVaultAndSecretReferenceArgs) ToKeyVaultAndSecretReferenceOutput() KeyVaultAndSecretReferenceOutput {
	return i.ToKeyVaultAndSecretReferenceOutputWithContext(context.Background())
}

func (i KeyVaultAndSecretReferenceArgs) ToKeyVaultAndSecretReferenceOutputWithContext(ctx context.Context) KeyVaultAndSecretReferenceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyVaultAndSecretReferenceOutput)
}

func (i KeyVaultAndSecretReferenceArgs) ToKeyVaultAndSecretReferencePtrOutput() KeyVaultAndSecretReferencePtrOutput {
	return i.ToKeyVaultAndSecretReferencePtrOutputWithContext(context.Background())
}

func (i KeyVaultAndSecretReferenceArgs) ToKeyVaultAndSecretReferencePtrOutputWithContext(ctx context.Context) KeyVaultAndSecretReferencePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyVaultAndSecretReferenceOutput).ToKeyVaultAndSecretReferencePtrOutputWithContext(ctx)
}









type KeyVaultAndSecretReferencePtrInput interface {
	pulumi.Input

	ToKeyVaultAndSecretReferencePtrOutput() KeyVaultAndSecretReferencePtrOutput
	ToKeyVaultAndSecretReferencePtrOutputWithContext(context.Context) KeyVaultAndSecretReferencePtrOutput
}

type keyVaultAndSecretReferencePtrType KeyVaultAndSecretReferenceArgs

func KeyVaultAndSecretReferencePtr(v *KeyVaultAndSecretReferenceArgs) KeyVaultAndSecretReferencePtrInput {
	return (*keyVaultAndSecretReferencePtrType)(v)
}

func (*keyVaultAndSecretReferencePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**KeyVaultAndSecretReference)(nil)).Elem()
}

func (i *keyVaultAndSecretReferencePtrType) ToKeyVaultAndSecretReferencePtrOutput() KeyVaultAndSecretReferencePtrOutput {
	return i.ToKeyVaultAndSecretReferencePtrOutputWithContext(context.Background())
}

func (i *keyVaultAndSecretReferencePtrType) ToKeyVaultAndSecretReferencePtrOutputWithContext(ctx context.Context) KeyVaultAndSecretReferencePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyVaultAndSecretReferencePtrOutput)
}

type KeyVaultAndSecretReferenceOutput struct{ *pulumi.OutputState }

func (KeyVaultAndSecretReferenceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*KeyVaultAndSecretReference)(nil)).Elem()
}

func (o KeyVaultAndSecretReferenceOutput) ToKeyVaultAndSecretReferenceOutput() KeyVaultAndSecretReferenceOutput {
	return o
}

func (o KeyVaultAndSecretReferenceOutput) ToKeyVaultAndSecretReferenceOutputWithContext(ctx context.Context) KeyVaultAndSecretReferenceOutput {
	return o
}

func (o KeyVaultAndSecretReferenceOutput) ToKeyVaultAndSecretReferencePtrOutput() KeyVaultAndSecretReferencePtrOutput {
	return o.ToKeyVaultAndSecretReferencePtrOutputWithContext(context.Background())
}

func (o KeyVaultAndSecretReferenceOutput) ToKeyVaultAndSecretReferencePtrOutputWithContext(ctx context.Context) KeyVaultAndSecretReferencePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v KeyVaultAndSecretReference) *KeyVaultAndSecretReference {
		return &v
	}).(KeyVaultAndSecretReferencePtrOutput)
}

func (o KeyVaultAndSecretReferenceOutput) SecretUrl() pulumi.StringOutput {
	return o.ApplyT(func(v KeyVaultAndSecretReference) string { return v.SecretUrl }).(pulumi.StringOutput)
}

func (o KeyVaultAndSecretReferenceOutput) SourceVault() SourceVaultOutput {
	return o.ApplyT(func(v KeyVaultAndSecretReference) SourceVault { return v.SourceVault }).(SourceVaultOutput)
}

type KeyVaultAndSecretReferencePtrOutput struct{ *pulumi.OutputState }

func (KeyVaultAndSecretReferencePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**KeyVaultAndSecretReference)(nil)).Elem()
}

func (o KeyVaultAndSecretReferencePtrOutput) ToKeyVaultAndSecretReferencePtrOutput() KeyVaultAndSecretReferencePtrOutput {
	return o
}

func (o KeyVaultAndSecretReferencePtrOutput) ToKeyVaultAndSecretReferencePtrOutputWithContext(ctx context.Context) KeyVaultAndSecretReferencePtrOutput {
	return o
}

func (o KeyVaultAndSecretReferencePtrOutput) Elem() KeyVaultAndSecretReferenceOutput {
	return o.ApplyT(func(v *KeyVaultAndSecretReference) KeyVaultAndSecretReference {
		if v != nil {
			return *v
		}
		var ret KeyVaultAndSecretReference
		return ret
	}).(KeyVaultAndSecretReferenceOutput)
}

func (o KeyVaultAndSecretReferencePtrOutput) SecretUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KeyVaultAndSecretReference) *string {
		if v == nil {
			return nil
		}
		return &v.SecretUrl
	}).(pulumi.StringPtrOutput)
}

func (o KeyVaultAndSecretReferencePtrOutput) SourceVault() SourceVaultPtrOutput {
	return o.ApplyT(func(v *KeyVaultAndSecretReference) *SourceVault {
		if v == nil {
			return nil
		}
		return &v.SourceVault
	}).(SourceVaultPtrOutput)
}

type KeyVaultAndSecretReferenceResponse struct {
	SecretUrl   string              `pulumi:"secretUrl"`
	SourceVault SourceVaultResponse `pulumi:"sourceVault"`
}

type KeyVaultAndSecretReferenceResponseOutput struct{ *pulumi.OutputState }

func (KeyVaultAndSecretReferenceResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*KeyVaultAndSecretReferenceResponse)(nil)).Elem()
}

func (o KeyVaultAndSecretReferenceResponseOutput) ToKeyVaultAndSecretReferenceResponseOutput() KeyVaultAndSecretReferenceResponseOutput {
	return o
}

func (o KeyVaultAndSecretReferenceResponseOutput) ToKeyVaultAndSecretReferenceResponseOutputWithContext(ctx context.Context) KeyVaultAndSecretReferenceResponseOutput {
	return o
}

func (o KeyVaultAndSecretReferenceResponseOutput) SecretUrl() pulumi.StringOutput {
	return o.ApplyT(func(v KeyVaultAndSecretReferenceResponse) string { return v.SecretUrl }).(pulumi.StringOutput)
}

func (o KeyVaultAndSecretReferenceResponseOutput) SourceVault() SourceVaultResponseOutput {
	return o.ApplyT(func(v KeyVaultAndSecretReferenceResponse) SourceVaultResponse { return v.SourceVault }).(SourceVaultResponseOutput)
}

type KeyVaultAndSecretReferenceResponsePtrOutput struct{ *pulumi.OutputState }

func (KeyVaultAndSecretReferenceResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**KeyVaultAndSecretReferenceResponse)(nil)).Elem()
}

func (o KeyVaultAndSecretReferenceResponsePtrOutput) ToKeyVaultAndSecretReferenceResponsePtrOutput() KeyVaultAndSecretReferenceResponsePtrOutput {
	return o
}

func (o KeyVaultAndSecretReferenceResponsePtrOutput) ToKeyVaultAndSecretReferenceResponsePtrOutputWithContext(ctx context.Context) KeyVaultAndSecretReferenceResponsePtrOutput {
	return o
}

func (o KeyVaultAndSecretReferenceResponsePtrOutput) Elem() KeyVaultAndSecretReferenceResponseOutput {
	return o.ApplyT(func(v *KeyVaultAndSecretReferenceResponse) KeyVaultAndSecretReferenceResponse {
		if v != nil {
			return *v
		}
		var ret KeyVaultAndSecretReferenceResponse
		return ret
	}).(KeyVaultAndSecretReferenceResponseOutput)
}

func (o KeyVaultAndSecretReferenceResponsePtrOutput) SecretUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KeyVaultAndSecretReferenceResponse) *string {
		if v == nil {
			return nil
		}
		return &v.SecretUrl
	}).(pulumi.StringPtrOutput)
}

func (o KeyVaultAndSecretReferenceResponsePtrOutput) SourceVault() SourceVaultResponsePtrOutput {
	return o.ApplyT(func(v *KeyVaultAndSecretReferenceResponse) *SourceVaultResponse {
		if v == nil {
			return nil
		}
		return &v.SourceVault
	}).(SourceVaultResponsePtrOutput)
}

type KeyVaultKeyReference struct {
	KeyUrl      string      `pulumi:"keyUrl"`
	SourceVault SubResource `pulumi:"sourceVault"`
}





type KeyVaultKeyReferenceInput interface {
	pulumi.Input

	ToKeyVaultKeyReferenceOutput() KeyVaultKeyReferenceOutput
	ToKeyVaultKeyReferenceOutputWithContext(context.Context) KeyVaultKeyReferenceOutput
}

type KeyVaultKeyReferenceArgs struct {
	KeyUrl      pulumi.StringInput `pulumi:"keyUrl"`
	SourceVault SubResourceInput   `pulumi:"sourceVault"`
}

func (KeyVaultKeyReferenceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*KeyVaultKeyReference)(nil)).Elem()
}

func (i KeyVaultKeyReferenceArgs) ToKeyVaultKeyReferenceOutput() KeyVaultKeyReferenceOutput {
	return i.ToKeyVaultKeyReferenceOutputWithContext(context.Background())
}

func (i KeyVaultKeyReferenceArgs) ToKeyVaultKeyReferenceOutputWithContext(ctx context.Context) KeyVaultKeyReferenceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyVaultKeyReferenceOutput)
}

func (i KeyVaultKeyReferenceArgs) ToKeyVaultKeyReferencePtrOutput() KeyVaultKeyReferencePtrOutput {
	return i.ToKeyVaultKeyReferencePtrOutputWithContext(context.Background())
}

func (i KeyVaultKeyReferenceArgs) ToKeyVaultKeyReferencePtrOutputWithContext(ctx context.Context) KeyVaultKeyReferencePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyVaultKeyReferenceOutput).ToKeyVaultKeyReferencePtrOutputWithContext(ctx)
}









type KeyVaultKeyReferencePtrInput interface {
	pulumi.Input

	ToKeyVaultKeyReferencePtrOutput() KeyVaultKeyReferencePtrOutput
	ToKeyVaultKeyReferencePtrOutputWithContext(context.Context) KeyVaultKeyReferencePtrOutput
}

type keyVaultKeyReferencePtrType KeyVaultKeyReferenceArgs

func KeyVaultKeyReferencePtr(v *KeyVaultKeyReferenceArgs) KeyVaultKeyReferencePtrInput {
	return (*keyVaultKeyReferencePtrType)(v)
}

func (*keyVaultKeyReferencePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**KeyVaultKeyReference)(nil)).Elem()
}

func (i *keyVaultKeyReferencePtrType) ToKeyVaultKeyReferencePtrOutput() KeyVaultKeyReferencePtrOutput {
	return i.ToKeyVaultKeyReferencePtrOutputWithContext(context.Background())
}

func (i *keyVaultKeyReferencePtrType) ToKeyVaultKeyReferencePtrOutputWithContext(ctx context.Context) KeyVaultKeyReferencePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyVaultKeyReferencePtrOutput)
}

type KeyVaultKeyReferenceOutput struct{ *pulumi.OutputState }

func (KeyVaultKeyReferenceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*KeyVaultKeyReference)(nil)).Elem()
}

func (o KeyVaultKeyReferenceOutput) ToKeyVaultKeyReferenceOutput() KeyVaultKeyReferenceOutput {
	return o
}

func (o KeyVaultKeyReferenceOutput) ToKeyVaultKeyReferenceOutputWithContext(ctx context.Context) KeyVaultKeyReferenceOutput {
	return o
}

func (o KeyVaultKeyReferenceOutput) ToKeyVaultKeyReferencePtrOutput() KeyVaultKeyReferencePtrOutput {
	return o.ToKeyVaultKeyReferencePtrOutputWithContext(context.Background())
}

func (o KeyVaultKeyReferenceOutput) ToKeyVaultKeyReferencePtrOutputWithContext(ctx context.Context) KeyVaultKeyReferencePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v KeyVaultKeyReference) *KeyVaultKeyReference {
		return &v
	}).(KeyVaultKeyReferencePtrOutput)
}

func (o KeyVaultKeyReferenceOutput) KeyUrl() pulumi.StringOutput {
	return o.ApplyT(func(v KeyVaultKeyReference) string { return v.KeyUrl }).(pulumi.StringOutput)
}

func (o KeyVaultKeyReferenceOutput) SourceVault() SubResourceOutput {
	return o.ApplyT(func(v KeyVaultKeyReference) SubResource { return v.SourceVault }).(SubResourceOutput)
}

type KeyVaultKeyReferencePtrOutput struct{ *pulumi.OutputState }

func (KeyVaultKeyReferencePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**KeyVaultKeyReference)(nil)).Elem()
}

func (o KeyVaultKeyReferencePtrOutput) ToKeyVaultKeyReferencePtrOutput() KeyVaultKeyReferencePtrOutput {
	return o
}

func (o KeyVaultKeyReferencePtrOutput) ToKeyVaultKeyReferencePtrOutputWithContext(ctx context.Context) KeyVaultKeyReferencePtrOutput {
	return o
}

func (o KeyVaultKeyReferencePtrOutput) Elem() KeyVaultKeyReferenceOutput {
	return o.ApplyT(func(v *KeyVaultKeyReference) KeyVaultKeyReference {
		if v != nil {
			return *v
		}
		var ret KeyVaultKeyReference
		return ret
	}).(KeyVaultKeyReferenceOutput)
}

func (o KeyVaultKeyReferencePtrOutput) KeyUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KeyVaultKeyReference) *string {
		if v == nil {
			return nil
		}
		return &v.KeyUrl
	}).(pulumi.StringPtrOutput)
}

func (o KeyVaultKeyReferencePtrOutput) SourceVault() SubResourcePtrOutput {
	return o.ApplyT(func(v *KeyVaultKeyReference) *SubResource {
		if v == nil {
			return nil
		}
		return &v.SourceVault
	}).(SubResourcePtrOutput)
}

type KeyVaultKeyReferenceResponse struct {
	KeyUrl      string              `pulumi:"keyUrl"`
	SourceVault SubResourceResponse `pulumi:"sourceVault"`
}

type KeyVaultKeyReferenceResponseOutput struct{ *pulumi.OutputState }

func (KeyVaultKeyReferenceResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*KeyVaultKeyReferenceResponse)(nil)).Elem()
}

func (o KeyVaultKeyReferenceResponseOutput) ToKeyVaultKeyReferenceResponseOutput() KeyVaultKeyReferenceResponseOutput {
	return o
}

func (o KeyVaultKeyReferenceResponseOutput) ToKeyVaultKeyReferenceResponseOutputWithContext(ctx context.Context) KeyVaultKeyReferenceResponseOutput {
	return o
}

func (o KeyVaultKeyReferenceResponseOutput) KeyUrl() pulumi.StringOutput {
	return o.ApplyT(func(v KeyVaultKeyReferenceResponse) string { return v.KeyUrl }).(pulumi.StringOutput)
}

func (o KeyVaultKeyReferenceResponseOutput) SourceVault() SubResourceResponseOutput {
	return o.ApplyT(func(v KeyVaultKeyReferenceResponse) SubResourceResponse { return v.SourceVault }).(SubResourceResponseOutput)
}

type KeyVaultKeyReferenceResponsePtrOutput struct{ *pulumi.OutputState }

func (KeyVaultKeyReferenceResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**KeyVaultKeyReferenceResponse)(nil)).Elem()
}

func (o KeyVaultKeyReferenceResponsePtrOutput) ToKeyVaultKeyReferenceResponsePtrOutput() KeyVaultKeyReferenceResponsePtrOutput {
	return o
}

func (o KeyVaultKeyReferenceResponsePtrOutput) ToKeyVaultKeyReferenceResponsePtrOutputWithContext(ctx context.Context) KeyVaultKeyReferenceResponsePtrOutput {
	return o
}

func (o KeyVaultKeyReferenceResponsePtrOutput) Elem() KeyVaultKeyReferenceResponseOutput {
	return o.ApplyT(func(v *KeyVaultKeyReferenceResponse) KeyVaultKeyReferenceResponse {
		if v != nil {
			return *v
		}
		var ret KeyVaultKeyReferenceResponse
		return ret
	}).(KeyVaultKeyReferenceResponseOutput)
}

func (o KeyVaultKeyReferenceResponsePtrOutput) KeyUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KeyVaultKeyReferenceResponse) *string {
		if v == nil {
			return nil
		}
		return &v.KeyUrl
	}).(pulumi.StringPtrOutput)
}

func (o KeyVaultKeyReferenceResponsePtrOutput) SourceVault() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v *KeyVaultKeyReferenceResponse) *SubResourceResponse {
		if v == nil {
			return nil
		}
		return &v.SourceVault
	}).(SubResourceResponsePtrOutput)
}

type KeyVaultSecretReference struct {
	SecretUrl   string      `pulumi:"secretUrl"`
	SourceVault SubResource `pulumi:"sourceVault"`
}





type KeyVaultSecretReferenceInput interface {
	pulumi.Input

	ToKeyVaultSecretReferenceOutput() KeyVaultSecretReferenceOutput
	ToKeyVaultSecretReferenceOutputWithContext(context.Context) KeyVaultSecretReferenceOutput
}

type KeyVaultSecretReferenceArgs struct {
	SecretUrl   pulumi.StringInput `pulumi:"secretUrl"`
	SourceVault SubResourceInput   `pulumi:"sourceVault"`
}

func (KeyVaultSecretReferenceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*KeyVaultSecretReference)(nil)).Elem()
}

func (i KeyVaultSecretReferenceArgs) ToKeyVaultSecretReferenceOutput() KeyVaultSecretReferenceOutput {
	return i.ToKeyVaultSecretReferenceOutputWithContext(context.Background())
}

func (i KeyVaultSecretReferenceArgs) ToKeyVaultSecretReferenceOutputWithContext(ctx context.Context) KeyVaultSecretReferenceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyVaultSecretReferenceOutput)
}

func (i KeyVaultSecretReferenceArgs) ToKeyVaultSecretReferencePtrOutput() KeyVaultSecretReferencePtrOutput {
	return i.ToKeyVaultSecretReferencePtrOutputWithContext(context.Background())
}

func (i KeyVaultSecretReferenceArgs) ToKeyVaultSecretReferencePtrOutputWithContext(ctx context.Context) KeyVaultSecretReferencePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyVaultSecretReferenceOutput).ToKeyVaultSecretReferencePtrOutputWithContext(ctx)
}









type KeyVaultSecretReferencePtrInput interface {
	pulumi.Input

	ToKeyVaultSecretReferencePtrOutput() KeyVaultSecretReferencePtrOutput
	ToKeyVaultSecretReferencePtrOutputWithContext(context.Context) KeyVaultSecretReferencePtrOutput
}

type keyVaultSecretReferencePtrType KeyVaultSecretReferenceArgs

func KeyVaultSecretReferencePtr(v *KeyVaultSecretReferenceArgs) KeyVaultSecretReferencePtrInput {
	return (*keyVaultSecretReferencePtrType)(v)
}

func (*keyVaultSecretReferencePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**KeyVaultSecretReference)(nil)).Elem()
}

func (i *keyVaultSecretReferencePtrType) ToKeyVaultSecretReferencePtrOutput() KeyVaultSecretReferencePtrOutput {
	return i.ToKeyVaultSecretReferencePtrOutputWithContext(context.Background())
}

func (i *keyVaultSecretReferencePtrType) ToKeyVaultSecretReferencePtrOutputWithContext(ctx context.Context) KeyVaultSecretReferencePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyVaultSecretReferencePtrOutput)
}

type KeyVaultSecretReferenceOutput struct{ *pulumi.OutputState }

func (KeyVaultSecretReferenceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*KeyVaultSecretReference)(nil)).Elem()
}

func (o KeyVaultSecretReferenceOutput) ToKeyVaultSecretReferenceOutput() KeyVaultSecretReferenceOutput {
	return o
}

func (o KeyVaultSecretReferenceOutput) ToKeyVaultSecretReferenceOutputWithContext(ctx context.Context) KeyVaultSecretReferenceOutput {
	return o
}

func (o KeyVaultSecretReferenceOutput) ToKeyVaultSecretReferencePtrOutput() KeyVaultSecretReferencePtrOutput {
	return o.ToKeyVaultSecretReferencePtrOutputWithContext(context.Background())
}

func (o KeyVaultSecretReferenceOutput) ToKeyVaultSecretReferencePtrOutputWithContext(ctx context.Context) KeyVaultSecretReferencePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v KeyVaultSecretReference) *KeyVaultSecretReference {
		return &v
	}).(KeyVaultSecretReferencePtrOutput)
}

func (o KeyVaultSecretReferenceOutput) SecretUrl() pulumi.StringOutput {
	return o.ApplyT(func(v KeyVaultSecretReference) string { return v.SecretUrl }).(pulumi.StringOutput)
}

func (o KeyVaultSecretReferenceOutput) SourceVault() SubResourceOutput {
	return o.ApplyT(func(v KeyVaultSecretReference) SubResource { return v.SourceVault }).(SubResourceOutput)
}

type KeyVaultSecretReferencePtrOutput struct{ *pulumi.OutputState }

func (KeyVaultSecretReferencePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**KeyVaultSecretReference)(nil)).Elem()
}

func (o KeyVaultSecretReferencePtrOutput) ToKeyVaultSecretReferencePtrOutput() KeyVaultSecretReferencePtrOutput {
	return o
}

func (o KeyVaultSecretReferencePtrOutput) ToKeyVaultSecretReferencePtrOutputWithContext(ctx context.Context) KeyVaultSecretReferencePtrOutput {
	return o
}

func (o KeyVaultSecretReferencePtrOutput) Elem() KeyVaultSecretReferenceOutput {
	return o.ApplyT(func(v *KeyVaultSecretReference) KeyVaultSecretReference {
		if v != nil {
			return *v
		}
		var ret KeyVaultSecretReference
		return ret
	}).(KeyVaultSecretReferenceOutput)
}

func (o KeyVaultSecretReferencePtrOutput) SecretUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KeyVaultSecretReference) *string {
		if v == nil {
			return nil
		}
		return &v.SecretUrl
	}).(pulumi.StringPtrOutput)
}

func (o KeyVaultSecretReferencePtrOutput) SourceVault() SubResourcePtrOutput {
	return o.ApplyT(func(v *KeyVaultSecretReference) *SubResource {
		if v == nil {
			return nil
		}
		return &v.SourceVault
	}).(SubResourcePtrOutput)
}

type KeyVaultSecretReferenceResponse struct {
	SecretUrl   string              `pulumi:"secretUrl"`
	SourceVault SubResourceResponse `pulumi:"sourceVault"`
}

type KeyVaultSecretReferenceResponseOutput struct{ *pulumi.OutputState }

func (KeyVaultSecretReferenceResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*KeyVaultSecretReferenceResponse)(nil)).Elem()
}

func (o KeyVaultSecretReferenceResponseOutput) ToKeyVaultSecretReferenceResponseOutput() KeyVaultSecretReferenceResponseOutput {
	return o
}

func (o KeyVaultSecretReferenceResponseOutput) ToKeyVaultSecretReferenceResponseOutputWithContext(ctx context.Context) KeyVaultSecretReferenceResponseOutput {
	return o
}

func (o KeyVaultSecretReferenceResponseOutput) SecretUrl() pulumi.StringOutput {
	return o.ApplyT(func(v KeyVaultSecretReferenceResponse) string { return v.SecretUrl }).(pulumi.StringOutput)
}

func (o KeyVaultSecretReferenceResponseOutput) SourceVault() SubResourceResponseOutput {
	return o.ApplyT(func(v KeyVaultSecretReferenceResponse) SubResourceResponse { return v.SourceVault }).(SubResourceResponseOutput)
}

type KeyVaultSecretReferenceResponsePtrOutput struct{ *pulumi.OutputState }

func (KeyVaultSecretReferenceResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**KeyVaultSecretReferenceResponse)(nil)).Elem()
}

func (o KeyVaultSecretReferenceResponsePtrOutput) ToKeyVaultSecretReferenceResponsePtrOutput() KeyVaultSecretReferenceResponsePtrOutput {
	return o
}

func (o KeyVaultSecretReferenceResponsePtrOutput) ToKeyVaultSecretReferenceResponsePtrOutputWithContext(ctx context.Context) KeyVaultSecretReferenceResponsePtrOutput {
	return o
}

func (o KeyVaultSecretReferenceResponsePtrOutput) Elem() KeyVaultSecretReferenceResponseOutput {
	return o.ApplyT(func(v *KeyVaultSecretReferenceResponse) KeyVaultSecretReferenceResponse {
		if v != nil {
			return *v
		}
		var ret KeyVaultSecretReferenceResponse
		return ret
	}).(KeyVaultSecretReferenceResponseOutput)
}

func (o KeyVaultSecretReferenceResponsePtrOutput) SecretUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KeyVaultSecretReferenceResponse) *string {
		if v == nil {
			return nil
		}
		return &v.SecretUrl
	}).(pulumi.StringPtrOutput)
}

func (o KeyVaultSecretReferenceResponsePtrOutput) SourceVault() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v *KeyVaultSecretReferenceResponse) *SubResourceResponse {
		if v == nil {
			return nil
		}
		return &v.SourceVault
	}).(SubResourceResponsePtrOutput)
}

type LinuxConfiguration struct {
	DisablePasswordAuthentication *bool             `pulumi:"disablePasswordAuthentication"`
	Ssh                           *SshConfiguration `pulumi:"ssh"`
}





type LinuxConfigurationInput interface {
	pulumi.Input

	ToLinuxConfigurationOutput() LinuxConfigurationOutput
	ToLinuxConfigurationOutputWithContext(context.Context) LinuxConfigurationOutput
}

type LinuxConfigurationArgs struct {
	DisablePasswordAuthentication pulumi.BoolPtrInput      `pulumi:"disablePasswordAuthentication"`
	Ssh                           SshConfigurationPtrInput `pulumi:"ssh"`
}

func (LinuxConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LinuxConfiguration)(nil)).Elem()
}

func (i LinuxConfigurationArgs) ToLinuxConfigurationOutput() LinuxConfigurationOutput {
	return i.ToLinuxConfigurationOutputWithContext(context.Background())
}

func (i LinuxConfigurationArgs) ToLinuxConfigurationOutputWithContext(ctx context.Context) LinuxConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LinuxConfigurationOutput)
}

func (i LinuxConfigurationArgs) ToLinuxConfigurationPtrOutput() LinuxConfigurationPtrOutput {
	return i.ToLinuxConfigurationPtrOutputWithContext(context.Background())
}

func (i LinuxConfigurationArgs) ToLinuxConfigurationPtrOutputWithContext(ctx context.Context) LinuxConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LinuxConfigurationOutput).ToLinuxConfigurationPtrOutputWithContext(ctx)
}









type LinuxConfigurationPtrInput interface {
	pulumi.Input

	ToLinuxConfigurationPtrOutput() LinuxConfigurationPtrOutput
	ToLinuxConfigurationPtrOutputWithContext(context.Context) LinuxConfigurationPtrOutput
}

type linuxConfigurationPtrType LinuxConfigurationArgs

func LinuxConfigurationPtr(v *LinuxConfigurationArgs) LinuxConfigurationPtrInput {
	return (*linuxConfigurationPtrType)(v)
}

func (*linuxConfigurationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**LinuxConfiguration)(nil)).Elem()
}

func (i *linuxConfigurationPtrType) ToLinuxConfigurationPtrOutput() LinuxConfigurationPtrOutput {
	return i.ToLinuxConfigurationPtrOutputWithContext(context.Background())
}

func (i *linuxConfigurationPtrType) ToLinuxConfigurationPtrOutputWithContext(ctx context.Context) LinuxConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LinuxConfigurationPtrOutput)
}

type LinuxConfigurationOutput struct{ *pulumi.OutputState }

func (LinuxConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LinuxConfiguration)(nil)).Elem()
}

func (o LinuxConfigurationOutput) ToLinuxConfigurationOutput() LinuxConfigurationOutput {
	return o
}

func (o LinuxConfigurationOutput) ToLinuxConfigurationOutputWithContext(ctx context.Context) LinuxConfigurationOutput {
	return o
}

func (o LinuxConfigurationOutput) ToLinuxConfigurationPtrOutput() LinuxConfigurationPtrOutput {
	return o.ToLinuxConfigurationPtrOutputWithContext(context.Background())
}

func (o LinuxConfigurationOutput) ToLinuxConfigurationPtrOutputWithContext(ctx context.Context) LinuxConfigurationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v LinuxConfiguration) *LinuxConfiguration {
		return &v
	}).(LinuxConfigurationPtrOutput)
}

func (o LinuxConfigurationOutput) DisablePasswordAuthentication() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LinuxConfiguration) *bool { return v.DisablePasswordAuthentication }).(pulumi.BoolPtrOutput)
}

func (o LinuxConfigurationOutput) Ssh() SshConfigurationPtrOutput {
	return o.ApplyT(func(v LinuxConfiguration) *SshConfiguration { return v.Ssh }).(SshConfigurationPtrOutput)
}

type LinuxConfigurationPtrOutput struct{ *pulumi.OutputState }

func (LinuxConfigurationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LinuxConfiguration)(nil)).Elem()
}

func (o LinuxConfigurationPtrOutput) ToLinuxConfigurationPtrOutput() LinuxConfigurationPtrOutput {
	return o
}

func (o LinuxConfigurationPtrOutput) ToLinuxConfigurationPtrOutputWithContext(ctx context.Context) LinuxConfigurationPtrOutput {
	return o
}

func (o LinuxConfigurationPtrOutput) Elem() LinuxConfigurationOutput {
	return o.ApplyT(func(v *LinuxConfiguration) LinuxConfiguration {
		if v != nil {
			return *v
		}
		var ret LinuxConfiguration
		return ret
	}).(LinuxConfigurationOutput)
}

func (o LinuxConfigurationPtrOutput) DisablePasswordAuthentication() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *LinuxConfiguration) *bool {
		if v == nil {
			return nil
		}
		return v.DisablePasswordAuthentication
	}).(pulumi.BoolPtrOutput)
}

func (o LinuxConfigurationPtrOutput) Ssh() SshConfigurationPtrOutput {
	return o.ApplyT(func(v *LinuxConfiguration) *SshConfiguration {
		if v == nil {
			return nil
		}
		return v.Ssh
	}).(SshConfigurationPtrOutput)
}

type LinuxConfigurationResponse struct {
	DisablePasswordAuthentication *bool                     `pulumi:"disablePasswordAuthentication"`
	Ssh                           *SshConfigurationResponse `pulumi:"ssh"`
}

type LinuxConfigurationResponseOutput struct{ *pulumi.OutputState }

func (LinuxConfigurationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LinuxConfigurationResponse)(nil)).Elem()
}

func (o LinuxConfigurationResponseOutput) ToLinuxConfigurationResponseOutput() LinuxConfigurationResponseOutput {
	return o
}

func (o LinuxConfigurationResponseOutput) ToLinuxConfigurationResponseOutputWithContext(ctx context.Context) LinuxConfigurationResponseOutput {
	return o
}

func (o LinuxConfigurationResponseOutput) DisablePasswordAuthentication() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LinuxConfigurationResponse) *bool { return v.DisablePasswordAuthentication }).(pulumi.BoolPtrOutput)
}

func (o LinuxConfigurationResponseOutput) Ssh() SshConfigurationResponsePtrOutput {
	return o.ApplyT(func(v LinuxConfigurationResponse) *SshConfigurationResponse { return v.Ssh }).(SshConfigurationResponsePtrOutput)
}

type LinuxConfigurationResponsePtrOutput struct{ *pulumi.OutputState }

func (LinuxConfigurationResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LinuxConfigurationResponse)(nil)).Elem()
}

func (o LinuxConfigurationResponsePtrOutput) ToLinuxConfigurationResponsePtrOutput() LinuxConfigurationResponsePtrOutput {
	return o
}

func (o LinuxConfigurationResponsePtrOutput) ToLinuxConfigurationResponsePtrOutputWithContext(ctx context.Context) LinuxConfigurationResponsePtrOutput {
	return o
}

func (o LinuxConfigurationResponsePtrOutput) Elem() LinuxConfigurationResponseOutput {
	return o.ApplyT(func(v *LinuxConfigurationResponse) LinuxConfigurationResponse {
		if v != nil {
			return *v
		}
		var ret LinuxConfigurationResponse
		return ret
	}).(LinuxConfigurationResponseOutput)
}

func (o LinuxConfigurationResponsePtrOutput) DisablePasswordAuthentication() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *LinuxConfigurationResponse) *bool {
		if v == nil {
			return nil
		}
		return v.DisablePasswordAuthentication
	}).(pulumi.BoolPtrOutput)
}

func (o LinuxConfigurationResponsePtrOutput) Ssh() SshConfigurationResponsePtrOutput {
	return o.ApplyT(func(v *LinuxConfigurationResponse) *SshConfigurationResponse {
		if v == nil {
			return nil
		}
		return v.Ssh
	}).(SshConfigurationResponsePtrOutput)
}

type LogAnalyticsOutputResponse struct {
	Output string `pulumi:"output"`
}

type LogAnalyticsOutputResponseOutput struct{ *pulumi.OutputState }

func (LogAnalyticsOutputResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LogAnalyticsOutputResponse)(nil)).Elem()
}

func (o LogAnalyticsOutputResponseOutput) ToLogAnalyticsOutputResponseOutput() LogAnalyticsOutputResponseOutput {
	return o
}

func (o LogAnalyticsOutputResponseOutput) ToLogAnalyticsOutputResponseOutputWithContext(ctx context.Context) LogAnalyticsOutputResponseOutput {
	return o
}

func (o LogAnalyticsOutputResponseOutput) Output() pulumi.StringOutput {
	return o.ApplyT(func(v LogAnalyticsOutputResponse) string { return v.Output }).(pulumi.StringOutput)
}

type MaintenanceRedeployStatusResponse struct {
	IsCustomerInitiatedMaintenanceAllowed *bool   `pulumi:"isCustomerInitiatedMaintenanceAllowed"`
	LastOperationMessage                  *string `pulumi:"lastOperationMessage"`
	LastOperationResultCode               *string `pulumi:"lastOperationResultCode"`
	MaintenanceWindowEndTime              *string `pulumi:"maintenanceWindowEndTime"`
	MaintenanceWindowStartTime            *string `pulumi:"maintenanceWindowStartTime"`
	PreMaintenanceWindowEndTime           *string `pulumi:"preMaintenanceWindowEndTime"`
	PreMaintenanceWindowStartTime         *string `pulumi:"preMaintenanceWindowStartTime"`
}

type MaintenanceRedeployStatusResponseOutput struct{ *pulumi.OutputState }

func (MaintenanceRedeployStatusResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MaintenanceRedeployStatusResponse)(nil)).Elem()
}

func (o MaintenanceRedeployStatusResponseOutput) ToMaintenanceRedeployStatusResponseOutput() MaintenanceRedeployStatusResponseOutput {
	return o
}

func (o MaintenanceRedeployStatusResponseOutput) ToMaintenanceRedeployStatusResponseOutputWithContext(ctx context.Context) MaintenanceRedeployStatusResponseOutput {
	return o
}

func (o MaintenanceRedeployStatusResponseOutput) IsCustomerInitiatedMaintenanceAllowed() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v MaintenanceRedeployStatusResponse) *bool { return v.IsCustomerInitiatedMaintenanceAllowed }).(pulumi.BoolPtrOutput)
}

func (o MaintenanceRedeployStatusResponseOutput) LastOperationMessage() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MaintenanceRedeployStatusResponse) *string { return v.LastOperationMessage }).(pulumi.StringPtrOutput)
}

func (o MaintenanceRedeployStatusResponseOutput) LastOperationResultCode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MaintenanceRedeployStatusResponse) *string { return v.LastOperationResultCode }).(pulumi.StringPtrOutput)
}

func (o MaintenanceRedeployStatusResponseOutput) MaintenanceWindowEndTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MaintenanceRedeployStatusResponse) *string { return v.MaintenanceWindowEndTime }).(pulumi.StringPtrOutput)
}

func (o MaintenanceRedeployStatusResponseOutput) MaintenanceWindowStartTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MaintenanceRedeployStatusResponse) *string { return v.MaintenanceWindowStartTime }).(pulumi.StringPtrOutput)
}

func (o MaintenanceRedeployStatusResponseOutput) PreMaintenanceWindowEndTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MaintenanceRedeployStatusResponse) *string { return v.PreMaintenanceWindowEndTime }).(pulumi.StringPtrOutput)
}

func (o MaintenanceRedeployStatusResponseOutput) PreMaintenanceWindowStartTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MaintenanceRedeployStatusResponse) *string { return v.PreMaintenanceWindowStartTime }).(pulumi.StringPtrOutput)
}

type MaintenanceRedeployStatusResponsePtrOutput struct{ *pulumi.OutputState }

func (MaintenanceRedeployStatusResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MaintenanceRedeployStatusResponse)(nil)).Elem()
}

func (o MaintenanceRedeployStatusResponsePtrOutput) ToMaintenanceRedeployStatusResponsePtrOutput() MaintenanceRedeployStatusResponsePtrOutput {
	return o
}

func (o MaintenanceRedeployStatusResponsePtrOutput) ToMaintenanceRedeployStatusResponsePtrOutputWithContext(ctx context.Context) MaintenanceRedeployStatusResponsePtrOutput {
	return o
}

func (o MaintenanceRedeployStatusResponsePtrOutput) Elem() MaintenanceRedeployStatusResponseOutput {
	return o.ApplyT(func(v *MaintenanceRedeployStatusResponse) MaintenanceRedeployStatusResponse {
		if v != nil {
			return *v
		}
		var ret MaintenanceRedeployStatusResponse
		return ret
	}).(MaintenanceRedeployStatusResponseOutput)
}

func (o MaintenanceRedeployStatusResponsePtrOutput) IsCustomerInitiatedMaintenanceAllowed() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *MaintenanceRedeployStatusResponse) *bool {
		if v == nil {
			return nil
		}
		return v.IsCustomerInitiatedMaintenanceAllowed
	}).(pulumi.BoolPtrOutput)
}

func (o MaintenanceRedeployStatusResponsePtrOutput) LastOperationMessage() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MaintenanceRedeployStatusResponse) *string {
		if v == nil {
			return nil
		}
		return v.LastOperationMessage
	}).(pulumi.StringPtrOutput)
}

func (o MaintenanceRedeployStatusResponsePtrOutput) LastOperationResultCode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MaintenanceRedeployStatusResponse) *string {
		if v == nil {
			return nil
		}
		return v.LastOperationResultCode
	}).(pulumi.StringPtrOutput)
}

func (o MaintenanceRedeployStatusResponsePtrOutput) MaintenanceWindowEndTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MaintenanceRedeployStatusResponse) *string {
		if v == nil {
			return nil
		}
		return v.MaintenanceWindowEndTime
	}).(pulumi.StringPtrOutput)
}

func (o MaintenanceRedeployStatusResponsePtrOutput) MaintenanceWindowStartTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MaintenanceRedeployStatusResponse) *string {
		if v == nil {
			return nil
		}
		return v.MaintenanceWindowStartTime
	}).(pulumi.StringPtrOutput)
}

func (o MaintenanceRedeployStatusResponsePtrOutput) PreMaintenanceWindowEndTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MaintenanceRedeployStatusResponse) *string {
		if v == nil {
			return nil
		}
		return v.PreMaintenanceWindowEndTime
	}).(pulumi.StringPtrOutput)
}

func (o MaintenanceRedeployStatusResponsePtrOutput) PreMaintenanceWindowStartTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MaintenanceRedeployStatusResponse) *string {
		if v == nil {
			return nil
		}
		return v.PreMaintenanceWindowStartTime
	}).(pulumi.StringPtrOutput)
}

type ManagedDiskParameters struct {
	Id                 *string `pulumi:"id"`
	StorageAccountType *string `pulumi:"storageAccountType"`
}





type ManagedDiskParametersInput interface {
	pulumi.Input

	ToManagedDiskParametersOutput() ManagedDiskParametersOutput
	ToManagedDiskParametersOutputWithContext(context.Context) ManagedDiskParametersOutput
}

type ManagedDiskParametersArgs struct {
	Id                 pulumi.StringPtrInput `pulumi:"id"`
	StorageAccountType pulumi.StringPtrInput `pulumi:"storageAccountType"`
}

func (ManagedDiskParametersArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedDiskParameters)(nil)).Elem()
}

func (i ManagedDiskParametersArgs) ToManagedDiskParametersOutput() ManagedDiskParametersOutput {
	return i.ToManagedDiskParametersOutputWithContext(context.Background())
}

func (i ManagedDiskParametersArgs) ToManagedDiskParametersOutputWithContext(ctx context.Context) ManagedDiskParametersOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedDiskParametersOutput)
}

func (i ManagedDiskParametersArgs) ToManagedDiskParametersPtrOutput() ManagedDiskParametersPtrOutput {
	return i.ToManagedDiskParametersPtrOutputWithContext(context.Background())
}

func (i ManagedDiskParametersArgs) ToManagedDiskParametersPtrOutputWithContext(ctx context.Context) ManagedDiskParametersPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedDiskParametersOutput).ToManagedDiskParametersPtrOutputWithContext(ctx)
}









type ManagedDiskParametersPtrInput interface {
	pulumi.Input

	ToManagedDiskParametersPtrOutput() ManagedDiskParametersPtrOutput
	ToManagedDiskParametersPtrOutputWithContext(context.Context) ManagedDiskParametersPtrOutput
}

type managedDiskParametersPtrType ManagedDiskParametersArgs

func ManagedDiskParametersPtr(v *ManagedDiskParametersArgs) ManagedDiskParametersPtrInput {
	return (*managedDiskParametersPtrType)(v)
}

func (*managedDiskParametersPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedDiskParameters)(nil)).Elem()
}

func (i *managedDiskParametersPtrType) ToManagedDiskParametersPtrOutput() ManagedDiskParametersPtrOutput {
	return i.ToManagedDiskParametersPtrOutputWithContext(context.Background())
}

func (i *managedDiskParametersPtrType) ToManagedDiskParametersPtrOutputWithContext(ctx context.Context) ManagedDiskParametersPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedDiskParametersPtrOutput)
}

type ManagedDiskParametersOutput struct{ *pulumi.OutputState }

func (ManagedDiskParametersOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedDiskParameters)(nil)).Elem()
}

func (o ManagedDiskParametersOutput) ToManagedDiskParametersOutput() ManagedDiskParametersOutput {
	return o
}

func (o ManagedDiskParametersOutput) ToManagedDiskParametersOutputWithContext(ctx context.Context) ManagedDiskParametersOutput {
	return o
}

func (o ManagedDiskParametersOutput) ToManagedDiskParametersPtrOutput() ManagedDiskParametersPtrOutput {
	return o.ToManagedDiskParametersPtrOutputWithContext(context.Background())
}

func (o ManagedDiskParametersOutput) ToManagedDiskParametersPtrOutputWithContext(ctx context.Context) ManagedDiskParametersPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ManagedDiskParameters) *ManagedDiskParameters {
		return &v
	}).(ManagedDiskParametersPtrOutput)
}

func (o ManagedDiskParametersOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagedDiskParameters) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o ManagedDiskParametersOutput) StorageAccountType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagedDiskParameters) *string { return v.StorageAccountType }).(pulumi.StringPtrOutput)
}

type ManagedDiskParametersPtrOutput struct{ *pulumi.OutputState }

func (ManagedDiskParametersPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedDiskParameters)(nil)).Elem()
}

func (o ManagedDiskParametersPtrOutput) ToManagedDiskParametersPtrOutput() ManagedDiskParametersPtrOutput {
	return o
}

func (o ManagedDiskParametersPtrOutput) ToManagedDiskParametersPtrOutputWithContext(ctx context.Context) ManagedDiskParametersPtrOutput {
	return o
}

func (o ManagedDiskParametersPtrOutput) Elem() ManagedDiskParametersOutput {
	return o.ApplyT(func(v *ManagedDiskParameters) ManagedDiskParameters {
		if v != nil {
			return *v
		}
		var ret ManagedDiskParameters
		return ret
	}).(ManagedDiskParametersOutput)
}

func (o ManagedDiskParametersPtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedDiskParameters) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

func (o ManagedDiskParametersPtrOutput) StorageAccountType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedDiskParameters) *string {
		if v == nil {
			return nil
		}
		return v.StorageAccountType
	}).(pulumi.StringPtrOutput)
}

type ManagedDiskParametersResponse struct {
	Id                 *string `pulumi:"id"`
	StorageAccountType *string `pulumi:"storageAccountType"`
}

type ManagedDiskParametersResponseOutput struct{ *pulumi.OutputState }

func (ManagedDiskParametersResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedDiskParametersResponse)(nil)).Elem()
}

func (o ManagedDiskParametersResponseOutput) ToManagedDiskParametersResponseOutput() ManagedDiskParametersResponseOutput {
	return o
}

func (o ManagedDiskParametersResponseOutput) ToManagedDiskParametersResponseOutputWithContext(ctx context.Context) ManagedDiskParametersResponseOutput {
	return o
}

func (o ManagedDiskParametersResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagedDiskParametersResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o ManagedDiskParametersResponseOutput) StorageAccountType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagedDiskParametersResponse) *string { return v.StorageAccountType }).(pulumi.StringPtrOutput)
}

type ManagedDiskParametersResponsePtrOutput struct{ *pulumi.OutputState }

func (ManagedDiskParametersResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedDiskParametersResponse)(nil)).Elem()
}

func (o ManagedDiskParametersResponsePtrOutput) ToManagedDiskParametersResponsePtrOutput() ManagedDiskParametersResponsePtrOutput {
	return o
}

func (o ManagedDiskParametersResponsePtrOutput) ToManagedDiskParametersResponsePtrOutputWithContext(ctx context.Context) ManagedDiskParametersResponsePtrOutput {
	return o
}

func (o ManagedDiskParametersResponsePtrOutput) Elem() ManagedDiskParametersResponseOutput {
	return o.ApplyT(func(v *ManagedDiskParametersResponse) ManagedDiskParametersResponse {
		if v != nil {
			return *v
		}
		var ret ManagedDiskParametersResponse
		return ret
	}).(ManagedDiskParametersResponseOutput)
}

func (o ManagedDiskParametersResponsePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedDiskParametersResponse) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

func (o ManagedDiskParametersResponsePtrOutput) StorageAccountType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedDiskParametersResponse) *string {
		if v == nil {
			return nil
		}
		return v.StorageAccountType
	}).(pulumi.StringPtrOutput)
}

type NetworkInterfaceReference struct {
	Id      *string `pulumi:"id"`
	Primary *bool   `pulumi:"primary"`
}





type NetworkInterfaceReferenceInput interface {
	pulumi.Input

	ToNetworkInterfaceReferenceOutput() NetworkInterfaceReferenceOutput
	ToNetworkInterfaceReferenceOutputWithContext(context.Context) NetworkInterfaceReferenceOutput
}

type NetworkInterfaceReferenceArgs struct {
	Id      pulumi.StringPtrInput `pulumi:"id"`
	Primary pulumi.BoolPtrInput   `pulumi:"primary"`
}

func (NetworkInterfaceReferenceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkInterfaceReference)(nil)).Elem()
}

func (i NetworkInterfaceReferenceArgs) ToNetworkInterfaceReferenceOutput() NetworkInterfaceReferenceOutput {
	return i.ToNetworkInterfaceReferenceOutputWithContext(context.Background())
}

func (i NetworkInterfaceReferenceArgs) ToNetworkInterfaceReferenceOutputWithContext(ctx context.Context) NetworkInterfaceReferenceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkInterfaceReferenceOutput)
}





type NetworkInterfaceReferenceArrayInput interface {
	pulumi.Input

	ToNetworkInterfaceReferenceArrayOutput() NetworkInterfaceReferenceArrayOutput
	ToNetworkInterfaceReferenceArrayOutputWithContext(context.Context) NetworkInterfaceReferenceArrayOutput
}

type NetworkInterfaceReferenceArray []NetworkInterfaceReferenceInput

func (NetworkInterfaceReferenceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]NetworkInterfaceReference)(nil)).Elem()
}

func (i NetworkInterfaceReferenceArray) ToNetworkInterfaceReferenceArrayOutput() NetworkInterfaceReferenceArrayOutput {
	return i.ToNetworkInterfaceReferenceArrayOutputWithContext(context.Background())
}

func (i NetworkInterfaceReferenceArray) ToNetworkInterfaceReferenceArrayOutputWithContext(ctx context.Context) NetworkInterfaceReferenceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkInterfaceReferenceArrayOutput)
}

type NetworkInterfaceReferenceOutput struct{ *pulumi.OutputState }

func (NetworkInterfaceReferenceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkInterfaceReference)(nil)).Elem()
}

func (o NetworkInterfaceReferenceOutput) ToNetworkInterfaceReferenceOutput() NetworkInterfaceReferenceOutput {
	return o
}

func (o NetworkInterfaceReferenceOutput) ToNetworkInterfaceReferenceOutputWithContext(ctx context.Context) NetworkInterfaceReferenceOutput {
	return o
}

func (o NetworkInterfaceReferenceOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkInterfaceReference) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o NetworkInterfaceReferenceOutput) Primary() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v NetworkInterfaceReference) *bool { return v.Primary }).(pulumi.BoolPtrOutput)
}

type NetworkInterfaceReferenceArrayOutput struct{ *pulumi.OutputState }

func (NetworkInterfaceReferenceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]NetworkInterfaceReference)(nil)).Elem()
}

func (o NetworkInterfaceReferenceArrayOutput) ToNetworkInterfaceReferenceArrayOutput() NetworkInterfaceReferenceArrayOutput {
	return o
}

func (o NetworkInterfaceReferenceArrayOutput) ToNetworkInterfaceReferenceArrayOutputWithContext(ctx context.Context) NetworkInterfaceReferenceArrayOutput {
	return o
}

func (o NetworkInterfaceReferenceArrayOutput) Index(i pulumi.IntInput) NetworkInterfaceReferenceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) NetworkInterfaceReference {
		return vs[0].([]NetworkInterfaceReference)[vs[1].(int)]
	}).(NetworkInterfaceReferenceOutput)
}

type NetworkInterfaceReferenceResponse struct {
	Id      *string `pulumi:"id"`
	Primary *bool   `pulumi:"primary"`
}

type NetworkInterfaceReferenceResponseOutput struct{ *pulumi.OutputState }

func (NetworkInterfaceReferenceResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkInterfaceReferenceResponse)(nil)).Elem()
}

func (o NetworkInterfaceReferenceResponseOutput) ToNetworkInterfaceReferenceResponseOutput() NetworkInterfaceReferenceResponseOutput {
	return o
}

func (o NetworkInterfaceReferenceResponseOutput) ToNetworkInterfaceReferenceResponseOutputWithContext(ctx context.Context) NetworkInterfaceReferenceResponseOutput {
	return o
}

func (o NetworkInterfaceReferenceResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkInterfaceReferenceResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o NetworkInterfaceReferenceResponseOutput) Primary() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v NetworkInterfaceReferenceResponse) *bool { return v.Primary }).(pulumi.BoolPtrOutput)
}

type NetworkInterfaceReferenceResponseArrayOutput struct{ *pulumi.OutputState }

func (NetworkInterfaceReferenceResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]NetworkInterfaceReferenceResponse)(nil)).Elem()
}

func (o NetworkInterfaceReferenceResponseArrayOutput) ToNetworkInterfaceReferenceResponseArrayOutput() NetworkInterfaceReferenceResponseArrayOutput {
	return o
}

func (o NetworkInterfaceReferenceResponseArrayOutput) ToNetworkInterfaceReferenceResponseArrayOutputWithContext(ctx context.Context) NetworkInterfaceReferenceResponseArrayOutput {
	return o
}

func (o NetworkInterfaceReferenceResponseArrayOutput) Index(i pulumi.IntInput) NetworkInterfaceReferenceResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) NetworkInterfaceReferenceResponse {
		return vs[0].([]NetworkInterfaceReferenceResponse)[vs[1].(int)]
	}).(NetworkInterfaceReferenceResponseOutput)
}

type NetworkProfile struct {
	NetworkInterfaces []NetworkInterfaceReference `pulumi:"networkInterfaces"`
}





type NetworkProfileInput interface {
	pulumi.Input

	ToNetworkProfileOutput() NetworkProfileOutput
	ToNetworkProfileOutputWithContext(context.Context) NetworkProfileOutput
}

type NetworkProfileArgs struct {
	NetworkInterfaces NetworkInterfaceReferenceArrayInput `pulumi:"networkInterfaces"`
}

func (NetworkProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkProfile)(nil)).Elem()
}

func (i NetworkProfileArgs) ToNetworkProfileOutput() NetworkProfileOutput {
	return i.ToNetworkProfileOutputWithContext(context.Background())
}

func (i NetworkProfileArgs) ToNetworkProfileOutputWithContext(ctx context.Context) NetworkProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkProfileOutput)
}

func (i NetworkProfileArgs) ToNetworkProfilePtrOutput() NetworkProfilePtrOutput {
	return i.ToNetworkProfilePtrOutputWithContext(context.Background())
}

func (i NetworkProfileArgs) ToNetworkProfilePtrOutputWithContext(ctx context.Context) NetworkProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkProfileOutput).ToNetworkProfilePtrOutputWithContext(ctx)
}









type NetworkProfilePtrInput interface {
	pulumi.Input

	ToNetworkProfilePtrOutput() NetworkProfilePtrOutput
	ToNetworkProfilePtrOutputWithContext(context.Context) NetworkProfilePtrOutput
}

type networkProfilePtrType NetworkProfileArgs

func NetworkProfilePtr(v *NetworkProfileArgs) NetworkProfilePtrInput {
	return (*networkProfilePtrType)(v)
}

func (*networkProfilePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkProfile)(nil)).Elem()
}

func (i *networkProfilePtrType) ToNetworkProfilePtrOutput() NetworkProfilePtrOutput {
	return i.ToNetworkProfilePtrOutputWithContext(context.Background())
}

func (i *networkProfilePtrType) ToNetworkProfilePtrOutputWithContext(ctx context.Context) NetworkProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkProfilePtrOutput)
}

type NetworkProfileOutput struct{ *pulumi.OutputState }

func (NetworkProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkProfile)(nil)).Elem()
}

func (o NetworkProfileOutput) ToNetworkProfileOutput() NetworkProfileOutput {
	return o
}

func (o NetworkProfileOutput) ToNetworkProfileOutputWithContext(ctx context.Context) NetworkProfileOutput {
	return o
}

func (o NetworkProfileOutput) ToNetworkProfilePtrOutput() NetworkProfilePtrOutput {
	return o.ToNetworkProfilePtrOutputWithContext(context.Background())
}

func (o NetworkProfileOutput) ToNetworkProfilePtrOutputWithContext(ctx context.Context) NetworkProfilePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v NetworkProfile) *NetworkProfile {
		return &v
	}).(NetworkProfilePtrOutput)
}

func (o NetworkProfileOutput) NetworkInterfaces() NetworkInterfaceReferenceArrayOutput {
	return o.ApplyT(func(v NetworkProfile) []NetworkInterfaceReference { return v.NetworkInterfaces }).(NetworkInterfaceReferenceArrayOutput)
}

type NetworkProfilePtrOutput struct{ *pulumi.OutputState }

func (NetworkProfilePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkProfile)(nil)).Elem()
}

func (o NetworkProfilePtrOutput) ToNetworkProfilePtrOutput() NetworkProfilePtrOutput {
	return o
}

func (o NetworkProfilePtrOutput) ToNetworkProfilePtrOutputWithContext(ctx context.Context) NetworkProfilePtrOutput {
	return o
}

func (o NetworkProfilePtrOutput) Elem() NetworkProfileOutput {
	return o.ApplyT(func(v *NetworkProfile) NetworkProfile {
		if v != nil {
			return *v
		}
		var ret NetworkProfile
		return ret
	}).(NetworkProfileOutput)
}

func (o NetworkProfilePtrOutput) NetworkInterfaces() NetworkInterfaceReferenceArrayOutput {
	return o.ApplyT(func(v *NetworkProfile) []NetworkInterfaceReference {
		if v == nil {
			return nil
		}
		return v.NetworkInterfaces
	}).(NetworkInterfaceReferenceArrayOutput)
}

type NetworkProfileResponse struct {
	NetworkInterfaces []NetworkInterfaceReferenceResponse `pulumi:"networkInterfaces"`
}

type NetworkProfileResponseOutput struct{ *pulumi.OutputState }

func (NetworkProfileResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkProfileResponse)(nil)).Elem()
}

func (o NetworkProfileResponseOutput) ToNetworkProfileResponseOutput() NetworkProfileResponseOutput {
	return o
}

func (o NetworkProfileResponseOutput) ToNetworkProfileResponseOutputWithContext(ctx context.Context) NetworkProfileResponseOutput {
	return o
}

func (o NetworkProfileResponseOutput) NetworkInterfaces() NetworkInterfaceReferenceResponseArrayOutput {
	return o.ApplyT(func(v NetworkProfileResponse) []NetworkInterfaceReferenceResponse { return v.NetworkInterfaces }).(NetworkInterfaceReferenceResponseArrayOutput)
}

type NetworkProfileResponsePtrOutput struct{ *pulumi.OutputState }

func (NetworkProfileResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkProfileResponse)(nil)).Elem()
}

func (o NetworkProfileResponsePtrOutput) ToNetworkProfileResponsePtrOutput() NetworkProfileResponsePtrOutput {
	return o
}

func (o NetworkProfileResponsePtrOutput) ToNetworkProfileResponsePtrOutputWithContext(ctx context.Context) NetworkProfileResponsePtrOutput {
	return o
}

func (o NetworkProfileResponsePtrOutput) Elem() NetworkProfileResponseOutput {
	return o.ApplyT(func(v *NetworkProfileResponse) NetworkProfileResponse {
		if v != nil {
			return *v
		}
		var ret NetworkProfileResponse
		return ret
	}).(NetworkProfileResponseOutput)
}

func (o NetworkProfileResponsePtrOutput) NetworkInterfaces() NetworkInterfaceReferenceResponseArrayOutput {
	return o.ApplyT(func(v *NetworkProfileResponse) []NetworkInterfaceReferenceResponse {
		if v == nil {
			return nil
		}
		return v.NetworkInterfaces
	}).(NetworkInterfaceReferenceResponseArrayOutput)
}

type OSDisk struct {
	Caching                 *CachingTypes           `pulumi:"caching"`
	CreateOption            string                  `pulumi:"createOption"`
	DiskSizeGB              *int                    `pulumi:"diskSizeGB"`
	EncryptionSettings      *DiskEncryptionSettings `pulumi:"encryptionSettings"`
	Image                   *VirtualHardDisk        `pulumi:"image"`
	ManagedDisk             *ManagedDiskParameters  `pulumi:"managedDisk"`
	Name                    *string                 `pulumi:"name"`
	OsType                  *OperatingSystemTypes   `pulumi:"osType"`
	Vhd                     *VirtualHardDisk        `pulumi:"vhd"`
	WriteAcceleratorEnabled *bool                   `pulumi:"writeAcceleratorEnabled"`
}





type OSDiskInput interface {
	pulumi.Input

	ToOSDiskOutput() OSDiskOutput
	ToOSDiskOutputWithContext(context.Context) OSDiskOutput
}

type OSDiskArgs struct {
	Caching                 CachingTypesPtrInput           `pulumi:"caching"`
	CreateOption            pulumi.StringInput             `pulumi:"createOption"`
	DiskSizeGB              pulumi.IntPtrInput             `pulumi:"diskSizeGB"`
	EncryptionSettings      DiskEncryptionSettingsPtrInput `pulumi:"encryptionSettings"`
	Image                   VirtualHardDiskPtrInput        `pulumi:"image"`
	ManagedDisk             ManagedDiskParametersPtrInput  `pulumi:"managedDisk"`
	Name                    pulumi.StringPtrInput          `pulumi:"name"`
	OsType                  OperatingSystemTypesPtrInput   `pulumi:"osType"`
	Vhd                     VirtualHardDiskPtrInput        `pulumi:"vhd"`
	WriteAcceleratorEnabled pulumi.BoolPtrInput            `pulumi:"writeAcceleratorEnabled"`
}

func (OSDiskArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*OSDisk)(nil)).Elem()
}

func (i OSDiskArgs) ToOSDiskOutput() OSDiskOutput {
	return i.ToOSDiskOutputWithContext(context.Background())
}

func (i OSDiskArgs) ToOSDiskOutputWithContext(ctx context.Context) OSDiskOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OSDiskOutput)
}

func (i OSDiskArgs) ToOSDiskPtrOutput() OSDiskPtrOutput {
	return i.ToOSDiskPtrOutputWithContext(context.Background())
}

func (i OSDiskArgs) ToOSDiskPtrOutputWithContext(ctx context.Context) OSDiskPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OSDiskOutput).ToOSDiskPtrOutputWithContext(ctx)
}









type OSDiskPtrInput interface {
	pulumi.Input

	ToOSDiskPtrOutput() OSDiskPtrOutput
	ToOSDiskPtrOutputWithContext(context.Context) OSDiskPtrOutput
}

type osdiskPtrType OSDiskArgs

func OSDiskPtr(v *OSDiskArgs) OSDiskPtrInput {
	return (*osdiskPtrType)(v)
}

func (*osdiskPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**OSDisk)(nil)).Elem()
}

func (i *osdiskPtrType) ToOSDiskPtrOutput() OSDiskPtrOutput {
	return i.ToOSDiskPtrOutputWithContext(context.Background())
}

func (i *osdiskPtrType) ToOSDiskPtrOutputWithContext(ctx context.Context) OSDiskPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OSDiskPtrOutput)
}

type OSDiskOutput struct{ *pulumi.OutputState }

func (OSDiskOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*OSDisk)(nil)).Elem()
}

func (o OSDiskOutput) ToOSDiskOutput() OSDiskOutput {
	return o
}

func (o OSDiskOutput) ToOSDiskOutputWithContext(ctx context.Context) OSDiskOutput {
	return o
}

func (o OSDiskOutput) ToOSDiskPtrOutput() OSDiskPtrOutput {
	return o.ToOSDiskPtrOutputWithContext(context.Background())
}

func (o OSDiskOutput) ToOSDiskPtrOutputWithContext(ctx context.Context) OSDiskPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v OSDisk) *OSDisk {
		return &v
	}).(OSDiskPtrOutput)
}

func (o OSDiskOutput) Caching() CachingTypesPtrOutput {
	return o.ApplyT(func(v OSDisk) *CachingTypes { return v.Caching }).(CachingTypesPtrOutput)
}

func (o OSDiskOutput) CreateOption() pulumi.StringOutput {
	return o.ApplyT(func(v OSDisk) string { return v.CreateOption }).(pulumi.StringOutput)
}

func (o OSDiskOutput) DiskSizeGB() pulumi.IntPtrOutput {
	return o.ApplyT(func(v OSDisk) *int { return v.DiskSizeGB }).(pulumi.IntPtrOutput)
}

func (o OSDiskOutput) EncryptionSettings() DiskEncryptionSettingsPtrOutput {
	return o.ApplyT(func(v OSDisk) *DiskEncryptionSettings { return v.EncryptionSettings }).(DiskEncryptionSettingsPtrOutput)
}

func (o OSDiskOutput) Image() VirtualHardDiskPtrOutput {
	return o.ApplyT(func(v OSDisk) *VirtualHardDisk { return v.Image }).(VirtualHardDiskPtrOutput)
}

func (o OSDiskOutput) ManagedDisk() ManagedDiskParametersPtrOutput {
	return o.ApplyT(func(v OSDisk) *ManagedDiskParameters { return v.ManagedDisk }).(ManagedDiskParametersPtrOutput)
}

func (o OSDiskOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v OSDisk) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o OSDiskOutput) OsType() OperatingSystemTypesPtrOutput {
	return o.ApplyT(func(v OSDisk) *OperatingSystemTypes { return v.OsType }).(OperatingSystemTypesPtrOutput)
}

func (o OSDiskOutput) Vhd() VirtualHardDiskPtrOutput {
	return o.ApplyT(func(v OSDisk) *VirtualHardDisk { return v.Vhd }).(VirtualHardDiskPtrOutput)
}

func (o OSDiskOutput) WriteAcceleratorEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v OSDisk) *bool { return v.WriteAcceleratorEnabled }).(pulumi.BoolPtrOutput)
}

type OSDiskPtrOutput struct{ *pulumi.OutputState }

func (OSDiskPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OSDisk)(nil)).Elem()
}

func (o OSDiskPtrOutput) ToOSDiskPtrOutput() OSDiskPtrOutput {
	return o
}

func (o OSDiskPtrOutput) ToOSDiskPtrOutputWithContext(ctx context.Context) OSDiskPtrOutput {
	return o
}

func (o OSDiskPtrOutput) Elem() OSDiskOutput {
	return o.ApplyT(func(v *OSDisk) OSDisk {
		if v != nil {
			return *v
		}
		var ret OSDisk
		return ret
	}).(OSDiskOutput)
}

func (o OSDiskPtrOutput) Caching() CachingTypesPtrOutput {
	return o.ApplyT(func(v *OSDisk) *CachingTypes {
		if v == nil {
			return nil
		}
		return v.Caching
	}).(CachingTypesPtrOutput)
}

func (o OSDiskPtrOutput) CreateOption() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OSDisk) *string {
		if v == nil {
			return nil
		}
		return &v.CreateOption
	}).(pulumi.StringPtrOutput)
}

func (o OSDiskPtrOutput) DiskSizeGB() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *OSDisk) *int {
		if v == nil {
			return nil
		}
		return v.DiskSizeGB
	}).(pulumi.IntPtrOutput)
}

func (o OSDiskPtrOutput) EncryptionSettings() DiskEncryptionSettingsPtrOutput {
	return o.ApplyT(func(v *OSDisk) *DiskEncryptionSettings {
		if v == nil {
			return nil
		}
		return v.EncryptionSettings
	}).(DiskEncryptionSettingsPtrOutput)
}

func (o OSDiskPtrOutput) Image() VirtualHardDiskPtrOutput {
	return o.ApplyT(func(v *OSDisk) *VirtualHardDisk {
		if v == nil {
			return nil
		}
		return v.Image
	}).(VirtualHardDiskPtrOutput)
}

func (o OSDiskPtrOutput) ManagedDisk() ManagedDiskParametersPtrOutput {
	return o.ApplyT(func(v *OSDisk) *ManagedDiskParameters {
		if v == nil {
			return nil
		}
		return v.ManagedDisk
	}).(ManagedDiskParametersPtrOutput)
}

func (o OSDiskPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OSDisk) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o OSDiskPtrOutput) OsType() OperatingSystemTypesPtrOutput {
	return o.ApplyT(func(v *OSDisk) *OperatingSystemTypes {
		if v == nil {
			return nil
		}
		return v.OsType
	}).(OperatingSystemTypesPtrOutput)
}

func (o OSDiskPtrOutput) Vhd() VirtualHardDiskPtrOutput {
	return o.ApplyT(func(v *OSDisk) *VirtualHardDisk {
		if v == nil {
			return nil
		}
		return v.Vhd
	}).(VirtualHardDiskPtrOutput)
}

func (o OSDiskPtrOutput) WriteAcceleratorEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *OSDisk) *bool {
		if v == nil {
			return nil
		}
		return v.WriteAcceleratorEnabled
	}).(pulumi.BoolPtrOutput)
}

type OSDiskResponse struct {
	Caching                 *string                         `pulumi:"caching"`
	CreateOption            string                          `pulumi:"createOption"`
	DiskSizeGB              *int                            `pulumi:"diskSizeGB"`
	EncryptionSettings      *DiskEncryptionSettingsResponse `pulumi:"encryptionSettings"`
	Image                   *VirtualHardDiskResponse        `pulumi:"image"`
	ManagedDisk             *ManagedDiskParametersResponse  `pulumi:"managedDisk"`
	Name                    *string                         `pulumi:"name"`
	OsType                  *string                         `pulumi:"osType"`
	Vhd                     *VirtualHardDiskResponse        `pulumi:"vhd"`
	WriteAcceleratorEnabled *bool                           `pulumi:"writeAcceleratorEnabled"`
}

type OSDiskResponseOutput struct{ *pulumi.OutputState }

func (OSDiskResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*OSDiskResponse)(nil)).Elem()
}

func (o OSDiskResponseOutput) ToOSDiskResponseOutput() OSDiskResponseOutput {
	return o
}

func (o OSDiskResponseOutput) ToOSDiskResponseOutputWithContext(ctx context.Context) OSDiskResponseOutput {
	return o
}

func (o OSDiskResponseOutput) Caching() pulumi.StringPtrOutput {
	return o.ApplyT(func(v OSDiskResponse) *string { return v.Caching }).(pulumi.StringPtrOutput)
}

func (o OSDiskResponseOutput) CreateOption() pulumi.StringOutput {
	return o.ApplyT(func(v OSDiskResponse) string { return v.CreateOption }).(pulumi.StringOutput)
}

func (o OSDiskResponseOutput) DiskSizeGB() pulumi.IntPtrOutput {
	return o.ApplyT(func(v OSDiskResponse) *int { return v.DiskSizeGB }).(pulumi.IntPtrOutput)
}

func (o OSDiskResponseOutput) EncryptionSettings() DiskEncryptionSettingsResponsePtrOutput {
	return o.ApplyT(func(v OSDiskResponse) *DiskEncryptionSettingsResponse { return v.EncryptionSettings }).(DiskEncryptionSettingsResponsePtrOutput)
}

func (o OSDiskResponseOutput) Image() VirtualHardDiskResponsePtrOutput {
	return o.ApplyT(func(v OSDiskResponse) *VirtualHardDiskResponse { return v.Image }).(VirtualHardDiskResponsePtrOutput)
}

func (o OSDiskResponseOutput) ManagedDisk() ManagedDiskParametersResponsePtrOutput {
	return o.ApplyT(func(v OSDiskResponse) *ManagedDiskParametersResponse { return v.ManagedDisk }).(ManagedDiskParametersResponsePtrOutput)
}

func (o OSDiskResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v OSDiskResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o OSDiskResponseOutput) OsType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v OSDiskResponse) *string { return v.OsType }).(pulumi.StringPtrOutput)
}

func (o OSDiskResponseOutput) Vhd() VirtualHardDiskResponsePtrOutput {
	return o.ApplyT(func(v OSDiskResponse) *VirtualHardDiskResponse { return v.Vhd }).(VirtualHardDiskResponsePtrOutput)
}

func (o OSDiskResponseOutput) WriteAcceleratorEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v OSDiskResponse) *bool { return v.WriteAcceleratorEnabled }).(pulumi.BoolPtrOutput)
}

type OSDiskResponsePtrOutput struct{ *pulumi.OutputState }

func (OSDiskResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OSDiskResponse)(nil)).Elem()
}

func (o OSDiskResponsePtrOutput) ToOSDiskResponsePtrOutput() OSDiskResponsePtrOutput {
	return o
}

func (o OSDiskResponsePtrOutput) ToOSDiskResponsePtrOutputWithContext(ctx context.Context) OSDiskResponsePtrOutput {
	return o
}

func (o OSDiskResponsePtrOutput) Elem() OSDiskResponseOutput {
	return o.ApplyT(func(v *OSDiskResponse) OSDiskResponse {
		if v != nil {
			return *v
		}
		var ret OSDiskResponse
		return ret
	}).(OSDiskResponseOutput)
}

func (o OSDiskResponsePtrOutput) Caching() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OSDiskResponse) *string {
		if v == nil {
			return nil
		}
		return v.Caching
	}).(pulumi.StringPtrOutput)
}

func (o OSDiskResponsePtrOutput) CreateOption() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OSDiskResponse) *string {
		if v == nil {
			return nil
		}
		return &v.CreateOption
	}).(pulumi.StringPtrOutput)
}

func (o OSDiskResponsePtrOutput) DiskSizeGB() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *OSDiskResponse) *int {
		if v == nil {
			return nil
		}
		return v.DiskSizeGB
	}).(pulumi.IntPtrOutput)
}

func (o OSDiskResponsePtrOutput) EncryptionSettings() DiskEncryptionSettingsResponsePtrOutput {
	return o.ApplyT(func(v *OSDiskResponse) *DiskEncryptionSettingsResponse {
		if v == nil {
			return nil
		}
		return v.EncryptionSettings
	}).(DiskEncryptionSettingsResponsePtrOutput)
}

func (o OSDiskResponsePtrOutput) Image() VirtualHardDiskResponsePtrOutput {
	return o.ApplyT(func(v *OSDiskResponse) *VirtualHardDiskResponse {
		if v == nil {
			return nil
		}
		return v.Image
	}).(VirtualHardDiskResponsePtrOutput)
}

func (o OSDiskResponsePtrOutput) ManagedDisk() ManagedDiskParametersResponsePtrOutput {
	return o.ApplyT(func(v *OSDiskResponse) *ManagedDiskParametersResponse {
		if v == nil {
			return nil
		}
		return v.ManagedDisk
	}).(ManagedDiskParametersResponsePtrOutput)
}

func (o OSDiskResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OSDiskResponse) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o OSDiskResponsePtrOutput) OsType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OSDiskResponse) *string {
		if v == nil {
			return nil
		}
		return v.OsType
	}).(pulumi.StringPtrOutput)
}

func (o OSDiskResponsePtrOutput) Vhd() VirtualHardDiskResponsePtrOutput {
	return o.ApplyT(func(v *OSDiskResponse) *VirtualHardDiskResponse {
		if v == nil {
			return nil
		}
		return v.Vhd
	}).(VirtualHardDiskResponsePtrOutput)
}

func (o OSDiskResponsePtrOutput) WriteAcceleratorEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *OSDiskResponse) *bool {
		if v == nil {
			return nil
		}
		return v.WriteAcceleratorEnabled
	}).(pulumi.BoolPtrOutput)
}

type OSProfile struct {
	AdminPassword        *string               `pulumi:"adminPassword"`
	AdminUsername        *string               `pulumi:"adminUsername"`
	ComputerName         *string               `pulumi:"computerName"`
	CustomData           *string               `pulumi:"customData"`
	LinuxConfiguration   *LinuxConfiguration   `pulumi:"linuxConfiguration"`
	Secrets              []VaultSecretGroup    `pulumi:"secrets"`
	WindowsConfiguration *WindowsConfiguration `pulumi:"windowsConfiguration"`
}





type OSProfileInput interface {
	pulumi.Input

	ToOSProfileOutput() OSProfileOutput
	ToOSProfileOutputWithContext(context.Context) OSProfileOutput
}

type OSProfileArgs struct {
	AdminPassword        pulumi.StringPtrInput        `pulumi:"adminPassword"`
	AdminUsername        pulumi.StringPtrInput        `pulumi:"adminUsername"`
	ComputerName         pulumi.StringPtrInput        `pulumi:"computerName"`
	CustomData           pulumi.StringPtrInput        `pulumi:"customData"`
	LinuxConfiguration   LinuxConfigurationPtrInput   `pulumi:"linuxConfiguration"`
	Secrets              VaultSecretGroupArrayInput   `pulumi:"secrets"`
	WindowsConfiguration WindowsConfigurationPtrInput `pulumi:"windowsConfiguration"`
}

func (OSProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*OSProfile)(nil)).Elem()
}

func (i OSProfileArgs) ToOSProfileOutput() OSProfileOutput {
	return i.ToOSProfileOutputWithContext(context.Background())
}

func (i OSProfileArgs) ToOSProfileOutputWithContext(ctx context.Context) OSProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OSProfileOutput)
}

func (i OSProfileArgs) ToOSProfilePtrOutput() OSProfilePtrOutput {
	return i.ToOSProfilePtrOutputWithContext(context.Background())
}

func (i OSProfileArgs) ToOSProfilePtrOutputWithContext(ctx context.Context) OSProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OSProfileOutput).ToOSProfilePtrOutputWithContext(ctx)
}









type OSProfilePtrInput interface {
	pulumi.Input

	ToOSProfilePtrOutput() OSProfilePtrOutput
	ToOSProfilePtrOutputWithContext(context.Context) OSProfilePtrOutput
}

type osprofilePtrType OSProfileArgs

func OSProfilePtr(v *OSProfileArgs) OSProfilePtrInput {
	return (*osprofilePtrType)(v)
}

func (*osprofilePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**OSProfile)(nil)).Elem()
}

func (i *osprofilePtrType) ToOSProfilePtrOutput() OSProfilePtrOutput {
	return i.ToOSProfilePtrOutputWithContext(context.Background())
}

func (i *osprofilePtrType) ToOSProfilePtrOutputWithContext(ctx context.Context) OSProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OSProfilePtrOutput)
}

type OSProfileOutput struct{ *pulumi.OutputState }

func (OSProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*OSProfile)(nil)).Elem()
}

func (o OSProfileOutput) ToOSProfileOutput() OSProfileOutput {
	return o
}

func (o OSProfileOutput) ToOSProfileOutputWithContext(ctx context.Context) OSProfileOutput {
	return o
}

func (o OSProfileOutput) ToOSProfilePtrOutput() OSProfilePtrOutput {
	return o.ToOSProfilePtrOutputWithContext(context.Background())
}

func (o OSProfileOutput) ToOSProfilePtrOutputWithContext(ctx context.Context) OSProfilePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v OSProfile) *OSProfile {
		return &v
	}).(OSProfilePtrOutput)
}

func (o OSProfileOutput) AdminPassword() pulumi.StringPtrOutput {
	return o.ApplyT(func(v OSProfile) *string { return v.AdminPassword }).(pulumi.StringPtrOutput)
}

func (o OSProfileOutput) AdminUsername() pulumi.StringPtrOutput {
	return o.ApplyT(func(v OSProfile) *string { return v.AdminUsername }).(pulumi.StringPtrOutput)
}

func (o OSProfileOutput) ComputerName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v OSProfile) *string { return v.ComputerName }).(pulumi.StringPtrOutput)
}

func (o OSProfileOutput) CustomData() pulumi.StringPtrOutput {
	return o.ApplyT(func(v OSProfile) *string { return v.CustomData }).(pulumi.StringPtrOutput)
}

func (o OSProfileOutput) LinuxConfiguration() LinuxConfigurationPtrOutput {
	return o.ApplyT(func(v OSProfile) *LinuxConfiguration { return v.LinuxConfiguration }).(LinuxConfigurationPtrOutput)
}

func (o OSProfileOutput) Secrets() VaultSecretGroupArrayOutput {
	return o.ApplyT(func(v OSProfile) []VaultSecretGroup { return v.Secrets }).(VaultSecretGroupArrayOutput)
}

func (o OSProfileOutput) WindowsConfiguration() WindowsConfigurationPtrOutput {
	return o.ApplyT(func(v OSProfile) *WindowsConfiguration { return v.WindowsConfiguration }).(WindowsConfigurationPtrOutput)
}

type OSProfilePtrOutput struct{ *pulumi.OutputState }

func (OSProfilePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OSProfile)(nil)).Elem()
}

func (o OSProfilePtrOutput) ToOSProfilePtrOutput() OSProfilePtrOutput {
	return o
}

func (o OSProfilePtrOutput) ToOSProfilePtrOutputWithContext(ctx context.Context) OSProfilePtrOutput {
	return o
}

func (o OSProfilePtrOutput) Elem() OSProfileOutput {
	return o.ApplyT(func(v *OSProfile) OSProfile {
		if v != nil {
			return *v
		}
		var ret OSProfile
		return ret
	}).(OSProfileOutput)
}

func (o OSProfilePtrOutput) AdminPassword() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OSProfile) *string {
		if v == nil {
			return nil
		}
		return v.AdminPassword
	}).(pulumi.StringPtrOutput)
}

func (o OSProfilePtrOutput) AdminUsername() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OSProfile) *string {
		if v == nil {
			return nil
		}
		return v.AdminUsername
	}).(pulumi.StringPtrOutput)
}

func (o OSProfilePtrOutput) ComputerName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OSProfile) *string {
		if v == nil {
			return nil
		}
		return v.ComputerName
	}).(pulumi.StringPtrOutput)
}

func (o OSProfilePtrOutput) CustomData() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OSProfile) *string {
		if v == nil {
			return nil
		}
		return v.CustomData
	}).(pulumi.StringPtrOutput)
}

func (o OSProfilePtrOutput) LinuxConfiguration() LinuxConfigurationPtrOutput {
	return o.ApplyT(func(v *OSProfile) *LinuxConfiguration {
		if v == nil {
			return nil
		}
		return v.LinuxConfiguration
	}).(LinuxConfigurationPtrOutput)
}

func (o OSProfilePtrOutput) Secrets() VaultSecretGroupArrayOutput {
	return o.ApplyT(func(v *OSProfile) []VaultSecretGroup {
		if v == nil {
			return nil
		}
		return v.Secrets
	}).(VaultSecretGroupArrayOutput)
}

func (o OSProfilePtrOutput) WindowsConfiguration() WindowsConfigurationPtrOutput {
	return o.ApplyT(func(v *OSProfile) *WindowsConfiguration {
		if v == nil {
			return nil
		}
		return v.WindowsConfiguration
	}).(WindowsConfigurationPtrOutput)
}

type OSProfileResponse struct {
	AdminPassword        *string                       `pulumi:"adminPassword"`
	AdminUsername        *string                       `pulumi:"adminUsername"`
	ComputerName         *string                       `pulumi:"computerName"`
	CustomData           *string                       `pulumi:"customData"`
	LinuxConfiguration   *LinuxConfigurationResponse   `pulumi:"linuxConfiguration"`
	Secrets              []VaultSecretGroupResponse    `pulumi:"secrets"`
	WindowsConfiguration *WindowsConfigurationResponse `pulumi:"windowsConfiguration"`
}

type OSProfileResponseOutput struct{ *pulumi.OutputState }

func (OSProfileResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*OSProfileResponse)(nil)).Elem()
}

func (o OSProfileResponseOutput) ToOSProfileResponseOutput() OSProfileResponseOutput {
	return o
}

func (o OSProfileResponseOutput) ToOSProfileResponseOutputWithContext(ctx context.Context) OSProfileResponseOutput {
	return o
}

func (o OSProfileResponseOutput) AdminPassword() pulumi.StringPtrOutput {
	return o.ApplyT(func(v OSProfileResponse) *string { return v.AdminPassword }).(pulumi.StringPtrOutput)
}

func (o OSProfileResponseOutput) AdminUsername() pulumi.StringPtrOutput {
	return o.ApplyT(func(v OSProfileResponse) *string { return v.AdminUsername }).(pulumi.StringPtrOutput)
}

func (o OSProfileResponseOutput) ComputerName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v OSProfileResponse) *string { return v.ComputerName }).(pulumi.StringPtrOutput)
}

func (o OSProfileResponseOutput) CustomData() pulumi.StringPtrOutput {
	return o.ApplyT(func(v OSProfileResponse) *string { return v.CustomData }).(pulumi.StringPtrOutput)
}

func (o OSProfileResponseOutput) LinuxConfiguration() LinuxConfigurationResponsePtrOutput {
	return o.ApplyT(func(v OSProfileResponse) *LinuxConfigurationResponse { return v.LinuxConfiguration }).(LinuxConfigurationResponsePtrOutput)
}

func (o OSProfileResponseOutput) Secrets() VaultSecretGroupResponseArrayOutput {
	return o.ApplyT(func(v OSProfileResponse) []VaultSecretGroupResponse { return v.Secrets }).(VaultSecretGroupResponseArrayOutput)
}

func (o OSProfileResponseOutput) WindowsConfiguration() WindowsConfigurationResponsePtrOutput {
	return o.ApplyT(func(v OSProfileResponse) *WindowsConfigurationResponse { return v.WindowsConfiguration }).(WindowsConfigurationResponsePtrOutput)
}

type OSProfileResponsePtrOutput struct{ *pulumi.OutputState }

func (OSProfileResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OSProfileResponse)(nil)).Elem()
}

func (o OSProfileResponsePtrOutput) ToOSProfileResponsePtrOutput() OSProfileResponsePtrOutput {
	return o
}

func (o OSProfileResponsePtrOutput) ToOSProfileResponsePtrOutputWithContext(ctx context.Context) OSProfileResponsePtrOutput {
	return o
}

func (o OSProfileResponsePtrOutput) Elem() OSProfileResponseOutput {
	return o.ApplyT(func(v *OSProfileResponse) OSProfileResponse {
		if v != nil {
			return *v
		}
		var ret OSProfileResponse
		return ret
	}).(OSProfileResponseOutput)
}

func (o OSProfileResponsePtrOutput) AdminPassword() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OSProfileResponse) *string {
		if v == nil {
			return nil
		}
		return v.AdminPassword
	}).(pulumi.StringPtrOutput)
}

func (o OSProfileResponsePtrOutput) AdminUsername() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OSProfileResponse) *string {
		if v == nil {
			return nil
		}
		return v.AdminUsername
	}).(pulumi.StringPtrOutput)
}

func (o OSProfileResponsePtrOutput) ComputerName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OSProfileResponse) *string {
		if v == nil {
			return nil
		}
		return v.ComputerName
	}).(pulumi.StringPtrOutput)
}

func (o OSProfileResponsePtrOutput) CustomData() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OSProfileResponse) *string {
		if v == nil {
			return nil
		}
		return v.CustomData
	}).(pulumi.StringPtrOutput)
}

func (o OSProfileResponsePtrOutput) LinuxConfiguration() LinuxConfigurationResponsePtrOutput {
	return o.ApplyT(func(v *OSProfileResponse) *LinuxConfigurationResponse {
		if v == nil {
			return nil
		}
		return v.LinuxConfiguration
	}).(LinuxConfigurationResponsePtrOutput)
}

func (o OSProfileResponsePtrOutput) Secrets() VaultSecretGroupResponseArrayOutput {
	return o.ApplyT(func(v *OSProfileResponse) []VaultSecretGroupResponse {
		if v == nil {
			return nil
		}
		return v.Secrets
	}).(VaultSecretGroupResponseArrayOutput)
}

func (o OSProfileResponsePtrOutput) WindowsConfiguration() WindowsConfigurationResponsePtrOutput {
	return o.ApplyT(func(v *OSProfileResponse) *WindowsConfigurationResponse {
		if v == nil {
			return nil
		}
		return v.WindowsConfiguration
	}).(WindowsConfigurationResponsePtrOutput)
}

type Plan struct {
	Name          *string `pulumi:"name"`
	Product       *string `pulumi:"product"`
	PromotionCode *string `pulumi:"promotionCode"`
	Publisher     *string `pulumi:"publisher"`
}





type PlanInput interface {
	pulumi.Input

	ToPlanOutput() PlanOutput
	ToPlanOutputWithContext(context.Context) PlanOutput
}

type PlanArgs struct {
	Name          pulumi.StringPtrInput `pulumi:"name"`
	Product       pulumi.StringPtrInput `pulumi:"product"`
	PromotionCode pulumi.StringPtrInput `pulumi:"promotionCode"`
	Publisher     pulumi.StringPtrInput `pulumi:"publisher"`
}

func (PlanArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Plan)(nil)).Elem()
}

func (i PlanArgs) ToPlanOutput() PlanOutput {
	return i.ToPlanOutputWithContext(context.Background())
}

func (i PlanArgs) ToPlanOutputWithContext(ctx context.Context) PlanOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PlanOutput)
}

func (i PlanArgs) ToPlanPtrOutput() PlanPtrOutput {
	return i.ToPlanPtrOutputWithContext(context.Background())
}

func (i PlanArgs) ToPlanPtrOutputWithContext(ctx context.Context) PlanPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PlanOutput).ToPlanPtrOutputWithContext(ctx)
}









type PlanPtrInput interface {
	pulumi.Input

	ToPlanPtrOutput() PlanPtrOutput
	ToPlanPtrOutputWithContext(context.Context) PlanPtrOutput
}

type planPtrType PlanArgs

func PlanPtr(v *PlanArgs) PlanPtrInput {
	return (*planPtrType)(v)
}

func (*planPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Plan)(nil)).Elem()
}

func (i *planPtrType) ToPlanPtrOutput() PlanPtrOutput {
	return i.ToPlanPtrOutputWithContext(context.Background())
}

func (i *planPtrType) ToPlanPtrOutputWithContext(ctx context.Context) PlanPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PlanPtrOutput)
}

type PlanOutput struct{ *pulumi.OutputState }

func (PlanOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Plan)(nil)).Elem()
}

func (o PlanOutput) ToPlanOutput() PlanOutput {
	return o
}

func (o PlanOutput) ToPlanOutputWithContext(ctx context.Context) PlanOutput {
	return o
}

func (o PlanOutput) ToPlanPtrOutput() PlanPtrOutput {
	return o.ToPlanPtrOutputWithContext(context.Background())
}

func (o PlanOutput) ToPlanPtrOutputWithContext(ctx context.Context) PlanPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Plan) *Plan {
		return &v
	}).(PlanPtrOutput)
}

func (o PlanOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Plan) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o PlanOutput) Product() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Plan) *string { return v.Product }).(pulumi.StringPtrOutput)
}

func (o PlanOutput) PromotionCode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Plan) *string { return v.PromotionCode }).(pulumi.StringPtrOutput)
}

func (o PlanOutput) Publisher() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Plan) *string { return v.Publisher }).(pulumi.StringPtrOutput)
}

type PlanPtrOutput struct{ *pulumi.OutputState }

func (PlanPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Plan)(nil)).Elem()
}

func (o PlanPtrOutput) ToPlanPtrOutput() PlanPtrOutput {
	return o
}

func (o PlanPtrOutput) ToPlanPtrOutputWithContext(ctx context.Context) PlanPtrOutput {
	return o
}

func (o PlanPtrOutput) Elem() PlanOutput {
	return o.ApplyT(func(v *Plan) Plan {
		if v != nil {
			return *v
		}
		var ret Plan
		return ret
	}).(PlanOutput)
}

func (o PlanPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Plan) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o PlanPtrOutput) Product() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Plan) *string {
		if v == nil {
			return nil
		}
		return v.Product
	}).(pulumi.StringPtrOutput)
}

func (o PlanPtrOutput) PromotionCode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Plan) *string {
		if v == nil {
			return nil
		}
		return v.PromotionCode
	}).(pulumi.StringPtrOutput)
}

func (o PlanPtrOutput) Publisher() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Plan) *string {
		if v == nil {
			return nil
		}
		return v.Publisher
	}).(pulumi.StringPtrOutput)
}

type PlanResponse struct {
	Name          *string `pulumi:"name"`
	Product       *string `pulumi:"product"`
	PromotionCode *string `pulumi:"promotionCode"`
	Publisher     *string `pulumi:"publisher"`
}

type PlanResponseOutput struct{ *pulumi.OutputState }

func (PlanResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PlanResponse)(nil)).Elem()
}

func (o PlanResponseOutput) ToPlanResponseOutput() PlanResponseOutput {
	return o
}

func (o PlanResponseOutput) ToPlanResponseOutputWithContext(ctx context.Context) PlanResponseOutput {
	return o
}

func (o PlanResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PlanResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o PlanResponseOutput) Product() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PlanResponse) *string { return v.Product }).(pulumi.StringPtrOutput)
}

func (o PlanResponseOutput) PromotionCode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PlanResponse) *string { return v.PromotionCode }).(pulumi.StringPtrOutput)
}

func (o PlanResponseOutput) Publisher() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PlanResponse) *string { return v.Publisher }).(pulumi.StringPtrOutput)
}

type PlanResponsePtrOutput struct{ *pulumi.OutputState }

func (PlanResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PlanResponse)(nil)).Elem()
}

func (o PlanResponsePtrOutput) ToPlanResponsePtrOutput() PlanResponsePtrOutput {
	return o
}

func (o PlanResponsePtrOutput) ToPlanResponsePtrOutputWithContext(ctx context.Context) PlanResponsePtrOutput {
	return o
}

func (o PlanResponsePtrOutput) Elem() PlanResponseOutput {
	return o.ApplyT(func(v *PlanResponse) PlanResponse {
		if v != nil {
			return *v
		}
		var ret PlanResponse
		return ret
	}).(PlanResponseOutput)
}

func (o PlanResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PlanResponse) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o PlanResponsePtrOutput) Product() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PlanResponse) *string {
		if v == nil {
			return nil
		}
		return v.Product
	}).(pulumi.StringPtrOutput)
}

func (o PlanResponsePtrOutput) PromotionCode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PlanResponse) *string {
		if v == nil {
			return nil
		}
		return v.PromotionCode
	}).(pulumi.StringPtrOutput)
}

func (o PlanResponsePtrOutput) Publisher() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PlanResponse) *string {
		if v == nil {
			return nil
		}
		return v.Publisher
	}).(pulumi.StringPtrOutput)
}

type RollingUpgradePolicy struct {
	MaxBatchInstancePercent             *int    `pulumi:"maxBatchInstancePercent"`
	MaxUnhealthyInstancePercent         *int    `pulumi:"maxUnhealthyInstancePercent"`
	MaxUnhealthyUpgradedInstancePercent *int    `pulumi:"maxUnhealthyUpgradedInstancePercent"`
	PauseTimeBetweenBatches             *string `pulumi:"pauseTimeBetweenBatches"`
}





type RollingUpgradePolicyInput interface {
	pulumi.Input

	ToRollingUpgradePolicyOutput() RollingUpgradePolicyOutput
	ToRollingUpgradePolicyOutputWithContext(context.Context) RollingUpgradePolicyOutput
}

type RollingUpgradePolicyArgs struct {
	MaxBatchInstancePercent             pulumi.IntPtrInput    `pulumi:"maxBatchInstancePercent"`
	MaxUnhealthyInstancePercent         pulumi.IntPtrInput    `pulumi:"maxUnhealthyInstancePercent"`
	MaxUnhealthyUpgradedInstancePercent pulumi.IntPtrInput    `pulumi:"maxUnhealthyUpgradedInstancePercent"`
	PauseTimeBetweenBatches             pulumi.StringPtrInput `pulumi:"pauseTimeBetweenBatches"`
}

func (RollingUpgradePolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RollingUpgradePolicy)(nil)).Elem()
}

func (i RollingUpgradePolicyArgs) ToRollingUpgradePolicyOutput() RollingUpgradePolicyOutput {
	return i.ToRollingUpgradePolicyOutputWithContext(context.Background())
}

func (i RollingUpgradePolicyArgs) ToRollingUpgradePolicyOutputWithContext(ctx context.Context) RollingUpgradePolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RollingUpgradePolicyOutput)
}

func (i RollingUpgradePolicyArgs) ToRollingUpgradePolicyPtrOutput() RollingUpgradePolicyPtrOutput {
	return i.ToRollingUpgradePolicyPtrOutputWithContext(context.Background())
}

func (i RollingUpgradePolicyArgs) ToRollingUpgradePolicyPtrOutputWithContext(ctx context.Context) RollingUpgradePolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RollingUpgradePolicyOutput).ToRollingUpgradePolicyPtrOutputWithContext(ctx)
}









type RollingUpgradePolicyPtrInput interface {
	pulumi.Input

	ToRollingUpgradePolicyPtrOutput() RollingUpgradePolicyPtrOutput
	ToRollingUpgradePolicyPtrOutputWithContext(context.Context) RollingUpgradePolicyPtrOutput
}

type rollingUpgradePolicyPtrType RollingUpgradePolicyArgs

func RollingUpgradePolicyPtr(v *RollingUpgradePolicyArgs) RollingUpgradePolicyPtrInput {
	return (*rollingUpgradePolicyPtrType)(v)
}

func (*rollingUpgradePolicyPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**RollingUpgradePolicy)(nil)).Elem()
}

func (i *rollingUpgradePolicyPtrType) ToRollingUpgradePolicyPtrOutput() RollingUpgradePolicyPtrOutput {
	return i.ToRollingUpgradePolicyPtrOutputWithContext(context.Background())
}

func (i *rollingUpgradePolicyPtrType) ToRollingUpgradePolicyPtrOutputWithContext(ctx context.Context) RollingUpgradePolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RollingUpgradePolicyPtrOutput)
}

type RollingUpgradePolicyOutput struct{ *pulumi.OutputState }

func (RollingUpgradePolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RollingUpgradePolicy)(nil)).Elem()
}

func (o RollingUpgradePolicyOutput) ToRollingUpgradePolicyOutput() RollingUpgradePolicyOutput {
	return o
}

func (o RollingUpgradePolicyOutput) ToRollingUpgradePolicyOutputWithContext(ctx context.Context) RollingUpgradePolicyOutput {
	return o
}

func (o RollingUpgradePolicyOutput) ToRollingUpgradePolicyPtrOutput() RollingUpgradePolicyPtrOutput {
	return o.ToRollingUpgradePolicyPtrOutputWithContext(context.Background())
}

func (o RollingUpgradePolicyOutput) ToRollingUpgradePolicyPtrOutputWithContext(ctx context.Context) RollingUpgradePolicyPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v RollingUpgradePolicy) *RollingUpgradePolicy {
		return &v
	}).(RollingUpgradePolicyPtrOutput)
}

func (o RollingUpgradePolicyOutput) MaxBatchInstancePercent() pulumi.IntPtrOutput {
	return o.ApplyT(func(v RollingUpgradePolicy) *int { return v.MaxBatchInstancePercent }).(pulumi.IntPtrOutput)
}

func (o RollingUpgradePolicyOutput) MaxUnhealthyInstancePercent() pulumi.IntPtrOutput {
	return o.ApplyT(func(v RollingUpgradePolicy) *int { return v.MaxUnhealthyInstancePercent }).(pulumi.IntPtrOutput)
}

func (o RollingUpgradePolicyOutput) MaxUnhealthyUpgradedInstancePercent() pulumi.IntPtrOutput {
	return o.ApplyT(func(v RollingUpgradePolicy) *int { return v.MaxUnhealthyUpgradedInstancePercent }).(pulumi.IntPtrOutput)
}

func (o RollingUpgradePolicyOutput) PauseTimeBetweenBatches() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RollingUpgradePolicy) *string { return v.PauseTimeBetweenBatches }).(pulumi.StringPtrOutput)
}

type RollingUpgradePolicyPtrOutput struct{ *pulumi.OutputState }

func (RollingUpgradePolicyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RollingUpgradePolicy)(nil)).Elem()
}

func (o RollingUpgradePolicyPtrOutput) ToRollingUpgradePolicyPtrOutput() RollingUpgradePolicyPtrOutput {
	return o
}

func (o RollingUpgradePolicyPtrOutput) ToRollingUpgradePolicyPtrOutputWithContext(ctx context.Context) RollingUpgradePolicyPtrOutput {
	return o
}

func (o RollingUpgradePolicyPtrOutput) Elem() RollingUpgradePolicyOutput {
	return o.ApplyT(func(v *RollingUpgradePolicy) RollingUpgradePolicy {
		if v != nil {
			return *v
		}
		var ret RollingUpgradePolicy
		return ret
	}).(RollingUpgradePolicyOutput)
}

func (o RollingUpgradePolicyPtrOutput) MaxBatchInstancePercent() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *RollingUpgradePolicy) *int {
		if v == nil {
			return nil
		}
		return v.MaxBatchInstancePercent
	}).(pulumi.IntPtrOutput)
}

func (o RollingUpgradePolicyPtrOutput) MaxUnhealthyInstancePercent() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *RollingUpgradePolicy) *int {
		if v == nil {
			return nil
		}
		return v.MaxUnhealthyInstancePercent
	}).(pulumi.IntPtrOutput)
}

func (o RollingUpgradePolicyPtrOutput) MaxUnhealthyUpgradedInstancePercent() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *RollingUpgradePolicy) *int {
		if v == nil {
			return nil
		}
		return v.MaxUnhealthyUpgradedInstancePercent
	}).(pulumi.IntPtrOutput)
}

func (o RollingUpgradePolicyPtrOutput) PauseTimeBetweenBatches() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RollingUpgradePolicy) *string {
		if v == nil {
			return nil
		}
		return v.PauseTimeBetweenBatches
	}).(pulumi.StringPtrOutput)
}

type RollingUpgradePolicyResponse struct {
	MaxBatchInstancePercent             *int    `pulumi:"maxBatchInstancePercent"`
	MaxUnhealthyInstancePercent         *int    `pulumi:"maxUnhealthyInstancePercent"`
	MaxUnhealthyUpgradedInstancePercent *int    `pulumi:"maxUnhealthyUpgradedInstancePercent"`
	PauseTimeBetweenBatches             *string `pulumi:"pauseTimeBetweenBatches"`
}

type RollingUpgradePolicyResponseOutput struct{ *pulumi.OutputState }

func (RollingUpgradePolicyResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RollingUpgradePolicyResponse)(nil)).Elem()
}

func (o RollingUpgradePolicyResponseOutput) ToRollingUpgradePolicyResponseOutput() RollingUpgradePolicyResponseOutput {
	return o
}

func (o RollingUpgradePolicyResponseOutput) ToRollingUpgradePolicyResponseOutputWithContext(ctx context.Context) RollingUpgradePolicyResponseOutput {
	return o
}

func (o RollingUpgradePolicyResponseOutput) MaxBatchInstancePercent() pulumi.IntPtrOutput {
	return o.ApplyT(func(v RollingUpgradePolicyResponse) *int { return v.MaxBatchInstancePercent }).(pulumi.IntPtrOutput)
}

func (o RollingUpgradePolicyResponseOutput) MaxUnhealthyInstancePercent() pulumi.IntPtrOutput {
	return o.ApplyT(func(v RollingUpgradePolicyResponse) *int { return v.MaxUnhealthyInstancePercent }).(pulumi.IntPtrOutput)
}

func (o RollingUpgradePolicyResponseOutput) MaxUnhealthyUpgradedInstancePercent() pulumi.IntPtrOutput {
	return o.ApplyT(func(v RollingUpgradePolicyResponse) *int { return v.MaxUnhealthyUpgradedInstancePercent }).(pulumi.IntPtrOutput)
}

func (o RollingUpgradePolicyResponseOutput) PauseTimeBetweenBatches() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RollingUpgradePolicyResponse) *string { return v.PauseTimeBetweenBatches }).(pulumi.StringPtrOutput)
}

type RollingUpgradePolicyResponsePtrOutput struct{ *pulumi.OutputState }

func (RollingUpgradePolicyResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RollingUpgradePolicyResponse)(nil)).Elem()
}

func (o RollingUpgradePolicyResponsePtrOutput) ToRollingUpgradePolicyResponsePtrOutput() RollingUpgradePolicyResponsePtrOutput {
	return o
}

func (o RollingUpgradePolicyResponsePtrOutput) ToRollingUpgradePolicyResponsePtrOutputWithContext(ctx context.Context) RollingUpgradePolicyResponsePtrOutput {
	return o
}

func (o RollingUpgradePolicyResponsePtrOutput) Elem() RollingUpgradePolicyResponseOutput {
	return o.ApplyT(func(v *RollingUpgradePolicyResponse) RollingUpgradePolicyResponse {
		if v != nil {
			return *v
		}
		var ret RollingUpgradePolicyResponse
		return ret
	}).(RollingUpgradePolicyResponseOutput)
}

func (o RollingUpgradePolicyResponsePtrOutput) MaxBatchInstancePercent() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *RollingUpgradePolicyResponse) *int {
		if v == nil {
			return nil
		}
		return v.MaxBatchInstancePercent
	}).(pulumi.IntPtrOutput)
}

func (o RollingUpgradePolicyResponsePtrOutput) MaxUnhealthyInstancePercent() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *RollingUpgradePolicyResponse) *int {
		if v == nil {
			return nil
		}
		return v.MaxUnhealthyInstancePercent
	}).(pulumi.IntPtrOutput)
}

func (o RollingUpgradePolicyResponsePtrOutput) MaxUnhealthyUpgradedInstancePercent() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *RollingUpgradePolicyResponse) *int {
		if v == nil {
			return nil
		}
		return v.MaxUnhealthyUpgradedInstancePercent
	}).(pulumi.IntPtrOutput)
}

func (o RollingUpgradePolicyResponsePtrOutput) PauseTimeBetweenBatches() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RollingUpgradePolicyResponse) *string {
		if v == nil {
			return nil
		}
		return v.PauseTimeBetweenBatches
	}).(pulumi.StringPtrOutput)
}

type Sku struct {
	Capacity *float64 `pulumi:"capacity"`
	Name     *string  `pulumi:"name"`
	Tier     *string  `pulumi:"tier"`
}





type SkuInput interface {
	pulumi.Input

	ToSkuOutput() SkuOutput
	ToSkuOutputWithContext(context.Context) SkuOutput
}

type SkuArgs struct {
	Capacity pulumi.Float64PtrInput `pulumi:"capacity"`
	Name     pulumi.StringPtrInput  `pulumi:"name"`
	Tier     pulumi.StringPtrInput  `pulumi:"tier"`
}

func (SkuArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Sku)(nil)).Elem()
}

func (i SkuArgs) ToSkuOutput() SkuOutput {
	return i.ToSkuOutputWithContext(context.Background())
}

func (i SkuArgs) ToSkuOutputWithContext(ctx context.Context) SkuOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuOutput)
}

func (i SkuArgs) ToSkuPtrOutput() SkuPtrOutput {
	return i.ToSkuPtrOutputWithContext(context.Background())
}

func (i SkuArgs) ToSkuPtrOutputWithContext(ctx context.Context) SkuPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuOutput).ToSkuPtrOutputWithContext(ctx)
}









type SkuPtrInput interface {
	pulumi.Input

	ToSkuPtrOutput() SkuPtrOutput
	ToSkuPtrOutputWithContext(context.Context) SkuPtrOutput
}

type skuPtrType SkuArgs

func SkuPtr(v *SkuArgs) SkuPtrInput {
	return (*skuPtrType)(v)
}

func (*skuPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Sku)(nil)).Elem()
}

func (i *skuPtrType) ToSkuPtrOutput() SkuPtrOutput {
	return i.ToSkuPtrOutputWithContext(context.Background())
}

func (i *skuPtrType) ToSkuPtrOutputWithContext(ctx context.Context) SkuPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuPtrOutput)
}

type SkuOutput struct{ *pulumi.OutputState }

func (SkuOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Sku)(nil)).Elem()
}

func (o SkuOutput) ToSkuOutput() SkuOutput {
	return o
}

func (o SkuOutput) ToSkuOutputWithContext(ctx context.Context) SkuOutput {
	return o
}

func (o SkuOutput) ToSkuPtrOutput() SkuPtrOutput {
	return o.ToSkuPtrOutputWithContext(context.Background())
}

func (o SkuOutput) ToSkuPtrOutputWithContext(ctx context.Context) SkuPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Sku) *Sku {
		return &v
	}).(SkuPtrOutput)
}

func (o SkuOutput) Capacity() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v Sku) *float64 { return v.Capacity }).(pulumi.Float64PtrOutput)
}

func (o SkuOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Sku) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o SkuOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Sku) *string { return v.Tier }).(pulumi.StringPtrOutput)
}

type SkuPtrOutput struct{ *pulumi.OutputState }

func (SkuPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Sku)(nil)).Elem()
}

func (o SkuPtrOutput) ToSkuPtrOutput() SkuPtrOutput {
	return o
}

func (o SkuPtrOutput) ToSkuPtrOutputWithContext(ctx context.Context) SkuPtrOutput {
	return o
}

func (o SkuPtrOutput) Elem() SkuOutput {
	return o.ApplyT(func(v *Sku) Sku {
		if v != nil {
			return *v
		}
		var ret Sku
		return ret
	}).(SkuOutput)
}

func (o SkuPtrOutput) Capacity() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *Sku) *float64 {
		if v == nil {
			return nil
		}
		return v.Capacity
	}).(pulumi.Float64PtrOutput)
}

func (o SkuPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Sku) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o SkuPtrOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Sku) *string {
		if v == nil {
			return nil
		}
		return v.Tier
	}).(pulumi.StringPtrOutput)
}

type SkuResponse struct {
	Capacity *float64 `pulumi:"capacity"`
	Name     *string  `pulumi:"name"`
	Tier     *string  `pulumi:"tier"`
}

type SkuResponseOutput struct{ *pulumi.OutputState }

func (SkuResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SkuResponse)(nil)).Elem()
}

func (o SkuResponseOutput) ToSkuResponseOutput() SkuResponseOutput {
	return o
}

func (o SkuResponseOutput) ToSkuResponseOutputWithContext(ctx context.Context) SkuResponseOutput {
	return o
}

func (o SkuResponseOutput) Capacity() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v SkuResponse) *float64 { return v.Capacity }).(pulumi.Float64PtrOutput)
}

func (o SkuResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SkuResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o SkuResponseOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SkuResponse) *string { return v.Tier }).(pulumi.StringPtrOutput)
}

type SkuResponsePtrOutput struct{ *pulumi.OutputState }

func (SkuResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SkuResponse)(nil)).Elem()
}

func (o SkuResponsePtrOutput) ToSkuResponsePtrOutput() SkuResponsePtrOutput {
	return o
}

func (o SkuResponsePtrOutput) ToSkuResponsePtrOutputWithContext(ctx context.Context) SkuResponsePtrOutput {
	return o
}

func (o SkuResponsePtrOutput) Elem() SkuResponseOutput {
	return o.ApplyT(func(v *SkuResponse) SkuResponse {
		if v != nil {
			return *v
		}
		var ret SkuResponse
		return ret
	}).(SkuResponseOutput)
}

func (o SkuResponsePtrOutput) Capacity() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *SkuResponse) *float64 {
		if v == nil {
			return nil
		}
		return v.Capacity
	}).(pulumi.Float64PtrOutput)
}

func (o SkuResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SkuResponse) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o SkuResponsePtrOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SkuResponse) *string {
		if v == nil {
			return nil
		}
		return v.Tier
	}).(pulumi.StringPtrOutput)
}

type SnapshotSku struct {
	Name *string `pulumi:"name"`
}





type SnapshotSkuInput interface {
	pulumi.Input

	ToSnapshotSkuOutput() SnapshotSkuOutput
	ToSnapshotSkuOutputWithContext(context.Context) SnapshotSkuOutput
}

type SnapshotSkuArgs struct {
	Name pulumi.StringPtrInput `pulumi:"name"`
}

func (SnapshotSkuArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SnapshotSku)(nil)).Elem()
}

func (i SnapshotSkuArgs) ToSnapshotSkuOutput() SnapshotSkuOutput {
	return i.ToSnapshotSkuOutputWithContext(context.Background())
}

func (i SnapshotSkuArgs) ToSnapshotSkuOutputWithContext(ctx context.Context) SnapshotSkuOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SnapshotSkuOutput)
}

func (i SnapshotSkuArgs) ToSnapshotSkuPtrOutput() SnapshotSkuPtrOutput {
	return i.ToSnapshotSkuPtrOutputWithContext(context.Background())
}

func (i SnapshotSkuArgs) ToSnapshotSkuPtrOutputWithContext(ctx context.Context) SnapshotSkuPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SnapshotSkuOutput).ToSnapshotSkuPtrOutputWithContext(ctx)
}









type SnapshotSkuPtrInput interface {
	pulumi.Input

	ToSnapshotSkuPtrOutput() SnapshotSkuPtrOutput
	ToSnapshotSkuPtrOutputWithContext(context.Context) SnapshotSkuPtrOutput
}

type snapshotSkuPtrType SnapshotSkuArgs

func SnapshotSkuPtr(v *SnapshotSkuArgs) SnapshotSkuPtrInput {
	return (*snapshotSkuPtrType)(v)
}

func (*snapshotSkuPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SnapshotSku)(nil)).Elem()
}

func (i *snapshotSkuPtrType) ToSnapshotSkuPtrOutput() SnapshotSkuPtrOutput {
	return i.ToSnapshotSkuPtrOutputWithContext(context.Background())
}

func (i *snapshotSkuPtrType) ToSnapshotSkuPtrOutputWithContext(ctx context.Context) SnapshotSkuPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SnapshotSkuPtrOutput)
}

type SnapshotSkuOutput struct{ *pulumi.OutputState }

func (SnapshotSkuOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SnapshotSku)(nil)).Elem()
}

func (o SnapshotSkuOutput) ToSnapshotSkuOutput() SnapshotSkuOutput {
	return o
}

func (o SnapshotSkuOutput) ToSnapshotSkuOutputWithContext(ctx context.Context) SnapshotSkuOutput {
	return o
}

func (o SnapshotSkuOutput) ToSnapshotSkuPtrOutput() SnapshotSkuPtrOutput {
	return o.ToSnapshotSkuPtrOutputWithContext(context.Background())
}

func (o SnapshotSkuOutput) ToSnapshotSkuPtrOutputWithContext(ctx context.Context) SnapshotSkuPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SnapshotSku) *SnapshotSku {
		return &v
	}).(SnapshotSkuPtrOutput)
}

func (o SnapshotSkuOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SnapshotSku) *string { return v.Name }).(pulumi.StringPtrOutput)
}

type SnapshotSkuPtrOutput struct{ *pulumi.OutputState }

func (SnapshotSkuPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SnapshotSku)(nil)).Elem()
}

func (o SnapshotSkuPtrOutput) ToSnapshotSkuPtrOutput() SnapshotSkuPtrOutput {
	return o
}

func (o SnapshotSkuPtrOutput) ToSnapshotSkuPtrOutputWithContext(ctx context.Context) SnapshotSkuPtrOutput {
	return o
}

func (o SnapshotSkuPtrOutput) Elem() SnapshotSkuOutput {
	return o.ApplyT(func(v *SnapshotSku) SnapshotSku {
		if v != nil {
			return *v
		}
		var ret SnapshotSku
		return ret
	}).(SnapshotSkuOutput)
}

func (o SnapshotSkuPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SnapshotSku) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

type SnapshotSkuResponse struct {
	Name *string `pulumi:"name"`
	Tier string  `pulumi:"tier"`
}


func (val *SnapshotSkuResponse) Defaults() *SnapshotSkuResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Tier) {
		tmp.Tier = "Standard"
	}
	return &tmp
}

type SnapshotSkuResponseOutput struct{ *pulumi.OutputState }

func (SnapshotSkuResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SnapshotSkuResponse)(nil)).Elem()
}

func (o SnapshotSkuResponseOutput) ToSnapshotSkuResponseOutput() SnapshotSkuResponseOutput {
	return o
}

func (o SnapshotSkuResponseOutput) ToSnapshotSkuResponseOutputWithContext(ctx context.Context) SnapshotSkuResponseOutput {
	return o
}

func (o SnapshotSkuResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SnapshotSkuResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o SnapshotSkuResponseOutput) Tier() pulumi.StringOutput {
	return o.ApplyT(func(v SnapshotSkuResponse) string { return v.Tier }).(pulumi.StringOutput)
}

type SnapshotSkuResponsePtrOutput struct{ *pulumi.OutputState }

func (SnapshotSkuResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SnapshotSkuResponse)(nil)).Elem()
}

func (o SnapshotSkuResponsePtrOutput) ToSnapshotSkuResponsePtrOutput() SnapshotSkuResponsePtrOutput {
	return o
}

func (o SnapshotSkuResponsePtrOutput) ToSnapshotSkuResponsePtrOutputWithContext(ctx context.Context) SnapshotSkuResponsePtrOutput {
	return o
}

func (o SnapshotSkuResponsePtrOutput) Elem() SnapshotSkuResponseOutput {
	return o.ApplyT(func(v *SnapshotSkuResponse) SnapshotSkuResponse {
		if v != nil {
			return *v
		}
		var ret SnapshotSkuResponse
		return ret
	}).(SnapshotSkuResponseOutput)
}

func (o SnapshotSkuResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SnapshotSkuResponse) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o SnapshotSkuResponsePtrOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SnapshotSkuResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Tier
	}).(pulumi.StringPtrOutput)
}

type SourceVault struct {
	Id *string `pulumi:"id"`
}





type SourceVaultInput interface {
	pulumi.Input

	ToSourceVaultOutput() SourceVaultOutput
	ToSourceVaultOutputWithContext(context.Context) SourceVaultOutput
}

type SourceVaultArgs struct {
	Id pulumi.StringPtrInput `pulumi:"id"`
}

func (SourceVaultArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SourceVault)(nil)).Elem()
}

func (i SourceVaultArgs) ToSourceVaultOutput() SourceVaultOutput {
	return i.ToSourceVaultOutputWithContext(context.Background())
}

func (i SourceVaultArgs) ToSourceVaultOutputWithContext(ctx context.Context) SourceVaultOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceVaultOutput)
}

func (i SourceVaultArgs) ToSourceVaultPtrOutput() SourceVaultPtrOutput {
	return i.ToSourceVaultPtrOutputWithContext(context.Background())
}

func (i SourceVaultArgs) ToSourceVaultPtrOutputWithContext(ctx context.Context) SourceVaultPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceVaultOutput).ToSourceVaultPtrOutputWithContext(ctx)
}









type SourceVaultPtrInput interface {
	pulumi.Input

	ToSourceVaultPtrOutput() SourceVaultPtrOutput
	ToSourceVaultPtrOutputWithContext(context.Context) SourceVaultPtrOutput
}

type sourceVaultPtrType SourceVaultArgs

func SourceVaultPtr(v *SourceVaultArgs) SourceVaultPtrInput {
	return (*sourceVaultPtrType)(v)
}

func (*sourceVaultPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SourceVault)(nil)).Elem()
}

func (i *sourceVaultPtrType) ToSourceVaultPtrOutput() SourceVaultPtrOutput {
	return i.ToSourceVaultPtrOutputWithContext(context.Background())
}

func (i *sourceVaultPtrType) ToSourceVaultPtrOutputWithContext(ctx context.Context) SourceVaultPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceVaultPtrOutput)
}

type SourceVaultOutput struct{ *pulumi.OutputState }

func (SourceVaultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SourceVault)(nil)).Elem()
}

func (o SourceVaultOutput) ToSourceVaultOutput() SourceVaultOutput {
	return o
}

func (o SourceVaultOutput) ToSourceVaultOutputWithContext(ctx context.Context) SourceVaultOutput {
	return o
}

func (o SourceVaultOutput) ToSourceVaultPtrOutput() SourceVaultPtrOutput {
	return o.ToSourceVaultPtrOutputWithContext(context.Background())
}

func (o SourceVaultOutput) ToSourceVaultPtrOutputWithContext(ctx context.Context) SourceVaultPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SourceVault) *SourceVault {
		return &v
	}).(SourceVaultPtrOutput)
}

func (o SourceVaultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SourceVault) *string { return v.Id }).(pulumi.StringPtrOutput)
}

type SourceVaultPtrOutput struct{ *pulumi.OutputState }

func (SourceVaultPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SourceVault)(nil)).Elem()
}

func (o SourceVaultPtrOutput) ToSourceVaultPtrOutput() SourceVaultPtrOutput {
	return o
}

func (o SourceVaultPtrOutput) ToSourceVaultPtrOutputWithContext(ctx context.Context) SourceVaultPtrOutput {
	return o
}

func (o SourceVaultPtrOutput) Elem() SourceVaultOutput {
	return o.ApplyT(func(v *SourceVault) SourceVault {
		if v != nil {
			return *v
		}
		var ret SourceVault
		return ret
	}).(SourceVaultOutput)
}

func (o SourceVaultPtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SourceVault) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

type SourceVaultResponse struct {
	Id *string `pulumi:"id"`
}

type SourceVaultResponseOutput struct{ *pulumi.OutputState }

func (SourceVaultResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SourceVaultResponse)(nil)).Elem()
}

func (o SourceVaultResponseOutput) ToSourceVaultResponseOutput() SourceVaultResponseOutput {
	return o
}

func (o SourceVaultResponseOutput) ToSourceVaultResponseOutputWithContext(ctx context.Context) SourceVaultResponseOutput {
	return o
}

func (o SourceVaultResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SourceVaultResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

type SourceVaultResponsePtrOutput struct{ *pulumi.OutputState }

func (SourceVaultResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SourceVaultResponse)(nil)).Elem()
}

func (o SourceVaultResponsePtrOutput) ToSourceVaultResponsePtrOutput() SourceVaultResponsePtrOutput {
	return o
}

func (o SourceVaultResponsePtrOutput) ToSourceVaultResponsePtrOutputWithContext(ctx context.Context) SourceVaultResponsePtrOutput {
	return o
}

func (o SourceVaultResponsePtrOutput) Elem() SourceVaultResponseOutput {
	return o.ApplyT(func(v *SourceVaultResponse) SourceVaultResponse {
		if v != nil {
			return *v
		}
		var ret SourceVaultResponse
		return ret
	}).(SourceVaultResponseOutput)
}

func (o SourceVaultResponsePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SourceVaultResponse) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

type SshConfiguration struct {
	PublicKeys []SshPublicKey `pulumi:"publicKeys"`
}





type SshConfigurationInput interface {
	pulumi.Input

	ToSshConfigurationOutput() SshConfigurationOutput
	ToSshConfigurationOutputWithContext(context.Context) SshConfigurationOutput
}

type SshConfigurationArgs struct {
	PublicKeys SshPublicKeyArrayInput `pulumi:"publicKeys"`
}

func (SshConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SshConfiguration)(nil)).Elem()
}

func (i SshConfigurationArgs) ToSshConfigurationOutput() SshConfigurationOutput {
	return i.ToSshConfigurationOutputWithContext(context.Background())
}

func (i SshConfigurationArgs) ToSshConfigurationOutputWithContext(ctx context.Context) SshConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SshConfigurationOutput)
}

func (i SshConfigurationArgs) ToSshConfigurationPtrOutput() SshConfigurationPtrOutput {
	return i.ToSshConfigurationPtrOutputWithContext(context.Background())
}

func (i SshConfigurationArgs) ToSshConfigurationPtrOutputWithContext(ctx context.Context) SshConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SshConfigurationOutput).ToSshConfigurationPtrOutputWithContext(ctx)
}









type SshConfigurationPtrInput interface {
	pulumi.Input

	ToSshConfigurationPtrOutput() SshConfigurationPtrOutput
	ToSshConfigurationPtrOutputWithContext(context.Context) SshConfigurationPtrOutput
}

type sshConfigurationPtrType SshConfigurationArgs

func SshConfigurationPtr(v *SshConfigurationArgs) SshConfigurationPtrInput {
	return (*sshConfigurationPtrType)(v)
}

func (*sshConfigurationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SshConfiguration)(nil)).Elem()
}

func (i *sshConfigurationPtrType) ToSshConfigurationPtrOutput() SshConfigurationPtrOutput {
	return i.ToSshConfigurationPtrOutputWithContext(context.Background())
}

func (i *sshConfigurationPtrType) ToSshConfigurationPtrOutputWithContext(ctx context.Context) SshConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SshConfigurationPtrOutput)
}

type SshConfigurationOutput struct{ *pulumi.OutputState }

func (SshConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SshConfiguration)(nil)).Elem()
}

func (o SshConfigurationOutput) ToSshConfigurationOutput() SshConfigurationOutput {
	return o
}

func (o SshConfigurationOutput) ToSshConfigurationOutputWithContext(ctx context.Context) SshConfigurationOutput {
	return o
}

func (o SshConfigurationOutput) ToSshConfigurationPtrOutput() SshConfigurationPtrOutput {
	return o.ToSshConfigurationPtrOutputWithContext(context.Background())
}

func (o SshConfigurationOutput) ToSshConfigurationPtrOutputWithContext(ctx context.Context) SshConfigurationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SshConfiguration) *SshConfiguration {
		return &v
	}).(SshConfigurationPtrOutput)
}

func (o SshConfigurationOutput) PublicKeys() SshPublicKeyArrayOutput {
	return o.ApplyT(func(v SshConfiguration) []SshPublicKey { return v.PublicKeys }).(SshPublicKeyArrayOutput)
}

type SshConfigurationPtrOutput struct{ *pulumi.OutputState }

func (SshConfigurationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SshConfiguration)(nil)).Elem()
}

func (o SshConfigurationPtrOutput) ToSshConfigurationPtrOutput() SshConfigurationPtrOutput {
	return o
}

func (o SshConfigurationPtrOutput) ToSshConfigurationPtrOutputWithContext(ctx context.Context) SshConfigurationPtrOutput {
	return o
}

func (o SshConfigurationPtrOutput) Elem() SshConfigurationOutput {
	return o.ApplyT(func(v *SshConfiguration) SshConfiguration {
		if v != nil {
			return *v
		}
		var ret SshConfiguration
		return ret
	}).(SshConfigurationOutput)
}

func (o SshConfigurationPtrOutput) PublicKeys() SshPublicKeyArrayOutput {
	return o.ApplyT(func(v *SshConfiguration) []SshPublicKey {
		if v == nil {
			return nil
		}
		return v.PublicKeys
	}).(SshPublicKeyArrayOutput)
}

type SshConfigurationResponse struct {
	PublicKeys []SshPublicKeyResponse `pulumi:"publicKeys"`
}

type SshConfigurationResponseOutput struct{ *pulumi.OutputState }

func (SshConfigurationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SshConfigurationResponse)(nil)).Elem()
}

func (o SshConfigurationResponseOutput) ToSshConfigurationResponseOutput() SshConfigurationResponseOutput {
	return o
}

func (o SshConfigurationResponseOutput) ToSshConfigurationResponseOutputWithContext(ctx context.Context) SshConfigurationResponseOutput {
	return o
}

func (o SshConfigurationResponseOutput) PublicKeys() SshPublicKeyResponseArrayOutput {
	return o.ApplyT(func(v SshConfigurationResponse) []SshPublicKeyResponse { return v.PublicKeys }).(SshPublicKeyResponseArrayOutput)
}

type SshConfigurationResponsePtrOutput struct{ *pulumi.OutputState }

func (SshConfigurationResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SshConfigurationResponse)(nil)).Elem()
}

func (o SshConfigurationResponsePtrOutput) ToSshConfigurationResponsePtrOutput() SshConfigurationResponsePtrOutput {
	return o
}

func (o SshConfigurationResponsePtrOutput) ToSshConfigurationResponsePtrOutputWithContext(ctx context.Context) SshConfigurationResponsePtrOutput {
	return o
}

func (o SshConfigurationResponsePtrOutput) Elem() SshConfigurationResponseOutput {
	return o.ApplyT(func(v *SshConfigurationResponse) SshConfigurationResponse {
		if v != nil {
			return *v
		}
		var ret SshConfigurationResponse
		return ret
	}).(SshConfigurationResponseOutput)
}

func (o SshConfigurationResponsePtrOutput) PublicKeys() SshPublicKeyResponseArrayOutput {
	return o.ApplyT(func(v *SshConfigurationResponse) []SshPublicKeyResponse {
		if v == nil {
			return nil
		}
		return v.PublicKeys
	}).(SshPublicKeyResponseArrayOutput)
}

type SshPublicKey struct {
	KeyData *string `pulumi:"keyData"`
	Path    *string `pulumi:"path"`
}





type SshPublicKeyInput interface {
	pulumi.Input

	ToSshPublicKeyOutput() SshPublicKeyOutput
	ToSshPublicKeyOutputWithContext(context.Context) SshPublicKeyOutput
}

type SshPublicKeyArgs struct {
	KeyData pulumi.StringPtrInput `pulumi:"keyData"`
	Path    pulumi.StringPtrInput `pulumi:"path"`
}

func (SshPublicKeyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SshPublicKey)(nil)).Elem()
}

func (i SshPublicKeyArgs) ToSshPublicKeyOutput() SshPublicKeyOutput {
	return i.ToSshPublicKeyOutputWithContext(context.Background())
}

func (i SshPublicKeyArgs) ToSshPublicKeyOutputWithContext(ctx context.Context) SshPublicKeyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SshPublicKeyOutput)
}





type SshPublicKeyArrayInput interface {
	pulumi.Input

	ToSshPublicKeyArrayOutput() SshPublicKeyArrayOutput
	ToSshPublicKeyArrayOutputWithContext(context.Context) SshPublicKeyArrayOutput
}

type SshPublicKeyArray []SshPublicKeyInput

func (SshPublicKeyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SshPublicKey)(nil)).Elem()
}

func (i SshPublicKeyArray) ToSshPublicKeyArrayOutput() SshPublicKeyArrayOutput {
	return i.ToSshPublicKeyArrayOutputWithContext(context.Background())
}

func (i SshPublicKeyArray) ToSshPublicKeyArrayOutputWithContext(ctx context.Context) SshPublicKeyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SshPublicKeyArrayOutput)
}

type SshPublicKeyOutput struct{ *pulumi.OutputState }

func (SshPublicKeyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SshPublicKey)(nil)).Elem()
}

func (o SshPublicKeyOutput) ToSshPublicKeyOutput() SshPublicKeyOutput {
	return o
}

func (o SshPublicKeyOutput) ToSshPublicKeyOutputWithContext(ctx context.Context) SshPublicKeyOutput {
	return o
}

func (o SshPublicKeyOutput) KeyData() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SshPublicKey) *string { return v.KeyData }).(pulumi.StringPtrOutput)
}

func (o SshPublicKeyOutput) Path() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SshPublicKey) *string { return v.Path }).(pulumi.StringPtrOutput)
}

type SshPublicKeyArrayOutput struct{ *pulumi.OutputState }

func (SshPublicKeyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SshPublicKey)(nil)).Elem()
}

func (o SshPublicKeyArrayOutput) ToSshPublicKeyArrayOutput() SshPublicKeyArrayOutput {
	return o
}

func (o SshPublicKeyArrayOutput) ToSshPublicKeyArrayOutputWithContext(ctx context.Context) SshPublicKeyArrayOutput {
	return o
}

func (o SshPublicKeyArrayOutput) Index(i pulumi.IntInput) SshPublicKeyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SshPublicKey {
		return vs[0].([]SshPublicKey)[vs[1].(int)]
	}).(SshPublicKeyOutput)
}

type SshPublicKeyResponse struct {
	KeyData *string `pulumi:"keyData"`
	Path    *string `pulumi:"path"`
}

type SshPublicKeyResponseOutput struct{ *pulumi.OutputState }

func (SshPublicKeyResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SshPublicKeyResponse)(nil)).Elem()
}

func (o SshPublicKeyResponseOutput) ToSshPublicKeyResponseOutput() SshPublicKeyResponseOutput {
	return o
}

func (o SshPublicKeyResponseOutput) ToSshPublicKeyResponseOutputWithContext(ctx context.Context) SshPublicKeyResponseOutput {
	return o
}

func (o SshPublicKeyResponseOutput) KeyData() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SshPublicKeyResponse) *string { return v.KeyData }).(pulumi.StringPtrOutput)
}

func (o SshPublicKeyResponseOutput) Path() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SshPublicKeyResponse) *string { return v.Path }).(pulumi.StringPtrOutput)
}

type SshPublicKeyResponseArrayOutput struct{ *pulumi.OutputState }

func (SshPublicKeyResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SshPublicKeyResponse)(nil)).Elem()
}

func (o SshPublicKeyResponseArrayOutput) ToSshPublicKeyResponseArrayOutput() SshPublicKeyResponseArrayOutput {
	return o
}

func (o SshPublicKeyResponseArrayOutput) ToSshPublicKeyResponseArrayOutputWithContext(ctx context.Context) SshPublicKeyResponseArrayOutput {
	return o
}

func (o SshPublicKeyResponseArrayOutput) Index(i pulumi.IntInput) SshPublicKeyResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SshPublicKeyResponse {
		return vs[0].([]SshPublicKeyResponse)[vs[1].(int)]
	}).(SshPublicKeyResponseOutput)
}

type StorageProfile struct {
	DataDisks      []DataDisk      `pulumi:"dataDisks"`
	ImageReference *ImageReference `pulumi:"imageReference"`
	OsDisk         *OSDisk         `pulumi:"osDisk"`
}





type StorageProfileInput interface {
	pulumi.Input

	ToStorageProfileOutput() StorageProfileOutput
	ToStorageProfileOutputWithContext(context.Context) StorageProfileOutput
}

type StorageProfileArgs struct {
	DataDisks      DataDiskArrayInput     `pulumi:"dataDisks"`
	ImageReference ImageReferencePtrInput `pulumi:"imageReference"`
	OsDisk         OSDiskPtrInput         `pulumi:"osDisk"`
}

func (StorageProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*StorageProfile)(nil)).Elem()
}

func (i StorageProfileArgs) ToStorageProfileOutput() StorageProfileOutput {
	return i.ToStorageProfileOutputWithContext(context.Background())
}

func (i StorageProfileArgs) ToStorageProfileOutputWithContext(ctx context.Context) StorageProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageProfileOutput)
}

func (i StorageProfileArgs) ToStorageProfilePtrOutput() StorageProfilePtrOutput {
	return i.ToStorageProfilePtrOutputWithContext(context.Background())
}

func (i StorageProfileArgs) ToStorageProfilePtrOutputWithContext(ctx context.Context) StorageProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageProfileOutput).ToStorageProfilePtrOutputWithContext(ctx)
}









type StorageProfilePtrInput interface {
	pulumi.Input

	ToStorageProfilePtrOutput() StorageProfilePtrOutput
	ToStorageProfilePtrOutputWithContext(context.Context) StorageProfilePtrOutput
}

type storageProfilePtrType StorageProfileArgs

func StorageProfilePtr(v *StorageProfileArgs) StorageProfilePtrInput {
	return (*storageProfilePtrType)(v)
}

func (*storageProfilePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**StorageProfile)(nil)).Elem()
}

func (i *storageProfilePtrType) ToStorageProfilePtrOutput() StorageProfilePtrOutput {
	return i.ToStorageProfilePtrOutputWithContext(context.Background())
}

func (i *storageProfilePtrType) ToStorageProfilePtrOutputWithContext(ctx context.Context) StorageProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageProfilePtrOutput)
}

type StorageProfileOutput struct{ *pulumi.OutputState }

func (StorageProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StorageProfile)(nil)).Elem()
}

func (o StorageProfileOutput) ToStorageProfileOutput() StorageProfileOutput {
	return o
}

func (o StorageProfileOutput) ToStorageProfileOutputWithContext(ctx context.Context) StorageProfileOutput {
	return o
}

func (o StorageProfileOutput) ToStorageProfilePtrOutput() StorageProfilePtrOutput {
	return o.ToStorageProfilePtrOutputWithContext(context.Background())
}

func (o StorageProfileOutput) ToStorageProfilePtrOutputWithContext(ctx context.Context) StorageProfilePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v StorageProfile) *StorageProfile {
		return &v
	}).(StorageProfilePtrOutput)
}

func (o StorageProfileOutput) DataDisks() DataDiskArrayOutput {
	return o.ApplyT(func(v StorageProfile) []DataDisk { return v.DataDisks }).(DataDiskArrayOutput)
}

func (o StorageProfileOutput) ImageReference() ImageReferencePtrOutput {
	return o.ApplyT(func(v StorageProfile) *ImageReference { return v.ImageReference }).(ImageReferencePtrOutput)
}

func (o StorageProfileOutput) OsDisk() OSDiskPtrOutput {
	return o.ApplyT(func(v StorageProfile) *OSDisk { return v.OsDisk }).(OSDiskPtrOutput)
}

type StorageProfilePtrOutput struct{ *pulumi.OutputState }

func (StorageProfilePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StorageProfile)(nil)).Elem()
}

func (o StorageProfilePtrOutput) ToStorageProfilePtrOutput() StorageProfilePtrOutput {
	return o
}

func (o StorageProfilePtrOutput) ToStorageProfilePtrOutputWithContext(ctx context.Context) StorageProfilePtrOutput {
	return o
}

func (o StorageProfilePtrOutput) Elem() StorageProfileOutput {
	return o.ApplyT(func(v *StorageProfile) StorageProfile {
		if v != nil {
			return *v
		}
		var ret StorageProfile
		return ret
	}).(StorageProfileOutput)
}

func (o StorageProfilePtrOutput) DataDisks() DataDiskArrayOutput {
	return o.ApplyT(func(v *StorageProfile) []DataDisk {
		if v == nil {
			return nil
		}
		return v.DataDisks
	}).(DataDiskArrayOutput)
}

func (o StorageProfilePtrOutput) ImageReference() ImageReferencePtrOutput {
	return o.ApplyT(func(v *StorageProfile) *ImageReference {
		if v == nil {
			return nil
		}
		return v.ImageReference
	}).(ImageReferencePtrOutput)
}

func (o StorageProfilePtrOutput) OsDisk() OSDiskPtrOutput {
	return o.ApplyT(func(v *StorageProfile) *OSDisk {
		if v == nil {
			return nil
		}
		return v.OsDisk
	}).(OSDiskPtrOutput)
}

type StorageProfileResponse struct {
	DataDisks      []DataDiskResponse      `pulumi:"dataDisks"`
	ImageReference *ImageReferenceResponse `pulumi:"imageReference"`
	OsDisk         *OSDiskResponse         `pulumi:"osDisk"`
}

type StorageProfileResponseOutput struct{ *pulumi.OutputState }

func (StorageProfileResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StorageProfileResponse)(nil)).Elem()
}

func (o StorageProfileResponseOutput) ToStorageProfileResponseOutput() StorageProfileResponseOutput {
	return o
}

func (o StorageProfileResponseOutput) ToStorageProfileResponseOutputWithContext(ctx context.Context) StorageProfileResponseOutput {
	return o
}

func (o StorageProfileResponseOutput) DataDisks() DataDiskResponseArrayOutput {
	return o.ApplyT(func(v StorageProfileResponse) []DataDiskResponse { return v.DataDisks }).(DataDiskResponseArrayOutput)
}

func (o StorageProfileResponseOutput) ImageReference() ImageReferenceResponsePtrOutput {
	return o.ApplyT(func(v StorageProfileResponse) *ImageReferenceResponse { return v.ImageReference }).(ImageReferenceResponsePtrOutput)
}

func (o StorageProfileResponseOutput) OsDisk() OSDiskResponsePtrOutput {
	return o.ApplyT(func(v StorageProfileResponse) *OSDiskResponse { return v.OsDisk }).(OSDiskResponsePtrOutput)
}

type StorageProfileResponsePtrOutput struct{ *pulumi.OutputState }

func (StorageProfileResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StorageProfileResponse)(nil)).Elem()
}

func (o StorageProfileResponsePtrOutput) ToStorageProfileResponsePtrOutput() StorageProfileResponsePtrOutput {
	return o
}

func (o StorageProfileResponsePtrOutput) ToStorageProfileResponsePtrOutputWithContext(ctx context.Context) StorageProfileResponsePtrOutput {
	return o
}

func (o StorageProfileResponsePtrOutput) Elem() StorageProfileResponseOutput {
	return o.ApplyT(func(v *StorageProfileResponse) StorageProfileResponse {
		if v != nil {
			return *v
		}
		var ret StorageProfileResponse
		return ret
	}).(StorageProfileResponseOutput)
}

func (o StorageProfileResponsePtrOutput) DataDisks() DataDiskResponseArrayOutput {
	return o.ApplyT(func(v *StorageProfileResponse) []DataDiskResponse {
		if v == nil {
			return nil
		}
		return v.DataDisks
	}).(DataDiskResponseArrayOutput)
}

func (o StorageProfileResponsePtrOutput) ImageReference() ImageReferenceResponsePtrOutput {
	return o.ApplyT(func(v *StorageProfileResponse) *ImageReferenceResponse {
		if v == nil {
			return nil
		}
		return v.ImageReference
	}).(ImageReferenceResponsePtrOutput)
}

func (o StorageProfileResponsePtrOutput) OsDisk() OSDiskResponsePtrOutput {
	return o.ApplyT(func(v *StorageProfileResponse) *OSDiskResponse {
		if v == nil {
			return nil
		}
		return v.OsDisk
	}).(OSDiskResponsePtrOutput)
}

type SubResource struct {
	Id *string `pulumi:"id"`
}





type SubResourceInput interface {
	pulumi.Input

	ToSubResourceOutput() SubResourceOutput
	ToSubResourceOutputWithContext(context.Context) SubResourceOutput
}

type SubResourceArgs struct {
	Id pulumi.StringPtrInput `pulumi:"id"`
}

func (SubResourceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SubResource)(nil)).Elem()
}

func (i SubResourceArgs) ToSubResourceOutput() SubResourceOutput {
	return i.ToSubResourceOutputWithContext(context.Background())
}

func (i SubResourceArgs) ToSubResourceOutputWithContext(ctx context.Context) SubResourceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubResourceOutput)
}

func (i SubResourceArgs) ToSubResourcePtrOutput() SubResourcePtrOutput {
	return i.ToSubResourcePtrOutputWithContext(context.Background())
}

func (i SubResourceArgs) ToSubResourcePtrOutputWithContext(ctx context.Context) SubResourcePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubResourceOutput).ToSubResourcePtrOutputWithContext(ctx)
}









type SubResourcePtrInput interface {
	pulumi.Input

	ToSubResourcePtrOutput() SubResourcePtrOutput
	ToSubResourcePtrOutputWithContext(context.Context) SubResourcePtrOutput
}

type subResourcePtrType SubResourceArgs

func SubResourcePtr(v *SubResourceArgs) SubResourcePtrInput {
	return (*subResourcePtrType)(v)
}

func (*subResourcePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SubResource)(nil)).Elem()
}

func (i *subResourcePtrType) ToSubResourcePtrOutput() SubResourcePtrOutput {
	return i.ToSubResourcePtrOutputWithContext(context.Background())
}

func (i *subResourcePtrType) ToSubResourcePtrOutputWithContext(ctx context.Context) SubResourcePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubResourcePtrOutput)
}





type SubResourceArrayInput interface {
	pulumi.Input

	ToSubResourceArrayOutput() SubResourceArrayOutput
	ToSubResourceArrayOutputWithContext(context.Context) SubResourceArrayOutput
}

type SubResourceArray []SubResourceInput

func (SubResourceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SubResource)(nil)).Elem()
}

func (i SubResourceArray) ToSubResourceArrayOutput() SubResourceArrayOutput {
	return i.ToSubResourceArrayOutputWithContext(context.Background())
}

func (i SubResourceArray) ToSubResourceArrayOutputWithContext(ctx context.Context) SubResourceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubResourceArrayOutput)
}

type SubResourceOutput struct{ *pulumi.OutputState }

func (SubResourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SubResource)(nil)).Elem()
}

func (o SubResourceOutput) ToSubResourceOutput() SubResourceOutput {
	return o
}

func (o SubResourceOutput) ToSubResourceOutputWithContext(ctx context.Context) SubResourceOutput {
	return o
}

func (o SubResourceOutput) ToSubResourcePtrOutput() SubResourcePtrOutput {
	return o.ToSubResourcePtrOutputWithContext(context.Background())
}

func (o SubResourceOutput) ToSubResourcePtrOutputWithContext(ctx context.Context) SubResourcePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SubResource) *SubResource {
		return &v
	}).(SubResourcePtrOutput)
}

func (o SubResourceOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SubResource) *string { return v.Id }).(pulumi.StringPtrOutput)
}

type SubResourcePtrOutput struct{ *pulumi.OutputState }

func (SubResourcePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SubResource)(nil)).Elem()
}

func (o SubResourcePtrOutput) ToSubResourcePtrOutput() SubResourcePtrOutput {
	return o
}

func (o SubResourcePtrOutput) ToSubResourcePtrOutputWithContext(ctx context.Context) SubResourcePtrOutput {
	return o
}

func (o SubResourcePtrOutput) Elem() SubResourceOutput {
	return o.ApplyT(func(v *SubResource) SubResource {
		if v != nil {
			return *v
		}
		var ret SubResource
		return ret
	}).(SubResourceOutput)
}

func (o SubResourcePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SubResource) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

type SubResourceArrayOutput struct{ *pulumi.OutputState }

func (SubResourceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SubResource)(nil)).Elem()
}

func (o SubResourceArrayOutput) ToSubResourceArrayOutput() SubResourceArrayOutput {
	return o
}

func (o SubResourceArrayOutput) ToSubResourceArrayOutputWithContext(ctx context.Context) SubResourceArrayOutput {
	return o
}

func (o SubResourceArrayOutput) Index(i pulumi.IntInput) SubResourceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SubResource {
		return vs[0].([]SubResource)[vs[1].(int)]
	}).(SubResourceOutput)
}

type SubResourceResponse struct {
	Id *string `pulumi:"id"`
}

type SubResourceResponseOutput struct{ *pulumi.OutputState }

func (SubResourceResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SubResourceResponse)(nil)).Elem()
}

func (o SubResourceResponseOutput) ToSubResourceResponseOutput() SubResourceResponseOutput {
	return o
}

func (o SubResourceResponseOutput) ToSubResourceResponseOutputWithContext(ctx context.Context) SubResourceResponseOutput {
	return o
}

func (o SubResourceResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SubResourceResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

type SubResourceResponsePtrOutput struct{ *pulumi.OutputState }

func (SubResourceResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SubResourceResponse)(nil)).Elem()
}

func (o SubResourceResponsePtrOutput) ToSubResourceResponsePtrOutput() SubResourceResponsePtrOutput {
	return o
}

func (o SubResourceResponsePtrOutput) ToSubResourceResponsePtrOutputWithContext(ctx context.Context) SubResourceResponsePtrOutput {
	return o
}

func (o SubResourceResponsePtrOutput) Elem() SubResourceResponseOutput {
	return o.ApplyT(func(v *SubResourceResponse) SubResourceResponse {
		if v != nil {
			return *v
		}
		var ret SubResourceResponse
		return ret
	}).(SubResourceResponseOutput)
}

func (o SubResourceResponsePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SubResourceResponse) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

type SubResourceResponseArrayOutput struct{ *pulumi.OutputState }

func (SubResourceResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SubResourceResponse)(nil)).Elem()
}

func (o SubResourceResponseArrayOutput) ToSubResourceResponseArrayOutput() SubResourceResponseArrayOutput {
	return o
}

func (o SubResourceResponseArrayOutput) ToSubResourceResponseArrayOutputWithContext(ctx context.Context) SubResourceResponseArrayOutput {
	return o
}

func (o SubResourceResponseArrayOutput) Index(i pulumi.IntInput) SubResourceResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SubResourceResponse {
		return vs[0].([]SubResourceResponse)[vs[1].(int)]
	}).(SubResourceResponseOutput)
}

type UpgradePolicy struct {
	AutoOSUpgradePolicy  *AutoOSUpgradePolicy  `pulumi:"autoOSUpgradePolicy"`
	AutomaticOSUpgrade   *bool                 `pulumi:"automaticOSUpgrade"`
	Mode                 *UpgradeMode          `pulumi:"mode"`
	RollingUpgradePolicy *RollingUpgradePolicy `pulumi:"rollingUpgradePolicy"`
}





type UpgradePolicyInput interface {
	pulumi.Input

	ToUpgradePolicyOutput() UpgradePolicyOutput
	ToUpgradePolicyOutputWithContext(context.Context) UpgradePolicyOutput
}

type UpgradePolicyArgs struct {
	AutoOSUpgradePolicy  AutoOSUpgradePolicyPtrInput  `pulumi:"autoOSUpgradePolicy"`
	AutomaticOSUpgrade   pulumi.BoolPtrInput          `pulumi:"automaticOSUpgrade"`
	Mode                 UpgradeModePtrInput          `pulumi:"mode"`
	RollingUpgradePolicy RollingUpgradePolicyPtrInput `pulumi:"rollingUpgradePolicy"`
}

func (UpgradePolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*UpgradePolicy)(nil)).Elem()
}

func (i UpgradePolicyArgs) ToUpgradePolicyOutput() UpgradePolicyOutput {
	return i.ToUpgradePolicyOutputWithContext(context.Background())
}

func (i UpgradePolicyArgs) ToUpgradePolicyOutputWithContext(ctx context.Context) UpgradePolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UpgradePolicyOutput)
}

func (i UpgradePolicyArgs) ToUpgradePolicyPtrOutput() UpgradePolicyPtrOutput {
	return i.ToUpgradePolicyPtrOutputWithContext(context.Background())
}

func (i UpgradePolicyArgs) ToUpgradePolicyPtrOutputWithContext(ctx context.Context) UpgradePolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UpgradePolicyOutput).ToUpgradePolicyPtrOutputWithContext(ctx)
}









type UpgradePolicyPtrInput interface {
	pulumi.Input

	ToUpgradePolicyPtrOutput() UpgradePolicyPtrOutput
	ToUpgradePolicyPtrOutputWithContext(context.Context) UpgradePolicyPtrOutput
}

type upgradePolicyPtrType UpgradePolicyArgs

func UpgradePolicyPtr(v *UpgradePolicyArgs) UpgradePolicyPtrInput {
	return (*upgradePolicyPtrType)(v)
}

func (*upgradePolicyPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**UpgradePolicy)(nil)).Elem()
}

func (i *upgradePolicyPtrType) ToUpgradePolicyPtrOutput() UpgradePolicyPtrOutput {
	return i.ToUpgradePolicyPtrOutputWithContext(context.Background())
}

func (i *upgradePolicyPtrType) ToUpgradePolicyPtrOutputWithContext(ctx context.Context) UpgradePolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UpgradePolicyPtrOutput)
}

type UpgradePolicyOutput struct{ *pulumi.OutputState }

func (UpgradePolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UpgradePolicy)(nil)).Elem()
}

func (o UpgradePolicyOutput) ToUpgradePolicyOutput() UpgradePolicyOutput {
	return o
}

func (o UpgradePolicyOutput) ToUpgradePolicyOutputWithContext(ctx context.Context) UpgradePolicyOutput {
	return o
}

func (o UpgradePolicyOutput) ToUpgradePolicyPtrOutput() UpgradePolicyPtrOutput {
	return o.ToUpgradePolicyPtrOutputWithContext(context.Background())
}

func (o UpgradePolicyOutput) ToUpgradePolicyPtrOutputWithContext(ctx context.Context) UpgradePolicyPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v UpgradePolicy) *UpgradePolicy {
		return &v
	}).(UpgradePolicyPtrOutput)
}

func (o UpgradePolicyOutput) AutoOSUpgradePolicy() AutoOSUpgradePolicyPtrOutput {
	return o.ApplyT(func(v UpgradePolicy) *AutoOSUpgradePolicy { return v.AutoOSUpgradePolicy }).(AutoOSUpgradePolicyPtrOutput)
}

func (o UpgradePolicyOutput) AutomaticOSUpgrade() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v UpgradePolicy) *bool { return v.AutomaticOSUpgrade }).(pulumi.BoolPtrOutput)
}

func (o UpgradePolicyOutput) Mode() UpgradeModePtrOutput {
	return o.ApplyT(func(v UpgradePolicy) *UpgradeMode { return v.Mode }).(UpgradeModePtrOutput)
}

func (o UpgradePolicyOutput) RollingUpgradePolicy() RollingUpgradePolicyPtrOutput {
	return o.ApplyT(func(v UpgradePolicy) *RollingUpgradePolicy { return v.RollingUpgradePolicy }).(RollingUpgradePolicyPtrOutput)
}

type UpgradePolicyPtrOutput struct{ *pulumi.OutputState }

func (UpgradePolicyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**UpgradePolicy)(nil)).Elem()
}

func (o UpgradePolicyPtrOutput) ToUpgradePolicyPtrOutput() UpgradePolicyPtrOutput {
	return o
}

func (o UpgradePolicyPtrOutput) ToUpgradePolicyPtrOutputWithContext(ctx context.Context) UpgradePolicyPtrOutput {
	return o
}

func (o UpgradePolicyPtrOutput) Elem() UpgradePolicyOutput {
	return o.ApplyT(func(v *UpgradePolicy) UpgradePolicy {
		if v != nil {
			return *v
		}
		var ret UpgradePolicy
		return ret
	}).(UpgradePolicyOutput)
}

func (o UpgradePolicyPtrOutput) AutoOSUpgradePolicy() AutoOSUpgradePolicyPtrOutput {
	return o.ApplyT(func(v *UpgradePolicy) *AutoOSUpgradePolicy {
		if v == nil {
			return nil
		}
		return v.AutoOSUpgradePolicy
	}).(AutoOSUpgradePolicyPtrOutput)
}

func (o UpgradePolicyPtrOutput) AutomaticOSUpgrade() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *UpgradePolicy) *bool {
		if v == nil {
			return nil
		}
		return v.AutomaticOSUpgrade
	}).(pulumi.BoolPtrOutput)
}

func (o UpgradePolicyPtrOutput) Mode() UpgradeModePtrOutput {
	return o.ApplyT(func(v *UpgradePolicy) *UpgradeMode {
		if v == nil {
			return nil
		}
		return v.Mode
	}).(UpgradeModePtrOutput)
}

func (o UpgradePolicyPtrOutput) RollingUpgradePolicy() RollingUpgradePolicyPtrOutput {
	return o.ApplyT(func(v *UpgradePolicy) *RollingUpgradePolicy {
		if v == nil {
			return nil
		}
		return v.RollingUpgradePolicy
	}).(RollingUpgradePolicyPtrOutput)
}

type UpgradePolicyResponse struct {
	AutoOSUpgradePolicy  *AutoOSUpgradePolicyResponse  `pulumi:"autoOSUpgradePolicy"`
	AutomaticOSUpgrade   *bool                         `pulumi:"automaticOSUpgrade"`
	Mode                 *string                       `pulumi:"mode"`
	RollingUpgradePolicy *RollingUpgradePolicyResponse `pulumi:"rollingUpgradePolicy"`
}

type UpgradePolicyResponseOutput struct{ *pulumi.OutputState }

func (UpgradePolicyResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UpgradePolicyResponse)(nil)).Elem()
}

func (o UpgradePolicyResponseOutput) ToUpgradePolicyResponseOutput() UpgradePolicyResponseOutput {
	return o
}

func (o UpgradePolicyResponseOutput) ToUpgradePolicyResponseOutputWithContext(ctx context.Context) UpgradePolicyResponseOutput {
	return o
}

func (o UpgradePolicyResponseOutput) AutoOSUpgradePolicy() AutoOSUpgradePolicyResponsePtrOutput {
	return o.ApplyT(func(v UpgradePolicyResponse) *AutoOSUpgradePolicyResponse { return v.AutoOSUpgradePolicy }).(AutoOSUpgradePolicyResponsePtrOutput)
}

func (o UpgradePolicyResponseOutput) AutomaticOSUpgrade() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v UpgradePolicyResponse) *bool { return v.AutomaticOSUpgrade }).(pulumi.BoolPtrOutput)
}

func (o UpgradePolicyResponseOutput) Mode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v UpgradePolicyResponse) *string { return v.Mode }).(pulumi.StringPtrOutput)
}

func (o UpgradePolicyResponseOutput) RollingUpgradePolicy() RollingUpgradePolicyResponsePtrOutput {
	return o.ApplyT(func(v UpgradePolicyResponse) *RollingUpgradePolicyResponse { return v.RollingUpgradePolicy }).(RollingUpgradePolicyResponsePtrOutput)
}

type UpgradePolicyResponsePtrOutput struct{ *pulumi.OutputState }

func (UpgradePolicyResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**UpgradePolicyResponse)(nil)).Elem()
}

func (o UpgradePolicyResponsePtrOutput) ToUpgradePolicyResponsePtrOutput() UpgradePolicyResponsePtrOutput {
	return o
}

func (o UpgradePolicyResponsePtrOutput) ToUpgradePolicyResponsePtrOutputWithContext(ctx context.Context) UpgradePolicyResponsePtrOutput {
	return o
}

func (o UpgradePolicyResponsePtrOutput) Elem() UpgradePolicyResponseOutput {
	return o.ApplyT(func(v *UpgradePolicyResponse) UpgradePolicyResponse {
		if v != nil {
			return *v
		}
		var ret UpgradePolicyResponse
		return ret
	}).(UpgradePolicyResponseOutput)
}

func (o UpgradePolicyResponsePtrOutput) AutoOSUpgradePolicy() AutoOSUpgradePolicyResponsePtrOutput {
	return o.ApplyT(func(v *UpgradePolicyResponse) *AutoOSUpgradePolicyResponse {
		if v == nil {
			return nil
		}
		return v.AutoOSUpgradePolicy
	}).(AutoOSUpgradePolicyResponsePtrOutput)
}

func (o UpgradePolicyResponsePtrOutput) AutomaticOSUpgrade() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *UpgradePolicyResponse) *bool {
		if v == nil {
			return nil
		}
		return v.AutomaticOSUpgrade
	}).(pulumi.BoolPtrOutput)
}

func (o UpgradePolicyResponsePtrOutput) Mode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UpgradePolicyResponse) *string {
		if v == nil {
			return nil
		}
		return v.Mode
	}).(pulumi.StringPtrOutput)
}

func (o UpgradePolicyResponsePtrOutput) RollingUpgradePolicy() RollingUpgradePolicyResponsePtrOutput {
	return o.ApplyT(func(v *UpgradePolicyResponse) *RollingUpgradePolicyResponse {
		if v == nil {
			return nil
		}
		return v.RollingUpgradePolicy
	}).(RollingUpgradePolicyResponsePtrOutput)
}

type VaultCertificate struct {
	CertificateStore *string `pulumi:"certificateStore"`
	CertificateUrl   *string `pulumi:"certificateUrl"`
}





type VaultCertificateInput interface {
	pulumi.Input

	ToVaultCertificateOutput() VaultCertificateOutput
	ToVaultCertificateOutputWithContext(context.Context) VaultCertificateOutput
}

type VaultCertificateArgs struct {
	CertificateStore pulumi.StringPtrInput `pulumi:"certificateStore"`
	CertificateUrl   pulumi.StringPtrInput `pulumi:"certificateUrl"`
}

func (VaultCertificateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VaultCertificate)(nil)).Elem()
}

func (i VaultCertificateArgs) ToVaultCertificateOutput() VaultCertificateOutput {
	return i.ToVaultCertificateOutputWithContext(context.Background())
}

func (i VaultCertificateArgs) ToVaultCertificateOutputWithContext(ctx context.Context) VaultCertificateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VaultCertificateOutput)
}





type VaultCertificateArrayInput interface {
	pulumi.Input

	ToVaultCertificateArrayOutput() VaultCertificateArrayOutput
	ToVaultCertificateArrayOutputWithContext(context.Context) VaultCertificateArrayOutput
}

type VaultCertificateArray []VaultCertificateInput

func (VaultCertificateArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VaultCertificate)(nil)).Elem()
}

func (i VaultCertificateArray) ToVaultCertificateArrayOutput() VaultCertificateArrayOutput {
	return i.ToVaultCertificateArrayOutputWithContext(context.Background())
}

func (i VaultCertificateArray) ToVaultCertificateArrayOutputWithContext(ctx context.Context) VaultCertificateArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VaultCertificateArrayOutput)
}

type VaultCertificateOutput struct{ *pulumi.OutputState }

func (VaultCertificateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VaultCertificate)(nil)).Elem()
}

func (o VaultCertificateOutput) ToVaultCertificateOutput() VaultCertificateOutput {
	return o
}

func (o VaultCertificateOutput) ToVaultCertificateOutputWithContext(ctx context.Context) VaultCertificateOutput {
	return o
}

func (o VaultCertificateOutput) CertificateStore() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VaultCertificate) *string { return v.CertificateStore }).(pulumi.StringPtrOutput)
}

func (o VaultCertificateOutput) CertificateUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VaultCertificate) *string { return v.CertificateUrl }).(pulumi.StringPtrOutput)
}

type VaultCertificateArrayOutput struct{ *pulumi.OutputState }

func (VaultCertificateArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VaultCertificate)(nil)).Elem()
}

func (o VaultCertificateArrayOutput) ToVaultCertificateArrayOutput() VaultCertificateArrayOutput {
	return o
}

func (o VaultCertificateArrayOutput) ToVaultCertificateArrayOutputWithContext(ctx context.Context) VaultCertificateArrayOutput {
	return o
}

func (o VaultCertificateArrayOutput) Index(i pulumi.IntInput) VaultCertificateOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VaultCertificate {
		return vs[0].([]VaultCertificate)[vs[1].(int)]
	}).(VaultCertificateOutput)
}

type VaultCertificateResponse struct {
	CertificateStore *string `pulumi:"certificateStore"`
	CertificateUrl   *string `pulumi:"certificateUrl"`
}

type VaultCertificateResponseOutput struct{ *pulumi.OutputState }

func (VaultCertificateResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VaultCertificateResponse)(nil)).Elem()
}

func (o VaultCertificateResponseOutput) ToVaultCertificateResponseOutput() VaultCertificateResponseOutput {
	return o
}

func (o VaultCertificateResponseOutput) ToVaultCertificateResponseOutputWithContext(ctx context.Context) VaultCertificateResponseOutput {
	return o
}

func (o VaultCertificateResponseOutput) CertificateStore() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VaultCertificateResponse) *string { return v.CertificateStore }).(pulumi.StringPtrOutput)
}

func (o VaultCertificateResponseOutput) CertificateUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VaultCertificateResponse) *string { return v.CertificateUrl }).(pulumi.StringPtrOutput)
}

type VaultCertificateResponseArrayOutput struct{ *pulumi.OutputState }

func (VaultCertificateResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VaultCertificateResponse)(nil)).Elem()
}

func (o VaultCertificateResponseArrayOutput) ToVaultCertificateResponseArrayOutput() VaultCertificateResponseArrayOutput {
	return o
}

func (o VaultCertificateResponseArrayOutput) ToVaultCertificateResponseArrayOutputWithContext(ctx context.Context) VaultCertificateResponseArrayOutput {
	return o
}

func (o VaultCertificateResponseArrayOutput) Index(i pulumi.IntInput) VaultCertificateResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VaultCertificateResponse {
		return vs[0].([]VaultCertificateResponse)[vs[1].(int)]
	}).(VaultCertificateResponseOutput)
}

type VaultSecretGroup struct {
	SourceVault       *SubResource       `pulumi:"sourceVault"`
	VaultCertificates []VaultCertificate `pulumi:"vaultCertificates"`
}





type VaultSecretGroupInput interface {
	pulumi.Input

	ToVaultSecretGroupOutput() VaultSecretGroupOutput
	ToVaultSecretGroupOutputWithContext(context.Context) VaultSecretGroupOutput
}

type VaultSecretGroupArgs struct {
	SourceVault       SubResourcePtrInput        `pulumi:"sourceVault"`
	VaultCertificates VaultCertificateArrayInput `pulumi:"vaultCertificates"`
}

func (VaultSecretGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VaultSecretGroup)(nil)).Elem()
}

func (i VaultSecretGroupArgs) ToVaultSecretGroupOutput() VaultSecretGroupOutput {
	return i.ToVaultSecretGroupOutputWithContext(context.Background())
}

func (i VaultSecretGroupArgs) ToVaultSecretGroupOutputWithContext(ctx context.Context) VaultSecretGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VaultSecretGroupOutput)
}





type VaultSecretGroupArrayInput interface {
	pulumi.Input

	ToVaultSecretGroupArrayOutput() VaultSecretGroupArrayOutput
	ToVaultSecretGroupArrayOutputWithContext(context.Context) VaultSecretGroupArrayOutput
}

type VaultSecretGroupArray []VaultSecretGroupInput

func (VaultSecretGroupArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VaultSecretGroup)(nil)).Elem()
}

func (i VaultSecretGroupArray) ToVaultSecretGroupArrayOutput() VaultSecretGroupArrayOutput {
	return i.ToVaultSecretGroupArrayOutputWithContext(context.Background())
}

func (i VaultSecretGroupArray) ToVaultSecretGroupArrayOutputWithContext(ctx context.Context) VaultSecretGroupArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VaultSecretGroupArrayOutput)
}

type VaultSecretGroupOutput struct{ *pulumi.OutputState }

func (VaultSecretGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VaultSecretGroup)(nil)).Elem()
}

func (o VaultSecretGroupOutput) ToVaultSecretGroupOutput() VaultSecretGroupOutput {
	return o
}

func (o VaultSecretGroupOutput) ToVaultSecretGroupOutputWithContext(ctx context.Context) VaultSecretGroupOutput {
	return o
}

func (o VaultSecretGroupOutput) SourceVault() SubResourcePtrOutput {
	return o.ApplyT(func(v VaultSecretGroup) *SubResource { return v.SourceVault }).(SubResourcePtrOutput)
}

func (o VaultSecretGroupOutput) VaultCertificates() VaultCertificateArrayOutput {
	return o.ApplyT(func(v VaultSecretGroup) []VaultCertificate { return v.VaultCertificates }).(VaultCertificateArrayOutput)
}

type VaultSecretGroupArrayOutput struct{ *pulumi.OutputState }

func (VaultSecretGroupArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VaultSecretGroup)(nil)).Elem()
}

func (o VaultSecretGroupArrayOutput) ToVaultSecretGroupArrayOutput() VaultSecretGroupArrayOutput {
	return o
}

func (o VaultSecretGroupArrayOutput) ToVaultSecretGroupArrayOutputWithContext(ctx context.Context) VaultSecretGroupArrayOutput {
	return o
}

func (o VaultSecretGroupArrayOutput) Index(i pulumi.IntInput) VaultSecretGroupOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VaultSecretGroup {
		return vs[0].([]VaultSecretGroup)[vs[1].(int)]
	}).(VaultSecretGroupOutput)
}

type VaultSecretGroupResponse struct {
	SourceVault       *SubResourceResponse       `pulumi:"sourceVault"`
	VaultCertificates []VaultCertificateResponse `pulumi:"vaultCertificates"`
}

type VaultSecretGroupResponseOutput struct{ *pulumi.OutputState }

func (VaultSecretGroupResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VaultSecretGroupResponse)(nil)).Elem()
}

func (o VaultSecretGroupResponseOutput) ToVaultSecretGroupResponseOutput() VaultSecretGroupResponseOutput {
	return o
}

func (o VaultSecretGroupResponseOutput) ToVaultSecretGroupResponseOutputWithContext(ctx context.Context) VaultSecretGroupResponseOutput {
	return o
}

func (o VaultSecretGroupResponseOutput) SourceVault() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v VaultSecretGroupResponse) *SubResourceResponse { return v.SourceVault }).(SubResourceResponsePtrOutput)
}

func (o VaultSecretGroupResponseOutput) VaultCertificates() VaultCertificateResponseArrayOutput {
	return o.ApplyT(func(v VaultSecretGroupResponse) []VaultCertificateResponse { return v.VaultCertificates }).(VaultCertificateResponseArrayOutput)
}

type VaultSecretGroupResponseArrayOutput struct{ *pulumi.OutputState }

func (VaultSecretGroupResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VaultSecretGroupResponse)(nil)).Elem()
}

func (o VaultSecretGroupResponseArrayOutput) ToVaultSecretGroupResponseArrayOutput() VaultSecretGroupResponseArrayOutput {
	return o
}

func (o VaultSecretGroupResponseArrayOutput) ToVaultSecretGroupResponseArrayOutputWithContext(ctx context.Context) VaultSecretGroupResponseArrayOutput {
	return o
}

func (o VaultSecretGroupResponseArrayOutput) Index(i pulumi.IntInput) VaultSecretGroupResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VaultSecretGroupResponse {
		return vs[0].([]VaultSecretGroupResponse)[vs[1].(int)]
	}).(VaultSecretGroupResponseOutput)
}

type VirtualHardDisk struct {
	Uri *string `pulumi:"uri"`
}





type VirtualHardDiskInput interface {
	pulumi.Input

	ToVirtualHardDiskOutput() VirtualHardDiskOutput
	ToVirtualHardDiskOutputWithContext(context.Context) VirtualHardDiskOutput
}

type VirtualHardDiskArgs struct {
	Uri pulumi.StringPtrInput `pulumi:"uri"`
}

func (VirtualHardDiskArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualHardDisk)(nil)).Elem()
}

func (i VirtualHardDiskArgs) ToVirtualHardDiskOutput() VirtualHardDiskOutput {
	return i.ToVirtualHardDiskOutputWithContext(context.Background())
}

func (i VirtualHardDiskArgs) ToVirtualHardDiskOutputWithContext(ctx context.Context) VirtualHardDiskOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualHardDiskOutput)
}

func (i VirtualHardDiskArgs) ToVirtualHardDiskPtrOutput() VirtualHardDiskPtrOutput {
	return i.ToVirtualHardDiskPtrOutputWithContext(context.Background())
}

func (i VirtualHardDiskArgs) ToVirtualHardDiskPtrOutputWithContext(ctx context.Context) VirtualHardDiskPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualHardDiskOutput).ToVirtualHardDiskPtrOutputWithContext(ctx)
}









type VirtualHardDiskPtrInput interface {
	pulumi.Input

	ToVirtualHardDiskPtrOutput() VirtualHardDiskPtrOutput
	ToVirtualHardDiskPtrOutputWithContext(context.Context) VirtualHardDiskPtrOutput
}

type virtualHardDiskPtrType VirtualHardDiskArgs

func VirtualHardDiskPtr(v *VirtualHardDiskArgs) VirtualHardDiskPtrInput {
	return (*virtualHardDiskPtrType)(v)
}

func (*virtualHardDiskPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualHardDisk)(nil)).Elem()
}

func (i *virtualHardDiskPtrType) ToVirtualHardDiskPtrOutput() VirtualHardDiskPtrOutput {
	return i.ToVirtualHardDiskPtrOutputWithContext(context.Background())
}

func (i *virtualHardDiskPtrType) ToVirtualHardDiskPtrOutputWithContext(ctx context.Context) VirtualHardDiskPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualHardDiskPtrOutput)
}

type VirtualHardDiskOutput struct{ *pulumi.OutputState }

func (VirtualHardDiskOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualHardDisk)(nil)).Elem()
}

func (o VirtualHardDiskOutput) ToVirtualHardDiskOutput() VirtualHardDiskOutput {
	return o
}

func (o VirtualHardDiskOutput) ToVirtualHardDiskOutputWithContext(ctx context.Context) VirtualHardDiskOutput {
	return o
}

func (o VirtualHardDiskOutput) ToVirtualHardDiskPtrOutput() VirtualHardDiskPtrOutput {
	return o.ToVirtualHardDiskPtrOutputWithContext(context.Background())
}

func (o VirtualHardDiskOutput) ToVirtualHardDiskPtrOutputWithContext(ctx context.Context) VirtualHardDiskPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v VirtualHardDisk) *VirtualHardDisk {
		return &v
	}).(VirtualHardDiskPtrOutput)
}

func (o VirtualHardDiskOutput) Uri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualHardDisk) *string { return v.Uri }).(pulumi.StringPtrOutput)
}

type VirtualHardDiskPtrOutput struct{ *pulumi.OutputState }

func (VirtualHardDiskPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualHardDisk)(nil)).Elem()
}

func (o VirtualHardDiskPtrOutput) ToVirtualHardDiskPtrOutput() VirtualHardDiskPtrOutput {
	return o
}

func (o VirtualHardDiskPtrOutput) ToVirtualHardDiskPtrOutputWithContext(ctx context.Context) VirtualHardDiskPtrOutput {
	return o
}

func (o VirtualHardDiskPtrOutput) Elem() VirtualHardDiskOutput {
	return o.ApplyT(func(v *VirtualHardDisk) VirtualHardDisk {
		if v != nil {
			return *v
		}
		var ret VirtualHardDisk
		return ret
	}).(VirtualHardDiskOutput)
}

func (o VirtualHardDiskPtrOutput) Uri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualHardDisk) *string {
		if v == nil {
			return nil
		}
		return v.Uri
	}).(pulumi.StringPtrOutput)
}

type VirtualHardDiskResponse struct {
	Uri *string `pulumi:"uri"`
}

type VirtualHardDiskResponseOutput struct{ *pulumi.OutputState }

func (VirtualHardDiskResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualHardDiskResponse)(nil)).Elem()
}

func (o VirtualHardDiskResponseOutput) ToVirtualHardDiskResponseOutput() VirtualHardDiskResponseOutput {
	return o
}

func (o VirtualHardDiskResponseOutput) ToVirtualHardDiskResponseOutputWithContext(ctx context.Context) VirtualHardDiskResponseOutput {
	return o
}

func (o VirtualHardDiskResponseOutput) Uri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualHardDiskResponse) *string { return v.Uri }).(pulumi.StringPtrOutput)
}

type VirtualHardDiskResponsePtrOutput struct{ *pulumi.OutputState }

func (VirtualHardDiskResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualHardDiskResponse)(nil)).Elem()
}

func (o VirtualHardDiskResponsePtrOutput) ToVirtualHardDiskResponsePtrOutput() VirtualHardDiskResponsePtrOutput {
	return o
}

func (o VirtualHardDiskResponsePtrOutput) ToVirtualHardDiskResponsePtrOutputWithContext(ctx context.Context) VirtualHardDiskResponsePtrOutput {
	return o
}

func (o VirtualHardDiskResponsePtrOutput) Elem() VirtualHardDiskResponseOutput {
	return o.ApplyT(func(v *VirtualHardDiskResponse) VirtualHardDiskResponse {
		if v != nil {
			return *v
		}
		var ret VirtualHardDiskResponse
		return ret
	}).(VirtualHardDiskResponseOutput)
}

func (o VirtualHardDiskResponsePtrOutput) Uri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualHardDiskResponse) *string {
		if v == nil {
			return nil
		}
		return v.Uri
	}).(pulumi.StringPtrOutput)
}

type VirtualMachineAgentInstanceViewResponse struct {
	ExtensionHandlers []VirtualMachineExtensionHandlerInstanceViewResponse `pulumi:"extensionHandlers"`
	Statuses          []InstanceViewStatusResponse                         `pulumi:"statuses"`
	VmAgentVersion    *string                                              `pulumi:"vmAgentVersion"`
}

type VirtualMachineAgentInstanceViewResponseOutput struct{ *pulumi.OutputState }

func (VirtualMachineAgentInstanceViewResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualMachineAgentInstanceViewResponse)(nil)).Elem()
}

func (o VirtualMachineAgentInstanceViewResponseOutput) ToVirtualMachineAgentInstanceViewResponseOutput() VirtualMachineAgentInstanceViewResponseOutput {
	return o
}

func (o VirtualMachineAgentInstanceViewResponseOutput) ToVirtualMachineAgentInstanceViewResponseOutputWithContext(ctx context.Context) VirtualMachineAgentInstanceViewResponseOutput {
	return o
}

func (o VirtualMachineAgentInstanceViewResponseOutput) ExtensionHandlers() VirtualMachineExtensionHandlerInstanceViewResponseArrayOutput {
	return o.ApplyT(func(v VirtualMachineAgentInstanceViewResponse) []VirtualMachineExtensionHandlerInstanceViewResponse {
		return v.ExtensionHandlers
	}).(VirtualMachineExtensionHandlerInstanceViewResponseArrayOutput)
}

func (o VirtualMachineAgentInstanceViewResponseOutput) Statuses() InstanceViewStatusResponseArrayOutput {
	return o.ApplyT(func(v VirtualMachineAgentInstanceViewResponse) []InstanceViewStatusResponse { return v.Statuses }).(InstanceViewStatusResponseArrayOutput)
}

func (o VirtualMachineAgentInstanceViewResponseOutput) VmAgentVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualMachineAgentInstanceViewResponse) *string { return v.VmAgentVersion }).(pulumi.StringPtrOutput)
}

type VirtualMachineAgentInstanceViewResponsePtrOutput struct{ *pulumi.OutputState }

func (VirtualMachineAgentInstanceViewResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualMachineAgentInstanceViewResponse)(nil)).Elem()
}

func (o VirtualMachineAgentInstanceViewResponsePtrOutput) ToVirtualMachineAgentInstanceViewResponsePtrOutput() VirtualMachineAgentInstanceViewResponsePtrOutput {
	return o
}

func (o VirtualMachineAgentInstanceViewResponsePtrOutput) ToVirtualMachineAgentInstanceViewResponsePtrOutputWithContext(ctx context.Context) VirtualMachineAgentInstanceViewResponsePtrOutput {
	return o
}

func (o VirtualMachineAgentInstanceViewResponsePtrOutput) Elem() VirtualMachineAgentInstanceViewResponseOutput {
	return o.ApplyT(func(v *VirtualMachineAgentInstanceViewResponse) VirtualMachineAgentInstanceViewResponse {
		if v != nil {
			return *v
		}
		var ret VirtualMachineAgentInstanceViewResponse
		return ret
	}).(VirtualMachineAgentInstanceViewResponseOutput)
}

func (o VirtualMachineAgentInstanceViewResponsePtrOutput) ExtensionHandlers() VirtualMachineExtensionHandlerInstanceViewResponseArrayOutput {
	return o.ApplyT(func(v *VirtualMachineAgentInstanceViewResponse) []VirtualMachineExtensionHandlerInstanceViewResponse {
		if v == nil {
			return nil
		}
		return v.ExtensionHandlers
	}).(VirtualMachineExtensionHandlerInstanceViewResponseArrayOutput)
}

func (o VirtualMachineAgentInstanceViewResponsePtrOutput) Statuses() InstanceViewStatusResponseArrayOutput {
	return o.ApplyT(func(v *VirtualMachineAgentInstanceViewResponse) []InstanceViewStatusResponse {
		if v == nil {
			return nil
		}
		return v.Statuses
	}).(InstanceViewStatusResponseArrayOutput)
}

func (o VirtualMachineAgentInstanceViewResponsePtrOutput) VmAgentVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachineAgentInstanceViewResponse) *string {
		if v == nil {
			return nil
		}
		return v.VmAgentVersion
	}).(pulumi.StringPtrOutput)
}

type VirtualMachineExtensionHandlerInstanceViewResponse struct {
	Status             *InstanceViewStatusResponse `pulumi:"status"`
	Type               *string                     `pulumi:"type"`
	TypeHandlerVersion *string                     `pulumi:"typeHandlerVersion"`
}

type VirtualMachineExtensionHandlerInstanceViewResponseOutput struct{ *pulumi.OutputState }

func (VirtualMachineExtensionHandlerInstanceViewResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualMachineExtensionHandlerInstanceViewResponse)(nil)).Elem()
}

func (o VirtualMachineExtensionHandlerInstanceViewResponseOutput) ToVirtualMachineExtensionHandlerInstanceViewResponseOutput() VirtualMachineExtensionHandlerInstanceViewResponseOutput {
	return o
}

func (o VirtualMachineExtensionHandlerInstanceViewResponseOutput) ToVirtualMachineExtensionHandlerInstanceViewResponseOutputWithContext(ctx context.Context) VirtualMachineExtensionHandlerInstanceViewResponseOutput {
	return o
}

func (o VirtualMachineExtensionHandlerInstanceViewResponseOutput) Status() InstanceViewStatusResponsePtrOutput {
	return o.ApplyT(func(v VirtualMachineExtensionHandlerInstanceViewResponse) *InstanceViewStatusResponse {
		return v.Status
	}).(InstanceViewStatusResponsePtrOutput)
}

func (o VirtualMachineExtensionHandlerInstanceViewResponseOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualMachineExtensionHandlerInstanceViewResponse) *string { return v.Type }).(pulumi.StringPtrOutput)
}

func (o VirtualMachineExtensionHandlerInstanceViewResponseOutput) TypeHandlerVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualMachineExtensionHandlerInstanceViewResponse) *string { return v.TypeHandlerVersion }).(pulumi.StringPtrOutput)
}

type VirtualMachineExtensionHandlerInstanceViewResponseArrayOutput struct{ *pulumi.OutputState }

func (VirtualMachineExtensionHandlerInstanceViewResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VirtualMachineExtensionHandlerInstanceViewResponse)(nil)).Elem()
}

func (o VirtualMachineExtensionHandlerInstanceViewResponseArrayOutput) ToVirtualMachineExtensionHandlerInstanceViewResponseArrayOutput() VirtualMachineExtensionHandlerInstanceViewResponseArrayOutput {
	return o
}

func (o VirtualMachineExtensionHandlerInstanceViewResponseArrayOutput) ToVirtualMachineExtensionHandlerInstanceViewResponseArrayOutputWithContext(ctx context.Context) VirtualMachineExtensionHandlerInstanceViewResponseArrayOutput {
	return o
}

func (o VirtualMachineExtensionHandlerInstanceViewResponseArrayOutput) Index(i pulumi.IntInput) VirtualMachineExtensionHandlerInstanceViewResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VirtualMachineExtensionHandlerInstanceViewResponse {
		return vs[0].([]VirtualMachineExtensionHandlerInstanceViewResponse)[vs[1].(int)]
	}).(VirtualMachineExtensionHandlerInstanceViewResponseOutput)
}

type VirtualMachineExtensionInstanceView struct {
	Name               *string              `pulumi:"name"`
	Statuses           []InstanceViewStatus `pulumi:"statuses"`
	Substatuses        []InstanceViewStatus `pulumi:"substatuses"`
	Type               *string              `pulumi:"type"`
	TypeHandlerVersion *string              `pulumi:"typeHandlerVersion"`
}





type VirtualMachineExtensionInstanceViewInput interface {
	pulumi.Input

	ToVirtualMachineExtensionInstanceViewOutput() VirtualMachineExtensionInstanceViewOutput
	ToVirtualMachineExtensionInstanceViewOutputWithContext(context.Context) VirtualMachineExtensionInstanceViewOutput
}

type VirtualMachineExtensionInstanceViewArgs struct {
	Name               pulumi.StringPtrInput        `pulumi:"name"`
	Statuses           InstanceViewStatusArrayInput `pulumi:"statuses"`
	Substatuses        InstanceViewStatusArrayInput `pulumi:"substatuses"`
	Type               pulumi.StringPtrInput        `pulumi:"type"`
	TypeHandlerVersion pulumi.StringPtrInput        `pulumi:"typeHandlerVersion"`
}

func (VirtualMachineExtensionInstanceViewArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualMachineExtensionInstanceView)(nil)).Elem()
}

func (i VirtualMachineExtensionInstanceViewArgs) ToVirtualMachineExtensionInstanceViewOutput() VirtualMachineExtensionInstanceViewOutput {
	return i.ToVirtualMachineExtensionInstanceViewOutputWithContext(context.Background())
}

func (i VirtualMachineExtensionInstanceViewArgs) ToVirtualMachineExtensionInstanceViewOutputWithContext(ctx context.Context) VirtualMachineExtensionInstanceViewOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualMachineExtensionInstanceViewOutput)
}

func (i VirtualMachineExtensionInstanceViewArgs) ToVirtualMachineExtensionInstanceViewPtrOutput() VirtualMachineExtensionInstanceViewPtrOutput {
	return i.ToVirtualMachineExtensionInstanceViewPtrOutputWithContext(context.Background())
}

func (i VirtualMachineExtensionInstanceViewArgs) ToVirtualMachineExtensionInstanceViewPtrOutputWithContext(ctx context.Context) VirtualMachineExtensionInstanceViewPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualMachineExtensionInstanceViewOutput).ToVirtualMachineExtensionInstanceViewPtrOutputWithContext(ctx)
}









type VirtualMachineExtensionInstanceViewPtrInput interface {
	pulumi.Input

	ToVirtualMachineExtensionInstanceViewPtrOutput() VirtualMachineExtensionInstanceViewPtrOutput
	ToVirtualMachineExtensionInstanceViewPtrOutputWithContext(context.Context) VirtualMachineExtensionInstanceViewPtrOutput
}

type virtualMachineExtensionInstanceViewPtrType VirtualMachineExtensionInstanceViewArgs

func VirtualMachineExtensionInstanceViewPtr(v *VirtualMachineExtensionInstanceViewArgs) VirtualMachineExtensionInstanceViewPtrInput {
	return (*virtualMachineExtensionInstanceViewPtrType)(v)
}

func (*virtualMachineExtensionInstanceViewPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualMachineExtensionInstanceView)(nil)).Elem()
}

func (i *virtualMachineExtensionInstanceViewPtrType) ToVirtualMachineExtensionInstanceViewPtrOutput() VirtualMachineExtensionInstanceViewPtrOutput {
	return i.ToVirtualMachineExtensionInstanceViewPtrOutputWithContext(context.Background())
}

func (i *virtualMachineExtensionInstanceViewPtrType) ToVirtualMachineExtensionInstanceViewPtrOutputWithContext(ctx context.Context) VirtualMachineExtensionInstanceViewPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualMachineExtensionInstanceViewPtrOutput)
}

type VirtualMachineExtensionInstanceViewOutput struct{ *pulumi.OutputState }

func (VirtualMachineExtensionInstanceViewOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualMachineExtensionInstanceView)(nil)).Elem()
}

func (o VirtualMachineExtensionInstanceViewOutput) ToVirtualMachineExtensionInstanceViewOutput() VirtualMachineExtensionInstanceViewOutput {
	return o
}

func (o VirtualMachineExtensionInstanceViewOutput) ToVirtualMachineExtensionInstanceViewOutputWithContext(ctx context.Context) VirtualMachineExtensionInstanceViewOutput {
	return o
}

func (o VirtualMachineExtensionInstanceViewOutput) ToVirtualMachineExtensionInstanceViewPtrOutput() VirtualMachineExtensionInstanceViewPtrOutput {
	return o.ToVirtualMachineExtensionInstanceViewPtrOutputWithContext(context.Background())
}

func (o VirtualMachineExtensionInstanceViewOutput) ToVirtualMachineExtensionInstanceViewPtrOutputWithContext(ctx context.Context) VirtualMachineExtensionInstanceViewPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v VirtualMachineExtensionInstanceView) *VirtualMachineExtensionInstanceView {
		return &v
	}).(VirtualMachineExtensionInstanceViewPtrOutput)
}

func (o VirtualMachineExtensionInstanceViewOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualMachineExtensionInstanceView) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o VirtualMachineExtensionInstanceViewOutput) Statuses() InstanceViewStatusArrayOutput {
	return o.ApplyT(func(v VirtualMachineExtensionInstanceView) []InstanceViewStatus { return v.Statuses }).(InstanceViewStatusArrayOutput)
}

func (o VirtualMachineExtensionInstanceViewOutput) Substatuses() InstanceViewStatusArrayOutput {
	return o.ApplyT(func(v VirtualMachineExtensionInstanceView) []InstanceViewStatus { return v.Substatuses }).(InstanceViewStatusArrayOutput)
}

func (o VirtualMachineExtensionInstanceViewOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualMachineExtensionInstanceView) *string { return v.Type }).(pulumi.StringPtrOutput)
}

func (o VirtualMachineExtensionInstanceViewOutput) TypeHandlerVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualMachineExtensionInstanceView) *string { return v.TypeHandlerVersion }).(pulumi.StringPtrOutput)
}

type VirtualMachineExtensionInstanceViewPtrOutput struct{ *pulumi.OutputState }

func (VirtualMachineExtensionInstanceViewPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualMachineExtensionInstanceView)(nil)).Elem()
}

func (o VirtualMachineExtensionInstanceViewPtrOutput) ToVirtualMachineExtensionInstanceViewPtrOutput() VirtualMachineExtensionInstanceViewPtrOutput {
	return o
}

func (o VirtualMachineExtensionInstanceViewPtrOutput) ToVirtualMachineExtensionInstanceViewPtrOutputWithContext(ctx context.Context) VirtualMachineExtensionInstanceViewPtrOutput {
	return o
}

func (o VirtualMachineExtensionInstanceViewPtrOutput) Elem() VirtualMachineExtensionInstanceViewOutput {
	return o.ApplyT(func(v *VirtualMachineExtensionInstanceView) VirtualMachineExtensionInstanceView {
		if v != nil {
			return *v
		}
		var ret VirtualMachineExtensionInstanceView
		return ret
	}).(VirtualMachineExtensionInstanceViewOutput)
}

func (o VirtualMachineExtensionInstanceViewPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachineExtensionInstanceView) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o VirtualMachineExtensionInstanceViewPtrOutput) Statuses() InstanceViewStatusArrayOutput {
	return o.ApplyT(func(v *VirtualMachineExtensionInstanceView) []InstanceViewStatus {
		if v == nil {
			return nil
		}
		return v.Statuses
	}).(InstanceViewStatusArrayOutput)
}

func (o VirtualMachineExtensionInstanceViewPtrOutput) Substatuses() InstanceViewStatusArrayOutput {
	return o.ApplyT(func(v *VirtualMachineExtensionInstanceView) []InstanceViewStatus {
		if v == nil {
			return nil
		}
		return v.Substatuses
	}).(InstanceViewStatusArrayOutput)
}

func (o VirtualMachineExtensionInstanceViewPtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachineExtensionInstanceView) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

func (o VirtualMachineExtensionInstanceViewPtrOutput) TypeHandlerVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachineExtensionInstanceView) *string {
		if v == nil {
			return nil
		}
		return v.TypeHandlerVersion
	}).(pulumi.StringPtrOutput)
}

type VirtualMachineExtensionInstanceViewResponse struct {
	Name               *string                      `pulumi:"name"`
	Statuses           []InstanceViewStatusResponse `pulumi:"statuses"`
	Substatuses        []InstanceViewStatusResponse `pulumi:"substatuses"`
	Type               *string                      `pulumi:"type"`
	TypeHandlerVersion *string                      `pulumi:"typeHandlerVersion"`
}

type VirtualMachineExtensionInstanceViewResponseOutput struct{ *pulumi.OutputState }

func (VirtualMachineExtensionInstanceViewResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualMachineExtensionInstanceViewResponse)(nil)).Elem()
}

func (o VirtualMachineExtensionInstanceViewResponseOutput) ToVirtualMachineExtensionInstanceViewResponseOutput() VirtualMachineExtensionInstanceViewResponseOutput {
	return o
}

func (o VirtualMachineExtensionInstanceViewResponseOutput) ToVirtualMachineExtensionInstanceViewResponseOutputWithContext(ctx context.Context) VirtualMachineExtensionInstanceViewResponseOutput {
	return o
}

func (o VirtualMachineExtensionInstanceViewResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualMachineExtensionInstanceViewResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o VirtualMachineExtensionInstanceViewResponseOutput) Statuses() InstanceViewStatusResponseArrayOutput {
	return o.ApplyT(func(v VirtualMachineExtensionInstanceViewResponse) []InstanceViewStatusResponse { return v.Statuses }).(InstanceViewStatusResponseArrayOutput)
}

func (o VirtualMachineExtensionInstanceViewResponseOutput) Substatuses() InstanceViewStatusResponseArrayOutput {
	return o.ApplyT(func(v VirtualMachineExtensionInstanceViewResponse) []InstanceViewStatusResponse { return v.Substatuses }).(InstanceViewStatusResponseArrayOutput)
}

func (o VirtualMachineExtensionInstanceViewResponseOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualMachineExtensionInstanceViewResponse) *string { return v.Type }).(pulumi.StringPtrOutput)
}

func (o VirtualMachineExtensionInstanceViewResponseOutput) TypeHandlerVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualMachineExtensionInstanceViewResponse) *string { return v.TypeHandlerVersion }).(pulumi.StringPtrOutput)
}

type VirtualMachineExtensionInstanceViewResponsePtrOutput struct{ *pulumi.OutputState }

func (VirtualMachineExtensionInstanceViewResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualMachineExtensionInstanceViewResponse)(nil)).Elem()
}

func (o VirtualMachineExtensionInstanceViewResponsePtrOutput) ToVirtualMachineExtensionInstanceViewResponsePtrOutput() VirtualMachineExtensionInstanceViewResponsePtrOutput {
	return o
}

func (o VirtualMachineExtensionInstanceViewResponsePtrOutput) ToVirtualMachineExtensionInstanceViewResponsePtrOutputWithContext(ctx context.Context) VirtualMachineExtensionInstanceViewResponsePtrOutput {
	return o
}

func (o VirtualMachineExtensionInstanceViewResponsePtrOutput) Elem() VirtualMachineExtensionInstanceViewResponseOutput {
	return o.ApplyT(func(v *VirtualMachineExtensionInstanceViewResponse) VirtualMachineExtensionInstanceViewResponse {
		if v != nil {
			return *v
		}
		var ret VirtualMachineExtensionInstanceViewResponse
		return ret
	}).(VirtualMachineExtensionInstanceViewResponseOutput)
}

func (o VirtualMachineExtensionInstanceViewResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachineExtensionInstanceViewResponse) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o VirtualMachineExtensionInstanceViewResponsePtrOutput) Statuses() InstanceViewStatusResponseArrayOutput {
	return o.ApplyT(func(v *VirtualMachineExtensionInstanceViewResponse) []InstanceViewStatusResponse {
		if v == nil {
			return nil
		}
		return v.Statuses
	}).(InstanceViewStatusResponseArrayOutput)
}

func (o VirtualMachineExtensionInstanceViewResponsePtrOutput) Substatuses() InstanceViewStatusResponseArrayOutput {
	return o.ApplyT(func(v *VirtualMachineExtensionInstanceViewResponse) []InstanceViewStatusResponse {
		if v == nil {
			return nil
		}
		return v.Substatuses
	}).(InstanceViewStatusResponseArrayOutput)
}

func (o VirtualMachineExtensionInstanceViewResponsePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachineExtensionInstanceViewResponse) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

func (o VirtualMachineExtensionInstanceViewResponsePtrOutput) TypeHandlerVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachineExtensionInstanceViewResponse) *string {
		if v == nil {
			return nil
		}
		return v.TypeHandlerVersion
	}).(pulumi.StringPtrOutput)
}

type VirtualMachineExtensionInstanceViewResponseArrayOutput struct{ *pulumi.OutputState }

func (VirtualMachineExtensionInstanceViewResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VirtualMachineExtensionInstanceViewResponse)(nil)).Elem()
}

func (o VirtualMachineExtensionInstanceViewResponseArrayOutput) ToVirtualMachineExtensionInstanceViewResponseArrayOutput() VirtualMachineExtensionInstanceViewResponseArrayOutput {
	return o
}

func (o VirtualMachineExtensionInstanceViewResponseArrayOutput) ToVirtualMachineExtensionInstanceViewResponseArrayOutputWithContext(ctx context.Context) VirtualMachineExtensionInstanceViewResponseArrayOutput {
	return o
}

func (o VirtualMachineExtensionInstanceViewResponseArrayOutput) Index(i pulumi.IntInput) VirtualMachineExtensionInstanceViewResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VirtualMachineExtensionInstanceViewResponse {
		return vs[0].([]VirtualMachineExtensionInstanceViewResponse)[vs[1].(int)]
	}).(VirtualMachineExtensionInstanceViewResponseOutput)
}

type VirtualMachineExtensionResponse struct {
	AutoUpgradeMinorVersion *bool                                        `pulumi:"autoUpgradeMinorVersion"`
	ForceUpdateTag          *string                                      `pulumi:"forceUpdateTag"`
	Id                      string                                       `pulumi:"id"`
	InstanceView            *VirtualMachineExtensionInstanceViewResponse `pulumi:"instanceView"`
	Location                string                                       `pulumi:"location"`
	Name                    string                                       `pulumi:"name"`
	ProtectedSettings       interface{}                                  `pulumi:"protectedSettings"`
	ProvisioningState       string                                       `pulumi:"provisioningState"`
	Publisher               *string                                      `pulumi:"publisher"`
	Settings                interface{}                                  `pulumi:"settings"`
	Tags                    map[string]string                            `pulumi:"tags"`
	Type                    string                                       `pulumi:"type"`
	TypeHandlerVersion      *string                                      `pulumi:"typeHandlerVersion"`
}

type VirtualMachineExtensionResponseOutput struct{ *pulumi.OutputState }

func (VirtualMachineExtensionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualMachineExtensionResponse)(nil)).Elem()
}

func (o VirtualMachineExtensionResponseOutput) ToVirtualMachineExtensionResponseOutput() VirtualMachineExtensionResponseOutput {
	return o
}

func (o VirtualMachineExtensionResponseOutput) ToVirtualMachineExtensionResponseOutputWithContext(ctx context.Context) VirtualMachineExtensionResponseOutput {
	return o
}

func (o VirtualMachineExtensionResponseOutput) AutoUpgradeMinorVersion() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VirtualMachineExtensionResponse) *bool { return v.AutoUpgradeMinorVersion }).(pulumi.BoolPtrOutput)
}

func (o VirtualMachineExtensionResponseOutput) ForceUpdateTag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualMachineExtensionResponse) *string { return v.ForceUpdateTag }).(pulumi.StringPtrOutput)
}

func (o VirtualMachineExtensionResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v VirtualMachineExtensionResponse) string { return v.Id }).(pulumi.StringOutput)
}

func (o VirtualMachineExtensionResponseOutput) InstanceView() VirtualMachineExtensionInstanceViewResponsePtrOutput {
	return o.ApplyT(func(v VirtualMachineExtensionResponse) *VirtualMachineExtensionInstanceViewResponse {
		return v.InstanceView
	}).(VirtualMachineExtensionInstanceViewResponsePtrOutput)
}

func (o VirtualMachineExtensionResponseOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v VirtualMachineExtensionResponse) string { return v.Location }).(pulumi.StringOutput)
}

func (o VirtualMachineExtensionResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v VirtualMachineExtensionResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o VirtualMachineExtensionResponseOutput) ProtectedSettings() pulumi.AnyOutput {
	return o.ApplyT(func(v VirtualMachineExtensionResponse) interface{} { return v.ProtectedSettings }).(pulumi.AnyOutput)
}

func (o VirtualMachineExtensionResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v VirtualMachineExtensionResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o VirtualMachineExtensionResponseOutput) Publisher() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualMachineExtensionResponse) *string { return v.Publisher }).(pulumi.StringPtrOutput)
}

func (o VirtualMachineExtensionResponseOutput) Settings() pulumi.AnyOutput {
	return o.ApplyT(func(v VirtualMachineExtensionResponse) interface{} { return v.Settings }).(pulumi.AnyOutput)
}

func (o VirtualMachineExtensionResponseOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v VirtualMachineExtensionResponse) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o VirtualMachineExtensionResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v VirtualMachineExtensionResponse) string { return v.Type }).(pulumi.StringOutput)
}

func (o VirtualMachineExtensionResponseOutput) TypeHandlerVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualMachineExtensionResponse) *string { return v.TypeHandlerVersion }).(pulumi.StringPtrOutput)
}

type VirtualMachineExtensionResponseArrayOutput struct{ *pulumi.OutputState }

func (VirtualMachineExtensionResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VirtualMachineExtensionResponse)(nil)).Elem()
}

func (o VirtualMachineExtensionResponseArrayOutput) ToVirtualMachineExtensionResponseArrayOutput() VirtualMachineExtensionResponseArrayOutput {
	return o
}

func (o VirtualMachineExtensionResponseArrayOutput) ToVirtualMachineExtensionResponseArrayOutputWithContext(ctx context.Context) VirtualMachineExtensionResponseArrayOutput {
	return o
}

func (o VirtualMachineExtensionResponseArrayOutput) Index(i pulumi.IntInput) VirtualMachineExtensionResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VirtualMachineExtensionResponse {
		return vs[0].([]VirtualMachineExtensionResponse)[vs[1].(int)]
	}).(VirtualMachineExtensionResponseOutput)
}

type VirtualMachineHealthStatusResponse struct {
	Status InstanceViewStatusResponse `pulumi:"status"`
}

type VirtualMachineHealthStatusResponseOutput struct{ *pulumi.OutputState }

func (VirtualMachineHealthStatusResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualMachineHealthStatusResponse)(nil)).Elem()
}

func (o VirtualMachineHealthStatusResponseOutput) ToVirtualMachineHealthStatusResponseOutput() VirtualMachineHealthStatusResponseOutput {
	return o
}

func (o VirtualMachineHealthStatusResponseOutput) ToVirtualMachineHealthStatusResponseOutputWithContext(ctx context.Context) VirtualMachineHealthStatusResponseOutput {
	return o
}

func (o VirtualMachineHealthStatusResponseOutput) Status() InstanceViewStatusResponseOutput {
	return o.ApplyT(func(v VirtualMachineHealthStatusResponse) InstanceViewStatusResponse { return v.Status }).(InstanceViewStatusResponseOutput)
}

type VirtualMachineIdentity struct {
	IdentityIds []string              `pulumi:"identityIds"`
	Type        *ResourceIdentityType `pulumi:"type"`
}





type VirtualMachineIdentityInput interface {
	pulumi.Input

	ToVirtualMachineIdentityOutput() VirtualMachineIdentityOutput
	ToVirtualMachineIdentityOutputWithContext(context.Context) VirtualMachineIdentityOutput
}

type VirtualMachineIdentityArgs struct {
	IdentityIds pulumi.StringArrayInput      `pulumi:"identityIds"`
	Type        ResourceIdentityTypePtrInput `pulumi:"type"`
}

func (VirtualMachineIdentityArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualMachineIdentity)(nil)).Elem()
}

func (i VirtualMachineIdentityArgs) ToVirtualMachineIdentityOutput() VirtualMachineIdentityOutput {
	return i.ToVirtualMachineIdentityOutputWithContext(context.Background())
}

func (i VirtualMachineIdentityArgs) ToVirtualMachineIdentityOutputWithContext(ctx context.Context) VirtualMachineIdentityOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualMachineIdentityOutput)
}

func (i VirtualMachineIdentityArgs) ToVirtualMachineIdentityPtrOutput() VirtualMachineIdentityPtrOutput {
	return i.ToVirtualMachineIdentityPtrOutputWithContext(context.Background())
}

func (i VirtualMachineIdentityArgs) ToVirtualMachineIdentityPtrOutputWithContext(ctx context.Context) VirtualMachineIdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualMachineIdentityOutput).ToVirtualMachineIdentityPtrOutputWithContext(ctx)
}









type VirtualMachineIdentityPtrInput interface {
	pulumi.Input

	ToVirtualMachineIdentityPtrOutput() VirtualMachineIdentityPtrOutput
	ToVirtualMachineIdentityPtrOutputWithContext(context.Context) VirtualMachineIdentityPtrOutput
}

type virtualMachineIdentityPtrType VirtualMachineIdentityArgs

func VirtualMachineIdentityPtr(v *VirtualMachineIdentityArgs) VirtualMachineIdentityPtrInput {
	return (*virtualMachineIdentityPtrType)(v)
}

func (*virtualMachineIdentityPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualMachineIdentity)(nil)).Elem()
}

func (i *virtualMachineIdentityPtrType) ToVirtualMachineIdentityPtrOutput() VirtualMachineIdentityPtrOutput {
	return i.ToVirtualMachineIdentityPtrOutputWithContext(context.Background())
}

func (i *virtualMachineIdentityPtrType) ToVirtualMachineIdentityPtrOutputWithContext(ctx context.Context) VirtualMachineIdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualMachineIdentityPtrOutput)
}

type VirtualMachineIdentityOutput struct{ *pulumi.OutputState }

func (VirtualMachineIdentityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualMachineIdentity)(nil)).Elem()
}

func (o VirtualMachineIdentityOutput) ToVirtualMachineIdentityOutput() VirtualMachineIdentityOutput {
	return o
}

func (o VirtualMachineIdentityOutput) ToVirtualMachineIdentityOutputWithContext(ctx context.Context) VirtualMachineIdentityOutput {
	return o
}

func (o VirtualMachineIdentityOutput) ToVirtualMachineIdentityPtrOutput() VirtualMachineIdentityPtrOutput {
	return o.ToVirtualMachineIdentityPtrOutputWithContext(context.Background())
}

func (o VirtualMachineIdentityOutput) ToVirtualMachineIdentityPtrOutputWithContext(ctx context.Context) VirtualMachineIdentityPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v VirtualMachineIdentity) *VirtualMachineIdentity {
		return &v
	}).(VirtualMachineIdentityPtrOutput)
}

func (o VirtualMachineIdentityOutput) IdentityIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v VirtualMachineIdentity) []string { return v.IdentityIds }).(pulumi.StringArrayOutput)
}

func (o VirtualMachineIdentityOutput) Type() ResourceIdentityTypePtrOutput {
	return o.ApplyT(func(v VirtualMachineIdentity) *ResourceIdentityType { return v.Type }).(ResourceIdentityTypePtrOutput)
}

type VirtualMachineIdentityPtrOutput struct{ *pulumi.OutputState }

func (VirtualMachineIdentityPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualMachineIdentity)(nil)).Elem()
}

func (o VirtualMachineIdentityPtrOutput) ToVirtualMachineIdentityPtrOutput() VirtualMachineIdentityPtrOutput {
	return o
}

func (o VirtualMachineIdentityPtrOutput) ToVirtualMachineIdentityPtrOutputWithContext(ctx context.Context) VirtualMachineIdentityPtrOutput {
	return o
}

func (o VirtualMachineIdentityPtrOutput) Elem() VirtualMachineIdentityOutput {
	return o.ApplyT(func(v *VirtualMachineIdentity) VirtualMachineIdentity {
		if v != nil {
			return *v
		}
		var ret VirtualMachineIdentity
		return ret
	}).(VirtualMachineIdentityOutput)
}

func (o VirtualMachineIdentityPtrOutput) IdentityIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *VirtualMachineIdentity) []string {
		if v == nil {
			return nil
		}
		return v.IdentityIds
	}).(pulumi.StringArrayOutput)
}

func (o VirtualMachineIdentityPtrOutput) Type() ResourceIdentityTypePtrOutput {
	return o.ApplyT(func(v *VirtualMachineIdentity) *ResourceIdentityType {
		if v == nil {
			return nil
		}
		return v.Type
	}).(ResourceIdentityTypePtrOutput)
}

type VirtualMachineIdentityResponse struct {
	IdentityIds []string `pulumi:"identityIds"`
	PrincipalId string   `pulumi:"principalId"`
	TenantId    string   `pulumi:"tenantId"`
	Type        *string  `pulumi:"type"`
}

type VirtualMachineIdentityResponseOutput struct{ *pulumi.OutputState }

func (VirtualMachineIdentityResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualMachineIdentityResponse)(nil)).Elem()
}

func (o VirtualMachineIdentityResponseOutput) ToVirtualMachineIdentityResponseOutput() VirtualMachineIdentityResponseOutput {
	return o
}

func (o VirtualMachineIdentityResponseOutput) ToVirtualMachineIdentityResponseOutputWithContext(ctx context.Context) VirtualMachineIdentityResponseOutput {
	return o
}

func (o VirtualMachineIdentityResponseOutput) IdentityIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v VirtualMachineIdentityResponse) []string { return v.IdentityIds }).(pulumi.StringArrayOutput)
}

func (o VirtualMachineIdentityResponseOutput) PrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v VirtualMachineIdentityResponse) string { return v.PrincipalId }).(pulumi.StringOutput)
}

func (o VirtualMachineIdentityResponseOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v VirtualMachineIdentityResponse) string { return v.TenantId }).(pulumi.StringOutput)
}

func (o VirtualMachineIdentityResponseOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualMachineIdentityResponse) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type VirtualMachineIdentityResponsePtrOutput struct{ *pulumi.OutputState }

func (VirtualMachineIdentityResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualMachineIdentityResponse)(nil)).Elem()
}

func (o VirtualMachineIdentityResponsePtrOutput) ToVirtualMachineIdentityResponsePtrOutput() VirtualMachineIdentityResponsePtrOutput {
	return o
}

func (o VirtualMachineIdentityResponsePtrOutput) ToVirtualMachineIdentityResponsePtrOutputWithContext(ctx context.Context) VirtualMachineIdentityResponsePtrOutput {
	return o
}

func (o VirtualMachineIdentityResponsePtrOutput) Elem() VirtualMachineIdentityResponseOutput {
	return o.ApplyT(func(v *VirtualMachineIdentityResponse) VirtualMachineIdentityResponse {
		if v != nil {
			return *v
		}
		var ret VirtualMachineIdentityResponse
		return ret
	}).(VirtualMachineIdentityResponseOutput)
}

func (o VirtualMachineIdentityResponsePtrOutput) IdentityIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *VirtualMachineIdentityResponse) []string {
		if v == nil {
			return nil
		}
		return v.IdentityIds
	}).(pulumi.StringArrayOutput)
}

func (o VirtualMachineIdentityResponsePtrOutput) PrincipalId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachineIdentityResponse) *string {
		if v == nil {
			return nil
		}
		return &v.PrincipalId
	}).(pulumi.StringPtrOutput)
}

func (o VirtualMachineIdentityResponsePtrOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachineIdentityResponse) *string {
		if v == nil {
			return nil
		}
		return &v.TenantId
	}).(pulumi.StringPtrOutput)
}

func (o VirtualMachineIdentityResponsePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachineIdentityResponse) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

type VirtualMachineInstanceViewResponse struct {
	BootDiagnostics           *BootDiagnosticsInstanceViewResponse          `pulumi:"bootDiagnostics"`
	ComputerName              *string                                       `pulumi:"computerName"`
	Disks                     []DiskInstanceViewResponse                    `pulumi:"disks"`
	Extensions                []VirtualMachineExtensionInstanceViewResponse `pulumi:"extensions"`
	MaintenanceRedeployStatus *MaintenanceRedeployStatusResponse            `pulumi:"maintenanceRedeployStatus"`
	OsName                    *string                                       `pulumi:"osName"`
	OsVersion                 *string                                       `pulumi:"osVersion"`
	PlatformFaultDomain       *int                                          `pulumi:"platformFaultDomain"`
	PlatformUpdateDomain      *int                                          `pulumi:"platformUpdateDomain"`
	RdpThumbPrint             *string                                       `pulumi:"rdpThumbPrint"`
	Statuses                  []InstanceViewStatusResponse                  `pulumi:"statuses"`
	VmAgent                   *VirtualMachineAgentInstanceViewResponse      `pulumi:"vmAgent"`
}

type VirtualMachineInstanceViewResponseOutput struct{ *pulumi.OutputState }

func (VirtualMachineInstanceViewResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualMachineInstanceViewResponse)(nil)).Elem()
}

func (o VirtualMachineInstanceViewResponseOutput) ToVirtualMachineInstanceViewResponseOutput() VirtualMachineInstanceViewResponseOutput {
	return o
}

func (o VirtualMachineInstanceViewResponseOutput) ToVirtualMachineInstanceViewResponseOutputWithContext(ctx context.Context) VirtualMachineInstanceViewResponseOutput {
	return o
}

func (o VirtualMachineInstanceViewResponseOutput) BootDiagnostics() BootDiagnosticsInstanceViewResponsePtrOutput {
	return o.ApplyT(func(v VirtualMachineInstanceViewResponse) *BootDiagnosticsInstanceViewResponse {
		return v.BootDiagnostics
	}).(BootDiagnosticsInstanceViewResponsePtrOutput)
}

func (o VirtualMachineInstanceViewResponseOutput) ComputerName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualMachineInstanceViewResponse) *string { return v.ComputerName }).(pulumi.StringPtrOutput)
}

func (o VirtualMachineInstanceViewResponseOutput) Disks() DiskInstanceViewResponseArrayOutput {
	return o.ApplyT(func(v VirtualMachineInstanceViewResponse) []DiskInstanceViewResponse { return v.Disks }).(DiskInstanceViewResponseArrayOutput)
}

func (o VirtualMachineInstanceViewResponseOutput) Extensions() VirtualMachineExtensionInstanceViewResponseArrayOutput {
	return o.ApplyT(func(v VirtualMachineInstanceViewResponse) []VirtualMachineExtensionInstanceViewResponse {
		return v.Extensions
	}).(VirtualMachineExtensionInstanceViewResponseArrayOutput)
}

func (o VirtualMachineInstanceViewResponseOutput) MaintenanceRedeployStatus() MaintenanceRedeployStatusResponsePtrOutput {
	return o.ApplyT(func(v VirtualMachineInstanceViewResponse) *MaintenanceRedeployStatusResponse {
		return v.MaintenanceRedeployStatus
	}).(MaintenanceRedeployStatusResponsePtrOutput)
}

func (o VirtualMachineInstanceViewResponseOutput) OsName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualMachineInstanceViewResponse) *string { return v.OsName }).(pulumi.StringPtrOutput)
}

func (o VirtualMachineInstanceViewResponseOutput) OsVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualMachineInstanceViewResponse) *string { return v.OsVersion }).(pulumi.StringPtrOutput)
}

func (o VirtualMachineInstanceViewResponseOutput) PlatformFaultDomain() pulumi.IntPtrOutput {
	return o.ApplyT(func(v VirtualMachineInstanceViewResponse) *int { return v.PlatformFaultDomain }).(pulumi.IntPtrOutput)
}

func (o VirtualMachineInstanceViewResponseOutput) PlatformUpdateDomain() pulumi.IntPtrOutput {
	return o.ApplyT(func(v VirtualMachineInstanceViewResponse) *int { return v.PlatformUpdateDomain }).(pulumi.IntPtrOutput)
}

func (o VirtualMachineInstanceViewResponseOutput) RdpThumbPrint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualMachineInstanceViewResponse) *string { return v.RdpThumbPrint }).(pulumi.StringPtrOutput)
}

func (o VirtualMachineInstanceViewResponseOutput) Statuses() InstanceViewStatusResponseArrayOutput {
	return o.ApplyT(func(v VirtualMachineInstanceViewResponse) []InstanceViewStatusResponse { return v.Statuses }).(InstanceViewStatusResponseArrayOutput)
}

func (o VirtualMachineInstanceViewResponseOutput) VmAgent() VirtualMachineAgentInstanceViewResponsePtrOutput {
	return o.ApplyT(func(v VirtualMachineInstanceViewResponse) *VirtualMachineAgentInstanceViewResponse { return v.VmAgent }).(VirtualMachineAgentInstanceViewResponsePtrOutput)
}

type VirtualMachineScaleSetDataDisk struct {
	Caching                 *CachingTypes                                `pulumi:"caching"`
	CreateOption            string                                       `pulumi:"createOption"`
	DiskSizeGB              *int                                         `pulumi:"diskSizeGB"`
	Lun                     int                                          `pulumi:"lun"`
	ManagedDisk             *VirtualMachineScaleSetManagedDiskParameters `pulumi:"managedDisk"`
	Name                    *string                                      `pulumi:"name"`
	WriteAcceleratorEnabled *bool                                        `pulumi:"writeAcceleratorEnabled"`
}





type VirtualMachineScaleSetDataDiskInput interface {
	pulumi.Input

	ToVirtualMachineScaleSetDataDiskOutput() VirtualMachineScaleSetDataDiskOutput
	ToVirtualMachineScaleSetDataDiskOutputWithContext(context.Context) VirtualMachineScaleSetDataDiskOutput
}

type VirtualMachineScaleSetDataDiskArgs struct {
	Caching                 CachingTypesPtrInput                                `pulumi:"caching"`
	CreateOption            pulumi.StringInput                                  `pulumi:"createOption"`
	DiskSizeGB              pulumi.IntPtrInput                                  `pulumi:"diskSizeGB"`
	Lun                     pulumi.IntInput                                     `pulumi:"lun"`
	ManagedDisk             VirtualMachineScaleSetManagedDiskParametersPtrInput `pulumi:"managedDisk"`
	Name                    pulumi.StringPtrInput                               `pulumi:"name"`
	WriteAcceleratorEnabled pulumi.BoolPtrInput                                 `pulumi:"writeAcceleratorEnabled"`
}

func (VirtualMachineScaleSetDataDiskArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualMachineScaleSetDataDisk)(nil)).Elem()
}

func (i VirtualMachineScaleSetDataDiskArgs) ToVirtualMachineScaleSetDataDiskOutput() VirtualMachineScaleSetDataDiskOutput {
	return i.ToVirtualMachineScaleSetDataDiskOutputWithContext(context.Background())
}

func (i VirtualMachineScaleSetDataDiskArgs) ToVirtualMachineScaleSetDataDiskOutputWithContext(ctx context.Context) VirtualMachineScaleSetDataDiskOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualMachineScaleSetDataDiskOutput)
}





type VirtualMachineScaleSetDataDiskArrayInput interface {
	pulumi.Input

	ToVirtualMachineScaleSetDataDiskArrayOutput() VirtualMachineScaleSetDataDiskArrayOutput
	ToVirtualMachineScaleSetDataDiskArrayOutputWithContext(context.Context) VirtualMachineScaleSetDataDiskArrayOutput
}

type VirtualMachineScaleSetDataDiskArray []VirtualMachineScaleSetDataDiskInput

func (VirtualMachineScaleSetDataDiskArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VirtualMachineScaleSetDataDisk)(nil)).Elem()
}

func (i VirtualMachineScaleSetDataDiskArray) ToVirtualMachineScaleSetDataDiskArrayOutput() VirtualMachineScaleSetDataDiskArrayOutput {
	return i.ToVirtualMachineScaleSetDataDiskArrayOutputWithContext(context.Background())
}

func (i VirtualMachineScaleSetDataDiskArray) ToVirtualMachineScaleSetDataDiskArrayOutputWithContext(ctx context.Context) VirtualMachineScaleSetDataDiskArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualMachineScaleSetDataDiskArrayOutput)
}

type VirtualMachineScaleSetDataDiskOutput struct{ *pulumi.OutputState }

func (VirtualMachineScaleSetDataDiskOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualMachineScaleSetDataDisk)(nil)).Elem()
}

func (o VirtualMachineScaleSetDataDiskOutput) ToVirtualMachineScaleSetDataDiskOutput() VirtualMachineScaleSetDataDiskOutput {
	return o
}

func (o VirtualMachineScaleSetDataDiskOutput) ToVirtualMachineScaleSetDataDiskOutputWithContext(ctx context.Context) VirtualMachineScaleSetDataDiskOutput {
	return o
}

func (o VirtualMachineScaleSetDataDiskOutput) Caching() CachingTypesPtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetDataDisk) *CachingTypes { return v.Caching }).(CachingTypesPtrOutput)
}

func (o VirtualMachineScaleSetDataDiskOutput) CreateOption() pulumi.StringOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetDataDisk) string { return v.CreateOption }).(pulumi.StringOutput)
}

func (o VirtualMachineScaleSetDataDiskOutput) DiskSizeGB() pulumi.IntPtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetDataDisk) *int { return v.DiskSizeGB }).(pulumi.IntPtrOutput)
}

func (o VirtualMachineScaleSetDataDiskOutput) Lun() pulumi.IntOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetDataDisk) int { return v.Lun }).(pulumi.IntOutput)
}

func (o VirtualMachineScaleSetDataDiskOutput) ManagedDisk() VirtualMachineScaleSetManagedDiskParametersPtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetDataDisk) *VirtualMachineScaleSetManagedDiskParameters {
		return v.ManagedDisk
	}).(VirtualMachineScaleSetManagedDiskParametersPtrOutput)
}

func (o VirtualMachineScaleSetDataDiskOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetDataDisk) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o VirtualMachineScaleSetDataDiskOutput) WriteAcceleratorEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetDataDisk) *bool { return v.WriteAcceleratorEnabled }).(pulumi.BoolPtrOutput)
}

type VirtualMachineScaleSetDataDiskArrayOutput struct{ *pulumi.OutputState }

func (VirtualMachineScaleSetDataDiskArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VirtualMachineScaleSetDataDisk)(nil)).Elem()
}

func (o VirtualMachineScaleSetDataDiskArrayOutput) ToVirtualMachineScaleSetDataDiskArrayOutput() VirtualMachineScaleSetDataDiskArrayOutput {
	return o
}

func (o VirtualMachineScaleSetDataDiskArrayOutput) ToVirtualMachineScaleSetDataDiskArrayOutputWithContext(ctx context.Context) VirtualMachineScaleSetDataDiskArrayOutput {
	return o
}

func (o VirtualMachineScaleSetDataDiskArrayOutput) Index(i pulumi.IntInput) VirtualMachineScaleSetDataDiskOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VirtualMachineScaleSetDataDisk {
		return vs[0].([]VirtualMachineScaleSetDataDisk)[vs[1].(int)]
	}).(VirtualMachineScaleSetDataDiskOutput)
}

type VirtualMachineScaleSetDataDiskResponse struct {
	Caching                 *string                                              `pulumi:"caching"`
	CreateOption            string                                               `pulumi:"createOption"`
	DiskSizeGB              *int                                                 `pulumi:"diskSizeGB"`
	Lun                     int                                                  `pulumi:"lun"`
	ManagedDisk             *VirtualMachineScaleSetManagedDiskParametersResponse `pulumi:"managedDisk"`
	Name                    *string                                              `pulumi:"name"`
	WriteAcceleratorEnabled *bool                                                `pulumi:"writeAcceleratorEnabled"`
}

type VirtualMachineScaleSetDataDiskResponseOutput struct{ *pulumi.OutputState }

func (VirtualMachineScaleSetDataDiskResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualMachineScaleSetDataDiskResponse)(nil)).Elem()
}

func (o VirtualMachineScaleSetDataDiskResponseOutput) ToVirtualMachineScaleSetDataDiskResponseOutput() VirtualMachineScaleSetDataDiskResponseOutput {
	return o
}

func (o VirtualMachineScaleSetDataDiskResponseOutput) ToVirtualMachineScaleSetDataDiskResponseOutputWithContext(ctx context.Context) VirtualMachineScaleSetDataDiskResponseOutput {
	return o
}

func (o VirtualMachineScaleSetDataDiskResponseOutput) Caching() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetDataDiskResponse) *string { return v.Caching }).(pulumi.StringPtrOutput)
}

func (o VirtualMachineScaleSetDataDiskResponseOutput) CreateOption() pulumi.StringOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetDataDiskResponse) string { return v.CreateOption }).(pulumi.StringOutput)
}

func (o VirtualMachineScaleSetDataDiskResponseOutput) DiskSizeGB() pulumi.IntPtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetDataDiskResponse) *int { return v.DiskSizeGB }).(pulumi.IntPtrOutput)
}

func (o VirtualMachineScaleSetDataDiskResponseOutput) Lun() pulumi.IntOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetDataDiskResponse) int { return v.Lun }).(pulumi.IntOutput)
}

func (o VirtualMachineScaleSetDataDiskResponseOutput) ManagedDisk() VirtualMachineScaleSetManagedDiskParametersResponsePtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetDataDiskResponse) *VirtualMachineScaleSetManagedDiskParametersResponse {
		return v.ManagedDisk
	}).(VirtualMachineScaleSetManagedDiskParametersResponsePtrOutput)
}

func (o VirtualMachineScaleSetDataDiskResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetDataDiskResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o VirtualMachineScaleSetDataDiskResponseOutput) WriteAcceleratorEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetDataDiskResponse) *bool { return v.WriteAcceleratorEnabled }).(pulumi.BoolPtrOutput)
}

type VirtualMachineScaleSetDataDiskResponseArrayOutput struct{ *pulumi.OutputState }

func (VirtualMachineScaleSetDataDiskResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VirtualMachineScaleSetDataDiskResponse)(nil)).Elem()
}

func (o VirtualMachineScaleSetDataDiskResponseArrayOutput) ToVirtualMachineScaleSetDataDiskResponseArrayOutput() VirtualMachineScaleSetDataDiskResponseArrayOutput {
	return o
}

func (o VirtualMachineScaleSetDataDiskResponseArrayOutput) ToVirtualMachineScaleSetDataDiskResponseArrayOutputWithContext(ctx context.Context) VirtualMachineScaleSetDataDiskResponseArrayOutput {
	return o
}

func (o VirtualMachineScaleSetDataDiskResponseArrayOutput) Index(i pulumi.IntInput) VirtualMachineScaleSetDataDiskResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VirtualMachineScaleSetDataDiskResponse {
		return vs[0].([]VirtualMachineScaleSetDataDiskResponse)[vs[1].(int)]
	}).(VirtualMachineScaleSetDataDiskResponseOutput)
}

type VirtualMachineScaleSetExtensionType struct {
	AutoUpgradeMinorVersion *bool       `pulumi:"autoUpgradeMinorVersion"`
	ForceUpdateTag          *string     `pulumi:"forceUpdateTag"`
	Name                    *string     `pulumi:"name"`
	ProtectedSettings       interface{} `pulumi:"protectedSettings"`
	Publisher               *string     `pulumi:"publisher"`
	Settings                interface{} `pulumi:"settings"`
	Type                    *string     `pulumi:"type"`
	TypeHandlerVersion      *string     `pulumi:"typeHandlerVersion"`
}





type VirtualMachineScaleSetExtensionTypeInput interface {
	pulumi.Input

	ToVirtualMachineScaleSetExtensionTypeOutput() VirtualMachineScaleSetExtensionTypeOutput
	ToVirtualMachineScaleSetExtensionTypeOutputWithContext(context.Context) VirtualMachineScaleSetExtensionTypeOutput
}

type VirtualMachineScaleSetExtensionTypeArgs struct {
	AutoUpgradeMinorVersion pulumi.BoolPtrInput   `pulumi:"autoUpgradeMinorVersion"`
	ForceUpdateTag          pulumi.StringPtrInput `pulumi:"forceUpdateTag"`
	Name                    pulumi.StringPtrInput `pulumi:"name"`
	ProtectedSettings       pulumi.Input          `pulumi:"protectedSettings"`
	Publisher               pulumi.StringPtrInput `pulumi:"publisher"`
	Settings                pulumi.Input          `pulumi:"settings"`
	Type                    pulumi.StringPtrInput `pulumi:"type"`
	TypeHandlerVersion      pulumi.StringPtrInput `pulumi:"typeHandlerVersion"`
}

func (VirtualMachineScaleSetExtensionTypeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualMachineScaleSetExtensionType)(nil)).Elem()
}

func (i VirtualMachineScaleSetExtensionTypeArgs) ToVirtualMachineScaleSetExtensionTypeOutput() VirtualMachineScaleSetExtensionTypeOutput {
	return i.ToVirtualMachineScaleSetExtensionTypeOutputWithContext(context.Background())
}

func (i VirtualMachineScaleSetExtensionTypeArgs) ToVirtualMachineScaleSetExtensionTypeOutputWithContext(ctx context.Context) VirtualMachineScaleSetExtensionTypeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualMachineScaleSetExtensionTypeOutput)
}





type VirtualMachineScaleSetExtensionTypeArrayInput interface {
	pulumi.Input

	ToVirtualMachineScaleSetExtensionTypeArrayOutput() VirtualMachineScaleSetExtensionTypeArrayOutput
	ToVirtualMachineScaleSetExtensionTypeArrayOutputWithContext(context.Context) VirtualMachineScaleSetExtensionTypeArrayOutput
}

type VirtualMachineScaleSetExtensionTypeArray []VirtualMachineScaleSetExtensionTypeInput

func (VirtualMachineScaleSetExtensionTypeArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VirtualMachineScaleSetExtensionType)(nil)).Elem()
}

func (i VirtualMachineScaleSetExtensionTypeArray) ToVirtualMachineScaleSetExtensionTypeArrayOutput() VirtualMachineScaleSetExtensionTypeArrayOutput {
	return i.ToVirtualMachineScaleSetExtensionTypeArrayOutputWithContext(context.Background())
}

func (i VirtualMachineScaleSetExtensionTypeArray) ToVirtualMachineScaleSetExtensionTypeArrayOutputWithContext(ctx context.Context) VirtualMachineScaleSetExtensionTypeArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualMachineScaleSetExtensionTypeArrayOutput)
}

type VirtualMachineScaleSetExtensionTypeOutput struct{ *pulumi.OutputState }

func (VirtualMachineScaleSetExtensionTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualMachineScaleSetExtensionType)(nil)).Elem()
}

func (o VirtualMachineScaleSetExtensionTypeOutput) ToVirtualMachineScaleSetExtensionTypeOutput() VirtualMachineScaleSetExtensionTypeOutput {
	return o
}

func (o VirtualMachineScaleSetExtensionTypeOutput) ToVirtualMachineScaleSetExtensionTypeOutputWithContext(ctx context.Context) VirtualMachineScaleSetExtensionTypeOutput {
	return o
}

func (o VirtualMachineScaleSetExtensionTypeOutput) AutoUpgradeMinorVersion() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetExtensionType) *bool { return v.AutoUpgradeMinorVersion }).(pulumi.BoolPtrOutput)
}

func (o VirtualMachineScaleSetExtensionTypeOutput) ForceUpdateTag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetExtensionType) *string { return v.ForceUpdateTag }).(pulumi.StringPtrOutput)
}

func (o VirtualMachineScaleSetExtensionTypeOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetExtensionType) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o VirtualMachineScaleSetExtensionTypeOutput) ProtectedSettings() pulumi.AnyOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetExtensionType) interface{} { return v.ProtectedSettings }).(pulumi.AnyOutput)
}

func (o VirtualMachineScaleSetExtensionTypeOutput) Publisher() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetExtensionType) *string { return v.Publisher }).(pulumi.StringPtrOutput)
}

func (o VirtualMachineScaleSetExtensionTypeOutput) Settings() pulumi.AnyOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetExtensionType) interface{} { return v.Settings }).(pulumi.AnyOutput)
}

func (o VirtualMachineScaleSetExtensionTypeOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetExtensionType) *string { return v.Type }).(pulumi.StringPtrOutput)
}

func (o VirtualMachineScaleSetExtensionTypeOutput) TypeHandlerVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetExtensionType) *string { return v.TypeHandlerVersion }).(pulumi.StringPtrOutput)
}

type VirtualMachineScaleSetExtensionTypeArrayOutput struct{ *pulumi.OutputState }

func (VirtualMachineScaleSetExtensionTypeArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VirtualMachineScaleSetExtensionType)(nil)).Elem()
}

func (o VirtualMachineScaleSetExtensionTypeArrayOutput) ToVirtualMachineScaleSetExtensionTypeArrayOutput() VirtualMachineScaleSetExtensionTypeArrayOutput {
	return o
}

func (o VirtualMachineScaleSetExtensionTypeArrayOutput) ToVirtualMachineScaleSetExtensionTypeArrayOutputWithContext(ctx context.Context) VirtualMachineScaleSetExtensionTypeArrayOutput {
	return o
}

func (o VirtualMachineScaleSetExtensionTypeArrayOutput) Index(i pulumi.IntInput) VirtualMachineScaleSetExtensionTypeOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VirtualMachineScaleSetExtensionType {
		return vs[0].([]VirtualMachineScaleSetExtensionType)[vs[1].(int)]
	}).(VirtualMachineScaleSetExtensionTypeOutput)
}

type VirtualMachineScaleSetExtensionProfile struct {
	Extensions []VirtualMachineScaleSetExtensionType `pulumi:"extensions"`
}





type VirtualMachineScaleSetExtensionProfileInput interface {
	pulumi.Input

	ToVirtualMachineScaleSetExtensionProfileOutput() VirtualMachineScaleSetExtensionProfileOutput
	ToVirtualMachineScaleSetExtensionProfileOutputWithContext(context.Context) VirtualMachineScaleSetExtensionProfileOutput
}

type VirtualMachineScaleSetExtensionProfileArgs struct {
	Extensions VirtualMachineScaleSetExtensionTypeArrayInput `pulumi:"extensions"`
}

func (VirtualMachineScaleSetExtensionProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualMachineScaleSetExtensionProfile)(nil)).Elem()
}

func (i VirtualMachineScaleSetExtensionProfileArgs) ToVirtualMachineScaleSetExtensionProfileOutput() VirtualMachineScaleSetExtensionProfileOutput {
	return i.ToVirtualMachineScaleSetExtensionProfileOutputWithContext(context.Background())
}

func (i VirtualMachineScaleSetExtensionProfileArgs) ToVirtualMachineScaleSetExtensionProfileOutputWithContext(ctx context.Context) VirtualMachineScaleSetExtensionProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualMachineScaleSetExtensionProfileOutput)
}

func (i VirtualMachineScaleSetExtensionProfileArgs) ToVirtualMachineScaleSetExtensionProfilePtrOutput() VirtualMachineScaleSetExtensionProfilePtrOutput {
	return i.ToVirtualMachineScaleSetExtensionProfilePtrOutputWithContext(context.Background())
}

func (i VirtualMachineScaleSetExtensionProfileArgs) ToVirtualMachineScaleSetExtensionProfilePtrOutputWithContext(ctx context.Context) VirtualMachineScaleSetExtensionProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualMachineScaleSetExtensionProfileOutput).ToVirtualMachineScaleSetExtensionProfilePtrOutputWithContext(ctx)
}









type VirtualMachineScaleSetExtensionProfilePtrInput interface {
	pulumi.Input

	ToVirtualMachineScaleSetExtensionProfilePtrOutput() VirtualMachineScaleSetExtensionProfilePtrOutput
	ToVirtualMachineScaleSetExtensionProfilePtrOutputWithContext(context.Context) VirtualMachineScaleSetExtensionProfilePtrOutput
}

type virtualMachineScaleSetExtensionProfilePtrType VirtualMachineScaleSetExtensionProfileArgs

func VirtualMachineScaleSetExtensionProfilePtr(v *VirtualMachineScaleSetExtensionProfileArgs) VirtualMachineScaleSetExtensionProfilePtrInput {
	return (*virtualMachineScaleSetExtensionProfilePtrType)(v)
}

func (*virtualMachineScaleSetExtensionProfilePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualMachineScaleSetExtensionProfile)(nil)).Elem()
}

func (i *virtualMachineScaleSetExtensionProfilePtrType) ToVirtualMachineScaleSetExtensionProfilePtrOutput() VirtualMachineScaleSetExtensionProfilePtrOutput {
	return i.ToVirtualMachineScaleSetExtensionProfilePtrOutputWithContext(context.Background())
}

func (i *virtualMachineScaleSetExtensionProfilePtrType) ToVirtualMachineScaleSetExtensionProfilePtrOutputWithContext(ctx context.Context) VirtualMachineScaleSetExtensionProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualMachineScaleSetExtensionProfilePtrOutput)
}

type VirtualMachineScaleSetExtensionProfileOutput struct{ *pulumi.OutputState }

func (VirtualMachineScaleSetExtensionProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualMachineScaleSetExtensionProfile)(nil)).Elem()
}

func (o VirtualMachineScaleSetExtensionProfileOutput) ToVirtualMachineScaleSetExtensionProfileOutput() VirtualMachineScaleSetExtensionProfileOutput {
	return o
}

func (o VirtualMachineScaleSetExtensionProfileOutput) ToVirtualMachineScaleSetExtensionProfileOutputWithContext(ctx context.Context) VirtualMachineScaleSetExtensionProfileOutput {
	return o
}

func (o VirtualMachineScaleSetExtensionProfileOutput) ToVirtualMachineScaleSetExtensionProfilePtrOutput() VirtualMachineScaleSetExtensionProfilePtrOutput {
	return o.ToVirtualMachineScaleSetExtensionProfilePtrOutputWithContext(context.Background())
}

func (o VirtualMachineScaleSetExtensionProfileOutput) ToVirtualMachineScaleSetExtensionProfilePtrOutputWithContext(ctx context.Context) VirtualMachineScaleSetExtensionProfilePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v VirtualMachineScaleSetExtensionProfile) *VirtualMachineScaleSetExtensionProfile {
		return &v
	}).(VirtualMachineScaleSetExtensionProfilePtrOutput)
}

func (o VirtualMachineScaleSetExtensionProfileOutput) Extensions() VirtualMachineScaleSetExtensionTypeArrayOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetExtensionProfile) []VirtualMachineScaleSetExtensionType {
		return v.Extensions
	}).(VirtualMachineScaleSetExtensionTypeArrayOutput)
}

type VirtualMachineScaleSetExtensionProfilePtrOutput struct{ *pulumi.OutputState }

func (VirtualMachineScaleSetExtensionProfilePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualMachineScaleSetExtensionProfile)(nil)).Elem()
}

func (o VirtualMachineScaleSetExtensionProfilePtrOutput) ToVirtualMachineScaleSetExtensionProfilePtrOutput() VirtualMachineScaleSetExtensionProfilePtrOutput {
	return o
}

func (o VirtualMachineScaleSetExtensionProfilePtrOutput) ToVirtualMachineScaleSetExtensionProfilePtrOutputWithContext(ctx context.Context) VirtualMachineScaleSetExtensionProfilePtrOutput {
	return o
}

func (o VirtualMachineScaleSetExtensionProfilePtrOutput) Elem() VirtualMachineScaleSetExtensionProfileOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetExtensionProfile) VirtualMachineScaleSetExtensionProfile {
		if v != nil {
			return *v
		}
		var ret VirtualMachineScaleSetExtensionProfile
		return ret
	}).(VirtualMachineScaleSetExtensionProfileOutput)
}

func (o VirtualMachineScaleSetExtensionProfilePtrOutput) Extensions() VirtualMachineScaleSetExtensionTypeArrayOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetExtensionProfile) []VirtualMachineScaleSetExtensionType {
		if v == nil {
			return nil
		}
		return v.Extensions
	}).(VirtualMachineScaleSetExtensionTypeArrayOutput)
}

type VirtualMachineScaleSetExtensionProfileResponse struct {
	Extensions []VirtualMachineScaleSetExtensionResponse `pulumi:"extensions"`
}

type VirtualMachineScaleSetExtensionProfileResponseOutput struct{ *pulumi.OutputState }

func (VirtualMachineScaleSetExtensionProfileResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualMachineScaleSetExtensionProfileResponse)(nil)).Elem()
}

func (o VirtualMachineScaleSetExtensionProfileResponseOutput) ToVirtualMachineScaleSetExtensionProfileResponseOutput() VirtualMachineScaleSetExtensionProfileResponseOutput {
	return o
}

func (o VirtualMachineScaleSetExtensionProfileResponseOutput) ToVirtualMachineScaleSetExtensionProfileResponseOutputWithContext(ctx context.Context) VirtualMachineScaleSetExtensionProfileResponseOutput {
	return o
}

func (o VirtualMachineScaleSetExtensionProfileResponseOutput) Extensions() VirtualMachineScaleSetExtensionResponseArrayOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetExtensionProfileResponse) []VirtualMachineScaleSetExtensionResponse {
		return v.Extensions
	}).(VirtualMachineScaleSetExtensionResponseArrayOutput)
}

type VirtualMachineScaleSetExtensionProfileResponsePtrOutput struct{ *pulumi.OutputState }

func (VirtualMachineScaleSetExtensionProfileResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualMachineScaleSetExtensionProfileResponse)(nil)).Elem()
}

func (o VirtualMachineScaleSetExtensionProfileResponsePtrOutput) ToVirtualMachineScaleSetExtensionProfileResponsePtrOutput() VirtualMachineScaleSetExtensionProfileResponsePtrOutput {
	return o
}

func (o VirtualMachineScaleSetExtensionProfileResponsePtrOutput) ToVirtualMachineScaleSetExtensionProfileResponsePtrOutputWithContext(ctx context.Context) VirtualMachineScaleSetExtensionProfileResponsePtrOutput {
	return o
}

func (o VirtualMachineScaleSetExtensionProfileResponsePtrOutput) Elem() VirtualMachineScaleSetExtensionProfileResponseOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetExtensionProfileResponse) VirtualMachineScaleSetExtensionProfileResponse {
		if v != nil {
			return *v
		}
		var ret VirtualMachineScaleSetExtensionProfileResponse
		return ret
	}).(VirtualMachineScaleSetExtensionProfileResponseOutput)
}

func (o VirtualMachineScaleSetExtensionProfileResponsePtrOutput) Extensions() VirtualMachineScaleSetExtensionResponseArrayOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetExtensionProfileResponse) []VirtualMachineScaleSetExtensionResponse {
		if v == nil {
			return nil
		}
		return v.Extensions
	}).(VirtualMachineScaleSetExtensionResponseArrayOutput)
}

type VirtualMachineScaleSetExtensionResponse struct {
	AutoUpgradeMinorVersion *bool       `pulumi:"autoUpgradeMinorVersion"`
	ForceUpdateTag          *string     `pulumi:"forceUpdateTag"`
	Id                      string      `pulumi:"id"`
	Name                    *string     `pulumi:"name"`
	ProtectedSettings       interface{} `pulumi:"protectedSettings"`
	ProvisioningState       string      `pulumi:"provisioningState"`
	Publisher               *string     `pulumi:"publisher"`
	Settings                interface{} `pulumi:"settings"`
	Type                    *string     `pulumi:"type"`
	TypeHandlerVersion      *string     `pulumi:"typeHandlerVersion"`
}

type VirtualMachineScaleSetExtensionResponseOutput struct{ *pulumi.OutputState }

func (VirtualMachineScaleSetExtensionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualMachineScaleSetExtensionResponse)(nil)).Elem()
}

func (o VirtualMachineScaleSetExtensionResponseOutput) ToVirtualMachineScaleSetExtensionResponseOutput() VirtualMachineScaleSetExtensionResponseOutput {
	return o
}

func (o VirtualMachineScaleSetExtensionResponseOutput) ToVirtualMachineScaleSetExtensionResponseOutputWithContext(ctx context.Context) VirtualMachineScaleSetExtensionResponseOutput {
	return o
}

func (o VirtualMachineScaleSetExtensionResponseOutput) AutoUpgradeMinorVersion() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetExtensionResponse) *bool { return v.AutoUpgradeMinorVersion }).(pulumi.BoolPtrOutput)
}

func (o VirtualMachineScaleSetExtensionResponseOutput) ForceUpdateTag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetExtensionResponse) *string { return v.ForceUpdateTag }).(pulumi.StringPtrOutput)
}

func (o VirtualMachineScaleSetExtensionResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetExtensionResponse) string { return v.Id }).(pulumi.StringOutput)
}

func (o VirtualMachineScaleSetExtensionResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetExtensionResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o VirtualMachineScaleSetExtensionResponseOutput) ProtectedSettings() pulumi.AnyOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetExtensionResponse) interface{} { return v.ProtectedSettings }).(pulumi.AnyOutput)
}

func (o VirtualMachineScaleSetExtensionResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetExtensionResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o VirtualMachineScaleSetExtensionResponseOutput) Publisher() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetExtensionResponse) *string { return v.Publisher }).(pulumi.StringPtrOutput)
}

func (o VirtualMachineScaleSetExtensionResponseOutput) Settings() pulumi.AnyOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetExtensionResponse) interface{} { return v.Settings }).(pulumi.AnyOutput)
}

func (o VirtualMachineScaleSetExtensionResponseOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetExtensionResponse) *string { return v.Type }).(pulumi.StringPtrOutput)
}

func (o VirtualMachineScaleSetExtensionResponseOutput) TypeHandlerVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetExtensionResponse) *string { return v.TypeHandlerVersion }).(pulumi.StringPtrOutput)
}

type VirtualMachineScaleSetExtensionResponseArrayOutput struct{ *pulumi.OutputState }

func (VirtualMachineScaleSetExtensionResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VirtualMachineScaleSetExtensionResponse)(nil)).Elem()
}

func (o VirtualMachineScaleSetExtensionResponseArrayOutput) ToVirtualMachineScaleSetExtensionResponseArrayOutput() VirtualMachineScaleSetExtensionResponseArrayOutput {
	return o
}

func (o VirtualMachineScaleSetExtensionResponseArrayOutput) ToVirtualMachineScaleSetExtensionResponseArrayOutputWithContext(ctx context.Context) VirtualMachineScaleSetExtensionResponseArrayOutput {
	return o
}

func (o VirtualMachineScaleSetExtensionResponseArrayOutput) Index(i pulumi.IntInput) VirtualMachineScaleSetExtensionResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VirtualMachineScaleSetExtensionResponse {
		return vs[0].([]VirtualMachineScaleSetExtensionResponse)[vs[1].(int)]
	}).(VirtualMachineScaleSetExtensionResponseOutput)
}

type VirtualMachineScaleSetIPConfiguration struct {
	ApplicationGatewayBackendAddressPools []SubResource                                       `pulumi:"applicationGatewayBackendAddressPools"`
	Id                                    *string                                             `pulumi:"id"`
	LoadBalancerBackendAddressPools       []SubResource                                       `pulumi:"loadBalancerBackendAddressPools"`
	LoadBalancerInboundNatPools           []SubResource                                       `pulumi:"loadBalancerInboundNatPools"`
	Name                                  string                                              `pulumi:"name"`
	Primary                               *bool                                               `pulumi:"primary"`
	PrivateIPAddressVersion               *string                                             `pulumi:"privateIPAddressVersion"`
	PublicIPAddressConfiguration          *VirtualMachineScaleSetPublicIPAddressConfiguration `pulumi:"publicIPAddressConfiguration"`
	Subnet                                *ApiEntityReference                                 `pulumi:"subnet"`
}





type VirtualMachineScaleSetIPConfigurationInput interface {
	pulumi.Input

	ToVirtualMachineScaleSetIPConfigurationOutput() VirtualMachineScaleSetIPConfigurationOutput
	ToVirtualMachineScaleSetIPConfigurationOutputWithContext(context.Context) VirtualMachineScaleSetIPConfigurationOutput
}

type VirtualMachineScaleSetIPConfigurationArgs struct {
	ApplicationGatewayBackendAddressPools SubResourceArrayInput                                      `pulumi:"applicationGatewayBackendAddressPools"`
	Id                                    pulumi.StringPtrInput                                      `pulumi:"id"`
	LoadBalancerBackendAddressPools       SubResourceArrayInput                                      `pulumi:"loadBalancerBackendAddressPools"`
	LoadBalancerInboundNatPools           SubResourceArrayInput                                      `pulumi:"loadBalancerInboundNatPools"`
	Name                                  pulumi.StringInput                                         `pulumi:"name"`
	Primary                               pulumi.BoolPtrInput                                        `pulumi:"primary"`
	PrivateIPAddressVersion               pulumi.StringPtrInput                                      `pulumi:"privateIPAddressVersion"`
	PublicIPAddressConfiguration          VirtualMachineScaleSetPublicIPAddressConfigurationPtrInput `pulumi:"publicIPAddressConfiguration"`
	Subnet                                ApiEntityReferencePtrInput                                 `pulumi:"subnet"`
}

func (VirtualMachineScaleSetIPConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualMachineScaleSetIPConfiguration)(nil)).Elem()
}

func (i VirtualMachineScaleSetIPConfigurationArgs) ToVirtualMachineScaleSetIPConfigurationOutput() VirtualMachineScaleSetIPConfigurationOutput {
	return i.ToVirtualMachineScaleSetIPConfigurationOutputWithContext(context.Background())
}

func (i VirtualMachineScaleSetIPConfigurationArgs) ToVirtualMachineScaleSetIPConfigurationOutputWithContext(ctx context.Context) VirtualMachineScaleSetIPConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualMachineScaleSetIPConfigurationOutput)
}





type VirtualMachineScaleSetIPConfigurationArrayInput interface {
	pulumi.Input

	ToVirtualMachineScaleSetIPConfigurationArrayOutput() VirtualMachineScaleSetIPConfigurationArrayOutput
	ToVirtualMachineScaleSetIPConfigurationArrayOutputWithContext(context.Context) VirtualMachineScaleSetIPConfigurationArrayOutput
}

type VirtualMachineScaleSetIPConfigurationArray []VirtualMachineScaleSetIPConfigurationInput

func (VirtualMachineScaleSetIPConfigurationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VirtualMachineScaleSetIPConfiguration)(nil)).Elem()
}

func (i VirtualMachineScaleSetIPConfigurationArray) ToVirtualMachineScaleSetIPConfigurationArrayOutput() VirtualMachineScaleSetIPConfigurationArrayOutput {
	return i.ToVirtualMachineScaleSetIPConfigurationArrayOutputWithContext(context.Background())
}

func (i VirtualMachineScaleSetIPConfigurationArray) ToVirtualMachineScaleSetIPConfigurationArrayOutputWithContext(ctx context.Context) VirtualMachineScaleSetIPConfigurationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualMachineScaleSetIPConfigurationArrayOutput)
}

type VirtualMachineScaleSetIPConfigurationOutput struct{ *pulumi.OutputState }

func (VirtualMachineScaleSetIPConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualMachineScaleSetIPConfiguration)(nil)).Elem()
}

func (o VirtualMachineScaleSetIPConfigurationOutput) ToVirtualMachineScaleSetIPConfigurationOutput() VirtualMachineScaleSetIPConfigurationOutput {
	return o
}

func (o VirtualMachineScaleSetIPConfigurationOutput) ToVirtualMachineScaleSetIPConfigurationOutputWithContext(ctx context.Context) VirtualMachineScaleSetIPConfigurationOutput {
	return o
}

func (o VirtualMachineScaleSetIPConfigurationOutput) ApplicationGatewayBackendAddressPools() SubResourceArrayOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetIPConfiguration) []SubResource {
		return v.ApplicationGatewayBackendAddressPools
	}).(SubResourceArrayOutput)
}

func (o VirtualMachineScaleSetIPConfigurationOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetIPConfiguration) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o VirtualMachineScaleSetIPConfigurationOutput) LoadBalancerBackendAddressPools() SubResourceArrayOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetIPConfiguration) []SubResource { return v.LoadBalancerBackendAddressPools }).(SubResourceArrayOutput)
}

func (o VirtualMachineScaleSetIPConfigurationOutput) LoadBalancerInboundNatPools() SubResourceArrayOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetIPConfiguration) []SubResource { return v.LoadBalancerInboundNatPools }).(SubResourceArrayOutput)
}

func (o VirtualMachineScaleSetIPConfigurationOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetIPConfiguration) string { return v.Name }).(pulumi.StringOutput)
}

func (o VirtualMachineScaleSetIPConfigurationOutput) Primary() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetIPConfiguration) *bool { return v.Primary }).(pulumi.BoolPtrOutput)
}

func (o VirtualMachineScaleSetIPConfigurationOutput) PrivateIPAddressVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetIPConfiguration) *string { return v.PrivateIPAddressVersion }).(pulumi.StringPtrOutput)
}

func (o VirtualMachineScaleSetIPConfigurationOutput) PublicIPAddressConfiguration() VirtualMachineScaleSetPublicIPAddressConfigurationPtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetIPConfiguration) *VirtualMachineScaleSetPublicIPAddressConfiguration {
		return v.PublicIPAddressConfiguration
	}).(VirtualMachineScaleSetPublicIPAddressConfigurationPtrOutput)
}

func (o VirtualMachineScaleSetIPConfigurationOutput) Subnet() ApiEntityReferencePtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetIPConfiguration) *ApiEntityReference { return v.Subnet }).(ApiEntityReferencePtrOutput)
}

type VirtualMachineScaleSetIPConfigurationArrayOutput struct{ *pulumi.OutputState }

func (VirtualMachineScaleSetIPConfigurationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VirtualMachineScaleSetIPConfiguration)(nil)).Elem()
}

func (o VirtualMachineScaleSetIPConfigurationArrayOutput) ToVirtualMachineScaleSetIPConfigurationArrayOutput() VirtualMachineScaleSetIPConfigurationArrayOutput {
	return o
}

func (o VirtualMachineScaleSetIPConfigurationArrayOutput) ToVirtualMachineScaleSetIPConfigurationArrayOutputWithContext(ctx context.Context) VirtualMachineScaleSetIPConfigurationArrayOutput {
	return o
}

func (o VirtualMachineScaleSetIPConfigurationArrayOutput) Index(i pulumi.IntInput) VirtualMachineScaleSetIPConfigurationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VirtualMachineScaleSetIPConfiguration {
		return vs[0].([]VirtualMachineScaleSetIPConfiguration)[vs[1].(int)]
	}).(VirtualMachineScaleSetIPConfigurationOutput)
}

type VirtualMachineScaleSetIPConfigurationResponse struct {
	ApplicationGatewayBackendAddressPools []SubResourceResponse                                       `pulumi:"applicationGatewayBackendAddressPools"`
	Id                                    *string                                                     `pulumi:"id"`
	LoadBalancerBackendAddressPools       []SubResourceResponse                                       `pulumi:"loadBalancerBackendAddressPools"`
	LoadBalancerInboundNatPools           []SubResourceResponse                                       `pulumi:"loadBalancerInboundNatPools"`
	Name                                  string                                                      `pulumi:"name"`
	Primary                               *bool                                                       `pulumi:"primary"`
	PrivateIPAddressVersion               *string                                                     `pulumi:"privateIPAddressVersion"`
	PublicIPAddressConfiguration          *VirtualMachineScaleSetPublicIPAddressConfigurationResponse `pulumi:"publicIPAddressConfiguration"`
	Subnet                                *ApiEntityReferenceResponse                                 `pulumi:"subnet"`
}

type VirtualMachineScaleSetIPConfigurationResponseOutput struct{ *pulumi.OutputState }

func (VirtualMachineScaleSetIPConfigurationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualMachineScaleSetIPConfigurationResponse)(nil)).Elem()
}

func (o VirtualMachineScaleSetIPConfigurationResponseOutput) ToVirtualMachineScaleSetIPConfigurationResponseOutput() VirtualMachineScaleSetIPConfigurationResponseOutput {
	return o
}

func (o VirtualMachineScaleSetIPConfigurationResponseOutput) ToVirtualMachineScaleSetIPConfigurationResponseOutputWithContext(ctx context.Context) VirtualMachineScaleSetIPConfigurationResponseOutput {
	return o
}

func (o VirtualMachineScaleSetIPConfigurationResponseOutput) ApplicationGatewayBackendAddressPools() SubResourceResponseArrayOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetIPConfigurationResponse) []SubResourceResponse {
		return v.ApplicationGatewayBackendAddressPools
	}).(SubResourceResponseArrayOutput)
}

func (o VirtualMachineScaleSetIPConfigurationResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetIPConfigurationResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o VirtualMachineScaleSetIPConfigurationResponseOutput) LoadBalancerBackendAddressPools() SubResourceResponseArrayOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetIPConfigurationResponse) []SubResourceResponse {
		return v.LoadBalancerBackendAddressPools
	}).(SubResourceResponseArrayOutput)
}

func (o VirtualMachineScaleSetIPConfigurationResponseOutput) LoadBalancerInboundNatPools() SubResourceResponseArrayOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetIPConfigurationResponse) []SubResourceResponse {
		return v.LoadBalancerInboundNatPools
	}).(SubResourceResponseArrayOutput)
}

func (o VirtualMachineScaleSetIPConfigurationResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetIPConfigurationResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o VirtualMachineScaleSetIPConfigurationResponseOutput) Primary() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetIPConfigurationResponse) *bool { return v.Primary }).(pulumi.BoolPtrOutput)
}

func (o VirtualMachineScaleSetIPConfigurationResponseOutput) PrivateIPAddressVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetIPConfigurationResponse) *string { return v.PrivateIPAddressVersion }).(pulumi.StringPtrOutput)
}

func (o VirtualMachineScaleSetIPConfigurationResponseOutput) PublicIPAddressConfiguration() VirtualMachineScaleSetPublicIPAddressConfigurationResponsePtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetIPConfigurationResponse) *VirtualMachineScaleSetPublicIPAddressConfigurationResponse {
		return v.PublicIPAddressConfiguration
	}).(VirtualMachineScaleSetPublicIPAddressConfigurationResponsePtrOutput)
}

func (o VirtualMachineScaleSetIPConfigurationResponseOutput) Subnet() ApiEntityReferenceResponsePtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetIPConfigurationResponse) *ApiEntityReferenceResponse { return v.Subnet }).(ApiEntityReferenceResponsePtrOutput)
}

type VirtualMachineScaleSetIPConfigurationResponseArrayOutput struct{ *pulumi.OutputState }

func (VirtualMachineScaleSetIPConfigurationResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VirtualMachineScaleSetIPConfigurationResponse)(nil)).Elem()
}

func (o VirtualMachineScaleSetIPConfigurationResponseArrayOutput) ToVirtualMachineScaleSetIPConfigurationResponseArrayOutput() VirtualMachineScaleSetIPConfigurationResponseArrayOutput {
	return o
}

func (o VirtualMachineScaleSetIPConfigurationResponseArrayOutput) ToVirtualMachineScaleSetIPConfigurationResponseArrayOutputWithContext(ctx context.Context) VirtualMachineScaleSetIPConfigurationResponseArrayOutput {
	return o
}

func (o VirtualMachineScaleSetIPConfigurationResponseArrayOutput) Index(i pulumi.IntInput) VirtualMachineScaleSetIPConfigurationResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VirtualMachineScaleSetIPConfigurationResponse {
		return vs[0].([]VirtualMachineScaleSetIPConfigurationResponse)[vs[1].(int)]
	}).(VirtualMachineScaleSetIPConfigurationResponseOutput)
}

type VirtualMachineScaleSetIdentity struct {
	IdentityIds []string              `pulumi:"identityIds"`
	Type        *ResourceIdentityType `pulumi:"type"`
}





type VirtualMachineScaleSetIdentityInput interface {
	pulumi.Input

	ToVirtualMachineScaleSetIdentityOutput() VirtualMachineScaleSetIdentityOutput
	ToVirtualMachineScaleSetIdentityOutputWithContext(context.Context) VirtualMachineScaleSetIdentityOutput
}

type VirtualMachineScaleSetIdentityArgs struct {
	IdentityIds pulumi.StringArrayInput      `pulumi:"identityIds"`
	Type        ResourceIdentityTypePtrInput `pulumi:"type"`
}

func (VirtualMachineScaleSetIdentityArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualMachineScaleSetIdentity)(nil)).Elem()
}

func (i VirtualMachineScaleSetIdentityArgs) ToVirtualMachineScaleSetIdentityOutput() VirtualMachineScaleSetIdentityOutput {
	return i.ToVirtualMachineScaleSetIdentityOutputWithContext(context.Background())
}

func (i VirtualMachineScaleSetIdentityArgs) ToVirtualMachineScaleSetIdentityOutputWithContext(ctx context.Context) VirtualMachineScaleSetIdentityOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualMachineScaleSetIdentityOutput)
}

func (i VirtualMachineScaleSetIdentityArgs) ToVirtualMachineScaleSetIdentityPtrOutput() VirtualMachineScaleSetIdentityPtrOutput {
	return i.ToVirtualMachineScaleSetIdentityPtrOutputWithContext(context.Background())
}

func (i VirtualMachineScaleSetIdentityArgs) ToVirtualMachineScaleSetIdentityPtrOutputWithContext(ctx context.Context) VirtualMachineScaleSetIdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualMachineScaleSetIdentityOutput).ToVirtualMachineScaleSetIdentityPtrOutputWithContext(ctx)
}









type VirtualMachineScaleSetIdentityPtrInput interface {
	pulumi.Input

	ToVirtualMachineScaleSetIdentityPtrOutput() VirtualMachineScaleSetIdentityPtrOutput
	ToVirtualMachineScaleSetIdentityPtrOutputWithContext(context.Context) VirtualMachineScaleSetIdentityPtrOutput
}

type virtualMachineScaleSetIdentityPtrType VirtualMachineScaleSetIdentityArgs

func VirtualMachineScaleSetIdentityPtr(v *VirtualMachineScaleSetIdentityArgs) VirtualMachineScaleSetIdentityPtrInput {
	return (*virtualMachineScaleSetIdentityPtrType)(v)
}

func (*virtualMachineScaleSetIdentityPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualMachineScaleSetIdentity)(nil)).Elem()
}

func (i *virtualMachineScaleSetIdentityPtrType) ToVirtualMachineScaleSetIdentityPtrOutput() VirtualMachineScaleSetIdentityPtrOutput {
	return i.ToVirtualMachineScaleSetIdentityPtrOutputWithContext(context.Background())
}

func (i *virtualMachineScaleSetIdentityPtrType) ToVirtualMachineScaleSetIdentityPtrOutputWithContext(ctx context.Context) VirtualMachineScaleSetIdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualMachineScaleSetIdentityPtrOutput)
}

type VirtualMachineScaleSetIdentityOutput struct{ *pulumi.OutputState }

func (VirtualMachineScaleSetIdentityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualMachineScaleSetIdentity)(nil)).Elem()
}

func (o VirtualMachineScaleSetIdentityOutput) ToVirtualMachineScaleSetIdentityOutput() VirtualMachineScaleSetIdentityOutput {
	return o
}

func (o VirtualMachineScaleSetIdentityOutput) ToVirtualMachineScaleSetIdentityOutputWithContext(ctx context.Context) VirtualMachineScaleSetIdentityOutput {
	return o
}

func (o VirtualMachineScaleSetIdentityOutput) ToVirtualMachineScaleSetIdentityPtrOutput() VirtualMachineScaleSetIdentityPtrOutput {
	return o.ToVirtualMachineScaleSetIdentityPtrOutputWithContext(context.Background())
}

func (o VirtualMachineScaleSetIdentityOutput) ToVirtualMachineScaleSetIdentityPtrOutputWithContext(ctx context.Context) VirtualMachineScaleSetIdentityPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v VirtualMachineScaleSetIdentity) *VirtualMachineScaleSetIdentity {
		return &v
	}).(VirtualMachineScaleSetIdentityPtrOutput)
}

func (o VirtualMachineScaleSetIdentityOutput) IdentityIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetIdentity) []string { return v.IdentityIds }).(pulumi.StringArrayOutput)
}

func (o VirtualMachineScaleSetIdentityOutput) Type() ResourceIdentityTypePtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetIdentity) *ResourceIdentityType { return v.Type }).(ResourceIdentityTypePtrOutput)
}

type VirtualMachineScaleSetIdentityPtrOutput struct{ *pulumi.OutputState }

func (VirtualMachineScaleSetIdentityPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualMachineScaleSetIdentity)(nil)).Elem()
}

func (o VirtualMachineScaleSetIdentityPtrOutput) ToVirtualMachineScaleSetIdentityPtrOutput() VirtualMachineScaleSetIdentityPtrOutput {
	return o
}

func (o VirtualMachineScaleSetIdentityPtrOutput) ToVirtualMachineScaleSetIdentityPtrOutputWithContext(ctx context.Context) VirtualMachineScaleSetIdentityPtrOutput {
	return o
}

func (o VirtualMachineScaleSetIdentityPtrOutput) Elem() VirtualMachineScaleSetIdentityOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetIdentity) VirtualMachineScaleSetIdentity {
		if v != nil {
			return *v
		}
		var ret VirtualMachineScaleSetIdentity
		return ret
	}).(VirtualMachineScaleSetIdentityOutput)
}

func (o VirtualMachineScaleSetIdentityPtrOutput) IdentityIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetIdentity) []string {
		if v == nil {
			return nil
		}
		return v.IdentityIds
	}).(pulumi.StringArrayOutput)
}

func (o VirtualMachineScaleSetIdentityPtrOutput) Type() ResourceIdentityTypePtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetIdentity) *ResourceIdentityType {
		if v == nil {
			return nil
		}
		return v.Type
	}).(ResourceIdentityTypePtrOutput)
}

type VirtualMachineScaleSetIdentityResponse struct {
	IdentityIds []string `pulumi:"identityIds"`
	PrincipalId string   `pulumi:"principalId"`
	TenantId    string   `pulumi:"tenantId"`
	Type        *string  `pulumi:"type"`
}

type VirtualMachineScaleSetIdentityResponseOutput struct{ *pulumi.OutputState }

func (VirtualMachineScaleSetIdentityResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualMachineScaleSetIdentityResponse)(nil)).Elem()
}

func (o VirtualMachineScaleSetIdentityResponseOutput) ToVirtualMachineScaleSetIdentityResponseOutput() VirtualMachineScaleSetIdentityResponseOutput {
	return o
}

func (o VirtualMachineScaleSetIdentityResponseOutput) ToVirtualMachineScaleSetIdentityResponseOutputWithContext(ctx context.Context) VirtualMachineScaleSetIdentityResponseOutput {
	return o
}

func (o VirtualMachineScaleSetIdentityResponseOutput) IdentityIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetIdentityResponse) []string { return v.IdentityIds }).(pulumi.StringArrayOutput)
}

func (o VirtualMachineScaleSetIdentityResponseOutput) PrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetIdentityResponse) string { return v.PrincipalId }).(pulumi.StringOutput)
}

func (o VirtualMachineScaleSetIdentityResponseOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetIdentityResponse) string { return v.TenantId }).(pulumi.StringOutput)
}

func (o VirtualMachineScaleSetIdentityResponseOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetIdentityResponse) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type VirtualMachineScaleSetIdentityResponsePtrOutput struct{ *pulumi.OutputState }

func (VirtualMachineScaleSetIdentityResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualMachineScaleSetIdentityResponse)(nil)).Elem()
}

func (o VirtualMachineScaleSetIdentityResponsePtrOutput) ToVirtualMachineScaleSetIdentityResponsePtrOutput() VirtualMachineScaleSetIdentityResponsePtrOutput {
	return o
}

func (o VirtualMachineScaleSetIdentityResponsePtrOutput) ToVirtualMachineScaleSetIdentityResponsePtrOutputWithContext(ctx context.Context) VirtualMachineScaleSetIdentityResponsePtrOutput {
	return o
}

func (o VirtualMachineScaleSetIdentityResponsePtrOutput) Elem() VirtualMachineScaleSetIdentityResponseOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetIdentityResponse) VirtualMachineScaleSetIdentityResponse {
		if v != nil {
			return *v
		}
		var ret VirtualMachineScaleSetIdentityResponse
		return ret
	}).(VirtualMachineScaleSetIdentityResponseOutput)
}

func (o VirtualMachineScaleSetIdentityResponsePtrOutput) IdentityIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetIdentityResponse) []string {
		if v == nil {
			return nil
		}
		return v.IdentityIds
	}).(pulumi.StringArrayOutput)
}

func (o VirtualMachineScaleSetIdentityResponsePtrOutput) PrincipalId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetIdentityResponse) *string {
		if v == nil {
			return nil
		}
		return &v.PrincipalId
	}).(pulumi.StringPtrOutput)
}

func (o VirtualMachineScaleSetIdentityResponsePtrOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetIdentityResponse) *string {
		if v == nil {
			return nil
		}
		return &v.TenantId
	}).(pulumi.StringPtrOutput)
}

func (o VirtualMachineScaleSetIdentityResponsePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetIdentityResponse) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

type VirtualMachineScaleSetIpTag struct {
	IpTagType *string `pulumi:"ipTagType"`
	Tag       *string `pulumi:"tag"`
}





type VirtualMachineScaleSetIpTagInput interface {
	pulumi.Input

	ToVirtualMachineScaleSetIpTagOutput() VirtualMachineScaleSetIpTagOutput
	ToVirtualMachineScaleSetIpTagOutputWithContext(context.Context) VirtualMachineScaleSetIpTagOutput
}

type VirtualMachineScaleSetIpTagArgs struct {
	IpTagType pulumi.StringPtrInput `pulumi:"ipTagType"`
	Tag       pulumi.StringPtrInput `pulumi:"tag"`
}

func (VirtualMachineScaleSetIpTagArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualMachineScaleSetIpTag)(nil)).Elem()
}

func (i VirtualMachineScaleSetIpTagArgs) ToVirtualMachineScaleSetIpTagOutput() VirtualMachineScaleSetIpTagOutput {
	return i.ToVirtualMachineScaleSetIpTagOutputWithContext(context.Background())
}

func (i VirtualMachineScaleSetIpTagArgs) ToVirtualMachineScaleSetIpTagOutputWithContext(ctx context.Context) VirtualMachineScaleSetIpTagOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualMachineScaleSetIpTagOutput)
}





type VirtualMachineScaleSetIpTagArrayInput interface {
	pulumi.Input

	ToVirtualMachineScaleSetIpTagArrayOutput() VirtualMachineScaleSetIpTagArrayOutput
	ToVirtualMachineScaleSetIpTagArrayOutputWithContext(context.Context) VirtualMachineScaleSetIpTagArrayOutput
}

type VirtualMachineScaleSetIpTagArray []VirtualMachineScaleSetIpTagInput

func (VirtualMachineScaleSetIpTagArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VirtualMachineScaleSetIpTag)(nil)).Elem()
}

func (i VirtualMachineScaleSetIpTagArray) ToVirtualMachineScaleSetIpTagArrayOutput() VirtualMachineScaleSetIpTagArrayOutput {
	return i.ToVirtualMachineScaleSetIpTagArrayOutputWithContext(context.Background())
}

func (i VirtualMachineScaleSetIpTagArray) ToVirtualMachineScaleSetIpTagArrayOutputWithContext(ctx context.Context) VirtualMachineScaleSetIpTagArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualMachineScaleSetIpTagArrayOutput)
}

type VirtualMachineScaleSetIpTagOutput struct{ *pulumi.OutputState }

func (VirtualMachineScaleSetIpTagOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualMachineScaleSetIpTag)(nil)).Elem()
}

func (o VirtualMachineScaleSetIpTagOutput) ToVirtualMachineScaleSetIpTagOutput() VirtualMachineScaleSetIpTagOutput {
	return o
}

func (o VirtualMachineScaleSetIpTagOutput) ToVirtualMachineScaleSetIpTagOutputWithContext(ctx context.Context) VirtualMachineScaleSetIpTagOutput {
	return o
}

func (o VirtualMachineScaleSetIpTagOutput) IpTagType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetIpTag) *string { return v.IpTagType }).(pulumi.StringPtrOutput)
}

func (o VirtualMachineScaleSetIpTagOutput) Tag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetIpTag) *string { return v.Tag }).(pulumi.StringPtrOutput)
}

type VirtualMachineScaleSetIpTagArrayOutput struct{ *pulumi.OutputState }

func (VirtualMachineScaleSetIpTagArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VirtualMachineScaleSetIpTag)(nil)).Elem()
}

func (o VirtualMachineScaleSetIpTagArrayOutput) ToVirtualMachineScaleSetIpTagArrayOutput() VirtualMachineScaleSetIpTagArrayOutput {
	return o
}

func (o VirtualMachineScaleSetIpTagArrayOutput) ToVirtualMachineScaleSetIpTagArrayOutputWithContext(ctx context.Context) VirtualMachineScaleSetIpTagArrayOutput {
	return o
}

func (o VirtualMachineScaleSetIpTagArrayOutput) Index(i pulumi.IntInput) VirtualMachineScaleSetIpTagOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VirtualMachineScaleSetIpTag {
		return vs[0].([]VirtualMachineScaleSetIpTag)[vs[1].(int)]
	}).(VirtualMachineScaleSetIpTagOutput)
}

type VirtualMachineScaleSetIpTagResponse struct {
	IpTagType *string `pulumi:"ipTagType"`
	Tag       *string `pulumi:"tag"`
}

type VirtualMachineScaleSetIpTagResponseOutput struct{ *pulumi.OutputState }

func (VirtualMachineScaleSetIpTagResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualMachineScaleSetIpTagResponse)(nil)).Elem()
}

func (o VirtualMachineScaleSetIpTagResponseOutput) ToVirtualMachineScaleSetIpTagResponseOutput() VirtualMachineScaleSetIpTagResponseOutput {
	return o
}

func (o VirtualMachineScaleSetIpTagResponseOutput) ToVirtualMachineScaleSetIpTagResponseOutputWithContext(ctx context.Context) VirtualMachineScaleSetIpTagResponseOutput {
	return o
}

func (o VirtualMachineScaleSetIpTagResponseOutput) IpTagType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetIpTagResponse) *string { return v.IpTagType }).(pulumi.StringPtrOutput)
}

func (o VirtualMachineScaleSetIpTagResponseOutput) Tag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetIpTagResponse) *string { return v.Tag }).(pulumi.StringPtrOutput)
}

type VirtualMachineScaleSetIpTagResponseArrayOutput struct{ *pulumi.OutputState }

func (VirtualMachineScaleSetIpTagResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VirtualMachineScaleSetIpTagResponse)(nil)).Elem()
}

func (o VirtualMachineScaleSetIpTagResponseArrayOutput) ToVirtualMachineScaleSetIpTagResponseArrayOutput() VirtualMachineScaleSetIpTagResponseArrayOutput {
	return o
}

func (o VirtualMachineScaleSetIpTagResponseArrayOutput) ToVirtualMachineScaleSetIpTagResponseArrayOutputWithContext(ctx context.Context) VirtualMachineScaleSetIpTagResponseArrayOutput {
	return o
}

func (o VirtualMachineScaleSetIpTagResponseArrayOutput) Index(i pulumi.IntInput) VirtualMachineScaleSetIpTagResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VirtualMachineScaleSetIpTagResponse {
		return vs[0].([]VirtualMachineScaleSetIpTagResponse)[vs[1].(int)]
	}).(VirtualMachineScaleSetIpTagResponseOutput)
}

type VirtualMachineScaleSetManagedDiskParameters struct {
	StorageAccountType *string `pulumi:"storageAccountType"`
}





type VirtualMachineScaleSetManagedDiskParametersInput interface {
	pulumi.Input

	ToVirtualMachineScaleSetManagedDiskParametersOutput() VirtualMachineScaleSetManagedDiskParametersOutput
	ToVirtualMachineScaleSetManagedDiskParametersOutputWithContext(context.Context) VirtualMachineScaleSetManagedDiskParametersOutput
}

type VirtualMachineScaleSetManagedDiskParametersArgs struct {
	StorageAccountType pulumi.StringPtrInput `pulumi:"storageAccountType"`
}

func (VirtualMachineScaleSetManagedDiskParametersArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualMachineScaleSetManagedDiskParameters)(nil)).Elem()
}

func (i VirtualMachineScaleSetManagedDiskParametersArgs) ToVirtualMachineScaleSetManagedDiskParametersOutput() VirtualMachineScaleSetManagedDiskParametersOutput {
	return i.ToVirtualMachineScaleSetManagedDiskParametersOutputWithContext(context.Background())
}

func (i VirtualMachineScaleSetManagedDiskParametersArgs) ToVirtualMachineScaleSetManagedDiskParametersOutputWithContext(ctx context.Context) VirtualMachineScaleSetManagedDiskParametersOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualMachineScaleSetManagedDiskParametersOutput)
}

func (i VirtualMachineScaleSetManagedDiskParametersArgs) ToVirtualMachineScaleSetManagedDiskParametersPtrOutput() VirtualMachineScaleSetManagedDiskParametersPtrOutput {
	return i.ToVirtualMachineScaleSetManagedDiskParametersPtrOutputWithContext(context.Background())
}

func (i VirtualMachineScaleSetManagedDiskParametersArgs) ToVirtualMachineScaleSetManagedDiskParametersPtrOutputWithContext(ctx context.Context) VirtualMachineScaleSetManagedDiskParametersPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualMachineScaleSetManagedDiskParametersOutput).ToVirtualMachineScaleSetManagedDiskParametersPtrOutputWithContext(ctx)
}









type VirtualMachineScaleSetManagedDiskParametersPtrInput interface {
	pulumi.Input

	ToVirtualMachineScaleSetManagedDiskParametersPtrOutput() VirtualMachineScaleSetManagedDiskParametersPtrOutput
	ToVirtualMachineScaleSetManagedDiskParametersPtrOutputWithContext(context.Context) VirtualMachineScaleSetManagedDiskParametersPtrOutput
}

type virtualMachineScaleSetManagedDiskParametersPtrType VirtualMachineScaleSetManagedDiskParametersArgs

func VirtualMachineScaleSetManagedDiskParametersPtr(v *VirtualMachineScaleSetManagedDiskParametersArgs) VirtualMachineScaleSetManagedDiskParametersPtrInput {
	return (*virtualMachineScaleSetManagedDiskParametersPtrType)(v)
}

func (*virtualMachineScaleSetManagedDiskParametersPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualMachineScaleSetManagedDiskParameters)(nil)).Elem()
}

func (i *virtualMachineScaleSetManagedDiskParametersPtrType) ToVirtualMachineScaleSetManagedDiskParametersPtrOutput() VirtualMachineScaleSetManagedDiskParametersPtrOutput {
	return i.ToVirtualMachineScaleSetManagedDiskParametersPtrOutputWithContext(context.Background())
}

func (i *virtualMachineScaleSetManagedDiskParametersPtrType) ToVirtualMachineScaleSetManagedDiskParametersPtrOutputWithContext(ctx context.Context) VirtualMachineScaleSetManagedDiskParametersPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualMachineScaleSetManagedDiskParametersPtrOutput)
}

type VirtualMachineScaleSetManagedDiskParametersOutput struct{ *pulumi.OutputState }

func (VirtualMachineScaleSetManagedDiskParametersOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualMachineScaleSetManagedDiskParameters)(nil)).Elem()
}

func (o VirtualMachineScaleSetManagedDiskParametersOutput) ToVirtualMachineScaleSetManagedDiskParametersOutput() VirtualMachineScaleSetManagedDiskParametersOutput {
	return o
}

func (o VirtualMachineScaleSetManagedDiskParametersOutput) ToVirtualMachineScaleSetManagedDiskParametersOutputWithContext(ctx context.Context) VirtualMachineScaleSetManagedDiskParametersOutput {
	return o
}

func (o VirtualMachineScaleSetManagedDiskParametersOutput) ToVirtualMachineScaleSetManagedDiskParametersPtrOutput() VirtualMachineScaleSetManagedDiskParametersPtrOutput {
	return o.ToVirtualMachineScaleSetManagedDiskParametersPtrOutputWithContext(context.Background())
}

func (o VirtualMachineScaleSetManagedDiskParametersOutput) ToVirtualMachineScaleSetManagedDiskParametersPtrOutputWithContext(ctx context.Context) VirtualMachineScaleSetManagedDiskParametersPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v VirtualMachineScaleSetManagedDiskParameters) *VirtualMachineScaleSetManagedDiskParameters {
		return &v
	}).(VirtualMachineScaleSetManagedDiskParametersPtrOutput)
}

func (o VirtualMachineScaleSetManagedDiskParametersOutput) StorageAccountType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetManagedDiskParameters) *string { return v.StorageAccountType }).(pulumi.StringPtrOutput)
}

type VirtualMachineScaleSetManagedDiskParametersPtrOutput struct{ *pulumi.OutputState }

func (VirtualMachineScaleSetManagedDiskParametersPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualMachineScaleSetManagedDiskParameters)(nil)).Elem()
}

func (o VirtualMachineScaleSetManagedDiskParametersPtrOutput) ToVirtualMachineScaleSetManagedDiskParametersPtrOutput() VirtualMachineScaleSetManagedDiskParametersPtrOutput {
	return o
}

func (o VirtualMachineScaleSetManagedDiskParametersPtrOutput) ToVirtualMachineScaleSetManagedDiskParametersPtrOutputWithContext(ctx context.Context) VirtualMachineScaleSetManagedDiskParametersPtrOutput {
	return o
}

func (o VirtualMachineScaleSetManagedDiskParametersPtrOutput) Elem() VirtualMachineScaleSetManagedDiskParametersOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetManagedDiskParameters) VirtualMachineScaleSetManagedDiskParameters {
		if v != nil {
			return *v
		}
		var ret VirtualMachineScaleSetManagedDiskParameters
		return ret
	}).(VirtualMachineScaleSetManagedDiskParametersOutput)
}

func (o VirtualMachineScaleSetManagedDiskParametersPtrOutput) StorageAccountType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetManagedDiskParameters) *string {
		if v == nil {
			return nil
		}
		return v.StorageAccountType
	}).(pulumi.StringPtrOutput)
}

type VirtualMachineScaleSetManagedDiskParametersResponse struct {
	StorageAccountType *string `pulumi:"storageAccountType"`
}

type VirtualMachineScaleSetManagedDiskParametersResponseOutput struct{ *pulumi.OutputState }

func (VirtualMachineScaleSetManagedDiskParametersResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualMachineScaleSetManagedDiskParametersResponse)(nil)).Elem()
}

func (o VirtualMachineScaleSetManagedDiskParametersResponseOutput) ToVirtualMachineScaleSetManagedDiskParametersResponseOutput() VirtualMachineScaleSetManagedDiskParametersResponseOutput {
	return o
}

func (o VirtualMachineScaleSetManagedDiskParametersResponseOutput) ToVirtualMachineScaleSetManagedDiskParametersResponseOutputWithContext(ctx context.Context) VirtualMachineScaleSetManagedDiskParametersResponseOutput {
	return o
}

func (o VirtualMachineScaleSetManagedDiskParametersResponseOutput) StorageAccountType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetManagedDiskParametersResponse) *string { return v.StorageAccountType }).(pulumi.StringPtrOutput)
}

type VirtualMachineScaleSetManagedDiskParametersResponsePtrOutput struct{ *pulumi.OutputState }

func (VirtualMachineScaleSetManagedDiskParametersResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualMachineScaleSetManagedDiskParametersResponse)(nil)).Elem()
}

func (o VirtualMachineScaleSetManagedDiskParametersResponsePtrOutput) ToVirtualMachineScaleSetManagedDiskParametersResponsePtrOutput() VirtualMachineScaleSetManagedDiskParametersResponsePtrOutput {
	return o
}

func (o VirtualMachineScaleSetManagedDiskParametersResponsePtrOutput) ToVirtualMachineScaleSetManagedDiskParametersResponsePtrOutputWithContext(ctx context.Context) VirtualMachineScaleSetManagedDiskParametersResponsePtrOutput {
	return o
}

func (o VirtualMachineScaleSetManagedDiskParametersResponsePtrOutput) Elem() VirtualMachineScaleSetManagedDiskParametersResponseOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetManagedDiskParametersResponse) VirtualMachineScaleSetManagedDiskParametersResponse {
		if v != nil {
			return *v
		}
		var ret VirtualMachineScaleSetManagedDiskParametersResponse
		return ret
	}).(VirtualMachineScaleSetManagedDiskParametersResponseOutput)
}

func (o VirtualMachineScaleSetManagedDiskParametersResponsePtrOutput) StorageAccountType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetManagedDiskParametersResponse) *string {
		if v == nil {
			return nil
		}
		return v.StorageAccountType
	}).(pulumi.StringPtrOutput)
}

type VirtualMachineScaleSetNetworkConfiguration struct {
	DnsSettings                 *VirtualMachineScaleSetNetworkConfigurationDnsSettings `pulumi:"dnsSettings"`
	EnableAcceleratedNetworking *bool                                                  `pulumi:"enableAcceleratedNetworking"`
	EnableIPForwarding          *bool                                                  `pulumi:"enableIPForwarding"`
	Id                          *string                                                `pulumi:"id"`
	IpConfigurations            []VirtualMachineScaleSetIPConfiguration                `pulumi:"ipConfigurations"`
	Name                        string                                                 `pulumi:"name"`
	NetworkSecurityGroup        *SubResource                                           `pulumi:"networkSecurityGroup"`
	Primary                     *bool                                                  `pulumi:"primary"`
}





type VirtualMachineScaleSetNetworkConfigurationInput interface {
	pulumi.Input

	ToVirtualMachineScaleSetNetworkConfigurationOutput() VirtualMachineScaleSetNetworkConfigurationOutput
	ToVirtualMachineScaleSetNetworkConfigurationOutputWithContext(context.Context) VirtualMachineScaleSetNetworkConfigurationOutput
}

type VirtualMachineScaleSetNetworkConfigurationArgs struct {
	DnsSettings                 VirtualMachineScaleSetNetworkConfigurationDnsSettingsPtrInput `pulumi:"dnsSettings"`
	EnableAcceleratedNetworking pulumi.BoolPtrInput                                           `pulumi:"enableAcceleratedNetworking"`
	EnableIPForwarding          pulumi.BoolPtrInput                                           `pulumi:"enableIPForwarding"`
	Id                          pulumi.StringPtrInput                                         `pulumi:"id"`
	IpConfigurations            VirtualMachineScaleSetIPConfigurationArrayInput               `pulumi:"ipConfigurations"`
	Name                        pulumi.StringInput                                            `pulumi:"name"`
	NetworkSecurityGroup        SubResourcePtrInput                                           `pulumi:"networkSecurityGroup"`
	Primary                     pulumi.BoolPtrInput                                           `pulumi:"primary"`
}

func (VirtualMachineScaleSetNetworkConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualMachineScaleSetNetworkConfiguration)(nil)).Elem()
}

func (i VirtualMachineScaleSetNetworkConfigurationArgs) ToVirtualMachineScaleSetNetworkConfigurationOutput() VirtualMachineScaleSetNetworkConfigurationOutput {
	return i.ToVirtualMachineScaleSetNetworkConfigurationOutputWithContext(context.Background())
}

func (i VirtualMachineScaleSetNetworkConfigurationArgs) ToVirtualMachineScaleSetNetworkConfigurationOutputWithContext(ctx context.Context) VirtualMachineScaleSetNetworkConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualMachineScaleSetNetworkConfigurationOutput)
}





type VirtualMachineScaleSetNetworkConfigurationArrayInput interface {
	pulumi.Input

	ToVirtualMachineScaleSetNetworkConfigurationArrayOutput() VirtualMachineScaleSetNetworkConfigurationArrayOutput
	ToVirtualMachineScaleSetNetworkConfigurationArrayOutputWithContext(context.Context) VirtualMachineScaleSetNetworkConfigurationArrayOutput
}

type VirtualMachineScaleSetNetworkConfigurationArray []VirtualMachineScaleSetNetworkConfigurationInput

func (VirtualMachineScaleSetNetworkConfigurationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VirtualMachineScaleSetNetworkConfiguration)(nil)).Elem()
}

func (i VirtualMachineScaleSetNetworkConfigurationArray) ToVirtualMachineScaleSetNetworkConfigurationArrayOutput() VirtualMachineScaleSetNetworkConfigurationArrayOutput {
	return i.ToVirtualMachineScaleSetNetworkConfigurationArrayOutputWithContext(context.Background())
}

func (i VirtualMachineScaleSetNetworkConfigurationArray) ToVirtualMachineScaleSetNetworkConfigurationArrayOutputWithContext(ctx context.Context) VirtualMachineScaleSetNetworkConfigurationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualMachineScaleSetNetworkConfigurationArrayOutput)
}

type VirtualMachineScaleSetNetworkConfigurationOutput struct{ *pulumi.OutputState }

func (VirtualMachineScaleSetNetworkConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualMachineScaleSetNetworkConfiguration)(nil)).Elem()
}

func (o VirtualMachineScaleSetNetworkConfigurationOutput) ToVirtualMachineScaleSetNetworkConfigurationOutput() VirtualMachineScaleSetNetworkConfigurationOutput {
	return o
}

func (o VirtualMachineScaleSetNetworkConfigurationOutput) ToVirtualMachineScaleSetNetworkConfigurationOutputWithContext(ctx context.Context) VirtualMachineScaleSetNetworkConfigurationOutput {
	return o
}

func (o VirtualMachineScaleSetNetworkConfigurationOutput) DnsSettings() VirtualMachineScaleSetNetworkConfigurationDnsSettingsPtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetNetworkConfiguration) *VirtualMachineScaleSetNetworkConfigurationDnsSettings {
		return v.DnsSettings
	}).(VirtualMachineScaleSetNetworkConfigurationDnsSettingsPtrOutput)
}

func (o VirtualMachineScaleSetNetworkConfigurationOutput) EnableAcceleratedNetworking() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetNetworkConfiguration) *bool { return v.EnableAcceleratedNetworking }).(pulumi.BoolPtrOutput)
}

func (o VirtualMachineScaleSetNetworkConfigurationOutput) EnableIPForwarding() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetNetworkConfiguration) *bool { return v.EnableIPForwarding }).(pulumi.BoolPtrOutput)
}

func (o VirtualMachineScaleSetNetworkConfigurationOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetNetworkConfiguration) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o VirtualMachineScaleSetNetworkConfigurationOutput) IpConfigurations() VirtualMachineScaleSetIPConfigurationArrayOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetNetworkConfiguration) []VirtualMachineScaleSetIPConfiguration {
		return v.IpConfigurations
	}).(VirtualMachineScaleSetIPConfigurationArrayOutput)
}

func (o VirtualMachineScaleSetNetworkConfigurationOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetNetworkConfiguration) string { return v.Name }).(pulumi.StringOutput)
}

func (o VirtualMachineScaleSetNetworkConfigurationOutput) NetworkSecurityGroup() SubResourcePtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetNetworkConfiguration) *SubResource { return v.NetworkSecurityGroup }).(SubResourcePtrOutput)
}

func (o VirtualMachineScaleSetNetworkConfigurationOutput) Primary() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetNetworkConfiguration) *bool { return v.Primary }).(pulumi.BoolPtrOutput)
}

type VirtualMachineScaleSetNetworkConfigurationArrayOutput struct{ *pulumi.OutputState }

func (VirtualMachineScaleSetNetworkConfigurationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VirtualMachineScaleSetNetworkConfiguration)(nil)).Elem()
}

func (o VirtualMachineScaleSetNetworkConfigurationArrayOutput) ToVirtualMachineScaleSetNetworkConfigurationArrayOutput() VirtualMachineScaleSetNetworkConfigurationArrayOutput {
	return o
}

func (o VirtualMachineScaleSetNetworkConfigurationArrayOutput) ToVirtualMachineScaleSetNetworkConfigurationArrayOutputWithContext(ctx context.Context) VirtualMachineScaleSetNetworkConfigurationArrayOutput {
	return o
}

func (o VirtualMachineScaleSetNetworkConfigurationArrayOutput) Index(i pulumi.IntInput) VirtualMachineScaleSetNetworkConfigurationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VirtualMachineScaleSetNetworkConfiguration {
		return vs[0].([]VirtualMachineScaleSetNetworkConfiguration)[vs[1].(int)]
	}).(VirtualMachineScaleSetNetworkConfigurationOutput)
}

type VirtualMachineScaleSetNetworkConfigurationDnsSettings struct {
	DnsServers []string `pulumi:"dnsServers"`
}





type VirtualMachineScaleSetNetworkConfigurationDnsSettingsInput interface {
	pulumi.Input

	ToVirtualMachineScaleSetNetworkConfigurationDnsSettingsOutput() VirtualMachineScaleSetNetworkConfigurationDnsSettingsOutput
	ToVirtualMachineScaleSetNetworkConfigurationDnsSettingsOutputWithContext(context.Context) VirtualMachineScaleSetNetworkConfigurationDnsSettingsOutput
}

type VirtualMachineScaleSetNetworkConfigurationDnsSettingsArgs struct {
	DnsServers pulumi.StringArrayInput `pulumi:"dnsServers"`
}

func (VirtualMachineScaleSetNetworkConfigurationDnsSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualMachineScaleSetNetworkConfigurationDnsSettings)(nil)).Elem()
}

func (i VirtualMachineScaleSetNetworkConfigurationDnsSettingsArgs) ToVirtualMachineScaleSetNetworkConfigurationDnsSettingsOutput() VirtualMachineScaleSetNetworkConfigurationDnsSettingsOutput {
	return i.ToVirtualMachineScaleSetNetworkConfigurationDnsSettingsOutputWithContext(context.Background())
}

func (i VirtualMachineScaleSetNetworkConfigurationDnsSettingsArgs) ToVirtualMachineScaleSetNetworkConfigurationDnsSettingsOutputWithContext(ctx context.Context) VirtualMachineScaleSetNetworkConfigurationDnsSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualMachineScaleSetNetworkConfigurationDnsSettingsOutput)
}

func (i VirtualMachineScaleSetNetworkConfigurationDnsSettingsArgs) ToVirtualMachineScaleSetNetworkConfigurationDnsSettingsPtrOutput() VirtualMachineScaleSetNetworkConfigurationDnsSettingsPtrOutput {
	return i.ToVirtualMachineScaleSetNetworkConfigurationDnsSettingsPtrOutputWithContext(context.Background())
}

func (i VirtualMachineScaleSetNetworkConfigurationDnsSettingsArgs) ToVirtualMachineScaleSetNetworkConfigurationDnsSettingsPtrOutputWithContext(ctx context.Context) VirtualMachineScaleSetNetworkConfigurationDnsSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualMachineScaleSetNetworkConfigurationDnsSettingsOutput).ToVirtualMachineScaleSetNetworkConfigurationDnsSettingsPtrOutputWithContext(ctx)
}









type VirtualMachineScaleSetNetworkConfigurationDnsSettingsPtrInput interface {
	pulumi.Input

	ToVirtualMachineScaleSetNetworkConfigurationDnsSettingsPtrOutput() VirtualMachineScaleSetNetworkConfigurationDnsSettingsPtrOutput
	ToVirtualMachineScaleSetNetworkConfigurationDnsSettingsPtrOutputWithContext(context.Context) VirtualMachineScaleSetNetworkConfigurationDnsSettingsPtrOutput
}

type virtualMachineScaleSetNetworkConfigurationDnsSettingsPtrType VirtualMachineScaleSetNetworkConfigurationDnsSettingsArgs

func VirtualMachineScaleSetNetworkConfigurationDnsSettingsPtr(v *VirtualMachineScaleSetNetworkConfigurationDnsSettingsArgs) VirtualMachineScaleSetNetworkConfigurationDnsSettingsPtrInput {
	return (*virtualMachineScaleSetNetworkConfigurationDnsSettingsPtrType)(v)
}

func (*virtualMachineScaleSetNetworkConfigurationDnsSettingsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualMachineScaleSetNetworkConfigurationDnsSettings)(nil)).Elem()
}

func (i *virtualMachineScaleSetNetworkConfigurationDnsSettingsPtrType) ToVirtualMachineScaleSetNetworkConfigurationDnsSettingsPtrOutput() VirtualMachineScaleSetNetworkConfigurationDnsSettingsPtrOutput {
	return i.ToVirtualMachineScaleSetNetworkConfigurationDnsSettingsPtrOutputWithContext(context.Background())
}

func (i *virtualMachineScaleSetNetworkConfigurationDnsSettingsPtrType) ToVirtualMachineScaleSetNetworkConfigurationDnsSettingsPtrOutputWithContext(ctx context.Context) VirtualMachineScaleSetNetworkConfigurationDnsSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualMachineScaleSetNetworkConfigurationDnsSettingsPtrOutput)
}

type VirtualMachineScaleSetNetworkConfigurationDnsSettingsOutput struct{ *pulumi.OutputState }

func (VirtualMachineScaleSetNetworkConfigurationDnsSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualMachineScaleSetNetworkConfigurationDnsSettings)(nil)).Elem()
}

func (o VirtualMachineScaleSetNetworkConfigurationDnsSettingsOutput) ToVirtualMachineScaleSetNetworkConfigurationDnsSettingsOutput() VirtualMachineScaleSetNetworkConfigurationDnsSettingsOutput {
	return o
}

func (o VirtualMachineScaleSetNetworkConfigurationDnsSettingsOutput) ToVirtualMachineScaleSetNetworkConfigurationDnsSettingsOutputWithContext(ctx context.Context) VirtualMachineScaleSetNetworkConfigurationDnsSettingsOutput {
	return o
}

func (o VirtualMachineScaleSetNetworkConfigurationDnsSettingsOutput) ToVirtualMachineScaleSetNetworkConfigurationDnsSettingsPtrOutput() VirtualMachineScaleSetNetworkConfigurationDnsSettingsPtrOutput {
	return o.ToVirtualMachineScaleSetNetworkConfigurationDnsSettingsPtrOutputWithContext(context.Background())
}

func (o VirtualMachineScaleSetNetworkConfigurationDnsSettingsOutput) ToVirtualMachineScaleSetNetworkConfigurationDnsSettingsPtrOutputWithContext(ctx context.Context) VirtualMachineScaleSetNetworkConfigurationDnsSettingsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v VirtualMachineScaleSetNetworkConfigurationDnsSettings) *VirtualMachineScaleSetNetworkConfigurationDnsSettings {
		return &v
	}).(VirtualMachineScaleSetNetworkConfigurationDnsSettingsPtrOutput)
}

func (o VirtualMachineScaleSetNetworkConfigurationDnsSettingsOutput) DnsServers() pulumi.StringArrayOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetNetworkConfigurationDnsSettings) []string { return v.DnsServers }).(pulumi.StringArrayOutput)
}

type VirtualMachineScaleSetNetworkConfigurationDnsSettingsPtrOutput struct{ *pulumi.OutputState }

func (VirtualMachineScaleSetNetworkConfigurationDnsSettingsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualMachineScaleSetNetworkConfigurationDnsSettings)(nil)).Elem()
}

func (o VirtualMachineScaleSetNetworkConfigurationDnsSettingsPtrOutput) ToVirtualMachineScaleSetNetworkConfigurationDnsSettingsPtrOutput() VirtualMachineScaleSetNetworkConfigurationDnsSettingsPtrOutput {
	return o
}

func (o VirtualMachineScaleSetNetworkConfigurationDnsSettingsPtrOutput) ToVirtualMachineScaleSetNetworkConfigurationDnsSettingsPtrOutputWithContext(ctx context.Context) VirtualMachineScaleSetNetworkConfigurationDnsSettingsPtrOutput {
	return o
}

func (o VirtualMachineScaleSetNetworkConfigurationDnsSettingsPtrOutput) Elem() VirtualMachineScaleSetNetworkConfigurationDnsSettingsOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetNetworkConfigurationDnsSettings) VirtualMachineScaleSetNetworkConfigurationDnsSettings {
		if v != nil {
			return *v
		}
		var ret VirtualMachineScaleSetNetworkConfigurationDnsSettings
		return ret
	}).(VirtualMachineScaleSetNetworkConfigurationDnsSettingsOutput)
}

func (o VirtualMachineScaleSetNetworkConfigurationDnsSettingsPtrOutput) DnsServers() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetNetworkConfigurationDnsSettings) []string {
		if v == nil {
			return nil
		}
		return v.DnsServers
	}).(pulumi.StringArrayOutput)
}

type VirtualMachineScaleSetNetworkConfigurationDnsSettingsResponse struct {
	DnsServers []string `pulumi:"dnsServers"`
}

type VirtualMachineScaleSetNetworkConfigurationDnsSettingsResponseOutput struct{ *pulumi.OutputState }

func (VirtualMachineScaleSetNetworkConfigurationDnsSettingsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualMachineScaleSetNetworkConfigurationDnsSettingsResponse)(nil)).Elem()
}

func (o VirtualMachineScaleSetNetworkConfigurationDnsSettingsResponseOutput) ToVirtualMachineScaleSetNetworkConfigurationDnsSettingsResponseOutput() VirtualMachineScaleSetNetworkConfigurationDnsSettingsResponseOutput {
	return o
}

func (o VirtualMachineScaleSetNetworkConfigurationDnsSettingsResponseOutput) ToVirtualMachineScaleSetNetworkConfigurationDnsSettingsResponseOutputWithContext(ctx context.Context) VirtualMachineScaleSetNetworkConfigurationDnsSettingsResponseOutput {
	return o
}

func (o VirtualMachineScaleSetNetworkConfigurationDnsSettingsResponseOutput) DnsServers() pulumi.StringArrayOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetNetworkConfigurationDnsSettingsResponse) []string { return v.DnsServers }).(pulumi.StringArrayOutput)
}

type VirtualMachineScaleSetNetworkConfigurationDnsSettingsResponsePtrOutput struct{ *pulumi.OutputState }

func (VirtualMachineScaleSetNetworkConfigurationDnsSettingsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualMachineScaleSetNetworkConfigurationDnsSettingsResponse)(nil)).Elem()
}

func (o VirtualMachineScaleSetNetworkConfigurationDnsSettingsResponsePtrOutput) ToVirtualMachineScaleSetNetworkConfigurationDnsSettingsResponsePtrOutput() VirtualMachineScaleSetNetworkConfigurationDnsSettingsResponsePtrOutput {
	return o
}

func (o VirtualMachineScaleSetNetworkConfigurationDnsSettingsResponsePtrOutput) ToVirtualMachineScaleSetNetworkConfigurationDnsSettingsResponsePtrOutputWithContext(ctx context.Context) VirtualMachineScaleSetNetworkConfigurationDnsSettingsResponsePtrOutput {
	return o
}

func (o VirtualMachineScaleSetNetworkConfigurationDnsSettingsResponsePtrOutput) Elem() VirtualMachineScaleSetNetworkConfigurationDnsSettingsResponseOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetNetworkConfigurationDnsSettingsResponse) VirtualMachineScaleSetNetworkConfigurationDnsSettingsResponse {
		if v != nil {
			return *v
		}
		var ret VirtualMachineScaleSetNetworkConfigurationDnsSettingsResponse
		return ret
	}).(VirtualMachineScaleSetNetworkConfigurationDnsSettingsResponseOutput)
}

func (o VirtualMachineScaleSetNetworkConfigurationDnsSettingsResponsePtrOutput) DnsServers() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetNetworkConfigurationDnsSettingsResponse) []string {
		if v == nil {
			return nil
		}
		return v.DnsServers
	}).(pulumi.StringArrayOutput)
}

type VirtualMachineScaleSetNetworkConfigurationResponse struct {
	DnsSettings                 *VirtualMachineScaleSetNetworkConfigurationDnsSettingsResponse `pulumi:"dnsSettings"`
	EnableAcceleratedNetworking *bool                                                          `pulumi:"enableAcceleratedNetworking"`
	EnableIPForwarding          *bool                                                          `pulumi:"enableIPForwarding"`
	Id                          *string                                                        `pulumi:"id"`
	IpConfigurations            []VirtualMachineScaleSetIPConfigurationResponse                `pulumi:"ipConfigurations"`
	Name                        string                                                         `pulumi:"name"`
	NetworkSecurityGroup        *SubResourceResponse                                           `pulumi:"networkSecurityGroup"`
	Primary                     *bool                                                          `pulumi:"primary"`
}

type VirtualMachineScaleSetNetworkConfigurationResponseOutput struct{ *pulumi.OutputState }

func (VirtualMachineScaleSetNetworkConfigurationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualMachineScaleSetNetworkConfigurationResponse)(nil)).Elem()
}

func (o VirtualMachineScaleSetNetworkConfigurationResponseOutput) ToVirtualMachineScaleSetNetworkConfigurationResponseOutput() VirtualMachineScaleSetNetworkConfigurationResponseOutput {
	return o
}

func (o VirtualMachineScaleSetNetworkConfigurationResponseOutput) ToVirtualMachineScaleSetNetworkConfigurationResponseOutputWithContext(ctx context.Context) VirtualMachineScaleSetNetworkConfigurationResponseOutput {
	return o
}

func (o VirtualMachineScaleSetNetworkConfigurationResponseOutput) DnsSettings() VirtualMachineScaleSetNetworkConfigurationDnsSettingsResponsePtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetNetworkConfigurationResponse) *VirtualMachineScaleSetNetworkConfigurationDnsSettingsResponse {
		return v.DnsSettings
	}).(VirtualMachineScaleSetNetworkConfigurationDnsSettingsResponsePtrOutput)
}

func (o VirtualMachineScaleSetNetworkConfigurationResponseOutput) EnableAcceleratedNetworking() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetNetworkConfigurationResponse) *bool { return v.EnableAcceleratedNetworking }).(pulumi.BoolPtrOutput)
}

func (o VirtualMachineScaleSetNetworkConfigurationResponseOutput) EnableIPForwarding() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetNetworkConfigurationResponse) *bool { return v.EnableIPForwarding }).(pulumi.BoolPtrOutput)
}

func (o VirtualMachineScaleSetNetworkConfigurationResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetNetworkConfigurationResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o VirtualMachineScaleSetNetworkConfigurationResponseOutput) IpConfigurations() VirtualMachineScaleSetIPConfigurationResponseArrayOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetNetworkConfigurationResponse) []VirtualMachineScaleSetIPConfigurationResponse {
		return v.IpConfigurations
	}).(VirtualMachineScaleSetIPConfigurationResponseArrayOutput)
}

func (o VirtualMachineScaleSetNetworkConfigurationResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetNetworkConfigurationResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o VirtualMachineScaleSetNetworkConfigurationResponseOutput) NetworkSecurityGroup() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetNetworkConfigurationResponse) *SubResourceResponse {
		return v.NetworkSecurityGroup
	}).(SubResourceResponsePtrOutput)
}

func (o VirtualMachineScaleSetNetworkConfigurationResponseOutput) Primary() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetNetworkConfigurationResponse) *bool { return v.Primary }).(pulumi.BoolPtrOutput)
}

type VirtualMachineScaleSetNetworkConfigurationResponseArrayOutput struct{ *pulumi.OutputState }

func (VirtualMachineScaleSetNetworkConfigurationResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VirtualMachineScaleSetNetworkConfigurationResponse)(nil)).Elem()
}

func (o VirtualMachineScaleSetNetworkConfigurationResponseArrayOutput) ToVirtualMachineScaleSetNetworkConfigurationResponseArrayOutput() VirtualMachineScaleSetNetworkConfigurationResponseArrayOutput {
	return o
}

func (o VirtualMachineScaleSetNetworkConfigurationResponseArrayOutput) ToVirtualMachineScaleSetNetworkConfigurationResponseArrayOutputWithContext(ctx context.Context) VirtualMachineScaleSetNetworkConfigurationResponseArrayOutput {
	return o
}

func (o VirtualMachineScaleSetNetworkConfigurationResponseArrayOutput) Index(i pulumi.IntInput) VirtualMachineScaleSetNetworkConfigurationResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VirtualMachineScaleSetNetworkConfigurationResponse {
		return vs[0].([]VirtualMachineScaleSetNetworkConfigurationResponse)[vs[1].(int)]
	}).(VirtualMachineScaleSetNetworkConfigurationResponseOutput)
}

type VirtualMachineScaleSetNetworkProfile struct {
	HealthProbe                    *ApiEntityReference                          `pulumi:"healthProbe"`
	NetworkInterfaceConfigurations []VirtualMachineScaleSetNetworkConfiguration `pulumi:"networkInterfaceConfigurations"`
}





type VirtualMachineScaleSetNetworkProfileInput interface {
	pulumi.Input

	ToVirtualMachineScaleSetNetworkProfileOutput() VirtualMachineScaleSetNetworkProfileOutput
	ToVirtualMachineScaleSetNetworkProfileOutputWithContext(context.Context) VirtualMachineScaleSetNetworkProfileOutput
}

type VirtualMachineScaleSetNetworkProfileArgs struct {
	HealthProbe                    ApiEntityReferencePtrInput                           `pulumi:"healthProbe"`
	NetworkInterfaceConfigurations VirtualMachineScaleSetNetworkConfigurationArrayInput `pulumi:"networkInterfaceConfigurations"`
}

func (VirtualMachineScaleSetNetworkProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualMachineScaleSetNetworkProfile)(nil)).Elem()
}

func (i VirtualMachineScaleSetNetworkProfileArgs) ToVirtualMachineScaleSetNetworkProfileOutput() VirtualMachineScaleSetNetworkProfileOutput {
	return i.ToVirtualMachineScaleSetNetworkProfileOutputWithContext(context.Background())
}

func (i VirtualMachineScaleSetNetworkProfileArgs) ToVirtualMachineScaleSetNetworkProfileOutputWithContext(ctx context.Context) VirtualMachineScaleSetNetworkProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualMachineScaleSetNetworkProfileOutput)
}

func (i VirtualMachineScaleSetNetworkProfileArgs) ToVirtualMachineScaleSetNetworkProfilePtrOutput() VirtualMachineScaleSetNetworkProfilePtrOutput {
	return i.ToVirtualMachineScaleSetNetworkProfilePtrOutputWithContext(context.Background())
}

func (i VirtualMachineScaleSetNetworkProfileArgs) ToVirtualMachineScaleSetNetworkProfilePtrOutputWithContext(ctx context.Context) VirtualMachineScaleSetNetworkProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualMachineScaleSetNetworkProfileOutput).ToVirtualMachineScaleSetNetworkProfilePtrOutputWithContext(ctx)
}









type VirtualMachineScaleSetNetworkProfilePtrInput interface {
	pulumi.Input

	ToVirtualMachineScaleSetNetworkProfilePtrOutput() VirtualMachineScaleSetNetworkProfilePtrOutput
	ToVirtualMachineScaleSetNetworkProfilePtrOutputWithContext(context.Context) VirtualMachineScaleSetNetworkProfilePtrOutput
}

type virtualMachineScaleSetNetworkProfilePtrType VirtualMachineScaleSetNetworkProfileArgs

func VirtualMachineScaleSetNetworkProfilePtr(v *VirtualMachineScaleSetNetworkProfileArgs) VirtualMachineScaleSetNetworkProfilePtrInput {
	return (*virtualMachineScaleSetNetworkProfilePtrType)(v)
}

func (*virtualMachineScaleSetNetworkProfilePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualMachineScaleSetNetworkProfile)(nil)).Elem()
}

func (i *virtualMachineScaleSetNetworkProfilePtrType) ToVirtualMachineScaleSetNetworkProfilePtrOutput() VirtualMachineScaleSetNetworkProfilePtrOutput {
	return i.ToVirtualMachineScaleSetNetworkProfilePtrOutputWithContext(context.Background())
}

func (i *virtualMachineScaleSetNetworkProfilePtrType) ToVirtualMachineScaleSetNetworkProfilePtrOutputWithContext(ctx context.Context) VirtualMachineScaleSetNetworkProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualMachineScaleSetNetworkProfilePtrOutput)
}

type VirtualMachineScaleSetNetworkProfileOutput struct{ *pulumi.OutputState }

func (VirtualMachineScaleSetNetworkProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualMachineScaleSetNetworkProfile)(nil)).Elem()
}

func (o VirtualMachineScaleSetNetworkProfileOutput) ToVirtualMachineScaleSetNetworkProfileOutput() VirtualMachineScaleSetNetworkProfileOutput {
	return o
}

func (o VirtualMachineScaleSetNetworkProfileOutput) ToVirtualMachineScaleSetNetworkProfileOutputWithContext(ctx context.Context) VirtualMachineScaleSetNetworkProfileOutput {
	return o
}

func (o VirtualMachineScaleSetNetworkProfileOutput) ToVirtualMachineScaleSetNetworkProfilePtrOutput() VirtualMachineScaleSetNetworkProfilePtrOutput {
	return o.ToVirtualMachineScaleSetNetworkProfilePtrOutputWithContext(context.Background())
}

func (o VirtualMachineScaleSetNetworkProfileOutput) ToVirtualMachineScaleSetNetworkProfilePtrOutputWithContext(ctx context.Context) VirtualMachineScaleSetNetworkProfilePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v VirtualMachineScaleSetNetworkProfile) *VirtualMachineScaleSetNetworkProfile {
		return &v
	}).(VirtualMachineScaleSetNetworkProfilePtrOutput)
}

func (o VirtualMachineScaleSetNetworkProfileOutput) HealthProbe() ApiEntityReferencePtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetNetworkProfile) *ApiEntityReference { return v.HealthProbe }).(ApiEntityReferencePtrOutput)
}

func (o VirtualMachineScaleSetNetworkProfileOutput) NetworkInterfaceConfigurations() VirtualMachineScaleSetNetworkConfigurationArrayOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetNetworkProfile) []VirtualMachineScaleSetNetworkConfiguration {
		return v.NetworkInterfaceConfigurations
	}).(VirtualMachineScaleSetNetworkConfigurationArrayOutput)
}

type VirtualMachineScaleSetNetworkProfilePtrOutput struct{ *pulumi.OutputState }

func (VirtualMachineScaleSetNetworkProfilePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualMachineScaleSetNetworkProfile)(nil)).Elem()
}

func (o VirtualMachineScaleSetNetworkProfilePtrOutput) ToVirtualMachineScaleSetNetworkProfilePtrOutput() VirtualMachineScaleSetNetworkProfilePtrOutput {
	return o
}

func (o VirtualMachineScaleSetNetworkProfilePtrOutput) ToVirtualMachineScaleSetNetworkProfilePtrOutputWithContext(ctx context.Context) VirtualMachineScaleSetNetworkProfilePtrOutput {
	return o
}

func (o VirtualMachineScaleSetNetworkProfilePtrOutput) Elem() VirtualMachineScaleSetNetworkProfileOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetNetworkProfile) VirtualMachineScaleSetNetworkProfile {
		if v != nil {
			return *v
		}
		var ret VirtualMachineScaleSetNetworkProfile
		return ret
	}).(VirtualMachineScaleSetNetworkProfileOutput)
}

func (o VirtualMachineScaleSetNetworkProfilePtrOutput) HealthProbe() ApiEntityReferencePtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetNetworkProfile) *ApiEntityReference {
		if v == nil {
			return nil
		}
		return v.HealthProbe
	}).(ApiEntityReferencePtrOutput)
}

func (o VirtualMachineScaleSetNetworkProfilePtrOutput) NetworkInterfaceConfigurations() VirtualMachineScaleSetNetworkConfigurationArrayOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetNetworkProfile) []VirtualMachineScaleSetNetworkConfiguration {
		if v == nil {
			return nil
		}
		return v.NetworkInterfaceConfigurations
	}).(VirtualMachineScaleSetNetworkConfigurationArrayOutput)
}

type VirtualMachineScaleSetNetworkProfileResponse struct {
	HealthProbe                    *ApiEntityReferenceResponse                          `pulumi:"healthProbe"`
	NetworkInterfaceConfigurations []VirtualMachineScaleSetNetworkConfigurationResponse `pulumi:"networkInterfaceConfigurations"`
}

type VirtualMachineScaleSetNetworkProfileResponseOutput struct{ *pulumi.OutputState }

func (VirtualMachineScaleSetNetworkProfileResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualMachineScaleSetNetworkProfileResponse)(nil)).Elem()
}

func (o VirtualMachineScaleSetNetworkProfileResponseOutput) ToVirtualMachineScaleSetNetworkProfileResponseOutput() VirtualMachineScaleSetNetworkProfileResponseOutput {
	return o
}

func (o VirtualMachineScaleSetNetworkProfileResponseOutput) ToVirtualMachineScaleSetNetworkProfileResponseOutputWithContext(ctx context.Context) VirtualMachineScaleSetNetworkProfileResponseOutput {
	return o
}

func (o VirtualMachineScaleSetNetworkProfileResponseOutput) HealthProbe() ApiEntityReferenceResponsePtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetNetworkProfileResponse) *ApiEntityReferenceResponse { return v.HealthProbe }).(ApiEntityReferenceResponsePtrOutput)
}

func (o VirtualMachineScaleSetNetworkProfileResponseOutput) NetworkInterfaceConfigurations() VirtualMachineScaleSetNetworkConfigurationResponseArrayOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetNetworkProfileResponse) []VirtualMachineScaleSetNetworkConfigurationResponse {
		return v.NetworkInterfaceConfigurations
	}).(VirtualMachineScaleSetNetworkConfigurationResponseArrayOutput)
}

type VirtualMachineScaleSetNetworkProfileResponsePtrOutput struct{ *pulumi.OutputState }

func (VirtualMachineScaleSetNetworkProfileResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualMachineScaleSetNetworkProfileResponse)(nil)).Elem()
}

func (o VirtualMachineScaleSetNetworkProfileResponsePtrOutput) ToVirtualMachineScaleSetNetworkProfileResponsePtrOutput() VirtualMachineScaleSetNetworkProfileResponsePtrOutput {
	return o
}

func (o VirtualMachineScaleSetNetworkProfileResponsePtrOutput) ToVirtualMachineScaleSetNetworkProfileResponsePtrOutputWithContext(ctx context.Context) VirtualMachineScaleSetNetworkProfileResponsePtrOutput {
	return o
}

func (o VirtualMachineScaleSetNetworkProfileResponsePtrOutput) Elem() VirtualMachineScaleSetNetworkProfileResponseOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetNetworkProfileResponse) VirtualMachineScaleSetNetworkProfileResponse {
		if v != nil {
			return *v
		}
		var ret VirtualMachineScaleSetNetworkProfileResponse
		return ret
	}).(VirtualMachineScaleSetNetworkProfileResponseOutput)
}

func (o VirtualMachineScaleSetNetworkProfileResponsePtrOutput) HealthProbe() ApiEntityReferenceResponsePtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetNetworkProfileResponse) *ApiEntityReferenceResponse {
		if v == nil {
			return nil
		}
		return v.HealthProbe
	}).(ApiEntityReferenceResponsePtrOutput)
}

func (o VirtualMachineScaleSetNetworkProfileResponsePtrOutput) NetworkInterfaceConfigurations() VirtualMachineScaleSetNetworkConfigurationResponseArrayOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetNetworkProfileResponse) []VirtualMachineScaleSetNetworkConfigurationResponse {
		if v == nil {
			return nil
		}
		return v.NetworkInterfaceConfigurations
	}).(VirtualMachineScaleSetNetworkConfigurationResponseArrayOutput)
}

type VirtualMachineScaleSetOSDisk struct {
	Caching                 *CachingTypes                                `pulumi:"caching"`
	CreateOption            string                                       `pulumi:"createOption"`
	DiskSizeGB              *int                                         `pulumi:"diskSizeGB"`
	Image                   *VirtualHardDisk                             `pulumi:"image"`
	ManagedDisk             *VirtualMachineScaleSetManagedDiskParameters `pulumi:"managedDisk"`
	Name                    *string                                      `pulumi:"name"`
	OsType                  *OperatingSystemTypes                        `pulumi:"osType"`
	VhdContainers           []string                                     `pulumi:"vhdContainers"`
	WriteAcceleratorEnabled *bool                                        `pulumi:"writeAcceleratorEnabled"`
}





type VirtualMachineScaleSetOSDiskInput interface {
	pulumi.Input

	ToVirtualMachineScaleSetOSDiskOutput() VirtualMachineScaleSetOSDiskOutput
	ToVirtualMachineScaleSetOSDiskOutputWithContext(context.Context) VirtualMachineScaleSetOSDiskOutput
}

type VirtualMachineScaleSetOSDiskArgs struct {
	Caching                 CachingTypesPtrInput                                `pulumi:"caching"`
	CreateOption            pulumi.StringInput                                  `pulumi:"createOption"`
	DiskSizeGB              pulumi.IntPtrInput                                  `pulumi:"diskSizeGB"`
	Image                   VirtualHardDiskPtrInput                             `pulumi:"image"`
	ManagedDisk             VirtualMachineScaleSetManagedDiskParametersPtrInput `pulumi:"managedDisk"`
	Name                    pulumi.StringPtrInput                               `pulumi:"name"`
	OsType                  OperatingSystemTypesPtrInput                        `pulumi:"osType"`
	VhdContainers           pulumi.StringArrayInput                             `pulumi:"vhdContainers"`
	WriteAcceleratorEnabled pulumi.BoolPtrInput                                 `pulumi:"writeAcceleratorEnabled"`
}

func (VirtualMachineScaleSetOSDiskArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualMachineScaleSetOSDisk)(nil)).Elem()
}

func (i VirtualMachineScaleSetOSDiskArgs) ToVirtualMachineScaleSetOSDiskOutput() VirtualMachineScaleSetOSDiskOutput {
	return i.ToVirtualMachineScaleSetOSDiskOutputWithContext(context.Background())
}

func (i VirtualMachineScaleSetOSDiskArgs) ToVirtualMachineScaleSetOSDiskOutputWithContext(ctx context.Context) VirtualMachineScaleSetOSDiskOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualMachineScaleSetOSDiskOutput)
}

func (i VirtualMachineScaleSetOSDiskArgs) ToVirtualMachineScaleSetOSDiskPtrOutput() VirtualMachineScaleSetOSDiskPtrOutput {
	return i.ToVirtualMachineScaleSetOSDiskPtrOutputWithContext(context.Background())
}

func (i VirtualMachineScaleSetOSDiskArgs) ToVirtualMachineScaleSetOSDiskPtrOutputWithContext(ctx context.Context) VirtualMachineScaleSetOSDiskPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualMachineScaleSetOSDiskOutput).ToVirtualMachineScaleSetOSDiskPtrOutputWithContext(ctx)
}









type VirtualMachineScaleSetOSDiskPtrInput interface {
	pulumi.Input

	ToVirtualMachineScaleSetOSDiskPtrOutput() VirtualMachineScaleSetOSDiskPtrOutput
	ToVirtualMachineScaleSetOSDiskPtrOutputWithContext(context.Context) VirtualMachineScaleSetOSDiskPtrOutput
}

type virtualMachineScaleSetOSDiskPtrType VirtualMachineScaleSetOSDiskArgs

func VirtualMachineScaleSetOSDiskPtr(v *VirtualMachineScaleSetOSDiskArgs) VirtualMachineScaleSetOSDiskPtrInput {
	return (*virtualMachineScaleSetOSDiskPtrType)(v)
}

func (*virtualMachineScaleSetOSDiskPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualMachineScaleSetOSDisk)(nil)).Elem()
}

func (i *virtualMachineScaleSetOSDiskPtrType) ToVirtualMachineScaleSetOSDiskPtrOutput() VirtualMachineScaleSetOSDiskPtrOutput {
	return i.ToVirtualMachineScaleSetOSDiskPtrOutputWithContext(context.Background())
}

func (i *virtualMachineScaleSetOSDiskPtrType) ToVirtualMachineScaleSetOSDiskPtrOutputWithContext(ctx context.Context) VirtualMachineScaleSetOSDiskPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualMachineScaleSetOSDiskPtrOutput)
}

type VirtualMachineScaleSetOSDiskOutput struct{ *pulumi.OutputState }

func (VirtualMachineScaleSetOSDiskOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualMachineScaleSetOSDisk)(nil)).Elem()
}

func (o VirtualMachineScaleSetOSDiskOutput) ToVirtualMachineScaleSetOSDiskOutput() VirtualMachineScaleSetOSDiskOutput {
	return o
}

func (o VirtualMachineScaleSetOSDiskOutput) ToVirtualMachineScaleSetOSDiskOutputWithContext(ctx context.Context) VirtualMachineScaleSetOSDiskOutput {
	return o
}

func (o VirtualMachineScaleSetOSDiskOutput) ToVirtualMachineScaleSetOSDiskPtrOutput() VirtualMachineScaleSetOSDiskPtrOutput {
	return o.ToVirtualMachineScaleSetOSDiskPtrOutputWithContext(context.Background())
}

func (o VirtualMachineScaleSetOSDiskOutput) ToVirtualMachineScaleSetOSDiskPtrOutputWithContext(ctx context.Context) VirtualMachineScaleSetOSDiskPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v VirtualMachineScaleSetOSDisk) *VirtualMachineScaleSetOSDisk {
		return &v
	}).(VirtualMachineScaleSetOSDiskPtrOutput)
}

func (o VirtualMachineScaleSetOSDiskOutput) Caching() CachingTypesPtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetOSDisk) *CachingTypes { return v.Caching }).(CachingTypesPtrOutput)
}

func (o VirtualMachineScaleSetOSDiskOutput) CreateOption() pulumi.StringOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetOSDisk) string { return v.CreateOption }).(pulumi.StringOutput)
}

func (o VirtualMachineScaleSetOSDiskOutput) DiskSizeGB() pulumi.IntPtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetOSDisk) *int { return v.DiskSizeGB }).(pulumi.IntPtrOutput)
}

func (o VirtualMachineScaleSetOSDiskOutput) Image() VirtualHardDiskPtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetOSDisk) *VirtualHardDisk { return v.Image }).(VirtualHardDiskPtrOutput)
}

func (o VirtualMachineScaleSetOSDiskOutput) ManagedDisk() VirtualMachineScaleSetManagedDiskParametersPtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetOSDisk) *VirtualMachineScaleSetManagedDiskParameters {
		return v.ManagedDisk
	}).(VirtualMachineScaleSetManagedDiskParametersPtrOutput)
}

func (o VirtualMachineScaleSetOSDiskOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetOSDisk) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o VirtualMachineScaleSetOSDiskOutput) OsType() OperatingSystemTypesPtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetOSDisk) *OperatingSystemTypes { return v.OsType }).(OperatingSystemTypesPtrOutput)
}

func (o VirtualMachineScaleSetOSDiskOutput) VhdContainers() pulumi.StringArrayOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetOSDisk) []string { return v.VhdContainers }).(pulumi.StringArrayOutput)
}

func (o VirtualMachineScaleSetOSDiskOutput) WriteAcceleratorEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetOSDisk) *bool { return v.WriteAcceleratorEnabled }).(pulumi.BoolPtrOutput)
}

type VirtualMachineScaleSetOSDiskPtrOutput struct{ *pulumi.OutputState }

func (VirtualMachineScaleSetOSDiskPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualMachineScaleSetOSDisk)(nil)).Elem()
}

func (o VirtualMachineScaleSetOSDiskPtrOutput) ToVirtualMachineScaleSetOSDiskPtrOutput() VirtualMachineScaleSetOSDiskPtrOutput {
	return o
}

func (o VirtualMachineScaleSetOSDiskPtrOutput) ToVirtualMachineScaleSetOSDiskPtrOutputWithContext(ctx context.Context) VirtualMachineScaleSetOSDiskPtrOutput {
	return o
}

func (o VirtualMachineScaleSetOSDiskPtrOutput) Elem() VirtualMachineScaleSetOSDiskOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetOSDisk) VirtualMachineScaleSetOSDisk {
		if v != nil {
			return *v
		}
		var ret VirtualMachineScaleSetOSDisk
		return ret
	}).(VirtualMachineScaleSetOSDiskOutput)
}

func (o VirtualMachineScaleSetOSDiskPtrOutput) Caching() CachingTypesPtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetOSDisk) *CachingTypes {
		if v == nil {
			return nil
		}
		return v.Caching
	}).(CachingTypesPtrOutput)
}

func (o VirtualMachineScaleSetOSDiskPtrOutput) CreateOption() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetOSDisk) *string {
		if v == nil {
			return nil
		}
		return &v.CreateOption
	}).(pulumi.StringPtrOutput)
}

func (o VirtualMachineScaleSetOSDiskPtrOutput) DiskSizeGB() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetOSDisk) *int {
		if v == nil {
			return nil
		}
		return v.DiskSizeGB
	}).(pulumi.IntPtrOutput)
}

func (o VirtualMachineScaleSetOSDiskPtrOutput) Image() VirtualHardDiskPtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetOSDisk) *VirtualHardDisk {
		if v == nil {
			return nil
		}
		return v.Image
	}).(VirtualHardDiskPtrOutput)
}

func (o VirtualMachineScaleSetOSDiskPtrOutput) ManagedDisk() VirtualMachineScaleSetManagedDiskParametersPtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetOSDisk) *VirtualMachineScaleSetManagedDiskParameters {
		if v == nil {
			return nil
		}
		return v.ManagedDisk
	}).(VirtualMachineScaleSetManagedDiskParametersPtrOutput)
}

func (o VirtualMachineScaleSetOSDiskPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetOSDisk) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o VirtualMachineScaleSetOSDiskPtrOutput) OsType() OperatingSystemTypesPtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetOSDisk) *OperatingSystemTypes {
		if v == nil {
			return nil
		}
		return v.OsType
	}).(OperatingSystemTypesPtrOutput)
}

func (o VirtualMachineScaleSetOSDiskPtrOutput) VhdContainers() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetOSDisk) []string {
		if v == nil {
			return nil
		}
		return v.VhdContainers
	}).(pulumi.StringArrayOutput)
}

func (o VirtualMachineScaleSetOSDiskPtrOutput) WriteAcceleratorEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetOSDisk) *bool {
		if v == nil {
			return nil
		}
		return v.WriteAcceleratorEnabled
	}).(pulumi.BoolPtrOutput)
}

type VirtualMachineScaleSetOSDiskResponse struct {
	Caching                 *string                                              `pulumi:"caching"`
	CreateOption            string                                               `pulumi:"createOption"`
	DiskSizeGB              *int                                                 `pulumi:"diskSizeGB"`
	Image                   *VirtualHardDiskResponse                             `pulumi:"image"`
	ManagedDisk             *VirtualMachineScaleSetManagedDiskParametersResponse `pulumi:"managedDisk"`
	Name                    *string                                              `pulumi:"name"`
	OsType                  *string                                              `pulumi:"osType"`
	VhdContainers           []string                                             `pulumi:"vhdContainers"`
	WriteAcceleratorEnabled *bool                                                `pulumi:"writeAcceleratorEnabled"`
}

type VirtualMachineScaleSetOSDiskResponseOutput struct{ *pulumi.OutputState }

func (VirtualMachineScaleSetOSDiskResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualMachineScaleSetOSDiskResponse)(nil)).Elem()
}

func (o VirtualMachineScaleSetOSDiskResponseOutput) ToVirtualMachineScaleSetOSDiskResponseOutput() VirtualMachineScaleSetOSDiskResponseOutput {
	return o
}

func (o VirtualMachineScaleSetOSDiskResponseOutput) ToVirtualMachineScaleSetOSDiskResponseOutputWithContext(ctx context.Context) VirtualMachineScaleSetOSDiskResponseOutput {
	return o
}

func (o VirtualMachineScaleSetOSDiskResponseOutput) Caching() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetOSDiskResponse) *string { return v.Caching }).(pulumi.StringPtrOutput)
}

func (o VirtualMachineScaleSetOSDiskResponseOutput) CreateOption() pulumi.StringOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetOSDiskResponse) string { return v.CreateOption }).(pulumi.StringOutput)
}

func (o VirtualMachineScaleSetOSDiskResponseOutput) DiskSizeGB() pulumi.IntPtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetOSDiskResponse) *int { return v.DiskSizeGB }).(pulumi.IntPtrOutput)
}

func (o VirtualMachineScaleSetOSDiskResponseOutput) Image() VirtualHardDiskResponsePtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetOSDiskResponse) *VirtualHardDiskResponse { return v.Image }).(VirtualHardDiskResponsePtrOutput)
}

func (o VirtualMachineScaleSetOSDiskResponseOutput) ManagedDisk() VirtualMachineScaleSetManagedDiskParametersResponsePtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetOSDiskResponse) *VirtualMachineScaleSetManagedDiskParametersResponse {
		return v.ManagedDisk
	}).(VirtualMachineScaleSetManagedDiskParametersResponsePtrOutput)
}

func (o VirtualMachineScaleSetOSDiskResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetOSDiskResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o VirtualMachineScaleSetOSDiskResponseOutput) OsType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetOSDiskResponse) *string { return v.OsType }).(pulumi.StringPtrOutput)
}

func (o VirtualMachineScaleSetOSDiskResponseOutput) VhdContainers() pulumi.StringArrayOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetOSDiskResponse) []string { return v.VhdContainers }).(pulumi.StringArrayOutput)
}

func (o VirtualMachineScaleSetOSDiskResponseOutput) WriteAcceleratorEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetOSDiskResponse) *bool { return v.WriteAcceleratorEnabled }).(pulumi.BoolPtrOutput)
}

type VirtualMachineScaleSetOSDiskResponsePtrOutput struct{ *pulumi.OutputState }

func (VirtualMachineScaleSetOSDiskResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualMachineScaleSetOSDiskResponse)(nil)).Elem()
}

func (o VirtualMachineScaleSetOSDiskResponsePtrOutput) ToVirtualMachineScaleSetOSDiskResponsePtrOutput() VirtualMachineScaleSetOSDiskResponsePtrOutput {
	return o
}

func (o VirtualMachineScaleSetOSDiskResponsePtrOutput) ToVirtualMachineScaleSetOSDiskResponsePtrOutputWithContext(ctx context.Context) VirtualMachineScaleSetOSDiskResponsePtrOutput {
	return o
}

func (o VirtualMachineScaleSetOSDiskResponsePtrOutput) Elem() VirtualMachineScaleSetOSDiskResponseOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetOSDiskResponse) VirtualMachineScaleSetOSDiskResponse {
		if v != nil {
			return *v
		}
		var ret VirtualMachineScaleSetOSDiskResponse
		return ret
	}).(VirtualMachineScaleSetOSDiskResponseOutput)
}

func (o VirtualMachineScaleSetOSDiskResponsePtrOutput) Caching() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetOSDiskResponse) *string {
		if v == nil {
			return nil
		}
		return v.Caching
	}).(pulumi.StringPtrOutput)
}

func (o VirtualMachineScaleSetOSDiskResponsePtrOutput) CreateOption() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetOSDiskResponse) *string {
		if v == nil {
			return nil
		}
		return &v.CreateOption
	}).(pulumi.StringPtrOutput)
}

func (o VirtualMachineScaleSetOSDiskResponsePtrOutput) DiskSizeGB() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetOSDiskResponse) *int {
		if v == nil {
			return nil
		}
		return v.DiskSizeGB
	}).(pulumi.IntPtrOutput)
}

func (o VirtualMachineScaleSetOSDiskResponsePtrOutput) Image() VirtualHardDiskResponsePtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetOSDiskResponse) *VirtualHardDiskResponse {
		if v == nil {
			return nil
		}
		return v.Image
	}).(VirtualHardDiskResponsePtrOutput)
}

func (o VirtualMachineScaleSetOSDiskResponsePtrOutput) ManagedDisk() VirtualMachineScaleSetManagedDiskParametersResponsePtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetOSDiskResponse) *VirtualMachineScaleSetManagedDiskParametersResponse {
		if v == nil {
			return nil
		}
		return v.ManagedDisk
	}).(VirtualMachineScaleSetManagedDiskParametersResponsePtrOutput)
}

func (o VirtualMachineScaleSetOSDiskResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetOSDiskResponse) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o VirtualMachineScaleSetOSDiskResponsePtrOutput) OsType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetOSDiskResponse) *string {
		if v == nil {
			return nil
		}
		return v.OsType
	}).(pulumi.StringPtrOutput)
}

func (o VirtualMachineScaleSetOSDiskResponsePtrOutput) VhdContainers() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetOSDiskResponse) []string {
		if v == nil {
			return nil
		}
		return v.VhdContainers
	}).(pulumi.StringArrayOutput)
}

func (o VirtualMachineScaleSetOSDiskResponsePtrOutput) WriteAcceleratorEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetOSDiskResponse) *bool {
		if v == nil {
			return nil
		}
		return v.WriteAcceleratorEnabled
	}).(pulumi.BoolPtrOutput)
}

type VirtualMachineScaleSetOSProfile struct {
	AdminPassword        *string               `pulumi:"adminPassword"`
	AdminUsername        *string               `pulumi:"adminUsername"`
	ComputerNamePrefix   *string               `pulumi:"computerNamePrefix"`
	CustomData           *string               `pulumi:"customData"`
	LinuxConfiguration   *LinuxConfiguration   `pulumi:"linuxConfiguration"`
	Secrets              []VaultSecretGroup    `pulumi:"secrets"`
	WindowsConfiguration *WindowsConfiguration `pulumi:"windowsConfiguration"`
}





type VirtualMachineScaleSetOSProfileInput interface {
	pulumi.Input

	ToVirtualMachineScaleSetOSProfileOutput() VirtualMachineScaleSetOSProfileOutput
	ToVirtualMachineScaleSetOSProfileOutputWithContext(context.Context) VirtualMachineScaleSetOSProfileOutput
}

type VirtualMachineScaleSetOSProfileArgs struct {
	AdminPassword        pulumi.StringPtrInput        `pulumi:"adminPassword"`
	AdminUsername        pulumi.StringPtrInput        `pulumi:"adminUsername"`
	ComputerNamePrefix   pulumi.StringPtrInput        `pulumi:"computerNamePrefix"`
	CustomData           pulumi.StringPtrInput        `pulumi:"customData"`
	LinuxConfiguration   LinuxConfigurationPtrInput   `pulumi:"linuxConfiguration"`
	Secrets              VaultSecretGroupArrayInput   `pulumi:"secrets"`
	WindowsConfiguration WindowsConfigurationPtrInput `pulumi:"windowsConfiguration"`
}

func (VirtualMachineScaleSetOSProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualMachineScaleSetOSProfile)(nil)).Elem()
}

func (i VirtualMachineScaleSetOSProfileArgs) ToVirtualMachineScaleSetOSProfileOutput() VirtualMachineScaleSetOSProfileOutput {
	return i.ToVirtualMachineScaleSetOSProfileOutputWithContext(context.Background())
}

func (i VirtualMachineScaleSetOSProfileArgs) ToVirtualMachineScaleSetOSProfileOutputWithContext(ctx context.Context) VirtualMachineScaleSetOSProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualMachineScaleSetOSProfileOutput)
}

func (i VirtualMachineScaleSetOSProfileArgs) ToVirtualMachineScaleSetOSProfilePtrOutput() VirtualMachineScaleSetOSProfilePtrOutput {
	return i.ToVirtualMachineScaleSetOSProfilePtrOutputWithContext(context.Background())
}

func (i VirtualMachineScaleSetOSProfileArgs) ToVirtualMachineScaleSetOSProfilePtrOutputWithContext(ctx context.Context) VirtualMachineScaleSetOSProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualMachineScaleSetOSProfileOutput).ToVirtualMachineScaleSetOSProfilePtrOutputWithContext(ctx)
}









type VirtualMachineScaleSetOSProfilePtrInput interface {
	pulumi.Input

	ToVirtualMachineScaleSetOSProfilePtrOutput() VirtualMachineScaleSetOSProfilePtrOutput
	ToVirtualMachineScaleSetOSProfilePtrOutputWithContext(context.Context) VirtualMachineScaleSetOSProfilePtrOutput
}

type virtualMachineScaleSetOSProfilePtrType VirtualMachineScaleSetOSProfileArgs

func VirtualMachineScaleSetOSProfilePtr(v *VirtualMachineScaleSetOSProfileArgs) VirtualMachineScaleSetOSProfilePtrInput {
	return (*virtualMachineScaleSetOSProfilePtrType)(v)
}

func (*virtualMachineScaleSetOSProfilePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualMachineScaleSetOSProfile)(nil)).Elem()
}

func (i *virtualMachineScaleSetOSProfilePtrType) ToVirtualMachineScaleSetOSProfilePtrOutput() VirtualMachineScaleSetOSProfilePtrOutput {
	return i.ToVirtualMachineScaleSetOSProfilePtrOutputWithContext(context.Background())
}

func (i *virtualMachineScaleSetOSProfilePtrType) ToVirtualMachineScaleSetOSProfilePtrOutputWithContext(ctx context.Context) VirtualMachineScaleSetOSProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualMachineScaleSetOSProfilePtrOutput)
}

type VirtualMachineScaleSetOSProfileOutput struct{ *pulumi.OutputState }

func (VirtualMachineScaleSetOSProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualMachineScaleSetOSProfile)(nil)).Elem()
}

func (o VirtualMachineScaleSetOSProfileOutput) ToVirtualMachineScaleSetOSProfileOutput() VirtualMachineScaleSetOSProfileOutput {
	return o
}

func (o VirtualMachineScaleSetOSProfileOutput) ToVirtualMachineScaleSetOSProfileOutputWithContext(ctx context.Context) VirtualMachineScaleSetOSProfileOutput {
	return o
}

func (o VirtualMachineScaleSetOSProfileOutput) ToVirtualMachineScaleSetOSProfilePtrOutput() VirtualMachineScaleSetOSProfilePtrOutput {
	return o.ToVirtualMachineScaleSetOSProfilePtrOutputWithContext(context.Background())
}

func (o VirtualMachineScaleSetOSProfileOutput) ToVirtualMachineScaleSetOSProfilePtrOutputWithContext(ctx context.Context) VirtualMachineScaleSetOSProfilePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v VirtualMachineScaleSetOSProfile) *VirtualMachineScaleSetOSProfile {
		return &v
	}).(VirtualMachineScaleSetOSProfilePtrOutput)
}

func (o VirtualMachineScaleSetOSProfileOutput) AdminPassword() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetOSProfile) *string { return v.AdminPassword }).(pulumi.StringPtrOutput)
}

func (o VirtualMachineScaleSetOSProfileOutput) AdminUsername() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetOSProfile) *string { return v.AdminUsername }).(pulumi.StringPtrOutput)
}

func (o VirtualMachineScaleSetOSProfileOutput) ComputerNamePrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetOSProfile) *string { return v.ComputerNamePrefix }).(pulumi.StringPtrOutput)
}

func (o VirtualMachineScaleSetOSProfileOutput) CustomData() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetOSProfile) *string { return v.CustomData }).(pulumi.StringPtrOutput)
}

func (o VirtualMachineScaleSetOSProfileOutput) LinuxConfiguration() LinuxConfigurationPtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetOSProfile) *LinuxConfiguration { return v.LinuxConfiguration }).(LinuxConfigurationPtrOutput)
}

func (o VirtualMachineScaleSetOSProfileOutput) Secrets() VaultSecretGroupArrayOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetOSProfile) []VaultSecretGroup { return v.Secrets }).(VaultSecretGroupArrayOutput)
}

func (o VirtualMachineScaleSetOSProfileOutput) WindowsConfiguration() WindowsConfigurationPtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetOSProfile) *WindowsConfiguration { return v.WindowsConfiguration }).(WindowsConfigurationPtrOutput)
}

type VirtualMachineScaleSetOSProfilePtrOutput struct{ *pulumi.OutputState }

func (VirtualMachineScaleSetOSProfilePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualMachineScaleSetOSProfile)(nil)).Elem()
}

func (o VirtualMachineScaleSetOSProfilePtrOutput) ToVirtualMachineScaleSetOSProfilePtrOutput() VirtualMachineScaleSetOSProfilePtrOutput {
	return o
}

func (o VirtualMachineScaleSetOSProfilePtrOutput) ToVirtualMachineScaleSetOSProfilePtrOutputWithContext(ctx context.Context) VirtualMachineScaleSetOSProfilePtrOutput {
	return o
}

func (o VirtualMachineScaleSetOSProfilePtrOutput) Elem() VirtualMachineScaleSetOSProfileOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetOSProfile) VirtualMachineScaleSetOSProfile {
		if v != nil {
			return *v
		}
		var ret VirtualMachineScaleSetOSProfile
		return ret
	}).(VirtualMachineScaleSetOSProfileOutput)
}

func (o VirtualMachineScaleSetOSProfilePtrOutput) AdminPassword() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetOSProfile) *string {
		if v == nil {
			return nil
		}
		return v.AdminPassword
	}).(pulumi.StringPtrOutput)
}

func (o VirtualMachineScaleSetOSProfilePtrOutput) AdminUsername() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetOSProfile) *string {
		if v == nil {
			return nil
		}
		return v.AdminUsername
	}).(pulumi.StringPtrOutput)
}

func (o VirtualMachineScaleSetOSProfilePtrOutput) ComputerNamePrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetOSProfile) *string {
		if v == nil {
			return nil
		}
		return v.ComputerNamePrefix
	}).(pulumi.StringPtrOutput)
}

func (o VirtualMachineScaleSetOSProfilePtrOutput) CustomData() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetOSProfile) *string {
		if v == nil {
			return nil
		}
		return v.CustomData
	}).(pulumi.StringPtrOutput)
}

func (o VirtualMachineScaleSetOSProfilePtrOutput) LinuxConfiguration() LinuxConfigurationPtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetOSProfile) *LinuxConfiguration {
		if v == nil {
			return nil
		}
		return v.LinuxConfiguration
	}).(LinuxConfigurationPtrOutput)
}

func (o VirtualMachineScaleSetOSProfilePtrOutput) Secrets() VaultSecretGroupArrayOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetOSProfile) []VaultSecretGroup {
		if v == nil {
			return nil
		}
		return v.Secrets
	}).(VaultSecretGroupArrayOutput)
}

func (o VirtualMachineScaleSetOSProfilePtrOutput) WindowsConfiguration() WindowsConfigurationPtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetOSProfile) *WindowsConfiguration {
		if v == nil {
			return nil
		}
		return v.WindowsConfiguration
	}).(WindowsConfigurationPtrOutput)
}

type VirtualMachineScaleSetOSProfileResponse struct {
	AdminPassword        *string                       `pulumi:"adminPassword"`
	AdminUsername        *string                       `pulumi:"adminUsername"`
	ComputerNamePrefix   *string                       `pulumi:"computerNamePrefix"`
	CustomData           *string                       `pulumi:"customData"`
	LinuxConfiguration   *LinuxConfigurationResponse   `pulumi:"linuxConfiguration"`
	Secrets              []VaultSecretGroupResponse    `pulumi:"secrets"`
	WindowsConfiguration *WindowsConfigurationResponse `pulumi:"windowsConfiguration"`
}

type VirtualMachineScaleSetOSProfileResponseOutput struct{ *pulumi.OutputState }

func (VirtualMachineScaleSetOSProfileResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualMachineScaleSetOSProfileResponse)(nil)).Elem()
}

func (o VirtualMachineScaleSetOSProfileResponseOutput) ToVirtualMachineScaleSetOSProfileResponseOutput() VirtualMachineScaleSetOSProfileResponseOutput {
	return o
}

func (o VirtualMachineScaleSetOSProfileResponseOutput) ToVirtualMachineScaleSetOSProfileResponseOutputWithContext(ctx context.Context) VirtualMachineScaleSetOSProfileResponseOutput {
	return o
}

func (o VirtualMachineScaleSetOSProfileResponseOutput) AdminPassword() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetOSProfileResponse) *string { return v.AdminPassword }).(pulumi.StringPtrOutput)
}

func (o VirtualMachineScaleSetOSProfileResponseOutput) AdminUsername() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetOSProfileResponse) *string { return v.AdminUsername }).(pulumi.StringPtrOutput)
}

func (o VirtualMachineScaleSetOSProfileResponseOutput) ComputerNamePrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetOSProfileResponse) *string { return v.ComputerNamePrefix }).(pulumi.StringPtrOutput)
}

func (o VirtualMachineScaleSetOSProfileResponseOutput) CustomData() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetOSProfileResponse) *string { return v.CustomData }).(pulumi.StringPtrOutput)
}

func (o VirtualMachineScaleSetOSProfileResponseOutput) LinuxConfiguration() LinuxConfigurationResponsePtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetOSProfileResponse) *LinuxConfigurationResponse {
		return v.LinuxConfiguration
	}).(LinuxConfigurationResponsePtrOutput)
}

func (o VirtualMachineScaleSetOSProfileResponseOutput) Secrets() VaultSecretGroupResponseArrayOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetOSProfileResponse) []VaultSecretGroupResponse { return v.Secrets }).(VaultSecretGroupResponseArrayOutput)
}

func (o VirtualMachineScaleSetOSProfileResponseOutput) WindowsConfiguration() WindowsConfigurationResponsePtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetOSProfileResponse) *WindowsConfigurationResponse {
		return v.WindowsConfiguration
	}).(WindowsConfigurationResponsePtrOutput)
}

type VirtualMachineScaleSetOSProfileResponsePtrOutput struct{ *pulumi.OutputState }

func (VirtualMachineScaleSetOSProfileResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualMachineScaleSetOSProfileResponse)(nil)).Elem()
}

func (o VirtualMachineScaleSetOSProfileResponsePtrOutput) ToVirtualMachineScaleSetOSProfileResponsePtrOutput() VirtualMachineScaleSetOSProfileResponsePtrOutput {
	return o
}

func (o VirtualMachineScaleSetOSProfileResponsePtrOutput) ToVirtualMachineScaleSetOSProfileResponsePtrOutputWithContext(ctx context.Context) VirtualMachineScaleSetOSProfileResponsePtrOutput {
	return o
}

func (o VirtualMachineScaleSetOSProfileResponsePtrOutput) Elem() VirtualMachineScaleSetOSProfileResponseOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetOSProfileResponse) VirtualMachineScaleSetOSProfileResponse {
		if v != nil {
			return *v
		}
		var ret VirtualMachineScaleSetOSProfileResponse
		return ret
	}).(VirtualMachineScaleSetOSProfileResponseOutput)
}

func (o VirtualMachineScaleSetOSProfileResponsePtrOutput) AdminPassword() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetOSProfileResponse) *string {
		if v == nil {
			return nil
		}
		return v.AdminPassword
	}).(pulumi.StringPtrOutput)
}

func (o VirtualMachineScaleSetOSProfileResponsePtrOutput) AdminUsername() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetOSProfileResponse) *string {
		if v == nil {
			return nil
		}
		return v.AdminUsername
	}).(pulumi.StringPtrOutput)
}

func (o VirtualMachineScaleSetOSProfileResponsePtrOutput) ComputerNamePrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetOSProfileResponse) *string {
		if v == nil {
			return nil
		}
		return v.ComputerNamePrefix
	}).(pulumi.StringPtrOutput)
}

func (o VirtualMachineScaleSetOSProfileResponsePtrOutput) CustomData() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetOSProfileResponse) *string {
		if v == nil {
			return nil
		}
		return v.CustomData
	}).(pulumi.StringPtrOutput)
}

func (o VirtualMachineScaleSetOSProfileResponsePtrOutput) LinuxConfiguration() LinuxConfigurationResponsePtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetOSProfileResponse) *LinuxConfigurationResponse {
		if v == nil {
			return nil
		}
		return v.LinuxConfiguration
	}).(LinuxConfigurationResponsePtrOutput)
}

func (o VirtualMachineScaleSetOSProfileResponsePtrOutput) Secrets() VaultSecretGroupResponseArrayOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetOSProfileResponse) []VaultSecretGroupResponse {
		if v == nil {
			return nil
		}
		return v.Secrets
	}).(VaultSecretGroupResponseArrayOutput)
}

func (o VirtualMachineScaleSetOSProfileResponsePtrOutput) WindowsConfiguration() WindowsConfigurationResponsePtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetOSProfileResponse) *WindowsConfigurationResponse {
		if v == nil {
			return nil
		}
		return v.WindowsConfiguration
	}).(WindowsConfigurationResponsePtrOutput)
}

type VirtualMachineScaleSetPublicIPAddressConfiguration struct {
	DnsSettings          *VirtualMachineScaleSetPublicIPAddressConfigurationDnsSettings `pulumi:"dnsSettings"`
	IdleTimeoutInMinutes *int                                                           `pulumi:"idleTimeoutInMinutes"`
	IpTags               []VirtualMachineScaleSetIpTag                                  `pulumi:"ipTags"`
	Name                 string                                                         `pulumi:"name"`
}





type VirtualMachineScaleSetPublicIPAddressConfigurationInput interface {
	pulumi.Input

	ToVirtualMachineScaleSetPublicIPAddressConfigurationOutput() VirtualMachineScaleSetPublicIPAddressConfigurationOutput
	ToVirtualMachineScaleSetPublicIPAddressConfigurationOutputWithContext(context.Context) VirtualMachineScaleSetPublicIPAddressConfigurationOutput
}

type VirtualMachineScaleSetPublicIPAddressConfigurationArgs struct {
	DnsSettings          VirtualMachineScaleSetPublicIPAddressConfigurationDnsSettingsPtrInput `pulumi:"dnsSettings"`
	IdleTimeoutInMinutes pulumi.IntPtrInput                                                    `pulumi:"idleTimeoutInMinutes"`
	IpTags               VirtualMachineScaleSetIpTagArrayInput                                 `pulumi:"ipTags"`
	Name                 pulumi.StringInput                                                    `pulumi:"name"`
}

func (VirtualMachineScaleSetPublicIPAddressConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualMachineScaleSetPublicIPAddressConfiguration)(nil)).Elem()
}

func (i VirtualMachineScaleSetPublicIPAddressConfigurationArgs) ToVirtualMachineScaleSetPublicIPAddressConfigurationOutput() VirtualMachineScaleSetPublicIPAddressConfigurationOutput {
	return i.ToVirtualMachineScaleSetPublicIPAddressConfigurationOutputWithContext(context.Background())
}

func (i VirtualMachineScaleSetPublicIPAddressConfigurationArgs) ToVirtualMachineScaleSetPublicIPAddressConfigurationOutputWithContext(ctx context.Context) VirtualMachineScaleSetPublicIPAddressConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualMachineScaleSetPublicIPAddressConfigurationOutput)
}

func (i VirtualMachineScaleSetPublicIPAddressConfigurationArgs) ToVirtualMachineScaleSetPublicIPAddressConfigurationPtrOutput() VirtualMachineScaleSetPublicIPAddressConfigurationPtrOutput {
	return i.ToVirtualMachineScaleSetPublicIPAddressConfigurationPtrOutputWithContext(context.Background())
}

func (i VirtualMachineScaleSetPublicIPAddressConfigurationArgs) ToVirtualMachineScaleSetPublicIPAddressConfigurationPtrOutputWithContext(ctx context.Context) VirtualMachineScaleSetPublicIPAddressConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualMachineScaleSetPublicIPAddressConfigurationOutput).ToVirtualMachineScaleSetPublicIPAddressConfigurationPtrOutputWithContext(ctx)
}









type VirtualMachineScaleSetPublicIPAddressConfigurationPtrInput interface {
	pulumi.Input

	ToVirtualMachineScaleSetPublicIPAddressConfigurationPtrOutput() VirtualMachineScaleSetPublicIPAddressConfigurationPtrOutput
	ToVirtualMachineScaleSetPublicIPAddressConfigurationPtrOutputWithContext(context.Context) VirtualMachineScaleSetPublicIPAddressConfigurationPtrOutput
}

type virtualMachineScaleSetPublicIPAddressConfigurationPtrType VirtualMachineScaleSetPublicIPAddressConfigurationArgs

func VirtualMachineScaleSetPublicIPAddressConfigurationPtr(v *VirtualMachineScaleSetPublicIPAddressConfigurationArgs) VirtualMachineScaleSetPublicIPAddressConfigurationPtrInput {
	return (*virtualMachineScaleSetPublicIPAddressConfigurationPtrType)(v)
}

func (*virtualMachineScaleSetPublicIPAddressConfigurationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualMachineScaleSetPublicIPAddressConfiguration)(nil)).Elem()
}

func (i *virtualMachineScaleSetPublicIPAddressConfigurationPtrType) ToVirtualMachineScaleSetPublicIPAddressConfigurationPtrOutput() VirtualMachineScaleSetPublicIPAddressConfigurationPtrOutput {
	return i.ToVirtualMachineScaleSetPublicIPAddressConfigurationPtrOutputWithContext(context.Background())
}

func (i *virtualMachineScaleSetPublicIPAddressConfigurationPtrType) ToVirtualMachineScaleSetPublicIPAddressConfigurationPtrOutputWithContext(ctx context.Context) VirtualMachineScaleSetPublicIPAddressConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualMachineScaleSetPublicIPAddressConfigurationPtrOutput)
}

type VirtualMachineScaleSetPublicIPAddressConfigurationOutput struct{ *pulumi.OutputState }

func (VirtualMachineScaleSetPublicIPAddressConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualMachineScaleSetPublicIPAddressConfiguration)(nil)).Elem()
}

func (o VirtualMachineScaleSetPublicIPAddressConfigurationOutput) ToVirtualMachineScaleSetPublicIPAddressConfigurationOutput() VirtualMachineScaleSetPublicIPAddressConfigurationOutput {
	return o
}

func (o VirtualMachineScaleSetPublicIPAddressConfigurationOutput) ToVirtualMachineScaleSetPublicIPAddressConfigurationOutputWithContext(ctx context.Context) VirtualMachineScaleSetPublicIPAddressConfigurationOutput {
	return o
}

func (o VirtualMachineScaleSetPublicIPAddressConfigurationOutput) ToVirtualMachineScaleSetPublicIPAddressConfigurationPtrOutput() VirtualMachineScaleSetPublicIPAddressConfigurationPtrOutput {
	return o.ToVirtualMachineScaleSetPublicIPAddressConfigurationPtrOutputWithContext(context.Background())
}

func (o VirtualMachineScaleSetPublicIPAddressConfigurationOutput) ToVirtualMachineScaleSetPublicIPAddressConfigurationPtrOutputWithContext(ctx context.Context) VirtualMachineScaleSetPublicIPAddressConfigurationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v VirtualMachineScaleSetPublicIPAddressConfiguration) *VirtualMachineScaleSetPublicIPAddressConfiguration {
		return &v
	}).(VirtualMachineScaleSetPublicIPAddressConfigurationPtrOutput)
}

func (o VirtualMachineScaleSetPublicIPAddressConfigurationOutput) DnsSettings() VirtualMachineScaleSetPublicIPAddressConfigurationDnsSettingsPtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetPublicIPAddressConfiguration) *VirtualMachineScaleSetPublicIPAddressConfigurationDnsSettings {
		return v.DnsSettings
	}).(VirtualMachineScaleSetPublicIPAddressConfigurationDnsSettingsPtrOutput)
}

func (o VirtualMachineScaleSetPublicIPAddressConfigurationOutput) IdleTimeoutInMinutes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetPublicIPAddressConfiguration) *int { return v.IdleTimeoutInMinutes }).(pulumi.IntPtrOutput)
}

func (o VirtualMachineScaleSetPublicIPAddressConfigurationOutput) IpTags() VirtualMachineScaleSetIpTagArrayOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetPublicIPAddressConfiguration) []VirtualMachineScaleSetIpTag {
		return v.IpTags
	}).(VirtualMachineScaleSetIpTagArrayOutput)
}

func (o VirtualMachineScaleSetPublicIPAddressConfigurationOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetPublicIPAddressConfiguration) string { return v.Name }).(pulumi.StringOutput)
}

type VirtualMachineScaleSetPublicIPAddressConfigurationPtrOutput struct{ *pulumi.OutputState }

func (VirtualMachineScaleSetPublicIPAddressConfigurationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualMachineScaleSetPublicIPAddressConfiguration)(nil)).Elem()
}

func (o VirtualMachineScaleSetPublicIPAddressConfigurationPtrOutput) ToVirtualMachineScaleSetPublicIPAddressConfigurationPtrOutput() VirtualMachineScaleSetPublicIPAddressConfigurationPtrOutput {
	return o
}

func (o VirtualMachineScaleSetPublicIPAddressConfigurationPtrOutput) ToVirtualMachineScaleSetPublicIPAddressConfigurationPtrOutputWithContext(ctx context.Context) VirtualMachineScaleSetPublicIPAddressConfigurationPtrOutput {
	return o
}

func (o VirtualMachineScaleSetPublicIPAddressConfigurationPtrOutput) Elem() VirtualMachineScaleSetPublicIPAddressConfigurationOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetPublicIPAddressConfiguration) VirtualMachineScaleSetPublicIPAddressConfiguration {
		if v != nil {
			return *v
		}
		var ret VirtualMachineScaleSetPublicIPAddressConfiguration
		return ret
	}).(VirtualMachineScaleSetPublicIPAddressConfigurationOutput)
}

func (o VirtualMachineScaleSetPublicIPAddressConfigurationPtrOutput) DnsSettings() VirtualMachineScaleSetPublicIPAddressConfigurationDnsSettingsPtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetPublicIPAddressConfiguration) *VirtualMachineScaleSetPublicIPAddressConfigurationDnsSettings {
		if v == nil {
			return nil
		}
		return v.DnsSettings
	}).(VirtualMachineScaleSetPublicIPAddressConfigurationDnsSettingsPtrOutput)
}

func (o VirtualMachineScaleSetPublicIPAddressConfigurationPtrOutput) IdleTimeoutInMinutes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetPublicIPAddressConfiguration) *int {
		if v == nil {
			return nil
		}
		return v.IdleTimeoutInMinutes
	}).(pulumi.IntPtrOutput)
}

func (o VirtualMachineScaleSetPublicIPAddressConfigurationPtrOutput) IpTags() VirtualMachineScaleSetIpTagArrayOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetPublicIPAddressConfiguration) []VirtualMachineScaleSetIpTag {
		if v == nil {
			return nil
		}
		return v.IpTags
	}).(VirtualMachineScaleSetIpTagArrayOutput)
}

func (o VirtualMachineScaleSetPublicIPAddressConfigurationPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetPublicIPAddressConfiguration) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

type VirtualMachineScaleSetPublicIPAddressConfigurationDnsSettings struct {
	DomainNameLabel string `pulumi:"domainNameLabel"`
}





type VirtualMachineScaleSetPublicIPAddressConfigurationDnsSettingsInput interface {
	pulumi.Input

	ToVirtualMachineScaleSetPublicIPAddressConfigurationDnsSettingsOutput() VirtualMachineScaleSetPublicIPAddressConfigurationDnsSettingsOutput
	ToVirtualMachineScaleSetPublicIPAddressConfigurationDnsSettingsOutputWithContext(context.Context) VirtualMachineScaleSetPublicIPAddressConfigurationDnsSettingsOutput
}

type VirtualMachineScaleSetPublicIPAddressConfigurationDnsSettingsArgs struct {
	DomainNameLabel pulumi.StringInput `pulumi:"domainNameLabel"`
}

func (VirtualMachineScaleSetPublicIPAddressConfigurationDnsSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualMachineScaleSetPublicIPAddressConfigurationDnsSettings)(nil)).Elem()
}

func (i VirtualMachineScaleSetPublicIPAddressConfigurationDnsSettingsArgs) ToVirtualMachineScaleSetPublicIPAddressConfigurationDnsSettingsOutput() VirtualMachineScaleSetPublicIPAddressConfigurationDnsSettingsOutput {
	return i.ToVirtualMachineScaleSetPublicIPAddressConfigurationDnsSettingsOutputWithContext(context.Background())
}

func (i VirtualMachineScaleSetPublicIPAddressConfigurationDnsSettingsArgs) ToVirtualMachineScaleSetPublicIPAddressConfigurationDnsSettingsOutputWithContext(ctx context.Context) VirtualMachineScaleSetPublicIPAddressConfigurationDnsSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualMachineScaleSetPublicIPAddressConfigurationDnsSettingsOutput)
}

func (i VirtualMachineScaleSetPublicIPAddressConfigurationDnsSettingsArgs) ToVirtualMachineScaleSetPublicIPAddressConfigurationDnsSettingsPtrOutput() VirtualMachineScaleSetPublicIPAddressConfigurationDnsSettingsPtrOutput {
	return i.ToVirtualMachineScaleSetPublicIPAddressConfigurationDnsSettingsPtrOutputWithContext(context.Background())
}

func (i VirtualMachineScaleSetPublicIPAddressConfigurationDnsSettingsArgs) ToVirtualMachineScaleSetPublicIPAddressConfigurationDnsSettingsPtrOutputWithContext(ctx context.Context) VirtualMachineScaleSetPublicIPAddressConfigurationDnsSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualMachineScaleSetPublicIPAddressConfigurationDnsSettingsOutput).ToVirtualMachineScaleSetPublicIPAddressConfigurationDnsSettingsPtrOutputWithContext(ctx)
}









type VirtualMachineScaleSetPublicIPAddressConfigurationDnsSettingsPtrInput interface {
	pulumi.Input

	ToVirtualMachineScaleSetPublicIPAddressConfigurationDnsSettingsPtrOutput() VirtualMachineScaleSetPublicIPAddressConfigurationDnsSettingsPtrOutput
	ToVirtualMachineScaleSetPublicIPAddressConfigurationDnsSettingsPtrOutputWithContext(context.Context) VirtualMachineScaleSetPublicIPAddressConfigurationDnsSettingsPtrOutput
}

type virtualMachineScaleSetPublicIPAddressConfigurationDnsSettingsPtrType VirtualMachineScaleSetPublicIPAddressConfigurationDnsSettingsArgs

func VirtualMachineScaleSetPublicIPAddressConfigurationDnsSettingsPtr(v *VirtualMachineScaleSetPublicIPAddressConfigurationDnsSettingsArgs) VirtualMachineScaleSetPublicIPAddressConfigurationDnsSettingsPtrInput {
	return (*virtualMachineScaleSetPublicIPAddressConfigurationDnsSettingsPtrType)(v)
}

func (*virtualMachineScaleSetPublicIPAddressConfigurationDnsSettingsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualMachineScaleSetPublicIPAddressConfigurationDnsSettings)(nil)).Elem()
}

func (i *virtualMachineScaleSetPublicIPAddressConfigurationDnsSettingsPtrType) ToVirtualMachineScaleSetPublicIPAddressConfigurationDnsSettingsPtrOutput() VirtualMachineScaleSetPublicIPAddressConfigurationDnsSettingsPtrOutput {
	return i.ToVirtualMachineScaleSetPublicIPAddressConfigurationDnsSettingsPtrOutputWithContext(context.Background())
}

func (i *virtualMachineScaleSetPublicIPAddressConfigurationDnsSettingsPtrType) ToVirtualMachineScaleSetPublicIPAddressConfigurationDnsSettingsPtrOutputWithContext(ctx context.Context) VirtualMachineScaleSetPublicIPAddressConfigurationDnsSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualMachineScaleSetPublicIPAddressConfigurationDnsSettingsPtrOutput)
}

type VirtualMachineScaleSetPublicIPAddressConfigurationDnsSettingsOutput struct{ *pulumi.OutputState }

func (VirtualMachineScaleSetPublicIPAddressConfigurationDnsSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualMachineScaleSetPublicIPAddressConfigurationDnsSettings)(nil)).Elem()
}

func (o VirtualMachineScaleSetPublicIPAddressConfigurationDnsSettingsOutput) ToVirtualMachineScaleSetPublicIPAddressConfigurationDnsSettingsOutput() VirtualMachineScaleSetPublicIPAddressConfigurationDnsSettingsOutput {
	return o
}

func (o VirtualMachineScaleSetPublicIPAddressConfigurationDnsSettingsOutput) ToVirtualMachineScaleSetPublicIPAddressConfigurationDnsSettingsOutputWithContext(ctx context.Context) VirtualMachineScaleSetPublicIPAddressConfigurationDnsSettingsOutput {
	return o
}

func (o VirtualMachineScaleSetPublicIPAddressConfigurationDnsSettingsOutput) ToVirtualMachineScaleSetPublicIPAddressConfigurationDnsSettingsPtrOutput() VirtualMachineScaleSetPublicIPAddressConfigurationDnsSettingsPtrOutput {
	return o.ToVirtualMachineScaleSetPublicIPAddressConfigurationDnsSettingsPtrOutputWithContext(context.Background())
}

func (o VirtualMachineScaleSetPublicIPAddressConfigurationDnsSettingsOutput) ToVirtualMachineScaleSetPublicIPAddressConfigurationDnsSettingsPtrOutputWithContext(ctx context.Context) VirtualMachineScaleSetPublicIPAddressConfigurationDnsSettingsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v VirtualMachineScaleSetPublicIPAddressConfigurationDnsSettings) *VirtualMachineScaleSetPublicIPAddressConfigurationDnsSettings {
		return &v
	}).(VirtualMachineScaleSetPublicIPAddressConfigurationDnsSettingsPtrOutput)
}

func (o VirtualMachineScaleSetPublicIPAddressConfigurationDnsSettingsOutput) DomainNameLabel() pulumi.StringOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetPublicIPAddressConfigurationDnsSettings) string { return v.DomainNameLabel }).(pulumi.StringOutput)
}

type VirtualMachineScaleSetPublicIPAddressConfigurationDnsSettingsPtrOutput struct{ *pulumi.OutputState }

func (VirtualMachineScaleSetPublicIPAddressConfigurationDnsSettingsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualMachineScaleSetPublicIPAddressConfigurationDnsSettings)(nil)).Elem()
}

func (o VirtualMachineScaleSetPublicIPAddressConfigurationDnsSettingsPtrOutput) ToVirtualMachineScaleSetPublicIPAddressConfigurationDnsSettingsPtrOutput() VirtualMachineScaleSetPublicIPAddressConfigurationDnsSettingsPtrOutput {
	return o
}

func (o VirtualMachineScaleSetPublicIPAddressConfigurationDnsSettingsPtrOutput) ToVirtualMachineScaleSetPublicIPAddressConfigurationDnsSettingsPtrOutputWithContext(ctx context.Context) VirtualMachineScaleSetPublicIPAddressConfigurationDnsSettingsPtrOutput {
	return o
}

func (o VirtualMachineScaleSetPublicIPAddressConfigurationDnsSettingsPtrOutput) Elem() VirtualMachineScaleSetPublicIPAddressConfigurationDnsSettingsOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetPublicIPAddressConfigurationDnsSettings) VirtualMachineScaleSetPublicIPAddressConfigurationDnsSettings {
		if v != nil {
			return *v
		}
		var ret VirtualMachineScaleSetPublicIPAddressConfigurationDnsSettings
		return ret
	}).(VirtualMachineScaleSetPublicIPAddressConfigurationDnsSettingsOutput)
}

func (o VirtualMachineScaleSetPublicIPAddressConfigurationDnsSettingsPtrOutput) DomainNameLabel() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetPublicIPAddressConfigurationDnsSettings) *string {
		if v == nil {
			return nil
		}
		return &v.DomainNameLabel
	}).(pulumi.StringPtrOutput)
}

type VirtualMachineScaleSetPublicIPAddressConfigurationDnsSettingsResponse struct {
	DomainNameLabel string `pulumi:"domainNameLabel"`
}

type VirtualMachineScaleSetPublicIPAddressConfigurationDnsSettingsResponseOutput struct{ *pulumi.OutputState }

func (VirtualMachineScaleSetPublicIPAddressConfigurationDnsSettingsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualMachineScaleSetPublicIPAddressConfigurationDnsSettingsResponse)(nil)).Elem()
}

func (o VirtualMachineScaleSetPublicIPAddressConfigurationDnsSettingsResponseOutput) ToVirtualMachineScaleSetPublicIPAddressConfigurationDnsSettingsResponseOutput() VirtualMachineScaleSetPublicIPAddressConfigurationDnsSettingsResponseOutput {
	return o
}

func (o VirtualMachineScaleSetPublicIPAddressConfigurationDnsSettingsResponseOutput) ToVirtualMachineScaleSetPublicIPAddressConfigurationDnsSettingsResponseOutputWithContext(ctx context.Context) VirtualMachineScaleSetPublicIPAddressConfigurationDnsSettingsResponseOutput {
	return o
}

func (o VirtualMachineScaleSetPublicIPAddressConfigurationDnsSettingsResponseOutput) DomainNameLabel() pulumi.StringOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetPublicIPAddressConfigurationDnsSettingsResponse) string {
		return v.DomainNameLabel
	}).(pulumi.StringOutput)
}

type VirtualMachineScaleSetPublicIPAddressConfigurationDnsSettingsResponsePtrOutput struct{ *pulumi.OutputState }

func (VirtualMachineScaleSetPublicIPAddressConfigurationDnsSettingsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualMachineScaleSetPublicIPAddressConfigurationDnsSettingsResponse)(nil)).Elem()
}

func (o VirtualMachineScaleSetPublicIPAddressConfigurationDnsSettingsResponsePtrOutput) ToVirtualMachineScaleSetPublicIPAddressConfigurationDnsSettingsResponsePtrOutput() VirtualMachineScaleSetPublicIPAddressConfigurationDnsSettingsResponsePtrOutput {
	return o
}

func (o VirtualMachineScaleSetPublicIPAddressConfigurationDnsSettingsResponsePtrOutput) ToVirtualMachineScaleSetPublicIPAddressConfigurationDnsSettingsResponsePtrOutputWithContext(ctx context.Context) VirtualMachineScaleSetPublicIPAddressConfigurationDnsSettingsResponsePtrOutput {
	return o
}

func (o VirtualMachineScaleSetPublicIPAddressConfigurationDnsSettingsResponsePtrOutput) Elem() VirtualMachineScaleSetPublicIPAddressConfigurationDnsSettingsResponseOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetPublicIPAddressConfigurationDnsSettingsResponse) VirtualMachineScaleSetPublicIPAddressConfigurationDnsSettingsResponse {
		if v != nil {
			return *v
		}
		var ret VirtualMachineScaleSetPublicIPAddressConfigurationDnsSettingsResponse
		return ret
	}).(VirtualMachineScaleSetPublicIPAddressConfigurationDnsSettingsResponseOutput)
}

func (o VirtualMachineScaleSetPublicIPAddressConfigurationDnsSettingsResponsePtrOutput) DomainNameLabel() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetPublicIPAddressConfigurationDnsSettingsResponse) *string {
		if v == nil {
			return nil
		}
		return &v.DomainNameLabel
	}).(pulumi.StringPtrOutput)
}

type VirtualMachineScaleSetPublicIPAddressConfigurationResponse struct {
	DnsSettings          *VirtualMachineScaleSetPublicIPAddressConfigurationDnsSettingsResponse `pulumi:"dnsSettings"`
	IdleTimeoutInMinutes *int                                                                   `pulumi:"idleTimeoutInMinutes"`
	IpTags               []VirtualMachineScaleSetIpTagResponse                                  `pulumi:"ipTags"`
	Name                 string                                                                 `pulumi:"name"`
}

type VirtualMachineScaleSetPublicIPAddressConfigurationResponseOutput struct{ *pulumi.OutputState }

func (VirtualMachineScaleSetPublicIPAddressConfigurationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualMachineScaleSetPublicIPAddressConfigurationResponse)(nil)).Elem()
}

func (o VirtualMachineScaleSetPublicIPAddressConfigurationResponseOutput) ToVirtualMachineScaleSetPublicIPAddressConfigurationResponseOutput() VirtualMachineScaleSetPublicIPAddressConfigurationResponseOutput {
	return o
}

func (o VirtualMachineScaleSetPublicIPAddressConfigurationResponseOutput) ToVirtualMachineScaleSetPublicIPAddressConfigurationResponseOutputWithContext(ctx context.Context) VirtualMachineScaleSetPublicIPAddressConfigurationResponseOutput {
	return o
}

func (o VirtualMachineScaleSetPublicIPAddressConfigurationResponseOutput) DnsSettings() VirtualMachineScaleSetPublicIPAddressConfigurationDnsSettingsResponsePtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetPublicIPAddressConfigurationResponse) *VirtualMachineScaleSetPublicIPAddressConfigurationDnsSettingsResponse {
		return v.DnsSettings
	}).(VirtualMachineScaleSetPublicIPAddressConfigurationDnsSettingsResponsePtrOutput)
}

func (o VirtualMachineScaleSetPublicIPAddressConfigurationResponseOutput) IdleTimeoutInMinutes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetPublicIPAddressConfigurationResponse) *int { return v.IdleTimeoutInMinutes }).(pulumi.IntPtrOutput)
}

func (o VirtualMachineScaleSetPublicIPAddressConfigurationResponseOutput) IpTags() VirtualMachineScaleSetIpTagResponseArrayOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetPublicIPAddressConfigurationResponse) []VirtualMachineScaleSetIpTagResponse {
		return v.IpTags
	}).(VirtualMachineScaleSetIpTagResponseArrayOutput)
}

func (o VirtualMachineScaleSetPublicIPAddressConfigurationResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetPublicIPAddressConfigurationResponse) string { return v.Name }).(pulumi.StringOutput)
}

type VirtualMachineScaleSetPublicIPAddressConfigurationResponsePtrOutput struct{ *pulumi.OutputState }

func (VirtualMachineScaleSetPublicIPAddressConfigurationResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualMachineScaleSetPublicIPAddressConfigurationResponse)(nil)).Elem()
}

func (o VirtualMachineScaleSetPublicIPAddressConfigurationResponsePtrOutput) ToVirtualMachineScaleSetPublicIPAddressConfigurationResponsePtrOutput() VirtualMachineScaleSetPublicIPAddressConfigurationResponsePtrOutput {
	return o
}

func (o VirtualMachineScaleSetPublicIPAddressConfigurationResponsePtrOutput) ToVirtualMachineScaleSetPublicIPAddressConfigurationResponsePtrOutputWithContext(ctx context.Context) VirtualMachineScaleSetPublicIPAddressConfigurationResponsePtrOutput {
	return o
}

func (o VirtualMachineScaleSetPublicIPAddressConfigurationResponsePtrOutput) Elem() VirtualMachineScaleSetPublicIPAddressConfigurationResponseOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetPublicIPAddressConfigurationResponse) VirtualMachineScaleSetPublicIPAddressConfigurationResponse {
		if v != nil {
			return *v
		}
		var ret VirtualMachineScaleSetPublicIPAddressConfigurationResponse
		return ret
	}).(VirtualMachineScaleSetPublicIPAddressConfigurationResponseOutput)
}

func (o VirtualMachineScaleSetPublicIPAddressConfigurationResponsePtrOutput) DnsSettings() VirtualMachineScaleSetPublicIPAddressConfigurationDnsSettingsResponsePtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetPublicIPAddressConfigurationResponse) *VirtualMachineScaleSetPublicIPAddressConfigurationDnsSettingsResponse {
		if v == nil {
			return nil
		}
		return v.DnsSettings
	}).(VirtualMachineScaleSetPublicIPAddressConfigurationDnsSettingsResponsePtrOutput)
}

func (o VirtualMachineScaleSetPublicIPAddressConfigurationResponsePtrOutput) IdleTimeoutInMinutes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetPublicIPAddressConfigurationResponse) *int {
		if v == nil {
			return nil
		}
		return v.IdleTimeoutInMinutes
	}).(pulumi.IntPtrOutput)
}

func (o VirtualMachineScaleSetPublicIPAddressConfigurationResponsePtrOutput) IpTags() VirtualMachineScaleSetIpTagResponseArrayOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetPublicIPAddressConfigurationResponse) []VirtualMachineScaleSetIpTagResponse {
		if v == nil {
			return nil
		}
		return v.IpTags
	}).(VirtualMachineScaleSetIpTagResponseArrayOutput)
}

func (o VirtualMachineScaleSetPublicIPAddressConfigurationResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetPublicIPAddressConfigurationResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

type VirtualMachineScaleSetStorageProfile struct {
	DataDisks      []VirtualMachineScaleSetDataDisk `pulumi:"dataDisks"`
	ImageReference *ImageReference                  `pulumi:"imageReference"`
	OsDisk         *VirtualMachineScaleSetOSDisk    `pulumi:"osDisk"`
}





type VirtualMachineScaleSetStorageProfileInput interface {
	pulumi.Input

	ToVirtualMachineScaleSetStorageProfileOutput() VirtualMachineScaleSetStorageProfileOutput
	ToVirtualMachineScaleSetStorageProfileOutputWithContext(context.Context) VirtualMachineScaleSetStorageProfileOutput
}

type VirtualMachineScaleSetStorageProfileArgs struct {
	DataDisks      VirtualMachineScaleSetDataDiskArrayInput `pulumi:"dataDisks"`
	ImageReference ImageReferencePtrInput                   `pulumi:"imageReference"`
	OsDisk         VirtualMachineScaleSetOSDiskPtrInput     `pulumi:"osDisk"`
}

func (VirtualMachineScaleSetStorageProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualMachineScaleSetStorageProfile)(nil)).Elem()
}

func (i VirtualMachineScaleSetStorageProfileArgs) ToVirtualMachineScaleSetStorageProfileOutput() VirtualMachineScaleSetStorageProfileOutput {
	return i.ToVirtualMachineScaleSetStorageProfileOutputWithContext(context.Background())
}

func (i VirtualMachineScaleSetStorageProfileArgs) ToVirtualMachineScaleSetStorageProfileOutputWithContext(ctx context.Context) VirtualMachineScaleSetStorageProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualMachineScaleSetStorageProfileOutput)
}

func (i VirtualMachineScaleSetStorageProfileArgs) ToVirtualMachineScaleSetStorageProfilePtrOutput() VirtualMachineScaleSetStorageProfilePtrOutput {
	return i.ToVirtualMachineScaleSetStorageProfilePtrOutputWithContext(context.Background())
}

func (i VirtualMachineScaleSetStorageProfileArgs) ToVirtualMachineScaleSetStorageProfilePtrOutputWithContext(ctx context.Context) VirtualMachineScaleSetStorageProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualMachineScaleSetStorageProfileOutput).ToVirtualMachineScaleSetStorageProfilePtrOutputWithContext(ctx)
}









type VirtualMachineScaleSetStorageProfilePtrInput interface {
	pulumi.Input

	ToVirtualMachineScaleSetStorageProfilePtrOutput() VirtualMachineScaleSetStorageProfilePtrOutput
	ToVirtualMachineScaleSetStorageProfilePtrOutputWithContext(context.Context) VirtualMachineScaleSetStorageProfilePtrOutput
}

type virtualMachineScaleSetStorageProfilePtrType VirtualMachineScaleSetStorageProfileArgs

func VirtualMachineScaleSetStorageProfilePtr(v *VirtualMachineScaleSetStorageProfileArgs) VirtualMachineScaleSetStorageProfilePtrInput {
	return (*virtualMachineScaleSetStorageProfilePtrType)(v)
}

func (*virtualMachineScaleSetStorageProfilePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualMachineScaleSetStorageProfile)(nil)).Elem()
}

func (i *virtualMachineScaleSetStorageProfilePtrType) ToVirtualMachineScaleSetStorageProfilePtrOutput() VirtualMachineScaleSetStorageProfilePtrOutput {
	return i.ToVirtualMachineScaleSetStorageProfilePtrOutputWithContext(context.Background())
}

func (i *virtualMachineScaleSetStorageProfilePtrType) ToVirtualMachineScaleSetStorageProfilePtrOutputWithContext(ctx context.Context) VirtualMachineScaleSetStorageProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualMachineScaleSetStorageProfilePtrOutput)
}

type VirtualMachineScaleSetStorageProfileOutput struct{ *pulumi.OutputState }

func (VirtualMachineScaleSetStorageProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualMachineScaleSetStorageProfile)(nil)).Elem()
}

func (o VirtualMachineScaleSetStorageProfileOutput) ToVirtualMachineScaleSetStorageProfileOutput() VirtualMachineScaleSetStorageProfileOutput {
	return o
}

func (o VirtualMachineScaleSetStorageProfileOutput) ToVirtualMachineScaleSetStorageProfileOutputWithContext(ctx context.Context) VirtualMachineScaleSetStorageProfileOutput {
	return o
}

func (o VirtualMachineScaleSetStorageProfileOutput) ToVirtualMachineScaleSetStorageProfilePtrOutput() VirtualMachineScaleSetStorageProfilePtrOutput {
	return o.ToVirtualMachineScaleSetStorageProfilePtrOutputWithContext(context.Background())
}

func (o VirtualMachineScaleSetStorageProfileOutput) ToVirtualMachineScaleSetStorageProfilePtrOutputWithContext(ctx context.Context) VirtualMachineScaleSetStorageProfilePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v VirtualMachineScaleSetStorageProfile) *VirtualMachineScaleSetStorageProfile {
		return &v
	}).(VirtualMachineScaleSetStorageProfilePtrOutput)
}

func (o VirtualMachineScaleSetStorageProfileOutput) DataDisks() VirtualMachineScaleSetDataDiskArrayOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetStorageProfile) []VirtualMachineScaleSetDataDisk { return v.DataDisks }).(VirtualMachineScaleSetDataDiskArrayOutput)
}

func (o VirtualMachineScaleSetStorageProfileOutput) ImageReference() ImageReferencePtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetStorageProfile) *ImageReference { return v.ImageReference }).(ImageReferencePtrOutput)
}

func (o VirtualMachineScaleSetStorageProfileOutput) OsDisk() VirtualMachineScaleSetOSDiskPtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetStorageProfile) *VirtualMachineScaleSetOSDisk { return v.OsDisk }).(VirtualMachineScaleSetOSDiskPtrOutput)
}

type VirtualMachineScaleSetStorageProfilePtrOutput struct{ *pulumi.OutputState }

func (VirtualMachineScaleSetStorageProfilePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualMachineScaleSetStorageProfile)(nil)).Elem()
}

func (o VirtualMachineScaleSetStorageProfilePtrOutput) ToVirtualMachineScaleSetStorageProfilePtrOutput() VirtualMachineScaleSetStorageProfilePtrOutput {
	return o
}

func (o VirtualMachineScaleSetStorageProfilePtrOutput) ToVirtualMachineScaleSetStorageProfilePtrOutputWithContext(ctx context.Context) VirtualMachineScaleSetStorageProfilePtrOutput {
	return o
}

func (o VirtualMachineScaleSetStorageProfilePtrOutput) Elem() VirtualMachineScaleSetStorageProfileOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetStorageProfile) VirtualMachineScaleSetStorageProfile {
		if v != nil {
			return *v
		}
		var ret VirtualMachineScaleSetStorageProfile
		return ret
	}).(VirtualMachineScaleSetStorageProfileOutput)
}

func (o VirtualMachineScaleSetStorageProfilePtrOutput) DataDisks() VirtualMachineScaleSetDataDiskArrayOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetStorageProfile) []VirtualMachineScaleSetDataDisk {
		if v == nil {
			return nil
		}
		return v.DataDisks
	}).(VirtualMachineScaleSetDataDiskArrayOutput)
}

func (o VirtualMachineScaleSetStorageProfilePtrOutput) ImageReference() ImageReferencePtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetStorageProfile) *ImageReference {
		if v == nil {
			return nil
		}
		return v.ImageReference
	}).(ImageReferencePtrOutput)
}

func (o VirtualMachineScaleSetStorageProfilePtrOutput) OsDisk() VirtualMachineScaleSetOSDiskPtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetStorageProfile) *VirtualMachineScaleSetOSDisk {
		if v == nil {
			return nil
		}
		return v.OsDisk
	}).(VirtualMachineScaleSetOSDiskPtrOutput)
}

type VirtualMachineScaleSetStorageProfileResponse struct {
	DataDisks      []VirtualMachineScaleSetDataDiskResponse `pulumi:"dataDisks"`
	ImageReference *ImageReferenceResponse                  `pulumi:"imageReference"`
	OsDisk         *VirtualMachineScaleSetOSDiskResponse    `pulumi:"osDisk"`
}

type VirtualMachineScaleSetStorageProfileResponseOutput struct{ *pulumi.OutputState }

func (VirtualMachineScaleSetStorageProfileResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualMachineScaleSetStorageProfileResponse)(nil)).Elem()
}

func (o VirtualMachineScaleSetStorageProfileResponseOutput) ToVirtualMachineScaleSetStorageProfileResponseOutput() VirtualMachineScaleSetStorageProfileResponseOutput {
	return o
}

func (o VirtualMachineScaleSetStorageProfileResponseOutput) ToVirtualMachineScaleSetStorageProfileResponseOutputWithContext(ctx context.Context) VirtualMachineScaleSetStorageProfileResponseOutput {
	return o
}

func (o VirtualMachineScaleSetStorageProfileResponseOutput) DataDisks() VirtualMachineScaleSetDataDiskResponseArrayOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetStorageProfileResponse) []VirtualMachineScaleSetDataDiskResponse {
		return v.DataDisks
	}).(VirtualMachineScaleSetDataDiskResponseArrayOutput)
}

func (o VirtualMachineScaleSetStorageProfileResponseOutput) ImageReference() ImageReferenceResponsePtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetStorageProfileResponse) *ImageReferenceResponse { return v.ImageReference }).(ImageReferenceResponsePtrOutput)
}

func (o VirtualMachineScaleSetStorageProfileResponseOutput) OsDisk() VirtualMachineScaleSetOSDiskResponsePtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetStorageProfileResponse) *VirtualMachineScaleSetOSDiskResponse {
		return v.OsDisk
	}).(VirtualMachineScaleSetOSDiskResponsePtrOutput)
}

type VirtualMachineScaleSetStorageProfileResponsePtrOutput struct{ *pulumi.OutputState }

func (VirtualMachineScaleSetStorageProfileResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualMachineScaleSetStorageProfileResponse)(nil)).Elem()
}

func (o VirtualMachineScaleSetStorageProfileResponsePtrOutput) ToVirtualMachineScaleSetStorageProfileResponsePtrOutput() VirtualMachineScaleSetStorageProfileResponsePtrOutput {
	return o
}

func (o VirtualMachineScaleSetStorageProfileResponsePtrOutput) ToVirtualMachineScaleSetStorageProfileResponsePtrOutputWithContext(ctx context.Context) VirtualMachineScaleSetStorageProfileResponsePtrOutput {
	return o
}

func (o VirtualMachineScaleSetStorageProfileResponsePtrOutput) Elem() VirtualMachineScaleSetStorageProfileResponseOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetStorageProfileResponse) VirtualMachineScaleSetStorageProfileResponse {
		if v != nil {
			return *v
		}
		var ret VirtualMachineScaleSetStorageProfileResponse
		return ret
	}).(VirtualMachineScaleSetStorageProfileResponseOutput)
}

func (o VirtualMachineScaleSetStorageProfileResponsePtrOutput) DataDisks() VirtualMachineScaleSetDataDiskResponseArrayOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetStorageProfileResponse) []VirtualMachineScaleSetDataDiskResponse {
		if v == nil {
			return nil
		}
		return v.DataDisks
	}).(VirtualMachineScaleSetDataDiskResponseArrayOutput)
}

func (o VirtualMachineScaleSetStorageProfileResponsePtrOutput) ImageReference() ImageReferenceResponsePtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetStorageProfileResponse) *ImageReferenceResponse {
		if v == nil {
			return nil
		}
		return v.ImageReference
	}).(ImageReferenceResponsePtrOutput)
}

func (o VirtualMachineScaleSetStorageProfileResponsePtrOutput) OsDisk() VirtualMachineScaleSetOSDiskResponsePtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetStorageProfileResponse) *VirtualMachineScaleSetOSDiskResponse {
		if v == nil {
			return nil
		}
		return v.OsDisk
	}).(VirtualMachineScaleSetOSDiskResponsePtrOutput)
}

type VirtualMachineScaleSetVMInstanceViewResponse struct {
	BootDiagnostics           *BootDiagnosticsInstanceViewResponse          `pulumi:"bootDiagnostics"`
	Disks                     []DiskInstanceViewResponse                    `pulumi:"disks"`
	Extensions                []VirtualMachineExtensionInstanceViewResponse `pulumi:"extensions"`
	MaintenanceRedeployStatus *MaintenanceRedeployStatusResponse            `pulumi:"maintenanceRedeployStatus"`
	PlacementGroupId          *string                                       `pulumi:"placementGroupId"`
	PlatformFaultDomain       *int                                          `pulumi:"platformFaultDomain"`
	PlatformUpdateDomain      *int                                          `pulumi:"platformUpdateDomain"`
	RdpThumbPrint             *string                                       `pulumi:"rdpThumbPrint"`
	Statuses                  []InstanceViewStatusResponse                  `pulumi:"statuses"`
	VmAgent                   *VirtualMachineAgentInstanceViewResponse      `pulumi:"vmAgent"`
	VmHealth                  VirtualMachineHealthStatusResponse            `pulumi:"vmHealth"`
}

type VirtualMachineScaleSetVMInstanceViewResponseOutput struct{ *pulumi.OutputState }

func (VirtualMachineScaleSetVMInstanceViewResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualMachineScaleSetVMInstanceViewResponse)(nil)).Elem()
}

func (o VirtualMachineScaleSetVMInstanceViewResponseOutput) ToVirtualMachineScaleSetVMInstanceViewResponseOutput() VirtualMachineScaleSetVMInstanceViewResponseOutput {
	return o
}

func (o VirtualMachineScaleSetVMInstanceViewResponseOutput) ToVirtualMachineScaleSetVMInstanceViewResponseOutputWithContext(ctx context.Context) VirtualMachineScaleSetVMInstanceViewResponseOutput {
	return o
}

func (o VirtualMachineScaleSetVMInstanceViewResponseOutput) BootDiagnostics() BootDiagnosticsInstanceViewResponsePtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetVMInstanceViewResponse) *BootDiagnosticsInstanceViewResponse {
		return v.BootDiagnostics
	}).(BootDiagnosticsInstanceViewResponsePtrOutput)
}

func (o VirtualMachineScaleSetVMInstanceViewResponseOutput) Disks() DiskInstanceViewResponseArrayOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetVMInstanceViewResponse) []DiskInstanceViewResponse { return v.Disks }).(DiskInstanceViewResponseArrayOutput)
}

func (o VirtualMachineScaleSetVMInstanceViewResponseOutput) Extensions() VirtualMachineExtensionInstanceViewResponseArrayOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetVMInstanceViewResponse) []VirtualMachineExtensionInstanceViewResponse {
		return v.Extensions
	}).(VirtualMachineExtensionInstanceViewResponseArrayOutput)
}

func (o VirtualMachineScaleSetVMInstanceViewResponseOutput) MaintenanceRedeployStatus() MaintenanceRedeployStatusResponsePtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetVMInstanceViewResponse) *MaintenanceRedeployStatusResponse {
		return v.MaintenanceRedeployStatus
	}).(MaintenanceRedeployStatusResponsePtrOutput)
}

func (o VirtualMachineScaleSetVMInstanceViewResponseOutput) PlacementGroupId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetVMInstanceViewResponse) *string { return v.PlacementGroupId }).(pulumi.StringPtrOutput)
}

func (o VirtualMachineScaleSetVMInstanceViewResponseOutput) PlatformFaultDomain() pulumi.IntPtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetVMInstanceViewResponse) *int { return v.PlatformFaultDomain }).(pulumi.IntPtrOutput)
}

func (o VirtualMachineScaleSetVMInstanceViewResponseOutput) PlatformUpdateDomain() pulumi.IntPtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetVMInstanceViewResponse) *int { return v.PlatformUpdateDomain }).(pulumi.IntPtrOutput)
}

func (o VirtualMachineScaleSetVMInstanceViewResponseOutput) RdpThumbPrint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetVMInstanceViewResponse) *string { return v.RdpThumbPrint }).(pulumi.StringPtrOutput)
}

func (o VirtualMachineScaleSetVMInstanceViewResponseOutput) Statuses() InstanceViewStatusResponseArrayOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetVMInstanceViewResponse) []InstanceViewStatusResponse { return v.Statuses }).(InstanceViewStatusResponseArrayOutput)
}

func (o VirtualMachineScaleSetVMInstanceViewResponseOutput) VmAgent() VirtualMachineAgentInstanceViewResponsePtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetVMInstanceViewResponse) *VirtualMachineAgentInstanceViewResponse {
		return v.VmAgent
	}).(VirtualMachineAgentInstanceViewResponsePtrOutput)
}

func (o VirtualMachineScaleSetVMInstanceViewResponseOutput) VmHealth() VirtualMachineHealthStatusResponseOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetVMInstanceViewResponse) VirtualMachineHealthStatusResponse {
		return v.VmHealth
	}).(VirtualMachineHealthStatusResponseOutput)
}

type VirtualMachineScaleSetVMProfile struct {
	DiagnosticsProfile *DiagnosticsProfile                     `pulumi:"diagnosticsProfile"`
	EvictionPolicy     *string                                 `pulumi:"evictionPolicy"`
	ExtensionProfile   *VirtualMachineScaleSetExtensionProfile `pulumi:"extensionProfile"`
	LicenseType        *string                                 `pulumi:"licenseType"`
	NetworkProfile     *VirtualMachineScaleSetNetworkProfile   `pulumi:"networkProfile"`
	OsProfile          *VirtualMachineScaleSetOSProfile        `pulumi:"osProfile"`
	Priority           *string                                 `pulumi:"priority"`
	StorageProfile     *VirtualMachineScaleSetStorageProfile   `pulumi:"storageProfile"`
}





type VirtualMachineScaleSetVMProfileInput interface {
	pulumi.Input

	ToVirtualMachineScaleSetVMProfileOutput() VirtualMachineScaleSetVMProfileOutput
	ToVirtualMachineScaleSetVMProfileOutputWithContext(context.Context) VirtualMachineScaleSetVMProfileOutput
}

type VirtualMachineScaleSetVMProfileArgs struct {
	DiagnosticsProfile DiagnosticsProfilePtrInput                     `pulumi:"diagnosticsProfile"`
	EvictionPolicy     pulumi.StringPtrInput                          `pulumi:"evictionPolicy"`
	ExtensionProfile   VirtualMachineScaleSetExtensionProfilePtrInput `pulumi:"extensionProfile"`
	LicenseType        pulumi.StringPtrInput                          `pulumi:"licenseType"`
	NetworkProfile     VirtualMachineScaleSetNetworkProfilePtrInput   `pulumi:"networkProfile"`
	OsProfile          VirtualMachineScaleSetOSProfilePtrInput        `pulumi:"osProfile"`
	Priority           pulumi.StringPtrInput                          `pulumi:"priority"`
	StorageProfile     VirtualMachineScaleSetStorageProfilePtrInput   `pulumi:"storageProfile"`
}

func (VirtualMachineScaleSetVMProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualMachineScaleSetVMProfile)(nil)).Elem()
}

func (i VirtualMachineScaleSetVMProfileArgs) ToVirtualMachineScaleSetVMProfileOutput() VirtualMachineScaleSetVMProfileOutput {
	return i.ToVirtualMachineScaleSetVMProfileOutputWithContext(context.Background())
}

func (i VirtualMachineScaleSetVMProfileArgs) ToVirtualMachineScaleSetVMProfileOutputWithContext(ctx context.Context) VirtualMachineScaleSetVMProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualMachineScaleSetVMProfileOutput)
}

func (i VirtualMachineScaleSetVMProfileArgs) ToVirtualMachineScaleSetVMProfilePtrOutput() VirtualMachineScaleSetVMProfilePtrOutput {
	return i.ToVirtualMachineScaleSetVMProfilePtrOutputWithContext(context.Background())
}

func (i VirtualMachineScaleSetVMProfileArgs) ToVirtualMachineScaleSetVMProfilePtrOutputWithContext(ctx context.Context) VirtualMachineScaleSetVMProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualMachineScaleSetVMProfileOutput).ToVirtualMachineScaleSetVMProfilePtrOutputWithContext(ctx)
}









type VirtualMachineScaleSetVMProfilePtrInput interface {
	pulumi.Input

	ToVirtualMachineScaleSetVMProfilePtrOutput() VirtualMachineScaleSetVMProfilePtrOutput
	ToVirtualMachineScaleSetVMProfilePtrOutputWithContext(context.Context) VirtualMachineScaleSetVMProfilePtrOutput
}

type virtualMachineScaleSetVMProfilePtrType VirtualMachineScaleSetVMProfileArgs

func VirtualMachineScaleSetVMProfilePtr(v *VirtualMachineScaleSetVMProfileArgs) VirtualMachineScaleSetVMProfilePtrInput {
	return (*virtualMachineScaleSetVMProfilePtrType)(v)
}

func (*virtualMachineScaleSetVMProfilePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualMachineScaleSetVMProfile)(nil)).Elem()
}

func (i *virtualMachineScaleSetVMProfilePtrType) ToVirtualMachineScaleSetVMProfilePtrOutput() VirtualMachineScaleSetVMProfilePtrOutput {
	return i.ToVirtualMachineScaleSetVMProfilePtrOutputWithContext(context.Background())
}

func (i *virtualMachineScaleSetVMProfilePtrType) ToVirtualMachineScaleSetVMProfilePtrOutputWithContext(ctx context.Context) VirtualMachineScaleSetVMProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualMachineScaleSetVMProfilePtrOutput)
}

type VirtualMachineScaleSetVMProfileOutput struct{ *pulumi.OutputState }

func (VirtualMachineScaleSetVMProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualMachineScaleSetVMProfile)(nil)).Elem()
}

func (o VirtualMachineScaleSetVMProfileOutput) ToVirtualMachineScaleSetVMProfileOutput() VirtualMachineScaleSetVMProfileOutput {
	return o
}

func (o VirtualMachineScaleSetVMProfileOutput) ToVirtualMachineScaleSetVMProfileOutputWithContext(ctx context.Context) VirtualMachineScaleSetVMProfileOutput {
	return o
}

func (o VirtualMachineScaleSetVMProfileOutput) ToVirtualMachineScaleSetVMProfilePtrOutput() VirtualMachineScaleSetVMProfilePtrOutput {
	return o.ToVirtualMachineScaleSetVMProfilePtrOutputWithContext(context.Background())
}

func (o VirtualMachineScaleSetVMProfileOutput) ToVirtualMachineScaleSetVMProfilePtrOutputWithContext(ctx context.Context) VirtualMachineScaleSetVMProfilePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v VirtualMachineScaleSetVMProfile) *VirtualMachineScaleSetVMProfile {
		return &v
	}).(VirtualMachineScaleSetVMProfilePtrOutput)
}

func (o VirtualMachineScaleSetVMProfileOutput) DiagnosticsProfile() DiagnosticsProfilePtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetVMProfile) *DiagnosticsProfile { return v.DiagnosticsProfile }).(DiagnosticsProfilePtrOutput)
}

func (o VirtualMachineScaleSetVMProfileOutput) EvictionPolicy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetVMProfile) *string { return v.EvictionPolicy }).(pulumi.StringPtrOutput)
}

func (o VirtualMachineScaleSetVMProfileOutput) ExtensionProfile() VirtualMachineScaleSetExtensionProfilePtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetVMProfile) *VirtualMachineScaleSetExtensionProfile {
		return v.ExtensionProfile
	}).(VirtualMachineScaleSetExtensionProfilePtrOutput)
}

func (o VirtualMachineScaleSetVMProfileOutput) LicenseType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetVMProfile) *string { return v.LicenseType }).(pulumi.StringPtrOutput)
}

func (o VirtualMachineScaleSetVMProfileOutput) NetworkProfile() VirtualMachineScaleSetNetworkProfilePtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetVMProfile) *VirtualMachineScaleSetNetworkProfile { return v.NetworkProfile }).(VirtualMachineScaleSetNetworkProfilePtrOutput)
}

func (o VirtualMachineScaleSetVMProfileOutput) OsProfile() VirtualMachineScaleSetOSProfilePtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetVMProfile) *VirtualMachineScaleSetOSProfile { return v.OsProfile }).(VirtualMachineScaleSetOSProfilePtrOutput)
}

func (o VirtualMachineScaleSetVMProfileOutput) Priority() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetVMProfile) *string { return v.Priority }).(pulumi.StringPtrOutput)
}

func (o VirtualMachineScaleSetVMProfileOutput) StorageProfile() VirtualMachineScaleSetStorageProfilePtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetVMProfile) *VirtualMachineScaleSetStorageProfile { return v.StorageProfile }).(VirtualMachineScaleSetStorageProfilePtrOutput)
}

type VirtualMachineScaleSetVMProfilePtrOutput struct{ *pulumi.OutputState }

func (VirtualMachineScaleSetVMProfilePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualMachineScaleSetVMProfile)(nil)).Elem()
}

func (o VirtualMachineScaleSetVMProfilePtrOutput) ToVirtualMachineScaleSetVMProfilePtrOutput() VirtualMachineScaleSetVMProfilePtrOutput {
	return o
}

func (o VirtualMachineScaleSetVMProfilePtrOutput) ToVirtualMachineScaleSetVMProfilePtrOutputWithContext(ctx context.Context) VirtualMachineScaleSetVMProfilePtrOutput {
	return o
}

func (o VirtualMachineScaleSetVMProfilePtrOutput) Elem() VirtualMachineScaleSetVMProfileOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetVMProfile) VirtualMachineScaleSetVMProfile {
		if v != nil {
			return *v
		}
		var ret VirtualMachineScaleSetVMProfile
		return ret
	}).(VirtualMachineScaleSetVMProfileOutput)
}

func (o VirtualMachineScaleSetVMProfilePtrOutput) DiagnosticsProfile() DiagnosticsProfilePtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetVMProfile) *DiagnosticsProfile {
		if v == nil {
			return nil
		}
		return v.DiagnosticsProfile
	}).(DiagnosticsProfilePtrOutput)
}

func (o VirtualMachineScaleSetVMProfilePtrOutput) EvictionPolicy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetVMProfile) *string {
		if v == nil {
			return nil
		}
		return v.EvictionPolicy
	}).(pulumi.StringPtrOutput)
}

func (o VirtualMachineScaleSetVMProfilePtrOutput) ExtensionProfile() VirtualMachineScaleSetExtensionProfilePtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetVMProfile) *VirtualMachineScaleSetExtensionProfile {
		if v == nil {
			return nil
		}
		return v.ExtensionProfile
	}).(VirtualMachineScaleSetExtensionProfilePtrOutput)
}

func (o VirtualMachineScaleSetVMProfilePtrOutput) LicenseType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetVMProfile) *string {
		if v == nil {
			return nil
		}
		return v.LicenseType
	}).(pulumi.StringPtrOutput)
}

func (o VirtualMachineScaleSetVMProfilePtrOutput) NetworkProfile() VirtualMachineScaleSetNetworkProfilePtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetVMProfile) *VirtualMachineScaleSetNetworkProfile {
		if v == nil {
			return nil
		}
		return v.NetworkProfile
	}).(VirtualMachineScaleSetNetworkProfilePtrOutput)
}

func (o VirtualMachineScaleSetVMProfilePtrOutput) OsProfile() VirtualMachineScaleSetOSProfilePtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetVMProfile) *VirtualMachineScaleSetOSProfile {
		if v == nil {
			return nil
		}
		return v.OsProfile
	}).(VirtualMachineScaleSetOSProfilePtrOutput)
}

func (o VirtualMachineScaleSetVMProfilePtrOutput) Priority() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetVMProfile) *string {
		if v == nil {
			return nil
		}
		return v.Priority
	}).(pulumi.StringPtrOutput)
}

func (o VirtualMachineScaleSetVMProfilePtrOutput) StorageProfile() VirtualMachineScaleSetStorageProfilePtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetVMProfile) *VirtualMachineScaleSetStorageProfile {
		if v == nil {
			return nil
		}
		return v.StorageProfile
	}).(VirtualMachineScaleSetStorageProfilePtrOutput)
}

type VirtualMachineScaleSetVMProfileResponse struct {
	DiagnosticsProfile *DiagnosticsProfileResponse                     `pulumi:"diagnosticsProfile"`
	EvictionPolicy     *string                                         `pulumi:"evictionPolicy"`
	ExtensionProfile   *VirtualMachineScaleSetExtensionProfileResponse `pulumi:"extensionProfile"`
	LicenseType        *string                                         `pulumi:"licenseType"`
	NetworkProfile     *VirtualMachineScaleSetNetworkProfileResponse   `pulumi:"networkProfile"`
	OsProfile          *VirtualMachineScaleSetOSProfileResponse        `pulumi:"osProfile"`
	Priority           *string                                         `pulumi:"priority"`
	StorageProfile     *VirtualMachineScaleSetStorageProfileResponse   `pulumi:"storageProfile"`
}

type VirtualMachineScaleSetVMProfileResponseOutput struct{ *pulumi.OutputState }

func (VirtualMachineScaleSetVMProfileResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualMachineScaleSetVMProfileResponse)(nil)).Elem()
}

func (o VirtualMachineScaleSetVMProfileResponseOutput) ToVirtualMachineScaleSetVMProfileResponseOutput() VirtualMachineScaleSetVMProfileResponseOutput {
	return o
}

func (o VirtualMachineScaleSetVMProfileResponseOutput) ToVirtualMachineScaleSetVMProfileResponseOutputWithContext(ctx context.Context) VirtualMachineScaleSetVMProfileResponseOutput {
	return o
}

func (o VirtualMachineScaleSetVMProfileResponseOutput) DiagnosticsProfile() DiagnosticsProfileResponsePtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetVMProfileResponse) *DiagnosticsProfileResponse {
		return v.DiagnosticsProfile
	}).(DiagnosticsProfileResponsePtrOutput)
}

func (o VirtualMachineScaleSetVMProfileResponseOutput) EvictionPolicy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetVMProfileResponse) *string { return v.EvictionPolicy }).(pulumi.StringPtrOutput)
}

func (o VirtualMachineScaleSetVMProfileResponseOutput) ExtensionProfile() VirtualMachineScaleSetExtensionProfileResponsePtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetVMProfileResponse) *VirtualMachineScaleSetExtensionProfileResponse {
		return v.ExtensionProfile
	}).(VirtualMachineScaleSetExtensionProfileResponsePtrOutput)
}

func (o VirtualMachineScaleSetVMProfileResponseOutput) LicenseType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetVMProfileResponse) *string { return v.LicenseType }).(pulumi.StringPtrOutput)
}

func (o VirtualMachineScaleSetVMProfileResponseOutput) NetworkProfile() VirtualMachineScaleSetNetworkProfileResponsePtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetVMProfileResponse) *VirtualMachineScaleSetNetworkProfileResponse {
		return v.NetworkProfile
	}).(VirtualMachineScaleSetNetworkProfileResponsePtrOutput)
}

func (o VirtualMachineScaleSetVMProfileResponseOutput) OsProfile() VirtualMachineScaleSetOSProfileResponsePtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetVMProfileResponse) *VirtualMachineScaleSetOSProfileResponse {
		return v.OsProfile
	}).(VirtualMachineScaleSetOSProfileResponsePtrOutput)
}

func (o VirtualMachineScaleSetVMProfileResponseOutput) Priority() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetVMProfileResponse) *string { return v.Priority }).(pulumi.StringPtrOutput)
}

func (o VirtualMachineScaleSetVMProfileResponseOutput) StorageProfile() VirtualMachineScaleSetStorageProfileResponsePtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetVMProfileResponse) *VirtualMachineScaleSetStorageProfileResponse {
		return v.StorageProfile
	}).(VirtualMachineScaleSetStorageProfileResponsePtrOutput)
}

type VirtualMachineScaleSetVMProfileResponsePtrOutput struct{ *pulumi.OutputState }

func (VirtualMachineScaleSetVMProfileResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualMachineScaleSetVMProfileResponse)(nil)).Elem()
}

func (o VirtualMachineScaleSetVMProfileResponsePtrOutput) ToVirtualMachineScaleSetVMProfileResponsePtrOutput() VirtualMachineScaleSetVMProfileResponsePtrOutput {
	return o
}

func (o VirtualMachineScaleSetVMProfileResponsePtrOutput) ToVirtualMachineScaleSetVMProfileResponsePtrOutputWithContext(ctx context.Context) VirtualMachineScaleSetVMProfileResponsePtrOutput {
	return o
}

func (o VirtualMachineScaleSetVMProfileResponsePtrOutput) Elem() VirtualMachineScaleSetVMProfileResponseOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetVMProfileResponse) VirtualMachineScaleSetVMProfileResponse {
		if v != nil {
			return *v
		}
		var ret VirtualMachineScaleSetVMProfileResponse
		return ret
	}).(VirtualMachineScaleSetVMProfileResponseOutput)
}

func (o VirtualMachineScaleSetVMProfileResponsePtrOutput) DiagnosticsProfile() DiagnosticsProfileResponsePtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetVMProfileResponse) *DiagnosticsProfileResponse {
		if v == nil {
			return nil
		}
		return v.DiagnosticsProfile
	}).(DiagnosticsProfileResponsePtrOutput)
}

func (o VirtualMachineScaleSetVMProfileResponsePtrOutput) EvictionPolicy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetVMProfileResponse) *string {
		if v == nil {
			return nil
		}
		return v.EvictionPolicy
	}).(pulumi.StringPtrOutput)
}

func (o VirtualMachineScaleSetVMProfileResponsePtrOutput) ExtensionProfile() VirtualMachineScaleSetExtensionProfileResponsePtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetVMProfileResponse) *VirtualMachineScaleSetExtensionProfileResponse {
		if v == nil {
			return nil
		}
		return v.ExtensionProfile
	}).(VirtualMachineScaleSetExtensionProfileResponsePtrOutput)
}

func (o VirtualMachineScaleSetVMProfileResponsePtrOutput) LicenseType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetVMProfileResponse) *string {
		if v == nil {
			return nil
		}
		return v.LicenseType
	}).(pulumi.StringPtrOutput)
}

func (o VirtualMachineScaleSetVMProfileResponsePtrOutput) NetworkProfile() VirtualMachineScaleSetNetworkProfileResponsePtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetVMProfileResponse) *VirtualMachineScaleSetNetworkProfileResponse {
		if v == nil {
			return nil
		}
		return v.NetworkProfile
	}).(VirtualMachineScaleSetNetworkProfileResponsePtrOutput)
}

func (o VirtualMachineScaleSetVMProfileResponsePtrOutput) OsProfile() VirtualMachineScaleSetOSProfileResponsePtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetVMProfileResponse) *VirtualMachineScaleSetOSProfileResponse {
		if v == nil {
			return nil
		}
		return v.OsProfile
	}).(VirtualMachineScaleSetOSProfileResponsePtrOutput)
}

func (o VirtualMachineScaleSetVMProfileResponsePtrOutput) Priority() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetVMProfileResponse) *string {
		if v == nil {
			return nil
		}
		return v.Priority
	}).(pulumi.StringPtrOutput)
}

func (o VirtualMachineScaleSetVMProfileResponsePtrOutput) StorageProfile() VirtualMachineScaleSetStorageProfileResponsePtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetVMProfileResponse) *VirtualMachineScaleSetStorageProfileResponse {
		if v == nil {
			return nil
		}
		return v.StorageProfile
	}).(VirtualMachineScaleSetStorageProfileResponsePtrOutput)
}

type WinRMConfiguration struct {
	Listeners []WinRMListener `pulumi:"listeners"`
}





type WinRMConfigurationInput interface {
	pulumi.Input

	ToWinRMConfigurationOutput() WinRMConfigurationOutput
	ToWinRMConfigurationOutputWithContext(context.Context) WinRMConfigurationOutput
}

type WinRMConfigurationArgs struct {
	Listeners WinRMListenerArrayInput `pulumi:"listeners"`
}

func (WinRMConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*WinRMConfiguration)(nil)).Elem()
}

func (i WinRMConfigurationArgs) ToWinRMConfigurationOutput() WinRMConfigurationOutput {
	return i.ToWinRMConfigurationOutputWithContext(context.Background())
}

func (i WinRMConfigurationArgs) ToWinRMConfigurationOutputWithContext(ctx context.Context) WinRMConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WinRMConfigurationOutput)
}

func (i WinRMConfigurationArgs) ToWinRMConfigurationPtrOutput() WinRMConfigurationPtrOutput {
	return i.ToWinRMConfigurationPtrOutputWithContext(context.Background())
}

func (i WinRMConfigurationArgs) ToWinRMConfigurationPtrOutputWithContext(ctx context.Context) WinRMConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WinRMConfigurationOutput).ToWinRMConfigurationPtrOutputWithContext(ctx)
}









type WinRMConfigurationPtrInput interface {
	pulumi.Input

	ToWinRMConfigurationPtrOutput() WinRMConfigurationPtrOutput
	ToWinRMConfigurationPtrOutputWithContext(context.Context) WinRMConfigurationPtrOutput
}

type winRMConfigurationPtrType WinRMConfigurationArgs

func WinRMConfigurationPtr(v *WinRMConfigurationArgs) WinRMConfigurationPtrInput {
	return (*winRMConfigurationPtrType)(v)
}

func (*winRMConfigurationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**WinRMConfiguration)(nil)).Elem()
}

func (i *winRMConfigurationPtrType) ToWinRMConfigurationPtrOutput() WinRMConfigurationPtrOutput {
	return i.ToWinRMConfigurationPtrOutputWithContext(context.Background())
}

func (i *winRMConfigurationPtrType) ToWinRMConfigurationPtrOutputWithContext(ctx context.Context) WinRMConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WinRMConfigurationPtrOutput)
}

type WinRMConfigurationOutput struct{ *pulumi.OutputState }

func (WinRMConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WinRMConfiguration)(nil)).Elem()
}

func (o WinRMConfigurationOutput) ToWinRMConfigurationOutput() WinRMConfigurationOutput {
	return o
}

func (o WinRMConfigurationOutput) ToWinRMConfigurationOutputWithContext(ctx context.Context) WinRMConfigurationOutput {
	return o
}

func (o WinRMConfigurationOutput) ToWinRMConfigurationPtrOutput() WinRMConfigurationPtrOutput {
	return o.ToWinRMConfigurationPtrOutputWithContext(context.Background())
}

func (o WinRMConfigurationOutput) ToWinRMConfigurationPtrOutputWithContext(ctx context.Context) WinRMConfigurationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v WinRMConfiguration) *WinRMConfiguration {
		return &v
	}).(WinRMConfigurationPtrOutput)
}

func (o WinRMConfigurationOutput) Listeners() WinRMListenerArrayOutput {
	return o.ApplyT(func(v WinRMConfiguration) []WinRMListener { return v.Listeners }).(WinRMListenerArrayOutput)
}

type WinRMConfigurationPtrOutput struct{ *pulumi.OutputState }

func (WinRMConfigurationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WinRMConfiguration)(nil)).Elem()
}

func (o WinRMConfigurationPtrOutput) ToWinRMConfigurationPtrOutput() WinRMConfigurationPtrOutput {
	return o
}

func (o WinRMConfigurationPtrOutput) ToWinRMConfigurationPtrOutputWithContext(ctx context.Context) WinRMConfigurationPtrOutput {
	return o
}

func (o WinRMConfigurationPtrOutput) Elem() WinRMConfigurationOutput {
	return o.ApplyT(func(v *WinRMConfiguration) WinRMConfiguration {
		if v != nil {
			return *v
		}
		var ret WinRMConfiguration
		return ret
	}).(WinRMConfigurationOutput)
}

func (o WinRMConfigurationPtrOutput) Listeners() WinRMListenerArrayOutput {
	return o.ApplyT(func(v *WinRMConfiguration) []WinRMListener {
		if v == nil {
			return nil
		}
		return v.Listeners
	}).(WinRMListenerArrayOutput)
}

type WinRMConfigurationResponse struct {
	Listeners []WinRMListenerResponse `pulumi:"listeners"`
}

type WinRMConfigurationResponseOutput struct{ *pulumi.OutputState }

func (WinRMConfigurationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WinRMConfigurationResponse)(nil)).Elem()
}

func (o WinRMConfigurationResponseOutput) ToWinRMConfigurationResponseOutput() WinRMConfigurationResponseOutput {
	return o
}

func (o WinRMConfigurationResponseOutput) ToWinRMConfigurationResponseOutputWithContext(ctx context.Context) WinRMConfigurationResponseOutput {
	return o
}

func (o WinRMConfigurationResponseOutput) Listeners() WinRMListenerResponseArrayOutput {
	return o.ApplyT(func(v WinRMConfigurationResponse) []WinRMListenerResponse { return v.Listeners }).(WinRMListenerResponseArrayOutput)
}

type WinRMConfigurationResponsePtrOutput struct{ *pulumi.OutputState }

func (WinRMConfigurationResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WinRMConfigurationResponse)(nil)).Elem()
}

func (o WinRMConfigurationResponsePtrOutput) ToWinRMConfigurationResponsePtrOutput() WinRMConfigurationResponsePtrOutput {
	return o
}

func (o WinRMConfigurationResponsePtrOutput) ToWinRMConfigurationResponsePtrOutputWithContext(ctx context.Context) WinRMConfigurationResponsePtrOutput {
	return o
}

func (o WinRMConfigurationResponsePtrOutput) Elem() WinRMConfigurationResponseOutput {
	return o.ApplyT(func(v *WinRMConfigurationResponse) WinRMConfigurationResponse {
		if v != nil {
			return *v
		}
		var ret WinRMConfigurationResponse
		return ret
	}).(WinRMConfigurationResponseOutput)
}

func (o WinRMConfigurationResponsePtrOutput) Listeners() WinRMListenerResponseArrayOutput {
	return o.ApplyT(func(v *WinRMConfigurationResponse) []WinRMListenerResponse {
		if v == nil {
			return nil
		}
		return v.Listeners
	}).(WinRMListenerResponseArrayOutput)
}

type WinRMListener struct {
	CertificateUrl *string        `pulumi:"certificateUrl"`
	Protocol       *ProtocolTypes `pulumi:"protocol"`
}





type WinRMListenerInput interface {
	pulumi.Input

	ToWinRMListenerOutput() WinRMListenerOutput
	ToWinRMListenerOutputWithContext(context.Context) WinRMListenerOutput
}

type WinRMListenerArgs struct {
	CertificateUrl pulumi.StringPtrInput `pulumi:"certificateUrl"`
	Protocol       ProtocolTypesPtrInput `pulumi:"protocol"`
}

func (WinRMListenerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*WinRMListener)(nil)).Elem()
}

func (i WinRMListenerArgs) ToWinRMListenerOutput() WinRMListenerOutput {
	return i.ToWinRMListenerOutputWithContext(context.Background())
}

func (i WinRMListenerArgs) ToWinRMListenerOutputWithContext(ctx context.Context) WinRMListenerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WinRMListenerOutput)
}





type WinRMListenerArrayInput interface {
	pulumi.Input

	ToWinRMListenerArrayOutput() WinRMListenerArrayOutput
	ToWinRMListenerArrayOutputWithContext(context.Context) WinRMListenerArrayOutput
}

type WinRMListenerArray []WinRMListenerInput

func (WinRMListenerArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]WinRMListener)(nil)).Elem()
}

func (i WinRMListenerArray) ToWinRMListenerArrayOutput() WinRMListenerArrayOutput {
	return i.ToWinRMListenerArrayOutputWithContext(context.Background())
}

func (i WinRMListenerArray) ToWinRMListenerArrayOutputWithContext(ctx context.Context) WinRMListenerArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WinRMListenerArrayOutput)
}

type WinRMListenerOutput struct{ *pulumi.OutputState }

func (WinRMListenerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WinRMListener)(nil)).Elem()
}

func (o WinRMListenerOutput) ToWinRMListenerOutput() WinRMListenerOutput {
	return o
}

func (o WinRMListenerOutput) ToWinRMListenerOutputWithContext(ctx context.Context) WinRMListenerOutput {
	return o
}

func (o WinRMListenerOutput) CertificateUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WinRMListener) *string { return v.CertificateUrl }).(pulumi.StringPtrOutput)
}

func (o WinRMListenerOutput) Protocol() ProtocolTypesPtrOutput {
	return o.ApplyT(func(v WinRMListener) *ProtocolTypes { return v.Protocol }).(ProtocolTypesPtrOutput)
}

type WinRMListenerArrayOutput struct{ *pulumi.OutputState }

func (WinRMListenerArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]WinRMListener)(nil)).Elem()
}

func (o WinRMListenerArrayOutput) ToWinRMListenerArrayOutput() WinRMListenerArrayOutput {
	return o
}

func (o WinRMListenerArrayOutput) ToWinRMListenerArrayOutputWithContext(ctx context.Context) WinRMListenerArrayOutput {
	return o
}

func (o WinRMListenerArrayOutput) Index(i pulumi.IntInput) WinRMListenerOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) WinRMListener {
		return vs[0].([]WinRMListener)[vs[1].(int)]
	}).(WinRMListenerOutput)
}

type WinRMListenerResponse struct {
	CertificateUrl *string `pulumi:"certificateUrl"`
	Protocol       *string `pulumi:"protocol"`
}

type WinRMListenerResponseOutput struct{ *pulumi.OutputState }

func (WinRMListenerResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WinRMListenerResponse)(nil)).Elem()
}

func (o WinRMListenerResponseOutput) ToWinRMListenerResponseOutput() WinRMListenerResponseOutput {
	return o
}

func (o WinRMListenerResponseOutput) ToWinRMListenerResponseOutputWithContext(ctx context.Context) WinRMListenerResponseOutput {
	return o
}

func (o WinRMListenerResponseOutput) CertificateUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WinRMListenerResponse) *string { return v.CertificateUrl }).(pulumi.StringPtrOutput)
}

func (o WinRMListenerResponseOutput) Protocol() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WinRMListenerResponse) *string { return v.Protocol }).(pulumi.StringPtrOutput)
}

type WinRMListenerResponseArrayOutput struct{ *pulumi.OutputState }

func (WinRMListenerResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]WinRMListenerResponse)(nil)).Elem()
}

func (o WinRMListenerResponseArrayOutput) ToWinRMListenerResponseArrayOutput() WinRMListenerResponseArrayOutput {
	return o
}

func (o WinRMListenerResponseArrayOutput) ToWinRMListenerResponseArrayOutputWithContext(ctx context.Context) WinRMListenerResponseArrayOutput {
	return o
}

func (o WinRMListenerResponseArrayOutput) Index(i pulumi.IntInput) WinRMListenerResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) WinRMListenerResponse {
		return vs[0].([]WinRMListenerResponse)[vs[1].(int)]
	}).(WinRMListenerResponseOutput)
}

type WindowsConfiguration struct {
	AdditionalUnattendContent []AdditionalUnattendContent `pulumi:"additionalUnattendContent"`
	EnableAutomaticUpdates    *bool                       `pulumi:"enableAutomaticUpdates"`
	ProvisionVMAgent          *bool                       `pulumi:"provisionVMAgent"`
	TimeZone                  *string                     `pulumi:"timeZone"`
	WinRM                     *WinRMConfiguration         `pulumi:"winRM"`
}





type WindowsConfigurationInput interface {
	pulumi.Input

	ToWindowsConfigurationOutput() WindowsConfigurationOutput
	ToWindowsConfigurationOutputWithContext(context.Context) WindowsConfigurationOutput
}

type WindowsConfigurationArgs struct {
	AdditionalUnattendContent AdditionalUnattendContentArrayInput `pulumi:"additionalUnattendContent"`
	EnableAutomaticUpdates    pulumi.BoolPtrInput                 `pulumi:"enableAutomaticUpdates"`
	ProvisionVMAgent          pulumi.BoolPtrInput                 `pulumi:"provisionVMAgent"`
	TimeZone                  pulumi.StringPtrInput               `pulumi:"timeZone"`
	WinRM                     WinRMConfigurationPtrInput          `pulumi:"winRM"`
}

func (WindowsConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*WindowsConfiguration)(nil)).Elem()
}

func (i WindowsConfigurationArgs) ToWindowsConfigurationOutput() WindowsConfigurationOutput {
	return i.ToWindowsConfigurationOutputWithContext(context.Background())
}

func (i WindowsConfigurationArgs) ToWindowsConfigurationOutputWithContext(ctx context.Context) WindowsConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WindowsConfigurationOutput)
}

func (i WindowsConfigurationArgs) ToWindowsConfigurationPtrOutput() WindowsConfigurationPtrOutput {
	return i.ToWindowsConfigurationPtrOutputWithContext(context.Background())
}

func (i WindowsConfigurationArgs) ToWindowsConfigurationPtrOutputWithContext(ctx context.Context) WindowsConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WindowsConfigurationOutput).ToWindowsConfigurationPtrOutputWithContext(ctx)
}









type WindowsConfigurationPtrInput interface {
	pulumi.Input

	ToWindowsConfigurationPtrOutput() WindowsConfigurationPtrOutput
	ToWindowsConfigurationPtrOutputWithContext(context.Context) WindowsConfigurationPtrOutput
}

type windowsConfigurationPtrType WindowsConfigurationArgs

func WindowsConfigurationPtr(v *WindowsConfigurationArgs) WindowsConfigurationPtrInput {
	return (*windowsConfigurationPtrType)(v)
}

func (*windowsConfigurationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**WindowsConfiguration)(nil)).Elem()
}

func (i *windowsConfigurationPtrType) ToWindowsConfigurationPtrOutput() WindowsConfigurationPtrOutput {
	return i.ToWindowsConfigurationPtrOutputWithContext(context.Background())
}

func (i *windowsConfigurationPtrType) ToWindowsConfigurationPtrOutputWithContext(ctx context.Context) WindowsConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WindowsConfigurationPtrOutput)
}

type WindowsConfigurationOutput struct{ *pulumi.OutputState }

func (WindowsConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WindowsConfiguration)(nil)).Elem()
}

func (o WindowsConfigurationOutput) ToWindowsConfigurationOutput() WindowsConfigurationOutput {
	return o
}

func (o WindowsConfigurationOutput) ToWindowsConfigurationOutputWithContext(ctx context.Context) WindowsConfigurationOutput {
	return o
}

func (o WindowsConfigurationOutput) ToWindowsConfigurationPtrOutput() WindowsConfigurationPtrOutput {
	return o.ToWindowsConfigurationPtrOutputWithContext(context.Background())
}

func (o WindowsConfigurationOutput) ToWindowsConfigurationPtrOutputWithContext(ctx context.Context) WindowsConfigurationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v WindowsConfiguration) *WindowsConfiguration {
		return &v
	}).(WindowsConfigurationPtrOutput)
}

func (o WindowsConfigurationOutput) AdditionalUnattendContent() AdditionalUnattendContentArrayOutput {
	return o.ApplyT(func(v WindowsConfiguration) []AdditionalUnattendContent { return v.AdditionalUnattendContent }).(AdditionalUnattendContentArrayOutput)
}

func (o WindowsConfigurationOutput) EnableAutomaticUpdates() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v WindowsConfiguration) *bool { return v.EnableAutomaticUpdates }).(pulumi.BoolPtrOutput)
}

func (o WindowsConfigurationOutput) ProvisionVMAgent() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v WindowsConfiguration) *bool { return v.ProvisionVMAgent }).(pulumi.BoolPtrOutput)
}

func (o WindowsConfigurationOutput) TimeZone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WindowsConfiguration) *string { return v.TimeZone }).(pulumi.StringPtrOutput)
}

func (o WindowsConfigurationOutput) WinRM() WinRMConfigurationPtrOutput {
	return o.ApplyT(func(v WindowsConfiguration) *WinRMConfiguration { return v.WinRM }).(WinRMConfigurationPtrOutput)
}

type WindowsConfigurationPtrOutput struct{ *pulumi.OutputState }

func (WindowsConfigurationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WindowsConfiguration)(nil)).Elem()
}

func (o WindowsConfigurationPtrOutput) ToWindowsConfigurationPtrOutput() WindowsConfigurationPtrOutput {
	return o
}

func (o WindowsConfigurationPtrOutput) ToWindowsConfigurationPtrOutputWithContext(ctx context.Context) WindowsConfigurationPtrOutput {
	return o
}

func (o WindowsConfigurationPtrOutput) Elem() WindowsConfigurationOutput {
	return o.ApplyT(func(v *WindowsConfiguration) WindowsConfiguration {
		if v != nil {
			return *v
		}
		var ret WindowsConfiguration
		return ret
	}).(WindowsConfigurationOutput)
}

func (o WindowsConfigurationPtrOutput) AdditionalUnattendContent() AdditionalUnattendContentArrayOutput {
	return o.ApplyT(func(v *WindowsConfiguration) []AdditionalUnattendContent {
		if v == nil {
			return nil
		}
		return v.AdditionalUnattendContent
	}).(AdditionalUnattendContentArrayOutput)
}

func (o WindowsConfigurationPtrOutput) EnableAutomaticUpdates() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *WindowsConfiguration) *bool {
		if v == nil {
			return nil
		}
		return v.EnableAutomaticUpdates
	}).(pulumi.BoolPtrOutput)
}

func (o WindowsConfigurationPtrOutput) ProvisionVMAgent() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *WindowsConfiguration) *bool {
		if v == nil {
			return nil
		}
		return v.ProvisionVMAgent
	}).(pulumi.BoolPtrOutput)
}

func (o WindowsConfigurationPtrOutput) TimeZone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WindowsConfiguration) *string {
		if v == nil {
			return nil
		}
		return v.TimeZone
	}).(pulumi.StringPtrOutput)
}

func (o WindowsConfigurationPtrOutput) WinRM() WinRMConfigurationPtrOutput {
	return o.ApplyT(func(v *WindowsConfiguration) *WinRMConfiguration {
		if v == nil {
			return nil
		}
		return v.WinRM
	}).(WinRMConfigurationPtrOutput)
}

type WindowsConfigurationResponse struct {
	AdditionalUnattendContent []AdditionalUnattendContentResponse `pulumi:"additionalUnattendContent"`
	EnableAutomaticUpdates    *bool                               `pulumi:"enableAutomaticUpdates"`
	ProvisionVMAgent          *bool                               `pulumi:"provisionVMAgent"`
	TimeZone                  *string                             `pulumi:"timeZone"`
	WinRM                     *WinRMConfigurationResponse         `pulumi:"winRM"`
}

type WindowsConfigurationResponseOutput struct{ *pulumi.OutputState }

func (WindowsConfigurationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WindowsConfigurationResponse)(nil)).Elem()
}

func (o WindowsConfigurationResponseOutput) ToWindowsConfigurationResponseOutput() WindowsConfigurationResponseOutput {
	return o
}

func (o WindowsConfigurationResponseOutput) ToWindowsConfigurationResponseOutputWithContext(ctx context.Context) WindowsConfigurationResponseOutput {
	return o
}

func (o WindowsConfigurationResponseOutput) AdditionalUnattendContent() AdditionalUnattendContentResponseArrayOutput {
	return o.ApplyT(func(v WindowsConfigurationResponse) []AdditionalUnattendContentResponse {
		return v.AdditionalUnattendContent
	}).(AdditionalUnattendContentResponseArrayOutput)
}

func (o WindowsConfigurationResponseOutput) EnableAutomaticUpdates() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v WindowsConfigurationResponse) *bool { return v.EnableAutomaticUpdates }).(pulumi.BoolPtrOutput)
}

func (o WindowsConfigurationResponseOutput) ProvisionVMAgent() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v WindowsConfigurationResponse) *bool { return v.ProvisionVMAgent }).(pulumi.BoolPtrOutput)
}

func (o WindowsConfigurationResponseOutput) TimeZone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WindowsConfigurationResponse) *string { return v.TimeZone }).(pulumi.StringPtrOutput)
}

func (o WindowsConfigurationResponseOutput) WinRM() WinRMConfigurationResponsePtrOutput {
	return o.ApplyT(func(v WindowsConfigurationResponse) *WinRMConfigurationResponse { return v.WinRM }).(WinRMConfigurationResponsePtrOutput)
}

type WindowsConfigurationResponsePtrOutput struct{ *pulumi.OutputState }

func (WindowsConfigurationResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WindowsConfigurationResponse)(nil)).Elem()
}

func (o WindowsConfigurationResponsePtrOutput) ToWindowsConfigurationResponsePtrOutput() WindowsConfigurationResponsePtrOutput {
	return o
}

func (o WindowsConfigurationResponsePtrOutput) ToWindowsConfigurationResponsePtrOutputWithContext(ctx context.Context) WindowsConfigurationResponsePtrOutput {
	return o
}

func (o WindowsConfigurationResponsePtrOutput) Elem() WindowsConfigurationResponseOutput {
	return o.ApplyT(func(v *WindowsConfigurationResponse) WindowsConfigurationResponse {
		if v != nil {
			return *v
		}
		var ret WindowsConfigurationResponse
		return ret
	}).(WindowsConfigurationResponseOutput)
}

func (o WindowsConfigurationResponsePtrOutput) AdditionalUnattendContent() AdditionalUnattendContentResponseArrayOutput {
	return o.ApplyT(func(v *WindowsConfigurationResponse) []AdditionalUnattendContentResponse {
		if v == nil {
			return nil
		}
		return v.AdditionalUnattendContent
	}).(AdditionalUnattendContentResponseArrayOutput)
}

func (o WindowsConfigurationResponsePtrOutput) EnableAutomaticUpdates() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *WindowsConfigurationResponse) *bool {
		if v == nil {
			return nil
		}
		return v.EnableAutomaticUpdates
	}).(pulumi.BoolPtrOutput)
}

func (o WindowsConfigurationResponsePtrOutput) ProvisionVMAgent() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *WindowsConfigurationResponse) *bool {
		if v == nil {
			return nil
		}
		return v.ProvisionVMAgent
	}).(pulumi.BoolPtrOutput)
}

func (o WindowsConfigurationResponsePtrOutput) TimeZone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WindowsConfigurationResponse) *string {
		if v == nil {
			return nil
		}
		return v.TimeZone
	}).(pulumi.StringPtrOutput)
}

func (o WindowsConfigurationResponsePtrOutput) WinRM() WinRMConfigurationResponsePtrOutput {
	return o.ApplyT(func(v *WindowsConfigurationResponse) *WinRMConfigurationResponse {
		if v == nil {
			return nil
		}
		return v.WinRM
	}).(WinRMConfigurationResponsePtrOutput)
}

func init() {
	pulumi.RegisterOutputType(AdditionalUnattendContentOutput{})
	pulumi.RegisterOutputType(AdditionalUnattendContentArrayOutput{})
	pulumi.RegisterOutputType(AdditionalUnattendContentResponseOutput{})
	pulumi.RegisterOutputType(AdditionalUnattendContentResponseArrayOutput{})
	pulumi.RegisterOutputType(ApiEntityReferenceOutput{})
	pulumi.RegisterOutputType(ApiEntityReferencePtrOutput{})
	pulumi.RegisterOutputType(ApiEntityReferenceResponseOutput{})
	pulumi.RegisterOutputType(ApiEntityReferenceResponsePtrOutput{})
	pulumi.RegisterOutputType(AutoOSUpgradePolicyOutput{})
	pulumi.RegisterOutputType(AutoOSUpgradePolicyPtrOutput{})
	pulumi.RegisterOutputType(AutoOSUpgradePolicyResponseOutput{})
	pulumi.RegisterOutputType(AutoOSUpgradePolicyResponsePtrOutput{})
	pulumi.RegisterOutputType(BootDiagnosticsOutput{})
	pulumi.RegisterOutputType(BootDiagnosticsPtrOutput{})
	pulumi.RegisterOutputType(BootDiagnosticsInstanceViewResponseOutput{})
	pulumi.RegisterOutputType(BootDiagnosticsInstanceViewResponsePtrOutput{})
	pulumi.RegisterOutputType(BootDiagnosticsResponseOutput{})
	pulumi.RegisterOutputType(BootDiagnosticsResponsePtrOutput{})
	pulumi.RegisterOutputType(CreationDataOutput{})
	pulumi.RegisterOutputType(CreationDataResponseOutput{})
	pulumi.RegisterOutputType(DataDiskOutput{})
	pulumi.RegisterOutputType(DataDiskArrayOutput{})
	pulumi.RegisterOutputType(DataDiskResponseOutput{})
	pulumi.RegisterOutputType(DataDiskResponseArrayOutput{})
	pulumi.RegisterOutputType(DiagnosticsProfileOutput{})
	pulumi.RegisterOutputType(DiagnosticsProfilePtrOutput{})
	pulumi.RegisterOutputType(DiagnosticsProfileResponseOutput{})
	pulumi.RegisterOutputType(DiagnosticsProfileResponsePtrOutput{})
	pulumi.RegisterOutputType(DiskEncryptionSettingsOutput{})
	pulumi.RegisterOutputType(DiskEncryptionSettingsPtrOutput{})
	pulumi.RegisterOutputType(DiskEncryptionSettingsResponseOutput{})
	pulumi.RegisterOutputType(DiskEncryptionSettingsResponsePtrOutput{})
	pulumi.RegisterOutputType(DiskEncryptionSettingsResponseArrayOutput{})
	pulumi.RegisterOutputType(DiskInstanceViewResponseOutput{})
	pulumi.RegisterOutputType(DiskInstanceViewResponseArrayOutput{})
	pulumi.RegisterOutputType(DiskSkuOutput{})
	pulumi.RegisterOutputType(DiskSkuPtrOutput{})
	pulumi.RegisterOutputType(DiskSkuResponseOutput{})
	pulumi.RegisterOutputType(DiskSkuResponsePtrOutput{})
	pulumi.RegisterOutputType(EncryptionSettingsOutput{})
	pulumi.RegisterOutputType(EncryptionSettingsPtrOutput{})
	pulumi.RegisterOutputType(EncryptionSettingsResponseOutput{})
	pulumi.RegisterOutputType(EncryptionSettingsResponsePtrOutput{})
	pulumi.RegisterOutputType(HardwareProfileOutput{})
	pulumi.RegisterOutputType(HardwareProfilePtrOutput{})
	pulumi.RegisterOutputType(HardwareProfileResponseOutput{})
	pulumi.RegisterOutputType(HardwareProfileResponsePtrOutput{})
	pulumi.RegisterOutputType(ImageDataDiskOutput{})
	pulumi.RegisterOutputType(ImageDataDiskArrayOutput{})
	pulumi.RegisterOutputType(ImageDataDiskResponseOutput{})
	pulumi.RegisterOutputType(ImageDataDiskResponseArrayOutput{})
	pulumi.RegisterOutputType(ImageDiskReferenceOutput{})
	pulumi.RegisterOutputType(ImageDiskReferencePtrOutput{})
	pulumi.RegisterOutputType(ImageDiskReferenceResponseOutput{})
	pulumi.RegisterOutputType(ImageDiskReferenceResponsePtrOutput{})
	pulumi.RegisterOutputType(ImageOSDiskOutput{})
	pulumi.RegisterOutputType(ImageOSDiskPtrOutput{})
	pulumi.RegisterOutputType(ImageOSDiskResponseOutput{})
	pulumi.RegisterOutputType(ImageOSDiskResponsePtrOutput{})
	pulumi.RegisterOutputType(ImageReferenceOutput{})
	pulumi.RegisterOutputType(ImageReferencePtrOutput{})
	pulumi.RegisterOutputType(ImageReferenceResponseOutput{})
	pulumi.RegisterOutputType(ImageReferenceResponsePtrOutput{})
	pulumi.RegisterOutputType(ImageStorageProfileOutput{})
	pulumi.RegisterOutputType(ImageStorageProfilePtrOutput{})
	pulumi.RegisterOutputType(ImageStorageProfileResponseOutput{})
	pulumi.RegisterOutputType(ImageStorageProfileResponsePtrOutput{})
	pulumi.RegisterOutputType(InstanceViewStatusOutput{})
	pulumi.RegisterOutputType(InstanceViewStatusArrayOutput{})
	pulumi.RegisterOutputType(InstanceViewStatusResponseOutput{})
	pulumi.RegisterOutputType(InstanceViewStatusResponsePtrOutput{})
	pulumi.RegisterOutputType(InstanceViewStatusResponseArrayOutput{})
	pulumi.RegisterOutputType(KeyVaultAndKeyReferenceOutput{})
	pulumi.RegisterOutputType(KeyVaultAndKeyReferencePtrOutput{})
	pulumi.RegisterOutputType(KeyVaultAndKeyReferenceResponseOutput{})
	pulumi.RegisterOutputType(KeyVaultAndKeyReferenceResponsePtrOutput{})
	pulumi.RegisterOutputType(KeyVaultAndSecretReferenceOutput{})
	pulumi.RegisterOutputType(KeyVaultAndSecretReferencePtrOutput{})
	pulumi.RegisterOutputType(KeyVaultAndSecretReferenceResponseOutput{})
	pulumi.RegisterOutputType(KeyVaultAndSecretReferenceResponsePtrOutput{})
	pulumi.RegisterOutputType(KeyVaultKeyReferenceOutput{})
	pulumi.RegisterOutputType(KeyVaultKeyReferencePtrOutput{})
	pulumi.RegisterOutputType(KeyVaultKeyReferenceResponseOutput{})
	pulumi.RegisterOutputType(KeyVaultKeyReferenceResponsePtrOutput{})
	pulumi.RegisterOutputType(KeyVaultSecretReferenceOutput{})
	pulumi.RegisterOutputType(KeyVaultSecretReferencePtrOutput{})
	pulumi.RegisterOutputType(KeyVaultSecretReferenceResponseOutput{})
	pulumi.RegisterOutputType(KeyVaultSecretReferenceResponsePtrOutput{})
	pulumi.RegisterOutputType(LinuxConfigurationOutput{})
	pulumi.RegisterOutputType(LinuxConfigurationPtrOutput{})
	pulumi.RegisterOutputType(LinuxConfigurationResponseOutput{})
	pulumi.RegisterOutputType(LinuxConfigurationResponsePtrOutput{})
	pulumi.RegisterOutputType(LogAnalyticsOutputResponseOutput{})
	pulumi.RegisterOutputType(MaintenanceRedeployStatusResponseOutput{})
	pulumi.RegisterOutputType(MaintenanceRedeployStatusResponsePtrOutput{})
	pulumi.RegisterOutputType(ManagedDiskParametersOutput{})
	pulumi.RegisterOutputType(ManagedDiskParametersPtrOutput{})
	pulumi.RegisterOutputType(ManagedDiskParametersResponseOutput{})
	pulumi.RegisterOutputType(ManagedDiskParametersResponsePtrOutput{})
	pulumi.RegisterOutputType(NetworkInterfaceReferenceOutput{})
	pulumi.RegisterOutputType(NetworkInterfaceReferenceArrayOutput{})
	pulumi.RegisterOutputType(NetworkInterfaceReferenceResponseOutput{})
	pulumi.RegisterOutputType(NetworkInterfaceReferenceResponseArrayOutput{})
	pulumi.RegisterOutputType(NetworkProfileOutput{})
	pulumi.RegisterOutputType(NetworkProfilePtrOutput{})
	pulumi.RegisterOutputType(NetworkProfileResponseOutput{})
	pulumi.RegisterOutputType(NetworkProfileResponsePtrOutput{})
	pulumi.RegisterOutputType(OSDiskOutput{})
	pulumi.RegisterOutputType(OSDiskPtrOutput{})
	pulumi.RegisterOutputType(OSDiskResponseOutput{})
	pulumi.RegisterOutputType(OSDiskResponsePtrOutput{})
	pulumi.RegisterOutputType(OSProfileOutput{})
	pulumi.RegisterOutputType(OSProfilePtrOutput{})
	pulumi.RegisterOutputType(OSProfileResponseOutput{})
	pulumi.RegisterOutputType(OSProfileResponsePtrOutput{})
	pulumi.RegisterOutputType(PlanOutput{})
	pulumi.RegisterOutputType(PlanPtrOutput{})
	pulumi.RegisterOutputType(PlanResponseOutput{})
	pulumi.RegisterOutputType(PlanResponsePtrOutput{})
	pulumi.RegisterOutputType(RollingUpgradePolicyOutput{})
	pulumi.RegisterOutputType(RollingUpgradePolicyPtrOutput{})
	pulumi.RegisterOutputType(RollingUpgradePolicyResponseOutput{})
	pulumi.RegisterOutputType(RollingUpgradePolicyResponsePtrOutput{})
	pulumi.RegisterOutputType(SkuOutput{})
	pulumi.RegisterOutputType(SkuPtrOutput{})
	pulumi.RegisterOutputType(SkuResponseOutput{})
	pulumi.RegisterOutputType(SkuResponsePtrOutput{})
	pulumi.RegisterOutputType(SnapshotSkuOutput{})
	pulumi.RegisterOutputType(SnapshotSkuPtrOutput{})
	pulumi.RegisterOutputType(SnapshotSkuResponseOutput{})
	pulumi.RegisterOutputType(SnapshotSkuResponsePtrOutput{})
	pulumi.RegisterOutputType(SourceVaultOutput{})
	pulumi.RegisterOutputType(SourceVaultPtrOutput{})
	pulumi.RegisterOutputType(SourceVaultResponseOutput{})
	pulumi.RegisterOutputType(SourceVaultResponsePtrOutput{})
	pulumi.RegisterOutputType(SshConfigurationOutput{})
	pulumi.RegisterOutputType(SshConfigurationPtrOutput{})
	pulumi.RegisterOutputType(SshConfigurationResponseOutput{})
	pulumi.RegisterOutputType(SshConfigurationResponsePtrOutput{})
	pulumi.RegisterOutputType(SshPublicKeyOutput{})
	pulumi.RegisterOutputType(SshPublicKeyArrayOutput{})
	pulumi.RegisterOutputType(SshPublicKeyResponseOutput{})
	pulumi.RegisterOutputType(SshPublicKeyResponseArrayOutput{})
	pulumi.RegisterOutputType(StorageProfileOutput{})
	pulumi.RegisterOutputType(StorageProfilePtrOutput{})
	pulumi.RegisterOutputType(StorageProfileResponseOutput{})
	pulumi.RegisterOutputType(StorageProfileResponsePtrOutput{})
	pulumi.RegisterOutputType(SubResourceOutput{})
	pulumi.RegisterOutputType(SubResourcePtrOutput{})
	pulumi.RegisterOutputType(SubResourceArrayOutput{})
	pulumi.RegisterOutputType(SubResourceResponseOutput{})
	pulumi.RegisterOutputType(SubResourceResponsePtrOutput{})
	pulumi.RegisterOutputType(SubResourceResponseArrayOutput{})
	pulumi.RegisterOutputType(UpgradePolicyOutput{})
	pulumi.RegisterOutputType(UpgradePolicyPtrOutput{})
	pulumi.RegisterOutputType(UpgradePolicyResponseOutput{})
	pulumi.RegisterOutputType(UpgradePolicyResponsePtrOutput{})
	pulumi.RegisterOutputType(VaultCertificateOutput{})
	pulumi.RegisterOutputType(VaultCertificateArrayOutput{})
	pulumi.RegisterOutputType(VaultCertificateResponseOutput{})
	pulumi.RegisterOutputType(VaultCertificateResponseArrayOutput{})
	pulumi.RegisterOutputType(VaultSecretGroupOutput{})
	pulumi.RegisterOutputType(VaultSecretGroupArrayOutput{})
	pulumi.RegisterOutputType(VaultSecretGroupResponseOutput{})
	pulumi.RegisterOutputType(VaultSecretGroupResponseArrayOutput{})
	pulumi.RegisterOutputType(VirtualHardDiskOutput{})
	pulumi.RegisterOutputType(VirtualHardDiskPtrOutput{})
	pulumi.RegisterOutputType(VirtualHardDiskResponseOutput{})
	pulumi.RegisterOutputType(VirtualHardDiskResponsePtrOutput{})
	pulumi.RegisterOutputType(VirtualMachineAgentInstanceViewResponseOutput{})
	pulumi.RegisterOutputType(VirtualMachineAgentInstanceViewResponsePtrOutput{})
	pulumi.RegisterOutputType(VirtualMachineExtensionHandlerInstanceViewResponseOutput{})
	pulumi.RegisterOutputType(VirtualMachineExtensionHandlerInstanceViewResponseArrayOutput{})
	pulumi.RegisterOutputType(VirtualMachineExtensionInstanceViewOutput{})
	pulumi.RegisterOutputType(VirtualMachineExtensionInstanceViewPtrOutput{})
	pulumi.RegisterOutputType(VirtualMachineExtensionInstanceViewResponseOutput{})
	pulumi.RegisterOutputType(VirtualMachineExtensionInstanceViewResponsePtrOutput{})
	pulumi.RegisterOutputType(VirtualMachineExtensionInstanceViewResponseArrayOutput{})
	pulumi.RegisterOutputType(VirtualMachineExtensionResponseOutput{})
	pulumi.RegisterOutputType(VirtualMachineExtensionResponseArrayOutput{})
	pulumi.RegisterOutputType(VirtualMachineHealthStatusResponseOutput{})
	pulumi.RegisterOutputType(VirtualMachineIdentityOutput{})
	pulumi.RegisterOutputType(VirtualMachineIdentityPtrOutput{})
	pulumi.RegisterOutputType(VirtualMachineIdentityResponseOutput{})
	pulumi.RegisterOutputType(VirtualMachineIdentityResponsePtrOutput{})
	pulumi.RegisterOutputType(VirtualMachineInstanceViewResponseOutput{})
	pulumi.RegisterOutputType(VirtualMachineScaleSetDataDiskOutput{})
	pulumi.RegisterOutputType(VirtualMachineScaleSetDataDiskArrayOutput{})
	pulumi.RegisterOutputType(VirtualMachineScaleSetDataDiskResponseOutput{})
	pulumi.RegisterOutputType(VirtualMachineScaleSetDataDiskResponseArrayOutput{})
	pulumi.RegisterOutputType(VirtualMachineScaleSetExtensionTypeOutput{})
	pulumi.RegisterOutputType(VirtualMachineScaleSetExtensionTypeArrayOutput{})
	pulumi.RegisterOutputType(VirtualMachineScaleSetExtensionProfileOutput{})
	pulumi.RegisterOutputType(VirtualMachineScaleSetExtensionProfilePtrOutput{})
	pulumi.RegisterOutputType(VirtualMachineScaleSetExtensionProfileResponseOutput{})
	pulumi.RegisterOutputType(VirtualMachineScaleSetExtensionProfileResponsePtrOutput{})
	pulumi.RegisterOutputType(VirtualMachineScaleSetExtensionResponseOutput{})
	pulumi.RegisterOutputType(VirtualMachineScaleSetExtensionResponseArrayOutput{})
	pulumi.RegisterOutputType(VirtualMachineScaleSetIPConfigurationOutput{})
	pulumi.RegisterOutputType(VirtualMachineScaleSetIPConfigurationArrayOutput{})
	pulumi.RegisterOutputType(VirtualMachineScaleSetIPConfigurationResponseOutput{})
	pulumi.RegisterOutputType(VirtualMachineScaleSetIPConfigurationResponseArrayOutput{})
	pulumi.RegisterOutputType(VirtualMachineScaleSetIdentityOutput{})
	pulumi.RegisterOutputType(VirtualMachineScaleSetIdentityPtrOutput{})
	pulumi.RegisterOutputType(VirtualMachineScaleSetIdentityResponseOutput{})
	pulumi.RegisterOutputType(VirtualMachineScaleSetIdentityResponsePtrOutput{})
	pulumi.RegisterOutputType(VirtualMachineScaleSetIpTagOutput{})
	pulumi.RegisterOutputType(VirtualMachineScaleSetIpTagArrayOutput{})
	pulumi.RegisterOutputType(VirtualMachineScaleSetIpTagResponseOutput{})
	pulumi.RegisterOutputType(VirtualMachineScaleSetIpTagResponseArrayOutput{})
	pulumi.RegisterOutputType(VirtualMachineScaleSetManagedDiskParametersOutput{})
	pulumi.RegisterOutputType(VirtualMachineScaleSetManagedDiskParametersPtrOutput{})
	pulumi.RegisterOutputType(VirtualMachineScaleSetManagedDiskParametersResponseOutput{})
	pulumi.RegisterOutputType(VirtualMachineScaleSetManagedDiskParametersResponsePtrOutput{})
	pulumi.RegisterOutputType(VirtualMachineScaleSetNetworkConfigurationOutput{})
	pulumi.RegisterOutputType(VirtualMachineScaleSetNetworkConfigurationArrayOutput{})
	pulumi.RegisterOutputType(VirtualMachineScaleSetNetworkConfigurationDnsSettingsOutput{})
	pulumi.RegisterOutputType(VirtualMachineScaleSetNetworkConfigurationDnsSettingsPtrOutput{})
	pulumi.RegisterOutputType(VirtualMachineScaleSetNetworkConfigurationDnsSettingsResponseOutput{})
	pulumi.RegisterOutputType(VirtualMachineScaleSetNetworkConfigurationDnsSettingsResponsePtrOutput{})
	pulumi.RegisterOutputType(VirtualMachineScaleSetNetworkConfigurationResponseOutput{})
	pulumi.RegisterOutputType(VirtualMachineScaleSetNetworkConfigurationResponseArrayOutput{})
	pulumi.RegisterOutputType(VirtualMachineScaleSetNetworkProfileOutput{})
	pulumi.RegisterOutputType(VirtualMachineScaleSetNetworkProfilePtrOutput{})
	pulumi.RegisterOutputType(VirtualMachineScaleSetNetworkProfileResponseOutput{})
	pulumi.RegisterOutputType(VirtualMachineScaleSetNetworkProfileResponsePtrOutput{})
	pulumi.RegisterOutputType(VirtualMachineScaleSetOSDiskOutput{})
	pulumi.RegisterOutputType(VirtualMachineScaleSetOSDiskPtrOutput{})
	pulumi.RegisterOutputType(VirtualMachineScaleSetOSDiskResponseOutput{})
	pulumi.RegisterOutputType(VirtualMachineScaleSetOSDiskResponsePtrOutput{})
	pulumi.RegisterOutputType(VirtualMachineScaleSetOSProfileOutput{})
	pulumi.RegisterOutputType(VirtualMachineScaleSetOSProfilePtrOutput{})
	pulumi.RegisterOutputType(VirtualMachineScaleSetOSProfileResponseOutput{})
	pulumi.RegisterOutputType(VirtualMachineScaleSetOSProfileResponsePtrOutput{})
	pulumi.RegisterOutputType(VirtualMachineScaleSetPublicIPAddressConfigurationOutput{})
	pulumi.RegisterOutputType(VirtualMachineScaleSetPublicIPAddressConfigurationPtrOutput{})
	pulumi.RegisterOutputType(VirtualMachineScaleSetPublicIPAddressConfigurationDnsSettingsOutput{})
	pulumi.RegisterOutputType(VirtualMachineScaleSetPublicIPAddressConfigurationDnsSettingsPtrOutput{})
	pulumi.RegisterOutputType(VirtualMachineScaleSetPublicIPAddressConfigurationDnsSettingsResponseOutput{})
	pulumi.RegisterOutputType(VirtualMachineScaleSetPublicIPAddressConfigurationDnsSettingsResponsePtrOutput{})
	pulumi.RegisterOutputType(VirtualMachineScaleSetPublicIPAddressConfigurationResponseOutput{})
	pulumi.RegisterOutputType(VirtualMachineScaleSetPublicIPAddressConfigurationResponsePtrOutput{})
	pulumi.RegisterOutputType(VirtualMachineScaleSetStorageProfileOutput{})
	pulumi.RegisterOutputType(VirtualMachineScaleSetStorageProfilePtrOutput{})
	pulumi.RegisterOutputType(VirtualMachineScaleSetStorageProfileResponseOutput{})
	pulumi.RegisterOutputType(VirtualMachineScaleSetStorageProfileResponsePtrOutput{})
	pulumi.RegisterOutputType(VirtualMachineScaleSetVMInstanceViewResponseOutput{})
	pulumi.RegisterOutputType(VirtualMachineScaleSetVMProfileOutput{})
	pulumi.RegisterOutputType(VirtualMachineScaleSetVMProfilePtrOutput{})
	pulumi.RegisterOutputType(VirtualMachineScaleSetVMProfileResponseOutput{})
	pulumi.RegisterOutputType(VirtualMachineScaleSetVMProfileResponsePtrOutput{})
	pulumi.RegisterOutputType(WinRMConfigurationOutput{})
	pulumi.RegisterOutputType(WinRMConfigurationPtrOutput{})
	pulumi.RegisterOutputType(WinRMConfigurationResponseOutput{})
	pulumi.RegisterOutputType(WinRMConfigurationResponsePtrOutput{})
	pulumi.RegisterOutputType(WinRMListenerOutput{})
	pulumi.RegisterOutputType(WinRMListenerArrayOutput{})
	pulumi.RegisterOutputType(WinRMListenerResponseOutput{})
	pulumi.RegisterOutputType(WinRMListenerResponseArrayOutput{})
	pulumi.RegisterOutputType(WindowsConfigurationOutput{})
	pulumi.RegisterOutputType(WindowsConfigurationPtrOutput{})
	pulumi.RegisterOutputType(WindowsConfigurationResponseOutput{})
	pulumi.RegisterOutputType(WindowsConfigurationResponsePtrOutput{})
}
