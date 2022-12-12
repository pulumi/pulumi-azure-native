


package v20170426

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupConnectorMapping(ctx *pulumi.Context, args *LookupConnectorMappingArgs, opts ...pulumi.InvokeOption) (*LookupConnectorMappingResult, error) {
	var rv LookupConnectorMappingResult
	err := ctx.Invoke("azure-native:customerinsights/v20170426:getConnectorMapping", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupConnectorMappingArgs struct {
	ConnectorName     string `pulumi:"connectorName"`
	HubName           string `pulumi:"hubName"`
	MappingName       string `pulumi:"mappingName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupConnectorMappingResult struct {
	ConnectorMappingName string                             `pulumi:"connectorMappingName"`
	ConnectorName        string                             `pulumi:"connectorName"`
	ConnectorType        *string                            `pulumi:"connectorType"`
	Created              string                             `pulumi:"created"`
	DataFormatId         string                             `pulumi:"dataFormatId"`
	Description          *string                            `pulumi:"description"`
	DisplayName          *string                            `pulumi:"displayName"`
	EntityType           string                             `pulumi:"entityType"`
	EntityTypeName       string                             `pulumi:"entityTypeName"`
	Id                   string                             `pulumi:"id"`
	LastModified         string                             `pulumi:"lastModified"`
	MappingProperties    ConnectorMappingPropertiesResponse `pulumi:"mappingProperties"`
	Name                 string                             `pulumi:"name"`
	NextRunTime          string                             `pulumi:"nextRunTime"`
	RunId                string                             `pulumi:"runId"`
	State                string                             `pulumi:"state"`
	TenantId             string                             `pulumi:"tenantId"`
	Type                 string                             `pulumi:"type"`
}

func LookupConnectorMappingOutput(ctx *pulumi.Context, args LookupConnectorMappingOutputArgs, opts ...pulumi.InvokeOption) LookupConnectorMappingResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupConnectorMappingResult, error) {
			args := v.(LookupConnectorMappingArgs)
			r, err := LookupConnectorMapping(ctx, &args, opts...)
			var s LookupConnectorMappingResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupConnectorMappingResultOutput)
}

type LookupConnectorMappingOutputArgs struct {
	ConnectorName     pulumi.StringInput `pulumi:"connectorName"`
	HubName           pulumi.StringInput `pulumi:"hubName"`
	MappingName       pulumi.StringInput `pulumi:"mappingName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupConnectorMappingOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupConnectorMappingArgs)(nil)).Elem()
}


type LookupConnectorMappingResultOutput struct{ *pulumi.OutputState }

func (LookupConnectorMappingResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupConnectorMappingResult)(nil)).Elem()
}

func (o LookupConnectorMappingResultOutput) ToLookupConnectorMappingResultOutput() LookupConnectorMappingResultOutput {
	return o
}

func (o LookupConnectorMappingResultOutput) ToLookupConnectorMappingResultOutputWithContext(ctx context.Context) LookupConnectorMappingResultOutput {
	return o
}

func (o LookupConnectorMappingResultOutput) ConnectorMappingName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConnectorMappingResult) string { return v.ConnectorMappingName }).(pulumi.StringOutput)
}

func (o LookupConnectorMappingResultOutput) ConnectorName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConnectorMappingResult) string { return v.ConnectorName }).(pulumi.StringOutput)
}

func (o LookupConnectorMappingResultOutput) ConnectorType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupConnectorMappingResult) *string { return v.ConnectorType }).(pulumi.StringPtrOutput)
}

func (o LookupConnectorMappingResultOutput) Created() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConnectorMappingResult) string { return v.Created }).(pulumi.StringOutput)
}

func (o LookupConnectorMappingResultOutput) DataFormatId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConnectorMappingResult) string { return v.DataFormatId }).(pulumi.StringOutput)
}

func (o LookupConnectorMappingResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupConnectorMappingResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupConnectorMappingResultOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupConnectorMappingResult) *string { return v.DisplayName }).(pulumi.StringPtrOutput)
}

func (o LookupConnectorMappingResultOutput) EntityType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConnectorMappingResult) string { return v.EntityType }).(pulumi.StringOutput)
}

func (o LookupConnectorMappingResultOutput) EntityTypeName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConnectorMappingResult) string { return v.EntityTypeName }).(pulumi.StringOutput)
}

func (o LookupConnectorMappingResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConnectorMappingResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupConnectorMappingResultOutput) LastModified() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConnectorMappingResult) string { return v.LastModified }).(pulumi.StringOutput)
}

func (o LookupConnectorMappingResultOutput) MappingProperties() ConnectorMappingPropertiesResponseOutput {
	return o.ApplyT(func(v LookupConnectorMappingResult) ConnectorMappingPropertiesResponse { return v.MappingProperties }).(ConnectorMappingPropertiesResponseOutput)
}

func (o LookupConnectorMappingResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConnectorMappingResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupConnectorMappingResultOutput) NextRunTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConnectorMappingResult) string { return v.NextRunTime }).(pulumi.StringOutput)
}

func (o LookupConnectorMappingResultOutput) RunId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConnectorMappingResult) string { return v.RunId }).(pulumi.StringOutput)
}

func (o LookupConnectorMappingResultOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConnectorMappingResult) string { return v.State }).(pulumi.StringOutput)
}

func (o LookupConnectorMappingResultOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConnectorMappingResult) string { return v.TenantId }).(pulumi.StringOutput)
}

func (o LookupConnectorMappingResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConnectorMappingResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupConnectorMappingResultOutput{})
}
