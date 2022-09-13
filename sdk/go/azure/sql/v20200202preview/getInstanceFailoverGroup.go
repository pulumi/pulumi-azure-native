


package v20200202preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupInstanceFailoverGroup(ctx *pulumi.Context, args *LookupInstanceFailoverGroupArgs, opts ...pulumi.InvokeOption) (*LookupInstanceFailoverGroupResult, error) {
	var rv LookupInstanceFailoverGroupResult
	err := ctx.Invoke("azure-native:sql/v20200202preview:getInstanceFailoverGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupInstanceFailoverGroupArgs struct {
	FailoverGroupName string `pulumi:"failoverGroupName"`
	LocationName      string `pulumi:"locationName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupInstanceFailoverGroupResult struct {
	Id                   string                                         `pulumi:"id"`
	ManagedInstancePairs []ManagedInstancePairInfoResponse              `pulumi:"managedInstancePairs"`
	Name                 string                                         `pulumi:"name"`
	PartnerRegions       []PartnerRegionInfoResponse                    `pulumi:"partnerRegions"`
	ReadOnlyEndpoint     *InstanceFailoverGroupReadOnlyEndpointResponse `pulumi:"readOnlyEndpoint"`
	ReadWriteEndpoint    InstanceFailoverGroupReadWriteEndpointResponse `pulumi:"readWriteEndpoint"`
	ReplicationRole      string                                         `pulumi:"replicationRole"`
	ReplicationState     string                                         `pulumi:"replicationState"`
	Type                 string                                         `pulumi:"type"`
}

func LookupInstanceFailoverGroupOutput(ctx *pulumi.Context, args LookupInstanceFailoverGroupOutputArgs, opts ...pulumi.InvokeOption) LookupInstanceFailoverGroupResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupInstanceFailoverGroupResult, error) {
			args := v.(LookupInstanceFailoverGroupArgs)
			r, err := LookupInstanceFailoverGroup(ctx, &args, opts...)
			var s LookupInstanceFailoverGroupResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupInstanceFailoverGroupResultOutput)
}

type LookupInstanceFailoverGroupOutputArgs struct {
	FailoverGroupName pulumi.StringInput `pulumi:"failoverGroupName"`
	LocationName      pulumi.StringInput `pulumi:"locationName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupInstanceFailoverGroupOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupInstanceFailoverGroupArgs)(nil)).Elem()
}


type LookupInstanceFailoverGroupResultOutput struct{ *pulumi.OutputState }

func (LookupInstanceFailoverGroupResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupInstanceFailoverGroupResult)(nil)).Elem()
}

func (o LookupInstanceFailoverGroupResultOutput) ToLookupInstanceFailoverGroupResultOutput() LookupInstanceFailoverGroupResultOutput {
	return o
}

func (o LookupInstanceFailoverGroupResultOutput) ToLookupInstanceFailoverGroupResultOutputWithContext(ctx context.Context) LookupInstanceFailoverGroupResultOutput {
	return o
}

func (o LookupInstanceFailoverGroupResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceFailoverGroupResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupInstanceFailoverGroupResultOutput) ManagedInstancePairs() ManagedInstancePairInfoResponseArrayOutput {
	return o.ApplyT(func(v LookupInstanceFailoverGroupResult) []ManagedInstancePairInfoResponse {
		return v.ManagedInstancePairs
	}).(ManagedInstancePairInfoResponseArrayOutput)
}

func (o LookupInstanceFailoverGroupResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceFailoverGroupResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupInstanceFailoverGroupResultOutput) PartnerRegions() PartnerRegionInfoResponseArrayOutput {
	return o.ApplyT(func(v LookupInstanceFailoverGroupResult) []PartnerRegionInfoResponse { return v.PartnerRegions }).(PartnerRegionInfoResponseArrayOutput)
}

func (o LookupInstanceFailoverGroupResultOutput) ReadOnlyEndpoint() InstanceFailoverGroupReadOnlyEndpointResponsePtrOutput {
	return o.ApplyT(func(v LookupInstanceFailoverGroupResult) *InstanceFailoverGroupReadOnlyEndpointResponse {
		return v.ReadOnlyEndpoint
	}).(InstanceFailoverGroupReadOnlyEndpointResponsePtrOutput)
}

func (o LookupInstanceFailoverGroupResultOutput) ReadWriteEndpoint() InstanceFailoverGroupReadWriteEndpointResponseOutput {
	return o.ApplyT(func(v LookupInstanceFailoverGroupResult) InstanceFailoverGroupReadWriteEndpointResponse {
		return v.ReadWriteEndpoint
	}).(InstanceFailoverGroupReadWriteEndpointResponseOutput)
}

func (o LookupInstanceFailoverGroupResultOutput) ReplicationRole() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceFailoverGroupResult) string { return v.ReplicationRole }).(pulumi.StringOutput)
}

func (o LookupInstanceFailoverGroupResultOutput) ReplicationState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceFailoverGroupResult) string { return v.ReplicationState }).(pulumi.StringOutput)
}

func (o LookupInstanceFailoverGroupResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceFailoverGroupResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupInstanceFailoverGroupResultOutput{})
}
