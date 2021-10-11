


package v20210601preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type BasicLoginInformation struct {
	Password *string `pulumi:"password"`
	Username *string `pulumi:"username"`
}





type BasicLoginInformationInput interface {
	pulumi.Input

	ToBasicLoginInformationOutput() BasicLoginInformationOutput
	ToBasicLoginInformationOutputWithContext(context.Context) BasicLoginInformationOutput
}

type BasicLoginInformationArgs struct {
	Password pulumi.StringPtrInput `pulumi:"password"`
	Username pulumi.StringPtrInput `pulumi:"username"`
}

func (BasicLoginInformationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*BasicLoginInformation)(nil)).Elem()
}

func (i BasicLoginInformationArgs) ToBasicLoginInformationOutput() BasicLoginInformationOutput {
	return i.ToBasicLoginInformationOutputWithContext(context.Background())
}

func (i BasicLoginInformationArgs) ToBasicLoginInformationOutputWithContext(ctx context.Context) BasicLoginInformationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BasicLoginInformationOutput)
}

func (i BasicLoginInformationArgs) ToBasicLoginInformationPtrOutput() BasicLoginInformationPtrOutput {
	return i.ToBasicLoginInformationPtrOutputWithContext(context.Background())
}

func (i BasicLoginInformationArgs) ToBasicLoginInformationPtrOutputWithContext(ctx context.Context) BasicLoginInformationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BasicLoginInformationOutput).ToBasicLoginInformationPtrOutputWithContext(ctx)
}









type BasicLoginInformationPtrInput interface {
	pulumi.Input

	ToBasicLoginInformationPtrOutput() BasicLoginInformationPtrOutput
	ToBasicLoginInformationPtrOutputWithContext(context.Context) BasicLoginInformationPtrOutput
}

type basicLoginInformationPtrType BasicLoginInformationArgs

func BasicLoginInformationPtr(v *BasicLoginInformationArgs) BasicLoginInformationPtrInput {
	return (*basicLoginInformationPtrType)(v)
}

func (*basicLoginInformationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**BasicLoginInformation)(nil)).Elem()
}

func (i *basicLoginInformationPtrType) ToBasicLoginInformationPtrOutput() BasicLoginInformationPtrOutput {
	return i.ToBasicLoginInformationPtrOutputWithContext(context.Background())
}

func (i *basicLoginInformationPtrType) ToBasicLoginInformationPtrOutputWithContext(ctx context.Context) BasicLoginInformationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BasicLoginInformationPtrOutput)
}

type BasicLoginInformationOutput struct{ *pulumi.OutputState }

func (BasicLoginInformationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BasicLoginInformation)(nil)).Elem()
}

func (o BasicLoginInformationOutput) ToBasicLoginInformationOutput() BasicLoginInformationOutput {
	return o
}

func (o BasicLoginInformationOutput) ToBasicLoginInformationOutputWithContext(ctx context.Context) BasicLoginInformationOutput {
	return o
}

func (o BasicLoginInformationOutput) ToBasicLoginInformationPtrOutput() BasicLoginInformationPtrOutput {
	return o.ToBasicLoginInformationPtrOutputWithContext(context.Background())
}

func (o BasicLoginInformationOutput) ToBasicLoginInformationPtrOutputWithContext(ctx context.Context) BasicLoginInformationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v BasicLoginInformation) *BasicLoginInformation {
		return &v
	}).(BasicLoginInformationPtrOutput)
}

func (o BasicLoginInformationOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BasicLoginInformation) *string { return v.Password }).(pulumi.StringPtrOutput)
}

func (o BasicLoginInformationOutput) Username() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BasicLoginInformation) *string { return v.Username }).(pulumi.StringPtrOutput)
}

type BasicLoginInformationPtrOutput struct{ *pulumi.OutputState }

func (BasicLoginInformationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BasicLoginInformation)(nil)).Elem()
}

func (o BasicLoginInformationPtrOutput) ToBasicLoginInformationPtrOutput() BasicLoginInformationPtrOutput {
	return o
}

func (o BasicLoginInformationPtrOutput) ToBasicLoginInformationPtrOutputWithContext(ctx context.Context) BasicLoginInformationPtrOutput {
	return o
}

func (o BasicLoginInformationPtrOutput) Elem() BasicLoginInformationOutput {
	return o.ApplyT(func(v *BasicLoginInformation) BasicLoginInformation {
		if v != nil {
			return *v
		}
		var ret BasicLoginInformation
		return ret
	}).(BasicLoginInformationOutput)
}

func (o BasicLoginInformationPtrOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BasicLoginInformation) *string {
		if v == nil {
			return nil
		}
		return v.Password
	}).(pulumi.StringPtrOutput)
}

func (o BasicLoginInformationPtrOutput) Username() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BasicLoginInformation) *string {
		if v == nil {
			return nil
		}
		return v.Username
	}).(pulumi.StringPtrOutput)
}

type BasicLoginInformationResponse struct {
	Username *string `pulumi:"username"`
}





type BasicLoginInformationResponseInput interface {
	pulumi.Input

	ToBasicLoginInformationResponseOutput() BasicLoginInformationResponseOutput
	ToBasicLoginInformationResponseOutputWithContext(context.Context) BasicLoginInformationResponseOutput
}

type BasicLoginInformationResponseArgs struct {
	Username pulumi.StringPtrInput `pulumi:"username"`
}

func (BasicLoginInformationResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*BasicLoginInformationResponse)(nil)).Elem()
}

func (i BasicLoginInformationResponseArgs) ToBasicLoginInformationResponseOutput() BasicLoginInformationResponseOutput {
	return i.ToBasicLoginInformationResponseOutputWithContext(context.Background())
}

func (i BasicLoginInformationResponseArgs) ToBasicLoginInformationResponseOutputWithContext(ctx context.Context) BasicLoginInformationResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BasicLoginInformationResponseOutput)
}

func (i BasicLoginInformationResponseArgs) ToBasicLoginInformationResponsePtrOutput() BasicLoginInformationResponsePtrOutput {
	return i.ToBasicLoginInformationResponsePtrOutputWithContext(context.Background())
}

func (i BasicLoginInformationResponseArgs) ToBasicLoginInformationResponsePtrOutputWithContext(ctx context.Context) BasicLoginInformationResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BasicLoginInformationResponseOutput).ToBasicLoginInformationResponsePtrOutputWithContext(ctx)
}









type BasicLoginInformationResponsePtrInput interface {
	pulumi.Input

	ToBasicLoginInformationResponsePtrOutput() BasicLoginInformationResponsePtrOutput
	ToBasicLoginInformationResponsePtrOutputWithContext(context.Context) BasicLoginInformationResponsePtrOutput
}

type basicLoginInformationResponsePtrType BasicLoginInformationResponseArgs

func BasicLoginInformationResponsePtr(v *BasicLoginInformationResponseArgs) BasicLoginInformationResponsePtrInput {
	return (*basicLoginInformationResponsePtrType)(v)
}

func (*basicLoginInformationResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**BasicLoginInformationResponse)(nil)).Elem()
}

func (i *basicLoginInformationResponsePtrType) ToBasicLoginInformationResponsePtrOutput() BasicLoginInformationResponsePtrOutput {
	return i.ToBasicLoginInformationResponsePtrOutputWithContext(context.Background())
}

func (i *basicLoginInformationResponsePtrType) ToBasicLoginInformationResponsePtrOutputWithContext(ctx context.Context) BasicLoginInformationResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BasicLoginInformationResponsePtrOutput)
}

type BasicLoginInformationResponseOutput struct{ *pulumi.OutputState }

func (BasicLoginInformationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BasicLoginInformationResponse)(nil)).Elem()
}

func (o BasicLoginInformationResponseOutput) ToBasicLoginInformationResponseOutput() BasicLoginInformationResponseOutput {
	return o
}

func (o BasicLoginInformationResponseOutput) ToBasicLoginInformationResponseOutputWithContext(ctx context.Context) BasicLoginInformationResponseOutput {
	return o
}

func (o BasicLoginInformationResponseOutput) ToBasicLoginInformationResponsePtrOutput() BasicLoginInformationResponsePtrOutput {
	return o.ToBasicLoginInformationResponsePtrOutputWithContext(context.Background())
}

func (o BasicLoginInformationResponseOutput) ToBasicLoginInformationResponsePtrOutputWithContext(ctx context.Context) BasicLoginInformationResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v BasicLoginInformationResponse) *BasicLoginInformationResponse {
		return &v
	}).(BasicLoginInformationResponsePtrOutput)
}

func (o BasicLoginInformationResponseOutput) Username() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BasicLoginInformationResponse) *string { return v.Username }).(pulumi.StringPtrOutput)
}

type BasicLoginInformationResponsePtrOutput struct{ *pulumi.OutputState }

func (BasicLoginInformationResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BasicLoginInformationResponse)(nil)).Elem()
}

func (o BasicLoginInformationResponsePtrOutput) ToBasicLoginInformationResponsePtrOutput() BasicLoginInformationResponsePtrOutput {
	return o
}

func (o BasicLoginInformationResponsePtrOutput) ToBasicLoginInformationResponsePtrOutputWithContext(ctx context.Context) BasicLoginInformationResponsePtrOutput {
	return o
}

func (o BasicLoginInformationResponsePtrOutput) Elem() BasicLoginInformationResponseOutput {
	return o.ApplyT(func(v *BasicLoginInformationResponse) BasicLoginInformationResponse {
		if v != nil {
			return *v
		}
		var ret BasicLoginInformationResponse
		return ret
	}).(BasicLoginInformationResponseOutput)
}

func (o BasicLoginInformationResponsePtrOutput) Username() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BasicLoginInformationResponse) *string {
		if v == nil {
			return nil
		}
		return v.Username
	}).(pulumi.StringPtrOutput)
}

type DataControllerProperties struct {
	BasicLoginInformation       *BasicLoginInformation       `pulumi:"basicLoginInformation"`
	K8sRaw                      interface{}                  `pulumi:"k8sRaw"`
	LastUploadedDate            *string                      `pulumi:"lastUploadedDate"`
	LogAnalyticsWorkspaceConfig *LogAnalyticsWorkspaceConfig `pulumi:"logAnalyticsWorkspaceConfig"`
	OnPremiseProperty           *OnPremiseProperty           `pulumi:"onPremiseProperty"`
	UploadServicePrincipal      *UploadServicePrincipal      `pulumi:"uploadServicePrincipal"`
	UploadWatermark             *UploadWatermark             `pulumi:"uploadWatermark"`
}





type DataControllerPropertiesInput interface {
	pulumi.Input

	ToDataControllerPropertiesOutput() DataControllerPropertiesOutput
	ToDataControllerPropertiesOutputWithContext(context.Context) DataControllerPropertiesOutput
}

type DataControllerPropertiesArgs struct {
	BasicLoginInformation       BasicLoginInformationPtrInput       `pulumi:"basicLoginInformation"`
	K8sRaw                      pulumi.Input                        `pulumi:"k8sRaw"`
	LastUploadedDate            pulumi.StringPtrInput               `pulumi:"lastUploadedDate"`
	LogAnalyticsWorkspaceConfig LogAnalyticsWorkspaceConfigPtrInput `pulumi:"logAnalyticsWorkspaceConfig"`
	OnPremiseProperty           OnPremisePropertyPtrInput           `pulumi:"onPremiseProperty"`
	UploadServicePrincipal      UploadServicePrincipalPtrInput      `pulumi:"uploadServicePrincipal"`
	UploadWatermark             UploadWatermarkPtrInput             `pulumi:"uploadWatermark"`
}

func (DataControllerPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DataControllerProperties)(nil)).Elem()
}

func (i DataControllerPropertiesArgs) ToDataControllerPropertiesOutput() DataControllerPropertiesOutput {
	return i.ToDataControllerPropertiesOutputWithContext(context.Background())
}

func (i DataControllerPropertiesArgs) ToDataControllerPropertiesOutputWithContext(ctx context.Context) DataControllerPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataControllerPropertiesOutput)
}

func (i DataControllerPropertiesArgs) ToDataControllerPropertiesPtrOutput() DataControllerPropertiesPtrOutput {
	return i.ToDataControllerPropertiesPtrOutputWithContext(context.Background())
}

func (i DataControllerPropertiesArgs) ToDataControllerPropertiesPtrOutputWithContext(ctx context.Context) DataControllerPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataControllerPropertiesOutput).ToDataControllerPropertiesPtrOutputWithContext(ctx)
}









type DataControllerPropertiesPtrInput interface {
	pulumi.Input

	ToDataControllerPropertiesPtrOutput() DataControllerPropertiesPtrOutput
	ToDataControllerPropertiesPtrOutputWithContext(context.Context) DataControllerPropertiesPtrOutput
}

type dataControllerPropertiesPtrType DataControllerPropertiesArgs

func DataControllerPropertiesPtr(v *DataControllerPropertiesArgs) DataControllerPropertiesPtrInput {
	return (*dataControllerPropertiesPtrType)(v)
}

func (*dataControllerPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DataControllerProperties)(nil)).Elem()
}

func (i *dataControllerPropertiesPtrType) ToDataControllerPropertiesPtrOutput() DataControllerPropertiesPtrOutput {
	return i.ToDataControllerPropertiesPtrOutputWithContext(context.Background())
}

func (i *dataControllerPropertiesPtrType) ToDataControllerPropertiesPtrOutputWithContext(ctx context.Context) DataControllerPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataControllerPropertiesPtrOutput)
}

type DataControllerPropertiesOutput struct{ *pulumi.OutputState }

func (DataControllerPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DataControllerProperties)(nil)).Elem()
}

func (o DataControllerPropertiesOutput) ToDataControllerPropertiesOutput() DataControllerPropertiesOutput {
	return o
}

func (o DataControllerPropertiesOutput) ToDataControllerPropertiesOutputWithContext(ctx context.Context) DataControllerPropertiesOutput {
	return o
}

func (o DataControllerPropertiesOutput) ToDataControllerPropertiesPtrOutput() DataControllerPropertiesPtrOutput {
	return o.ToDataControllerPropertiesPtrOutputWithContext(context.Background())
}

func (o DataControllerPropertiesOutput) ToDataControllerPropertiesPtrOutputWithContext(ctx context.Context) DataControllerPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DataControllerProperties) *DataControllerProperties {
		return &v
	}).(DataControllerPropertiesPtrOutput)
}

func (o DataControllerPropertiesOutput) BasicLoginInformation() BasicLoginInformationPtrOutput {
	return o.ApplyT(func(v DataControllerProperties) *BasicLoginInformation { return v.BasicLoginInformation }).(BasicLoginInformationPtrOutput)
}

func (o DataControllerPropertiesOutput) K8sRaw() pulumi.AnyOutput {
	return o.ApplyT(func(v DataControllerProperties) interface{} { return v.K8sRaw }).(pulumi.AnyOutput)
}

func (o DataControllerPropertiesOutput) LastUploadedDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DataControllerProperties) *string { return v.LastUploadedDate }).(pulumi.StringPtrOutput)
}

func (o DataControllerPropertiesOutput) LogAnalyticsWorkspaceConfig() LogAnalyticsWorkspaceConfigPtrOutput {
	return o.ApplyT(func(v DataControllerProperties) *LogAnalyticsWorkspaceConfig { return v.LogAnalyticsWorkspaceConfig }).(LogAnalyticsWorkspaceConfigPtrOutput)
}

func (o DataControllerPropertiesOutput) OnPremiseProperty() OnPremisePropertyPtrOutput {
	return o.ApplyT(func(v DataControllerProperties) *OnPremiseProperty { return v.OnPremiseProperty }).(OnPremisePropertyPtrOutput)
}

func (o DataControllerPropertiesOutput) UploadServicePrincipal() UploadServicePrincipalPtrOutput {
	return o.ApplyT(func(v DataControllerProperties) *UploadServicePrincipal { return v.UploadServicePrincipal }).(UploadServicePrincipalPtrOutput)
}

func (o DataControllerPropertiesOutput) UploadWatermark() UploadWatermarkPtrOutput {
	return o.ApplyT(func(v DataControllerProperties) *UploadWatermark { return v.UploadWatermark }).(UploadWatermarkPtrOutput)
}

type DataControllerPropertiesPtrOutput struct{ *pulumi.OutputState }

func (DataControllerPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DataControllerProperties)(nil)).Elem()
}

func (o DataControllerPropertiesPtrOutput) ToDataControllerPropertiesPtrOutput() DataControllerPropertiesPtrOutput {
	return o
}

func (o DataControllerPropertiesPtrOutput) ToDataControllerPropertiesPtrOutputWithContext(ctx context.Context) DataControllerPropertiesPtrOutput {
	return o
}

func (o DataControllerPropertiesPtrOutput) Elem() DataControllerPropertiesOutput {
	return o.ApplyT(func(v *DataControllerProperties) DataControllerProperties {
		if v != nil {
			return *v
		}
		var ret DataControllerProperties
		return ret
	}).(DataControllerPropertiesOutput)
}

func (o DataControllerPropertiesPtrOutput) BasicLoginInformation() BasicLoginInformationPtrOutput {
	return o.ApplyT(func(v *DataControllerProperties) *BasicLoginInformation {
		if v == nil {
			return nil
		}
		return v.BasicLoginInformation
	}).(BasicLoginInformationPtrOutput)
}

func (o DataControllerPropertiesPtrOutput) K8sRaw() pulumi.AnyOutput {
	return o.ApplyT(func(v *DataControllerProperties) interface{} {
		if v == nil {
			return nil
		}
		return v.K8sRaw
	}).(pulumi.AnyOutput)
}

func (o DataControllerPropertiesPtrOutput) LastUploadedDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DataControllerProperties) *string {
		if v == nil {
			return nil
		}
		return v.LastUploadedDate
	}).(pulumi.StringPtrOutput)
}

func (o DataControllerPropertiesPtrOutput) LogAnalyticsWorkspaceConfig() LogAnalyticsWorkspaceConfigPtrOutput {
	return o.ApplyT(func(v *DataControllerProperties) *LogAnalyticsWorkspaceConfig {
		if v == nil {
			return nil
		}
		return v.LogAnalyticsWorkspaceConfig
	}).(LogAnalyticsWorkspaceConfigPtrOutput)
}

func (o DataControllerPropertiesPtrOutput) OnPremiseProperty() OnPremisePropertyPtrOutput {
	return o.ApplyT(func(v *DataControllerProperties) *OnPremiseProperty {
		if v == nil {
			return nil
		}
		return v.OnPremiseProperty
	}).(OnPremisePropertyPtrOutput)
}

func (o DataControllerPropertiesPtrOutput) UploadServicePrincipal() UploadServicePrincipalPtrOutput {
	return o.ApplyT(func(v *DataControllerProperties) *UploadServicePrincipal {
		if v == nil {
			return nil
		}
		return v.UploadServicePrincipal
	}).(UploadServicePrincipalPtrOutput)
}

func (o DataControllerPropertiesPtrOutput) UploadWatermark() UploadWatermarkPtrOutput {
	return o.ApplyT(func(v *DataControllerProperties) *UploadWatermark {
		if v == nil {
			return nil
		}
		return v.UploadWatermark
	}).(UploadWatermarkPtrOutput)
}

type DataControllerPropertiesResponse struct {
	BasicLoginInformation       *BasicLoginInformationResponse       `pulumi:"basicLoginInformation"`
	K8sRaw                      interface{}                          `pulumi:"k8sRaw"`
	LastUploadedDate            *string                              `pulumi:"lastUploadedDate"`
	LogAnalyticsWorkspaceConfig *LogAnalyticsWorkspaceConfigResponse `pulumi:"logAnalyticsWorkspaceConfig"`
	OnPremiseProperty           *OnPremisePropertyResponse           `pulumi:"onPremiseProperty"`
	ProvisioningState           string                               `pulumi:"provisioningState"`
	UploadServicePrincipal      *UploadServicePrincipalResponse      `pulumi:"uploadServicePrincipal"`
	UploadWatermark             *UploadWatermarkResponse             `pulumi:"uploadWatermark"`
}





type DataControllerPropertiesResponseInput interface {
	pulumi.Input

	ToDataControllerPropertiesResponseOutput() DataControllerPropertiesResponseOutput
	ToDataControllerPropertiesResponseOutputWithContext(context.Context) DataControllerPropertiesResponseOutput
}

type DataControllerPropertiesResponseArgs struct {
	BasicLoginInformation       BasicLoginInformationResponsePtrInput       `pulumi:"basicLoginInformation"`
	K8sRaw                      pulumi.Input                                `pulumi:"k8sRaw"`
	LastUploadedDate            pulumi.StringPtrInput                       `pulumi:"lastUploadedDate"`
	LogAnalyticsWorkspaceConfig LogAnalyticsWorkspaceConfigResponsePtrInput `pulumi:"logAnalyticsWorkspaceConfig"`
	OnPremiseProperty           OnPremisePropertyResponsePtrInput           `pulumi:"onPremiseProperty"`
	ProvisioningState           pulumi.StringInput                          `pulumi:"provisioningState"`
	UploadServicePrincipal      UploadServicePrincipalResponsePtrInput      `pulumi:"uploadServicePrincipal"`
	UploadWatermark             UploadWatermarkResponsePtrInput             `pulumi:"uploadWatermark"`
}

func (DataControllerPropertiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DataControllerPropertiesResponse)(nil)).Elem()
}

func (i DataControllerPropertiesResponseArgs) ToDataControllerPropertiesResponseOutput() DataControllerPropertiesResponseOutput {
	return i.ToDataControllerPropertiesResponseOutputWithContext(context.Background())
}

func (i DataControllerPropertiesResponseArgs) ToDataControllerPropertiesResponseOutputWithContext(ctx context.Context) DataControllerPropertiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataControllerPropertiesResponseOutput)
}

func (i DataControllerPropertiesResponseArgs) ToDataControllerPropertiesResponsePtrOutput() DataControllerPropertiesResponsePtrOutput {
	return i.ToDataControllerPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i DataControllerPropertiesResponseArgs) ToDataControllerPropertiesResponsePtrOutputWithContext(ctx context.Context) DataControllerPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataControllerPropertiesResponseOutput).ToDataControllerPropertiesResponsePtrOutputWithContext(ctx)
}









type DataControllerPropertiesResponsePtrInput interface {
	pulumi.Input

	ToDataControllerPropertiesResponsePtrOutput() DataControllerPropertiesResponsePtrOutput
	ToDataControllerPropertiesResponsePtrOutputWithContext(context.Context) DataControllerPropertiesResponsePtrOutput
}

type dataControllerPropertiesResponsePtrType DataControllerPropertiesResponseArgs

func DataControllerPropertiesResponsePtr(v *DataControllerPropertiesResponseArgs) DataControllerPropertiesResponsePtrInput {
	return (*dataControllerPropertiesResponsePtrType)(v)
}

func (*dataControllerPropertiesResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DataControllerPropertiesResponse)(nil)).Elem()
}

func (i *dataControllerPropertiesResponsePtrType) ToDataControllerPropertiesResponsePtrOutput() DataControllerPropertiesResponsePtrOutput {
	return i.ToDataControllerPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i *dataControllerPropertiesResponsePtrType) ToDataControllerPropertiesResponsePtrOutputWithContext(ctx context.Context) DataControllerPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataControllerPropertiesResponsePtrOutput)
}

