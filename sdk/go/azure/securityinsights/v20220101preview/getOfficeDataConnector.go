


package v20220101preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupOfficeDataConnector(ctx *pulumi.Context, args *LookupOfficeDataConnectorArgs, opts ...pulumi.InvokeOption) (*LookupOfficeDataConnectorResult, error) {
	var rv LookupOfficeDataConnectorResult
	err := ctx.Invoke("azure-native:securityinsights/v20220101preview:getOfficeDataConnector", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupOfficeDataConnectorArgs struct {
	DataConnectorId   string `pulumi:"dataConnectorId"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	WorkspaceName     string `pulumi:"workspaceName"`
}


type LookupOfficeDataConnectorResult struct {
	DataTypes  OfficeDataConnectorDataTypesResponse `pulumi:"dataTypes"`
	Etag       *string                              `pulumi:"etag"`
	Id         string                               `pulumi:"id"`
	Kind       string                               `pulumi:"kind"`
	Name       string                               `pulumi:"name"`
	SystemData SystemDataResponse                   `pulumi:"systemData"`
	TenantId   string                               `pulumi:"tenantId"`
	Type       string                               `pulumi:"type"`
}

func LookupOfficeDataConnectorOutput(ctx *pulumi.Context, args LookupOfficeDataConnectorOutputArgs, opts ...pulumi.InvokeOption) LookupOfficeDataConnectorResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupOfficeDataConnectorResult, error) {
			args := v.(LookupOfficeDataConnectorArgs)
			r, err := LookupOfficeDataConnector(ctx, &args, opts...)
			var s LookupOfficeDataConnectorResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupOfficeDataConnectorResultOutput)
}

type LookupOfficeDataConnectorOutputArgs struct {
	DataConnectorId   pulumi.StringInput `pulumi:"dataConnectorId"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	WorkspaceName     pulumi.StringInput `pulumi:"workspaceName"`
}

func (LookupOfficeDataConnectorOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupOfficeDataConnectorArgs)(nil)).Elem()
}


type LookupOfficeDataConnectorResultOutput struct{ *pulumi.OutputState }

func (LookupOfficeDataConnectorResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupOfficeDataConnectorResult)(nil)).Elem()
}

func (o LookupOfficeDataConnectorResultOutput) ToLookupOfficeDataConnectorResultOutput() LookupOfficeDataConnectorResultOutput {
	return o
}

func (o LookupOfficeDataConnectorResultOutput) ToLookupOfficeDataConnectorResultOutputWithContext(ctx context.Context) LookupOfficeDataConnectorResultOutput {
	return o
}

func (o LookupOfficeDataConnectorResultOutput) DataTypes() OfficeDataConnectorDataTypesResponseOutput {
	return o.ApplyT(func(v LookupOfficeDataConnectorResult) OfficeDataConnectorDataTypesResponse { return v.DataTypes }).(OfficeDataConnectorDataTypesResponseOutput)
}

func (o LookupOfficeDataConnectorResultOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupOfficeDataConnectorResult) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o LookupOfficeDataConnectorResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOfficeDataConnectorResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupOfficeDataConnectorResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOfficeDataConnectorResult) string { return v.Kind }).(pulumi.StringOutput)
}

func (o LookupOfficeDataConnectorResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOfficeDataConnectorResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupOfficeDataConnectorResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupOfficeDataConnectorResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupOfficeDataConnectorResultOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOfficeDataConnectorResult) string { return v.TenantId }).(pulumi.StringOutput)
}

func (o LookupOfficeDataConnectorResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOfficeDataConnectorResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupOfficeDataConnectorResultOutput{})
}
