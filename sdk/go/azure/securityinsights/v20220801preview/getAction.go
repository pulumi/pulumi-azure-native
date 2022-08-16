


package v20220801preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupAction(ctx *pulumi.Context, args *LookupActionArgs, opts ...pulumi.InvokeOption) (*LookupActionResult, error) {
	var rv LookupActionResult
	err := ctx.Invoke("azure-native:securityinsights/v20220801preview:getAction", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupActionArgs struct {
	ActionId          string `pulumi:"actionId"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	RuleId            string `pulumi:"ruleId"`
	WorkspaceName     string `pulumi:"workspaceName"`
}


type LookupActionResult struct {
	Etag               *string            `pulumi:"etag"`
	Id                 string             `pulumi:"id"`
	LogicAppResourceId string             `pulumi:"logicAppResourceId"`
	Name               string             `pulumi:"name"`
	SystemData         SystemDataResponse `pulumi:"systemData"`
	Type               string             `pulumi:"type"`
	WorkflowId         *string            `pulumi:"workflowId"`
}

func LookupActionOutput(ctx *pulumi.Context, args LookupActionOutputArgs, opts ...pulumi.InvokeOption) LookupActionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupActionResult, error) {
			args := v.(LookupActionArgs)
			r, err := LookupAction(ctx, &args, opts...)
			var s LookupActionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupActionResultOutput)
}

type LookupActionOutputArgs struct {
	ActionId          pulumi.StringInput `pulumi:"actionId"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	RuleId            pulumi.StringInput `pulumi:"ruleId"`
	WorkspaceName     pulumi.StringInput `pulumi:"workspaceName"`
}

func (LookupActionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupActionArgs)(nil)).Elem()
}


type LookupActionResultOutput struct{ *pulumi.OutputState }

func (LookupActionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupActionResult)(nil)).Elem()
}

func (o LookupActionResultOutput) ToLookupActionResultOutput() LookupActionResultOutput {
	return o
}

func (o LookupActionResultOutput) ToLookupActionResultOutputWithContext(ctx context.Context) LookupActionResultOutput {
	return o
}

func (o LookupActionResultOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupActionResult) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o LookupActionResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupActionResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupActionResultOutput) LogicAppResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupActionResult) string { return v.LogicAppResourceId }).(pulumi.StringOutput)
}

func (o LookupActionResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupActionResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupActionResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupActionResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupActionResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupActionResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupActionResultOutput) WorkflowId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupActionResult) *string { return v.WorkflowId }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupActionResultOutput{})
}
