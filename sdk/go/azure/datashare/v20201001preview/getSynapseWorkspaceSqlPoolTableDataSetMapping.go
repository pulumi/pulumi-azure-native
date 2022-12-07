


package v20201001preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupSynapseWorkspaceSqlPoolTableDataSetMapping(ctx *pulumi.Context, args *LookupSynapseWorkspaceSqlPoolTableDataSetMappingArgs, opts ...pulumi.InvokeOption) (*LookupSynapseWorkspaceSqlPoolTableDataSetMappingResult, error) {
	var rv LookupSynapseWorkspaceSqlPoolTableDataSetMappingResult
	err := ctx.Invoke("azure-native:datashare/v20201001preview:getSynapseWorkspaceSqlPoolTableDataSetMapping", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSynapseWorkspaceSqlPoolTableDataSetMappingArgs struct {
	AccountName           string `pulumi:"accountName"`
	DataSetMappingName    string `pulumi:"dataSetMappingName"`
	ResourceGroupName     string `pulumi:"resourceGroupName"`
	ShareSubscriptionName string `pulumi:"shareSubscriptionName"`
}


type LookupSynapseWorkspaceSqlPoolTableDataSetMappingResult struct {
	DataSetId                              string             `pulumi:"dataSetId"`
	DataSetMappingStatus                   string             `pulumi:"dataSetMappingStatus"`
	Id                                     string             `pulumi:"id"`
	Kind                                   string             `pulumi:"kind"`
	Name                                   string             `pulumi:"name"`
	ProvisioningState                      string             `pulumi:"provisioningState"`
	SynapseWorkspaceSqlPoolTableResourceId string             `pulumi:"synapseWorkspaceSqlPoolTableResourceId"`
	SystemData                             SystemDataResponse `pulumi:"systemData"`
	Type                                   string             `pulumi:"type"`
}

func LookupSynapseWorkspaceSqlPoolTableDataSetMappingOutput(ctx *pulumi.Context, args LookupSynapseWorkspaceSqlPoolTableDataSetMappingOutputArgs, opts ...pulumi.InvokeOption) LookupSynapseWorkspaceSqlPoolTableDataSetMappingResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSynapseWorkspaceSqlPoolTableDataSetMappingResult, error) {
			args := v.(LookupSynapseWorkspaceSqlPoolTableDataSetMappingArgs)
			r, err := LookupSynapseWorkspaceSqlPoolTableDataSetMapping(ctx, &args, opts...)
			var s LookupSynapseWorkspaceSqlPoolTableDataSetMappingResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSynapseWorkspaceSqlPoolTableDataSetMappingResultOutput)
}

type LookupSynapseWorkspaceSqlPoolTableDataSetMappingOutputArgs struct {
	AccountName           pulumi.StringInput `pulumi:"accountName"`
	DataSetMappingName    pulumi.StringInput `pulumi:"dataSetMappingName"`
	ResourceGroupName     pulumi.StringInput `pulumi:"resourceGroupName"`
	ShareSubscriptionName pulumi.StringInput `pulumi:"shareSubscriptionName"`
}

func (LookupSynapseWorkspaceSqlPoolTableDataSetMappingOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSynapseWorkspaceSqlPoolTableDataSetMappingArgs)(nil)).Elem()
}


type LookupSynapseWorkspaceSqlPoolTableDataSetMappingResultOutput struct{ *pulumi.OutputState }

func (LookupSynapseWorkspaceSqlPoolTableDataSetMappingResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSynapseWorkspaceSqlPoolTableDataSetMappingResult)(nil)).Elem()
}

func (o LookupSynapseWorkspaceSqlPoolTableDataSetMappingResultOutput) ToLookupSynapseWorkspaceSqlPoolTableDataSetMappingResultOutput() LookupSynapseWorkspaceSqlPoolTableDataSetMappingResultOutput {
	return o
}

func (o LookupSynapseWorkspaceSqlPoolTableDataSetMappingResultOutput) ToLookupSynapseWorkspaceSqlPoolTableDataSetMappingResultOutputWithContext(ctx context.Context) LookupSynapseWorkspaceSqlPoolTableDataSetMappingResultOutput {
	return o
}

func (o LookupSynapseWorkspaceSqlPoolTableDataSetMappingResultOutput) DataSetId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSynapseWorkspaceSqlPoolTableDataSetMappingResult) string { return v.DataSetId }).(pulumi.StringOutput)
}

func (o LookupSynapseWorkspaceSqlPoolTableDataSetMappingResultOutput) DataSetMappingStatus() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSynapseWorkspaceSqlPoolTableDataSetMappingResult) string { return v.DataSetMappingStatus }).(pulumi.StringOutput)
}

func (o LookupSynapseWorkspaceSqlPoolTableDataSetMappingResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSynapseWorkspaceSqlPoolTableDataSetMappingResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupSynapseWorkspaceSqlPoolTableDataSetMappingResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSynapseWorkspaceSqlPoolTableDataSetMappingResult) string { return v.Kind }).(pulumi.StringOutput)
}

func (o LookupSynapseWorkspaceSqlPoolTableDataSetMappingResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSynapseWorkspaceSqlPoolTableDataSetMappingResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupSynapseWorkspaceSqlPoolTableDataSetMappingResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSynapseWorkspaceSqlPoolTableDataSetMappingResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupSynapseWorkspaceSqlPoolTableDataSetMappingResultOutput) SynapseWorkspaceSqlPoolTableResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSynapseWorkspaceSqlPoolTableDataSetMappingResult) string {
		return v.SynapseWorkspaceSqlPoolTableResourceId
	}).(pulumi.StringOutput)
}

func (o LookupSynapseWorkspaceSqlPoolTableDataSetMappingResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupSynapseWorkspaceSqlPoolTableDataSetMappingResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupSynapseWorkspaceSqlPoolTableDataSetMappingResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSynapseWorkspaceSqlPoolTableDataSetMappingResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSynapseWorkspaceSqlPoolTableDataSetMappingResultOutput{})
}
