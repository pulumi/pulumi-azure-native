


package v20220501preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupAATPDataConnector(ctx *pulumi.Context, args *LookupAATPDataConnectorArgs, opts ...pulumi.InvokeOption) (*LookupAATPDataConnectorResult, error) {
	var rv LookupAATPDataConnectorResult
	err := ctx.Invoke("azure-native:securityinsights/v20220501preview:getAATPDataConnector", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupAATPDataConnectorArgs struct {
	DataConnectorId   string `pulumi:"dataConnectorId"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	WorkspaceName     string `pulumi:"workspaceName"`
}


type LookupAATPDataConnectorResult struct {
	DataTypes  *AlertsDataTypeOfDataConnectorResponse `pulumi:"dataTypes"`
	Etag       *string                                `pulumi:"etag"`
	Id         string                                 `pulumi:"id"`
	Kind       string                                 `pulumi:"kind"`
	Name       string                                 `pulumi:"name"`
	SystemData SystemDataResponse                     `pulumi:"systemData"`
	TenantId   string                                 `pulumi:"tenantId"`
	Type       string                                 `pulumi:"type"`
}

func LookupAATPDataConnectorOutput(ctx *pulumi.Context, args LookupAATPDataConnectorOutputArgs, opts ...pulumi.InvokeOption) LookupAATPDataConnectorResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupAATPDataConnectorResult, error) {
			args := v.(LookupAATPDataConnectorArgs)
			r, err := LookupAATPDataConnector(ctx, &args, opts...)
			var s LookupAATPDataConnectorResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupAATPDataConnectorResultOutput)
}

type LookupAATPDataConnectorOutputArgs struct {
	DataConnectorId   pulumi.StringInput `pulumi:"dataConnectorId"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	WorkspaceName     pulumi.StringInput `pulumi:"workspaceName"`
}

func (LookupAATPDataConnectorOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAATPDataConnectorArgs)(nil)).Elem()
}


type LookupAATPDataConnectorResultOutput struct{ *pulumi.OutputState }

func (LookupAATPDataConnectorResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAATPDataConnectorResult)(nil)).Elem()
}

func (o LookupAATPDataConnectorResultOutput) ToLookupAATPDataConnectorResultOutput() LookupAATPDataConnectorResultOutput {
	return o
}

func (o LookupAATPDataConnectorResultOutput) ToLookupAATPDataConnectorResultOutputWithContext(ctx context.Context) LookupAATPDataConnectorResultOutput {
	return o
}

func (o LookupAATPDataConnectorResultOutput) DataTypes() AlertsDataTypeOfDataConnectorResponsePtrOutput {
	return o.ApplyT(func(v LookupAATPDataConnectorResult) *AlertsDataTypeOfDataConnectorResponse { return v.DataTypes }).(AlertsDataTypeOfDataConnectorResponsePtrOutput)
}

func (o LookupAATPDataConnectorResultOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAATPDataConnectorResult) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o LookupAATPDataConnectorResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAATPDataConnectorResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupAATPDataConnectorResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAATPDataConnectorResult) string { return v.Kind }).(pulumi.StringOutput)
}

func (o LookupAATPDataConnectorResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAATPDataConnectorResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupAATPDataConnectorResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupAATPDataConnectorResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupAATPDataConnectorResultOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAATPDataConnectorResult) string { return v.TenantId }).(pulumi.StringOutput)
}

func (o LookupAATPDataConnectorResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAATPDataConnectorResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupAATPDataConnectorResultOutput{})
}
