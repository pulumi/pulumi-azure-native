


package v20220901preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupOffice365ProjectDataConnector(ctx *pulumi.Context, args *LookupOffice365ProjectDataConnectorArgs, opts ...pulumi.InvokeOption) (*LookupOffice365ProjectDataConnectorResult, error) {
	var rv LookupOffice365ProjectDataConnectorResult
	err := ctx.Invoke("azure-native:securityinsights/v20220901preview:getOffice365ProjectDataConnector", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupOffice365ProjectDataConnectorArgs struct {
	DataConnectorId   string `pulumi:"dataConnectorId"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	WorkspaceName     string `pulumi:"workspaceName"`
}


type LookupOffice365ProjectDataConnectorResult struct {
	DataTypes  Office365ProjectConnectorDataTypesResponse `pulumi:"dataTypes"`
	Etag       *string                                    `pulumi:"etag"`
	Id         string                                     `pulumi:"id"`
	Kind       string                                     `pulumi:"kind"`
	Name       string                                     `pulumi:"name"`
	SystemData SystemDataResponse                         `pulumi:"systemData"`
	TenantId   string                                     `pulumi:"tenantId"`
	Type       string                                     `pulumi:"type"`
}

func LookupOffice365ProjectDataConnectorOutput(ctx *pulumi.Context, args LookupOffice365ProjectDataConnectorOutputArgs, opts ...pulumi.InvokeOption) LookupOffice365ProjectDataConnectorResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupOffice365ProjectDataConnectorResult, error) {
			args := v.(LookupOffice365ProjectDataConnectorArgs)
			r, err := LookupOffice365ProjectDataConnector(ctx, &args, opts...)
			var s LookupOffice365ProjectDataConnectorResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupOffice365ProjectDataConnectorResultOutput)
}

type LookupOffice365ProjectDataConnectorOutputArgs struct {
	DataConnectorId   pulumi.StringInput `pulumi:"dataConnectorId"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	WorkspaceName     pulumi.StringInput `pulumi:"workspaceName"`
}

func (LookupOffice365ProjectDataConnectorOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupOffice365ProjectDataConnectorArgs)(nil)).Elem()
}


type LookupOffice365ProjectDataConnectorResultOutput struct{ *pulumi.OutputState }

func (LookupOffice365ProjectDataConnectorResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupOffice365ProjectDataConnectorResult)(nil)).Elem()
}

func (o LookupOffice365ProjectDataConnectorResultOutput) ToLookupOffice365ProjectDataConnectorResultOutput() LookupOffice365ProjectDataConnectorResultOutput {
	return o
}

func (o LookupOffice365ProjectDataConnectorResultOutput) ToLookupOffice365ProjectDataConnectorResultOutputWithContext(ctx context.Context) LookupOffice365ProjectDataConnectorResultOutput {
	return o
}

func (o LookupOffice365ProjectDataConnectorResultOutput) DataTypes() Office365ProjectConnectorDataTypesResponseOutput {
	return o.ApplyT(func(v LookupOffice365ProjectDataConnectorResult) Office365ProjectConnectorDataTypesResponse {
		return v.DataTypes
	}).(Office365ProjectConnectorDataTypesResponseOutput)
}

func (o LookupOffice365ProjectDataConnectorResultOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupOffice365ProjectDataConnectorResult) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o LookupOffice365ProjectDataConnectorResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOffice365ProjectDataConnectorResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupOffice365ProjectDataConnectorResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOffice365ProjectDataConnectorResult) string { return v.Kind }).(pulumi.StringOutput)
}

func (o LookupOffice365ProjectDataConnectorResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOffice365ProjectDataConnectorResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupOffice365ProjectDataConnectorResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupOffice365ProjectDataConnectorResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupOffice365ProjectDataConnectorResultOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOffice365ProjectDataConnectorResult) string { return v.TenantId }).(pulumi.StringOutput)
}

func (o LookupOffice365ProjectDataConnectorResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOffice365ProjectDataConnectorResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupOffice365ProjectDataConnectorResultOutput{})
}
