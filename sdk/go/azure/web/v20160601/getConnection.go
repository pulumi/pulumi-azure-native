


package v20160601

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupConnection(ctx *pulumi.Context, args *LookupConnectionArgs, opts ...pulumi.InvokeOption) (*LookupConnectionResult, error) {
	var rv LookupConnectionResult
	err := ctx.Invoke("azure-native:web/v20160601:getConnection", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupConnectionArgs struct {
	ConnectionName    string  `pulumi:"connectionName"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
	SubscriptionId    *string `pulumi:"subscriptionId"`
}


type LookupConnectionResult struct {
	Etag       *string                                   `pulumi:"etag"`
	Id         string                                    `pulumi:"id"`
	Location   *string                                   `pulumi:"location"`
	Name       string                                    `pulumi:"name"`
	Properties ApiConnectionDefinitionResponseProperties `pulumi:"properties"`
	Tags       map[string]string                         `pulumi:"tags"`
	Type       string                                    `pulumi:"type"`
}

func LookupConnectionOutput(ctx *pulumi.Context, args LookupConnectionOutputArgs, opts ...pulumi.InvokeOption) LookupConnectionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupConnectionResult, error) {
			args := v.(LookupConnectionArgs)
			r, err := LookupConnection(ctx, &args, opts...)
			var s LookupConnectionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupConnectionResultOutput)
}

type LookupConnectionOutputArgs struct {
	ConnectionName    pulumi.StringInput    `pulumi:"connectionName"`
	ResourceGroupName pulumi.StringInput    `pulumi:"resourceGroupName"`
	SubscriptionId    pulumi.StringPtrInput `pulumi:"subscriptionId"`
}

func (LookupConnectionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupConnectionArgs)(nil)).Elem()
}


type LookupConnectionResultOutput struct{ *pulumi.OutputState }

func (LookupConnectionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupConnectionResult)(nil)).Elem()
}

func (o LookupConnectionResultOutput) ToLookupConnectionResultOutput() LookupConnectionResultOutput {
	return o
}

func (o LookupConnectionResultOutput) ToLookupConnectionResultOutputWithContext(ctx context.Context) LookupConnectionResultOutput {
	return o
}

func (o LookupConnectionResultOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupConnectionResult) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o LookupConnectionResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConnectionResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupConnectionResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupConnectionResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupConnectionResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConnectionResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupConnectionResultOutput) Properties() ApiConnectionDefinitionResponsePropertiesOutput {
	return o.ApplyT(func(v LookupConnectionResult) ApiConnectionDefinitionResponseProperties { return v.Properties }).(ApiConnectionDefinitionResponsePropertiesOutput)
}

func (o LookupConnectionResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupConnectionResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupConnectionResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConnectionResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupConnectionResultOutput{})
}
