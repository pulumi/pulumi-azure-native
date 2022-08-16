


package v20210801

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupKustoTableDataSetMapping(ctx *pulumi.Context, args *LookupKustoTableDataSetMappingArgs, opts ...pulumi.InvokeOption) (*LookupKustoTableDataSetMappingResult, error) {
	var rv LookupKustoTableDataSetMappingResult
	err := ctx.Invoke("azure-native:datashare/v20210801:getKustoTableDataSetMapping", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupKustoTableDataSetMappingArgs struct {
	AccountName           string `pulumi:"accountName"`
	DataSetMappingName    string `pulumi:"dataSetMappingName"`
	ResourceGroupName     string `pulumi:"resourceGroupName"`
	ShareSubscriptionName string `pulumi:"shareSubscriptionName"`
}


type LookupKustoTableDataSetMappingResult struct {
	DataSetId              string             `pulumi:"dataSetId"`
	DataSetMappingStatus   string             `pulumi:"dataSetMappingStatus"`
	Id                     string             `pulumi:"id"`
	Kind                   string             `pulumi:"kind"`
	KustoClusterResourceId string             `pulumi:"kustoClusterResourceId"`
	Location               string             `pulumi:"location"`
	Name                   string             `pulumi:"name"`
	ProvisioningState      string             `pulumi:"provisioningState"`
	SystemData             SystemDataResponse `pulumi:"systemData"`
	Type                   string             `pulumi:"type"`
}

func LookupKustoTableDataSetMappingOutput(ctx *pulumi.Context, args LookupKustoTableDataSetMappingOutputArgs, opts ...pulumi.InvokeOption) LookupKustoTableDataSetMappingResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupKustoTableDataSetMappingResult, error) {
			args := v.(LookupKustoTableDataSetMappingArgs)
			r, err := LookupKustoTableDataSetMapping(ctx, &args, opts...)
			var s LookupKustoTableDataSetMappingResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupKustoTableDataSetMappingResultOutput)
}

type LookupKustoTableDataSetMappingOutputArgs struct {
	AccountName           pulumi.StringInput `pulumi:"accountName"`
	DataSetMappingName    pulumi.StringInput `pulumi:"dataSetMappingName"`
	ResourceGroupName     pulumi.StringInput `pulumi:"resourceGroupName"`
	ShareSubscriptionName pulumi.StringInput `pulumi:"shareSubscriptionName"`
}

func (LookupKustoTableDataSetMappingOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupKustoTableDataSetMappingArgs)(nil)).Elem()
}


type LookupKustoTableDataSetMappingResultOutput struct{ *pulumi.OutputState }

func (LookupKustoTableDataSetMappingResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupKustoTableDataSetMappingResult)(nil)).Elem()
}

func (o LookupKustoTableDataSetMappingResultOutput) ToLookupKustoTableDataSetMappingResultOutput() LookupKustoTableDataSetMappingResultOutput {
	return o
}

func (o LookupKustoTableDataSetMappingResultOutput) ToLookupKustoTableDataSetMappingResultOutputWithContext(ctx context.Context) LookupKustoTableDataSetMappingResultOutput {
	return o
}

func (o LookupKustoTableDataSetMappingResultOutput) DataSetId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKustoTableDataSetMappingResult) string { return v.DataSetId }).(pulumi.StringOutput)
}

func (o LookupKustoTableDataSetMappingResultOutput) DataSetMappingStatus() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKustoTableDataSetMappingResult) string { return v.DataSetMappingStatus }).(pulumi.StringOutput)
}

func (o LookupKustoTableDataSetMappingResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKustoTableDataSetMappingResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupKustoTableDataSetMappingResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKustoTableDataSetMappingResult) string { return v.Kind }).(pulumi.StringOutput)
}

func (o LookupKustoTableDataSetMappingResultOutput) KustoClusterResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKustoTableDataSetMappingResult) string { return v.KustoClusterResourceId }).(pulumi.StringOutput)
}

func (o LookupKustoTableDataSetMappingResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKustoTableDataSetMappingResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupKustoTableDataSetMappingResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKustoTableDataSetMappingResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupKustoTableDataSetMappingResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKustoTableDataSetMappingResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupKustoTableDataSetMappingResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupKustoTableDataSetMappingResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupKustoTableDataSetMappingResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKustoTableDataSetMappingResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupKustoTableDataSetMappingResultOutput{})
}
