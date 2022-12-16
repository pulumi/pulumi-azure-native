


package v20221201preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupOfficePowerBIDataConnector(ctx *pulumi.Context, args *LookupOfficePowerBIDataConnectorArgs, opts ...pulumi.InvokeOption) (*LookupOfficePowerBIDataConnectorResult, error) {
	var rv LookupOfficePowerBIDataConnectorResult
	err := ctx.Invoke("azure-native:securityinsights/v20221201preview:getOfficePowerBIDataConnector", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupOfficePowerBIDataConnectorArgs struct {
	DataConnectorId   string `pulumi:"dataConnectorId"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	WorkspaceName     string `pulumi:"workspaceName"`
}


type LookupOfficePowerBIDataConnectorResult struct {
	DataTypes  OfficePowerBIConnectorDataTypesResponse `pulumi:"dataTypes"`
	Etag       *string                                 `pulumi:"etag"`
	Id         string                                  `pulumi:"id"`
	Kind       string                                  `pulumi:"kind"`
	Name       string                                  `pulumi:"name"`
	SystemData SystemDataResponse                      `pulumi:"systemData"`
	TenantId   string                                  `pulumi:"tenantId"`
	Type       string                                  `pulumi:"type"`
}

func LookupOfficePowerBIDataConnectorOutput(ctx *pulumi.Context, args LookupOfficePowerBIDataConnectorOutputArgs, opts ...pulumi.InvokeOption) LookupOfficePowerBIDataConnectorResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupOfficePowerBIDataConnectorResult, error) {
			args := v.(LookupOfficePowerBIDataConnectorArgs)
			r, err := LookupOfficePowerBIDataConnector(ctx, &args, opts...)
			var s LookupOfficePowerBIDataConnectorResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupOfficePowerBIDataConnectorResultOutput)
}

type LookupOfficePowerBIDataConnectorOutputArgs struct {
	DataConnectorId   pulumi.StringInput `pulumi:"dataConnectorId"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	WorkspaceName     pulumi.StringInput `pulumi:"workspaceName"`
}

func (LookupOfficePowerBIDataConnectorOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupOfficePowerBIDataConnectorArgs)(nil)).Elem()
}


type LookupOfficePowerBIDataConnectorResultOutput struct{ *pulumi.OutputState }

func (LookupOfficePowerBIDataConnectorResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupOfficePowerBIDataConnectorResult)(nil)).Elem()
}

func (o LookupOfficePowerBIDataConnectorResultOutput) ToLookupOfficePowerBIDataConnectorResultOutput() LookupOfficePowerBIDataConnectorResultOutput {
	return o
}

func (o LookupOfficePowerBIDataConnectorResultOutput) ToLookupOfficePowerBIDataConnectorResultOutputWithContext(ctx context.Context) LookupOfficePowerBIDataConnectorResultOutput {
	return o
}

func (o LookupOfficePowerBIDataConnectorResultOutput) DataTypes() OfficePowerBIConnectorDataTypesResponseOutput {
	return o.ApplyT(func(v LookupOfficePowerBIDataConnectorResult) OfficePowerBIConnectorDataTypesResponse {
		return v.DataTypes
	}).(OfficePowerBIConnectorDataTypesResponseOutput)
}

func (o LookupOfficePowerBIDataConnectorResultOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupOfficePowerBIDataConnectorResult) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o LookupOfficePowerBIDataConnectorResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOfficePowerBIDataConnectorResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupOfficePowerBIDataConnectorResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOfficePowerBIDataConnectorResult) string { return v.Kind }).(pulumi.StringOutput)
}

func (o LookupOfficePowerBIDataConnectorResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOfficePowerBIDataConnectorResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupOfficePowerBIDataConnectorResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupOfficePowerBIDataConnectorResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupOfficePowerBIDataConnectorResultOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOfficePowerBIDataConnectorResult) string { return v.TenantId }).(pulumi.StringOutput)
}

func (o LookupOfficePowerBIDataConnectorResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOfficePowerBIDataConnectorResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupOfficePowerBIDataConnectorResultOutput{})
}
