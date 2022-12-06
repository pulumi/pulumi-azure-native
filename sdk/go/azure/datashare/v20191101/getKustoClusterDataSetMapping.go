


package v20191101

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupKustoClusterDataSetMapping(ctx *pulumi.Context, args *LookupKustoClusterDataSetMappingArgs, opts ...pulumi.InvokeOption) (*LookupKustoClusterDataSetMappingResult, error) {
	var rv LookupKustoClusterDataSetMappingResult
	err := ctx.Invoke("azure-native:datashare/v20191101:getKustoClusterDataSetMapping", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupKustoClusterDataSetMappingArgs struct {
	AccountName           string `pulumi:"accountName"`
	DataSetMappingName    string `pulumi:"dataSetMappingName"`
	ResourceGroupName     string `pulumi:"resourceGroupName"`
	ShareSubscriptionName string `pulumi:"shareSubscriptionName"`
}


type LookupKustoClusterDataSetMappingResult struct {
	DataSetId              string `pulumi:"dataSetId"`
	DataSetMappingStatus   string `pulumi:"dataSetMappingStatus"`
	Id                     string `pulumi:"id"`
	Kind                   string `pulumi:"kind"`
	KustoClusterResourceId string `pulumi:"kustoClusterResourceId"`
	Location               string `pulumi:"location"`
	Name                   string `pulumi:"name"`
	ProvisioningState      string `pulumi:"provisioningState"`
	Type                   string `pulumi:"type"`
}

func LookupKustoClusterDataSetMappingOutput(ctx *pulumi.Context, args LookupKustoClusterDataSetMappingOutputArgs, opts ...pulumi.InvokeOption) LookupKustoClusterDataSetMappingResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupKustoClusterDataSetMappingResult, error) {
			args := v.(LookupKustoClusterDataSetMappingArgs)
			r, err := LookupKustoClusterDataSetMapping(ctx, &args, opts...)
			var s LookupKustoClusterDataSetMappingResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupKustoClusterDataSetMappingResultOutput)
}

type LookupKustoClusterDataSetMappingOutputArgs struct {
	AccountName           pulumi.StringInput `pulumi:"accountName"`
	DataSetMappingName    pulumi.StringInput `pulumi:"dataSetMappingName"`
	ResourceGroupName     pulumi.StringInput `pulumi:"resourceGroupName"`
	ShareSubscriptionName pulumi.StringInput `pulumi:"shareSubscriptionName"`
}

func (LookupKustoClusterDataSetMappingOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupKustoClusterDataSetMappingArgs)(nil)).Elem()
}


type LookupKustoClusterDataSetMappingResultOutput struct{ *pulumi.OutputState }

func (LookupKustoClusterDataSetMappingResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupKustoClusterDataSetMappingResult)(nil)).Elem()
}

func (o LookupKustoClusterDataSetMappingResultOutput) ToLookupKustoClusterDataSetMappingResultOutput() LookupKustoClusterDataSetMappingResultOutput {
	return o
}

func (o LookupKustoClusterDataSetMappingResultOutput) ToLookupKustoClusterDataSetMappingResultOutputWithContext(ctx context.Context) LookupKustoClusterDataSetMappingResultOutput {
	return o
}

func (o LookupKustoClusterDataSetMappingResultOutput) DataSetId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKustoClusterDataSetMappingResult) string { return v.DataSetId }).(pulumi.StringOutput)
}

func (o LookupKustoClusterDataSetMappingResultOutput) DataSetMappingStatus() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKustoClusterDataSetMappingResult) string { return v.DataSetMappingStatus }).(pulumi.StringOutput)
}

func (o LookupKustoClusterDataSetMappingResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKustoClusterDataSetMappingResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupKustoClusterDataSetMappingResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKustoClusterDataSetMappingResult) string { return v.Kind }).(pulumi.StringOutput)
}

func (o LookupKustoClusterDataSetMappingResultOutput) KustoClusterResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKustoClusterDataSetMappingResult) string { return v.KustoClusterResourceId }).(pulumi.StringOutput)
}

func (o LookupKustoClusterDataSetMappingResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKustoClusterDataSetMappingResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupKustoClusterDataSetMappingResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKustoClusterDataSetMappingResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupKustoClusterDataSetMappingResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKustoClusterDataSetMappingResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupKustoClusterDataSetMappingResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKustoClusterDataSetMappingResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupKustoClusterDataSetMappingResultOutput{})
}
