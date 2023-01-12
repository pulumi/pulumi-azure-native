


package v20220101preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupIoTDataConnector(ctx *pulumi.Context, args *LookupIoTDataConnectorArgs, opts ...pulumi.InvokeOption) (*LookupIoTDataConnectorResult, error) {
	var rv LookupIoTDataConnectorResult
	err := ctx.Invoke("azure-native:securityinsights/v20220101preview:getIoTDataConnector", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupIoTDataConnectorArgs struct {
	DataConnectorId   string `pulumi:"dataConnectorId"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	WorkspaceName     string `pulumi:"workspaceName"`
}


type LookupIoTDataConnectorResult struct {
	DataTypes      *AlertsDataTypeOfDataConnectorResponse `pulumi:"dataTypes"`
	Etag           *string                                `pulumi:"etag"`
	Id             string                                 `pulumi:"id"`
	Kind           string                                 `pulumi:"kind"`
	Name           string                                 `pulumi:"name"`
	SubscriptionId *string                                `pulumi:"subscriptionId"`
	SystemData     SystemDataResponse                     `pulumi:"systemData"`
	Type           string                                 `pulumi:"type"`
}

func LookupIoTDataConnectorOutput(ctx *pulumi.Context, args LookupIoTDataConnectorOutputArgs, opts ...pulumi.InvokeOption) LookupIoTDataConnectorResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupIoTDataConnectorResult, error) {
			args := v.(LookupIoTDataConnectorArgs)
			r, err := LookupIoTDataConnector(ctx, &args, opts...)
			var s LookupIoTDataConnectorResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupIoTDataConnectorResultOutput)
}

type LookupIoTDataConnectorOutputArgs struct {
	DataConnectorId   pulumi.StringInput `pulumi:"dataConnectorId"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	WorkspaceName     pulumi.StringInput `pulumi:"workspaceName"`
}

func (LookupIoTDataConnectorOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupIoTDataConnectorArgs)(nil)).Elem()
}


type LookupIoTDataConnectorResultOutput struct{ *pulumi.OutputState }

func (LookupIoTDataConnectorResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupIoTDataConnectorResult)(nil)).Elem()
}

func (o LookupIoTDataConnectorResultOutput) ToLookupIoTDataConnectorResultOutput() LookupIoTDataConnectorResultOutput {
	return o
}

func (o LookupIoTDataConnectorResultOutput) ToLookupIoTDataConnectorResultOutputWithContext(ctx context.Context) LookupIoTDataConnectorResultOutput {
	return o
}

func (o LookupIoTDataConnectorResultOutput) DataTypes() AlertsDataTypeOfDataConnectorResponsePtrOutput {
	return o.ApplyT(func(v LookupIoTDataConnectorResult) *AlertsDataTypeOfDataConnectorResponse { return v.DataTypes }).(AlertsDataTypeOfDataConnectorResponsePtrOutput)
}

func (o LookupIoTDataConnectorResultOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupIoTDataConnectorResult) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o LookupIoTDataConnectorResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIoTDataConnectorResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupIoTDataConnectorResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIoTDataConnectorResult) string { return v.Kind }).(pulumi.StringOutput)
}

func (o LookupIoTDataConnectorResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIoTDataConnectorResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupIoTDataConnectorResultOutput) SubscriptionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupIoTDataConnectorResult) *string { return v.SubscriptionId }).(pulumi.StringPtrOutput)
}

func (o LookupIoTDataConnectorResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupIoTDataConnectorResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupIoTDataConnectorResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIoTDataConnectorResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupIoTDataConnectorResultOutput{})
}
