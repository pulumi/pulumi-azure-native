


package v20191101

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


func LookupStorageTarget(ctx *pulumi.Context, args *LookupStorageTargetArgs, opts ...pulumi.InvokeOption) (*LookupStorageTargetResult, error) {
	var rv LookupStorageTargetResult
	err := ctx.Invoke("azure-native:storagecache/v20191101:getStorageTarget", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupStorageTargetArgs struct {
	CacheName         string `pulumi:"cacheName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	StorageTargetName string `pulumi:"storageTargetName"`
}


type LookupStorageTargetResult struct {
	Clfs              *ClfsTargetResponse         `pulumi:"clfs"`
	Id                string                      `pulumi:"id"`
	Junctions         []NamespaceJunctionResponse `pulumi:"junctions"`
	Name              string                      `pulumi:"name"`
	Nfs3              *Nfs3TargetResponse         `pulumi:"nfs3"`
	ProvisioningState *string                     `pulumi:"provisioningState"`
	TargetType        *string                     `pulumi:"targetType"`
	Type              string                      `pulumi:"type"`
	Unknown           *UnknownTargetResponse      `pulumi:"unknown"`
}

func LookupStorageTargetOutput(ctx *pulumi.Context, args LookupStorageTargetOutputArgs, opts ...pulumi.InvokeOption) LookupStorageTargetResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupStorageTargetResult, error) {
			args := v.(LookupStorageTargetArgs)
			r, err := LookupStorageTarget(ctx, &args, opts...)
			var s LookupStorageTargetResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupStorageTargetResultOutput)
}

type LookupStorageTargetOutputArgs struct {
	CacheName         pulumi.StringInput `pulumi:"cacheName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	StorageTargetName pulumi.StringInput `pulumi:"storageTargetName"`
}

func (LookupStorageTargetOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupStorageTargetArgs)(nil)).Elem()
}


type LookupStorageTargetResultOutput struct{ *pulumi.OutputState }

func (LookupStorageTargetResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupStorageTargetResult)(nil)).Elem()
}

func (o LookupStorageTargetResultOutput) ToLookupStorageTargetResultOutput() LookupStorageTargetResultOutput {
	return o
}

func (o LookupStorageTargetResultOutput) ToLookupStorageTargetResultOutputWithContext(ctx context.Context) LookupStorageTargetResultOutput {
	return o
}

func (o LookupStorageTargetResultOutput) Clfs() ClfsTargetResponsePtrOutput {
	return o.ApplyT(func(v LookupStorageTargetResult) *ClfsTargetResponse { return v.Clfs }).(ClfsTargetResponsePtrOutput)
}

func (o LookupStorageTargetResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStorageTargetResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupStorageTargetResultOutput) Junctions() NamespaceJunctionResponseArrayOutput {
	return o.ApplyT(func(v LookupStorageTargetResult) []NamespaceJunctionResponse { return v.Junctions }).(NamespaceJunctionResponseArrayOutput)
}

func (o LookupStorageTargetResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStorageTargetResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupStorageTargetResultOutput) Nfs3() Nfs3TargetResponsePtrOutput {
	return o.ApplyT(func(v LookupStorageTargetResult) *Nfs3TargetResponse { return v.Nfs3 }).(Nfs3TargetResponsePtrOutput)
}

func (o LookupStorageTargetResultOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupStorageTargetResult) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

func (o LookupStorageTargetResultOutput) TargetType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupStorageTargetResult) *string { return v.TargetType }).(pulumi.StringPtrOutput)
}

func (o LookupStorageTargetResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStorageTargetResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupStorageTargetResultOutput) Unknown() UnknownTargetResponsePtrOutput {
	return o.ApplyT(func(v LookupStorageTargetResult) *UnknownTargetResponse { return v.Unknown }).(UnknownTargetResponsePtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupStorageTargetResultOutput{})
}
