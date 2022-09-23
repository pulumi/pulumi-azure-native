


package v20220301

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupProtectionIntent(ctx *pulumi.Context, args *LookupProtectionIntentArgs, opts ...pulumi.InvokeOption) (*LookupProtectionIntentResult, error) {
	var rv LookupProtectionIntentResult
	err := ctx.Invoke("azure-native:recoveryservices/v20220301:getProtectionIntent", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupProtectionIntentArgs struct {
	FabricName        string `pulumi:"fabricName"`
	IntentObjectName  string `pulumi:"intentObjectName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	VaultName         string `pulumi:"vaultName"`
}


type LookupProtectionIntentResult struct {
	ETag       *string           `pulumi:"eTag"`
	Id         string            `pulumi:"id"`
	Location   *string           `pulumi:"location"`
	Name       string            `pulumi:"name"`
	Properties interface{}       `pulumi:"properties"`
	Tags       map[string]string `pulumi:"tags"`
	Type       string            `pulumi:"type"`
}

func LookupProtectionIntentOutput(ctx *pulumi.Context, args LookupProtectionIntentOutputArgs, opts ...pulumi.InvokeOption) LookupProtectionIntentResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupProtectionIntentResult, error) {
			args := v.(LookupProtectionIntentArgs)
			r, err := LookupProtectionIntent(ctx, &args, opts...)
			var s LookupProtectionIntentResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupProtectionIntentResultOutput)
}

type LookupProtectionIntentOutputArgs struct {
	FabricName        pulumi.StringInput `pulumi:"fabricName"`
	IntentObjectName  pulumi.StringInput `pulumi:"intentObjectName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	VaultName         pulumi.StringInput `pulumi:"vaultName"`
}

func (LookupProtectionIntentOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupProtectionIntentArgs)(nil)).Elem()
}


type LookupProtectionIntentResultOutput struct{ *pulumi.OutputState }

func (LookupProtectionIntentResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupProtectionIntentResult)(nil)).Elem()
}

func (o LookupProtectionIntentResultOutput) ToLookupProtectionIntentResultOutput() LookupProtectionIntentResultOutput {
	return o
}

func (o LookupProtectionIntentResultOutput) ToLookupProtectionIntentResultOutputWithContext(ctx context.Context) LookupProtectionIntentResultOutput {
	return o
}

func (o LookupProtectionIntentResultOutput) ETag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupProtectionIntentResult) *string { return v.ETag }).(pulumi.StringPtrOutput)
}

func (o LookupProtectionIntentResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProtectionIntentResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupProtectionIntentResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupProtectionIntentResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupProtectionIntentResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProtectionIntentResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupProtectionIntentResultOutput) Properties() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupProtectionIntentResult) interface{} { return v.Properties }).(pulumi.AnyOutput)
}

func (o LookupProtectionIntentResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupProtectionIntentResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupProtectionIntentResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProtectionIntentResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupProtectionIntentResultOutput{})
}
