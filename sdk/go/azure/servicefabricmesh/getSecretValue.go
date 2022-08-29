


package servicefabricmesh

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupSecretValue(ctx *pulumi.Context, args *LookupSecretValueArgs, opts ...pulumi.InvokeOption) (*LookupSecretValueResult, error) {
	var rv LookupSecretValueResult
	err := ctx.Invoke("azure-native:servicefabricmesh:getSecretValue", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSecretValueArgs struct {
	ResourceGroupName       string `pulumi:"resourceGroupName"`
	SecretResourceName      string `pulumi:"secretResourceName"`
	SecretValueResourceName string `pulumi:"secretValueResourceName"`
}


type LookupSecretValueResult struct {
	Id                string            `pulumi:"id"`
	Location          string            `pulumi:"location"`
	Name              string            `pulumi:"name"`
	ProvisioningState string            `pulumi:"provisioningState"`
	Tags              map[string]string `pulumi:"tags"`
	Type              string            `pulumi:"type"`
	Value             *string           `pulumi:"value"`
}

func LookupSecretValueOutput(ctx *pulumi.Context, args LookupSecretValueOutputArgs, opts ...pulumi.InvokeOption) LookupSecretValueResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSecretValueResult, error) {
			args := v.(LookupSecretValueArgs)
			r, err := LookupSecretValue(ctx, &args, opts...)
			var s LookupSecretValueResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSecretValueResultOutput)
}

type LookupSecretValueOutputArgs struct {
	ResourceGroupName       pulumi.StringInput `pulumi:"resourceGroupName"`
	SecretResourceName      pulumi.StringInput `pulumi:"secretResourceName"`
	SecretValueResourceName pulumi.StringInput `pulumi:"secretValueResourceName"`
}

func (LookupSecretValueOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSecretValueArgs)(nil)).Elem()
}


type LookupSecretValueResultOutput struct{ *pulumi.OutputState }

func (LookupSecretValueResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSecretValueResult)(nil)).Elem()
}

func (o LookupSecretValueResultOutput) ToLookupSecretValueResultOutput() LookupSecretValueResultOutput {
	return o
}

func (o LookupSecretValueResultOutput) ToLookupSecretValueResultOutputWithContext(ctx context.Context) LookupSecretValueResultOutput {
	return o
}

func (o LookupSecretValueResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSecretValueResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupSecretValueResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSecretValueResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupSecretValueResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSecretValueResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupSecretValueResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSecretValueResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupSecretValueResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupSecretValueResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupSecretValueResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSecretValueResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupSecretValueResultOutput) Value() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSecretValueResult) *string { return v.Value }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSecretValueResultOutput{})
}
