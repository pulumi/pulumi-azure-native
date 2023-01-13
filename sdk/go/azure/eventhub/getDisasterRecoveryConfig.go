


package eventhub

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupDisasterRecoveryConfig(ctx *pulumi.Context, args *LookupDisasterRecoveryConfigArgs, opts ...pulumi.InvokeOption) (*LookupDisasterRecoveryConfigResult, error) {
	var rv LookupDisasterRecoveryConfigResult
	err := ctx.Invoke("azure-native:eventhub:getDisasterRecoveryConfig", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDisasterRecoveryConfigArgs struct {
	Alias             string `pulumi:"alias"`
	NamespaceName     string `pulumi:"namespaceName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupDisasterRecoveryConfigResult struct {
	AlternateName                     *string `pulumi:"alternateName"`
	Id                                string  `pulumi:"id"`
	Name                              string  `pulumi:"name"`
	PartnerNamespace                  *string `pulumi:"partnerNamespace"`
	PendingReplicationOperationsCount float64 `pulumi:"pendingReplicationOperationsCount"`
	ProvisioningState                 string  `pulumi:"provisioningState"`
	Role                              string  `pulumi:"role"`
	Type                              string  `pulumi:"type"`
}

func LookupDisasterRecoveryConfigOutput(ctx *pulumi.Context, args LookupDisasterRecoveryConfigOutputArgs, opts ...pulumi.InvokeOption) LookupDisasterRecoveryConfigResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDisasterRecoveryConfigResult, error) {
			args := v.(LookupDisasterRecoveryConfigArgs)
			r, err := LookupDisasterRecoveryConfig(ctx, &args, opts...)
			var s LookupDisasterRecoveryConfigResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupDisasterRecoveryConfigResultOutput)
}

type LookupDisasterRecoveryConfigOutputArgs struct {
	Alias             pulumi.StringInput `pulumi:"alias"`
	NamespaceName     pulumi.StringInput `pulumi:"namespaceName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupDisasterRecoveryConfigOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDisasterRecoveryConfigArgs)(nil)).Elem()
}


type LookupDisasterRecoveryConfigResultOutput struct{ *pulumi.OutputState }

func (LookupDisasterRecoveryConfigResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDisasterRecoveryConfigResult)(nil)).Elem()
}

func (o LookupDisasterRecoveryConfigResultOutput) ToLookupDisasterRecoveryConfigResultOutput() LookupDisasterRecoveryConfigResultOutput {
	return o
}

func (o LookupDisasterRecoveryConfigResultOutput) ToLookupDisasterRecoveryConfigResultOutputWithContext(ctx context.Context) LookupDisasterRecoveryConfigResultOutput {
	return o
}

func (o LookupDisasterRecoveryConfigResultOutput) AlternateName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDisasterRecoveryConfigResult) *string { return v.AlternateName }).(pulumi.StringPtrOutput)
}

func (o LookupDisasterRecoveryConfigResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDisasterRecoveryConfigResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupDisasterRecoveryConfigResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDisasterRecoveryConfigResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupDisasterRecoveryConfigResultOutput) PartnerNamespace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDisasterRecoveryConfigResult) *string { return v.PartnerNamespace }).(pulumi.StringPtrOutput)
}

func (o LookupDisasterRecoveryConfigResultOutput) PendingReplicationOperationsCount() pulumi.Float64Output {
	return o.ApplyT(func(v LookupDisasterRecoveryConfigResult) float64 { return v.PendingReplicationOperationsCount }).(pulumi.Float64Output)
}

func (o LookupDisasterRecoveryConfigResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDisasterRecoveryConfigResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupDisasterRecoveryConfigResultOutput) Role() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDisasterRecoveryConfigResult) string { return v.Role }).(pulumi.StringOutput)
}

func (o LookupDisasterRecoveryConfigResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDisasterRecoveryConfigResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDisasterRecoveryConfigResultOutput{})
}
