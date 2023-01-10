


package v20190201

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


func LookupAgentPool(ctx *pulumi.Context, args *LookupAgentPoolArgs, opts ...pulumi.InvokeOption) (*LookupAgentPoolResult, error) {
	var rv LookupAgentPoolResult
	err := ctx.Invoke("azure-native:containerservice/v20190201:getAgentPool", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupAgentPoolArgs struct {
	AgentPoolName      string `pulumi:"agentPoolName"`
	ManagedClusterName string `pulumi:"managedClusterName"`
	ResourceGroupName  string `pulumi:"resourceGroupName"`
}


type LookupAgentPoolResult struct {
	AvailabilityZones   []string `pulumi:"availabilityZones"`
	Count               int      `pulumi:"count"`
	EnableAutoScaling   *bool    `pulumi:"enableAutoScaling"`
	Id                  string   `pulumi:"id"`
	MaxCount            *int     `pulumi:"maxCount"`
	MaxPods             *int     `pulumi:"maxPods"`
	MinCount            *int     `pulumi:"minCount"`
	Name                string   `pulumi:"name"`
	OrchestratorVersion *string  `pulumi:"orchestratorVersion"`
	OsDiskSizeGB        *int     `pulumi:"osDiskSizeGB"`
	OsType              *string  `pulumi:"osType"`
	ProvisioningState   string   `pulumi:"provisioningState"`
	Type                string   `pulumi:"type"`
	VmSize              string   `pulumi:"vmSize"`
	VnetSubnetID        *string  `pulumi:"vnetSubnetID"`
}


func (val *LookupAgentPoolResult) Defaults() *LookupAgentPoolResult {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Count) {
		tmp.Count = 1
	}
	return &tmp
}

func LookupAgentPoolOutput(ctx *pulumi.Context, args LookupAgentPoolOutputArgs, opts ...pulumi.InvokeOption) LookupAgentPoolResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupAgentPoolResult, error) {
			args := v.(LookupAgentPoolArgs)
			r, err := LookupAgentPool(ctx, &args, opts...)
			var s LookupAgentPoolResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupAgentPoolResultOutput)
}

type LookupAgentPoolOutputArgs struct {
	AgentPoolName      pulumi.StringInput `pulumi:"agentPoolName"`
	ManagedClusterName pulumi.StringInput `pulumi:"managedClusterName"`
	ResourceGroupName  pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupAgentPoolOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAgentPoolArgs)(nil)).Elem()
}


type LookupAgentPoolResultOutput struct{ *pulumi.OutputState }

func (LookupAgentPoolResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAgentPoolResult)(nil)).Elem()
}

func (o LookupAgentPoolResultOutput) ToLookupAgentPoolResultOutput() LookupAgentPoolResultOutput {
	return o
}

func (o LookupAgentPoolResultOutput) ToLookupAgentPoolResultOutputWithContext(ctx context.Context) LookupAgentPoolResultOutput {
	return o
}

func (o LookupAgentPoolResultOutput) AvailabilityZones() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupAgentPoolResult) []string { return v.AvailabilityZones }).(pulumi.StringArrayOutput)
}

func (o LookupAgentPoolResultOutput) Count() pulumi.IntOutput {
	return o.ApplyT(func(v LookupAgentPoolResult) int { return v.Count }).(pulumi.IntOutput)
}

func (o LookupAgentPoolResultOutput) EnableAutoScaling() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupAgentPoolResult) *bool { return v.EnableAutoScaling }).(pulumi.BoolPtrOutput)
}

func (o LookupAgentPoolResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAgentPoolResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupAgentPoolResultOutput) MaxCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupAgentPoolResult) *int { return v.MaxCount }).(pulumi.IntPtrOutput)
}

func (o LookupAgentPoolResultOutput) MaxPods() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupAgentPoolResult) *int { return v.MaxPods }).(pulumi.IntPtrOutput)
}

func (o LookupAgentPoolResultOutput) MinCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupAgentPoolResult) *int { return v.MinCount }).(pulumi.IntPtrOutput)
}

func (o LookupAgentPoolResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAgentPoolResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupAgentPoolResultOutput) OrchestratorVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAgentPoolResult) *string { return v.OrchestratorVersion }).(pulumi.StringPtrOutput)
}

func (o LookupAgentPoolResultOutput) OsDiskSizeGB() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupAgentPoolResult) *int { return v.OsDiskSizeGB }).(pulumi.IntPtrOutput)
}

func (o LookupAgentPoolResultOutput) OsType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAgentPoolResult) *string { return v.OsType }).(pulumi.StringPtrOutput)
}

func (o LookupAgentPoolResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAgentPoolResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupAgentPoolResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAgentPoolResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupAgentPoolResultOutput) VmSize() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAgentPoolResult) string { return v.VmSize }).(pulumi.StringOutput)
}

func (o LookupAgentPoolResultOutput) VnetSubnetID() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAgentPoolResult) *string { return v.VnetSubnetID }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupAgentPoolResultOutput{})
}
