


package insights

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type DataCollectionEndpoint struct {
	pulumi.CustomResourceState

	ConfigurationAccess DataCollectionEndpointResponseConfigurationAccessPtrOutput `pulumi:"configurationAccess"`
	Description         pulumi.StringPtrOutput                                     `pulumi:"description"`
	Etag                pulumi.StringOutput                                        `pulumi:"etag"`
	ImmutableId         pulumi.StringPtrOutput                                     `pulumi:"immutableId"`
	Kind                pulumi.StringPtrOutput                                     `pulumi:"kind"`
	Location            pulumi.StringOutput                                        `pulumi:"location"`
	LogsIngestion       DataCollectionEndpointResponseLogsIngestionPtrOutput       `pulumi:"logsIngestion"`
	Name                pulumi.StringOutput                                        `pulumi:"name"`
	NetworkAcls         DataCollectionEndpointResponseNetworkAclsPtrOutput         `pulumi:"networkAcls"`
	ProvisioningState   pulumi.StringOutput                                        `pulumi:"provisioningState"`
	SystemData          DataCollectionEndpointResourceResponseSystemDataOutput     `pulumi:"systemData"`
	Tags                pulumi.StringMapOutput                                     `pulumi:"tags"`
	Type                pulumi.StringOutput                                        `pulumi:"type"`
}


func NewDataCollectionEndpoint(ctx *pulumi.Context,
	name string, args *DataCollectionEndpointArgs, opts ...pulumi.ResourceOption) (*DataCollectionEndpoint, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:insights:DataCollectionEndpoint"),
		},
		{
			Type: pulumi.String("azure-native:insights/v20210401:DataCollectionEndpoint"),
		},
		{
			Type: pulumi.String("azure-nextgen:insights/v20210401:DataCollectionEndpoint"),
		},
	})
	opts = append(opts, aliases)
	var resource DataCollectionEndpoint
	err := ctx.RegisterResource("azure-native:insights:DataCollectionEndpoint", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetDataCollectionEndpoint(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DataCollectionEndpointState, opts ...pulumi.ResourceOption) (*DataCollectionEndpoint, error) {
	var resource DataCollectionEndpoint
	err := ctx.ReadResource("azure-native:insights:DataCollectionEndpoint", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type dataCollectionEndpointState struct {
}

type DataCollectionEndpointState struct {
}

func (DataCollectionEndpointState) ElementType() reflect.Type {
	return reflect.TypeOf((*dataCollectionEndpointState)(nil)).Elem()
}

type dataCollectionEndpointArgs struct {
	DataCollectionEndpointName *string                            `pulumi:"dataCollectionEndpointName"`
	Description                *string                            `pulumi:"description"`
	ImmutableId                *string                            `pulumi:"immutableId"`
	Kind                       *string                            `pulumi:"kind"`
	Location                   *string                            `pulumi:"location"`
	NetworkAcls                *DataCollectionEndpointNetworkAcls `pulumi:"networkAcls"`
	ResourceGroupName          string                             `pulumi:"resourceGroupName"`
	Tags                       map[string]string                  `pulumi:"tags"`
}


type DataCollectionEndpointArgs struct {
	DataCollectionEndpointName pulumi.StringPtrInput
	Description                pulumi.StringPtrInput
	ImmutableId                pulumi.StringPtrInput
	Kind                       pulumi.StringPtrInput
	Location                   pulumi.StringPtrInput
	NetworkAcls                DataCollectionEndpointNetworkAclsPtrInput
	ResourceGroupName          pulumi.StringInput
	Tags                       pulumi.StringMapInput
}

func (DataCollectionEndpointArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dataCollectionEndpointArgs)(nil)).Elem()
}

type DataCollectionEndpointInput interface {
	pulumi.Input

	ToDataCollectionEndpointOutput() DataCollectionEndpointOutput
	ToDataCollectionEndpointOutputWithContext(ctx context.Context) DataCollectionEndpointOutput
}

func (*DataCollectionEndpoint) ElementType() reflect.Type {
	return reflect.TypeOf((*DataCollectionEndpoint)(nil))
}

func (i *DataCollectionEndpoint) ToDataCollectionEndpointOutput() DataCollectionEndpointOutput {
	return i.ToDataCollectionEndpointOutputWithContext(context.Background())
}

func (i *DataCollectionEndpoint) ToDataCollectionEndpointOutputWithContext(ctx context.Context) DataCollectionEndpointOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataCollectionEndpointOutput)
}

type DataCollectionEndpointOutput struct{ *pulumi.OutputState }

func (DataCollectionEndpointOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DataCollectionEndpoint)(nil))
}

func (o DataCollectionEndpointOutput) ToDataCollectionEndpointOutput() DataCollectionEndpointOutput {
	return o
}

func (o DataCollectionEndpointOutput) ToDataCollectionEndpointOutputWithContext(ctx context.Context) DataCollectionEndpointOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DataCollectionEndpointInput)(nil)).Elem(), &DataCollectionEndpoint{})
	pulumi.RegisterOutputType(DataCollectionEndpointOutput{})
}
