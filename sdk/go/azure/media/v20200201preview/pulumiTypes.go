


package v20200201preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type MediaGraphAssetSink struct {
	AssetName string   `pulumi:"assetName"`
	Inputs    []string `pulumi:"inputs"`
	Name      string   `pulumi:"name"`
	OdataType string   `pulumi:"odataType"`
}





type MediaGraphAssetSinkInput interface {
	pulumi.Input

	ToMediaGraphAssetSinkOutput() MediaGraphAssetSinkOutput
	ToMediaGraphAssetSinkOutputWithContext(context.Context) MediaGraphAssetSinkOutput
}

type MediaGraphAssetSinkArgs struct {
	AssetName pulumi.StringInput      `pulumi:"assetName"`
	Inputs    pulumi.StringArrayInput `pulumi:"inputs"`
	Name      pulumi.StringInput      `pulumi:"name"`
	OdataType pulumi.StringInput      `pulumi:"odataType"`
}

func (MediaGraphAssetSinkArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MediaGraphAssetSink)(nil)).Elem()
}

func (i MediaGraphAssetSinkArgs) ToMediaGraphAssetSinkOutput() MediaGraphAssetSinkOutput {
	return i.ToMediaGraphAssetSinkOutputWithContext(context.Background())
}

func (i MediaGraphAssetSinkArgs) ToMediaGraphAssetSinkOutputWithContext(ctx context.Context) MediaGraphAssetSinkOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MediaGraphAssetSinkOutput)
}





type MediaGraphAssetSinkArrayInput interface {
	pulumi.Input

	ToMediaGraphAssetSinkArrayOutput() MediaGraphAssetSinkArrayOutput
	ToMediaGraphAssetSinkArrayOutputWithContext(context.Context) MediaGraphAssetSinkArrayOutput
}

type MediaGraphAssetSinkArray []MediaGraphAssetSinkInput

func (MediaGraphAssetSinkArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]MediaGraphAssetSink)(nil)).Elem()
}

func (i MediaGraphAssetSinkArray) ToMediaGraphAssetSinkArrayOutput() MediaGraphAssetSinkArrayOutput {
	return i.ToMediaGraphAssetSinkArrayOutputWithContext(context.Background())
}

func (i MediaGraphAssetSinkArray) ToMediaGraphAssetSinkArrayOutputWithContext(ctx context.Context) MediaGraphAssetSinkArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MediaGraphAssetSinkArrayOutput)
}

type MediaGraphAssetSinkOutput struct{ *pulumi.OutputState }

func (MediaGraphAssetSinkOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MediaGraphAssetSink)(nil)).Elem()
}

func (o MediaGraphAssetSinkOutput) ToMediaGraphAssetSinkOutput() MediaGraphAssetSinkOutput {
	return o
}

func (o MediaGraphAssetSinkOutput) ToMediaGraphAssetSinkOutputWithContext(ctx context.Context) MediaGraphAssetSinkOutput {
	return o
}

func (o MediaGraphAssetSinkOutput) AssetName() pulumi.StringOutput {
	return o.ApplyT(func(v MediaGraphAssetSink) string { return v.AssetName }).(pulumi.StringOutput)
}

func (o MediaGraphAssetSinkOutput) Inputs() pulumi.StringArrayOutput {
	return o.ApplyT(func(v MediaGraphAssetSink) []string { return v.Inputs }).(pulumi.StringArrayOutput)
}

func (o MediaGraphAssetSinkOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v MediaGraphAssetSink) string { return v.Name }).(pulumi.StringOutput)
}

func (o MediaGraphAssetSinkOutput) OdataType() pulumi.StringOutput {
	return o.ApplyT(func(v MediaGraphAssetSink) string { return v.OdataType }).(pulumi.StringOutput)
}

type MediaGraphAssetSinkArrayOutput struct{ *pulumi.OutputState }

func (MediaGraphAssetSinkArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]MediaGraphAssetSink)(nil)).Elem()
}

func (o MediaGraphAssetSinkArrayOutput) ToMediaGraphAssetSinkArrayOutput() MediaGraphAssetSinkArrayOutput {
	return o
}

func (o MediaGraphAssetSinkArrayOutput) ToMediaGraphAssetSinkArrayOutputWithContext(ctx context.Context) MediaGraphAssetSinkArrayOutput {
	return o
}

