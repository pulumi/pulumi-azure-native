


package v20151031

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupVariable(ctx *pulumi.Context, args *LookupVariableArgs, opts ...pulumi.InvokeOption) (*LookupVariableResult, error) {
	var rv LookupVariableResult
	err := ctx.Invoke("azure-native:automation/v20151031:getVariable", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupVariableArgs struct {
	AutomationAccountName string `pulumi:"automationAccountName"`
	ResourceGroupName     string `pulumi:"resourceGroupName"`
	VariableName          string `pulumi:"variableName"`
}


type LookupVariableResult struct {
	CreationTime     *string `pulumi:"creationTime"`
	Description      *string `pulumi:"description"`
	Id               string  `pulumi:"id"`
	IsEncrypted      *bool   `pulumi:"isEncrypted"`
	LastModifiedTime *string `pulumi:"lastModifiedTime"`
	Name             string  `pulumi:"name"`
	Type             string  `pulumi:"type"`
	Value            *string `pulumi:"value"`
}

func LookupVariableOutput(ctx *pulumi.Context, args LookupVariableOutputArgs, opts ...pulumi.InvokeOption) LookupVariableResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupVariableResult, error) {
			args := v.(LookupVariableArgs)
			r, err := LookupVariable(ctx, &args, opts...)
			var s LookupVariableResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupVariableResultOutput)
}

type LookupVariableOutputArgs struct {
	AutomationAccountName pulumi.StringInput `pulumi:"automationAccountName"`
	ResourceGroupName     pulumi.StringInput `pulumi:"resourceGroupName"`
	VariableName          pulumi.StringInput `pulumi:"variableName"`
}

func (LookupVariableOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVariableArgs)(nil)).Elem()
}


type LookupVariableResultOutput struct{ *pulumi.OutputState }

func (LookupVariableResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVariableResult)(nil)).Elem()
}

func (o LookupVariableResultOutput) ToLookupVariableResultOutput() LookupVariableResultOutput {
	return o
}

func (o LookupVariableResultOutput) ToLookupVariableResultOutputWithContext(ctx context.Context) LookupVariableResultOutput {
	return o
}

func (o LookupVariableResultOutput) CreationTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVariableResult) *string { return v.CreationTime }).(pulumi.StringPtrOutput)
}

func (o LookupVariableResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVariableResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupVariableResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVariableResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupVariableResultOutput) IsEncrypted() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupVariableResult) *bool { return v.IsEncrypted }).(pulumi.BoolPtrOutput)
}

func (o LookupVariableResultOutput) LastModifiedTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVariableResult) *string { return v.LastModifiedTime }).(pulumi.StringPtrOutput)
}

func (o LookupVariableResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVariableResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupVariableResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVariableResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupVariableResultOutput) Value() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVariableResult) *string { return v.Value }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupVariableResultOutput{})
}
