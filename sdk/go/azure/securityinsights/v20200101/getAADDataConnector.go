


package v20200101

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupAADDataConnector(ctx *pulumi.Context, args *LookupAADDataConnectorArgs, opts ...pulumi.InvokeOption) (*LookupAADDataConnectorResult, error) {
	var rv LookupAADDataConnectorResult
	err := ctx.Invoke("azure-native:securityinsights/v20200101:getAADDataConnector", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupAADDataConnectorArgs struct {
	DataConnectorId   string `pulumi:"dataConnectorId"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	WorkspaceName     string `pulumi:"workspaceName"`
}


type LookupAADDataConnectorResult struct {
	DataTypes *AlertsDataTypeOfDataConnectorResponse `pulumi:"dataTypes"`
	Etag      *string                                `pulumi:"etag"`
	Id        string                                 `pulumi:"id"`
	Kind      string                                 `pulumi:"kind"`
	Name      string                                 `pulumi:"name"`
	TenantId  *string                                `pulumi:"tenantId"`
	Type      string                                 `pulumi:"type"`
}

func LookupAADDataConnectorOutput(ctx *pulumi.Context, args LookupAADDataConnectorOutputArgs, opts ...pulumi.InvokeOption) LookupAADDataConnectorResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupAADDataConnectorResult, error) {
			args := v.(LookupAADDataConnectorArgs)
			r, err := LookupAADDataConnector(ctx, &args, opts...)
			var s LookupAADDataConnectorResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupAADDataConnectorResultOutput)
}

type LookupAADDataConnectorOutputArgs struct {
	DataConnectorId   pulumi.StringInput `pulumi:"dataConnectorId"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	WorkspaceName     pulumi.StringInput `pulumi:"workspaceName"`
}

func (LookupAADDataConnectorOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAADDataConnectorArgs)(nil)).Elem()
}


type LookupAADDataConnectorResultOutput struct{ *pulumi.OutputState }

func (LookupAADDataConnectorResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAADDataConnectorResult)(nil)).Elem()
}

func (o LookupAADDataConnectorResultOutput) ToLookupAADDataConnectorResultOutput() LookupAADDataConnectorResultOutput {
	return o
}

func (o LookupAADDataConnectorResultOutput) ToLookupAADDataConnectorResultOutputWithContext(ctx context.Context) LookupAADDataConnectorResultOutput {
	return o
}

func (o LookupAADDataConnectorResultOutput) DataTypes() AlertsDataTypeOfDataConnectorResponsePtrOutput {
	return o.ApplyT(func(v LookupAADDataConnectorResult) *AlertsDataTypeOfDataConnectorResponse { return v.DataTypes }).(AlertsDataTypeOfDataConnectorResponsePtrOutput)
}

func (o LookupAADDataConnectorResultOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAADDataConnectorResult) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o LookupAADDataConnectorResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAADDataConnectorResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupAADDataConnectorResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAADDataConnectorResult) string { return v.Kind }).(pulumi.StringOutput)
}

func (o LookupAADDataConnectorResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAADDataConnectorResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupAADDataConnectorResultOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAADDataConnectorResult) *string { return v.TenantId }).(pulumi.StringPtrOutput)
}

func (o LookupAADDataConnectorResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAADDataConnectorResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupAADDataConnectorResultOutput{})
}
