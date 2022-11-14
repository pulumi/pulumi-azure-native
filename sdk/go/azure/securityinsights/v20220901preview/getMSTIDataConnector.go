


package v20220901preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupMSTIDataConnector(ctx *pulumi.Context, args *LookupMSTIDataConnectorArgs, opts ...pulumi.InvokeOption) (*LookupMSTIDataConnectorResult, error) {
	var rv LookupMSTIDataConnectorResult
	err := ctx.Invoke("azure-native:securityinsights/v20220901preview:getMSTIDataConnector", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupMSTIDataConnectorArgs struct {
	DataConnectorId   string `pulumi:"dataConnectorId"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	WorkspaceName     string `pulumi:"workspaceName"`
}


type LookupMSTIDataConnectorResult struct {
	DataTypes  MSTIDataConnectorDataTypesResponse `pulumi:"dataTypes"`
	Etag       *string                            `pulumi:"etag"`
	Id         string                             `pulumi:"id"`
	Kind       string                             `pulumi:"kind"`
	Name       string                             `pulumi:"name"`
	SystemData SystemDataResponse                 `pulumi:"systemData"`
	TenantId   string                             `pulumi:"tenantId"`
	Type       string                             `pulumi:"type"`
}

func LookupMSTIDataConnectorOutput(ctx *pulumi.Context, args LookupMSTIDataConnectorOutputArgs, opts ...pulumi.InvokeOption) LookupMSTIDataConnectorResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupMSTIDataConnectorResult, error) {
			args := v.(LookupMSTIDataConnectorArgs)
			r, err := LookupMSTIDataConnector(ctx, &args, opts...)
			var s LookupMSTIDataConnectorResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupMSTIDataConnectorResultOutput)
}

type LookupMSTIDataConnectorOutputArgs struct {
	DataConnectorId   pulumi.StringInput `pulumi:"dataConnectorId"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	WorkspaceName     pulumi.StringInput `pulumi:"workspaceName"`
}

func (LookupMSTIDataConnectorOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupMSTIDataConnectorArgs)(nil)).Elem()
}


type LookupMSTIDataConnectorResultOutput struct{ *pulumi.OutputState }

func (LookupMSTIDataConnectorResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupMSTIDataConnectorResult)(nil)).Elem()
}

func (o LookupMSTIDataConnectorResultOutput) ToLookupMSTIDataConnectorResultOutput() LookupMSTIDataConnectorResultOutput {
	return o
}

func (o LookupMSTIDataConnectorResultOutput) ToLookupMSTIDataConnectorResultOutputWithContext(ctx context.Context) LookupMSTIDataConnectorResultOutput {
	return o
}

func (o LookupMSTIDataConnectorResultOutput) DataTypes() MSTIDataConnectorDataTypesResponseOutput {
	return o.ApplyT(func(v LookupMSTIDataConnectorResult) MSTIDataConnectorDataTypesResponse { return v.DataTypes }).(MSTIDataConnectorDataTypesResponseOutput)
}

func (o LookupMSTIDataConnectorResultOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMSTIDataConnectorResult) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o LookupMSTIDataConnectorResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMSTIDataConnectorResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupMSTIDataConnectorResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMSTIDataConnectorResult) string { return v.Kind }).(pulumi.StringOutput)
}

func (o LookupMSTIDataConnectorResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMSTIDataConnectorResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupMSTIDataConnectorResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupMSTIDataConnectorResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupMSTIDataConnectorResultOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMSTIDataConnectorResult) string { return v.TenantId }).(pulumi.StringOutput)
}

func (o LookupMSTIDataConnectorResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMSTIDataConnectorResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupMSTIDataConnectorResultOutput{})
}
