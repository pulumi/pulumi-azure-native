


package v20221111

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type CosmosDbDataConnection struct {
	pulumi.CustomResourceState

	CosmosDbAccountResourceId pulumi.StringOutput    `pulumi:"cosmosDbAccountResourceId"`
	CosmosDbContainer         pulumi.StringOutput    `pulumi:"cosmosDbContainer"`
	CosmosDbDatabase          pulumi.StringOutput    `pulumi:"cosmosDbDatabase"`
	Kind                      pulumi.StringOutput    `pulumi:"kind"`
	Location                  pulumi.StringPtrOutput `pulumi:"location"`
	ManagedIdentityObjectId   pulumi.StringOutput    `pulumi:"managedIdentityObjectId"`
	ManagedIdentityResourceId pulumi.StringOutput    `pulumi:"managedIdentityResourceId"`
	MappingRuleName           pulumi.StringPtrOutput `pulumi:"mappingRuleName"`
	Name                      pulumi.StringOutput    `pulumi:"name"`
	ProvisioningState         pulumi.StringOutput    `pulumi:"provisioningState"`
	RetrievalStartDate        pulumi.StringPtrOutput `pulumi:"retrievalStartDate"`
	TableName                 pulumi.StringOutput    `pulumi:"tableName"`
	Type                      pulumi.StringOutput    `pulumi:"type"`
}


func NewCosmosDbDataConnection(ctx *pulumi.Context,
	name string, args *CosmosDbDataConnectionArgs, opts ...pulumi.ResourceOption) (*CosmosDbDataConnection, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ClusterName == nil {
		return nil, errors.New("invalid value for required argument 'ClusterName'")
	}
	if args.CosmosDbAccountResourceId == nil {
		return nil, errors.New("invalid value for required argument 'CosmosDbAccountResourceId'")
	}
	if args.CosmosDbContainer == nil {
		return nil, errors.New("invalid value for required argument 'CosmosDbContainer'")
	}
	if args.CosmosDbDatabase == nil {
		return nil, errors.New("invalid value for required argument 'CosmosDbDatabase'")
	}
	if args.DatabaseName == nil {
		return nil, errors.New("invalid value for required argument 'DatabaseName'")
	}
	if args.Kind == nil {
		return nil, errors.New("invalid value for required argument 'Kind'")
	}
	if args.ManagedIdentityResourceId == nil {
		return nil, errors.New("invalid value for required argument 'ManagedIdentityResourceId'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.TableName == nil {
		return nil, errors.New("invalid value for required argument 'TableName'")
	}
	args.Kind = pulumi.String("CosmosDb")
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:kusto:CosmosDbDataConnection"),
		},
		{
			Type: pulumi.String("azure-native:kusto/v20190121:CosmosDbDataConnection"),
		},
		{
			Type: pulumi.String("azure-native:kusto/v20190515:CosmosDbDataConnection"),
		},
		{
			Type: pulumi.String("azure-native:kusto/v20190907:CosmosDbDataConnection"),
		},
		{
			Type: pulumi.String("azure-native:kusto/v20191109:CosmosDbDataConnection"),
		},
		{
			Type: pulumi.String("azure-native:kusto/v20200215:CosmosDbDataConnection"),
		},
		{
			Type: pulumi.String("azure-native:kusto/v20200614:CosmosDbDataConnection"),
		},
		{
			Type: pulumi.String("azure-native:kusto/v20200918:CosmosDbDataConnection"),
		},
		{
			Type: pulumi.String("azure-native:kusto/v20210101:CosmosDbDataConnection"),
		},
		{
			Type: pulumi.String("azure-native:kusto/v20210827:CosmosDbDataConnection"),
		},
		{
			Type: pulumi.String("azure-native:kusto/v20220201:CosmosDbDataConnection"),
		},
		{
			Type: pulumi.String("azure-native:kusto/v20220707:CosmosDbDataConnection"),
		},
	})
	opts = append(opts, aliases)
	var resource CosmosDbDataConnection
	err := ctx.RegisterResource("azure-native:kusto/v20221111:CosmosDbDataConnection", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetCosmosDbDataConnection(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CosmosDbDataConnectionState, opts ...pulumi.ResourceOption) (*CosmosDbDataConnection, error) {
	var resource CosmosDbDataConnection
	err := ctx.ReadResource("azure-native:kusto/v20221111:CosmosDbDataConnection", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type cosmosDbDataConnectionState struct {
}

type CosmosDbDataConnectionState struct {
}

func (CosmosDbDataConnectionState) ElementType() reflect.Type {
	return reflect.TypeOf((*cosmosDbDataConnectionState)(nil)).Elem()
}

type cosmosDbDataConnectionArgs struct {
	ClusterName               string  `pulumi:"clusterName"`
	CosmosDbAccountResourceId string  `pulumi:"cosmosDbAccountResourceId"`
	CosmosDbContainer         string  `pulumi:"cosmosDbContainer"`
	CosmosDbDatabase          string  `pulumi:"cosmosDbDatabase"`
	DataConnectionName        *string `pulumi:"dataConnectionName"`
	DatabaseName              string  `pulumi:"databaseName"`
	Kind                      string  `pulumi:"kind"`
	Location                  *string `pulumi:"location"`
	ManagedIdentityResourceId string  `pulumi:"managedIdentityResourceId"`
	MappingRuleName           *string `pulumi:"mappingRuleName"`
	ResourceGroupName         string  `pulumi:"resourceGroupName"`
	RetrievalStartDate        *string `pulumi:"retrievalStartDate"`
	TableName                 string  `pulumi:"tableName"`
}


type CosmosDbDataConnectionArgs struct {
	ClusterName               pulumi.StringInput
	CosmosDbAccountResourceId pulumi.StringInput
	CosmosDbContainer         pulumi.StringInput
	CosmosDbDatabase          pulumi.StringInput
	DataConnectionName        pulumi.StringPtrInput
	DatabaseName              pulumi.StringInput
	Kind                      pulumi.StringInput
	Location                  pulumi.StringPtrInput
	ManagedIdentityResourceId pulumi.StringInput
	MappingRuleName           pulumi.StringPtrInput
	ResourceGroupName         pulumi.StringInput
	RetrievalStartDate        pulumi.StringPtrInput
	TableName                 pulumi.StringInput
}

func (CosmosDbDataConnectionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*cosmosDbDataConnectionArgs)(nil)).Elem()
}

type CosmosDbDataConnectionInput interface {
	pulumi.Input

	ToCosmosDbDataConnectionOutput() CosmosDbDataConnectionOutput
	ToCosmosDbDataConnectionOutputWithContext(ctx context.Context) CosmosDbDataConnectionOutput
}

func (*CosmosDbDataConnection) ElementType() reflect.Type {
	return reflect.TypeOf((**CosmosDbDataConnection)(nil)).Elem()
}

func (i *CosmosDbDataConnection) ToCosmosDbDataConnectionOutput() CosmosDbDataConnectionOutput {
	return i.ToCosmosDbDataConnectionOutputWithContext(context.Background())
}

func (i *CosmosDbDataConnection) ToCosmosDbDataConnectionOutputWithContext(ctx context.Context) CosmosDbDataConnectionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CosmosDbDataConnectionOutput)
}

