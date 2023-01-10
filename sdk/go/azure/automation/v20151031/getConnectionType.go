


package v20151031

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupConnectionType(ctx *pulumi.Context, args *LookupConnectionTypeArgs, opts ...pulumi.InvokeOption) (*LookupConnectionTypeResult, error) {
	var rv LookupConnectionTypeResult
	err := ctx.Invoke("azure-native:automation/v20151031:getConnectionType", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupConnectionTypeArgs struct {
	AutomationAccountName string `pulumi:"automationAccountName"`
	ConnectionTypeName    string `pulumi:"connectionTypeName"`
	ResourceGroupName     string `pulumi:"resourceGroupName"`
}


type LookupConnectionTypeResult struct {
	CreationTime     string                             `pulumi:"creationTime"`
	Description      *string                            `pulumi:"description"`
	FieldDefinitions map[string]FieldDefinitionResponse `pulumi:"fieldDefinitions"`
	Id               string                             `pulumi:"id"`
	IsGlobal         *bool                              `pulumi:"isGlobal"`
	LastModifiedTime *string                            `pulumi:"lastModifiedTime"`
	Name             string                             `pulumi:"name"`
	Type             string                             `pulumi:"type"`
}

func LookupConnectionTypeOutput(ctx *pulumi.Context, args LookupConnectionTypeOutputArgs, opts ...pulumi.InvokeOption) LookupConnectionTypeResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupConnectionTypeResult, error) {
			args := v.(LookupConnectionTypeArgs)
			r, err := LookupConnectionType(ctx, &args, opts...)
			var s LookupConnectionTypeResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupConnectionTypeResultOutput)
}

type LookupConnectionTypeOutputArgs struct {
	AutomationAccountName pulumi.StringInput `pulumi:"automationAccountName"`
	ConnectionTypeName    pulumi.StringInput `pulumi:"connectionTypeName"`
	ResourceGroupName     pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupConnectionTypeOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupConnectionTypeArgs)(nil)).Elem()
}


type LookupConnectionTypeResultOutput struct{ *pulumi.OutputState }

func (LookupConnectionTypeResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupConnectionTypeResult)(nil)).Elem()
}

func (o LookupConnectionTypeResultOutput) ToLookupConnectionTypeResultOutput() LookupConnectionTypeResultOutput {
	return o
}

func (o LookupConnectionTypeResultOutput) ToLookupConnectionTypeResultOutputWithContext(ctx context.Context) LookupConnectionTypeResultOutput {
	return o
}

func (o LookupConnectionTypeResultOutput) CreationTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConnectionTypeResult) string { return v.CreationTime }).(pulumi.StringOutput)
}

func (o LookupConnectionTypeResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupConnectionTypeResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupConnectionTypeResultOutput) FieldDefinitions() FieldDefinitionResponseMapOutput {
	return o.ApplyT(func(v LookupConnectionTypeResult) map[string]FieldDefinitionResponse { return v.FieldDefinitions }).(FieldDefinitionResponseMapOutput)
}

func (o LookupConnectionTypeResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConnectionTypeResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupConnectionTypeResultOutput) IsGlobal() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupConnectionTypeResult) *bool { return v.IsGlobal }).(pulumi.BoolPtrOutput)
}

func (o LookupConnectionTypeResultOutput) LastModifiedTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupConnectionTypeResult) *string { return v.LastModifiedTime }).(pulumi.StringPtrOutput)
}

func (o LookupConnectionTypeResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConnectionTypeResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupConnectionTypeResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConnectionTypeResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupConnectionTypeResultOutput{})
}
