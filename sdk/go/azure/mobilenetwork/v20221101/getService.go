


package v20221101

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupService(ctx *pulumi.Context, args *LookupServiceArgs, opts ...pulumi.InvokeOption) (*LookupServiceResult, error) {
	var rv LookupServiceResult
	err := ctx.Invoke("azure-native:mobilenetwork/v20221101:getService", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupServiceArgs struct {
	MobileNetworkName string `pulumi:"mobileNetworkName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ServiceName       string `pulumi:"serviceName"`
}


type LookupServiceResult struct {
	Id                string                         `pulumi:"id"`
	Location          string                         `pulumi:"location"`
	Name              string                         `pulumi:"name"`
	PccRules          []PccRuleConfigurationResponse `pulumi:"pccRules"`
	ProvisioningState string                         `pulumi:"provisioningState"`
	ServicePrecedence int                            `pulumi:"servicePrecedence"`
	ServiceQosPolicy  *QosPolicyResponse             `pulumi:"serviceQosPolicy"`
	SystemData        SystemDataResponse             `pulumi:"systemData"`
	Tags              map[string]string              `pulumi:"tags"`
	Type              string                         `pulumi:"type"`
}


func (val *LookupServiceResult) Defaults() *LookupServiceResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.ServiceQosPolicy = tmp.ServiceQosPolicy.Defaults()

	return &tmp
}

func LookupServiceOutput(ctx *pulumi.Context, args LookupServiceOutputArgs, opts ...pulumi.InvokeOption) LookupServiceResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupServiceResult, error) {
			args := v.(LookupServiceArgs)
			r, err := LookupService(ctx, &args, opts...)
			var s LookupServiceResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupServiceResultOutput)
}

type LookupServiceOutputArgs struct {
	MobileNetworkName pulumi.StringInput `pulumi:"mobileNetworkName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	ServiceName       pulumi.StringInput `pulumi:"serviceName"`
}

func (LookupServiceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupServiceArgs)(nil)).Elem()
}


type LookupServiceResultOutput struct{ *pulumi.OutputState }

func (LookupServiceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupServiceResult)(nil)).Elem()
}

func (o LookupServiceResultOutput) ToLookupServiceResultOutput() LookupServiceResultOutput {
	return o
}

func (o LookupServiceResultOutput) ToLookupServiceResultOutputWithContext(ctx context.Context) LookupServiceResultOutput {
	return o
}

func (o LookupServiceResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServiceResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupServiceResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServiceResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupServiceResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServiceResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupServiceResultOutput) PccRules() PccRuleConfigurationResponseArrayOutput {
	return o.ApplyT(func(v LookupServiceResult) []PccRuleConfigurationResponse { return v.PccRules }).(PccRuleConfigurationResponseArrayOutput)
}

func (o LookupServiceResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServiceResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupServiceResultOutput) ServicePrecedence() pulumi.IntOutput {
	return o.ApplyT(func(v LookupServiceResult) int { return v.ServicePrecedence }).(pulumi.IntOutput)
}

func (o LookupServiceResultOutput) ServiceQosPolicy() QosPolicyResponsePtrOutput {
	return o.ApplyT(func(v LookupServiceResult) *QosPolicyResponse { return v.ServiceQosPolicy }).(QosPolicyResponsePtrOutput)
}

func (o LookupServiceResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupServiceResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupServiceResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupServiceResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupServiceResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServiceResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupServiceResultOutput{})
}
