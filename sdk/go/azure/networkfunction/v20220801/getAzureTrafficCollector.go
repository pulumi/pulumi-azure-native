


package v20220801

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupAzureTrafficCollector(ctx *pulumi.Context, args *LookupAzureTrafficCollectorArgs, opts ...pulumi.InvokeOption) (*LookupAzureTrafficCollectorResult, error) {
	var rv LookupAzureTrafficCollectorResult
	err := ctx.Invoke("azure-native:networkfunction/v20220801:getAzureTrafficCollector", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupAzureTrafficCollectorArgs struct {
	AzureTrafficCollectorName string `pulumi:"azureTrafficCollectorName"`
	ResourceGroupName         string `pulumi:"resourceGroupName"`
}


type LookupAzureTrafficCollectorResult struct {
	CollectorPolicies []CollectorPolicyResponse         `pulumi:"collectorPolicies"`
	Etag              string                            `pulumi:"etag"`
	Id                string                            `pulumi:"id"`
	Location          *string                           `pulumi:"location"`
	Name              string                            `pulumi:"name"`
	ProvisioningState string                            `pulumi:"provisioningState"`
	SystemData        TrackedResourceResponseSystemData `pulumi:"systemData"`
	Tags              map[string]string                 `pulumi:"tags"`
	Type              string                            `pulumi:"type"`
	VirtualHub        *ResourceReferenceResponse        `pulumi:"virtualHub"`
}

func LookupAzureTrafficCollectorOutput(ctx *pulumi.Context, args LookupAzureTrafficCollectorOutputArgs, opts ...pulumi.InvokeOption) LookupAzureTrafficCollectorResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupAzureTrafficCollectorResult, error) {
			args := v.(LookupAzureTrafficCollectorArgs)
			r, err := LookupAzureTrafficCollector(ctx, &args, opts...)
			var s LookupAzureTrafficCollectorResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupAzureTrafficCollectorResultOutput)
}

type LookupAzureTrafficCollectorOutputArgs struct {
	AzureTrafficCollectorName pulumi.StringInput `pulumi:"azureTrafficCollectorName"`
	ResourceGroupName         pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupAzureTrafficCollectorOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAzureTrafficCollectorArgs)(nil)).Elem()
}


type LookupAzureTrafficCollectorResultOutput struct{ *pulumi.OutputState }

func (LookupAzureTrafficCollectorResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAzureTrafficCollectorResult)(nil)).Elem()
}

func (o LookupAzureTrafficCollectorResultOutput) ToLookupAzureTrafficCollectorResultOutput() LookupAzureTrafficCollectorResultOutput {
	return o
}

func (o LookupAzureTrafficCollectorResultOutput) ToLookupAzureTrafficCollectorResultOutputWithContext(ctx context.Context) LookupAzureTrafficCollectorResultOutput {
	return o
}

func (o LookupAzureTrafficCollectorResultOutput) CollectorPolicies() CollectorPolicyResponseArrayOutput {
	return o.ApplyT(func(v LookupAzureTrafficCollectorResult) []CollectorPolicyResponse { return v.CollectorPolicies }).(CollectorPolicyResponseArrayOutput)
}

func (o LookupAzureTrafficCollectorResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAzureTrafficCollectorResult) string { return v.Etag }).(pulumi.StringOutput)
}

func (o LookupAzureTrafficCollectorResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAzureTrafficCollectorResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupAzureTrafficCollectorResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAzureTrafficCollectorResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupAzureTrafficCollectorResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAzureTrafficCollectorResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupAzureTrafficCollectorResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAzureTrafficCollectorResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupAzureTrafficCollectorResultOutput) SystemData() TrackedResourceResponseSystemDataOutput {
	return o.ApplyT(func(v LookupAzureTrafficCollectorResult) TrackedResourceResponseSystemData { return v.SystemData }).(TrackedResourceResponseSystemDataOutput)
}

func (o LookupAzureTrafficCollectorResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupAzureTrafficCollectorResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupAzureTrafficCollectorResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAzureTrafficCollectorResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupAzureTrafficCollectorResultOutput) VirtualHub() ResourceReferenceResponsePtrOutput {
	return o.ApplyT(func(v LookupAzureTrafficCollectorResult) *ResourceReferenceResponse { return v.VirtualHub }).(ResourceReferenceResponsePtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupAzureTrafficCollectorResultOutput{})
}
