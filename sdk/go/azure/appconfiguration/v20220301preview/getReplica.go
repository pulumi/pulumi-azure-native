


package v20220301preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupReplica(ctx *pulumi.Context, args *LookupReplicaArgs, opts ...pulumi.InvokeOption) (*LookupReplicaResult, error) {
	var rv LookupReplicaResult
	err := ctx.Invoke("azure-native:appconfiguration/v20220301preview:getReplica", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupReplicaArgs struct {
	ConfigStoreName   string `pulumi:"configStoreName"`
	ReplicaName       string `pulumi:"replicaName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupReplicaResult struct {
	Endpoint          string             `pulumi:"endpoint"`
	Id                string             `pulumi:"id"`
	Location          *string            `pulumi:"location"`
	Name              string             `pulumi:"name"`
	ProvisioningState string             `pulumi:"provisioningState"`
	SystemData        SystemDataResponse `pulumi:"systemData"`
	Type              string             `pulumi:"type"`
}

func LookupReplicaOutput(ctx *pulumi.Context, args LookupReplicaOutputArgs, opts ...pulumi.InvokeOption) LookupReplicaResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupReplicaResult, error) {
			args := v.(LookupReplicaArgs)
			r, err := LookupReplica(ctx, &args, opts...)
			var s LookupReplicaResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupReplicaResultOutput)
}

type LookupReplicaOutputArgs struct {
	ConfigStoreName   pulumi.StringInput `pulumi:"configStoreName"`
	ReplicaName       pulumi.StringInput `pulumi:"replicaName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupReplicaOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupReplicaArgs)(nil)).Elem()
}


type LookupReplicaResultOutput struct{ *pulumi.OutputState }

func (LookupReplicaResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupReplicaResult)(nil)).Elem()
}

func (o LookupReplicaResultOutput) ToLookupReplicaResultOutput() LookupReplicaResultOutput {
	return o
}

func (o LookupReplicaResultOutput) ToLookupReplicaResultOutputWithContext(ctx context.Context) LookupReplicaResultOutput {
	return o
}

func (o LookupReplicaResultOutput) Endpoint() pulumi.StringOutput {
	return o.ApplyT(func(v LookupReplicaResult) string { return v.Endpoint }).(pulumi.StringOutput)
}

func (o LookupReplicaResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupReplicaResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupReplicaResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupReplicaResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupReplicaResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupReplicaResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupReplicaResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupReplicaResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupReplicaResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupReplicaResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupReplicaResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupReplicaResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupReplicaResultOutput{})
}
