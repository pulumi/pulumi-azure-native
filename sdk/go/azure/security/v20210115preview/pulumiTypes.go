


package v20210115preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type IngestionConnectionStringResponse struct {
	Location string `pulumi:"location"`
	Value    string `pulumi:"value"`
}

type IngestionConnectionStringResponseOutput struct{ *pulumi.OutputState }

func (IngestionConnectionStringResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IngestionConnectionStringResponse)(nil)).Elem()
}

func (o IngestionConnectionStringResponseOutput) ToIngestionConnectionStringResponseOutput() IngestionConnectionStringResponseOutput {
	return o
}

func (o IngestionConnectionStringResponseOutput) ToIngestionConnectionStringResponseOutputWithContext(ctx context.Context) IngestionConnectionStringResponseOutput {
	return o
}

func (o IngestionConnectionStringResponseOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v IngestionConnectionStringResponse) string { return v.Location }).(pulumi.StringOutput)
}

func (o IngestionConnectionStringResponseOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v IngestionConnectionStringResponse) string { return v.Value }).(pulumi.StringOutput)
}

type IngestionConnectionStringResponseArrayOutput struct{ *pulumi.OutputState }

func (IngestionConnectionStringResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]IngestionConnectionStringResponse)(nil)).Elem()
}

func (o IngestionConnectionStringResponseArrayOutput) ToIngestionConnectionStringResponseArrayOutput() IngestionConnectionStringResponseArrayOutput {
	return o
}

func (o IngestionConnectionStringResponseArrayOutput) ToIngestionConnectionStringResponseArrayOutputWithContext(ctx context.Context) IngestionConnectionStringResponseArrayOutput {
	return o
}

func (o IngestionConnectionStringResponseArrayOutput) Index(i pulumi.IntInput) IngestionConnectionStringResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) IngestionConnectionStringResponse {
		return vs[0].([]IngestionConnectionStringResponse)[vs[1].(int)]
	}).(IngestionConnectionStringResponseOutput)
}

func init() {
	pulumi.RegisterOutputType(IngestionConnectionStringResponseOutput{})
	pulumi.RegisterOutputType(IngestionConnectionStringResponseArrayOutput{})
}
