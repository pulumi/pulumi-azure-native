


package v20220702preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupManagedClusterSnapshot(ctx *pulumi.Context, args *LookupManagedClusterSnapshotArgs, opts ...pulumi.InvokeOption) (*LookupManagedClusterSnapshotResult, error) {
	var rv LookupManagedClusterSnapshotResult
	err := ctx.Invoke("azure-native:containerservice/v20220702preview:getManagedClusterSnapshot", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupManagedClusterSnapshotArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ResourceName      string `pulumi:"resourceName"`
}


type LookupManagedClusterSnapshotResult struct {
	CreationData                     *CreationDataResponse                       `pulumi:"creationData"`
	Id                               string                                      `pulumi:"id"`
	Location                         string                                      `pulumi:"location"`
	ManagedClusterPropertiesReadOnly ManagedClusterPropertiesForSnapshotResponse `pulumi:"managedClusterPropertiesReadOnly"`
	Name                             string                                      `pulumi:"name"`
	SnapshotType                     *string                                     `pulumi:"snapshotType"`
	SystemData                       SystemDataResponse                          `pulumi:"systemData"`
	Tags                             map[string]string                           `pulumi:"tags"`
	Type                             string                                      `pulumi:"type"`
}

func LookupManagedClusterSnapshotOutput(ctx *pulumi.Context, args LookupManagedClusterSnapshotOutputArgs, opts ...pulumi.InvokeOption) LookupManagedClusterSnapshotResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupManagedClusterSnapshotResult, error) {
			args := v.(LookupManagedClusterSnapshotArgs)
			r, err := LookupManagedClusterSnapshot(ctx, &args, opts...)
			var s LookupManagedClusterSnapshotResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupManagedClusterSnapshotResultOutput)
}

type LookupManagedClusterSnapshotOutputArgs struct {
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	ResourceName      pulumi.StringInput `pulumi:"resourceName"`
}

func (LookupManagedClusterSnapshotOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupManagedClusterSnapshotArgs)(nil)).Elem()
}


type LookupManagedClusterSnapshotResultOutput struct{ *pulumi.OutputState }

func (LookupManagedClusterSnapshotResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupManagedClusterSnapshotResult)(nil)).Elem()
}

func (o LookupManagedClusterSnapshotResultOutput) ToLookupManagedClusterSnapshotResultOutput() LookupManagedClusterSnapshotResultOutput {
	return o
}

func (o LookupManagedClusterSnapshotResultOutput) ToLookupManagedClusterSnapshotResultOutputWithContext(ctx context.Context) LookupManagedClusterSnapshotResultOutput {
	return o
}

func (o LookupManagedClusterSnapshotResultOutput) CreationData() CreationDataResponsePtrOutput {
	return o.ApplyT(func(v LookupManagedClusterSnapshotResult) *CreationDataResponse { return v.CreationData }).(CreationDataResponsePtrOutput)
}

func (o LookupManagedClusterSnapshotResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagedClusterSnapshotResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupManagedClusterSnapshotResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagedClusterSnapshotResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupManagedClusterSnapshotResultOutput) ManagedClusterPropertiesReadOnly() ManagedClusterPropertiesForSnapshotResponseOutput {
	return o.ApplyT(func(v LookupManagedClusterSnapshotResult) ManagedClusterPropertiesForSnapshotResponse {
		return v.ManagedClusterPropertiesReadOnly
	}).(ManagedClusterPropertiesForSnapshotResponseOutput)
}

func (o LookupManagedClusterSnapshotResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagedClusterSnapshotResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupManagedClusterSnapshotResultOutput) SnapshotType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupManagedClusterSnapshotResult) *string { return v.SnapshotType }).(pulumi.StringPtrOutput)
}

func (o LookupManagedClusterSnapshotResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupManagedClusterSnapshotResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupManagedClusterSnapshotResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupManagedClusterSnapshotResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupManagedClusterSnapshotResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagedClusterSnapshotResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupManagedClusterSnapshotResultOutput{})
}
