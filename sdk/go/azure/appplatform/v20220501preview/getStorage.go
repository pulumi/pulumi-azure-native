


package v20220501preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupStorage(ctx *pulumi.Context, args *LookupStorageArgs, opts ...pulumi.InvokeOption) (*LookupStorageResult, error) {
	var rv LookupStorageResult
	err := ctx.Invoke("azure-native:appplatform/v20220501preview:getStorage", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupStorageArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ServiceName       string `pulumi:"serviceName"`
	StorageName       string `pulumi:"storageName"`
}


type LookupStorageResult struct {
	Id         string                 `pulumi:"id"`
	Name       string                 `pulumi:"name"`
	Properties StorageAccountResponse `pulumi:"properties"`
	SystemData SystemDataResponse     `pulumi:"systemData"`
	Type       string                 `pulumi:"type"`
}

func LookupStorageOutput(ctx *pulumi.Context, args LookupStorageOutputArgs, opts ...pulumi.InvokeOption) LookupStorageResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupStorageResult, error) {
			args := v.(LookupStorageArgs)
			r, err := LookupStorage(ctx, &args, opts...)
			var s LookupStorageResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupStorageResultOutput)
}

type LookupStorageOutputArgs struct {
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	ServiceName       pulumi.StringInput `pulumi:"serviceName"`
	StorageName       pulumi.StringInput `pulumi:"storageName"`
}

func (LookupStorageOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupStorageArgs)(nil)).Elem()
}


type LookupStorageResultOutput struct{ *pulumi.OutputState }

func (LookupStorageResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupStorageResult)(nil)).Elem()
}

func (o LookupStorageResultOutput) ToLookupStorageResultOutput() LookupStorageResultOutput {
	return o
}

func (o LookupStorageResultOutput) ToLookupStorageResultOutputWithContext(ctx context.Context) LookupStorageResultOutput {
	return o
}

func (o LookupStorageResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStorageResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupStorageResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStorageResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupStorageResultOutput) Properties() StorageAccountResponseOutput {
	return o.ApplyT(func(v LookupStorageResult) StorageAccountResponse { return v.Properties }).(StorageAccountResponseOutput)
}

func (o LookupStorageResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupStorageResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupStorageResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStorageResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupStorageResultOutput{})
}
