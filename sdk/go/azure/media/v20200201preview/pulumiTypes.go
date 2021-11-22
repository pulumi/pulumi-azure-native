


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





type MediaGraphAssetSinkResponseInput interface {
	pulumi.Input

	ToMediaGraphAssetSinkResponseOutput() MediaGraphAssetSinkResponseOutput
	ToMediaGraphAssetSinkResponseOutputWithContext(context.Context) MediaGraphAssetSinkResponseOutput
}

type MediaGraphAssetSinkResponseArgs struct {
	AssetName pulumi.StringInput      `pulumi:"assetName"`
	Inputs    pulumi.StringArrayInput `pulumi:"inputs"`
	Name      pulumi.StringInput      `pulumi:"name"`
	OdataType pulumi.StringInput      `pulumi:"odataType"`
}

func (MediaGraphAssetSinkResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MediaGraphAssetSinkResponse)(nil)).Elem()
}

func (i MediaGraphAssetSinkResponseArgs) ToMediaGraphAssetSinkResponseOutput() MediaGraphAssetSinkResponseOutput {
	return i.ToMediaGraphAssetSinkResponseOutputWithContext(context.Background())
}

func (i MediaGraphAssetSinkResponseArgs) ToMediaGraphAssetSinkResponseOutputWithContext(ctx context.Context) MediaGraphAssetSinkResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MediaGraphAssetSinkResponseOutput)
}





type MediaGraphAssetSinkResponseArrayInput interface {
	pulumi.Input

	ToMediaGraphAssetSinkResponseArrayOutput() MediaGraphAssetSinkResponseArrayOutput
	ToMediaGraphAssetSinkResponseArrayOutputWithContext(context.Context) MediaGraphAssetSinkResponseArrayOutput
}

type MediaGraphAssetSinkResponseArray []MediaGraphAssetSinkResponseInput

func (MediaGraphAssetSinkResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]MediaGraphAssetSinkResponse)(nil)).Elem()
}

func (i MediaGraphAssetSinkResponseArray) ToMediaGraphAssetSinkResponseArrayOutput() MediaGraphAssetSinkResponseArrayOutput {
	return i.ToMediaGraphAssetSinkResponseArrayOutputWithContext(context.Background())
}

func (i MediaGraphAssetSinkResponseArray) ToMediaGraphAssetSinkResponseArrayOutputWithContext(ctx context.Context) MediaGraphAssetSinkResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MediaGraphAssetSinkResponseArrayOutput)
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





type MediaGraphClearEndpointInput interface {
	pulumi.Input

	ToMediaGraphClearEndpointOutput() MediaGraphClearEndpointOutput
	ToMediaGraphClearEndpointOutputWithContext(context.Context) MediaGraphClearEndpointOutput
}

type MediaGraphClearEndpointArgs struct {
	Credentials MediaGraphUsernamePasswordCredentialsPtrInput `pulumi:"credentials"`
	OdataType   pulumi.StringInput                            `pulumi:"odataType"`
	Url         pulumi.StringInput                            `pulumi:"url"`
}

func (MediaGraphClearEndpointArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MediaGraphClearEndpoint)(nil)).Elem()
}

func (i MediaGraphClearEndpointArgs) ToMediaGraphClearEndpointOutput() MediaGraphClearEndpointOutput {
	return i.ToMediaGraphClearEndpointOutputWithContext(context.Background())
}

func (i MediaGraphClearEndpointArgs) ToMediaGraphClearEndpointOutputWithContext(ctx context.Context) MediaGraphClearEndpointOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MediaGraphClearEndpointOutput)
}

type MediaGraphClearEndpointOutput struct{ *pulumi.OutputState }

func (MediaGraphClearEndpointOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MediaGraphClearEndpoint)(nil)).Elem()
}

func (o MediaGraphClearEndpointOutput) ToMediaGraphClearEndpointOutput() MediaGraphClearEndpointOutput {
	return o
}

func (o MediaGraphClearEndpointOutput) ToMediaGraphClearEndpointOutputWithContext(ctx context.Context) MediaGraphClearEndpointOutput {
	return o
}

func (o MediaGraphClearEndpointOutput) Credentials() MediaGraphUsernamePasswordCredentialsPtrOutput {
	return o.ApplyT(func(v MediaGraphClearEndpoint) *MediaGraphUsernamePasswordCredentials { return v.Credentials }).(MediaGraphUsernamePasswordCredentialsPtrOutput)
}

func (o MediaGraphClearEndpointOutput) OdataType() pulumi.StringOutput {
	return o.ApplyT(func(v MediaGraphClearEndpoint) string { return v.OdataType }).(pulumi.StringOutput)
}

func (o MediaGraphClearEndpointOutput) Url() pulumi.StringOutput {
	return o.ApplyT(func(v MediaGraphClearEndpoint) string { return v.Url }).(pulumi.StringOutput)
}

type MediaGraphClearEndpointResponse struct {
	Credentials *MediaGraphUsernamePasswordCredentialsResponse `pulumi:"credentials"`
	OdataType   string                                         `pulumi:"odataType"`
	Url         string                                         `pulumi:"url"`
}





type MediaGraphClearEndpointResponseInput interface {
	pulumi.Input

	ToMediaGraphClearEndpointResponseOutput() MediaGraphClearEndpointResponseOutput
	ToMediaGraphClearEndpointResponseOutputWithContext(context.Context) MediaGraphClearEndpointResponseOutput
}

type MediaGraphClearEndpointResponseArgs struct {
	Credentials MediaGraphUsernamePasswordCredentialsResponsePtrInput `pulumi:"credentials"`
	OdataType   pulumi.StringInput                                    `pulumi:"odataType"`
	Url         pulumi.StringInput                                    `pulumi:"url"`
}

func (MediaGraphClearEndpointResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MediaGraphClearEndpointResponse)(nil)).Elem()
}

func (i MediaGraphClearEndpointResponseArgs) ToMediaGraphClearEndpointResponseOutput() MediaGraphClearEndpointResponseOutput {
	return i.ToMediaGraphClearEndpointResponseOutputWithContext(context.Background())
}

func (i MediaGraphClearEndpointResponseArgs) ToMediaGraphClearEndpointResponseOutputWithContext(ctx context.Context) MediaGraphClearEndpointResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MediaGraphClearEndpointResponseOutput)
}

type MediaGraphClearEndpointResponseOutput struct{ *pulumi.OutputState }

func (MediaGraphClearEndpointResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MediaGraphClearEndpointResponse)(nil)).Elem()
}

func (o MediaGraphClearEndpointResponseOutput) ToMediaGraphClearEndpointResponseOutput() MediaGraphClearEndpointResponseOutput {
	return o
}

func (o MediaGraphClearEndpointResponseOutput) ToMediaGraphClearEndpointResponseOutputWithContext(ctx context.Context) MediaGraphClearEndpointResponseOutput {
	return o
}

func (o MediaGraphClearEndpointResponseOutput) Credentials() MediaGraphUsernamePasswordCredentialsResponsePtrOutput {
	return o.ApplyT(func(v MediaGraphClearEndpointResponse) *MediaGraphUsernamePasswordCredentialsResponse {
		return v.Credentials
	}).(MediaGraphUsernamePasswordCredentialsResponsePtrOutput)
}

func (o MediaGraphClearEndpointResponseOutput) OdataType() pulumi.StringOutput {
	return o.ApplyT(func(v MediaGraphClearEndpointResponse) string { return v.OdataType }).(pulumi.StringOutput)
}

func (o MediaGraphClearEndpointResponseOutput) Url() pulumi.StringOutput {
	return o.ApplyT(func(v MediaGraphClearEndpointResponse) string { return v.Url }).(pulumi.StringOutput)
}

