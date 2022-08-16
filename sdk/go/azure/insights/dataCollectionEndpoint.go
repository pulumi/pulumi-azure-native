


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
			Type: pulumi.String("azure-native:insights/v20210401:DataCollectionEndpoint"),
		},
		{
			Type: pulumi.String("azure-native:insights/v20210901preview:DataCollectionEndpoint"),
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
	return reflect.TypeOf((**DataCollectionEndpoint)(nil)).Elem()
}

func (i *DataCollectionEndpoint) ToDataCollectionEndpointOutput() DataCollectionEndpointOutput {
	return i.ToDataCollectionEndpointOutputWithContext(context.Background())
}

func (i *DataCollectionEndpoint) ToDataCollectionEndpointOutputWithContext(ctx context.Context) DataCollectionEndpointOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataCollectionEndpointOutput)
}

type DataCollectionEndpointOutput struct{ *pulumi.OutputState }

func (DataCollectionEndpointOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DataCollectionEndpoint)(nil)).Elem()
}

func (o DataCollectionEndpointOutput) ToDataCollectionEndpointOutput() DataCollectionEndpointOutput {
	return o
}

func (o DataCollectionEndpointOutput) ToDataCollectionEndpointOutputWithContext(ctx context.Context) DataCollectionEndpointOutput {
	return o
}

func (o DataCollectionEndpointOutput) ConfigurationAccess() DataCollectionEndpointResponseConfigurationAccessPtrOutput {
	return o.ApplyT(func(v *DataCollectionEndpoint) DataCollectionEndpointResponseConfigurationAccessPtrOutput {
		return v.ConfigurationAccess
	}).(DataCollectionEndpointResponseConfigurationAccessPtrOutput)
}

func (o DataCollectionEndpointOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DataCollectionEndpoint) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

func (o DataCollectionEndpointOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *DataCollectionEndpoint) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

func (o DataCollectionEndpointOutput) ImmutableId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DataCollectionEndpoint) pulumi.StringPtrOutput { return v.ImmutableId }).(pulumi.StringPtrOutput)
}

func (o DataCollectionEndpointOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DataCollectionEndpoint) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o DataCollectionEndpointOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *DataCollectionEndpoint) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o DataCollectionEndpointOutput) LogsIngestion() DataCollectionEndpointResponseLogsIngestionPtrOutput {
	return o.ApplyT(func(v *DataCollectionEndpoint) DataCollectionEndpointResponseLogsIngestionPtrOutput {
		return v.LogsIngestion
	}).(DataCollectionEndpointResponseLogsIngestionPtrOutput)
}

func (o DataCollectionEndpointOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *DataCollectionEndpoint) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o DataCollectionEndpointOutput) NetworkAcls() DataCollectionEndpointResponseNetworkAclsPtrOutput {
	return o.ApplyT(func(v *DataCollectionEndpoint) DataCollectionEndpointResponseNetworkAclsPtrOutput {
		return v.NetworkAcls
	}).(DataCollectionEndpointResponseNetworkAclsPtrOutput)
}

func (o DataCollectionEndpointOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *DataCollectionEndpoint) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o DataCollectionEndpointOutput) SystemData() DataCollectionEndpointResourceResponseSystemDataOutput {
	return o.ApplyT(func(v *DataCollectionEndpoint) DataCollectionEndpointResourceResponseSystemDataOutput {
		return v.SystemData
	}).(DataCollectionEndpointResourceResponseSystemDataOutput)
}

func (o DataCollectionEndpointOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *DataCollectionEndpoint) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o DataCollectionEndpointOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *DataCollectionEndpoint) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(DataCollectionEndpointOutput{})
}
