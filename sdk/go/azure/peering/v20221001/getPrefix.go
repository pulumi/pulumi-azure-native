


package v20221001

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupPrefix(ctx *pulumi.Context, args *LookupPrefixArgs, opts ...pulumi.InvokeOption) (*LookupPrefixResult, error) {
	var rv LookupPrefixResult
	err := ctx.Invoke("azure-native:peering/v20221001:getPrefix", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupPrefixArgs struct {
	Expand             *string `pulumi:"expand"`
	PeeringServiceName string  `pulumi:"peeringServiceName"`
	PrefixName         string  `pulumi:"prefixName"`
	ResourceGroupName  string  `pulumi:"resourceGroupName"`
}


type LookupPrefixResult struct {
	ErrorMessage            string                              `pulumi:"errorMessage"`
	Events                  []PeeringServicePrefixEventResponse `pulumi:"events"`
	Id                      string                              `pulumi:"id"`
	LearnedType             string                              `pulumi:"learnedType"`
	Name                    string                              `pulumi:"name"`
	PeeringServicePrefixKey *string                             `pulumi:"peeringServicePrefixKey"`
	Prefix                  *string                             `pulumi:"prefix"`
	PrefixValidationState   string                              `pulumi:"prefixValidationState"`
	ProvisioningState       string                              `pulumi:"provisioningState"`
	Type                    string                              `pulumi:"type"`
}

func LookupPrefixOutput(ctx *pulumi.Context, args LookupPrefixOutputArgs, opts ...pulumi.InvokeOption) LookupPrefixResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupPrefixResult, error) {
			args := v.(LookupPrefixArgs)
			r, err := LookupPrefix(ctx, &args, opts...)
			var s LookupPrefixResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupPrefixResultOutput)
}

type LookupPrefixOutputArgs struct {
	Expand             pulumi.StringPtrInput `pulumi:"expand"`
	PeeringServiceName pulumi.StringInput    `pulumi:"peeringServiceName"`
	PrefixName         pulumi.StringInput    `pulumi:"prefixName"`
	ResourceGroupName  pulumi.StringInput    `pulumi:"resourceGroupName"`
}

func (LookupPrefixOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPrefixArgs)(nil)).Elem()
}


type LookupPrefixResultOutput struct{ *pulumi.OutputState }

func (LookupPrefixResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPrefixResult)(nil)).Elem()
}

func (o LookupPrefixResultOutput) ToLookupPrefixResultOutput() LookupPrefixResultOutput {
	return o
}

func (o LookupPrefixResultOutput) ToLookupPrefixResultOutputWithContext(ctx context.Context) LookupPrefixResultOutput {
	return o
}

func (o LookupPrefixResultOutput) ErrorMessage() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrefixResult) string { return v.ErrorMessage }).(pulumi.StringOutput)
}

func (o LookupPrefixResultOutput) Events() PeeringServicePrefixEventResponseArrayOutput {
	return o.ApplyT(func(v LookupPrefixResult) []PeeringServicePrefixEventResponse { return v.Events }).(PeeringServicePrefixEventResponseArrayOutput)
}

func (o LookupPrefixResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrefixResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupPrefixResultOutput) LearnedType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrefixResult) string { return v.LearnedType }).(pulumi.StringOutput)
}

func (o LookupPrefixResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrefixResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupPrefixResultOutput) PeeringServicePrefixKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPrefixResult) *string { return v.PeeringServicePrefixKey }).(pulumi.StringPtrOutput)
}

func (o LookupPrefixResultOutput) Prefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPrefixResult) *string { return v.Prefix }).(pulumi.StringPtrOutput)
}

func (o LookupPrefixResultOutput) PrefixValidationState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrefixResult) string { return v.PrefixValidationState }).(pulumi.StringOutput)
}

func (o LookupPrefixResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrefixResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupPrefixResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrefixResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupPrefixResultOutput{})
}