func (o MediaGraphAssetSinkArrayOutput) Index(i pulumi.IntInput) MediaGraphAssetSinkOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) MediaGraphAssetSink {
		return vs[0].([]MediaGraphAssetSink)[vs[1].(int)]
	}).(MediaGraphAssetSinkOutput)
}

type MediaGraphAssetSinkResponse struct {
	AssetName string   `pulumi:"assetName"`
	Inputs    []string `pulumi:"inputs"`
	Name      string   `pulumi:"name"`
	OdataType string   `pulumi:"odataType"`
}

type MediaGraphAssetSinkResponseOutput struct{ *pulumi.OutputState }

func (MediaGraphAssetSinkResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MediaGraphAssetSinkResponse)(nil)).Elem()
}

func (o MediaGraphAssetSinkResponseOutput) ToMediaGraphAssetSinkResponseOutput() MediaGraphAssetSinkResponseOutput {
	return o
}

func (o MediaGraphAssetSinkResponseOutput) ToMediaGraphAssetSinkResponseOutputWithContext(ctx context.Context) MediaGraphAssetSinkResponseOutput {
	return o
}

func (o MediaGraphAssetSinkResponseOutput) AssetName() pulumi.StringOutput {
	return o.ApplyT(func(v MediaGraphAssetSinkResponse) string { return v.AssetName }).(pulumi.StringOutput)
}

func (o MediaGraphAssetSinkResponseOutput) Inputs() pulumi.StringArrayOutput {
	return o.ApplyT(func(v MediaGraphAssetSinkResponse) []string { return v.Inputs }).(pulumi.StringArrayOutput)
}

func (o MediaGraphAssetSinkResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v MediaGraphAssetSinkResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o MediaGraphAssetSinkResponseOutput) OdataType() pulumi.StringOutput {
	return o.ApplyT(func(v MediaGraphAssetSinkResponse) string { return v.OdataType }).(pulumi.StringOutput)
}

type MediaGraphAssetSinkResponseArrayOutput struct{ *pulumi.OutputState }

func (MediaGraphAssetSinkResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]MediaGraphAssetSinkResponse)(nil)).Elem()
}

func (o MediaGraphAssetSinkResponseArrayOutput) ToMediaGraphAssetSinkResponseArrayOutput() MediaGraphAssetSinkResponseArrayOutput {
	return o
}

func (o MediaGraphAssetSinkResponseArrayOutput) ToMediaGraphAssetSinkResponseArrayOutputWithContext(ctx context.Context) MediaGraphAssetSinkResponseArrayOutput {
	return o
}

func (o MediaGraphAssetSinkResponseArrayOutput) Index(i pulumi.IntInput) MediaGraphAssetSinkResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) MediaGraphAssetSinkResponse {
		return vs[0].([]MediaGraphAssetSinkResponse)[vs[1].(int)]
	}).(MediaGraphAssetSinkResponseOutput)
}

type MediaGraphClearEndpoint struct {
	Credentials *MediaGraphUsernamePasswordCredentials `pulumi:"credentials"`
	OdataType   string                                 `pulumi:"odataType"`
	Url         string                                 `pulumi:"url"`
}

type MediaGraphClearEndpointResponse struct {
	Credentials *MediaGraphUsernamePasswordCredentialsResponse `pulumi:"credentials"`
	OdataType   string                                         `pulumi:"odataType"`
	Url         string                                         `pulumi:"url"`
}

type MediaGraphPemCertificateList struct {
	Certificates []string `pulumi:"certificates"`
	OdataType    string   `pulumi:"odataType"`
}

type MediaGraphPemCertificateListResponse struct {
	Certificates []string `pulumi:"certificates"`
	OdataType    string   `pulumi:"odataType"`
}

type MediaGraphRtspSource struct {
	Endpoint  interface{} `pulumi:"endpoint"`
	Name      string      `pulumi:"name"`
	OdataType string      `pulumi:"odataType"`
	Transport string      `pulumi:"transport"`
}





type MediaGraphRtspSourceInput interface {
	pulumi.Input

	ToMediaGraphRtspSourceOutput() MediaGraphRtspSourceOutput
	ToMediaGraphRtspSourceOutputWithContext(context.Context) MediaGraphRtspSourceOutput
}

type MediaGraphRtspSourceArgs struct {
	Endpoint  pulumi.Input       `pulumi:"endpoint"`
	Name      pulumi.StringInput `pulumi:"name"`
	OdataType pulumi.StringInput `pulumi:"odataType"`
	Transport pulumi.StringInput `pulumi:"transport"`
}

