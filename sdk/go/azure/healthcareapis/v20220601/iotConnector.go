


package v20220601

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type IotConnector struct {
	pulumi.CustomResourceState

	DeviceMapping                  IotMappingPropertiesResponsePtrOutput                      `pulumi:"deviceMapping"`
	Etag                           pulumi.StringPtrOutput                                     `pulumi:"etag"`
	Identity                       ServiceManagedIdentityResponseIdentityPtrOutput            `pulumi:"identity"`
	IngestionEndpointConfiguration IotEventHubIngestionEndpointConfigurationResponsePtrOutput `pulumi:"ingestionEndpointConfiguration"`
	Location                       pulumi.StringPtrOutput                                     `pulumi:"location"`
	Name                           pulumi.StringOutput                                        `pulumi:"name"`
	ProvisioningState              pulumi.StringOutput                                        `pulumi:"provisioningState"`
	SystemData                     SystemDataResponseOutput                                   `pulumi:"systemData"`
	Tags                           pulumi.StringMapOutput                                     `pulumi:"tags"`
	Type                           pulumi.StringOutput                                        `pulumi:"type"`
}


func NewIotConnector(ctx *pulumi.Context,
	name string, args *IotConnectorArgs, opts ...pulumi.ResourceOption) (*IotConnector, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.WorkspaceName == nil {
		return nil, errors.New("invalid value for required argument 'WorkspaceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:healthcareapis:IotConnector"),
		},
		{
			Type: pulumi.String("azure-native:healthcareapis/v20210601preview:IotConnector"),
		},
		{
			Type: pulumi.String("azure-native:healthcareapis/v20211101:IotConnector"),
		},
		{
			Type: pulumi.String("azure-native:healthcareapis/v20220131preview:IotConnector"),
		},
		{
			Type: pulumi.String("azure-native:healthcareapis/v20220515:IotConnector"),
		},
		{
			Type: pulumi.String("azure-native:healthcareapis/v20221001preview:IotConnector"),
		},
	})
	opts = append(opts, aliases)
	var resource IotConnector
	err := ctx.RegisterResource("azure-native:healthcareapis/v20220601:IotConnector", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetIotConnector(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IotConnectorState, opts ...pulumi.ResourceOption) (*IotConnector, error) {
	var resource IotConnector
	err := ctx.ReadResource("azure-native:healthcareapis/v20220601:IotConnector", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type iotConnectorState struct {
}

type IotConnectorState struct {
}

func (IotConnectorState) ElementType() reflect.Type {
	return reflect.TypeOf((*iotConnectorState)(nil)).Elem()
}

type iotConnectorArgs struct {
	DeviceMapping                  *IotMappingProperties                      `pulumi:"deviceMapping"`
	Identity                       *ServiceManagedIdentityIdentity            `pulumi:"identity"`
	IngestionEndpointConfiguration *IotEventHubIngestionEndpointConfiguration `pulumi:"ingestionEndpointConfiguration"`
	IotConnectorName               *string                                    `pulumi:"iotConnectorName"`
	Location                       *string                                    `pulumi:"location"`
	ResourceGroupName              string                                     `pulumi:"resourceGroupName"`
	Tags                           map[string]string                          `pulumi:"tags"`
	WorkspaceName                  string                                     `pulumi:"workspaceName"`
}


type IotConnectorArgs struct {
	DeviceMapping                  IotMappingPropertiesPtrInput
	Identity                       ServiceManagedIdentityIdentityPtrInput
	IngestionEndpointConfiguration IotEventHubIngestionEndpointConfigurationPtrInput
	IotConnectorName               pulumi.StringPtrInput
	Location                       pulumi.StringPtrInput
	ResourceGroupName              pulumi.StringInput
	Tags                           pulumi.StringMapInput
	WorkspaceName                  pulumi.StringInput
}

func (IotConnectorArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*iotConnectorArgs)(nil)).Elem()
}

type IotConnectorInput interface {
	pulumi.Input

	ToIotConnectorOutput() IotConnectorOutput
	ToIotConnectorOutputWithContext(ctx context.Context) IotConnectorOutput
}

func (*IotConnector) ElementType() reflect.Type {
	return reflect.TypeOf((**IotConnector)(nil)).Elem()
}

func (i *IotConnector) ToIotConnectorOutput() IotConnectorOutput {
	return i.ToIotConnectorOutputWithContext(context.Background())
}

func (i *IotConnector) ToIotConnectorOutputWithContext(ctx context.Context) IotConnectorOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IotConnectorOutput)
}

type IotConnectorOutput struct{ *pulumi.OutputState }

func (IotConnectorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IotConnector)(nil)).Elem()
}

func (o IotConnectorOutput) ToIotConnectorOutput() IotConnectorOutput {
	return o
}

func (o IotConnectorOutput) ToIotConnectorOutputWithContext(ctx context.Context) IotConnectorOutput {
	return o
}

func (o IotConnectorOutput) DeviceMapping() IotMappingPropertiesResponsePtrOutput {
	return o.ApplyT(func(v *IotConnector) IotMappingPropertiesResponsePtrOutput { return v.DeviceMapping }).(IotMappingPropertiesResponsePtrOutput)
}

func (o IotConnectorOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IotConnector) pulumi.StringPtrOutput { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o IotConnectorOutput) Identity() ServiceManagedIdentityResponseIdentityPtrOutput {
	return o.ApplyT(func(v *IotConnector) ServiceManagedIdentityResponseIdentityPtrOutput { return v.Identity }).(ServiceManagedIdentityResponseIdentityPtrOutput)
}

func (o IotConnectorOutput) IngestionEndpointConfiguration() IotEventHubIngestionEndpointConfigurationResponsePtrOutput {
	return o.ApplyT(func(v *IotConnector) IotEventHubIngestionEndpointConfigurationResponsePtrOutput {
		return v.IngestionEndpointConfiguration
	}).(IotEventHubIngestionEndpointConfigurationResponsePtrOutput)
}

func (o IotConnectorOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IotConnector) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

func (o IotConnectorOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *IotConnector) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o IotConnectorOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *IotConnector) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o IotConnectorOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *IotConnector) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o IotConnectorOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *IotConnector) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o IotConnectorOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *IotConnector) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(IotConnectorOutput{})
}