type DataControllerPropertiesResponseOutput struct{ *pulumi.OutputState }

func (DataControllerPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DataControllerPropertiesResponse)(nil)).Elem()
}

func (o DataControllerPropertiesResponseOutput) ToDataControllerPropertiesResponseOutput() DataControllerPropertiesResponseOutput {
	return o
}

func (o DataControllerPropertiesResponseOutput) ToDataControllerPropertiesResponseOutputWithContext(ctx context.Context) DataControllerPropertiesResponseOutput {
	return o
}

func (o DataControllerPropertiesResponseOutput) ToDataControllerPropertiesResponsePtrOutput() DataControllerPropertiesResponsePtrOutput {
	return o.ToDataControllerPropertiesResponsePtrOutputWithContext(context.Background())
}

func (o DataControllerPropertiesResponseOutput) ToDataControllerPropertiesResponsePtrOutputWithContext(ctx context.Context) DataControllerPropertiesResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DataControllerPropertiesResponse) *DataControllerPropertiesResponse {
		return &v
	}).(DataControllerPropertiesResponsePtrOutput)
}

func (o DataControllerPropertiesResponseOutput) BasicLoginInformation() BasicLoginInformationResponsePtrOutput {
	return o.ApplyT(func(v DataControllerPropertiesResponse) *BasicLoginInformationResponse {
		return v.BasicLoginInformation
	}).(BasicLoginInformationResponsePtrOutput)
}

func (o DataControllerPropertiesResponseOutput) K8sRaw() pulumi.AnyOutput {
	return o.ApplyT(func(v DataControllerPropertiesResponse) interface{} { return v.K8sRaw }).(pulumi.AnyOutput)
}

func (o DataControllerPropertiesResponseOutput) LastUploadedDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DataControllerPropertiesResponse) *string { return v.LastUploadedDate }).(pulumi.StringPtrOutput)
}

func (o DataControllerPropertiesResponseOutput) LogAnalyticsWorkspaceConfig() LogAnalyticsWorkspaceConfigResponsePtrOutput {
	return o.ApplyT(func(v DataControllerPropertiesResponse) *LogAnalyticsWorkspaceConfigResponse {
		return v.LogAnalyticsWorkspaceConfig
	}).(LogAnalyticsWorkspaceConfigResponsePtrOutput)
}

func (o DataControllerPropertiesResponseOutput) OnPremiseProperty() OnPremisePropertyResponsePtrOutput {
	return o.ApplyT(func(v DataControllerPropertiesResponse) *OnPremisePropertyResponse { return v.OnPremiseProperty }).(OnPremisePropertyResponsePtrOutput)
}

func (o DataControllerPropertiesResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v DataControllerPropertiesResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o DataControllerPropertiesResponseOutput) UploadServicePrincipal() UploadServicePrincipalResponsePtrOutput {
	return o.ApplyT(func(v DataControllerPropertiesResponse) *UploadServicePrincipalResponse {
		return v.UploadServicePrincipal
	}).(UploadServicePrincipalResponsePtrOutput)
}

func (o DataControllerPropertiesResponseOutput) UploadWatermark() UploadWatermarkResponsePtrOutput {
	return o.ApplyT(func(v DataControllerPropertiesResponse) *UploadWatermarkResponse { return v.UploadWatermark }).(UploadWatermarkResponsePtrOutput)
}

type DataControllerPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (DataControllerPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DataControllerPropertiesResponse)(nil)).Elem()
}

func (o DataControllerPropertiesResponsePtrOutput) ToDataControllerPropertiesResponsePtrOutput() DataControllerPropertiesResponsePtrOutput {
	return o
}

func (o DataControllerPropertiesResponsePtrOutput) ToDataControllerPropertiesResponsePtrOutputWithContext(ctx context.Context) DataControllerPropertiesResponsePtrOutput {
	return o
}

func (o DataControllerPropertiesResponsePtrOutput) Elem() DataControllerPropertiesResponseOutput {
	return o.ApplyT(func(v *DataControllerPropertiesResponse) DataControllerPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret DataControllerPropertiesResponse
		return ret
	}).(DataControllerPropertiesResponseOutput)
}

func (o DataControllerPropertiesResponsePtrOutput) BasicLoginInformation() BasicLoginInformationResponsePtrOutput {
	return o.ApplyT(func(v *DataControllerPropertiesResponse) *BasicLoginInformationResponse {
		if v == nil {
			return nil
		}
		return v.BasicLoginInformation
	}).(BasicLoginInformationResponsePtrOutput)
}

func (o DataControllerPropertiesResponsePtrOutput) K8sRaw() pulumi.AnyOutput {
	return o.ApplyT(func(v *DataControllerPropertiesResponse) interface{} {
		if v == nil {
			return nil
		}
		return v.K8sRaw
	}).(pulumi.AnyOutput)
}

func (o DataControllerPropertiesResponsePtrOutput) LastUploadedDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DataControllerPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.LastUploadedDate
	}).(pulumi.StringPtrOutput)
}

func (o DataControllerPropertiesResponsePtrOutput) LogAnalyticsWorkspaceConfig() LogAnalyticsWorkspaceConfigResponsePtrOutput {
	return o.ApplyT(func(v *DataControllerPropertiesResponse) *LogAnalyticsWorkspaceConfigResponse {
		if v == nil {
			return nil
		}
		return v.LogAnalyticsWorkspaceConfig
	}).(LogAnalyticsWorkspaceConfigResponsePtrOutput)
}

func (o DataControllerPropertiesResponsePtrOutput) OnPremiseProperty() OnPremisePropertyResponsePtrOutput {
	return o.ApplyT(func(v *DataControllerPropertiesResponse) *OnPremisePropertyResponse {
		if v == nil {
			return nil
		}
		return v.OnPremiseProperty
	}).(OnPremisePropertyResponsePtrOutput)
}

func (o DataControllerPropertiesResponsePtrOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DataControllerPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.ProvisioningState
	}).(pulumi.StringPtrOutput)
}

func (o DataControllerPropertiesResponsePtrOutput) UploadServicePrincipal() UploadServicePrincipalResponsePtrOutput {
	return o.ApplyT(func(v *DataControllerPropertiesResponse) *UploadServicePrincipalResponse {
		if v == nil {
			return nil
		}
		return v.UploadServicePrincipal
	}).(UploadServicePrincipalResponsePtrOutput)
}

func (o DataControllerPropertiesResponsePtrOutput) UploadWatermark() UploadWatermarkResponsePtrOutput {
	return o.ApplyT(func(v *DataControllerPropertiesResponse) *UploadWatermarkResponse {
		if v == nil {
			return nil
		}
		return v.UploadWatermark
	}).(UploadWatermarkResponsePtrOutput)
}

type ExtendedLocation struct {
	Name *string `pulumi:"name"`
	Type *string `pulumi:"type"`
}





type ExtendedLocationInput interface {
	pulumi.Input

	ToExtendedLocationOutput() ExtendedLocationOutput
	ToExtendedLocationOutputWithContext(context.Context) ExtendedLocationOutput
}

type ExtendedLocationArgs struct {
	Name pulumi.StringPtrInput `pulumi:"name"`
	Type pulumi.StringPtrInput `pulumi:"type"`
}

func (ExtendedLocationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ExtendedLocation)(nil)).Elem()
}

func (i ExtendedLocationArgs) ToExtendedLocationOutput() ExtendedLocationOutput {
	return i.ToExtendedLocationOutputWithContext(context.Background())
}

func (i ExtendedLocationArgs) ToExtendedLocationOutputWithContext(ctx context.Context) ExtendedLocationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExtendedLocationOutput)
}

func (i ExtendedLocationArgs) ToExtendedLocationPtrOutput() ExtendedLocationPtrOutput {
	return i.ToExtendedLocationPtrOutputWithContext(context.Background())
}

func (i ExtendedLocationArgs) ToExtendedLocationPtrOutputWithContext(ctx context.Context) ExtendedLocationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExtendedLocationOutput).ToExtendedLocationPtrOutputWithContext(ctx)
}









type ExtendedLocationPtrInput interface {
	pulumi.Input

	ToExtendedLocationPtrOutput() ExtendedLocationPtrOutput
	ToExtendedLocationPtrOutputWithContext(context.Context) ExtendedLocationPtrOutput
}

type extendedLocationPtrType ExtendedLocationArgs

func ExtendedLocationPtr(v *ExtendedLocationArgs) ExtendedLocationPtrInput {
	return (*extendedLocationPtrType)(v)
}

func (*extendedLocationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ExtendedLocation)(nil)).Elem()
}

func (i *extendedLocationPtrType) ToExtendedLocationPtrOutput() ExtendedLocationPtrOutput {
	return i.ToExtendedLocationPtrOutputWithContext(context.Background())
}

func (i *extendedLocationPtrType) ToExtendedLocationPtrOutputWithContext(ctx context.Context) ExtendedLocationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExtendedLocationPtrOutput)
}

type ExtendedLocationOutput struct{ *pulumi.OutputState }

func (ExtendedLocationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ExtendedLocation)(nil)).Elem()
}

func (o ExtendedLocationOutput) ToExtendedLocationOutput() ExtendedLocationOutput {
	return o
}

func (o ExtendedLocationOutput) ToExtendedLocationOutputWithContext(ctx context.Context) ExtendedLocationOutput {
	return o
}

func (o ExtendedLocationOutput) ToExtendedLocationPtrOutput() ExtendedLocationPtrOutput {
	return o.ToExtendedLocationPtrOutputWithContext(context.Background())
}

func (o ExtendedLocationOutput) ToExtendedLocationPtrOutputWithContext(ctx context.Context) ExtendedLocationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ExtendedLocation) *ExtendedLocation {
		return &v
	}).(ExtendedLocationPtrOutput)
}

func (o ExtendedLocationOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExtendedLocation) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ExtendedLocationOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExtendedLocation) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type ExtendedLocationPtrOutput struct{ *pulumi.OutputState }

func (ExtendedLocationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ExtendedLocation)(nil)).Elem()
}

func (o ExtendedLocationPtrOutput) ToExtendedLocationPtrOutput() ExtendedLocationPtrOutput {
	return o
}

func (o ExtendedLocationPtrOutput) ToExtendedLocationPtrOutputWithContext(ctx context.Context) ExtendedLocationPtrOutput {
	return o
}

func (o ExtendedLocationPtrOutput) Elem() ExtendedLocationOutput {
	return o.ApplyT(func(v *ExtendedLocation) ExtendedLocation {
		if v != nil {
			return *v
		}
		var ret ExtendedLocation
		return ret
	}).(ExtendedLocationOutput)
}

func (o ExtendedLocationPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ExtendedLocation) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o ExtendedLocationPtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ExtendedLocation) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

type ExtendedLocationResponse struct {
	Name *string `pulumi:"name"`
	Type *string `pulumi:"type"`
}





type ExtendedLocationResponseInput interface {
	pulumi.Input

	ToExtendedLocationResponseOutput() ExtendedLocationResponseOutput
	ToExtendedLocationResponseOutputWithContext(context.Context) ExtendedLocationResponseOutput
}

type ExtendedLocationResponseArgs struct {
	Name pulumi.StringPtrInput `pulumi:"name"`
	Type pulumi.StringPtrInput `pulumi:"type"`
}

func (ExtendedLocationResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ExtendedLocationResponse)(nil)).Elem()
}

func (i ExtendedLocationResponseArgs) ToExtendedLocationResponseOutput() ExtendedLocationResponseOutput {
	return i.ToExtendedLocationResponseOutputWithContext(context.Background())
}

func (i ExtendedLocationResponseArgs) ToExtendedLocationResponseOutputWithContext(ctx context.Context) ExtendedLocationResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExtendedLocationResponseOutput)
}

func (i ExtendedLocationResponseArgs) ToExtendedLocationResponsePtrOutput() ExtendedLocationResponsePtrOutput {
	return i.ToExtendedLocationResponsePtrOutputWithContext(context.Background())
}

func (i ExtendedLocationResponseArgs) ToExtendedLocationResponsePtrOutputWithContext(ctx context.Context) ExtendedLocationResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExtendedLocationResponseOutput).ToExtendedLocationResponsePtrOutputWithContext(ctx)
}









type ExtendedLocationResponsePtrInput interface {
	pulumi.Input

	ToExtendedLocationResponsePtrOutput() ExtendedLocationResponsePtrOutput
	ToExtendedLocationResponsePtrOutputWithContext(context.Context) ExtendedLocationResponsePtrOutput
}

type extendedLocationResponsePtrType ExtendedLocationResponseArgs

func ExtendedLocationResponsePtr(v *ExtendedLocationResponseArgs) ExtendedLocationResponsePtrInput {
	return (*extendedLocationResponsePtrType)(v)
}

func (*extendedLocationResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ExtendedLocationResponse)(nil)).Elem()
}

func (i *extendedLocationResponsePtrType) ToExtendedLocationResponsePtrOutput() ExtendedLocationResponsePtrOutput {
	return i.ToExtendedLocationResponsePtrOutputWithContext(context.Background())
}

func (i *extendedLocationResponsePtrType) ToExtendedLocationResponsePtrOutputWithContext(ctx context.Context) ExtendedLocationResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExtendedLocationResponsePtrOutput)
}

type ExtendedLocationResponseOutput struct{ *pulumi.OutputState }

func (ExtendedLocationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ExtendedLocationResponse)(nil)).Elem()
}

func (o ExtendedLocationResponseOutput) ToExtendedLocationResponseOutput() ExtendedLocationResponseOutput {
	return o
}

func (o ExtendedLocationResponseOutput) ToExtendedLocationResponseOutputWithContext(ctx context.Context) ExtendedLocationResponseOutput {
	return o
}

func (o ExtendedLocationResponseOutput) ToExtendedLocationResponsePtrOutput() ExtendedLocationResponsePtrOutput {
	return o.ToExtendedLocationResponsePtrOutputWithContext(context.Background())
}

func (o ExtendedLocationResponseOutput) ToExtendedLocationResponsePtrOutputWithContext(ctx context.Context) ExtendedLocationResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ExtendedLocationResponse) *ExtendedLocationResponse {
		return &v
	}).(ExtendedLocationResponsePtrOutput)
}

func (o ExtendedLocationResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExtendedLocationResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ExtendedLocationResponseOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExtendedLocationResponse) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type ExtendedLocationResponsePtrOutput struct{ *pulumi.OutputState }

func (ExtendedLocationResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ExtendedLocationResponse)(nil)).Elem()
}

func (o ExtendedLocationResponsePtrOutput) ToExtendedLocationResponsePtrOutput() ExtendedLocationResponsePtrOutput {
	return o
}

func (o ExtendedLocationResponsePtrOutput) ToExtendedLocationResponsePtrOutputWithContext(ctx context.Context) ExtendedLocationResponsePtrOutput {
	return o
}

func (o ExtendedLocationResponsePtrOutput) Elem() ExtendedLocationResponseOutput {
	return o.ApplyT(func(v *ExtendedLocationResponse) ExtendedLocationResponse {
		if v != nil {
			return *v
		}
		var ret ExtendedLocationResponse
		return ret
	}).(ExtendedLocationResponseOutput)
}

func (o ExtendedLocationResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ExtendedLocationResponse) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o ExtendedLocationResponsePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ExtendedLocationResponse) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

type LogAnalyticsWorkspaceConfig struct {
	PrimaryKey  *string `pulumi:"primaryKey"`
	WorkspaceId *string `pulumi:"workspaceId"`
}





type LogAnalyticsWorkspaceConfigInput interface {
	pulumi.Input

	ToLogAnalyticsWorkspaceConfigOutput() LogAnalyticsWorkspaceConfigOutput
	ToLogAnalyticsWorkspaceConfigOutputWithContext(context.Context) LogAnalyticsWorkspaceConfigOutput
}

type LogAnalyticsWorkspaceConfigArgs struct {
	PrimaryKey  pulumi.StringPtrInput `pulumi:"primaryKey"`
	WorkspaceId pulumi.StringPtrInput `pulumi:"workspaceId"`
}

func (LogAnalyticsWorkspaceConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LogAnalyticsWorkspaceConfig)(nil)).Elem()
}

func (i LogAnalyticsWorkspaceConfigArgs) ToLogAnalyticsWorkspaceConfigOutput() LogAnalyticsWorkspaceConfigOutput {
	return i.ToLogAnalyticsWorkspaceConfigOutputWithContext(context.Background())
}

func (i LogAnalyticsWorkspaceConfigArgs) ToLogAnalyticsWorkspaceConfigOutputWithContext(ctx context.Context) LogAnalyticsWorkspaceConfigOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogAnalyticsWorkspaceConfigOutput)
}

func (i LogAnalyticsWorkspaceConfigArgs) ToLogAnalyticsWorkspaceConfigPtrOutput() LogAnalyticsWorkspaceConfigPtrOutput {
	return i.ToLogAnalyticsWorkspaceConfigPtrOutputWithContext(context.Background())
}

func (i LogAnalyticsWorkspaceConfigArgs) ToLogAnalyticsWorkspaceConfigPtrOutputWithContext(ctx context.Context) LogAnalyticsWorkspaceConfigPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogAnalyticsWorkspaceConfigOutput).ToLogAnalyticsWorkspaceConfigPtrOutputWithContext(ctx)
}









type LogAnalyticsWorkspaceConfigPtrInput interface {
	pulumi.Input

	ToLogAnalyticsWorkspaceConfigPtrOutput() LogAnalyticsWorkspaceConfigPtrOutput
	ToLogAnalyticsWorkspaceConfigPtrOutputWithContext(context.Context) LogAnalyticsWorkspaceConfigPtrOutput
}

type logAnalyticsWorkspaceConfigPtrType LogAnalyticsWorkspaceConfigArgs

func LogAnalyticsWorkspaceConfigPtr(v *LogAnalyticsWorkspaceConfigArgs) LogAnalyticsWorkspaceConfigPtrInput {
	return (*logAnalyticsWorkspaceConfigPtrType)(v)
}

func (*logAnalyticsWorkspaceConfigPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**LogAnalyticsWorkspaceConfig)(nil)).Elem()
}

func (i *logAnalyticsWorkspaceConfigPtrType) ToLogAnalyticsWorkspaceConfigPtrOutput() LogAnalyticsWorkspaceConfigPtrOutput {
	return i.ToLogAnalyticsWorkspaceConfigPtrOutputWithContext(context.Background())
}

func (i *logAnalyticsWorkspaceConfigPtrType) ToLogAnalyticsWorkspaceConfigPtrOutputWithContext(ctx context.Context) LogAnalyticsWorkspaceConfigPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogAnalyticsWorkspaceConfigPtrOutput)
}

type LogAnalyticsWorkspaceConfigOutput struct{ *pulumi.OutputState }

func (LogAnalyticsWorkspaceConfigOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LogAnalyticsWorkspaceConfig)(nil)).Elem()
}

func (o LogAnalyticsWorkspaceConfigOutput) ToLogAnalyticsWorkspaceConfigOutput() LogAnalyticsWorkspaceConfigOutput {
	return o
}

func (o LogAnalyticsWorkspaceConfigOutput) ToLogAnalyticsWorkspaceConfigOutputWithContext(ctx context.Context) LogAnalyticsWorkspaceConfigOutput {
	return o
}

func (o LogAnalyticsWorkspaceConfigOutput) ToLogAnalyticsWorkspaceConfigPtrOutput() LogAnalyticsWorkspaceConfigPtrOutput {
	return o.ToLogAnalyticsWorkspaceConfigPtrOutputWithContext(context.Background())
}

func (o LogAnalyticsWorkspaceConfigOutput) ToLogAnalyticsWorkspaceConfigPtrOutputWithContext(ctx context.Context) LogAnalyticsWorkspaceConfigPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v LogAnalyticsWorkspaceConfig) *LogAnalyticsWorkspaceConfig {
		return &v
	}).(LogAnalyticsWorkspaceConfigPtrOutput)
}

func (o LogAnalyticsWorkspaceConfigOutput) PrimaryKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LogAnalyticsWorkspaceConfig) *string { return v.PrimaryKey }).(pulumi.StringPtrOutput)
}

func (o LogAnalyticsWorkspaceConfigOutput) WorkspaceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LogAnalyticsWorkspaceConfig) *string { return v.WorkspaceId }).(pulumi.StringPtrOutput)
}

type LogAnalyticsWorkspaceConfigPtrOutput struct{ *pulumi.OutputState }

func (LogAnalyticsWorkspaceConfigPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LogAnalyticsWorkspaceConfig)(nil)).Elem()
}

func (o LogAnalyticsWorkspaceConfigPtrOutput) ToLogAnalyticsWorkspaceConfigPtrOutput() LogAnalyticsWorkspaceConfigPtrOutput {
	return o
}

func (o LogAnalyticsWorkspaceConfigPtrOutput) ToLogAnalyticsWorkspaceConfigPtrOutputWithContext(ctx context.Context) LogAnalyticsWorkspaceConfigPtrOutput {
	return o
}

func (o LogAnalyticsWorkspaceConfigPtrOutput) Elem() LogAnalyticsWorkspaceConfigOutput {
	return o.ApplyT(func(v *LogAnalyticsWorkspaceConfig) LogAnalyticsWorkspaceConfig {
		if v != nil {
			return *v
		}
		var ret LogAnalyticsWorkspaceConfig
		return ret
	}).(LogAnalyticsWorkspaceConfigOutput)
}

func (o LogAnalyticsWorkspaceConfigPtrOutput) PrimaryKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LogAnalyticsWorkspaceConfig) *string {
		if v == nil {
			return nil
		}
		return v.PrimaryKey
	}).(pulumi.StringPtrOutput)
}

func (o LogAnalyticsWorkspaceConfigPtrOutput) WorkspaceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LogAnalyticsWorkspaceConfig) *string {
		if v == nil {
			return nil
		}
		return v.WorkspaceId
	}).(pulumi.StringPtrOutput)
}

type LogAnalyticsWorkspaceConfigResponse struct {
	WorkspaceId *string `pulumi:"workspaceId"`
}





type LogAnalyticsWorkspaceConfigResponseInput interface {
	pulumi.Input

	ToLogAnalyticsWorkspaceConfigResponseOutput() LogAnalyticsWorkspaceConfigResponseOutput
	ToLogAnalyticsWorkspaceConfigResponseOutputWithContext(context.Context) LogAnalyticsWorkspaceConfigResponseOutput
}

type LogAnalyticsWorkspaceConfigResponseArgs struct {
	WorkspaceId pulumi.StringPtrInput `pulumi:"workspaceId"`
}

func (LogAnalyticsWorkspaceConfigResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LogAnalyticsWorkspaceConfigResponse)(nil)).Elem()
}

func (i LogAnalyticsWorkspaceConfigResponseArgs) ToLogAnalyticsWorkspaceConfigResponseOutput() LogAnalyticsWorkspaceConfigResponseOutput {
	return i.ToLogAnalyticsWorkspaceConfigResponseOutputWithContext(context.Background())
}

func (i LogAnalyticsWorkspaceConfigResponseArgs) ToLogAnalyticsWorkspaceConfigResponseOutputWithContext(ctx context.Context) LogAnalyticsWorkspaceConfigResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogAnalyticsWorkspaceConfigResponseOutput)
}