func (MediaGraphRtspSourceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MediaGraphRtspSource)(nil)).Elem()
}

func (i MediaGraphRtspSourceArgs) ToMediaGraphRtspSourceOutput() MediaGraphRtspSourceOutput {
	return i.ToMediaGraphRtspSourceOutputWithContext(context.Background())
}

func (i MediaGraphRtspSourceArgs) ToMediaGraphRtspSourceOutputWithContext(ctx context.Context) MediaGraphRtspSourceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MediaGraphRtspSourceOutput)
}





type MediaGraphRtspSourceArrayInput interface {
	pulumi.Input

	ToMediaGraphRtspSourceArrayOutput() MediaGraphRtspSourceArrayOutput
	ToMediaGraphRtspSourceArrayOutputWithContext(context.Context) MediaGraphRtspSourceArrayOutput
}

type MediaGraphRtspSourceArray []MediaGraphRtspSourceInput

func (MediaGraphRtspSourceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]MediaGraphRtspSource)(nil)).Elem()
}

func (i MediaGraphRtspSourceArray) ToMediaGraphRtspSourceArrayOutput() MediaGraphRtspSourceArrayOutput {
	return i.ToMediaGraphRtspSourceArrayOutputWithContext(context.Background())
}

func (i MediaGraphRtspSourceArray) ToMediaGraphRtspSourceArrayOutputWithContext(ctx context.Context) MediaGraphRtspSourceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MediaGraphRtspSourceArrayOutput)
}

type MediaGraphRtspSourceOutput struct{ *pulumi.OutputState }

func (MediaGraphRtspSourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MediaGraphRtspSource)(nil)).Elem()
}

func (o MediaGraphRtspSourceOutput) ToMediaGraphRtspSourceOutput() MediaGraphRtspSourceOutput {
	return o
}

func (o MediaGraphRtspSourceOutput) ToMediaGraphRtspSourceOutputWithContext(ctx context.Context) MediaGraphRtspSourceOutput {
	return o
}

func (o MediaGraphRtspSourceOutput) Endpoint() pulumi.AnyOutput {
	return o.ApplyT(func(v MediaGraphRtspSource) interface{} { return v.Endpoint }).(pulumi.AnyOutput)
}

func (o MediaGraphRtspSourceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v MediaGraphRtspSource) string { return v.Name }).(pulumi.StringOutput)
}

func (o MediaGraphRtspSourceOutput) OdataType() pulumi.StringOutput {
	return o.ApplyT(func(v MediaGraphRtspSource) string { return v.OdataType }).(pulumi.StringOutput)
}

func (o MediaGraphRtspSourceOutput) Transport() pulumi.StringOutput {
	return o.ApplyT(func(v MediaGraphRtspSource) string { return v.Transport }).(pulumi.StringOutput)
}

type MediaGraphRtspSourceArrayOutput struct{ *pulumi.OutputState }

func (MediaGraphRtspSourceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]MediaGraphRtspSource)(nil)).Elem()
}

func (o MediaGraphRtspSourceArrayOutput) ToMediaGraphRtspSourceArrayOutput() MediaGraphRtspSourceArrayOutput {
	return o
}

func (o MediaGraphRtspSourceArrayOutput) ToMediaGraphRtspSourceArrayOutputWithContext(ctx context.Context) MediaGraphRtspSourceArrayOutput {
	return o
}

func (o MediaGraphRtspSourceArrayOutput) Index(i pulumi.IntInput) MediaGraphRtspSourceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) MediaGraphRtspSource {
		return vs[0].([]MediaGraphRtspSource)[vs[1].(int)]
	}).(MediaGraphRtspSourceOutput)
}

type MediaGraphRtspSourceResponse struct {
	Endpoint  interface{} `pulumi:"endpoint"`
	Name      string      `pulumi:"name"`
	OdataType string      `pulumi:"odataType"`
	Transport string      `pulumi:"transport"`
}

type MediaGraphRtspSourceResponseOutput struct{ *pulumi.OutputState }

func (MediaGraphRtspSourceResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MediaGraphRtspSourceResponse)(nil)).Elem()
}

func (o MediaGraphRtspSourceResponseOutput) ToMediaGraphRtspSourceResponseOutput() MediaGraphRtspSourceResponseOutput {
	return o
}

