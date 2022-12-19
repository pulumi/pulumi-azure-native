


package v20221001preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupAnalyticsConnector(ctx *pulumi.Context, args *LookupAnalyticsConnectorArgs, opts ...pulumi.InvokeOption) (*LookupAnalyticsConnectorResult, error) {
	var rv LookupAnalyticsConnectorResult
	err := ctx.Invoke("azure-native:healthcareapis/v20221001preview:getAnalyticsConnector", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupAnalyticsConnectorArgs struct {
	AnalyticsConnectorName string `pulumi:"analyticsConnectorName"`
	ResourceGroupName      string `pulumi:"resourceGroupName"`
	WorkspaceName          string `pulumi:"workspaceName"`
}


type LookupAnalyticsConnectorResult struct {
	DataDestinationConfiguration AnalyticsConnectorDataLakeDataDestinationResponse `pulumi:"dataDestinationConfiguration"`
	DataMappingConfiguration     AnalyticsConnectorFhirToParquetMappingResponse    `pulumi:"dataMappingConfiguration"`
	DataSourceConfiguration      AnalyticsConnectorFhirServiceDataSourceResponse   `pulumi:"dataSourceConfiguration"`
	Etag                         *string                                           `pulumi:"etag"`
	Id                           string                                            `pulumi:"id"`
	Identity                     *ServiceManagedIdentityResponseIdentity           `pulumi:"identity"`
	Location                     *string                                           `pulumi:"location"`
	Name                         string                                            `pulumi:"name"`
	ProvisioningState            string                                            `pulumi:"provisioningState"`
	SystemData                   SystemDataResponse                                `pulumi:"systemData"`
	Tags                         map[string]string                                 `pulumi:"tags"`
	Type                         string                                            `pulumi:"type"`
}

func LookupAnalyticsConnectorOutput(ctx *pulumi.Context, args LookupAnalyticsConnectorOutputArgs, opts ...pulumi.InvokeOption) LookupAnalyticsConnectorResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupAnalyticsConnectorResult, error) {
			args := v.(LookupAnalyticsConnectorArgs)
			r, err := LookupAnalyticsConnector(ctx, &args, opts...)
			var s LookupAnalyticsConnectorResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupAnalyticsConnectorResultOutput)
}

type LookupAnalyticsConnectorOutputArgs struct {
	AnalyticsConnectorName pulumi.StringInput `pulumi:"analyticsConnectorName"`
	ResourceGroupName      pulumi.StringInput `pulumi:"resourceGroupName"`
	WorkspaceName          pulumi.StringInput `pulumi:"workspaceName"`
}

func (LookupAnalyticsConnectorOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAnalyticsConnectorArgs)(nil)).Elem()
}


type LookupAnalyticsConnectorResultOutput struct{ *pulumi.OutputState }

func (LookupAnalyticsConnectorResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAnalyticsConnectorResult)(nil)).Elem()
}

func (o LookupAnalyticsConnectorResultOutput) ToLookupAnalyticsConnectorResultOutput() LookupAnalyticsConnectorResultOutput {
	return o
}

func (o LookupAnalyticsConnectorResultOutput) ToLookupAnalyticsConnectorResultOutputWithContext(ctx context.Context) LookupAnalyticsConnectorResultOutput {
	return o
}

func (o LookupAnalyticsConnectorResultOutput) DataDestinationConfiguration() AnalyticsConnectorDataLakeDataDestinationResponseOutput {
	return o.ApplyT(func(v LookupAnalyticsConnectorResult) AnalyticsConnectorDataLakeDataDestinationResponse {
		return v.DataDestinationConfiguration
	}).(AnalyticsConnectorDataLakeDataDestinationResponseOutput)
}

func (o LookupAnalyticsConnectorResultOutput) DataMappingConfiguration() AnalyticsConnectorFhirToParquetMappingResponseOutput {
	return o.ApplyT(func(v LookupAnalyticsConnectorResult) AnalyticsConnectorFhirToParquetMappingResponse {
		return v.DataMappingConfiguration
	}).(AnalyticsConnectorFhirToParquetMappingResponseOutput)
}

func (o LookupAnalyticsConnectorResultOutput) DataSourceConfiguration() AnalyticsConnectorFhirServiceDataSourceResponseOutput {
	return o.ApplyT(func(v LookupAnalyticsConnectorResult) AnalyticsConnectorFhirServiceDataSourceResponse {
		return v.DataSourceConfiguration
	}).(AnalyticsConnectorFhirServiceDataSourceResponseOutput)
}

func (o LookupAnalyticsConnectorResultOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAnalyticsConnectorResult) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o LookupAnalyticsConnectorResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAnalyticsConnectorResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupAnalyticsConnectorResultOutput) Identity() ServiceManagedIdentityResponseIdentityPtrOutput {
	return o.ApplyT(func(v LookupAnalyticsConnectorResult) *ServiceManagedIdentityResponseIdentity { return v.Identity }).(ServiceManagedIdentityResponseIdentityPtrOutput)
}

func (o LookupAnalyticsConnectorResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAnalyticsConnectorResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupAnalyticsConnectorResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAnalyticsConnectorResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupAnalyticsConnectorResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAnalyticsConnectorResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupAnalyticsConnectorResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupAnalyticsConnectorResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupAnalyticsConnectorResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupAnalyticsConnectorResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupAnalyticsConnectorResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAnalyticsConnectorResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupAnalyticsConnectorResultOutput{})
}
