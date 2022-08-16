


package v20210101

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupRegisteredPrefix(ctx *pulumi.Context, args *LookupRegisteredPrefixArgs, opts ...pulumi.InvokeOption) (*LookupRegisteredPrefixResult, error) {
	var rv LookupRegisteredPrefixResult
	err := ctx.Invoke("azure-native:peering/v20210101:getRegisteredPrefix", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupRegisteredPrefixArgs struct {
	PeeringName          string `pulumi:"peeringName"`
	RegisteredPrefixName string `pulumi:"registeredPrefixName"`
	ResourceGroupName    string `pulumi:"resourceGroupName"`
}


type LookupRegisteredPrefixResult struct {
	ErrorMessage            string  `pulumi:"errorMessage"`
	Id                      string  `pulumi:"id"`
	Name                    string  `pulumi:"name"`
	PeeringServicePrefixKey string  `pulumi:"peeringServicePrefixKey"`
	Prefix                  *string `pulumi:"prefix"`
	PrefixValidationState   string  `pulumi:"prefixValidationState"`
	ProvisioningState       string  `pulumi:"provisioningState"`
	Type                    string  `pulumi:"type"`
}

func LookupRegisteredPrefixOutput(ctx *pulumi.Context, args LookupRegisteredPrefixOutputArgs, opts ...pulumi.InvokeOption) LookupRegisteredPrefixResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupRegisteredPrefixResult, error) {
			args := v.(LookupRegisteredPrefixArgs)
			r, err := LookupRegisteredPrefix(ctx, &args, opts...)
			var s LookupRegisteredPrefixResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupRegisteredPrefixResultOutput)
}

type LookupRegisteredPrefixOutputArgs struct {
	PeeringName          pulumi.StringInput `pulumi:"peeringName"`
	RegisteredPrefixName pulumi.StringInput `pulumi:"registeredPrefixName"`
	ResourceGroupName    pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupRegisteredPrefixOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRegisteredPrefixArgs)(nil)).Elem()
}


type LookupRegisteredPrefixResultOutput struct{ *pulumi.OutputState }

func (LookupRegisteredPrefixResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRegisteredPrefixResult)(nil)).Elem()
}

func (o LookupRegisteredPrefixResultOutput) ToLookupRegisteredPrefixResultOutput() LookupRegisteredPrefixResultOutput {
	return o
}

func (o LookupRegisteredPrefixResultOutput) ToLookupRegisteredPrefixResultOutputWithContext(ctx context.Context) LookupRegisteredPrefixResultOutput {
	return o
}

func (o LookupRegisteredPrefixResultOutput) ErrorMessage() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRegisteredPrefixResult) string { return v.ErrorMessage }).(pulumi.StringOutput)
}

func (o LookupRegisteredPrefixResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRegisteredPrefixResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupRegisteredPrefixResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRegisteredPrefixResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupRegisteredPrefixResultOutput) PeeringServicePrefixKey() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRegisteredPrefixResult) string { return v.PeeringServicePrefixKey }).(pulumi.StringOutput)
}

func (o LookupRegisteredPrefixResultOutput) Prefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRegisteredPrefixResult) *string { return v.Prefix }).(pulumi.StringPtrOutput)
}

func (o LookupRegisteredPrefixResultOutput) PrefixValidationState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRegisteredPrefixResult) string { return v.PrefixValidationState }).(pulumi.StringOutput)
}

func (o LookupRegisteredPrefixResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRegisteredPrefixResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupRegisteredPrefixResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRegisteredPrefixResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupRegisteredPrefixResultOutput{})
}
