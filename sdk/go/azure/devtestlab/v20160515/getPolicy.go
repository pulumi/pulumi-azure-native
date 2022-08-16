


package v20160515

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


func LookupPolicy(ctx *pulumi.Context, args *LookupPolicyArgs, opts ...pulumi.InvokeOption) (*LookupPolicyResult, error) {
	var rv LookupPolicyResult
	err := ctx.Invoke("azure-native:devtestlab/v20160515:getPolicy", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupPolicyArgs struct {
	Expand            *string `pulumi:"expand"`
	LabName           string  `pulumi:"labName"`
	Name              string  `pulumi:"name"`
	PolicySetName     string  `pulumi:"policySetName"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
}


type LookupPolicyResult struct {
	CreatedDate       string            `pulumi:"createdDate"`
	Description       *string           `pulumi:"description"`
	EvaluatorType     *string           `pulumi:"evaluatorType"`
	FactData          *string           `pulumi:"factData"`
	FactName          *string           `pulumi:"factName"`
	Id                string            `pulumi:"id"`
	Location          *string           `pulumi:"location"`
	Name              string            `pulumi:"name"`
	ProvisioningState *string           `pulumi:"provisioningState"`
	Status            *string           `pulumi:"status"`
	Tags              map[string]string `pulumi:"tags"`
	Threshold         *string           `pulumi:"threshold"`
	Type              string            `pulumi:"type"`
	UniqueIdentifier  *string           `pulumi:"uniqueIdentifier"`
}

func LookupPolicyOutput(ctx *pulumi.Context, args LookupPolicyOutputArgs, opts ...pulumi.InvokeOption) LookupPolicyResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupPolicyResult, error) {
			args := v.(LookupPolicyArgs)
			r, err := LookupPolicy(ctx, &args, opts...)
			var s LookupPolicyResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupPolicyResultOutput)
}

type LookupPolicyOutputArgs struct {
	Expand            pulumi.StringPtrInput `pulumi:"expand"`
	LabName           pulumi.StringInput    `pulumi:"labName"`
	Name              pulumi.StringInput    `pulumi:"name"`
	PolicySetName     pulumi.StringInput    `pulumi:"policySetName"`
	ResourceGroupName pulumi.StringInput    `pulumi:"resourceGroupName"`
}

func (LookupPolicyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPolicyArgs)(nil)).Elem()
}


type LookupPolicyResultOutput struct{ *pulumi.OutputState }

func (LookupPolicyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPolicyResult)(nil)).Elem()
}

func (o LookupPolicyResultOutput) ToLookupPolicyResultOutput() LookupPolicyResultOutput {
	return o
}

func (o LookupPolicyResultOutput) ToLookupPolicyResultOutputWithContext(ctx context.Context) LookupPolicyResultOutput {
	return o
}

func (o LookupPolicyResultOutput) CreatedDate() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPolicyResult) string { return v.CreatedDate }).(pulumi.StringOutput)
}

func (o LookupPolicyResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPolicyResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupPolicyResultOutput) EvaluatorType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPolicyResult) *string { return v.EvaluatorType }).(pulumi.StringPtrOutput)
}

func (o LookupPolicyResultOutput) FactData() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPolicyResult) *string { return v.FactData }).(pulumi.StringPtrOutput)
}

func (o LookupPolicyResultOutput) FactName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPolicyResult) *string { return v.FactName }).(pulumi.StringPtrOutput)
}

func (o LookupPolicyResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPolicyResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupPolicyResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPolicyResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupPolicyResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPolicyResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupPolicyResultOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPolicyResult) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

func (o LookupPolicyResultOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPolicyResult) *string { return v.Status }).(pulumi.StringPtrOutput)
}

func (o LookupPolicyResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupPolicyResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupPolicyResultOutput) Threshold() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPolicyResult) *string { return v.Threshold }).(pulumi.StringPtrOutput)
}

func (o LookupPolicyResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPolicyResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupPolicyResultOutput) UniqueIdentifier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPolicyResult) *string { return v.UniqueIdentifier }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupPolicyResultOutput{})
}
