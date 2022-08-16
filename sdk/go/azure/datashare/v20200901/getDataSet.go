


package v20200901

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


func LookupDataSet(ctx *pulumi.Context, args *LookupDataSetArgs, opts ...pulumi.InvokeOption) (*LookupDataSetResult, error) {
	var rv LookupDataSetResult
	err := ctx.Invoke("azure-native:datashare/v20200901:getDataSet", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDataSetArgs struct {
	AccountName       string `pulumi:"accountName"`
	DataSetName       string `pulumi:"dataSetName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ShareName         string `pulumi:"shareName"`
}


type LookupDataSetResult struct {
	Id         string             `pulumi:"id"`
	Kind       string             `pulumi:"kind"`
	Name       string             `pulumi:"name"`
	SystemData SystemDataResponse `pulumi:"systemData"`
	Type       string             `pulumi:"type"`
}

func LookupDataSetOutput(ctx *pulumi.Context, args LookupDataSetOutputArgs, opts ...pulumi.InvokeOption) LookupDataSetResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDataSetResult, error) {
			args := v.(LookupDataSetArgs)
			r, err := LookupDataSet(ctx, &args, opts...)
			var s LookupDataSetResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupDataSetResultOutput)
}

type LookupDataSetOutputArgs struct {
	AccountName       pulumi.StringInput `pulumi:"accountName"`
	DataSetName       pulumi.StringInput `pulumi:"dataSetName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	ShareName         pulumi.StringInput `pulumi:"shareName"`
}

func (LookupDataSetOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDataSetArgs)(nil)).Elem()
}


type LookupDataSetResultOutput struct{ *pulumi.OutputState }

func (LookupDataSetResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDataSetResult)(nil)).Elem()
}

func (o LookupDataSetResultOutput) ToLookupDataSetResultOutput() LookupDataSetResultOutput {
	return o
}

func (o LookupDataSetResultOutput) ToLookupDataSetResultOutputWithContext(ctx context.Context) LookupDataSetResultOutput {
	return o
}

func (o LookupDataSetResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDataSetResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupDataSetResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDataSetResult) string { return v.Kind }).(pulumi.StringOutput)
}

func (o LookupDataSetResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDataSetResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupDataSetResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupDataSetResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupDataSetResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDataSetResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDataSetResultOutput{})
}
