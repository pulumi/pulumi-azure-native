


package v20190901preview

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

type MediaGraphRtspSource struct {
	Credentials *MediaGraphUserCredentials `pulumi:"credentials"`
	Name        string                     `pulumi:"name"`
	OdataType   string                     `pulumi:"odataType"`
	RtspUrl     string                     `pulumi:"rtspUrl"`
}





type MediaGraphRtspSourceInput interface {
	pulumi.Input

	ToMediaGraphRtspSourceOutput() MediaGraphRtspSourceOutput
	ToMediaGraphRtspSourceOutputWithContext(context.Context) MediaGraphRtspSourceOutput
}

type MediaGraphRtspSourceArgs struct {
	Credentials MediaGraphUserCredentialsPtrInput `pulumi:"credentials"`
	Name        pulumi.StringInput                `pulumi:"name"`
	OdataType   pulumi.StringInput                `pulumi:"odataType"`
	RtspUrl     pulumi.StringInput                `pulumi:"rtspUrl"`
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

func (o MediaGraphRtspSourceOutput) Credentials() MediaGraphUserCredentialsPtrOutput {
	return o.ApplyT(func(v MediaGraphRtspSource) *MediaGraphUserCredentials { return v.Credentials }).(MediaGraphUserCredentialsPtrOutput)
}

func (o MediaGraphRtspSourceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v MediaGraphRtspSource) string { return v.Name }).(pulumi.StringOutput)
}

func (o MediaGraphRtspSourceOutput) OdataType() pulumi.StringOutput {
	return o.ApplyT(func(v MediaGraphRtspSource) string { return v.OdataType }).(pulumi.StringOutput)
}

func (o MediaGraphRtspSourceOutput) RtspUrl() pulumi.StringOutput {
	return o.ApplyT(func(v MediaGraphRtspSource) string { return v.RtspUrl }).(pulumi.StringOutput)
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
	Credentials *MediaGraphUserCredentialsResponse `pulumi:"credentials"`
	Name        string                             `pulumi:"name"`
	OdataType   string                             `pulumi:"odataType"`
	RtspUrl     string                             `pulumi:"rtspUrl"`
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

func (o MediaGraphRtspSourceResponseOutput) Credentials() MediaGraphUserCredentialsResponsePtrOutput {
	return o.ApplyT(func(v MediaGraphRtspSourceResponse) *MediaGraphUserCredentialsResponse { return v.Credentials }).(MediaGraphUserCredentialsResponsePtrOutput)
}

func (o MediaGraphRtspSourceResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v MediaGraphRtspSourceResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o MediaGraphRtspSourceResponseOutput) OdataType() pulumi.StringOutput {
	return o.ApplyT(func(v MediaGraphRtspSourceResponse) string { return v.OdataType }).(pulumi.StringOutput)
}

func (o MediaGraphRtspSourceResponseOutput) RtspUrl() pulumi.StringOutput {
	return o.ApplyT(func(v MediaGraphRtspSourceResponse) string { return v.RtspUrl }).(pulumi.StringOutput)
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

type MediaGraphUserCredentials struct {
	Password string `pulumi:"password"`
	Username string `pulumi:"username"`
}





type MediaGraphUserCredentialsInput interface {
	pulumi.Input

	ToMediaGraphUserCredentialsOutput() MediaGraphUserCredentialsOutput
	ToMediaGraphUserCredentialsOutputWithContext(context.Context) MediaGraphUserCredentialsOutput
}

type MediaGraphUserCredentialsArgs struct {
	Password pulumi.StringInput `pulumi:"password"`
	Username pulumi.StringInput `pulumi:"username"`
}

func (MediaGraphUserCredentialsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MediaGraphUserCredentials)(nil)).Elem()
}

func (i MediaGraphUserCredentialsArgs) ToMediaGraphUserCredentialsOutput() MediaGraphUserCredentialsOutput {
	return i.ToMediaGraphUserCredentialsOutputWithContext(context.Background())
}