func (i LogAnalyticsWorkspaceConfigResponseArgs) ToLogAnalyticsWorkspaceConfigResponsePtrOutput() LogAnalyticsWorkspaceConfigResponsePtrOutput {
	return i.ToLogAnalyticsWorkspaceConfigResponsePtrOutputWithContext(context.Background())
}

func (i LogAnalyticsWorkspaceConfigResponseArgs) ToLogAnalyticsWorkspaceConfigResponsePtrOutputWithContext(ctx context.Context) LogAnalyticsWorkspaceConfigResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogAnalyticsWorkspaceConfigResponseOutput).ToLogAnalyticsWorkspaceConfigResponsePtrOutputWithContext(ctx)
}









type LogAnalyticsWorkspaceConfigResponsePtrInput interface {
	pulumi.Input

	ToLogAnalyticsWorkspaceConfigResponsePtrOutput() LogAnalyticsWorkspaceConfigResponsePtrOutput
	ToLogAnalyticsWorkspaceConfigResponsePtrOutputWithContext(context.Context) LogAnalyticsWorkspaceConfigResponsePtrOutput
}

type logAnalyticsWorkspaceConfigResponsePtrType LogAnalyticsWorkspaceConfigResponseArgs

func LogAnalyticsWorkspaceConfigResponsePtr(v *LogAnalyticsWorkspaceConfigResponseArgs) LogAnalyticsWorkspaceConfigResponsePtrInput {
	return (*logAnalyticsWorkspaceConfigResponsePtrType)(v)
}

func (*logAnalyticsWorkspaceConfigResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**LogAnalyticsWorkspaceConfigResponse)(nil)).Elem()
}

func (i *logAnalyticsWorkspaceConfigResponsePtrType) ToLogAnalyticsWorkspaceConfigResponsePtrOutput() LogAnalyticsWorkspaceConfigResponsePtrOutput {
	return i.ToLogAnalyticsWorkspaceConfigResponsePtrOutputWithContext(context.Background())
}

func (i *logAnalyticsWorkspaceConfigResponsePtrType) ToLogAnalyticsWorkspaceConfigResponsePtrOutputWithContext(ctx context.Context) LogAnalyticsWorkspaceConfigResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogAnalyticsWorkspaceConfigResponsePtrOutput)
}

type LogAnalyticsWorkspaceConfigResponseOutput struct{ *pulumi.OutputState }

func (LogAnalyticsWorkspaceConfigResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LogAnalyticsWorkspaceConfigResponse)(nil)).Elem()
}

func (o LogAnalyticsWorkspaceConfigResponseOutput) ToLogAnalyticsWorkspaceConfigResponseOutput() LogAnalyticsWorkspaceConfigResponseOutput {
	return o
}

func (o LogAnalyticsWorkspaceConfigResponseOutput) ToLogAnalyticsWorkspaceConfigResponseOutputWithContext(ctx context.Context) LogAnalyticsWorkspaceConfigResponseOutput {
	return o
}

func (o LogAnalyticsWorkspaceConfigResponseOutput) ToLogAnalyticsWorkspaceConfigResponsePtrOutput() LogAnalyticsWorkspaceConfigResponsePtrOutput {
	return o.ToLogAnalyticsWorkspaceConfigResponsePtrOutputWithContext(context.Background())
}

func (o LogAnalyticsWorkspaceConfigResponseOutput) ToLogAnalyticsWorkspaceConfigResponsePtrOutputWithContext(ctx context.Context) LogAnalyticsWorkspaceConfigResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v LogAnalyticsWorkspaceConfigResponse) *LogAnalyticsWorkspaceConfigResponse {
		return &v
	}).(LogAnalyticsWorkspaceConfigResponsePtrOutput)
}

func (o LogAnalyticsWorkspaceConfigResponseOutput) WorkspaceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LogAnalyticsWorkspaceConfigResponse) *string { return v.WorkspaceId }).(pulumi.StringPtrOutput)
}

type LogAnalyticsWorkspaceConfigResponsePtrOutput struct{ *pulumi.OutputState }

func (LogAnalyticsWorkspaceConfigResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LogAnalyticsWorkspaceConfigResponse)(nil)).Elem()
}

func (o LogAnalyticsWorkspaceConfigResponsePtrOutput) ToLogAnalyticsWorkspaceConfigResponsePtrOutput() LogAnalyticsWorkspaceConfigResponsePtrOutput {
	return o
}

func (o LogAnalyticsWorkspaceConfigResponsePtrOutput) ToLogAnalyticsWorkspaceConfigResponsePtrOutputWithContext(ctx context.Context) LogAnalyticsWorkspaceConfigResponsePtrOutput {
	return o
}

func (o LogAnalyticsWorkspaceConfigResponsePtrOutput) Elem() LogAnalyticsWorkspaceConfigResponseOutput {
	return o.ApplyT(func(v *LogAnalyticsWorkspaceConfigResponse) LogAnalyticsWorkspaceConfigResponse {
		if v != nil {
			return *v
		}
		var ret LogAnalyticsWorkspaceConfigResponse
		return ret
	}).(LogAnalyticsWorkspaceConfigResponseOutput)
}

func (o LogAnalyticsWorkspaceConfigResponsePtrOutput) WorkspaceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LogAnalyticsWorkspaceConfigResponse) *string {
		if v == nil {
			return nil
		}
		return v.WorkspaceId
	}).(pulumi.StringPtrOutput)
}

type OnPremiseProperty struct {
	Id                           string  `pulumi:"id"`
	PublicSigningKey             string  `pulumi:"publicSigningKey"`
	SigningCertificateThumbprint *string `pulumi:"signingCertificateThumbprint"`
}





type OnPremisePropertyInput interface {
	pulumi.Input

	ToOnPremisePropertyOutput() OnPremisePropertyOutput
	ToOnPremisePropertyOutputWithContext(context.Context) OnPremisePropertyOutput
}

type OnPremisePropertyArgs struct {
	Id                           pulumi.StringInput    `pulumi:"id"`
	PublicSigningKey             pulumi.StringInput    `pulumi:"publicSigningKey"`
	SigningCertificateThumbprint pulumi.StringPtrInput `pulumi:"signingCertificateThumbprint"`
}

func (OnPremisePropertyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*OnPremiseProperty)(nil)).Elem()
}

func (i OnPremisePropertyArgs) ToOnPremisePropertyOutput() OnPremisePropertyOutput {
	return i.ToOnPremisePropertyOutputWithContext(context.Background())
}

func (i OnPremisePropertyArgs) ToOnPremisePropertyOutputWithContext(ctx context.Context) OnPremisePropertyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OnPremisePropertyOutput)
}

func (i OnPremisePropertyArgs) ToOnPremisePropertyPtrOutput() OnPremisePropertyPtrOutput {
	return i.ToOnPremisePropertyPtrOutputWithContext(context.Background())
}

func (i OnPremisePropertyArgs) ToOnPremisePropertyPtrOutputWithContext(ctx context.Context) OnPremisePropertyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OnPremisePropertyOutput).ToOnPremisePropertyPtrOutputWithContext(ctx)
}









type OnPremisePropertyPtrInput interface {
	pulumi.Input

	ToOnPremisePropertyPtrOutput() OnPremisePropertyPtrOutput
	ToOnPremisePropertyPtrOutputWithContext(context.Context) OnPremisePropertyPtrOutput
}

type onPremisePropertyPtrType OnPremisePropertyArgs

func OnPremisePropertyPtr(v *OnPremisePropertyArgs) OnPremisePropertyPtrInput {
	return (*onPremisePropertyPtrType)(v)
}

func (*onPremisePropertyPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**OnPremiseProperty)(nil)).Elem()
}

func (i *onPremisePropertyPtrType) ToOnPremisePropertyPtrOutput() OnPremisePropertyPtrOutput {
	return i.ToOnPremisePropertyPtrOutputWithContext(context.Background())
}

func (i *onPremisePropertyPtrType) ToOnPremisePropertyPtrOutputWithContext(ctx context.Context) OnPremisePropertyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OnPremisePropertyPtrOutput)
}

type OnPremisePropertyOutput struct{ *pulumi.OutputState }

func (OnPremisePropertyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*OnPremiseProperty)(nil)).Elem()
}

func (o OnPremisePropertyOutput) ToOnPremisePropertyOutput() OnPremisePropertyOutput {
	return o
}

func (o OnPremisePropertyOutput) ToOnPremisePropertyOutputWithContext(ctx context.Context) OnPremisePropertyOutput {
	return o
}

func (o OnPremisePropertyOutput) ToOnPremisePropertyPtrOutput() OnPremisePropertyPtrOutput {
	return o.ToOnPremisePropertyPtrOutputWithContext(context.Background())
}

func (o OnPremisePropertyOutput) ToOnPremisePropertyPtrOutputWithContext(ctx context.Context) OnPremisePropertyPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v OnPremiseProperty) *OnPremiseProperty {
		return &v
	}).(OnPremisePropertyPtrOutput)
}

func (o OnPremisePropertyOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v OnPremiseProperty) string { return v.Id }).(pulumi.StringOutput)
}

func (o OnPremisePropertyOutput) PublicSigningKey() pulumi.StringOutput {
	return o.ApplyT(func(v OnPremiseProperty) string { return v.PublicSigningKey }).(pulumi.StringOutput)
}

func (o OnPremisePropertyOutput) SigningCertificateThumbprint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v OnPremiseProperty) *string { return v.SigningCertificateThumbprint }).(pulumi.StringPtrOutput)
}

type OnPremisePropertyPtrOutput struct{ *pulumi.OutputState }

func (OnPremisePropertyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OnPremiseProperty)(nil)).Elem()
}

func (o OnPremisePropertyPtrOutput) ToOnPremisePropertyPtrOutput() OnPremisePropertyPtrOutput {
	return o
}

func (o OnPremisePropertyPtrOutput) ToOnPremisePropertyPtrOutputWithContext(ctx context.Context) OnPremisePropertyPtrOutput {
	return o
}

func (o OnPremisePropertyPtrOutput) Elem() OnPremisePropertyOutput {
	return o.ApplyT(func(v *OnPremiseProperty) OnPremiseProperty {
		if v != nil {
			return *v
		}
		var ret OnPremiseProperty
		return ret
	}).(OnPremisePropertyOutput)
}

func (o OnPremisePropertyPtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OnPremiseProperty) *string {
		if v == nil {
			return nil
		}
		return &v.Id
	}).(pulumi.StringPtrOutput)
}

func (o OnPremisePropertyPtrOutput) PublicSigningKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OnPremiseProperty) *string {
		if v == nil {
			return nil
		}
		return &v.PublicSigningKey
	}).(pulumi.StringPtrOutput)
}

func (o OnPremisePropertyPtrOutput) SigningCertificateThumbprint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OnPremiseProperty) *string {
		if v == nil {
			return nil
		}
		return v.SigningCertificateThumbprint
	}).(pulumi.StringPtrOutput)
}

type OnPremisePropertyResponse struct {
	Id                           string  `pulumi:"id"`
	PublicSigningKey             string  `pulumi:"publicSigningKey"`
	SigningCertificateThumbprint *string `pulumi:"signingCertificateThumbprint"`
}





type OnPremisePropertyResponseInput interface {
	pulumi.Input

	ToOnPremisePropertyResponseOutput() OnPremisePropertyResponseOutput
	ToOnPremisePropertyResponseOutputWithContext(context.Context) OnPremisePropertyResponseOutput
}

type OnPremisePropertyResponseArgs struct {
	Id                           pulumi.StringInput    `pulumi:"id"`
	PublicSigningKey             pulumi.StringInput    `pulumi:"publicSigningKey"`
	SigningCertificateThumbprint pulumi.StringPtrInput `pulumi:"signingCertificateThumbprint"`
}

func (OnPremisePropertyResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*OnPremisePropertyResponse)(nil)).Elem()
}

func (i OnPremisePropertyResponseArgs) ToOnPremisePropertyResponseOutput() OnPremisePropertyResponseOutput {
	return i.ToOnPremisePropertyResponseOutputWithContext(context.Background())
}

func (i OnPremisePropertyResponseArgs) ToOnPremisePropertyResponseOutputWithContext(ctx context.Context) OnPremisePropertyResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OnPremisePropertyResponseOutput)
}

func (i OnPremisePropertyResponseArgs) ToOnPremisePropertyResponsePtrOutput() OnPremisePropertyResponsePtrOutput {
	return i.ToOnPremisePropertyResponsePtrOutputWithContext(context.Background())
}

func (i OnPremisePropertyResponseArgs) ToOnPremisePropertyResponsePtrOutputWithContext(ctx context.Context) OnPremisePropertyResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OnPremisePropertyResponseOutput).ToOnPremisePropertyResponsePtrOutputWithContext(ctx)
}









type OnPremisePropertyResponsePtrInput interface {
	pulumi.Input

	ToOnPremisePropertyResponsePtrOutput() OnPremisePropertyResponsePtrOutput
	ToOnPremisePropertyResponsePtrOutputWithContext(context.Context) OnPremisePropertyResponsePtrOutput
}

type onPremisePropertyResponsePtrType OnPremisePropertyResponseArgs

func OnPremisePropertyResponsePtr(v *OnPremisePropertyResponseArgs) OnPremisePropertyResponsePtrInput {
	return (*onPremisePropertyResponsePtrType)(v)
}

func (*onPremisePropertyResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**OnPremisePropertyResponse)(nil)).Elem()
}

func (i *onPremisePropertyResponsePtrType) ToOnPremisePropertyResponsePtrOutput() OnPremisePropertyResponsePtrOutput {
	return i.ToOnPremisePropertyResponsePtrOutputWithContext(context.Background())
}

func (i *onPremisePropertyResponsePtrType) ToOnPremisePropertyResponsePtrOutputWithContext(ctx context.Context) OnPremisePropertyResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OnPremisePropertyResponsePtrOutput)
}

type OnPremisePropertyResponseOutput struct{ *pulumi.OutputState }

func (OnPremisePropertyResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*OnPremisePropertyResponse)(nil)).Elem()
}

func (o OnPremisePropertyResponseOutput) ToOnPremisePropertyResponseOutput() OnPremisePropertyResponseOutput {
	return o
}

func (o OnPremisePropertyResponseOutput) ToOnPremisePropertyResponseOutputWithContext(ctx context.Context) OnPremisePropertyResponseOutput {
	return o
}

func (o OnPremisePropertyResponseOutput) ToOnPremisePropertyResponsePtrOutput() OnPremisePropertyResponsePtrOutput {
	return o.ToOnPremisePropertyResponsePtrOutputWithContext(context.Background())
}

func (o OnPremisePropertyResponseOutput) ToOnPremisePropertyResponsePtrOutputWithContext(ctx context.Context) OnPremisePropertyResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v OnPremisePropertyResponse) *OnPremisePropertyResponse {
		return &v
	}).(OnPremisePropertyResponsePtrOutput)
}

func (o OnPremisePropertyResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v OnPremisePropertyResponse) string { return v.Id }).(pulumi.StringOutput)
}

func (o OnPremisePropertyResponseOutput) PublicSigningKey() pulumi.StringOutput {
	return o.ApplyT(func(v OnPremisePropertyResponse) string { return v.PublicSigningKey }).(pulumi.StringOutput)
}

func (o OnPremisePropertyResponseOutput) SigningCertificateThumbprint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v OnPremisePropertyResponse) *string { return v.SigningCertificateThumbprint }).(pulumi.StringPtrOutput)
}

type OnPremisePropertyResponsePtrOutput struct{ *pulumi.OutputState }

func (OnPremisePropertyResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OnPremisePropertyResponse)(nil)).Elem()
}

func (o OnPremisePropertyResponsePtrOutput) ToOnPremisePropertyResponsePtrOutput() OnPremisePropertyResponsePtrOutput {
	return o
}

func (o OnPremisePropertyResponsePtrOutput) ToOnPremisePropertyResponsePtrOutputWithContext(ctx context.Context) OnPremisePropertyResponsePtrOutput {
	return o
}

func (o OnPremisePropertyResponsePtrOutput) Elem() OnPremisePropertyResponseOutput {
	return o.ApplyT(func(v *OnPremisePropertyResponse) OnPremisePropertyResponse {
		if v != nil {
			return *v
		}
		var ret OnPremisePropertyResponse
		return ret
	}).(OnPremisePropertyResponseOutput)
}

func (o OnPremisePropertyResponsePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OnPremisePropertyResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Id
	}).(pulumi.StringPtrOutput)
}

func (o OnPremisePropertyResponsePtrOutput) PublicSigningKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OnPremisePropertyResponse) *string {
		if v == nil {
			return nil
		}
		return &v.PublicSigningKey
	}).(pulumi.StringPtrOutput)
}

func (o OnPremisePropertyResponsePtrOutput) SigningCertificateThumbprint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OnPremisePropertyResponse) *string {
		if v == nil {
			return nil
		}
		return v.SigningCertificateThumbprint
	}).(pulumi.StringPtrOutput)
}

type PostgresInstanceProperties struct {
	Admin                 *string                `pulumi:"admin"`
	BasicLoginInformation *BasicLoginInformation `pulumi:"basicLoginInformation"`
	DataControllerId      *string                `pulumi:"dataControllerId"`
	K8sRaw                interface{}            `pulumi:"k8sRaw"`
	LastUploadedDate      *string                `pulumi:"lastUploadedDate"`
}





type PostgresInstancePropertiesInput interface {
	pulumi.Input

	ToPostgresInstancePropertiesOutput() PostgresInstancePropertiesOutput
	ToPostgresInstancePropertiesOutputWithContext(context.Context) PostgresInstancePropertiesOutput
}

type PostgresInstancePropertiesArgs struct {
	Admin                 pulumi.StringPtrInput         `pulumi:"admin"`
	BasicLoginInformation BasicLoginInformationPtrInput `pulumi:"basicLoginInformation"`
	DataControllerId      pulumi.StringPtrInput         `pulumi:"dataControllerId"`
	K8sRaw                pulumi.Input                  `pulumi:"k8sRaw"`
	LastUploadedDate      pulumi.StringPtrInput         `pulumi:"lastUploadedDate"`
}

func (PostgresInstancePropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PostgresInstanceProperties)(nil)).Elem()
}

func (i PostgresInstancePropertiesArgs) ToPostgresInstancePropertiesOutput() PostgresInstancePropertiesOutput {
	return i.ToPostgresInstancePropertiesOutputWithContext(context.Background())
}

func (i PostgresInstancePropertiesArgs) ToPostgresInstancePropertiesOutputWithContext(ctx context.Context) PostgresInstancePropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PostgresInstancePropertiesOutput)
}

func (i PostgresInstancePropertiesArgs) ToPostgresInstancePropertiesPtrOutput() PostgresInstancePropertiesPtrOutput {
	return i.ToPostgresInstancePropertiesPtrOutputWithContext(context.Background())
}

func (i PostgresInstancePropertiesArgs) ToPostgresInstancePropertiesPtrOutputWithContext(ctx context.Context) PostgresInstancePropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PostgresInstancePropertiesOutput).ToPostgresInstancePropertiesPtrOutputWithContext(ctx)
}









type PostgresInstancePropertiesPtrInput interface {
	pulumi.Input

	ToPostgresInstancePropertiesPtrOutput() PostgresInstancePropertiesPtrOutput
	ToPostgresInstancePropertiesPtrOutputWithContext(context.Context) PostgresInstancePropertiesPtrOutput
}

type postgresInstancePropertiesPtrType PostgresInstancePropertiesArgs

func PostgresInstancePropertiesPtr(v *PostgresInstancePropertiesArgs) PostgresInstancePropertiesPtrInput {
	return (*postgresInstancePropertiesPtrType)(v)
}

func (*postgresInstancePropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PostgresInstanceProperties)(nil)).Elem()
}

func (i *postgresInstancePropertiesPtrType) ToPostgresInstancePropertiesPtrOutput() PostgresInstancePropertiesPtrOutput {
	return i.ToPostgresInstancePropertiesPtrOutputWithContext(context.Background())
}

func (i *postgresInstancePropertiesPtrType) ToPostgresInstancePropertiesPtrOutputWithContext(ctx context.Context) PostgresInstancePropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PostgresInstancePropertiesPtrOutput)
}

type PostgresInstancePropertiesOutput struct{ *pulumi.OutputState }

func (PostgresInstancePropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PostgresInstanceProperties)(nil)).Elem()
}

func (o PostgresInstancePropertiesOutput) ToPostgresInstancePropertiesOutput() PostgresInstancePropertiesOutput {
	return o
}

func (o PostgresInstancePropertiesOutput) ToPostgresInstancePropertiesOutputWithContext(ctx context.Context) PostgresInstancePropertiesOutput {
	return o
}

func (o PostgresInstancePropertiesOutput) ToPostgresInstancePropertiesPtrOutput() PostgresInstancePropertiesPtrOutput {
	return o.ToPostgresInstancePropertiesPtrOutputWithContext(context.Background())
}

func (o PostgresInstancePropertiesOutput) ToPostgresInstancePropertiesPtrOutputWithContext(ctx context.Context) PostgresInstancePropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PostgresInstanceProperties) *PostgresInstanceProperties {
		return &v
	}).(PostgresInstancePropertiesPtrOutput)
}

func (o PostgresInstancePropertiesOutput) Admin() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PostgresInstanceProperties) *string { return v.Admin }).(pulumi.StringPtrOutput)
}

func (o PostgresInstancePropertiesOutput) BasicLoginInformation() BasicLoginInformationPtrOutput {
	return o.ApplyT(func(v PostgresInstanceProperties) *BasicLoginInformation { return v.BasicLoginInformation }).(BasicLoginInformationPtrOutput)
}

func (o PostgresInstancePropertiesOutput) DataControllerId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PostgresInstanceProperties) *string { return v.DataControllerId }).(pulumi.StringPtrOutput)
}

func (o PostgresInstancePropertiesOutput) K8sRaw() pulumi.AnyOutput {
	return o.ApplyT(func(v PostgresInstanceProperties) interface{} { return v.K8sRaw }).(pulumi.AnyOutput)
}

func (o PostgresInstancePropertiesOutput) LastUploadedDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PostgresInstanceProperties) *string { return v.LastUploadedDate }).(pulumi.StringPtrOutput)
}

type PostgresInstancePropertiesPtrOutput struct{ *pulumi.OutputState }

func (PostgresInstancePropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PostgresInstanceProperties)(nil)).Elem()
}

func (o PostgresInstancePropertiesPtrOutput) ToPostgresInstancePropertiesPtrOutput() PostgresInstancePropertiesPtrOutput {
	return o
}

func (o PostgresInstancePropertiesPtrOutput) ToPostgresInstancePropertiesPtrOutputWithContext(ctx context.Context) PostgresInstancePropertiesPtrOutput {
	return o
}

func (o PostgresInstancePropertiesPtrOutput) Elem() PostgresInstancePropertiesOutput {
	return o.ApplyT(func(v *PostgresInstanceProperties) PostgresInstanceProperties {
		if v != nil {
			return *v
		}
		var ret PostgresInstanceProperties
		return ret
	}).(PostgresInstancePropertiesOutput)
}

