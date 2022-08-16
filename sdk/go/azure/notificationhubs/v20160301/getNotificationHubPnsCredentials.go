


package v20160301

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetNotificationHubPnsCredentials(ctx *pulumi.Context, args *GetNotificationHubPnsCredentialsArgs, opts ...pulumi.InvokeOption) (*GetNotificationHubPnsCredentialsResult, error) {
	var rv GetNotificationHubPnsCredentialsResult
	err := ctx.Invoke("azure-native:notificationhubs/v20160301:getNotificationHubPnsCredentials", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetNotificationHubPnsCredentialsArgs struct {
	NamespaceName       string `pulumi:"namespaceName"`
	NotificationHubName string `pulumi:"notificationHubName"`
	ResourceGroupName   string `pulumi:"resourceGroupName"`
}


type GetNotificationHubPnsCredentialsResult struct {
	AdmCredential   *AdmCredentialResponse   `pulumi:"admCredential"`
	ApnsCredential  *ApnsCredentialResponse  `pulumi:"apnsCredential"`
	BaiduCredential *BaiduCredentialResponse `pulumi:"baiduCredential"`
	GcmCredential   *GcmCredentialResponse   `pulumi:"gcmCredential"`
	Id              string                   `pulumi:"id"`
	Location        string                   `pulumi:"location"`
	MpnsCredential  *MpnsCredentialResponse  `pulumi:"mpnsCredential"`
	Name            string                   `pulumi:"name"`
	Sku             *SkuResponse             `pulumi:"sku"`
	Tags            map[string]string        `pulumi:"tags"`
	Type            string                   `pulumi:"type"`
	WnsCredential   *WnsCredentialResponse   `pulumi:"wnsCredential"`
}

func GetNotificationHubPnsCredentialsOutput(ctx *pulumi.Context, args GetNotificationHubPnsCredentialsOutputArgs, opts ...pulumi.InvokeOption) GetNotificationHubPnsCredentialsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetNotificationHubPnsCredentialsResult, error) {
			args := v.(GetNotificationHubPnsCredentialsArgs)
			r, err := GetNotificationHubPnsCredentials(ctx, &args, opts...)
			var s GetNotificationHubPnsCredentialsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetNotificationHubPnsCredentialsResultOutput)
}

type GetNotificationHubPnsCredentialsOutputArgs struct {
	NamespaceName       pulumi.StringInput `pulumi:"namespaceName"`
	NotificationHubName pulumi.StringInput `pulumi:"notificationHubName"`
	ResourceGroupName   pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (GetNotificationHubPnsCredentialsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetNotificationHubPnsCredentialsArgs)(nil)).Elem()
}


type GetNotificationHubPnsCredentialsResultOutput struct{ *pulumi.OutputState }

func (GetNotificationHubPnsCredentialsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetNotificationHubPnsCredentialsResult)(nil)).Elem()
}

func (o GetNotificationHubPnsCredentialsResultOutput) ToGetNotificationHubPnsCredentialsResultOutput() GetNotificationHubPnsCredentialsResultOutput {
	return o
}

func (o GetNotificationHubPnsCredentialsResultOutput) ToGetNotificationHubPnsCredentialsResultOutputWithContext(ctx context.Context) GetNotificationHubPnsCredentialsResultOutput {
	return o
}

func (o GetNotificationHubPnsCredentialsResultOutput) AdmCredential() AdmCredentialResponsePtrOutput {
	return o.ApplyT(func(v GetNotificationHubPnsCredentialsResult) *AdmCredentialResponse { return v.AdmCredential }).(AdmCredentialResponsePtrOutput)
}

func (o GetNotificationHubPnsCredentialsResultOutput) ApnsCredential() ApnsCredentialResponsePtrOutput {
	return o.ApplyT(func(v GetNotificationHubPnsCredentialsResult) *ApnsCredentialResponse { return v.ApnsCredential }).(ApnsCredentialResponsePtrOutput)
}

func (o GetNotificationHubPnsCredentialsResultOutput) BaiduCredential() BaiduCredentialResponsePtrOutput {
	return o.ApplyT(func(v GetNotificationHubPnsCredentialsResult) *BaiduCredentialResponse { return v.BaiduCredential }).(BaiduCredentialResponsePtrOutput)
}

func (o GetNotificationHubPnsCredentialsResultOutput) GcmCredential() GcmCredentialResponsePtrOutput {
	return o.ApplyT(func(v GetNotificationHubPnsCredentialsResult) *GcmCredentialResponse { return v.GcmCredential }).(GcmCredentialResponsePtrOutput)
}

func (o GetNotificationHubPnsCredentialsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetNotificationHubPnsCredentialsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetNotificationHubPnsCredentialsResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v GetNotificationHubPnsCredentialsResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o GetNotificationHubPnsCredentialsResultOutput) MpnsCredential() MpnsCredentialResponsePtrOutput {
	return o.ApplyT(func(v GetNotificationHubPnsCredentialsResult) *MpnsCredentialResponse { return v.MpnsCredential }).(MpnsCredentialResponsePtrOutput)
}

func (o GetNotificationHubPnsCredentialsResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetNotificationHubPnsCredentialsResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o GetNotificationHubPnsCredentialsResultOutput) Sku() SkuResponsePtrOutput {
	return o.ApplyT(func(v GetNotificationHubPnsCredentialsResult) *SkuResponse { return v.Sku }).(SkuResponsePtrOutput)
}

func (o GetNotificationHubPnsCredentialsResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v GetNotificationHubPnsCredentialsResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o GetNotificationHubPnsCredentialsResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v GetNotificationHubPnsCredentialsResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o GetNotificationHubPnsCredentialsResultOutput) WnsCredential() WnsCredentialResponsePtrOutput {
	return o.ApplyT(func(v GetNotificationHubPnsCredentialsResult) *WnsCredentialResponse { return v.WnsCredential }).(WnsCredentialResponsePtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetNotificationHubPnsCredentialsResultOutput{})
}
