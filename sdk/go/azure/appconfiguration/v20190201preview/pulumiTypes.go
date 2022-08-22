


package v20190201preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ApiKeyResponse struct {
	ConnectionString string `pulumi:"connectionString"`
	Id               string `pulumi:"id"`
	LastModified     string `pulumi:"lastModified"`
	Name             string `pulumi:"name"`
	ReadOnly         bool   `pulumi:"readOnly"`
	Value            string `pulumi:"value"`
}

type ApiKeyResponseOutput struct{ *pulumi.OutputState }

func (ApiKeyResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApiKeyResponse)(nil)).Elem()
}

func (o ApiKeyResponseOutput) ToApiKeyResponseOutput() ApiKeyResponseOutput {
	return o
}

func (o ApiKeyResponseOutput) ToApiKeyResponseOutputWithContext(ctx context.Context) ApiKeyResponseOutput {
	return o
}

func (o ApiKeyResponseOutput) ConnectionString() pulumi.StringOutput {
	return o.ApplyT(func(v ApiKeyResponse) string { return v.ConnectionString }).(pulumi.StringOutput)
}

func (o ApiKeyResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v ApiKeyResponse) string { return v.Id }).(pulumi.StringOutput)
}

func (o ApiKeyResponseOutput) LastModified() pulumi.StringOutput {
	return o.ApplyT(func(v ApiKeyResponse) string { return v.LastModified }).(pulumi.StringOutput)
}

func (o ApiKeyResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ApiKeyResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o ApiKeyResponseOutput) ReadOnly() pulumi.BoolOutput {
	return o.ApplyT(func(v ApiKeyResponse) bool { return v.ReadOnly }).(pulumi.BoolOutput)
}

func (o ApiKeyResponseOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v ApiKeyResponse) string { return v.Value }).(pulumi.StringOutput)
}

type ApiKeyResponseArrayOutput struct{ *pulumi.OutputState }

func (ApiKeyResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ApiKeyResponse)(nil)).Elem()
}

func (o ApiKeyResponseArrayOutput) ToApiKeyResponseArrayOutput() ApiKeyResponseArrayOutput {
	return o
}

func (o ApiKeyResponseArrayOutput) ToApiKeyResponseArrayOutputWithContext(ctx context.Context) ApiKeyResponseArrayOutput {
	return o
}

func (o ApiKeyResponseArrayOutput) Index(i pulumi.IntInput) ApiKeyResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ApiKeyResponse {
		return vs[0].([]ApiKeyResponse)[vs[1].(int)]
	}).(ApiKeyResponseOutput)
}

func init() {
	pulumi.RegisterOutputType(ApiKeyResponseOutput{})
	pulumi.RegisterOutputType(ApiKeyResponseArrayOutput{})
}
