


package servicefabricmesh

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupSecret(ctx *pulumi.Context, args *LookupSecretArgs, opts ...pulumi.InvokeOption) (*LookupSecretResult, error) {
	var rv LookupSecretResult
	err := ctx.Invoke("azure-native:servicefabricmesh:getSecret", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSecretArgs struct {
	ResourceGroupName  string `pulumi:"resourceGroupName"`
	SecretResourceName string `pulumi:"secretResourceName"`
}


type LookupSecretResult struct {
	Id         string                           `pulumi:"id"`
	Location   string                           `pulumi:"location"`
	Name       string                           `pulumi:"name"`
	Properties SecretResourcePropertiesResponse `pulumi:"properties"`
	Tags       map[string]string                `pulumi:"tags"`
	Type       string                           `pulumi:"type"`
}

func LookupSecretOutput(ctx *pulumi.Context, args LookupSecretOutputArgs, opts ...pulumi.InvokeOption) LookupSecretResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSecretResult, error) {
			args := v.(LookupSecretArgs)
			r, err := LookupSecret(ctx, &args, opts...)
			var s LookupSecretResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSecretResultOutput)
}

type LookupSecretOutputArgs struct {
	ResourceGroupName  pulumi.StringInput `pulumi:"resourceGroupName"`
	SecretResourceName pulumi.StringInput `pulumi:"secretResourceName"`
}

func (LookupSecretOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSecretArgs)(nil)).Elem()
}


type LookupSecretResultOutput struct{ *pulumi.OutputState }

func (LookupSecretResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSecretResult)(nil)).Elem()
}

func (o LookupSecretResultOutput) ToLookupSecretResultOutput() LookupSecretResultOutput {
	return o
}

func (o LookupSecretResultOutput) ToLookupSecretResultOutputWithContext(ctx context.Context) LookupSecretResultOutput {
	return o
}

func (o LookupSecretResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSecretResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupSecretResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSecretResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupSecretResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSecretResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupSecretResultOutput) Properties() SecretResourcePropertiesResponseOutput {
	return o.ApplyT(func(v LookupSecretResult) SecretResourcePropertiesResponse { return v.Properties }).(SecretResourcePropertiesResponseOutput)
}

func (o LookupSecretResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupSecretResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupSecretResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSecretResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSecretResultOutput{})
}