func (o PostgresInstancePropertiesPtrOutput) Admin() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PostgresInstanceProperties) *string {
		if v == nil {
			return nil
		}
		return v.Admin
	}).(pulumi.StringPtrOutput)
}

func (o PostgresInstancePropertiesPtrOutput) BasicLoginInformation() BasicLoginInformationPtrOutput {
	return o.ApplyT(func(v *PostgresInstanceProperties) *BasicLoginInformation {
		if v == nil {
			return nil
		}
		return v.BasicLoginInformation
	}).(BasicLoginInformationPtrOutput)
}

func (o PostgresInstancePropertiesPtrOutput) DataControllerId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PostgresInstanceProperties) *string {
		if v == nil {
			return nil
		}
		return v.DataControllerId
	}).(pulumi.StringPtrOutput)
}

func (o PostgresInstancePropertiesPtrOutput) K8sRaw() pulumi.AnyOutput {
	return o.ApplyT(func(v *PostgresInstanceProperties) interface{} {
		if v == nil {
			return nil
		}
		return v.K8sRaw
	}).(pulumi.AnyOutput)
}

func (o PostgresInstancePropertiesPtrOutput) LastUploadedDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PostgresInstanceProperties) *string {
		if v == nil {
			return nil
		}
		return v.LastUploadedDate
	}).(pulumi.StringPtrOutput)
}

type PostgresInstancePropertiesResponse struct {
	Admin                 *string                        `pulumi:"admin"`
	BasicLoginInformation *BasicLoginInformationResponse `pulumi:"basicLoginInformation"`
	DataControllerId      *string                        `pulumi:"dataControllerId"`
	K8sRaw                interface{}                    `pulumi:"k8sRaw"`
	LastUploadedDate      *string                        `pulumi:"lastUploadedDate"`
	ProvisioningState     string                         `pulumi:"provisioningState"`
}





type PostgresInstancePropertiesResponseInput interface {
	pulumi.Input

	ToPostgresInstancePropertiesResponseOutput() PostgresInstancePropertiesResponseOutput
	ToPostgresInstancePropertiesResponseOutputWithContext(context.Context) PostgresInstancePropertiesResponseOutput
}

type PostgresInstancePropertiesResponseArgs struct {
	Admin                 pulumi.StringPtrInput                 `pulumi:"admin"`
	BasicLoginInformation BasicLoginInformationResponsePtrInput `pulumi:"basicLoginInformation"`
	DataControllerId      pulumi.StringPtrInput                 `pulumi:"dataControllerId"`
	K8sRaw                pulumi.Input                          `pulumi:"k8sRaw"`
	LastUploadedDate      pulumi.StringPtrInput                 `pulumi:"lastUploadedDate"`
	ProvisioningState     pulumi.StringInput                    `pulumi:"provisioningState"`
}

func (PostgresInstancePropertiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PostgresInstancePropertiesResponse)(nil)).Elem()
}

func (i PostgresInstancePropertiesResponseArgs) ToPostgresInstancePropertiesResponseOutput() PostgresInstancePropertiesResponseOutput {
	return i.ToPostgresInstancePropertiesResponseOutputWithContext(context.Background())
}

func (i PostgresInstancePropertiesResponseArgs) ToPostgresInstancePropertiesResponseOutputWithContext(ctx context.Context) PostgresInstancePropertiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PostgresInstancePropertiesResponseOutput)
}

func (i PostgresInstancePropertiesResponseArgs) ToPostgresInstancePropertiesResponsePtrOutput() PostgresInstancePropertiesResponsePtrOutput {
	return i.ToPostgresInstancePropertiesResponsePtrOutputWithContext(context.Background())
}

func (i PostgresInstancePropertiesResponseArgs) ToPostgresInstancePropertiesResponsePtrOutputWithContext(ctx context.Context) PostgresInstancePropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PostgresInstancePropertiesResponseOutput).ToPostgresInstancePropertiesResponsePtrOutputWithContext(ctx)
}









type PostgresInstancePropertiesResponsePtrInput interface {
	pulumi.Input

	ToPostgresInstancePropertiesResponsePtrOutput() PostgresInstancePropertiesResponsePtrOutput
	ToPostgresInstancePropertiesResponsePtrOutputWithContext(context.Context) PostgresInstancePropertiesResponsePtrOutput
}

type postgresInstancePropertiesResponsePtrType PostgresInstancePropertiesResponseArgs

func PostgresInstancePropertiesResponsePtr(v *PostgresInstancePropertiesResponseArgs) PostgresInstancePropertiesResponsePtrInput {
	return (*postgresInstancePropertiesResponsePtrType)(v)
}

func (*postgresInstancePropertiesResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PostgresInstancePropertiesResponse)(nil)).Elem()
}

func (i *postgresInstancePropertiesResponsePtrType) ToPostgresInstancePropertiesResponsePtrOutput() PostgresInstancePropertiesResponsePtrOutput {
	return i.ToPostgresInstancePropertiesResponsePtrOutputWithContext(context.Background())
}

func (i *postgresInstancePropertiesResponsePtrType) ToPostgresInstancePropertiesResponsePtrOutputWithContext(ctx context.Context) PostgresInstancePropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PostgresInstancePropertiesResponsePtrOutput)
}

type PostgresInstancePropertiesResponseOutput struct{ *pulumi.OutputState }

func (PostgresInstancePropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PostgresInstancePropertiesResponse)(nil)).Elem()
}

func (o PostgresInstancePropertiesResponseOutput) ToPostgresInstancePropertiesResponseOutput() PostgresInstancePropertiesResponseOutput {
	return o
}

func (o PostgresInstancePropertiesResponseOutput) ToPostgresInstancePropertiesResponseOutputWithContext(ctx context.Context) PostgresInstancePropertiesResponseOutput {
	return o
}

func (o PostgresInstancePropertiesResponseOutput) ToPostgresInstancePropertiesResponsePtrOutput() PostgresInstancePropertiesResponsePtrOutput {
	return o.ToPostgresInstancePropertiesResponsePtrOutputWithContext(context.Background())
}

func (o PostgresInstancePropertiesResponseOutput) ToPostgresInstancePropertiesResponsePtrOutputWithContext(ctx context.Context) PostgresInstancePropertiesResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PostgresInstancePropertiesResponse) *PostgresInstancePropertiesResponse {
		return &v
	}).(PostgresInstancePropertiesResponsePtrOutput)
}

func (o PostgresInstancePropertiesResponseOutput) Admin() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PostgresInstancePropertiesResponse) *string { return v.Admin }).(pulumi.StringPtrOutput)
}

func (o PostgresInstancePropertiesResponseOutput) BasicLoginInformation() BasicLoginInformationResponsePtrOutput {
	return o.ApplyT(func(v PostgresInstancePropertiesResponse) *BasicLoginInformationResponse {
		return v.BasicLoginInformation
	}).(BasicLoginInformationResponsePtrOutput)
}

func (o PostgresInstancePropertiesResponseOutput) DataControllerId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PostgresInstancePropertiesResponse) *string { return v.DataControllerId }).(pulumi.StringPtrOutput)
}

func (o PostgresInstancePropertiesResponseOutput) K8sRaw() pulumi.AnyOutput {
	return o.ApplyT(func(v PostgresInstancePropertiesResponse) interface{} { return v.K8sRaw }).(pulumi.AnyOutput)
}

func (o PostgresInstancePropertiesResponseOutput) LastUploadedDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PostgresInstancePropertiesResponse) *string { return v.LastUploadedDate }).(pulumi.StringPtrOutput)
}

func (o PostgresInstancePropertiesResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v PostgresInstancePropertiesResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

type PostgresInstancePropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (PostgresInstancePropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PostgresInstancePropertiesResponse)(nil)).Elem()
}

func (o PostgresInstancePropertiesResponsePtrOutput) ToPostgresInstancePropertiesResponsePtrOutput() PostgresInstancePropertiesResponsePtrOutput {
	return o
}

func (o PostgresInstancePropertiesResponsePtrOutput) ToPostgresInstancePropertiesResponsePtrOutputWithContext(ctx context.Context) PostgresInstancePropertiesResponsePtrOutput {
	return o
}

func (o PostgresInstancePropertiesResponsePtrOutput) Elem() PostgresInstancePropertiesResponseOutput {
	return o.ApplyT(func(v *PostgresInstancePropertiesResponse) PostgresInstancePropertiesResponse {
		if v != nil {
			return *v
		}
		var ret PostgresInstancePropertiesResponse
		return ret
	}).(PostgresInstancePropertiesResponseOutput)
}

func (o PostgresInstancePropertiesResponsePtrOutput) Admin() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PostgresInstancePropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.Admin
	}).(pulumi.StringPtrOutput)
}

func (o PostgresInstancePropertiesResponsePtrOutput) BasicLoginInformation() BasicLoginInformationResponsePtrOutput {
	return o.ApplyT(func(v *PostgresInstancePropertiesResponse) *BasicLoginInformationResponse {
		if v == nil {
			return nil
		}
		return v.BasicLoginInformation
	}).(BasicLoginInformationResponsePtrOutput)
}

func (o PostgresInstancePropertiesResponsePtrOutput) DataControllerId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PostgresInstancePropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.DataControllerId
	}).(pulumi.StringPtrOutput)
}

func (o PostgresInstancePropertiesResponsePtrOutput) K8sRaw() pulumi.AnyOutput {
	return o.ApplyT(func(v *PostgresInstancePropertiesResponse) interface{} {
		if v == nil {
			return nil
		}
		return v.K8sRaw
	}).(pulumi.AnyOutput)
}

func (o PostgresInstancePropertiesResponsePtrOutput) LastUploadedDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PostgresInstancePropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.LastUploadedDate
	}).(pulumi.StringPtrOutput)
}

func (o PostgresInstancePropertiesResponsePtrOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PostgresInstancePropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.ProvisioningState
	}).(pulumi.StringPtrOutput)
}

type PostgresInstanceSku struct {
	Capacity *int                     `pulumi:"capacity"`
	Dev      *bool                    `pulumi:"dev"`
	Family   *string                  `pulumi:"family"`
	Name     string                   `pulumi:"name"`
	Size     *string                  `pulumi:"size"`
	Tier     *PostgresInstanceSkuTier `pulumi:"tier"`
}





type PostgresInstanceSkuInput interface {
	pulumi.Input

	ToPostgresInstanceSkuOutput() PostgresInstanceSkuOutput
	ToPostgresInstanceSkuOutputWithContext(context.Context) PostgresInstanceSkuOutput
}

type PostgresInstanceSkuArgs struct {
	Capacity pulumi.IntPtrInput              `pulumi:"capacity"`
	Dev      pulumi.BoolPtrInput             `pulumi:"dev"`
	Family   pulumi.StringPtrInput           `pulumi:"family"`
	Name     pulumi.StringInput              `pulumi:"name"`
	Size     pulumi.StringPtrInput           `pulumi:"size"`
	Tier     PostgresInstanceSkuTierPtrInput `pulumi:"tier"`
}

func (PostgresInstanceSkuArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PostgresInstanceSku)(nil)).Elem()
}

func (i PostgresInstanceSkuArgs) ToPostgresInstanceSkuOutput() PostgresInstanceSkuOutput {
	return i.ToPostgresInstanceSkuOutputWithContext(context.Background())
}

func (i PostgresInstanceSkuArgs) ToPostgresInstanceSkuOutputWithContext(ctx context.Context) PostgresInstanceSkuOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PostgresInstanceSkuOutput)
}

func (i PostgresInstanceSkuArgs) ToPostgresInstanceSkuPtrOutput() PostgresInstanceSkuPtrOutput {
	return i.ToPostgresInstanceSkuPtrOutputWithContext(context.Background())
}

func (i PostgresInstanceSkuArgs) ToPostgresInstanceSkuPtrOutputWithContext(ctx context.Context) PostgresInstanceSkuPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PostgresInstanceSkuOutput).ToPostgresInstanceSkuPtrOutputWithContext(ctx)
}









type PostgresInstanceSkuPtrInput interface {
	pulumi.Input

	ToPostgresInstanceSkuPtrOutput() PostgresInstanceSkuPtrOutput
	ToPostgresInstanceSkuPtrOutputWithContext(context.Context) PostgresInstanceSkuPtrOutput
}

type postgresInstanceSkuPtrType PostgresInstanceSkuArgs

func PostgresInstanceSkuPtr(v *PostgresInstanceSkuArgs) PostgresInstanceSkuPtrInput {
	return (*postgresInstanceSkuPtrType)(v)
}

func (*postgresInstanceSkuPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PostgresInstanceSku)(nil)).Elem()
}

func (i *postgresInstanceSkuPtrType) ToPostgresInstanceSkuPtrOutput() PostgresInstanceSkuPtrOutput {
	return i.ToPostgresInstanceSkuPtrOutputWithContext(context.Background())
}

func (i *postgresInstanceSkuPtrType) ToPostgresInstanceSkuPtrOutputWithContext(ctx context.Context) PostgresInstanceSkuPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PostgresInstanceSkuPtrOutput)
}

type PostgresInstanceSkuOutput struct{ *pulumi.OutputState }

func (PostgresInstanceSkuOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PostgresInstanceSku)(nil)).Elem()
}

func (o PostgresInstanceSkuOutput) ToPostgresInstanceSkuOutput() PostgresInstanceSkuOutput {
	return o
}

func (o PostgresInstanceSkuOutput) ToPostgresInstanceSkuOutputWithContext(ctx context.Context) PostgresInstanceSkuOutput {
	return o
}

func (o PostgresInstanceSkuOutput) ToPostgresInstanceSkuPtrOutput() PostgresInstanceSkuPtrOutput {
	return o.ToPostgresInstanceSkuPtrOutputWithContext(context.Background())
}

func (o PostgresInstanceSkuOutput) ToPostgresInstanceSkuPtrOutputWithContext(ctx context.Context) PostgresInstanceSkuPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PostgresInstanceSku) *PostgresInstanceSku {
		return &v
	}).(PostgresInstanceSkuPtrOutput)
}

func (o PostgresInstanceSkuOutput) Capacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v PostgresInstanceSku) *int { return v.Capacity }).(pulumi.IntPtrOutput)
}

func (o PostgresInstanceSkuOutput) Dev() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v PostgresInstanceSku) *bool { return v.Dev }).(pulumi.BoolPtrOutput)
}

func (o PostgresInstanceSkuOutput) Family() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PostgresInstanceSku) *string { return v.Family }).(pulumi.StringPtrOutput)
}

func (o PostgresInstanceSkuOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v PostgresInstanceSku) string { return v.Name }).(pulumi.StringOutput)
}

func (o PostgresInstanceSkuOutput) Size() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PostgresInstanceSku) *string { return v.Size }).(pulumi.StringPtrOutput)
}

func (o PostgresInstanceSkuOutput) Tier() PostgresInstanceSkuTierPtrOutput {
	return o.ApplyT(func(v PostgresInstanceSku) *PostgresInstanceSkuTier { return v.Tier }).(PostgresInstanceSkuTierPtrOutput)
}

type PostgresInstanceSkuPtrOutput struct{ *pulumi.OutputState }

func (PostgresInstanceSkuPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PostgresInstanceSku)(nil)).Elem()
}

func (o PostgresInstanceSkuPtrOutput) ToPostgresInstanceSkuPtrOutput() PostgresInstanceSkuPtrOutput {
	return o
}

func (o PostgresInstanceSkuPtrOutput) ToPostgresInstanceSkuPtrOutputWithContext(ctx context.Context) PostgresInstanceSkuPtrOutput {
	return o
}

func (o PostgresInstanceSkuPtrOutput) Elem() PostgresInstanceSkuOutput {
	return o.ApplyT(func(v *PostgresInstanceSku) PostgresInstanceSku {
		if v != nil {
			return *v
		}
		var ret PostgresInstanceSku
		return ret
	}).(PostgresInstanceSkuOutput)
}

func (o PostgresInstanceSkuPtrOutput) Capacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *PostgresInstanceSku) *int {
		if v == nil {
			return nil
		}
		return v.Capacity
	}).(pulumi.IntPtrOutput)
}

func (o PostgresInstanceSkuPtrOutput) Dev() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *PostgresInstanceSku) *bool {
		if v == nil {
			return nil
		}
		return v.Dev
	}).(pulumi.BoolPtrOutput)
}

func (o PostgresInstanceSkuPtrOutput) Family() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PostgresInstanceSku) *string {
		if v == nil {
			return nil
		}
		return v.Family
	}).(pulumi.StringPtrOutput)
}

func (o PostgresInstanceSkuPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PostgresInstanceSku) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

func (o PostgresInstanceSkuPtrOutput) Size() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PostgresInstanceSku) *string {
		if v == nil {
			return nil
		}
		return v.Size
	}).(pulumi.StringPtrOutput)
}

func (o PostgresInstanceSkuPtrOutput) Tier() PostgresInstanceSkuTierPtrOutput {
	return o.ApplyT(func(v *PostgresInstanceSku) *PostgresInstanceSkuTier {
		if v == nil {
			return nil
		}
		return v.Tier
	}).(PostgresInstanceSkuTierPtrOutput)
}

type PostgresInstanceSkuResponse struct {
	Capacity *int    `pulumi:"capacity"`
	Dev      *bool   `pulumi:"dev"`
	Family   *string `pulumi:"family"`
	Name     string  `pulumi:"name"`
	Size     *string `pulumi:"size"`
	Tier     *string `pulumi:"tier"`
}





type PostgresInstanceSkuResponseInput interface {
	pulumi.Input

	ToPostgresInstanceSkuResponseOutput() PostgresInstanceSkuResponseOutput
	ToPostgresInstanceSkuResponseOutputWithContext(context.Context) PostgresInstanceSkuResponseOutput
}

type PostgresInstanceSkuResponseArgs struct {
	Capacity pulumi.IntPtrInput    `pulumi:"capacity"`
	Dev      pulumi.BoolPtrInput   `pulumi:"dev"`
	Family   pulumi.StringPtrInput `pulumi:"family"`
	Name     pulumi.StringInput    `pulumi:"name"`
	Size     pulumi.StringPtrInput `pulumi:"size"`
	Tier     pulumi.StringPtrInput `pulumi:"tier"`
}

func (PostgresInstanceSkuResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PostgresInstanceSkuResponse)(nil)).Elem()
}

func (i PostgresInstanceSkuResponseArgs) ToPostgresInstanceSkuResponseOutput() PostgresInstanceSkuResponseOutput {
	return i.ToPostgresInstanceSkuResponseOutputWithContext(context.Background())
}

func (i PostgresInstanceSkuResponseArgs) ToPostgresInstanceSkuResponseOutputWithContext(ctx context.Context) PostgresInstanceSkuResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PostgresInstanceSkuResponseOutput)
}

func (i PostgresInstanceSkuResponseArgs) ToPostgresInstanceSkuResponsePtrOutput() PostgresInstanceSkuResponsePtrOutput {
	return i.ToPostgresInstanceSkuResponsePtrOutputWithContext(context.Background())
}

func (i PostgresInstanceSkuResponseArgs) ToPostgresInstanceSkuResponsePtrOutputWithContext(ctx context.Context) PostgresInstanceSkuResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PostgresInstanceSkuResponseOutput).ToPostgresInstanceSkuResponsePtrOutputWithContext(ctx)
}









type PostgresInstanceSkuResponsePtrInput interface {
	pulumi.Input

	ToPostgresInstanceSkuResponsePtrOutput() PostgresInstanceSkuResponsePtrOutput
	ToPostgresInstanceSkuResponsePtrOutputWithContext(context.Context) PostgresInstanceSkuResponsePtrOutput
}

type postgresInstanceSkuResponsePtrType PostgresInstanceSkuResponseArgs

func PostgresInstanceSkuResponsePtr(v *PostgresInstanceSkuResponseArgs) PostgresInstanceSkuResponsePtrInput {
	return (*postgresInstanceSkuResponsePtrType)(v)
}

func (*postgresInstanceSkuResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PostgresInstanceSkuResponse)(nil)).Elem()
}

func (i *postgresInstanceSkuResponsePtrType) ToPostgresInstanceSkuResponsePtrOutput() PostgresInstanceSkuResponsePtrOutput {
	return i.ToPostgresInstanceSkuResponsePtrOutputWithContext(context.Background())
}

func (i *postgresInstanceSkuResponsePtrType) ToPostgresInstanceSkuResponsePtrOutputWithContext(ctx context.Context) PostgresInstanceSkuResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PostgresInstanceSkuResponsePtrOutput)
}

type PostgresInstanceSkuResponseOutput struct{ *pulumi.OutputState }

func (PostgresInstanceSkuResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PostgresInstanceSkuResponse)(nil)).Elem()
}

func (o PostgresInstanceSkuResponseOutput) ToPostgresInstanceSkuResponseOutput() PostgresInstanceSkuResponseOutput {
	return o
}

func (o PostgresInstanceSkuResponseOutput) ToPostgresInstanceSkuResponseOutputWithContext(ctx context.Context) PostgresInstanceSkuResponseOutput {
	return o
}

func (o PostgresInstanceSkuResponseOutput) ToPostgresInstanceSkuResponsePtrOutput() PostgresInstanceSkuResponsePtrOutput {
	return o.ToPostgresInstanceSkuResponsePtrOutputWithContext(context.Background())
}

func (o PostgresInstanceSkuResponseOutput) ToPostgresInstanceSkuResponsePtrOutputWithContext(ctx context.Context) PostgresInstanceSkuResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PostgresInstanceSkuResponse) *PostgresInstanceSkuResponse {
		return &v
	}).(PostgresInstanceSkuResponsePtrOutput)
}

func (o PostgresInstanceSkuResponseOutput) Capacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v PostgresInstanceSkuResponse) *int { return v.Capacity }).(pulumi.IntPtrOutput)
}

func (o PostgresInstanceSkuResponseOutput) Dev() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v PostgresInstanceSkuResponse) *bool { return v.Dev }).(pulumi.BoolPtrOutput)
}

func (o PostgresInstanceSkuResponseOutput) Family() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PostgresInstanceSkuResponse) *string { return v.Family }).(pulumi.StringPtrOutput)
}

func (o PostgresInstanceSkuResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v PostgresInstanceSkuResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o PostgresInstanceSkuResponseOutput) Size() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PostgresInstanceSkuResponse) *string { return v.Size }).(pulumi.StringPtrOutput)
}

func (o PostgresInstanceSkuResponseOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PostgresInstanceSkuResponse) *string { return v.Tier }).(pulumi.StringPtrOutput)
}

type PostgresInstanceSkuResponsePtrOutput struct{ *pulumi.OutputState }

func (PostgresInstanceSkuResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PostgresInstanceSkuResponse)(nil)).Elem()
}

func (o PostgresInstanceSkuResponsePtrOutput) ToPostgresInstanceSkuResponsePtrOutput() PostgresInstanceSkuResponsePtrOutput {
	return o
}

func (o PostgresInstanceSkuResponsePtrOutput) ToPostgresInstanceSkuResponsePtrOutputWithContext(ctx context.Context) PostgresInstanceSkuResponsePtrOutput {
	return o
}

