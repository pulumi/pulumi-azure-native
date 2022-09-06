


package v20200201preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupMediaGraph(ctx *pulumi.Context, args *LookupMediaGraphArgs, opts ...pulumi.InvokeOption) (*LookupMediaGraphResult, error) {
	var rv LookupMediaGraphResult
	err := ctx.Invoke("azure-native:media/v20200201preview:getMediaGraph", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupMediaGraphArgs struct {
	AccountName       string `pulumi:"accountName"`
	MediaGraphName    string `pulumi:"mediaGraphName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupMediaGraphResult struct {
	Created      string                         `pulumi:"created"`
	Description  *string                        `pulumi:"description"`
	Id           string                         `pulumi:"id"`
	LastModified string                         `pulumi:"lastModified"`
	Name         string                         `pulumi:"name"`
	Sinks        []MediaGraphAssetSinkResponse  `pulumi:"sinks"`
	Sources      []MediaGraphRtspSourceResponse `pulumi:"sources"`
	State        string                         `pulumi:"state"`
	Type         string                         `pulumi:"type"`
}

func LookupMediaGraphOutput(ctx *pulumi.Context, args LookupMediaGraphOutputArgs, opts ...pulumi.InvokeOption) LookupMediaGraphResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupMediaGraphResult, error) {
			args := v.(LookupMediaGraphArgs)
			r, err := LookupMediaGraph(ctx, &args, opts...)
			var s LookupMediaGraphResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupMediaGraphResultOutput)
}

type LookupMediaGraphOutputArgs struct {
	AccountName       pulumi.StringInput `pulumi:"accountName"`
	MediaGraphName    pulumi.StringInput `pulumi:"mediaGraphName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupMediaGraphOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupMediaGraphArgs)(nil)).Elem()
}


type LookupMediaGraphResultOutput struct{ *pulumi.OutputState }

func (LookupMediaGraphResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupMediaGraphResult)(nil)).Elem()
}

func (o LookupMediaGraphResultOutput) ToLookupMediaGraphResultOutput() LookupMediaGraphResultOutput {
	return o
}

func (o LookupMediaGraphResultOutput) ToLookupMediaGraphResultOutputWithContext(ctx context.Context) LookupMediaGraphResultOutput {
	return o
}

func (o LookupMediaGraphResultOutput) Created() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMediaGraphResult) string { return v.Created }).(pulumi.StringOutput)
}

func (o LookupMediaGraphResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMediaGraphResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupMediaGraphResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMediaGraphResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupMediaGraphResultOutput) LastModified() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMediaGraphResult) string { return v.LastModified }).(pulumi.StringOutput)
}

func (o LookupMediaGraphResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMediaGraphResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupMediaGraphResultOutput) Sinks() MediaGraphAssetSinkResponseArrayOutput {
	return o.ApplyT(func(v LookupMediaGraphResult) []MediaGraphAssetSinkResponse { return v.Sinks }).(MediaGraphAssetSinkResponseArrayOutput)
}

func (o LookupMediaGraphResultOutput) Sources() MediaGraphRtspSourceResponseArrayOutput {
	return o.ApplyT(func(v LookupMediaGraphResult) []MediaGraphRtspSourceResponse { return v.Sources }).(MediaGraphRtspSourceResponseArrayOutput)
}

func (o LookupMediaGraphResultOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMediaGraphResult) string { return v.State }).(pulumi.StringOutput)
}

func (o LookupMediaGraphResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMediaGraphResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupMediaGraphResultOutput{})
}
