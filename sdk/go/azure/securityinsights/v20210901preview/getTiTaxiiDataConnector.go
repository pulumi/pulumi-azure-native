


package v20210901preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupTiTaxiiDataConnector(ctx *pulumi.Context, args *LookupTiTaxiiDataConnectorArgs, opts ...pulumi.InvokeOption) (*LookupTiTaxiiDataConnectorResult, error) {
	var rv LookupTiTaxiiDataConnectorResult
	err := ctx.Invoke("azure-native:securityinsights/v20210901preview:getTiTaxiiDataConnector", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupTiTaxiiDataConnectorArgs struct {
	DataConnectorId   string `pulumi:"dataConnectorId"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	WorkspaceName     string `pulumi:"workspaceName"`
}


type LookupTiTaxiiDataConnectorResult struct {
	CollectionId        *string                               `pulumi:"collectionId"`
	DataTypes           TiTaxiiDataConnectorDataTypesResponse `pulumi:"dataTypes"`
	Etag                *string                               `pulumi:"etag"`
	FriendlyName        *string                               `pulumi:"friendlyName"`
	Id                  string                                `pulumi:"id"`
	Kind                string                                `pulumi:"kind"`
	Name                string                                `pulumi:"name"`
	Password            *string                               `pulumi:"password"`
	PollingFrequency    string                                `pulumi:"pollingFrequency"`
	SystemData          SystemDataResponse                    `pulumi:"systemData"`
	TaxiiLookbackPeriod *string                               `pulumi:"taxiiLookbackPeriod"`
	TaxiiServer         *string                               `pulumi:"taxiiServer"`
	TenantId            string                                `pulumi:"tenantId"`
	Type                string                                `pulumi:"type"`
	UserName            *string                               `pulumi:"userName"`
	WorkspaceId         *string                               `pulumi:"workspaceId"`
}

func LookupTiTaxiiDataConnectorOutput(ctx *pulumi.Context, args LookupTiTaxiiDataConnectorOutputArgs, opts ...pulumi.InvokeOption) LookupTiTaxiiDataConnectorResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupTiTaxiiDataConnectorResult, error) {
			args := v.(LookupTiTaxiiDataConnectorArgs)
			r, err := LookupTiTaxiiDataConnector(ctx, &args, opts...)
			var s LookupTiTaxiiDataConnectorResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupTiTaxiiDataConnectorResultOutput)
}

type LookupTiTaxiiDataConnectorOutputArgs struct {
	DataConnectorId   pulumi.StringInput `pulumi:"dataConnectorId"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	WorkspaceName     pulumi.StringInput `pulumi:"workspaceName"`
}

func (LookupTiTaxiiDataConnectorOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTiTaxiiDataConnectorArgs)(nil)).Elem()
}


type LookupTiTaxiiDataConnectorResultOutput struct{ *pulumi.OutputState }

func (LookupTiTaxiiDataConnectorResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTiTaxiiDataConnectorResult)(nil)).Elem()
}

func (o LookupTiTaxiiDataConnectorResultOutput) ToLookupTiTaxiiDataConnectorResultOutput() LookupTiTaxiiDataConnectorResultOutput {
	return o
}

func (o LookupTiTaxiiDataConnectorResultOutput) ToLookupTiTaxiiDataConnectorResultOutputWithContext(ctx context.Context) LookupTiTaxiiDataConnectorResultOutput {
	return o
}

func (o LookupTiTaxiiDataConnectorResultOutput) CollectionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupTiTaxiiDataConnectorResult) *string { return v.CollectionId }).(pulumi.StringPtrOutput)
}

func (o LookupTiTaxiiDataConnectorResultOutput) DataTypes() TiTaxiiDataConnectorDataTypesResponseOutput {
	return o.ApplyT(func(v LookupTiTaxiiDataConnectorResult) TiTaxiiDataConnectorDataTypesResponse { return v.DataTypes }).(TiTaxiiDataConnectorDataTypesResponseOutput)
}

func (o LookupTiTaxiiDataConnectorResultOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupTiTaxiiDataConnectorResult) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o LookupTiTaxiiDataConnectorResultOutput) FriendlyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupTiTaxiiDataConnectorResult) *string { return v.FriendlyName }).(pulumi.StringPtrOutput)
}

func (o LookupTiTaxiiDataConnectorResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTiTaxiiDataConnectorResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupTiTaxiiDataConnectorResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTiTaxiiDataConnectorResult) string { return v.Kind }).(pulumi.StringOutput)
}

func (o LookupTiTaxiiDataConnectorResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTiTaxiiDataConnectorResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupTiTaxiiDataConnectorResultOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupTiTaxiiDataConnectorResult) *string { return v.Password }).(pulumi.StringPtrOutput)
}

func (o LookupTiTaxiiDataConnectorResultOutput) PollingFrequency() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTiTaxiiDataConnectorResult) string { return v.PollingFrequency }).(pulumi.StringOutput)
}

func (o LookupTiTaxiiDataConnectorResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupTiTaxiiDataConnectorResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupTiTaxiiDataConnectorResultOutput) TaxiiLookbackPeriod() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupTiTaxiiDataConnectorResult) *string { return v.TaxiiLookbackPeriod }).(pulumi.StringPtrOutput)
}

func (o LookupTiTaxiiDataConnectorResultOutput) TaxiiServer() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupTiTaxiiDataConnectorResult) *string { return v.TaxiiServer }).(pulumi.StringPtrOutput)
}

func (o LookupTiTaxiiDataConnectorResultOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTiTaxiiDataConnectorResult) string { return v.TenantId }).(pulumi.StringOutput)
}

func (o LookupTiTaxiiDataConnectorResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTiTaxiiDataConnectorResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupTiTaxiiDataConnectorResultOutput) UserName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupTiTaxiiDataConnectorResult) *string { return v.UserName }).(pulumi.StringPtrOutput)
}

func (o LookupTiTaxiiDataConnectorResultOutput) WorkspaceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupTiTaxiiDataConnectorResult) *string { return v.WorkspaceId }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupTiTaxiiDataConnectorResultOutput{})
}
