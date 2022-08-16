


package v20190601preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupTask(ctx *pulumi.Context, args *LookupTaskArgs, opts ...pulumi.InvokeOption) (*LookupTaskResult, error) {
	var rv LookupTaskResult
	err := ctx.Invoke("azure-native:containerregistry/v20190601preview:getTask", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupTaskArgs struct {
	RegistryName      string `pulumi:"registryName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	TaskName          string `pulumi:"taskName"`
}



type LookupTaskResult struct {
	AgentConfiguration *AgentPropertiesResponse    `pulumi:"agentConfiguration"`
	AgentPoolName      *string                     `pulumi:"agentPoolName"`
	CreationDate       string                      `pulumi:"creationDate"`
	Credentials        *CredentialsResponse        `pulumi:"credentials"`
	Id                 string                      `pulumi:"id"`
	Identity           *IdentityPropertiesResponse `pulumi:"identity"`
	IsSystemTask       *bool                       `pulumi:"isSystemTask"`
	Location           string                      `pulumi:"location"`
	LogTemplate        *string                     `pulumi:"logTemplate"`
	Name               string                      `pulumi:"name"`
	Platform           *PlatformPropertiesResponse `pulumi:"platform"`
	ProvisioningState  string                      `pulumi:"provisioningState"`
	Status             *string                     `pulumi:"status"`
	Step               interface{}                 `pulumi:"step"`
	SystemData         SystemDataResponse          `pulumi:"systemData"`
	Tags               map[string]string           `pulumi:"tags"`
	Timeout            *int                        `pulumi:"timeout"`
	Trigger            *TriggerPropertiesResponse  `pulumi:"trigger"`
	Type               string                      `pulumi:"type"`
}


func (val *LookupTaskResult) Defaults() *LookupTaskResult {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.IsSystemTask) {
		isSystemTask_ := false
		tmp.IsSystemTask = &isSystemTask_
	}
	if isZero(tmp.Timeout) {
		timeout_ := 3600
		tmp.Timeout = &timeout_
	}
	tmp.Trigger = tmp.Trigger.Defaults()

	return &tmp
}

func LookupTaskOutput(ctx *pulumi.Context, args LookupTaskOutputArgs, opts ...pulumi.InvokeOption) LookupTaskResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupTaskResult, error) {
			args := v.(LookupTaskArgs)
			r, err := LookupTask(ctx, &args, opts...)
			var s LookupTaskResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupTaskResultOutput)
}

type LookupTaskOutputArgs struct {
	RegistryName      pulumi.StringInput `pulumi:"registryName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	TaskName          pulumi.StringInput `pulumi:"taskName"`
}

func (LookupTaskOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTaskArgs)(nil)).Elem()
}



type LookupTaskResultOutput struct{ *pulumi.OutputState }

func (LookupTaskResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTaskResult)(nil)).Elem()
}

func (o LookupTaskResultOutput) ToLookupTaskResultOutput() LookupTaskResultOutput {
	return o
}

func (o LookupTaskResultOutput) ToLookupTaskResultOutputWithContext(ctx context.Context) LookupTaskResultOutput {
	return o
}

func (o LookupTaskResultOutput) AgentConfiguration() AgentPropertiesResponsePtrOutput {
	return o.ApplyT(func(v LookupTaskResult) *AgentPropertiesResponse { return v.AgentConfiguration }).(AgentPropertiesResponsePtrOutput)
}

func (o LookupTaskResultOutput) AgentPoolName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupTaskResult) *string { return v.AgentPoolName }).(pulumi.StringPtrOutput)
}

func (o LookupTaskResultOutput) CreationDate() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTaskResult) string { return v.CreationDate }).(pulumi.StringOutput)
}

func (o LookupTaskResultOutput) Credentials() CredentialsResponsePtrOutput {
	return o.ApplyT(func(v LookupTaskResult) *CredentialsResponse { return v.Credentials }).(CredentialsResponsePtrOutput)
}

func (o LookupTaskResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTaskResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupTaskResultOutput) Identity() IdentityPropertiesResponsePtrOutput {
	return o.ApplyT(func(v LookupTaskResult) *IdentityPropertiesResponse { return v.Identity }).(IdentityPropertiesResponsePtrOutput)
}

func (o LookupTaskResultOutput) IsSystemTask() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupTaskResult) *bool { return v.IsSystemTask }).(pulumi.BoolPtrOutput)
}

func (o LookupTaskResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTaskResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupTaskResultOutput) LogTemplate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupTaskResult) *string { return v.LogTemplate }).(pulumi.StringPtrOutput)
}

func (o LookupTaskResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTaskResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupTaskResultOutput) Platform() PlatformPropertiesResponsePtrOutput {
	return o.ApplyT(func(v LookupTaskResult) *PlatformPropertiesResponse { return v.Platform }).(PlatformPropertiesResponsePtrOutput)
}

func (o LookupTaskResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTaskResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupTaskResultOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupTaskResult) *string { return v.Status }).(pulumi.StringPtrOutput)
}

func (o LookupTaskResultOutput) Step() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupTaskResult) interface{} { return v.Step }).(pulumi.AnyOutput)
}

func (o LookupTaskResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupTaskResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupTaskResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupTaskResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupTaskResultOutput) Timeout() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupTaskResult) *int { return v.Timeout }).(pulumi.IntPtrOutput)
}

func (o LookupTaskResultOutput) Trigger() TriggerPropertiesResponsePtrOutput {
	return o.ApplyT(func(v LookupTaskResult) *TriggerPropertiesResponse { return v.Trigger }).(TriggerPropertiesResponsePtrOutput)
}

func (o LookupTaskResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTaskResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupTaskResultOutput{})
}
