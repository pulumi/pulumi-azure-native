


package v20221201preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupOfficeATPDataConnector(ctx *pulumi.Context, args *LookupOfficeATPDataConnectorArgs, opts ...pulumi.InvokeOption) (*LookupOfficeATPDataConnectorResult, error) {
	var rv LookupOfficeATPDataConnectorResult
	err := ctx.Invoke("azure-native:securityinsights/v20221201preview:getOfficeATPDataConnector", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupOfficeATPDataConnectorArgs struct {
	DataConnectorId   string `pulumi:"dataConnectorId"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	WorkspaceName     string `pulumi:"workspaceName"`
}


type LookupOfficeATPDataConnectorResult struct {
	DataTypes  *AlertsDataTypeOfDataConnectorResponse `pulumi:"dataTypes"`
	Etag       *string                                `pulumi:"etag"`
	Id         string                                 `pulumi:"id"`
	Kind       string                                 `pulumi:"kind"`
	Name       string                                 `pulumi:"name"`
	SystemData SystemDataResponse                     `pulumi:"systemData"`
	TenantId   string                                 `pulumi:"tenantId"`
	Type       string                                 `pulumi:"type"`
}

func LookupOfficeATPDataConnectorOutput(ctx *pulumi.Context, args LookupOfficeATPDataConnectorOutputArgs, opts ...pulumi.InvokeOption) LookupOfficeATPDataConnectorResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupOfficeATPDataConnectorResult, error) {
			args := v.(LookupOfficeATPDataConnectorArgs)
			r, err := LookupOfficeATPDataConnector(ctx, &args, opts...)
			var s LookupOfficeATPDataConnectorResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupOfficeATPDataConnectorResultOutput)
}

type LookupOfficeATPDataConnectorOutputArgs struct {
	DataConnectorId   pulumi.StringInput `pulumi:"dataConnectorId"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	WorkspaceName     pulumi.StringInput `pulumi:"workspaceName"`
}

func (LookupOfficeATPDataConnectorOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupOfficeATPDataConnectorArgs)(nil)).Elem()
}


type LookupOfficeATPDataConnectorResultOutput struct{ *pulumi.OutputState }

func (LookupOfficeATPDataConnectorResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupOfficeATPDataConnectorResult)(nil)).Elem()
}

func (o LookupOfficeATPDataConnectorResultOutput) ToLookupOfficeATPDataConnectorResultOutput() LookupOfficeATPDataConnectorResultOutput {
	return o
}

func (o LookupOfficeATPDataConnectorResultOutput) ToLookupOfficeATPDataConnectorResultOutputWithContext(ctx context.Context) LookupOfficeATPDataConnectorResultOutput {
	return o
}

func (o LookupOfficeATPDataConnectorResultOutput) DataTypes() AlertsDataTypeOfDataConnectorResponsePtrOutput {
	return o.ApplyT(func(v LookupOfficeATPDataConnectorResult) *AlertsDataTypeOfDataConnectorResponse { return v.DataTypes }).(AlertsDataTypeOfDataConnectorResponsePtrOutput)
}

func (o LookupOfficeATPDataConnectorResultOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupOfficeATPDataConnectorResult) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o LookupOfficeATPDataConnectorResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOfficeATPDataConnectorResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupOfficeATPDataConnectorResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOfficeATPDataConnectorResult) string { return v.Kind }).(pulumi.StringOutput)
}

func (o LookupOfficeATPDataConnectorResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOfficeATPDataConnectorResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupOfficeATPDataConnectorResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupOfficeATPDataConnectorResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupOfficeATPDataConnectorResultOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOfficeATPDataConnectorResult) string { return v.TenantId }).(pulumi.StringOutput)
}

func (o LookupOfficeATPDataConnectorResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOfficeATPDataConnectorResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupOfficeATPDataConnectorResultOutput{})
}
