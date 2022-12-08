


package v20220701preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupJobDefinition(ctx *pulumi.Context, args *LookupJobDefinitionArgs, opts ...pulumi.InvokeOption) (*LookupJobDefinitionResult, error) {
	var rv LookupJobDefinitionResult
	err := ctx.Invoke("azure-native:storagemover/v20220701preview:getJobDefinition", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupJobDefinitionArgs struct {
	JobDefinitionName string `pulumi:"jobDefinitionName"`
	ProjectName       string `pulumi:"projectName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	StorageMoverName  string `pulumi:"storageMoverName"`
}


type LookupJobDefinitionResult struct {
	AgentName              *string            `pulumi:"agentName"`
	AgentResourceId        string             `pulumi:"agentResourceId"`
	CopyMode               string             `pulumi:"copyMode"`
	Description            *string            `pulumi:"description"`
	Id                     string             `pulumi:"id"`
	LatestJobRunName       string             `pulumi:"latestJobRunName"`
	LatestJobRunResourceId string             `pulumi:"latestJobRunResourceId"`
	LatestJobRunStatus     string             `pulumi:"latestJobRunStatus"`
	Name                   string             `pulumi:"name"`
	ProvisioningState      string             `pulumi:"provisioningState"`
	SourceName             string             `pulumi:"sourceName"`
	SourceResourceId       string             `pulumi:"sourceResourceId"`
	SourceSubpath          *string            `pulumi:"sourceSubpath"`
	SystemData             SystemDataResponse `pulumi:"systemData"`
	TargetName             string             `pulumi:"targetName"`
	TargetResourceId       string             `pulumi:"targetResourceId"`
	TargetSubpath          *string            `pulumi:"targetSubpath"`
	Type                   string             `pulumi:"type"`
}

func LookupJobDefinitionOutput(ctx *pulumi.Context, args LookupJobDefinitionOutputArgs, opts ...pulumi.InvokeOption) LookupJobDefinitionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupJobDefinitionResult, error) {
			args := v.(LookupJobDefinitionArgs)
			r, err := LookupJobDefinition(ctx, &args, opts...)
			var s LookupJobDefinitionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupJobDefinitionResultOutput)
}

type LookupJobDefinitionOutputArgs struct {
	JobDefinitionName pulumi.StringInput `pulumi:"jobDefinitionName"`
	ProjectName       pulumi.StringInput `pulumi:"projectName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	StorageMoverName  pulumi.StringInput `pulumi:"storageMoverName"`
}

func (LookupJobDefinitionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupJobDefinitionArgs)(nil)).Elem()
}


type LookupJobDefinitionResultOutput struct{ *pulumi.OutputState }

func (LookupJobDefinitionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupJobDefinitionResult)(nil)).Elem()
}

func (o LookupJobDefinitionResultOutput) ToLookupJobDefinitionResultOutput() LookupJobDefinitionResultOutput {
	return o
}

func (o LookupJobDefinitionResultOutput) ToLookupJobDefinitionResultOutputWithContext(ctx context.Context) LookupJobDefinitionResultOutput {
	return o
}

func (o LookupJobDefinitionResultOutput) AgentName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupJobDefinitionResult) *string { return v.AgentName }).(pulumi.StringPtrOutput)
}

func (o LookupJobDefinitionResultOutput) AgentResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupJobDefinitionResult) string { return v.AgentResourceId }).(pulumi.StringOutput)
}

func (o LookupJobDefinitionResultOutput) CopyMode() pulumi.StringOutput {
	return o.ApplyT(func(v LookupJobDefinitionResult) string { return v.CopyMode }).(pulumi.StringOutput)
}

func (o LookupJobDefinitionResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupJobDefinitionResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupJobDefinitionResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupJobDefinitionResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupJobDefinitionResultOutput) LatestJobRunName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupJobDefinitionResult) string { return v.LatestJobRunName }).(pulumi.StringOutput)
}

func (o LookupJobDefinitionResultOutput) LatestJobRunResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupJobDefinitionResult) string { return v.LatestJobRunResourceId }).(pulumi.StringOutput)
}

func (o LookupJobDefinitionResultOutput) LatestJobRunStatus() pulumi.StringOutput {
	return o.ApplyT(func(v LookupJobDefinitionResult) string { return v.LatestJobRunStatus }).(pulumi.StringOutput)
}

func (o LookupJobDefinitionResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupJobDefinitionResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupJobDefinitionResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupJobDefinitionResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupJobDefinitionResultOutput) SourceName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupJobDefinitionResult) string { return v.SourceName }).(pulumi.StringOutput)
}

func (o LookupJobDefinitionResultOutput) SourceResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupJobDefinitionResult) string { return v.SourceResourceId }).(pulumi.StringOutput)
}

func (o LookupJobDefinitionResultOutput) SourceSubpath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupJobDefinitionResult) *string { return v.SourceSubpath }).(pulumi.StringPtrOutput)
}

func (o LookupJobDefinitionResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupJobDefinitionResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupJobDefinitionResultOutput) TargetName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupJobDefinitionResult) string { return v.TargetName }).(pulumi.StringOutput)
}

func (o LookupJobDefinitionResultOutput) TargetResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupJobDefinitionResult) string { return v.TargetResourceId }).(pulumi.StringOutput)
}

func (o LookupJobDefinitionResultOutput) TargetSubpath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupJobDefinitionResult) *string { return v.TargetSubpath }).(pulumi.StringPtrOutput)
}

func (o LookupJobDefinitionResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupJobDefinitionResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupJobDefinitionResultOutput{})
}
