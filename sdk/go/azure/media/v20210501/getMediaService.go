


package v20210501

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupMediaService(ctx *pulumi.Context, args *LookupMediaServiceArgs, opts ...pulumi.InvokeOption) (*LookupMediaServiceResult, error) {
	var rv LookupMediaServiceResult
	err := ctx.Invoke("azure-native:media/v20210501:getMediaService", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupMediaServiceArgs struct {
	AccountName       string `pulumi:"accountName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupMediaServiceResult struct {
	Encryption            *AccountEncryptionResponse    `pulumi:"encryption"`
	Id                    string                        `pulumi:"id"`
	Identity              *MediaServiceIdentityResponse `pulumi:"identity"`
	KeyDelivery           *KeyDeliveryResponse          `pulumi:"keyDelivery"`
	Location              string                        `pulumi:"location"`
	MediaServiceId        string                        `pulumi:"mediaServiceId"`
	Name                  string                        `pulumi:"name"`
	StorageAccounts       []StorageAccountResponse      `pulumi:"storageAccounts"`
	StorageAuthentication *string                       `pulumi:"storageAuthentication"`
	SystemData            SystemDataResponse            `pulumi:"systemData"`
	Tags                  map[string]string             `pulumi:"tags"`
	Type                  string                        `pulumi:"type"`
}

func LookupMediaServiceOutput(ctx *pulumi.Context, args LookupMediaServiceOutputArgs, opts ...pulumi.InvokeOption) LookupMediaServiceResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupMediaServiceResult, error) {
			args := v.(LookupMediaServiceArgs)
			r, err := LookupMediaService(ctx, &args, opts...)
			var s LookupMediaServiceResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupMediaServiceResultOutput)
}

type LookupMediaServiceOutputArgs struct {
	AccountName       pulumi.StringInput `pulumi:"accountName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupMediaServiceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupMediaServiceArgs)(nil)).Elem()
}


type LookupMediaServiceResultOutput struct{ *pulumi.OutputState }

func (LookupMediaServiceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupMediaServiceResult)(nil)).Elem()
}

func (o LookupMediaServiceResultOutput) ToLookupMediaServiceResultOutput() LookupMediaServiceResultOutput {
	return o
}

func (o LookupMediaServiceResultOutput) ToLookupMediaServiceResultOutputWithContext(ctx context.Context) LookupMediaServiceResultOutput {
	return o
}

func (o LookupMediaServiceResultOutput) Encryption() AccountEncryptionResponsePtrOutput {
	return o.ApplyT(func(v LookupMediaServiceResult) *AccountEncryptionResponse { return v.Encryption }).(AccountEncryptionResponsePtrOutput)
}

func (o LookupMediaServiceResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMediaServiceResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupMediaServiceResultOutput) Identity() MediaServiceIdentityResponsePtrOutput {
	return o.ApplyT(func(v LookupMediaServiceResult) *MediaServiceIdentityResponse { return v.Identity }).(MediaServiceIdentityResponsePtrOutput)
}

func (o LookupMediaServiceResultOutput) KeyDelivery() KeyDeliveryResponsePtrOutput {
	return o.ApplyT(func(v LookupMediaServiceResult) *KeyDeliveryResponse { return v.KeyDelivery }).(KeyDeliveryResponsePtrOutput)
}

func (o LookupMediaServiceResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMediaServiceResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupMediaServiceResultOutput) MediaServiceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMediaServiceResult) string { return v.MediaServiceId }).(pulumi.StringOutput)
}

func (o LookupMediaServiceResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMediaServiceResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupMediaServiceResultOutput) StorageAccounts() StorageAccountResponseArrayOutput {
	return o.ApplyT(func(v LookupMediaServiceResult) []StorageAccountResponse { return v.StorageAccounts }).(StorageAccountResponseArrayOutput)
}

func (o LookupMediaServiceResultOutput) StorageAuthentication() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMediaServiceResult) *string { return v.StorageAuthentication }).(pulumi.StringPtrOutput)
}

func (o LookupMediaServiceResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupMediaServiceResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupMediaServiceResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupMediaServiceResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupMediaServiceResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMediaServiceResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupMediaServiceResultOutput{})
}
