


package v20190801preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


func LookupPeeringServicePrefix(ctx *pulumi.Context, args *LookupPeeringServicePrefixArgs, opts ...pulumi.InvokeOption) (*LookupPeeringServicePrefixResult, error) {
	var rv LookupPeeringServicePrefixResult
	err := ctx.Invoke("azure-native:peering/v20190801preview:getPeeringServicePrefix", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupPeeringServicePrefixArgs struct {
	PeeringServiceName string `pulumi:"peeringServiceName"`
	PrefixName         string `pulumi:"prefixName"`
	ResourceGroupName  string `pulumi:"resourceGroupName"`
}


type LookupPeeringServicePrefixResult struct {
	Id                    string  `pulumi:"id"`
	LearnedType           *string `pulumi:"learnedType"`
	Name                  string  `pulumi:"name"`
	Prefix                *string `pulumi:"prefix"`
	PrefixValidationState *string `pulumi:"prefixValidationState"`
	ProvisioningState     string  `pulumi:"provisioningState"`
	Type                  string  `pulumi:"type"`
}

func LookupPeeringServicePrefixOutput(ctx *pulumi.Context, args LookupPeeringServicePrefixOutputArgs, opts ...pulumi.InvokeOption) LookupPeeringServicePrefixResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupPeeringServicePrefixResult, error) {
			args := v.(LookupPeeringServicePrefixArgs)
			r, err := LookupPeeringServicePrefix(ctx, &args, opts...)
			var s LookupPeeringServicePrefixResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupPeeringServicePrefixResultOutput)
}

type LookupPeeringServicePrefixOutputArgs struct {
	PeeringServiceName pulumi.StringInput `pulumi:"peeringServiceName"`
	PrefixName         pulumi.StringInput `pulumi:"prefixName"`
	ResourceGroupName  pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupPeeringServicePrefixOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPeeringServicePrefixArgs)(nil)).Elem()
}


type LookupPeeringServicePrefixResultOutput struct{ *pulumi.OutputState }

func (LookupPeeringServicePrefixResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPeeringServicePrefixResult)(nil)).Elem()
}

func (o LookupPeeringServicePrefixResultOutput) ToLookupPeeringServicePrefixResultOutput() LookupPeeringServicePrefixResultOutput {
	return o
}

func (o LookupPeeringServicePrefixResultOutput) ToLookupPeeringServicePrefixResultOutputWithContext(ctx context.Context) LookupPeeringServicePrefixResultOutput {
	return o
}

func (o LookupPeeringServicePrefixResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPeeringServicePrefixResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupPeeringServicePrefixResultOutput) LearnedType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPeeringServicePrefixResult) *string { return v.LearnedType }).(pulumi.StringPtrOutput)
}

func (o LookupPeeringServicePrefixResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPeeringServicePrefixResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupPeeringServicePrefixResultOutput) Prefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPeeringServicePrefixResult) *string { return v.Prefix }).(pulumi.StringPtrOutput)
}

func (o LookupPeeringServicePrefixResultOutput) PrefixValidationState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPeeringServicePrefixResult) *string { return v.PrefixValidationState }).(pulumi.StringPtrOutput)
}

func (o LookupPeeringServicePrefixResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPeeringServicePrefixResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupPeeringServicePrefixResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPeeringServicePrefixResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupPeeringServicePrefixResultOutput{})
}
