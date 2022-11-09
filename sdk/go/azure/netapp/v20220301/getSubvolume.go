


package v20220301

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupSubvolume(ctx *pulumi.Context, args *LookupSubvolumeArgs, opts ...pulumi.InvokeOption) (*LookupSubvolumeResult, error) {
	var rv LookupSubvolumeResult
	err := ctx.Invoke("azure-native:netapp/v20220301:getSubvolume", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSubvolumeArgs struct {
	AccountName       string `pulumi:"accountName"`
	PoolName          string `pulumi:"poolName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	SubvolumeName     string `pulumi:"subvolumeName"`
	VolumeName        string `pulumi:"volumeName"`
}


type LookupSubvolumeResult struct {
	Id                string             `pulumi:"id"`
	Name              string             `pulumi:"name"`
	ParentPath        *string            `pulumi:"parentPath"`
	Path              *string            `pulumi:"path"`
	ProvisioningState string             `pulumi:"provisioningState"`
	SystemData        SystemDataResponse `pulumi:"systemData"`
	Type              string             `pulumi:"type"`
}

func LookupSubvolumeOutput(ctx *pulumi.Context, args LookupSubvolumeOutputArgs, opts ...pulumi.InvokeOption) LookupSubvolumeResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSubvolumeResult, error) {
			args := v.(LookupSubvolumeArgs)
			r, err := LookupSubvolume(ctx, &args, opts...)
			var s LookupSubvolumeResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSubvolumeResultOutput)
}

type LookupSubvolumeOutputArgs struct {
	AccountName       pulumi.StringInput `pulumi:"accountName"`
	PoolName          pulumi.StringInput `pulumi:"poolName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	SubvolumeName     pulumi.StringInput `pulumi:"subvolumeName"`
	VolumeName        pulumi.StringInput `pulumi:"volumeName"`
}

func (LookupSubvolumeOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSubvolumeArgs)(nil)).Elem()
}


type LookupSubvolumeResultOutput struct{ *pulumi.OutputState }

func (LookupSubvolumeResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSubvolumeResult)(nil)).Elem()
}

func (o LookupSubvolumeResultOutput) ToLookupSubvolumeResultOutput() LookupSubvolumeResultOutput {
	return o
}

func (o LookupSubvolumeResultOutput) ToLookupSubvolumeResultOutputWithContext(ctx context.Context) LookupSubvolumeResultOutput {
	return o
}

func (o LookupSubvolumeResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSubvolumeResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupSubvolumeResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSubvolumeResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupSubvolumeResultOutput) ParentPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSubvolumeResult) *string { return v.ParentPath }).(pulumi.StringPtrOutput)
}

func (o LookupSubvolumeResultOutput) Path() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSubvolumeResult) *string { return v.Path }).(pulumi.StringPtrOutput)
}

func (o LookupSubvolumeResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSubvolumeResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupSubvolumeResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupSubvolumeResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupSubvolumeResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSubvolumeResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSubvolumeResultOutput{})
}
