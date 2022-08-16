


package v20171101preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupExtension(ctx *pulumi.Context, args *LookupExtensionArgs, opts ...pulumi.InvokeOption) (*LookupExtensionResult, error) {
	var rv LookupExtensionResult
	err := ctx.Invoke("azure-native:visualstudio/v20171101preview:getExtension", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupExtensionArgs struct {
	AccountResourceName   string `pulumi:"accountResourceName"`
	ExtensionResourceName string `pulumi:"extensionResourceName"`
	ResourceGroupName     string `pulumi:"resourceGroupName"`
}


type LookupExtensionResult struct {
	Id         string                         `pulumi:"id"`
	Location   *string                        `pulumi:"location"`
	Name       string                         `pulumi:"name"`
	Plan       *ExtensionResourcePlanResponse `pulumi:"plan"`
	Properties map[string]string              `pulumi:"properties"`
	Tags       map[string]string              `pulumi:"tags"`
	Type       string                         `pulumi:"type"`
}

func LookupExtensionOutput(ctx *pulumi.Context, args LookupExtensionOutputArgs, opts ...pulumi.InvokeOption) LookupExtensionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupExtensionResult, error) {
			args := v.(LookupExtensionArgs)
			r, err := LookupExtension(ctx, &args, opts...)
			var s LookupExtensionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupExtensionResultOutput)
}

type LookupExtensionOutputArgs struct {
	AccountResourceName   pulumi.StringInput `pulumi:"accountResourceName"`
	ExtensionResourceName pulumi.StringInput `pulumi:"extensionResourceName"`
	ResourceGroupName     pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupExtensionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupExtensionArgs)(nil)).Elem()
}


type LookupExtensionResultOutput struct{ *pulumi.OutputState }

func (LookupExtensionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupExtensionResult)(nil)).Elem()
}

func (o LookupExtensionResultOutput) ToLookupExtensionResultOutput() LookupExtensionResultOutput {
	return o
}

func (o LookupExtensionResultOutput) ToLookupExtensionResultOutputWithContext(ctx context.Context) LookupExtensionResultOutput {
	return o
}

func (o LookupExtensionResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupExtensionResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupExtensionResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupExtensionResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupExtensionResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupExtensionResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupExtensionResultOutput) Plan() ExtensionResourcePlanResponsePtrOutput {
	return o.ApplyT(func(v LookupExtensionResult) *ExtensionResourcePlanResponse { return v.Plan }).(ExtensionResourcePlanResponsePtrOutput)
}

func (o LookupExtensionResultOutput) Properties() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupExtensionResult) map[string]string { return v.Properties }).(pulumi.StringMapOutput)
}

func (o LookupExtensionResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupExtensionResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupExtensionResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupExtensionResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupExtensionResultOutput{})
}
