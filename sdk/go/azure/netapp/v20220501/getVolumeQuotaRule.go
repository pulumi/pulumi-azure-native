


package v20220501

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupVolumeQuotaRule(ctx *pulumi.Context, args *LookupVolumeQuotaRuleArgs, opts ...pulumi.InvokeOption) (*LookupVolumeQuotaRuleResult, error) {
	var rv LookupVolumeQuotaRuleResult
	err := ctx.Invoke("azure-native:netapp/v20220501:getVolumeQuotaRule", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupVolumeQuotaRuleArgs struct {
	AccountName         string `pulumi:"accountName"`
	PoolName            string `pulumi:"poolName"`
	ResourceGroupName   string `pulumi:"resourceGroupName"`
	VolumeName          string `pulumi:"volumeName"`
	VolumeQuotaRuleName string `pulumi:"volumeQuotaRuleName"`
}


type LookupVolumeQuotaRuleResult struct {
	Id                string             `pulumi:"id"`
	Location          string             `pulumi:"location"`
	Name              string             `pulumi:"name"`
	ProvisioningState string             `pulumi:"provisioningState"`
	QuotaSizeInKiBs   *float64           `pulumi:"quotaSizeInKiBs"`
	QuotaTarget       *string            `pulumi:"quotaTarget"`
	QuotaType         *string            `pulumi:"quotaType"`
	SystemData        SystemDataResponse `pulumi:"systemData"`
	Tags              map[string]string  `pulumi:"tags"`
	Type              string             `pulumi:"type"`
}

func LookupVolumeQuotaRuleOutput(ctx *pulumi.Context, args LookupVolumeQuotaRuleOutputArgs, opts ...pulumi.InvokeOption) LookupVolumeQuotaRuleResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupVolumeQuotaRuleResult, error) {
			args := v.(LookupVolumeQuotaRuleArgs)
			r, err := LookupVolumeQuotaRule(ctx, &args, opts...)
			var s LookupVolumeQuotaRuleResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupVolumeQuotaRuleResultOutput)
}

type LookupVolumeQuotaRuleOutputArgs struct {
	AccountName         pulumi.StringInput `pulumi:"accountName"`
	PoolName            pulumi.StringInput `pulumi:"poolName"`
	ResourceGroupName   pulumi.StringInput `pulumi:"resourceGroupName"`
	VolumeName          pulumi.StringInput `pulumi:"volumeName"`
	VolumeQuotaRuleName pulumi.StringInput `pulumi:"volumeQuotaRuleName"`
}

func (LookupVolumeQuotaRuleOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVolumeQuotaRuleArgs)(nil)).Elem()
}


type LookupVolumeQuotaRuleResultOutput struct{ *pulumi.OutputState }

func (LookupVolumeQuotaRuleResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVolumeQuotaRuleResult)(nil)).Elem()
}

func (o LookupVolumeQuotaRuleResultOutput) ToLookupVolumeQuotaRuleResultOutput() LookupVolumeQuotaRuleResultOutput {
	return o
}

func (o LookupVolumeQuotaRuleResultOutput) ToLookupVolumeQuotaRuleResultOutputWithContext(ctx context.Context) LookupVolumeQuotaRuleResultOutput {
	return o
}

func (o LookupVolumeQuotaRuleResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVolumeQuotaRuleResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupVolumeQuotaRuleResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVolumeQuotaRuleResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupVolumeQuotaRuleResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVolumeQuotaRuleResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupVolumeQuotaRuleResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVolumeQuotaRuleResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupVolumeQuotaRuleResultOutput) QuotaSizeInKiBs() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v LookupVolumeQuotaRuleResult) *float64 { return v.QuotaSizeInKiBs }).(pulumi.Float64PtrOutput)
}

func (o LookupVolumeQuotaRuleResultOutput) QuotaTarget() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVolumeQuotaRuleResult) *string { return v.QuotaTarget }).(pulumi.StringPtrOutput)
}

func (o LookupVolumeQuotaRuleResultOutput) QuotaType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVolumeQuotaRuleResult) *string { return v.QuotaType }).(pulumi.StringPtrOutput)
}

func (o LookupVolumeQuotaRuleResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupVolumeQuotaRuleResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupVolumeQuotaRuleResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupVolumeQuotaRuleResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupVolumeQuotaRuleResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVolumeQuotaRuleResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupVolumeQuotaRuleResultOutput{})
}