func (o PostgresInstanceSkuResponsePtrOutput) Elem() PostgresInstanceSkuResponseOutput {
	return o.ApplyT(func(v *PostgresInstanceSkuResponse) PostgresInstanceSkuResponse {
		if v != nil {
			return *v
		}
		var ret PostgresInstanceSkuResponse
		return ret
	}).(PostgresInstanceSkuResponseOutput)
}

func (o PostgresInstanceSkuResponsePtrOutput) Capacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *PostgresInstanceSkuResponse) *int {
		if v == nil {
			return nil
		}
		return v.Capacity
	}).(pulumi.IntPtrOutput)
}

func (o PostgresInstanceSkuResponsePtrOutput) Dev() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *PostgresInstanceSkuResponse) *bool {
		if v == nil {
			return nil
		}
		return v.Dev
	}).(pulumi.BoolPtrOutput)
}

func (o PostgresInstanceSkuResponsePtrOutput) Family() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PostgresInstanceSkuResponse) *string {
		if v == nil {
			return nil
		}
		return v.Family
	}).(pulumi.StringPtrOutput)
}

func (o PostgresInstanceSkuResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PostgresInstanceSkuResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

func (o PostgresInstanceSkuResponsePtrOutput) Size() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PostgresInstanceSkuResponse) *string {
		if v == nil {
			return nil
		}
		return v.Size
	}).(pulumi.StringPtrOutput)
}

func (o PostgresInstanceSkuResponsePtrOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PostgresInstanceSkuResponse) *string {
		if v == nil {
			return nil
		}
		return v.Tier
	}).(pulumi.StringPtrOutput)
}

type SqlManagedInstanceProperties struct {
	Admin                 *string                `pulumi:"admin"`
	BasicLoginInformation *BasicLoginInformation `pulumi:"basicLoginInformation"`
	DataControllerId      *string                `pulumi:"dataControllerId"`
	EndTime               *string                `pulumi:"endTime"`
	K8sRaw                interface{}            `pulumi:"k8sRaw"`
	LastUploadedDate      *string                `pulumi:"lastUploadedDate"`
	StartTime             *string                `pulumi:"startTime"`
}





type SqlManagedInstancePropertiesInput interface {
	pulumi.Input

	ToSqlManagedInstancePropertiesOutput() SqlManagedInstancePropertiesOutput
	ToSqlManagedInstancePropertiesOutputWithContext(context.Context) SqlManagedInstancePropertiesOutput
}

type SqlManagedInstancePropertiesArgs struct {
	Admin                 pulumi.StringPtrInput         `pulumi:"admin"`
	BasicLoginInformation BasicLoginInformationPtrInput `pulumi:"basicLoginInformation"`
	DataControllerId      pulumi.StringPtrInput         `pulumi:"dataControllerId"`
	EndTime               pulumi.StringPtrInput         `pulumi:"endTime"`
	K8sRaw                pulumi.Input                  `pulumi:"k8sRaw"`
	LastUploadedDate      pulumi.StringPtrInput         `pulumi:"lastUploadedDate"`
	StartTime             pulumi.StringPtrInput         `pulumi:"startTime"`
}

func (SqlManagedInstancePropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SqlManagedInstanceProperties)(nil)).Elem()
}

func (i SqlManagedInstancePropertiesArgs) ToSqlManagedInstancePropertiesOutput() SqlManagedInstancePropertiesOutput {
	return i.ToSqlManagedInstancePropertiesOutputWithContext(context.Background())
}

func (i SqlManagedInstancePropertiesArgs) ToSqlManagedInstancePropertiesOutputWithContext(ctx context.Context) SqlManagedInstancePropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SqlManagedInstancePropertiesOutput)
}

func (i SqlManagedInstancePropertiesArgs) ToSqlManagedInstancePropertiesPtrOutput() SqlManagedInstancePropertiesPtrOutput {
	return i.ToSqlManagedInstancePropertiesPtrOutputWithContext(context.Background())
}

func (i SqlManagedInstancePropertiesArgs) ToSqlManagedInstancePropertiesPtrOutputWithContext(ctx context.Context) SqlManagedInstancePropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SqlManagedInstancePropertiesOutput).ToSqlManagedInstancePropertiesPtrOutputWithContext(ctx)
}









type SqlManagedInstancePropertiesPtrInput interface {
	pulumi.Input

	ToSqlManagedInstancePropertiesPtrOutput() SqlManagedInstancePropertiesPtrOutput
	ToSqlManagedInstancePropertiesPtrOutputWithContext(context.Context) SqlManagedInstancePropertiesPtrOutput
}

type sqlManagedInstancePropertiesPtrType SqlManagedInstancePropertiesArgs

func SqlManagedInstancePropertiesPtr(v *SqlManagedInstancePropertiesArgs) SqlManagedInstancePropertiesPtrInput {
	return (*sqlManagedInstancePropertiesPtrType)(v)
}

func (*sqlManagedInstancePropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SqlManagedInstanceProperties)(nil)).Elem()
}

func (i *sqlManagedInstancePropertiesPtrType) ToSqlManagedInstancePropertiesPtrOutput() SqlManagedInstancePropertiesPtrOutput {
	return i.ToSqlManagedInstancePropertiesPtrOutputWithContext(context.Background())
}

func (i *sqlManagedInstancePropertiesPtrType) ToSqlManagedInstancePropertiesPtrOutputWithContext(ctx context.Context) SqlManagedInstancePropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SqlManagedInstancePropertiesPtrOutput)
}

type SqlManagedInstancePropertiesOutput struct{ *pulumi.OutputState }

func (SqlManagedInstancePropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SqlManagedInstanceProperties)(nil)).Elem()
}

func (o SqlManagedInstancePropertiesOutput) ToSqlManagedInstancePropertiesOutput() SqlManagedInstancePropertiesOutput {
	return o
}

func (o SqlManagedInstancePropertiesOutput) ToSqlManagedInstancePropertiesOutputWithContext(ctx context.Context) SqlManagedInstancePropertiesOutput {
	return o
}

func (o SqlManagedInstancePropertiesOutput) ToSqlManagedInstancePropertiesPtrOutput() SqlManagedInstancePropertiesPtrOutput {
	return o.ToSqlManagedInstancePropertiesPtrOutputWithContext(context.Background())
}

func (o SqlManagedInstancePropertiesOutput) ToSqlManagedInstancePropertiesPtrOutputWithContext(ctx context.Context) SqlManagedInstancePropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SqlManagedInstanceProperties) *SqlManagedInstanceProperties {
		return &v
	}).(SqlManagedInstancePropertiesPtrOutput)
}

func (o SqlManagedInstancePropertiesOutput) Admin() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SqlManagedInstanceProperties) *string { return v.Admin }).(pulumi.StringPtrOutput)
}

func (o SqlManagedInstancePropertiesOutput) BasicLoginInformation() BasicLoginInformationPtrOutput {
	return o.ApplyT(func(v SqlManagedInstanceProperties) *BasicLoginInformation { return v.BasicLoginInformation }).(BasicLoginInformationPtrOutput)
}

func (o SqlManagedInstancePropertiesOutput) DataControllerId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SqlManagedInstanceProperties) *string { return v.DataControllerId }).(pulumi.StringPtrOutput)
}

func (o SqlManagedInstancePropertiesOutput) EndTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SqlManagedInstanceProperties) *string { return v.EndTime }).(pulumi.StringPtrOutput)
}

func (o SqlManagedInstancePropertiesOutput) K8sRaw() pulumi.AnyOutput {
	return o.ApplyT(func(v SqlManagedInstanceProperties) interface{} { return v.K8sRaw }).(pulumi.AnyOutput)
}

func (o SqlManagedInstancePropertiesOutput) LastUploadedDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SqlManagedInstanceProperties) *string { return v.LastUploadedDate }).(pulumi.StringPtrOutput)
}

func (o SqlManagedInstancePropertiesOutput) StartTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SqlManagedInstanceProperties) *string { return v.StartTime }).(pulumi.StringPtrOutput)
}

type SqlManagedInstancePropertiesPtrOutput struct{ *pulumi.OutputState }

func (SqlManagedInstancePropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SqlManagedInstanceProperties)(nil)).Elem()
}

func (o SqlManagedInstancePropertiesPtrOutput) ToSqlManagedInstancePropertiesPtrOutput() SqlManagedInstancePropertiesPtrOutput {
	return o
}

func (o SqlManagedInstancePropertiesPtrOutput) ToSqlManagedInstancePropertiesPtrOutputWithContext(ctx context.Context) SqlManagedInstancePropertiesPtrOutput {
	return o
}

func (o SqlManagedInstancePropertiesPtrOutput) Elem() SqlManagedInstancePropertiesOutput {
	return o.ApplyT(func(v *SqlManagedInstanceProperties) SqlManagedInstanceProperties {
		if v != nil {
			return *v
		}
		var ret SqlManagedInstanceProperties
		return ret
	}).(SqlManagedInstancePropertiesOutput)
}

func (o SqlManagedInstancePropertiesPtrOutput) Admin() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SqlManagedInstanceProperties) *string {
		if v == nil {
			return nil
		}
		return v.Admin
	}).(pulumi.StringPtrOutput)
}

func (o SqlManagedInstancePropertiesPtrOutput) BasicLoginInformation() BasicLoginInformationPtrOutput {
	return o.ApplyT(func(v *SqlManagedInstanceProperties) *BasicLoginInformation {
		if v == nil {
			return nil
		}
		return v.BasicLoginInformation
	}).(BasicLoginInformationPtrOutput)
}

func (o SqlManagedInstancePropertiesPtrOutput) DataControllerId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SqlManagedInstanceProperties) *string {
		if v == nil {
			return nil
		}
		return v.DataControllerId
	}).(pulumi.StringPtrOutput)
}

func (o SqlManagedInstancePropertiesPtrOutput) EndTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SqlManagedInstanceProperties) *string {
		if v == nil {
			return nil
		}
		return v.EndTime
	}).(pulumi.StringPtrOutput)
}

func (o SqlManagedInstancePropertiesPtrOutput) K8sRaw() pulumi.AnyOutput {
	return o.ApplyT(func(v *SqlManagedInstanceProperties) interface{} {
		if v == nil {
			return nil
		}
		return v.K8sRaw
	}).(pulumi.AnyOutput)
}

func (o SqlManagedInstancePropertiesPtrOutput) LastUploadedDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SqlManagedInstanceProperties) *string {
		if v == nil {
			return nil
		}
		return v.LastUploadedDate
	}).(pulumi.StringPtrOutput)
}

func (o SqlManagedInstancePropertiesPtrOutput) StartTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SqlManagedInstanceProperties) *string {
		if v == nil {
			return nil
		}
		return v.StartTime
	}).(pulumi.StringPtrOutput)
}

type SqlManagedInstancePropertiesResponse struct {
	Admin                 *string                        `pulumi:"admin"`
	BasicLoginInformation *BasicLoginInformationResponse `pulumi:"basicLoginInformation"`
	DataControllerId      *string                        `pulumi:"dataControllerId"`
	EndTime               *string                        `pulumi:"endTime"`
	K8sRaw                interface{}                    `pulumi:"k8sRaw"`
	LastUploadedDate      *string                        `pulumi:"lastUploadedDate"`
	ProvisioningState     string                         `pulumi:"provisioningState"`
	StartTime             *string                        `pulumi:"startTime"`
}





type SqlManagedInstancePropertiesResponseInput interface {
	pulumi.Input

	ToSqlManagedInstancePropertiesResponseOutput() SqlManagedInstancePropertiesResponseOutput
	ToSqlManagedInstancePropertiesResponseOutputWithContext(context.Context) SqlManagedInstancePropertiesResponseOutput
}

type SqlManagedInstancePropertiesResponseArgs struct {
	Admin                 pulumi.StringPtrInput                 `pulumi:"admin"`
	BasicLoginInformation BasicLoginInformationResponsePtrInput `pulumi:"basicLoginInformation"`
	DataControllerId      pulumi.StringPtrInput                 `pulumi:"dataControllerId"`
	EndTime               pulumi.StringPtrInput                 `pulumi:"endTime"`
	K8sRaw                pulumi.Input                          `pulumi:"k8sRaw"`
	LastUploadedDate      pulumi.StringPtrInput                 `pulumi:"lastUploadedDate"`
	ProvisioningState     pulumi.StringInput                    `pulumi:"provisioningState"`
	StartTime             pulumi.StringPtrInput                 `pulumi:"startTime"`
}

func (SqlManagedInstancePropertiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SqlManagedInstancePropertiesResponse)(nil)).Elem()
}

func (i SqlManagedInstancePropertiesResponseArgs) ToSqlManagedInstancePropertiesResponseOutput() SqlManagedInstancePropertiesResponseOutput {
	return i.ToSqlManagedInstancePropertiesResponseOutputWithContext(context.Background())
}

func (i SqlManagedInstancePropertiesResponseArgs) ToSqlManagedInstancePropertiesResponseOutputWithContext(ctx context.Context) SqlManagedInstancePropertiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SqlManagedInstancePropertiesResponseOutput)
}

func (i SqlManagedInstancePropertiesResponseArgs) ToSqlManagedInstancePropertiesResponsePtrOutput() SqlManagedInstancePropertiesResponsePtrOutput {
	return i.ToSqlManagedInstancePropertiesResponsePtrOutputWithContext(context.Background())
}

func (i SqlManagedInstancePropertiesResponseArgs) ToSqlManagedInstancePropertiesResponsePtrOutputWithContext(ctx context.Context) SqlManagedInstancePropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SqlManagedInstancePropertiesResponseOutput).ToSqlManagedInstancePropertiesResponsePtrOutputWithContext(ctx)
}









type SqlManagedInstancePropertiesResponsePtrInput interface {
	pulumi.Input

	ToSqlManagedInstancePropertiesResponsePtrOutput() SqlManagedInstancePropertiesResponsePtrOutput
	ToSqlManagedInstancePropertiesResponsePtrOutputWithContext(context.Context) SqlManagedInstancePropertiesResponsePtrOutput
}

type sqlManagedInstancePropertiesResponsePtrType SqlManagedInstancePropertiesResponseArgs

func SqlManagedInstancePropertiesResponsePtr(v *SqlManagedInstancePropertiesResponseArgs) SqlManagedInstancePropertiesResponsePtrInput {
	return (*sqlManagedInstancePropertiesResponsePtrType)(v)
}

func (*sqlManagedInstancePropertiesResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SqlManagedInstancePropertiesResponse)(nil)).Elem()
}

func (i *sqlManagedInstancePropertiesResponsePtrType) ToSqlManagedInstancePropertiesResponsePtrOutput() SqlManagedInstancePropertiesResponsePtrOutput {
	return i.ToSqlManagedInstancePropertiesResponsePtrOutputWithContext(context.Background())
}

func (i *sqlManagedInstancePropertiesResponsePtrType) ToSqlManagedInstancePropertiesResponsePtrOutputWithContext(ctx context.Context) SqlManagedInstancePropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SqlManagedInstancePropertiesResponsePtrOutput)
}

type SqlManagedInstancePropertiesResponseOutput struct{ *pulumi.OutputState }

func (SqlManagedInstancePropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SqlManagedInstancePropertiesResponse)(nil)).Elem()
}

func (o SqlManagedInstancePropertiesResponseOutput) ToSqlManagedInstancePropertiesResponseOutput() SqlManagedInstancePropertiesResponseOutput {
	return o
}

func (o SqlManagedInstancePropertiesResponseOutput) ToSqlManagedInstancePropertiesResponseOutputWithContext(ctx context.Context) SqlManagedInstancePropertiesResponseOutput {
	return o
}

func (o SqlManagedInstancePropertiesResponseOutput) ToSqlManagedInstancePropertiesResponsePtrOutput() SqlManagedInstancePropertiesResponsePtrOutput {
	return o.ToSqlManagedInstancePropertiesResponsePtrOutputWithContext(context.Background())
}

func (o SqlManagedInstancePropertiesResponseOutput) ToSqlManagedInstancePropertiesResponsePtrOutputWithContext(ctx context.Context) SqlManagedInstancePropertiesResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SqlManagedInstancePropertiesResponse) *SqlManagedInstancePropertiesResponse {
		return &v
	}).(SqlManagedInstancePropertiesResponsePtrOutput)
}

func (o SqlManagedInstancePropertiesResponseOutput) Admin() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SqlManagedInstancePropertiesResponse) *string { return v.Admin }).(pulumi.StringPtrOutput)
}

func (o SqlManagedInstancePropertiesResponseOutput) BasicLoginInformation() BasicLoginInformationResponsePtrOutput {
	return o.ApplyT(func(v SqlManagedInstancePropertiesResponse) *BasicLoginInformationResponse {
		return v.BasicLoginInformation
	}).(BasicLoginInformationResponsePtrOutput)
}

func (o SqlManagedInstancePropertiesResponseOutput) DataControllerId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SqlManagedInstancePropertiesResponse) *string { return v.DataControllerId }).(pulumi.StringPtrOutput)
}

func (o SqlManagedInstancePropertiesResponseOutput) EndTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SqlManagedInstancePropertiesResponse) *string { return v.EndTime }).(pulumi.StringPtrOutput)
}

func (o SqlManagedInstancePropertiesResponseOutput) K8sRaw() pulumi.AnyOutput {
	return o.ApplyT(func(v SqlManagedInstancePropertiesResponse) interface{} { return v.K8sRaw }).(pulumi.AnyOutput)
}

func (o SqlManagedInstancePropertiesResponseOutput) LastUploadedDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SqlManagedInstancePropertiesResponse) *string { return v.LastUploadedDate }).(pulumi.StringPtrOutput)
}

func (o SqlManagedInstancePropertiesResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v SqlManagedInstancePropertiesResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o SqlManagedInstancePropertiesResponseOutput) StartTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SqlManagedInstancePropertiesResponse) *string { return v.StartTime }).(pulumi.StringPtrOutput)
}

type SqlManagedInstancePropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (SqlManagedInstancePropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SqlManagedInstancePropertiesResponse)(nil)).Elem()
}

func (o SqlManagedInstancePropertiesResponsePtrOutput) ToSqlManagedInstancePropertiesResponsePtrOutput() SqlManagedInstancePropertiesResponsePtrOutput {
	return o
}

func (o SqlManagedInstancePropertiesResponsePtrOutput) ToSqlManagedInstancePropertiesResponsePtrOutputWithContext(ctx context.Context) SqlManagedInstancePropertiesResponsePtrOutput {
	return o
}

func (o SqlManagedInstancePropertiesResponsePtrOutput) Elem() SqlManagedInstancePropertiesResponseOutput {
	return o.ApplyT(func(v *SqlManagedInstancePropertiesResponse) SqlManagedInstancePropertiesResponse {
		if v != nil {
			return *v
		}
		var ret SqlManagedInstancePropertiesResponse
		return ret
	}).(SqlManagedInstancePropertiesResponseOutput)
}

func (o SqlManagedInstancePropertiesResponsePtrOutput) Admin() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SqlManagedInstancePropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.Admin
	}).(pulumi.StringPtrOutput)
}

func (o SqlManagedInstancePropertiesResponsePtrOutput) BasicLoginInformation() BasicLoginInformationResponsePtrOutput {
	return o.ApplyT(func(v *SqlManagedInstancePropertiesResponse) *BasicLoginInformationResponse {
		if v == nil {
			return nil
		}
		return v.BasicLoginInformation
	}).(BasicLoginInformationResponsePtrOutput)
}

func (o SqlManagedInstancePropertiesResponsePtrOutput) DataControllerId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SqlManagedInstancePropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.DataControllerId
	}).(pulumi.StringPtrOutput)
}

func (o SqlManagedInstancePropertiesResponsePtrOutput) EndTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SqlManagedInstancePropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.EndTime
	}).(pulumi.StringPtrOutput)
}

func (o SqlManagedInstancePropertiesResponsePtrOutput) K8sRaw() pulumi.AnyOutput {
	return o.ApplyT(func(v *SqlManagedInstancePropertiesResponse) interface{} {
		if v == nil {
			return nil
		}
		return v.K8sRaw
	}).(pulumi.AnyOutput)
}

func (o SqlManagedInstancePropertiesResponsePtrOutput) LastUploadedDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SqlManagedInstancePropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.LastUploadedDate
	}).(pulumi.StringPtrOutput)
}

func (o SqlManagedInstancePropertiesResponsePtrOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SqlManagedInstancePropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.ProvisioningState
	}).(pulumi.StringPtrOutput)
}

func (o SqlManagedInstancePropertiesResponsePtrOutput) StartTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SqlManagedInstancePropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.StartTime
	}).(pulumi.StringPtrOutput)
}

type SqlManagedInstanceSku struct {
	Capacity *int                       `pulumi:"capacity"`
	Dev      *bool                      `pulumi:"dev"`
	Family   *string                    `pulumi:"family"`
	Name     string                     `pulumi:"name"`
	Size     *string                    `pulumi:"size"`
	Tier     *SqlManagedInstanceSkuTier `pulumi:"tier"`
}





type SqlManagedInstanceSkuInput interface {
	pulumi.Input

	ToSqlManagedInstanceSkuOutput() SqlManagedInstanceSkuOutput
	ToSqlManagedInstanceSkuOutputWithContext(context.Context) SqlManagedInstanceSkuOutput
}

type SqlManagedInstanceSkuArgs struct {
	Capacity pulumi.IntPtrInput                `pulumi:"capacity"`
	Dev      pulumi.BoolPtrInput               `pulumi:"dev"`
	Family   pulumi.StringPtrInput             `pulumi:"family"`
	Name     pulumi.StringInput                `pulumi:"name"`
	Size     pulumi.StringPtrInput             `pulumi:"size"`
	Tier     SqlManagedInstanceSkuTierPtrInput `pulumi:"tier"`
}

func (SqlManagedInstanceSkuArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SqlManagedInstanceSku)(nil)).Elem()
}

func (i SqlManagedInstanceSkuArgs) ToSqlManagedInstanceSkuOutput() SqlManagedInstanceSkuOutput {
	return i.ToSqlManagedInstanceSkuOutputWithContext(context.Background())
}

func (i SqlManagedInstanceSkuArgs) ToSqlManagedInstanceSkuOutputWithContext(ctx context.Context) SqlManagedInstanceSkuOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SqlManagedInstanceSkuOutput)
}

func (i SqlManagedInstanceSkuArgs) ToSqlManagedInstanceSkuPtrOutput() SqlManagedInstanceSkuPtrOutput {
	return i.ToSqlManagedInstanceSkuPtrOutputWithContext(context.Background())
}

func (i SqlManagedInstanceSkuArgs) ToSqlManagedInstanceSkuPtrOutputWithContext(ctx context.Context) SqlManagedInstanceSkuPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SqlManagedInstanceSkuOutput).ToSqlManagedInstanceSkuPtrOutputWithContext(ctx)
}









type SqlManagedInstanceSkuPtrInput interface {
	pulumi.Input

	ToSqlManagedInstanceSkuPtrOutput() SqlManagedInstanceSkuPtrOutput
	ToSqlManagedInstanceSkuPtrOutputWithContext(context.Context) SqlManagedInstanceSkuPtrOutput
}

type sqlManagedInstanceSkuPtrType SqlManagedInstanceSkuArgs

