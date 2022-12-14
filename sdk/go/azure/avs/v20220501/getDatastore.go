


package v20220501

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupDatastore(ctx *pulumi.Context, args *LookupDatastoreArgs, opts ...pulumi.InvokeOption) (*LookupDatastoreResult, error) {
	var rv LookupDatastoreResult
	err := ctx.Invoke("azure-native:avs/v20220501:getDatastore", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupDatastoreArgs struct {
	ClusterName       string `pulumi:"clusterName"`
	DatastoreName     string `pulumi:"datastoreName"`
	PrivateCloudName  string `pulumi:"privateCloudName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupDatastoreResult struct {
	DiskPoolVolume    *DiskPoolVolumeResponse `pulumi:"diskPoolVolume"`
	Id                string                  `pulumi:"id"`
	Name              string                  `pulumi:"name"`
	NetAppVolume      *NetAppVolumeResponse   `pulumi:"netAppVolume"`
	ProvisioningState string                  `pulumi:"provisioningState"`
	Status            string                  `pulumi:"status"`
	Type              string                  `pulumi:"type"`
}


func (val *LookupDatastoreResult) Defaults() *LookupDatastoreResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.DiskPoolVolume = tmp.DiskPoolVolume.Defaults()

	return &tmp
}

func LookupDatastoreOutput(ctx *pulumi.Context, args LookupDatastoreOutputArgs, opts ...pulumi.InvokeOption) LookupDatastoreResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDatastoreResult, error) {
			args := v.(LookupDatastoreArgs)
			r, err := LookupDatastore(ctx, &args, opts...)
			var s LookupDatastoreResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupDatastoreResultOutput)
}

type LookupDatastoreOutputArgs struct {
	ClusterName       pulumi.StringInput `pulumi:"clusterName"`
	DatastoreName     pulumi.StringInput `pulumi:"datastoreName"`
	PrivateCloudName  pulumi.StringInput `pulumi:"privateCloudName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupDatastoreOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDatastoreArgs)(nil)).Elem()
}


type LookupDatastoreResultOutput struct{ *pulumi.OutputState }

func (LookupDatastoreResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDatastoreResult)(nil)).Elem()
}

func (o LookupDatastoreResultOutput) ToLookupDatastoreResultOutput() LookupDatastoreResultOutput {
	return o
}

func (o LookupDatastoreResultOutput) ToLookupDatastoreResultOutputWithContext(ctx context.Context) LookupDatastoreResultOutput {
	return o
}

func (o LookupDatastoreResultOutput) DiskPoolVolume() DiskPoolVolumeResponsePtrOutput {
	return o.ApplyT(func(v LookupDatastoreResult) *DiskPoolVolumeResponse { return v.DiskPoolVolume }).(DiskPoolVolumeResponsePtrOutput)
}

func (o LookupDatastoreResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatastoreResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupDatastoreResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatastoreResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupDatastoreResultOutput) NetAppVolume() NetAppVolumeResponsePtrOutput {
	return o.ApplyT(func(v LookupDatastoreResult) *NetAppVolumeResponse { return v.NetAppVolume }).(NetAppVolumeResponsePtrOutput)
}

func (o LookupDatastoreResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatastoreResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupDatastoreResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatastoreResult) string { return v.Status }).(pulumi.StringOutput)
}

func (o LookupDatastoreResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatastoreResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDatastoreResultOutput{})
}
