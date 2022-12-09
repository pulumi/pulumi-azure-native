


package v20210301preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupMDATPDataConnector(ctx *pulumi.Context, args *LookupMDATPDataConnectorArgs, opts ...pulumi.InvokeOption) (*LookupMDATPDataConnectorResult, error) {
	var rv LookupMDATPDataConnectorResult
	err := ctx.Invoke("azure-native:securityinsights/v20210301preview:getMDATPDataConnector", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupMDATPDataConnectorArgs struct {
	DataConnectorId                     string `pulumi:"dataConnectorId"`
	OperationalInsightsResourceProvider string `pulumi:"operationalInsightsResourceProvider"`
	ResourceGroupName                   string `pulumi:"resourceGroupName"`
	WorkspaceName                       string `pulumi:"workspaceName"`
}


type LookupMDATPDataConnectorResult struct {
	DataTypes  *AlertsDataTypeOfDataConnectorResponse `pulumi:"dataTypes"`
	Etag       *string                                `pulumi:"etag"`
	Id         string                                 `pulumi:"id"`
	Kind       string                                 `pulumi:"kind"`
	Name       string                                 `pulumi:"name"`
	SystemData SystemDataResponse                     `pulumi:"systemData"`
	TenantId   string                                 `pulumi:"tenantId"`
	Type       string                                 `pulumi:"type"`
}

func LookupMDATPDataConnectorOutput(ctx *pulumi.Context, args LookupMDATPDataConnectorOutputArgs, opts ...pulumi.InvokeOption) LookupMDATPDataConnectorResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupMDATPDataConnectorResult, error) {
			args := v.(LookupMDATPDataConnectorArgs)
			r, err := LookupMDATPDataConnector(ctx, &args, opts...)
			var s LookupMDATPDataConnectorResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupMDATPDataConnectorResultOutput)
}

type LookupMDATPDataConnectorOutputArgs struct {
	DataConnectorId                     pulumi.StringInput `pulumi:"dataConnectorId"`
	OperationalInsightsResourceProvider pulumi.StringInput `pulumi:"operationalInsightsResourceProvider"`
	ResourceGroupName                   pulumi.StringInput `pulumi:"resourceGroupName"`
	WorkspaceName                       pulumi.StringInput `pulumi:"workspaceName"`
}

func (LookupMDATPDataConnectorOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupMDATPDataConnectorArgs)(nil)).Elem()
}


type LookupMDATPDataConnectorResultOutput struct{ *pulumi.OutputState }

func (LookupMDATPDataConnectorResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupMDATPDataConnectorResult)(nil)).Elem()
}

func (o LookupMDATPDataConnectorResultOutput) ToLookupMDATPDataConnectorResultOutput() LookupMDATPDataConnectorResultOutput {
	return o
}

func (o LookupMDATPDataConnectorResultOutput) ToLookupMDATPDataConnectorResultOutputWithContext(ctx context.Context) LookupMDATPDataConnectorResultOutput {
	return o
}

func (o LookupMDATPDataConnectorResultOutput) DataTypes() AlertsDataTypeOfDataConnectorResponsePtrOutput {
	return o.ApplyT(func(v LookupMDATPDataConnectorResult) *AlertsDataTypeOfDataConnectorResponse { return v.DataTypes }).(AlertsDataTypeOfDataConnectorResponsePtrOutput)
}

func (o LookupMDATPDataConnectorResultOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMDATPDataConnectorResult) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o LookupMDATPDataConnectorResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMDATPDataConnectorResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupMDATPDataConnectorResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMDATPDataConnectorResult) string { return v.Kind }).(pulumi.StringOutput)
}

func (o LookupMDATPDataConnectorResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMDATPDataConnectorResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupMDATPDataConnectorResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupMDATPDataConnectorResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupMDATPDataConnectorResultOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMDATPDataConnectorResult) string { return v.TenantId }).(pulumi.StringOutput)
}

func (o LookupMDATPDataConnectorResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMDATPDataConnectorResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupMDATPDataConnectorResultOutput{})
}
