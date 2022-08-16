


package maps

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupPrivateAtlase(ctx *pulumi.Context, args *LookupPrivateAtlaseArgs, opts ...pulumi.InvokeOption) (*LookupPrivateAtlaseResult, error) {
	var rv LookupPrivateAtlaseResult
	err := ctx.Invoke("azure-native:maps:getPrivateAtlase", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupPrivateAtlaseArgs struct {
	AccountName       string `pulumi:"accountName"`
	PrivateAtlasName  string `pulumi:"privateAtlasName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupPrivateAtlaseResult struct {
	Id         string                         `pulumi:"id"`
	Location   string                         `pulumi:"location"`
	Name       string                         `pulumi:"name"`
	Properties PrivateAtlasPropertiesResponse `pulumi:"properties"`
	Tags       map[string]string              `pulumi:"tags"`
	Type       string                         `pulumi:"type"`
}

func LookupPrivateAtlaseOutput(ctx *pulumi.Context, args LookupPrivateAtlaseOutputArgs, opts ...pulumi.InvokeOption) LookupPrivateAtlaseResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupPrivateAtlaseResult, error) {
			args := v.(LookupPrivateAtlaseArgs)
			r, err := LookupPrivateAtlase(ctx, &args, opts...)
			var s LookupPrivateAtlaseResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupPrivateAtlaseResultOutput)
}

type LookupPrivateAtlaseOutputArgs struct {
	AccountName       pulumi.StringInput `pulumi:"accountName"`
	PrivateAtlasName  pulumi.StringInput `pulumi:"privateAtlasName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupPrivateAtlaseOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPrivateAtlaseArgs)(nil)).Elem()
}


type LookupPrivateAtlaseResultOutput struct{ *pulumi.OutputState }

func (LookupPrivateAtlaseResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPrivateAtlaseResult)(nil)).Elem()
}

func (o LookupPrivateAtlaseResultOutput) ToLookupPrivateAtlaseResultOutput() LookupPrivateAtlaseResultOutput {
	return o
}

func (o LookupPrivateAtlaseResultOutput) ToLookupPrivateAtlaseResultOutputWithContext(ctx context.Context) LookupPrivateAtlaseResultOutput {
	return o
}

func (o LookupPrivateAtlaseResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateAtlaseResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupPrivateAtlaseResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateAtlaseResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupPrivateAtlaseResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateAtlaseResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupPrivateAtlaseResultOutput) Properties() PrivateAtlasPropertiesResponseOutput {
	return o.ApplyT(func(v LookupPrivateAtlaseResult) PrivateAtlasPropertiesResponse { return v.Properties }).(PrivateAtlasPropertiesResponseOutput)
}

func (o LookupPrivateAtlaseResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupPrivateAtlaseResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupPrivateAtlaseResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateAtlaseResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupPrivateAtlaseResultOutput{})
}
