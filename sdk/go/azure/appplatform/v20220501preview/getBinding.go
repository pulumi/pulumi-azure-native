


package v20220501preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupBinding(ctx *pulumi.Context, args *LookupBindingArgs, opts ...pulumi.InvokeOption) (*LookupBindingResult, error) {
	var rv LookupBindingResult
	err := ctx.Invoke("azure-native:appplatform/v20220501preview:getBinding", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupBindingArgs struct {
	AppName           string `pulumi:"appName"`
	BindingName       string `pulumi:"bindingName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ServiceName       string `pulumi:"serviceName"`
}


type LookupBindingResult struct {
	Id         string                            `pulumi:"id"`
	Name       string                            `pulumi:"name"`
	Properties BindingResourcePropertiesResponse `pulumi:"properties"`
	SystemData SystemDataResponse                `pulumi:"systemData"`
	Type       string                            `pulumi:"type"`
}

func LookupBindingOutput(ctx *pulumi.Context, args LookupBindingOutputArgs, opts ...pulumi.InvokeOption) LookupBindingResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupBindingResult, error) {
			args := v.(LookupBindingArgs)
			r, err := LookupBinding(ctx, &args, opts...)
			var s LookupBindingResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupBindingResultOutput)
}

type LookupBindingOutputArgs struct {
	AppName           pulumi.StringInput `pulumi:"appName"`
	BindingName       pulumi.StringInput `pulumi:"bindingName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	ServiceName       pulumi.StringInput `pulumi:"serviceName"`
}

func (LookupBindingOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupBindingArgs)(nil)).Elem()
}


type LookupBindingResultOutput struct{ *pulumi.OutputState }

func (LookupBindingResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupBindingResult)(nil)).Elem()
}

func (o LookupBindingResultOutput) ToLookupBindingResultOutput() LookupBindingResultOutput {
	return o
}

func (o LookupBindingResultOutput) ToLookupBindingResultOutputWithContext(ctx context.Context) LookupBindingResultOutput {
	return o
}

func (o LookupBindingResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBindingResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupBindingResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBindingResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupBindingResultOutput) Properties() BindingResourcePropertiesResponseOutput {
	return o.ApplyT(func(v LookupBindingResult) BindingResourcePropertiesResponse { return v.Properties }).(BindingResourcePropertiesResponseOutput)
}

func (o LookupBindingResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupBindingResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupBindingResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBindingResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupBindingResultOutput{})
}
