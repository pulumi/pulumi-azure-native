


package v20150521preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


func LookupPolicyResource(ctx *pulumi.Context, args *LookupPolicyResourceArgs, opts ...pulumi.InvokeOption) (*LookupPolicyResourceResult, error) {
	var rv LookupPolicyResourceResult
	err := ctx.Invoke("azure-native:devtestlab/v20150521preview:getPolicyResource", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupPolicyResourceArgs struct {
	LabName           string `pulumi:"labName"`
	Name              string `pulumi:"name"`
	PolicySetName     string `pulumi:"policySetName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupPolicyResourceResult struct {
	Description       *string           `pulumi:"description"`
	EvaluatorType     *string           `pulumi:"evaluatorType"`
	FactData          *string           `pulumi:"factData"`
	FactName          *string           `pulumi:"factName"`
	Id                *string           `pulumi:"id"`
	Location          *string           `pulumi:"location"`
	Name              *string           `pulumi:"name"`
	ProvisioningState *string           `pulumi:"provisioningState"`
	Status            *string           `pulumi:"status"`
	Tags              map[string]string `pulumi:"tags"`
	Threshold         *string           `pulumi:"threshold"`
	Type              *string           `pulumi:"type"`
}

func LookupPolicyResourceOutput(ctx *pulumi.Context, args LookupPolicyResourceOutputArgs, opts ...pulumi.InvokeOption) LookupPolicyResourceResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupPolicyResourceResult, error) {
			args := v.(LookupPolicyResourceArgs)
			r, err := LookupPolicyResource(ctx, &args, opts...)
			var s LookupPolicyResourceResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupPolicyResourceResultOutput)
}

type LookupPolicyResourceOutputArgs struct {
	LabName           pulumi.StringInput `pulumi:"labName"`
	Name              pulumi.StringInput `pulumi:"name"`
	PolicySetName     pulumi.StringInput `pulumi:"policySetName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupPolicyResourceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPolicyResourceArgs)(nil)).Elem()
}


type LookupPolicyResourceResultOutput struct{ *pulumi.OutputState }

func (LookupPolicyResourceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPolicyResourceResult)(nil)).Elem()
}

func (o LookupPolicyResourceResultOutput) ToLookupPolicyResourceResultOutput() LookupPolicyResourceResultOutput {
	return o
}

func (o LookupPolicyResourceResultOutput) ToLookupPolicyResourceResultOutputWithContext(ctx context.Context) LookupPolicyResourceResultOutput {
	return o
}

func (o LookupPolicyResourceResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPolicyResourceResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupPolicyResourceResultOutput) EvaluatorType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPolicyResourceResult) *string { return v.EvaluatorType }).(pulumi.StringPtrOutput)
}

func (o LookupPolicyResourceResultOutput) FactData() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPolicyResourceResult) *string { return v.FactData }).(pulumi.StringPtrOutput)
}

func (o LookupPolicyResourceResultOutput) FactName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPolicyResourceResult) *string { return v.FactName }).(pulumi.StringPtrOutput)
}

func (o LookupPolicyResourceResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPolicyResourceResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupPolicyResourceResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPolicyResourceResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupPolicyResourceResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPolicyResourceResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o LookupPolicyResourceResultOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPolicyResourceResult) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

func (o LookupPolicyResourceResultOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPolicyResourceResult) *string { return v.Status }).(pulumi.StringPtrOutput)
}

func (o LookupPolicyResourceResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupPolicyResourceResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupPolicyResourceResultOutput) Threshold() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPolicyResourceResult) *string { return v.Threshold }).(pulumi.StringPtrOutput)
}

func (o LookupPolicyResourceResultOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPolicyResourceResult) *string { return v.Type }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupPolicyResourceResultOutput{})
}
