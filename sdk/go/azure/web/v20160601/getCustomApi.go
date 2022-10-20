


package v20160601

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupCustomApi(ctx *pulumi.Context, args *LookupCustomApiArgs, opts ...pulumi.InvokeOption) (*LookupCustomApiResult, error) {
	var rv LookupCustomApiResult
	err := ctx.Invoke("azure-native:web/v20160601:getCustomApi", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupCustomApiArgs struct {
	ApiName           string  `pulumi:"apiName"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
	SubscriptionId    *string `pulumi:"subscriptionId"`
}


type LookupCustomApiResult struct {
	Etag       *string                               `pulumi:"etag"`
	Id         string                                `pulumi:"id"`
	Location   *string                               `pulumi:"location"`
	Name       string                                `pulumi:"name"`
	Properties CustomApiPropertiesDefinitionResponse `pulumi:"properties"`
	Tags       map[string]string                     `pulumi:"tags"`
	Type       string                                `pulumi:"type"`
}

func LookupCustomApiOutput(ctx *pulumi.Context, args LookupCustomApiOutputArgs, opts ...pulumi.InvokeOption) LookupCustomApiResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupCustomApiResult, error) {
			args := v.(LookupCustomApiArgs)
			r, err := LookupCustomApi(ctx, &args, opts...)
			var s LookupCustomApiResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupCustomApiResultOutput)
}

type LookupCustomApiOutputArgs struct {
	ApiName           pulumi.StringInput    `pulumi:"apiName"`
	ResourceGroupName pulumi.StringInput    `pulumi:"resourceGroupName"`
	SubscriptionId    pulumi.StringPtrInput `pulumi:"subscriptionId"`
}

func (LookupCustomApiOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCustomApiArgs)(nil)).Elem()
}


type LookupCustomApiResultOutput struct{ *pulumi.OutputState }

func (LookupCustomApiResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCustomApiResult)(nil)).Elem()
}

func (o LookupCustomApiResultOutput) ToLookupCustomApiResultOutput() LookupCustomApiResultOutput {
	return o
}

func (o LookupCustomApiResultOutput) ToLookupCustomApiResultOutputWithContext(ctx context.Context) LookupCustomApiResultOutput {
	return o
}

func (o LookupCustomApiResultOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCustomApiResult) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o LookupCustomApiResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCustomApiResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupCustomApiResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCustomApiResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupCustomApiResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCustomApiResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupCustomApiResultOutput) Properties() CustomApiPropertiesDefinitionResponseOutput {
	return o.ApplyT(func(v LookupCustomApiResult) CustomApiPropertiesDefinitionResponse { return v.Properties }).(CustomApiPropertiesDefinitionResponseOutput)
}

func (o LookupCustomApiResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupCustomApiResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupCustomApiResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCustomApiResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupCustomApiResultOutput{})
}
