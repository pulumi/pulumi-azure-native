


package v20211001preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupOfficeIRMDataConnector(ctx *pulumi.Context, args *LookupOfficeIRMDataConnectorArgs, opts ...pulumi.InvokeOption) (*LookupOfficeIRMDataConnectorResult, error) {
	var rv LookupOfficeIRMDataConnectorResult
	err := ctx.Invoke("azure-native:securityinsights/v20211001preview:getOfficeIRMDataConnector", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupOfficeIRMDataConnectorArgs struct {
	DataConnectorId   string `pulumi:"dataConnectorId"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	WorkspaceName     string `pulumi:"workspaceName"`
}


type LookupOfficeIRMDataConnectorResult struct {
	DataTypes  *AlertsDataTypeOfDataConnectorResponse `pulumi:"dataTypes"`
	Etag       *string                                `pulumi:"etag"`
	Id         string                                 `pulumi:"id"`
	Kind       string                                 `pulumi:"kind"`
	Name       string                                 `pulumi:"name"`
	SystemData SystemDataResponse                     `pulumi:"systemData"`
	TenantId   string                                 `pulumi:"tenantId"`
	Type       string                                 `pulumi:"type"`
}

func LookupOfficeIRMDataConnectorOutput(ctx *pulumi.Context, args LookupOfficeIRMDataConnectorOutputArgs, opts ...pulumi.InvokeOption) LookupOfficeIRMDataConnectorResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupOfficeIRMDataConnectorResult, error) {
			args := v.(LookupOfficeIRMDataConnectorArgs)
			r, err := LookupOfficeIRMDataConnector(ctx, &args, opts...)
			var s LookupOfficeIRMDataConnectorResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupOfficeIRMDataConnectorResultOutput)
}

type LookupOfficeIRMDataConnectorOutputArgs struct {
	DataConnectorId   pulumi.StringInput `pulumi:"dataConnectorId"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	WorkspaceName     pulumi.StringInput `pulumi:"workspaceName"`
}

func (LookupOfficeIRMDataConnectorOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupOfficeIRMDataConnectorArgs)(nil)).Elem()
}


type LookupOfficeIRMDataConnectorResultOutput struct{ *pulumi.OutputState }

func (LookupOfficeIRMDataConnectorResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupOfficeIRMDataConnectorResult)(nil)).Elem()
}

func (o LookupOfficeIRMDataConnectorResultOutput) ToLookupOfficeIRMDataConnectorResultOutput() LookupOfficeIRMDataConnectorResultOutput {
	return o
}

func (o LookupOfficeIRMDataConnectorResultOutput) ToLookupOfficeIRMDataConnectorResultOutputWithContext(ctx context.Context) LookupOfficeIRMDataConnectorResultOutput {
	return o
}

func (o LookupOfficeIRMDataConnectorResultOutput) DataTypes() AlertsDataTypeOfDataConnectorResponsePtrOutput {
	return o.ApplyT(func(v LookupOfficeIRMDataConnectorResult) *AlertsDataTypeOfDataConnectorResponse { return v.DataTypes }).(AlertsDataTypeOfDataConnectorResponsePtrOutput)
}

func (o LookupOfficeIRMDataConnectorResultOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupOfficeIRMDataConnectorResult) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o LookupOfficeIRMDataConnectorResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOfficeIRMDataConnectorResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupOfficeIRMDataConnectorResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOfficeIRMDataConnectorResult) string { return v.Kind }).(pulumi.StringOutput)
}

func (o LookupOfficeIRMDataConnectorResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOfficeIRMDataConnectorResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupOfficeIRMDataConnectorResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupOfficeIRMDataConnectorResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupOfficeIRMDataConnectorResultOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOfficeIRMDataConnectorResult) string { return v.TenantId }).(pulumi.StringOutput)
}

func (o LookupOfficeIRMDataConnectorResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOfficeIRMDataConnectorResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupOfficeIRMDataConnectorResultOutput{})
}
