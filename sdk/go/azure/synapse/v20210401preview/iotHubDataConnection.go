


package v20210401preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type IotHubDataConnection struct {
	pulumi.CustomResourceState

	ConsumerGroup          pulumi.StringOutput      `pulumi:"consumerGroup"`
	DataFormat             pulumi.StringPtrOutput   `pulumi:"dataFormat"`
	EventSystemProperties  pulumi.StringArrayOutput `pulumi:"eventSystemProperties"`
	IotHubResourceId       pulumi.StringOutput      `pulumi:"iotHubResourceId"`
	Kind                   pulumi.StringOutput      `pulumi:"kind"`
	Location               pulumi.StringPtrOutput   `pulumi:"location"`
	MappingRuleName        pulumi.StringPtrOutput   `pulumi:"mappingRuleName"`
	Name                   pulumi.StringOutput      `pulumi:"name"`
	ProvisioningState      pulumi.StringOutput      `pulumi:"provisioningState"`
	SharedAccessPolicyName pulumi.StringOutput      `pulumi:"sharedAccessPolicyName"`
	SystemData             SystemDataResponseOutput `pulumi:"systemData"`
	TableName              pulumi.StringPtrOutput   `pulumi:"tableName"`
	Type                   pulumi.StringOutput      `pulumi:"type"`
}


func NewIotHubDataConnection(ctx *pulumi.Context,
	name string, args *IotHubDataConnectionArgs, opts ...pulumi.ResourceOption) (*IotHubDataConnection, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ConsumerGroup == nil {
		return nil, errors.New("invalid value for required argument 'ConsumerGroup'")
	}
	if args.DatabaseName == nil {
		return nil, errors.New("invalid value for required argument 'DatabaseName'")
	}
	if args.IotHubResourceId == nil {
		return nil, errors.New("invalid value for required argument 'IotHubResourceId'")
	}
	if args.Kind == nil {
		return nil, errors.New("invalid value for required argument 'Kind'")
	}
	if args.KustoPoolName == nil {
		return nil, errors.New("invalid value for required argument 'KustoPoolName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.SharedAccessPolicyName == nil {
		return nil, errors.New("invalid value for required argument 'SharedAccessPolicyName'")
	}
	if args.WorkspaceName == nil {
		return nil, errors.New("invalid value for required argument 'WorkspaceName'")
	}
	args.Kind = pulumi.String("IotHub")
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:synapse/v20210401preview:IotHubDataConnection"),
		},
		{
			Type: pulumi.String("azure-native:synapse:IotHubDataConnection"),
		},
		{
			Type: pulumi.String("azure-nextgen:synapse:IotHubDataConnection"),
		},
		{
			Type: pulumi.String("azure-native:synapse/v20210601preview:IotHubDataConnection"),
		},
		{
			Type: pulumi.String("azure-nextgen:synapse/v20210601preview:IotHubDataConnection"),
		},
	})
	opts = append(opts, aliases)
	var resource IotHubDataConnection
	err := ctx.RegisterResource("azure-native:synapse/v20210401preview:IotHubDataConnection", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetIotHubDataConnection(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IotHubDataConnectionState, opts ...pulumi.ResourceOption) (*IotHubDataConnection, error) {
	var resource IotHubDataConnection
	err := ctx.ReadResource("azure-native:synapse/v20210401preview:IotHubDataConnection", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type iotHubDataConnectionState struct {
}

type IotHubDataConnectionState struct {
}

func (IotHubDataConnectionState) ElementType() reflect.Type {
	return reflect.TypeOf((*iotHubDataConnectionState)(nil)).Elem()
}

type iotHubDataConnectionArgs struct {
	ConsumerGroup          string   `pulumi:"consumerGroup"`
	DataConnectionName     *string  `pulumi:"dataConnectionName"`
	DataFormat             *string  `pulumi:"dataFormat"`
	DatabaseName           string   `pulumi:"databaseName"`
	EventSystemProperties  []string `pulumi:"eventSystemProperties"`
	IotHubResourceId       string   `pulumi:"iotHubResourceId"`
	Kind                   string   `pulumi:"kind"`
	KustoPoolName          string   `pulumi:"kustoPoolName"`
	Location               *string  `pulumi:"location"`
	MappingRuleName        *string  `pulumi:"mappingRuleName"`
	ResourceGroupName      string   `pulumi:"resourceGroupName"`
	SharedAccessPolicyName string   `pulumi:"sharedAccessPolicyName"`
	TableName              *string  `pulumi:"tableName"`
	WorkspaceName          string   `pulumi:"workspaceName"`
}


type IotHubDataConnectionArgs struct {
	ConsumerGroup          pulumi.StringInput
	DataConnectionName     pulumi.StringPtrInput
	DataFormat             pulumi.StringPtrInput
	DatabaseName           pulumi.StringInput
	EventSystemProperties  pulumi.StringArrayInput
	IotHubResourceId       pulumi.StringInput
	Kind                   pulumi.StringInput
	KustoPoolName          pulumi.StringInput
	Location               pulumi.StringPtrInput
	MappingRuleName        pulumi.StringPtrInput
	ResourceGroupName      pulumi.StringInput
	SharedAccessPolicyName pulumi.StringInput
	TableName              pulumi.StringPtrInput
	WorkspaceName          pulumi.StringInput
}

func (IotHubDataConnectionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*iotHubDataConnectionArgs)(nil)).Elem()
}

type IotHubDataConnectionInput interface {
	pulumi.Input

	ToIotHubDataConnectionOutput() IotHubDataConnectionOutput
	ToIotHubDataConnectionOutputWithContext(ctx context.Context) IotHubDataConnectionOutput
}

func (*IotHubDataConnection) ElementType() reflect.Type {
	return reflect.TypeOf((*IotHubDataConnection)(nil))
}

func (i *IotHubDataConnection) ToIotHubDataConnectionOutput() IotHubDataConnectionOutput {
	return i.ToIotHubDataConnectionOutputWithContext(context.Background())
}

func (i *IotHubDataConnection) ToIotHubDataConnectionOutputWithContext(ctx context.Context) IotHubDataConnectionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IotHubDataConnectionOutput)
}

type IotHubDataConnectionOutput struct{ *pulumi.OutputState }

func (IotHubDataConnectionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IotHubDataConnection)(nil))
}

func (o IotHubDataConnectionOutput) ToIotHubDataConnectionOutput() IotHubDataConnectionOutput {
	return o
}

func (o IotHubDataConnectionOutput) ToIotHubDataConnectionOutputWithContext(ctx context.Context) IotHubDataConnectionOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*IotHubDataConnectionInput)(nil)).Elem(), &IotHubDataConnection{})
	pulumi.RegisterOutputType(IotHubDataConnectionOutput{})
}