type MediaGraphPemCertificateList struct {
	Certificates []string `pulumi:"certificates"`
	OdataType    string   `pulumi:"odataType"`
}





type MediaGraphPemCertificateListInput interface {
	pulumi.Input

	ToMediaGraphPemCertificateListOutput() MediaGraphPemCertificateListOutput
	ToMediaGraphPemCertificateListOutputWithContext(context.Context) MediaGraphPemCertificateListOutput
}

type MediaGraphPemCertificateListArgs struct {
	Certificates pulumi.StringArrayInput `pulumi:"certificates"`
	OdataType    pulumi.StringInput      `pulumi:"odataType"`
}

func (MediaGraphPemCertificateListArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MediaGraphPemCertificateList)(nil)).Elem()
}

func (i MediaGraphPemCertificateListArgs) ToMediaGraphPemCertificateListOutput() MediaGraphPemCertificateListOutput {
	return i.ToMediaGraphPemCertificateListOutputWithContext(context.Background())
}

func (i MediaGraphPemCertificateListArgs) ToMediaGraphPemCertificateListOutputWithContext(ctx context.Context) MediaGraphPemCertificateListOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MediaGraphPemCertificateListOutput)
}

func (i MediaGraphPemCertificateListArgs) ToMediaGraphPemCertificateListPtrOutput() MediaGraphPemCertificateListPtrOutput {
	return i.ToMediaGraphPemCertificateListPtrOutputWithContext(context.Background())
}

func (i MediaGraphPemCertificateListArgs) ToMediaGraphPemCertificateListPtrOutputWithContext(ctx context.Context) MediaGraphPemCertificateListPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MediaGraphPemCertificateListOutput).ToMediaGraphPemCertificateListPtrOutputWithContext(ctx)
}









type MediaGraphPemCertificateListPtrInput interface {
	pulumi.Input

	ToMediaGraphPemCertificateListPtrOutput() MediaGraphPemCertificateListPtrOutput
	ToMediaGraphPemCertificateListPtrOutputWithContext(context.Context) MediaGraphPemCertificateListPtrOutput
}

type mediaGraphPemCertificateListPtrType MediaGraphPemCertificateListArgs

func MediaGraphPemCertificateListPtr(v *MediaGraphPemCertificateListArgs) MediaGraphPemCertificateListPtrInput {
	return (*mediaGraphPemCertificateListPtrType)(v)
}

func (*mediaGraphPemCertificateListPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**MediaGraphPemCertificateList)(nil)).Elem()
}

func (i *mediaGraphPemCertificateListPtrType) ToMediaGraphPemCertificateListPtrOutput() MediaGraphPemCertificateListPtrOutput {
	return i.ToMediaGraphPemCertificateListPtrOutputWithContext(context.Background())
}

func (i *mediaGraphPemCertificateListPtrType) ToMediaGraphPemCertificateListPtrOutputWithContext(ctx context.Context) MediaGraphPemCertificateListPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MediaGraphPemCertificateListPtrOutput)
}

type MediaGraphPemCertificateListOutput struct{ *pulumi.OutputState }

func (MediaGraphPemCertificateListOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MediaGraphPemCertificateList)(nil)).Elem()
}

func (o MediaGraphPemCertificateListOutput) ToMediaGraphPemCertificateListOutput() MediaGraphPemCertificateListOutput {
	return o
}

func (o MediaGraphPemCertificateListOutput) ToMediaGraphPemCertificateListOutputWithContext(ctx context.Context) MediaGraphPemCertificateListOutput {
	return o
}

func (o MediaGraphPemCertificateListOutput) ToMediaGraphPemCertificateListPtrOutput() MediaGraphPemCertificateListPtrOutput {
	return o.ToMediaGraphPemCertificateListPtrOutputWithContext(context.Background())
}

func (o MediaGraphPemCertificateListOutput) ToMediaGraphPemCertificateListPtrOutputWithContext(ctx context.Context) MediaGraphPemCertificateListPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v MediaGraphPemCertificateList) *MediaGraphPemCertificateList {
		return &v
	}).(MediaGraphPemCertificateListPtrOutput)
}

func (o MediaGraphPemCertificateListOutput) Certificates() pulumi.StringArrayOutput {
	return o.ApplyT(func(v MediaGraphPemCertificateList) []string { return v.Certificates }).(pulumi.StringArrayOutput)
}

func (o MediaGraphPemCertificateListOutput) OdataType() pulumi.StringOutput {
	return o.ApplyT(func(v MediaGraphPemCertificateList) string { return v.OdataType }).(pulumi.StringOutput)
}

type MediaGraphPemCertificateListPtrOutput struct{ *pulumi.OutputState }

func (MediaGraphPemCertificateListPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MediaGraphPemCertificateList)(nil)).Elem()
}

func (o MediaGraphPemCertificateListPtrOutput) ToMediaGraphPemCertificateListPtrOutput() MediaGraphPemCertificateListPtrOutput {
	return o
}

func (o MediaGraphPemCertificateListPtrOutput) ToMediaGraphPemCertificateListPtrOutputWithContext(ctx context.Context) MediaGraphPemCertificateListPtrOutput {
	return o
}

func (o MediaGraphPemCertificateListPtrOutput) Elem() MediaGraphPemCertificateListOutput {
	return o.ApplyT(func(v *MediaGraphPemCertificateList) MediaGraphPemCertificateList {
		if v != nil {
			return *v
		}
		var ret MediaGraphPemCertificateList
		return ret
	}).(MediaGraphPemCertificateListOutput)
}

func (o MediaGraphPemCertificateListPtrOutput) Certificates() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *MediaGraphPemCertificateList) []string {
		if v == nil {
			return nil
		}
		return v.Certificates
	}).(pulumi.StringArrayOutput)
}

func (o MediaGraphPemCertificateListPtrOutput) OdataType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MediaGraphPemCertificateList) *string {
		if v == nil {
			return nil
		}
		return &v.OdataType
	}).(pulumi.StringPtrOutput)
}

type MediaGraphPemCertificateListResponse struct {
	Certificates []string `pulumi:"certificates"`
	OdataType    string   `pulumi:"odataType"`
}





type MediaGraphPemCertificateListResponseInput interface {
	pulumi.Input

	ToMediaGraphPemCertificateListResponseOutput() MediaGraphPemCertificateListResponseOutput
	ToMediaGraphPemCertificateListResponseOutputWithContext(context.Context) MediaGraphPemCertificateListResponseOutput
}

type MediaGraphPemCertificateListResponseArgs struct {
	Certificates pulumi.StringArrayInput `pulumi:"certificates"`
	OdataType    pulumi.StringInput      `pulumi:"odataType"`
}

func (MediaGraphPemCertificateListResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MediaGraphPemCertificateListResponse)(nil)).Elem()
}

func (i MediaGraphPemCertificateListResponseArgs) ToMediaGraphPemCertificateListResponseOutput() MediaGraphPemCertificateListResponseOutput {
	return i.ToMediaGraphPemCertificateListResponseOutputWithContext(context.Background())
}

func (i MediaGraphPemCertificateListResponseArgs) ToMediaGraphPemCertificateListResponseOutputWithContext(ctx context.Context) MediaGraphPemCertificateListResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MediaGraphPemCertificateListResponseOutput)
}

func (i MediaGraphPemCertificateListResponseArgs) ToMediaGraphPemCertificateListResponsePtrOutput() MediaGraphPemCertificateListResponsePtrOutput {
	return i.ToMediaGraphPemCertificateListResponsePtrOutputWithContext(context.Background())
}

func (i MediaGraphPemCertificateListResponseArgs) ToMediaGraphPemCertificateListResponsePtrOutputWithContext(ctx context.Context) MediaGraphPemCertificateListResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MediaGraphPemCertificateListResponseOutput).ToMediaGraphPemCertificateListResponsePtrOutputWithContext(ctx)
}