func (i MediaGraphUserCredentialsArgs) ToMediaGraphUserCredentialsOutputWithContext(ctx context.Context) MediaGraphUserCredentialsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MediaGraphUserCredentialsOutput)
}

func (i MediaGraphUserCredentialsArgs) ToMediaGraphUserCredentialsPtrOutput() MediaGraphUserCredentialsPtrOutput {
	return i.ToMediaGraphUserCredentialsPtrOutputWithContext(context.Background())
}

func (i MediaGraphUserCredentialsArgs) ToMediaGraphUserCredentialsPtrOutputWithContext(ctx context.Context) MediaGraphUserCredentialsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MediaGraphUserCredentialsOutput).ToMediaGraphUserCredentialsPtrOutputWithContext(ctx)
}









type MediaGraphUserCredentialsPtrInput interface {
	pulumi.Input

	ToMediaGraphUserCredentialsPtrOutput() MediaGraphUserCredentialsPtrOutput
	ToMediaGraphUserCredentialsPtrOutputWithContext(context.Context) MediaGraphUserCredentialsPtrOutput
}

type mediaGraphUserCredentialsPtrType MediaGraphUserCredentialsArgs

func MediaGraphUserCredentialsPtr(v *MediaGraphUserCredentialsArgs) MediaGraphUserCredentialsPtrInput {
	return (*mediaGraphUserCredentialsPtrType)(v)
}

func (*mediaGraphUserCredentialsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**MediaGraphUserCredentials)(nil)).Elem()
}

func (i *mediaGraphUserCredentialsPtrType) ToMediaGraphUserCredentialsPtrOutput() MediaGraphUserCredentialsPtrOutput {
	return i.ToMediaGraphUserCredentialsPtrOutputWithContext(context.Background())
}

func (i *mediaGraphUserCredentialsPtrType) ToMediaGraphUserCredentialsPtrOutputWithContext(ctx context.Context) MediaGraphUserCredentialsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MediaGraphUserCredentialsPtrOutput)
}

type MediaGraphUserCredentialsOutput struct{ *pulumi.OutputState }

func (MediaGraphUserCredentialsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MediaGraphUserCredentials)(nil)).Elem()
}

func (o MediaGraphUserCredentialsOutput) ToMediaGraphUserCredentialsOutput() MediaGraphUserCredentialsOutput {
	return o
}

func (o MediaGraphUserCredentialsOutput) ToMediaGraphUserCredentialsOutputWithContext(ctx context.Context) MediaGraphUserCredentialsOutput {
	return o
}

func (o MediaGraphUserCredentialsOutput) ToMediaGraphUserCredentialsPtrOutput() MediaGraphUserCredentialsPtrOutput {
	return o.ToMediaGraphUserCredentialsPtrOutputWithContext(context.Background())
}

func (o MediaGraphUserCredentialsOutput) ToMediaGraphUserCredentialsPtrOutputWithContext(ctx context.Context) MediaGraphUserCredentialsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v MediaGraphUserCredentials) *MediaGraphUserCredentials {
		return &v
	}).(MediaGraphUserCredentialsPtrOutput)
}

func (o MediaGraphUserCredentialsOutput) Password() pulumi.StringOutput {
	return o.ApplyT(func(v MediaGraphUserCredentials) string { return v.Password }).(pulumi.StringOutput)
}

func (o MediaGraphUserCredentialsOutput) Username() pulumi.StringOutput {
	return o.ApplyT(func(v MediaGraphUserCredentials) string { return v.Username }).(pulumi.StringOutput)
}

type MediaGraphUserCredentialsPtrOutput struct{ *pulumi.OutputState }

func (MediaGraphUserCredentialsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MediaGraphUserCredentials)(nil)).Elem()
}

func (o MediaGraphUserCredentialsPtrOutput) ToMediaGraphUserCredentialsPtrOutput() MediaGraphUserCredentialsPtrOutput {
	return o
}

func (o MediaGraphUserCredentialsPtrOutput) ToMediaGraphUserCredentialsPtrOutputWithContext(ctx context.Context) MediaGraphUserCredentialsPtrOutput {
	return o
}

