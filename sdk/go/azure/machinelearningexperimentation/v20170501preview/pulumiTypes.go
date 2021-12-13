


package v20170501preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type StorageAccountProperties struct {
	AccessKey        string `pulumi:"accessKey"`
	StorageAccountId string `pulumi:"storageAccountId"`
}





type StorageAccountPropertiesInput interface {
	pulumi.Input

	ToStorageAccountPropertiesOutput() StorageAccountPropertiesOutput
	ToStorageAccountPropertiesOutputWithContext(context.Context) StorageAccountPropertiesOutput
}

type StorageAccountPropertiesArgs struct {
	AccessKey        pulumi.StringInput `pulumi:"accessKey"`
	StorageAccountId pulumi.StringInput `pulumi:"storageAccountId"`
}

func (StorageAccountPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*StorageAccountProperties)(nil)).Elem()
}

func (i StorageAccountPropertiesArgs) ToStorageAccountPropertiesOutput() StorageAccountPropertiesOutput {
	return i.ToStorageAccountPropertiesOutputWithContext(context.Background())
}

func (i StorageAccountPropertiesArgs) ToStorageAccountPropertiesOutputWithContext(ctx context.Context) StorageAccountPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageAccountPropertiesOutput)
}

type StorageAccountPropertiesOutput struct{ *pulumi.OutputState }

func (StorageAccountPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StorageAccountProperties)(nil)).Elem()
}

func (o StorageAccountPropertiesOutput) ToStorageAccountPropertiesOutput() StorageAccountPropertiesOutput {
	return o
}

func (o StorageAccountPropertiesOutput) ToStorageAccountPropertiesOutputWithContext(ctx context.Context) StorageAccountPropertiesOutput {
	return o
}

func (o StorageAccountPropertiesOutput) AccessKey() pulumi.StringOutput {
	return o.ApplyT(func(v StorageAccountProperties) string { return v.AccessKey }).(pulumi.StringOutput)
}

func (o StorageAccountPropertiesOutput) StorageAccountId() pulumi.StringOutput {
	return o.ApplyT(func(v StorageAccountProperties) string { return v.StorageAccountId }).(pulumi.StringOutput)
}

type StorageAccountPropertiesResponse struct {
	AccessKey        string `pulumi:"accessKey"`
	StorageAccountId string `pulumi:"storageAccountId"`
}

type StorageAccountPropertiesResponseOutput struct{ *pulumi.OutputState }

func (StorageAccountPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StorageAccountPropertiesResponse)(nil)).Elem()
}

func (o StorageAccountPropertiesResponseOutput) ToStorageAccountPropertiesResponseOutput() StorageAccountPropertiesResponseOutput {
	return o
}

func (o StorageAccountPropertiesResponseOutput) ToStorageAccountPropertiesResponseOutputWithContext(ctx context.Context) StorageAccountPropertiesResponseOutput {
	return o
}

func (o StorageAccountPropertiesResponseOutput) AccessKey() pulumi.StringOutput {
	return o.ApplyT(func(v StorageAccountPropertiesResponse) string { return v.AccessKey }).(pulumi.StringOutput)
}

func (o StorageAccountPropertiesResponseOutput) StorageAccountId() pulumi.StringOutput {
	return o.ApplyT(func(v StorageAccountPropertiesResponse) string { return v.StorageAccountId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(StorageAccountPropertiesOutput{})
	pulumi.RegisterOutputType(StorageAccountPropertiesResponseOutput{})
}
