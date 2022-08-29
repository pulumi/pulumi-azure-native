


package v20180901preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


func LookupAttestationProvider(ctx *pulumi.Context, args *LookupAttestationProviderArgs, opts ...pulumi.InvokeOption) (*LookupAttestationProviderResult, error) {
	var rv LookupAttestationProviderResult
	err := ctx.Invoke("azure-native:attestation/v20180901preview:getAttestationProvider", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupAttestationProviderArgs struct {
	ProviderName      string `pulumi:"providerName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupAttestationProviderResult struct {
	AttestUri  *string           `pulumi:"attestUri"`
	Id         string            `pulumi:"id"`
	Location   string            `pulumi:"location"`
	Name       string            `pulumi:"name"`
	Status     *string           `pulumi:"status"`
	Tags       map[string]string `pulumi:"tags"`
	TrustModel *string           `pulumi:"trustModel"`
	Type       string            `pulumi:"type"`
}

func LookupAttestationProviderOutput(ctx *pulumi.Context, args LookupAttestationProviderOutputArgs, opts ...pulumi.InvokeOption) LookupAttestationProviderResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupAttestationProviderResult, error) {
			args := v.(LookupAttestationProviderArgs)
			r, err := LookupAttestationProvider(ctx, &args, opts...)
			var s LookupAttestationProviderResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupAttestationProviderResultOutput)
}

type LookupAttestationProviderOutputArgs struct {
	ProviderName      pulumi.StringInput `pulumi:"providerName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupAttestationProviderOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAttestationProviderArgs)(nil)).Elem()
}


type LookupAttestationProviderResultOutput struct{ *pulumi.OutputState }

func (LookupAttestationProviderResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAttestationProviderResult)(nil)).Elem()
}

func (o LookupAttestationProviderResultOutput) ToLookupAttestationProviderResultOutput() LookupAttestationProviderResultOutput {
	return o
}

func (o LookupAttestationProviderResultOutput) ToLookupAttestationProviderResultOutputWithContext(ctx context.Context) LookupAttestationProviderResultOutput {
	return o
}

func (o LookupAttestationProviderResultOutput) AttestUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAttestationProviderResult) *string { return v.AttestUri }).(pulumi.StringPtrOutput)
}

func (o LookupAttestationProviderResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAttestationProviderResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupAttestationProviderResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAttestationProviderResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupAttestationProviderResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAttestationProviderResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupAttestationProviderResultOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAttestationProviderResult) *string { return v.Status }).(pulumi.StringPtrOutput)
}

func (o LookupAttestationProviderResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupAttestationProviderResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupAttestationProviderResultOutput) TrustModel() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAttestationProviderResult) *string { return v.TrustModel }).(pulumi.StringPtrOutput)
}

func (o LookupAttestationProviderResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAttestationProviderResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupAttestationProviderResultOutput{})
}
