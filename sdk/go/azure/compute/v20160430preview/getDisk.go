


package v20160430preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


func LookupDisk(ctx *pulumi.Context, args *LookupDiskArgs, opts ...pulumi.InvokeOption) (*LookupDiskResult, error) {
	var rv LookupDiskResult
	err := ctx.Invoke("azure-native:compute/v20160430preview:getDisk", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDiskArgs struct {
	DiskName          string `pulumi:"diskName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupDiskResult struct {
	AccountType        *string                     `pulumi:"accountType"`
	CreationData       CreationDataResponse        `pulumi:"creationData"`
	DiskSizeGB         *int                        `pulumi:"diskSizeGB"`
	EncryptionSettings *EncryptionSettingsResponse `pulumi:"encryptionSettings"`
	Id                 string                      `pulumi:"id"`
	Location           string                      `pulumi:"location"`
	Name               string                      `pulumi:"name"`
	OsType             *string                     `pulumi:"osType"`
	OwnerId            string                      `pulumi:"ownerId"`
	ProvisioningState  string                      `pulumi:"provisioningState"`
	Tags               map[string]string           `pulumi:"tags"`
	TimeCreated        string                      `pulumi:"timeCreated"`
	Type               string                      `pulumi:"type"`
}

func LookupDiskOutput(ctx *pulumi.Context, args LookupDiskOutputArgs, opts ...pulumi.InvokeOption) LookupDiskResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDiskResult, error) {
			args := v.(LookupDiskArgs)
			r, err := LookupDisk(ctx, &args, opts...)
			var s LookupDiskResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupDiskResultOutput)
}

type LookupDiskOutputArgs struct {
	DiskName          pulumi.StringInput `pulumi:"diskName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupDiskOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDiskArgs)(nil)).Elem()
}


type LookupDiskResultOutput struct{ *pulumi.OutputState }

func (LookupDiskResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDiskResult)(nil)).Elem()
}

func (o LookupDiskResultOutput) ToLookupDiskResultOutput() LookupDiskResultOutput {
	return o
}

func (o LookupDiskResultOutput) ToLookupDiskResultOutputWithContext(ctx context.Context) LookupDiskResultOutput {
	return o
}

func (o LookupDiskResultOutput) AccountType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDiskResult) *string { return v.AccountType }).(pulumi.StringPtrOutput)
}

func (o LookupDiskResultOutput) CreationData() CreationDataResponseOutput {
	return o.ApplyT(func(v LookupDiskResult) CreationDataResponse { return v.CreationData }).(CreationDataResponseOutput)
}

func (o LookupDiskResultOutput) DiskSizeGB() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupDiskResult) *int { return v.DiskSizeGB }).(pulumi.IntPtrOutput)
}

func (o LookupDiskResultOutput) EncryptionSettings() EncryptionSettingsResponsePtrOutput {
	return o.ApplyT(func(v LookupDiskResult) *EncryptionSettingsResponse { return v.EncryptionSettings }).(EncryptionSettingsResponsePtrOutput)
}

func (o LookupDiskResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDiskResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupDiskResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDiskResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupDiskResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDiskResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupDiskResultOutput) OsType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDiskResult) *string { return v.OsType }).(pulumi.StringPtrOutput)
}

func (o LookupDiskResultOutput) OwnerId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDiskResult) string { return v.OwnerId }).(pulumi.StringOutput)
}

func (o LookupDiskResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDiskResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupDiskResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupDiskResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupDiskResultOutput) TimeCreated() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDiskResult) string { return v.TimeCreated }).(pulumi.StringOutput)
}

func (o LookupDiskResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDiskResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDiskResultOutput{})
}
