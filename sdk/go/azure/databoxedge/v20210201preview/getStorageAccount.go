


package v20210201preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupStorageAccount(ctx *pulumi.Context, args *LookupStorageAccountArgs, opts ...pulumi.InvokeOption) (*LookupStorageAccountResult, error) {
	var rv LookupStorageAccountResult
	err := ctx.Invoke("azure-native:databoxedge/v20210201preview:getStorageAccount", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupStorageAccountArgs struct {
	DeviceName         string `pulumi:"deviceName"`
	ResourceGroupName  string `pulumi:"resourceGroupName"`
	StorageAccountName string `pulumi:"storageAccountName"`
}


type LookupStorageAccountResult struct {
	BlobEndpoint               string             `pulumi:"blobEndpoint"`
	ContainerCount             int                `pulumi:"containerCount"`
	DataPolicy                 string             `pulumi:"dataPolicy"`
	Description                *string            `pulumi:"description"`
	Id                         string             `pulumi:"id"`
	Name                       string             `pulumi:"name"`
	StorageAccountCredentialId *string            `pulumi:"storageAccountCredentialId"`
	StorageAccountStatus       *string            `pulumi:"storageAccountStatus"`
	SystemData                 SystemDataResponse `pulumi:"systemData"`
	Type                       string             `pulumi:"type"`
}

func LookupStorageAccountOutput(ctx *pulumi.Context, args LookupStorageAccountOutputArgs, opts ...pulumi.InvokeOption) LookupStorageAccountResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupStorageAccountResult, error) {
			args := v.(LookupStorageAccountArgs)
			r, err := LookupStorageAccount(ctx, &args, opts...)
			var s LookupStorageAccountResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupStorageAccountResultOutput)
}

type LookupStorageAccountOutputArgs struct {
	DeviceName         pulumi.StringInput `pulumi:"deviceName"`
	ResourceGroupName  pulumi.StringInput `pulumi:"resourceGroupName"`
	StorageAccountName pulumi.StringInput `pulumi:"storageAccountName"`
}

func (LookupStorageAccountOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupStorageAccountArgs)(nil)).Elem()
}


type LookupStorageAccountResultOutput struct{ *pulumi.OutputState }

func (LookupStorageAccountResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupStorageAccountResult)(nil)).Elem()
}

func (o LookupStorageAccountResultOutput) ToLookupStorageAccountResultOutput() LookupStorageAccountResultOutput {
	return o
}

func (o LookupStorageAccountResultOutput) ToLookupStorageAccountResultOutputWithContext(ctx context.Context) LookupStorageAccountResultOutput {
	return o
}

func (o LookupStorageAccountResultOutput) BlobEndpoint() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStorageAccountResult) string { return v.BlobEndpoint }).(pulumi.StringOutput)
}

func (o LookupStorageAccountResultOutput) ContainerCount() pulumi.IntOutput {
	return o.ApplyT(func(v LookupStorageAccountResult) int { return v.ContainerCount }).(pulumi.IntOutput)
}

func (o LookupStorageAccountResultOutput) DataPolicy() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStorageAccountResult) string { return v.DataPolicy }).(pulumi.StringOutput)
}

func (o LookupStorageAccountResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupStorageAccountResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupStorageAccountResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStorageAccountResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupStorageAccountResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStorageAccountResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupStorageAccountResultOutput) StorageAccountCredentialId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupStorageAccountResult) *string { return v.StorageAccountCredentialId }).(pulumi.StringPtrOutput)
}

func (o LookupStorageAccountResultOutput) StorageAccountStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupStorageAccountResult) *string { return v.StorageAccountStatus }).(pulumi.StringPtrOutput)
}

func (o LookupStorageAccountResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupStorageAccountResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupStorageAccountResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStorageAccountResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupStorageAccountResultOutput{})
}
