


package securitydevops

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupAzureDevOpsConnector(ctx *pulumi.Context, args *LookupAzureDevOpsConnectorArgs, opts ...pulumi.InvokeOption) (*LookupAzureDevOpsConnectorResult, error) {
	var rv LookupAzureDevOpsConnectorResult
	err := ctx.Invoke("azure-native:securitydevops:getAzureDevOpsConnector", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupAzureDevOpsConnectorArgs struct {
	AzureDevOpsConnectorName string `pulumi:"azureDevOpsConnectorName"`
	ResourceGroupName        string `pulumi:"resourceGroupName"`
}

type LookupAzureDevOpsConnectorResult struct {
	Id         string                                 `pulumi:"id"`
	Location   string                                 `pulumi:"location"`
	Name       string                                 `pulumi:"name"`
	Properties AzureDevOpsConnectorPropertiesResponse `pulumi:"properties"`
	SystemData SystemDataResponse                     `pulumi:"systemData"`
	Tags       map[string]string                      `pulumi:"tags"`
	Type       string                                 `pulumi:"type"`
}

func LookupAzureDevOpsConnectorOutput(ctx *pulumi.Context, args LookupAzureDevOpsConnectorOutputArgs, opts ...pulumi.InvokeOption) LookupAzureDevOpsConnectorResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupAzureDevOpsConnectorResult, error) {
			args := v.(LookupAzureDevOpsConnectorArgs)
			r, err := LookupAzureDevOpsConnector(ctx, &args, opts...)
			var s LookupAzureDevOpsConnectorResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupAzureDevOpsConnectorResultOutput)
}

type LookupAzureDevOpsConnectorOutputArgs struct {
	AzureDevOpsConnectorName pulumi.StringInput `pulumi:"azureDevOpsConnectorName"`
	ResourceGroupName        pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupAzureDevOpsConnectorOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAzureDevOpsConnectorArgs)(nil)).Elem()
}

type LookupAzureDevOpsConnectorResultOutput struct{ *pulumi.OutputState }

func (LookupAzureDevOpsConnectorResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAzureDevOpsConnectorResult)(nil)).Elem()
}

func (o LookupAzureDevOpsConnectorResultOutput) ToLookupAzureDevOpsConnectorResultOutput() LookupAzureDevOpsConnectorResultOutput {
	return o
}

func (o LookupAzureDevOpsConnectorResultOutput) ToLookupAzureDevOpsConnectorResultOutputWithContext(ctx context.Context) LookupAzureDevOpsConnectorResultOutput {
	return o
}

func (o LookupAzureDevOpsConnectorResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAzureDevOpsConnectorResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupAzureDevOpsConnectorResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAzureDevOpsConnectorResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupAzureDevOpsConnectorResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAzureDevOpsConnectorResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupAzureDevOpsConnectorResultOutput) Properties() AzureDevOpsConnectorPropertiesResponseOutput {
	return o.ApplyT(func(v LookupAzureDevOpsConnectorResult) AzureDevOpsConnectorPropertiesResponse { return v.Properties }).(AzureDevOpsConnectorPropertiesResponseOutput)
}

func (o LookupAzureDevOpsConnectorResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupAzureDevOpsConnectorResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupAzureDevOpsConnectorResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupAzureDevOpsConnectorResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupAzureDevOpsConnectorResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAzureDevOpsConnectorResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupAzureDevOpsConnectorResultOutput{})
}