func SqlManagedInstanceSkuPtr(v *SqlManagedInstanceSkuArgs) SqlManagedInstanceSkuPtrInput {
	return (*sqlManagedInstanceSkuPtrType)(v)
}

func (*sqlManagedInstanceSkuPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SqlManagedInstanceSku)(nil)).Elem()
}

func (i *sqlManagedInstanceSkuPtrType) ToSqlManagedInstanceSkuPtrOutput() SqlManagedInstanceSkuPtrOutput {
	return i.ToSqlManagedInstanceSkuPtrOutputWithContext(context.Background())
}

func (i *sqlManagedInstanceSkuPtrType) ToSqlManagedInstanceSkuPtrOutputWithContext(ctx context.Context) SqlManagedInstanceSkuPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SqlManagedInstanceSkuPtrOutput)
}

type SqlManagedInstanceSkuOutput struct{ *pulumi.OutputState }

func (SqlManagedInstanceSkuOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SqlManagedInstanceSku)(nil)).Elem()
}

func (o SqlManagedInstanceSkuOutput) ToSqlManagedInstanceSkuOutput() SqlManagedInstanceSkuOutput {
	return o
}

func (o SqlManagedInstanceSkuOutput) ToSqlManagedInstanceSkuOutputWithContext(ctx context.Context) SqlManagedInstanceSkuOutput {
	return o
}

func (o SqlManagedInstanceSkuOutput) ToSqlManagedInstanceSkuPtrOutput() SqlManagedInstanceSkuPtrOutput {
	return o.ToSqlManagedInstanceSkuPtrOutputWithContext(context.Background())
}

func (o SqlManagedInstanceSkuOutput) ToSqlManagedInstanceSkuPtrOutputWithContext(ctx context.Context) SqlManagedInstanceSkuPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SqlManagedInstanceSku) *SqlManagedInstanceSku {
		return &v
	}).(SqlManagedInstanceSkuPtrOutput)
}

func (o SqlManagedInstanceSkuOutput) Capacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SqlManagedInstanceSku) *int { return v.Capacity }).(pulumi.IntPtrOutput)
}

func (o SqlManagedInstanceSkuOutput) Dev() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v SqlManagedInstanceSku) *bool { return v.Dev }).(pulumi.BoolPtrOutput)
}

func (o SqlManagedInstanceSkuOutput) Family() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SqlManagedInstanceSku) *string { return v.Family }).(pulumi.StringPtrOutput)
}

func (o SqlManagedInstanceSkuOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v SqlManagedInstanceSku) string { return v.Name }).(pulumi.StringOutput)
}

func (o SqlManagedInstanceSkuOutput) Size() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SqlManagedInstanceSku) *string { return v.Size }).(pulumi.StringPtrOutput)
}

func (o SqlManagedInstanceSkuOutput) Tier() SqlManagedInstanceSkuTierPtrOutput {
	return o.ApplyT(func(v SqlManagedInstanceSku) *SqlManagedInstanceSkuTier { return v.Tier }).(SqlManagedInstanceSkuTierPtrOutput)
}

type SqlManagedInstanceSkuPtrOutput struct{ *pulumi.OutputState }

func (SqlManagedInstanceSkuPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SqlManagedInstanceSku)(nil)).Elem()
}

func (o SqlManagedInstanceSkuPtrOutput) ToSqlManagedInstanceSkuPtrOutput() SqlManagedInstanceSkuPtrOutput {
	return o
}

func (o SqlManagedInstanceSkuPtrOutput) ToSqlManagedInstanceSkuPtrOutputWithContext(ctx context.Context) SqlManagedInstanceSkuPtrOutput {
	return o
}

func (o SqlManagedInstanceSkuPtrOutput) Elem() SqlManagedInstanceSkuOutput {
	return o.ApplyT(func(v *SqlManagedInstanceSku) SqlManagedInstanceSku {
		if v != nil {
			return *v
		}
		var ret SqlManagedInstanceSku
		return ret
	}).(SqlManagedInstanceSkuOutput)
}

func (o SqlManagedInstanceSkuPtrOutput) Capacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SqlManagedInstanceSku) *int {
		if v == nil {
			return nil
		}
		return v.Capacity
	}).(pulumi.IntPtrOutput)
}

func (o SqlManagedInstanceSkuPtrOutput) Dev() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SqlManagedInstanceSku) *bool {
		if v == nil {
			return nil
		}
		return v.Dev
	}).(pulumi.BoolPtrOutput)
}

func (o SqlManagedInstanceSkuPtrOutput) Family() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SqlManagedInstanceSku) *string {
		if v == nil {
			return nil
		}
		return v.Family
	}).(pulumi.StringPtrOutput)
}

func (o SqlManagedInstanceSkuPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SqlManagedInstanceSku) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

func (o SqlManagedInstanceSkuPtrOutput) Size() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SqlManagedInstanceSku) *string {
		if v == nil {
			return nil
		}
		return v.Size
	}).(pulumi.StringPtrOutput)
}

func (o SqlManagedInstanceSkuPtrOutput) Tier() SqlManagedInstanceSkuTierPtrOutput {
	return o.ApplyT(func(v *SqlManagedInstanceSku) *SqlManagedInstanceSkuTier {
		if v == nil {
			return nil
		}
		return v.Tier
	}).(SqlManagedInstanceSkuTierPtrOutput)
}

type SqlManagedInstanceSkuResponse struct {
	Capacity *int    `pulumi:"capacity"`
	Dev      *bool   `pulumi:"dev"`
	Family   *string `pulumi:"family"`
	Name     string  `pulumi:"name"`
	Size     *string `pulumi:"size"`
	Tier     *string `pulumi:"tier"`
}





type SqlManagedInstanceSkuResponseInput interface {
	pulumi.Input

	ToSqlManagedInstanceSkuResponseOutput() SqlManagedInstanceSkuResponseOutput
	ToSqlManagedInstanceSkuResponseOutputWithContext(context.Context) SqlManagedInstanceSkuResponseOutput
}

type SqlManagedInstanceSkuResponseArgs struct {
	Capacity pulumi.IntPtrInput    `pulumi:"capacity"`
	Dev      pulumi.BoolPtrInput   `pulumi:"dev"`
	Family   pulumi.StringPtrInput `pulumi:"family"`
	Name     pulumi.StringInput    `pulumi:"name"`
	Size     pulumi.StringPtrInput `pulumi:"size"`
	Tier     pulumi.StringPtrInput `pulumi:"tier"`
}

func (SqlManagedInstanceSkuResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SqlManagedInstanceSkuResponse)(nil)).Elem()
}

func (i SqlManagedInstanceSkuResponseArgs) ToSqlManagedInstanceSkuResponseOutput() SqlManagedInstanceSkuResponseOutput {
	return i.ToSqlManagedInstanceSkuResponseOutputWithContext(context.Background())
}

func (i SqlManagedInstanceSkuResponseArgs) ToSqlManagedInstanceSkuResponseOutputWithContext(ctx context.Context) SqlManagedInstanceSkuResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SqlManagedInstanceSkuResponseOutput)
}

func (i SqlManagedInstanceSkuResponseArgs) ToSqlManagedInstanceSkuResponsePtrOutput() SqlManagedInstanceSkuResponsePtrOutput {
	return i.ToSqlManagedInstanceSkuResponsePtrOutputWithContext(context.Background())
}

func (i SqlManagedInstanceSkuResponseArgs) ToSqlManagedInstanceSkuResponsePtrOutputWithContext(ctx context.Context) SqlManagedInstanceSkuResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SqlManagedInstanceSkuResponseOutput).ToSqlManagedInstanceSkuResponsePtrOutputWithContext(ctx)
}









type SqlManagedInstanceSkuResponsePtrInput interface {
	pulumi.Input

	ToSqlManagedInstanceSkuResponsePtrOutput() SqlManagedInstanceSkuResponsePtrOutput
	ToSqlManagedInstanceSkuResponsePtrOutputWithContext(context.Context) SqlManagedInstanceSkuResponsePtrOutput
}

type sqlManagedInstanceSkuResponsePtrType SqlManagedInstanceSkuResponseArgs

func SqlManagedInstanceSkuResponsePtr(v *SqlManagedInstanceSkuResponseArgs) SqlManagedInstanceSkuResponsePtrInput {
	return (*sqlManagedInstanceSkuResponsePtrType)(v)
}

func (*sqlManagedInstanceSkuResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SqlManagedInstanceSkuResponse)(nil)).Elem()
}

func (i *sqlManagedInstanceSkuResponsePtrType) ToSqlManagedInstanceSkuResponsePtrOutput() SqlManagedInstanceSkuResponsePtrOutput {
	return i.ToSqlManagedInstanceSkuResponsePtrOutputWithContext(context.Background())
}

func (i *sqlManagedInstanceSkuResponsePtrType) ToSqlManagedInstanceSkuResponsePtrOutputWithContext(ctx context.Context) SqlManagedInstanceSkuResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SqlManagedInstanceSkuResponsePtrOutput)
}

type SqlManagedInstanceSkuResponseOutput struct{ *pulumi.OutputState }

func (SqlManagedInstanceSkuResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SqlManagedInstanceSkuResponse)(nil)).Elem()
}

func (o SqlManagedInstanceSkuResponseOutput) ToSqlManagedInstanceSkuResponseOutput() SqlManagedInstanceSkuResponseOutput {
	return o
}

func (o SqlManagedInstanceSkuResponseOutput) ToSqlManagedInstanceSkuResponseOutputWithContext(ctx context.Context) SqlManagedInstanceSkuResponseOutput {
	return o
}

func (o SqlManagedInstanceSkuResponseOutput) ToSqlManagedInstanceSkuResponsePtrOutput() SqlManagedInstanceSkuResponsePtrOutput {
	return o.ToSqlManagedInstanceSkuResponsePtrOutputWithContext(context.Background())
}

func (o SqlManagedInstanceSkuResponseOutput) ToSqlManagedInstanceSkuResponsePtrOutputWithContext(ctx context.Context) SqlManagedInstanceSkuResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SqlManagedInstanceSkuResponse) *SqlManagedInstanceSkuResponse {
		return &v
	}).(SqlManagedInstanceSkuResponsePtrOutput)
}

func (o SqlManagedInstanceSkuResponseOutput) Capacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SqlManagedInstanceSkuResponse) *int { return v.Capacity }).(pulumi.IntPtrOutput)
}

func (o SqlManagedInstanceSkuResponseOutput) Dev() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v SqlManagedInstanceSkuResponse) *bool { return v.Dev }).(pulumi.BoolPtrOutput)
}

func (o SqlManagedInstanceSkuResponseOutput) Family() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SqlManagedInstanceSkuResponse) *string { return v.Family }).(pulumi.StringPtrOutput)
}

func (o SqlManagedInstanceSkuResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v SqlManagedInstanceSkuResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o SqlManagedInstanceSkuResponseOutput) Size() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SqlManagedInstanceSkuResponse) *string { return v.Size }).(pulumi.StringPtrOutput)
}

func (o SqlManagedInstanceSkuResponseOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SqlManagedInstanceSkuResponse) *string { return v.Tier }).(pulumi.StringPtrOutput)
}

type SqlManagedInstanceSkuResponsePtrOutput struct{ *pulumi.OutputState }

func (SqlManagedInstanceSkuResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SqlManagedInstanceSkuResponse)(nil)).Elem()
}

func (o SqlManagedInstanceSkuResponsePtrOutput) ToSqlManagedInstanceSkuResponsePtrOutput() SqlManagedInstanceSkuResponsePtrOutput {
	return o
}

func (o SqlManagedInstanceSkuResponsePtrOutput) ToSqlManagedInstanceSkuResponsePtrOutputWithContext(ctx context.Context) SqlManagedInstanceSkuResponsePtrOutput {
	return o
}

func (o SqlManagedInstanceSkuResponsePtrOutput) Elem() SqlManagedInstanceSkuResponseOutput {
	return o.ApplyT(func(v *SqlManagedInstanceSkuResponse) SqlManagedInstanceSkuResponse {
		if v != nil {
			return *v
		}
		var ret SqlManagedInstanceSkuResponse
		return ret
	}).(SqlManagedInstanceSkuResponseOutput)
}

func (o SqlManagedInstanceSkuResponsePtrOutput) Capacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SqlManagedInstanceSkuResponse) *int {
		if v == nil {
			return nil
		}
		return v.Capacity
	}).(pulumi.IntPtrOutput)
}

func (o SqlManagedInstanceSkuResponsePtrOutput) Dev() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SqlManagedInstanceSkuResponse) *bool {
		if v == nil {
			return nil
		}
		return v.Dev
	}).(pulumi.BoolPtrOutput)
}

func (o SqlManagedInstanceSkuResponsePtrOutput) Family() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SqlManagedInstanceSkuResponse) *string {
		if v == nil {
			return nil
		}
		return v.Family
	}).(pulumi.StringPtrOutput)
}

func (o SqlManagedInstanceSkuResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SqlManagedInstanceSkuResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

func (o SqlManagedInstanceSkuResponsePtrOutput) Size() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SqlManagedInstanceSkuResponse) *string {
		if v == nil {
			return nil
		}
		return v.Size
	}).(pulumi.StringPtrOutput)
}

func (o SqlManagedInstanceSkuResponsePtrOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SqlManagedInstanceSkuResponse) *string {
		if v == nil {
			return nil
		}
		return v.Tier
	}).(pulumi.StringPtrOutput)
}

type SqlServerInstanceProperties struct {
	Collation           *string `pulumi:"collation"`
	ContainerResourceId string  `pulumi:"containerResourceId"`
	CurrentVersion      *string `pulumi:"currentVersion"`
	Edition             *string `pulumi:"edition"`
	InstanceName        *string `pulumi:"instanceName"`
	LicenseType         *string `pulumi:"licenseType"`
	PatchLevel          *string `pulumi:"patchLevel"`
	ProductId           *string `pulumi:"productId"`
	Status              string  `pulumi:"status"`
	TcpDynamicPorts     *string `pulumi:"tcpDynamicPorts"`
	TcpStaticPorts      *string `pulumi:"tcpStaticPorts"`
	VCore               *string `pulumi:"vCore"`
	Version             *string `pulumi:"version"`
}





type SqlServerInstancePropertiesInput interface {
	pulumi.Input

	ToSqlServerInstancePropertiesOutput() SqlServerInstancePropertiesOutput
	ToSqlServerInstancePropertiesOutputWithContext(context.Context) SqlServerInstancePropertiesOutput
}

type SqlServerInstancePropertiesArgs struct {
	Collation           pulumi.StringPtrInput `pulumi:"collation"`
	ContainerResourceId pulumi.StringInput    `pulumi:"containerResourceId"`
	CurrentVersion      pulumi.StringPtrInput `pulumi:"currentVersion"`
	Edition             pulumi.StringPtrInput `pulumi:"edition"`
	InstanceName        pulumi.StringPtrInput `pulumi:"instanceName"`
	LicenseType         pulumi.StringPtrInput `pulumi:"licenseType"`
	PatchLevel          pulumi.StringPtrInput `pulumi:"patchLevel"`
	ProductId           pulumi.StringPtrInput `pulumi:"productId"`
	Status              pulumi.StringInput    `pulumi:"status"`
	TcpDynamicPorts     pulumi.StringPtrInput `pulumi:"tcpDynamicPorts"`
	TcpStaticPorts      pulumi.StringPtrInput `pulumi:"tcpStaticPorts"`
	VCore               pulumi.StringPtrInput `pulumi:"vCore"`
	Version             pulumi.StringPtrInput `pulumi:"version"`
}

func (SqlServerInstancePropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SqlServerInstanceProperties)(nil)).Elem()
}

func (i SqlServerInstancePropertiesArgs) ToSqlServerInstancePropertiesOutput() SqlServerInstancePropertiesOutput {
	return i.ToSqlServerInstancePropertiesOutputWithContext(context.Background())
}

func (i SqlServerInstancePropertiesArgs) ToSqlServerInstancePropertiesOutputWithContext(ctx context.Context) SqlServerInstancePropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SqlServerInstancePropertiesOutput)
}

func (i SqlServerInstancePropertiesArgs) ToSqlServerInstancePropertiesPtrOutput() SqlServerInstancePropertiesPtrOutput {
	return i.ToSqlServerInstancePropertiesPtrOutputWithContext(context.Background())
}

func (i SqlServerInstancePropertiesArgs) ToSqlServerInstancePropertiesPtrOutputWithContext(ctx context.Context) SqlServerInstancePropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SqlServerInstancePropertiesOutput).ToSqlServerInstancePropertiesPtrOutputWithContext(ctx)
}









type SqlServerInstancePropertiesPtrInput interface {
	pulumi.Input

	ToSqlServerInstancePropertiesPtrOutput() SqlServerInstancePropertiesPtrOutput
	ToSqlServerInstancePropertiesPtrOutputWithContext(context.Context) SqlServerInstancePropertiesPtrOutput
}

type sqlServerInstancePropertiesPtrType SqlServerInstancePropertiesArgs

func SqlServerInstancePropertiesPtr(v *SqlServerInstancePropertiesArgs) SqlServerInstancePropertiesPtrInput {
	return (*sqlServerInstancePropertiesPtrType)(v)
}

func (*sqlServerInstancePropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SqlServerInstanceProperties)(nil)).Elem()
}

func (i *sqlServerInstancePropertiesPtrType) ToSqlServerInstancePropertiesPtrOutput() SqlServerInstancePropertiesPtrOutput {
	return i.ToSqlServerInstancePropertiesPtrOutputWithContext(context.Background())
}

func (i *sqlServerInstancePropertiesPtrType) ToSqlServerInstancePropertiesPtrOutputWithContext(ctx context.Context) SqlServerInstancePropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SqlServerInstancePropertiesPtrOutput)
}

type SqlServerInstancePropertiesOutput struct{ *pulumi.OutputState }

func (SqlServerInstancePropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SqlServerInstanceProperties)(nil)).Elem()
}

func (o SqlServerInstancePropertiesOutput) ToSqlServerInstancePropertiesOutput() SqlServerInstancePropertiesOutput {
	return o
}

func (o SqlServerInstancePropertiesOutput) ToSqlServerInstancePropertiesOutputWithContext(ctx context.Context) SqlServerInstancePropertiesOutput {
	return o
}

func (o SqlServerInstancePropertiesOutput) ToSqlServerInstancePropertiesPtrOutput() SqlServerInstancePropertiesPtrOutput {
	return o.ToSqlServerInstancePropertiesPtrOutputWithContext(context.Background())
}

func (o SqlServerInstancePropertiesOutput) ToSqlServerInstancePropertiesPtrOutputWithContext(ctx context.Context) SqlServerInstancePropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SqlServerInstanceProperties) *SqlServerInstanceProperties {
		return &v
	}).(SqlServerInstancePropertiesPtrOutput)
}

func (o SqlServerInstancePropertiesOutput) Collation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SqlServerInstanceProperties) *string { return v.Collation }).(pulumi.StringPtrOutput)
}

func (o SqlServerInstancePropertiesOutput) ContainerResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v SqlServerInstanceProperties) string { return v.ContainerResourceId }).(pulumi.StringOutput)
}

func (o SqlServerInstancePropertiesOutput) CurrentVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SqlServerInstanceProperties) *string { return v.CurrentVersion }).(pulumi.StringPtrOutput)
}

func (o SqlServerInstancePropertiesOutput) Edition() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SqlServerInstanceProperties) *string { return v.Edition }).(pulumi.StringPtrOutput)
}

func (o SqlServerInstancePropertiesOutput) InstanceName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SqlServerInstanceProperties) *string { return v.InstanceName }).(pulumi.StringPtrOutput)
}

func (o SqlServerInstancePropertiesOutput) LicenseType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SqlServerInstanceProperties) *string { return v.LicenseType }).(pulumi.StringPtrOutput)
}

func (o SqlServerInstancePropertiesOutput) PatchLevel() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SqlServerInstanceProperties) *string { return v.PatchLevel }).(pulumi.StringPtrOutput)
}

func (o SqlServerInstancePropertiesOutput) ProductId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SqlServerInstanceProperties) *string { return v.ProductId }).(pulumi.StringPtrOutput)
}

func (o SqlServerInstancePropertiesOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v SqlServerInstanceProperties) string { return v.Status }).(pulumi.StringOutput)
}

func (o SqlServerInstancePropertiesOutput) TcpDynamicPorts() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SqlServerInstanceProperties) *string { return v.TcpDynamicPorts }).(pulumi.StringPtrOutput)
}

func (o SqlServerInstancePropertiesOutput) TcpStaticPorts() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SqlServerInstanceProperties) *string { return v.TcpStaticPorts }).(pulumi.StringPtrOutput)
}

func (o SqlServerInstancePropertiesOutput) VCore() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SqlServerInstanceProperties) *string { return v.VCore }).(pulumi.StringPtrOutput)
}

func (o SqlServerInstancePropertiesOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SqlServerInstanceProperties) *string { return v.Version }).(pulumi.StringPtrOutput)
}

type SqlServerInstancePropertiesPtrOutput struct{ *pulumi.OutputState }

func (SqlServerInstancePropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SqlServerInstanceProperties)(nil)).Elem()
}

func (o SqlServerInstancePropertiesPtrOutput) ToSqlServerInstancePropertiesPtrOutput() SqlServerInstancePropertiesPtrOutput {
	return o
}

func (o SqlServerInstancePropertiesPtrOutput) ToSqlServerInstancePropertiesPtrOutputWithContext(ctx context.Context) SqlServerInstancePropertiesPtrOutput {
	return o
}

func (o SqlServerInstancePropertiesPtrOutput) Elem() SqlServerInstancePropertiesOutput {
	return o.ApplyT(func(v *SqlServerInstanceProperties) SqlServerInstanceProperties {
		if v != nil {
			return *v
		}
		var ret SqlServerInstanceProperties
		return ret
	}).(SqlServerInstancePropertiesOutput)
}

func (o SqlServerInstancePropertiesPtrOutput) Collation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SqlServerInstanceProperties) *string {
		if v == nil {
			return nil
		}
		return v.Collation
	}).(pulumi.StringPtrOutput)
}

func (o SqlServerInstancePropertiesPtrOutput) ContainerResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SqlServerInstanceProperties) *string {
		if v == nil {
			return nil
		}
		return &v.ContainerResourceId
	}).(pulumi.StringPtrOutput)
}

func (o SqlServerInstancePropertiesPtrOutput) CurrentVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SqlServerInstanceProperties) *string {
		if v == nil {
			return nil
		}
		return v.CurrentVersion
	}).(pulumi.StringPtrOutput)
}

func (o SqlServerInstancePropertiesPtrOutput) Edition() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SqlServerInstanceProperties) *string {
		if v == nil {
			return nil
		}
		return v.Edition
	}).(pulumi.StringPtrOutput)
}

func (o SqlServerInstancePropertiesPtrOutput) InstanceName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SqlServerInstanceProperties) *string {
		if v == nil {
			return nil
		}
		return v.InstanceName
	}).(pulumi.StringPtrOutput)
}

func (o SqlServerInstancePropertiesPtrOutput) LicenseType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SqlServerInstanceProperties) *string {
		if v == nil {
			return nil
		}
		return v.LicenseType
	}).(pulumi.StringPtrOutput)
}

