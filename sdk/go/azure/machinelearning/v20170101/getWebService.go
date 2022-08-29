


package v20170101

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupWebService(ctx *pulumi.Context, args *LookupWebServiceArgs, opts ...pulumi.InvokeOption) (*LookupWebServiceResult, error) {
	var rv LookupWebServiceResult
	err := ctx.Invoke("azure-native:machinelearning/v20170101:getWebService", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupWebServiceArgs struct {
	Region            *string `pulumi:"region"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
	WebServiceName    string  `pulumi:"webServiceName"`
}


type LookupWebServiceResult struct {
	Id         string                               `pulumi:"id"`
	Location   string                               `pulumi:"location"`
	Name       string                               `pulumi:"name"`
	Properties WebServicePropertiesForGraphResponse `pulumi:"properties"`
	Tags       map[string]string                    `pulumi:"tags"`
	Type       string                               `pulumi:"type"`
}


func (val *LookupWebServiceResult) Defaults() *LookupWebServiceResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.Properties = *tmp.Properties.Defaults()

	return &tmp
}

func LookupWebServiceOutput(ctx *pulumi.Context, args LookupWebServiceOutputArgs, opts ...pulumi.InvokeOption) LookupWebServiceResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupWebServiceResult, error) {
			args := v.(LookupWebServiceArgs)
			r, err := LookupWebService(ctx, &args, opts...)
			var s LookupWebServiceResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupWebServiceResultOutput)
}

type LookupWebServiceOutputArgs struct {
	Region            pulumi.StringPtrInput `pulumi:"region"`
	ResourceGroupName pulumi.StringInput    `pulumi:"resourceGroupName"`
	WebServiceName    pulumi.StringInput    `pulumi:"webServiceName"`
}

func (LookupWebServiceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWebServiceArgs)(nil)).Elem()
}


type LookupWebServiceResultOutput struct{ *pulumi.OutputState }

func (LookupWebServiceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWebServiceResult)(nil)).Elem()
}

func (o LookupWebServiceResultOutput) ToLookupWebServiceResultOutput() LookupWebServiceResultOutput {
	return o
}

func (o LookupWebServiceResultOutput) ToLookupWebServiceResultOutputWithContext(ctx context.Context) LookupWebServiceResultOutput {
	return o
}

func (o LookupWebServiceResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebServiceResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupWebServiceResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebServiceResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupWebServiceResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebServiceResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupWebServiceResultOutput) Properties() WebServicePropertiesForGraphResponseOutput {
	return o.ApplyT(func(v LookupWebServiceResult) WebServicePropertiesForGraphResponse { return v.Properties }).(WebServicePropertiesForGraphResponseOutput)
}

func (o LookupWebServiceResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupWebServiceResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupWebServiceResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebServiceResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupWebServiceResultOutput{})
}
