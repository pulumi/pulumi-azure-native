


package v20220131preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupIotConnector(ctx *pulumi.Context, args *LookupIotConnectorArgs, opts ...pulumi.InvokeOption) (*LookupIotConnectorResult, error) {
	var rv LookupIotConnectorResult
	err := ctx.Invoke("azure-native:healthcareapis/v20220131preview:getIotConnector", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupIotConnectorArgs struct {
	IotConnectorName  string `pulumi:"iotConnectorName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	WorkspaceName     string `pulumi:"workspaceName"`
}


type LookupIotConnectorResult struct {
	DeviceMapping                  *IotMappingPropertiesResponse                      `pulumi:"deviceMapping"`
	Etag                           *string                                            `pulumi:"etag"`
	Id                             string                                             `pulumi:"id"`
	Identity                       *ServiceManagedIdentityResponseIdentity            `pulumi:"identity"`
	IngestionEndpointConfiguration *IotEventHubIngestionEndpointConfigurationResponse `pulumi:"ingestionEndpointConfiguration"`
	Location                       *string                                            `pulumi:"location"`
	Name                           string                                             `pulumi:"name"`
	ProvisioningState              string                                             `pulumi:"provisioningState"`
	SystemData                     SystemDataResponse                                 `pulumi:"systemData"`
	Tags                           map[string]string                                  `pulumi:"tags"`
	Type                           string                                             `pulumi:"type"`
}

func LookupIotConnectorOutput(ctx *pulumi.Context, args LookupIotConnectorOutputArgs, opts ...pulumi.InvokeOption) LookupIotConnectorResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupIotConnectorResult, error) {
			args := v.(LookupIotConnectorArgs)
			r, err := LookupIotConnector(ctx, &args, opts...)
			var s LookupIotConnectorResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupIotConnectorResultOutput)
}

type LookupIotConnectorOutputArgs struct {
	IotConnectorName  pulumi.StringInput `pulumi:"iotConnectorName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	WorkspaceName     pulumi.StringInput `pulumi:"workspaceName"`
}

func (LookupIotConnectorOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupIotConnectorArgs)(nil)).Elem()
}


type LookupIotConnectorResultOutput struct{ *pulumi.OutputState }

func (LookupIotConnectorResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupIotConnectorResult)(nil)).Elem()
}

func (o LookupIotConnectorResultOutput) ToLookupIotConnectorResultOutput() LookupIotConnectorResultOutput {
	return o
}

func (o LookupIotConnectorResultOutput) ToLookupIotConnectorResultOutputWithContext(ctx context.Context) LookupIotConnectorResultOutput {
	return o
}

func (o LookupIotConnectorResultOutput) DeviceMapping() IotMappingPropertiesResponsePtrOutput {
	return o.ApplyT(func(v LookupIotConnectorResult) *IotMappingPropertiesResponse { return v.DeviceMapping }).(IotMappingPropertiesResponsePtrOutput)
}

func (o LookupIotConnectorResultOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupIotConnectorResult) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o LookupIotConnectorResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIotConnectorResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupIotConnectorResultOutput) Identity() ServiceManagedIdentityResponseIdentityPtrOutput {
	return o.ApplyT(func(v LookupIotConnectorResult) *ServiceManagedIdentityResponseIdentity { return v.Identity }).(ServiceManagedIdentityResponseIdentityPtrOutput)
}

func (o LookupIotConnectorResultOutput) IngestionEndpointConfiguration() IotEventHubIngestionEndpointConfigurationResponsePtrOutput {
	return o.ApplyT(func(v LookupIotConnectorResult) *IotEventHubIngestionEndpointConfigurationResponse {
		return v.IngestionEndpointConfiguration
	}).(IotEventHubIngestionEndpointConfigurationResponsePtrOutput)
}

func (o LookupIotConnectorResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupIotConnectorResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupIotConnectorResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIotConnectorResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupIotConnectorResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIotConnectorResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupIotConnectorResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupIotConnectorResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupIotConnectorResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupIotConnectorResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupIotConnectorResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIotConnectorResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupIotConnectorResultOutput{})
}
