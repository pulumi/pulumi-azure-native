


package v20210801

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupKustoTableDataSet(ctx *pulumi.Context, args *LookupKustoTableDataSetArgs, opts ...pulumi.InvokeOption) (*LookupKustoTableDataSetResult, error) {
	var rv LookupKustoTableDataSetResult
	err := ctx.Invoke("azure-native:datashare/v20210801:getKustoTableDataSet", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupKustoTableDataSetArgs struct {
	AccountName       string `pulumi:"accountName"`
	DataSetName       string `pulumi:"dataSetName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ShareName         string `pulumi:"shareName"`
}


type LookupKustoTableDataSetResult struct {
	DataSetId                   string                              `pulumi:"dataSetId"`
	Id                          string                              `pulumi:"id"`
	Kind                        string                              `pulumi:"kind"`
	KustoDatabaseResourceId     string                              `pulumi:"kustoDatabaseResourceId"`
	Location                    string                              `pulumi:"location"`
	Name                        string                              `pulumi:"name"`
	ProvisioningState           string                              `pulumi:"provisioningState"`
	SystemData                  SystemDataResponse                  `pulumi:"systemData"`
	TableLevelSharingProperties TableLevelSharingPropertiesResponse `pulumi:"tableLevelSharingProperties"`
	Type                        string                              `pulumi:"type"`
}

func LookupKustoTableDataSetOutput(ctx *pulumi.Context, args LookupKustoTableDataSetOutputArgs, opts ...pulumi.InvokeOption) LookupKustoTableDataSetResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupKustoTableDataSetResult, error) {
			args := v.(LookupKustoTableDataSetArgs)
			r, err := LookupKustoTableDataSet(ctx, &args, opts...)
			var s LookupKustoTableDataSetResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupKustoTableDataSetResultOutput)
}

type LookupKustoTableDataSetOutputArgs struct {
	AccountName       pulumi.StringInput `pulumi:"accountName"`
	DataSetName       pulumi.StringInput `pulumi:"dataSetName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	ShareName         pulumi.StringInput `pulumi:"shareName"`
}

func (LookupKustoTableDataSetOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupKustoTableDataSetArgs)(nil)).Elem()
}


type LookupKustoTableDataSetResultOutput struct{ *pulumi.OutputState }

func (LookupKustoTableDataSetResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupKustoTableDataSetResult)(nil)).Elem()
}

func (o LookupKustoTableDataSetResultOutput) ToLookupKustoTableDataSetResultOutput() LookupKustoTableDataSetResultOutput {
	return o
}

func (o LookupKustoTableDataSetResultOutput) ToLookupKustoTableDataSetResultOutputWithContext(ctx context.Context) LookupKustoTableDataSetResultOutput {
	return o
}

func (o LookupKustoTableDataSetResultOutput) DataSetId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKustoTableDataSetResult) string { return v.DataSetId }).(pulumi.StringOutput)
}

func (o LookupKustoTableDataSetResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKustoTableDataSetResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupKustoTableDataSetResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKustoTableDataSetResult) string { return v.Kind }).(pulumi.StringOutput)
}

func (o LookupKustoTableDataSetResultOutput) KustoDatabaseResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKustoTableDataSetResult) string { return v.KustoDatabaseResourceId }).(pulumi.StringOutput)
}

func (o LookupKustoTableDataSetResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKustoTableDataSetResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupKustoTableDataSetResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKustoTableDataSetResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupKustoTableDataSetResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKustoTableDataSetResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupKustoTableDataSetResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupKustoTableDataSetResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupKustoTableDataSetResultOutput) TableLevelSharingProperties() TableLevelSharingPropertiesResponseOutput {
	return o.ApplyT(func(v LookupKustoTableDataSetResult) TableLevelSharingPropertiesResponse {
		return v.TableLevelSharingProperties
	}).(TableLevelSharingPropertiesResponseOutput)
}

func (o LookupKustoTableDataSetResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKustoTableDataSetResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupKustoTableDataSetResultOutput{})
}