type CosmosDbDataConnectionOutput struct{ *pulumi.OutputState }

func (CosmosDbDataConnectionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CosmosDbDataConnection)(nil)).Elem()
}

func (o CosmosDbDataConnectionOutput) ToCosmosDbDataConnectionOutput() CosmosDbDataConnectionOutput {
	return o
}

func (o CosmosDbDataConnectionOutput) ToCosmosDbDataConnectionOutputWithContext(ctx context.Context) CosmosDbDataConnectionOutput {
	return o
}

func (o CosmosDbDataConnectionOutput) CosmosDbAccountResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v *CosmosDbDataConnection) pulumi.StringOutput { return v.CosmosDbAccountResourceId }).(pulumi.StringOutput)
}

func (o CosmosDbDataConnectionOutput) CosmosDbContainer() pulumi.StringOutput {
	return o.ApplyT(func(v *CosmosDbDataConnection) pulumi.StringOutput { return v.CosmosDbContainer }).(pulumi.StringOutput)
}

func (o CosmosDbDataConnectionOutput) CosmosDbDatabase() pulumi.StringOutput {
	return o.ApplyT(func(v *CosmosDbDataConnection) pulumi.StringOutput { return v.CosmosDbDatabase }).(pulumi.StringOutput)
}

func (o CosmosDbDataConnectionOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *CosmosDbDataConnection) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

func (o CosmosDbDataConnectionOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CosmosDbDataConnection) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

func (o CosmosDbDataConnectionOutput) ManagedIdentityObjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *CosmosDbDataConnection) pulumi.StringOutput { return v.ManagedIdentityObjectId }).(pulumi.StringOutput)
}

func (o CosmosDbDataConnectionOutput) ManagedIdentityResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v *CosmosDbDataConnection) pulumi.StringOutput { return v.ManagedIdentityResourceId }).(pulumi.StringOutput)
}

func (o CosmosDbDataConnectionOutput) MappingRuleName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CosmosDbDataConnection) pulumi.StringPtrOutput { return v.MappingRuleName }).(pulumi.StringPtrOutput)
}

func (o CosmosDbDataConnectionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *CosmosDbDataConnection) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o CosmosDbDataConnectionOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *CosmosDbDataConnection) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o CosmosDbDataConnectionOutput) RetrievalStartDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CosmosDbDataConnection) pulumi.StringPtrOutput { return v.RetrievalStartDate }).(pulumi.StringPtrOutput)
}

func (o CosmosDbDataConnectionOutput) TableName() pulumi.StringOutput {
	return o.ApplyT(func(v *CosmosDbDataConnection) pulumi.StringOutput { return v.TableName }).(pulumi.StringOutput)
}

func (o CosmosDbDataConnectionOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *CosmosDbDataConnection) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(CosmosDbDataConnectionOutput{})
}
