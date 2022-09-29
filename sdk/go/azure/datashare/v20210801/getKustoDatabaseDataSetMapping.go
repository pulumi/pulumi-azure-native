


package v20210801

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupKustoDatabaseDataSetMapping(ctx *pulumi.Context, args *LookupKustoDatabaseDataSetMappingArgs, opts ...pulumi.InvokeOption) (*LookupKustoDatabaseDataSetMappingResult, error) {
	var rv LookupKustoDatabaseDataSetMappingResult
	err := ctx.Invoke("azure-native:datashare/v20210801:getKustoDatabaseDataSetMapping", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupKustoDatabaseDataSetMappingArgs struct {
	AccountName           string `pulumi:"accountName"`
	DataSetMappingName    string `pulumi:"dataSetMappingName"`
	ResourceGroupName     string `pulumi:"resourceGroupName"`
	ShareSubscriptionName string `pulumi:"shareSubscriptionName"`
}


type LookupKustoDatabaseDataSetMappingResult struct {
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

func LookupKustoDatabaseDataSetMappingOutput(ctx *pulumi.Context, args LookupKustoDatabaseDataSetMappingOutputArgs, opts ...pulumi.InvokeOption) LookupKustoDatabaseDataSetMappingResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupKustoDatabaseDataSetMappingResult, error) {
			args := v.(LookupKustoDatabaseDataSetMappingArgs)
			r, err := LookupKustoDatabaseDataSetMapping(ctx, &args, opts...)
			var s LookupKustoDatabaseDataSetMappingResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupKustoDatabaseDataSetMappingResultOutput)
}

type LookupKustoDatabaseDataSetMappingOutputArgs struct {
	AccountName           pulumi.StringInput `pulumi:"accountName"`
	DataSetMappingName    pulumi.StringInput `pulumi:"dataSetMappingName"`
	ResourceGroupName     pulumi.StringInput `pulumi:"resourceGroupName"`
	ShareSubscriptionName pulumi.StringInput `pulumi:"shareSubscriptionName"`
}

func (LookupKustoDatabaseDataSetMappingOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupKustoDatabaseDataSetMappingArgs)(nil)).Elem()
}


type LookupKustoDatabaseDataSetMappingResultOutput struct{ *pulumi.OutputState }

func (LookupKustoDatabaseDataSetMappingResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupKustoDatabaseDataSetMappingResult)(nil)).Elem()
}

func (o LookupKustoDatabaseDataSetMappingResultOutput) ToLookupKustoDatabaseDataSetMappingResultOutput() LookupKustoDatabaseDataSetMappingResultOutput {
	return o
}

func (o LookupKustoDatabaseDataSetMappingResultOutput) ToLookupKustoDatabaseDataSetMappingResultOutputWithContext(ctx context.Context) LookupKustoDatabaseDataSetMappingResultOutput {
	return o
}

func (o LookupKustoDatabaseDataSetMappingResultOutput) DataSetId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKustoDatabaseDataSetMappingResult) string { return v.DataSetId }).(pulumi.StringOutput)
}

func (o LookupKustoDatabaseDataSetMappingResultOutput) DataSetMappingStatus() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKustoDatabaseDataSetMappingResult) string { return v.DataSetMappingStatus }).(pulumi.StringOutput)
}

func (o LookupKustoDatabaseDataSetMappingResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKustoDatabaseDataSetMappingResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupKustoDatabaseDataSetMappingResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKustoDatabaseDataSetMappingResult) string { return v.Kind }).(pulumi.StringOutput)
}

func (o LookupKustoDatabaseDataSetMappingResultOutput) KustoClusterResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKustoDatabaseDataSetMappingResult) string { return v.KustoClusterResourceId }).(pulumi.StringOutput)
}

func (o LookupKustoDatabaseDataSetMappingResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKustoDatabaseDataSetMappingResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupKustoDatabaseDataSetMappingResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKustoDatabaseDataSetMappingResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupKustoDatabaseDataSetMappingResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKustoDatabaseDataSetMappingResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupKustoDatabaseDataSetMappingResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupKustoDatabaseDataSetMappingResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupKustoDatabaseDataSetMappingResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKustoDatabaseDataSetMappingResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupKustoDatabaseDataSetMappingResultOutput{})
}
