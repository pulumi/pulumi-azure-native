


package v20210601preview

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
			Type: pulumi.String("azure-nextgen:healthcareapis/v20210601preview:IotConnector"),
		},
		{
			Type: pulumi.String("azure-native:healthcareapis:IotConnector"),
		},
		{
			Type: pulumi.String("azure-nextgen:healthcareapis:IotConnector"),
		},
	})
	opts = append(opts, aliases)
	var resource IotConnector
	err := ctx.RegisterResource("azure-native:healthcareapis/v20210601preview:IotConnector", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetIotConnector(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IotConnectorState, opts ...pulumi.ResourceOption) (*IotConnector, error) {
	var resource IotConnector
	err := ctx.ReadResource("azure-native:healthcareapis/v20210601preview:IotConnector", name, id, state, &resource, opts...)
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
	Etag                           *string                                    `pulumi:"etag"`
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
	Etag                           pulumi.StringPtrInput
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
	return reflect.TypeOf((*IotConnector)(nil))
}

func (i *IotConnector) ToIotConnectorOutput() IotConnectorOutput {
	return i.ToIotConnectorOutputWithContext(context.Background())
}

func (i *IotConnector) ToIotConnectorOutputWithContext(ctx context.Context) IotConnectorOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IotConnectorOutput)
}

type IotConnectorOutput struct{ *pulumi.OutputState }

func (IotConnectorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IotConnector)(nil))
}

func (o IotConnectorOutput) ToIotConnectorOutput() IotConnectorOutput {
	return o
}

func (o IotConnectorOutput) ToIotConnectorOutputWithContext(ctx context.Context) IotConnectorOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*IotConnectorInput)(nil)).Elem(), &IotConnector{})
	pulumi.RegisterOutputType(IotConnectorOutput{})
}
