


package v20220801preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupMCASDataConnector(ctx *pulumi.Context, args *LookupMCASDataConnectorArgs, opts ...pulumi.InvokeOption) (*LookupMCASDataConnectorResult, error) {
	var rv LookupMCASDataConnectorResult
	err := ctx.Invoke("azure-native:securityinsights/v20220801preview:getMCASDataConnector", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupMCASDataConnectorArgs struct {
	DataConnectorId   string `pulumi:"dataConnectorId"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	WorkspaceName     string `pulumi:"workspaceName"`
}


type LookupMCASDataConnectorResult struct {
	DataTypes  MCASDataConnectorDataTypesResponse `pulumi:"dataTypes"`
	Etag       *string                            `pulumi:"etag"`
	Id         string                             `pulumi:"id"`
	Kind       string                             `pulumi:"kind"`
	Name       string                             `pulumi:"name"`
	SystemData SystemDataResponse                 `pulumi:"systemData"`
	TenantId   string                             `pulumi:"tenantId"`
	Type       string                             `pulumi:"type"`
}

func LookupMCASDataConnectorOutput(ctx *pulumi.Context, args LookupMCASDataConnectorOutputArgs, opts ...pulumi.InvokeOption) LookupMCASDataConnectorResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupMCASDataConnectorResult, error) {
			args := v.(LookupMCASDataConnectorArgs)
			r, err := LookupMCASDataConnector(ctx, &args, opts...)
			var s LookupMCASDataConnectorResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupMCASDataConnectorResultOutput)
}

type LookupMCASDataConnectorOutputArgs struct {
	DataConnectorId   pulumi.StringInput `pulumi:"dataConnectorId"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	WorkspaceName     pulumi.StringInput `pulumi:"workspaceName"`
}

func (LookupMCASDataConnectorOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupMCASDataConnectorArgs)(nil)).Elem()
}


type LookupMCASDataConnectorResultOutput struct{ *pulumi.OutputState }

func (LookupMCASDataConnectorResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupMCASDataConnectorResult)(nil)).Elem()
}

func (o LookupMCASDataConnectorResultOutput) ToLookupMCASDataConnectorResultOutput() LookupMCASDataConnectorResultOutput {
	return o
}

func (o LookupMCASDataConnectorResultOutput) ToLookupMCASDataConnectorResultOutputWithContext(ctx context.Context) LookupMCASDataConnectorResultOutput {
	return o
}

func (o LookupMCASDataConnectorResultOutput) DataTypes() MCASDataConnectorDataTypesResponseOutput {
	return o.ApplyT(func(v LookupMCASDataConnectorResult) MCASDataConnectorDataTypesResponse { return v.DataTypes }).(MCASDataConnectorDataTypesResponseOutput)
}

func (o LookupMCASDataConnectorResultOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMCASDataConnectorResult) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o LookupMCASDataConnectorResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMCASDataConnectorResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupMCASDataConnectorResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMCASDataConnectorResult) string { return v.Kind }).(pulumi.StringOutput)
}

func (o LookupMCASDataConnectorResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMCASDataConnectorResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupMCASDataConnectorResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupMCASDataConnectorResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupMCASDataConnectorResultOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMCASDataConnectorResult) string { return v.TenantId }).(pulumi.StringOutput)
}

func (o LookupMCASDataConnectorResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMCASDataConnectorResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupMCASDataConnectorResultOutput{})
}
