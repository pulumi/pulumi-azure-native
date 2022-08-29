


package v20220101

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupVolumeGroup(ctx *pulumi.Context, args *LookupVolumeGroupArgs, opts ...pulumi.InvokeOption) (*LookupVolumeGroupResult, error) {
	var rv LookupVolumeGroupResult
	err := ctx.Invoke("azure-native:netapp/v20220101:getVolumeGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupVolumeGroupArgs struct {
	AccountName       string `pulumi:"accountName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	VolumeGroupName   string `pulumi:"volumeGroupName"`
}


type LookupVolumeGroupResult struct {
	GroupMetaData     *VolumeGroupMetaDataResponse          `pulumi:"groupMetaData"`
	Id                string                                `pulumi:"id"`
	Location          *string                               `pulumi:"location"`
	Name              string                                `pulumi:"name"`
	ProvisioningState string                                `pulumi:"provisioningState"`
	Type              string                                `pulumi:"type"`
	Volumes           []VolumeGroupVolumePropertiesResponse `pulumi:"volumes"`
}

func LookupVolumeGroupOutput(ctx *pulumi.Context, args LookupVolumeGroupOutputArgs, opts ...pulumi.InvokeOption) LookupVolumeGroupResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupVolumeGroupResult, error) {
			args := v.(LookupVolumeGroupArgs)
			r, err := LookupVolumeGroup(ctx, &args, opts...)
			var s LookupVolumeGroupResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupVolumeGroupResultOutput)
}

type LookupVolumeGroupOutputArgs struct {
	AccountName       pulumi.StringInput `pulumi:"accountName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	VolumeGroupName   pulumi.StringInput `pulumi:"volumeGroupName"`
}

func (LookupVolumeGroupOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVolumeGroupArgs)(nil)).Elem()
}


type LookupVolumeGroupResultOutput struct{ *pulumi.OutputState }

func (LookupVolumeGroupResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVolumeGroupResult)(nil)).Elem()
}

func (o LookupVolumeGroupResultOutput) ToLookupVolumeGroupResultOutput() LookupVolumeGroupResultOutput {
	return o
}

func (o LookupVolumeGroupResultOutput) ToLookupVolumeGroupResultOutputWithContext(ctx context.Context) LookupVolumeGroupResultOutput {
	return o
}

func (o LookupVolumeGroupResultOutput) GroupMetaData() VolumeGroupMetaDataResponsePtrOutput {
	return o.ApplyT(func(v LookupVolumeGroupResult) *VolumeGroupMetaDataResponse { return v.GroupMetaData }).(VolumeGroupMetaDataResponsePtrOutput)
}

func (o LookupVolumeGroupResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVolumeGroupResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupVolumeGroupResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVolumeGroupResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupVolumeGroupResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVolumeGroupResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupVolumeGroupResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVolumeGroupResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupVolumeGroupResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVolumeGroupResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupVolumeGroupResultOutput) Volumes() VolumeGroupVolumePropertiesResponseArrayOutput {
	return o.ApplyT(func(v LookupVolumeGroupResult) []VolumeGroupVolumePropertiesResponse { return v.Volumes }).(VolumeGroupVolumePropertiesResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupVolumeGroupResultOutput{})
}
