


package v20211120preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupVolume(ctx *pulumi.Context, args *LookupVolumeArgs, opts ...pulumi.InvokeOption) (*LookupVolumeResult, error) {
	var rv LookupVolumeResult
	err := ctx.Invoke("azure-native:elasticsan/v20211120preview:getVolume", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupVolumeArgs struct {
	ElasticSanName    string `pulumi:"elasticSanName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	VolumeGroupName   string `pulumi:"volumeGroupName"`
	VolumeName        string `pulumi:"volumeName"`
}


type LookupVolumeResult struct {
	CreationData  *SourceCreationDataResponse `pulumi:"creationData"`
	Id            string                      `pulumi:"id"`
	Name          string                      `pulumi:"name"`
	SizeGiB       *float64                    `pulumi:"sizeGiB"`
	StorageTarget IscsiTargetInfoResponse     `pulumi:"storageTarget"`
	SystemData    SystemDataResponse          `pulumi:"systemData"`
	Tags          map[string]string           `pulumi:"tags"`
	Type          string                      `pulumi:"type"`
	VolumeId      string                      `pulumi:"volumeId"`
}

func LookupVolumeOutput(ctx *pulumi.Context, args LookupVolumeOutputArgs, opts ...pulumi.InvokeOption) LookupVolumeResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupVolumeResult, error) {
			args := v.(LookupVolumeArgs)
			r, err := LookupVolume(ctx, &args, opts...)
			var s LookupVolumeResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupVolumeResultOutput)
}

type LookupVolumeOutputArgs struct {
	ElasticSanName    pulumi.StringInput `pulumi:"elasticSanName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	VolumeGroupName   pulumi.StringInput `pulumi:"volumeGroupName"`
	VolumeName        pulumi.StringInput `pulumi:"volumeName"`
}

func (LookupVolumeOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVolumeArgs)(nil)).Elem()
}


type LookupVolumeResultOutput struct{ *pulumi.OutputState }

func (LookupVolumeResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVolumeResult)(nil)).Elem()
}

func (o LookupVolumeResultOutput) ToLookupVolumeResultOutput() LookupVolumeResultOutput {
	return o
}

func (o LookupVolumeResultOutput) ToLookupVolumeResultOutputWithContext(ctx context.Context) LookupVolumeResultOutput {
	return o
}

func (o LookupVolumeResultOutput) CreationData() SourceCreationDataResponsePtrOutput {
	return o.ApplyT(func(v LookupVolumeResult) *SourceCreationDataResponse { return v.CreationData }).(SourceCreationDataResponsePtrOutput)
}

func (o LookupVolumeResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVolumeResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupVolumeResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVolumeResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupVolumeResultOutput) SizeGiB() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v LookupVolumeResult) *float64 { return v.SizeGiB }).(pulumi.Float64PtrOutput)
}

func (o LookupVolumeResultOutput) StorageTarget() IscsiTargetInfoResponseOutput {
	return o.ApplyT(func(v LookupVolumeResult) IscsiTargetInfoResponse { return v.StorageTarget }).(IscsiTargetInfoResponseOutput)
}

func (o LookupVolumeResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupVolumeResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupVolumeResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupVolumeResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupVolumeResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVolumeResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupVolumeResultOutput) VolumeId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVolumeResult) string { return v.VolumeId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupVolumeResultOutput{})
}
