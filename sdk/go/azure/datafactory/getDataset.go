


package datafactory

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupDataset(ctx *pulumi.Context, args *LookupDatasetArgs, opts ...pulumi.InvokeOption) (*LookupDatasetResult, error) {
	var rv LookupDatasetResult
	err := ctx.Invoke("azure-native:datafactory:getDataset", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDatasetArgs struct {
	DatasetName       string `pulumi:"datasetName"`
	FactoryName       string `pulumi:"factoryName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupDatasetResult struct {
	Etag       string      `pulumi:"etag"`
	Id         string      `pulumi:"id"`
	Name       string      `pulumi:"name"`
	Properties interface{} `pulumi:"properties"`
	Type       string      `pulumi:"type"`
}

func LookupDatasetOutput(ctx *pulumi.Context, args LookupDatasetOutputArgs, opts ...pulumi.InvokeOption) LookupDatasetResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDatasetResult, error) {
			args := v.(LookupDatasetArgs)
			r, err := LookupDataset(ctx, &args, opts...)
			var s LookupDatasetResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupDatasetResultOutput)
}

type LookupDatasetOutputArgs struct {
	DatasetName       pulumi.StringInput `pulumi:"datasetName"`
	FactoryName       pulumi.StringInput `pulumi:"factoryName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupDatasetOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDatasetArgs)(nil)).Elem()
}


type LookupDatasetResultOutput struct{ *pulumi.OutputState }

func (LookupDatasetResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDatasetResult)(nil)).Elem()
}

func (o LookupDatasetResultOutput) ToLookupDatasetResultOutput() LookupDatasetResultOutput {
	return o
}

func (o LookupDatasetResultOutput) ToLookupDatasetResultOutputWithContext(ctx context.Context) LookupDatasetResultOutput {
	return o
}

func (o LookupDatasetResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatasetResult) string { return v.Etag }).(pulumi.StringOutput)
}

func (o LookupDatasetResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatasetResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupDatasetResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatasetResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupDatasetResultOutput) Properties() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupDatasetResult) interface{} { return v.Properties }).(pulumi.AnyOutput)
}

func (o LookupDatasetResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatasetResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDatasetResultOutput{})
}