type MediaGraphPemCertificateListResponsePtrInput interface {
	pulumi.Input

	ToMediaGraphPemCertificateListResponsePtrOutput() MediaGraphPemCertificateListResponsePtrOutput
	ToMediaGraphPemCertificateListResponsePtrOutputWithContext(context.Context) MediaGraphPemCertificateListResponsePtrOutput
}

type mediaGraphPemCertificateListResponsePtrType MediaGraphPemCertificateListResponseArgs

func MediaGraphPemCertificateListResponsePtr(v *MediaGraphPemCertificateListResponseArgs) MediaGraphPemCertificateListResponsePtrInput {
	return (*mediaGraphPemCertificateListResponsePtrType)(v)
}

func (*mediaGraphPemCertificateListResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**MediaGraphPemCertificateListResponse)(nil)).Elem()
}

func (i *mediaGraphPemCertificateListResponsePtrType) ToMediaGraphPemCertificateListResponsePtrOutput() MediaGraphPemCertificateListResponsePtrOutput {
	return i.ToMediaGraphPemCertificateListResponsePtrOutputWithContext(context.Background())
}

func (i *mediaGraphPemCertificateListResponsePtrType) ToMediaGraphPemCertificateListResponsePtrOutputWithContext(ctx context.Context) MediaGraphPemCertificateListResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MediaGraphPemCertificateListResponsePtrOutput)
}

type MediaGraphPemCertificateListResponseOutput struct{ *pulumi.OutputState }

func (MediaGraphPemCertificateListResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MediaGraphPemCertificateListResponse)(nil)).Elem()
}

func (o MediaGraphPemCertificateListResponseOutput) ToMediaGraphPemCertificateListResponseOutput() MediaGraphPemCertificateListResponseOutput {
	return o
}

func (o MediaGraphPemCertificateListResponseOutput) ToMediaGraphPemCertificateListResponseOutputWithContext(ctx context.Context) MediaGraphPemCertificateListResponseOutput {
	return o
}

func (o MediaGraphPemCertificateListResponseOutput) ToMediaGraphPemCertificateListResponsePtrOutput() MediaGraphPemCertificateListResponsePtrOutput {
	return o.ToMediaGraphPemCertificateListResponsePtrOutputWithContext(context.Background())
}

func (o MediaGraphPemCertificateListResponseOutput) ToMediaGraphPemCertificateListResponsePtrOutputWithContext(ctx context.Context) MediaGraphPemCertificateListResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v MediaGraphPemCertificateListResponse) *MediaGraphPemCertificateListResponse {
		return &v
	}).(MediaGraphPemCertificateListResponsePtrOutput)
}

func (o MediaGraphPemCertificateListResponseOutput) Certificates() pulumi.StringArrayOutput {
	return o.ApplyT(func(v MediaGraphPemCertificateListResponse) []string { return v.Certificates }).(pulumi.StringArrayOutput)
}

func (o MediaGraphPemCertificateListResponseOutput) OdataType() pulumi.StringOutput {
	return o.ApplyT(func(v MediaGraphPemCertificateListResponse) string { return v.OdataType }).(pulumi.StringOutput)
}

type MediaGraphPemCertificateListResponsePtrOutput struct{ *pulumi.OutputState }

func (MediaGraphPemCertificateListResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MediaGraphPemCertificateListResponse)(nil)).Elem()
}

func (o MediaGraphPemCertificateListResponsePtrOutput) ToMediaGraphPemCertificateListResponsePtrOutput() MediaGraphPemCertificateListResponsePtrOutput {
	return o
}

func (o MediaGraphPemCertificateListResponsePtrOutput) ToMediaGraphPemCertificateListResponsePtrOutputWithContext(ctx context.Context) MediaGraphPemCertificateListResponsePtrOutput {
	return o
}

func (o MediaGraphPemCertificateListResponsePtrOutput) Elem() MediaGraphPemCertificateListResponseOutput {
	return o.ApplyT(func(v *MediaGraphPemCertificateListResponse) MediaGraphPemCertificateListResponse {
		if v != nil {
			return *v
		}
		var ret MediaGraphPemCertificateListResponse
		return ret
	}).(MediaGraphPemCertificateListResponseOutput)
}

func (o MediaGraphPemCertificateListResponsePtrOutput) Certificates() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *MediaGraphPemCertificateListResponse) []string {
		if v == nil {
			return nil
		}
		return v.Certificates
	}).(pulumi.StringArrayOutput)
}

func (o MediaGraphPemCertificateListResponsePtrOutput) OdataType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MediaGraphPemCertificateListResponse) *string {
		if v == nil {
			return nil
		}
		return &v.OdataType
	}).(pulumi.StringPtrOutput)
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





type MediaGraphRtspSourceResponseInput interface {
	pulumi.Input

	ToMediaGraphRtspSourceResponseOutput() MediaGraphRtspSourceResponseOutput
	ToMediaGraphRtspSourceResponseOutputWithContext(context.Context) MediaGraphRtspSourceResponseOutput
}

type MediaGraphRtspSourceResponseArgs struct {
	Endpoint  pulumi.Input       `pulumi:"endpoint"`
	Name      pulumi.StringInput `pulumi:"name"`
	OdataType pulumi.StringInput `pulumi:"odataType"`
	Transport pulumi.StringInput `pulumi:"transport"`
}

func (MediaGraphRtspSourceResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MediaGraphRtspSourceResponse)(nil)).Elem()
}

func (i MediaGraphRtspSourceResponseArgs) ToMediaGraphRtspSourceResponseOutput() MediaGraphRtspSourceResponseOutput {
	return i.ToMediaGraphRtspSourceResponseOutputWithContext(context.Background())
}

func (i MediaGraphRtspSourceResponseArgs) ToMediaGraphRtspSourceResponseOutputWithContext(ctx context.Context) MediaGraphRtspSourceResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MediaGraphRtspSourceResponseOutput)
}





type MediaGraphRtspSourceResponseArrayInput interface {
	pulumi.Input

	ToMediaGraphRtspSourceResponseArrayOutput() MediaGraphRtspSourceResponseArrayOutput
	ToMediaGraphRtspSourceResponseArrayOutputWithContext(context.Context) MediaGraphRtspSourceResponseArrayOutput
}

type MediaGraphRtspSourceResponseArray []MediaGraphRtspSourceResponseInput

func (MediaGraphRtspSourceResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]MediaGraphRtspSourceResponse)(nil)).Elem()
}

func (i MediaGraphRtspSourceResponseArray) ToMediaGraphRtspSourceResponseArrayOutput() MediaGraphRtspSourceResponseArrayOutput {
	return i.ToMediaGraphRtspSourceResponseArrayOutputWithContext(context.Background())
}

func (i MediaGraphRtspSourceResponseArray) ToMediaGraphRtspSourceResponseArrayOutputWithContext(ctx context.Context) MediaGraphRtspSourceResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MediaGraphRtspSourceResponseArrayOutput)
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





type MediaGraphTlsEndpointInput interface {
	pulumi.Input

	ToMediaGraphTlsEndpointOutput() MediaGraphTlsEndpointOutput
	ToMediaGraphTlsEndpointOutputWithContext(context.Context) MediaGraphTlsEndpointOutput
}

type MediaGraphTlsEndpointArgs struct {
	Credentials         MediaGraphUsernamePasswordCredentialsPtrInput `pulumi:"credentials"`
	OdataType           pulumi.StringInput                            `pulumi:"odataType"`
	TrustedCertificates MediaGraphPemCertificateListPtrInput          `pulumi:"trustedCertificates"`
	Url                 pulumi.StringInput                            `pulumi:"url"`
	ValidationOptions   MediaGraphTlsValidationOptionsPtrInput        `pulumi:"validationOptions"`
}

func (MediaGraphTlsEndpointArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MediaGraphTlsEndpoint)(nil)).Elem()
}

func (i MediaGraphTlsEndpointArgs) ToMediaGraphTlsEndpointOutput() MediaGraphTlsEndpointOutput {
	return i.ToMediaGraphTlsEndpointOutputWithContext(context.Background())
}

func (i MediaGraphTlsEndpointArgs) ToMediaGraphTlsEndpointOutputWithContext(ctx context.Context) MediaGraphTlsEndpointOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MediaGraphTlsEndpointOutput)
}

type MediaGraphTlsEndpointOutput struct{ *pulumi.OutputState }

func (MediaGraphTlsEndpointOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MediaGraphTlsEndpoint)(nil)).Elem()
}

func (o MediaGraphTlsEndpointOutput) ToMediaGraphTlsEndpointOutput() MediaGraphTlsEndpointOutput {
	return o
}

func (o MediaGraphTlsEndpointOutput) ToMediaGraphTlsEndpointOutputWithContext(ctx context.Context) MediaGraphTlsEndpointOutput {
	return o
}

func (o MediaGraphTlsEndpointOutput) Credentials() MediaGraphUsernamePasswordCredentialsPtrOutput {
	return o.ApplyT(func(v MediaGraphTlsEndpoint) *MediaGraphUsernamePasswordCredentials { return v.Credentials }).(MediaGraphUsernamePasswordCredentialsPtrOutput)
}

func (o MediaGraphTlsEndpointOutput) OdataType() pulumi.StringOutput {
	return o.ApplyT(func(v MediaGraphTlsEndpoint) string { return v.OdataType }).(pulumi.StringOutput)
}

func (o MediaGraphTlsEndpointOutput) TrustedCertificates() MediaGraphPemCertificateListPtrOutput {
	return o.ApplyT(func(v MediaGraphTlsEndpoint) *MediaGraphPemCertificateList { return v.TrustedCertificates }).(MediaGraphPemCertificateListPtrOutput)
}

func (o MediaGraphTlsEndpointOutput) Url() pulumi.StringOutput {
	return o.ApplyT(func(v MediaGraphTlsEndpoint) string { return v.Url }).(pulumi.StringOutput)
}

func (o MediaGraphTlsEndpointOutput) ValidationOptions() MediaGraphTlsValidationOptionsPtrOutput {
	return o.ApplyT(func(v MediaGraphTlsEndpoint) *MediaGraphTlsValidationOptions { return v.ValidationOptions }).(MediaGraphTlsValidationOptionsPtrOutput)
}

type MediaGraphTlsEndpointResponse struct {
	Credentials         *MediaGraphUsernamePasswordCredentialsResponse `pulumi:"credentials"`
	OdataType           string                                         `pulumi:"odataType"`
	TrustedCertificates *MediaGraphPemCertificateListResponse          `pulumi:"trustedCertificates"`
	Url                 string                                         `pulumi:"url"`
	ValidationOptions   *MediaGraphTlsValidationOptionsResponse        `pulumi:"validationOptions"`
}





type MediaGraphTlsEndpointResponseInput interface {
	pulumi.Input

	ToMediaGraphTlsEndpointResponseOutput() MediaGraphTlsEndpointResponseOutput
	ToMediaGraphTlsEndpointResponseOutputWithContext(context.Context) MediaGraphTlsEndpointResponseOutput
}

type MediaGraphTlsEndpointResponseArgs struct {
	Credentials         MediaGraphUsernamePasswordCredentialsResponsePtrInput `pulumi:"credentials"`
	OdataType           pulumi.StringInput                                    `pulumi:"odataType"`
	TrustedCertificates MediaGraphPemCertificateListResponsePtrInput          `pulumi:"trustedCertificates"`
	Url                 pulumi.StringInput                                    `pulumi:"url"`
	ValidationOptions   MediaGraphTlsValidationOptionsResponsePtrInput        `pulumi:"validationOptions"`
}

func (MediaGraphTlsEndpointResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MediaGraphTlsEndpointResponse)(nil)).Elem()
}

func (i MediaGraphTlsEndpointResponseArgs) ToMediaGraphTlsEndpointResponseOutput() MediaGraphTlsEndpointResponseOutput {
	return i.ToMediaGraphTlsEndpointResponseOutputWithContext(context.Background())
}

func (i MediaGraphTlsEndpointResponseArgs) ToMediaGraphTlsEndpointResponseOutputWithContext(ctx context.Context) MediaGraphTlsEndpointResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MediaGraphTlsEndpointResponseOutput)
}

type MediaGraphTlsEndpointResponseOutput struct{ *pulumi.OutputState }

func (MediaGraphTlsEndpointResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MediaGraphTlsEndpointResponse)(nil)).Elem()
}

func (o MediaGraphTlsEndpointResponseOutput) ToMediaGraphTlsEndpointResponseOutput() MediaGraphTlsEndpointResponseOutput {
	return o
}

func (o MediaGraphTlsEndpointResponseOutput) ToMediaGraphTlsEndpointResponseOutputWithContext(ctx context.Context) MediaGraphTlsEndpointResponseOutput {
	return o
}

func (o MediaGraphTlsEndpointResponseOutput) Credentials() MediaGraphUsernamePasswordCredentialsResponsePtrOutput {
	return o.ApplyT(func(v MediaGraphTlsEndpointResponse) *MediaGraphUsernamePasswordCredentialsResponse {
		return v.Credentials
	}).(MediaGraphUsernamePasswordCredentialsResponsePtrOutput)
}

func (o MediaGraphTlsEndpointResponseOutput) OdataType() pulumi.StringOutput {
	return o.ApplyT(func(v MediaGraphTlsEndpointResponse) string { return v.OdataType }).(pulumi.StringOutput)
}

func (o MediaGraphTlsEndpointResponseOutput) TrustedCertificates() MediaGraphPemCertificateListResponsePtrOutput {
	return o.ApplyT(func(v MediaGraphTlsEndpointResponse) *MediaGraphPemCertificateListResponse {
		return v.TrustedCertificates
	}).(MediaGraphPemCertificateListResponsePtrOutput)
}

func (o MediaGraphTlsEndpointResponseOutput) Url() pulumi.StringOutput {
	return o.ApplyT(func(v MediaGraphTlsEndpointResponse) string { return v.Url }).(pulumi.StringOutput)
}

func (o MediaGraphTlsEndpointResponseOutput) ValidationOptions() MediaGraphTlsValidationOptionsResponsePtrOutput {
	return o.ApplyT(func(v MediaGraphTlsEndpointResponse) *MediaGraphTlsValidationOptionsResponse {
		return v.ValidationOptions
	}).(MediaGraphTlsValidationOptionsResponsePtrOutput)
}

type MediaGraphTlsValidationOptions struct {
	IgnoreHostname  bool `pulumi:"ignoreHostname"`
	IgnoreSignature bool `pulumi:"ignoreSignature"`
}





type MediaGraphTlsValidationOptionsInput interface {
	pulumi.Input

	ToMediaGraphTlsValidationOptionsOutput() MediaGraphTlsValidationOptionsOutput
	ToMediaGraphTlsValidationOptionsOutputWithContext(context.Context) MediaGraphTlsValidationOptionsOutput
}

type MediaGraphTlsValidationOptionsArgs struct {
	IgnoreHostname  pulumi.BoolInput `pulumi:"ignoreHostname"`
	IgnoreSignature pulumi.BoolInput `pulumi:"ignoreSignature"`
}

func (MediaGraphTlsValidationOptionsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MediaGraphTlsValidationOptions)(nil)).Elem()
}

func (i MediaGraphTlsValidationOptionsArgs) ToMediaGraphTlsValidationOptionsOutput() MediaGraphTlsValidationOptionsOutput {
	return i.ToMediaGraphTlsValidationOptionsOutputWithContext(context.Background())
}

