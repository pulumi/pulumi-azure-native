


package v20190301preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


func LookupApplication(ctx *pulumi.Context, args *LookupApplicationArgs, opts ...pulumi.InvokeOption) (*LookupApplicationResult, error) {
	var rv LookupApplicationResult
	err := ctx.Invoke("azure-native:servicefabric/v20190301preview:getApplication", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupApplicationArgs struct {
	ApplicationName   string `pulumi:"applicationName"`
	ClusterName       string `pulumi:"clusterName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupApplicationResult struct {
	Etag                      string                                 `pulumi:"etag"`
	Id                        string                                 `pulumi:"id"`
	Location                  *string                                `pulumi:"location"`
	MaximumNodes              *float64                               `pulumi:"maximumNodes"`
	Metrics                   []ApplicationMetricDescriptionResponse `pulumi:"metrics"`
	MinimumNodes              *float64                               `pulumi:"minimumNodes"`
	Name                      string                                 `pulumi:"name"`
	Parameters                map[string]string                      `pulumi:"parameters"`
	ProvisioningState         string                                 `pulumi:"provisioningState"`
	RemoveApplicationCapacity *bool                                  `pulumi:"removeApplicationCapacity"`
	Tags                      map[string]string                      `pulumi:"tags"`
	Type                      string                                 `pulumi:"type"`
	TypeName                  *string                                `pulumi:"typeName"`
	TypeVersion               *string                                `pulumi:"typeVersion"`
	UpgradePolicy             *ApplicationUpgradePolicyResponse      `pulumi:"upgradePolicy"`
}


func (val *LookupApplicationResult) Defaults() *LookupApplicationResult {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.MaximumNodes) {
		maximumNodes_ := 0.0
		tmp.MaximumNodes = &maximumNodes_
	}
	tmp.UpgradePolicy = tmp.UpgradePolicy.Defaults()

	return &tmp
}

func LookupApplicationOutput(ctx *pulumi.Context, args LookupApplicationOutputArgs, opts ...pulumi.InvokeOption) LookupApplicationResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupApplicationResult, error) {
			args := v.(LookupApplicationArgs)
			r, err := LookupApplication(ctx, &args, opts...)
			var s LookupApplicationResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupApplicationResultOutput)
}

type LookupApplicationOutputArgs struct {
	ApplicationName   pulumi.StringInput `pulumi:"applicationName"`
	ClusterName       pulumi.StringInput `pulumi:"clusterName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupApplicationOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupApplicationArgs)(nil)).Elem()
}


type LookupApplicationResultOutput struct{ *pulumi.OutputState }

func (LookupApplicationResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupApplicationResult)(nil)).Elem()
}

func (o LookupApplicationResultOutput) ToLookupApplicationResultOutput() LookupApplicationResultOutput {
	return o
}

func (o LookupApplicationResultOutput) ToLookupApplicationResultOutputWithContext(ctx context.Context) LookupApplicationResultOutput {
	return o
}

func (o LookupApplicationResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApplicationResult) string { return v.Etag }).(pulumi.StringOutput)
}

func (o LookupApplicationResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApplicationResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupApplicationResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupApplicationResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupApplicationResultOutput) MaximumNodes() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v LookupApplicationResult) *float64 { return v.MaximumNodes }).(pulumi.Float64PtrOutput)
}

func (o LookupApplicationResultOutput) Metrics() ApplicationMetricDescriptionResponseArrayOutput {
	return o.ApplyT(func(v LookupApplicationResult) []ApplicationMetricDescriptionResponse { return v.Metrics }).(ApplicationMetricDescriptionResponseArrayOutput)
}

func (o LookupApplicationResultOutput) MinimumNodes() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v LookupApplicationResult) *float64 { return v.MinimumNodes }).(pulumi.Float64PtrOutput)
}

func (o LookupApplicationResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApplicationResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupApplicationResultOutput) Parameters() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupApplicationResult) map[string]string { return v.Parameters }).(pulumi.StringMapOutput)
}

func (o LookupApplicationResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApplicationResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupApplicationResultOutput) RemoveApplicationCapacity() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupApplicationResult) *bool { return v.RemoveApplicationCapacity }).(pulumi.BoolPtrOutput)
}

func (o LookupApplicationResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupApplicationResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupApplicationResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApplicationResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupApplicationResultOutput) TypeName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupApplicationResult) *string { return v.TypeName }).(pulumi.StringPtrOutput)
}

func (o LookupApplicationResultOutput) TypeVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupApplicationResult) *string { return v.TypeVersion }).(pulumi.StringPtrOutput)
}

func (o LookupApplicationResultOutput) UpgradePolicy() ApplicationUpgradePolicyResponsePtrOutput {
	return o.ApplyT(func(v LookupApplicationResult) *ApplicationUpgradePolicyResponse { return v.UpgradePolicy }).(ApplicationUpgradePolicyResponsePtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupApplicationResultOutput{})
}
