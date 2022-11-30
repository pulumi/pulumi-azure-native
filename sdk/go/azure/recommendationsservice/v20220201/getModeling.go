


package v20220201

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupModeling(ctx *pulumi.Context, args *LookupModelingArgs, opts ...pulumi.InvokeOption) (*LookupModelingResult, error) {
	var rv LookupModelingResult
	err := ctx.Invoke("azure-native:recommendationsservice/v20220201:getModeling", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupModelingArgs struct {
	AccountName       string `pulumi:"accountName"`
	ModelingName      string `pulumi:"modelingName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupModelingResult struct {
	Id         string                             `pulumi:"id"`
	Location   string                             `pulumi:"location"`
	Name       string                             `pulumi:"name"`
	Properties ModelingResourceResponseProperties `pulumi:"properties"`
	SystemData SystemDataResponse                 `pulumi:"systemData"`
	Tags       map[string]string                  `pulumi:"tags"`
	Type       string                             `pulumi:"type"`
}

func LookupModelingOutput(ctx *pulumi.Context, args LookupModelingOutputArgs, opts ...pulumi.InvokeOption) LookupModelingResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupModelingResult, error) {
			args := v.(LookupModelingArgs)
			r, err := LookupModeling(ctx, &args, opts...)
			var s LookupModelingResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupModelingResultOutput)
}

type LookupModelingOutputArgs struct {
	AccountName       pulumi.StringInput `pulumi:"accountName"`
	ModelingName      pulumi.StringInput `pulumi:"modelingName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupModelingOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupModelingArgs)(nil)).Elem()
}


type LookupModelingResultOutput struct{ *pulumi.OutputState }

func (LookupModelingResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupModelingResult)(nil)).Elem()
}

func (o LookupModelingResultOutput) ToLookupModelingResultOutput() LookupModelingResultOutput {
	return o
}

func (o LookupModelingResultOutput) ToLookupModelingResultOutputWithContext(ctx context.Context) LookupModelingResultOutput {
	return o
}

func (o LookupModelingResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupModelingResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupModelingResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupModelingResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupModelingResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupModelingResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupModelingResultOutput) Properties() ModelingResourceResponsePropertiesOutput {
	return o.ApplyT(func(v LookupModelingResult) ModelingResourceResponseProperties { return v.Properties }).(ModelingResourceResponsePropertiesOutput)
}

func (o LookupModelingResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupModelingResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupModelingResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupModelingResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupModelingResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupModelingResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupModelingResultOutput{})
}
