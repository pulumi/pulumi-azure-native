


package hybridcontainerservice

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetagentPool(ctx *pulumi.Context, args *GetagentPoolArgs, opts ...pulumi.InvokeOption) (*GetagentPoolResult, error) {
	var rv GetagentPoolResult
	err := ctx.Invoke("azure-native:hybridcontainerservice:getagentPool", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type GetagentPoolArgs struct {
	AgentPoolName           string `pulumi:"agentPoolName"`
	ProvisionedClustersName string `pulumi:"provisionedClustersName"`
	ResourceGroupName       string `pulumi:"resourceGroupName"`
}


type GetagentPoolResult struct {
	AvailabilityZones    []string                                   `pulumi:"availabilityZones"`
	CloudProviderProfile *CloudProviderProfileResponse              `pulumi:"cloudProviderProfile"`
	Count                *int                                       `pulumi:"count"`
	ExtendedLocation     *AgentPoolResponseExtendedLocation         `pulumi:"extendedLocation"`
	Id                   string                                     `pulumi:"id"`
	Location             *string                                    `pulumi:"location"`
	MaxCount             *int                                       `pulumi:"maxCount"`
	MaxPods              *int                                       `pulumi:"maxPods"`
	MinCount             *int                                       `pulumi:"minCount"`
	Mode                 *string                                    `pulumi:"mode"`
	Name                 string                                     `pulumi:"name"`
	NodeImageVersion     *string                                    `pulumi:"nodeImageVersion"`
	NodeLabels           map[string]string                          `pulumi:"nodeLabels"`
	NodeTaints           []string                                   `pulumi:"nodeTaints"`
	OsType               *string                                    `pulumi:"osType"`
	ProvisioningState    string                                     `pulumi:"provisioningState"`
	Status               *AgentPoolProvisioningStatusResponseStatus `pulumi:"status"`
	SystemData           SystemDataResponse                         `pulumi:"systemData"`
	Tags                 map[string]string                          `pulumi:"tags"`
	Type                 string                                     `pulumi:"type"`
	VmSize               *string                                    `pulumi:"vmSize"`
}


func (val *GetagentPoolResult) Defaults() *GetagentPoolResult {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Count) {
		count_ := 1
		tmp.Count = &count_
	}
	if isZero(tmp.Mode) {
		mode_ := "User"
		tmp.Mode = &mode_
	}
	if isZero(tmp.OsType) {
		osType_ := "Linux"
		tmp.OsType = &osType_
	}
	return &tmp
}

func GetagentPoolOutput(ctx *pulumi.Context, args GetagentPoolOutputArgs, opts ...pulumi.InvokeOption) GetagentPoolResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetagentPoolResult, error) {
			args := v.(GetagentPoolArgs)
			r, err := GetagentPool(ctx, &args, opts...)
			var s GetagentPoolResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetagentPoolResultOutput)
}

type GetagentPoolOutputArgs struct {
	AgentPoolName           pulumi.StringInput `pulumi:"agentPoolName"`
	ProvisionedClustersName pulumi.StringInput `pulumi:"provisionedClustersName"`
	ResourceGroupName       pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (GetagentPoolOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetagentPoolArgs)(nil)).Elem()
}


type GetagentPoolResultOutput struct{ *pulumi.OutputState }

func (GetagentPoolResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetagentPoolResult)(nil)).Elem()
}

func (o GetagentPoolResultOutput) ToGetagentPoolResultOutput() GetagentPoolResultOutput {
	return o
}

func (o GetagentPoolResultOutput) ToGetagentPoolResultOutputWithContext(ctx context.Context) GetagentPoolResultOutput {
	return o
}

func (o GetagentPoolResultOutput) AvailabilityZones() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetagentPoolResult) []string { return v.AvailabilityZones }).(pulumi.StringArrayOutput)
}

func (o GetagentPoolResultOutput) CloudProviderProfile() CloudProviderProfileResponsePtrOutput {
	return o.ApplyT(func(v GetagentPoolResult) *CloudProviderProfileResponse { return v.CloudProviderProfile }).(CloudProviderProfileResponsePtrOutput)
}

func (o GetagentPoolResultOutput) Count() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GetagentPoolResult) *int { return v.Count }).(pulumi.IntPtrOutput)
}

func (o GetagentPoolResultOutput) ExtendedLocation() AgentPoolResponseExtendedLocationPtrOutput {
	return o.ApplyT(func(v GetagentPoolResult) *AgentPoolResponseExtendedLocation { return v.ExtendedLocation }).(AgentPoolResponseExtendedLocationPtrOutput)
}

func (o GetagentPoolResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetagentPoolResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetagentPoolResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetagentPoolResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o GetagentPoolResultOutput) MaxCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GetagentPoolResult) *int { return v.MaxCount }).(pulumi.IntPtrOutput)
}

func (o GetagentPoolResultOutput) MaxPods() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GetagentPoolResult) *int { return v.MaxPods }).(pulumi.IntPtrOutput)
}

func (o GetagentPoolResultOutput) MinCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GetagentPoolResult) *int { return v.MinCount }).(pulumi.IntPtrOutput)
}

func (o GetagentPoolResultOutput) Mode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetagentPoolResult) *string { return v.Mode }).(pulumi.StringPtrOutput)
}

func (o GetagentPoolResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetagentPoolResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o GetagentPoolResultOutput) NodeImageVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetagentPoolResult) *string { return v.NodeImageVersion }).(pulumi.StringPtrOutput)
}

func (o GetagentPoolResultOutput) NodeLabels() pulumi.StringMapOutput {
	return o.ApplyT(func(v GetagentPoolResult) map[string]string { return v.NodeLabels }).(pulumi.StringMapOutput)
}

func (o GetagentPoolResultOutput) NodeTaints() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetagentPoolResult) []string { return v.NodeTaints }).(pulumi.StringArrayOutput)
}

func (o GetagentPoolResultOutput) OsType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetagentPoolResult) *string { return v.OsType }).(pulumi.StringPtrOutput)
}

func (o GetagentPoolResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v GetagentPoolResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o GetagentPoolResultOutput) Status() AgentPoolProvisioningStatusResponseStatusPtrOutput {
	return o.ApplyT(func(v GetagentPoolResult) *AgentPoolProvisioningStatusResponseStatus { return v.Status }).(AgentPoolProvisioningStatusResponseStatusPtrOutput)
}

func (o GetagentPoolResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v GetagentPoolResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o GetagentPoolResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v GetagentPoolResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o GetagentPoolResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v GetagentPoolResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o GetagentPoolResultOutput) VmSize() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetagentPoolResult) *string { return v.VmSize }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetagentPoolResultOutput{})
}