func (o MediaGraphRtspSourceResponseOutput) ToMediaGraphRtspSourceResponseOutputWithContext(ctx context.Context) MediaGraphRtspSourceResponseOutput {
	return o
}

func (o MediaGraphRtspSourceResponseOutput) Endpoint() pulumi.AnyOutput {
	return o.ApplyT(func(v MediaGraphRtspSourceResponse) interface{} { return v.Endpoint }).(pulumi.AnyOutput)
}

func (o MediaGraphRtspSourceResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v MediaGraphRtspSourceResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o MediaGraphRtspSourceResponseOutput) OdataType() pulumi.StringOutput {
	return o.ApplyT(func(v MediaGraphRtspSourceResponse) string { return v.OdataType }).(pulumi.StringOutput)
}

func (o MediaGraphRtspSourceResponseOutput) Transport() pulumi.StringOutput {
	return o.ApplyT(func(v MediaGraphRtspSourceResponse) string { return v.Transport }).(pulumi.StringOutput)
}

type MediaGraphRtspSourceResponseArrayOutput struct{ *pulumi.OutputState }

func (MediaGraphRtspSourceResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]MediaGraphRtspSourceResponse)(nil)).Elem()
}

func (o MediaGraphRtspSourceResponseArrayOutput) ToMediaGraphRtspSourceResponseArrayOutput() MediaGraphRtspSourceResponseArrayOutput {
	return o
}

func (o MediaGraphRtspSourceResponseArrayOutput) ToMediaGraphRtspSourceResponseArrayOutputWithContext(ctx context.Context) MediaGraphRtspSourceResponseArrayOutput {
	return o
}

func (o MediaGraphRtspSourceResponseArrayOutput) Index(i pulumi.IntInput) MediaGraphRtspSourceResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) MediaGraphRtspSourceResponse {
		return vs[0].([]MediaGraphRtspSourceResponse)[vs[1].(int)]
	}).(MediaGraphRtspSourceResponseOutput)
}

type MediaGraphTlsEndpoint struct {
	Credentials         *MediaGraphUsernamePasswordCredentials `pulumi:"credentials"`
	OdataType           string                                 `pulumi:"odataType"`
	TrustedCertificates *MediaGraphPemCertificateList          `pulumi:"trustedCertificates"`
	Url                 string                                 `pulumi:"url"`
	ValidationOptions   *MediaGraphTlsValidationOptions        `pulumi:"validationOptions"`
}

type MediaGraphTlsEndpointResponse struct {
	Credentials         *MediaGraphUsernamePasswordCredentialsResponse `pulumi:"credentials"`
	OdataType           string                                         `pulumi:"odataType"`
	TrustedCertificates *MediaGraphPemCertificateListResponse          `pulumi:"trustedCertificates"`
	Url                 string                                         `pulumi:"url"`
	ValidationOptions   *MediaGraphTlsValidationOptionsResponse        `pulumi:"validationOptions"`
}

type MediaGraphTlsValidationOptions struct {
	IgnoreHostname  bool `pulumi:"ignoreHostname"`
	IgnoreSignature bool `pulumi:"ignoreSignature"`
}

type MediaGraphTlsValidationOptionsResponse struct {
	IgnoreHostname  bool `pulumi:"ignoreHostname"`
	IgnoreSignature bool `pulumi:"ignoreSignature"`
}

type MediaGraphUsernamePasswordCredentials struct {
	OdataType string `pulumi:"odataType"`
	Password  string `pulumi:"password"`
	Username  string `pulumi:"username"`
}

type MediaGraphUsernamePasswordCredentialsResponse struct {
	OdataType string `pulumi:"odataType"`
	Password  string `pulumi:"password"`
	Username  string `pulumi:"username"`
}

func init() {
	pulumi.RegisterOutputType(MediaGraphAssetSinkOutput{})
	pulumi.RegisterOutputType(MediaGraphAssetSinkArrayOutput{})
	pulumi.RegisterOutputType(MediaGraphAssetSinkResponseOutput{})
	pulumi.RegisterOutputType(MediaGraphAssetSinkResponseArrayOutput{})
	pulumi.RegisterOutputType(MediaGraphRtspSourceOutput{})
	pulumi.RegisterOutputType(MediaGraphRtspSourceArrayOutput{})
	pulumi.RegisterOutputType(MediaGraphRtspSourceResponseOutput{})
	pulumi.RegisterOutputType(MediaGraphRtspSourceResponseArrayOutput{})
}