func (i MediaGraphTlsValidationOptionsArgs) ToMediaGraphTlsValidationOptionsOutputWithContext(ctx context.Context) MediaGraphTlsValidationOptionsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MediaGraphTlsValidationOptionsOutput)
}

func (i MediaGraphTlsValidationOptionsArgs) ToMediaGraphTlsValidationOptionsPtrOutput() MediaGraphTlsValidationOptionsPtrOutput {
	return i.ToMediaGraphTlsValidationOptionsPtrOutputWithContext(context.Background())
}

func (i MediaGraphTlsValidationOptionsArgs) ToMediaGraphTlsValidationOptionsPtrOutputWithContext(ctx context.Context) MediaGraphTlsValidationOptionsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MediaGraphTlsValidationOptionsOutput).ToMediaGraphTlsValidationOptionsPtrOutputWithContext(ctx)
}









type MediaGraphTlsValidationOptionsPtrInput interface {
	pulumi.Input

	ToMediaGraphTlsValidationOptionsPtrOutput() MediaGraphTlsValidationOptionsPtrOutput
	ToMediaGraphTlsValidationOptionsPtrOutputWithContext(context.Context) MediaGraphTlsValidationOptionsPtrOutput
}

type mediaGraphTlsValidationOptionsPtrType MediaGraphTlsValidationOptionsArgs

func MediaGraphTlsValidationOptionsPtr(v *MediaGraphTlsValidationOptionsArgs) MediaGraphTlsValidationOptionsPtrInput {
	return (*mediaGraphTlsValidationOptionsPtrType)(v)
}

func (*mediaGraphTlsValidationOptionsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**MediaGraphTlsValidationOptions)(nil)).Elem()
}

func (i *mediaGraphTlsValidationOptionsPtrType) ToMediaGraphTlsValidationOptionsPtrOutput() MediaGraphTlsValidationOptionsPtrOutput {
	return i.ToMediaGraphTlsValidationOptionsPtrOutputWithContext(context.Background())
}

func (i *mediaGraphTlsValidationOptionsPtrType) ToMediaGraphTlsValidationOptionsPtrOutputWithContext(ctx context.Context) MediaGraphTlsValidationOptionsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MediaGraphTlsValidationOptionsPtrOutput)
}

type MediaGraphTlsValidationOptionsOutput struct{ *pulumi.OutputState }

func (MediaGraphTlsValidationOptionsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MediaGraphTlsValidationOptions)(nil)).Elem()
}

func (o MediaGraphTlsValidationOptionsOutput) ToMediaGraphTlsValidationOptionsOutput() MediaGraphTlsValidationOptionsOutput {
	return o
}

func (o MediaGraphTlsValidationOptionsOutput) ToMediaGraphTlsValidationOptionsOutputWithContext(ctx context.Context) MediaGraphTlsValidationOptionsOutput {
	return o
}

func (o MediaGraphTlsValidationOptionsOutput) ToMediaGraphTlsValidationOptionsPtrOutput() MediaGraphTlsValidationOptionsPtrOutput {
	return o.ToMediaGraphTlsValidationOptionsPtrOutputWithContext(context.Background())
}

func (o MediaGraphTlsValidationOptionsOutput) ToMediaGraphTlsValidationOptionsPtrOutputWithContext(ctx context.Context) MediaGraphTlsValidationOptionsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v MediaGraphTlsValidationOptions) *MediaGraphTlsValidationOptions {
		return &v
	}).(MediaGraphTlsValidationOptionsPtrOutput)
}

func (o MediaGraphTlsValidationOptionsOutput) IgnoreHostname() pulumi.BoolOutput {
	return o.ApplyT(func(v MediaGraphTlsValidationOptions) bool { return v.IgnoreHostname }).(pulumi.BoolOutput)
}

func (o MediaGraphTlsValidationOptionsOutput) IgnoreSignature() pulumi.BoolOutput {
	return o.ApplyT(func(v MediaGraphTlsValidationOptions) bool { return v.IgnoreSignature }).(pulumi.BoolOutput)
}

type MediaGraphTlsValidationOptionsPtrOutput struct{ *pulumi.OutputState }

func (MediaGraphTlsValidationOptionsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MediaGraphTlsValidationOptions)(nil)).Elem()
}

func (o MediaGraphTlsValidationOptionsPtrOutput) ToMediaGraphTlsValidationOptionsPtrOutput() MediaGraphTlsValidationOptionsPtrOutput {
	return o
}

func (o MediaGraphTlsValidationOptionsPtrOutput) ToMediaGraphTlsValidationOptionsPtrOutputWithContext(ctx context.Context) MediaGraphTlsValidationOptionsPtrOutput {
	return o
}

func (o MediaGraphTlsValidationOptionsPtrOutput) Elem() MediaGraphTlsValidationOptionsOutput {
	return o.ApplyT(func(v *MediaGraphTlsValidationOptions) MediaGraphTlsValidationOptions {
		if v != nil {
			return *v
		}
		var ret MediaGraphTlsValidationOptions
		return ret
	}).(MediaGraphTlsValidationOptionsOutput)
}

func (o MediaGraphTlsValidationOptionsPtrOutput) IgnoreHostname() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *MediaGraphTlsValidationOptions) *bool {
		if v == nil {
			return nil
		}
		return &v.IgnoreHostname
	}).(pulumi.BoolPtrOutput)
}

func (o MediaGraphTlsValidationOptionsPtrOutput) IgnoreSignature() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *MediaGraphTlsValidationOptions) *bool {
		if v == nil {
			return nil
		}
		return &v.IgnoreSignature
	}).(pulumi.BoolPtrOutput)
}

type MediaGraphTlsValidationOptionsResponse struct {
	IgnoreHostname  bool `pulumi:"ignoreHostname"`
	IgnoreSignature bool `pulumi:"ignoreSignature"`
}





type MediaGraphTlsValidationOptionsResponseInput interface {
	pulumi.Input

	ToMediaGraphTlsValidationOptionsResponseOutput() MediaGraphTlsValidationOptionsResponseOutput
	ToMediaGraphTlsValidationOptionsResponseOutputWithContext(context.Context) MediaGraphTlsValidationOptionsResponseOutput
}

type MediaGraphTlsValidationOptionsResponseArgs struct {
	IgnoreHostname  pulumi.BoolInput `pulumi:"ignoreHostname"`
	IgnoreSignature pulumi.BoolInput `pulumi:"ignoreSignature"`
}

func (MediaGraphTlsValidationOptionsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MediaGraphTlsValidationOptionsResponse)(nil)).Elem()
}

func (i MediaGraphTlsValidationOptionsResponseArgs) ToMediaGraphTlsValidationOptionsResponseOutput() MediaGraphTlsValidationOptionsResponseOutput {
	return i.ToMediaGraphTlsValidationOptionsResponseOutputWithContext(context.Background())
}

func (i MediaGraphTlsValidationOptionsResponseArgs) ToMediaGraphTlsValidationOptionsResponseOutputWithContext(ctx context.Context) MediaGraphTlsValidationOptionsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MediaGraphTlsValidationOptionsResponseOutput)
}

func (i MediaGraphTlsValidationOptionsResponseArgs) ToMediaGraphTlsValidationOptionsResponsePtrOutput() MediaGraphTlsValidationOptionsResponsePtrOutput {
	return i.ToMediaGraphTlsValidationOptionsResponsePtrOutputWithContext(context.Background())
}

func (i MediaGraphTlsValidationOptionsResponseArgs) ToMediaGraphTlsValidationOptionsResponsePtrOutputWithContext(ctx context.Context) MediaGraphTlsValidationOptionsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MediaGraphTlsValidationOptionsResponseOutput).ToMediaGraphTlsValidationOptionsResponsePtrOutputWithContext(ctx)
}









