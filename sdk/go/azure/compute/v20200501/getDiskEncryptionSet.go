


package v20200501

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


func LookupDiskEncryptionSet(ctx *pulumi.Context, args *LookupDiskEncryptionSetArgs, opts ...pulumi.InvokeOption) (*LookupDiskEncryptionSetResult, error) {
	var rv LookupDiskEncryptionSetResult
	err := ctx.Invoke("azure-native:compute/v20200501:getDiskEncryptionSet", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDiskEncryptionSetArgs struct {
	DiskEncryptionSetName string `pulumi:"diskEncryptionSetName"`
	ResourceGroupName     string `pulumi:"resourceGroupName"`
}


type LookupDiskEncryptionSetResult struct {
	ActiveKey         *KeyVaultAndKeyReferenceResponse  `pulumi:"activeKey"`
	EncryptionType    *string                           `pulumi:"encryptionType"`
	Id                string                            `pulumi:"id"`
	Identity          *EncryptionSetIdentityResponse    `pulumi:"identity"`
	Location          string                            `pulumi:"location"`
	Name              string                            `pulumi:"name"`
	PreviousKeys      []KeyVaultAndKeyReferenceResponse `pulumi:"previousKeys"`
	ProvisioningState string                            `pulumi:"provisioningState"`
	Tags              map[string]string                 `pulumi:"tags"`
	Type              string                            `pulumi:"type"`
}

func LookupDiskEncryptionSetOutput(ctx *pulumi.Context, args LookupDiskEncryptionSetOutputArgs, opts ...pulumi.InvokeOption) LookupDiskEncryptionSetResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDiskEncryptionSetResult, error) {
			args := v.(LookupDiskEncryptionSetArgs)
			r, err := LookupDiskEncryptionSet(ctx, &args, opts...)
			var s LookupDiskEncryptionSetResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupDiskEncryptionSetResultOutput)
}

type LookupDiskEncryptionSetOutputArgs struct {
	DiskEncryptionSetName pulumi.StringInput `pulumi:"diskEncryptionSetName"`
	ResourceGroupName     pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupDiskEncryptionSetOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDiskEncryptionSetArgs)(nil)).Elem()
}


type LookupDiskEncryptionSetResultOutput struct{ *pulumi.OutputState }

func (LookupDiskEncryptionSetResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDiskEncryptionSetResult)(nil)).Elem()
}

func (o LookupDiskEncryptionSetResultOutput) ToLookupDiskEncryptionSetResultOutput() LookupDiskEncryptionSetResultOutput {
	return o
}

func (o LookupDiskEncryptionSetResultOutput) ToLookupDiskEncryptionSetResultOutputWithContext(ctx context.Context) LookupDiskEncryptionSetResultOutput {
	return o
}

func (o LookupDiskEncryptionSetResultOutput) ActiveKey() KeyVaultAndKeyReferenceResponsePtrOutput {
	return o.ApplyT(func(v LookupDiskEncryptionSetResult) *KeyVaultAndKeyReferenceResponse { return v.ActiveKey }).(KeyVaultAndKeyReferenceResponsePtrOutput)
}

func (o LookupDiskEncryptionSetResultOutput) EncryptionType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDiskEncryptionSetResult) *string { return v.EncryptionType }).(pulumi.StringPtrOutput)
}

func (o LookupDiskEncryptionSetResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDiskEncryptionSetResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupDiskEncryptionSetResultOutput) Identity() EncryptionSetIdentityResponsePtrOutput {
	return o.ApplyT(func(v LookupDiskEncryptionSetResult) *EncryptionSetIdentityResponse { return v.Identity }).(EncryptionSetIdentityResponsePtrOutput)
}

func (o LookupDiskEncryptionSetResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDiskEncryptionSetResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupDiskEncryptionSetResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDiskEncryptionSetResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupDiskEncryptionSetResultOutput) PreviousKeys() KeyVaultAndKeyReferenceResponseArrayOutput {
	return o.ApplyT(func(v LookupDiskEncryptionSetResult) []KeyVaultAndKeyReferenceResponse { return v.PreviousKeys }).(KeyVaultAndKeyReferenceResponseArrayOutput)
}

func (o LookupDiskEncryptionSetResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDiskEncryptionSetResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupDiskEncryptionSetResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupDiskEncryptionSetResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupDiskEncryptionSetResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDiskEncryptionSetResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDiskEncryptionSetResultOutput{})
}
