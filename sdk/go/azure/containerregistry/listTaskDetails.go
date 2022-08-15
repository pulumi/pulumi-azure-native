


package containerregistry

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListTaskDetails(ctx *pulumi.Context, args *ListTaskDetailsArgs, opts ...pulumi.InvokeOption) (*ListTaskDetailsResult, error) {
	var rv ListTaskDetailsResult
	err := ctx.Invoke("azure-native:containerregistry:listTaskDetails", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type ListTaskDetailsArgs struct {
	RegistryName      string `pulumi:"registryName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	TaskName          string `pulumi:"taskName"`
}



type ListTaskDetailsResult struct {
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


func (val *ListTaskDetailsResult) Defaults() *ListTaskDetailsResult {
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

func ListTaskDetailsOutput(ctx *pulumi.Context, args ListTaskDetailsOutputArgs, opts ...pulumi.InvokeOption) ListTaskDetailsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListTaskDetailsResult, error) {
			args := v.(ListTaskDetailsArgs)
			r, err := ListTaskDetails(ctx, &args, opts...)
			var s ListTaskDetailsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListTaskDetailsResultOutput)
}

type ListTaskDetailsOutputArgs struct {
	RegistryName      pulumi.StringInput `pulumi:"registryName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	TaskName          pulumi.StringInput `pulumi:"taskName"`
}

func (ListTaskDetailsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListTaskDetailsArgs)(nil)).Elem()
}



type ListTaskDetailsResultOutput struct{ *pulumi.OutputState }

func (ListTaskDetailsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListTaskDetailsResult)(nil)).Elem()
}

func (o ListTaskDetailsResultOutput) ToListTaskDetailsResultOutput() ListTaskDetailsResultOutput {
	return o
}

func (o ListTaskDetailsResultOutput) ToListTaskDetailsResultOutputWithContext(ctx context.Context) ListTaskDetailsResultOutput {
	return o
}

func (o ListTaskDetailsResultOutput) AgentConfiguration() AgentPropertiesResponsePtrOutput {
	return o.ApplyT(func(v ListTaskDetailsResult) *AgentPropertiesResponse { return v.AgentConfiguration }).(AgentPropertiesResponsePtrOutput)
}

func (o ListTaskDetailsResultOutput) AgentPoolName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListTaskDetailsResult) *string { return v.AgentPoolName }).(pulumi.StringPtrOutput)
}

func (o ListTaskDetailsResultOutput) CreationDate() pulumi.StringOutput {
	return o.ApplyT(func(v ListTaskDetailsResult) string { return v.CreationDate }).(pulumi.StringOutput)
}

func (o ListTaskDetailsResultOutput) Credentials() CredentialsResponsePtrOutput {
	return o.ApplyT(func(v ListTaskDetailsResult) *CredentialsResponse { return v.Credentials }).(CredentialsResponsePtrOutput)
}

func (o ListTaskDetailsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v ListTaskDetailsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o ListTaskDetailsResultOutput) Identity() IdentityPropertiesResponsePtrOutput {
	return o.ApplyT(func(v ListTaskDetailsResult) *IdentityPropertiesResponse { return v.Identity }).(IdentityPropertiesResponsePtrOutput)
}

func (o ListTaskDetailsResultOutput) IsSystemTask() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ListTaskDetailsResult) *bool { return v.IsSystemTask }).(pulumi.BoolPtrOutput)
}

func (o ListTaskDetailsResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v ListTaskDetailsResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o ListTaskDetailsResultOutput) LogTemplate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListTaskDetailsResult) *string { return v.LogTemplate }).(pulumi.StringPtrOutput)
}

func (o ListTaskDetailsResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ListTaskDetailsResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o ListTaskDetailsResultOutput) Platform() PlatformPropertiesResponsePtrOutput {
	return o.ApplyT(func(v ListTaskDetailsResult) *PlatformPropertiesResponse { return v.Platform }).(PlatformPropertiesResponsePtrOutput)
}

func (o ListTaskDetailsResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v ListTaskDetailsResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o ListTaskDetailsResultOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListTaskDetailsResult) *string { return v.Status }).(pulumi.StringPtrOutput)
}

func (o ListTaskDetailsResultOutput) Step() pulumi.AnyOutput {
	return o.ApplyT(func(v ListTaskDetailsResult) interface{} { return v.Step }).(pulumi.AnyOutput)
}

func (o ListTaskDetailsResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v ListTaskDetailsResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o ListTaskDetailsResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v ListTaskDetailsResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o ListTaskDetailsResultOutput) Timeout() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ListTaskDetailsResult) *int { return v.Timeout }).(pulumi.IntPtrOutput)
}

func (o ListTaskDetailsResultOutput) Trigger() TriggerPropertiesResponsePtrOutput {
	return o.ApplyT(func(v ListTaskDetailsResult) *TriggerPropertiesResponse { return v.Trigger }).(TriggerPropertiesResponsePtrOutput)
}

func (o ListTaskDetailsResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ListTaskDetailsResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ListTaskDetailsResultOutput{})
}
