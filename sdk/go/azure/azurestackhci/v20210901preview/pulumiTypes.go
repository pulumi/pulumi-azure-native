


package v20210901preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ClusterDesiredProperties struct {
	DiagnosticLevel           *string `pulumi:"diagnosticLevel"`
	WindowsServerSubscription *string `pulumi:"windowsServerSubscription"`
}





type ClusterDesiredPropertiesInput interface {
	pulumi.Input

	ToClusterDesiredPropertiesOutput() ClusterDesiredPropertiesOutput
	ToClusterDesiredPropertiesOutputWithContext(context.Context) ClusterDesiredPropertiesOutput
}

type ClusterDesiredPropertiesArgs struct {
	DiagnosticLevel           pulumi.StringPtrInput `pulumi:"diagnosticLevel"`
	WindowsServerSubscription pulumi.StringPtrInput `pulumi:"windowsServerSubscription"`
}

func (ClusterDesiredPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ClusterDesiredProperties)(nil)).Elem()
}

func (i ClusterDesiredPropertiesArgs) ToClusterDesiredPropertiesOutput() ClusterDesiredPropertiesOutput {
	return i.ToClusterDesiredPropertiesOutputWithContext(context.Background())
}

func (i ClusterDesiredPropertiesArgs) ToClusterDesiredPropertiesOutputWithContext(ctx context.Context) ClusterDesiredPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterDesiredPropertiesOutput)
}

func (i ClusterDesiredPropertiesArgs) ToClusterDesiredPropertiesPtrOutput() ClusterDesiredPropertiesPtrOutput {
	return i.ToClusterDesiredPropertiesPtrOutputWithContext(context.Background())
}

func (i ClusterDesiredPropertiesArgs) ToClusterDesiredPropertiesPtrOutputWithContext(ctx context.Context) ClusterDesiredPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterDesiredPropertiesOutput).ToClusterDesiredPropertiesPtrOutputWithContext(ctx)
}









type ClusterDesiredPropertiesPtrInput interface {
	pulumi.Input

	ToClusterDesiredPropertiesPtrOutput() ClusterDesiredPropertiesPtrOutput
	ToClusterDesiredPropertiesPtrOutputWithContext(context.Context) ClusterDesiredPropertiesPtrOutput
}

type clusterDesiredPropertiesPtrType ClusterDesiredPropertiesArgs

func ClusterDesiredPropertiesPtr(v *ClusterDesiredPropertiesArgs) ClusterDesiredPropertiesPtrInput {
	return (*clusterDesiredPropertiesPtrType)(v)
}

func (*clusterDesiredPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ClusterDesiredProperties)(nil)).Elem()
}

func (i *clusterDesiredPropertiesPtrType) ToClusterDesiredPropertiesPtrOutput() ClusterDesiredPropertiesPtrOutput {
	return i.ToClusterDesiredPropertiesPtrOutputWithContext(context.Background())
}

func (i *clusterDesiredPropertiesPtrType) ToClusterDesiredPropertiesPtrOutputWithContext(ctx context.Context) ClusterDesiredPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterDesiredPropertiesPtrOutput)
}

type ClusterDesiredPropertiesOutput struct{ *pulumi.OutputState }

func (ClusterDesiredPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ClusterDesiredProperties)(nil)).Elem()
}

func (o ClusterDesiredPropertiesOutput) ToClusterDesiredPropertiesOutput() ClusterDesiredPropertiesOutput {
	return o
}

func (o ClusterDesiredPropertiesOutput) ToClusterDesiredPropertiesOutputWithContext(ctx context.Context) ClusterDesiredPropertiesOutput {
	return o
}

func (o ClusterDesiredPropertiesOutput) ToClusterDesiredPropertiesPtrOutput() ClusterDesiredPropertiesPtrOutput {
	return o.ToClusterDesiredPropertiesPtrOutputWithContext(context.Background())
}

func (o ClusterDesiredPropertiesOutput) ToClusterDesiredPropertiesPtrOutputWithContext(ctx context.Context) ClusterDesiredPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ClusterDesiredProperties) *ClusterDesiredProperties {
		return &v
	}).(ClusterDesiredPropertiesPtrOutput)
}

func (o ClusterDesiredPropertiesOutput) DiagnosticLevel() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ClusterDesiredProperties) *string { return v.DiagnosticLevel }).(pulumi.StringPtrOutput)
}

func (o ClusterDesiredPropertiesOutput) WindowsServerSubscription() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ClusterDesiredProperties) *string { return v.WindowsServerSubscription }).(pulumi.StringPtrOutput)
}

type ClusterDesiredPropertiesPtrOutput struct{ *pulumi.OutputState }

func (ClusterDesiredPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ClusterDesiredProperties)(nil)).Elem()
}

func (o ClusterDesiredPropertiesPtrOutput) ToClusterDesiredPropertiesPtrOutput() ClusterDesiredPropertiesPtrOutput {
	return o
}

func (o ClusterDesiredPropertiesPtrOutput) ToClusterDesiredPropertiesPtrOutputWithContext(ctx context.Context) ClusterDesiredPropertiesPtrOutput {
	return o
}

func (o ClusterDesiredPropertiesPtrOutput) Elem() ClusterDesiredPropertiesOutput {
	return o.ApplyT(func(v *ClusterDesiredProperties) ClusterDesiredProperties {
		if v != nil {
			return *v
		}
		var ret ClusterDesiredProperties
		return ret
	}).(ClusterDesiredPropertiesOutput)
}

func (o ClusterDesiredPropertiesPtrOutput) DiagnosticLevel() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ClusterDesiredProperties) *string {
		if v == nil {
			return nil
		}
		return v.DiagnosticLevel
	}).(pulumi.StringPtrOutput)
}

func (o ClusterDesiredPropertiesPtrOutput) WindowsServerSubscription() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ClusterDesiredProperties) *string {
		if v == nil {
			return nil
		}
		return v.WindowsServerSubscription
	}).(pulumi.StringPtrOutput)
}

type ClusterDesiredPropertiesResponse struct {
	DiagnosticLevel           *string `pulumi:"diagnosticLevel"`
	WindowsServerSubscription *string `pulumi:"windowsServerSubscription"`
}

type ClusterDesiredPropertiesResponseOutput struct{ *pulumi.OutputState }

func (ClusterDesiredPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ClusterDesiredPropertiesResponse)(nil)).Elem()
}

func (o ClusterDesiredPropertiesResponseOutput) ToClusterDesiredPropertiesResponseOutput() ClusterDesiredPropertiesResponseOutput {
	return o
}

func (o ClusterDesiredPropertiesResponseOutput) ToClusterDesiredPropertiesResponseOutputWithContext(ctx context.Context) ClusterDesiredPropertiesResponseOutput {
	return o
}

func (o ClusterDesiredPropertiesResponseOutput) DiagnosticLevel() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ClusterDesiredPropertiesResponse) *string { return v.DiagnosticLevel }).(pulumi.StringPtrOutput)
}

func (o ClusterDesiredPropertiesResponseOutput) WindowsServerSubscription() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ClusterDesiredPropertiesResponse) *string { return v.WindowsServerSubscription }).(pulumi.StringPtrOutput)
}

type ClusterDesiredPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (ClusterDesiredPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ClusterDesiredPropertiesResponse)(nil)).Elem()
}

func (o ClusterDesiredPropertiesResponsePtrOutput) ToClusterDesiredPropertiesResponsePtrOutput() ClusterDesiredPropertiesResponsePtrOutput {
	return o
}

func (o ClusterDesiredPropertiesResponsePtrOutput) ToClusterDesiredPropertiesResponsePtrOutputWithContext(ctx context.Context) ClusterDesiredPropertiesResponsePtrOutput {
	return o
}

func (o ClusterDesiredPropertiesResponsePtrOutput) Elem() ClusterDesiredPropertiesResponseOutput {
	return o.ApplyT(func(v *ClusterDesiredPropertiesResponse) ClusterDesiredPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret ClusterDesiredPropertiesResponse
		return ret
	}).(ClusterDesiredPropertiesResponseOutput)
}

func (o ClusterDesiredPropertiesResponsePtrOutput) DiagnosticLevel() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ClusterDesiredPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.DiagnosticLevel
	}).(pulumi.StringPtrOutput)
}

func (o ClusterDesiredPropertiesResponsePtrOutput) WindowsServerSubscription() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ClusterDesiredPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.WindowsServerSubscription
	}).(pulumi.StringPtrOutput)
}

type ClusterNodeResponse struct {
	CoreCount                 float64 `pulumi:"coreCount"`
	Id                        float64 `pulumi:"id"`
	Manufacturer              string  `pulumi:"manufacturer"`
	MemoryInGiB               float64 `pulumi:"memoryInGiB"`
	Model                     string  `pulumi:"model"`
	Name                      string  `pulumi:"name"`
	OsName                    string  `pulumi:"osName"`
	OsVersion                 string  `pulumi:"osVersion"`
	SerialNumber              string  `pulumi:"serialNumber"`
	WindowsServerSubscription string  `pulumi:"windowsServerSubscription"`
}

type ClusterNodeResponseOutput struct{ *pulumi.OutputState }

func (ClusterNodeResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ClusterNodeResponse)(nil)).Elem()
}

func (o ClusterNodeResponseOutput) ToClusterNodeResponseOutput() ClusterNodeResponseOutput {
	return o
}

func (o ClusterNodeResponseOutput) ToClusterNodeResponseOutputWithContext(ctx context.Context) ClusterNodeResponseOutput {
	return o
}

func (o ClusterNodeResponseOutput) CoreCount() pulumi.Float64Output {
	return o.ApplyT(func(v ClusterNodeResponse) float64 { return v.CoreCount }).(pulumi.Float64Output)
}

func (o ClusterNodeResponseOutput) Id() pulumi.Float64Output {
	return o.ApplyT(func(v ClusterNodeResponse) float64 { return v.Id }).(pulumi.Float64Output)
}

func (o ClusterNodeResponseOutput) Manufacturer() pulumi.StringOutput {
	return o.ApplyT(func(v ClusterNodeResponse) string { return v.Manufacturer }).(pulumi.StringOutput)
}

func (o ClusterNodeResponseOutput) MemoryInGiB() pulumi.Float64Output {
	return o.ApplyT(func(v ClusterNodeResponse) float64 { return v.MemoryInGiB }).(pulumi.Float64Output)
}

func (o ClusterNodeResponseOutput) Model() pulumi.StringOutput {
	return o.ApplyT(func(v ClusterNodeResponse) string { return v.Model }).(pulumi.StringOutput)
}

func (o ClusterNodeResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ClusterNodeResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o ClusterNodeResponseOutput) OsName() pulumi.StringOutput {
	return o.ApplyT(func(v ClusterNodeResponse) string { return v.OsName }).(pulumi.StringOutput)
}

func (o ClusterNodeResponseOutput) OsVersion() pulumi.StringOutput {
	return o.ApplyT(func(v ClusterNodeResponse) string { return v.OsVersion }).(pulumi.StringOutput)
}

func (o ClusterNodeResponseOutput) SerialNumber() pulumi.StringOutput {
	return o.ApplyT(func(v ClusterNodeResponse) string { return v.SerialNumber }).(pulumi.StringOutput)
}

func (o ClusterNodeResponseOutput) WindowsServerSubscription() pulumi.StringOutput {
	return o.ApplyT(func(v ClusterNodeResponse) string { return v.WindowsServerSubscription }).(pulumi.StringOutput)
}

type ClusterNodeResponseArrayOutput struct{ *pulumi.OutputState }

func (ClusterNodeResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ClusterNodeResponse)(nil)).Elem()
}

func (o ClusterNodeResponseArrayOutput) ToClusterNodeResponseArrayOutput() ClusterNodeResponseArrayOutput {
	return o
}

func (o ClusterNodeResponseArrayOutput) ToClusterNodeResponseArrayOutputWithContext(ctx context.Context) ClusterNodeResponseArrayOutput {
	return o
}

func (o ClusterNodeResponseArrayOutput) Index(i pulumi.IntInput) ClusterNodeResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ClusterNodeResponse {
		return vs[0].([]ClusterNodeResponse)[vs[1].(int)]
	}).(ClusterNodeResponseOutput)
}

type ClusterReportedPropertiesResponse struct {
	ClusterId       string                `pulumi:"clusterId"`
	ClusterName     string                `pulumi:"clusterName"`
	ClusterVersion  string                `pulumi:"clusterVersion"`
	DiagnosticLevel *string               `pulumi:"diagnosticLevel"`
	ImdsAttestation string                `pulumi:"imdsAttestation"`
	LastUpdated     string                `pulumi:"lastUpdated"`
	Nodes           []ClusterNodeResponse `pulumi:"nodes"`
}

type ClusterReportedPropertiesResponseOutput struct{ *pulumi.OutputState }

func (ClusterReportedPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ClusterReportedPropertiesResponse)(nil)).Elem()
}

func (o ClusterReportedPropertiesResponseOutput) ToClusterReportedPropertiesResponseOutput() ClusterReportedPropertiesResponseOutput {
	return o
}

func (o ClusterReportedPropertiesResponseOutput) ToClusterReportedPropertiesResponseOutputWithContext(ctx context.Context) ClusterReportedPropertiesResponseOutput {
	return o
}

func (o ClusterReportedPropertiesResponseOutput) ClusterId() pulumi.StringOutput {
	return o.ApplyT(func(v ClusterReportedPropertiesResponse) string { return v.ClusterId }).(pulumi.StringOutput)
}

func (o ClusterReportedPropertiesResponseOutput) ClusterName() pulumi.StringOutput {
	return o.ApplyT(func(v ClusterReportedPropertiesResponse) string { return v.ClusterName }).(pulumi.StringOutput)
}

func (o ClusterReportedPropertiesResponseOutput) ClusterVersion() pulumi.StringOutput {
	return o.ApplyT(func(v ClusterReportedPropertiesResponse) string { return v.ClusterVersion }).(pulumi.StringOutput)
}

func (o ClusterReportedPropertiesResponseOutput) DiagnosticLevel() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ClusterReportedPropertiesResponse) *string { return v.DiagnosticLevel }).(pulumi.StringPtrOutput)
}

func (o ClusterReportedPropertiesResponseOutput) ImdsAttestation() pulumi.StringOutput {
	return o.ApplyT(func(v ClusterReportedPropertiesResponse) string { return v.ImdsAttestation }).(pulumi.StringOutput)
}

func (o ClusterReportedPropertiesResponseOutput) LastUpdated() pulumi.StringOutput {
	return o.ApplyT(func(v ClusterReportedPropertiesResponse) string { return v.LastUpdated }).(pulumi.StringOutput)
}

func (o ClusterReportedPropertiesResponseOutput) Nodes() ClusterNodeResponseArrayOutput {
	return o.ApplyT(func(v ClusterReportedPropertiesResponse) []ClusterNodeResponse { return v.Nodes }).(ClusterNodeResponseArrayOutput)
}

type ErrorAdditionalInfoResponse struct {
	Info interface{} `pulumi:"info"`
	Type string      `pulumi:"type"`
}

type ErrorAdditionalInfoResponseOutput struct{ *pulumi.OutputState }

func (ErrorAdditionalInfoResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ErrorAdditionalInfoResponse)(nil)).Elem()
}

func (o ErrorAdditionalInfoResponseOutput) ToErrorAdditionalInfoResponseOutput() ErrorAdditionalInfoResponseOutput {
	return o
}

func (o ErrorAdditionalInfoResponseOutput) ToErrorAdditionalInfoResponseOutputWithContext(ctx context.Context) ErrorAdditionalInfoResponseOutput {
	return o
}

func (o ErrorAdditionalInfoResponseOutput) Info() pulumi.AnyOutput {
	return o.ApplyT(func(v ErrorAdditionalInfoResponse) interface{} { return v.Info }).(pulumi.AnyOutput)
}

func (o ErrorAdditionalInfoResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ErrorAdditionalInfoResponse) string { return v.Type }).(pulumi.StringOutput)
}

type ErrorAdditionalInfoResponseArrayOutput struct{ *pulumi.OutputState }

func (ErrorAdditionalInfoResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ErrorAdditionalInfoResponse)(nil)).Elem()
}

func (o ErrorAdditionalInfoResponseArrayOutput) ToErrorAdditionalInfoResponseArrayOutput() ErrorAdditionalInfoResponseArrayOutput {
	return o
}

func (o ErrorAdditionalInfoResponseArrayOutput) ToErrorAdditionalInfoResponseArrayOutputWithContext(ctx context.Context) ErrorAdditionalInfoResponseArrayOutput {
	return o
}

func (o ErrorAdditionalInfoResponseArrayOutput) Index(i pulumi.IntInput) ErrorAdditionalInfoResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ErrorAdditionalInfoResponse {
		return vs[0].([]ErrorAdditionalInfoResponse)[vs[1].(int)]
	}).(ErrorAdditionalInfoResponseOutput)
}

type ErrorDetailResponse struct {
	AdditionalInfo []ErrorAdditionalInfoResponse `pulumi:"additionalInfo"`
	Code           string                        `pulumi:"code"`
	Details        []ErrorDetailResponse         `pulumi:"details"`
	Message        string                        `pulumi:"message"`
	Target         string                        `pulumi:"target"`
}

type ErrorDetailResponseOutput struct{ *pulumi.OutputState }

func (ErrorDetailResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ErrorDetailResponse)(nil)).Elem()
}

func (o ErrorDetailResponseOutput) ToErrorDetailResponseOutput() ErrorDetailResponseOutput {
	return o
}

func (o ErrorDetailResponseOutput) ToErrorDetailResponseOutputWithContext(ctx context.Context) ErrorDetailResponseOutput {
	return o
}

func (o ErrorDetailResponseOutput) AdditionalInfo() ErrorAdditionalInfoResponseArrayOutput {
	return o.ApplyT(func(v ErrorDetailResponse) []ErrorAdditionalInfoResponse { return v.AdditionalInfo }).(ErrorAdditionalInfoResponseArrayOutput)
}

func (o ErrorDetailResponseOutput) Code() pulumi.StringOutput {
	return o.ApplyT(func(v ErrorDetailResponse) string { return v.Code }).(pulumi.StringOutput)
}

func (o ErrorDetailResponseOutput) Details() ErrorDetailResponseArrayOutput {
	return o.ApplyT(func(v ErrorDetailResponse) []ErrorDetailResponse { return v.Details }).(ErrorDetailResponseArrayOutput)
}

func (o ErrorDetailResponseOutput) Message() pulumi.StringOutput {
	return o.ApplyT(func(v ErrorDetailResponse) string { return v.Message }).(pulumi.StringOutput)
}

func (o ErrorDetailResponseOutput) Target() pulumi.StringOutput {
	return o.ApplyT(func(v ErrorDetailResponse) string { return v.Target }).(pulumi.StringOutput)
}

type ErrorDetailResponseArrayOutput struct{ *pulumi.OutputState }

func (ErrorDetailResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ErrorDetailResponse)(nil)).Elem()
}

func (o ErrorDetailResponseArrayOutput) ToErrorDetailResponseArrayOutput() ErrorDetailResponseArrayOutput {
	return o
}

func (o ErrorDetailResponseArrayOutput) ToErrorDetailResponseArrayOutputWithContext(ctx context.Context) ErrorDetailResponseArrayOutput {
	return o
}

func (o ErrorDetailResponseArrayOutput) Index(i pulumi.IntInput) ErrorDetailResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ErrorDetailResponse {
		return vs[0].([]ErrorDetailResponse)[vs[1].(int)]
	}).(ErrorDetailResponseOutput)
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

type GalleryImageIdentifier struct {
	Offer     string `pulumi:"offer"`
	Publisher string `pulumi:"publisher"`
	Sku       string `pulumi:"sku"`
}





type GalleryImageIdentifierInput interface {
	pulumi.Input

	ToGalleryImageIdentifierOutput() GalleryImageIdentifierOutput
	ToGalleryImageIdentifierOutputWithContext(context.Context) GalleryImageIdentifierOutput
}

type GalleryImageIdentifierArgs struct {
	Offer     pulumi.StringInput `pulumi:"offer"`
	Publisher pulumi.StringInput `pulumi:"publisher"`
	Sku       pulumi.StringInput `pulumi:"sku"`
}

func (GalleryImageIdentifierArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GalleryImageIdentifier)(nil)).Elem()
}

func (i GalleryImageIdentifierArgs) ToGalleryImageIdentifierOutput() GalleryImageIdentifierOutput {
	return i.ToGalleryImageIdentifierOutputWithContext(context.Background())
}

func (i GalleryImageIdentifierArgs) ToGalleryImageIdentifierOutputWithContext(ctx context.Context) GalleryImageIdentifierOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GalleryImageIdentifierOutput)
}

func (i GalleryImageIdentifierArgs) ToGalleryImageIdentifierPtrOutput() GalleryImageIdentifierPtrOutput {
	return i.ToGalleryImageIdentifierPtrOutputWithContext(context.Background())
}

func (i GalleryImageIdentifierArgs) ToGalleryImageIdentifierPtrOutputWithContext(ctx context.Context) GalleryImageIdentifierPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GalleryImageIdentifierOutput).ToGalleryImageIdentifierPtrOutputWithContext(ctx)
}









type GalleryImageIdentifierPtrInput interface {
	pulumi.Input

	ToGalleryImageIdentifierPtrOutput() GalleryImageIdentifierPtrOutput
	ToGalleryImageIdentifierPtrOutputWithContext(context.Context) GalleryImageIdentifierPtrOutput
}

type galleryImageIdentifierPtrType GalleryImageIdentifierArgs

func GalleryImageIdentifierPtr(v *GalleryImageIdentifierArgs) GalleryImageIdentifierPtrInput {
	return (*galleryImageIdentifierPtrType)(v)
}

func (*galleryImageIdentifierPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**GalleryImageIdentifier)(nil)).Elem()
}

func (i *galleryImageIdentifierPtrType) ToGalleryImageIdentifierPtrOutput() GalleryImageIdentifierPtrOutput {
	return i.ToGalleryImageIdentifierPtrOutputWithContext(context.Background())
}

func (i *galleryImageIdentifierPtrType) ToGalleryImageIdentifierPtrOutputWithContext(ctx context.Context) GalleryImageIdentifierPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GalleryImageIdentifierPtrOutput)
}

type GalleryImageIdentifierOutput struct{ *pulumi.OutputState }

func (GalleryImageIdentifierOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GalleryImageIdentifier)(nil)).Elem()
}

func (o GalleryImageIdentifierOutput) ToGalleryImageIdentifierOutput() GalleryImageIdentifierOutput {
	return o
}

func (o GalleryImageIdentifierOutput) ToGalleryImageIdentifierOutputWithContext(ctx context.Context) GalleryImageIdentifierOutput {
	return o
}

func (o GalleryImageIdentifierOutput) ToGalleryImageIdentifierPtrOutput() GalleryImageIdentifierPtrOutput {
	return o.ToGalleryImageIdentifierPtrOutputWithContext(context.Background())
}

func (o GalleryImageIdentifierOutput) ToGalleryImageIdentifierPtrOutputWithContext(ctx context.Context) GalleryImageIdentifierPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v GalleryImageIdentifier) *GalleryImageIdentifier {
		return &v
	}).(GalleryImageIdentifierPtrOutput)
}

func (o GalleryImageIdentifierOutput) Offer() pulumi.StringOutput {
	return o.ApplyT(func(v GalleryImageIdentifier) string { return v.Offer }).(pulumi.StringOutput)
}

func (o GalleryImageIdentifierOutput) Publisher() pulumi.StringOutput {
	return o.ApplyT(func(v GalleryImageIdentifier) string { return v.Publisher }).(pulumi.StringOutput)
}

func (o GalleryImageIdentifierOutput) Sku() pulumi.StringOutput {
	return o.ApplyT(func(v GalleryImageIdentifier) string { return v.Sku }).(pulumi.StringOutput)
}

type GalleryImageIdentifierPtrOutput struct{ *pulumi.OutputState }

func (GalleryImageIdentifierPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GalleryImageIdentifier)(nil)).Elem()
}

func (o GalleryImageIdentifierPtrOutput) ToGalleryImageIdentifierPtrOutput() GalleryImageIdentifierPtrOutput {
	return o
}

func (o GalleryImageIdentifierPtrOutput) ToGalleryImageIdentifierPtrOutputWithContext(ctx context.Context) GalleryImageIdentifierPtrOutput {
	return o
}

func (o GalleryImageIdentifierPtrOutput) Elem() GalleryImageIdentifierOutput {
	return o.ApplyT(func(v *GalleryImageIdentifier) GalleryImageIdentifier {
		if v != nil {
			return *v
		}
		var ret GalleryImageIdentifier
		return ret
	}).(GalleryImageIdentifierOutput)
}

func (o GalleryImageIdentifierPtrOutput) Offer() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GalleryImageIdentifier) *string {
		if v == nil {
			return nil
		}
		return &v.Offer
	}).(pulumi.StringPtrOutput)
}

func (o GalleryImageIdentifierPtrOutput) Publisher() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GalleryImageIdentifier) *string {
		if v == nil {
			return nil
		}
		return &v.Publisher
	}).(pulumi.StringPtrOutput)
}

func (o GalleryImageIdentifierPtrOutput) Sku() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GalleryImageIdentifier) *string {
		if v == nil {
			return nil
		}
		return &v.Sku
	}).(pulumi.StringPtrOutput)
}

type GalleryImageIdentifierResponse struct {
	Offer     string `pulumi:"offer"`
	Publisher string `pulumi:"publisher"`
	Sku       string `pulumi:"sku"`
}

type GalleryImageIdentifierResponseOutput struct{ *pulumi.OutputState }

func (GalleryImageIdentifierResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GalleryImageIdentifierResponse)(nil)).Elem()
}

func (o GalleryImageIdentifierResponseOutput) ToGalleryImageIdentifierResponseOutput() GalleryImageIdentifierResponseOutput {
	return o
}

func (o GalleryImageIdentifierResponseOutput) ToGalleryImageIdentifierResponseOutputWithContext(ctx context.Context) GalleryImageIdentifierResponseOutput {
	return o
}

func (o GalleryImageIdentifierResponseOutput) Offer() pulumi.StringOutput {
	return o.ApplyT(func(v GalleryImageIdentifierResponse) string { return v.Offer }).(pulumi.StringOutput)
}

func (o GalleryImageIdentifierResponseOutput) Publisher() pulumi.StringOutput {
	return o.ApplyT(func(v GalleryImageIdentifierResponse) string { return v.Publisher }).(pulumi.StringOutput)
}

func (o GalleryImageIdentifierResponseOutput) Sku() pulumi.StringOutput {
	return o.ApplyT(func(v GalleryImageIdentifierResponse) string { return v.Sku }).(pulumi.StringOutput)
}

type GalleryImageIdentifierResponsePtrOutput struct{ *pulumi.OutputState }

func (GalleryImageIdentifierResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GalleryImageIdentifierResponse)(nil)).Elem()
}

func (o GalleryImageIdentifierResponsePtrOutput) ToGalleryImageIdentifierResponsePtrOutput() GalleryImageIdentifierResponsePtrOutput {
	return o
}

func (o GalleryImageIdentifierResponsePtrOutput) ToGalleryImageIdentifierResponsePtrOutputWithContext(ctx context.Context) GalleryImageIdentifierResponsePtrOutput {
	return o
}

func (o GalleryImageIdentifierResponsePtrOutput) Elem() GalleryImageIdentifierResponseOutput {
	return o.ApplyT(func(v *GalleryImageIdentifierResponse) GalleryImageIdentifierResponse {
		if v != nil {
			return *v
		}
		var ret GalleryImageIdentifierResponse
		return ret
	}).(GalleryImageIdentifierResponseOutput)
}

func (o GalleryImageIdentifierResponsePtrOutput) Offer() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GalleryImageIdentifierResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Offer
	}).(pulumi.StringPtrOutput)
}

func (o GalleryImageIdentifierResponsePtrOutput) Publisher() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GalleryImageIdentifierResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Publisher
	}).(pulumi.StringPtrOutput)
}

func (o GalleryImageIdentifierResponsePtrOutput) Sku() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GalleryImageIdentifierResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Sku
	}).(pulumi.StringPtrOutput)
}

type GalleryImageStatusResponse struct {
	DownloadStatus     *GalleryImageStatusResponseDownloadStatus     `pulumi:"downloadStatus"`
	ErrorCode          *string                                       `pulumi:"errorCode"`
	ErrorMessage       *string                                       `pulumi:"errorMessage"`
	ProgressPercentage *float64                                      `pulumi:"progressPercentage"`
	ProvisioningStatus *GalleryImageStatusResponseProvisioningStatus `pulumi:"provisioningStatus"`
}

type GalleryImageStatusResponseOutput struct{ *pulumi.OutputState }

func (GalleryImageStatusResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GalleryImageStatusResponse)(nil)).Elem()
}

func (o GalleryImageStatusResponseOutput) ToGalleryImageStatusResponseOutput() GalleryImageStatusResponseOutput {
	return o
}

func (o GalleryImageStatusResponseOutput) ToGalleryImageStatusResponseOutputWithContext(ctx context.Context) GalleryImageStatusResponseOutput {
	return o
}

func (o GalleryImageStatusResponseOutput) DownloadStatus() GalleryImageStatusResponseDownloadStatusPtrOutput {
	return o.ApplyT(func(v GalleryImageStatusResponse) *GalleryImageStatusResponseDownloadStatus { return v.DownloadStatus }).(GalleryImageStatusResponseDownloadStatusPtrOutput)
}

func (o GalleryImageStatusResponseOutput) ErrorCode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GalleryImageStatusResponse) *string { return v.ErrorCode }).(pulumi.StringPtrOutput)
}

func (o GalleryImageStatusResponseOutput) ErrorMessage() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GalleryImageStatusResponse) *string { return v.ErrorMessage }).(pulumi.StringPtrOutput)
}

func (o GalleryImageStatusResponseOutput) ProgressPercentage() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v GalleryImageStatusResponse) *float64 { return v.ProgressPercentage }).(pulumi.Float64PtrOutput)
}

func (o GalleryImageStatusResponseOutput) ProvisioningStatus() GalleryImageStatusResponseProvisioningStatusPtrOutput {
	return o.ApplyT(func(v GalleryImageStatusResponse) *GalleryImageStatusResponseProvisioningStatus {
		return v.ProvisioningStatus
	}).(GalleryImageStatusResponseProvisioningStatusPtrOutput)
}

type GalleryImageStatusResponseDownloadStatus struct {
	DownloadSizeInMB *float64 `pulumi:"downloadSizeInMB"`
}

type GalleryImageStatusResponseDownloadStatusOutput struct{ *pulumi.OutputState }

func (GalleryImageStatusResponseDownloadStatusOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GalleryImageStatusResponseDownloadStatus)(nil)).Elem()
}

func (o GalleryImageStatusResponseDownloadStatusOutput) ToGalleryImageStatusResponseDownloadStatusOutput() GalleryImageStatusResponseDownloadStatusOutput {
	return o
}

func (o GalleryImageStatusResponseDownloadStatusOutput) ToGalleryImageStatusResponseDownloadStatusOutputWithContext(ctx context.Context) GalleryImageStatusResponseDownloadStatusOutput {
	return o
}

func (o GalleryImageStatusResponseDownloadStatusOutput) DownloadSizeInMB() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v GalleryImageStatusResponseDownloadStatus) *float64 { return v.DownloadSizeInMB }).(pulumi.Float64PtrOutput)
}

type GalleryImageStatusResponseDownloadStatusPtrOutput struct{ *pulumi.OutputState }

func (GalleryImageStatusResponseDownloadStatusPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GalleryImageStatusResponseDownloadStatus)(nil)).Elem()
}

func (o GalleryImageStatusResponseDownloadStatusPtrOutput) ToGalleryImageStatusResponseDownloadStatusPtrOutput() GalleryImageStatusResponseDownloadStatusPtrOutput {
	return o
}

func (o GalleryImageStatusResponseDownloadStatusPtrOutput) ToGalleryImageStatusResponseDownloadStatusPtrOutputWithContext(ctx context.Context) GalleryImageStatusResponseDownloadStatusPtrOutput {
	return o
}

func (o GalleryImageStatusResponseDownloadStatusPtrOutput) Elem() GalleryImageStatusResponseDownloadStatusOutput {
	return o.ApplyT(func(v *GalleryImageStatusResponseDownloadStatus) GalleryImageStatusResponseDownloadStatus {
		if v != nil {
			return *v
		}
		var ret GalleryImageStatusResponseDownloadStatus
		return ret
	}).(GalleryImageStatusResponseDownloadStatusOutput)
}

func (o GalleryImageStatusResponseDownloadStatusPtrOutput) DownloadSizeInMB() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *GalleryImageStatusResponseDownloadStatus) *float64 {
		if v == nil {
			return nil
		}
		return v.DownloadSizeInMB
	}).(pulumi.Float64PtrOutput)
}

type GalleryImageStatusResponseProvisioningStatus struct {
	OperationId *string `pulumi:"operationId"`
	Status      *string `pulumi:"status"`
}

type GalleryImageStatusResponseProvisioningStatusOutput struct{ *pulumi.OutputState }

func (GalleryImageStatusResponseProvisioningStatusOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GalleryImageStatusResponseProvisioningStatus)(nil)).Elem()
}

func (o GalleryImageStatusResponseProvisioningStatusOutput) ToGalleryImageStatusResponseProvisioningStatusOutput() GalleryImageStatusResponseProvisioningStatusOutput {
	return o
}

func (o GalleryImageStatusResponseProvisioningStatusOutput) ToGalleryImageStatusResponseProvisioningStatusOutputWithContext(ctx context.Context) GalleryImageStatusResponseProvisioningStatusOutput {
	return o
}

func (o GalleryImageStatusResponseProvisioningStatusOutput) OperationId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GalleryImageStatusResponseProvisioningStatus) *string { return v.OperationId }).(pulumi.StringPtrOutput)
}

func (o GalleryImageStatusResponseProvisioningStatusOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GalleryImageStatusResponseProvisioningStatus) *string { return v.Status }).(pulumi.StringPtrOutput)
}

type GalleryImageStatusResponseProvisioningStatusPtrOutput struct{ *pulumi.OutputState }

func (GalleryImageStatusResponseProvisioningStatusPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GalleryImageStatusResponseProvisioningStatus)(nil)).Elem()
}

func (o GalleryImageStatusResponseProvisioningStatusPtrOutput) ToGalleryImageStatusResponseProvisioningStatusPtrOutput() GalleryImageStatusResponseProvisioningStatusPtrOutput {
	return o
}

func (o GalleryImageStatusResponseProvisioningStatusPtrOutput) ToGalleryImageStatusResponseProvisioningStatusPtrOutputWithContext(ctx context.Context) GalleryImageStatusResponseProvisioningStatusPtrOutput {
	return o
}

func (o GalleryImageStatusResponseProvisioningStatusPtrOutput) Elem() GalleryImageStatusResponseProvisioningStatusOutput {
	return o.ApplyT(func(v *GalleryImageStatusResponseProvisioningStatus) GalleryImageStatusResponseProvisioningStatus {
		if v != nil {
			return *v
		}
		var ret GalleryImageStatusResponseProvisioningStatus
		return ret
	}).(GalleryImageStatusResponseProvisioningStatusOutput)
}

func (o GalleryImageStatusResponseProvisioningStatusPtrOutput) OperationId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GalleryImageStatusResponseProvisioningStatus) *string {
		if v == nil {
			return nil
		}
		return v.OperationId
	}).(pulumi.StringPtrOutput)
}

func (o GalleryImageStatusResponseProvisioningStatusPtrOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GalleryImageStatusResponseProvisioningStatus) *string {
		if v == nil {
			return nil
		}
		return v.Status
	}).(pulumi.StringPtrOutput)
}

type GalleryImageVersion struct {
	Name *string `pulumi:"name"`
}





type GalleryImageVersionInput interface {
	pulumi.Input

	ToGalleryImageVersionOutput() GalleryImageVersionOutput
	ToGalleryImageVersionOutputWithContext(context.Context) GalleryImageVersionOutput
}

type GalleryImageVersionArgs struct {
	Name pulumi.StringPtrInput `pulumi:"name"`
}

func (GalleryImageVersionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GalleryImageVersion)(nil)).Elem()
}

func (i GalleryImageVersionArgs) ToGalleryImageVersionOutput() GalleryImageVersionOutput {
	return i.ToGalleryImageVersionOutputWithContext(context.Background())
}

func (i GalleryImageVersionArgs) ToGalleryImageVersionOutputWithContext(ctx context.Context) GalleryImageVersionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GalleryImageVersionOutput)
}

func (i GalleryImageVersionArgs) ToGalleryImageVersionPtrOutput() GalleryImageVersionPtrOutput {
	return i.ToGalleryImageVersionPtrOutputWithContext(context.Background())
}

func (i GalleryImageVersionArgs) ToGalleryImageVersionPtrOutputWithContext(ctx context.Context) GalleryImageVersionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GalleryImageVersionOutput).ToGalleryImageVersionPtrOutputWithContext(ctx)
}









type GalleryImageVersionPtrInput interface {
	pulumi.Input

	ToGalleryImageVersionPtrOutput() GalleryImageVersionPtrOutput
	ToGalleryImageVersionPtrOutputWithContext(context.Context) GalleryImageVersionPtrOutput
}

type galleryImageVersionPtrType GalleryImageVersionArgs

func GalleryImageVersionPtr(v *GalleryImageVersionArgs) GalleryImageVersionPtrInput {
	return (*galleryImageVersionPtrType)(v)
}

func (*galleryImageVersionPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**GalleryImageVersion)(nil)).Elem()
}

func (i *galleryImageVersionPtrType) ToGalleryImageVersionPtrOutput() GalleryImageVersionPtrOutput {
	return i.ToGalleryImageVersionPtrOutputWithContext(context.Background())
}

func (i *galleryImageVersionPtrType) ToGalleryImageVersionPtrOutputWithContext(ctx context.Context) GalleryImageVersionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GalleryImageVersionPtrOutput)
}

type GalleryImageVersionOutput struct{ *pulumi.OutputState }

func (GalleryImageVersionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GalleryImageVersion)(nil)).Elem()
}

func (o GalleryImageVersionOutput) ToGalleryImageVersionOutput() GalleryImageVersionOutput {
	return o
}

func (o GalleryImageVersionOutput) ToGalleryImageVersionOutputWithContext(ctx context.Context) GalleryImageVersionOutput {
	return o
}

func (o GalleryImageVersionOutput) ToGalleryImageVersionPtrOutput() GalleryImageVersionPtrOutput {
	return o.ToGalleryImageVersionPtrOutputWithContext(context.Background())
}

func (o GalleryImageVersionOutput) ToGalleryImageVersionPtrOutputWithContext(ctx context.Context) GalleryImageVersionPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v GalleryImageVersion) *GalleryImageVersion {
		return &v
	}).(GalleryImageVersionPtrOutput)
}

func (o GalleryImageVersionOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GalleryImageVersion) *string { return v.Name }).(pulumi.StringPtrOutput)
}

type GalleryImageVersionPtrOutput struct{ *pulumi.OutputState }

func (GalleryImageVersionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GalleryImageVersion)(nil)).Elem()
}

func (o GalleryImageVersionPtrOutput) ToGalleryImageVersionPtrOutput() GalleryImageVersionPtrOutput {
	return o
}

func (o GalleryImageVersionPtrOutput) ToGalleryImageVersionPtrOutputWithContext(ctx context.Context) GalleryImageVersionPtrOutput {
	return o
}

func (o GalleryImageVersionPtrOutput) Elem() GalleryImageVersionOutput {
	return o.ApplyT(func(v *GalleryImageVersion) GalleryImageVersion {
		if v != nil {
			return *v
		}
		var ret GalleryImageVersion
		return ret
	}).(GalleryImageVersionOutput)
}

func (o GalleryImageVersionPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GalleryImageVersion) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

type GalleryImageVersionResponse struct {
	Name           *string                                   `pulumi:"name"`
	StorageProfile GalleryImageVersionStorageProfileResponse `pulumi:"storageProfile"`
}

type GalleryImageVersionResponseOutput struct{ *pulumi.OutputState }

func (GalleryImageVersionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GalleryImageVersionResponse)(nil)).Elem()
}

func (o GalleryImageVersionResponseOutput) ToGalleryImageVersionResponseOutput() GalleryImageVersionResponseOutput {
	return o
}

func (o GalleryImageVersionResponseOutput) ToGalleryImageVersionResponseOutputWithContext(ctx context.Context) GalleryImageVersionResponseOutput {
	return o
}

func (o GalleryImageVersionResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GalleryImageVersionResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o GalleryImageVersionResponseOutput) StorageProfile() GalleryImageVersionStorageProfileResponseOutput {
	return o.ApplyT(func(v GalleryImageVersionResponse) GalleryImageVersionStorageProfileResponse { return v.StorageProfile }).(GalleryImageVersionStorageProfileResponseOutput)
}

type GalleryImageVersionResponsePtrOutput struct{ *pulumi.OutputState }

func (GalleryImageVersionResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GalleryImageVersionResponse)(nil)).Elem()
}

func (o GalleryImageVersionResponsePtrOutput) ToGalleryImageVersionResponsePtrOutput() GalleryImageVersionResponsePtrOutput {
	return o
}

func (o GalleryImageVersionResponsePtrOutput) ToGalleryImageVersionResponsePtrOutputWithContext(ctx context.Context) GalleryImageVersionResponsePtrOutput {
	return o
}

func (o GalleryImageVersionResponsePtrOutput) Elem() GalleryImageVersionResponseOutput {
	return o.ApplyT(func(v *GalleryImageVersionResponse) GalleryImageVersionResponse {
		if v != nil {
			return *v
		}
		var ret GalleryImageVersionResponse
		return ret
	}).(GalleryImageVersionResponseOutput)
}

func (o GalleryImageVersionResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GalleryImageVersionResponse) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o GalleryImageVersionResponsePtrOutput) StorageProfile() GalleryImageVersionStorageProfileResponsePtrOutput {
	return o.ApplyT(func(v *GalleryImageVersionResponse) *GalleryImageVersionStorageProfileResponse {
		if v == nil {
			return nil
		}
		return &v.StorageProfile
	}).(GalleryImageVersionStorageProfileResponsePtrOutput)
}

type GalleryImageVersionStorageProfileResponse struct {
	OsDiskImage *GalleryOSDiskImageResponse `pulumi:"osDiskImage"`
}

type GalleryImageVersionStorageProfileResponseOutput struct{ *pulumi.OutputState }

func (GalleryImageVersionStorageProfileResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GalleryImageVersionStorageProfileResponse)(nil)).Elem()
}

func (o GalleryImageVersionStorageProfileResponseOutput) ToGalleryImageVersionStorageProfileResponseOutput() GalleryImageVersionStorageProfileResponseOutput {
	return o
}

func (o GalleryImageVersionStorageProfileResponseOutput) ToGalleryImageVersionStorageProfileResponseOutputWithContext(ctx context.Context) GalleryImageVersionStorageProfileResponseOutput {
	return o
}

func (o GalleryImageVersionStorageProfileResponseOutput) OsDiskImage() GalleryOSDiskImageResponsePtrOutput {
	return o.ApplyT(func(v GalleryImageVersionStorageProfileResponse) *GalleryOSDiskImageResponse { return v.OsDiskImage }).(GalleryOSDiskImageResponsePtrOutput)
}

type GalleryImageVersionStorageProfileResponsePtrOutput struct{ *pulumi.OutputState }

func (GalleryImageVersionStorageProfileResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GalleryImageVersionStorageProfileResponse)(nil)).Elem()
}

func (o GalleryImageVersionStorageProfileResponsePtrOutput) ToGalleryImageVersionStorageProfileResponsePtrOutput() GalleryImageVersionStorageProfileResponsePtrOutput {
	return o
}

func (o GalleryImageVersionStorageProfileResponsePtrOutput) ToGalleryImageVersionStorageProfileResponsePtrOutputWithContext(ctx context.Context) GalleryImageVersionStorageProfileResponsePtrOutput {
	return o
}

func (o GalleryImageVersionStorageProfileResponsePtrOutput) Elem() GalleryImageVersionStorageProfileResponseOutput {
	return o.ApplyT(func(v *GalleryImageVersionStorageProfileResponse) GalleryImageVersionStorageProfileResponse {
		if v != nil {
			return *v
		}
		var ret GalleryImageVersionStorageProfileResponse
		return ret
	}).(GalleryImageVersionStorageProfileResponseOutput)
}

func (o GalleryImageVersionStorageProfileResponsePtrOutput) OsDiskImage() GalleryOSDiskImageResponsePtrOutput {
	return o.ApplyT(func(v *GalleryImageVersionStorageProfileResponse) *GalleryOSDiskImageResponse {
		if v == nil {
			return nil
		}
		return v.OsDiskImage
	}).(GalleryOSDiskImageResponsePtrOutput)
}

type GalleryOSDiskImageResponse struct {
	SizeInMB float64 `pulumi:"sizeInMB"`
}

type GalleryOSDiskImageResponseOutput struct{ *pulumi.OutputState }

func (GalleryOSDiskImageResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GalleryOSDiskImageResponse)(nil)).Elem()
}

func (o GalleryOSDiskImageResponseOutput) ToGalleryOSDiskImageResponseOutput() GalleryOSDiskImageResponseOutput {
	return o
}

func (o GalleryOSDiskImageResponseOutput) ToGalleryOSDiskImageResponseOutputWithContext(ctx context.Context) GalleryOSDiskImageResponseOutput {
	return o
}

func (o GalleryOSDiskImageResponseOutput) SizeInMB() pulumi.Float64Output {
	return o.ApplyT(func(v GalleryOSDiskImageResponse) float64 { return v.SizeInMB }).(pulumi.Float64Output)
}

type GalleryOSDiskImageResponsePtrOutput struct{ *pulumi.OutputState }

func (GalleryOSDiskImageResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GalleryOSDiskImageResponse)(nil)).Elem()
}

func (o GalleryOSDiskImageResponsePtrOutput) ToGalleryOSDiskImageResponsePtrOutput() GalleryOSDiskImageResponsePtrOutput {
	return o
}

func (o GalleryOSDiskImageResponsePtrOutput) ToGalleryOSDiskImageResponsePtrOutputWithContext(ctx context.Context) GalleryOSDiskImageResponsePtrOutput {
	return o
}

func (o GalleryOSDiskImageResponsePtrOutput) Elem() GalleryOSDiskImageResponseOutput {
	return o.ApplyT(func(v *GalleryOSDiskImageResponse) GalleryOSDiskImageResponse {
		if v != nil {
			return *v
		}
		var ret GalleryOSDiskImageResponse
		return ret
	}).(GalleryOSDiskImageResponseOutput)
}

func (o GalleryOSDiskImageResponsePtrOutput) SizeInMB() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *GalleryOSDiskImageResponse) *float64 {
		if v == nil {
			return nil
		}
		return &v.SizeInMB
	}).(pulumi.Float64PtrOutput)
}

type GuestAgentProfileResponse struct {
	AgentVersion     string                `pulumi:"agentVersion"`
	ErrorDetails     []ErrorDetailResponse `pulumi:"errorDetails"`
	LastStatusChange string                `pulumi:"lastStatusChange"`
	Status           string                `pulumi:"status"`
	VmUuid           string                `pulumi:"vmUuid"`
}

type GuestAgentProfileResponseOutput struct{ *pulumi.OutputState }

func (GuestAgentProfileResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GuestAgentProfileResponse)(nil)).Elem()
}

func (o GuestAgentProfileResponseOutput) ToGuestAgentProfileResponseOutput() GuestAgentProfileResponseOutput {
	return o
}

func (o GuestAgentProfileResponseOutput) ToGuestAgentProfileResponseOutputWithContext(ctx context.Context) GuestAgentProfileResponseOutput {
	return o
}

func (o GuestAgentProfileResponseOutput) AgentVersion() pulumi.StringOutput {
	return o.ApplyT(func(v GuestAgentProfileResponse) string { return v.AgentVersion }).(pulumi.StringOutput)
}

func (o GuestAgentProfileResponseOutput) ErrorDetails() ErrorDetailResponseArrayOutput {
	return o.ApplyT(func(v GuestAgentProfileResponse) []ErrorDetailResponse { return v.ErrorDetails }).(ErrorDetailResponseArrayOutput)
}

func (o GuestAgentProfileResponseOutput) LastStatusChange() pulumi.StringOutput {
	return o.ApplyT(func(v GuestAgentProfileResponse) string { return v.LastStatusChange }).(pulumi.StringOutput)
}

func (o GuestAgentProfileResponseOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v GuestAgentProfileResponse) string { return v.Status }).(pulumi.StringOutput)
}

func (o GuestAgentProfileResponseOutput) VmUuid() pulumi.StringOutput {
	return o.ApplyT(func(v GuestAgentProfileResponse) string { return v.VmUuid }).(pulumi.StringOutput)
}

type GuestAgentProfileResponsePtrOutput struct{ *pulumi.OutputState }

func (GuestAgentProfileResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GuestAgentProfileResponse)(nil)).Elem()
}

func (o GuestAgentProfileResponsePtrOutput) ToGuestAgentProfileResponsePtrOutput() GuestAgentProfileResponsePtrOutput {
	return o
}

func (o GuestAgentProfileResponsePtrOutput) ToGuestAgentProfileResponsePtrOutputWithContext(ctx context.Context) GuestAgentProfileResponsePtrOutput {
	return o
}

func (o GuestAgentProfileResponsePtrOutput) Elem() GuestAgentProfileResponseOutput {
	return o.ApplyT(func(v *GuestAgentProfileResponse) GuestAgentProfileResponse {
		if v != nil {
			return *v
		}
		var ret GuestAgentProfileResponse
		return ret
	}).(GuestAgentProfileResponseOutput)
}

func (o GuestAgentProfileResponsePtrOutput) AgentVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GuestAgentProfileResponse) *string {
		if v == nil {
			return nil
		}
		return &v.AgentVersion
	}).(pulumi.StringPtrOutput)
}

func (o GuestAgentProfileResponsePtrOutput) ErrorDetails() ErrorDetailResponseArrayOutput {
	return o.ApplyT(func(v *GuestAgentProfileResponse) []ErrorDetailResponse {
		if v == nil {
			return nil
		}
		return v.ErrorDetails
	}).(ErrorDetailResponseArrayOutput)
}

func (o GuestAgentProfileResponsePtrOutput) LastStatusChange() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GuestAgentProfileResponse) *string {
		if v == nil {
			return nil
		}
		return &v.LastStatusChange
	}).(pulumi.StringPtrOutput)
}

func (o GuestAgentProfileResponsePtrOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GuestAgentProfileResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Status
	}).(pulumi.StringPtrOutput)
}

func (o GuestAgentProfileResponsePtrOutput) VmUuid() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GuestAgentProfileResponse) *string {
		if v == nil {
			return nil
		}
		return &v.VmUuid
	}).(pulumi.StringPtrOutput)
}

type GuestCredential struct {
	Password *string `pulumi:"password"`
	Username *string `pulumi:"username"`
}





type GuestCredentialInput interface {
	pulumi.Input

	ToGuestCredentialOutput() GuestCredentialOutput
	ToGuestCredentialOutputWithContext(context.Context) GuestCredentialOutput
}

type GuestCredentialArgs struct {
	Password pulumi.StringPtrInput `pulumi:"password"`
	Username pulumi.StringPtrInput `pulumi:"username"`
}

func (GuestCredentialArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GuestCredential)(nil)).Elem()
}

func (i GuestCredentialArgs) ToGuestCredentialOutput() GuestCredentialOutput {
	return i.ToGuestCredentialOutputWithContext(context.Background())
}

func (i GuestCredentialArgs) ToGuestCredentialOutputWithContext(ctx context.Context) GuestCredentialOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GuestCredentialOutput)
}

func (i GuestCredentialArgs) ToGuestCredentialPtrOutput() GuestCredentialPtrOutput {
	return i.ToGuestCredentialPtrOutputWithContext(context.Background())
}

func (i GuestCredentialArgs) ToGuestCredentialPtrOutputWithContext(ctx context.Context) GuestCredentialPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GuestCredentialOutput).ToGuestCredentialPtrOutputWithContext(ctx)
}









type GuestCredentialPtrInput interface {
	pulumi.Input

	ToGuestCredentialPtrOutput() GuestCredentialPtrOutput
	ToGuestCredentialPtrOutputWithContext(context.Context) GuestCredentialPtrOutput
}

type guestCredentialPtrType GuestCredentialArgs

func GuestCredentialPtr(v *GuestCredentialArgs) GuestCredentialPtrInput {
	return (*guestCredentialPtrType)(v)
}

func (*guestCredentialPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**GuestCredential)(nil)).Elem()
}

func (i *guestCredentialPtrType) ToGuestCredentialPtrOutput() GuestCredentialPtrOutput {
	return i.ToGuestCredentialPtrOutputWithContext(context.Background())
}

func (i *guestCredentialPtrType) ToGuestCredentialPtrOutputWithContext(ctx context.Context) GuestCredentialPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GuestCredentialPtrOutput)
}

type GuestCredentialOutput struct{ *pulumi.OutputState }

func (GuestCredentialOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GuestCredential)(nil)).Elem()
}

func (o GuestCredentialOutput) ToGuestCredentialOutput() GuestCredentialOutput {
	return o
}

func (o GuestCredentialOutput) ToGuestCredentialOutputWithContext(ctx context.Context) GuestCredentialOutput {
	return o
}

func (o GuestCredentialOutput) ToGuestCredentialPtrOutput() GuestCredentialPtrOutput {
	return o.ToGuestCredentialPtrOutputWithContext(context.Background())
}

func (o GuestCredentialOutput) ToGuestCredentialPtrOutputWithContext(ctx context.Context) GuestCredentialPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v GuestCredential) *GuestCredential {
		return &v
	}).(GuestCredentialPtrOutput)
}

func (o GuestCredentialOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GuestCredential) *string { return v.Password }).(pulumi.StringPtrOutput)
}

func (o GuestCredentialOutput) Username() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GuestCredential) *string { return v.Username }).(pulumi.StringPtrOutput)
}

type GuestCredentialPtrOutput struct{ *pulumi.OutputState }

func (GuestCredentialPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GuestCredential)(nil)).Elem()
}

func (o GuestCredentialPtrOutput) ToGuestCredentialPtrOutput() GuestCredentialPtrOutput {
	return o
}

func (o GuestCredentialPtrOutput) ToGuestCredentialPtrOutputWithContext(ctx context.Context) GuestCredentialPtrOutput {
	return o
}

func (o GuestCredentialPtrOutput) Elem() GuestCredentialOutput {
	return o.ApplyT(func(v *GuestCredential) GuestCredential {
		if v != nil {
			return *v
		}
		var ret GuestCredential
		return ret
	}).(GuestCredentialOutput)
}

func (o GuestCredentialPtrOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GuestCredential) *string {
		if v == nil {
			return nil
		}
		return v.Password
	}).(pulumi.StringPtrOutput)
}

func (o GuestCredentialPtrOutput) Username() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GuestCredential) *string {
		if v == nil {
			return nil
		}
		return v.Username
	}).(pulumi.StringPtrOutput)
}

type GuestCredentialResponse struct {
	Username *string `pulumi:"username"`
}

type GuestCredentialResponseOutput struct{ *pulumi.OutputState }

func (GuestCredentialResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GuestCredentialResponse)(nil)).Elem()
}

func (o GuestCredentialResponseOutput) ToGuestCredentialResponseOutput() GuestCredentialResponseOutput {
	return o
}

func (o GuestCredentialResponseOutput) ToGuestCredentialResponseOutputWithContext(ctx context.Context) GuestCredentialResponseOutput {
	return o
}

func (o GuestCredentialResponseOutput) Username() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GuestCredentialResponse) *string { return v.Username }).(pulumi.StringPtrOutput)
}

type GuestCredentialResponsePtrOutput struct{ *pulumi.OutputState }

func (GuestCredentialResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GuestCredentialResponse)(nil)).Elem()
}

func (o GuestCredentialResponsePtrOutput) ToGuestCredentialResponsePtrOutput() GuestCredentialResponsePtrOutput {
	return o
}

func (o GuestCredentialResponsePtrOutput) ToGuestCredentialResponsePtrOutputWithContext(ctx context.Context) GuestCredentialResponsePtrOutput {
	return o
}

func (o GuestCredentialResponsePtrOutput) Elem() GuestCredentialResponseOutput {
	return o.ApplyT(func(v *GuestCredentialResponse) GuestCredentialResponse {
		if v != nil {
			return *v
		}
		var ret GuestCredentialResponse
		return ret
	}).(GuestCredentialResponseOutput)
}

func (o GuestCredentialResponsePtrOutput) Username() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GuestCredentialResponse) *string {
		if v == nil {
			return nil
		}
		return v.Username
	}).(pulumi.StringPtrOutput)
}

type HttpProxyConfiguration struct {
	HttpsProxy *string `pulumi:"httpsProxy"`
}





type HttpProxyConfigurationInput interface {
	pulumi.Input

	ToHttpProxyConfigurationOutput() HttpProxyConfigurationOutput
	ToHttpProxyConfigurationOutputWithContext(context.Context) HttpProxyConfigurationOutput
}

type HttpProxyConfigurationArgs struct {
	HttpsProxy pulumi.StringPtrInput `pulumi:"httpsProxy"`
}

func (HttpProxyConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*HttpProxyConfiguration)(nil)).Elem()
}

func (i HttpProxyConfigurationArgs) ToHttpProxyConfigurationOutput() HttpProxyConfigurationOutput {
	return i.ToHttpProxyConfigurationOutputWithContext(context.Background())
}

func (i HttpProxyConfigurationArgs) ToHttpProxyConfigurationOutputWithContext(ctx context.Context) HttpProxyConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HttpProxyConfigurationOutput)
}

func (i HttpProxyConfigurationArgs) ToHttpProxyConfigurationPtrOutput() HttpProxyConfigurationPtrOutput {
	return i.ToHttpProxyConfigurationPtrOutputWithContext(context.Background())
}

func (i HttpProxyConfigurationArgs) ToHttpProxyConfigurationPtrOutputWithContext(ctx context.Context) HttpProxyConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HttpProxyConfigurationOutput).ToHttpProxyConfigurationPtrOutputWithContext(ctx)
}









type HttpProxyConfigurationPtrInput interface {
	pulumi.Input

	ToHttpProxyConfigurationPtrOutput() HttpProxyConfigurationPtrOutput
	ToHttpProxyConfigurationPtrOutputWithContext(context.Context) HttpProxyConfigurationPtrOutput
}

type httpProxyConfigurationPtrType HttpProxyConfigurationArgs

func HttpProxyConfigurationPtr(v *HttpProxyConfigurationArgs) HttpProxyConfigurationPtrInput {
	return (*httpProxyConfigurationPtrType)(v)
}

func (*httpProxyConfigurationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**HttpProxyConfiguration)(nil)).Elem()
}

func (i *httpProxyConfigurationPtrType) ToHttpProxyConfigurationPtrOutput() HttpProxyConfigurationPtrOutput {
	return i.ToHttpProxyConfigurationPtrOutputWithContext(context.Background())
}

func (i *httpProxyConfigurationPtrType) ToHttpProxyConfigurationPtrOutputWithContext(ctx context.Context) HttpProxyConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HttpProxyConfigurationPtrOutput)
}

type HttpProxyConfigurationOutput struct{ *pulumi.OutputState }

func (HttpProxyConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*HttpProxyConfiguration)(nil)).Elem()
}

func (o HttpProxyConfigurationOutput) ToHttpProxyConfigurationOutput() HttpProxyConfigurationOutput {
	return o
}

func (o HttpProxyConfigurationOutput) ToHttpProxyConfigurationOutputWithContext(ctx context.Context) HttpProxyConfigurationOutput {
	return o
}

func (o HttpProxyConfigurationOutput) ToHttpProxyConfigurationPtrOutput() HttpProxyConfigurationPtrOutput {
	return o.ToHttpProxyConfigurationPtrOutputWithContext(context.Background())
}

func (o HttpProxyConfigurationOutput) ToHttpProxyConfigurationPtrOutputWithContext(ctx context.Context) HttpProxyConfigurationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v HttpProxyConfiguration) *HttpProxyConfiguration {
		return &v
	}).(HttpProxyConfigurationPtrOutput)
}

func (o HttpProxyConfigurationOutput) HttpsProxy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HttpProxyConfiguration) *string { return v.HttpsProxy }).(pulumi.StringPtrOutput)
}

type HttpProxyConfigurationPtrOutput struct{ *pulumi.OutputState }

func (HttpProxyConfigurationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**HttpProxyConfiguration)(nil)).Elem()
}

func (o HttpProxyConfigurationPtrOutput) ToHttpProxyConfigurationPtrOutput() HttpProxyConfigurationPtrOutput {
	return o
}

func (o HttpProxyConfigurationPtrOutput) ToHttpProxyConfigurationPtrOutputWithContext(ctx context.Context) HttpProxyConfigurationPtrOutput {
	return o
}

func (o HttpProxyConfigurationPtrOutput) Elem() HttpProxyConfigurationOutput {
	return o.ApplyT(func(v *HttpProxyConfiguration) HttpProxyConfiguration {
		if v != nil {
			return *v
		}
		var ret HttpProxyConfiguration
		return ret
	}).(HttpProxyConfigurationOutput)
}

func (o HttpProxyConfigurationPtrOutput) HttpsProxy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HttpProxyConfiguration) *string {
		if v == nil {
			return nil
		}
		return v.HttpsProxy
	}).(pulumi.StringPtrOutput)
}

type HttpProxyConfigurationResponse struct {
	HttpsProxy *string `pulumi:"httpsProxy"`
}

type HttpProxyConfigurationResponseOutput struct{ *pulumi.OutputState }

func (HttpProxyConfigurationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*HttpProxyConfigurationResponse)(nil)).Elem()
}

func (o HttpProxyConfigurationResponseOutput) ToHttpProxyConfigurationResponseOutput() HttpProxyConfigurationResponseOutput {
	return o
}

func (o HttpProxyConfigurationResponseOutput) ToHttpProxyConfigurationResponseOutputWithContext(ctx context.Context) HttpProxyConfigurationResponseOutput {
	return o
}

func (o HttpProxyConfigurationResponseOutput) HttpsProxy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HttpProxyConfigurationResponse) *string { return v.HttpsProxy }).(pulumi.StringPtrOutput)
}

type HttpProxyConfigurationResponsePtrOutput struct{ *pulumi.OutputState }

func (HttpProxyConfigurationResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**HttpProxyConfigurationResponse)(nil)).Elem()
}

func (o HttpProxyConfigurationResponsePtrOutput) ToHttpProxyConfigurationResponsePtrOutput() HttpProxyConfigurationResponsePtrOutput {
	return o
}

func (o HttpProxyConfigurationResponsePtrOutput) ToHttpProxyConfigurationResponsePtrOutputWithContext(ctx context.Context) HttpProxyConfigurationResponsePtrOutput {
	return o
}

func (o HttpProxyConfigurationResponsePtrOutput) Elem() HttpProxyConfigurationResponseOutput {
	return o.ApplyT(func(v *HttpProxyConfigurationResponse) HttpProxyConfigurationResponse {
		if v != nil {
			return *v
		}
		var ret HttpProxyConfigurationResponse
		return ret
	}).(HttpProxyConfigurationResponseOutput)
}

func (o HttpProxyConfigurationResponsePtrOutput) HttpsProxy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HttpProxyConfigurationResponse) *string {
		if v == nil {
			return nil
		}
		return v.HttpsProxy
	}).(pulumi.StringPtrOutput)
}

type IPPool struct {
	End        *string         `pulumi:"end"`
	IpPoolType *IPPoolTypeEnum `pulumi:"ipPoolType"`
	Start      *string         `pulumi:"start"`
}





type IPPoolInput interface {
	pulumi.Input

	ToIPPoolOutput() IPPoolOutput
	ToIPPoolOutputWithContext(context.Context) IPPoolOutput
}

type IPPoolArgs struct {
	End        pulumi.StringPtrInput  `pulumi:"end"`
	IpPoolType IPPoolTypeEnumPtrInput `pulumi:"ipPoolType"`
	Start      pulumi.StringPtrInput  `pulumi:"start"`
}

func (IPPoolArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*IPPool)(nil)).Elem()
}

func (i IPPoolArgs) ToIPPoolOutput() IPPoolOutput {
	return i.ToIPPoolOutputWithContext(context.Background())
}

func (i IPPoolArgs) ToIPPoolOutputWithContext(ctx context.Context) IPPoolOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IPPoolOutput)
}





type IPPoolArrayInput interface {
	pulumi.Input

	ToIPPoolArrayOutput() IPPoolArrayOutput
	ToIPPoolArrayOutputWithContext(context.Context) IPPoolArrayOutput
}

type IPPoolArray []IPPoolInput

func (IPPoolArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]IPPool)(nil)).Elem()
}

func (i IPPoolArray) ToIPPoolArrayOutput() IPPoolArrayOutput {
	return i.ToIPPoolArrayOutputWithContext(context.Background())
}

func (i IPPoolArray) ToIPPoolArrayOutputWithContext(ctx context.Context) IPPoolArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IPPoolArrayOutput)
}

type IPPoolOutput struct{ *pulumi.OutputState }

func (IPPoolOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IPPool)(nil)).Elem()
}

func (o IPPoolOutput) ToIPPoolOutput() IPPoolOutput {
	return o
}

func (o IPPoolOutput) ToIPPoolOutputWithContext(ctx context.Context) IPPoolOutput {
	return o
}

func (o IPPoolOutput) End() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IPPool) *string { return v.End }).(pulumi.StringPtrOutput)
}

func (o IPPoolOutput) IpPoolType() IPPoolTypeEnumPtrOutput {
	return o.ApplyT(func(v IPPool) *IPPoolTypeEnum { return v.IpPoolType }).(IPPoolTypeEnumPtrOutput)
}

func (o IPPoolOutput) Start() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IPPool) *string { return v.Start }).(pulumi.StringPtrOutput)
}

type IPPoolArrayOutput struct{ *pulumi.OutputState }

func (IPPoolArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]IPPool)(nil)).Elem()
}

func (o IPPoolArrayOutput) ToIPPoolArrayOutput() IPPoolArrayOutput {
	return o
}

func (o IPPoolArrayOutput) ToIPPoolArrayOutputWithContext(ctx context.Context) IPPoolArrayOutput {
	return o
}

func (o IPPoolArrayOutput) Index(i pulumi.IntInput) IPPoolOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) IPPool {
		return vs[0].([]IPPool)[vs[1].(int)]
	}).(IPPoolOutput)
}

type IPPoolInfoResponse struct {
	Available string `pulumi:"available"`
	Used      string `pulumi:"used"`
}

type IPPoolInfoResponseOutput struct{ *pulumi.OutputState }

func (IPPoolInfoResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IPPoolInfoResponse)(nil)).Elem()
}

func (o IPPoolInfoResponseOutput) ToIPPoolInfoResponseOutput() IPPoolInfoResponseOutput {
	return o
}

func (o IPPoolInfoResponseOutput) ToIPPoolInfoResponseOutputWithContext(ctx context.Context) IPPoolInfoResponseOutput {
	return o
}

func (o IPPoolInfoResponseOutput) Available() pulumi.StringOutput {
	return o.ApplyT(func(v IPPoolInfoResponse) string { return v.Available }).(pulumi.StringOutput)
}

func (o IPPoolInfoResponseOutput) Used() pulumi.StringOutput {
	return o.ApplyT(func(v IPPoolInfoResponse) string { return v.Used }).(pulumi.StringOutput)
}

type IPPoolInfoResponsePtrOutput struct{ *pulumi.OutputState }

func (IPPoolInfoResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IPPoolInfoResponse)(nil)).Elem()
}

func (o IPPoolInfoResponsePtrOutput) ToIPPoolInfoResponsePtrOutput() IPPoolInfoResponsePtrOutput {
	return o
}

func (o IPPoolInfoResponsePtrOutput) ToIPPoolInfoResponsePtrOutputWithContext(ctx context.Context) IPPoolInfoResponsePtrOutput {
	return o
}

func (o IPPoolInfoResponsePtrOutput) Elem() IPPoolInfoResponseOutput {
	return o.ApplyT(func(v *IPPoolInfoResponse) IPPoolInfoResponse {
		if v != nil {
			return *v
		}
		var ret IPPoolInfoResponse
		return ret
	}).(IPPoolInfoResponseOutput)
}

func (o IPPoolInfoResponsePtrOutput) Available() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IPPoolInfoResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Available
	}).(pulumi.StringPtrOutput)
}

func (o IPPoolInfoResponsePtrOutput) Used() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IPPoolInfoResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Used
	}).(pulumi.StringPtrOutput)
}

type IPPoolResponse struct {
	End        *string             `pulumi:"end"`
	Info       *IPPoolInfoResponse `pulumi:"info"`
	IpPoolType *string             `pulumi:"ipPoolType"`
	Start      *string             `pulumi:"start"`
}

type IPPoolResponseOutput struct{ *pulumi.OutputState }

func (IPPoolResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IPPoolResponse)(nil)).Elem()
}

func (o IPPoolResponseOutput) ToIPPoolResponseOutput() IPPoolResponseOutput {
	return o
}

func (o IPPoolResponseOutput) ToIPPoolResponseOutputWithContext(ctx context.Context) IPPoolResponseOutput {
	return o
}

func (o IPPoolResponseOutput) End() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IPPoolResponse) *string { return v.End }).(pulumi.StringPtrOutput)
}

func (o IPPoolResponseOutput) Info() IPPoolInfoResponsePtrOutput {
	return o.ApplyT(func(v IPPoolResponse) *IPPoolInfoResponse { return v.Info }).(IPPoolInfoResponsePtrOutput)
}

func (o IPPoolResponseOutput) IpPoolType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IPPoolResponse) *string { return v.IpPoolType }).(pulumi.StringPtrOutput)
}

func (o IPPoolResponseOutput) Start() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IPPoolResponse) *string { return v.Start }).(pulumi.StringPtrOutput)
}

type IPPoolResponseArrayOutput struct{ *pulumi.OutputState }

func (IPPoolResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]IPPoolResponse)(nil)).Elem()
}

func (o IPPoolResponseArrayOutput) ToIPPoolResponseArrayOutput() IPPoolResponseArrayOutput {
	return o
}

func (o IPPoolResponseArrayOutput) ToIPPoolResponseArrayOutputWithContext(ctx context.Context) IPPoolResponseArrayOutput {
	return o
}

func (o IPPoolResponseArrayOutput) Index(i pulumi.IntInput) IPPoolResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) IPPoolResponse {
		return vs[0].([]IPPoolResponse)[vs[1].(int)]
	}).(IPPoolResponseOutput)
}

type Identity struct {
	Type *ResourceIdentityType `pulumi:"type"`
}





type IdentityInput interface {
	pulumi.Input

	ToIdentityOutput() IdentityOutput
	ToIdentityOutputWithContext(context.Context) IdentityOutput
}

type IdentityArgs struct {
	Type ResourceIdentityTypePtrInput `pulumi:"type"`
}

func (IdentityArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Identity)(nil)).Elem()
}

func (i IdentityArgs) ToIdentityOutput() IdentityOutput {
	return i.ToIdentityOutputWithContext(context.Background())
}

func (i IdentityArgs) ToIdentityOutputWithContext(ctx context.Context) IdentityOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentityOutput)
}

func (i IdentityArgs) ToIdentityPtrOutput() IdentityPtrOutput {
	return i.ToIdentityPtrOutputWithContext(context.Background())
}

func (i IdentityArgs) ToIdentityPtrOutputWithContext(ctx context.Context) IdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentityOutput).ToIdentityPtrOutputWithContext(ctx)
}









type IdentityPtrInput interface {
	pulumi.Input

	ToIdentityPtrOutput() IdentityPtrOutput
	ToIdentityPtrOutputWithContext(context.Context) IdentityPtrOutput
}

type identityPtrType IdentityArgs

func IdentityPtr(v *IdentityArgs) IdentityPtrInput {
	return (*identityPtrType)(v)
}

func (*identityPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Identity)(nil)).Elem()
}

func (i *identityPtrType) ToIdentityPtrOutput() IdentityPtrOutput {
	return i.ToIdentityPtrOutputWithContext(context.Background())
}

func (i *identityPtrType) ToIdentityPtrOutputWithContext(ctx context.Context) IdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentityPtrOutput)
}

type IdentityOutput struct{ *pulumi.OutputState }

func (IdentityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Identity)(nil)).Elem()
}

func (o IdentityOutput) ToIdentityOutput() IdentityOutput {
	return o
}

func (o IdentityOutput) ToIdentityOutputWithContext(ctx context.Context) IdentityOutput {
	return o
}

func (o IdentityOutput) ToIdentityPtrOutput() IdentityPtrOutput {
	return o.ToIdentityPtrOutputWithContext(context.Background())
}

func (o IdentityOutput) ToIdentityPtrOutputWithContext(ctx context.Context) IdentityPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Identity) *Identity {
		return &v
	}).(IdentityPtrOutput)
}

func (o IdentityOutput) Type() ResourceIdentityTypePtrOutput {
	return o.ApplyT(func(v Identity) *ResourceIdentityType { return v.Type }).(ResourceIdentityTypePtrOutput)
}

type IdentityPtrOutput struct{ *pulumi.OutputState }

func (IdentityPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Identity)(nil)).Elem()
}

func (o IdentityPtrOutput) ToIdentityPtrOutput() IdentityPtrOutput {
	return o
}

func (o IdentityPtrOutput) ToIdentityPtrOutputWithContext(ctx context.Context) IdentityPtrOutput {
	return o
}

func (o IdentityPtrOutput) Elem() IdentityOutput {
	return o.ApplyT(func(v *Identity) Identity {
		if v != nil {
			return *v
		}
		var ret Identity
		return ret
	}).(IdentityOutput)
}

func (o IdentityPtrOutput) Type() ResourceIdentityTypePtrOutput {
	return o.ApplyT(func(v *Identity) *ResourceIdentityType {
		if v == nil {
			return nil
		}
		return v.Type
	}).(ResourceIdentityTypePtrOutput)
}

type IdentityResponse struct {
	PrincipalId string  `pulumi:"principalId"`
	TenantId    string  `pulumi:"tenantId"`
	Type        *string `pulumi:"type"`
}

type IdentityResponseOutput struct{ *pulumi.OutputState }

func (IdentityResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IdentityResponse)(nil)).Elem()
}

func (o IdentityResponseOutput) ToIdentityResponseOutput() IdentityResponseOutput {
	return o
}

func (o IdentityResponseOutput) ToIdentityResponseOutputWithContext(ctx context.Context) IdentityResponseOutput {
	return o
}

func (o IdentityResponseOutput) PrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v IdentityResponse) string { return v.PrincipalId }).(pulumi.StringOutput)
}

func (o IdentityResponseOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v IdentityResponse) string { return v.TenantId }).(pulumi.StringOutput)
}

func (o IdentityResponseOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IdentityResponse) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type IdentityResponsePtrOutput struct{ *pulumi.OutputState }

func (IdentityResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IdentityResponse)(nil)).Elem()
}

func (o IdentityResponsePtrOutput) ToIdentityResponsePtrOutput() IdentityResponsePtrOutput {
	return o
}

func (o IdentityResponsePtrOutput) ToIdentityResponsePtrOutputWithContext(ctx context.Context) IdentityResponsePtrOutput {
	return o
}

func (o IdentityResponsePtrOutput) Elem() IdentityResponseOutput {
	return o.ApplyT(func(v *IdentityResponse) IdentityResponse {
		if v != nil {
			return *v
		}
		var ret IdentityResponse
		return ret
	}).(IdentityResponseOutput)
}

func (o IdentityResponsePtrOutput) PrincipalId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IdentityResponse) *string {
		if v == nil {
			return nil
		}
		return &v.PrincipalId
	}).(pulumi.StringPtrOutput)
}

func (o IdentityResponsePtrOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IdentityResponse) *string {
		if v == nil {
			return nil
		}
		return &v.TenantId
	}).(pulumi.StringPtrOutput)
}

func (o IdentityResponsePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IdentityResponse) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

type InterfaceDNSSettings struct {
	DnsServers []string `pulumi:"dnsServers"`
}





type InterfaceDNSSettingsInput interface {
	pulumi.Input

	ToInterfaceDNSSettingsOutput() InterfaceDNSSettingsOutput
	ToInterfaceDNSSettingsOutputWithContext(context.Context) InterfaceDNSSettingsOutput
}

type InterfaceDNSSettingsArgs struct {
	DnsServers pulumi.StringArrayInput `pulumi:"dnsServers"`
}

func (InterfaceDNSSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*InterfaceDNSSettings)(nil)).Elem()
}

func (i InterfaceDNSSettingsArgs) ToInterfaceDNSSettingsOutput() InterfaceDNSSettingsOutput {
	return i.ToInterfaceDNSSettingsOutputWithContext(context.Background())
}

func (i InterfaceDNSSettingsArgs) ToInterfaceDNSSettingsOutputWithContext(ctx context.Context) InterfaceDNSSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InterfaceDNSSettingsOutput)
}

func (i InterfaceDNSSettingsArgs) ToInterfaceDNSSettingsPtrOutput() InterfaceDNSSettingsPtrOutput {
	return i.ToInterfaceDNSSettingsPtrOutputWithContext(context.Background())
}

func (i InterfaceDNSSettingsArgs) ToInterfaceDNSSettingsPtrOutputWithContext(ctx context.Context) InterfaceDNSSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InterfaceDNSSettingsOutput).ToInterfaceDNSSettingsPtrOutputWithContext(ctx)
}









type InterfaceDNSSettingsPtrInput interface {
	pulumi.Input

	ToInterfaceDNSSettingsPtrOutput() InterfaceDNSSettingsPtrOutput
	ToInterfaceDNSSettingsPtrOutputWithContext(context.Context) InterfaceDNSSettingsPtrOutput
}

type interfaceDNSSettingsPtrType InterfaceDNSSettingsArgs

func InterfaceDNSSettingsPtr(v *InterfaceDNSSettingsArgs) InterfaceDNSSettingsPtrInput {
	return (*interfaceDNSSettingsPtrType)(v)
}

func (*interfaceDNSSettingsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**InterfaceDNSSettings)(nil)).Elem()
}

func (i *interfaceDNSSettingsPtrType) ToInterfaceDNSSettingsPtrOutput() InterfaceDNSSettingsPtrOutput {
	return i.ToInterfaceDNSSettingsPtrOutputWithContext(context.Background())
}

func (i *interfaceDNSSettingsPtrType) ToInterfaceDNSSettingsPtrOutputWithContext(ctx context.Context) InterfaceDNSSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InterfaceDNSSettingsPtrOutput)
}

type InterfaceDNSSettingsOutput struct{ *pulumi.OutputState }

func (InterfaceDNSSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*InterfaceDNSSettings)(nil)).Elem()
}

func (o InterfaceDNSSettingsOutput) ToInterfaceDNSSettingsOutput() InterfaceDNSSettingsOutput {
	return o
}

func (o InterfaceDNSSettingsOutput) ToInterfaceDNSSettingsOutputWithContext(ctx context.Context) InterfaceDNSSettingsOutput {
	return o
}

func (o InterfaceDNSSettingsOutput) ToInterfaceDNSSettingsPtrOutput() InterfaceDNSSettingsPtrOutput {
	return o.ToInterfaceDNSSettingsPtrOutputWithContext(context.Background())
}

func (o InterfaceDNSSettingsOutput) ToInterfaceDNSSettingsPtrOutputWithContext(ctx context.Context) InterfaceDNSSettingsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v InterfaceDNSSettings) *InterfaceDNSSettings {
		return &v
	}).(InterfaceDNSSettingsPtrOutput)
}

func (o InterfaceDNSSettingsOutput) DnsServers() pulumi.StringArrayOutput {
	return o.ApplyT(func(v InterfaceDNSSettings) []string { return v.DnsServers }).(pulumi.StringArrayOutput)
}

type InterfaceDNSSettingsPtrOutput struct{ *pulumi.OutputState }

func (InterfaceDNSSettingsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**InterfaceDNSSettings)(nil)).Elem()
}

func (o InterfaceDNSSettingsPtrOutput) ToInterfaceDNSSettingsPtrOutput() InterfaceDNSSettingsPtrOutput {
	return o
}

func (o InterfaceDNSSettingsPtrOutput) ToInterfaceDNSSettingsPtrOutputWithContext(ctx context.Context) InterfaceDNSSettingsPtrOutput {
	return o
}

func (o InterfaceDNSSettingsPtrOutput) Elem() InterfaceDNSSettingsOutput {
	return o.ApplyT(func(v *InterfaceDNSSettings) InterfaceDNSSettings {
		if v != nil {
			return *v
		}
		var ret InterfaceDNSSettings
		return ret
	}).(InterfaceDNSSettingsOutput)
}

func (o InterfaceDNSSettingsPtrOutput) DnsServers() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *InterfaceDNSSettings) []string {
		if v == nil {
			return nil
		}
		return v.DnsServers
	}).(pulumi.StringArrayOutput)
}

type InterfaceDNSSettingsResponse struct {
	DnsServers []string `pulumi:"dnsServers"`
}

type InterfaceDNSSettingsResponseOutput struct{ *pulumi.OutputState }

func (InterfaceDNSSettingsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*InterfaceDNSSettingsResponse)(nil)).Elem()
}

func (o InterfaceDNSSettingsResponseOutput) ToInterfaceDNSSettingsResponseOutput() InterfaceDNSSettingsResponseOutput {
	return o
}

func (o InterfaceDNSSettingsResponseOutput) ToInterfaceDNSSettingsResponseOutputWithContext(ctx context.Context) InterfaceDNSSettingsResponseOutput {
	return o
}

func (o InterfaceDNSSettingsResponseOutput) DnsServers() pulumi.StringArrayOutput {
	return o.ApplyT(func(v InterfaceDNSSettingsResponse) []string { return v.DnsServers }).(pulumi.StringArrayOutput)
}

type InterfaceDNSSettingsResponsePtrOutput struct{ *pulumi.OutputState }

func (InterfaceDNSSettingsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**InterfaceDNSSettingsResponse)(nil)).Elem()
}

func (o InterfaceDNSSettingsResponsePtrOutput) ToInterfaceDNSSettingsResponsePtrOutput() InterfaceDNSSettingsResponsePtrOutput {
	return o
}

func (o InterfaceDNSSettingsResponsePtrOutput) ToInterfaceDNSSettingsResponsePtrOutputWithContext(ctx context.Context) InterfaceDNSSettingsResponsePtrOutput {
	return o
}

func (o InterfaceDNSSettingsResponsePtrOutput) Elem() InterfaceDNSSettingsResponseOutput {
	return o.ApplyT(func(v *InterfaceDNSSettingsResponse) InterfaceDNSSettingsResponse {
		if v != nil {
			return *v
		}
		var ret InterfaceDNSSettingsResponse
		return ret
	}).(InterfaceDNSSettingsResponseOutput)
}

func (o InterfaceDNSSettingsResponsePtrOutput) DnsServers() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *InterfaceDNSSettingsResponse) []string {
		if v == nil {
			return nil
		}
		return v.DnsServers
	}).(pulumi.StringArrayOutput)
}

type IpConfiguration struct {
	Name       *string                    `pulumi:"name"`
	Properties *IpConfigurationProperties `pulumi:"properties"`
}





type IpConfigurationInput interface {
	pulumi.Input

	ToIpConfigurationOutput() IpConfigurationOutput
	ToIpConfigurationOutputWithContext(context.Context) IpConfigurationOutput
}

type IpConfigurationArgs struct {
	Name       pulumi.StringPtrInput             `pulumi:"name"`
	Properties IpConfigurationPropertiesPtrInput `pulumi:"properties"`
}

func (IpConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*IpConfiguration)(nil)).Elem()
}

func (i IpConfigurationArgs) ToIpConfigurationOutput() IpConfigurationOutput {
	return i.ToIpConfigurationOutputWithContext(context.Background())
}

func (i IpConfigurationArgs) ToIpConfigurationOutputWithContext(ctx context.Context) IpConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IpConfigurationOutput)
}





type IpConfigurationArrayInput interface {
	pulumi.Input

	ToIpConfigurationArrayOutput() IpConfigurationArrayOutput
	ToIpConfigurationArrayOutputWithContext(context.Context) IpConfigurationArrayOutput
}

type IpConfigurationArray []IpConfigurationInput

func (IpConfigurationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]IpConfiguration)(nil)).Elem()
}

func (i IpConfigurationArray) ToIpConfigurationArrayOutput() IpConfigurationArrayOutput {
	return i.ToIpConfigurationArrayOutputWithContext(context.Background())
}

func (i IpConfigurationArray) ToIpConfigurationArrayOutputWithContext(ctx context.Context) IpConfigurationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IpConfigurationArrayOutput)
}

type IpConfigurationOutput struct{ *pulumi.OutputState }

func (IpConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IpConfiguration)(nil)).Elem()
}

func (o IpConfigurationOutput) ToIpConfigurationOutput() IpConfigurationOutput {
	return o
}

func (o IpConfigurationOutput) ToIpConfigurationOutputWithContext(ctx context.Context) IpConfigurationOutput {
	return o
}

func (o IpConfigurationOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IpConfiguration) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o IpConfigurationOutput) Properties() IpConfigurationPropertiesPtrOutput {
	return o.ApplyT(func(v IpConfiguration) *IpConfigurationProperties { return v.Properties }).(IpConfigurationPropertiesPtrOutput)
}

type IpConfigurationArrayOutput struct{ *pulumi.OutputState }

func (IpConfigurationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]IpConfiguration)(nil)).Elem()
}

func (o IpConfigurationArrayOutput) ToIpConfigurationArrayOutput() IpConfigurationArrayOutput {
	return o
}

func (o IpConfigurationArrayOutput) ToIpConfigurationArrayOutputWithContext(ctx context.Context) IpConfigurationArrayOutput {
	return o
}

func (o IpConfigurationArrayOutput) Index(i pulumi.IntInput) IpConfigurationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) IpConfiguration {
		return vs[0].([]IpConfiguration)[vs[1].(int)]
	}).(IpConfigurationOutput)
}

type IpConfigurationProperties struct {
	PrefixLength              *string                `pulumi:"prefixLength"`
	PrivateIPAddress          *string                `pulumi:"privateIPAddress"`
	PrivateIPAllocationMethod *string                `pulumi:"privateIPAllocationMethod"`
	Subnet                    *IpConfigurationSubnet `pulumi:"subnet"`
}





type IpConfigurationPropertiesInput interface {
	pulumi.Input

	ToIpConfigurationPropertiesOutput() IpConfigurationPropertiesOutput
	ToIpConfigurationPropertiesOutputWithContext(context.Context) IpConfigurationPropertiesOutput
}

type IpConfigurationPropertiesArgs struct {
	PrefixLength              pulumi.StringPtrInput         `pulumi:"prefixLength"`
	PrivateIPAddress          pulumi.StringPtrInput         `pulumi:"privateIPAddress"`
	PrivateIPAllocationMethod pulumi.StringPtrInput         `pulumi:"privateIPAllocationMethod"`
	Subnet                    IpConfigurationSubnetPtrInput `pulumi:"subnet"`
}

func (IpConfigurationPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*IpConfigurationProperties)(nil)).Elem()
}

func (i IpConfigurationPropertiesArgs) ToIpConfigurationPropertiesOutput() IpConfigurationPropertiesOutput {
	return i.ToIpConfigurationPropertiesOutputWithContext(context.Background())
}

func (i IpConfigurationPropertiesArgs) ToIpConfigurationPropertiesOutputWithContext(ctx context.Context) IpConfigurationPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IpConfigurationPropertiesOutput)
}

func (i IpConfigurationPropertiesArgs) ToIpConfigurationPropertiesPtrOutput() IpConfigurationPropertiesPtrOutput {
	return i.ToIpConfigurationPropertiesPtrOutputWithContext(context.Background())
}

func (i IpConfigurationPropertiesArgs) ToIpConfigurationPropertiesPtrOutputWithContext(ctx context.Context) IpConfigurationPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IpConfigurationPropertiesOutput).ToIpConfigurationPropertiesPtrOutputWithContext(ctx)
}









type IpConfigurationPropertiesPtrInput interface {
	pulumi.Input

	ToIpConfigurationPropertiesPtrOutput() IpConfigurationPropertiesPtrOutput
	ToIpConfigurationPropertiesPtrOutputWithContext(context.Context) IpConfigurationPropertiesPtrOutput
}

type ipConfigurationPropertiesPtrType IpConfigurationPropertiesArgs

func IpConfigurationPropertiesPtr(v *IpConfigurationPropertiesArgs) IpConfigurationPropertiesPtrInput {
	return (*ipConfigurationPropertiesPtrType)(v)
}

func (*ipConfigurationPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**IpConfigurationProperties)(nil)).Elem()
}

func (i *ipConfigurationPropertiesPtrType) ToIpConfigurationPropertiesPtrOutput() IpConfigurationPropertiesPtrOutput {
	return i.ToIpConfigurationPropertiesPtrOutputWithContext(context.Background())
}

func (i *ipConfigurationPropertiesPtrType) ToIpConfigurationPropertiesPtrOutputWithContext(ctx context.Context) IpConfigurationPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IpConfigurationPropertiesPtrOutput)
}

type IpConfigurationPropertiesOutput struct{ *pulumi.OutputState }

func (IpConfigurationPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IpConfigurationProperties)(nil)).Elem()
}

func (o IpConfigurationPropertiesOutput) ToIpConfigurationPropertiesOutput() IpConfigurationPropertiesOutput {
	return o
}

func (o IpConfigurationPropertiesOutput) ToIpConfigurationPropertiesOutputWithContext(ctx context.Context) IpConfigurationPropertiesOutput {
	return o
}

func (o IpConfigurationPropertiesOutput) ToIpConfigurationPropertiesPtrOutput() IpConfigurationPropertiesPtrOutput {
	return o.ToIpConfigurationPropertiesPtrOutputWithContext(context.Background())
}

func (o IpConfigurationPropertiesOutput) ToIpConfigurationPropertiesPtrOutputWithContext(ctx context.Context) IpConfigurationPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v IpConfigurationProperties) *IpConfigurationProperties {
		return &v
	}).(IpConfigurationPropertiesPtrOutput)
}

func (o IpConfigurationPropertiesOutput) PrefixLength() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IpConfigurationProperties) *string { return v.PrefixLength }).(pulumi.StringPtrOutput)
}

func (o IpConfigurationPropertiesOutput) PrivateIPAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IpConfigurationProperties) *string { return v.PrivateIPAddress }).(pulumi.StringPtrOutput)
}

func (o IpConfigurationPropertiesOutput) PrivateIPAllocationMethod() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IpConfigurationProperties) *string { return v.PrivateIPAllocationMethod }).(pulumi.StringPtrOutput)
}

func (o IpConfigurationPropertiesOutput) Subnet() IpConfigurationSubnetPtrOutput {
	return o.ApplyT(func(v IpConfigurationProperties) *IpConfigurationSubnet { return v.Subnet }).(IpConfigurationSubnetPtrOutput)
}

type IpConfigurationPropertiesPtrOutput struct{ *pulumi.OutputState }

func (IpConfigurationPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IpConfigurationProperties)(nil)).Elem()
}

func (o IpConfigurationPropertiesPtrOutput) ToIpConfigurationPropertiesPtrOutput() IpConfigurationPropertiesPtrOutput {
	return o
}

func (o IpConfigurationPropertiesPtrOutput) ToIpConfigurationPropertiesPtrOutputWithContext(ctx context.Context) IpConfigurationPropertiesPtrOutput {
	return o
}

func (o IpConfigurationPropertiesPtrOutput) Elem() IpConfigurationPropertiesOutput {
	return o.ApplyT(func(v *IpConfigurationProperties) IpConfigurationProperties {
		if v != nil {
			return *v
		}
		var ret IpConfigurationProperties
		return ret
	}).(IpConfigurationPropertiesOutput)
}

func (o IpConfigurationPropertiesPtrOutput) PrefixLength() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IpConfigurationProperties) *string {
		if v == nil {
			return nil
		}
		return v.PrefixLength
	}).(pulumi.StringPtrOutput)
}

func (o IpConfigurationPropertiesPtrOutput) PrivateIPAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IpConfigurationProperties) *string {
		if v == nil {
			return nil
		}
		return v.PrivateIPAddress
	}).(pulumi.StringPtrOutput)
}

func (o IpConfigurationPropertiesPtrOutput) PrivateIPAllocationMethod() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IpConfigurationProperties) *string {
		if v == nil {
			return nil
		}
		return v.PrivateIPAllocationMethod
	}).(pulumi.StringPtrOutput)
}

func (o IpConfigurationPropertiesPtrOutput) Subnet() IpConfigurationSubnetPtrOutput {
	return o.ApplyT(func(v *IpConfigurationProperties) *IpConfigurationSubnet {
		if v == nil {
			return nil
		}
		return v.Subnet
	}).(IpConfigurationSubnetPtrOutput)
}

type IpConfigurationResponse struct {
	Name       *string                            `pulumi:"name"`
	Properties *IpConfigurationResponseProperties `pulumi:"properties"`
}

type IpConfigurationResponseOutput struct{ *pulumi.OutputState }

func (IpConfigurationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IpConfigurationResponse)(nil)).Elem()
}

func (o IpConfigurationResponseOutput) ToIpConfigurationResponseOutput() IpConfigurationResponseOutput {
	return o
}

func (o IpConfigurationResponseOutput) ToIpConfigurationResponseOutputWithContext(ctx context.Context) IpConfigurationResponseOutput {
	return o
}

func (o IpConfigurationResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IpConfigurationResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o IpConfigurationResponseOutput) Properties() IpConfigurationResponsePropertiesPtrOutput {
	return o.ApplyT(func(v IpConfigurationResponse) *IpConfigurationResponseProperties { return v.Properties }).(IpConfigurationResponsePropertiesPtrOutput)
}

type IpConfigurationResponseArrayOutput struct{ *pulumi.OutputState }

func (IpConfigurationResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]IpConfigurationResponse)(nil)).Elem()
}

func (o IpConfigurationResponseArrayOutput) ToIpConfigurationResponseArrayOutput() IpConfigurationResponseArrayOutput {
	return o
}

func (o IpConfigurationResponseArrayOutput) ToIpConfigurationResponseArrayOutputWithContext(ctx context.Context) IpConfigurationResponseArrayOutput {
	return o
}

func (o IpConfigurationResponseArrayOutput) Index(i pulumi.IntInput) IpConfigurationResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) IpConfigurationResponse {
		return vs[0].([]IpConfigurationResponse)[vs[1].(int)]
	}).(IpConfigurationResponseOutput)
}

type IpConfigurationResponseProperties struct {
	PrefixLength              *string                        `pulumi:"prefixLength"`
	PrivateIPAddress          *string                        `pulumi:"privateIPAddress"`
	PrivateIPAllocationMethod *string                        `pulumi:"privateIPAllocationMethod"`
	Subnet                    *IpConfigurationResponseSubnet `pulumi:"subnet"`
}

type IpConfigurationResponsePropertiesOutput struct{ *pulumi.OutputState }

func (IpConfigurationResponsePropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IpConfigurationResponseProperties)(nil)).Elem()
}

func (o IpConfigurationResponsePropertiesOutput) ToIpConfigurationResponsePropertiesOutput() IpConfigurationResponsePropertiesOutput {
	return o
}

func (o IpConfigurationResponsePropertiesOutput) ToIpConfigurationResponsePropertiesOutputWithContext(ctx context.Context) IpConfigurationResponsePropertiesOutput {
	return o
}

func (o IpConfigurationResponsePropertiesOutput) PrefixLength() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IpConfigurationResponseProperties) *string { return v.PrefixLength }).(pulumi.StringPtrOutput)
}

func (o IpConfigurationResponsePropertiesOutput) PrivateIPAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IpConfigurationResponseProperties) *string { return v.PrivateIPAddress }).(pulumi.StringPtrOutput)
}

func (o IpConfigurationResponsePropertiesOutput) PrivateIPAllocationMethod() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IpConfigurationResponseProperties) *string { return v.PrivateIPAllocationMethod }).(pulumi.StringPtrOutput)
}

func (o IpConfigurationResponsePropertiesOutput) Subnet() IpConfigurationResponseSubnetPtrOutput {
	return o.ApplyT(func(v IpConfigurationResponseProperties) *IpConfigurationResponseSubnet { return v.Subnet }).(IpConfigurationResponseSubnetPtrOutput)
}

type IpConfigurationResponsePropertiesPtrOutput struct{ *pulumi.OutputState }

func (IpConfigurationResponsePropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IpConfigurationResponseProperties)(nil)).Elem()
}

func (o IpConfigurationResponsePropertiesPtrOutput) ToIpConfigurationResponsePropertiesPtrOutput() IpConfigurationResponsePropertiesPtrOutput {
	return o
}

func (o IpConfigurationResponsePropertiesPtrOutput) ToIpConfigurationResponsePropertiesPtrOutputWithContext(ctx context.Context) IpConfigurationResponsePropertiesPtrOutput {
	return o
}

func (o IpConfigurationResponsePropertiesPtrOutput) Elem() IpConfigurationResponsePropertiesOutput {
	return o.ApplyT(func(v *IpConfigurationResponseProperties) IpConfigurationResponseProperties {
		if v != nil {
			return *v
		}
		var ret IpConfigurationResponseProperties
		return ret
	}).(IpConfigurationResponsePropertiesOutput)
}

func (o IpConfigurationResponsePropertiesPtrOutput) PrefixLength() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IpConfigurationResponseProperties) *string {
		if v == nil {
			return nil
		}
		return v.PrefixLength
	}).(pulumi.StringPtrOutput)
}

func (o IpConfigurationResponsePropertiesPtrOutput) PrivateIPAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IpConfigurationResponseProperties) *string {
		if v == nil {
			return nil
		}
		return v.PrivateIPAddress
	}).(pulumi.StringPtrOutput)
}

func (o IpConfigurationResponsePropertiesPtrOutput) PrivateIPAllocationMethod() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IpConfigurationResponseProperties) *string {
		if v == nil {
			return nil
		}
		return v.PrivateIPAllocationMethod
	}).(pulumi.StringPtrOutput)
}

func (o IpConfigurationResponsePropertiesPtrOutput) Subnet() IpConfigurationResponseSubnetPtrOutput {
	return o.ApplyT(func(v *IpConfigurationResponseProperties) *IpConfigurationResponseSubnet {
		if v == nil {
			return nil
		}
		return v.Subnet
	}).(IpConfigurationResponseSubnetPtrOutput)
}

type IpConfigurationResponseSubnet struct {
	Id *string `pulumi:"id"`
}

type IpConfigurationResponseSubnetOutput struct{ *pulumi.OutputState }

func (IpConfigurationResponseSubnetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IpConfigurationResponseSubnet)(nil)).Elem()
}

func (o IpConfigurationResponseSubnetOutput) ToIpConfigurationResponseSubnetOutput() IpConfigurationResponseSubnetOutput {
	return o
}

func (o IpConfigurationResponseSubnetOutput) ToIpConfigurationResponseSubnetOutputWithContext(ctx context.Context) IpConfigurationResponseSubnetOutput {
	return o
}

func (o IpConfigurationResponseSubnetOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IpConfigurationResponseSubnet) *string { return v.Id }).(pulumi.StringPtrOutput)
}

type IpConfigurationResponseSubnetPtrOutput struct{ *pulumi.OutputState }

func (IpConfigurationResponseSubnetPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IpConfigurationResponseSubnet)(nil)).Elem()
}

func (o IpConfigurationResponseSubnetPtrOutput) ToIpConfigurationResponseSubnetPtrOutput() IpConfigurationResponseSubnetPtrOutput {
	return o
}

func (o IpConfigurationResponseSubnetPtrOutput) ToIpConfigurationResponseSubnetPtrOutputWithContext(ctx context.Context) IpConfigurationResponseSubnetPtrOutput {
	return o
}

func (o IpConfigurationResponseSubnetPtrOutput) Elem() IpConfigurationResponseSubnetOutput {
	return o.ApplyT(func(v *IpConfigurationResponseSubnet) IpConfigurationResponseSubnet {
		if v != nil {
			return *v
		}
		var ret IpConfigurationResponseSubnet
		return ret
	}).(IpConfigurationResponseSubnetOutput)
}

func (o IpConfigurationResponseSubnetPtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IpConfigurationResponseSubnet) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

type IpConfigurationSubnet struct {
	Id *string `pulumi:"id"`
}





type IpConfigurationSubnetInput interface {
	pulumi.Input

	ToIpConfigurationSubnetOutput() IpConfigurationSubnetOutput
	ToIpConfigurationSubnetOutputWithContext(context.Context) IpConfigurationSubnetOutput
}

type IpConfigurationSubnetArgs struct {
	Id pulumi.StringPtrInput `pulumi:"id"`
}

func (IpConfigurationSubnetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*IpConfigurationSubnet)(nil)).Elem()
}

func (i IpConfigurationSubnetArgs) ToIpConfigurationSubnetOutput() IpConfigurationSubnetOutput {
	return i.ToIpConfigurationSubnetOutputWithContext(context.Background())
}

func (i IpConfigurationSubnetArgs) ToIpConfigurationSubnetOutputWithContext(ctx context.Context) IpConfigurationSubnetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IpConfigurationSubnetOutput)
}

func (i IpConfigurationSubnetArgs) ToIpConfigurationSubnetPtrOutput() IpConfigurationSubnetPtrOutput {
	return i.ToIpConfigurationSubnetPtrOutputWithContext(context.Background())
}

func (i IpConfigurationSubnetArgs) ToIpConfigurationSubnetPtrOutputWithContext(ctx context.Context) IpConfigurationSubnetPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IpConfigurationSubnetOutput).ToIpConfigurationSubnetPtrOutputWithContext(ctx)
}









type IpConfigurationSubnetPtrInput interface {
	pulumi.Input

	ToIpConfigurationSubnetPtrOutput() IpConfigurationSubnetPtrOutput
	ToIpConfigurationSubnetPtrOutputWithContext(context.Context) IpConfigurationSubnetPtrOutput
}

type ipConfigurationSubnetPtrType IpConfigurationSubnetArgs

func IpConfigurationSubnetPtr(v *IpConfigurationSubnetArgs) IpConfigurationSubnetPtrInput {
	return (*ipConfigurationSubnetPtrType)(v)
}

func (*ipConfigurationSubnetPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**IpConfigurationSubnet)(nil)).Elem()
}

func (i *ipConfigurationSubnetPtrType) ToIpConfigurationSubnetPtrOutput() IpConfigurationSubnetPtrOutput {
	return i.ToIpConfigurationSubnetPtrOutputWithContext(context.Background())
}

func (i *ipConfigurationSubnetPtrType) ToIpConfigurationSubnetPtrOutputWithContext(ctx context.Context) IpConfigurationSubnetPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IpConfigurationSubnetPtrOutput)
}

type IpConfigurationSubnetOutput struct{ *pulumi.OutputState }

func (IpConfigurationSubnetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IpConfigurationSubnet)(nil)).Elem()
}

func (o IpConfigurationSubnetOutput) ToIpConfigurationSubnetOutput() IpConfigurationSubnetOutput {
	return o
}

func (o IpConfigurationSubnetOutput) ToIpConfigurationSubnetOutputWithContext(ctx context.Context) IpConfigurationSubnetOutput {
	return o
}

func (o IpConfigurationSubnetOutput) ToIpConfigurationSubnetPtrOutput() IpConfigurationSubnetPtrOutput {
	return o.ToIpConfigurationSubnetPtrOutputWithContext(context.Background())
}

func (o IpConfigurationSubnetOutput) ToIpConfigurationSubnetPtrOutputWithContext(ctx context.Context) IpConfigurationSubnetPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v IpConfigurationSubnet) *IpConfigurationSubnet {
		return &v
	}).(IpConfigurationSubnetPtrOutput)
}

func (o IpConfigurationSubnetOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IpConfigurationSubnet) *string { return v.Id }).(pulumi.StringPtrOutput)
}

type IpConfigurationSubnetPtrOutput struct{ *pulumi.OutputState }

func (IpConfigurationSubnetPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IpConfigurationSubnet)(nil)).Elem()
}

func (o IpConfigurationSubnetPtrOutput) ToIpConfigurationSubnetPtrOutput() IpConfigurationSubnetPtrOutput {
	return o
}

func (o IpConfigurationSubnetPtrOutput) ToIpConfigurationSubnetPtrOutputWithContext(ctx context.Context) IpConfigurationSubnetPtrOutput {
	return o
}

func (o IpConfigurationSubnetPtrOutput) Elem() IpConfigurationSubnetOutput {
	return o.ApplyT(func(v *IpConfigurationSubnet) IpConfigurationSubnet {
		if v != nil {
			return *v
		}
		var ret IpConfigurationSubnet
		return ret
	}).(IpConfigurationSubnetOutput)
}

func (o IpConfigurationSubnetPtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IpConfigurationSubnet) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

type MachineExtensionInstanceViewResponseStatus struct {
	Code          string `pulumi:"code"`
	DisplayStatus string `pulumi:"displayStatus"`
	Level         string `pulumi:"level"`
	Message       string `pulumi:"message"`
	Time          string `pulumi:"time"`
}

type MachineExtensionInstanceViewResponseStatusOutput struct{ *pulumi.OutputState }

func (MachineExtensionInstanceViewResponseStatusOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MachineExtensionInstanceViewResponseStatus)(nil)).Elem()
}

func (o MachineExtensionInstanceViewResponseStatusOutput) ToMachineExtensionInstanceViewResponseStatusOutput() MachineExtensionInstanceViewResponseStatusOutput {
	return o
}

func (o MachineExtensionInstanceViewResponseStatusOutput) ToMachineExtensionInstanceViewResponseStatusOutputWithContext(ctx context.Context) MachineExtensionInstanceViewResponseStatusOutput {
	return o
}

func (o MachineExtensionInstanceViewResponseStatusOutput) Code() pulumi.StringOutput {
	return o.ApplyT(func(v MachineExtensionInstanceViewResponseStatus) string { return v.Code }).(pulumi.StringOutput)
}

func (o MachineExtensionInstanceViewResponseStatusOutput) DisplayStatus() pulumi.StringOutput {
	return o.ApplyT(func(v MachineExtensionInstanceViewResponseStatus) string { return v.DisplayStatus }).(pulumi.StringOutput)
}

func (o MachineExtensionInstanceViewResponseStatusOutput) Level() pulumi.StringOutput {
	return o.ApplyT(func(v MachineExtensionInstanceViewResponseStatus) string { return v.Level }).(pulumi.StringOutput)
}

func (o MachineExtensionInstanceViewResponseStatusOutput) Message() pulumi.StringOutput {
	return o.ApplyT(func(v MachineExtensionInstanceViewResponseStatus) string { return v.Message }).(pulumi.StringOutput)
}

func (o MachineExtensionInstanceViewResponseStatusOutput) Time() pulumi.StringOutput {
	return o.ApplyT(func(v MachineExtensionInstanceViewResponseStatus) string { return v.Time }).(pulumi.StringOutput)
}

type MachineExtensionInstanceViewResponseStatusPtrOutput struct{ *pulumi.OutputState }

func (MachineExtensionInstanceViewResponseStatusPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MachineExtensionInstanceViewResponseStatus)(nil)).Elem()
}

func (o MachineExtensionInstanceViewResponseStatusPtrOutput) ToMachineExtensionInstanceViewResponseStatusPtrOutput() MachineExtensionInstanceViewResponseStatusPtrOutput {
	return o
}

func (o MachineExtensionInstanceViewResponseStatusPtrOutput) ToMachineExtensionInstanceViewResponseStatusPtrOutputWithContext(ctx context.Context) MachineExtensionInstanceViewResponseStatusPtrOutput {
	return o
}

func (o MachineExtensionInstanceViewResponseStatusPtrOutput) Elem() MachineExtensionInstanceViewResponseStatusOutput {
	return o.ApplyT(func(v *MachineExtensionInstanceViewResponseStatus) MachineExtensionInstanceViewResponseStatus {
		if v != nil {
			return *v
		}
		var ret MachineExtensionInstanceViewResponseStatus
		return ret
	}).(MachineExtensionInstanceViewResponseStatusOutput)
}

func (o MachineExtensionInstanceViewResponseStatusPtrOutput) Code() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MachineExtensionInstanceViewResponseStatus) *string {
		if v == nil {
			return nil
		}
		return &v.Code
	}).(pulumi.StringPtrOutput)
}

func (o MachineExtensionInstanceViewResponseStatusPtrOutput) DisplayStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MachineExtensionInstanceViewResponseStatus) *string {
		if v == nil {
			return nil
		}
		return &v.DisplayStatus
	}).(pulumi.StringPtrOutput)
}

func (o MachineExtensionInstanceViewResponseStatusPtrOutput) Level() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MachineExtensionInstanceViewResponseStatus) *string {
		if v == nil {
			return nil
		}
		return &v.Level
	}).(pulumi.StringPtrOutput)
}

func (o MachineExtensionInstanceViewResponseStatusPtrOutput) Message() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MachineExtensionInstanceViewResponseStatus) *string {
		if v == nil {
			return nil
		}
		return &v.Message
	}).(pulumi.StringPtrOutput)
}

func (o MachineExtensionInstanceViewResponseStatusPtrOutput) Time() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MachineExtensionInstanceViewResponseStatus) *string {
		if v == nil {
			return nil
		}
		return &v.Time
	}).(pulumi.StringPtrOutput)
}

type MachineExtensionPropertiesResponseInstanceView struct {
	Name               string                                      `pulumi:"name"`
	Status             *MachineExtensionInstanceViewResponseStatus `pulumi:"status"`
	Type               string                                      `pulumi:"type"`
	TypeHandlerVersion string                                      `pulumi:"typeHandlerVersion"`
}

type MachineExtensionPropertiesResponseInstanceViewOutput struct{ *pulumi.OutputState }

func (MachineExtensionPropertiesResponseInstanceViewOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MachineExtensionPropertiesResponseInstanceView)(nil)).Elem()
}

func (o MachineExtensionPropertiesResponseInstanceViewOutput) ToMachineExtensionPropertiesResponseInstanceViewOutput() MachineExtensionPropertiesResponseInstanceViewOutput {
	return o
}

func (o MachineExtensionPropertiesResponseInstanceViewOutput) ToMachineExtensionPropertiesResponseInstanceViewOutputWithContext(ctx context.Context) MachineExtensionPropertiesResponseInstanceViewOutput {
	return o
}

func (o MachineExtensionPropertiesResponseInstanceViewOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v MachineExtensionPropertiesResponseInstanceView) string { return v.Name }).(pulumi.StringOutput)
}

func (o MachineExtensionPropertiesResponseInstanceViewOutput) Status() MachineExtensionInstanceViewResponseStatusPtrOutput {
	return o.ApplyT(func(v MachineExtensionPropertiesResponseInstanceView) *MachineExtensionInstanceViewResponseStatus {
		return v.Status
	}).(MachineExtensionInstanceViewResponseStatusPtrOutput)
}

func (o MachineExtensionPropertiesResponseInstanceViewOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v MachineExtensionPropertiesResponseInstanceView) string { return v.Type }).(pulumi.StringOutput)
}

func (o MachineExtensionPropertiesResponseInstanceViewOutput) TypeHandlerVersion() pulumi.StringOutput {
	return o.ApplyT(func(v MachineExtensionPropertiesResponseInstanceView) string { return v.TypeHandlerVersion }).(pulumi.StringOutput)
}

type MachineExtensionPropertiesResponseInstanceViewPtrOutput struct{ *pulumi.OutputState }

func (MachineExtensionPropertiesResponseInstanceViewPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MachineExtensionPropertiesResponseInstanceView)(nil)).Elem()
}

func (o MachineExtensionPropertiesResponseInstanceViewPtrOutput) ToMachineExtensionPropertiesResponseInstanceViewPtrOutput() MachineExtensionPropertiesResponseInstanceViewPtrOutput {
	return o
}

func (o MachineExtensionPropertiesResponseInstanceViewPtrOutput) ToMachineExtensionPropertiesResponseInstanceViewPtrOutputWithContext(ctx context.Context) MachineExtensionPropertiesResponseInstanceViewPtrOutput {
	return o
}

func (o MachineExtensionPropertiesResponseInstanceViewPtrOutput) Elem() MachineExtensionPropertiesResponseInstanceViewOutput {
	return o.ApplyT(func(v *MachineExtensionPropertiesResponseInstanceView) MachineExtensionPropertiesResponseInstanceView {
		if v != nil {
			return *v
		}
		var ret MachineExtensionPropertiesResponseInstanceView
		return ret
	}).(MachineExtensionPropertiesResponseInstanceViewOutput)
}

func (o MachineExtensionPropertiesResponseInstanceViewPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MachineExtensionPropertiesResponseInstanceView) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

func (o MachineExtensionPropertiesResponseInstanceViewPtrOutput) Status() MachineExtensionInstanceViewResponseStatusPtrOutput {
	return o.ApplyT(func(v *MachineExtensionPropertiesResponseInstanceView) *MachineExtensionInstanceViewResponseStatus {
		if v == nil {
			return nil
		}
		return v.Status
	}).(MachineExtensionInstanceViewResponseStatusPtrOutput)
}

func (o MachineExtensionPropertiesResponseInstanceViewPtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MachineExtensionPropertiesResponseInstanceView) *string {
		if v == nil {
			return nil
		}
		return &v.Type
	}).(pulumi.StringPtrOutput)
}

func (o MachineExtensionPropertiesResponseInstanceViewPtrOutput) TypeHandlerVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MachineExtensionPropertiesResponseInstanceView) *string {
		if v == nil {
			return nil
		}
		return &v.TypeHandlerVersion
	}).(pulumi.StringPtrOutput)
}

type MarketplaceGalleryImageStatusResponse struct {
	DownloadStatus     *MarketplaceGalleryImageStatusResponseDownloadStatus     `pulumi:"downloadStatus"`
	ErrorCode          *string                                                  `pulumi:"errorCode"`
	ErrorMessage       *string                                                  `pulumi:"errorMessage"`
	ProgressPercentage *float64                                                 `pulumi:"progressPercentage"`
	ProvisioningStatus *MarketplaceGalleryImageStatusResponseProvisioningStatus `pulumi:"provisioningStatus"`
}

type MarketplaceGalleryImageStatusResponseOutput struct{ *pulumi.OutputState }

func (MarketplaceGalleryImageStatusResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MarketplaceGalleryImageStatusResponse)(nil)).Elem()
}

func (o MarketplaceGalleryImageStatusResponseOutput) ToMarketplaceGalleryImageStatusResponseOutput() MarketplaceGalleryImageStatusResponseOutput {
	return o
}

func (o MarketplaceGalleryImageStatusResponseOutput) ToMarketplaceGalleryImageStatusResponseOutputWithContext(ctx context.Context) MarketplaceGalleryImageStatusResponseOutput {
	return o
}

func (o MarketplaceGalleryImageStatusResponseOutput) DownloadStatus() MarketplaceGalleryImageStatusResponseDownloadStatusPtrOutput {
	return o.ApplyT(func(v MarketplaceGalleryImageStatusResponse) *MarketplaceGalleryImageStatusResponseDownloadStatus {
		return v.DownloadStatus
	}).(MarketplaceGalleryImageStatusResponseDownloadStatusPtrOutput)
}

func (o MarketplaceGalleryImageStatusResponseOutput) ErrorCode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MarketplaceGalleryImageStatusResponse) *string { return v.ErrorCode }).(pulumi.StringPtrOutput)
}

func (o MarketplaceGalleryImageStatusResponseOutput) ErrorMessage() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MarketplaceGalleryImageStatusResponse) *string { return v.ErrorMessage }).(pulumi.StringPtrOutput)
}

func (o MarketplaceGalleryImageStatusResponseOutput) ProgressPercentage() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v MarketplaceGalleryImageStatusResponse) *float64 { return v.ProgressPercentage }).(pulumi.Float64PtrOutput)
}

func (o MarketplaceGalleryImageStatusResponseOutput) ProvisioningStatus() MarketplaceGalleryImageStatusResponseProvisioningStatusPtrOutput {
	return o.ApplyT(func(v MarketplaceGalleryImageStatusResponse) *MarketplaceGalleryImageStatusResponseProvisioningStatus {
		return v.ProvisioningStatus
	}).(MarketplaceGalleryImageStatusResponseProvisioningStatusPtrOutput)
}

type MarketplaceGalleryImageStatusResponseDownloadStatus struct {
	DownloadSizeInMB *float64 `pulumi:"downloadSizeInMB"`
}

type MarketplaceGalleryImageStatusResponseDownloadStatusOutput struct{ *pulumi.OutputState }

func (MarketplaceGalleryImageStatusResponseDownloadStatusOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MarketplaceGalleryImageStatusResponseDownloadStatus)(nil)).Elem()
}

func (o MarketplaceGalleryImageStatusResponseDownloadStatusOutput) ToMarketplaceGalleryImageStatusResponseDownloadStatusOutput() MarketplaceGalleryImageStatusResponseDownloadStatusOutput {
	return o
}

func (o MarketplaceGalleryImageStatusResponseDownloadStatusOutput) ToMarketplaceGalleryImageStatusResponseDownloadStatusOutputWithContext(ctx context.Context) MarketplaceGalleryImageStatusResponseDownloadStatusOutput {
	return o
}

func (o MarketplaceGalleryImageStatusResponseDownloadStatusOutput) DownloadSizeInMB() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v MarketplaceGalleryImageStatusResponseDownloadStatus) *float64 { return v.DownloadSizeInMB }).(pulumi.Float64PtrOutput)
}

type MarketplaceGalleryImageStatusResponseDownloadStatusPtrOutput struct{ *pulumi.OutputState }

func (MarketplaceGalleryImageStatusResponseDownloadStatusPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MarketplaceGalleryImageStatusResponseDownloadStatus)(nil)).Elem()
}

func (o MarketplaceGalleryImageStatusResponseDownloadStatusPtrOutput) ToMarketplaceGalleryImageStatusResponseDownloadStatusPtrOutput() MarketplaceGalleryImageStatusResponseDownloadStatusPtrOutput {
	return o
}

func (o MarketplaceGalleryImageStatusResponseDownloadStatusPtrOutput) ToMarketplaceGalleryImageStatusResponseDownloadStatusPtrOutputWithContext(ctx context.Context) MarketplaceGalleryImageStatusResponseDownloadStatusPtrOutput {
	return o
}

func (o MarketplaceGalleryImageStatusResponseDownloadStatusPtrOutput) Elem() MarketplaceGalleryImageStatusResponseDownloadStatusOutput {
	return o.ApplyT(func(v *MarketplaceGalleryImageStatusResponseDownloadStatus) MarketplaceGalleryImageStatusResponseDownloadStatus {
		if v != nil {
			return *v
		}
		var ret MarketplaceGalleryImageStatusResponseDownloadStatus
		return ret
	}).(MarketplaceGalleryImageStatusResponseDownloadStatusOutput)
}

func (o MarketplaceGalleryImageStatusResponseDownloadStatusPtrOutput) DownloadSizeInMB() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *MarketplaceGalleryImageStatusResponseDownloadStatus) *float64 {
		if v == nil {
			return nil
		}
		return v.DownloadSizeInMB
	}).(pulumi.Float64PtrOutput)
}

type MarketplaceGalleryImageStatusResponseProvisioningStatus struct {
	OperationId *string `pulumi:"operationId"`
	Status      *string `pulumi:"status"`
}

type MarketplaceGalleryImageStatusResponseProvisioningStatusOutput struct{ *pulumi.OutputState }

func (MarketplaceGalleryImageStatusResponseProvisioningStatusOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MarketplaceGalleryImageStatusResponseProvisioningStatus)(nil)).Elem()
}

func (o MarketplaceGalleryImageStatusResponseProvisioningStatusOutput) ToMarketplaceGalleryImageStatusResponseProvisioningStatusOutput() MarketplaceGalleryImageStatusResponseProvisioningStatusOutput {
	return o
}

func (o MarketplaceGalleryImageStatusResponseProvisioningStatusOutput) ToMarketplaceGalleryImageStatusResponseProvisioningStatusOutputWithContext(ctx context.Context) MarketplaceGalleryImageStatusResponseProvisioningStatusOutput {
	return o
}

func (o MarketplaceGalleryImageStatusResponseProvisioningStatusOutput) OperationId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MarketplaceGalleryImageStatusResponseProvisioningStatus) *string { return v.OperationId }).(pulumi.StringPtrOutput)
}

func (o MarketplaceGalleryImageStatusResponseProvisioningStatusOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MarketplaceGalleryImageStatusResponseProvisioningStatus) *string { return v.Status }).(pulumi.StringPtrOutput)
}

type MarketplaceGalleryImageStatusResponseProvisioningStatusPtrOutput struct{ *pulumi.OutputState }

func (MarketplaceGalleryImageStatusResponseProvisioningStatusPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MarketplaceGalleryImageStatusResponseProvisioningStatus)(nil)).Elem()
}

func (o MarketplaceGalleryImageStatusResponseProvisioningStatusPtrOutput) ToMarketplaceGalleryImageStatusResponseProvisioningStatusPtrOutput() MarketplaceGalleryImageStatusResponseProvisioningStatusPtrOutput {
	return o
}

func (o MarketplaceGalleryImageStatusResponseProvisioningStatusPtrOutput) ToMarketplaceGalleryImageStatusResponseProvisioningStatusPtrOutputWithContext(ctx context.Context) MarketplaceGalleryImageStatusResponseProvisioningStatusPtrOutput {
	return o
}

func (o MarketplaceGalleryImageStatusResponseProvisioningStatusPtrOutput) Elem() MarketplaceGalleryImageStatusResponseProvisioningStatusOutput {
	return o.ApplyT(func(v *MarketplaceGalleryImageStatusResponseProvisioningStatus) MarketplaceGalleryImageStatusResponseProvisioningStatus {
		if v != nil {
			return *v
		}
		var ret MarketplaceGalleryImageStatusResponseProvisioningStatus
		return ret
	}).(MarketplaceGalleryImageStatusResponseProvisioningStatusOutput)
}

func (o MarketplaceGalleryImageStatusResponseProvisioningStatusPtrOutput) OperationId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MarketplaceGalleryImageStatusResponseProvisioningStatus) *string {
		if v == nil {
			return nil
		}
		return v.OperationId
	}).(pulumi.StringPtrOutput)
}

func (o MarketplaceGalleryImageStatusResponseProvisioningStatusPtrOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MarketplaceGalleryImageStatusResponseProvisioningStatus) *string {
		if v == nil {
			return nil
		}
		return v.Status
	}).(pulumi.StringPtrOutput)
}

type NetworkInterfaceStatusResponse struct {
	ErrorCode          *string                                           `pulumi:"errorCode"`
	ErrorMessage       *string                                           `pulumi:"errorMessage"`
	ProvisioningStatus *NetworkInterfaceStatusResponseProvisioningStatus `pulumi:"provisioningStatus"`
}

type NetworkInterfaceStatusResponseOutput struct{ *pulumi.OutputState }

func (NetworkInterfaceStatusResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkInterfaceStatusResponse)(nil)).Elem()
}

func (o NetworkInterfaceStatusResponseOutput) ToNetworkInterfaceStatusResponseOutput() NetworkInterfaceStatusResponseOutput {
	return o
}

func (o NetworkInterfaceStatusResponseOutput) ToNetworkInterfaceStatusResponseOutputWithContext(ctx context.Context) NetworkInterfaceStatusResponseOutput {
	return o
}

func (o NetworkInterfaceStatusResponseOutput) ErrorCode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkInterfaceStatusResponse) *string { return v.ErrorCode }).(pulumi.StringPtrOutput)
}

func (o NetworkInterfaceStatusResponseOutput) ErrorMessage() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkInterfaceStatusResponse) *string { return v.ErrorMessage }).(pulumi.StringPtrOutput)
}

func (o NetworkInterfaceStatusResponseOutput) ProvisioningStatus() NetworkInterfaceStatusResponseProvisioningStatusPtrOutput {
	return o.ApplyT(func(v NetworkInterfaceStatusResponse) *NetworkInterfaceStatusResponseProvisioningStatus {
		return v.ProvisioningStatus
	}).(NetworkInterfaceStatusResponseProvisioningStatusPtrOutput)
}

type NetworkInterfaceStatusResponseProvisioningStatus struct {
	OperationId *string `pulumi:"operationId"`
	Status      *string `pulumi:"status"`
}

type NetworkInterfaceStatusResponseProvisioningStatusOutput struct{ *pulumi.OutputState }

func (NetworkInterfaceStatusResponseProvisioningStatusOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkInterfaceStatusResponseProvisioningStatus)(nil)).Elem()
}

func (o NetworkInterfaceStatusResponseProvisioningStatusOutput) ToNetworkInterfaceStatusResponseProvisioningStatusOutput() NetworkInterfaceStatusResponseProvisioningStatusOutput {
	return o
}

func (o NetworkInterfaceStatusResponseProvisioningStatusOutput) ToNetworkInterfaceStatusResponseProvisioningStatusOutputWithContext(ctx context.Context) NetworkInterfaceStatusResponseProvisioningStatusOutput {
	return o
}

func (o NetworkInterfaceStatusResponseProvisioningStatusOutput) OperationId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkInterfaceStatusResponseProvisioningStatus) *string { return v.OperationId }).(pulumi.StringPtrOutput)
}

func (o NetworkInterfaceStatusResponseProvisioningStatusOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkInterfaceStatusResponseProvisioningStatus) *string { return v.Status }).(pulumi.StringPtrOutput)
}

type NetworkInterfaceStatusResponseProvisioningStatusPtrOutput struct{ *pulumi.OutputState }

func (NetworkInterfaceStatusResponseProvisioningStatusPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkInterfaceStatusResponseProvisioningStatus)(nil)).Elem()
}

func (o NetworkInterfaceStatusResponseProvisioningStatusPtrOutput) ToNetworkInterfaceStatusResponseProvisioningStatusPtrOutput() NetworkInterfaceStatusResponseProvisioningStatusPtrOutput {
	return o
}

func (o NetworkInterfaceStatusResponseProvisioningStatusPtrOutput) ToNetworkInterfaceStatusResponseProvisioningStatusPtrOutputWithContext(ctx context.Context) NetworkInterfaceStatusResponseProvisioningStatusPtrOutput {
	return o
}

func (o NetworkInterfaceStatusResponseProvisioningStatusPtrOutput) Elem() NetworkInterfaceStatusResponseProvisioningStatusOutput {
	return o.ApplyT(func(v *NetworkInterfaceStatusResponseProvisioningStatus) NetworkInterfaceStatusResponseProvisioningStatus {
		if v != nil {
			return *v
		}
		var ret NetworkInterfaceStatusResponseProvisioningStatus
		return ret
	}).(NetworkInterfaceStatusResponseProvisioningStatusOutput)
}

func (o NetworkInterfaceStatusResponseProvisioningStatusPtrOutput) OperationId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkInterfaceStatusResponseProvisioningStatus) *string {
		if v == nil {
			return nil
		}
		return v.OperationId
	}).(pulumi.StringPtrOutput)
}

func (o NetworkInterfaceStatusResponseProvisioningStatusPtrOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkInterfaceStatusResponseProvisioningStatus) *string {
		if v == nil {
			return nil
		}
		return v.Status
	}).(pulumi.StringPtrOutput)
}

type PerNodeExtensionStateResponse struct {
	Extension string `pulumi:"extension"`
	Name      string `pulumi:"name"`
	State     string `pulumi:"state"`
}

type PerNodeExtensionStateResponseOutput struct{ *pulumi.OutputState }

func (PerNodeExtensionStateResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PerNodeExtensionStateResponse)(nil)).Elem()
}

func (o PerNodeExtensionStateResponseOutput) ToPerNodeExtensionStateResponseOutput() PerNodeExtensionStateResponseOutput {
	return o
}

func (o PerNodeExtensionStateResponseOutput) ToPerNodeExtensionStateResponseOutputWithContext(ctx context.Context) PerNodeExtensionStateResponseOutput {
	return o
}

func (o PerNodeExtensionStateResponseOutput) Extension() pulumi.StringOutput {
	return o.ApplyT(func(v PerNodeExtensionStateResponse) string { return v.Extension }).(pulumi.StringOutput)
}

func (o PerNodeExtensionStateResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v PerNodeExtensionStateResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o PerNodeExtensionStateResponseOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v PerNodeExtensionStateResponse) string { return v.State }).(pulumi.StringOutput)
}

type PerNodeExtensionStateResponseArrayOutput struct{ *pulumi.OutputState }

func (PerNodeExtensionStateResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PerNodeExtensionStateResponse)(nil)).Elem()
}

func (o PerNodeExtensionStateResponseArrayOutput) ToPerNodeExtensionStateResponseArrayOutput() PerNodeExtensionStateResponseArrayOutput {
	return o
}

func (o PerNodeExtensionStateResponseArrayOutput) ToPerNodeExtensionStateResponseArrayOutputWithContext(ctx context.Context) PerNodeExtensionStateResponseArrayOutput {
	return o
}

func (o PerNodeExtensionStateResponseArrayOutput) Index(i pulumi.IntInput) PerNodeExtensionStateResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) PerNodeExtensionStateResponse {
		return vs[0].([]PerNodeExtensionStateResponse)[vs[1].(int)]
	}).(PerNodeExtensionStateResponseOutput)
}

type PerNodeStateResponse struct {
	ArcInstance string `pulumi:"arcInstance"`
	Name        string `pulumi:"name"`
	State       string `pulumi:"state"`
}

type PerNodeStateResponseOutput struct{ *pulumi.OutputState }

func (PerNodeStateResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PerNodeStateResponse)(nil)).Elem()
}

func (o PerNodeStateResponseOutput) ToPerNodeStateResponseOutput() PerNodeStateResponseOutput {
	return o
}

func (o PerNodeStateResponseOutput) ToPerNodeStateResponseOutputWithContext(ctx context.Context) PerNodeStateResponseOutput {
	return o
}

func (o PerNodeStateResponseOutput) ArcInstance() pulumi.StringOutput {
	return o.ApplyT(func(v PerNodeStateResponse) string { return v.ArcInstance }).(pulumi.StringOutput)
}

func (o PerNodeStateResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v PerNodeStateResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o PerNodeStateResponseOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v PerNodeStateResponse) string { return v.State }).(pulumi.StringOutput)
}

type PerNodeStateResponseArrayOutput struct{ *pulumi.OutputState }

func (PerNodeStateResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PerNodeStateResponse)(nil)).Elem()
}

func (o PerNodeStateResponseArrayOutput) ToPerNodeStateResponseArrayOutput() PerNodeStateResponseArrayOutput {
	return o
}

func (o PerNodeStateResponseArrayOutput) ToPerNodeStateResponseArrayOutputWithContext(ctx context.Context) PerNodeStateResponseArrayOutput {
	return o
}

func (o PerNodeStateResponseArrayOutput) Index(i pulumi.IntInput) PerNodeStateResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) PerNodeStateResponse {
		return vs[0].([]PerNodeStateResponse)[vs[1].(int)]
	}).(PerNodeStateResponseOutput)
}

type StorageContainerStatusResponse struct {
	AvailableSizeMB    *float64                                          `pulumi:"availableSizeMB"`
	ContainerSizeMB    *float64                                          `pulumi:"containerSizeMB"`
	ErrorCode          *string                                           `pulumi:"errorCode"`
	ErrorMessage       *string                                           `pulumi:"errorMessage"`
	ProvisioningStatus *StorageContainerStatusResponseProvisioningStatus `pulumi:"provisioningStatus"`
}

type StorageContainerStatusResponseOutput struct{ *pulumi.OutputState }

func (StorageContainerStatusResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StorageContainerStatusResponse)(nil)).Elem()
}

func (o StorageContainerStatusResponseOutput) ToStorageContainerStatusResponseOutput() StorageContainerStatusResponseOutput {
	return o
}

func (o StorageContainerStatusResponseOutput) ToStorageContainerStatusResponseOutputWithContext(ctx context.Context) StorageContainerStatusResponseOutput {
	return o
}

func (o StorageContainerStatusResponseOutput) AvailableSizeMB() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v StorageContainerStatusResponse) *float64 { return v.AvailableSizeMB }).(pulumi.Float64PtrOutput)
}

func (o StorageContainerStatusResponseOutput) ContainerSizeMB() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v StorageContainerStatusResponse) *float64 { return v.ContainerSizeMB }).(pulumi.Float64PtrOutput)
}

func (o StorageContainerStatusResponseOutput) ErrorCode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StorageContainerStatusResponse) *string { return v.ErrorCode }).(pulumi.StringPtrOutput)
}

func (o StorageContainerStatusResponseOutput) ErrorMessage() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StorageContainerStatusResponse) *string { return v.ErrorMessage }).(pulumi.StringPtrOutput)
}

func (o StorageContainerStatusResponseOutput) ProvisioningStatus() StorageContainerStatusResponseProvisioningStatusPtrOutput {
	return o.ApplyT(func(v StorageContainerStatusResponse) *StorageContainerStatusResponseProvisioningStatus {
		return v.ProvisioningStatus
	}).(StorageContainerStatusResponseProvisioningStatusPtrOutput)
}

type StorageContainerStatusResponseProvisioningStatus struct {
	OperationId *string `pulumi:"operationId"`
	Status      *string `pulumi:"status"`
}

type StorageContainerStatusResponseProvisioningStatusOutput struct{ *pulumi.OutputState }

func (StorageContainerStatusResponseProvisioningStatusOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StorageContainerStatusResponseProvisioningStatus)(nil)).Elem()
}

func (o StorageContainerStatusResponseProvisioningStatusOutput) ToStorageContainerStatusResponseProvisioningStatusOutput() StorageContainerStatusResponseProvisioningStatusOutput {
	return o
}

func (o StorageContainerStatusResponseProvisioningStatusOutput) ToStorageContainerStatusResponseProvisioningStatusOutputWithContext(ctx context.Context) StorageContainerStatusResponseProvisioningStatusOutput {
	return o
}

func (o StorageContainerStatusResponseProvisioningStatusOutput) OperationId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StorageContainerStatusResponseProvisioningStatus) *string { return v.OperationId }).(pulumi.StringPtrOutput)
}

func (o StorageContainerStatusResponseProvisioningStatusOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StorageContainerStatusResponseProvisioningStatus) *string { return v.Status }).(pulumi.StringPtrOutput)
}

type StorageContainerStatusResponseProvisioningStatusPtrOutput struct{ *pulumi.OutputState }

func (StorageContainerStatusResponseProvisioningStatusPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StorageContainerStatusResponseProvisioningStatus)(nil)).Elem()
}

func (o StorageContainerStatusResponseProvisioningStatusPtrOutput) ToStorageContainerStatusResponseProvisioningStatusPtrOutput() StorageContainerStatusResponseProvisioningStatusPtrOutput {
	return o
}

func (o StorageContainerStatusResponseProvisioningStatusPtrOutput) ToStorageContainerStatusResponseProvisioningStatusPtrOutputWithContext(ctx context.Context) StorageContainerStatusResponseProvisioningStatusPtrOutput {
	return o
}

func (o StorageContainerStatusResponseProvisioningStatusPtrOutput) Elem() StorageContainerStatusResponseProvisioningStatusOutput {
	return o.ApplyT(func(v *StorageContainerStatusResponseProvisioningStatus) StorageContainerStatusResponseProvisioningStatus {
		if v != nil {
			return *v
		}
		var ret StorageContainerStatusResponseProvisioningStatus
		return ret
	}).(StorageContainerStatusResponseProvisioningStatusOutput)
}

func (o StorageContainerStatusResponseProvisioningStatusPtrOutput) OperationId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StorageContainerStatusResponseProvisioningStatus) *string {
		if v == nil {
			return nil
		}
		return v.OperationId
	}).(pulumi.StringPtrOutput)
}

func (o StorageContainerStatusResponseProvisioningStatusPtrOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StorageContainerStatusResponseProvisioningStatus) *string {
		if v == nil {
			return nil
		}
		return v.Status
	}).(pulumi.StringPtrOutput)
}

type StoragecontainersExtendedLocation struct {
	Name *string `pulumi:"name"`
	Type *string `pulumi:"type"`
}





type StoragecontainersExtendedLocationInput interface {
	pulumi.Input

	ToStoragecontainersExtendedLocationOutput() StoragecontainersExtendedLocationOutput
	ToStoragecontainersExtendedLocationOutputWithContext(context.Context) StoragecontainersExtendedLocationOutput
}

type StoragecontainersExtendedLocationArgs struct {
	Name pulumi.StringPtrInput `pulumi:"name"`
	Type pulumi.StringPtrInput `pulumi:"type"`
}

func (StoragecontainersExtendedLocationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*StoragecontainersExtendedLocation)(nil)).Elem()
}

func (i StoragecontainersExtendedLocationArgs) ToStoragecontainersExtendedLocationOutput() StoragecontainersExtendedLocationOutput {
	return i.ToStoragecontainersExtendedLocationOutputWithContext(context.Background())
}

func (i StoragecontainersExtendedLocationArgs) ToStoragecontainersExtendedLocationOutputWithContext(ctx context.Context) StoragecontainersExtendedLocationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StoragecontainersExtendedLocationOutput)
}

func (i StoragecontainersExtendedLocationArgs) ToStoragecontainersExtendedLocationPtrOutput() StoragecontainersExtendedLocationPtrOutput {
	return i.ToStoragecontainersExtendedLocationPtrOutputWithContext(context.Background())
}

func (i StoragecontainersExtendedLocationArgs) ToStoragecontainersExtendedLocationPtrOutputWithContext(ctx context.Context) StoragecontainersExtendedLocationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StoragecontainersExtendedLocationOutput).ToStoragecontainersExtendedLocationPtrOutputWithContext(ctx)
}









type StoragecontainersExtendedLocationPtrInput interface {
	pulumi.Input

	ToStoragecontainersExtendedLocationPtrOutput() StoragecontainersExtendedLocationPtrOutput
	ToStoragecontainersExtendedLocationPtrOutputWithContext(context.Context) StoragecontainersExtendedLocationPtrOutput
}

type storagecontainersExtendedLocationPtrType StoragecontainersExtendedLocationArgs

func StoragecontainersExtendedLocationPtr(v *StoragecontainersExtendedLocationArgs) StoragecontainersExtendedLocationPtrInput {
	return (*storagecontainersExtendedLocationPtrType)(v)
}

func (*storagecontainersExtendedLocationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**StoragecontainersExtendedLocation)(nil)).Elem()
}

func (i *storagecontainersExtendedLocationPtrType) ToStoragecontainersExtendedLocationPtrOutput() StoragecontainersExtendedLocationPtrOutput {
	return i.ToStoragecontainersExtendedLocationPtrOutputWithContext(context.Background())
}

func (i *storagecontainersExtendedLocationPtrType) ToStoragecontainersExtendedLocationPtrOutputWithContext(ctx context.Context) StoragecontainersExtendedLocationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StoragecontainersExtendedLocationPtrOutput)
}

type StoragecontainersExtendedLocationOutput struct{ *pulumi.OutputState }

func (StoragecontainersExtendedLocationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StoragecontainersExtendedLocation)(nil)).Elem()
}

func (o StoragecontainersExtendedLocationOutput) ToStoragecontainersExtendedLocationOutput() StoragecontainersExtendedLocationOutput {
	return o
}

func (o StoragecontainersExtendedLocationOutput) ToStoragecontainersExtendedLocationOutputWithContext(ctx context.Context) StoragecontainersExtendedLocationOutput {
	return o
}

func (o StoragecontainersExtendedLocationOutput) ToStoragecontainersExtendedLocationPtrOutput() StoragecontainersExtendedLocationPtrOutput {
	return o.ToStoragecontainersExtendedLocationPtrOutputWithContext(context.Background())
}

func (o StoragecontainersExtendedLocationOutput) ToStoragecontainersExtendedLocationPtrOutputWithContext(ctx context.Context) StoragecontainersExtendedLocationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v StoragecontainersExtendedLocation) *StoragecontainersExtendedLocation {
		return &v
	}).(StoragecontainersExtendedLocationPtrOutput)
}

func (o StoragecontainersExtendedLocationOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StoragecontainersExtendedLocation) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o StoragecontainersExtendedLocationOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StoragecontainersExtendedLocation) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type StoragecontainersExtendedLocationPtrOutput struct{ *pulumi.OutputState }

func (StoragecontainersExtendedLocationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StoragecontainersExtendedLocation)(nil)).Elem()
}

func (o StoragecontainersExtendedLocationPtrOutput) ToStoragecontainersExtendedLocationPtrOutput() StoragecontainersExtendedLocationPtrOutput {
	return o
}

func (o StoragecontainersExtendedLocationPtrOutput) ToStoragecontainersExtendedLocationPtrOutputWithContext(ctx context.Context) StoragecontainersExtendedLocationPtrOutput {
	return o
}

func (o StoragecontainersExtendedLocationPtrOutput) Elem() StoragecontainersExtendedLocationOutput {
	return o.ApplyT(func(v *StoragecontainersExtendedLocation) StoragecontainersExtendedLocation {
		if v != nil {
			return *v
		}
		var ret StoragecontainersExtendedLocation
		return ret
	}).(StoragecontainersExtendedLocationOutput)
}

func (o StoragecontainersExtendedLocationPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StoragecontainersExtendedLocation) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o StoragecontainersExtendedLocationPtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StoragecontainersExtendedLocation) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

type StoragecontainersResponseExtendedLocation struct {
	Name *string `pulumi:"name"`
	Type *string `pulumi:"type"`
}

type StoragecontainersResponseExtendedLocationOutput struct{ *pulumi.OutputState }

func (StoragecontainersResponseExtendedLocationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StoragecontainersResponseExtendedLocation)(nil)).Elem()
}

func (o StoragecontainersResponseExtendedLocationOutput) ToStoragecontainersResponseExtendedLocationOutput() StoragecontainersResponseExtendedLocationOutput {
	return o
}

func (o StoragecontainersResponseExtendedLocationOutput) ToStoragecontainersResponseExtendedLocationOutputWithContext(ctx context.Context) StoragecontainersResponseExtendedLocationOutput {
	return o
}

func (o StoragecontainersResponseExtendedLocationOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StoragecontainersResponseExtendedLocation) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o StoragecontainersResponseExtendedLocationOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StoragecontainersResponseExtendedLocation) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type StoragecontainersResponseExtendedLocationPtrOutput struct{ *pulumi.OutputState }

func (StoragecontainersResponseExtendedLocationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StoragecontainersResponseExtendedLocation)(nil)).Elem()
}

func (o StoragecontainersResponseExtendedLocationPtrOutput) ToStoragecontainersResponseExtendedLocationPtrOutput() StoragecontainersResponseExtendedLocationPtrOutput {
	return o
}

func (o StoragecontainersResponseExtendedLocationPtrOutput) ToStoragecontainersResponseExtendedLocationPtrOutputWithContext(ctx context.Context) StoragecontainersResponseExtendedLocationPtrOutput {
	return o
}

func (o StoragecontainersResponseExtendedLocationPtrOutput) Elem() StoragecontainersResponseExtendedLocationOutput {
	return o.ApplyT(func(v *StoragecontainersResponseExtendedLocation) StoragecontainersResponseExtendedLocation {
		if v != nil {
			return *v
		}
		var ret StoragecontainersResponseExtendedLocation
		return ret
	}).(StoragecontainersResponseExtendedLocationOutput)
}

func (o StoragecontainersResponseExtendedLocationPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StoragecontainersResponseExtendedLocation) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o StoragecontainersResponseExtendedLocationPtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StoragecontainersResponseExtendedLocation) *string {
		if v == nil {
			return nil
		}
		return v.Type
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

type VirtualHardDiskStatusResponse struct {
	ErrorCode          *string                                          `pulumi:"errorCode"`
	ErrorMessage       *string                                          `pulumi:"errorMessage"`
	ProvisioningStatus *VirtualHardDiskStatusResponseProvisioningStatus `pulumi:"provisioningStatus"`
}

type VirtualHardDiskStatusResponseOutput struct{ *pulumi.OutputState }

func (VirtualHardDiskStatusResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualHardDiskStatusResponse)(nil)).Elem()
}

func (o VirtualHardDiskStatusResponseOutput) ToVirtualHardDiskStatusResponseOutput() VirtualHardDiskStatusResponseOutput {
	return o
}

func (o VirtualHardDiskStatusResponseOutput) ToVirtualHardDiskStatusResponseOutputWithContext(ctx context.Context) VirtualHardDiskStatusResponseOutput {
	return o
}

func (o VirtualHardDiskStatusResponseOutput) ErrorCode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualHardDiskStatusResponse) *string { return v.ErrorCode }).(pulumi.StringPtrOutput)
}

func (o VirtualHardDiskStatusResponseOutput) ErrorMessage() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualHardDiskStatusResponse) *string { return v.ErrorMessage }).(pulumi.StringPtrOutput)
}

func (o VirtualHardDiskStatusResponseOutput) ProvisioningStatus() VirtualHardDiskStatusResponseProvisioningStatusPtrOutput {
	return o.ApplyT(func(v VirtualHardDiskStatusResponse) *VirtualHardDiskStatusResponseProvisioningStatus {
		return v.ProvisioningStatus
	}).(VirtualHardDiskStatusResponseProvisioningStatusPtrOutput)
}

type VirtualHardDiskStatusResponseProvisioningStatus struct {
	OperationId *string `pulumi:"operationId"`
	Status      *string `pulumi:"status"`
}

type VirtualHardDiskStatusResponseProvisioningStatusOutput struct{ *pulumi.OutputState }

func (VirtualHardDiskStatusResponseProvisioningStatusOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualHardDiskStatusResponseProvisioningStatus)(nil)).Elem()
}

func (o VirtualHardDiskStatusResponseProvisioningStatusOutput) ToVirtualHardDiskStatusResponseProvisioningStatusOutput() VirtualHardDiskStatusResponseProvisioningStatusOutput {
	return o
}

func (o VirtualHardDiskStatusResponseProvisioningStatusOutput) ToVirtualHardDiskStatusResponseProvisioningStatusOutputWithContext(ctx context.Context) VirtualHardDiskStatusResponseProvisioningStatusOutput {
	return o
}

func (o VirtualHardDiskStatusResponseProvisioningStatusOutput) OperationId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualHardDiskStatusResponseProvisioningStatus) *string { return v.OperationId }).(pulumi.StringPtrOutput)
}

func (o VirtualHardDiskStatusResponseProvisioningStatusOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualHardDiskStatusResponseProvisioningStatus) *string { return v.Status }).(pulumi.StringPtrOutput)
}

type VirtualHardDiskStatusResponseProvisioningStatusPtrOutput struct{ *pulumi.OutputState }

func (VirtualHardDiskStatusResponseProvisioningStatusPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualHardDiskStatusResponseProvisioningStatus)(nil)).Elem()
}

func (o VirtualHardDiskStatusResponseProvisioningStatusPtrOutput) ToVirtualHardDiskStatusResponseProvisioningStatusPtrOutput() VirtualHardDiskStatusResponseProvisioningStatusPtrOutput {
	return o
}

func (o VirtualHardDiskStatusResponseProvisioningStatusPtrOutput) ToVirtualHardDiskStatusResponseProvisioningStatusPtrOutputWithContext(ctx context.Context) VirtualHardDiskStatusResponseProvisioningStatusPtrOutput {
	return o
}

func (o VirtualHardDiskStatusResponseProvisioningStatusPtrOutput) Elem() VirtualHardDiskStatusResponseProvisioningStatusOutput {
	return o.ApplyT(func(v *VirtualHardDiskStatusResponseProvisioningStatus) VirtualHardDiskStatusResponseProvisioningStatus {
		if v != nil {
			return *v
		}
		var ret VirtualHardDiskStatusResponseProvisioningStatus
		return ret
	}).(VirtualHardDiskStatusResponseProvisioningStatusOutput)
}

func (o VirtualHardDiskStatusResponseProvisioningStatusPtrOutput) OperationId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualHardDiskStatusResponseProvisioningStatus) *string {
		if v == nil {
			return nil
		}
		return v.OperationId
	}).(pulumi.StringPtrOutput)
}

func (o VirtualHardDiskStatusResponseProvisioningStatusPtrOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualHardDiskStatusResponseProvisioningStatus) *string {
		if v == nil {
			return nil
		}
		return v.Status
	}).(pulumi.StringPtrOutput)
}

type VirtualMachineStatusResponse struct {
	ErrorCode          *string                                         `pulumi:"errorCode"`
	ErrorMessage       *string                                         `pulumi:"errorMessage"`
	PowerState         *string                                         `pulumi:"powerState"`
	ProvisioningStatus *VirtualMachineStatusResponseProvisioningStatus `pulumi:"provisioningStatus"`
}

type VirtualMachineStatusResponseOutput struct{ *pulumi.OutputState }

func (VirtualMachineStatusResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualMachineStatusResponse)(nil)).Elem()
}

func (o VirtualMachineStatusResponseOutput) ToVirtualMachineStatusResponseOutput() VirtualMachineStatusResponseOutput {
	return o
}

func (o VirtualMachineStatusResponseOutput) ToVirtualMachineStatusResponseOutputWithContext(ctx context.Context) VirtualMachineStatusResponseOutput {
	return o
}

func (o VirtualMachineStatusResponseOutput) ErrorCode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualMachineStatusResponse) *string { return v.ErrorCode }).(pulumi.StringPtrOutput)
}

func (o VirtualMachineStatusResponseOutput) ErrorMessage() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualMachineStatusResponse) *string { return v.ErrorMessage }).(pulumi.StringPtrOutput)
}

func (o VirtualMachineStatusResponseOutput) PowerState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualMachineStatusResponse) *string { return v.PowerState }).(pulumi.StringPtrOutput)
}

func (o VirtualMachineStatusResponseOutput) ProvisioningStatus() VirtualMachineStatusResponseProvisioningStatusPtrOutput {
	return o.ApplyT(func(v VirtualMachineStatusResponse) *VirtualMachineStatusResponseProvisioningStatus {
		return v.ProvisioningStatus
	}).(VirtualMachineStatusResponseProvisioningStatusPtrOutput)
}

type VirtualMachineStatusResponseProvisioningStatus struct {
	OperationId *string `pulumi:"operationId"`
	Status      *string `pulumi:"status"`
}

type VirtualMachineStatusResponseProvisioningStatusOutput struct{ *pulumi.OutputState }

func (VirtualMachineStatusResponseProvisioningStatusOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualMachineStatusResponseProvisioningStatus)(nil)).Elem()
}

func (o VirtualMachineStatusResponseProvisioningStatusOutput) ToVirtualMachineStatusResponseProvisioningStatusOutput() VirtualMachineStatusResponseProvisioningStatusOutput {
	return o
}

func (o VirtualMachineStatusResponseProvisioningStatusOutput) ToVirtualMachineStatusResponseProvisioningStatusOutputWithContext(ctx context.Context) VirtualMachineStatusResponseProvisioningStatusOutput {
	return o
}

func (o VirtualMachineStatusResponseProvisioningStatusOutput) OperationId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualMachineStatusResponseProvisioningStatus) *string { return v.OperationId }).(pulumi.StringPtrOutput)
}

func (o VirtualMachineStatusResponseProvisioningStatusOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualMachineStatusResponseProvisioningStatus) *string { return v.Status }).(pulumi.StringPtrOutput)
}

type VirtualMachineStatusResponseProvisioningStatusPtrOutput struct{ *pulumi.OutputState }

func (VirtualMachineStatusResponseProvisioningStatusPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualMachineStatusResponseProvisioningStatus)(nil)).Elem()
}

func (o VirtualMachineStatusResponseProvisioningStatusPtrOutput) ToVirtualMachineStatusResponseProvisioningStatusPtrOutput() VirtualMachineStatusResponseProvisioningStatusPtrOutput {
	return o
}

func (o VirtualMachineStatusResponseProvisioningStatusPtrOutput) ToVirtualMachineStatusResponseProvisioningStatusPtrOutputWithContext(ctx context.Context) VirtualMachineStatusResponseProvisioningStatusPtrOutput {
	return o
}

func (o VirtualMachineStatusResponseProvisioningStatusPtrOutput) Elem() VirtualMachineStatusResponseProvisioningStatusOutput {
	return o.ApplyT(func(v *VirtualMachineStatusResponseProvisioningStatus) VirtualMachineStatusResponseProvisioningStatus {
		if v != nil {
			return *v
		}
		var ret VirtualMachineStatusResponseProvisioningStatus
		return ret
	}).(VirtualMachineStatusResponseProvisioningStatusOutput)
}

func (o VirtualMachineStatusResponseProvisioningStatusPtrOutput) OperationId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachineStatusResponseProvisioningStatus) *string {
		if v == nil {
			return nil
		}
		return v.OperationId
	}).(pulumi.StringPtrOutput)
}

func (o VirtualMachineStatusResponseProvisioningStatusPtrOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachineStatusResponseProvisioningStatus) *string {
		if v == nil {
			return nil
		}
		return v.Status
	}).(pulumi.StringPtrOutput)
}

type VirtualNetworkStatusResponse struct {
	ErrorCode          *string                                         `pulumi:"errorCode"`
	ErrorMessage       *string                                         `pulumi:"errorMessage"`
	ProvisioningStatus *VirtualNetworkStatusResponseProvisioningStatus `pulumi:"provisioningStatus"`
}

type VirtualNetworkStatusResponseOutput struct{ *pulumi.OutputState }

func (VirtualNetworkStatusResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualNetworkStatusResponse)(nil)).Elem()
}

func (o VirtualNetworkStatusResponseOutput) ToVirtualNetworkStatusResponseOutput() VirtualNetworkStatusResponseOutput {
	return o
}

func (o VirtualNetworkStatusResponseOutput) ToVirtualNetworkStatusResponseOutputWithContext(ctx context.Context) VirtualNetworkStatusResponseOutput {
	return o
}

func (o VirtualNetworkStatusResponseOutput) ErrorCode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualNetworkStatusResponse) *string { return v.ErrorCode }).(pulumi.StringPtrOutput)
}

func (o VirtualNetworkStatusResponseOutput) ErrorMessage() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualNetworkStatusResponse) *string { return v.ErrorMessage }).(pulumi.StringPtrOutput)
}

func (o VirtualNetworkStatusResponseOutput) ProvisioningStatus() VirtualNetworkStatusResponseProvisioningStatusPtrOutput {
	return o.ApplyT(func(v VirtualNetworkStatusResponse) *VirtualNetworkStatusResponseProvisioningStatus {
		return v.ProvisioningStatus
	}).(VirtualNetworkStatusResponseProvisioningStatusPtrOutput)
}

type VirtualNetworkStatusResponseProvisioningStatus struct {
	OperationId *string `pulumi:"operationId"`
	Status      *string `pulumi:"status"`
}

type VirtualNetworkStatusResponseProvisioningStatusOutput struct{ *pulumi.OutputState }

func (VirtualNetworkStatusResponseProvisioningStatusOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualNetworkStatusResponseProvisioningStatus)(nil)).Elem()
}

func (o VirtualNetworkStatusResponseProvisioningStatusOutput) ToVirtualNetworkStatusResponseProvisioningStatusOutput() VirtualNetworkStatusResponseProvisioningStatusOutput {
	return o
}

func (o VirtualNetworkStatusResponseProvisioningStatusOutput) ToVirtualNetworkStatusResponseProvisioningStatusOutputWithContext(ctx context.Context) VirtualNetworkStatusResponseProvisioningStatusOutput {
	return o
}

func (o VirtualNetworkStatusResponseProvisioningStatusOutput) OperationId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualNetworkStatusResponseProvisioningStatus) *string { return v.OperationId }).(pulumi.StringPtrOutput)
}

func (o VirtualNetworkStatusResponseProvisioningStatusOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualNetworkStatusResponseProvisioningStatus) *string { return v.Status }).(pulumi.StringPtrOutput)
}

type VirtualNetworkStatusResponseProvisioningStatusPtrOutput struct{ *pulumi.OutputState }

func (VirtualNetworkStatusResponseProvisioningStatusPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualNetworkStatusResponseProvisioningStatus)(nil)).Elem()
}

func (o VirtualNetworkStatusResponseProvisioningStatusPtrOutput) ToVirtualNetworkStatusResponseProvisioningStatusPtrOutput() VirtualNetworkStatusResponseProvisioningStatusPtrOutput {
	return o
}

func (o VirtualNetworkStatusResponseProvisioningStatusPtrOutput) ToVirtualNetworkStatusResponseProvisioningStatusPtrOutputWithContext(ctx context.Context) VirtualNetworkStatusResponseProvisioningStatusPtrOutput {
	return o
}

func (o VirtualNetworkStatusResponseProvisioningStatusPtrOutput) Elem() VirtualNetworkStatusResponseProvisioningStatusOutput {
	return o.ApplyT(func(v *VirtualNetworkStatusResponseProvisioningStatus) VirtualNetworkStatusResponseProvisioningStatus {
		if v != nil {
			return *v
		}
		var ret VirtualNetworkStatusResponseProvisioningStatus
		return ret
	}).(VirtualNetworkStatusResponseProvisioningStatusOutput)
}

func (o VirtualNetworkStatusResponseProvisioningStatusPtrOutput) OperationId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualNetworkStatusResponseProvisioningStatus) *string {
		if v == nil {
			return nil
		}
		return v.OperationId
	}).(pulumi.StringPtrOutput)
}

func (o VirtualNetworkStatusResponseProvisioningStatusPtrOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualNetworkStatusResponseProvisioningStatus) *string {
		if v == nil {
			return nil
		}
		return v.Status
	}).(pulumi.StringPtrOutput)
}

type VirtualmachinesPropertiesDataDisks struct {
	Name *string `pulumi:"name"`
}





type VirtualmachinesPropertiesDataDisksInput interface {
	pulumi.Input

	ToVirtualmachinesPropertiesDataDisksOutput() VirtualmachinesPropertiesDataDisksOutput
	ToVirtualmachinesPropertiesDataDisksOutputWithContext(context.Context) VirtualmachinesPropertiesDataDisksOutput
}

type VirtualmachinesPropertiesDataDisksArgs struct {
	Name pulumi.StringPtrInput `pulumi:"name"`
}

func (VirtualmachinesPropertiesDataDisksArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualmachinesPropertiesDataDisks)(nil)).Elem()
}

func (i VirtualmachinesPropertiesDataDisksArgs) ToVirtualmachinesPropertiesDataDisksOutput() VirtualmachinesPropertiesDataDisksOutput {
	return i.ToVirtualmachinesPropertiesDataDisksOutputWithContext(context.Background())
}

func (i VirtualmachinesPropertiesDataDisksArgs) ToVirtualmachinesPropertiesDataDisksOutputWithContext(ctx context.Context) VirtualmachinesPropertiesDataDisksOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualmachinesPropertiesDataDisksOutput)
}





type VirtualmachinesPropertiesDataDisksArrayInput interface {
	pulumi.Input

	ToVirtualmachinesPropertiesDataDisksArrayOutput() VirtualmachinesPropertiesDataDisksArrayOutput
	ToVirtualmachinesPropertiesDataDisksArrayOutputWithContext(context.Context) VirtualmachinesPropertiesDataDisksArrayOutput
}

type VirtualmachinesPropertiesDataDisksArray []VirtualmachinesPropertiesDataDisksInput

func (VirtualmachinesPropertiesDataDisksArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VirtualmachinesPropertiesDataDisks)(nil)).Elem()
}

func (i VirtualmachinesPropertiesDataDisksArray) ToVirtualmachinesPropertiesDataDisksArrayOutput() VirtualmachinesPropertiesDataDisksArrayOutput {
	return i.ToVirtualmachinesPropertiesDataDisksArrayOutputWithContext(context.Background())
}

func (i VirtualmachinesPropertiesDataDisksArray) ToVirtualmachinesPropertiesDataDisksArrayOutputWithContext(ctx context.Context) VirtualmachinesPropertiesDataDisksArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualmachinesPropertiesDataDisksArrayOutput)
}

type VirtualmachinesPropertiesDataDisksOutput struct{ *pulumi.OutputState }

func (VirtualmachinesPropertiesDataDisksOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualmachinesPropertiesDataDisks)(nil)).Elem()
}

func (o VirtualmachinesPropertiesDataDisksOutput) ToVirtualmachinesPropertiesDataDisksOutput() VirtualmachinesPropertiesDataDisksOutput {
	return o
}

func (o VirtualmachinesPropertiesDataDisksOutput) ToVirtualmachinesPropertiesDataDisksOutputWithContext(ctx context.Context) VirtualmachinesPropertiesDataDisksOutput {
	return o
}

func (o VirtualmachinesPropertiesDataDisksOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualmachinesPropertiesDataDisks) *string { return v.Name }).(pulumi.StringPtrOutput)
}

type VirtualmachinesPropertiesDataDisksArrayOutput struct{ *pulumi.OutputState }

func (VirtualmachinesPropertiesDataDisksArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VirtualmachinesPropertiesDataDisks)(nil)).Elem()
}

func (o VirtualmachinesPropertiesDataDisksArrayOutput) ToVirtualmachinesPropertiesDataDisksArrayOutput() VirtualmachinesPropertiesDataDisksArrayOutput {
	return o
}

func (o VirtualmachinesPropertiesDataDisksArrayOutput) ToVirtualmachinesPropertiesDataDisksArrayOutputWithContext(ctx context.Context) VirtualmachinesPropertiesDataDisksArrayOutput {
	return o
}

func (o VirtualmachinesPropertiesDataDisksArrayOutput) Index(i pulumi.IntInput) VirtualmachinesPropertiesDataDisksOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VirtualmachinesPropertiesDataDisks {
		return vs[0].([]VirtualmachinesPropertiesDataDisks)[vs[1].(int)]
	}).(VirtualmachinesPropertiesDataDisksOutput)
}

type VirtualmachinesPropertiesDynamicMemoryConfig struct {
	MaximumMemoryGB    *float64 `pulumi:"maximumMemoryGB"`
	MinimumMemoryGB    *float64 `pulumi:"minimumMemoryGB"`
	TargetMemoryBuffer *int     `pulumi:"targetMemoryBuffer"`
}





type VirtualmachinesPropertiesDynamicMemoryConfigInput interface {
	pulumi.Input

	ToVirtualmachinesPropertiesDynamicMemoryConfigOutput() VirtualmachinesPropertiesDynamicMemoryConfigOutput
	ToVirtualmachinesPropertiesDynamicMemoryConfigOutputWithContext(context.Context) VirtualmachinesPropertiesDynamicMemoryConfigOutput
}

type VirtualmachinesPropertiesDynamicMemoryConfigArgs struct {
	MaximumMemoryGB    pulumi.Float64PtrInput `pulumi:"maximumMemoryGB"`
	MinimumMemoryGB    pulumi.Float64PtrInput `pulumi:"minimumMemoryGB"`
	TargetMemoryBuffer pulumi.IntPtrInput     `pulumi:"targetMemoryBuffer"`
}

func (VirtualmachinesPropertiesDynamicMemoryConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualmachinesPropertiesDynamicMemoryConfig)(nil)).Elem()
}

func (i VirtualmachinesPropertiesDynamicMemoryConfigArgs) ToVirtualmachinesPropertiesDynamicMemoryConfigOutput() VirtualmachinesPropertiesDynamicMemoryConfigOutput {
	return i.ToVirtualmachinesPropertiesDynamicMemoryConfigOutputWithContext(context.Background())
}

func (i VirtualmachinesPropertiesDynamicMemoryConfigArgs) ToVirtualmachinesPropertiesDynamicMemoryConfigOutputWithContext(ctx context.Context) VirtualmachinesPropertiesDynamicMemoryConfigOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualmachinesPropertiesDynamicMemoryConfigOutput)
}

func (i VirtualmachinesPropertiesDynamicMemoryConfigArgs) ToVirtualmachinesPropertiesDynamicMemoryConfigPtrOutput() VirtualmachinesPropertiesDynamicMemoryConfigPtrOutput {
	return i.ToVirtualmachinesPropertiesDynamicMemoryConfigPtrOutputWithContext(context.Background())
}

func (i VirtualmachinesPropertiesDynamicMemoryConfigArgs) ToVirtualmachinesPropertiesDynamicMemoryConfigPtrOutputWithContext(ctx context.Context) VirtualmachinesPropertiesDynamicMemoryConfigPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualmachinesPropertiesDynamicMemoryConfigOutput).ToVirtualmachinesPropertiesDynamicMemoryConfigPtrOutputWithContext(ctx)
}









type VirtualmachinesPropertiesDynamicMemoryConfigPtrInput interface {
	pulumi.Input

	ToVirtualmachinesPropertiesDynamicMemoryConfigPtrOutput() VirtualmachinesPropertiesDynamicMemoryConfigPtrOutput
	ToVirtualmachinesPropertiesDynamicMemoryConfigPtrOutputWithContext(context.Context) VirtualmachinesPropertiesDynamicMemoryConfigPtrOutput
}

type virtualmachinesPropertiesDynamicMemoryConfigPtrType VirtualmachinesPropertiesDynamicMemoryConfigArgs

func VirtualmachinesPropertiesDynamicMemoryConfigPtr(v *VirtualmachinesPropertiesDynamicMemoryConfigArgs) VirtualmachinesPropertiesDynamicMemoryConfigPtrInput {
	return (*virtualmachinesPropertiesDynamicMemoryConfigPtrType)(v)
}

func (*virtualmachinesPropertiesDynamicMemoryConfigPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualmachinesPropertiesDynamicMemoryConfig)(nil)).Elem()
}

func (i *virtualmachinesPropertiesDynamicMemoryConfigPtrType) ToVirtualmachinesPropertiesDynamicMemoryConfigPtrOutput() VirtualmachinesPropertiesDynamicMemoryConfigPtrOutput {
	return i.ToVirtualmachinesPropertiesDynamicMemoryConfigPtrOutputWithContext(context.Background())
}

func (i *virtualmachinesPropertiesDynamicMemoryConfigPtrType) ToVirtualmachinesPropertiesDynamicMemoryConfigPtrOutputWithContext(ctx context.Context) VirtualmachinesPropertiesDynamicMemoryConfigPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualmachinesPropertiesDynamicMemoryConfigPtrOutput)
}

type VirtualmachinesPropertiesDynamicMemoryConfigOutput struct{ *pulumi.OutputState }

func (VirtualmachinesPropertiesDynamicMemoryConfigOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualmachinesPropertiesDynamicMemoryConfig)(nil)).Elem()
}

func (o VirtualmachinesPropertiesDynamicMemoryConfigOutput) ToVirtualmachinesPropertiesDynamicMemoryConfigOutput() VirtualmachinesPropertiesDynamicMemoryConfigOutput {
	return o
}

func (o VirtualmachinesPropertiesDynamicMemoryConfigOutput) ToVirtualmachinesPropertiesDynamicMemoryConfigOutputWithContext(ctx context.Context) VirtualmachinesPropertiesDynamicMemoryConfigOutput {
	return o
}

func (o VirtualmachinesPropertiesDynamicMemoryConfigOutput) ToVirtualmachinesPropertiesDynamicMemoryConfigPtrOutput() VirtualmachinesPropertiesDynamicMemoryConfigPtrOutput {
	return o.ToVirtualmachinesPropertiesDynamicMemoryConfigPtrOutputWithContext(context.Background())
}

func (o VirtualmachinesPropertiesDynamicMemoryConfigOutput) ToVirtualmachinesPropertiesDynamicMemoryConfigPtrOutputWithContext(ctx context.Context) VirtualmachinesPropertiesDynamicMemoryConfigPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v VirtualmachinesPropertiesDynamicMemoryConfig) *VirtualmachinesPropertiesDynamicMemoryConfig {
		return &v
	}).(VirtualmachinesPropertiesDynamicMemoryConfigPtrOutput)
}

func (o VirtualmachinesPropertiesDynamicMemoryConfigOutput) MaximumMemoryGB() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v VirtualmachinesPropertiesDynamicMemoryConfig) *float64 { return v.MaximumMemoryGB }).(pulumi.Float64PtrOutput)
}

func (o VirtualmachinesPropertiesDynamicMemoryConfigOutput) MinimumMemoryGB() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v VirtualmachinesPropertiesDynamicMemoryConfig) *float64 { return v.MinimumMemoryGB }).(pulumi.Float64PtrOutput)
}

func (o VirtualmachinesPropertiesDynamicMemoryConfigOutput) TargetMemoryBuffer() pulumi.IntPtrOutput {
	return o.ApplyT(func(v VirtualmachinesPropertiesDynamicMemoryConfig) *int { return v.TargetMemoryBuffer }).(pulumi.IntPtrOutput)
}

type VirtualmachinesPropertiesDynamicMemoryConfigPtrOutput struct{ *pulumi.OutputState }

func (VirtualmachinesPropertiesDynamicMemoryConfigPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualmachinesPropertiesDynamicMemoryConfig)(nil)).Elem()
}

func (o VirtualmachinesPropertiesDynamicMemoryConfigPtrOutput) ToVirtualmachinesPropertiesDynamicMemoryConfigPtrOutput() VirtualmachinesPropertiesDynamicMemoryConfigPtrOutput {
	return o
}

func (o VirtualmachinesPropertiesDynamicMemoryConfigPtrOutput) ToVirtualmachinesPropertiesDynamicMemoryConfigPtrOutputWithContext(ctx context.Context) VirtualmachinesPropertiesDynamicMemoryConfigPtrOutput {
	return o
}

func (o VirtualmachinesPropertiesDynamicMemoryConfigPtrOutput) Elem() VirtualmachinesPropertiesDynamicMemoryConfigOutput {
	return o.ApplyT(func(v *VirtualmachinesPropertiesDynamicMemoryConfig) VirtualmachinesPropertiesDynamicMemoryConfig {
		if v != nil {
			return *v
		}
		var ret VirtualmachinesPropertiesDynamicMemoryConfig
		return ret
	}).(VirtualmachinesPropertiesDynamicMemoryConfigOutput)
}

func (o VirtualmachinesPropertiesDynamicMemoryConfigPtrOutput) MaximumMemoryGB() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *VirtualmachinesPropertiesDynamicMemoryConfig) *float64 {
		if v == nil {
			return nil
		}
		return v.MaximumMemoryGB
	}).(pulumi.Float64PtrOutput)
}

func (o VirtualmachinesPropertiesDynamicMemoryConfigPtrOutput) MinimumMemoryGB() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *VirtualmachinesPropertiesDynamicMemoryConfig) *float64 {
		if v == nil {
			return nil
		}
		return v.MinimumMemoryGB
	}).(pulumi.Float64PtrOutput)
}

func (o VirtualmachinesPropertiesDynamicMemoryConfigPtrOutput) TargetMemoryBuffer() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *VirtualmachinesPropertiesDynamicMemoryConfig) *int {
		if v == nil {
			return nil
		}
		return v.TargetMemoryBuffer
	}).(pulumi.IntPtrOutput)
}

type VirtualmachinesPropertiesHardwareProfile struct {
	DynamicMemoryConfig *VirtualmachinesPropertiesDynamicMemoryConfig `pulumi:"dynamicMemoryConfig"`
	MemoryGB            *int                                          `pulumi:"memoryGB"`
	Processors          *int                                          `pulumi:"processors"`
	VmSize              *string                                       `pulumi:"vmSize"`
}





type VirtualmachinesPropertiesHardwareProfileInput interface {
	pulumi.Input

	ToVirtualmachinesPropertiesHardwareProfileOutput() VirtualmachinesPropertiesHardwareProfileOutput
	ToVirtualmachinesPropertiesHardwareProfileOutputWithContext(context.Context) VirtualmachinesPropertiesHardwareProfileOutput
}

type VirtualmachinesPropertiesHardwareProfileArgs struct {
	DynamicMemoryConfig VirtualmachinesPropertiesDynamicMemoryConfigPtrInput `pulumi:"dynamicMemoryConfig"`
	MemoryGB            pulumi.IntPtrInput                                   `pulumi:"memoryGB"`
	Processors          pulumi.IntPtrInput                                   `pulumi:"processors"`
	VmSize              pulumi.StringPtrInput                                `pulumi:"vmSize"`
}

func (VirtualmachinesPropertiesHardwareProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualmachinesPropertiesHardwareProfile)(nil)).Elem()
}

func (i VirtualmachinesPropertiesHardwareProfileArgs) ToVirtualmachinesPropertiesHardwareProfileOutput() VirtualmachinesPropertiesHardwareProfileOutput {
	return i.ToVirtualmachinesPropertiesHardwareProfileOutputWithContext(context.Background())
}

func (i VirtualmachinesPropertiesHardwareProfileArgs) ToVirtualmachinesPropertiesHardwareProfileOutputWithContext(ctx context.Context) VirtualmachinesPropertiesHardwareProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualmachinesPropertiesHardwareProfileOutput)
}

func (i VirtualmachinesPropertiesHardwareProfileArgs) ToVirtualmachinesPropertiesHardwareProfilePtrOutput() VirtualmachinesPropertiesHardwareProfilePtrOutput {
	return i.ToVirtualmachinesPropertiesHardwareProfilePtrOutputWithContext(context.Background())
}

func (i VirtualmachinesPropertiesHardwareProfileArgs) ToVirtualmachinesPropertiesHardwareProfilePtrOutputWithContext(ctx context.Context) VirtualmachinesPropertiesHardwareProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualmachinesPropertiesHardwareProfileOutput).ToVirtualmachinesPropertiesHardwareProfilePtrOutputWithContext(ctx)
}









type VirtualmachinesPropertiesHardwareProfilePtrInput interface {
	pulumi.Input

	ToVirtualmachinesPropertiesHardwareProfilePtrOutput() VirtualmachinesPropertiesHardwareProfilePtrOutput
	ToVirtualmachinesPropertiesHardwareProfilePtrOutputWithContext(context.Context) VirtualmachinesPropertiesHardwareProfilePtrOutput
}

type virtualmachinesPropertiesHardwareProfilePtrType VirtualmachinesPropertiesHardwareProfileArgs

func VirtualmachinesPropertiesHardwareProfilePtr(v *VirtualmachinesPropertiesHardwareProfileArgs) VirtualmachinesPropertiesHardwareProfilePtrInput {
	return (*virtualmachinesPropertiesHardwareProfilePtrType)(v)
}

func (*virtualmachinesPropertiesHardwareProfilePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualmachinesPropertiesHardwareProfile)(nil)).Elem()
}

func (i *virtualmachinesPropertiesHardwareProfilePtrType) ToVirtualmachinesPropertiesHardwareProfilePtrOutput() VirtualmachinesPropertiesHardwareProfilePtrOutput {
	return i.ToVirtualmachinesPropertiesHardwareProfilePtrOutputWithContext(context.Background())
}

func (i *virtualmachinesPropertiesHardwareProfilePtrType) ToVirtualmachinesPropertiesHardwareProfilePtrOutputWithContext(ctx context.Context) VirtualmachinesPropertiesHardwareProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualmachinesPropertiesHardwareProfilePtrOutput)
}

type VirtualmachinesPropertiesHardwareProfileOutput struct{ *pulumi.OutputState }

func (VirtualmachinesPropertiesHardwareProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualmachinesPropertiesHardwareProfile)(nil)).Elem()
}

func (o VirtualmachinesPropertiesHardwareProfileOutput) ToVirtualmachinesPropertiesHardwareProfileOutput() VirtualmachinesPropertiesHardwareProfileOutput {
	return o
}

func (o VirtualmachinesPropertiesHardwareProfileOutput) ToVirtualmachinesPropertiesHardwareProfileOutputWithContext(ctx context.Context) VirtualmachinesPropertiesHardwareProfileOutput {
	return o
}

func (o VirtualmachinesPropertiesHardwareProfileOutput) ToVirtualmachinesPropertiesHardwareProfilePtrOutput() VirtualmachinesPropertiesHardwareProfilePtrOutput {
	return o.ToVirtualmachinesPropertiesHardwareProfilePtrOutputWithContext(context.Background())
}

func (o VirtualmachinesPropertiesHardwareProfileOutput) ToVirtualmachinesPropertiesHardwareProfilePtrOutputWithContext(ctx context.Context) VirtualmachinesPropertiesHardwareProfilePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v VirtualmachinesPropertiesHardwareProfile) *VirtualmachinesPropertiesHardwareProfile {
		return &v
	}).(VirtualmachinesPropertiesHardwareProfilePtrOutput)
}

func (o VirtualmachinesPropertiesHardwareProfileOutput) DynamicMemoryConfig() VirtualmachinesPropertiesDynamicMemoryConfigPtrOutput {
	return o.ApplyT(func(v VirtualmachinesPropertiesHardwareProfile) *VirtualmachinesPropertiesDynamicMemoryConfig {
		return v.DynamicMemoryConfig
	}).(VirtualmachinesPropertiesDynamicMemoryConfigPtrOutput)
}

func (o VirtualmachinesPropertiesHardwareProfileOutput) MemoryGB() pulumi.IntPtrOutput {
	return o.ApplyT(func(v VirtualmachinesPropertiesHardwareProfile) *int { return v.MemoryGB }).(pulumi.IntPtrOutput)
}

func (o VirtualmachinesPropertiesHardwareProfileOutput) Processors() pulumi.IntPtrOutput {
	return o.ApplyT(func(v VirtualmachinesPropertiesHardwareProfile) *int { return v.Processors }).(pulumi.IntPtrOutput)
}

func (o VirtualmachinesPropertiesHardwareProfileOutput) VmSize() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualmachinesPropertiesHardwareProfile) *string { return v.VmSize }).(pulumi.StringPtrOutput)
}

type VirtualmachinesPropertiesHardwareProfilePtrOutput struct{ *pulumi.OutputState }

func (VirtualmachinesPropertiesHardwareProfilePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualmachinesPropertiesHardwareProfile)(nil)).Elem()
}

func (o VirtualmachinesPropertiesHardwareProfilePtrOutput) ToVirtualmachinesPropertiesHardwareProfilePtrOutput() VirtualmachinesPropertiesHardwareProfilePtrOutput {
	return o
}

func (o VirtualmachinesPropertiesHardwareProfilePtrOutput) ToVirtualmachinesPropertiesHardwareProfilePtrOutputWithContext(ctx context.Context) VirtualmachinesPropertiesHardwareProfilePtrOutput {
	return o
}

func (o VirtualmachinesPropertiesHardwareProfilePtrOutput) Elem() VirtualmachinesPropertiesHardwareProfileOutput {
	return o.ApplyT(func(v *VirtualmachinesPropertiesHardwareProfile) VirtualmachinesPropertiesHardwareProfile {
		if v != nil {
			return *v
		}
		var ret VirtualmachinesPropertiesHardwareProfile
		return ret
	}).(VirtualmachinesPropertiesHardwareProfileOutput)
}

func (o VirtualmachinesPropertiesHardwareProfilePtrOutput) DynamicMemoryConfig() VirtualmachinesPropertiesDynamicMemoryConfigPtrOutput {
	return o.ApplyT(func(v *VirtualmachinesPropertiesHardwareProfile) *VirtualmachinesPropertiesDynamicMemoryConfig {
		if v == nil {
			return nil
		}
		return v.DynamicMemoryConfig
	}).(VirtualmachinesPropertiesDynamicMemoryConfigPtrOutput)
}

func (o VirtualmachinesPropertiesHardwareProfilePtrOutput) MemoryGB() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *VirtualmachinesPropertiesHardwareProfile) *int {
		if v == nil {
			return nil
		}
		return v.MemoryGB
	}).(pulumi.IntPtrOutput)
}

func (o VirtualmachinesPropertiesHardwareProfilePtrOutput) Processors() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *VirtualmachinesPropertiesHardwareProfile) *int {
		if v == nil {
			return nil
		}
		return v.Processors
	}).(pulumi.IntPtrOutput)
}

func (o VirtualmachinesPropertiesHardwareProfilePtrOutput) VmSize() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualmachinesPropertiesHardwareProfile) *string {
		if v == nil {
			return nil
		}
		return v.VmSize
	}).(pulumi.StringPtrOutput)
}

type VirtualmachinesPropertiesImageReference struct {
	Name *string `pulumi:"name"`
}





type VirtualmachinesPropertiesImageReferenceInput interface {
	pulumi.Input

	ToVirtualmachinesPropertiesImageReferenceOutput() VirtualmachinesPropertiesImageReferenceOutput
	ToVirtualmachinesPropertiesImageReferenceOutputWithContext(context.Context) VirtualmachinesPropertiesImageReferenceOutput
}

type VirtualmachinesPropertiesImageReferenceArgs struct {
	Name pulumi.StringPtrInput `pulumi:"name"`
}

func (VirtualmachinesPropertiesImageReferenceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualmachinesPropertiesImageReference)(nil)).Elem()
}

func (i VirtualmachinesPropertiesImageReferenceArgs) ToVirtualmachinesPropertiesImageReferenceOutput() VirtualmachinesPropertiesImageReferenceOutput {
	return i.ToVirtualmachinesPropertiesImageReferenceOutputWithContext(context.Background())
}

func (i VirtualmachinesPropertiesImageReferenceArgs) ToVirtualmachinesPropertiesImageReferenceOutputWithContext(ctx context.Context) VirtualmachinesPropertiesImageReferenceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualmachinesPropertiesImageReferenceOutput)
}

func (i VirtualmachinesPropertiesImageReferenceArgs) ToVirtualmachinesPropertiesImageReferencePtrOutput() VirtualmachinesPropertiesImageReferencePtrOutput {
	return i.ToVirtualmachinesPropertiesImageReferencePtrOutputWithContext(context.Background())
}

func (i VirtualmachinesPropertiesImageReferenceArgs) ToVirtualmachinesPropertiesImageReferencePtrOutputWithContext(ctx context.Context) VirtualmachinesPropertiesImageReferencePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualmachinesPropertiesImageReferenceOutput).ToVirtualmachinesPropertiesImageReferencePtrOutputWithContext(ctx)
}









type VirtualmachinesPropertiesImageReferencePtrInput interface {
	pulumi.Input

	ToVirtualmachinesPropertiesImageReferencePtrOutput() VirtualmachinesPropertiesImageReferencePtrOutput
	ToVirtualmachinesPropertiesImageReferencePtrOutputWithContext(context.Context) VirtualmachinesPropertiesImageReferencePtrOutput
}

type virtualmachinesPropertiesImageReferencePtrType VirtualmachinesPropertiesImageReferenceArgs

func VirtualmachinesPropertiesImageReferencePtr(v *VirtualmachinesPropertiesImageReferenceArgs) VirtualmachinesPropertiesImageReferencePtrInput {
	return (*virtualmachinesPropertiesImageReferencePtrType)(v)
}

func (*virtualmachinesPropertiesImageReferencePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualmachinesPropertiesImageReference)(nil)).Elem()
}

func (i *virtualmachinesPropertiesImageReferencePtrType) ToVirtualmachinesPropertiesImageReferencePtrOutput() VirtualmachinesPropertiesImageReferencePtrOutput {
	return i.ToVirtualmachinesPropertiesImageReferencePtrOutputWithContext(context.Background())
}

func (i *virtualmachinesPropertiesImageReferencePtrType) ToVirtualmachinesPropertiesImageReferencePtrOutputWithContext(ctx context.Context) VirtualmachinesPropertiesImageReferencePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualmachinesPropertiesImageReferencePtrOutput)
}

type VirtualmachinesPropertiesImageReferenceOutput struct{ *pulumi.OutputState }

func (VirtualmachinesPropertiesImageReferenceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualmachinesPropertiesImageReference)(nil)).Elem()
}

func (o VirtualmachinesPropertiesImageReferenceOutput) ToVirtualmachinesPropertiesImageReferenceOutput() VirtualmachinesPropertiesImageReferenceOutput {
	return o
}

func (o VirtualmachinesPropertiesImageReferenceOutput) ToVirtualmachinesPropertiesImageReferenceOutputWithContext(ctx context.Context) VirtualmachinesPropertiesImageReferenceOutput {
	return o
}

func (o VirtualmachinesPropertiesImageReferenceOutput) ToVirtualmachinesPropertiesImageReferencePtrOutput() VirtualmachinesPropertiesImageReferencePtrOutput {
	return o.ToVirtualmachinesPropertiesImageReferencePtrOutputWithContext(context.Background())
}

func (o VirtualmachinesPropertiesImageReferenceOutput) ToVirtualmachinesPropertiesImageReferencePtrOutputWithContext(ctx context.Context) VirtualmachinesPropertiesImageReferencePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v VirtualmachinesPropertiesImageReference) *VirtualmachinesPropertiesImageReference {
		return &v
	}).(VirtualmachinesPropertiesImageReferencePtrOutput)
}

func (o VirtualmachinesPropertiesImageReferenceOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualmachinesPropertiesImageReference) *string { return v.Name }).(pulumi.StringPtrOutput)
}

type VirtualmachinesPropertiesImageReferencePtrOutput struct{ *pulumi.OutputState }

func (VirtualmachinesPropertiesImageReferencePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualmachinesPropertiesImageReference)(nil)).Elem()
}

func (o VirtualmachinesPropertiesImageReferencePtrOutput) ToVirtualmachinesPropertiesImageReferencePtrOutput() VirtualmachinesPropertiesImageReferencePtrOutput {
	return o
}

func (o VirtualmachinesPropertiesImageReferencePtrOutput) ToVirtualmachinesPropertiesImageReferencePtrOutputWithContext(ctx context.Context) VirtualmachinesPropertiesImageReferencePtrOutput {
	return o
}

func (o VirtualmachinesPropertiesImageReferencePtrOutput) Elem() VirtualmachinesPropertiesImageReferenceOutput {
	return o.ApplyT(func(v *VirtualmachinesPropertiesImageReference) VirtualmachinesPropertiesImageReference {
		if v != nil {
			return *v
		}
		var ret VirtualmachinesPropertiesImageReference
		return ret
	}).(VirtualmachinesPropertiesImageReferenceOutput)
}

func (o VirtualmachinesPropertiesImageReferencePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualmachinesPropertiesImageReference) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

type VirtualmachinesPropertiesLinuxConfiguration struct {
	DisablePasswordAuthentication *bool                         `pulumi:"disablePasswordAuthentication"`
	ProvisionVMAgent              *bool                         `pulumi:"provisionVMAgent"`
	Ssh                           *VirtualmachinesPropertiesSsh `pulumi:"ssh"`
}





type VirtualmachinesPropertiesLinuxConfigurationInput interface {
	pulumi.Input

	ToVirtualmachinesPropertiesLinuxConfigurationOutput() VirtualmachinesPropertiesLinuxConfigurationOutput
	ToVirtualmachinesPropertiesLinuxConfigurationOutputWithContext(context.Context) VirtualmachinesPropertiesLinuxConfigurationOutput
}

type VirtualmachinesPropertiesLinuxConfigurationArgs struct {
	DisablePasswordAuthentication pulumi.BoolPtrInput                  `pulumi:"disablePasswordAuthentication"`
	ProvisionVMAgent              pulumi.BoolPtrInput                  `pulumi:"provisionVMAgent"`
	Ssh                           VirtualmachinesPropertiesSshPtrInput `pulumi:"ssh"`
}

func (VirtualmachinesPropertiesLinuxConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualmachinesPropertiesLinuxConfiguration)(nil)).Elem()
}

func (i VirtualmachinesPropertiesLinuxConfigurationArgs) ToVirtualmachinesPropertiesLinuxConfigurationOutput() VirtualmachinesPropertiesLinuxConfigurationOutput {
	return i.ToVirtualmachinesPropertiesLinuxConfigurationOutputWithContext(context.Background())
}

func (i VirtualmachinesPropertiesLinuxConfigurationArgs) ToVirtualmachinesPropertiesLinuxConfigurationOutputWithContext(ctx context.Context) VirtualmachinesPropertiesLinuxConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualmachinesPropertiesLinuxConfigurationOutput)
}

func (i VirtualmachinesPropertiesLinuxConfigurationArgs) ToVirtualmachinesPropertiesLinuxConfigurationPtrOutput() VirtualmachinesPropertiesLinuxConfigurationPtrOutput {
	return i.ToVirtualmachinesPropertiesLinuxConfigurationPtrOutputWithContext(context.Background())
}

func (i VirtualmachinesPropertiesLinuxConfigurationArgs) ToVirtualmachinesPropertiesLinuxConfigurationPtrOutputWithContext(ctx context.Context) VirtualmachinesPropertiesLinuxConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualmachinesPropertiesLinuxConfigurationOutput).ToVirtualmachinesPropertiesLinuxConfigurationPtrOutputWithContext(ctx)
}









type VirtualmachinesPropertiesLinuxConfigurationPtrInput interface {
	pulumi.Input

	ToVirtualmachinesPropertiesLinuxConfigurationPtrOutput() VirtualmachinesPropertiesLinuxConfigurationPtrOutput
	ToVirtualmachinesPropertiesLinuxConfigurationPtrOutputWithContext(context.Context) VirtualmachinesPropertiesLinuxConfigurationPtrOutput
}

type virtualmachinesPropertiesLinuxConfigurationPtrType VirtualmachinesPropertiesLinuxConfigurationArgs

func VirtualmachinesPropertiesLinuxConfigurationPtr(v *VirtualmachinesPropertiesLinuxConfigurationArgs) VirtualmachinesPropertiesLinuxConfigurationPtrInput {
	return (*virtualmachinesPropertiesLinuxConfigurationPtrType)(v)
}

func (*virtualmachinesPropertiesLinuxConfigurationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualmachinesPropertiesLinuxConfiguration)(nil)).Elem()
}

func (i *virtualmachinesPropertiesLinuxConfigurationPtrType) ToVirtualmachinesPropertiesLinuxConfigurationPtrOutput() VirtualmachinesPropertiesLinuxConfigurationPtrOutput {
	return i.ToVirtualmachinesPropertiesLinuxConfigurationPtrOutputWithContext(context.Background())
}

func (i *virtualmachinesPropertiesLinuxConfigurationPtrType) ToVirtualmachinesPropertiesLinuxConfigurationPtrOutputWithContext(ctx context.Context) VirtualmachinesPropertiesLinuxConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualmachinesPropertiesLinuxConfigurationPtrOutput)
}

type VirtualmachinesPropertiesLinuxConfigurationOutput struct{ *pulumi.OutputState }

func (VirtualmachinesPropertiesLinuxConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualmachinesPropertiesLinuxConfiguration)(nil)).Elem()
}

func (o VirtualmachinesPropertiesLinuxConfigurationOutput) ToVirtualmachinesPropertiesLinuxConfigurationOutput() VirtualmachinesPropertiesLinuxConfigurationOutput {
	return o
}

func (o VirtualmachinesPropertiesLinuxConfigurationOutput) ToVirtualmachinesPropertiesLinuxConfigurationOutputWithContext(ctx context.Context) VirtualmachinesPropertiesLinuxConfigurationOutput {
	return o
}

func (o VirtualmachinesPropertiesLinuxConfigurationOutput) ToVirtualmachinesPropertiesLinuxConfigurationPtrOutput() VirtualmachinesPropertiesLinuxConfigurationPtrOutput {
	return o.ToVirtualmachinesPropertiesLinuxConfigurationPtrOutputWithContext(context.Background())
}

func (o VirtualmachinesPropertiesLinuxConfigurationOutput) ToVirtualmachinesPropertiesLinuxConfigurationPtrOutputWithContext(ctx context.Context) VirtualmachinesPropertiesLinuxConfigurationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v VirtualmachinesPropertiesLinuxConfiguration) *VirtualmachinesPropertiesLinuxConfiguration {
		return &v
	}).(VirtualmachinesPropertiesLinuxConfigurationPtrOutput)
}

func (o VirtualmachinesPropertiesLinuxConfigurationOutput) DisablePasswordAuthentication() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VirtualmachinesPropertiesLinuxConfiguration) *bool { return v.DisablePasswordAuthentication }).(pulumi.BoolPtrOutput)
}

func (o VirtualmachinesPropertiesLinuxConfigurationOutput) ProvisionVMAgent() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VirtualmachinesPropertiesLinuxConfiguration) *bool { return v.ProvisionVMAgent }).(pulumi.BoolPtrOutput)
}

func (o VirtualmachinesPropertiesLinuxConfigurationOutput) Ssh() VirtualmachinesPropertiesSshPtrOutput {
	return o.ApplyT(func(v VirtualmachinesPropertiesLinuxConfiguration) *VirtualmachinesPropertiesSsh { return v.Ssh }).(VirtualmachinesPropertiesSshPtrOutput)
}

type VirtualmachinesPropertiesLinuxConfigurationPtrOutput struct{ *pulumi.OutputState }

func (VirtualmachinesPropertiesLinuxConfigurationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualmachinesPropertiesLinuxConfiguration)(nil)).Elem()
}

func (o VirtualmachinesPropertiesLinuxConfigurationPtrOutput) ToVirtualmachinesPropertiesLinuxConfigurationPtrOutput() VirtualmachinesPropertiesLinuxConfigurationPtrOutput {
	return o
}

func (o VirtualmachinesPropertiesLinuxConfigurationPtrOutput) ToVirtualmachinesPropertiesLinuxConfigurationPtrOutputWithContext(ctx context.Context) VirtualmachinesPropertiesLinuxConfigurationPtrOutput {
	return o
}

func (o VirtualmachinesPropertiesLinuxConfigurationPtrOutput) Elem() VirtualmachinesPropertiesLinuxConfigurationOutput {
	return o.ApplyT(func(v *VirtualmachinesPropertiesLinuxConfiguration) VirtualmachinesPropertiesLinuxConfiguration {
		if v != nil {
			return *v
		}
		var ret VirtualmachinesPropertiesLinuxConfiguration
		return ret
	}).(VirtualmachinesPropertiesLinuxConfigurationOutput)
}

func (o VirtualmachinesPropertiesLinuxConfigurationPtrOutput) DisablePasswordAuthentication() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *VirtualmachinesPropertiesLinuxConfiguration) *bool {
		if v == nil {
			return nil
		}
		return v.DisablePasswordAuthentication
	}).(pulumi.BoolPtrOutput)
}

func (o VirtualmachinesPropertiesLinuxConfigurationPtrOutput) ProvisionVMAgent() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *VirtualmachinesPropertiesLinuxConfiguration) *bool {
		if v == nil {
			return nil
		}
		return v.ProvisionVMAgent
	}).(pulumi.BoolPtrOutput)
}

func (o VirtualmachinesPropertiesLinuxConfigurationPtrOutput) Ssh() VirtualmachinesPropertiesSshPtrOutput {
	return o.ApplyT(func(v *VirtualmachinesPropertiesLinuxConfiguration) *VirtualmachinesPropertiesSsh {
		if v == nil {
			return nil
		}
		return v.Ssh
	}).(VirtualmachinesPropertiesSshPtrOutput)
}

type VirtualmachinesPropertiesNetworkInterfaces struct {
	Id *string `pulumi:"id"`
}





type VirtualmachinesPropertiesNetworkInterfacesInput interface {
	pulumi.Input

	ToVirtualmachinesPropertiesNetworkInterfacesOutput() VirtualmachinesPropertiesNetworkInterfacesOutput
	ToVirtualmachinesPropertiesNetworkInterfacesOutputWithContext(context.Context) VirtualmachinesPropertiesNetworkInterfacesOutput
}

type VirtualmachinesPropertiesNetworkInterfacesArgs struct {
	Id pulumi.StringPtrInput `pulumi:"id"`
}

func (VirtualmachinesPropertiesNetworkInterfacesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualmachinesPropertiesNetworkInterfaces)(nil)).Elem()
}

func (i VirtualmachinesPropertiesNetworkInterfacesArgs) ToVirtualmachinesPropertiesNetworkInterfacesOutput() VirtualmachinesPropertiesNetworkInterfacesOutput {
	return i.ToVirtualmachinesPropertiesNetworkInterfacesOutputWithContext(context.Background())
}

func (i VirtualmachinesPropertiesNetworkInterfacesArgs) ToVirtualmachinesPropertiesNetworkInterfacesOutputWithContext(ctx context.Context) VirtualmachinesPropertiesNetworkInterfacesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualmachinesPropertiesNetworkInterfacesOutput)
}





type VirtualmachinesPropertiesNetworkInterfacesArrayInput interface {
	pulumi.Input

	ToVirtualmachinesPropertiesNetworkInterfacesArrayOutput() VirtualmachinesPropertiesNetworkInterfacesArrayOutput
	ToVirtualmachinesPropertiesNetworkInterfacesArrayOutputWithContext(context.Context) VirtualmachinesPropertiesNetworkInterfacesArrayOutput
}

type VirtualmachinesPropertiesNetworkInterfacesArray []VirtualmachinesPropertiesNetworkInterfacesInput

func (VirtualmachinesPropertiesNetworkInterfacesArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VirtualmachinesPropertiesNetworkInterfaces)(nil)).Elem()
}

func (i VirtualmachinesPropertiesNetworkInterfacesArray) ToVirtualmachinesPropertiesNetworkInterfacesArrayOutput() VirtualmachinesPropertiesNetworkInterfacesArrayOutput {
	return i.ToVirtualmachinesPropertiesNetworkInterfacesArrayOutputWithContext(context.Background())
}

func (i VirtualmachinesPropertiesNetworkInterfacesArray) ToVirtualmachinesPropertiesNetworkInterfacesArrayOutputWithContext(ctx context.Context) VirtualmachinesPropertiesNetworkInterfacesArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualmachinesPropertiesNetworkInterfacesArrayOutput)
}

type VirtualmachinesPropertiesNetworkInterfacesOutput struct{ *pulumi.OutputState }

func (VirtualmachinesPropertiesNetworkInterfacesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualmachinesPropertiesNetworkInterfaces)(nil)).Elem()
}

func (o VirtualmachinesPropertiesNetworkInterfacesOutput) ToVirtualmachinesPropertiesNetworkInterfacesOutput() VirtualmachinesPropertiesNetworkInterfacesOutput {
	return o
}

func (o VirtualmachinesPropertiesNetworkInterfacesOutput) ToVirtualmachinesPropertiesNetworkInterfacesOutputWithContext(ctx context.Context) VirtualmachinesPropertiesNetworkInterfacesOutput {
	return o
}

func (o VirtualmachinesPropertiesNetworkInterfacesOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualmachinesPropertiesNetworkInterfaces) *string { return v.Id }).(pulumi.StringPtrOutput)
}

type VirtualmachinesPropertiesNetworkInterfacesArrayOutput struct{ *pulumi.OutputState }

func (VirtualmachinesPropertiesNetworkInterfacesArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VirtualmachinesPropertiesNetworkInterfaces)(nil)).Elem()
}

func (o VirtualmachinesPropertiesNetworkInterfacesArrayOutput) ToVirtualmachinesPropertiesNetworkInterfacesArrayOutput() VirtualmachinesPropertiesNetworkInterfacesArrayOutput {
	return o
}

func (o VirtualmachinesPropertiesNetworkInterfacesArrayOutput) ToVirtualmachinesPropertiesNetworkInterfacesArrayOutputWithContext(ctx context.Context) VirtualmachinesPropertiesNetworkInterfacesArrayOutput {
	return o
}

func (o VirtualmachinesPropertiesNetworkInterfacesArrayOutput) Index(i pulumi.IntInput) VirtualmachinesPropertiesNetworkInterfacesOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VirtualmachinesPropertiesNetworkInterfaces {
		return vs[0].([]VirtualmachinesPropertiesNetworkInterfaces)[vs[1].(int)]
	}).(VirtualmachinesPropertiesNetworkInterfacesOutput)
}

type VirtualmachinesPropertiesNetworkProfile struct {
	NetworkInterfaces []VirtualmachinesPropertiesNetworkInterfaces `pulumi:"networkInterfaces"`
}





type VirtualmachinesPropertiesNetworkProfileInput interface {
	pulumi.Input

	ToVirtualmachinesPropertiesNetworkProfileOutput() VirtualmachinesPropertiesNetworkProfileOutput
	ToVirtualmachinesPropertiesNetworkProfileOutputWithContext(context.Context) VirtualmachinesPropertiesNetworkProfileOutput
}

type VirtualmachinesPropertiesNetworkProfileArgs struct {
	NetworkInterfaces VirtualmachinesPropertiesNetworkInterfacesArrayInput `pulumi:"networkInterfaces"`
}

func (VirtualmachinesPropertiesNetworkProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualmachinesPropertiesNetworkProfile)(nil)).Elem()
}

func (i VirtualmachinesPropertiesNetworkProfileArgs) ToVirtualmachinesPropertiesNetworkProfileOutput() VirtualmachinesPropertiesNetworkProfileOutput {
	return i.ToVirtualmachinesPropertiesNetworkProfileOutputWithContext(context.Background())
}

func (i VirtualmachinesPropertiesNetworkProfileArgs) ToVirtualmachinesPropertiesNetworkProfileOutputWithContext(ctx context.Context) VirtualmachinesPropertiesNetworkProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualmachinesPropertiesNetworkProfileOutput)
}

func (i VirtualmachinesPropertiesNetworkProfileArgs) ToVirtualmachinesPropertiesNetworkProfilePtrOutput() VirtualmachinesPropertiesNetworkProfilePtrOutput {
	return i.ToVirtualmachinesPropertiesNetworkProfilePtrOutputWithContext(context.Background())
}

func (i VirtualmachinesPropertiesNetworkProfileArgs) ToVirtualmachinesPropertiesNetworkProfilePtrOutputWithContext(ctx context.Context) VirtualmachinesPropertiesNetworkProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualmachinesPropertiesNetworkProfileOutput).ToVirtualmachinesPropertiesNetworkProfilePtrOutputWithContext(ctx)
}









type VirtualmachinesPropertiesNetworkProfilePtrInput interface {
	pulumi.Input

	ToVirtualmachinesPropertiesNetworkProfilePtrOutput() VirtualmachinesPropertiesNetworkProfilePtrOutput
	ToVirtualmachinesPropertiesNetworkProfilePtrOutputWithContext(context.Context) VirtualmachinesPropertiesNetworkProfilePtrOutput
}

type virtualmachinesPropertiesNetworkProfilePtrType VirtualmachinesPropertiesNetworkProfileArgs

func VirtualmachinesPropertiesNetworkProfilePtr(v *VirtualmachinesPropertiesNetworkProfileArgs) VirtualmachinesPropertiesNetworkProfilePtrInput {
	return (*virtualmachinesPropertiesNetworkProfilePtrType)(v)
}

func (*virtualmachinesPropertiesNetworkProfilePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualmachinesPropertiesNetworkProfile)(nil)).Elem()
}

func (i *virtualmachinesPropertiesNetworkProfilePtrType) ToVirtualmachinesPropertiesNetworkProfilePtrOutput() VirtualmachinesPropertiesNetworkProfilePtrOutput {
	return i.ToVirtualmachinesPropertiesNetworkProfilePtrOutputWithContext(context.Background())
}

func (i *virtualmachinesPropertiesNetworkProfilePtrType) ToVirtualmachinesPropertiesNetworkProfilePtrOutputWithContext(ctx context.Context) VirtualmachinesPropertiesNetworkProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualmachinesPropertiesNetworkProfilePtrOutput)
}

type VirtualmachinesPropertiesNetworkProfileOutput struct{ *pulumi.OutputState }

func (VirtualmachinesPropertiesNetworkProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualmachinesPropertiesNetworkProfile)(nil)).Elem()
}

func (o VirtualmachinesPropertiesNetworkProfileOutput) ToVirtualmachinesPropertiesNetworkProfileOutput() VirtualmachinesPropertiesNetworkProfileOutput {
	return o
}

func (o VirtualmachinesPropertiesNetworkProfileOutput) ToVirtualmachinesPropertiesNetworkProfileOutputWithContext(ctx context.Context) VirtualmachinesPropertiesNetworkProfileOutput {
	return o
}

func (o VirtualmachinesPropertiesNetworkProfileOutput) ToVirtualmachinesPropertiesNetworkProfilePtrOutput() VirtualmachinesPropertiesNetworkProfilePtrOutput {
	return o.ToVirtualmachinesPropertiesNetworkProfilePtrOutputWithContext(context.Background())
}

func (o VirtualmachinesPropertiesNetworkProfileOutput) ToVirtualmachinesPropertiesNetworkProfilePtrOutputWithContext(ctx context.Context) VirtualmachinesPropertiesNetworkProfilePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v VirtualmachinesPropertiesNetworkProfile) *VirtualmachinesPropertiesNetworkProfile {
		return &v
	}).(VirtualmachinesPropertiesNetworkProfilePtrOutput)
}

func (o VirtualmachinesPropertiesNetworkProfileOutput) NetworkInterfaces() VirtualmachinesPropertiesNetworkInterfacesArrayOutput {
	return o.ApplyT(func(v VirtualmachinesPropertiesNetworkProfile) []VirtualmachinesPropertiesNetworkInterfaces {
		return v.NetworkInterfaces
	}).(VirtualmachinesPropertiesNetworkInterfacesArrayOutput)
}

type VirtualmachinesPropertiesNetworkProfilePtrOutput struct{ *pulumi.OutputState }

func (VirtualmachinesPropertiesNetworkProfilePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualmachinesPropertiesNetworkProfile)(nil)).Elem()
}

func (o VirtualmachinesPropertiesNetworkProfilePtrOutput) ToVirtualmachinesPropertiesNetworkProfilePtrOutput() VirtualmachinesPropertiesNetworkProfilePtrOutput {
	return o
}

func (o VirtualmachinesPropertiesNetworkProfilePtrOutput) ToVirtualmachinesPropertiesNetworkProfilePtrOutputWithContext(ctx context.Context) VirtualmachinesPropertiesNetworkProfilePtrOutput {
	return o
}

func (o VirtualmachinesPropertiesNetworkProfilePtrOutput) Elem() VirtualmachinesPropertiesNetworkProfileOutput {
	return o.ApplyT(func(v *VirtualmachinesPropertiesNetworkProfile) VirtualmachinesPropertiesNetworkProfile {
		if v != nil {
			return *v
		}
		var ret VirtualmachinesPropertiesNetworkProfile
		return ret
	}).(VirtualmachinesPropertiesNetworkProfileOutput)
}

func (o VirtualmachinesPropertiesNetworkProfilePtrOutput) NetworkInterfaces() VirtualmachinesPropertiesNetworkInterfacesArrayOutput {
	return o.ApplyT(func(v *VirtualmachinesPropertiesNetworkProfile) []VirtualmachinesPropertiesNetworkInterfaces {
		if v == nil {
			return nil
		}
		return v.NetworkInterfaces
	}).(VirtualmachinesPropertiesNetworkInterfacesArrayOutput)
}

type VirtualmachinesPropertiesOsDisk struct {
	Id *string `pulumi:"id"`
}





type VirtualmachinesPropertiesOsDiskInput interface {
	pulumi.Input

	ToVirtualmachinesPropertiesOsDiskOutput() VirtualmachinesPropertiesOsDiskOutput
	ToVirtualmachinesPropertiesOsDiskOutputWithContext(context.Context) VirtualmachinesPropertiesOsDiskOutput
}

type VirtualmachinesPropertiesOsDiskArgs struct {
	Id pulumi.StringPtrInput `pulumi:"id"`
}

func (VirtualmachinesPropertiesOsDiskArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualmachinesPropertiesOsDisk)(nil)).Elem()
}

func (i VirtualmachinesPropertiesOsDiskArgs) ToVirtualmachinesPropertiesOsDiskOutput() VirtualmachinesPropertiesOsDiskOutput {
	return i.ToVirtualmachinesPropertiesOsDiskOutputWithContext(context.Background())
}

func (i VirtualmachinesPropertiesOsDiskArgs) ToVirtualmachinesPropertiesOsDiskOutputWithContext(ctx context.Context) VirtualmachinesPropertiesOsDiskOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualmachinesPropertiesOsDiskOutput)
}

func (i VirtualmachinesPropertiesOsDiskArgs) ToVirtualmachinesPropertiesOsDiskPtrOutput() VirtualmachinesPropertiesOsDiskPtrOutput {
	return i.ToVirtualmachinesPropertiesOsDiskPtrOutputWithContext(context.Background())
}

func (i VirtualmachinesPropertiesOsDiskArgs) ToVirtualmachinesPropertiesOsDiskPtrOutputWithContext(ctx context.Context) VirtualmachinesPropertiesOsDiskPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualmachinesPropertiesOsDiskOutput).ToVirtualmachinesPropertiesOsDiskPtrOutputWithContext(ctx)
}









type VirtualmachinesPropertiesOsDiskPtrInput interface {
	pulumi.Input

	ToVirtualmachinesPropertiesOsDiskPtrOutput() VirtualmachinesPropertiesOsDiskPtrOutput
	ToVirtualmachinesPropertiesOsDiskPtrOutputWithContext(context.Context) VirtualmachinesPropertiesOsDiskPtrOutput
}

type virtualmachinesPropertiesOsDiskPtrType VirtualmachinesPropertiesOsDiskArgs

func VirtualmachinesPropertiesOsDiskPtr(v *VirtualmachinesPropertiesOsDiskArgs) VirtualmachinesPropertiesOsDiskPtrInput {
	return (*virtualmachinesPropertiesOsDiskPtrType)(v)
}

func (*virtualmachinesPropertiesOsDiskPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualmachinesPropertiesOsDisk)(nil)).Elem()
}

func (i *virtualmachinesPropertiesOsDiskPtrType) ToVirtualmachinesPropertiesOsDiskPtrOutput() VirtualmachinesPropertiesOsDiskPtrOutput {
	return i.ToVirtualmachinesPropertiesOsDiskPtrOutputWithContext(context.Background())
}

func (i *virtualmachinesPropertiesOsDiskPtrType) ToVirtualmachinesPropertiesOsDiskPtrOutputWithContext(ctx context.Context) VirtualmachinesPropertiesOsDiskPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualmachinesPropertiesOsDiskPtrOutput)
}

type VirtualmachinesPropertiesOsDiskOutput struct{ *pulumi.OutputState }

func (VirtualmachinesPropertiesOsDiskOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualmachinesPropertiesOsDisk)(nil)).Elem()
}

func (o VirtualmachinesPropertiesOsDiskOutput) ToVirtualmachinesPropertiesOsDiskOutput() VirtualmachinesPropertiesOsDiskOutput {
	return o
}

func (o VirtualmachinesPropertiesOsDiskOutput) ToVirtualmachinesPropertiesOsDiskOutputWithContext(ctx context.Context) VirtualmachinesPropertiesOsDiskOutput {
	return o
}

func (o VirtualmachinesPropertiesOsDiskOutput) ToVirtualmachinesPropertiesOsDiskPtrOutput() VirtualmachinesPropertiesOsDiskPtrOutput {
	return o.ToVirtualmachinesPropertiesOsDiskPtrOutputWithContext(context.Background())
}

func (o VirtualmachinesPropertiesOsDiskOutput) ToVirtualmachinesPropertiesOsDiskPtrOutputWithContext(ctx context.Context) VirtualmachinesPropertiesOsDiskPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v VirtualmachinesPropertiesOsDisk) *VirtualmachinesPropertiesOsDisk {
		return &v
	}).(VirtualmachinesPropertiesOsDiskPtrOutput)
}

func (o VirtualmachinesPropertiesOsDiskOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualmachinesPropertiesOsDisk) *string { return v.Id }).(pulumi.StringPtrOutput)
}

type VirtualmachinesPropertiesOsDiskPtrOutput struct{ *pulumi.OutputState }

func (VirtualmachinesPropertiesOsDiskPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualmachinesPropertiesOsDisk)(nil)).Elem()
}

func (o VirtualmachinesPropertiesOsDiskPtrOutput) ToVirtualmachinesPropertiesOsDiskPtrOutput() VirtualmachinesPropertiesOsDiskPtrOutput {
	return o
}

func (o VirtualmachinesPropertiesOsDiskPtrOutput) ToVirtualmachinesPropertiesOsDiskPtrOutputWithContext(ctx context.Context) VirtualmachinesPropertiesOsDiskPtrOutput {
	return o
}

func (o VirtualmachinesPropertiesOsDiskPtrOutput) Elem() VirtualmachinesPropertiesOsDiskOutput {
	return o.ApplyT(func(v *VirtualmachinesPropertiesOsDisk) VirtualmachinesPropertiesOsDisk {
		if v != nil {
			return *v
		}
		var ret VirtualmachinesPropertiesOsDisk
		return ret
	}).(VirtualmachinesPropertiesOsDiskOutput)
}

func (o VirtualmachinesPropertiesOsDiskPtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualmachinesPropertiesOsDisk) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

type VirtualmachinesPropertiesOsProfile struct {
	AdminPassword        *string                                        `pulumi:"adminPassword"`
	AdminUsername        *string                                        `pulumi:"adminUsername"`
	ComputerName         *string                                        `pulumi:"computerName"`
	LinuxConfiguration   *VirtualmachinesPropertiesLinuxConfiguration   `pulumi:"linuxConfiguration"`
	OsType               *string                                        `pulumi:"osType"`
	WindowsConfiguration *VirtualmachinesPropertiesWindowsConfiguration `pulumi:"windowsConfiguration"`
}





type VirtualmachinesPropertiesOsProfileInput interface {
	pulumi.Input

	ToVirtualmachinesPropertiesOsProfileOutput() VirtualmachinesPropertiesOsProfileOutput
	ToVirtualmachinesPropertiesOsProfileOutputWithContext(context.Context) VirtualmachinesPropertiesOsProfileOutput
}

type VirtualmachinesPropertiesOsProfileArgs struct {
	AdminPassword        pulumi.StringPtrInput                                 `pulumi:"adminPassword"`
	AdminUsername        pulumi.StringPtrInput                                 `pulumi:"adminUsername"`
	ComputerName         pulumi.StringPtrInput                                 `pulumi:"computerName"`
	LinuxConfiguration   VirtualmachinesPropertiesLinuxConfigurationPtrInput   `pulumi:"linuxConfiguration"`
	OsType               pulumi.StringPtrInput                                 `pulumi:"osType"`
	WindowsConfiguration VirtualmachinesPropertiesWindowsConfigurationPtrInput `pulumi:"windowsConfiguration"`
}

func (VirtualmachinesPropertiesOsProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualmachinesPropertiesOsProfile)(nil)).Elem()
}

func (i VirtualmachinesPropertiesOsProfileArgs) ToVirtualmachinesPropertiesOsProfileOutput() VirtualmachinesPropertiesOsProfileOutput {
	return i.ToVirtualmachinesPropertiesOsProfileOutputWithContext(context.Background())
}

func (i VirtualmachinesPropertiesOsProfileArgs) ToVirtualmachinesPropertiesOsProfileOutputWithContext(ctx context.Context) VirtualmachinesPropertiesOsProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualmachinesPropertiesOsProfileOutput)
}

func (i VirtualmachinesPropertiesOsProfileArgs) ToVirtualmachinesPropertiesOsProfilePtrOutput() VirtualmachinesPropertiesOsProfilePtrOutput {
	return i.ToVirtualmachinesPropertiesOsProfilePtrOutputWithContext(context.Background())
}

func (i VirtualmachinesPropertiesOsProfileArgs) ToVirtualmachinesPropertiesOsProfilePtrOutputWithContext(ctx context.Context) VirtualmachinesPropertiesOsProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualmachinesPropertiesOsProfileOutput).ToVirtualmachinesPropertiesOsProfilePtrOutputWithContext(ctx)
}









type VirtualmachinesPropertiesOsProfilePtrInput interface {
	pulumi.Input

	ToVirtualmachinesPropertiesOsProfilePtrOutput() VirtualmachinesPropertiesOsProfilePtrOutput
	ToVirtualmachinesPropertiesOsProfilePtrOutputWithContext(context.Context) VirtualmachinesPropertiesOsProfilePtrOutput
}

type virtualmachinesPropertiesOsProfilePtrType VirtualmachinesPropertiesOsProfileArgs

func VirtualmachinesPropertiesOsProfilePtr(v *VirtualmachinesPropertiesOsProfileArgs) VirtualmachinesPropertiesOsProfilePtrInput {
	return (*virtualmachinesPropertiesOsProfilePtrType)(v)
}

func (*virtualmachinesPropertiesOsProfilePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualmachinesPropertiesOsProfile)(nil)).Elem()
}

func (i *virtualmachinesPropertiesOsProfilePtrType) ToVirtualmachinesPropertiesOsProfilePtrOutput() VirtualmachinesPropertiesOsProfilePtrOutput {
	return i.ToVirtualmachinesPropertiesOsProfilePtrOutputWithContext(context.Background())
}

func (i *virtualmachinesPropertiesOsProfilePtrType) ToVirtualmachinesPropertiesOsProfilePtrOutputWithContext(ctx context.Context) VirtualmachinesPropertiesOsProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualmachinesPropertiesOsProfilePtrOutput)
}

type VirtualmachinesPropertiesOsProfileOutput struct{ *pulumi.OutputState }

func (VirtualmachinesPropertiesOsProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualmachinesPropertiesOsProfile)(nil)).Elem()
}

func (o VirtualmachinesPropertiesOsProfileOutput) ToVirtualmachinesPropertiesOsProfileOutput() VirtualmachinesPropertiesOsProfileOutput {
	return o
}

func (o VirtualmachinesPropertiesOsProfileOutput) ToVirtualmachinesPropertiesOsProfileOutputWithContext(ctx context.Context) VirtualmachinesPropertiesOsProfileOutput {
	return o
}

func (o VirtualmachinesPropertiesOsProfileOutput) ToVirtualmachinesPropertiesOsProfilePtrOutput() VirtualmachinesPropertiesOsProfilePtrOutput {
	return o.ToVirtualmachinesPropertiesOsProfilePtrOutputWithContext(context.Background())
}

func (o VirtualmachinesPropertiesOsProfileOutput) ToVirtualmachinesPropertiesOsProfilePtrOutputWithContext(ctx context.Context) VirtualmachinesPropertiesOsProfilePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v VirtualmachinesPropertiesOsProfile) *VirtualmachinesPropertiesOsProfile {
		return &v
	}).(VirtualmachinesPropertiesOsProfilePtrOutput)
}

func (o VirtualmachinesPropertiesOsProfileOutput) AdminPassword() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualmachinesPropertiesOsProfile) *string { return v.AdminPassword }).(pulumi.StringPtrOutput)
}

func (o VirtualmachinesPropertiesOsProfileOutput) AdminUsername() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualmachinesPropertiesOsProfile) *string { return v.AdminUsername }).(pulumi.StringPtrOutput)
}

func (o VirtualmachinesPropertiesOsProfileOutput) ComputerName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualmachinesPropertiesOsProfile) *string { return v.ComputerName }).(pulumi.StringPtrOutput)
}

func (o VirtualmachinesPropertiesOsProfileOutput) LinuxConfiguration() VirtualmachinesPropertiesLinuxConfigurationPtrOutput {
	return o.ApplyT(func(v VirtualmachinesPropertiesOsProfile) *VirtualmachinesPropertiesLinuxConfiguration {
		return v.LinuxConfiguration
	}).(VirtualmachinesPropertiesLinuxConfigurationPtrOutput)
}

func (o VirtualmachinesPropertiesOsProfileOutput) OsType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualmachinesPropertiesOsProfile) *string { return v.OsType }).(pulumi.StringPtrOutput)
}

func (o VirtualmachinesPropertiesOsProfileOutput) WindowsConfiguration() VirtualmachinesPropertiesWindowsConfigurationPtrOutput {
	return o.ApplyT(func(v VirtualmachinesPropertiesOsProfile) *VirtualmachinesPropertiesWindowsConfiguration {
		return v.WindowsConfiguration
	}).(VirtualmachinesPropertiesWindowsConfigurationPtrOutput)
}

type VirtualmachinesPropertiesOsProfilePtrOutput struct{ *pulumi.OutputState }

func (VirtualmachinesPropertiesOsProfilePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualmachinesPropertiesOsProfile)(nil)).Elem()
}

func (o VirtualmachinesPropertiesOsProfilePtrOutput) ToVirtualmachinesPropertiesOsProfilePtrOutput() VirtualmachinesPropertiesOsProfilePtrOutput {
	return o
}

func (o VirtualmachinesPropertiesOsProfilePtrOutput) ToVirtualmachinesPropertiesOsProfilePtrOutputWithContext(ctx context.Context) VirtualmachinesPropertiesOsProfilePtrOutput {
	return o
}

func (o VirtualmachinesPropertiesOsProfilePtrOutput) Elem() VirtualmachinesPropertiesOsProfileOutput {
	return o.ApplyT(func(v *VirtualmachinesPropertiesOsProfile) VirtualmachinesPropertiesOsProfile {
		if v != nil {
			return *v
		}
		var ret VirtualmachinesPropertiesOsProfile
		return ret
	}).(VirtualmachinesPropertiesOsProfileOutput)
}

func (o VirtualmachinesPropertiesOsProfilePtrOutput) AdminPassword() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualmachinesPropertiesOsProfile) *string {
		if v == nil {
			return nil
		}
		return v.AdminPassword
	}).(pulumi.StringPtrOutput)
}

func (o VirtualmachinesPropertiesOsProfilePtrOutput) AdminUsername() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualmachinesPropertiesOsProfile) *string {
		if v == nil {
			return nil
		}
		return v.AdminUsername
	}).(pulumi.StringPtrOutput)
}

func (o VirtualmachinesPropertiesOsProfilePtrOutput) ComputerName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualmachinesPropertiesOsProfile) *string {
		if v == nil {
			return nil
		}
		return v.ComputerName
	}).(pulumi.StringPtrOutput)
}

func (o VirtualmachinesPropertiesOsProfilePtrOutput) LinuxConfiguration() VirtualmachinesPropertiesLinuxConfigurationPtrOutput {
	return o.ApplyT(func(v *VirtualmachinesPropertiesOsProfile) *VirtualmachinesPropertiesLinuxConfiguration {
		if v == nil {
			return nil
		}
		return v.LinuxConfiguration
	}).(VirtualmachinesPropertiesLinuxConfigurationPtrOutput)
}

func (o VirtualmachinesPropertiesOsProfilePtrOutput) OsType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualmachinesPropertiesOsProfile) *string {
		if v == nil {
			return nil
		}
		return v.OsType
	}).(pulumi.StringPtrOutput)
}

func (o VirtualmachinesPropertiesOsProfilePtrOutput) WindowsConfiguration() VirtualmachinesPropertiesWindowsConfigurationPtrOutput {
	return o.ApplyT(func(v *VirtualmachinesPropertiesOsProfile) *VirtualmachinesPropertiesWindowsConfiguration {
		if v == nil {
			return nil
		}
		return v.WindowsConfiguration
	}).(VirtualmachinesPropertiesWindowsConfigurationPtrOutput)
}

type VirtualmachinesPropertiesPublicKeys struct {
	KeyData *string `pulumi:"keyData"`
	Path    *string `pulumi:"path"`
}





type VirtualmachinesPropertiesPublicKeysInput interface {
	pulumi.Input

	ToVirtualmachinesPropertiesPublicKeysOutput() VirtualmachinesPropertiesPublicKeysOutput
	ToVirtualmachinesPropertiesPublicKeysOutputWithContext(context.Context) VirtualmachinesPropertiesPublicKeysOutput
}

type VirtualmachinesPropertiesPublicKeysArgs struct {
	KeyData pulumi.StringPtrInput `pulumi:"keyData"`
	Path    pulumi.StringPtrInput `pulumi:"path"`
}

func (VirtualmachinesPropertiesPublicKeysArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualmachinesPropertiesPublicKeys)(nil)).Elem()
}

func (i VirtualmachinesPropertiesPublicKeysArgs) ToVirtualmachinesPropertiesPublicKeysOutput() VirtualmachinesPropertiesPublicKeysOutput {
	return i.ToVirtualmachinesPropertiesPublicKeysOutputWithContext(context.Background())
}

func (i VirtualmachinesPropertiesPublicKeysArgs) ToVirtualmachinesPropertiesPublicKeysOutputWithContext(ctx context.Context) VirtualmachinesPropertiesPublicKeysOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualmachinesPropertiesPublicKeysOutput)
}





type VirtualmachinesPropertiesPublicKeysArrayInput interface {
	pulumi.Input

	ToVirtualmachinesPropertiesPublicKeysArrayOutput() VirtualmachinesPropertiesPublicKeysArrayOutput
	ToVirtualmachinesPropertiesPublicKeysArrayOutputWithContext(context.Context) VirtualmachinesPropertiesPublicKeysArrayOutput
}

type VirtualmachinesPropertiesPublicKeysArray []VirtualmachinesPropertiesPublicKeysInput

func (VirtualmachinesPropertiesPublicKeysArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VirtualmachinesPropertiesPublicKeys)(nil)).Elem()
}

func (i VirtualmachinesPropertiesPublicKeysArray) ToVirtualmachinesPropertiesPublicKeysArrayOutput() VirtualmachinesPropertiesPublicKeysArrayOutput {
	return i.ToVirtualmachinesPropertiesPublicKeysArrayOutputWithContext(context.Background())
}

func (i VirtualmachinesPropertiesPublicKeysArray) ToVirtualmachinesPropertiesPublicKeysArrayOutputWithContext(ctx context.Context) VirtualmachinesPropertiesPublicKeysArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualmachinesPropertiesPublicKeysArrayOutput)
}

type VirtualmachinesPropertiesPublicKeysOutput struct{ *pulumi.OutputState }

func (VirtualmachinesPropertiesPublicKeysOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualmachinesPropertiesPublicKeys)(nil)).Elem()
}

func (o VirtualmachinesPropertiesPublicKeysOutput) ToVirtualmachinesPropertiesPublicKeysOutput() VirtualmachinesPropertiesPublicKeysOutput {
	return o
}

func (o VirtualmachinesPropertiesPublicKeysOutput) ToVirtualmachinesPropertiesPublicKeysOutputWithContext(ctx context.Context) VirtualmachinesPropertiesPublicKeysOutput {
	return o
}

func (o VirtualmachinesPropertiesPublicKeysOutput) KeyData() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualmachinesPropertiesPublicKeys) *string { return v.KeyData }).(pulumi.StringPtrOutput)
}

func (o VirtualmachinesPropertiesPublicKeysOutput) Path() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualmachinesPropertiesPublicKeys) *string { return v.Path }).(pulumi.StringPtrOutput)
}

type VirtualmachinesPropertiesPublicKeysArrayOutput struct{ *pulumi.OutputState }

func (VirtualmachinesPropertiesPublicKeysArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VirtualmachinesPropertiesPublicKeys)(nil)).Elem()
}

func (o VirtualmachinesPropertiesPublicKeysArrayOutput) ToVirtualmachinesPropertiesPublicKeysArrayOutput() VirtualmachinesPropertiesPublicKeysArrayOutput {
	return o
}

func (o VirtualmachinesPropertiesPublicKeysArrayOutput) ToVirtualmachinesPropertiesPublicKeysArrayOutputWithContext(ctx context.Context) VirtualmachinesPropertiesPublicKeysArrayOutput {
	return o
}

func (o VirtualmachinesPropertiesPublicKeysArrayOutput) Index(i pulumi.IntInput) VirtualmachinesPropertiesPublicKeysOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VirtualmachinesPropertiesPublicKeys {
		return vs[0].([]VirtualmachinesPropertiesPublicKeys)[vs[1].(int)]
	}).(VirtualmachinesPropertiesPublicKeysOutput)
}

type VirtualmachinesPropertiesPublicKeysPublicKeys struct {
	KeyData *string `pulumi:"keyData"`
	Path    *string `pulumi:"path"`
}





type VirtualmachinesPropertiesPublicKeysPublicKeysInput interface {
	pulumi.Input

	ToVirtualmachinesPropertiesPublicKeysPublicKeysOutput() VirtualmachinesPropertiesPublicKeysPublicKeysOutput
	ToVirtualmachinesPropertiesPublicKeysPublicKeysOutputWithContext(context.Context) VirtualmachinesPropertiesPublicKeysPublicKeysOutput
}

type VirtualmachinesPropertiesPublicKeysPublicKeysArgs struct {
	KeyData pulumi.StringPtrInput `pulumi:"keyData"`
	Path    pulumi.StringPtrInput `pulumi:"path"`
}

func (VirtualmachinesPropertiesPublicKeysPublicKeysArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualmachinesPropertiesPublicKeysPublicKeys)(nil)).Elem()
}

func (i VirtualmachinesPropertiesPublicKeysPublicKeysArgs) ToVirtualmachinesPropertiesPublicKeysPublicKeysOutput() VirtualmachinesPropertiesPublicKeysPublicKeysOutput {
	return i.ToVirtualmachinesPropertiesPublicKeysPublicKeysOutputWithContext(context.Background())
}

func (i VirtualmachinesPropertiesPublicKeysPublicKeysArgs) ToVirtualmachinesPropertiesPublicKeysPublicKeysOutputWithContext(ctx context.Context) VirtualmachinesPropertiesPublicKeysPublicKeysOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualmachinesPropertiesPublicKeysPublicKeysOutput)
}





type VirtualmachinesPropertiesPublicKeysPublicKeysArrayInput interface {
	pulumi.Input

	ToVirtualmachinesPropertiesPublicKeysPublicKeysArrayOutput() VirtualmachinesPropertiesPublicKeysPublicKeysArrayOutput
	ToVirtualmachinesPropertiesPublicKeysPublicKeysArrayOutputWithContext(context.Context) VirtualmachinesPropertiesPublicKeysPublicKeysArrayOutput
}

type VirtualmachinesPropertiesPublicKeysPublicKeysArray []VirtualmachinesPropertiesPublicKeysPublicKeysInput

func (VirtualmachinesPropertiesPublicKeysPublicKeysArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VirtualmachinesPropertiesPublicKeysPublicKeys)(nil)).Elem()
}

func (i VirtualmachinesPropertiesPublicKeysPublicKeysArray) ToVirtualmachinesPropertiesPublicKeysPublicKeysArrayOutput() VirtualmachinesPropertiesPublicKeysPublicKeysArrayOutput {
	return i.ToVirtualmachinesPropertiesPublicKeysPublicKeysArrayOutputWithContext(context.Background())
}

func (i VirtualmachinesPropertiesPublicKeysPublicKeysArray) ToVirtualmachinesPropertiesPublicKeysPublicKeysArrayOutputWithContext(ctx context.Context) VirtualmachinesPropertiesPublicKeysPublicKeysArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualmachinesPropertiesPublicKeysPublicKeysArrayOutput)
}

type VirtualmachinesPropertiesPublicKeysPublicKeysOutput struct{ *pulumi.OutputState }

func (VirtualmachinesPropertiesPublicKeysPublicKeysOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualmachinesPropertiesPublicKeysPublicKeys)(nil)).Elem()
}

func (o VirtualmachinesPropertiesPublicKeysPublicKeysOutput) ToVirtualmachinesPropertiesPublicKeysPublicKeysOutput() VirtualmachinesPropertiesPublicKeysPublicKeysOutput {
	return o
}

func (o VirtualmachinesPropertiesPublicKeysPublicKeysOutput) ToVirtualmachinesPropertiesPublicKeysPublicKeysOutputWithContext(ctx context.Context) VirtualmachinesPropertiesPublicKeysPublicKeysOutput {
	return o
}

func (o VirtualmachinesPropertiesPublicKeysPublicKeysOutput) KeyData() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualmachinesPropertiesPublicKeysPublicKeys) *string { return v.KeyData }).(pulumi.StringPtrOutput)
}

func (o VirtualmachinesPropertiesPublicKeysPublicKeysOutput) Path() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualmachinesPropertiesPublicKeysPublicKeys) *string { return v.Path }).(pulumi.StringPtrOutput)
}

type VirtualmachinesPropertiesPublicKeysPublicKeysArrayOutput struct{ *pulumi.OutputState }

func (VirtualmachinesPropertiesPublicKeysPublicKeysArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VirtualmachinesPropertiesPublicKeysPublicKeys)(nil)).Elem()
}

func (o VirtualmachinesPropertiesPublicKeysPublicKeysArrayOutput) ToVirtualmachinesPropertiesPublicKeysPublicKeysArrayOutput() VirtualmachinesPropertiesPublicKeysPublicKeysArrayOutput {
	return o
}

func (o VirtualmachinesPropertiesPublicKeysPublicKeysArrayOutput) ToVirtualmachinesPropertiesPublicKeysPublicKeysArrayOutputWithContext(ctx context.Context) VirtualmachinesPropertiesPublicKeysPublicKeysArrayOutput {
	return o
}

func (o VirtualmachinesPropertiesPublicKeysPublicKeysArrayOutput) Index(i pulumi.IntInput) VirtualmachinesPropertiesPublicKeysPublicKeysOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VirtualmachinesPropertiesPublicKeysPublicKeys {
		return vs[0].([]VirtualmachinesPropertiesPublicKeysPublicKeys)[vs[1].(int)]
	}).(VirtualmachinesPropertiesPublicKeysPublicKeysOutput)
}

type VirtualmachinesPropertiesResponseDataDisks struct {
	Name *string `pulumi:"name"`
}

type VirtualmachinesPropertiesResponseDataDisksOutput struct{ *pulumi.OutputState }

func (VirtualmachinesPropertiesResponseDataDisksOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualmachinesPropertiesResponseDataDisks)(nil)).Elem()
}

func (o VirtualmachinesPropertiesResponseDataDisksOutput) ToVirtualmachinesPropertiesResponseDataDisksOutput() VirtualmachinesPropertiesResponseDataDisksOutput {
	return o
}

func (o VirtualmachinesPropertiesResponseDataDisksOutput) ToVirtualmachinesPropertiesResponseDataDisksOutputWithContext(ctx context.Context) VirtualmachinesPropertiesResponseDataDisksOutput {
	return o
}

func (o VirtualmachinesPropertiesResponseDataDisksOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualmachinesPropertiesResponseDataDisks) *string { return v.Name }).(pulumi.StringPtrOutput)
}

type VirtualmachinesPropertiesResponseDataDisksArrayOutput struct{ *pulumi.OutputState }

func (VirtualmachinesPropertiesResponseDataDisksArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VirtualmachinesPropertiesResponseDataDisks)(nil)).Elem()
}

func (o VirtualmachinesPropertiesResponseDataDisksArrayOutput) ToVirtualmachinesPropertiesResponseDataDisksArrayOutput() VirtualmachinesPropertiesResponseDataDisksArrayOutput {
	return o
}

func (o VirtualmachinesPropertiesResponseDataDisksArrayOutput) ToVirtualmachinesPropertiesResponseDataDisksArrayOutputWithContext(ctx context.Context) VirtualmachinesPropertiesResponseDataDisksArrayOutput {
	return o
}

func (o VirtualmachinesPropertiesResponseDataDisksArrayOutput) Index(i pulumi.IntInput) VirtualmachinesPropertiesResponseDataDisksOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VirtualmachinesPropertiesResponseDataDisks {
		return vs[0].([]VirtualmachinesPropertiesResponseDataDisks)[vs[1].(int)]
	}).(VirtualmachinesPropertiesResponseDataDisksOutput)
}

type VirtualmachinesPropertiesResponseDynamicMemoryConfig struct {
	MaximumMemoryGB    *float64 `pulumi:"maximumMemoryGB"`
	MinimumMemoryGB    *float64 `pulumi:"minimumMemoryGB"`
	TargetMemoryBuffer *int     `pulumi:"targetMemoryBuffer"`
}

type VirtualmachinesPropertiesResponseDynamicMemoryConfigOutput struct{ *pulumi.OutputState }

func (VirtualmachinesPropertiesResponseDynamicMemoryConfigOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualmachinesPropertiesResponseDynamicMemoryConfig)(nil)).Elem()
}

func (o VirtualmachinesPropertiesResponseDynamicMemoryConfigOutput) ToVirtualmachinesPropertiesResponseDynamicMemoryConfigOutput() VirtualmachinesPropertiesResponseDynamicMemoryConfigOutput {
	return o
}

func (o VirtualmachinesPropertiesResponseDynamicMemoryConfigOutput) ToVirtualmachinesPropertiesResponseDynamicMemoryConfigOutputWithContext(ctx context.Context) VirtualmachinesPropertiesResponseDynamicMemoryConfigOutput {
	return o
}

func (o VirtualmachinesPropertiesResponseDynamicMemoryConfigOutput) MaximumMemoryGB() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v VirtualmachinesPropertiesResponseDynamicMemoryConfig) *float64 { return v.MaximumMemoryGB }).(pulumi.Float64PtrOutput)
}

func (o VirtualmachinesPropertiesResponseDynamicMemoryConfigOutput) MinimumMemoryGB() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v VirtualmachinesPropertiesResponseDynamicMemoryConfig) *float64 { return v.MinimumMemoryGB }).(pulumi.Float64PtrOutput)
}

func (o VirtualmachinesPropertiesResponseDynamicMemoryConfigOutput) TargetMemoryBuffer() pulumi.IntPtrOutput {
	return o.ApplyT(func(v VirtualmachinesPropertiesResponseDynamicMemoryConfig) *int { return v.TargetMemoryBuffer }).(pulumi.IntPtrOutput)
}

type VirtualmachinesPropertiesResponseDynamicMemoryConfigPtrOutput struct{ *pulumi.OutputState }

func (VirtualmachinesPropertiesResponseDynamicMemoryConfigPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualmachinesPropertiesResponseDynamicMemoryConfig)(nil)).Elem()
}

func (o VirtualmachinesPropertiesResponseDynamicMemoryConfigPtrOutput) ToVirtualmachinesPropertiesResponseDynamicMemoryConfigPtrOutput() VirtualmachinesPropertiesResponseDynamicMemoryConfigPtrOutput {
	return o
}

func (o VirtualmachinesPropertiesResponseDynamicMemoryConfigPtrOutput) ToVirtualmachinesPropertiesResponseDynamicMemoryConfigPtrOutputWithContext(ctx context.Context) VirtualmachinesPropertiesResponseDynamicMemoryConfigPtrOutput {
	return o
}

func (o VirtualmachinesPropertiesResponseDynamicMemoryConfigPtrOutput) Elem() VirtualmachinesPropertiesResponseDynamicMemoryConfigOutput {
	return o.ApplyT(func(v *VirtualmachinesPropertiesResponseDynamicMemoryConfig) VirtualmachinesPropertiesResponseDynamicMemoryConfig {
		if v != nil {
			return *v
		}
		var ret VirtualmachinesPropertiesResponseDynamicMemoryConfig
		return ret
	}).(VirtualmachinesPropertiesResponseDynamicMemoryConfigOutput)
}

func (o VirtualmachinesPropertiesResponseDynamicMemoryConfigPtrOutput) MaximumMemoryGB() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *VirtualmachinesPropertiesResponseDynamicMemoryConfig) *float64 {
		if v == nil {
			return nil
		}
		return v.MaximumMemoryGB
	}).(pulumi.Float64PtrOutput)
}

func (o VirtualmachinesPropertiesResponseDynamicMemoryConfigPtrOutput) MinimumMemoryGB() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *VirtualmachinesPropertiesResponseDynamicMemoryConfig) *float64 {
		if v == nil {
			return nil
		}
		return v.MinimumMemoryGB
	}).(pulumi.Float64PtrOutput)
}

func (o VirtualmachinesPropertiesResponseDynamicMemoryConfigPtrOutput) TargetMemoryBuffer() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *VirtualmachinesPropertiesResponseDynamicMemoryConfig) *int {
		if v == nil {
			return nil
		}
		return v.TargetMemoryBuffer
	}).(pulumi.IntPtrOutput)
}

type VirtualmachinesPropertiesResponseHardwareProfile struct {
	DynamicMemoryConfig *VirtualmachinesPropertiesResponseDynamicMemoryConfig `pulumi:"dynamicMemoryConfig"`
	MemoryGB            *int                                                  `pulumi:"memoryGB"`
	Processors          *int                                                  `pulumi:"processors"`
	VmSize              *string                                               `pulumi:"vmSize"`
}

type VirtualmachinesPropertiesResponseHardwareProfileOutput struct{ *pulumi.OutputState }

func (VirtualmachinesPropertiesResponseHardwareProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualmachinesPropertiesResponseHardwareProfile)(nil)).Elem()
}

func (o VirtualmachinesPropertiesResponseHardwareProfileOutput) ToVirtualmachinesPropertiesResponseHardwareProfileOutput() VirtualmachinesPropertiesResponseHardwareProfileOutput {
	return o
}

func (o VirtualmachinesPropertiesResponseHardwareProfileOutput) ToVirtualmachinesPropertiesResponseHardwareProfileOutputWithContext(ctx context.Context) VirtualmachinesPropertiesResponseHardwareProfileOutput {
	return o
}

func (o VirtualmachinesPropertiesResponseHardwareProfileOutput) DynamicMemoryConfig() VirtualmachinesPropertiesResponseDynamicMemoryConfigPtrOutput {
	return o.ApplyT(func(v VirtualmachinesPropertiesResponseHardwareProfile) *VirtualmachinesPropertiesResponseDynamicMemoryConfig {
		return v.DynamicMemoryConfig
	}).(VirtualmachinesPropertiesResponseDynamicMemoryConfigPtrOutput)
}

func (o VirtualmachinesPropertiesResponseHardwareProfileOutput) MemoryGB() pulumi.IntPtrOutput {
	return o.ApplyT(func(v VirtualmachinesPropertiesResponseHardwareProfile) *int { return v.MemoryGB }).(pulumi.IntPtrOutput)
}

func (o VirtualmachinesPropertiesResponseHardwareProfileOutput) Processors() pulumi.IntPtrOutput {
	return o.ApplyT(func(v VirtualmachinesPropertiesResponseHardwareProfile) *int { return v.Processors }).(pulumi.IntPtrOutput)
}

func (o VirtualmachinesPropertiesResponseHardwareProfileOutput) VmSize() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualmachinesPropertiesResponseHardwareProfile) *string { return v.VmSize }).(pulumi.StringPtrOutput)
}

type VirtualmachinesPropertiesResponseHardwareProfilePtrOutput struct{ *pulumi.OutputState }

func (VirtualmachinesPropertiesResponseHardwareProfilePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualmachinesPropertiesResponseHardwareProfile)(nil)).Elem()
}

func (o VirtualmachinesPropertiesResponseHardwareProfilePtrOutput) ToVirtualmachinesPropertiesResponseHardwareProfilePtrOutput() VirtualmachinesPropertiesResponseHardwareProfilePtrOutput {
	return o
}

func (o VirtualmachinesPropertiesResponseHardwareProfilePtrOutput) ToVirtualmachinesPropertiesResponseHardwareProfilePtrOutputWithContext(ctx context.Context) VirtualmachinesPropertiesResponseHardwareProfilePtrOutput {
	return o
}

func (o VirtualmachinesPropertiesResponseHardwareProfilePtrOutput) Elem() VirtualmachinesPropertiesResponseHardwareProfileOutput {
	return o.ApplyT(func(v *VirtualmachinesPropertiesResponseHardwareProfile) VirtualmachinesPropertiesResponseHardwareProfile {
		if v != nil {
			return *v
		}
		var ret VirtualmachinesPropertiesResponseHardwareProfile
		return ret
	}).(VirtualmachinesPropertiesResponseHardwareProfileOutput)
}

func (o VirtualmachinesPropertiesResponseHardwareProfilePtrOutput) DynamicMemoryConfig() VirtualmachinesPropertiesResponseDynamicMemoryConfigPtrOutput {
	return o.ApplyT(func(v *VirtualmachinesPropertiesResponseHardwareProfile) *VirtualmachinesPropertiesResponseDynamicMemoryConfig {
		if v == nil {
			return nil
		}
		return v.DynamicMemoryConfig
	}).(VirtualmachinesPropertiesResponseDynamicMemoryConfigPtrOutput)
}

func (o VirtualmachinesPropertiesResponseHardwareProfilePtrOutput) MemoryGB() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *VirtualmachinesPropertiesResponseHardwareProfile) *int {
		if v == nil {
			return nil
		}
		return v.MemoryGB
	}).(pulumi.IntPtrOutput)
}

func (o VirtualmachinesPropertiesResponseHardwareProfilePtrOutput) Processors() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *VirtualmachinesPropertiesResponseHardwareProfile) *int {
		if v == nil {
			return nil
		}
		return v.Processors
	}).(pulumi.IntPtrOutput)
}

func (o VirtualmachinesPropertiesResponseHardwareProfilePtrOutput) VmSize() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualmachinesPropertiesResponseHardwareProfile) *string {
		if v == nil {
			return nil
		}
		return v.VmSize
	}).(pulumi.StringPtrOutput)
}

type VirtualmachinesPropertiesResponseImageReference struct {
	Name *string `pulumi:"name"`
}

type VirtualmachinesPropertiesResponseImageReferenceOutput struct{ *pulumi.OutputState }

func (VirtualmachinesPropertiesResponseImageReferenceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualmachinesPropertiesResponseImageReference)(nil)).Elem()
}

func (o VirtualmachinesPropertiesResponseImageReferenceOutput) ToVirtualmachinesPropertiesResponseImageReferenceOutput() VirtualmachinesPropertiesResponseImageReferenceOutput {
	return o
}

func (o VirtualmachinesPropertiesResponseImageReferenceOutput) ToVirtualmachinesPropertiesResponseImageReferenceOutputWithContext(ctx context.Context) VirtualmachinesPropertiesResponseImageReferenceOutput {
	return o
}

func (o VirtualmachinesPropertiesResponseImageReferenceOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualmachinesPropertiesResponseImageReference) *string { return v.Name }).(pulumi.StringPtrOutput)
}

type VirtualmachinesPropertiesResponseImageReferencePtrOutput struct{ *pulumi.OutputState }

func (VirtualmachinesPropertiesResponseImageReferencePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualmachinesPropertiesResponseImageReference)(nil)).Elem()
}

func (o VirtualmachinesPropertiesResponseImageReferencePtrOutput) ToVirtualmachinesPropertiesResponseImageReferencePtrOutput() VirtualmachinesPropertiesResponseImageReferencePtrOutput {
	return o
}

func (o VirtualmachinesPropertiesResponseImageReferencePtrOutput) ToVirtualmachinesPropertiesResponseImageReferencePtrOutputWithContext(ctx context.Context) VirtualmachinesPropertiesResponseImageReferencePtrOutput {
	return o
}

func (o VirtualmachinesPropertiesResponseImageReferencePtrOutput) Elem() VirtualmachinesPropertiesResponseImageReferenceOutput {
	return o.ApplyT(func(v *VirtualmachinesPropertiesResponseImageReference) VirtualmachinesPropertiesResponseImageReference {
		if v != nil {
			return *v
		}
		var ret VirtualmachinesPropertiesResponseImageReference
		return ret
	}).(VirtualmachinesPropertiesResponseImageReferenceOutput)
}

func (o VirtualmachinesPropertiesResponseImageReferencePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualmachinesPropertiesResponseImageReference) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

type VirtualmachinesPropertiesResponseLinuxConfiguration struct {
	DisablePasswordAuthentication *bool                                 `pulumi:"disablePasswordAuthentication"`
	ProvisionVMAgent              *bool                                 `pulumi:"provisionVMAgent"`
	Ssh                           *VirtualmachinesPropertiesResponseSsh `pulumi:"ssh"`
}

type VirtualmachinesPropertiesResponseLinuxConfigurationOutput struct{ *pulumi.OutputState }

func (VirtualmachinesPropertiesResponseLinuxConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualmachinesPropertiesResponseLinuxConfiguration)(nil)).Elem()
}

func (o VirtualmachinesPropertiesResponseLinuxConfigurationOutput) ToVirtualmachinesPropertiesResponseLinuxConfigurationOutput() VirtualmachinesPropertiesResponseLinuxConfigurationOutput {
	return o
}

func (o VirtualmachinesPropertiesResponseLinuxConfigurationOutput) ToVirtualmachinesPropertiesResponseLinuxConfigurationOutputWithContext(ctx context.Context) VirtualmachinesPropertiesResponseLinuxConfigurationOutput {
	return o
}

func (o VirtualmachinesPropertiesResponseLinuxConfigurationOutput) DisablePasswordAuthentication() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VirtualmachinesPropertiesResponseLinuxConfiguration) *bool {
		return v.DisablePasswordAuthentication
	}).(pulumi.BoolPtrOutput)
}

func (o VirtualmachinesPropertiesResponseLinuxConfigurationOutput) ProvisionVMAgent() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VirtualmachinesPropertiesResponseLinuxConfiguration) *bool { return v.ProvisionVMAgent }).(pulumi.BoolPtrOutput)
}

func (o VirtualmachinesPropertiesResponseLinuxConfigurationOutput) Ssh() VirtualmachinesPropertiesResponseSshPtrOutput {
	return o.ApplyT(func(v VirtualmachinesPropertiesResponseLinuxConfiguration) *VirtualmachinesPropertiesResponseSsh {
		return v.Ssh
	}).(VirtualmachinesPropertiesResponseSshPtrOutput)
}

type VirtualmachinesPropertiesResponseLinuxConfigurationPtrOutput struct{ *pulumi.OutputState }

func (VirtualmachinesPropertiesResponseLinuxConfigurationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualmachinesPropertiesResponseLinuxConfiguration)(nil)).Elem()
}

func (o VirtualmachinesPropertiesResponseLinuxConfigurationPtrOutput) ToVirtualmachinesPropertiesResponseLinuxConfigurationPtrOutput() VirtualmachinesPropertiesResponseLinuxConfigurationPtrOutput {
	return o
}

func (o VirtualmachinesPropertiesResponseLinuxConfigurationPtrOutput) ToVirtualmachinesPropertiesResponseLinuxConfigurationPtrOutputWithContext(ctx context.Context) VirtualmachinesPropertiesResponseLinuxConfigurationPtrOutput {
	return o
}

func (o VirtualmachinesPropertiesResponseLinuxConfigurationPtrOutput) Elem() VirtualmachinesPropertiesResponseLinuxConfigurationOutput {
	return o.ApplyT(func(v *VirtualmachinesPropertiesResponseLinuxConfiguration) VirtualmachinesPropertiesResponseLinuxConfiguration {
		if v != nil {
			return *v
		}
		var ret VirtualmachinesPropertiesResponseLinuxConfiguration
		return ret
	}).(VirtualmachinesPropertiesResponseLinuxConfigurationOutput)
}

func (o VirtualmachinesPropertiesResponseLinuxConfigurationPtrOutput) DisablePasswordAuthentication() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *VirtualmachinesPropertiesResponseLinuxConfiguration) *bool {
		if v == nil {
			return nil
		}
		return v.DisablePasswordAuthentication
	}).(pulumi.BoolPtrOutput)
}

func (o VirtualmachinesPropertiesResponseLinuxConfigurationPtrOutput) ProvisionVMAgent() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *VirtualmachinesPropertiesResponseLinuxConfiguration) *bool {
		if v == nil {
			return nil
		}
		return v.ProvisionVMAgent
	}).(pulumi.BoolPtrOutput)
}

func (o VirtualmachinesPropertiesResponseLinuxConfigurationPtrOutput) Ssh() VirtualmachinesPropertiesResponseSshPtrOutput {
	return o.ApplyT(func(v *VirtualmachinesPropertiesResponseLinuxConfiguration) *VirtualmachinesPropertiesResponseSsh {
		if v == nil {
			return nil
		}
		return v.Ssh
	}).(VirtualmachinesPropertiesResponseSshPtrOutput)
}

type VirtualmachinesPropertiesResponseNetworkInterfaces struct {
	Id *string `pulumi:"id"`
}

type VirtualmachinesPropertiesResponseNetworkInterfacesOutput struct{ *pulumi.OutputState }

func (VirtualmachinesPropertiesResponseNetworkInterfacesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualmachinesPropertiesResponseNetworkInterfaces)(nil)).Elem()
}

func (o VirtualmachinesPropertiesResponseNetworkInterfacesOutput) ToVirtualmachinesPropertiesResponseNetworkInterfacesOutput() VirtualmachinesPropertiesResponseNetworkInterfacesOutput {
	return o
}

func (o VirtualmachinesPropertiesResponseNetworkInterfacesOutput) ToVirtualmachinesPropertiesResponseNetworkInterfacesOutputWithContext(ctx context.Context) VirtualmachinesPropertiesResponseNetworkInterfacesOutput {
	return o
}

func (o VirtualmachinesPropertiesResponseNetworkInterfacesOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualmachinesPropertiesResponseNetworkInterfaces) *string { return v.Id }).(pulumi.StringPtrOutput)
}

type VirtualmachinesPropertiesResponseNetworkInterfacesArrayOutput struct{ *pulumi.OutputState }

func (VirtualmachinesPropertiesResponseNetworkInterfacesArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VirtualmachinesPropertiesResponseNetworkInterfaces)(nil)).Elem()
}

func (o VirtualmachinesPropertiesResponseNetworkInterfacesArrayOutput) ToVirtualmachinesPropertiesResponseNetworkInterfacesArrayOutput() VirtualmachinesPropertiesResponseNetworkInterfacesArrayOutput {
	return o
}

func (o VirtualmachinesPropertiesResponseNetworkInterfacesArrayOutput) ToVirtualmachinesPropertiesResponseNetworkInterfacesArrayOutputWithContext(ctx context.Context) VirtualmachinesPropertiesResponseNetworkInterfacesArrayOutput {
	return o
}

func (o VirtualmachinesPropertiesResponseNetworkInterfacesArrayOutput) Index(i pulumi.IntInput) VirtualmachinesPropertiesResponseNetworkInterfacesOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VirtualmachinesPropertiesResponseNetworkInterfaces {
		return vs[0].([]VirtualmachinesPropertiesResponseNetworkInterfaces)[vs[1].(int)]
	}).(VirtualmachinesPropertiesResponseNetworkInterfacesOutput)
}

type VirtualmachinesPropertiesResponseNetworkProfile struct {
	NetworkInterfaces []VirtualmachinesPropertiesResponseNetworkInterfaces `pulumi:"networkInterfaces"`
}

type VirtualmachinesPropertiesResponseNetworkProfileOutput struct{ *pulumi.OutputState }

func (VirtualmachinesPropertiesResponseNetworkProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualmachinesPropertiesResponseNetworkProfile)(nil)).Elem()
}

func (o VirtualmachinesPropertiesResponseNetworkProfileOutput) ToVirtualmachinesPropertiesResponseNetworkProfileOutput() VirtualmachinesPropertiesResponseNetworkProfileOutput {
	return o
}

func (o VirtualmachinesPropertiesResponseNetworkProfileOutput) ToVirtualmachinesPropertiesResponseNetworkProfileOutputWithContext(ctx context.Context) VirtualmachinesPropertiesResponseNetworkProfileOutput {
	return o
}

func (o VirtualmachinesPropertiesResponseNetworkProfileOutput) NetworkInterfaces() VirtualmachinesPropertiesResponseNetworkInterfacesArrayOutput {
	return o.ApplyT(func(v VirtualmachinesPropertiesResponseNetworkProfile) []VirtualmachinesPropertiesResponseNetworkInterfaces {
		return v.NetworkInterfaces
	}).(VirtualmachinesPropertiesResponseNetworkInterfacesArrayOutput)
}

type VirtualmachinesPropertiesResponseNetworkProfilePtrOutput struct{ *pulumi.OutputState }

func (VirtualmachinesPropertiesResponseNetworkProfilePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualmachinesPropertiesResponseNetworkProfile)(nil)).Elem()
}

func (o VirtualmachinesPropertiesResponseNetworkProfilePtrOutput) ToVirtualmachinesPropertiesResponseNetworkProfilePtrOutput() VirtualmachinesPropertiesResponseNetworkProfilePtrOutput {
	return o
}

func (o VirtualmachinesPropertiesResponseNetworkProfilePtrOutput) ToVirtualmachinesPropertiesResponseNetworkProfilePtrOutputWithContext(ctx context.Context) VirtualmachinesPropertiesResponseNetworkProfilePtrOutput {
	return o
}

func (o VirtualmachinesPropertiesResponseNetworkProfilePtrOutput) Elem() VirtualmachinesPropertiesResponseNetworkProfileOutput {
	return o.ApplyT(func(v *VirtualmachinesPropertiesResponseNetworkProfile) VirtualmachinesPropertiesResponseNetworkProfile {
		if v != nil {
			return *v
		}
		var ret VirtualmachinesPropertiesResponseNetworkProfile
		return ret
	}).(VirtualmachinesPropertiesResponseNetworkProfileOutput)
}

func (o VirtualmachinesPropertiesResponseNetworkProfilePtrOutput) NetworkInterfaces() VirtualmachinesPropertiesResponseNetworkInterfacesArrayOutput {
	return o.ApplyT(func(v *VirtualmachinesPropertiesResponseNetworkProfile) []VirtualmachinesPropertiesResponseNetworkInterfaces {
		if v == nil {
			return nil
		}
		return v.NetworkInterfaces
	}).(VirtualmachinesPropertiesResponseNetworkInterfacesArrayOutput)
}

type VirtualmachinesPropertiesResponseOsDisk struct {
	Id *string `pulumi:"id"`
}

type VirtualmachinesPropertiesResponseOsDiskOutput struct{ *pulumi.OutputState }

func (VirtualmachinesPropertiesResponseOsDiskOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualmachinesPropertiesResponseOsDisk)(nil)).Elem()
}

func (o VirtualmachinesPropertiesResponseOsDiskOutput) ToVirtualmachinesPropertiesResponseOsDiskOutput() VirtualmachinesPropertiesResponseOsDiskOutput {
	return o
}

func (o VirtualmachinesPropertiesResponseOsDiskOutput) ToVirtualmachinesPropertiesResponseOsDiskOutputWithContext(ctx context.Context) VirtualmachinesPropertiesResponseOsDiskOutput {
	return o
}

func (o VirtualmachinesPropertiesResponseOsDiskOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualmachinesPropertiesResponseOsDisk) *string { return v.Id }).(pulumi.StringPtrOutput)
}

type VirtualmachinesPropertiesResponseOsDiskPtrOutput struct{ *pulumi.OutputState }

func (VirtualmachinesPropertiesResponseOsDiskPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualmachinesPropertiesResponseOsDisk)(nil)).Elem()
}

func (o VirtualmachinesPropertiesResponseOsDiskPtrOutput) ToVirtualmachinesPropertiesResponseOsDiskPtrOutput() VirtualmachinesPropertiesResponseOsDiskPtrOutput {
	return o
}

func (o VirtualmachinesPropertiesResponseOsDiskPtrOutput) ToVirtualmachinesPropertiesResponseOsDiskPtrOutputWithContext(ctx context.Context) VirtualmachinesPropertiesResponseOsDiskPtrOutput {
	return o
}

func (o VirtualmachinesPropertiesResponseOsDiskPtrOutput) Elem() VirtualmachinesPropertiesResponseOsDiskOutput {
	return o.ApplyT(func(v *VirtualmachinesPropertiesResponseOsDisk) VirtualmachinesPropertiesResponseOsDisk {
		if v != nil {
			return *v
		}
		var ret VirtualmachinesPropertiesResponseOsDisk
		return ret
	}).(VirtualmachinesPropertiesResponseOsDiskOutput)
}

func (o VirtualmachinesPropertiesResponseOsDiskPtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualmachinesPropertiesResponseOsDisk) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

type VirtualmachinesPropertiesResponseOsProfile struct {
	AdminUsername        *string                                                `pulumi:"adminUsername"`
	ComputerName         *string                                                `pulumi:"computerName"`
	LinuxConfiguration   *VirtualmachinesPropertiesResponseLinuxConfiguration   `pulumi:"linuxConfiguration"`
	OsType               *string                                                `pulumi:"osType"`
	WindowsConfiguration *VirtualmachinesPropertiesResponseWindowsConfiguration `pulumi:"windowsConfiguration"`
}

type VirtualmachinesPropertiesResponseOsProfileOutput struct{ *pulumi.OutputState }

func (VirtualmachinesPropertiesResponseOsProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualmachinesPropertiesResponseOsProfile)(nil)).Elem()
}

func (o VirtualmachinesPropertiesResponseOsProfileOutput) ToVirtualmachinesPropertiesResponseOsProfileOutput() VirtualmachinesPropertiesResponseOsProfileOutput {
	return o
}

func (o VirtualmachinesPropertiesResponseOsProfileOutput) ToVirtualmachinesPropertiesResponseOsProfileOutputWithContext(ctx context.Context) VirtualmachinesPropertiesResponseOsProfileOutput {
	return o
}

func (o VirtualmachinesPropertiesResponseOsProfileOutput) AdminUsername() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualmachinesPropertiesResponseOsProfile) *string { return v.AdminUsername }).(pulumi.StringPtrOutput)
}

func (o VirtualmachinesPropertiesResponseOsProfileOutput) ComputerName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualmachinesPropertiesResponseOsProfile) *string { return v.ComputerName }).(pulumi.StringPtrOutput)
}

func (o VirtualmachinesPropertiesResponseOsProfileOutput) LinuxConfiguration() VirtualmachinesPropertiesResponseLinuxConfigurationPtrOutput {
	return o.ApplyT(func(v VirtualmachinesPropertiesResponseOsProfile) *VirtualmachinesPropertiesResponseLinuxConfiguration {
		return v.LinuxConfiguration
	}).(VirtualmachinesPropertiesResponseLinuxConfigurationPtrOutput)
}

func (o VirtualmachinesPropertiesResponseOsProfileOutput) OsType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualmachinesPropertiesResponseOsProfile) *string { return v.OsType }).(pulumi.StringPtrOutput)
}

func (o VirtualmachinesPropertiesResponseOsProfileOutput) WindowsConfiguration() VirtualmachinesPropertiesResponseWindowsConfigurationPtrOutput {
	return o.ApplyT(func(v VirtualmachinesPropertiesResponseOsProfile) *VirtualmachinesPropertiesResponseWindowsConfiguration {
		return v.WindowsConfiguration
	}).(VirtualmachinesPropertiesResponseWindowsConfigurationPtrOutput)
}

type VirtualmachinesPropertiesResponseOsProfilePtrOutput struct{ *pulumi.OutputState }

func (VirtualmachinesPropertiesResponseOsProfilePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualmachinesPropertiesResponseOsProfile)(nil)).Elem()
}

func (o VirtualmachinesPropertiesResponseOsProfilePtrOutput) ToVirtualmachinesPropertiesResponseOsProfilePtrOutput() VirtualmachinesPropertiesResponseOsProfilePtrOutput {
	return o
}

func (o VirtualmachinesPropertiesResponseOsProfilePtrOutput) ToVirtualmachinesPropertiesResponseOsProfilePtrOutputWithContext(ctx context.Context) VirtualmachinesPropertiesResponseOsProfilePtrOutput {
	return o
}

func (o VirtualmachinesPropertiesResponseOsProfilePtrOutput) Elem() VirtualmachinesPropertiesResponseOsProfileOutput {
	return o.ApplyT(func(v *VirtualmachinesPropertiesResponseOsProfile) VirtualmachinesPropertiesResponseOsProfile {
		if v != nil {
			return *v
		}
		var ret VirtualmachinesPropertiesResponseOsProfile
		return ret
	}).(VirtualmachinesPropertiesResponseOsProfileOutput)
}

func (o VirtualmachinesPropertiesResponseOsProfilePtrOutput) AdminUsername() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualmachinesPropertiesResponseOsProfile) *string {
		if v == nil {
			return nil
		}
		return v.AdminUsername
	}).(pulumi.StringPtrOutput)
}

func (o VirtualmachinesPropertiesResponseOsProfilePtrOutput) ComputerName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualmachinesPropertiesResponseOsProfile) *string {
		if v == nil {
			return nil
		}
		return v.ComputerName
	}).(pulumi.StringPtrOutput)
}

func (o VirtualmachinesPropertiesResponseOsProfilePtrOutput) LinuxConfiguration() VirtualmachinesPropertiesResponseLinuxConfigurationPtrOutput {
	return o.ApplyT(func(v *VirtualmachinesPropertiesResponseOsProfile) *VirtualmachinesPropertiesResponseLinuxConfiguration {
		if v == nil {
			return nil
		}
		return v.LinuxConfiguration
	}).(VirtualmachinesPropertiesResponseLinuxConfigurationPtrOutput)
}

func (o VirtualmachinesPropertiesResponseOsProfilePtrOutput) OsType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualmachinesPropertiesResponseOsProfile) *string {
		if v == nil {
			return nil
		}
		return v.OsType
	}).(pulumi.StringPtrOutput)
}

func (o VirtualmachinesPropertiesResponseOsProfilePtrOutput) WindowsConfiguration() VirtualmachinesPropertiesResponseWindowsConfigurationPtrOutput {
	return o.ApplyT(func(v *VirtualmachinesPropertiesResponseOsProfile) *VirtualmachinesPropertiesResponseWindowsConfiguration {
		if v == nil {
			return nil
		}
		return v.WindowsConfiguration
	}).(VirtualmachinesPropertiesResponseWindowsConfigurationPtrOutput)
}

type VirtualmachinesPropertiesResponsePublicKeys struct {
	KeyData *string `pulumi:"keyData"`
	Path    *string `pulumi:"path"`
}

type VirtualmachinesPropertiesResponsePublicKeysOutput struct{ *pulumi.OutputState }

func (VirtualmachinesPropertiesResponsePublicKeysOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualmachinesPropertiesResponsePublicKeys)(nil)).Elem()
}

func (o VirtualmachinesPropertiesResponsePublicKeysOutput) ToVirtualmachinesPropertiesResponsePublicKeysOutput() VirtualmachinesPropertiesResponsePublicKeysOutput {
	return o
}

func (o VirtualmachinesPropertiesResponsePublicKeysOutput) ToVirtualmachinesPropertiesResponsePublicKeysOutputWithContext(ctx context.Context) VirtualmachinesPropertiesResponsePublicKeysOutput {
	return o
}

func (o VirtualmachinesPropertiesResponsePublicKeysOutput) KeyData() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualmachinesPropertiesResponsePublicKeys) *string { return v.KeyData }).(pulumi.StringPtrOutput)
}

func (o VirtualmachinesPropertiesResponsePublicKeysOutput) Path() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualmachinesPropertiesResponsePublicKeys) *string { return v.Path }).(pulumi.StringPtrOutput)
}

type VirtualmachinesPropertiesResponsePublicKeysArrayOutput struct{ *pulumi.OutputState }

func (VirtualmachinesPropertiesResponsePublicKeysArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VirtualmachinesPropertiesResponsePublicKeys)(nil)).Elem()
}

func (o VirtualmachinesPropertiesResponsePublicKeysArrayOutput) ToVirtualmachinesPropertiesResponsePublicKeysArrayOutput() VirtualmachinesPropertiesResponsePublicKeysArrayOutput {
	return o
}

func (o VirtualmachinesPropertiesResponsePublicKeysArrayOutput) ToVirtualmachinesPropertiesResponsePublicKeysArrayOutputWithContext(ctx context.Context) VirtualmachinesPropertiesResponsePublicKeysArrayOutput {
	return o
}

func (o VirtualmachinesPropertiesResponsePublicKeysArrayOutput) Index(i pulumi.IntInput) VirtualmachinesPropertiesResponsePublicKeysOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VirtualmachinesPropertiesResponsePublicKeys {
		return vs[0].([]VirtualmachinesPropertiesResponsePublicKeys)[vs[1].(int)]
	}).(VirtualmachinesPropertiesResponsePublicKeysOutput)
}

type VirtualmachinesPropertiesResponsePublicKeysPublicKeys struct {
	KeyData *string `pulumi:"keyData"`
	Path    *string `pulumi:"path"`
}

type VirtualmachinesPropertiesResponsePublicKeysPublicKeysOutput struct{ *pulumi.OutputState }

func (VirtualmachinesPropertiesResponsePublicKeysPublicKeysOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualmachinesPropertiesResponsePublicKeysPublicKeys)(nil)).Elem()
}

func (o VirtualmachinesPropertiesResponsePublicKeysPublicKeysOutput) ToVirtualmachinesPropertiesResponsePublicKeysPublicKeysOutput() VirtualmachinesPropertiesResponsePublicKeysPublicKeysOutput {
	return o
}

func (o VirtualmachinesPropertiesResponsePublicKeysPublicKeysOutput) ToVirtualmachinesPropertiesResponsePublicKeysPublicKeysOutputWithContext(ctx context.Context) VirtualmachinesPropertiesResponsePublicKeysPublicKeysOutput {
	return o
}

func (o VirtualmachinesPropertiesResponsePublicKeysPublicKeysOutput) KeyData() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualmachinesPropertiesResponsePublicKeysPublicKeys) *string { return v.KeyData }).(pulumi.StringPtrOutput)
}

func (o VirtualmachinesPropertiesResponsePublicKeysPublicKeysOutput) Path() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualmachinesPropertiesResponsePublicKeysPublicKeys) *string { return v.Path }).(pulumi.StringPtrOutput)
}

type VirtualmachinesPropertiesResponsePublicKeysPublicKeysArrayOutput struct{ *pulumi.OutputState }

func (VirtualmachinesPropertiesResponsePublicKeysPublicKeysArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VirtualmachinesPropertiesResponsePublicKeysPublicKeys)(nil)).Elem()
}

func (o VirtualmachinesPropertiesResponsePublicKeysPublicKeysArrayOutput) ToVirtualmachinesPropertiesResponsePublicKeysPublicKeysArrayOutput() VirtualmachinesPropertiesResponsePublicKeysPublicKeysArrayOutput {
	return o
}

func (o VirtualmachinesPropertiesResponsePublicKeysPublicKeysArrayOutput) ToVirtualmachinesPropertiesResponsePublicKeysPublicKeysArrayOutputWithContext(ctx context.Context) VirtualmachinesPropertiesResponsePublicKeysPublicKeysArrayOutput {
	return o
}

func (o VirtualmachinesPropertiesResponsePublicKeysPublicKeysArrayOutput) Index(i pulumi.IntInput) VirtualmachinesPropertiesResponsePublicKeysPublicKeysOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VirtualmachinesPropertiesResponsePublicKeysPublicKeys {
		return vs[0].([]VirtualmachinesPropertiesResponsePublicKeysPublicKeys)[vs[1].(int)]
	}).(VirtualmachinesPropertiesResponsePublicKeysPublicKeysOutput)
}

type VirtualmachinesPropertiesResponseSecurityProfile struct {
	EnableTPM    *bool                                          `pulumi:"enableTPM"`
	UefiSettings *VirtualmachinesPropertiesResponseUefiSettings `pulumi:"uefiSettings"`
}

type VirtualmachinesPropertiesResponseSecurityProfileOutput struct{ *pulumi.OutputState }

func (VirtualmachinesPropertiesResponseSecurityProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualmachinesPropertiesResponseSecurityProfile)(nil)).Elem()
}

func (o VirtualmachinesPropertiesResponseSecurityProfileOutput) ToVirtualmachinesPropertiesResponseSecurityProfileOutput() VirtualmachinesPropertiesResponseSecurityProfileOutput {
	return o
}

func (o VirtualmachinesPropertiesResponseSecurityProfileOutput) ToVirtualmachinesPropertiesResponseSecurityProfileOutputWithContext(ctx context.Context) VirtualmachinesPropertiesResponseSecurityProfileOutput {
	return o
}

func (o VirtualmachinesPropertiesResponseSecurityProfileOutput) EnableTPM() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VirtualmachinesPropertiesResponseSecurityProfile) *bool { return v.EnableTPM }).(pulumi.BoolPtrOutput)
}

func (o VirtualmachinesPropertiesResponseSecurityProfileOutput) UefiSettings() VirtualmachinesPropertiesResponseUefiSettingsPtrOutput {
	return o.ApplyT(func(v VirtualmachinesPropertiesResponseSecurityProfile) *VirtualmachinesPropertiesResponseUefiSettings {
		return v.UefiSettings
	}).(VirtualmachinesPropertiesResponseUefiSettingsPtrOutput)
}

type VirtualmachinesPropertiesResponseSecurityProfilePtrOutput struct{ *pulumi.OutputState }

func (VirtualmachinesPropertiesResponseSecurityProfilePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualmachinesPropertiesResponseSecurityProfile)(nil)).Elem()
}

func (o VirtualmachinesPropertiesResponseSecurityProfilePtrOutput) ToVirtualmachinesPropertiesResponseSecurityProfilePtrOutput() VirtualmachinesPropertiesResponseSecurityProfilePtrOutput {
	return o
}

func (o VirtualmachinesPropertiesResponseSecurityProfilePtrOutput) ToVirtualmachinesPropertiesResponseSecurityProfilePtrOutputWithContext(ctx context.Context) VirtualmachinesPropertiesResponseSecurityProfilePtrOutput {
	return o
}

func (o VirtualmachinesPropertiesResponseSecurityProfilePtrOutput) Elem() VirtualmachinesPropertiesResponseSecurityProfileOutput {
	return o.ApplyT(func(v *VirtualmachinesPropertiesResponseSecurityProfile) VirtualmachinesPropertiesResponseSecurityProfile {
		if v != nil {
			return *v
		}
		var ret VirtualmachinesPropertiesResponseSecurityProfile
		return ret
	}).(VirtualmachinesPropertiesResponseSecurityProfileOutput)
}

func (o VirtualmachinesPropertiesResponseSecurityProfilePtrOutput) EnableTPM() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *VirtualmachinesPropertiesResponseSecurityProfile) *bool {
		if v == nil {
			return nil
		}
		return v.EnableTPM
	}).(pulumi.BoolPtrOutput)
}

func (o VirtualmachinesPropertiesResponseSecurityProfilePtrOutput) UefiSettings() VirtualmachinesPropertiesResponseUefiSettingsPtrOutput {
	return o.ApplyT(func(v *VirtualmachinesPropertiesResponseSecurityProfile) *VirtualmachinesPropertiesResponseUefiSettings {
		if v == nil {
			return nil
		}
		return v.UefiSettings
	}).(VirtualmachinesPropertiesResponseUefiSettingsPtrOutput)
}

type VirtualmachinesPropertiesResponseSsh struct {
	PublicKeys []VirtualmachinesPropertiesResponsePublicKeys `pulumi:"publicKeys"`
}

type VirtualmachinesPropertiesResponseSshOutput struct{ *pulumi.OutputState }

func (VirtualmachinesPropertiesResponseSshOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualmachinesPropertiesResponseSsh)(nil)).Elem()
}

func (o VirtualmachinesPropertiesResponseSshOutput) ToVirtualmachinesPropertiesResponseSshOutput() VirtualmachinesPropertiesResponseSshOutput {
	return o
}

func (o VirtualmachinesPropertiesResponseSshOutput) ToVirtualmachinesPropertiesResponseSshOutputWithContext(ctx context.Context) VirtualmachinesPropertiesResponseSshOutput {
	return o
}

func (o VirtualmachinesPropertiesResponseSshOutput) PublicKeys() VirtualmachinesPropertiesResponsePublicKeysArrayOutput {
	return o.ApplyT(func(v VirtualmachinesPropertiesResponseSsh) []VirtualmachinesPropertiesResponsePublicKeys {
		return v.PublicKeys
	}).(VirtualmachinesPropertiesResponsePublicKeysArrayOutput)
}

type VirtualmachinesPropertiesResponseSshPtrOutput struct{ *pulumi.OutputState }

func (VirtualmachinesPropertiesResponseSshPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualmachinesPropertiesResponseSsh)(nil)).Elem()
}

func (o VirtualmachinesPropertiesResponseSshPtrOutput) ToVirtualmachinesPropertiesResponseSshPtrOutput() VirtualmachinesPropertiesResponseSshPtrOutput {
	return o
}

func (o VirtualmachinesPropertiesResponseSshPtrOutput) ToVirtualmachinesPropertiesResponseSshPtrOutputWithContext(ctx context.Context) VirtualmachinesPropertiesResponseSshPtrOutput {
	return o
}

func (o VirtualmachinesPropertiesResponseSshPtrOutput) Elem() VirtualmachinesPropertiesResponseSshOutput {
	return o.ApplyT(func(v *VirtualmachinesPropertiesResponseSsh) VirtualmachinesPropertiesResponseSsh {
		if v != nil {
			return *v
		}
		var ret VirtualmachinesPropertiesResponseSsh
		return ret
	}).(VirtualmachinesPropertiesResponseSshOutput)
}

func (o VirtualmachinesPropertiesResponseSshPtrOutput) PublicKeys() VirtualmachinesPropertiesResponsePublicKeysArrayOutput {
	return o.ApplyT(func(v *VirtualmachinesPropertiesResponseSsh) []VirtualmachinesPropertiesResponsePublicKeys {
		if v == nil {
			return nil
		}
		return v.PublicKeys
	}).(VirtualmachinesPropertiesResponsePublicKeysArrayOutput)
}

type VirtualmachinesPropertiesResponseSshSsh struct {
	PublicKeys []VirtualmachinesPropertiesResponsePublicKeysPublicKeys `pulumi:"publicKeys"`
}

type VirtualmachinesPropertiesResponseSshSshOutput struct{ *pulumi.OutputState }

func (VirtualmachinesPropertiesResponseSshSshOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualmachinesPropertiesResponseSshSsh)(nil)).Elem()
}

func (o VirtualmachinesPropertiesResponseSshSshOutput) ToVirtualmachinesPropertiesResponseSshSshOutput() VirtualmachinesPropertiesResponseSshSshOutput {
	return o
}

func (o VirtualmachinesPropertiesResponseSshSshOutput) ToVirtualmachinesPropertiesResponseSshSshOutputWithContext(ctx context.Context) VirtualmachinesPropertiesResponseSshSshOutput {
	return o
}

func (o VirtualmachinesPropertiesResponseSshSshOutput) PublicKeys() VirtualmachinesPropertiesResponsePublicKeysPublicKeysArrayOutput {
	return o.ApplyT(func(v VirtualmachinesPropertiesResponseSshSsh) []VirtualmachinesPropertiesResponsePublicKeysPublicKeys {
		return v.PublicKeys
	}).(VirtualmachinesPropertiesResponsePublicKeysPublicKeysArrayOutput)
}

type VirtualmachinesPropertiesResponseSshSshPtrOutput struct{ *pulumi.OutputState }

func (VirtualmachinesPropertiesResponseSshSshPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualmachinesPropertiesResponseSshSsh)(nil)).Elem()
}

func (o VirtualmachinesPropertiesResponseSshSshPtrOutput) ToVirtualmachinesPropertiesResponseSshSshPtrOutput() VirtualmachinesPropertiesResponseSshSshPtrOutput {
	return o
}

func (o VirtualmachinesPropertiesResponseSshSshPtrOutput) ToVirtualmachinesPropertiesResponseSshSshPtrOutputWithContext(ctx context.Context) VirtualmachinesPropertiesResponseSshSshPtrOutput {
	return o
}

func (o VirtualmachinesPropertiesResponseSshSshPtrOutput) Elem() VirtualmachinesPropertiesResponseSshSshOutput {
	return o.ApplyT(func(v *VirtualmachinesPropertiesResponseSshSsh) VirtualmachinesPropertiesResponseSshSsh {
		if v != nil {
			return *v
		}
		var ret VirtualmachinesPropertiesResponseSshSsh
		return ret
	}).(VirtualmachinesPropertiesResponseSshSshOutput)
}

func (o VirtualmachinesPropertiesResponseSshSshPtrOutput) PublicKeys() VirtualmachinesPropertiesResponsePublicKeysPublicKeysArrayOutput {
	return o.ApplyT(func(v *VirtualmachinesPropertiesResponseSshSsh) []VirtualmachinesPropertiesResponsePublicKeysPublicKeys {
		if v == nil {
			return nil
		}
		return v.PublicKeys
	}).(VirtualmachinesPropertiesResponsePublicKeysPublicKeysArrayOutput)
}

type VirtualmachinesPropertiesResponseStorageProfile struct {
	DataDisks             []VirtualmachinesPropertiesResponseDataDisks     `pulumi:"dataDisks"`
	ImageReference        *VirtualmachinesPropertiesResponseImageReference `pulumi:"imageReference"`
	OsDisk                *VirtualmachinesPropertiesResponseOsDisk         `pulumi:"osDisk"`
	VmConfigContainerName *string                                          `pulumi:"vmConfigContainerName"`
}

type VirtualmachinesPropertiesResponseStorageProfileOutput struct{ *pulumi.OutputState }

func (VirtualmachinesPropertiesResponseStorageProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualmachinesPropertiesResponseStorageProfile)(nil)).Elem()
}

func (o VirtualmachinesPropertiesResponseStorageProfileOutput) ToVirtualmachinesPropertiesResponseStorageProfileOutput() VirtualmachinesPropertiesResponseStorageProfileOutput {
	return o
}

func (o VirtualmachinesPropertiesResponseStorageProfileOutput) ToVirtualmachinesPropertiesResponseStorageProfileOutputWithContext(ctx context.Context) VirtualmachinesPropertiesResponseStorageProfileOutput {
	return o
}

func (o VirtualmachinesPropertiesResponseStorageProfileOutput) DataDisks() VirtualmachinesPropertiesResponseDataDisksArrayOutput {
	return o.ApplyT(func(v VirtualmachinesPropertiesResponseStorageProfile) []VirtualmachinesPropertiesResponseDataDisks {
		return v.DataDisks
	}).(VirtualmachinesPropertiesResponseDataDisksArrayOutput)
}

func (o VirtualmachinesPropertiesResponseStorageProfileOutput) ImageReference() VirtualmachinesPropertiesResponseImageReferencePtrOutput {
	return o.ApplyT(func(v VirtualmachinesPropertiesResponseStorageProfile) *VirtualmachinesPropertiesResponseImageReference {
		return v.ImageReference
	}).(VirtualmachinesPropertiesResponseImageReferencePtrOutput)
}

func (o VirtualmachinesPropertiesResponseStorageProfileOutput) OsDisk() VirtualmachinesPropertiesResponseOsDiskPtrOutput {
	return o.ApplyT(func(v VirtualmachinesPropertiesResponseStorageProfile) *VirtualmachinesPropertiesResponseOsDisk {
		return v.OsDisk
	}).(VirtualmachinesPropertiesResponseOsDiskPtrOutput)
}

func (o VirtualmachinesPropertiesResponseStorageProfileOutput) VmConfigContainerName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualmachinesPropertiesResponseStorageProfile) *string { return v.VmConfigContainerName }).(pulumi.StringPtrOutput)
}

type VirtualmachinesPropertiesResponseStorageProfilePtrOutput struct{ *pulumi.OutputState }

func (VirtualmachinesPropertiesResponseStorageProfilePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualmachinesPropertiesResponseStorageProfile)(nil)).Elem()
}

func (o VirtualmachinesPropertiesResponseStorageProfilePtrOutput) ToVirtualmachinesPropertiesResponseStorageProfilePtrOutput() VirtualmachinesPropertiesResponseStorageProfilePtrOutput {
	return o
}

func (o VirtualmachinesPropertiesResponseStorageProfilePtrOutput) ToVirtualmachinesPropertiesResponseStorageProfilePtrOutputWithContext(ctx context.Context) VirtualmachinesPropertiesResponseStorageProfilePtrOutput {
	return o
}

func (o VirtualmachinesPropertiesResponseStorageProfilePtrOutput) Elem() VirtualmachinesPropertiesResponseStorageProfileOutput {
	return o.ApplyT(func(v *VirtualmachinesPropertiesResponseStorageProfile) VirtualmachinesPropertiesResponseStorageProfile {
		if v != nil {
			return *v
		}
		var ret VirtualmachinesPropertiesResponseStorageProfile
		return ret
	}).(VirtualmachinesPropertiesResponseStorageProfileOutput)
}

func (o VirtualmachinesPropertiesResponseStorageProfilePtrOutput) DataDisks() VirtualmachinesPropertiesResponseDataDisksArrayOutput {
	return o.ApplyT(func(v *VirtualmachinesPropertiesResponseStorageProfile) []VirtualmachinesPropertiesResponseDataDisks {
		if v == nil {
			return nil
		}
		return v.DataDisks
	}).(VirtualmachinesPropertiesResponseDataDisksArrayOutput)
}

func (o VirtualmachinesPropertiesResponseStorageProfilePtrOutput) ImageReference() VirtualmachinesPropertiesResponseImageReferencePtrOutput {
	return o.ApplyT(func(v *VirtualmachinesPropertiesResponseStorageProfile) *VirtualmachinesPropertiesResponseImageReference {
		if v == nil {
			return nil
		}
		return v.ImageReference
	}).(VirtualmachinesPropertiesResponseImageReferencePtrOutput)
}

func (o VirtualmachinesPropertiesResponseStorageProfilePtrOutput) OsDisk() VirtualmachinesPropertiesResponseOsDiskPtrOutput {
	return o.ApplyT(func(v *VirtualmachinesPropertiesResponseStorageProfile) *VirtualmachinesPropertiesResponseOsDisk {
		if v == nil {
			return nil
		}
		return v.OsDisk
	}).(VirtualmachinesPropertiesResponseOsDiskPtrOutput)
}

func (o VirtualmachinesPropertiesResponseStorageProfilePtrOutput) VmConfigContainerName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualmachinesPropertiesResponseStorageProfile) *string {
		if v == nil {
			return nil
		}
		return v.VmConfigContainerName
	}).(pulumi.StringPtrOutput)
}

type VirtualmachinesPropertiesResponseUefiSettings struct {
	SecureBootEnabled *bool `pulumi:"secureBootEnabled"`
}

type VirtualmachinesPropertiesResponseUefiSettingsOutput struct{ *pulumi.OutputState }

func (VirtualmachinesPropertiesResponseUefiSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualmachinesPropertiesResponseUefiSettings)(nil)).Elem()
}

func (o VirtualmachinesPropertiesResponseUefiSettingsOutput) ToVirtualmachinesPropertiesResponseUefiSettingsOutput() VirtualmachinesPropertiesResponseUefiSettingsOutput {
	return o
}

func (o VirtualmachinesPropertiesResponseUefiSettingsOutput) ToVirtualmachinesPropertiesResponseUefiSettingsOutputWithContext(ctx context.Context) VirtualmachinesPropertiesResponseUefiSettingsOutput {
	return o
}

func (o VirtualmachinesPropertiesResponseUefiSettingsOutput) SecureBootEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VirtualmachinesPropertiesResponseUefiSettings) *bool { return v.SecureBootEnabled }).(pulumi.BoolPtrOutput)
}

type VirtualmachinesPropertiesResponseUefiSettingsPtrOutput struct{ *pulumi.OutputState }

func (VirtualmachinesPropertiesResponseUefiSettingsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualmachinesPropertiesResponseUefiSettings)(nil)).Elem()
}

func (o VirtualmachinesPropertiesResponseUefiSettingsPtrOutput) ToVirtualmachinesPropertiesResponseUefiSettingsPtrOutput() VirtualmachinesPropertiesResponseUefiSettingsPtrOutput {
	return o
}

func (o VirtualmachinesPropertiesResponseUefiSettingsPtrOutput) ToVirtualmachinesPropertiesResponseUefiSettingsPtrOutputWithContext(ctx context.Context) VirtualmachinesPropertiesResponseUefiSettingsPtrOutput {
	return o
}

func (o VirtualmachinesPropertiesResponseUefiSettingsPtrOutput) Elem() VirtualmachinesPropertiesResponseUefiSettingsOutput {
	return o.ApplyT(func(v *VirtualmachinesPropertiesResponseUefiSettings) VirtualmachinesPropertiesResponseUefiSettings {
		if v != nil {
			return *v
		}
		var ret VirtualmachinesPropertiesResponseUefiSettings
		return ret
	}).(VirtualmachinesPropertiesResponseUefiSettingsOutput)
}

func (o VirtualmachinesPropertiesResponseUefiSettingsPtrOutput) SecureBootEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *VirtualmachinesPropertiesResponseUefiSettings) *bool {
		if v == nil {
			return nil
		}
		return v.SecureBootEnabled
	}).(pulumi.BoolPtrOutput)
}

type VirtualmachinesPropertiesResponseWindowsConfiguration struct {
	EnableAutomaticUpdates *bool                                    `pulumi:"enableAutomaticUpdates"`
	ProvisionVMAgent       *bool                                    `pulumi:"provisionVMAgent"`
	Ssh                    *VirtualmachinesPropertiesResponseSshSsh `pulumi:"ssh"`
	TimeZone               *string                                  `pulumi:"timeZone"`
}

type VirtualmachinesPropertiesResponseWindowsConfigurationOutput struct{ *pulumi.OutputState }

func (VirtualmachinesPropertiesResponseWindowsConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualmachinesPropertiesResponseWindowsConfiguration)(nil)).Elem()
}

func (o VirtualmachinesPropertiesResponseWindowsConfigurationOutput) ToVirtualmachinesPropertiesResponseWindowsConfigurationOutput() VirtualmachinesPropertiesResponseWindowsConfigurationOutput {
	return o
}

func (o VirtualmachinesPropertiesResponseWindowsConfigurationOutput) ToVirtualmachinesPropertiesResponseWindowsConfigurationOutputWithContext(ctx context.Context) VirtualmachinesPropertiesResponseWindowsConfigurationOutput {
	return o
}

func (o VirtualmachinesPropertiesResponseWindowsConfigurationOutput) EnableAutomaticUpdates() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VirtualmachinesPropertiesResponseWindowsConfiguration) *bool { return v.EnableAutomaticUpdates }).(pulumi.BoolPtrOutput)
}

func (o VirtualmachinesPropertiesResponseWindowsConfigurationOutput) ProvisionVMAgent() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VirtualmachinesPropertiesResponseWindowsConfiguration) *bool { return v.ProvisionVMAgent }).(pulumi.BoolPtrOutput)
}

func (o VirtualmachinesPropertiesResponseWindowsConfigurationOutput) Ssh() VirtualmachinesPropertiesResponseSshSshPtrOutput {
	return o.ApplyT(func(v VirtualmachinesPropertiesResponseWindowsConfiguration) *VirtualmachinesPropertiesResponseSshSsh {
		return v.Ssh
	}).(VirtualmachinesPropertiesResponseSshSshPtrOutput)
}

func (o VirtualmachinesPropertiesResponseWindowsConfigurationOutput) TimeZone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualmachinesPropertiesResponseWindowsConfiguration) *string { return v.TimeZone }).(pulumi.StringPtrOutput)
}

type VirtualmachinesPropertiesResponseWindowsConfigurationPtrOutput struct{ *pulumi.OutputState }

func (VirtualmachinesPropertiesResponseWindowsConfigurationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualmachinesPropertiesResponseWindowsConfiguration)(nil)).Elem()
}

func (o VirtualmachinesPropertiesResponseWindowsConfigurationPtrOutput) ToVirtualmachinesPropertiesResponseWindowsConfigurationPtrOutput() VirtualmachinesPropertiesResponseWindowsConfigurationPtrOutput {
	return o
}

func (o VirtualmachinesPropertiesResponseWindowsConfigurationPtrOutput) ToVirtualmachinesPropertiesResponseWindowsConfigurationPtrOutputWithContext(ctx context.Context) VirtualmachinesPropertiesResponseWindowsConfigurationPtrOutput {
	return o
}

func (o VirtualmachinesPropertiesResponseWindowsConfigurationPtrOutput) Elem() VirtualmachinesPropertiesResponseWindowsConfigurationOutput {
	return o.ApplyT(func(v *VirtualmachinesPropertiesResponseWindowsConfiguration) VirtualmachinesPropertiesResponseWindowsConfiguration {
		if v != nil {
			return *v
		}
		var ret VirtualmachinesPropertiesResponseWindowsConfiguration
		return ret
	}).(VirtualmachinesPropertiesResponseWindowsConfigurationOutput)
}

func (o VirtualmachinesPropertiesResponseWindowsConfigurationPtrOutput) EnableAutomaticUpdates() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *VirtualmachinesPropertiesResponseWindowsConfiguration) *bool {
		if v == nil {
			return nil
		}
		return v.EnableAutomaticUpdates
	}).(pulumi.BoolPtrOutput)
}

func (o VirtualmachinesPropertiesResponseWindowsConfigurationPtrOutput) ProvisionVMAgent() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *VirtualmachinesPropertiesResponseWindowsConfiguration) *bool {
		if v == nil {
			return nil
		}
		return v.ProvisionVMAgent
	}).(pulumi.BoolPtrOutput)
}

func (o VirtualmachinesPropertiesResponseWindowsConfigurationPtrOutput) Ssh() VirtualmachinesPropertiesResponseSshSshPtrOutput {
	return o.ApplyT(func(v *VirtualmachinesPropertiesResponseWindowsConfiguration) *VirtualmachinesPropertiesResponseSshSsh {
		if v == nil {
			return nil
		}
		return v.Ssh
	}).(VirtualmachinesPropertiesResponseSshSshPtrOutput)
}

func (o VirtualmachinesPropertiesResponseWindowsConfigurationPtrOutput) TimeZone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualmachinesPropertiesResponseWindowsConfiguration) *string {
		if v == nil {
			return nil
		}
		return v.TimeZone
	}).(pulumi.StringPtrOutput)
}

type VirtualmachinesPropertiesSecurityProfile struct {
	EnableTPM    *bool                                  `pulumi:"enableTPM"`
	UefiSettings *VirtualmachinesPropertiesUefiSettings `pulumi:"uefiSettings"`
}





type VirtualmachinesPropertiesSecurityProfileInput interface {
	pulumi.Input

	ToVirtualmachinesPropertiesSecurityProfileOutput() VirtualmachinesPropertiesSecurityProfileOutput
	ToVirtualmachinesPropertiesSecurityProfileOutputWithContext(context.Context) VirtualmachinesPropertiesSecurityProfileOutput
}

type VirtualmachinesPropertiesSecurityProfileArgs struct {
	EnableTPM    pulumi.BoolPtrInput                           `pulumi:"enableTPM"`
	UefiSettings VirtualmachinesPropertiesUefiSettingsPtrInput `pulumi:"uefiSettings"`
}

func (VirtualmachinesPropertiesSecurityProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualmachinesPropertiesSecurityProfile)(nil)).Elem()
}

func (i VirtualmachinesPropertiesSecurityProfileArgs) ToVirtualmachinesPropertiesSecurityProfileOutput() VirtualmachinesPropertiesSecurityProfileOutput {
	return i.ToVirtualmachinesPropertiesSecurityProfileOutputWithContext(context.Background())
}

func (i VirtualmachinesPropertiesSecurityProfileArgs) ToVirtualmachinesPropertiesSecurityProfileOutputWithContext(ctx context.Context) VirtualmachinesPropertiesSecurityProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualmachinesPropertiesSecurityProfileOutput)
}

func (i VirtualmachinesPropertiesSecurityProfileArgs) ToVirtualmachinesPropertiesSecurityProfilePtrOutput() VirtualmachinesPropertiesSecurityProfilePtrOutput {
	return i.ToVirtualmachinesPropertiesSecurityProfilePtrOutputWithContext(context.Background())
}

func (i VirtualmachinesPropertiesSecurityProfileArgs) ToVirtualmachinesPropertiesSecurityProfilePtrOutputWithContext(ctx context.Context) VirtualmachinesPropertiesSecurityProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualmachinesPropertiesSecurityProfileOutput).ToVirtualmachinesPropertiesSecurityProfilePtrOutputWithContext(ctx)
}









type VirtualmachinesPropertiesSecurityProfilePtrInput interface {
	pulumi.Input

	ToVirtualmachinesPropertiesSecurityProfilePtrOutput() VirtualmachinesPropertiesSecurityProfilePtrOutput
	ToVirtualmachinesPropertiesSecurityProfilePtrOutputWithContext(context.Context) VirtualmachinesPropertiesSecurityProfilePtrOutput
}

type virtualmachinesPropertiesSecurityProfilePtrType VirtualmachinesPropertiesSecurityProfileArgs

func VirtualmachinesPropertiesSecurityProfilePtr(v *VirtualmachinesPropertiesSecurityProfileArgs) VirtualmachinesPropertiesSecurityProfilePtrInput {
	return (*virtualmachinesPropertiesSecurityProfilePtrType)(v)
}

func (*virtualmachinesPropertiesSecurityProfilePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualmachinesPropertiesSecurityProfile)(nil)).Elem()
}

func (i *virtualmachinesPropertiesSecurityProfilePtrType) ToVirtualmachinesPropertiesSecurityProfilePtrOutput() VirtualmachinesPropertiesSecurityProfilePtrOutput {
	return i.ToVirtualmachinesPropertiesSecurityProfilePtrOutputWithContext(context.Background())
}

func (i *virtualmachinesPropertiesSecurityProfilePtrType) ToVirtualmachinesPropertiesSecurityProfilePtrOutputWithContext(ctx context.Context) VirtualmachinesPropertiesSecurityProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualmachinesPropertiesSecurityProfilePtrOutput)
}

type VirtualmachinesPropertiesSecurityProfileOutput struct{ *pulumi.OutputState }

func (VirtualmachinesPropertiesSecurityProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualmachinesPropertiesSecurityProfile)(nil)).Elem()
}

func (o VirtualmachinesPropertiesSecurityProfileOutput) ToVirtualmachinesPropertiesSecurityProfileOutput() VirtualmachinesPropertiesSecurityProfileOutput {
	return o
}

func (o VirtualmachinesPropertiesSecurityProfileOutput) ToVirtualmachinesPropertiesSecurityProfileOutputWithContext(ctx context.Context) VirtualmachinesPropertiesSecurityProfileOutput {
	return o
}

func (o VirtualmachinesPropertiesSecurityProfileOutput) ToVirtualmachinesPropertiesSecurityProfilePtrOutput() VirtualmachinesPropertiesSecurityProfilePtrOutput {
	return o.ToVirtualmachinesPropertiesSecurityProfilePtrOutputWithContext(context.Background())
}

func (o VirtualmachinesPropertiesSecurityProfileOutput) ToVirtualmachinesPropertiesSecurityProfilePtrOutputWithContext(ctx context.Context) VirtualmachinesPropertiesSecurityProfilePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v VirtualmachinesPropertiesSecurityProfile) *VirtualmachinesPropertiesSecurityProfile {
		return &v
	}).(VirtualmachinesPropertiesSecurityProfilePtrOutput)
}

func (o VirtualmachinesPropertiesSecurityProfileOutput) EnableTPM() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VirtualmachinesPropertiesSecurityProfile) *bool { return v.EnableTPM }).(pulumi.BoolPtrOutput)
}

func (o VirtualmachinesPropertiesSecurityProfileOutput) UefiSettings() VirtualmachinesPropertiesUefiSettingsPtrOutput {
	return o.ApplyT(func(v VirtualmachinesPropertiesSecurityProfile) *VirtualmachinesPropertiesUefiSettings {
		return v.UefiSettings
	}).(VirtualmachinesPropertiesUefiSettingsPtrOutput)
}

type VirtualmachinesPropertiesSecurityProfilePtrOutput struct{ *pulumi.OutputState }

func (VirtualmachinesPropertiesSecurityProfilePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualmachinesPropertiesSecurityProfile)(nil)).Elem()
}

func (o VirtualmachinesPropertiesSecurityProfilePtrOutput) ToVirtualmachinesPropertiesSecurityProfilePtrOutput() VirtualmachinesPropertiesSecurityProfilePtrOutput {
	return o
}

func (o VirtualmachinesPropertiesSecurityProfilePtrOutput) ToVirtualmachinesPropertiesSecurityProfilePtrOutputWithContext(ctx context.Context) VirtualmachinesPropertiesSecurityProfilePtrOutput {
	return o
}

func (o VirtualmachinesPropertiesSecurityProfilePtrOutput) Elem() VirtualmachinesPropertiesSecurityProfileOutput {
	return o.ApplyT(func(v *VirtualmachinesPropertiesSecurityProfile) VirtualmachinesPropertiesSecurityProfile {
		if v != nil {
			return *v
		}
		var ret VirtualmachinesPropertiesSecurityProfile
		return ret
	}).(VirtualmachinesPropertiesSecurityProfileOutput)
}

func (o VirtualmachinesPropertiesSecurityProfilePtrOutput) EnableTPM() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *VirtualmachinesPropertiesSecurityProfile) *bool {
		if v == nil {
			return nil
		}
		return v.EnableTPM
	}).(pulumi.BoolPtrOutput)
}

func (o VirtualmachinesPropertiesSecurityProfilePtrOutput) UefiSettings() VirtualmachinesPropertiesUefiSettingsPtrOutput {
	return o.ApplyT(func(v *VirtualmachinesPropertiesSecurityProfile) *VirtualmachinesPropertiesUefiSettings {
		if v == nil {
			return nil
		}
		return v.UefiSettings
	}).(VirtualmachinesPropertiesUefiSettingsPtrOutput)
}

type VirtualmachinesPropertiesSsh struct {
	PublicKeys []VirtualmachinesPropertiesPublicKeys `pulumi:"publicKeys"`
}





type VirtualmachinesPropertiesSshInput interface {
	pulumi.Input

	ToVirtualmachinesPropertiesSshOutput() VirtualmachinesPropertiesSshOutput
	ToVirtualmachinesPropertiesSshOutputWithContext(context.Context) VirtualmachinesPropertiesSshOutput
}

type VirtualmachinesPropertiesSshArgs struct {
	PublicKeys VirtualmachinesPropertiesPublicKeysArrayInput `pulumi:"publicKeys"`
}

func (VirtualmachinesPropertiesSshArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualmachinesPropertiesSsh)(nil)).Elem()
}

func (i VirtualmachinesPropertiesSshArgs) ToVirtualmachinesPropertiesSshOutput() VirtualmachinesPropertiesSshOutput {
	return i.ToVirtualmachinesPropertiesSshOutputWithContext(context.Background())
}

func (i VirtualmachinesPropertiesSshArgs) ToVirtualmachinesPropertiesSshOutputWithContext(ctx context.Context) VirtualmachinesPropertiesSshOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualmachinesPropertiesSshOutput)
}

func (i VirtualmachinesPropertiesSshArgs) ToVirtualmachinesPropertiesSshPtrOutput() VirtualmachinesPropertiesSshPtrOutput {
	return i.ToVirtualmachinesPropertiesSshPtrOutputWithContext(context.Background())
}

func (i VirtualmachinesPropertiesSshArgs) ToVirtualmachinesPropertiesSshPtrOutputWithContext(ctx context.Context) VirtualmachinesPropertiesSshPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualmachinesPropertiesSshOutput).ToVirtualmachinesPropertiesSshPtrOutputWithContext(ctx)
}









type VirtualmachinesPropertiesSshPtrInput interface {
	pulumi.Input

	ToVirtualmachinesPropertiesSshPtrOutput() VirtualmachinesPropertiesSshPtrOutput
	ToVirtualmachinesPropertiesSshPtrOutputWithContext(context.Context) VirtualmachinesPropertiesSshPtrOutput
}

type virtualmachinesPropertiesSshPtrType VirtualmachinesPropertiesSshArgs

func VirtualmachinesPropertiesSshPtr(v *VirtualmachinesPropertiesSshArgs) VirtualmachinesPropertiesSshPtrInput {
	return (*virtualmachinesPropertiesSshPtrType)(v)
}

func (*virtualmachinesPropertiesSshPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualmachinesPropertiesSsh)(nil)).Elem()
}

func (i *virtualmachinesPropertiesSshPtrType) ToVirtualmachinesPropertiesSshPtrOutput() VirtualmachinesPropertiesSshPtrOutput {
	return i.ToVirtualmachinesPropertiesSshPtrOutputWithContext(context.Background())
}

func (i *virtualmachinesPropertiesSshPtrType) ToVirtualmachinesPropertiesSshPtrOutputWithContext(ctx context.Context) VirtualmachinesPropertiesSshPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualmachinesPropertiesSshPtrOutput)
}

type VirtualmachinesPropertiesSshOutput struct{ *pulumi.OutputState }

func (VirtualmachinesPropertiesSshOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualmachinesPropertiesSsh)(nil)).Elem()
}

func (o VirtualmachinesPropertiesSshOutput) ToVirtualmachinesPropertiesSshOutput() VirtualmachinesPropertiesSshOutput {
	return o
}

func (o VirtualmachinesPropertiesSshOutput) ToVirtualmachinesPropertiesSshOutputWithContext(ctx context.Context) VirtualmachinesPropertiesSshOutput {
	return o
}

func (o VirtualmachinesPropertiesSshOutput) ToVirtualmachinesPropertiesSshPtrOutput() VirtualmachinesPropertiesSshPtrOutput {
	return o.ToVirtualmachinesPropertiesSshPtrOutputWithContext(context.Background())
}

func (o VirtualmachinesPropertiesSshOutput) ToVirtualmachinesPropertiesSshPtrOutputWithContext(ctx context.Context) VirtualmachinesPropertiesSshPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v VirtualmachinesPropertiesSsh) *VirtualmachinesPropertiesSsh {
		return &v
	}).(VirtualmachinesPropertiesSshPtrOutput)
}

func (o VirtualmachinesPropertiesSshOutput) PublicKeys() VirtualmachinesPropertiesPublicKeysArrayOutput {
	return o.ApplyT(func(v VirtualmachinesPropertiesSsh) []VirtualmachinesPropertiesPublicKeys { return v.PublicKeys }).(VirtualmachinesPropertiesPublicKeysArrayOutput)
}

type VirtualmachinesPropertiesSshPtrOutput struct{ *pulumi.OutputState }

func (VirtualmachinesPropertiesSshPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualmachinesPropertiesSsh)(nil)).Elem()
}

func (o VirtualmachinesPropertiesSshPtrOutput) ToVirtualmachinesPropertiesSshPtrOutput() VirtualmachinesPropertiesSshPtrOutput {
	return o
}

func (o VirtualmachinesPropertiesSshPtrOutput) ToVirtualmachinesPropertiesSshPtrOutputWithContext(ctx context.Context) VirtualmachinesPropertiesSshPtrOutput {
	return o
}

func (o VirtualmachinesPropertiesSshPtrOutput) Elem() VirtualmachinesPropertiesSshOutput {
	return o.ApplyT(func(v *VirtualmachinesPropertiesSsh) VirtualmachinesPropertiesSsh {
		if v != nil {
			return *v
		}
		var ret VirtualmachinesPropertiesSsh
		return ret
	}).(VirtualmachinesPropertiesSshOutput)
}

func (o VirtualmachinesPropertiesSshPtrOutput) PublicKeys() VirtualmachinesPropertiesPublicKeysArrayOutput {
	return o.ApplyT(func(v *VirtualmachinesPropertiesSsh) []VirtualmachinesPropertiesPublicKeys {
		if v == nil {
			return nil
		}
		return v.PublicKeys
	}).(VirtualmachinesPropertiesPublicKeysArrayOutput)
}

type VirtualmachinesPropertiesSshSsh struct {
	PublicKeys []VirtualmachinesPropertiesPublicKeysPublicKeys `pulumi:"publicKeys"`
}





type VirtualmachinesPropertiesSshSshInput interface {
	pulumi.Input

	ToVirtualmachinesPropertiesSshSshOutput() VirtualmachinesPropertiesSshSshOutput
	ToVirtualmachinesPropertiesSshSshOutputWithContext(context.Context) VirtualmachinesPropertiesSshSshOutput
}

type VirtualmachinesPropertiesSshSshArgs struct {
	PublicKeys VirtualmachinesPropertiesPublicKeysPublicKeysArrayInput `pulumi:"publicKeys"`
}

func (VirtualmachinesPropertiesSshSshArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualmachinesPropertiesSshSsh)(nil)).Elem()
}

func (i VirtualmachinesPropertiesSshSshArgs) ToVirtualmachinesPropertiesSshSshOutput() VirtualmachinesPropertiesSshSshOutput {
	return i.ToVirtualmachinesPropertiesSshSshOutputWithContext(context.Background())
}

func (i VirtualmachinesPropertiesSshSshArgs) ToVirtualmachinesPropertiesSshSshOutputWithContext(ctx context.Context) VirtualmachinesPropertiesSshSshOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualmachinesPropertiesSshSshOutput)
}

func (i VirtualmachinesPropertiesSshSshArgs) ToVirtualmachinesPropertiesSshSshPtrOutput() VirtualmachinesPropertiesSshSshPtrOutput {
	return i.ToVirtualmachinesPropertiesSshSshPtrOutputWithContext(context.Background())
}

func (i VirtualmachinesPropertiesSshSshArgs) ToVirtualmachinesPropertiesSshSshPtrOutputWithContext(ctx context.Context) VirtualmachinesPropertiesSshSshPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualmachinesPropertiesSshSshOutput).ToVirtualmachinesPropertiesSshSshPtrOutputWithContext(ctx)
}









type VirtualmachinesPropertiesSshSshPtrInput interface {
	pulumi.Input

	ToVirtualmachinesPropertiesSshSshPtrOutput() VirtualmachinesPropertiesSshSshPtrOutput
	ToVirtualmachinesPropertiesSshSshPtrOutputWithContext(context.Context) VirtualmachinesPropertiesSshSshPtrOutput
}

type virtualmachinesPropertiesSshSshPtrType VirtualmachinesPropertiesSshSshArgs

func VirtualmachinesPropertiesSshSshPtr(v *VirtualmachinesPropertiesSshSshArgs) VirtualmachinesPropertiesSshSshPtrInput {
	return (*virtualmachinesPropertiesSshSshPtrType)(v)
}

func (*virtualmachinesPropertiesSshSshPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualmachinesPropertiesSshSsh)(nil)).Elem()
}

func (i *virtualmachinesPropertiesSshSshPtrType) ToVirtualmachinesPropertiesSshSshPtrOutput() VirtualmachinesPropertiesSshSshPtrOutput {
	return i.ToVirtualmachinesPropertiesSshSshPtrOutputWithContext(context.Background())
}

func (i *virtualmachinesPropertiesSshSshPtrType) ToVirtualmachinesPropertiesSshSshPtrOutputWithContext(ctx context.Context) VirtualmachinesPropertiesSshSshPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualmachinesPropertiesSshSshPtrOutput)
}

type VirtualmachinesPropertiesSshSshOutput struct{ *pulumi.OutputState }

func (VirtualmachinesPropertiesSshSshOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualmachinesPropertiesSshSsh)(nil)).Elem()
}

func (o VirtualmachinesPropertiesSshSshOutput) ToVirtualmachinesPropertiesSshSshOutput() VirtualmachinesPropertiesSshSshOutput {
	return o
}

func (o VirtualmachinesPropertiesSshSshOutput) ToVirtualmachinesPropertiesSshSshOutputWithContext(ctx context.Context) VirtualmachinesPropertiesSshSshOutput {
	return o
}

func (o VirtualmachinesPropertiesSshSshOutput) ToVirtualmachinesPropertiesSshSshPtrOutput() VirtualmachinesPropertiesSshSshPtrOutput {
	return o.ToVirtualmachinesPropertiesSshSshPtrOutputWithContext(context.Background())
}

func (o VirtualmachinesPropertiesSshSshOutput) ToVirtualmachinesPropertiesSshSshPtrOutputWithContext(ctx context.Context) VirtualmachinesPropertiesSshSshPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v VirtualmachinesPropertiesSshSsh) *VirtualmachinesPropertiesSshSsh {
		return &v
	}).(VirtualmachinesPropertiesSshSshPtrOutput)
}

func (o VirtualmachinesPropertiesSshSshOutput) PublicKeys() VirtualmachinesPropertiesPublicKeysPublicKeysArrayOutput {
	return o.ApplyT(func(v VirtualmachinesPropertiesSshSsh) []VirtualmachinesPropertiesPublicKeysPublicKeys {
		return v.PublicKeys
	}).(VirtualmachinesPropertiesPublicKeysPublicKeysArrayOutput)
}

type VirtualmachinesPropertiesSshSshPtrOutput struct{ *pulumi.OutputState }

func (VirtualmachinesPropertiesSshSshPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualmachinesPropertiesSshSsh)(nil)).Elem()
}

func (o VirtualmachinesPropertiesSshSshPtrOutput) ToVirtualmachinesPropertiesSshSshPtrOutput() VirtualmachinesPropertiesSshSshPtrOutput {
	return o
}

func (o VirtualmachinesPropertiesSshSshPtrOutput) ToVirtualmachinesPropertiesSshSshPtrOutputWithContext(ctx context.Context) VirtualmachinesPropertiesSshSshPtrOutput {
	return o
}

func (o VirtualmachinesPropertiesSshSshPtrOutput) Elem() VirtualmachinesPropertiesSshSshOutput {
	return o.ApplyT(func(v *VirtualmachinesPropertiesSshSsh) VirtualmachinesPropertiesSshSsh {
		if v != nil {
			return *v
		}
		var ret VirtualmachinesPropertiesSshSsh
		return ret
	}).(VirtualmachinesPropertiesSshSshOutput)
}

func (o VirtualmachinesPropertiesSshSshPtrOutput) PublicKeys() VirtualmachinesPropertiesPublicKeysPublicKeysArrayOutput {
	return o.ApplyT(func(v *VirtualmachinesPropertiesSshSsh) []VirtualmachinesPropertiesPublicKeysPublicKeys {
		if v == nil {
			return nil
		}
		return v.PublicKeys
	}).(VirtualmachinesPropertiesPublicKeysPublicKeysArrayOutput)
}

type VirtualmachinesPropertiesStorageProfile struct {
	DataDisks             []VirtualmachinesPropertiesDataDisks     `pulumi:"dataDisks"`
	ImageReference        *VirtualmachinesPropertiesImageReference `pulumi:"imageReference"`
	OsDisk                *VirtualmachinesPropertiesOsDisk         `pulumi:"osDisk"`
	VmConfigContainerName *string                                  `pulumi:"vmConfigContainerName"`
}





type VirtualmachinesPropertiesStorageProfileInput interface {
	pulumi.Input

	ToVirtualmachinesPropertiesStorageProfileOutput() VirtualmachinesPropertiesStorageProfileOutput
	ToVirtualmachinesPropertiesStorageProfileOutputWithContext(context.Context) VirtualmachinesPropertiesStorageProfileOutput
}

type VirtualmachinesPropertiesStorageProfileArgs struct {
	DataDisks             VirtualmachinesPropertiesDataDisksArrayInput    `pulumi:"dataDisks"`
	ImageReference        VirtualmachinesPropertiesImageReferencePtrInput `pulumi:"imageReference"`
	OsDisk                VirtualmachinesPropertiesOsDiskPtrInput         `pulumi:"osDisk"`
	VmConfigContainerName pulumi.StringPtrInput                           `pulumi:"vmConfigContainerName"`
}

func (VirtualmachinesPropertiesStorageProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualmachinesPropertiesStorageProfile)(nil)).Elem()
}

func (i VirtualmachinesPropertiesStorageProfileArgs) ToVirtualmachinesPropertiesStorageProfileOutput() VirtualmachinesPropertiesStorageProfileOutput {
	return i.ToVirtualmachinesPropertiesStorageProfileOutputWithContext(context.Background())
}

func (i VirtualmachinesPropertiesStorageProfileArgs) ToVirtualmachinesPropertiesStorageProfileOutputWithContext(ctx context.Context) VirtualmachinesPropertiesStorageProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualmachinesPropertiesStorageProfileOutput)
}

func (i VirtualmachinesPropertiesStorageProfileArgs) ToVirtualmachinesPropertiesStorageProfilePtrOutput() VirtualmachinesPropertiesStorageProfilePtrOutput {
	return i.ToVirtualmachinesPropertiesStorageProfilePtrOutputWithContext(context.Background())
}

func (i VirtualmachinesPropertiesStorageProfileArgs) ToVirtualmachinesPropertiesStorageProfilePtrOutputWithContext(ctx context.Context) VirtualmachinesPropertiesStorageProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualmachinesPropertiesStorageProfileOutput).ToVirtualmachinesPropertiesStorageProfilePtrOutputWithContext(ctx)
}









type VirtualmachinesPropertiesStorageProfilePtrInput interface {
	pulumi.Input

	ToVirtualmachinesPropertiesStorageProfilePtrOutput() VirtualmachinesPropertiesStorageProfilePtrOutput
	ToVirtualmachinesPropertiesStorageProfilePtrOutputWithContext(context.Context) VirtualmachinesPropertiesStorageProfilePtrOutput
}

type virtualmachinesPropertiesStorageProfilePtrType VirtualmachinesPropertiesStorageProfileArgs

func VirtualmachinesPropertiesStorageProfilePtr(v *VirtualmachinesPropertiesStorageProfileArgs) VirtualmachinesPropertiesStorageProfilePtrInput {
	return (*virtualmachinesPropertiesStorageProfilePtrType)(v)
}

func (*virtualmachinesPropertiesStorageProfilePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualmachinesPropertiesStorageProfile)(nil)).Elem()
}

func (i *virtualmachinesPropertiesStorageProfilePtrType) ToVirtualmachinesPropertiesStorageProfilePtrOutput() VirtualmachinesPropertiesStorageProfilePtrOutput {
	return i.ToVirtualmachinesPropertiesStorageProfilePtrOutputWithContext(context.Background())
}

func (i *virtualmachinesPropertiesStorageProfilePtrType) ToVirtualmachinesPropertiesStorageProfilePtrOutputWithContext(ctx context.Context) VirtualmachinesPropertiesStorageProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualmachinesPropertiesStorageProfilePtrOutput)
}

type VirtualmachinesPropertiesStorageProfileOutput struct{ *pulumi.OutputState }

func (VirtualmachinesPropertiesStorageProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualmachinesPropertiesStorageProfile)(nil)).Elem()
}

func (o VirtualmachinesPropertiesStorageProfileOutput) ToVirtualmachinesPropertiesStorageProfileOutput() VirtualmachinesPropertiesStorageProfileOutput {
	return o
}

func (o VirtualmachinesPropertiesStorageProfileOutput) ToVirtualmachinesPropertiesStorageProfileOutputWithContext(ctx context.Context) VirtualmachinesPropertiesStorageProfileOutput {
	return o
}

func (o VirtualmachinesPropertiesStorageProfileOutput) ToVirtualmachinesPropertiesStorageProfilePtrOutput() VirtualmachinesPropertiesStorageProfilePtrOutput {
	return o.ToVirtualmachinesPropertiesStorageProfilePtrOutputWithContext(context.Background())
}

func (o VirtualmachinesPropertiesStorageProfileOutput) ToVirtualmachinesPropertiesStorageProfilePtrOutputWithContext(ctx context.Context) VirtualmachinesPropertiesStorageProfilePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v VirtualmachinesPropertiesStorageProfile) *VirtualmachinesPropertiesStorageProfile {
		return &v
	}).(VirtualmachinesPropertiesStorageProfilePtrOutput)
}

func (o VirtualmachinesPropertiesStorageProfileOutput) DataDisks() VirtualmachinesPropertiesDataDisksArrayOutput {
	return o.ApplyT(func(v VirtualmachinesPropertiesStorageProfile) []VirtualmachinesPropertiesDataDisks {
		return v.DataDisks
	}).(VirtualmachinesPropertiesDataDisksArrayOutput)
}

func (o VirtualmachinesPropertiesStorageProfileOutput) ImageReference() VirtualmachinesPropertiesImageReferencePtrOutput {
	return o.ApplyT(func(v VirtualmachinesPropertiesStorageProfile) *VirtualmachinesPropertiesImageReference {
		return v.ImageReference
	}).(VirtualmachinesPropertiesImageReferencePtrOutput)
}

func (o VirtualmachinesPropertiesStorageProfileOutput) OsDisk() VirtualmachinesPropertiesOsDiskPtrOutput {
	return o.ApplyT(func(v VirtualmachinesPropertiesStorageProfile) *VirtualmachinesPropertiesOsDisk { return v.OsDisk }).(VirtualmachinesPropertiesOsDiskPtrOutput)
}

func (o VirtualmachinesPropertiesStorageProfileOutput) VmConfigContainerName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualmachinesPropertiesStorageProfile) *string { return v.VmConfigContainerName }).(pulumi.StringPtrOutput)
}

type VirtualmachinesPropertiesStorageProfilePtrOutput struct{ *pulumi.OutputState }

func (VirtualmachinesPropertiesStorageProfilePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualmachinesPropertiesStorageProfile)(nil)).Elem()
}

func (o VirtualmachinesPropertiesStorageProfilePtrOutput) ToVirtualmachinesPropertiesStorageProfilePtrOutput() VirtualmachinesPropertiesStorageProfilePtrOutput {
	return o
}

func (o VirtualmachinesPropertiesStorageProfilePtrOutput) ToVirtualmachinesPropertiesStorageProfilePtrOutputWithContext(ctx context.Context) VirtualmachinesPropertiesStorageProfilePtrOutput {
	return o
}

func (o VirtualmachinesPropertiesStorageProfilePtrOutput) Elem() VirtualmachinesPropertiesStorageProfileOutput {
	return o.ApplyT(func(v *VirtualmachinesPropertiesStorageProfile) VirtualmachinesPropertiesStorageProfile {
		if v != nil {
			return *v
		}
		var ret VirtualmachinesPropertiesStorageProfile
		return ret
	}).(VirtualmachinesPropertiesStorageProfileOutput)
}

func (o VirtualmachinesPropertiesStorageProfilePtrOutput) DataDisks() VirtualmachinesPropertiesDataDisksArrayOutput {
	return o.ApplyT(func(v *VirtualmachinesPropertiesStorageProfile) []VirtualmachinesPropertiesDataDisks {
		if v == nil {
			return nil
		}
		return v.DataDisks
	}).(VirtualmachinesPropertiesDataDisksArrayOutput)
}

func (o VirtualmachinesPropertiesStorageProfilePtrOutput) ImageReference() VirtualmachinesPropertiesImageReferencePtrOutput {
	return o.ApplyT(func(v *VirtualmachinesPropertiesStorageProfile) *VirtualmachinesPropertiesImageReference {
		if v == nil {
			return nil
		}
		return v.ImageReference
	}).(VirtualmachinesPropertiesImageReferencePtrOutput)
}

func (o VirtualmachinesPropertiesStorageProfilePtrOutput) OsDisk() VirtualmachinesPropertiesOsDiskPtrOutput {
	return o.ApplyT(func(v *VirtualmachinesPropertiesStorageProfile) *VirtualmachinesPropertiesOsDisk {
		if v == nil {
			return nil
		}
		return v.OsDisk
	}).(VirtualmachinesPropertiesOsDiskPtrOutput)
}

func (o VirtualmachinesPropertiesStorageProfilePtrOutput) VmConfigContainerName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualmachinesPropertiesStorageProfile) *string {
		if v == nil {
			return nil
		}
		return v.VmConfigContainerName
	}).(pulumi.StringPtrOutput)
}

type VirtualmachinesPropertiesUefiSettings struct {
	SecureBootEnabled *bool `pulumi:"secureBootEnabled"`
}





type VirtualmachinesPropertiesUefiSettingsInput interface {
	pulumi.Input

	ToVirtualmachinesPropertiesUefiSettingsOutput() VirtualmachinesPropertiesUefiSettingsOutput
	ToVirtualmachinesPropertiesUefiSettingsOutputWithContext(context.Context) VirtualmachinesPropertiesUefiSettingsOutput
}

type VirtualmachinesPropertiesUefiSettingsArgs struct {
	SecureBootEnabled pulumi.BoolPtrInput `pulumi:"secureBootEnabled"`
}

func (VirtualmachinesPropertiesUefiSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualmachinesPropertiesUefiSettings)(nil)).Elem()
}

func (i VirtualmachinesPropertiesUefiSettingsArgs) ToVirtualmachinesPropertiesUefiSettingsOutput() VirtualmachinesPropertiesUefiSettingsOutput {
	return i.ToVirtualmachinesPropertiesUefiSettingsOutputWithContext(context.Background())
}

func (i VirtualmachinesPropertiesUefiSettingsArgs) ToVirtualmachinesPropertiesUefiSettingsOutputWithContext(ctx context.Context) VirtualmachinesPropertiesUefiSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualmachinesPropertiesUefiSettingsOutput)
}

func (i VirtualmachinesPropertiesUefiSettingsArgs) ToVirtualmachinesPropertiesUefiSettingsPtrOutput() VirtualmachinesPropertiesUefiSettingsPtrOutput {
	return i.ToVirtualmachinesPropertiesUefiSettingsPtrOutputWithContext(context.Background())
}

func (i VirtualmachinesPropertiesUefiSettingsArgs) ToVirtualmachinesPropertiesUefiSettingsPtrOutputWithContext(ctx context.Context) VirtualmachinesPropertiesUefiSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualmachinesPropertiesUefiSettingsOutput).ToVirtualmachinesPropertiesUefiSettingsPtrOutputWithContext(ctx)
}









type VirtualmachinesPropertiesUefiSettingsPtrInput interface {
	pulumi.Input

	ToVirtualmachinesPropertiesUefiSettingsPtrOutput() VirtualmachinesPropertiesUefiSettingsPtrOutput
	ToVirtualmachinesPropertiesUefiSettingsPtrOutputWithContext(context.Context) VirtualmachinesPropertiesUefiSettingsPtrOutput
}

type virtualmachinesPropertiesUefiSettingsPtrType VirtualmachinesPropertiesUefiSettingsArgs

func VirtualmachinesPropertiesUefiSettingsPtr(v *VirtualmachinesPropertiesUefiSettingsArgs) VirtualmachinesPropertiesUefiSettingsPtrInput {
	return (*virtualmachinesPropertiesUefiSettingsPtrType)(v)
}

func (*virtualmachinesPropertiesUefiSettingsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualmachinesPropertiesUefiSettings)(nil)).Elem()
}

func (i *virtualmachinesPropertiesUefiSettingsPtrType) ToVirtualmachinesPropertiesUefiSettingsPtrOutput() VirtualmachinesPropertiesUefiSettingsPtrOutput {
	return i.ToVirtualmachinesPropertiesUefiSettingsPtrOutputWithContext(context.Background())
}

func (i *virtualmachinesPropertiesUefiSettingsPtrType) ToVirtualmachinesPropertiesUefiSettingsPtrOutputWithContext(ctx context.Context) VirtualmachinesPropertiesUefiSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualmachinesPropertiesUefiSettingsPtrOutput)
}

type VirtualmachinesPropertiesUefiSettingsOutput struct{ *pulumi.OutputState }

func (VirtualmachinesPropertiesUefiSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualmachinesPropertiesUefiSettings)(nil)).Elem()
}

func (o VirtualmachinesPropertiesUefiSettingsOutput) ToVirtualmachinesPropertiesUefiSettingsOutput() VirtualmachinesPropertiesUefiSettingsOutput {
	return o
}

func (o VirtualmachinesPropertiesUefiSettingsOutput) ToVirtualmachinesPropertiesUefiSettingsOutputWithContext(ctx context.Context) VirtualmachinesPropertiesUefiSettingsOutput {
	return o
}

func (o VirtualmachinesPropertiesUefiSettingsOutput) ToVirtualmachinesPropertiesUefiSettingsPtrOutput() VirtualmachinesPropertiesUefiSettingsPtrOutput {
	return o.ToVirtualmachinesPropertiesUefiSettingsPtrOutputWithContext(context.Background())
}

func (o VirtualmachinesPropertiesUefiSettingsOutput) ToVirtualmachinesPropertiesUefiSettingsPtrOutputWithContext(ctx context.Context) VirtualmachinesPropertiesUefiSettingsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v VirtualmachinesPropertiesUefiSettings) *VirtualmachinesPropertiesUefiSettings {
		return &v
	}).(VirtualmachinesPropertiesUefiSettingsPtrOutput)
}

func (o VirtualmachinesPropertiesUefiSettingsOutput) SecureBootEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VirtualmachinesPropertiesUefiSettings) *bool { return v.SecureBootEnabled }).(pulumi.BoolPtrOutput)
}

type VirtualmachinesPropertiesUefiSettingsPtrOutput struct{ *pulumi.OutputState }

func (VirtualmachinesPropertiesUefiSettingsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualmachinesPropertiesUefiSettings)(nil)).Elem()
}

func (o VirtualmachinesPropertiesUefiSettingsPtrOutput) ToVirtualmachinesPropertiesUefiSettingsPtrOutput() VirtualmachinesPropertiesUefiSettingsPtrOutput {
	return o
}

func (o VirtualmachinesPropertiesUefiSettingsPtrOutput) ToVirtualmachinesPropertiesUefiSettingsPtrOutputWithContext(ctx context.Context) VirtualmachinesPropertiesUefiSettingsPtrOutput {
	return o
}

func (o VirtualmachinesPropertiesUefiSettingsPtrOutput) Elem() VirtualmachinesPropertiesUefiSettingsOutput {
	return o.ApplyT(func(v *VirtualmachinesPropertiesUefiSettings) VirtualmachinesPropertiesUefiSettings {
		if v != nil {
			return *v
		}
		var ret VirtualmachinesPropertiesUefiSettings
		return ret
	}).(VirtualmachinesPropertiesUefiSettingsOutput)
}

func (o VirtualmachinesPropertiesUefiSettingsPtrOutput) SecureBootEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *VirtualmachinesPropertiesUefiSettings) *bool {
		if v == nil {
			return nil
		}
		return v.SecureBootEnabled
	}).(pulumi.BoolPtrOutput)
}

type VirtualmachinesPropertiesWindowsConfiguration struct {
	EnableAutomaticUpdates *bool                            `pulumi:"enableAutomaticUpdates"`
	ProvisionVMAgent       *bool                            `pulumi:"provisionVMAgent"`
	Ssh                    *VirtualmachinesPropertiesSshSsh `pulumi:"ssh"`
	TimeZone               *string                          `pulumi:"timeZone"`
}





type VirtualmachinesPropertiesWindowsConfigurationInput interface {
	pulumi.Input

	ToVirtualmachinesPropertiesWindowsConfigurationOutput() VirtualmachinesPropertiesWindowsConfigurationOutput
	ToVirtualmachinesPropertiesWindowsConfigurationOutputWithContext(context.Context) VirtualmachinesPropertiesWindowsConfigurationOutput
}

type VirtualmachinesPropertiesWindowsConfigurationArgs struct {
	EnableAutomaticUpdates pulumi.BoolPtrInput                     `pulumi:"enableAutomaticUpdates"`
	ProvisionVMAgent       pulumi.BoolPtrInput                     `pulumi:"provisionVMAgent"`
	Ssh                    VirtualmachinesPropertiesSshSshPtrInput `pulumi:"ssh"`
	TimeZone               pulumi.StringPtrInput                   `pulumi:"timeZone"`
}

func (VirtualmachinesPropertiesWindowsConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualmachinesPropertiesWindowsConfiguration)(nil)).Elem()
}

func (i VirtualmachinesPropertiesWindowsConfigurationArgs) ToVirtualmachinesPropertiesWindowsConfigurationOutput() VirtualmachinesPropertiesWindowsConfigurationOutput {
	return i.ToVirtualmachinesPropertiesWindowsConfigurationOutputWithContext(context.Background())
}

func (i VirtualmachinesPropertiesWindowsConfigurationArgs) ToVirtualmachinesPropertiesWindowsConfigurationOutputWithContext(ctx context.Context) VirtualmachinesPropertiesWindowsConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualmachinesPropertiesWindowsConfigurationOutput)
}

func (i VirtualmachinesPropertiesWindowsConfigurationArgs) ToVirtualmachinesPropertiesWindowsConfigurationPtrOutput() VirtualmachinesPropertiesWindowsConfigurationPtrOutput {
	return i.ToVirtualmachinesPropertiesWindowsConfigurationPtrOutputWithContext(context.Background())
}

func (i VirtualmachinesPropertiesWindowsConfigurationArgs) ToVirtualmachinesPropertiesWindowsConfigurationPtrOutputWithContext(ctx context.Context) VirtualmachinesPropertiesWindowsConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualmachinesPropertiesWindowsConfigurationOutput).ToVirtualmachinesPropertiesWindowsConfigurationPtrOutputWithContext(ctx)
}









type VirtualmachinesPropertiesWindowsConfigurationPtrInput interface {
	pulumi.Input

	ToVirtualmachinesPropertiesWindowsConfigurationPtrOutput() VirtualmachinesPropertiesWindowsConfigurationPtrOutput
	ToVirtualmachinesPropertiesWindowsConfigurationPtrOutputWithContext(context.Context) VirtualmachinesPropertiesWindowsConfigurationPtrOutput
}

type virtualmachinesPropertiesWindowsConfigurationPtrType VirtualmachinesPropertiesWindowsConfigurationArgs

func VirtualmachinesPropertiesWindowsConfigurationPtr(v *VirtualmachinesPropertiesWindowsConfigurationArgs) VirtualmachinesPropertiesWindowsConfigurationPtrInput {
	return (*virtualmachinesPropertiesWindowsConfigurationPtrType)(v)
}

func (*virtualmachinesPropertiesWindowsConfigurationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualmachinesPropertiesWindowsConfiguration)(nil)).Elem()
}

func (i *virtualmachinesPropertiesWindowsConfigurationPtrType) ToVirtualmachinesPropertiesWindowsConfigurationPtrOutput() VirtualmachinesPropertiesWindowsConfigurationPtrOutput {
	return i.ToVirtualmachinesPropertiesWindowsConfigurationPtrOutputWithContext(context.Background())
}

func (i *virtualmachinesPropertiesWindowsConfigurationPtrType) ToVirtualmachinesPropertiesWindowsConfigurationPtrOutputWithContext(ctx context.Context) VirtualmachinesPropertiesWindowsConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualmachinesPropertiesWindowsConfigurationPtrOutput)
}

type VirtualmachinesPropertiesWindowsConfigurationOutput struct{ *pulumi.OutputState }

func (VirtualmachinesPropertiesWindowsConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualmachinesPropertiesWindowsConfiguration)(nil)).Elem()
}

func (o VirtualmachinesPropertiesWindowsConfigurationOutput) ToVirtualmachinesPropertiesWindowsConfigurationOutput() VirtualmachinesPropertiesWindowsConfigurationOutput {
	return o
}

func (o VirtualmachinesPropertiesWindowsConfigurationOutput) ToVirtualmachinesPropertiesWindowsConfigurationOutputWithContext(ctx context.Context) VirtualmachinesPropertiesWindowsConfigurationOutput {
	return o
}

func (o VirtualmachinesPropertiesWindowsConfigurationOutput) ToVirtualmachinesPropertiesWindowsConfigurationPtrOutput() VirtualmachinesPropertiesWindowsConfigurationPtrOutput {
	return o.ToVirtualmachinesPropertiesWindowsConfigurationPtrOutputWithContext(context.Background())
}

func (o VirtualmachinesPropertiesWindowsConfigurationOutput) ToVirtualmachinesPropertiesWindowsConfigurationPtrOutputWithContext(ctx context.Context) VirtualmachinesPropertiesWindowsConfigurationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v VirtualmachinesPropertiesWindowsConfiguration) *VirtualmachinesPropertiesWindowsConfiguration {
		return &v
	}).(VirtualmachinesPropertiesWindowsConfigurationPtrOutput)
}

func (o VirtualmachinesPropertiesWindowsConfigurationOutput) EnableAutomaticUpdates() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VirtualmachinesPropertiesWindowsConfiguration) *bool { return v.EnableAutomaticUpdates }).(pulumi.BoolPtrOutput)
}

func (o VirtualmachinesPropertiesWindowsConfigurationOutput) ProvisionVMAgent() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VirtualmachinesPropertiesWindowsConfiguration) *bool { return v.ProvisionVMAgent }).(pulumi.BoolPtrOutput)
}

func (o VirtualmachinesPropertiesWindowsConfigurationOutput) Ssh() VirtualmachinesPropertiesSshSshPtrOutput {
	return o.ApplyT(func(v VirtualmachinesPropertiesWindowsConfiguration) *VirtualmachinesPropertiesSshSsh { return v.Ssh }).(VirtualmachinesPropertiesSshSshPtrOutput)
}

func (o VirtualmachinesPropertiesWindowsConfigurationOutput) TimeZone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualmachinesPropertiesWindowsConfiguration) *string { return v.TimeZone }).(pulumi.StringPtrOutput)
}

type VirtualmachinesPropertiesWindowsConfigurationPtrOutput struct{ *pulumi.OutputState }

func (VirtualmachinesPropertiesWindowsConfigurationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualmachinesPropertiesWindowsConfiguration)(nil)).Elem()
}

func (o VirtualmachinesPropertiesWindowsConfigurationPtrOutput) ToVirtualmachinesPropertiesWindowsConfigurationPtrOutput() VirtualmachinesPropertiesWindowsConfigurationPtrOutput {
	return o
}

func (o VirtualmachinesPropertiesWindowsConfigurationPtrOutput) ToVirtualmachinesPropertiesWindowsConfigurationPtrOutputWithContext(ctx context.Context) VirtualmachinesPropertiesWindowsConfigurationPtrOutput {
	return o
}

func (o VirtualmachinesPropertiesWindowsConfigurationPtrOutput) Elem() VirtualmachinesPropertiesWindowsConfigurationOutput {
	return o.ApplyT(func(v *VirtualmachinesPropertiesWindowsConfiguration) VirtualmachinesPropertiesWindowsConfiguration {
		if v != nil {
			return *v
		}
		var ret VirtualmachinesPropertiesWindowsConfiguration
		return ret
	}).(VirtualmachinesPropertiesWindowsConfigurationOutput)
}

func (o VirtualmachinesPropertiesWindowsConfigurationPtrOutput) EnableAutomaticUpdates() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *VirtualmachinesPropertiesWindowsConfiguration) *bool {
		if v == nil {
			return nil
		}
		return v.EnableAutomaticUpdates
	}).(pulumi.BoolPtrOutput)
}

func (o VirtualmachinesPropertiesWindowsConfigurationPtrOutput) ProvisionVMAgent() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *VirtualmachinesPropertiesWindowsConfiguration) *bool {
		if v == nil {
			return nil
		}
		return v.ProvisionVMAgent
	}).(pulumi.BoolPtrOutput)
}

func (o VirtualmachinesPropertiesWindowsConfigurationPtrOutput) Ssh() VirtualmachinesPropertiesSshSshPtrOutput {
	return o.ApplyT(func(v *VirtualmachinesPropertiesWindowsConfiguration) *VirtualmachinesPropertiesSshSsh {
		if v == nil {
			return nil
		}
		return v.Ssh
	}).(VirtualmachinesPropertiesSshSshPtrOutput)
}

func (o VirtualmachinesPropertiesWindowsConfigurationPtrOutput) TimeZone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualmachinesPropertiesWindowsConfiguration) *string {
		if v == nil {
			return nil
		}
		return v.TimeZone
	}).(pulumi.StringPtrOutput)
}

type VirtualnetworksPropertiesIpConfigurationReferences struct {
	Id *string `pulumi:"id"`
}





type VirtualnetworksPropertiesIpConfigurationReferencesInput interface {
	pulumi.Input

	ToVirtualnetworksPropertiesIpConfigurationReferencesOutput() VirtualnetworksPropertiesIpConfigurationReferencesOutput
	ToVirtualnetworksPropertiesIpConfigurationReferencesOutputWithContext(context.Context) VirtualnetworksPropertiesIpConfigurationReferencesOutput
}

type VirtualnetworksPropertiesIpConfigurationReferencesArgs struct {
	Id pulumi.StringPtrInput `pulumi:"id"`
}

func (VirtualnetworksPropertiesIpConfigurationReferencesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualnetworksPropertiesIpConfigurationReferences)(nil)).Elem()
}

func (i VirtualnetworksPropertiesIpConfigurationReferencesArgs) ToVirtualnetworksPropertiesIpConfigurationReferencesOutput() VirtualnetworksPropertiesIpConfigurationReferencesOutput {
	return i.ToVirtualnetworksPropertiesIpConfigurationReferencesOutputWithContext(context.Background())
}

func (i VirtualnetworksPropertiesIpConfigurationReferencesArgs) ToVirtualnetworksPropertiesIpConfigurationReferencesOutputWithContext(ctx context.Context) VirtualnetworksPropertiesIpConfigurationReferencesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualnetworksPropertiesIpConfigurationReferencesOutput)
}





type VirtualnetworksPropertiesIpConfigurationReferencesArrayInput interface {
	pulumi.Input

	ToVirtualnetworksPropertiesIpConfigurationReferencesArrayOutput() VirtualnetworksPropertiesIpConfigurationReferencesArrayOutput
	ToVirtualnetworksPropertiesIpConfigurationReferencesArrayOutputWithContext(context.Context) VirtualnetworksPropertiesIpConfigurationReferencesArrayOutput
}

type VirtualnetworksPropertiesIpConfigurationReferencesArray []VirtualnetworksPropertiesIpConfigurationReferencesInput

func (VirtualnetworksPropertiesIpConfigurationReferencesArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VirtualnetworksPropertiesIpConfigurationReferences)(nil)).Elem()
}

func (i VirtualnetworksPropertiesIpConfigurationReferencesArray) ToVirtualnetworksPropertiesIpConfigurationReferencesArrayOutput() VirtualnetworksPropertiesIpConfigurationReferencesArrayOutput {
	return i.ToVirtualnetworksPropertiesIpConfigurationReferencesArrayOutputWithContext(context.Background())
}

func (i VirtualnetworksPropertiesIpConfigurationReferencesArray) ToVirtualnetworksPropertiesIpConfigurationReferencesArrayOutputWithContext(ctx context.Context) VirtualnetworksPropertiesIpConfigurationReferencesArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualnetworksPropertiesIpConfigurationReferencesArrayOutput)
}

type VirtualnetworksPropertiesIpConfigurationReferencesOutput struct{ *pulumi.OutputState }

func (VirtualnetworksPropertiesIpConfigurationReferencesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualnetworksPropertiesIpConfigurationReferences)(nil)).Elem()
}

func (o VirtualnetworksPropertiesIpConfigurationReferencesOutput) ToVirtualnetworksPropertiesIpConfigurationReferencesOutput() VirtualnetworksPropertiesIpConfigurationReferencesOutput {
	return o
}

func (o VirtualnetworksPropertiesIpConfigurationReferencesOutput) ToVirtualnetworksPropertiesIpConfigurationReferencesOutputWithContext(ctx context.Context) VirtualnetworksPropertiesIpConfigurationReferencesOutput {
	return o
}

func (o VirtualnetworksPropertiesIpConfigurationReferencesOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualnetworksPropertiesIpConfigurationReferences) *string { return v.Id }).(pulumi.StringPtrOutput)
}

type VirtualnetworksPropertiesIpConfigurationReferencesArrayOutput struct{ *pulumi.OutputState }

func (VirtualnetworksPropertiesIpConfigurationReferencesArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VirtualnetworksPropertiesIpConfigurationReferences)(nil)).Elem()
}

func (o VirtualnetworksPropertiesIpConfigurationReferencesArrayOutput) ToVirtualnetworksPropertiesIpConfigurationReferencesArrayOutput() VirtualnetworksPropertiesIpConfigurationReferencesArrayOutput {
	return o
}

func (o VirtualnetworksPropertiesIpConfigurationReferencesArrayOutput) ToVirtualnetworksPropertiesIpConfigurationReferencesArrayOutputWithContext(ctx context.Context) VirtualnetworksPropertiesIpConfigurationReferencesArrayOutput {
	return o
}

func (o VirtualnetworksPropertiesIpConfigurationReferencesArrayOutput) Index(i pulumi.IntInput) VirtualnetworksPropertiesIpConfigurationReferencesOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VirtualnetworksPropertiesIpConfigurationReferences {
		return vs[0].([]VirtualnetworksPropertiesIpConfigurationReferences)[vs[1].(int)]
	}).(VirtualnetworksPropertiesIpConfigurationReferencesOutput)
}

type VirtualnetworksPropertiesResponseIpConfigurationReferences struct {
	Id *string `pulumi:"id"`
}

type VirtualnetworksPropertiesResponseIpConfigurationReferencesOutput struct{ *pulumi.OutputState }

func (VirtualnetworksPropertiesResponseIpConfigurationReferencesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualnetworksPropertiesResponseIpConfigurationReferences)(nil)).Elem()
}

func (o VirtualnetworksPropertiesResponseIpConfigurationReferencesOutput) ToVirtualnetworksPropertiesResponseIpConfigurationReferencesOutput() VirtualnetworksPropertiesResponseIpConfigurationReferencesOutput {
	return o
}

func (o VirtualnetworksPropertiesResponseIpConfigurationReferencesOutput) ToVirtualnetworksPropertiesResponseIpConfigurationReferencesOutputWithContext(ctx context.Context) VirtualnetworksPropertiesResponseIpConfigurationReferencesOutput {
	return o
}

func (o VirtualnetworksPropertiesResponseIpConfigurationReferencesOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualnetworksPropertiesResponseIpConfigurationReferences) *string { return v.Id }).(pulumi.StringPtrOutput)
}

type VirtualnetworksPropertiesResponseIpConfigurationReferencesArrayOutput struct{ *pulumi.OutputState }

func (VirtualnetworksPropertiesResponseIpConfigurationReferencesArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VirtualnetworksPropertiesResponseIpConfigurationReferences)(nil)).Elem()
}

func (o VirtualnetworksPropertiesResponseIpConfigurationReferencesArrayOutput) ToVirtualnetworksPropertiesResponseIpConfigurationReferencesArrayOutput() VirtualnetworksPropertiesResponseIpConfigurationReferencesArrayOutput {
	return o
}

func (o VirtualnetworksPropertiesResponseIpConfigurationReferencesArrayOutput) ToVirtualnetworksPropertiesResponseIpConfigurationReferencesArrayOutputWithContext(ctx context.Context) VirtualnetworksPropertiesResponseIpConfigurationReferencesArrayOutput {
	return o
}

func (o VirtualnetworksPropertiesResponseIpConfigurationReferencesArrayOutput) Index(i pulumi.IntInput) VirtualnetworksPropertiesResponseIpConfigurationReferencesOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VirtualnetworksPropertiesResponseIpConfigurationReferences {
		return vs[0].([]VirtualnetworksPropertiesResponseIpConfigurationReferences)[vs[1].(int)]
	}).(VirtualnetworksPropertiesResponseIpConfigurationReferencesOutput)
}

type VirtualnetworksPropertiesResponseRouteTable struct {
	Id     *string                                   `pulumi:"id"`
	Name   *string                                   `pulumi:"name"`
	Routes []VirtualnetworksPropertiesResponseRoutes `pulumi:"routes"`
	Type   *string                                   `pulumi:"type"`
}

type VirtualnetworksPropertiesResponseRouteTableOutput struct{ *pulumi.OutputState }

func (VirtualnetworksPropertiesResponseRouteTableOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualnetworksPropertiesResponseRouteTable)(nil)).Elem()
}

func (o VirtualnetworksPropertiesResponseRouteTableOutput) ToVirtualnetworksPropertiesResponseRouteTableOutput() VirtualnetworksPropertiesResponseRouteTableOutput {
	return o
}

func (o VirtualnetworksPropertiesResponseRouteTableOutput) ToVirtualnetworksPropertiesResponseRouteTableOutputWithContext(ctx context.Context) VirtualnetworksPropertiesResponseRouteTableOutput {
	return o
}

func (o VirtualnetworksPropertiesResponseRouteTableOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualnetworksPropertiesResponseRouteTable) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o VirtualnetworksPropertiesResponseRouteTableOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualnetworksPropertiesResponseRouteTable) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o VirtualnetworksPropertiesResponseRouteTableOutput) Routes() VirtualnetworksPropertiesResponseRoutesArrayOutput {
	return o.ApplyT(func(v VirtualnetworksPropertiesResponseRouteTable) []VirtualnetworksPropertiesResponseRoutes {
		return v.Routes
	}).(VirtualnetworksPropertiesResponseRoutesArrayOutput)
}

func (o VirtualnetworksPropertiesResponseRouteTableOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualnetworksPropertiesResponseRouteTable) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type VirtualnetworksPropertiesResponseRouteTablePtrOutput struct{ *pulumi.OutputState }

func (VirtualnetworksPropertiesResponseRouteTablePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualnetworksPropertiesResponseRouteTable)(nil)).Elem()
}

func (o VirtualnetworksPropertiesResponseRouteTablePtrOutput) ToVirtualnetworksPropertiesResponseRouteTablePtrOutput() VirtualnetworksPropertiesResponseRouteTablePtrOutput {
	return o
}

func (o VirtualnetworksPropertiesResponseRouteTablePtrOutput) ToVirtualnetworksPropertiesResponseRouteTablePtrOutputWithContext(ctx context.Context) VirtualnetworksPropertiesResponseRouteTablePtrOutput {
	return o
}

func (o VirtualnetworksPropertiesResponseRouteTablePtrOutput) Elem() VirtualnetworksPropertiesResponseRouteTableOutput {
	return o.ApplyT(func(v *VirtualnetworksPropertiesResponseRouteTable) VirtualnetworksPropertiesResponseRouteTable {
		if v != nil {
			return *v
		}
		var ret VirtualnetworksPropertiesResponseRouteTable
		return ret
	}).(VirtualnetworksPropertiesResponseRouteTableOutput)
}

func (o VirtualnetworksPropertiesResponseRouteTablePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualnetworksPropertiesResponseRouteTable) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

func (o VirtualnetworksPropertiesResponseRouteTablePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualnetworksPropertiesResponseRouteTable) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o VirtualnetworksPropertiesResponseRouteTablePtrOutput) Routes() VirtualnetworksPropertiesResponseRoutesArrayOutput {
	return o.ApplyT(func(v *VirtualnetworksPropertiesResponseRouteTable) []VirtualnetworksPropertiesResponseRoutes {
		if v == nil {
			return nil
		}
		return v.Routes
	}).(VirtualnetworksPropertiesResponseRoutesArrayOutput)
}

func (o VirtualnetworksPropertiesResponseRouteTablePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualnetworksPropertiesResponseRouteTable) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

type VirtualnetworksPropertiesResponseRoutes struct {
	AddressPrefix    *string `pulumi:"addressPrefix"`
	Name             *string `pulumi:"name"`
	NextHopIpAddress *string `pulumi:"nextHopIpAddress"`
}

type VirtualnetworksPropertiesResponseRoutesOutput struct{ *pulumi.OutputState }

func (VirtualnetworksPropertiesResponseRoutesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualnetworksPropertiesResponseRoutes)(nil)).Elem()
}

func (o VirtualnetworksPropertiesResponseRoutesOutput) ToVirtualnetworksPropertiesResponseRoutesOutput() VirtualnetworksPropertiesResponseRoutesOutput {
	return o
}

func (o VirtualnetworksPropertiesResponseRoutesOutput) ToVirtualnetworksPropertiesResponseRoutesOutputWithContext(ctx context.Context) VirtualnetworksPropertiesResponseRoutesOutput {
	return o
}

func (o VirtualnetworksPropertiesResponseRoutesOutput) AddressPrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualnetworksPropertiesResponseRoutes) *string { return v.AddressPrefix }).(pulumi.StringPtrOutput)
}

func (o VirtualnetworksPropertiesResponseRoutesOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualnetworksPropertiesResponseRoutes) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o VirtualnetworksPropertiesResponseRoutesOutput) NextHopIpAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualnetworksPropertiesResponseRoutes) *string { return v.NextHopIpAddress }).(pulumi.StringPtrOutput)
}

type VirtualnetworksPropertiesResponseRoutesArrayOutput struct{ *pulumi.OutputState }

func (VirtualnetworksPropertiesResponseRoutesArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VirtualnetworksPropertiesResponseRoutes)(nil)).Elem()
}

func (o VirtualnetworksPropertiesResponseRoutesArrayOutput) ToVirtualnetworksPropertiesResponseRoutesArrayOutput() VirtualnetworksPropertiesResponseRoutesArrayOutput {
	return o
}

func (o VirtualnetworksPropertiesResponseRoutesArrayOutput) ToVirtualnetworksPropertiesResponseRoutesArrayOutputWithContext(ctx context.Context) VirtualnetworksPropertiesResponseRoutesArrayOutput {
	return o
}

func (o VirtualnetworksPropertiesResponseRoutesArrayOutput) Index(i pulumi.IntInput) VirtualnetworksPropertiesResponseRoutesOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VirtualnetworksPropertiesResponseRoutes {
		return vs[0].([]VirtualnetworksPropertiesResponseRoutes)[vs[1].(int)]
	}).(VirtualnetworksPropertiesResponseRoutesOutput)
}

type VirtualnetworksPropertiesResponseSubnets struct {
	AddressPrefix             *string                                                      `pulumi:"addressPrefix"`
	AddressPrefixes           []string                                                     `pulumi:"addressPrefixes"`
	IpAllocationMethod        *string                                                      `pulumi:"ipAllocationMethod"`
	IpConfigurationReferences []VirtualnetworksPropertiesResponseIpConfigurationReferences `pulumi:"ipConfigurationReferences"`
	IpPools                   []IPPoolResponse                                             `pulumi:"ipPools"`
	Name                      *string                                                      `pulumi:"name"`
	RouteTable                *VirtualnetworksPropertiesResponseRouteTable                 `pulumi:"routeTable"`
	Vlan                      *int                                                         `pulumi:"vlan"`
}

type VirtualnetworksPropertiesResponseSubnetsOutput struct{ *pulumi.OutputState }

func (VirtualnetworksPropertiesResponseSubnetsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualnetworksPropertiesResponseSubnets)(nil)).Elem()
}

func (o VirtualnetworksPropertiesResponseSubnetsOutput) ToVirtualnetworksPropertiesResponseSubnetsOutput() VirtualnetworksPropertiesResponseSubnetsOutput {
	return o
}

func (o VirtualnetworksPropertiesResponseSubnetsOutput) ToVirtualnetworksPropertiesResponseSubnetsOutputWithContext(ctx context.Context) VirtualnetworksPropertiesResponseSubnetsOutput {
	return o
}

func (o VirtualnetworksPropertiesResponseSubnetsOutput) AddressPrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualnetworksPropertiesResponseSubnets) *string { return v.AddressPrefix }).(pulumi.StringPtrOutput)
}

func (o VirtualnetworksPropertiesResponseSubnetsOutput) AddressPrefixes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v VirtualnetworksPropertiesResponseSubnets) []string { return v.AddressPrefixes }).(pulumi.StringArrayOutput)
}

func (o VirtualnetworksPropertiesResponseSubnetsOutput) IpAllocationMethod() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualnetworksPropertiesResponseSubnets) *string { return v.IpAllocationMethod }).(pulumi.StringPtrOutput)
}

func (o VirtualnetworksPropertiesResponseSubnetsOutput) IpConfigurationReferences() VirtualnetworksPropertiesResponseIpConfigurationReferencesArrayOutput {
	return o.ApplyT(func(v VirtualnetworksPropertiesResponseSubnets) []VirtualnetworksPropertiesResponseIpConfigurationReferences {
		return v.IpConfigurationReferences
	}).(VirtualnetworksPropertiesResponseIpConfigurationReferencesArrayOutput)
}

func (o VirtualnetworksPropertiesResponseSubnetsOutput) IpPools() IPPoolResponseArrayOutput {
	return o.ApplyT(func(v VirtualnetworksPropertiesResponseSubnets) []IPPoolResponse { return v.IpPools }).(IPPoolResponseArrayOutput)
}

func (o VirtualnetworksPropertiesResponseSubnetsOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualnetworksPropertiesResponseSubnets) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o VirtualnetworksPropertiesResponseSubnetsOutput) RouteTable() VirtualnetworksPropertiesResponseRouteTablePtrOutput {
	return o.ApplyT(func(v VirtualnetworksPropertiesResponseSubnets) *VirtualnetworksPropertiesResponseRouteTable {
		return v.RouteTable
	}).(VirtualnetworksPropertiesResponseRouteTablePtrOutput)
}

func (o VirtualnetworksPropertiesResponseSubnetsOutput) Vlan() pulumi.IntPtrOutput {
	return o.ApplyT(func(v VirtualnetworksPropertiesResponseSubnets) *int { return v.Vlan }).(pulumi.IntPtrOutput)
}

type VirtualnetworksPropertiesResponseSubnetsArrayOutput struct{ *pulumi.OutputState }

func (VirtualnetworksPropertiesResponseSubnetsArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VirtualnetworksPropertiesResponseSubnets)(nil)).Elem()
}

func (o VirtualnetworksPropertiesResponseSubnetsArrayOutput) ToVirtualnetworksPropertiesResponseSubnetsArrayOutput() VirtualnetworksPropertiesResponseSubnetsArrayOutput {
	return o
}

func (o VirtualnetworksPropertiesResponseSubnetsArrayOutput) ToVirtualnetworksPropertiesResponseSubnetsArrayOutputWithContext(ctx context.Context) VirtualnetworksPropertiesResponseSubnetsArrayOutput {
	return o
}

func (o VirtualnetworksPropertiesResponseSubnetsArrayOutput) Index(i pulumi.IntInput) VirtualnetworksPropertiesResponseSubnetsOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VirtualnetworksPropertiesResponseSubnets {
		return vs[0].([]VirtualnetworksPropertiesResponseSubnets)[vs[1].(int)]
	}).(VirtualnetworksPropertiesResponseSubnetsOutput)
}

type VirtualnetworksPropertiesRouteTable struct {
	Id     *string                           `pulumi:"id"`
	Name   *string                           `pulumi:"name"`
	Routes []VirtualnetworksPropertiesRoutes `pulumi:"routes"`
	Type   *string                           `pulumi:"type"`
}





type VirtualnetworksPropertiesRouteTableInput interface {
	pulumi.Input

	ToVirtualnetworksPropertiesRouteTableOutput() VirtualnetworksPropertiesRouteTableOutput
	ToVirtualnetworksPropertiesRouteTableOutputWithContext(context.Context) VirtualnetworksPropertiesRouteTableOutput
}

type VirtualnetworksPropertiesRouteTableArgs struct {
	Id     pulumi.StringPtrInput                     `pulumi:"id"`
	Name   pulumi.StringPtrInput                     `pulumi:"name"`
	Routes VirtualnetworksPropertiesRoutesArrayInput `pulumi:"routes"`
	Type   pulumi.StringPtrInput                     `pulumi:"type"`
}

func (VirtualnetworksPropertiesRouteTableArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualnetworksPropertiesRouteTable)(nil)).Elem()
}

func (i VirtualnetworksPropertiesRouteTableArgs) ToVirtualnetworksPropertiesRouteTableOutput() VirtualnetworksPropertiesRouteTableOutput {
	return i.ToVirtualnetworksPropertiesRouteTableOutputWithContext(context.Background())
}

func (i VirtualnetworksPropertiesRouteTableArgs) ToVirtualnetworksPropertiesRouteTableOutputWithContext(ctx context.Context) VirtualnetworksPropertiesRouteTableOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualnetworksPropertiesRouteTableOutput)
}

func (i VirtualnetworksPropertiesRouteTableArgs) ToVirtualnetworksPropertiesRouteTablePtrOutput() VirtualnetworksPropertiesRouteTablePtrOutput {
	return i.ToVirtualnetworksPropertiesRouteTablePtrOutputWithContext(context.Background())
}

func (i VirtualnetworksPropertiesRouteTableArgs) ToVirtualnetworksPropertiesRouteTablePtrOutputWithContext(ctx context.Context) VirtualnetworksPropertiesRouteTablePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualnetworksPropertiesRouteTableOutput).ToVirtualnetworksPropertiesRouteTablePtrOutputWithContext(ctx)
}









type VirtualnetworksPropertiesRouteTablePtrInput interface {
	pulumi.Input

	ToVirtualnetworksPropertiesRouteTablePtrOutput() VirtualnetworksPropertiesRouteTablePtrOutput
	ToVirtualnetworksPropertiesRouteTablePtrOutputWithContext(context.Context) VirtualnetworksPropertiesRouteTablePtrOutput
}

type virtualnetworksPropertiesRouteTablePtrType VirtualnetworksPropertiesRouteTableArgs

func VirtualnetworksPropertiesRouteTablePtr(v *VirtualnetworksPropertiesRouteTableArgs) VirtualnetworksPropertiesRouteTablePtrInput {
	return (*virtualnetworksPropertiesRouteTablePtrType)(v)
}

func (*virtualnetworksPropertiesRouteTablePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualnetworksPropertiesRouteTable)(nil)).Elem()
}

func (i *virtualnetworksPropertiesRouteTablePtrType) ToVirtualnetworksPropertiesRouteTablePtrOutput() VirtualnetworksPropertiesRouteTablePtrOutput {
	return i.ToVirtualnetworksPropertiesRouteTablePtrOutputWithContext(context.Background())
}

func (i *virtualnetworksPropertiesRouteTablePtrType) ToVirtualnetworksPropertiesRouteTablePtrOutputWithContext(ctx context.Context) VirtualnetworksPropertiesRouteTablePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualnetworksPropertiesRouteTablePtrOutput)
}

type VirtualnetworksPropertiesRouteTableOutput struct{ *pulumi.OutputState }

func (VirtualnetworksPropertiesRouteTableOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualnetworksPropertiesRouteTable)(nil)).Elem()
}

func (o VirtualnetworksPropertiesRouteTableOutput) ToVirtualnetworksPropertiesRouteTableOutput() VirtualnetworksPropertiesRouteTableOutput {
	return o
}

func (o VirtualnetworksPropertiesRouteTableOutput) ToVirtualnetworksPropertiesRouteTableOutputWithContext(ctx context.Context) VirtualnetworksPropertiesRouteTableOutput {
	return o
}

func (o VirtualnetworksPropertiesRouteTableOutput) ToVirtualnetworksPropertiesRouteTablePtrOutput() VirtualnetworksPropertiesRouteTablePtrOutput {
	return o.ToVirtualnetworksPropertiesRouteTablePtrOutputWithContext(context.Background())
}

func (o VirtualnetworksPropertiesRouteTableOutput) ToVirtualnetworksPropertiesRouteTablePtrOutputWithContext(ctx context.Context) VirtualnetworksPropertiesRouteTablePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v VirtualnetworksPropertiesRouteTable) *VirtualnetworksPropertiesRouteTable {
		return &v
	}).(VirtualnetworksPropertiesRouteTablePtrOutput)
}

func (o VirtualnetworksPropertiesRouteTableOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualnetworksPropertiesRouteTable) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o VirtualnetworksPropertiesRouteTableOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualnetworksPropertiesRouteTable) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o VirtualnetworksPropertiesRouteTableOutput) Routes() VirtualnetworksPropertiesRoutesArrayOutput {
	return o.ApplyT(func(v VirtualnetworksPropertiesRouteTable) []VirtualnetworksPropertiesRoutes { return v.Routes }).(VirtualnetworksPropertiesRoutesArrayOutput)
}

func (o VirtualnetworksPropertiesRouteTableOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualnetworksPropertiesRouteTable) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type VirtualnetworksPropertiesRouteTablePtrOutput struct{ *pulumi.OutputState }

func (VirtualnetworksPropertiesRouteTablePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualnetworksPropertiesRouteTable)(nil)).Elem()
}

func (o VirtualnetworksPropertiesRouteTablePtrOutput) ToVirtualnetworksPropertiesRouteTablePtrOutput() VirtualnetworksPropertiesRouteTablePtrOutput {
	return o
}

func (o VirtualnetworksPropertiesRouteTablePtrOutput) ToVirtualnetworksPropertiesRouteTablePtrOutputWithContext(ctx context.Context) VirtualnetworksPropertiesRouteTablePtrOutput {
	return o
}

func (o VirtualnetworksPropertiesRouteTablePtrOutput) Elem() VirtualnetworksPropertiesRouteTableOutput {
	return o.ApplyT(func(v *VirtualnetworksPropertiesRouteTable) VirtualnetworksPropertiesRouteTable {
		if v != nil {
			return *v
		}
		var ret VirtualnetworksPropertiesRouteTable
		return ret
	}).(VirtualnetworksPropertiesRouteTableOutput)
}

func (o VirtualnetworksPropertiesRouteTablePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualnetworksPropertiesRouteTable) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

func (o VirtualnetworksPropertiesRouteTablePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualnetworksPropertiesRouteTable) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o VirtualnetworksPropertiesRouteTablePtrOutput) Routes() VirtualnetworksPropertiesRoutesArrayOutput {
	return o.ApplyT(func(v *VirtualnetworksPropertiesRouteTable) []VirtualnetworksPropertiesRoutes {
		if v == nil {
			return nil
		}
		return v.Routes
	}).(VirtualnetworksPropertiesRoutesArrayOutput)
}

func (o VirtualnetworksPropertiesRouteTablePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualnetworksPropertiesRouteTable) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

type VirtualnetworksPropertiesRoutes struct {
	AddressPrefix    *string `pulumi:"addressPrefix"`
	Name             *string `pulumi:"name"`
	NextHopIpAddress *string `pulumi:"nextHopIpAddress"`
}





type VirtualnetworksPropertiesRoutesInput interface {
	pulumi.Input

	ToVirtualnetworksPropertiesRoutesOutput() VirtualnetworksPropertiesRoutesOutput
	ToVirtualnetworksPropertiesRoutesOutputWithContext(context.Context) VirtualnetworksPropertiesRoutesOutput
}

type VirtualnetworksPropertiesRoutesArgs struct {
	AddressPrefix    pulumi.StringPtrInput `pulumi:"addressPrefix"`
	Name             pulumi.StringPtrInput `pulumi:"name"`
	NextHopIpAddress pulumi.StringPtrInput `pulumi:"nextHopIpAddress"`
}

func (VirtualnetworksPropertiesRoutesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualnetworksPropertiesRoutes)(nil)).Elem()
}

func (i VirtualnetworksPropertiesRoutesArgs) ToVirtualnetworksPropertiesRoutesOutput() VirtualnetworksPropertiesRoutesOutput {
	return i.ToVirtualnetworksPropertiesRoutesOutputWithContext(context.Background())
}

func (i VirtualnetworksPropertiesRoutesArgs) ToVirtualnetworksPropertiesRoutesOutputWithContext(ctx context.Context) VirtualnetworksPropertiesRoutesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualnetworksPropertiesRoutesOutput)
}





type VirtualnetworksPropertiesRoutesArrayInput interface {
	pulumi.Input

	ToVirtualnetworksPropertiesRoutesArrayOutput() VirtualnetworksPropertiesRoutesArrayOutput
	ToVirtualnetworksPropertiesRoutesArrayOutputWithContext(context.Context) VirtualnetworksPropertiesRoutesArrayOutput
}

type VirtualnetworksPropertiesRoutesArray []VirtualnetworksPropertiesRoutesInput

func (VirtualnetworksPropertiesRoutesArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VirtualnetworksPropertiesRoutes)(nil)).Elem()
}

func (i VirtualnetworksPropertiesRoutesArray) ToVirtualnetworksPropertiesRoutesArrayOutput() VirtualnetworksPropertiesRoutesArrayOutput {
	return i.ToVirtualnetworksPropertiesRoutesArrayOutputWithContext(context.Background())
}

func (i VirtualnetworksPropertiesRoutesArray) ToVirtualnetworksPropertiesRoutesArrayOutputWithContext(ctx context.Context) VirtualnetworksPropertiesRoutesArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualnetworksPropertiesRoutesArrayOutput)
}

type VirtualnetworksPropertiesRoutesOutput struct{ *pulumi.OutputState }

func (VirtualnetworksPropertiesRoutesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualnetworksPropertiesRoutes)(nil)).Elem()
}

func (o VirtualnetworksPropertiesRoutesOutput) ToVirtualnetworksPropertiesRoutesOutput() VirtualnetworksPropertiesRoutesOutput {
	return o
}

func (o VirtualnetworksPropertiesRoutesOutput) ToVirtualnetworksPropertiesRoutesOutputWithContext(ctx context.Context) VirtualnetworksPropertiesRoutesOutput {
	return o
}

func (o VirtualnetworksPropertiesRoutesOutput) AddressPrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualnetworksPropertiesRoutes) *string { return v.AddressPrefix }).(pulumi.StringPtrOutput)
}

func (o VirtualnetworksPropertiesRoutesOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualnetworksPropertiesRoutes) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o VirtualnetworksPropertiesRoutesOutput) NextHopIpAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualnetworksPropertiesRoutes) *string { return v.NextHopIpAddress }).(pulumi.StringPtrOutput)
}

type VirtualnetworksPropertiesRoutesArrayOutput struct{ *pulumi.OutputState }

func (VirtualnetworksPropertiesRoutesArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VirtualnetworksPropertiesRoutes)(nil)).Elem()
}

func (o VirtualnetworksPropertiesRoutesArrayOutput) ToVirtualnetworksPropertiesRoutesArrayOutput() VirtualnetworksPropertiesRoutesArrayOutput {
	return o
}

func (o VirtualnetworksPropertiesRoutesArrayOutput) ToVirtualnetworksPropertiesRoutesArrayOutputWithContext(ctx context.Context) VirtualnetworksPropertiesRoutesArrayOutput {
	return o
}

func (o VirtualnetworksPropertiesRoutesArrayOutput) Index(i pulumi.IntInput) VirtualnetworksPropertiesRoutesOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VirtualnetworksPropertiesRoutes {
		return vs[0].([]VirtualnetworksPropertiesRoutes)[vs[1].(int)]
	}).(VirtualnetworksPropertiesRoutesOutput)
}

type VirtualnetworksPropertiesSubnets struct {
	AddressPrefix             *string                                              `pulumi:"addressPrefix"`
	AddressPrefixes           []string                                             `pulumi:"addressPrefixes"`
	IpAllocationMethod        *string                                              `pulumi:"ipAllocationMethod"`
	IpConfigurationReferences []VirtualnetworksPropertiesIpConfigurationReferences `pulumi:"ipConfigurationReferences"`
	IpPools                   []IPPool                                             `pulumi:"ipPools"`
	Name                      *string                                              `pulumi:"name"`
	RouteTable                *VirtualnetworksPropertiesRouteTable                 `pulumi:"routeTable"`
	Vlan                      *int                                                 `pulumi:"vlan"`
}





type VirtualnetworksPropertiesSubnetsInput interface {
	pulumi.Input

	ToVirtualnetworksPropertiesSubnetsOutput() VirtualnetworksPropertiesSubnetsOutput
	ToVirtualnetworksPropertiesSubnetsOutputWithContext(context.Context) VirtualnetworksPropertiesSubnetsOutput
}

type VirtualnetworksPropertiesSubnetsArgs struct {
	AddressPrefix             pulumi.StringPtrInput                                        `pulumi:"addressPrefix"`
	AddressPrefixes           pulumi.StringArrayInput                                      `pulumi:"addressPrefixes"`
	IpAllocationMethod        pulumi.StringPtrInput                                        `pulumi:"ipAllocationMethod"`
	IpConfigurationReferences VirtualnetworksPropertiesIpConfigurationReferencesArrayInput `pulumi:"ipConfigurationReferences"`
	IpPools                   IPPoolArrayInput                                             `pulumi:"ipPools"`
	Name                      pulumi.StringPtrInput                                        `pulumi:"name"`
	RouteTable                VirtualnetworksPropertiesRouteTablePtrInput                  `pulumi:"routeTable"`
	Vlan                      pulumi.IntPtrInput                                           `pulumi:"vlan"`
}

func (VirtualnetworksPropertiesSubnetsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualnetworksPropertiesSubnets)(nil)).Elem()
}

func (i VirtualnetworksPropertiesSubnetsArgs) ToVirtualnetworksPropertiesSubnetsOutput() VirtualnetworksPropertiesSubnetsOutput {
	return i.ToVirtualnetworksPropertiesSubnetsOutputWithContext(context.Background())
}

func (i VirtualnetworksPropertiesSubnetsArgs) ToVirtualnetworksPropertiesSubnetsOutputWithContext(ctx context.Context) VirtualnetworksPropertiesSubnetsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualnetworksPropertiesSubnetsOutput)
}





type VirtualnetworksPropertiesSubnetsArrayInput interface {
	pulumi.Input

	ToVirtualnetworksPropertiesSubnetsArrayOutput() VirtualnetworksPropertiesSubnetsArrayOutput
	ToVirtualnetworksPropertiesSubnetsArrayOutputWithContext(context.Context) VirtualnetworksPropertiesSubnetsArrayOutput
}

type VirtualnetworksPropertiesSubnetsArray []VirtualnetworksPropertiesSubnetsInput

func (VirtualnetworksPropertiesSubnetsArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VirtualnetworksPropertiesSubnets)(nil)).Elem()
}

func (i VirtualnetworksPropertiesSubnetsArray) ToVirtualnetworksPropertiesSubnetsArrayOutput() VirtualnetworksPropertiesSubnetsArrayOutput {
	return i.ToVirtualnetworksPropertiesSubnetsArrayOutputWithContext(context.Background())
}

func (i VirtualnetworksPropertiesSubnetsArray) ToVirtualnetworksPropertiesSubnetsArrayOutputWithContext(ctx context.Context) VirtualnetworksPropertiesSubnetsArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualnetworksPropertiesSubnetsArrayOutput)
}

type VirtualnetworksPropertiesSubnetsOutput struct{ *pulumi.OutputState }

func (VirtualnetworksPropertiesSubnetsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualnetworksPropertiesSubnets)(nil)).Elem()
}

func (o VirtualnetworksPropertiesSubnetsOutput) ToVirtualnetworksPropertiesSubnetsOutput() VirtualnetworksPropertiesSubnetsOutput {
	return o
}

func (o VirtualnetworksPropertiesSubnetsOutput) ToVirtualnetworksPropertiesSubnetsOutputWithContext(ctx context.Context) VirtualnetworksPropertiesSubnetsOutput {
	return o
}

func (o VirtualnetworksPropertiesSubnetsOutput) AddressPrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualnetworksPropertiesSubnets) *string { return v.AddressPrefix }).(pulumi.StringPtrOutput)
}

func (o VirtualnetworksPropertiesSubnetsOutput) AddressPrefixes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v VirtualnetworksPropertiesSubnets) []string { return v.AddressPrefixes }).(pulumi.StringArrayOutput)
}

func (o VirtualnetworksPropertiesSubnetsOutput) IpAllocationMethod() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualnetworksPropertiesSubnets) *string { return v.IpAllocationMethod }).(pulumi.StringPtrOutput)
}

func (o VirtualnetworksPropertiesSubnetsOutput) IpConfigurationReferences() VirtualnetworksPropertiesIpConfigurationReferencesArrayOutput {
	return o.ApplyT(func(v VirtualnetworksPropertiesSubnets) []VirtualnetworksPropertiesIpConfigurationReferences {
		return v.IpConfigurationReferences
	}).(VirtualnetworksPropertiesIpConfigurationReferencesArrayOutput)
}

func (o VirtualnetworksPropertiesSubnetsOutput) IpPools() IPPoolArrayOutput {
	return o.ApplyT(func(v VirtualnetworksPropertiesSubnets) []IPPool { return v.IpPools }).(IPPoolArrayOutput)
}

func (o VirtualnetworksPropertiesSubnetsOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualnetworksPropertiesSubnets) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o VirtualnetworksPropertiesSubnetsOutput) RouteTable() VirtualnetworksPropertiesRouteTablePtrOutput {
	return o.ApplyT(func(v VirtualnetworksPropertiesSubnets) *VirtualnetworksPropertiesRouteTable { return v.RouteTable }).(VirtualnetworksPropertiesRouteTablePtrOutput)
}

func (o VirtualnetworksPropertiesSubnetsOutput) Vlan() pulumi.IntPtrOutput {
	return o.ApplyT(func(v VirtualnetworksPropertiesSubnets) *int { return v.Vlan }).(pulumi.IntPtrOutput)
}

type VirtualnetworksPropertiesSubnetsArrayOutput struct{ *pulumi.OutputState }

func (VirtualnetworksPropertiesSubnetsArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VirtualnetworksPropertiesSubnets)(nil)).Elem()
}

func (o VirtualnetworksPropertiesSubnetsArrayOutput) ToVirtualnetworksPropertiesSubnetsArrayOutput() VirtualnetworksPropertiesSubnetsArrayOutput {
	return o
}

func (o VirtualnetworksPropertiesSubnetsArrayOutput) ToVirtualnetworksPropertiesSubnetsArrayOutputWithContext(ctx context.Context) VirtualnetworksPropertiesSubnetsArrayOutput {
	return o
}

func (o VirtualnetworksPropertiesSubnetsArrayOutput) Index(i pulumi.IntInput) VirtualnetworksPropertiesSubnetsOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VirtualnetworksPropertiesSubnets {
		return vs[0].([]VirtualnetworksPropertiesSubnets)[vs[1].(int)]
	}).(VirtualnetworksPropertiesSubnetsOutput)
}

func init() {
	pulumi.RegisterOutputType(ClusterDesiredPropertiesOutput{})
	pulumi.RegisterOutputType(ClusterDesiredPropertiesPtrOutput{})
	pulumi.RegisterOutputType(ClusterDesiredPropertiesResponseOutput{})
	pulumi.RegisterOutputType(ClusterDesiredPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(ClusterNodeResponseOutput{})
	pulumi.RegisterOutputType(ClusterNodeResponseArrayOutput{})
	pulumi.RegisterOutputType(ClusterReportedPropertiesResponseOutput{})
	pulumi.RegisterOutputType(ErrorAdditionalInfoResponseOutput{})
	pulumi.RegisterOutputType(ErrorAdditionalInfoResponseArrayOutput{})
	pulumi.RegisterOutputType(ErrorDetailResponseOutput{})
	pulumi.RegisterOutputType(ErrorDetailResponseArrayOutput{})
	pulumi.RegisterOutputType(ExtendedLocationOutput{})
	pulumi.RegisterOutputType(ExtendedLocationPtrOutput{})
	pulumi.RegisterOutputType(ExtendedLocationResponseOutput{})
	pulumi.RegisterOutputType(ExtendedLocationResponsePtrOutput{})
	pulumi.RegisterOutputType(GalleryImageIdentifierOutput{})
	pulumi.RegisterOutputType(GalleryImageIdentifierPtrOutput{})
	pulumi.RegisterOutputType(GalleryImageIdentifierResponseOutput{})
	pulumi.RegisterOutputType(GalleryImageIdentifierResponsePtrOutput{})
	pulumi.RegisterOutputType(GalleryImageStatusResponseOutput{})
	pulumi.RegisterOutputType(GalleryImageStatusResponseDownloadStatusOutput{})
	pulumi.RegisterOutputType(GalleryImageStatusResponseDownloadStatusPtrOutput{})
	pulumi.RegisterOutputType(GalleryImageStatusResponseProvisioningStatusOutput{})
	pulumi.RegisterOutputType(GalleryImageStatusResponseProvisioningStatusPtrOutput{})
	pulumi.RegisterOutputType(GalleryImageVersionOutput{})
	pulumi.RegisterOutputType(GalleryImageVersionPtrOutput{})
	pulumi.RegisterOutputType(GalleryImageVersionResponseOutput{})
	pulumi.RegisterOutputType(GalleryImageVersionResponsePtrOutput{})
	pulumi.RegisterOutputType(GalleryImageVersionStorageProfileResponseOutput{})
	pulumi.RegisterOutputType(GalleryImageVersionStorageProfileResponsePtrOutput{})
	pulumi.RegisterOutputType(GalleryOSDiskImageResponseOutput{})
	pulumi.RegisterOutputType(GalleryOSDiskImageResponsePtrOutput{})
	pulumi.RegisterOutputType(GuestAgentProfileResponseOutput{})
	pulumi.RegisterOutputType(GuestAgentProfileResponsePtrOutput{})
	pulumi.RegisterOutputType(GuestCredentialOutput{})
	pulumi.RegisterOutputType(GuestCredentialPtrOutput{})
	pulumi.RegisterOutputType(GuestCredentialResponseOutput{})
	pulumi.RegisterOutputType(GuestCredentialResponsePtrOutput{})
	pulumi.RegisterOutputType(HttpProxyConfigurationOutput{})
	pulumi.RegisterOutputType(HttpProxyConfigurationPtrOutput{})
	pulumi.RegisterOutputType(HttpProxyConfigurationResponseOutput{})
	pulumi.RegisterOutputType(HttpProxyConfigurationResponsePtrOutput{})
	pulumi.RegisterOutputType(IPPoolOutput{})
	pulumi.RegisterOutputType(IPPoolArrayOutput{})
	pulumi.RegisterOutputType(IPPoolInfoResponseOutput{})
	pulumi.RegisterOutputType(IPPoolInfoResponsePtrOutput{})
	pulumi.RegisterOutputType(IPPoolResponseOutput{})
	pulumi.RegisterOutputType(IPPoolResponseArrayOutput{})
	pulumi.RegisterOutputType(IdentityOutput{})
	pulumi.RegisterOutputType(IdentityPtrOutput{})
	pulumi.RegisterOutputType(IdentityResponseOutput{})
	pulumi.RegisterOutputType(IdentityResponsePtrOutput{})
	pulumi.RegisterOutputType(InterfaceDNSSettingsOutput{})
	pulumi.RegisterOutputType(InterfaceDNSSettingsPtrOutput{})
	pulumi.RegisterOutputType(InterfaceDNSSettingsResponseOutput{})
	pulumi.RegisterOutputType(InterfaceDNSSettingsResponsePtrOutput{})
	pulumi.RegisterOutputType(IpConfigurationOutput{})
	pulumi.RegisterOutputType(IpConfigurationArrayOutput{})
	pulumi.RegisterOutputType(IpConfigurationPropertiesOutput{})
	pulumi.RegisterOutputType(IpConfigurationPropertiesPtrOutput{})
	pulumi.RegisterOutputType(IpConfigurationResponseOutput{})
	pulumi.RegisterOutputType(IpConfigurationResponseArrayOutput{})
	pulumi.RegisterOutputType(IpConfigurationResponsePropertiesOutput{})
	pulumi.RegisterOutputType(IpConfigurationResponsePropertiesPtrOutput{})
	pulumi.RegisterOutputType(IpConfigurationResponseSubnetOutput{})
	pulumi.RegisterOutputType(IpConfigurationResponseSubnetPtrOutput{})
	pulumi.RegisterOutputType(IpConfigurationSubnetOutput{})
	pulumi.RegisterOutputType(IpConfigurationSubnetPtrOutput{})
	pulumi.RegisterOutputType(MachineExtensionInstanceViewResponseStatusOutput{})
	pulumi.RegisterOutputType(MachineExtensionInstanceViewResponseStatusPtrOutput{})
	pulumi.RegisterOutputType(MachineExtensionPropertiesResponseInstanceViewOutput{})
	pulumi.RegisterOutputType(MachineExtensionPropertiesResponseInstanceViewPtrOutput{})
	pulumi.RegisterOutputType(MarketplaceGalleryImageStatusResponseOutput{})
	pulumi.RegisterOutputType(MarketplaceGalleryImageStatusResponseDownloadStatusOutput{})
	pulumi.RegisterOutputType(MarketplaceGalleryImageStatusResponseDownloadStatusPtrOutput{})
	pulumi.RegisterOutputType(MarketplaceGalleryImageStatusResponseProvisioningStatusOutput{})
	pulumi.RegisterOutputType(MarketplaceGalleryImageStatusResponseProvisioningStatusPtrOutput{})
	pulumi.RegisterOutputType(NetworkInterfaceStatusResponseOutput{})
	pulumi.RegisterOutputType(NetworkInterfaceStatusResponseProvisioningStatusOutput{})
	pulumi.RegisterOutputType(NetworkInterfaceStatusResponseProvisioningStatusPtrOutput{})
	pulumi.RegisterOutputType(PerNodeExtensionStateResponseOutput{})
	pulumi.RegisterOutputType(PerNodeExtensionStateResponseArrayOutput{})
	pulumi.RegisterOutputType(PerNodeStateResponseOutput{})
	pulumi.RegisterOutputType(PerNodeStateResponseArrayOutput{})
	pulumi.RegisterOutputType(StorageContainerStatusResponseOutput{})
	pulumi.RegisterOutputType(StorageContainerStatusResponseProvisioningStatusOutput{})
	pulumi.RegisterOutputType(StorageContainerStatusResponseProvisioningStatusPtrOutput{})
	pulumi.RegisterOutputType(StoragecontainersExtendedLocationOutput{})
	pulumi.RegisterOutputType(StoragecontainersExtendedLocationPtrOutput{})
	pulumi.RegisterOutputType(StoragecontainersResponseExtendedLocationOutput{})
	pulumi.RegisterOutputType(StoragecontainersResponseExtendedLocationPtrOutput{})
	pulumi.RegisterOutputType(SystemDataResponseOutput{})
	pulumi.RegisterOutputType(VirtualHardDiskStatusResponseOutput{})
	pulumi.RegisterOutputType(VirtualHardDiskStatusResponseProvisioningStatusOutput{})
	pulumi.RegisterOutputType(VirtualHardDiskStatusResponseProvisioningStatusPtrOutput{})
	pulumi.RegisterOutputType(VirtualMachineStatusResponseOutput{})
	pulumi.RegisterOutputType(VirtualMachineStatusResponseProvisioningStatusOutput{})
	pulumi.RegisterOutputType(VirtualMachineStatusResponseProvisioningStatusPtrOutput{})
	pulumi.RegisterOutputType(VirtualNetworkStatusResponseOutput{})
	pulumi.RegisterOutputType(VirtualNetworkStatusResponseProvisioningStatusOutput{})
	pulumi.RegisterOutputType(VirtualNetworkStatusResponseProvisioningStatusPtrOutput{})
	pulumi.RegisterOutputType(VirtualmachinesPropertiesDataDisksOutput{})
	pulumi.RegisterOutputType(VirtualmachinesPropertiesDataDisksArrayOutput{})
	pulumi.RegisterOutputType(VirtualmachinesPropertiesDynamicMemoryConfigOutput{})
	pulumi.RegisterOutputType(VirtualmachinesPropertiesDynamicMemoryConfigPtrOutput{})
	pulumi.RegisterOutputType(VirtualmachinesPropertiesHardwareProfileOutput{})
	pulumi.RegisterOutputType(VirtualmachinesPropertiesHardwareProfilePtrOutput{})
	pulumi.RegisterOutputType(VirtualmachinesPropertiesImageReferenceOutput{})
	pulumi.RegisterOutputType(VirtualmachinesPropertiesImageReferencePtrOutput{})
	pulumi.RegisterOutputType(VirtualmachinesPropertiesLinuxConfigurationOutput{})
	pulumi.RegisterOutputType(VirtualmachinesPropertiesLinuxConfigurationPtrOutput{})
	pulumi.RegisterOutputType(VirtualmachinesPropertiesNetworkInterfacesOutput{})
	pulumi.RegisterOutputType(VirtualmachinesPropertiesNetworkInterfacesArrayOutput{})
	pulumi.RegisterOutputType(VirtualmachinesPropertiesNetworkProfileOutput{})
	pulumi.RegisterOutputType(VirtualmachinesPropertiesNetworkProfilePtrOutput{})
	pulumi.RegisterOutputType(VirtualmachinesPropertiesOsDiskOutput{})
	pulumi.RegisterOutputType(VirtualmachinesPropertiesOsDiskPtrOutput{})
	pulumi.RegisterOutputType(VirtualmachinesPropertiesOsProfileOutput{})
	pulumi.RegisterOutputType(VirtualmachinesPropertiesOsProfilePtrOutput{})
	pulumi.RegisterOutputType(VirtualmachinesPropertiesPublicKeysOutput{})
	pulumi.RegisterOutputType(VirtualmachinesPropertiesPublicKeysArrayOutput{})
	pulumi.RegisterOutputType(VirtualmachinesPropertiesPublicKeysPublicKeysOutput{})
	pulumi.RegisterOutputType(VirtualmachinesPropertiesPublicKeysPublicKeysArrayOutput{})
	pulumi.RegisterOutputType(VirtualmachinesPropertiesResponseDataDisksOutput{})
	pulumi.RegisterOutputType(VirtualmachinesPropertiesResponseDataDisksArrayOutput{})
	pulumi.RegisterOutputType(VirtualmachinesPropertiesResponseDynamicMemoryConfigOutput{})
	pulumi.RegisterOutputType(VirtualmachinesPropertiesResponseDynamicMemoryConfigPtrOutput{})
	pulumi.RegisterOutputType(VirtualmachinesPropertiesResponseHardwareProfileOutput{})
	pulumi.RegisterOutputType(VirtualmachinesPropertiesResponseHardwareProfilePtrOutput{})
	pulumi.RegisterOutputType(VirtualmachinesPropertiesResponseImageReferenceOutput{})
	pulumi.RegisterOutputType(VirtualmachinesPropertiesResponseImageReferencePtrOutput{})
	pulumi.RegisterOutputType(VirtualmachinesPropertiesResponseLinuxConfigurationOutput{})
	pulumi.RegisterOutputType(VirtualmachinesPropertiesResponseLinuxConfigurationPtrOutput{})
	pulumi.RegisterOutputType(VirtualmachinesPropertiesResponseNetworkInterfacesOutput{})
	pulumi.RegisterOutputType(VirtualmachinesPropertiesResponseNetworkInterfacesArrayOutput{})
	pulumi.RegisterOutputType(VirtualmachinesPropertiesResponseNetworkProfileOutput{})
	pulumi.RegisterOutputType(VirtualmachinesPropertiesResponseNetworkProfilePtrOutput{})
	pulumi.RegisterOutputType(VirtualmachinesPropertiesResponseOsDiskOutput{})
	pulumi.RegisterOutputType(VirtualmachinesPropertiesResponseOsDiskPtrOutput{})
	pulumi.RegisterOutputType(VirtualmachinesPropertiesResponseOsProfileOutput{})
	pulumi.RegisterOutputType(VirtualmachinesPropertiesResponseOsProfilePtrOutput{})
	pulumi.RegisterOutputType(VirtualmachinesPropertiesResponsePublicKeysOutput{})
	pulumi.RegisterOutputType(VirtualmachinesPropertiesResponsePublicKeysArrayOutput{})
	pulumi.RegisterOutputType(VirtualmachinesPropertiesResponsePublicKeysPublicKeysOutput{})
	pulumi.RegisterOutputType(VirtualmachinesPropertiesResponsePublicKeysPublicKeysArrayOutput{})
	pulumi.RegisterOutputType(VirtualmachinesPropertiesResponseSecurityProfileOutput{})
	pulumi.RegisterOutputType(VirtualmachinesPropertiesResponseSecurityProfilePtrOutput{})
	pulumi.RegisterOutputType(VirtualmachinesPropertiesResponseSshOutput{})
	pulumi.RegisterOutputType(VirtualmachinesPropertiesResponseSshPtrOutput{})
	pulumi.RegisterOutputType(VirtualmachinesPropertiesResponseSshSshOutput{})
	pulumi.RegisterOutputType(VirtualmachinesPropertiesResponseSshSshPtrOutput{})
	pulumi.RegisterOutputType(VirtualmachinesPropertiesResponseStorageProfileOutput{})
	pulumi.RegisterOutputType(VirtualmachinesPropertiesResponseStorageProfilePtrOutput{})
	pulumi.RegisterOutputType(VirtualmachinesPropertiesResponseUefiSettingsOutput{})
	pulumi.RegisterOutputType(VirtualmachinesPropertiesResponseUefiSettingsPtrOutput{})
	pulumi.RegisterOutputType(VirtualmachinesPropertiesResponseWindowsConfigurationOutput{})
	pulumi.RegisterOutputType(VirtualmachinesPropertiesResponseWindowsConfigurationPtrOutput{})
	pulumi.RegisterOutputType(VirtualmachinesPropertiesSecurityProfileOutput{})
	pulumi.RegisterOutputType(VirtualmachinesPropertiesSecurityProfilePtrOutput{})
	pulumi.RegisterOutputType(VirtualmachinesPropertiesSshOutput{})
	pulumi.RegisterOutputType(VirtualmachinesPropertiesSshPtrOutput{})
	pulumi.RegisterOutputType(VirtualmachinesPropertiesSshSshOutput{})
	pulumi.RegisterOutputType(VirtualmachinesPropertiesSshSshPtrOutput{})
	pulumi.RegisterOutputType(VirtualmachinesPropertiesStorageProfileOutput{})
	pulumi.RegisterOutputType(VirtualmachinesPropertiesStorageProfilePtrOutput{})
	pulumi.RegisterOutputType(VirtualmachinesPropertiesUefiSettingsOutput{})
	pulumi.RegisterOutputType(VirtualmachinesPropertiesUefiSettingsPtrOutput{})
	pulumi.RegisterOutputType(VirtualmachinesPropertiesWindowsConfigurationOutput{})
	pulumi.RegisterOutputType(VirtualmachinesPropertiesWindowsConfigurationPtrOutput{})
	pulumi.RegisterOutputType(VirtualnetworksPropertiesIpConfigurationReferencesOutput{})
	pulumi.RegisterOutputType(VirtualnetworksPropertiesIpConfigurationReferencesArrayOutput{})
	pulumi.RegisterOutputType(VirtualnetworksPropertiesResponseIpConfigurationReferencesOutput{})
	pulumi.RegisterOutputType(VirtualnetworksPropertiesResponseIpConfigurationReferencesArrayOutput{})
	pulumi.RegisterOutputType(VirtualnetworksPropertiesResponseRouteTableOutput{})
	pulumi.RegisterOutputType(VirtualnetworksPropertiesResponseRouteTablePtrOutput{})
	pulumi.RegisterOutputType(VirtualnetworksPropertiesResponseRoutesOutput{})
	pulumi.RegisterOutputType(VirtualnetworksPropertiesResponseRoutesArrayOutput{})
	pulumi.RegisterOutputType(VirtualnetworksPropertiesResponseSubnetsOutput{})
	pulumi.RegisterOutputType(VirtualnetworksPropertiesResponseSubnetsArrayOutput{})
	pulumi.RegisterOutputType(VirtualnetworksPropertiesRouteTableOutput{})
	pulumi.RegisterOutputType(VirtualnetworksPropertiesRouteTablePtrOutput{})
	pulumi.RegisterOutputType(VirtualnetworksPropertiesRoutesOutput{})
	pulumi.RegisterOutputType(VirtualnetworksPropertiesRoutesArrayOutput{})
	pulumi.RegisterOutputType(VirtualnetworksPropertiesSubnetsOutput{})
	pulumi.RegisterOutputType(VirtualnetworksPropertiesSubnetsArrayOutput{})
}
