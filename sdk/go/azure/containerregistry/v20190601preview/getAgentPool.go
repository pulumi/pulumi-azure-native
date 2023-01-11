


package v20190601preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupAgentPool(ctx *pulumi.Context, args *LookupAgentPoolArgs, opts ...pulumi.InvokeOption) (*LookupAgentPoolResult, error) {
	var rv LookupAgentPoolResult
	err := ctx.Invoke("azure-native:containerregistry/v20190601preview:getAgentPool", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupAgentPoolArgs struct {
	AgentPoolName     string `pulumi:"agentPoolName"`
	RegistryName      string `pulumi:"registryName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}



type LookupAgentPoolResult struct {
	Count                          *int               `pulumi:"count"`
	Id                             string             `pulumi:"id"`
	Location                       string             `pulumi:"location"`
	Name                           string             `pulumi:"name"`
	Os                             *string            `pulumi:"os"`
	ProvisioningState              string             `pulumi:"provisioningState"`
	SystemData                     SystemDataResponse `pulumi:"systemData"`
	Tags                           map[string]string  `pulumi:"tags"`
	Tier                           *string            `pulumi:"tier"`
	Type                           string             `pulumi:"type"`
	VirtualNetworkSubnetResourceId *string            `pulumi:"virtualNetworkSubnetResourceId"`
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
	AgentPoolName     pulumi.StringInput `pulumi:"agentPoolName"`
	RegistryName      pulumi.StringInput `pulumi:"registryName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
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

func (o LookupAgentPoolResultOutput) Count() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupAgentPoolResult) *int { return v.Count }).(pulumi.IntPtrOutput)
}

func (o LookupAgentPoolResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAgentPoolResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupAgentPoolResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAgentPoolResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupAgentPoolResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAgentPoolResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupAgentPoolResultOutput) Os() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAgentPoolResult) *string { return v.Os }).(pulumi.StringPtrOutput)
}

func (o LookupAgentPoolResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAgentPoolResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupAgentPoolResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupAgentPoolResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupAgentPoolResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupAgentPoolResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupAgentPoolResultOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAgentPoolResult) *string { return v.Tier }).(pulumi.StringPtrOutput)
}

func (o LookupAgentPoolResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAgentPoolResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupAgentPoolResultOutput) VirtualNetworkSubnetResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAgentPoolResult) *string { return v.VirtualNetworkSubnetResourceId }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupAgentPoolResultOutput{})
}
