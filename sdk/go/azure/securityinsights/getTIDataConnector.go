


package securityinsights

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupTIDataConnector(ctx *pulumi.Context, args *LookupTIDataConnectorArgs, opts ...pulumi.InvokeOption) (*LookupTIDataConnectorResult, error) {
	var rv LookupTIDataConnectorResult
	err := ctx.Invoke("azure-native:securityinsights:getTIDataConnector", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupTIDataConnectorArgs struct {
	DataConnectorId   string `pulumi:"dataConnectorId"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	WorkspaceName     string `pulumi:"workspaceName"`
}


type LookupTIDataConnectorResult struct {
	DataTypes         *TIDataConnectorDataTypesResponse `pulumi:"dataTypes"`
	Etag              *string                           `pulumi:"etag"`
	Id                string                            `pulumi:"id"`
	Kind              string                            `pulumi:"kind"`
	Name              string                            `pulumi:"name"`
	TenantId          *string                           `pulumi:"tenantId"`
	TipLookbackPeriod *string                           `pulumi:"tipLookbackPeriod"`
	Type              string                            `pulumi:"type"`
}

func LookupTIDataConnectorOutput(ctx *pulumi.Context, args LookupTIDataConnectorOutputArgs, opts ...pulumi.InvokeOption) LookupTIDataConnectorResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupTIDataConnectorResult, error) {
			args := v.(LookupTIDataConnectorArgs)
			r, err := LookupTIDataConnector(ctx, &args, opts...)
			var s LookupTIDataConnectorResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupTIDataConnectorResultOutput)
}

type LookupTIDataConnectorOutputArgs struct {
	DataConnectorId   pulumi.StringInput `pulumi:"dataConnectorId"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	WorkspaceName     pulumi.StringInput `pulumi:"workspaceName"`
}

func (LookupTIDataConnectorOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTIDataConnectorArgs)(nil)).Elem()
}


type LookupTIDataConnectorResultOutput struct{ *pulumi.OutputState }

func (LookupTIDataConnectorResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTIDataConnectorResult)(nil)).Elem()
}

func (o LookupTIDataConnectorResultOutput) ToLookupTIDataConnectorResultOutput() LookupTIDataConnectorResultOutput {
	return o
}

func (o LookupTIDataConnectorResultOutput) ToLookupTIDataConnectorResultOutputWithContext(ctx context.Context) LookupTIDataConnectorResultOutput {
	return o
}

func (o LookupTIDataConnectorResultOutput) DataTypes() TIDataConnectorDataTypesResponsePtrOutput {
	return o.ApplyT(func(v LookupTIDataConnectorResult) *TIDataConnectorDataTypesResponse { return v.DataTypes }).(TIDataConnectorDataTypesResponsePtrOutput)
}

func (o LookupTIDataConnectorResultOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupTIDataConnectorResult) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o LookupTIDataConnectorResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTIDataConnectorResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupTIDataConnectorResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTIDataConnectorResult) string { return v.Kind }).(pulumi.StringOutput)
}

func (o LookupTIDataConnectorResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTIDataConnectorResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupTIDataConnectorResultOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupTIDataConnectorResult) *string { return v.TenantId }).(pulumi.StringPtrOutput)
}

func (o LookupTIDataConnectorResultOutput) TipLookbackPeriod() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupTIDataConnectorResult) *string { return v.TipLookbackPeriod }).(pulumi.StringPtrOutput)
}

func (o LookupTIDataConnectorResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTIDataConnectorResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupTIDataConnectorResultOutput{})
}
