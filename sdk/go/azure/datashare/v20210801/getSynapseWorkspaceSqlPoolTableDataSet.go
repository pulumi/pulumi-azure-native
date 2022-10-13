


package v20210801

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupSynapseWorkspaceSqlPoolTableDataSet(ctx *pulumi.Context, args *LookupSynapseWorkspaceSqlPoolTableDataSetArgs, opts ...pulumi.InvokeOption) (*LookupSynapseWorkspaceSqlPoolTableDataSetResult, error) {
	var rv LookupSynapseWorkspaceSqlPoolTableDataSetResult
	err := ctx.Invoke("azure-native:datashare/v20210801:getSynapseWorkspaceSqlPoolTableDataSet", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSynapseWorkspaceSqlPoolTableDataSetArgs struct {
	AccountName       string `pulumi:"accountName"`
	DataSetName       string `pulumi:"dataSetName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ShareName         string `pulumi:"shareName"`
}


type LookupSynapseWorkspaceSqlPoolTableDataSetResult struct {
	DataSetId                              string             `pulumi:"dataSetId"`
	Id                                     string             `pulumi:"id"`
	Kind                                   string             `pulumi:"kind"`
	Name                                   string             `pulumi:"name"`
	SynapseWorkspaceSqlPoolTableResourceId string             `pulumi:"synapseWorkspaceSqlPoolTableResourceId"`
	SystemData                             SystemDataResponse `pulumi:"systemData"`
	Type                                   string             `pulumi:"type"`
}

func LookupSynapseWorkspaceSqlPoolTableDataSetOutput(ctx *pulumi.Context, args LookupSynapseWorkspaceSqlPoolTableDataSetOutputArgs, opts ...pulumi.InvokeOption) LookupSynapseWorkspaceSqlPoolTableDataSetResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSynapseWorkspaceSqlPoolTableDataSetResult, error) {
			args := v.(LookupSynapseWorkspaceSqlPoolTableDataSetArgs)
			r, err := LookupSynapseWorkspaceSqlPoolTableDataSet(ctx, &args, opts...)
			var s LookupSynapseWorkspaceSqlPoolTableDataSetResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSynapseWorkspaceSqlPoolTableDataSetResultOutput)
}

type LookupSynapseWorkspaceSqlPoolTableDataSetOutputArgs struct {
	AccountName       pulumi.StringInput `pulumi:"accountName"`
	DataSetName       pulumi.StringInput `pulumi:"dataSetName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	ShareName         pulumi.StringInput `pulumi:"shareName"`
}

func (LookupSynapseWorkspaceSqlPoolTableDataSetOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSynapseWorkspaceSqlPoolTableDataSetArgs)(nil)).Elem()
}


type LookupSynapseWorkspaceSqlPoolTableDataSetResultOutput struct{ *pulumi.OutputState }

func (LookupSynapseWorkspaceSqlPoolTableDataSetResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSynapseWorkspaceSqlPoolTableDataSetResult)(nil)).Elem()
}

func (o LookupSynapseWorkspaceSqlPoolTableDataSetResultOutput) ToLookupSynapseWorkspaceSqlPoolTableDataSetResultOutput() LookupSynapseWorkspaceSqlPoolTableDataSetResultOutput {
	return o
}

func (o LookupSynapseWorkspaceSqlPoolTableDataSetResultOutput) ToLookupSynapseWorkspaceSqlPoolTableDataSetResultOutputWithContext(ctx context.Context) LookupSynapseWorkspaceSqlPoolTableDataSetResultOutput {
	return o
}

func (o LookupSynapseWorkspaceSqlPoolTableDataSetResultOutput) DataSetId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSynapseWorkspaceSqlPoolTableDataSetResult) string { return v.DataSetId }).(pulumi.StringOutput)
}

func (o LookupSynapseWorkspaceSqlPoolTableDataSetResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSynapseWorkspaceSqlPoolTableDataSetResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupSynapseWorkspaceSqlPoolTableDataSetResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSynapseWorkspaceSqlPoolTableDataSetResult) string { return v.Kind }).(pulumi.StringOutput)
}

func (o LookupSynapseWorkspaceSqlPoolTableDataSetResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSynapseWorkspaceSqlPoolTableDataSetResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupSynapseWorkspaceSqlPoolTableDataSetResultOutput) SynapseWorkspaceSqlPoolTableResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSynapseWorkspaceSqlPoolTableDataSetResult) string {
		return v.SynapseWorkspaceSqlPoolTableResourceId
	}).(pulumi.StringOutput)
}

func (o LookupSynapseWorkspaceSqlPoolTableDataSetResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupSynapseWorkspaceSqlPoolTableDataSetResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupSynapseWorkspaceSqlPoolTableDataSetResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSynapseWorkspaceSqlPoolTableDataSetResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSynapseWorkspaceSqlPoolTableDataSetResultOutput{})
}
