


package v20170601preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupReplication(ctx *pulumi.Context, args *LookupReplicationArgs, opts ...pulumi.InvokeOption) (*LookupReplicationResult, error) {
	var rv LookupReplicationResult
	err := ctx.Invoke("azure-native:containerregistry/v20170601preview:getReplication", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupReplicationArgs struct {
	RegistryName      string `pulumi:"registryName"`
	ReplicationName   string `pulumi:"replicationName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupReplicationResult struct {
	Id                string            `pulumi:"id"`
	Location          string            `pulumi:"location"`
	Name              string            `pulumi:"name"`
	ProvisioningState string            `pulumi:"provisioningState"`
	Status            StatusResponse    `pulumi:"status"`
	Tags              map[string]string `pulumi:"tags"`
	Type              string            `pulumi:"type"`
}

func LookupReplicationOutput(ctx *pulumi.Context, args LookupReplicationOutputArgs, opts ...pulumi.InvokeOption) LookupReplicationResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupReplicationResult, error) {
			args := v.(LookupReplicationArgs)
			r, err := LookupReplication(ctx, &args, opts...)
			var s LookupReplicationResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupReplicationResultOutput)
}

type LookupReplicationOutputArgs struct {
	RegistryName      pulumi.StringInput `pulumi:"registryName"`
	ReplicationName   pulumi.StringInput `pulumi:"replicationName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupReplicationOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupReplicationArgs)(nil)).Elem()
}


type LookupReplicationResultOutput struct{ *pulumi.OutputState }

func (LookupReplicationResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupReplicationResult)(nil)).Elem()
}

func (o LookupReplicationResultOutput) ToLookupReplicationResultOutput() LookupReplicationResultOutput {
	return o
}

func (o LookupReplicationResultOutput) ToLookupReplicationResultOutputWithContext(ctx context.Context) LookupReplicationResultOutput {
	return o
}

func (o LookupReplicationResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupReplicationResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupReplicationResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupReplicationResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupReplicationResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupReplicationResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupReplicationResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupReplicationResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupReplicationResultOutput) Status() StatusResponseOutput {
	return o.ApplyT(func(v LookupReplicationResult) StatusResponse { return v.Status }).(StatusResponseOutput)
}

func (o LookupReplicationResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupReplicationResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupReplicationResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupReplicationResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupReplicationResultOutput{})
}
