


package customerinsights

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupConnector(ctx *pulumi.Context, args *LookupConnectorArgs, opts ...pulumi.InvokeOption) (*LookupConnectorResult, error) {
	var rv LookupConnectorResult
	err := ctx.Invoke("azure-native:customerinsights:getConnector", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupConnectorArgs struct {
	ConnectorName     string `pulumi:"connectorName"`
	HubName           string `pulumi:"hubName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupConnectorResult struct {
	ConnectorId         int                    `pulumi:"connectorId"`
	ConnectorName       *string                `pulumi:"connectorName"`
	ConnectorProperties map[string]interface{} `pulumi:"connectorProperties"`
	ConnectorType       string                 `pulumi:"connectorType"`
	Created             string                 `pulumi:"created"`
	Description         *string                `pulumi:"description"`
	DisplayName         *string                `pulumi:"displayName"`
	Id                  string                 `pulumi:"id"`
	IsInternal          *bool                  `pulumi:"isInternal"`
	LastModified        string                 `pulumi:"lastModified"`
	Name                string                 `pulumi:"name"`
	State               string                 `pulumi:"state"`
	TenantId            string                 `pulumi:"tenantId"`
	Type                string                 `pulumi:"type"`
}

func LookupConnectorOutput(ctx *pulumi.Context, args LookupConnectorOutputArgs, opts ...pulumi.InvokeOption) LookupConnectorResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupConnectorResult, error) {
			args := v.(LookupConnectorArgs)
			r, err := LookupConnector(ctx, &args, opts...)
			var s LookupConnectorResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupConnectorResultOutput)
}

type LookupConnectorOutputArgs struct {
	ConnectorName     pulumi.StringInput `pulumi:"connectorName"`
	HubName           pulumi.StringInput `pulumi:"hubName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupConnectorOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupConnectorArgs)(nil)).Elem()
}


type LookupConnectorResultOutput struct{ *pulumi.OutputState }

func (LookupConnectorResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupConnectorResult)(nil)).Elem()
}

func (o LookupConnectorResultOutput) ToLookupConnectorResultOutput() LookupConnectorResultOutput {
	return o
}

func (o LookupConnectorResultOutput) ToLookupConnectorResultOutputWithContext(ctx context.Context) LookupConnectorResultOutput {
	return o
}

func (o LookupConnectorResultOutput) ConnectorId() pulumi.IntOutput {
	return o.ApplyT(func(v LookupConnectorResult) int { return v.ConnectorId }).(pulumi.IntOutput)
}

func (o LookupConnectorResultOutput) ConnectorName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupConnectorResult) *string { return v.ConnectorName }).(pulumi.StringPtrOutput)
}

func (o LookupConnectorResultOutput) ConnectorProperties() pulumi.MapOutput {
	return o.ApplyT(func(v LookupConnectorResult) map[string]interface{} { return v.ConnectorProperties }).(pulumi.MapOutput)
}

func (o LookupConnectorResultOutput) ConnectorType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConnectorResult) string { return v.ConnectorType }).(pulumi.StringOutput)
}

func (o LookupConnectorResultOutput) Created() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConnectorResult) string { return v.Created }).(pulumi.StringOutput)
}

func (o LookupConnectorResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupConnectorResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupConnectorResultOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupConnectorResult) *string { return v.DisplayName }).(pulumi.StringPtrOutput)
}

func (o LookupConnectorResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConnectorResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupConnectorResultOutput) IsInternal() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupConnectorResult) *bool { return v.IsInternal }).(pulumi.BoolPtrOutput)
}

func (o LookupConnectorResultOutput) LastModified() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConnectorResult) string { return v.LastModified }).(pulumi.StringOutput)
}

func (o LookupConnectorResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConnectorResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupConnectorResultOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConnectorResult) string { return v.State }).(pulumi.StringOutput)
}

func (o LookupConnectorResultOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConnectorResult) string { return v.TenantId }).(pulumi.StringOutput)
}

func (o LookupConnectorResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConnectorResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupConnectorResultOutput{})
}
