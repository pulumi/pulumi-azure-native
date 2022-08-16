


package netapp

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetSubvolumeMetadata(ctx *pulumi.Context, args *GetSubvolumeMetadataArgs, opts ...pulumi.InvokeOption) (*GetSubvolumeMetadataResult, error) {
	var rv GetSubvolumeMetadataResult
	err := ctx.Invoke("azure-native:netapp:getSubvolumeMetadata", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetSubvolumeMetadataArgs struct {
	AccountName       string `pulumi:"accountName"`
	PoolName          string `pulumi:"poolName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	SubvolumeName     string `pulumi:"subvolumeName"`
	VolumeName        string `pulumi:"volumeName"`
}


type GetSubvolumeMetadataResult struct {
	AccessedTimeStamp *string  `pulumi:"accessedTimeStamp"`
	BytesUsed         *float64 `pulumi:"bytesUsed"`
	ChangedTimeStamp  *string  `pulumi:"changedTimeStamp"`
	CreationTimeStamp *string  `pulumi:"creationTimeStamp"`
	Id                string   `pulumi:"id"`
	ModifiedTimeStamp *string  `pulumi:"modifiedTimeStamp"`
	Name              string   `pulumi:"name"`
	ParentPath        *string  `pulumi:"parentPath"`
	Path              *string  `pulumi:"path"`
	Permissions       *string  `pulumi:"permissions"`
	ProvisioningState *string  `pulumi:"provisioningState"`
	Size              *float64 `pulumi:"size"`
	Type              string   `pulumi:"type"`
}

func GetSubvolumeMetadataOutput(ctx *pulumi.Context, args GetSubvolumeMetadataOutputArgs, opts ...pulumi.InvokeOption) GetSubvolumeMetadataResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetSubvolumeMetadataResult, error) {
			args := v.(GetSubvolumeMetadataArgs)
			r, err := GetSubvolumeMetadata(ctx, &args, opts...)
			var s GetSubvolumeMetadataResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetSubvolumeMetadataResultOutput)
}

type GetSubvolumeMetadataOutputArgs struct {
	AccountName       pulumi.StringInput `pulumi:"accountName"`
	PoolName          pulumi.StringInput `pulumi:"poolName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	SubvolumeName     pulumi.StringInput `pulumi:"subvolumeName"`
	VolumeName        pulumi.StringInput `pulumi:"volumeName"`
}

func (GetSubvolumeMetadataOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSubvolumeMetadataArgs)(nil)).Elem()
}


type GetSubvolumeMetadataResultOutput struct{ *pulumi.OutputState }

func (GetSubvolumeMetadataResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSubvolumeMetadataResult)(nil)).Elem()
}

func (o GetSubvolumeMetadataResultOutput) ToGetSubvolumeMetadataResultOutput() GetSubvolumeMetadataResultOutput {
	return o
}

func (o GetSubvolumeMetadataResultOutput) ToGetSubvolumeMetadataResultOutputWithContext(ctx context.Context) GetSubvolumeMetadataResultOutput {
	return o
}

func (o GetSubvolumeMetadataResultOutput) AccessedTimeStamp() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetSubvolumeMetadataResult) *string { return v.AccessedTimeStamp }).(pulumi.StringPtrOutput)
}

func (o GetSubvolumeMetadataResultOutput) BytesUsed() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v GetSubvolumeMetadataResult) *float64 { return v.BytesUsed }).(pulumi.Float64PtrOutput)
}

func (o GetSubvolumeMetadataResultOutput) ChangedTimeStamp() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetSubvolumeMetadataResult) *string { return v.ChangedTimeStamp }).(pulumi.StringPtrOutput)
}

func (o GetSubvolumeMetadataResultOutput) CreationTimeStamp() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetSubvolumeMetadataResult) *string { return v.CreationTimeStamp }).(pulumi.StringPtrOutput)
}

func (o GetSubvolumeMetadataResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetSubvolumeMetadataResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetSubvolumeMetadataResultOutput) ModifiedTimeStamp() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetSubvolumeMetadataResult) *string { return v.ModifiedTimeStamp }).(pulumi.StringPtrOutput)
}

func (o GetSubvolumeMetadataResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetSubvolumeMetadataResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o GetSubvolumeMetadataResultOutput) ParentPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetSubvolumeMetadataResult) *string { return v.ParentPath }).(pulumi.StringPtrOutput)
}

func (o GetSubvolumeMetadataResultOutput) Path() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetSubvolumeMetadataResult) *string { return v.Path }).(pulumi.StringPtrOutput)
}

func (o GetSubvolumeMetadataResultOutput) Permissions() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetSubvolumeMetadataResult) *string { return v.Permissions }).(pulumi.StringPtrOutput)
}

func (o GetSubvolumeMetadataResultOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetSubvolumeMetadataResult) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

func (o GetSubvolumeMetadataResultOutput) Size() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v GetSubvolumeMetadataResult) *float64 { return v.Size }).(pulumi.Float64PtrOutput)
}

func (o GetSubvolumeMetadataResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v GetSubvolumeMetadataResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetSubvolumeMetadataResultOutput{})
}
