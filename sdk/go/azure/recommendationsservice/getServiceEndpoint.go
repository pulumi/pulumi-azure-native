


package recommendationsservice

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupServiceEndpoint(ctx *pulumi.Context, args *LookupServiceEndpointArgs, opts ...pulumi.InvokeOption) (*LookupServiceEndpointResult, error) {
	var rv LookupServiceEndpointResult
	err := ctx.Invoke("azure-native:recommendationsservice:getServiceEndpoint", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupServiceEndpointArgs struct {
	AccountName         string `pulumi:"accountName"`
	ResourceGroupName   string `pulumi:"resourceGroupName"`
	ServiceEndpointName string `pulumi:"serviceEndpointName"`
}


type LookupServiceEndpointResult struct {
	Id         string                                    `pulumi:"id"`
	Location   string                                    `pulumi:"location"`
	Name       string                                    `pulumi:"name"`
	Properties ServiceEndpointResourceResponseProperties `pulumi:"properties"`
	SystemData SystemDataResponse                        `pulumi:"systemData"`
	Tags       map[string]string                         `pulumi:"tags"`
	Type       string                                    `pulumi:"type"`
}

func LookupServiceEndpointOutput(ctx *pulumi.Context, args LookupServiceEndpointOutputArgs, opts ...pulumi.InvokeOption) LookupServiceEndpointResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupServiceEndpointResult, error) {
			args := v.(LookupServiceEndpointArgs)
			r, err := LookupServiceEndpoint(ctx, &args, opts...)
			var s LookupServiceEndpointResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupServiceEndpointResultOutput)
}

type LookupServiceEndpointOutputArgs struct {
	AccountName         pulumi.StringInput `pulumi:"accountName"`
	ResourceGroupName   pulumi.StringInput `pulumi:"resourceGroupName"`
	ServiceEndpointName pulumi.StringInput `pulumi:"serviceEndpointName"`
}

func (LookupServiceEndpointOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupServiceEndpointArgs)(nil)).Elem()
}


type LookupServiceEndpointResultOutput struct{ *pulumi.OutputState }

func (LookupServiceEndpointResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupServiceEndpointResult)(nil)).Elem()
}

func (o LookupServiceEndpointResultOutput) ToLookupServiceEndpointResultOutput() LookupServiceEndpointResultOutput {
	return o
}

func (o LookupServiceEndpointResultOutput) ToLookupServiceEndpointResultOutputWithContext(ctx context.Context) LookupServiceEndpointResultOutput {
	return o
}

func (o LookupServiceEndpointResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServiceEndpointResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupServiceEndpointResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServiceEndpointResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupServiceEndpointResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServiceEndpointResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupServiceEndpointResultOutput) Properties() ServiceEndpointResourceResponsePropertiesOutput {
	return o.ApplyT(func(v LookupServiceEndpointResult) ServiceEndpointResourceResponseProperties { return v.Properties }).(ServiceEndpointResourceResponsePropertiesOutput)
}

func (o LookupServiceEndpointResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupServiceEndpointResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupServiceEndpointResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupServiceEndpointResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupServiceEndpointResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServiceEndpointResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupServiceEndpointResultOutput{})
}
