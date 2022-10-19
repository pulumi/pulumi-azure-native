


package v20210801preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupDistributedAvailabilityGroup(ctx *pulumi.Context, args *LookupDistributedAvailabilityGroupArgs, opts ...pulumi.InvokeOption) (*LookupDistributedAvailabilityGroupResult, error) {
	var rv LookupDistributedAvailabilityGroupResult
	err := ctx.Invoke("azure-native:sql/v20210801preview:getDistributedAvailabilityGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDistributedAvailabilityGroupArgs struct {
	DistributedAvailabilityGroupName string `pulumi:"distributedAvailabilityGroupName"`
	ManagedInstanceName              string `pulumi:"managedInstanceName"`
	ResourceGroupName                string `pulumi:"resourceGroupName"`
}


type LookupDistributedAvailabilityGroupResult struct {
	DistributedAvailabilityGroupId string  `pulumi:"distributedAvailabilityGroupId"`
	Id                             string  `pulumi:"id"`
	LastHardenedLsn                string  `pulumi:"lastHardenedLsn"`
	LinkState                      string  `pulumi:"linkState"`
	Name                           string  `pulumi:"name"`
	PrimaryAvailabilityGroupName   *string `pulumi:"primaryAvailabilityGroupName"`
	ReplicationMode                *string `pulumi:"replicationMode"`
	SecondaryAvailabilityGroupName *string `pulumi:"secondaryAvailabilityGroupName"`
	SourceEndpoint                 *string `pulumi:"sourceEndpoint"`
	SourceReplicaId                string  `pulumi:"sourceReplicaId"`
	TargetDatabase                 *string `pulumi:"targetDatabase"`
	TargetReplicaId                string  `pulumi:"targetReplicaId"`
	Type                           string  `pulumi:"type"`
}

func LookupDistributedAvailabilityGroupOutput(ctx *pulumi.Context, args LookupDistributedAvailabilityGroupOutputArgs, opts ...pulumi.InvokeOption) LookupDistributedAvailabilityGroupResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDistributedAvailabilityGroupResult, error) {
			args := v.(LookupDistributedAvailabilityGroupArgs)
			r, err := LookupDistributedAvailabilityGroup(ctx, &args, opts...)
			var s LookupDistributedAvailabilityGroupResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupDistributedAvailabilityGroupResultOutput)
}

type LookupDistributedAvailabilityGroupOutputArgs struct {
	DistributedAvailabilityGroupName pulumi.StringInput `pulumi:"distributedAvailabilityGroupName"`
	ManagedInstanceName              pulumi.StringInput `pulumi:"managedInstanceName"`
	ResourceGroupName                pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupDistributedAvailabilityGroupOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDistributedAvailabilityGroupArgs)(nil)).Elem()
}


type LookupDistributedAvailabilityGroupResultOutput struct{ *pulumi.OutputState }

func (LookupDistributedAvailabilityGroupResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDistributedAvailabilityGroupResult)(nil)).Elem()
}

func (o LookupDistributedAvailabilityGroupResultOutput) ToLookupDistributedAvailabilityGroupResultOutput() LookupDistributedAvailabilityGroupResultOutput {
	return o
}

func (o LookupDistributedAvailabilityGroupResultOutput) ToLookupDistributedAvailabilityGroupResultOutputWithContext(ctx context.Context) LookupDistributedAvailabilityGroupResultOutput {
	return o
}

func (o LookupDistributedAvailabilityGroupResultOutput) DistributedAvailabilityGroupId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDistributedAvailabilityGroupResult) string { return v.DistributedAvailabilityGroupId }).(pulumi.StringOutput)
}

func (o LookupDistributedAvailabilityGroupResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDistributedAvailabilityGroupResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupDistributedAvailabilityGroupResultOutput) LastHardenedLsn() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDistributedAvailabilityGroupResult) string { return v.LastHardenedLsn }).(pulumi.StringOutput)
}

func (o LookupDistributedAvailabilityGroupResultOutput) LinkState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDistributedAvailabilityGroupResult) string { return v.LinkState }).(pulumi.StringOutput)
}

func (o LookupDistributedAvailabilityGroupResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDistributedAvailabilityGroupResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupDistributedAvailabilityGroupResultOutput) PrimaryAvailabilityGroupName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDistributedAvailabilityGroupResult) *string { return v.PrimaryAvailabilityGroupName }).(pulumi.StringPtrOutput)
}

func (o LookupDistributedAvailabilityGroupResultOutput) ReplicationMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDistributedAvailabilityGroupResult) *string { return v.ReplicationMode }).(pulumi.StringPtrOutput)
}

func (o LookupDistributedAvailabilityGroupResultOutput) SecondaryAvailabilityGroupName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDistributedAvailabilityGroupResult) *string { return v.SecondaryAvailabilityGroupName }).(pulumi.StringPtrOutput)
}

func (o LookupDistributedAvailabilityGroupResultOutput) SourceEndpoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDistributedAvailabilityGroupResult) *string { return v.SourceEndpoint }).(pulumi.StringPtrOutput)
}

func (o LookupDistributedAvailabilityGroupResultOutput) SourceReplicaId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDistributedAvailabilityGroupResult) string { return v.SourceReplicaId }).(pulumi.StringOutput)
}

func (o LookupDistributedAvailabilityGroupResultOutput) TargetDatabase() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDistributedAvailabilityGroupResult) *string { return v.TargetDatabase }).(pulumi.StringPtrOutput)
}

func (o LookupDistributedAvailabilityGroupResultOutput) TargetReplicaId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDistributedAvailabilityGroupResult) string { return v.TargetReplicaId }).(pulumi.StringOutput)
}

func (o LookupDistributedAvailabilityGroupResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDistributedAvailabilityGroupResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDistributedAvailabilityGroupResultOutput{})
}