type MediaGraphTlsValidationOptionsResponsePtrInput interface {
	pulumi.Input

	ToMediaGraphTlsValidationOptionsResponsePtrOutput() MediaGraphTlsValidationOptionsResponsePtrOutput
	ToMediaGraphTlsValidationOptionsResponsePtrOutputWithContext(context.Context) MediaGraphTlsValidationOptionsResponsePtrOutput
}

type mediaGraphTlsValidationOptionsResponsePtrType MediaGraphTlsValidationOptionsResponseArgs

func MediaGraphTlsValidationOptionsResponsePtr(v *MediaGraphTlsValidationOptionsResponseArgs) MediaGraphTlsValidationOptionsResponsePtrInput {
	return (*mediaGraphTlsValidationOptionsResponsePtrType)(v)
}

func (*mediaGraphTlsValidationOptionsResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**MediaGraphTlsValidationOptionsResponse)(nil)).Elem()
}

func (i *mediaGraphTlsValidationOptionsResponsePtrType) ToMediaGraphTlsValidationOptionsResponsePtrOutput() MediaGraphTlsValidationOptionsResponsePtrOutput {
	return i.ToMediaGraphTlsValidationOptionsResponsePtrOutputWithContext(context.Background())
}

func (i *mediaGraphTlsValidationOptionsResponsePtrType) ToMediaGraphTlsValidationOptionsResponsePtrOutputWithContext(ctx context.Context) MediaGraphTlsValidationOptionsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MediaGraphTlsValidationOptionsResponsePtrOutput)
}

type MediaGraphTlsValidationOptionsResponseOutput struct{ *pulumi.OutputState }

func (MediaGraphTlsValidationOptionsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MediaGraphTlsValidationOptionsResponse)(nil)).Elem()
}

func (o MediaGraphTlsValidationOptionsResponseOutput) ToMediaGraphTlsValidationOptionsResponseOutput() MediaGraphTlsValidationOptionsResponseOutput {
	return o
}

func (o MediaGraphTlsValidationOptionsResponseOutput) ToMediaGraphTlsValidationOptionsResponseOutputWithContext(ctx context.Context) MediaGraphTlsValidationOptionsResponseOutput {
	return o
}

func (o MediaGraphTlsValidationOptionsResponseOutput) ToMediaGraphTlsValidationOptionsResponsePtrOutput() MediaGraphTlsValidationOptionsResponsePtrOutput {
	return o.ToMediaGraphTlsValidationOptionsResponsePtrOutputWithContext(context.Background())
}

func (o MediaGraphTlsValidationOptionsResponseOutput) ToMediaGraphTlsValidationOptionsResponsePtrOutputWithContext(ctx context.Context) MediaGraphTlsValidationOptionsResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v MediaGraphTlsValidationOptionsResponse) *MediaGraphTlsValidationOptionsResponse {
		return &v
	}).(MediaGraphTlsValidationOptionsResponsePtrOutput)
}

func (o MediaGraphTlsValidationOptionsResponseOutput) IgnoreHostname() pulumi.BoolOutput {
	return o.ApplyT(func(v MediaGraphTlsValidationOptionsResponse) bool { return v.IgnoreHostname }).(pulumi.BoolOutput)
}

func (o MediaGraphTlsValidationOptionsResponseOutput) IgnoreSignature() pulumi.BoolOutput {
	return o.ApplyT(func(v MediaGraphTlsValidationOptionsResponse) bool { return v.IgnoreSignature }).(pulumi.BoolOutput)
}

type MediaGraphTlsValidationOptionsResponsePtrOutput struct{ *pulumi.OutputState }

func (MediaGraphTlsValidationOptionsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MediaGraphTlsValidationOptionsResponse)(nil)).Elem()
}

func (o MediaGraphTlsValidationOptionsResponsePtrOutput) ToMediaGraphTlsValidationOptionsResponsePtrOutput() MediaGraphTlsValidationOptionsResponsePtrOutput {
	return o
}

func (o MediaGraphTlsValidationOptionsResponsePtrOutput) ToMediaGraphTlsValidationOptionsResponsePtrOutputWithContext(ctx context.Context) MediaGraphTlsValidationOptionsResponsePtrOutput {
	return o
}

func (o MediaGraphTlsValidationOptionsResponsePtrOutput) Elem() MediaGraphTlsValidationOptionsResponseOutput {
	return o.ApplyT(func(v *MediaGraphTlsValidationOptionsResponse) MediaGraphTlsValidationOptionsResponse {
		if v != nil {
			return *v
		}
		var ret MediaGraphTlsValidationOptionsResponse
		return ret
	}).(MediaGraphTlsValidationOptionsResponseOutput)
}

func (o MediaGraphTlsValidationOptionsResponsePtrOutput) IgnoreHostname() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *MediaGraphTlsValidationOptionsResponse) *bool {
		if v == nil {
			return nil
		}
		return &v.IgnoreHostname
	}).(pulumi.BoolPtrOutput)
}

func (o MediaGraphTlsValidationOptionsResponsePtrOutput) IgnoreSignature() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *MediaGraphTlsValidationOptionsResponse) *bool {
		if v == nil {
			return nil
		}
		return &v.IgnoreSignature
	}).(pulumi.BoolPtrOutput)
}

type MediaGraphUsernamePasswordCredentials struct {
	OdataType string `pulumi:"odataType"`
	Password  string `pulumi:"password"`
	Username  string `pulumi:"username"`
}





type MediaGraphUsernamePasswordCredentialsInput interface {
	pulumi.Input

	ToMediaGraphUsernamePasswordCredentialsOutput() MediaGraphUsernamePasswordCredentialsOutput
	ToMediaGraphUsernamePasswordCredentialsOutputWithContext(context.Context) MediaGraphUsernamePasswordCredentialsOutput
}

type MediaGraphUsernamePasswordCredentialsArgs struct {
	OdataType pulumi.StringInput `pulumi:"odataType"`
	Password  pulumi.StringInput `pulumi:"password"`
	Username  pulumi.StringInput `pulumi:"username"`
}

func (MediaGraphUsernamePasswordCredentialsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MediaGraphUsernamePasswordCredentials)(nil)).Elem()
}

func (i MediaGraphUsernamePasswordCredentialsArgs) ToMediaGraphUsernamePasswordCredentialsOutput() MediaGraphUsernamePasswordCredentialsOutput {
	return i.ToMediaGraphUsernamePasswordCredentialsOutputWithContext(context.Background())
}

func (i MediaGraphUsernamePasswordCredentialsArgs) ToMediaGraphUsernamePasswordCredentialsOutputWithContext(ctx context.Context) MediaGraphUsernamePasswordCredentialsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MediaGraphUsernamePasswordCredentialsOutput)
}

func (i MediaGraphUsernamePasswordCredentialsArgs) ToMediaGraphUsernamePasswordCredentialsPtrOutput() MediaGraphUsernamePasswordCredentialsPtrOutput {
	return i.ToMediaGraphUsernamePasswordCredentialsPtrOutputWithContext(context.Background())
}

func (i MediaGraphUsernamePasswordCredentialsArgs) ToMediaGraphUsernamePasswordCredentialsPtrOutputWithContext(ctx context.Context) MediaGraphUsernamePasswordCredentialsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MediaGraphUsernamePasswordCredentialsOutput).ToMediaGraphUsernamePasswordCredentialsPtrOutputWithContext(ctx)
}









type MediaGraphUsernamePasswordCredentialsPtrInput interface {
	pulumi.Input

	ToMediaGraphUsernamePasswordCredentialsPtrOutput() MediaGraphUsernamePasswordCredentialsPtrOutput
	ToMediaGraphUsernamePasswordCredentialsPtrOutputWithContext(context.Context) MediaGraphUsernamePasswordCredentialsPtrOutput
}