func (o MediaGraphUserCredentialsPtrOutput) Elem() MediaGraphUserCredentialsOutput {
	return o.ApplyT(func(v *MediaGraphUserCredentials) MediaGraphUserCredentials {
		if v != nil {
			return *v
		}
		var ret MediaGraphUserCredentials
		return ret
	}).(MediaGraphUserCredentialsOutput)
}

func (o MediaGraphUserCredentialsPtrOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MediaGraphUserCredentials) *string {
		if v == nil {
			return nil
		}
		return &v.Password
	}).(pulumi.StringPtrOutput)
}

func (o MediaGraphUserCredentialsPtrOutput) Username() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MediaGraphUserCredentials) *string {
		if v == nil {
			return nil
		}
		return &v.Username
	}).(pulumi.StringPtrOutput)
}

type MediaGraphUserCredentialsResponse struct {
	Password string `pulumi:"password"`
	Username string `pulumi:"username"`
}

type MediaGraphUserCredentialsResponseOutput struct{ *pulumi.OutputState }

func (MediaGraphUserCredentialsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MediaGraphUserCredentialsResponse)(nil)).Elem()
}

func (o MediaGraphUserCredentialsResponseOutput) ToMediaGraphUserCredentialsResponseOutput() MediaGraphUserCredentialsResponseOutput {
	return o
}

func (o MediaGraphUserCredentialsResponseOutput) ToMediaGraphUserCredentialsResponseOutputWithContext(ctx context.Context) MediaGraphUserCredentialsResponseOutput {
	return o
}

func (o MediaGraphUserCredentialsResponseOutput) Password() pulumi.StringOutput {
	return o.ApplyT(func(v MediaGraphUserCredentialsResponse) string { return v.Password }).(pulumi.StringOutput)
}

func (o MediaGraphUserCredentialsResponseOutput) Username() pulumi.StringOutput {
	return o.ApplyT(func(v MediaGraphUserCredentialsResponse) string { return v.Username }).(pulumi.StringOutput)
}

type MediaGraphUserCredentialsResponsePtrOutput struct{ *pulumi.OutputState }

func (MediaGraphUserCredentialsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MediaGraphUserCredentialsResponse)(nil)).Elem()
}

func (o MediaGraphUserCredentialsResponsePtrOutput) ToMediaGraphUserCredentialsResponsePtrOutput() MediaGraphUserCredentialsResponsePtrOutput {
	return o
}

func (o MediaGraphUserCredentialsResponsePtrOutput) ToMediaGraphUserCredentialsResponsePtrOutputWithContext(ctx context.Context) MediaGraphUserCredentialsResponsePtrOutput {
	return o
}

func (o MediaGraphUserCredentialsResponsePtrOutput) Elem() MediaGraphUserCredentialsResponseOutput {
	return o.ApplyT(func(v *MediaGraphUserCredentialsResponse) MediaGraphUserCredentialsResponse {
		if v != nil {
			return *v
		}
		var ret MediaGraphUserCredentialsResponse
		return ret
	}).(MediaGraphUserCredentialsResponseOutput)
}

func (o MediaGraphUserCredentialsResponsePtrOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MediaGraphUserCredentialsResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Password
	}).(pulumi.StringPtrOutput)
}

func (o MediaGraphUserCredentialsResponsePtrOutput) Username() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MediaGraphUserCredentialsResponse) *string {
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
	pulumi.RegisterOutputType(MediaGraphRtspSourceOutput{})
	pulumi.RegisterOutputType(MediaGraphRtspSourceArrayOutput{})
	pulumi.RegisterOutputType(MediaGraphRtspSourceResponseOutput{})
	pulumi.RegisterOutputType(MediaGraphRtspSourceResponseArrayOutput{})
	pulumi.RegisterOutputType(MediaGraphUserCredentialsOutput{})
	pulumi.RegisterOutputType(MediaGraphUserCredentialsPtrOutput{})
	pulumi.RegisterOutputType(MediaGraphUserCredentialsResponseOutput{})
	pulumi.RegisterOutputType(MediaGraphUserCredentialsResponsePtrOutput{})
}
