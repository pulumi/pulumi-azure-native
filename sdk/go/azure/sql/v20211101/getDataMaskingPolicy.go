


package v20211101

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupDataMaskingPolicy(ctx *pulumi.Context, args *LookupDataMaskingPolicyArgs, opts ...pulumi.InvokeOption) (*LookupDataMaskingPolicyResult, error) {
	var rv LookupDataMaskingPolicyResult
	err := ctx.Invoke("azure-native:sql/v20211101:getDataMaskingPolicy", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDataMaskingPolicyArgs struct {
	DataMaskingPolicyName string `pulumi:"dataMaskingPolicyName"`
	DatabaseName          string `pulumi:"databaseName"`
	ResourceGroupName     string `pulumi:"resourceGroupName"`
	ServerName            string `pulumi:"serverName"`
}


type LookupDataMaskingPolicyResult struct {
	ApplicationPrincipals string  `pulumi:"applicationPrincipals"`
	DataMaskingState      string  `pulumi:"dataMaskingState"`
	ExemptPrincipals      *string `pulumi:"exemptPrincipals"`
	Id                    string  `pulumi:"id"`
	Kind                  string  `pulumi:"kind"`
	Location              string  `pulumi:"location"`
	MaskingLevel          string  `pulumi:"maskingLevel"`
	Name                  string  `pulumi:"name"`
	Type                  string  `pulumi:"type"`
}

func LookupDataMaskingPolicyOutput(ctx *pulumi.Context, args LookupDataMaskingPolicyOutputArgs, opts ...pulumi.InvokeOption) LookupDataMaskingPolicyResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDataMaskingPolicyResult, error) {
			args := v.(LookupDataMaskingPolicyArgs)
			r, err := LookupDataMaskingPolicy(ctx, &args, opts...)
			var s LookupDataMaskingPolicyResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupDataMaskingPolicyResultOutput)
}

type LookupDataMaskingPolicyOutputArgs struct {
	DataMaskingPolicyName pulumi.StringInput `pulumi:"dataMaskingPolicyName"`
	DatabaseName          pulumi.StringInput `pulumi:"databaseName"`
	ResourceGroupName     pulumi.StringInput `pulumi:"resourceGroupName"`
	ServerName            pulumi.StringInput `pulumi:"serverName"`
}

func (LookupDataMaskingPolicyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDataMaskingPolicyArgs)(nil)).Elem()
}


type LookupDataMaskingPolicyResultOutput struct{ *pulumi.OutputState }

func (LookupDataMaskingPolicyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDataMaskingPolicyResult)(nil)).Elem()
}

func (o LookupDataMaskingPolicyResultOutput) ToLookupDataMaskingPolicyResultOutput() LookupDataMaskingPolicyResultOutput {
	return o
}

func (o LookupDataMaskingPolicyResultOutput) ToLookupDataMaskingPolicyResultOutputWithContext(ctx context.Context) LookupDataMaskingPolicyResultOutput {
	return o
}

func (o LookupDataMaskingPolicyResultOutput) ApplicationPrincipals() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDataMaskingPolicyResult) string { return v.ApplicationPrincipals }).(pulumi.StringOutput)
}

func (o LookupDataMaskingPolicyResultOutput) DataMaskingState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDataMaskingPolicyResult) string { return v.DataMaskingState }).(pulumi.StringOutput)
}

func (o LookupDataMaskingPolicyResultOutput) ExemptPrincipals() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDataMaskingPolicyResult) *string { return v.ExemptPrincipals }).(pulumi.StringPtrOutput)
}

func (o LookupDataMaskingPolicyResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDataMaskingPolicyResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupDataMaskingPolicyResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDataMaskingPolicyResult) string { return v.Kind }).(pulumi.StringOutput)
}

func (o LookupDataMaskingPolicyResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDataMaskingPolicyResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupDataMaskingPolicyResultOutput) MaskingLevel() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDataMaskingPolicyResult) string { return v.MaskingLevel }).(pulumi.StringOutput)
}

func (o LookupDataMaskingPolicyResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDataMaskingPolicyResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupDataMaskingPolicyResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDataMaskingPolicyResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDataMaskingPolicyResultOutput{})
}
