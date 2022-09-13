


package v20220402preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupSnapshot(ctx *pulumi.Context, args *LookupSnapshotArgs, opts ...pulumi.InvokeOption) (*LookupSnapshotResult, error) {
	var rv LookupSnapshotResult
	err := ctx.Invoke("azure-native:containerservice/v20220402preview:getSnapshot", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSnapshotArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ResourceName      string `pulumi:"resourceName"`
}


type LookupSnapshotResult struct {
	CreationData      *CreationDataResponse `pulumi:"creationData"`
	EnableFIPS        bool                  `pulumi:"enableFIPS"`
	Id                string                `pulumi:"id"`
	KubernetesVersion string                `pulumi:"kubernetesVersion"`
	Location          string                `pulumi:"location"`
	Name              string                `pulumi:"name"`
	NodeImageVersion  string                `pulumi:"nodeImageVersion"`
	OsSku             string                `pulumi:"osSku"`
	OsType            string                `pulumi:"osType"`
	SnapshotType      *string               `pulumi:"snapshotType"`
	SystemData        SystemDataResponse    `pulumi:"systemData"`
	Tags              map[string]string     `pulumi:"tags"`
	Type              string                `pulumi:"type"`
	VmSize            string                `pulumi:"vmSize"`
}

func LookupSnapshotOutput(ctx *pulumi.Context, args LookupSnapshotOutputArgs, opts ...pulumi.InvokeOption) LookupSnapshotResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSnapshotResult, error) {
			args := v.(LookupSnapshotArgs)
			r, err := LookupSnapshot(ctx, &args, opts...)
			var s LookupSnapshotResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSnapshotResultOutput)
}

type LookupSnapshotOutputArgs struct {
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	ResourceName      pulumi.StringInput `pulumi:"resourceName"`
}

func (LookupSnapshotOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSnapshotArgs)(nil)).Elem()
}


type LookupSnapshotResultOutput struct{ *pulumi.OutputState }

func (LookupSnapshotResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSnapshotResult)(nil)).Elem()
}

func (o LookupSnapshotResultOutput) ToLookupSnapshotResultOutput() LookupSnapshotResultOutput {
	return o
}

func (o LookupSnapshotResultOutput) ToLookupSnapshotResultOutputWithContext(ctx context.Context) LookupSnapshotResultOutput {
	return o
}

func (o LookupSnapshotResultOutput) CreationData() CreationDataResponsePtrOutput {
	return o.ApplyT(func(v LookupSnapshotResult) *CreationDataResponse { return v.CreationData }).(CreationDataResponsePtrOutput)
}

func (o LookupSnapshotResultOutput) EnableFIPS() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupSnapshotResult) bool { return v.EnableFIPS }).(pulumi.BoolOutput)
}

func (o LookupSnapshotResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSnapshotResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupSnapshotResultOutput) KubernetesVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSnapshotResult) string { return v.KubernetesVersion }).(pulumi.StringOutput)
}

func (o LookupSnapshotResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSnapshotResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupSnapshotResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSnapshotResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupSnapshotResultOutput) NodeImageVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSnapshotResult) string { return v.NodeImageVersion }).(pulumi.StringOutput)
}

func (o LookupSnapshotResultOutput) OsSku() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSnapshotResult) string { return v.OsSku }).(pulumi.StringOutput)
}

func (o LookupSnapshotResultOutput) OsType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSnapshotResult) string { return v.OsType }).(pulumi.StringOutput)
}

func (o LookupSnapshotResultOutput) SnapshotType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSnapshotResult) *string { return v.SnapshotType }).(pulumi.StringPtrOutput)
}

func (o LookupSnapshotResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupSnapshotResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupSnapshotResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupSnapshotResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupSnapshotResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSnapshotResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupSnapshotResultOutput) VmSize() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSnapshotResult) string { return v.VmSize }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSnapshotResultOutput{})
}