func (o SqlServerInstancePropertiesPtrOutput) PatchLevel() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SqlServerInstanceProperties) *string {
		if v == nil {
			return nil
		}
		return v.PatchLevel
	}).(pulumi.StringPtrOutput)
}

func (o SqlServerInstancePropertiesPtrOutput) ProductId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SqlServerInstanceProperties) *string {
		if v == nil {
			return nil
		}
		return v.ProductId
	}).(pulumi.StringPtrOutput)
}

func (o SqlServerInstancePropertiesPtrOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SqlServerInstanceProperties) *string {
		if v == nil {
			return nil
		}
		return &v.Status
	}).(pulumi.StringPtrOutput)
}

func (o SqlServerInstancePropertiesPtrOutput) TcpDynamicPorts() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SqlServerInstanceProperties) *string {
		if v == nil {
			return nil
		}
		return v.TcpDynamicPorts
	}).(pulumi.StringPtrOutput)
}

func (o SqlServerInstancePropertiesPtrOutput) TcpStaticPorts() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SqlServerInstanceProperties) *string {
		if v == nil {
			return nil
		}
		return v.TcpStaticPorts
	}).(pulumi.StringPtrOutput)
}

func (o SqlServerInstancePropertiesPtrOutput) VCore() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SqlServerInstanceProperties) *string {
		if v == nil {
			return nil
		}
		return v.VCore
	}).(pulumi.StringPtrOutput)
}

func (o SqlServerInstancePropertiesPtrOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SqlServerInstanceProperties) *string {
		if v == nil {
			return nil
		}
		return v.Version
	}).(pulumi.StringPtrOutput)
}

type SqlServerInstancePropertiesResponse struct {
	Collation           *string `pulumi:"collation"`
	ContainerResourceId string  `pulumi:"containerResourceId"`
	CreateTime          string  `pulumi:"createTime"`
	CurrentVersion      *string `pulumi:"currentVersion"`
	Edition             *string `pulumi:"edition"`
	InstanceName        *string `pulumi:"instanceName"`
	LicenseType         *string `pulumi:"licenseType"`
	PatchLevel          *string `pulumi:"patchLevel"`
	ProductId           *string `pulumi:"productId"`
	ProvisioningState   string  `pulumi:"provisioningState"`
	Status              string  `pulumi:"status"`
	TcpDynamicPorts     *string `pulumi:"tcpDynamicPorts"`
	TcpStaticPorts      *string `pulumi:"tcpStaticPorts"`
	VCore               *string `pulumi:"vCore"`
	Version             *string `pulumi:"version"`
}





type SqlServerInstancePropertiesResponseInput interface {
	pulumi.Input

	ToSqlServerInstancePropertiesResponseOutput() SqlServerInstancePropertiesResponseOutput
	ToSqlServerInstancePropertiesResponseOutputWithContext(context.Context) SqlServerInstancePropertiesResponseOutput
}

type SqlServerInstancePropertiesResponseArgs struct {
	Collation           pulumi.StringPtrInput `pulumi:"collation"`
	ContainerResourceId pulumi.StringInput    `pulumi:"containerResourceId"`
	CreateTime          pulumi.StringInput    `pulumi:"createTime"`
	CurrentVersion      pulumi.StringPtrInput `pulumi:"currentVersion"`
	Edition             pulumi.StringPtrInput `pulumi:"edition"`
	InstanceName        pulumi.StringPtrInput `pulumi:"instanceName"`
	LicenseType         pulumi.StringPtrInput `pulumi:"licenseType"`
	PatchLevel          pulumi.StringPtrInput `pulumi:"patchLevel"`
	ProductId           pulumi.StringPtrInput `pulumi:"productId"`
	ProvisioningState   pulumi.StringInput    `pulumi:"provisioningState"`
	Status              pulumi.StringInput    `pulumi:"status"`
	TcpDynamicPorts     pulumi.StringPtrInput `pulumi:"tcpDynamicPorts"`
	TcpStaticPorts      pulumi.StringPtrInput `pulumi:"tcpStaticPorts"`
	VCore               pulumi.StringPtrInput `pulumi:"vCore"`
	Version             pulumi.StringPtrInput `pulumi:"version"`
}

func (SqlServerInstancePropertiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SqlServerInstancePropertiesResponse)(nil)).Elem()
}

func (i SqlServerInstancePropertiesResponseArgs) ToSqlServerInstancePropertiesResponseOutput() SqlServerInstancePropertiesResponseOutput {
	return i.ToSqlServerInstancePropertiesResponseOutputWithContext(context.Background())
}

func (i SqlServerInstancePropertiesResponseArgs) ToSqlServerInstancePropertiesResponseOutputWithContext(ctx context.Context) SqlServerInstancePropertiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SqlServerInstancePropertiesResponseOutput)
}

func (i SqlServerInstancePropertiesResponseArgs) ToSqlServerInstancePropertiesResponsePtrOutput() SqlServerInstancePropertiesResponsePtrOutput {
	return i.ToSqlServerInstancePropertiesResponsePtrOutputWithContext(context.Background())
}

func (i SqlServerInstancePropertiesResponseArgs) ToSqlServerInstancePropertiesResponsePtrOutputWithContext(ctx context.Context) SqlServerInstancePropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SqlServerInstancePropertiesResponseOutput).ToSqlServerInstancePropertiesResponsePtrOutputWithContext(ctx)
}









type SqlServerInstancePropertiesResponsePtrInput interface {
	pulumi.Input

	ToSqlServerInstancePropertiesResponsePtrOutput() SqlServerInstancePropertiesResponsePtrOutput
	ToSqlServerInstancePropertiesResponsePtrOutputWithContext(context.Context) SqlServerInstancePropertiesResponsePtrOutput
}

type sqlServerInstancePropertiesResponsePtrType SqlServerInstancePropertiesResponseArgs

func SqlServerInstancePropertiesResponsePtr(v *SqlServerInstancePropertiesResponseArgs) SqlServerInstancePropertiesResponsePtrInput {
	return (*sqlServerInstancePropertiesResponsePtrType)(v)
}

func (*sqlServerInstancePropertiesResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SqlServerInstancePropertiesResponse)(nil)).Elem()
}

func (i *sqlServerInstancePropertiesResponsePtrType) ToSqlServerInstancePropertiesResponsePtrOutput() SqlServerInstancePropertiesResponsePtrOutput {
	return i.ToSqlServerInstancePropertiesResponsePtrOutputWithContext(context.Background())
}

func (i *sqlServerInstancePropertiesResponsePtrType) ToSqlServerInstancePropertiesResponsePtrOutputWithContext(ctx context.Context) SqlServerInstancePropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SqlServerInstancePropertiesResponsePtrOutput)
}

type SqlServerInstancePropertiesResponseOutput struct{ *pulumi.OutputState }

func (SqlServerInstancePropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SqlServerInstancePropertiesResponse)(nil)).Elem()
}

func (o SqlServerInstancePropertiesResponseOutput) ToSqlServerInstancePropertiesResponseOutput() SqlServerInstancePropertiesResponseOutput {
	return o
}

func (o SqlServerInstancePropertiesResponseOutput) ToSqlServerInstancePropertiesResponseOutputWithContext(ctx context.Context) SqlServerInstancePropertiesResponseOutput {
	return o
}

func (o SqlServerInstancePropertiesResponseOutput) ToSqlServerInstancePropertiesResponsePtrOutput() SqlServerInstancePropertiesResponsePtrOutput {
	return o.ToSqlServerInstancePropertiesResponsePtrOutputWithContext(context.Background())
}

func (o SqlServerInstancePropertiesResponseOutput) ToSqlServerInstancePropertiesResponsePtrOutputWithContext(ctx context.Context) SqlServerInstancePropertiesResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SqlServerInstancePropertiesResponse) *SqlServerInstancePropertiesResponse {
		return &v
	}).(SqlServerInstancePropertiesResponsePtrOutput)
}

func (o SqlServerInstancePropertiesResponseOutput) Collation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SqlServerInstancePropertiesResponse) *string { return v.Collation }).(pulumi.StringPtrOutput)
}

func (o SqlServerInstancePropertiesResponseOutput) ContainerResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v SqlServerInstancePropertiesResponse) string { return v.ContainerResourceId }).(pulumi.StringOutput)
}

func (o SqlServerInstancePropertiesResponseOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v SqlServerInstancePropertiesResponse) string { return v.CreateTime }).(pulumi.StringOutput)
}

func (o SqlServerInstancePropertiesResponseOutput) CurrentVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SqlServerInstancePropertiesResponse) *string { return v.CurrentVersion }).(pulumi.StringPtrOutput)
}

func (o SqlServerInstancePropertiesResponseOutput) Edition() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SqlServerInstancePropertiesResponse) *string { return v.Edition }).(pulumi.StringPtrOutput)
}

func (o SqlServerInstancePropertiesResponseOutput) InstanceName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SqlServerInstancePropertiesResponse) *string { return v.InstanceName }).(pulumi.StringPtrOutput)
}

func (o SqlServerInstancePropertiesResponseOutput) LicenseType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SqlServerInstancePropertiesResponse) *string { return v.LicenseType }).(pulumi.StringPtrOutput)
}

func (o SqlServerInstancePropertiesResponseOutput) PatchLevel() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SqlServerInstancePropertiesResponse) *string { return v.PatchLevel }).(pulumi.StringPtrOutput)
}

func (o SqlServerInstancePropertiesResponseOutput) ProductId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SqlServerInstancePropertiesResponse) *string { return v.ProductId }).(pulumi.StringPtrOutput)
}

func (o SqlServerInstancePropertiesResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v SqlServerInstancePropertiesResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o SqlServerInstancePropertiesResponseOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v SqlServerInstancePropertiesResponse) string { return v.Status }).(pulumi.StringOutput)
}

func (o SqlServerInstancePropertiesResponseOutput) TcpDynamicPorts() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SqlServerInstancePropertiesResponse) *string { return v.TcpDynamicPorts }).(pulumi.StringPtrOutput)
}

func (o SqlServerInstancePropertiesResponseOutput) TcpStaticPorts() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SqlServerInstancePropertiesResponse) *string { return v.TcpStaticPorts }).(pulumi.StringPtrOutput)
}

func (o SqlServerInstancePropertiesResponseOutput) VCore() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SqlServerInstancePropertiesResponse) *string { return v.VCore }).(pulumi.StringPtrOutput)
}

func (o SqlServerInstancePropertiesResponseOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SqlServerInstancePropertiesResponse) *string { return v.Version }).(pulumi.StringPtrOutput)
}

type SqlServerInstancePropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (SqlServerInstancePropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SqlServerInstancePropertiesResponse)(nil)).Elem()
}

func (o SqlServerInstancePropertiesResponsePtrOutput) ToSqlServerInstancePropertiesResponsePtrOutput() SqlServerInstancePropertiesResponsePtrOutput {
	return o
}

func (o SqlServerInstancePropertiesResponsePtrOutput) ToSqlServerInstancePropertiesResponsePtrOutputWithContext(ctx context.Context) SqlServerInstancePropertiesResponsePtrOutput {
	return o
}

func (o SqlServerInstancePropertiesResponsePtrOutput) Elem() SqlServerInstancePropertiesResponseOutput {
	return o.ApplyT(func(v *SqlServerInstancePropertiesResponse) SqlServerInstancePropertiesResponse {
		if v != nil {
			return *v
		}
		var ret SqlServerInstancePropertiesResponse
		return ret
	}).(SqlServerInstancePropertiesResponseOutput)
}

func (o SqlServerInstancePropertiesResponsePtrOutput) Collation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SqlServerInstancePropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.Collation
	}).(pulumi.StringPtrOutput)
}

func (o SqlServerInstancePropertiesResponsePtrOutput) ContainerResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SqlServerInstancePropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.ContainerResourceId
	}).(pulumi.StringPtrOutput)
}

func (o SqlServerInstancePropertiesResponsePtrOutput) CreateTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SqlServerInstancePropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.CreateTime
	}).(pulumi.StringPtrOutput)
}

func (o SqlServerInstancePropertiesResponsePtrOutput) CurrentVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SqlServerInstancePropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.CurrentVersion
	}).(pulumi.StringPtrOutput)
}

func (o SqlServerInstancePropertiesResponsePtrOutput) Edition() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SqlServerInstancePropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.Edition
	}).(pulumi.StringPtrOutput)
}

func (o SqlServerInstancePropertiesResponsePtrOutput) InstanceName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SqlServerInstancePropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.InstanceName
	}).(pulumi.StringPtrOutput)
}

func (o SqlServerInstancePropertiesResponsePtrOutput) LicenseType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SqlServerInstancePropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.LicenseType
	}).(pulumi.StringPtrOutput)
}

func (o SqlServerInstancePropertiesResponsePtrOutput) PatchLevel() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SqlServerInstancePropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.PatchLevel
	}).(pulumi.StringPtrOutput)
}

func (o SqlServerInstancePropertiesResponsePtrOutput) ProductId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SqlServerInstancePropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.ProductId
	}).(pulumi.StringPtrOutput)
}

func (o SqlServerInstancePropertiesResponsePtrOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SqlServerInstancePropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.ProvisioningState
	}).(pulumi.StringPtrOutput)
}

func (o SqlServerInstancePropertiesResponsePtrOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SqlServerInstancePropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Status
	}).(pulumi.StringPtrOutput)
}

func (o SqlServerInstancePropertiesResponsePtrOutput) TcpDynamicPorts() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SqlServerInstancePropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.TcpDynamicPorts
	}).(pulumi.StringPtrOutput)
}

func (o SqlServerInstancePropertiesResponsePtrOutput) TcpStaticPorts() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SqlServerInstancePropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.TcpStaticPorts
	}).(pulumi.StringPtrOutput)
}

func (o SqlServerInstancePropertiesResponsePtrOutput) VCore() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SqlServerInstancePropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.VCore
	}).(pulumi.StringPtrOutput)
}

func (o SqlServerInstancePropertiesResponsePtrOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SqlServerInstancePropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.Version
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

type UploadServicePrincipal struct {
	Authority    *string `pulumi:"authority"`
	ClientId     *string `pulumi:"clientId"`
	ClientSecret *string `pulumi:"clientSecret"`
	TenantId     *string `pulumi:"tenantId"`
}





type UploadServicePrincipalInput interface {
	pulumi.Input

	ToUploadServicePrincipalOutput() UploadServicePrincipalOutput
	ToUploadServicePrincipalOutputWithContext(context.Context) UploadServicePrincipalOutput
}

type UploadServicePrincipalArgs struct {
	Authority    pulumi.StringPtrInput `pulumi:"authority"`
	ClientId     pulumi.StringPtrInput `pulumi:"clientId"`
	ClientSecret pulumi.StringPtrInput `pulumi:"clientSecret"`
	TenantId     pulumi.StringPtrInput `pulumi:"tenantId"`
}

func (UploadServicePrincipalArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*UploadServicePrincipal)(nil)).Elem()
}

func (i UploadServicePrincipalArgs) ToUploadServicePrincipalOutput() UploadServicePrincipalOutput {
	return i.ToUploadServicePrincipalOutputWithContext(context.Background())
}

func (i UploadServicePrincipalArgs) ToUploadServicePrincipalOutputWithContext(ctx context.Context) UploadServicePrincipalOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UploadServicePrincipalOutput)
}

func (i UploadServicePrincipalArgs) ToUploadServicePrincipalPtrOutput() UploadServicePrincipalPtrOutput {
	return i.ToUploadServicePrincipalPtrOutputWithContext(context.Background())
}

func (i UploadServicePrincipalArgs) ToUploadServicePrincipalPtrOutputWithContext(ctx context.Context) UploadServicePrincipalPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UploadServicePrincipalOutput).ToUploadServicePrincipalPtrOutputWithContext(ctx)
}









type UploadServicePrincipalPtrInput interface {
	pulumi.Input

	ToUploadServicePrincipalPtrOutput() UploadServicePrincipalPtrOutput
	ToUploadServicePrincipalPtrOutputWithContext(context.Context) UploadServicePrincipalPtrOutput
}

type uploadServicePrincipalPtrType UploadServicePrincipalArgs

func UploadServicePrincipalPtr(v *UploadServicePrincipalArgs) UploadServicePrincipalPtrInput {
	return (*uploadServicePrincipalPtrType)(v)
}

func (*uploadServicePrincipalPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**UploadServicePrincipal)(nil)).Elem()
}

func (i *uploadServicePrincipalPtrType) ToUploadServicePrincipalPtrOutput() UploadServicePrincipalPtrOutput {
	return i.ToUploadServicePrincipalPtrOutputWithContext(context.Background())
}

func (i *uploadServicePrincipalPtrType) ToUploadServicePrincipalPtrOutputWithContext(ctx context.Context) UploadServicePrincipalPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UploadServicePrincipalPtrOutput)
}

type UploadServicePrincipalOutput struct{ *pulumi.OutputState }

func (UploadServicePrincipalOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UploadServicePrincipal)(nil)).Elem()
}

func (o UploadServicePrincipalOutput) ToUploadServicePrincipalOutput() UploadServicePrincipalOutput {
	return o
}

func (o UploadServicePrincipalOutput) ToUploadServicePrincipalOutputWithContext(ctx context.Context) UploadServicePrincipalOutput {
	return o
}

func (o UploadServicePrincipalOutput) ToUploadServicePrincipalPtrOutput() UploadServicePrincipalPtrOutput {
	return o.ToUploadServicePrincipalPtrOutputWithContext(context.Background())
}

func (o UploadServicePrincipalOutput) ToUploadServicePrincipalPtrOutputWithContext(ctx context.Context) UploadServicePrincipalPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v UploadServicePrincipal) *UploadServicePrincipal {
		return &v
	}).(UploadServicePrincipalPtrOutput)
}

func (o UploadServicePrincipalOutput) Authority() pulumi.StringPtrOutput {
	return o.ApplyT(func(v UploadServicePrincipal) *string { return v.Authority }).(pulumi.StringPtrOutput)
}

func (o UploadServicePrincipalOutput) ClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v UploadServicePrincipal) *string { return v.ClientId }).(pulumi.StringPtrOutput)
}

func (o UploadServicePrincipalOutput) ClientSecret() pulumi.StringPtrOutput {
	return o.ApplyT(func(v UploadServicePrincipal) *string { return v.ClientSecret }).(pulumi.StringPtrOutput)
}

func (o UploadServicePrincipalOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v UploadServicePrincipal) *string { return v.TenantId }).(pulumi.StringPtrOutput)
}

type UploadServicePrincipalPtrOutput struct{ *pulumi.OutputState }

func (UploadServicePrincipalPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**UploadServicePrincipal)(nil)).Elem()
}

func (o UploadServicePrincipalPtrOutput) ToUploadServicePrincipalPtrOutput() UploadServicePrincipalPtrOutput {
	return o
}

func (o UploadServicePrincipalPtrOutput) ToUploadServicePrincipalPtrOutputWithContext(ctx context.Context) UploadServicePrincipalPtrOutput {
	return o
}

func (o UploadServicePrincipalPtrOutput) Elem() UploadServicePrincipalOutput {
	return o.ApplyT(func(v *UploadServicePrincipal) UploadServicePrincipal {
		if v != nil {
			return *v
		}
		var ret UploadServicePrincipal
		return ret
	}).(UploadServicePrincipalOutput)
}

func (o UploadServicePrincipalPtrOutput) Authority() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UploadServicePrincipal) *string {
		if v == nil {
			return nil
		}
		return v.Authority
	}).(pulumi.StringPtrOutput)
}

func (o UploadServicePrincipalPtrOutput) ClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UploadServicePrincipal) *string {
		if v == nil {
			return nil
		}
		return v.ClientId
	}).(pulumi.StringPtrOutput)
}

func (o UploadServicePrincipalPtrOutput) ClientSecret() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UploadServicePrincipal) *string {
		if v == nil {
			return nil
		}
		return v.ClientSecret
	}).(pulumi.StringPtrOutput)
}

func (o UploadServicePrincipalPtrOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UploadServicePrincipal) *string {
		if v == nil {
			return nil
		}
		return v.TenantId
	}).(pulumi.StringPtrOutput)
}

type UploadServicePrincipalResponse struct {
	Authority *string `pulumi:"authority"`
	ClientId  *string `pulumi:"clientId"`
	TenantId  *string `pulumi:"tenantId"`
}





type UploadServicePrincipalResponseInput interface {
	pulumi.Input

	ToUploadServicePrincipalResponseOutput() UploadServicePrincipalResponseOutput
	ToUploadServicePrincipalResponseOutputWithContext(context.Context) UploadServicePrincipalResponseOutput
}

type UploadServicePrincipalResponseArgs struct {
	Authority pulumi.StringPtrInput `pulumi:"authority"`
	ClientId  pulumi.StringPtrInput `pulumi:"clientId"`
	TenantId  pulumi.StringPtrInput `pulumi:"tenantId"`
}

func (UploadServicePrincipalResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*UploadServicePrincipalResponse)(nil)).Elem()
}

func (i UploadServicePrincipalResponseArgs) ToUploadServicePrincipalResponseOutput() UploadServicePrincipalResponseOutput {
	return i.ToUploadServicePrincipalResponseOutputWithContext(context.Background())
}

func (i UploadServicePrincipalResponseArgs) ToUploadServicePrincipalResponseOutputWithContext(ctx context.Context) UploadServicePrincipalResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UploadServicePrincipalResponseOutput)
}

func (i UploadServicePrincipalResponseArgs) ToUploadServicePrincipalResponsePtrOutput() UploadServicePrincipalResponsePtrOutput {
	return i.ToUploadServicePrincipalResponsePtrOutputWithContext(context.Background())
}

func (i UploadServicePrincipalResponseArgs) ToUploadServicePrincipalResponsePtrOutputWithContext(ctx context.Context) UploadServicePrincipalResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UploadServicePrincipalResponseOutput).ToUploadServicePrincipalResponsePtrOutputWithContext(ctx)
}









type UploadServicePrincipalResponsePtrInput interface {
	pulumi.Input

	ToUploadServicePrincipalResponsePtrOutput() UploadServicePrincipalResponsePtrOutput
	ToUploadServicePrincipalResponsePtrOutputWithContext(context.Context) UploadServicePrincipalResponsePtrOutput
}

type uploadServicePrincipalResponsePtrType UploadServicePrincipalResponseArgs

func UploadServicePrincipalResponsePtr(v *UploadServicePrincipalResponseArgs) UploadServicePrincipalResponsePtrInput {
	return (*uploadServicePrincipalResponsePtrType)(v)
}

func (*uploadServicePrincipalResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**UploadServicePrincipalResponse)(nil)).Elem()
}

func (i *uploadServicePrincipalResponsePtrType) ToUploadServicePrincipalResponsePtrOutput() UploadServicePrincipalResponsePtrOutput {
	return i.ToUploadServicePrincipalResponsePtrOutputWithContext(context.Background())
}

func (i *uploadServicePrincipalResponsePtrType) ToUploadServicePrincipalResponsePtrOutputWithContext(ctx context.Context) UploadServicePrincipalResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UploadServicePrincipalResponsePtrOutput)
}

type UploadServicePrincipalResponseOutput struct{ *pulumi.OutputState }

func (UploadServicePrincipalResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UploadServicePrincipalResponse)(nil)).Elem()
}