type mediaGraphUsernamePasswordCredentialsPtrType MediaGraphUsernamePasswordCredentialsArgs

func MediaGraphUsernamePasswordCredentialsPtr(v *MediaGraphUsernamePasswordCredentialsArgs) MediaGraphUsernamePasswordCredentialsPtrInput {
	return (*mediaGraphUsernamePasswordCredentialsPtrType)(v)
}

func (*mediaGraphUsernamePasswordCredentialsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**MediaGraphUsernamePasswordCredentials)(nil)).Elem()
}

func (i *mediaGraphUsernamePasswordCredentialsPtrType) ToMediaGraphUsernamePasswordCredentialsPtrOutput() MediaGraphUsernamePasswordCredentialsPtrOutput {
	return i.ToMediaGraphUsernamePasswordCredentialsPtrOutputWithContext(context.Background())
}

func (i *mediaGraphUsernamePasswordCredentialsPtrType) ToMediaGraphUsernamePasswordCredentialsPtrOutputWithContext(ctx context.Context) MediaGraphUsernamePasswordCredentialsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MediaGraphUsernamePasswordCredentialsPtrOutput)
}

type MediaGraphUsernamePasswordCredentialsOutput struct{ *pulumi.OutputState }

func (MediaGraphUsernamePasswordCredentialsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MediaGraphUsernamePasswordCredentials)(nil)).Elem()
}

func (o MediaGraphUsernamePasswordCredentialsOutput) ToMediaGraphUsernamePasswordCredentialsOutput() MediaGraphUsernamePasswordCredentialsOutput {
	return o
}

func (o MediaGraphUsernamePasswordCredentialsOutput) ToMediaGraphUsernamePasswordCredentialsOutputWithContext(ctx context.Context) MediaGraphUsernamePasswordCredentialsOutput {
	return o
}

func (o MediaGraphUsernamePasswordCredentialsOutput) ToMediaGraphUsernamePasswordCredentialsPtrOutput() MediaGraphUsernamePasswordCredentialsPtrOutput {
	return o.ToMediaGraphUsernamePasswordCredentialsPtrOutputWithContext(context.Background())
}

func (o MediaGraphUsernamePasswordCredentialsOutput) ToMediaGraphUsernamePasswordCredentialsPtrOutputWithContext(ctx context.Context) MediaGraphUsernamePasswordCredentialsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v MediaGraphUsernamePasswordCredentials) *MediaGraphUsernamePasswordCredentials {
		return &v
	}).(MediaGraphUsernamePasswordCredentialsPtrOutput)
}

func (o MediaGraphUsernamePasswordCredentialsOutput) OdataType() pulumi.StringOutput {
	return o.ApplyT(func(v MediaGraphUsernamePasswordCredentials) string { return v.OdataType }).(pulumi.StringOutput)
}

func (o MediaGraphUsernamePasswordCredentialsOutput) Password() pulumi.StringOutput {
	return o.ApplyT(func(v MediaGraphUsernamePasswordCredentials) string { return v.Password }).(pulumi.StringOutput)
}

func (o MediaGraphUsernamePasswordCredentialsOutput) Username() pulumi.StringOutput {
	return o.ApplyT(func(v MediaGraphUsernamePasswordCredentials) string { return v.Username }).(pulumi.StringOutput)
}

type MediaGraphUsernamePasswordCredentialsPtrOutput struct{ *pulumi.OutputState }

func (MediaGraphUsernamePasswordCredentialsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MediaGraphUsernamePasswordCredentials)(nil)).Elem()
}

func (o MediaGraphUsernamePasswordCredentialsPtrOutput) ToMediaGraphUsernamePasswordCredentialsPtrOutput() MediaGraphUsernamePasswordCredentialsPtrOutput {
	return o
}

func (o MediaGraphUsernamePasswordCredentialsPtrOutput) ToMediaGraphUsernamePasswordCredentialsPtrOutputWithContext(ctx context.Context) MediaGraphUsernamePasswordCredentialsPtrOutput {
	return o
}

func (o MediaGraphUsernamePasswordCredentialsPtrOutput) Elem() MediaGraphUsernamePasswordCredentialsOutput {
	return o.ApplyT(func(v *MediaGraphUsernamePasswordCredentials) MediaGraphUsernamePasswordCredentials {
		if v != nil {
			return *v
		}
		var ret MediaGraphUsernamePasswordCredentials
		return ret
	}).(MediaGraphUsernamePasswordCredentialsOutput)
}

func (o MediaGraphUsernamePasswordCredentialsPtrOutput) OdataType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MediaGraphUsernamePasswordCredentials) *string {
		if v == nil {
			return nil
		}
		return &v.OdataType
	}).(pulumi.StringPtrOutput)
}

func (o MediaGraphUsernamePasswordCredentialsPtrOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MediaGraphUsernamePasswordCredentials) *string {
		if v == nil {
			return nil
		}
		return &v.Password
	}).(pulumi.StringPtrOutput)
}

func (o MediaGraphUsernamePasswordCredentialsPtrOutput) Username() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MediaGraphUsernamePasswordCredentials) *string {
		if v == nil {
			return nil
		}
		return &v.Username
	}).(pulumi.StringPtrOutput)
}

type MediaGraphUsernamePasswordCredentialsResponse struct {
	OdataType string `pulumi:"odataType"`
	Password  string `pulumi:"password"`
	Username  string `pulumi:"username"`
}





type MediaGraphUsernamePasswordCredentialsResponseInput interface {
	pulumi.Input

	ToMediaGraphUsernamePasswordCredentialsResponseOutput() MediaGraphUsernamePasswordCredentialsResponseOutput
	ToMediaGraphUsernamePasswordCredentialsResponseOutputWithContext(context.Context) MediaGraphUsernamePasswordCredentialsResponseOutput
}

type MediaGraphUsernamePasswordCredentialsResponseArgs struct {
	OdataType pulumi.StringInput `pulumi:"odataType"`
	Password  pulumi.StringInput `pulumi:"password"`
	Username  pulumi.StringInput `pulumi:"username"`
}

func (MediaGraphUsernamePasswordCredentialsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MediaGraphUsernamePasswordCredentialsResponse)(nil)).Elem()
}

func (i MediaGraphUsernamePasswordCredentialsResponseArgs) ToMediaGraphUsernamePasswordCredentialsResponseOutput() MediaGraphUsernamePasswordCredentialsResponseOutput {
	return i.ToMediaGraphUsernamePasswordCredentialsResponseOutputWithContext(context.Background())
}

func (i MediaGraphUsernamePasswordCredentialsResponseArgs) ToMediaGraphUsernamePasswordCredentialsResponseOutputWithContext(ctx context.Context) MediaGraphUsernamePasswordCredentialsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MediaGraphUsernamePasswordCredentialsResponseOutput)
}

func (i MediaGraphUsernamePasswordCredentialsResponseArgs) ToMediaGraphUsernamePasswordCredentialsResponsePtrOutput() MediaGraphUsernamePasswordCredentialsResponsePtrOutput {
	return i.ToMediaGraphUsernamePasswordCredentialsResponsePtrOutputWithContext(context.Background())
}

func (i MediaGraphUsernamePasswordCredentialsResponseArgs) ToMediaGraphUsernamePasswordCredentialsResponsePtrOutputWithContext(ctx context.Context) MediaGraphUsernamePasswordCredentialsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MediaGraphUsernamePasswordCredentialsResponseOutput).ToMediaGraphUsernamePasswordCredentialsResponsePtrOutputWithContext(ctx)
}









type MediaGraphUsernamePasswordCredentialsResponsePtrInput interface {
	pulumi.Input

	ToMediaGraphUsernamePasswordCredentialsResponsePtrOutput() MediaGraphUsernamePasswordCredentialsResponsePtrOutput
	ToMediaGraphUsernamePasswordCredentialsResponsePtrOutputWithContext(context.Context) MediaGraphUsernamePasswordCredentialsResponsePtrOutput
}

