


package v20220501

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupCollectorPolicy(ctx *pulumi.Context, args *LookupCollectorPolicyArgs, opts ...pulumi.InvokeOption) (*LookupCollectorPolicyResult, error) {
	var rv LookupCollectorPolicyResult
	err := ctx.Invoke("azure-native:networkfunction/v20220501:getCollectorPolicy", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupCollectorPolicyArgs struct {
	AzureTrafficCollectorName string `pulumi:"azureTrafficCollectorName"`
	CollectorPolicyName       string `pulumi:"collectorPolicyName"`
	ResourceGroupName         string `pulumi:"resourceGroupName"`
}


type LookupCollectorPolicyResult struct {
	EmissionPolicies  []EmissionPoliciesPropertiesFormatResponse `pulumi:"emissionPolicies"`
	Etag              string                                     `pulumi:"etag"`
	Id                string                                     `pulumi:"id"`
	IngestionPolicy   *IngestionPolicyPropertiesFormatResponse   `pulumi:"ingestionPolicy"`
	Name              string                                     `pulumi:"name"`
	ProvisioningState string                                     `pulumi:"provisioningState"`
	SystemData        CollectorPolicyResponseSystemData          `pulumi:"systemData"`
	Type              string                                     `pulumi:"type"`
}

func LookupCollectorPolicyOutput(ctx *pulumi.Context, args LookupCollectorPolicyOutputArgs, opts ...pulumi.InvokeOption) LookupCollectorPolicyResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupCollectorPolicyResult, error) {
			args := v.(LookupCollectorPolicyArgs)
			r, err := LookupCollectorPolicy(ctx, &args, opts...)
			var s LookupCollectorPolicyResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupCollectorPolicyResultOutput)
}

type LookupCollectorPolicyOutputArgs struct {
	AzureTrafficCollectorName pulumi.StringInput `pulumi:"azureTrafficCollectorName"`
	CollectorPolicyName       pulumi.StringInput `pulumi:"collectorPolicyName"`
	ResourceGroupName         pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupCollectorPolicyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCollectorPolicyArgs)(nil)).Elem()
}


type LookupCollectorPolicyResultOutput struct{ *pulumi.OutputState }

func (LookupCollectorPolicyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCollectorPolicyResult)(nil)).Elem()
}

func (o LookupCollectorPolicyResultOutput) ToLookupCollectorPolicyResultOutput() LookupCollectorPolicyResultOutput {
	return o
}

func (o LookupCollectorPolicyResultOutput) ToLookupCollectorPolicyResultOutputWithContext(ctx context.Context) LookupCollectorPolicyResultOutput {
	return o
}

func (o LookupCollectorPolicyResultOutput) EmissionPolicies() EmissionPoliciesPropertiesFormatResponseArrayOutput {
	return o.ApplyT(func(v LookupCollectorPolicyResult) []EmissionPoliciesPropertiesFormatResponse {
		return v.EmissionPolicies
	}).(EmissionPoliciesPropertiesFormatResponseArrayOutput)
}

func (o LookupCollectorPolicyResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCollectorPolicyResult) string { return v.Etag }).(pulumi.StringOutput)
}

func (o LookupCollectorPolicyResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCollectorPolicyResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupCollectorPolicyResultOutput) IngestionPolicy() IngestionPolicyPropertiesFormatResponsePtrOutput {
	return o.ApplyT(func(v LookupCollectorPolicyResult) *IngestionPolicyPropertiesFormatResponse { return v.IngestionPolicy }).(IngestionPolicyPropertiesFormatResponsePtrOutput)
}

func (o LookupCollectorPolicyResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCollectorPolicyResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupCollectorPolicyResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCollectorPolicyResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupCollectorPolicyResultOutput) SystemData() CollectorPolicyResponseSystemDataOutput {
	return o.ApplyT(func(v LookupCollectorPolicyResult) CollectorPolicyResponseSystemData { return v.SystemData }).(CollectorPolicyResponseSystemDataOutput)
}

func (o LookupCollectorPolicyResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCollectorPolicyResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupCollectorPolicyResultOutput{})
}
