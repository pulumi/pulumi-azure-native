


package storagemover

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupAgent(ctx *pulumi.Context, args *LookupAgentArgs, opts ...pulumi.InvokeOption) (*LookupAgentResult, error) {
	var rv LookupAgentResult
	err := ctx.Invoke("azure-native:storagemover:getAgent", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupAgentArgs struct {
	AgentName         string `pulumi:"agentName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	StorageMoverName  string `pulumi:"storageMoverName"`
}


type LookupAgentResult struct {
	AgentStatus       string                              `pulumi:"agentStatus"`
	AgentVersion      string                              `pulumi:"agentVersion"`
	ArcResourceId     string                              `pulumi:"arcResourceId"`
	ArcVmUuid         string                              `pulumi:"arcVmUuid"`
	Description       *string                             `pulumi:"description"`
	ErrorDetails      AgentPropertiesResponseErrorDetails `pulumi:"errorDetails"`
	Id                string                              `pulumi:"id"`
	LastStatusUpdate  string                              `pulumi:"lastStatusUpdate"`
	LocalIPAddress    string                              `pulumi:"localIPAddress"`
	MemoryInMB        float64                             `pulumi:"memoryInMB"`
	Name              string                              `pulumi:"name"`
	NumberOfCores     float64                             `pulumi:"numberOfCores"`
	ProvisioningState string                              `pulumi:"provisioningState"`
	SystemData        SystemDataResponse                  `pulumi:"systemData"`
	Type              string                              `pulumi:"type"`
	UptimeInSeconds   float64                             `pulumi:"uptimeInSeconds"`
}

func LookupAgentOutput(ctx *pulumi.Context, args LookupAgentOutputArgs, opts ...pulumi.InvokeOption) LookupAgentResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupAgentResult, error) {
			args := v.(LookupAgentArgs)
			r, err := LookupAgent(ctx, &args, opts...)
			var s LookupAgentResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupAgentResultOutput)
}

type LookupAgentOutputArgs struct {
	AgentName         pulumi.StringInput `pulumi:"agentName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	StorageMoverName  pulumi.StringInput `pulumi:"storageMoverName"`
}

func (LookupAgentOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAgentArgs)(nil)).Elem()
}


type LookupAgentResultOutput struct{ *pulumi.OutputState }

func (LookupAgentResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAgentResult)(nil)).Elem()
}

func (o LookupAgentResultOutput) ToLookupAgentResultOutput() LookupAgentResultOutput {
	return o
}

func (o LookupAgentResultOutput) ToLookupAgentResultOutputWithContext(ctx context.Context) LookupAgentResultOutput {
	return o
}

func (o LookupAgentResultOutput) AgentStatus() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAgentResult) string { return v.AgentStatus }).(pulumi.StringOutput)
}

func (o LookupAgentResultOutput) AgentVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAgentResult) string { return v.AgentVersion }).(pulumi.StringOutput)
}

func (o LookupAgentResultOutput) ArcResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAgentResult) string { return v.ArcResourceId }).(pulumi.StringOutput)
}

func (o LookupAgentResultOutput) ArcVmUuid() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAgentResult) string { return v.ArcVmUuid }).(pulumi.StringOutput)
}

func (o LookupAgentResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAgentResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupAgentResultOutput) ErrorDetails() AgentPropertiesResponseErrorDetailsOutput {
	return o.ApplyT(func(v LookupAgentResult) AgentPropertiesResponseErrorDetails { return v.ErrorDetails }).(AgentPropertiesResponseErrorDetailsOutput)
}

func (o LookupAgentResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAgentResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupAgentResultOutput) LastStatusUpdate() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAgentResult) string { return v.LastStatusUpdate }).(pulumi.StringOutput)
}

func (o LookupAgentResultOutput) LocalIPAddress() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAgentResult) string { return v.LocalIPAddress }).(pulumi.StringOutput)
}

func (o LookupAgentResultOutput) MemoryInMB() pulumi.Float64Output {
	return o.ApplyT(func(v LookupAgentResult) float64 { return v.MemoryInMB }).(pulumi.Float64Output)
}

func (o LookupAgentResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAgentResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupAgentResultOutput) NumberOfCores() pulumi.Float64Output {
	return o.ApplyT(func(v LookupAgentResult) float64 { return v.NumberOfCores }).(pulumi.Float64Output)
}

func (o LookupAgentResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAgentResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupAgentResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupAgentResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupAgentResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAgentResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupAgentResultOutput) UptimeInSeconds() pulumi.Float64Output {
	return o.ApplyT(func(v LookupAgentResult) float64 { return v.UptimeInSeconds }).(pulumi.Float64Output)
}

func init() {
	pulumi.RegisterOutputType(LookupAgentResultOutput{})
}