type mediaGraphUsernamePasswordCredentialsResponsePtrType MediaGraphUsernamePasswordCredentialsResponseArgs

func MediaGraphUsernamePasswordCredentialsResponsePtr(v *MediaGraphUsernamePasswordCredentialsResponseArgs) MediaGraphUsernamePasswordCredentialsResponsePtrInput {
	return (*mediaGraphUsernamePasswordCredentialsResponsePtrType)(v)
}

func (*mediaGraphUsernamePasswordCredentialsResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**MediaGraphUsernamePasswordCredentialsResponse)(nil)).Elem()
}

func (i *mediaGraphUsernamePasswordCredentialsResponsePtrType) ToMediaGraphUsernamePasswordCredentialsResponsePtrOutput() MediaGraphUsernamePasswordCredentialsResponsePtrOutput {
	return i.ToMediaGraphUsernamePasswordCredentialsResponsePtrOutputWithContext(context.Background())
}

func (i *mediaGraphUsernamePasswordCredentialsResponsePtrType) ToMediaGraphUsernamePasswordCredentialsResponsePtrOutputWithContext(ctx context.Context) MediaGraphUsernamePasswordCredentialsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MediaGraphUsernamePasswordCredentialsResponsePtrOutput)
}

type MediaGraphUsernamePasswordCredentialsResponseOutput struct{ *pulumi.OutputState }

func (MediaGraphUsernamePasswordCredentialsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MediaGraphUsernamePasswordCredentialsResponse)(nil)).Elem()
}

func (o MediaGraphUsernamePasswordCredentialsResponseOutput) ToMediaGraphUsernamePasswordCredentialsResponseOutput() MediaGraphUsernamePasswordCredentialsResponseOutput {
	return o
}

func (o MediaGraphUsernamePasswordCredentialsResponseOutput) ToMediaGraphUsernamePasswordCredentialsResponseOutputWithContext(ctx context.Context) MediaGraphUsernamePasswordCredentialsResponseOutput {
	return o
}

func (o MediaGraphUsernamePasswordCredentialsResponseOutput) ToMediaGraphUsernamePasswordCredentialsResponsePtrOutput() MediaGraphUsernamePasswordCredentialsResponsePtrOutput {
	return o.ToMediaGraphUsernamePasswordCredentialsResponsePtrOutputWithContext(context.Background())
}

func (o MediaGraphUsernamePasswordCredentialsResponseOutput) ToMediaGraphUsernamePasswordCredentialsResponsePtrOutputWithContext(ctx context.Context) MediaGraphUsernamePasswordCredentialsResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v MediaGraphUsernamePasswordCredentialsResponse) *MediaGraphUsernamePasswordCredentialsResponse {
		return &v
	}).(MediaGraphUsernamePasswordCredentialsResponsePtrOutput)
}

func (o MediaGraphUsernamePasswordCredentialsResponseOutput) OdataType() pulumi.StringOutput {
	return o.ApplyT(func(v MediaGraphUsernamePasswordCredentialsResponse) string { return v.OdataType }).(pulumi.StringOutput)
}

func (o MediaGraphUsernamePasswordCredentialsResponseOutput) Password() pulumi.StringOutput {
	return o.ApplyT(func(v MediaGraphUsernamePasswordCredentialsResponse) string { return v.Password }).(pulumi.StringOutput)
}

func (o MediaGraphUsernamePasswordCredentialsResponseOutput) Username() pulumi.StringOutput {
	return o.ApplyT(func(v MediaGraphUsernamePasswordCredentialsResponse) string { return v.Username }).(pulumi.StringOutput)
}

type MediaGraphUsernamePasswordCredentialsResponsePtrOutput struct{ *pulumi.OutputState }

func (MediaGraphUsernamePasswordCredentialsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MediaGraphUsernamePasswordCredentialsResponse)(nil)).Elem()
}

func (o MediaGraphUsernamePasswordCredentialsResponsePtrOutput) ToMediaGraphUsernamePasswordCredentialsResponsePtrOutput() MediaGraphUsernamePasswordCredentialsResponsePtrOutput {
	return o
}

func (o MediaGraphUsernamePasswordCredentialsResponsePtrOutput) ToMediaGraphUsernamePasswordCredentialsResponsePtrOutputWithContext(ctx context.Context) MediaGraphUsernamePasswordCredentialsResponsePtrOutput {
	return o
}

func (o MediaGraphUsernamePasswordCredentialsResponsePtrOutput) Elem() MediaGraphUsernamePasswordCredentialsResponseOutput {
	return o.ApplyT(func(v *MediaGraphUsernamePasswordCredentialsResponse) MediaGraphUsernamePasswordCredentialsResponse {
		if v != nil {
			return *v
		}
		var ret MediaGraphUsernamePasswordCredentialsResponse
		return ret
	}).(MediaGraphUsernamePasswordCredentialsResponseOutput)
}

func (o MediaGraphUsernamePasswordCredentialsResponsePtrOutput) OdataType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MediaGraphUsernamePasswordCredentialsResponse) *string {
		if v == nil {
			return nil
		}
		return &v.OdataType
	}).(pulumi.StringPtrOutput)
}

func (o MediaGraphUsernamePasswordCredentialsResponsePtrOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MediaGraphUsernamePasswordCredentialsResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Password
	}).(pulumi.StringPtrOutput)
}

func (o MediaGraphUsernamePasswordCredentialsResponsePtrOutput) Username() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MediaGraphUsernamePasswordCredentialsResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Username
	}).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(MediaGraphAssetSinkOutput{})
	pulumi.RegisterOutputType(MediaGraphAssetSinkArrayOutput{})
	pulumi.RegisterOutputType(MediaGraphAssetSinkResponseOutput{})
	pulumi.RegisterOutputType(MediaGraphAssetSinkResponseArrayOutput{})
	pulumi.RegisterOutputType(MediaGraphClearEndpointOutput{})
	pulumi.RegisterOutputType(MediaGraphClearEndpointResponseOutput{})
	pulumi.RegisterOutputType(MediaGraphPemCertificateListOutput{})
	pulumi.RegisterOutputType(MediaGraphPemCertificateListPtrOutput{})
	pulumi.RegisterOutputType(MediaGraphPemCertificateListResponseOutput{})
	pulumi.RegisterOutputType(MediaGraphPemCertificateListResponsePtrOutput{})
	pulumi.RegisterOutputType(MediaGraphRtspSourceOutput{})
	pulumi.RegisterOutputType(MediaGraphRtspSourceArrayOutput{})
	pulumi.RegisterOutputType(MediaGraphRtspSourceResponseOutput{})
	pulumi.RegisterOutputType(MediaGraphRtspSourceResponseArrayOutput{})
	pulumi.RegisterOutputType(MediaGraphTlsEndpointOutput{})
	pulumi.RegisterOutputType(MediaGraphTlsEndpointResponseOutput{})
	pulumi.RegisterOutputType(MediaGraphTlsValidationOptionsOutput{})
	pulumi.RegisterOutputType(MediaGraphTlsValidationOptionsPtrOutput{})
	pulumi.RegisterOutputType(MediaGraphTlsValidationOptionsResponseOutput{})
	pulumi.RegisterOutputType(MediaGraphTlsValidationOptionsResponsePtrOutput{})
	pulumi.RegisterOutputType(MediaGraphUsernamePasswordCredentialsOutput{})
	pulumi.RegisterOutputType(MediaGraphUsernamePasswordCredentialsPtrOutput{})
	pulumi.RegisterOutputType(MediaGraphUsernamePasswordCredentialsResponseOutput{})
	pulumi.RegisterOutputType(MediaGraphUsernamePasswordCredentialsResponsePtrOutput{})
}
