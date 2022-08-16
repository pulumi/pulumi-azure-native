


package v20170401

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupNotificationHub(ctx *pulumi.Context, args *LookupNotificationHubArgs, opts ...pulumi.InvokeOption) (*LookupNotificationHubResult, error) {
	var rv LookupNotificationHubResult
	err := ctx.Invoke("azure-native:notificationhubs/v20170401:getNotificationHub", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupNotificationHubArgs struct {
	NamespaceName       string `pulumi:"namespaceName"`
	NotificationHubName string `pulumi:"notificationHubName"`
	ResourceGroupName   string `pulumi:"resourceGroupName"`
}


type LookupNotificationHubResult struct {
	AdmCredential      *AdmCredentialResponse                            `pulumi:"admCredential"`
	ApnsCredential     *ApnsCredentialResponse                           `pulumi:"apnsCredential"`
	AuthorizationRules []SharedAccessAuthorizationRulePropertiesResponse `pulumi:"authorizationRules"`
	BaiduCredential    *BaiduCredentialResponse                          `pulumi:"baiduCredential"`
	GcmCredential      *GcmCredentialResponse                            `pulumi:"gcmCredential"`
	Id                 string                                            `pulumi:"id"`
	Location           *string                                           `pulumi:"location"`
	MpnsCredential     *MpnsCredentialResponse                           `pulumi:"mpnsCredential"`
	Name               string                                            `pulumi:"name"`
	RegistrationTtl    *string                                           `pulumi:"registrationTtl"`
	Sku                *SkuResponse                                      `pulumi:"sku"`
	Tags               map[string]string                                 `pulumi:"tags"`
	Type               string                                            `pulumi:"type"`
	WnsCredential      *WnsCredentialResponse                            `pulumi:"wnsCredential"`
}

func LookupNotificationHubOutput(ctx *pulumi.Context, args LookupNotificationHubOutputArgs, opts ...pulumi.InvokeOption) LookupNotificationHubResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupNotificationHubResult, error) {
			args := v.(LookupNotificationHubArgs)
			r, err := LookupNotificationHub(ctx, &args, opts...)
			var s LookupNotificationHubResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupNotificationHubResultOutput)
}

type LookupNotificationHubOutputArgs struct {
	NamespaceName       pulumi.StringInput `pulumi:"namespaceName"`
	NotificationHubName pulumi.StringInput `pulumi:"notificationHubName"`
	ResourceGroupName   pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupNotificationHubOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNotificationHubArgs)(nil)).Elem()
}


type LookupNotificationHubResultOutput struct{ *pulumi.OutputState }

func (LookupNotificationHubResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNotificationHubResult)(nil)).Elem()
}

func (o LookupNotificationHubResultOutput) ToLookupNotificationHubResultOutput() LookupNotificationHubResultOutput {
	return o
}

func (o LookupNotificationHubResultOutput) ToLookupNotificationHubResultOutputWithContext(ctx context.Context) LookupNotificationHubResultOutput {
	return o
}

func (o LookupNotificationHubResultOutput) AdmCredential() AdmCredentialResponsePtrOutput {
	return o.ApplyT(func(v LookupNotificationHubResult) *AdmCredentialResponse { return v.AdmCredential }).(AdmCredentialResponsePtrOutput)
}

func (o LookupNotificationHubResultOutput) ApnsCredential() ApnsCredentialResponsePtrOutput {
	return o.ApplyT(func(v LookupNotificationHubResult) *ApnsCredentialResponse { return v.ApnsCredential }).(ApnsCredentialResponsePtrOutput)
}

func (o LookupNotificationHubResultOutput) AuthorizationRules() SharedAccessAuthorizationRulePropertiesResponseArrayOutput {
	return o.ApplyT(func(v LookupNotificationHubResult) []SharedAccessAuthorizationRulePropertiesResponse {
		return v.AuthorizationRules
	}).(SharedAccessAuthorizationRulePropertiesResponseArrayOutput)
}

func (o LookupNotificationHubResultOutput) BaiduCredential() BaiduCredentialResponsePtrOutput {
	return o.ApplyT(func(v LookupNotificationHubResult) *BaiduCredentialResponse { return v.BaiduCredential }).(BaiduCredentialResponsePtrOutput)
}

func (o LookupNotificationHubResultOutput) GcmCredential() GcmCredentialResponsePtrOutput {
	return o.ApplyT(func(v LookupNotificationHubResult) *GcmCredentialResponse { return v.GcmCredential }).(GcmCredentialResponsePtrOutput)
}

func (o LookupNotificationHubResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNotificationHubResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupNotificationHubResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNotificationHubResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupNotificationHubResultOutput) MpnsCredential() MpnsCredentialResponsePtrOutput {
	return o.ApplyT(func(v LookupNotificationHubResult) *MpnsCredentialResponse { return v.MpnsCredential }).(MpnsCredentialResponsePtrOutput)
}

func (o LookupNotificationHubResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNotificationHubResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupNotificationHubResultOutput) RegistrationTtl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNotificationHubResult) *string { return v.RegistrationTtl }).(pulumi.StringPtrOutput)
}

func (o LookupNotificationHubResultOutput) Sku() SkuResponsePtrOutput {
	return o.ApplyT(func(v LookupNotificationHubResult) *SkuResponse { return v.Sku }).(SkuResponsePtrOutput)
}

func (o LookupNotificationHubResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupNotificationHubResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupNotificationHubResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNotificationHubResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupNotificationHubResultOutput) WnsCredential() WnsCredentialResponsePtrOutput {
	return o.ApplyT(func(v LookupNotificationHubResult) *WnsCredentialResponse { return v.WnsCredential }).(WnsCredentialResponsePtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupNotificationHubResultOutput{})
}
