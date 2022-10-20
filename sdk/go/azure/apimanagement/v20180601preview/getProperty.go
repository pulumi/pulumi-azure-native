


package v20180601preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupProperty(ctx *pulumi.Context, args *LookupPropertyArgs, opts ...pulumi.InvokeOption) (*LookupPropertyResult, error) {
	var rv LookupPropertyResult
	err := ctx.Invoke("azure-native:apimanagement/v20180601preview:getProperty", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupPropertyArgs struct {
	PropId            string `pulumi:"propId"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ServiceName       string `pulumi:"serviceName"`
}


type LookupPropertyResult struct {
	DisplayName string   `pulumi:"displayName"`
	Id          string   `pulumi:"id"`
	Name        string   `pulumi:"name"`
	Secret      *bool    `pulumi:"secret"`
	Tags        []string `pulumi:"tags"`
	Type        string   `pulumi:"type"`
	Value       string   `pulumi:"value"`
}

func LookupPropertyOutput(ctx *pulumi.Context, args LookupPropertyOutputArgs, opts ...pulumi.InvokeOption) LookupPropertyResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupPropertyResult, error) {
			args := v.(LookupPropertyArgs)
			r, err := LookupProperty(ctx, &args, opts...)
			var s LookupPropertyResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupPropertyResultOutput)
}

type LookupPropertyOutputArgs struct {
	PropId            pulumi.StringInput `pulumi:"propId"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	ServiceName       pulumi.StringInput `pulumi:"serviceName"`
}

func (LookupPropertyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPropertyArgs)(nil)).Elem()
}


type LookupPropertyResultOutput struct{ *pulumi.OutputState }

func (LookupPropertyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPropertyResult)(nil)).Elem()
}

func (o LookupPropertyResultOutput) ToLookupPropertyResultOutput() LookupPropertyResultOutput {
	return o
}

func (o LookupPropertyResultOutput) ToLookupPropertyResultOutputWithContext(ctx context.Context) LookupPropertyResultOutput {
	return o
}

func (o LookupPropertyResultOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPropertyResult) string { return v.DisplayName }).(pulumi.StringOutput)
}

func (o LookupPropertyResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPropertyResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupPropertyResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPropertyResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupPropertyResultOutput) Secret() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupPropertyResult) *bool { return v.Secret }).(pulumi.BoolPtrOutput)
}

func (o LookupPropertyResultOutput) Tags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupPropertyResult) []string { return v.Tags }).(pulumi.StringArrayOutput)
}

func (o LookupPropertyResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPropertyResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupPropertyResultOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPropertyResult) string { return v.Value }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupPropertyResultOutput{})
}
