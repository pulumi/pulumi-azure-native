


package v20170101

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


func LookupHub(ctx *pulumi.Context, args *LookupHubArgs, opts ...pulumi.InvokeOption) (*LookupHubResult, error) {
	var rv LookupHubResult
	err := ctx.Invoke("azure-native:customerinsights/v20170101:getHub", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupHubArgs struct {
	HubName           string `pulumi:"hubName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupHubResult struct {
	ApiEndpoint       string                        `pulumi:"apiEndpoint"`
	HubBillingInfo    *HubBillingInfoFormatResponse `pulumi:"hubBillingInfo"`
	Id                string                        `pulumi:"id"`
	Location          *string                       `pulumi:"location"`
	Name              string                        `pulumi:"name"`
	ProvisioningState string                        `pulumi:"provisioningState"`
	Tags              map[string]string             `pulumi:"tags"`
	TenantFeatures    *int                          `pulumi:"tenantFeatures"`
	Type              string                        `pulumi:"type"`
	WebEndpoint       string                        `pulumi:"webEndpoint"`
}

func LookupHubOutput(ctx *pulumi.Context, args LookupHubOutputArgs, opts ...pulumi.InvokeOption) LookupHubResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupHubResult, error) {
			args := v.(LookupHubArgs)
			r, err := LookupHub(ctx, &args, opts...)
			var s LookupHubResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupHubResultOutput)
}

type LookupHubOutputArgs struct {
	HubName           pulumi.StringInput `pulumi:"hubName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupHubOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupHubArgs)(nil)).Elem()
}


type LookupHubResultOutput struct{ *pulumi.OutputState }

func (LookupHubResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupHubResult)(nil)).Elem()
}

func (o LookupHubResultOutput) ToLookupHubResultOutput() LookupHubResultOutput {
	return o
}

func (o LookupHubResultOutput) ToLookupHubResultOutputWithContext(ctx context.Context) LookupHubResultOutput {
	return o
}

func (o LookupHubResultOutput) ApiEndpoint() pulumi.StringOutput {
	return o.ApplyT(func(v LookupHubResult) string { return v.ApiEndpoint }).(pulumi.StringOutput)
}

func (o LookupHubResultOutput) HubBillingInfo() HubBillingInfoFormatResponsePtrOutput {
	return o.ApplyT(func(v LookupHubResult) *HubBillingInfoFormatResponse { return v.HubBillingInfo }).(HubBillingInfoFormatResponsePtrOutput)
}

func (o LookupHubResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupHubResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupHubResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupHubResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupHubResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupHubResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupHubResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupHubResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupHubResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupHubResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupHubResultOutput) TenantFeatures() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupHubResult) *int { return v.TenantFeatures }).(pulumi.IntPtrOutput)
}

func (o LookupHubResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupHubResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupHubResultOutput) WebEndpoint() pulumi.StringOutput {
	return o.ApplyT(func(v LookupHubResult) string { return v.WebEndpoint }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupHubResultOutput{})
}
