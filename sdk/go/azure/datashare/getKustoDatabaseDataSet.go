


package datashare

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupKustoDatabaseDataSet(ctx *pulumi.Context, args *LookupKustoDatabaseDataSetArgs, opts ...pulumi.InvokeOption) (*LookupKustoDatabaseDataSetResult, error) {
	var rv LookupKustoDatabaseDataSetResult
	err := ctx.Invoke("azure-native:datashare:getKustoDatabaseDataSet", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupKustoDatabaseDataSetArgs struct {
	AccountName       string `pulumi:"accountName"`
	DataSetName       string `pulumi:"dataSetName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ShareName         string `pulumi:"shareName"`
}


type LookupKustoDatabaseDataSetResult struct {
	DataSetId               string             `pulumi:"dataSetId"`
	Id                      string             `pulumi:"id"`
	Kind                    string             `pulumi:"kind"`
	KustoDatabaseResourceId string             `pulumi:"kustoDatabaseResourceId"`
	Location                string             `pulumi:"location"`
	Name                    string             `pulumi:"name"`
	ProvisioningState       string             `pulumi:"provisioningState"`
	SystemData              SystemDataResponse `pulumi:"systemData"`
	Type                    string             `pulumi:"type"`
}

func LookupKustoDatabaseDataSetOutput(ctx *pulumi.Context, args LookupKustoDatabaseDataSetOutputArgs, opts ...pulumi.InvokeOption) LookupKustoDatabaseDataSetResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupKustoDatabaseDataSetResult, error) {
			args := v.(LookupKustoDatabaseDataSetArgs)
			r, err := LookupKustoDatabaseDataSet(ctx, &args, opts...)
			var s LookupKustoDatabaseDataSetResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupKustoDatabaseDataSetResultOutput)
}

type LookupKustoDatabaseDataSetOutputArgs struct {
	AccountName       pulumi.StringInput `pulumi:"accountName"`
	DataSetName       pulumi.StringInput `pulumi:"dataSetName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	ShareName         pulumi.StringInput `pulumi:"shareName"`
}

func (LookupKustoDatabaseDataSetOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupKustoDatabaseDataSetArgs)(nil)).Elem()
}


type LookupKustoDatabaseDataSetResultOutput struct{ *pulumi.OutputState }

func (LookupKustoDatabaseDataSetResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupKustoDatabaseDataSetResult)(nil)).Elem()
}

func (o LookupKustoDatabaseDataSetResultOutput) ToLookupKustoDatabaseDataSetResultOutput() LookupKustoDatabaseDataSetResultOutput {
	return o
}

func (o LookupKustoDatabaseDataSetResultOutput) ToLookupKustoDatabaseDataSetResultOutputWithContext(ctx context.Context) LookupKustoDatabaseDataSetResultOutput {
	return o
}

func (o LookupKustoDatabaseDataSetResultOutput) DataSetId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKustoDatabaseDataSetResult) string { return v.DataSetId }).(pulumi.StringOutput)
}

func (o LookupKustoDatabaseDataSetResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKustoDatabaseDataSetResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupKustoDatabaseDataSetResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKustoDatabaseDataSetResult) string { return v.Kind }).(pulumi.StringOutput)
}

func (o LookupKustoDatabaseDataSetResultOutput) KustoDatabaseResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKustoDatabaseDataSetResult) string { return v.KustoDatabaseResourceId }).(pulumi.StringOutput)
}

func (o LookupKustoDatabaseDataSetResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKustoDatabaseDataSetResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupKustoDatabaseDataSetResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKustoDatabaseDataSetResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupKustoDatabaseDataSetResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKustoDatabaseDataSetResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupKustoDatabaseDataSetResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupKustoDatabaseDataSetResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupKustoDatabaseDataSetResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKustoDatabaseDataSetResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupKustoDatabaseDataSetResultOutput{})
}
