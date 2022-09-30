


package v20220101

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListVolumeReplications(ctx *pulumi.Context, args *ListVolumeReplicationsArgs, opts ...pulumi.InvokeOption) (*ListVolumeReplicationsResult, error) {
	var rv ListVolumeReplicationsResult
	err := ctx.Invoke("azure-native:netapp/v20220101:listVolumeReplications", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListVolumeReplicationsArgs struct {
	AccountName       string `pulumi:"accountName"`
	PoolName          string `pulumi:"poolName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	VolumeName        string `pulumi:"volumeName"`
}


type ListVolumeReplicationsResult struct {
	Value []ReplicationResponse `pulumi:"value"`
}

func ListVolumeReplicationsOutput(ctx *pulumi.Context, args ListVolumeReplicationsOutputArgs, opts ...pulumi.InvokeOption) ListVolumeReplicationsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListVolumeReplicationsResult, error) {
			args := v.(ListVolumeReplicationsArgs)
			r, err := ListVolumeReplications(ctx, &args, opts...)
			var s ListVolumeReplicationsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListVolumeReplicationsResultOutput)
}

type ListVolumeReplicationsOutputArgs struct {
	AccountName       pulumi.StringInput `pulumi:"accountName"`
	PoolName          pulumi.StringInput `pulumi:"poolName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	VolumeName        pulumi.StringInput `pulumi:"volumeName"`
}

func (ListVolumeReplicationsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListVolumeReplicationsArgs)(nil)).Elem()
}


type ListVolumeReplicationsResultOutput struct{ *pulumi.OutputState }

func (ListVolumeReplicationsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListVolumeReplicationsResult)(nil)).Elem()
}

func (o ListVolumeReplicationsResultOutput) ToListVolumeReplicationsResultOutput() ListVolumeReplicationsResultOutput {
	return o
}

func (o ListVolumeReplicationsResultOutput) ToListVolumeReplicationsResultOutputWithContext(ctx context.Context) ListVolumeReplicationsResultOutput {
	return o
}

func (o ListVolumeReplicationsResultOutput) Value() ReplicationResponseArrayOutput {
	return o.ApplyT(func(v ListVolumeReplicationsResult) []ReplicationResponse { return v.Value }).(ReplicationResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListVolumeReplicationsResultOutput{})
}
