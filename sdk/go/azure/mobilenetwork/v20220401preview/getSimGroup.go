


package v20220401preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupSimGroup(ctx *pulumi.Context, args *LookupSimGroupArgs, opts ...pulumi.InvokeOption) (*LookupSimGroupResult, error) {
	var rv LookupSimGroupResult
	err := ctx.Invoke("azure-native:mobilenetwork/v20220401preview:getSimGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSimGroupArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	SimGroupName      string `pulumi:"simGroupName"`
}


type LookupSimGroupResult struct {
	CreatedAt          *string                          `pulumi:"createdAt"`
	CreatedBy          *string                          `pulumi:"createdBy"`
	CreatedByType      *string                          `pulumi:"createdByType"`
	EncryptionKey      *KeyVaultKeyResponse             `pulumi:"encryptionKey"`
	Id                 string                           `pulumi:"id"`
	Identity           *ManagedServiceIdentityResponse  `pulumi:"identity"`
	LastModifiedAt     *string                          `pulumi:"lastModifiedAt"`
	LastModifiedBy     *string                          `pulumi:"lastModifiedBy"`
	LastModifiedByType *string                          `pulumi:"lastModifiedByType"`
	Location           string                           `pulumi:"location"`
	MobileNetwork      *MobileNetworkResourceIdResponse `pulumi:"mobileNetwork"`
	Name               string                           `pulumi:"name"`
	ProvisioningState  string                           `pulumi:"provisioningState"`
	SystemData         SystemDataResponse               `pulumi:"systemData"`
	Tags               map[string]string                `pulumi:"tags"`
	Type               string                           `pulumi:"type"`
}

func LookupSimGroupOutput(ctx *pulumi.Context, args LookupSimGroupOutputArgs, opts ...pulumi.InvokeOption) LookupSimGroupResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSimGroupResult, error) {
			args := v.(LookupSimGroupArgs)
			r, err := LookupSimGroup(ctx, &args, opts...)
			var s LookupSimGroupResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSimGroupResultOutput)
}

type LookupSimGroupOutputArgs struct {
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	SimGroupName      pulumi.StringInput `pulumi:"simGroupName"`
}

func (LookupSimGroupOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSimGroupArgs)(nil)).Elem()
}


type LookupSimGroupResultOutput struct{ *pulumi.OutputState }

func (LookupSimGroupResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSimGroupResult)(nil)).Elem()
}

func (o LookupSimGroupResultOutput) ToLookupSimGroupResultOutput() LookupSimGroupResultOutput {
	return o
}

func (o LookupSimGroupResultOutput) ToLookupSimGroupResultOutputWithContext(ctx context.Context) LookupSimGroupResultOutput {
	return o
}

func (o LookupSimGroupResultOutput) CreatedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSimGroupResult) *string { return v.CreatedAt }).(pulumi.StringPtrOutput)
}

func (o LookupSimGroupResultOutput) CreatedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSimGroupResult) *string { return v.CreatedBy }).(pulumi.StringPtrOutput)
}

func (o LookupSimGroupResultOutput) CreatedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSimGroupResult) *string { return v.CreatedByType }).(pulumi.StringPtrOutput)
}

func (o LookupSimGroupResultOutput) EncryptionKey() KeyVaultKeyResponsePtrOutput {
	return o.ApplyT(func(v LookupSimGroupResult) *KeyVaultKeyResponse { return v.EncryptionKey }).(KeyVaultKeyResponsePtrOutput)
}

func (o LookupSimGroupResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSimGroupResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupSimGroupResultOutput) Identity() ManagedServiceIdentityResponsePtrOutput {
	return o.ApplyT(func(v LookupSimGroupResult) *ManagedServiceIdentityResponse { return v.Identity }).(ManagedServiceIdentityResponsePtrOutput)
}

func (o LookupSimGroupResultOutput) LastModifiedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSimGroupResult) *string { return v.LastModifiedAt }).(pulumi.StringPtrOutput)
}

func (o LookupSimGroupResultOutput) LastModifiedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSimGroupResult) *string { return v.LastModifiedBy }).(pulumi.StringPtrOutput)
}

func (o LookupSimGroupResultOutput) LastModifiedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSimGroupResult) *string { return v.LastModifiedByType }).(pulumi.StringPtrOutput)
}

func (o LookupSimGroupResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSimGroupResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupSimGroupResultOutput) MobileNetwork() MobileNetworkResourceIdResponsePtrOutput {
	return o.ApplyT(func(v LookupSimGroupResult) *MobileNetworkResourceIdResponse { return v.MobileNetwork }).(MobileNetworkResourceIdResponsePtrOutput)
}

func (o LookupSimGroupResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSimGroupResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupSimGroupResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSimGroupResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupSimGroupResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupSimGroupResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupSimGroupResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupSimGroupResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupSimGroupResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSimGroupResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSimGroupResultOutput{})
}
