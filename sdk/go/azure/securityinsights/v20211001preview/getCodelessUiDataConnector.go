


package v20211001preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupCodelessUiDataConnector(ctx *pulumi.Context, args *LookupCodelessUiDataConnectorArgs, opts ...pulumi.InvokeOption) (*LookupCodelessUiDataConnectorResult, error) {
	var rv LookupCodelessUiDataConnectorResult
	err := ctx.Invoke("azure-native:securityinsights/v20211001preview:getCodelessUiDataConnector", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupCodelessUiDataConnectorArgs struct {
	DataConnectorId   string `pulumi:"dataConnectorId"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	WorkspaceName     string `pulumi:"workspaceName"`
}


type LookupCodelessUiDataConnectorResult struct {
	ConnectorUiConfig *CodelessUiConnectorConfigPropertiesResponse `pulumi:"connectorUiConfig"`
	Etag              *string                                      `pulumi:"etag"`
	Id                string                                       `pulumi:"id"`
	Kind              string                                       `pulumi:"kind"`
	Name              string                                       `pulumi:"name"`
	SystemData        SystemDataResponse                           `pulumi:"systemData"`
	Type              string                                       `pulumi:"type"`
}

func LookupCodelessUiDataConnectorOutput(ctx *pulumi.Context, args LookupCodelessUiDataConnectorOutputArgs, opts ...pulumi.InvokeOption) LookupCodelessUiDataConnectorResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupCodelessUiDataConnectorResult, error) {
			args := v.(LookupCodelessUiDataConnectorArgs)
			r, err := LookupCodelessUiDataConnector(ctx, &args, opts...)
			var s LookupCodelessUiDataConnectorResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupCodelessUiDataConnectorResultOutput)
}

type LookupCodelessUiDataConnectorOutputArgs struct {
	DataConnectorId   pulumi.StringInput `pulumi:"dataConnectorId"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	WorkspaceName     pulumi.StringInput `pulumi:"workspaceName"`
}

func (LookupCodelessUiDataConnectorOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCodelessUiDataConnectorArgs)(nil)).Elem()
}


type LookupCodelessUiDataConnectorResultOutput struct{ *pulumi.OutputState }

func (LookupCodelessUiDataConnectorResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCodelessUiDataConnectorResult)(nil)).Elem()
}

func (o LookupCodelessUiDataConnectorResultOutput) ToLookupCodelessUiDataConnectorResultOutput() LookupCodelessUiDataConnectorResultOutput {
	return o
}

func (o LookupCodelessUiDataConnectorResultOutput) ToLookupCodelessUiDataConnectorResultOutputWithContext(ctx context.Context) LookupCodelessUiDataConnectorResultOutput {
	return o
}

func (o LookupCodelessUiDataConnectorResultOutput) ConnectorUiConfig() CodelessUiConnectorConfigPropertiesResponsePtrOutput {
	return o.ApplyT(func(v LookupCodelessUiDataConnectorResult) *CodelessUiConnectorConfigPropertiesResponse {
		return v.ConnectorUiConfig
	}).(CodelessUiConnectorConfigPropertiesResponsePtrOutput)
}

func (o LookupCodelessUiDataConnectorResultOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCodelessUiDataConnectorResult) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o LookupCodelessUiDataConnectorResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCodelessUiDataConnectorResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupCodelessUiDataConnectorResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCodelessUiDataConnectorResult) string { return v.Kind }).(pulumi.StringOutput)
}

func (o LookupCodelessUiDataConnectorResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCodelessUiDataConnectorResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupCodelessUiDataConnectorResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupCodelessUiDataConnectorResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupCodelessUiDataConnectorResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCodelessUiDataConnectorResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupCodelessUiDataConnectorResultOutput{})
}
