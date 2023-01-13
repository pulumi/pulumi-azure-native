


package v20221201preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupMTPDataConnector(ctx *pulumi.Context, args *LookupMTPDataConnectorArgs, opts ...pulumi.InvokeOption) (*LookupMTPDataConnectorResult, error) {
	var rv LookupMTPDataConnectorResult
	err := ctx.Invoke("azure-native:securityinsights/v20221201preview:getMTPDataConnector", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupMTPDataConnectorArgs struct {
	DataConnectorId   string `pulumi:"dataConnectorId"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	WorkspaceName     string `pulumi:"workspaceName"`
}


type LookupMTPDataConnectorResult struct {
	DataTypes  MTPDataConnectorDataTypesResponse `pulumi:"dataTypes"`
	Etag       *string                           `pulumi:"etag"`
	Id         string                            `pulumi:"id"`
	Kind       string                            `pulumi:"kind"`
	Name       string                            `pulumi:"name"`
	SystemData SystemDataResponse                `pulumi:"systemData"`
	TenantId   string                            `pulumi:"tenantId"`
	Type       string                            `pulumi:"type"`
}

func LookupMTPDataConnectorOutput(ctx *pulumi.Context, args LookupMTPDataConnectorOutputArgs, opts ...pulumi.InvokeOption) LookupMTPDataConnectorResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupMTPDataConnectorResult, error) {
			args := v.(LookupMTPDataConnectorArgs)
			r, err := LookupMTPDataConnector(ctx, &args, opts...)
			var s LookupMTPDataConnectorResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupMTPDataConnectorResultOutput)
}

type LookupMTPDataConnectorOutputArgs struct {
	DataConnectorId   pulumi.StringInput `pulumi:"dataConnectorId"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	WorkspaceName     pulumi.StringInput `pulumi:"workspaceName"`
}

func (LookupMTPDataConnectorOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupMTPDataConnectorArgs)(nil)).Elem()
}


type LookupMTPDataConnectorResultOutput struct{ *pulumi.OutputState }

func (LookupMTPDataConnectorResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupMTPDataConnectorResult)(nil)).Elem()
}

func (o LookupMTPDataConnectorResultOutput) ToLookupMTPDataConnectorResultOutput() LookupMTPDataConnectorResultOutput {
	return o
}

func (o LookupMTPDataConnectorResultOutput) ToLookupMTPDataConnectorResultOutputWithContext(ctx context.Context) LookupMTPDataConnectorResultOutput {
	return o
}

func (o LookupMTPDataConnectorResultOutput) DataTypes() MTPDataConnectorDataTypesResponseOutput {
	return o.ApplyT(func(v LookupMTPDataConnectorResult) MTPDataConnectorDataTypesResponse { return v.DataTypes }).(MTPDataConnectorDataTypesResponseOutput)
}

func (o LookupMTPDataConnectorResultOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMTPDataConnectorResult) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o LookupMTPDataConnectorResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMTPDataConnectorResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupMTPDataConnectorResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMTPDataConnectorResult) string { return v.Kind }).(pulumi.StringOutput)
}

func (o LookupMTPDataConnectorResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMTPDataConnectorResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupMTPDataConnectorResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupMTPDataConnectorResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupMTPDataConnectorResultOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMTPDataConnectorResult) string { return v.TenantId }).(pulumi.StringOutput)
}

func (o LookupMTPDataConnectorResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMTPDataConnectorResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupMTPDataConnectorResultOutput{})
}
