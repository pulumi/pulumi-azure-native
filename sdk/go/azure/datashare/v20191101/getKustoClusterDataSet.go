


package v20191101

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupKustoClusterDataSet(ctx *pulumi.Context, args *LookupKustoClusterDataSetArgs, opts ...pulumi.InvokeOption) (*LookupKustoClusterDataSetResult, error) {
	var rv LookupKustoClusterDataSetResult
	err := ctx.Invoke("azure-native:datashare/v20191101:getKustoClusterDataSet", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupKustoClusterDataSetArgs struct {
	AccountName       string `pulumi:"accountName"`
	DataSetName       string `pulumi:"dataSetName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ShareName         string `pulumi:"shareName"`
}


type LookupKustoClusterDataSetResult struct {
	DataSetId              string `pulumi:"dataSetId"`
	Id                     string `pulumi:"id"`
	Kind                   string `pulumi:"kind"`
	KustoClusterResourceId string `pulumi:"kustoClusterResourceId"`
	Location               string `pulumi:"location"`
	Name                   string `pulumi:"name"`
	ProvisioningState      string `pulumi:"provisioningState"`
	Type                   string `pulumi:"type"`
}

func LookupKustoClusterDataSetOutput(ctx *pulumi.Context, args LookupKustoClusterDataSetOutputArgs, opts ...pulumi.InvokeOption) LookupKustoClusterDataSetResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupKustoClusterDataSetResult, error) {
			args := v.(LookupKustoClusterDataSetArgs)
			r, err := LookupKustoClusterDataSet(ctx, &args, opts...)
			var s LookupKustoClusterDataSetResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupKustoClusterDataSetResultOutput)
}

type LookupKustoClusterDataSetOutputArgs struct {
	AccountName       pulumi.StringInput `pulumi:"accountName"`
	DataSetName       pulumi.StringInput `pulumi:"dataSetName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	ShareName         pulumi.StringInput `pulumi:"shareName"`
}

func (LookupKustoClusterDataSetOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupKustoClusterDataSetArgs)(nil)).Elem()
}


type LookupKustoClusterDataSetResultOutput struct{ *pulumi.OutputState }

func (LookupKustoClusterDataSetResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupKustoClusterDataSetResult)(nil)).Elem()
}

func (o LookupKustoClusterDataSetResultOutput) ToLookupKustoClusterDataSetResultOutput() LookupKustoClusterDataSetResultOutput {
	return o
}

func (o LookupKustoClusterDataSetResultOutput) ToLookupKustoClusterDataSetResultOutputWithContext(ctx context.Context) LookupKustoClusterDataSetResultOutput {
	return o
}

func (o LookupKustoClusterDataSetResultOutput) DataSetId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKustoClusterDataSetResult) string { return v.DataSetId }).(pulumi.StringOutput)
}

func (o LookupKustoClusterDataSetResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKustoClusterDataSetResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupKustoClusterDataSetResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKustoClusterDataSetResult) string { return v.Kind }).(pulumi.StringOutput)
}

func (o LookupKustoClusterDataSetResultOutput) KustoClusterResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKustoClusterDataSetResult) string { return v.KustoClusterResourceId }).(pulumi.StringOutput)
}

func (o LookupKustoClusterDataSetResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKustoClusterDataSetResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupKustoClusterDataSetResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKustoClusterDataSetResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupKustoClusterDataSetResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKustoClusterDataSetResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupKustoClusterDataSetResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKustoClusterDataSetResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupKustoClusterDataSetResultOutput{})
}