func (o UploadServicePrincipalResponseOutput) ToUploadServicePrincipalResponseOutput() UploadServicePrincipalResponseOutput {
	return o
}

func (o UploadServicePrincipalResponseOutput) ToUploadServicePrincipalResponseOutputWithContext(ctx context.Context) UploadServicePrincipalResponseOutput {
	return o
}

func (o UploadServicePrincipalResponseOutput) ToUploadServicePrincipalResponsePtrOutput() UploadServicePrincipalResponsePtrOutput {
	return o.ToUploadServicePrincipalResponsePtrOutputWithContext(context.Background())
}

func (o UploadServicePrincipalResponseOutput) ToUploadServicePrincipalResponsePtrOutputWithContext(ctx context.Context) UploadServicePrincipalResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v UploadServicePrincipalResponse) *UploadServicePrincipalResponse {
		return &v
	}).(UploadServicePrincipalResponsePtrOutput)
}

func (o UploadServicePrincipalResponseOutput) Authority() pulumi.StringPtrOutput {
	return o.ApplyT(func(v UploadServicePrincipalResponse) *string { return v.Authority }).(pulumi.StringPtrOutput)
}

func (o UploadServicePrincipalResponseOutput) ClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v UploadServicePrincipalResponse) *string { return v.ClientId }).(pulumi.StringPtrOutput)
}

func (o UploadServicePrincipalResponseOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v UploadServicePrincipalResponse) *string { return v.TenantId }).(pulumi.StringPtrOutput)
}

type UploadServicePrincipalResponsePtrOutput struct{ *pulumi.OutputState }

func (UploadServicePrincipalResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**UploadServicePrincipalResponse)(nil)).Elem()
}

func (o UploadServicePrincipalResponsePtrOutput) ToUploadServicePrincipalResponsePtrOutput() UploadServicePrincipalResponsePtrOutput {
	return o
}

func (o UploadServicePrincipalResponsePtrOutput) ToUploadServicePrincipalResponsePtrOutputWithContext(ctx context.Context) UploadServicePrincipalResponsePtrOutput {
	return o
}

func (o UploadServicePrincipalResponsePtrOutput) Elem() UploadServicePrincipalResponseOutput {
	return o.ApplyT(func(v *UploadServicePrincipalResponse) UploadServicePrincipalResponse {
		if v != nil {
			return *v
		}
		var ret UploadServicePrincipalResponse
		return ret
	}).(UploadServicePrincipalResponseOutput)
}

func (o UploadServicePrincipalResponsePtrOutput) Authority() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UploadServicePrincipalResponse) *string {
		if v == nil {
			return nil
		}
		return v.Authority
	}).(pulumi.StringPtrOutput)
}

func (o UploadServicePrincipalResponsePtrOutput) ClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UploadServicePrincipalResponse) *string {
		if v == nil {
			return nil
		}
		return v.ClientId
	}).(pulumi.StringPtrOutput)
}

func (o UploadServicePrincipalResponsePtrOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UploadServicePrincipalResponse) *string {
		if v == nil {
			return nil
		}
		return v.TenantId
	}).(pulumi.StringPtrOutput)
}

type UploadWatermark struct {
	Logs    *string `pulumi:"logs"`
	Metrics *string `pulumi:"metrics"`
	Usages  *string `pulumi:"usages"`
}





type UploadWatermarkInput interface {
	pulumi.Input

	ToUploadWatermarkOutput() UploadWatermarkOutput
	ToUploadWatermarkOutputWithContext(context.Context) UploadWatermarkOutput
}

type UploadWatermarkArgs struct {
	Logs    pulumi.StringPtrInput `pulumi:"logs"`
	Metrics pulumi.StringPtrInput `pulumi:"metrics"`
	Usages  pulumi.StringPtrInput `pulumi:"usages"`
}

func (UploadWatermarkArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*UploadWatermark)(nil)).Elem()
}

func (i UploadWatermarkArgs) ToUploadWatermarkOutput() UploadWatermarkOutput {
	return i.ToUploadWatermarkOutputWithContext(context.Background())
}

func (i UploadWatermarkArgs) ToUploadWatermarkOutputWithContext(ctx context.Context) UploadWatermarkOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UploadWatermarkOutput)
}

func (i UploadWatermarkArgs) ToUploadWatermarkPtrOutput() UploadWatermarkPtrOutput {
	return i.ToUploadWatermarkPtrOutputWithContext(context.Background())
}

func (i UploadWatermarkArgs) ToUploadWatermarkPtrOutputWithContext(ctx context.Context) UploadWatermarkPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UploadWatermarkOutput).ToUploadWatermarkPtrOutputWithContext(ctx)
}









type UploadWatermarkPtrInput interface {
	pulumi.Input

	ToUploadWatermarkPtrOutput() UploadWatermarkPtrOutput
	ToUploadWatermarkPtrOutputWithContext(context.Context) UploadWatermarkPtrOutput
}

type uploadWatermarkPtrType UploadWatermarkArgs

func UploadWatermarkPtr(v *UploadWatermarkArgs) UploadWatermarkPtrInput {
	return (*uploadWatermarkPtrType)(v)
}

func (*uploadWatermarkPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**UploadWatermark)(nil)).Elem()
}

func (i *uploadWatermarkPtrType) ToUploadWatermarkPtrOutput() UploadWatermarkPtrOutput {
	return i.ToUploadWatermarkPtrOutputWithContext(context.Background())
}

func (i *uploadWatermarkPtrType) ToUploadWatermarkPtrOutputWithContext(ctx context.Context) UploadWatermarkPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UploadWatermarkPtrOutput)
}

type UploadWatermarkOutput struct{ *pulumi.OutputState }

func (UploadWatermarkOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UploadWatermark)(nil)).Elem()
}

func (o UploadWatermarkOutput) ToUploadWatermarkOutput() UploadWatermarkOutput {
	return o
}

func (o UploadWatermarkOutput) ToUploadWatermarkOutputWithContext(ctx context.Context) UploadWatermarkOutput {
	return o
}

func (o UploadWatermarkOutput) ToUploadWatermarkPtrOutput() UploadWatermarkPtrOutput {
	return o.ToUploadWatermarkPtrOutputWithContext(context.Background())
}

func (o UploadWatermarkOutput) ToUploadWatermarkPtrOutputWithContext(ctx context.Context) UploadWatermarkPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v UploadWatermark) *UploadWatermark {
		return &v
	}).(UploadWatermarkPtrOutput)
}

func (o UploadWatermarkOutput) Logs() pulumi.StringPtrOutput {
	return o.ApplyT(func(v UploadWatermark) *string { return v.Logs }).(pulumi.StringPtrOutput)
}

func (o UploadWatermarkOutput) Metrics() pulumi.StringPtrOutput {
	return o.ApplyT(func(v UploadWatermark) *string { return v.Metrics }).(pulumi.StringPtrOutput)
}

func (o UploadWatermarkOutput) Usages() pulumi.StringPtrOutput {
	return o.ApplyT(func(v UploadWatermark) *string { return v.Usages }).(pulumi.StringPtrOutput)
}

type UploadWatermarkPtrOutput struct{ *pulumi.OutputState }

func (UploadWatermarkPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**UploadWatermark)(nil)).Elem()
}

func (o UploadWatermarkPtrOutput) ToUploadWatermarkPtrOutput() UploadWatermarkPtrOutput {
	return o
}

func (o UploadWatermarkPtrOutput) ToUploadWatermarkPtrOutputWithContext(ctx context.Context) UploadWatermarkPtrOutput {
	return o
}

func (o UploadWatermarkPtrOutput) Elem() UploadWatermarkOutput {
	return o.ApplyT(func(v *UploadWatermark) UploadWatermark {
		if v != nil {
			return *v
		}
		var ret UploadWatermark
		return ret
	}).(UploadWatermarkOutput)
}

func (o UploadWatermarkPtrOutput) Logs() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UploadWatermark) *string {
		if v == nil {
			return nil
		}
		return v.Logs
	}).(pulumi.StringPtrOutput)
}

func (o UploadWatermarkPtrOutput) Metrics() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UploadWatermark) *string {
		if v == nil {
			return nil
		}
		return v.Metrics
	}).(pulumi.StringPtrOutput)
}

func (o UploadWatermarkPtrOutput) Usages() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UploadWatermark) *string {
		if v == nil {
			return nil
		}
		return v.Usages
	}).(pulumi.StringPtrOutput)
}

type UploadWatermarkResponse struct {
	Logs    *string `pulumi:"logs"`
	Metrics *string `pulumi:"metrics"`
	Usages  *string `pulumi:"usages"`
}





type UploadWatermarkResponseInput interface {
	pulumi.Input

	ToUploadWatermarkResponseOutput() UploadWatermarkResponseOutput
	ToUploadWatermarkResponseOutputWithContext(context.Context) UploadWatermarkResponseOutput
}

type UploadWatermarkResponseArgs struct {
	Logs    pulumi.StringPtrInput `pulumi:"logs"`
	Metrics pulumi.StringPtrInput `pulumi:"metrics"`
	Usages  pulumi.StringPtrInput `pulumi:"usages"`
}

func (UploadWatermarkResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*UploadWatermarkResponse)(nil)).Elem()
}

func (i UploadWatermarkResponseArgs) ToUploadWatermarkResponseOutput() UploadWatermarkResponseOutput {
	return i.ToUploadWatermarkResponseOutputWithContext(context.Background())
}

func (i UploadWatermarkResponseArgs) ToUploadWatermarkResponseOutputWithContext(ctx context.Context) UploadWatermarkResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UploadWatermarkResponseOutput)
}

func (i UploadWatermarkResponseArgs) ToUploadWatermarkResponsePtrOutput() UploadWatermarkResponsePtrOutput {
	return i.ToUploadWatermarkResponsePtrOutputWithContext(context.Background())
}

func (i UploadWatermarkResponseArgs) ToUploadWatermarkResponsePtrOutputWithContext(ctx context.Context) UploadWatermarkResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UploadWatermarkResponseOutput).ToUploadWatermarkResponsePtrOutputWithContext(ctx)
}









type UploadWatermarkResponsePtrInput interface {
	pulumi.Input

	ToUploadWatermarkResponsePtrOutput() UploadWatermarkResponsePtrOutput
	ToUploadWatermarkResponsePtrOutputWithContext(context.Context) UploadWatermarkResponsePtrOutput
}

type uploadWatermarkResponsePtrType UploadWatermarkResponseArgs

func UploadWatermarkResponsePtr(v *UploadWatermarkResponseArgs) UploadWatermarkResponsePtrInput {
	return (*uploadWatermarkResponsePtrType)(v)
}

func (*uploadWatermarkResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**UploadWatermarkResponse)(nil)).Elem()
}

func (i *uploadWatermarkResponsePtrType) ToUploadWatermarkResponsePtrOutput() UploadWatermarkResponsePtrOutput {
	return i.ToUploadWatermarkResponsePtrOutputWithContext(context.Background())
}

func (i *uploadWatermarkResponsePtrType) ToUploadWatermarkResponsePtrOutputWithContext(ctx context.Context) UploadWatermarkResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UploadWatermarkResponsePtrOutput)
}

type UploadWatermarkResponseOutput struct{ *pulumi.OutputState }

func (UploadWatermarkResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UploadWatermarkResponse)(nil)).Elem()
}

func (o UploadWatermarkResponseOutput) ToUploadWatermarkResponseOutput() UploadWatermarkResponseOutput {
	return o
}

func (o UploadWatermarkResponseOutput) ToUploadWatermarkResponseOutputWithContext(ctx context.Context) UploadWatermarkResponseOutput {
	return o
}

func (o UploadWatermarkResponseOutput) ToUploadWatermarkResponsePtrOutput() UploadWatermarkResponsePtrOutput {
	return o.ToUploadWatermarkResponsePtrOutputWithContext(context.Background())
}

func (o UploadWatermarkResponseOutput) ToUploadWatermarkResponsePtrOutputWithContext(ctx context.Context) UploadWatermarkResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v UploadWatermarkResponse) *UploadWatermarkResponse {
		return &v
	}).(UploadWatermarkResponsePtrOutput)
}

func (o UploadWatermarkResponseOutput) Logs() pulumi.StringPtrOutput {
	return o.ApplyT(func(v UploadWatermarkResponse) *string { return v.Logs }).(pulumi.StringPtrOutput)
}

func (o UploadWatermarkResponseOutput) Metrics() pulumi.StringPtrOutput {
	return o.ApplyT(func(v UploadWatermarkResponse) *string { return v.Metrics }).(pulumi.StringPtrOutput)
}

func (o UploadWatermarkResponseOutput) Usages() pulumi.StringPtrOutput {
	return o.ApplyT(func(v UploadWatermarkResponse) *string { return v.Usages }).(pulumi.StringPtrOutput)
}

type UploadWatermarkResponsePtrOutput struct{ *pulumi.OutputState }

func (UploadWatermarkResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**UploadWatermarkResponse)(nil)).Elem()
}

func (o UploadWatermarkResponsePtrOutput) ToUploadWatermarkResponsePtrOutput() UploadWatermarkResponsePtrOutput {
	return o
}

func (o UploadWatermarkResponsePtrOutput) ToUploadWatermarkResponsePtrOutputWithContext(ctx context.Context) UploadWatermarkResponsePtrOutput {
	return o
}

func (o UploadWatermarkResponsePtrOutput) Elem() UploadWatermarkResponseOutput {
	return o.ApplyT(func(v *UploadWatermarkResponse) UploadWatermarkResponse {
		if v != nil {
			return *v
		}
		var ret UploadWatermarkResponse
		return ret
	}).(UploadWatermarkResponseOutput)
}

func (o UploadWatermarkResponsePtrOutput) Logs() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UploadWatermarkResponse) *string {
		if v == nil {
			return nil
		}
		return v.Logs
	}).(pulumi.StringPtrOutput)
}

func (o UploadWatermarkResponsePtrOutput) Metrics() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UploadWatermarkResponse) *string {
		if v == nil {
			return nil
		}
		return v.Metrics
	}).(pulumi.StringPtrOutput)
}

func (o UploadWatermarkResponsePtrOutput) Usages() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UploadWatermarkResponse) *string {
		if v == nil {
			return nil
		}
		return v.Usages
	}).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*BasicLoginInformationInput)(nil)).Elem(), BasicLoginInformationArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*BasicLoginInformationPtrInput)(nil)).Elem(), BasicLoginInformationArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*BasicLoginInformationResponseInput)(nil)).Elem(), BasicLoginInformationResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*BasicLoginInformationResponsePtrInput)(nil)).Elem(), BasicLoginInformationResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*DataControllerPropertiesInput)(nil)).Elem(), DataControllerPropertiesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*DataControllerPropertiesPtrInput)(nil)).Elem(), DataControllerPropertiesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*DataControllerPropertiesResponseInput)(nil)).Elem(), DataControllerPropertiesResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*DataControllerPropertiesResponsePtrInput)(nil)).Elem(), DataControllerPropertiesResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ExtendedLocationInput)(nil)).Elem(), ExtendedLocationArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ExtendedLocationPtrInput)(nil)).Elem(), ExtendedLocationArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ExtendedLocationResponseInput)(nil)).Elem(), ExtendedLocationResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ExtendedLocationResponsePtrInput)(nil)).Elem(), ExtendedLocationResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*LogAnalyticsWorkspaceConfigInput)(nil)).Elem(), LogAnalyticsWorkspaceConfigArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*LogAnalyticsWorkspaceConfigPtrInput)(nil)).Elem(), LogAnalyticsWorkspaceConfigArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*LogAnalyticsWorkspaceConfigResponseInput)(nil)).Elem(), LogAnalyticsWorkspaceConfigResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*LogAnalyticsWorkspaceConfigResponsePtrInput)(nil)).Elem(), LogAnalyticsWorkspaceConfigResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*OnPremisePropertyInput)(nil)).Elem(), OnPremisePropertyArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*OnPremisePropertyPtrInput)(nil)).Elem(), OnPremisePropertyArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*OnPremisePropertyResponseInput)(nil)).Elem(), OnPremisePropertyResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*OnPremisePropertyResponsePtrInput)(nil)).Elem(), OnPremisePropertyResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*PostgresInstancePropertiesInput)(nil)).Elem(), PostgresInstancePropertiesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*PostgresInstancePropertiesPtrInput)(nil)).Elem(), PostgresInstancePropertiesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*PostgresInstancePropertiesResponseInput)(nil)).Elem(), PostgresInstancePropertiesResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*PostgresInstancePropertiesResponsePtrInput)(nil)).Elem(), PostgresInstancePropertiesResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*PostgresInstanceSkuInput)(nil)).Elem(), PostgresInstanceSkuArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*PostgresInstanceSkuPtrInput)(nil)).Elem(), PostgresInstanceSkuArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*PostgresInstanceSkuResponseInput)(nil)).Elem(), PostgresInstanceSkuResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*PostgresInstanceSkuResponsePtrInput)(nil)).Elem(), PostgresInstanceSkuResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SqlManagedInstancePropertiesInput)(nil)).Elem(), SqlManagedInstancePropertiesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SqlManagedInstancePropertiesPtrInput)(nil)).Elem(), SqlManagedInstancePropertiesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SqlManagedInstancePropertiesResponseInput)(nil)).Elem(), SqlManagedInstancePropertiesResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SqlManagedInstancePropertiesResponsePtrInput)(nil)).Elem(), SqlManagedInstancePropertiesResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SqlManagedInstanceSkuInput)(nil)).Elem(), SqlManagedInstanceSkuArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SqlManagedInstanceSkuPtrInput)(nil)).Elem(), SqlManagedInstanceSkuArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SqlManagedInstanceSkuResponseInput)(nil)).Elem(), SqlManagedInstanceSkuResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SqlManagedInstanceSkuResponsePtrInput)(nil)).Elem(), SqlManagedInstanceSkuResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SqlServerInstancePropertiesInput)(nil)).Elem(), SqlServerInstancePropertiesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SqlServerInstancePropertiesPtrInput)(nil)).Elem(), SqlServerInstancePropertiesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SqlServerInstancePropertiesResponseInput)(nil)).Elem(), SqlServerInstancePropertiesResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SqlServerInstancePropertiesResponsePtrInput)(nil)).Elem(), SqlServerInstancePropertiesResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemDataResponseInput)(nil)).Elem(), SystemDataResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemDataResponsePtrInput)(nil)).Elem(), SystemDataResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*UploadServicePrincipalInput)(nil)).Elem(), UploadServicePrincipalArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*UploadServicePrincipalPtrInput)(nil)).Elem(), UploadServicePrincipalArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*UploadServicePrincipalResponseInput)(nil)).Elem(), UploadServicePrincipalResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*UploadServicePrincipalResponsePtrInput)(nil)).Elem(), UploadServicePrincipalResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*UploadWatermarkInput)(nil)).Elem(), UploadWatermarkArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*UploadWatermarkPtrInput)(nil)).Elem(), UploadWatermarkArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*UploadWatermarkResponseInput)(nil)).Elem(), UploadWatermarkResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*UploadWatermarkResponsePtrInput)(nil)).Elem(), UploadWatermarkResponseArgs{})
	pulumi.RegisterOutputType(BasicLoginInformationOutput{})
	pulumi.RegisterOutputType(BasicLoginInformationPtrOutput{})
	pulumi.RegisterOutputType(BasicLoginInformationResponseOutput{})
	pulumi.RegisterOutputType(BasicLoginInformationResponsePtrOutput{})
	pulumi.RegisterOutputType(DataControllerPropertiesOutput{})
	pulumi.RegisterOutputType(DataControllerPropertiesPtrOutput{})
	pulumi.RegisterOutputType(DataControllerPropertiesResponseOutput{})
	pulumi.RegisterOutputType(DataControllerPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(ExtendedLocationOutput{})
	pulumi.RegisterOutputType(ExtendedLocationPtrOutput{})
	pulumi.RegisterOutputType(ExtendedLocationResponseOutput{})
	pulumi.RegisterOutputType(ExtendedLocationResponsePtrOutput{})
	pulumi.RegisterOutputType(LogAnalyticsWorkspaceConfigOutput{})
	pulumi.RegisterOutputType(LogAnalyticsWorkspaceConfigPtrOutput{})
	pulumi.RegisterOutputType(LogAnalyticsWorkspaceConfigResponseOutput{})
	pulumi.RegisterOutputType(LogAnalyticsWorkspaceConfigResponsePtrOutput{})
	pulumi.RegisterOutputType(OnPremisePropertyOutput{})
	pulumi.RegisterOutputType(OnPremisePropertyPtrOutput{})
	pulumi.RegisterOutputType(OnPremisePropertyResponseOutput{})
	pulumi.RegisterOutputType(OnPremisePropertyResponsePtrOutput{})
	pulumi.RegisterOutputType(PostgresInstancePropertiesOutput{})
	pulumi.RegisterOutputType(PostgresInstancePropertiesPtrOutput{})
	pulumi.RegisterOutputType(PostgresInstancePropertiesResponseOutput{})
	pulumi.RegisterOutputType(PostgresInstancePropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(PostgresInstanceSkuOutput{})
	pulumi.RegisterOutputType(PostgresInstanceSkuPtrOutput{})
	pulumi.RegisterOutputType(PostgresInstanceSkuResponseOutput{})
	pulumi.RegisterOutputType(PostgresInstanceSkuResponsePtrOutput{})
	pulumi.RegisterOutputType(SqlManagedInstancePropertiesOutput{})
	pulumi.RegisterOutputType(SqlManagedInstancePropertiesPtrOutput{})
	pulumi.RegisterOutputType(SqlManagedInstancePropertiesResponseOutput{})
	pulumi.RegisterOutputType(SqlManagedInstancePropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(SqlManagedInstanceSkuOutput{})
	pulumi.RegisterOutputType(SqlManagedInstanceSkuPtrOutput{})
	pulumi.RegisterOutputType(SqlManagedInstanceSkuResponseOutput{})
	pulumi.RegisterOutputType(SqlManagedInstanceSkuResponsePtrOutput{})
	pulumi.RegisterOutputType(SqlServerInstancePropertiesOutput{})
	pulumi.RegisterOutputType(SqlServerInstancePropertiesPtrOutput{})
	pulumi.RegisterOutputType(SqlServerInstancePropertiesResponseOutput{})
	pulumi.RegisterOutputType(SqlServerInstancePropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(SystemDataResponseOutput{})
	pulumi.RegisterOutputType(SystemDataResponsePtrOutput{})
	pulumi.RegisterOutputType(UploadServicePrincipalOutput{})
	pulumi.RegisterOutputType(UploadServicePrincipalPtrOutput{})
	pulumi.RegisterOutputType(UploadServicePrincipalResponseOutput{})
	pulumi.RegisterOutputType(UploadServicePrincipalResponsePtrOutput{})
	pulumi.RegisterOutputType(UploadWatermarkOutput{})
	pulumi.RegisterOutputType(UploadWatermarkPtrOutput{})
	pulumi.RegisterOutputType(UploadWatermarkResponseOutput{})
	pulumi.RegisterOutputType(UploadWatermarkResponsePtrOutput{})
}
