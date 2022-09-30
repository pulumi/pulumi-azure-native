


package v20220901preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupCodelessApiPollingDataConnector(ctx *pulumi.Context, args *LookupCodelessApiPollingDataConnectorArgs, opts ...pulumi.InvokeOption) (*LookupCodelessApiPollingDataConnectorResult, error) {
	var rv LookupCodelessApiPollingDataConnectorResult
	err := ctx.Invoke("azure-native:securityinsights/v20220901preview:getCodelessApiPollingDataConnector", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupCodelessApiPollingDataConnectorArgs struct {
	DataConnectorId   string `pulumi:"dataConnectorId"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	WorkspaceName     string `pulumi:"workspaceName"`
}


type LookupCodelessApiPollingDataConnectorResult struct {
	ConnectorUiConfig *CodelessUiConnectorConfigPropertiesResponse      `pulumi:"connectorUiConfig"`
	Etag              *string                                           `pulumi:"etag"`
	Id                string                                            `pulumi:"id"`
	Kind              string                                            `pulumi:"kind"`
	Name              string                                            `pulumi:"name"`
	PollingConfig     *CodelessConnectorPollingConfigPropertiesResponse `pulumi:"pollingConfig"`
	SystemData        SystemDataResponse                                `pulumi:"systemData"`
	Type              string                                            `pulumi:"type"`
}

func LookupCodelessApiPollingDataConnectorOutput(ctx *pulumi.Context, args LookupCodelessApiPollingDataConnectorOutputArgs, opts ...pulumi.InvokeOption) LookupCodelessApiPollingDataConnectorResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupCodelessApiPollingDataConnectorResult, error) {
			args := v.(LookupCodelessApiPollingDataConnectorArgs)
			r, err := LookupCodelessApiPollingDataConnector(ctx, &args, opts...)
			var s LookupCodelessApiPollingDataConnectorResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupCodelessApiPollingDataConnectorResultOutput)
}

type LookupCodelessApiPollingDataConnectorOutputArgs struct {
	DataConnectorId   pulumi.StringInput `pulumi:"dataConnectorId"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	WorkspaceName     pulumi.StringInput `pulumi:"workspaceName"`
}

func (LookupCodelessApiPollingDataConnectorOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCodelessApiPollingDataConnectorArgs)(nil)).Elem()
}


type LookupCodelessApiPollingDataConnectorResultOutput struct{ *pulumi.OutputState }

func (LookupCodelessApiPollingDataConnectorResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCodelessApiPollingDataConnectorResult)(nil)).Elem()
}

func (o LookupCodelessApiPollingDataConnectorResultOutput) ToLookupCodelessApiPollingDataConnectorResultOutput() LookupCodelessApiPollingDataConnectorResultOutput {
	return o
}

func (o LookupCodelessApiPollingDataConnectorResultOutput) ToLookupCodelessApiPollingDataConnectorResultOutputWithContext(ctx context.Context) LookupCodelessApiPollingDataConnectorResultOutput {
	return o
}

func (o LookupCodelessApiPollingDataConnectorResultOutput) ConnectorUiConfig() CodelessUiConnectorConfigPropertiesResponsePtrOutput {
	return o.ApplyT(func(v LookupCodelessApiPollingDataConnectorResult) *CodelessUiConnectorConfigPropertiesResponse {
		return v.ConnectorUiConfig
	}).(CodelessUiConnectorConfigPropertiesResponsePtrOutput)
}

func (o LookupCodelessApiPollingDataConnectorResultOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCodelessApiPollingDataConnectorResult) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o LookupCodelessApiPollingDataConnectorResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCodelessApiPollingDataConnectorResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupCodelessApiPollingDataConnectorResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCodelessApiPollingDataConnectorResult) string { return v.Kind }).(pulumi.StringOutput)
}

func (o LookupCodelessApiPollingDataConnectorResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCodelessApiPollingDataConnectorResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupCodelessApiPollingDataConnectorResultOutput) PollingConfig() CodelessConnectorPollingConfigPropertiesResponsePtrOutput {
	return o.ApplyT(func(v LookupCodelessApiPollingDataConnectorResult) *CodelessConnectorPollingConfigPropertiesResponse {
		return v.PollingConfig
	}).(CodelessConnectorPollingConfigPropertiesResponsePtrOutput)
}

func (o LookupCodelessApiPollingDataConnectorResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupCodelessApiPollingDataConnectorResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupCodelessApiPollingDataConnectorResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCodelessApiPollingDataConnectorResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupCodelessApiPollingDataConnectorResultOutput{})
}
