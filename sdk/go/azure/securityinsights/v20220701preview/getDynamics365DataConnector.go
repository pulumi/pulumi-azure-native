


package v20220701preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupDynamics365DataConnector(ctx *pulumi.Context, args *LookupDynamics365DataConnectorArgs, opts ...pulumi.InvokeOption) (*LookupDynamics365DataConnectorResult, error) {
	var rv LookupDynamics365DataConnectorResult
	err := ctx.Invoke("azure-native:securityinsights/v20220701preview:getDynamics365DataConnector", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDynamics365DataConnectorArgs struct {
	DataConnectorId   string `pulumi:"dataConnectorId"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	WorkspaceName     string `pulumi:"workspaceName"`
}


type LookupDynamics365DataConnectorResult struct {
	DataTypes  Dynamics365DataConnectorDataTypesResponse `pulumi:"dataTypes"`
	Etag       *string                                   `pulumi:"etag"`
	Id         string                                    `pulumi:"id"`
	Kind       string                                    `pulumi:"kind"`
	Name       string                                    `pulumi:"name"`
	SystemData SystemDataResponse                        `pulumi:"systemData"`
	TenantId   string                                    `pulumi:"tenantId"`
	Type       string                                    `pulumi:"type"`
}

func LookupDynamics365DataConnectorOutput(ctx *pulumi.Context, args LookupDynamics365DataConnectorOutputArgs, opts ...pulumi.InvokeOption) LookupDynamics365DataConnectorResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDynamics365DataConnectorResult, error) {
			args := v.(LookupDynamics365DataConnectorArgs)
			r, err := LookupDynamics365DataConnector(ctx, &args, opts...)
			var s LookupDynamics365DataConnectorResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupDynamics365DataConnectorResultOutput)
}

type LookupDynamics365DataConnectorOutputArgs struct {
	DataConnectorId   pulumi.StringInput `pulumi:"dataConnectorId"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	WorkspaceName     pulumi.StringInput `pulumi:"workspaceName"`
}

func (LookupDynamics365DataConnectorOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDynamics365DataConnectorArgs)(nil)).Elem()
}


type LookupDynamics365DataConnectorResultOutput struct{ *pulumi.OutputState }

func (LookupDynamics365DataConnectorResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDynamics365DataConnectorResult)(nil)).Elem()
}

func (o LookupDynamics365DataConnectorResultOutput) ToLookupDynamics365DataConnectorResultOutput() LookupDynamics365DataConnectorResultOutput {
	return o
}

func (o LookupDynamics365DataConnectorResultOutput) ToLookupDynamics365DataConnectorResultOutputWithContext(ctx context.Context) LookupDynamics365DataConnectorResultOutput {
	return o
}

func (o LookupDynamics365DataConnectorResultOutput) DataTypes() Dynamics365DataConnectorDataTypesResponseOutput {
	return o.ApplyT(func(v LookupDynamics365DataConnectorResult) Dynamics365DataConnectorDataTypesResponse {
		return v.DataTypes
	}).(Dynamics365DataConnectorDataTypesResponseOutput)
}

func (o LookupDynamics365DataConnectorResultOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDynamics365DataConnectorResult) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o LookupDynamics365DataConnectorResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDynamics365DataConnectorResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupDynamics365DataConnectorResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDynamics365DataConnectorResult) string { return v.Kind }).(pulumi.StringOutput)
}

func (o LookupDynamics365DataConnectorResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDynamics365DataConnectorResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupDynamics365DataConnectorResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupDynamics365DataConnectorResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupDynamics365DataConnectorResultOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDynamics365DataConnectorResult) string { return v.TenantId }).(pulumi.StringOutput)
}

func (o LookupDynamics365DataConnectorResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDynamics365DataConnectorResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDynamics365DataConnectorResultOutput{})
}
