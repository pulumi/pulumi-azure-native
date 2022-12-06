


package v20211001preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupASCDataConnector(ctx *pulumi.Context, args *LookupASCDataConnectorArgs, opts ...pulumi.InvokeOption) (*LookupASCDataConnectorResult, error) {
	var rv LookupASCDataConnectorResult
	err := ctx.Invoke("azure-native:securityinsights/v20211001preview:getASCDataConnector", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupASCDataConnectorArgs struct {
	DataConnectorId   string `pulumi:"dataConnectorId"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	WorkspaceName     string `pulumi:"workspaceName"`
}


type LookupASCDataConnectorResult struct {
	DataTypes      *AlertsDataTypeOfDataConnectorResponse `pulumi:"dataTypes"`
	Etag           *string                                `pulumi:"etag"`
	Id             string                                 `pulumi:"id"`
	Kind           string                                 `pulumi:"kind"`
	Name           string                                 `pulumi:"name"`
	SubscriptionId *string                                `pulumi:"subscriptionId"`
	SystemData     SystemDataResponse                     `pulumi:"systemData"`
	Type           string                                 `pulumi:"type"`
}

func LookupASCDataConnectorOutput(ctx *pulumi.Context, args LookupASCDataConnectorOutputArgs, opts ...pulumi.InvokeOption) LookupASCDataConnectorResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupASCDataConnectorResult, error) {
			args := v.(LookupASCDataConnectorArgs)
			r, err := LookupASCDataConnector(ctx, &args, opts...)
			var s LookupASCDataConnectorResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupASCDataConnectorResultOutput)
}

type LookupASCDataConnectorOutputArgs struct {
	DataConnectorId   pulumi.StringInput `pulumi:"dataConnectorId"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	WorkspaceName     pulumi.StringInput `pulumi:"workspaceName"`
}

func (LookupASCDataConnectorOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupASCDataConnectorArgs)(nil)).Elem()
}


type LookupASCDataConnectorResultOutput struct{ *pulumi.OutputState }

func (LookupASCDataConnectorResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupASCDataConnectorResult)(nil)).Elem()
}

func (o LookupASCDataConnectorResultOutput) ToLookupASCDataConnectorResultOutput() LookupASCDataConnectorResultOutput {
	return o
}

func (o LookupASCDataConnectorResultOutput) ToLookupASCDataConnectorResultOutputWithContext(ctx context.Context) LookupASCDataConnectorResultOutput {
	return o
}

func (o LookupASCDataConnectorResultOutput) DataTypes() AlertsDataTypeOfDataConnectorResponsePtrOutput {
	return o.ApplyT(func(v LookupASCDataConnectorResult) *AlertsDataTypeOfDataConnectorResponse { return v.DataTypes }).(AlertsDataTypeOfDataConnectorResponsePtrOutput)
}

func (o LookupASCDataConnectorResultOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupASCDataConnectorResult) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o LookupASCDataConnectorResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupASCDataConnectorResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupASCDataConnectorResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupASCDataConnectorResult) string { return v.Kind }).(pulumi.StringOutput)
}

func (o LookupASCDataConnectorResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupASCDataConnectorResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupASCDataConnectorResultOutput) SubscriptionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupASCDataConnectorResult) *string { return v.SubscriptionId }).(pulumi.StringPtrOutput)
}

func (o LookupASCDataConnectorResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupASCDataConnectorResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupASCDataConnectorResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupASCDataConnectorResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupASCDataConnectorResultOutput{})
}
