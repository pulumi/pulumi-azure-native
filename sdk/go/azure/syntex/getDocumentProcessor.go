


package syntex

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupDocumentProcessor(ctx *pulumi.Context, args *LookupDocumentProcessorArgs, opts ...pulumi.InvokeOption) (*LookupDocumentProcessorResult, error) {
	var rv LookupDocumentProcessorResult
	err := ctx.Invoke("azure-native:syntex:getDocumentProcessor", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDocumentProcessorArgs struct {
	ProcessorName     string `pulumi:"processorName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupDocumentProcessorResult struct {
	Id         string                              `pulumi:"id"`
	Location   string                              `pulumi:"location"`
	Name       string                              `pulumi:"name"`
	Properties DocumentProcessorPropertiesResponse `pulumi:"properties"`
	SystemData SystemDataResponse                  `pulumi:"systemData"`
	Tags       map[string]string                   `pulumi:"tags"`
	Type       string                              `pulumi:"type"`
}

func LookupDocumentProcessorOutput(ctx *pulumi.Context, args LookupDocumentProcessorOutputArgs, opts ...pulumi.InvokeOption) LookupDocumentProcessorResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDocumentProcessorResult, error) {
			args := v.(LookupDocumentProcessorArgs)
			r, err := LookupDocumentProcessor(ctx, &args, opts...)
			var s LookupDocumentProcessorResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupDocumentProcessorResultOutput)
}

type LookupDocumentProcessorOutputArgs struct {
	ProcessorName     pulumi.StringInput `pulumi:"processorName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupDocumentProcessorOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDocumentProcessorArgs)(nil)).Elem()
}


type LookupDocumentProcessorResultOutput struct{ *pulumi.OutputState }

func (LookupDocumentProcessorResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDocumentProcessorResult)(nil)).Elem()
}

func (o LookupDocumentProcessorResultOutput) ToLookupDocumentProcessorResultOutput() LookupDocumentProcessorResultOutput {
	return o
}

func (o LookupDocumentProcessorResultOutput) ToLookupDocumentProcessorResultOutputWithContext(ctx context.Context) LookupDocumentProcessorResultOutput {
	return o
}

func (o LookupDocumentProcessorResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDocumentProcessorResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupDocumentProcessorResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDocumentProcessorResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupDocumentProcessorResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDocumentProcessorResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupDocumentProcessorResultOutput) Properties() DocumentProcessorPropertiesResponseOutput {
	return o.ApplyT(func(v LookupDocumentProcessorResult) DocumentProcessorPropertiesResponse { return v.Properties }).(DocumentProcessorPropertiesResponseOutput)
}

func (o LookupDocumentProcessorResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupDocumentProcessorResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupDocumentProcessorResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupDocumentProcessorResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupDocumentProcessorResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDocumentProcessorResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDocumentProcessorResultOutput{})
}
