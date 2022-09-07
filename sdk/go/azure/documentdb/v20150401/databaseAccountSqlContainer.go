


package v20150401

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


type DatabaseAccountSqlContainer struct {
	pulumi.CustomResourceState

	ConflictResolutionPolicy ConflictResolutionPolicyResponsePtrOutput `pulumi:"conflictResolutionPolicy"`
	DefaultTtl               pulumi.IntPtrOutput                       `pulumi:"defaultTtl"`
	Etag                     pulumi.StringPtrOutput                    `pulumi:"etag"`
	IndexingPolicy           IndexingPolicyResponsePtrOutput           `pulumi:"indexingPolicy"`
	Location                 pulumi.StringPtrOutput                    `pulumi:"location"`
	Name                     pulumi.StringOutput                       `pulumi:"name"`
	PartitionKey             ContainerPartitionKeyResponsePtrOutput    `pulumi:"partitionKey"`
	Rid                      pulumi.StringPtrOutput                    `pulumi:"rid"`
	Tags                     pulumi.StringMapOutput                    `pulumi:"tags"`
	Ts                       pulumi.AnyOutput                          `pulumi:"ts"`
	Type                     pulumi.StringOutput                       `pulumi:"type"`
	UniqueKeyPolicy          UniqueKeyPolicyResponsePtrOutput          `pulumi:"uniqueKeyPolicy"`
}


func NewDatabaseAccountSqlContainer(ctx *pulumi.Context,
	name string, args *DatabaseAccountSqlContainerArgs, opts ...pulumi.ResourceOption) (*DatabaseAccountSqlContainer, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccountName == nil {
		return nil, errors.New("invalid value for required argument 'AccountName'")
	}
	if args.DatabaseName == nil {
		return nil, errors.New("invalid value for required argument 'DatabaseName'")
	}
	if args.Options == nil {
		return nil, errors.New("invalid value for required argument 'Options'")
	}
	if args.Resource == nil {
		return nil, errors.New("invalid value for required argument 'Resource'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	args.Resource = args.Resource.ToSqlContainerResourceOutput().ApplyT(func(v SqlContainerResource) SqlContainerResource { return *v.Defaults() }).(SqlContainerResourceOutput)
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:documentdb:DatabaseAccountSqlContainer"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20150408:DatabaseAccountSqlContainer"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20151106:DatabaseAccountSqlContainer"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20160319:DatabaseAccountSqlContainer"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20160331:DatabaseAccountSqlContainer"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20190801:DatabaseAccountSqlContainer"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20191212:DatabaseAccountSqlContainer"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20200301:DatabaseAccountSqlContainer"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20200401:DatabaseAccountSqlContainer"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20200601preview:DatabaseAccountSqlContainer"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20200901:DatabaseAccountSqlContainer"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210115:DatabaseAccountSqlContainer"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210301preview:DatabaseAccountSqlContainer"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210315:DatabaseAccountSqlContainer"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210401preview:DatabaseAccountSqlContainer"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210415:DatabaseAccountSqlContainer"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210515:DatabaseAccountSqlContainer"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210615:DatabaseAccountSqlContainer"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210701preview:DatabaseAccountSqlContainer"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20211015:DatabaseAccountSqlContainer"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20211015preview:DatabaseAccountSqlContainer"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20211115preview:DatabaseAccountSqlContainer"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20220215preview:DatabaseAccountSqlContainer"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20220515:DatabaseAccountSqlContainer"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20220515preview:DatabaseAccountSqlContainer"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20220815:DatabaseAccountSqlContainer"),
		},
	})
	opts = append(opts, aliases)
	var resource DatabaseAccountSqlContainer
	err := ctx.RegisterResource("azure-native:documentdb/v20150401:DatabaseAccountSqlContainer", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetDatabaseAccountSqlContainer(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DatabaseAccountSqlContainerState, opts ...pulumi.ResourceOption) (*DatabaseAccountSqlContainer, error) {
	var resource DatabaseAccountSqlContainer
	err := ctx.ReadResource("azure-native:documentdb/v20150401:DatabaseAccountSqlContainer", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type databaseAccountSqlContainerState struct {
}

type DatabaseAccountSqlContainerState struct {
}

func (DatabaseAccountSqlContainerState) ElementType() reflect.Type {
	return reflect.TypeOf((*databaseAccountSqlContainerState)(nil)).Elem()
}

type databaseAccountSqlContainerArgs struct {
	AccountName       string               `pulumi:"accountName"`
	ContainerName     *string              `pulumi:"containerName"`
	DatabaseName      string               `pulumi:"databaseName"`
	Options           map[string]string    `pulumi:"options"`
	Resource          SqlContainerResource `pulumi:"resource"`
	ResourceGroupName string               `pulumi:"resourceGroupName"`
}


type DatabaseAccountSqlContainerArgs struct {
	AccountName       pulumi.StringInput
	ContainerName     pulumi.StringPtrInput
	DatabaseName      pulumi.StringInput
	Options           pulumi.StringMapInput
	Resource          SqlContainerResourceInput
	ResourceGroupName pulumi.StringInput
}

func (DatabaseAccountSqlContainerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*databaseAccountSqlContainerArgs)(nil)).Elem()
}

type DatabaseAccountSqlContainerInput interface {
	pulumi.Input

	ToDatabaseAccountSqlContainerOutput() DatabaseAccountSqlContainerOutput
	ToDatabaseAccountSqlContainerOutputWithContext(ctx context.Context) DatabaseAccountSqlContainerOutput
}

func (*DatabaseAccountSqlContainer) ElementType() reflect.Type {
	return reflect.TypeOf((**DatabaseAccountSqlContainer)(nil)).Elem()
}

func (i *DatabaseAccountSqlContainer) ToDatabaseAccountSqlContainerOutput() DatabaseAccountSqlContainerOutput {
	return i.ToDatabaseAccountSqlContainerOutputWithContext(context.Background())
}

func (i *DatabaseAccountSqlContainer) ToDatabaseAccountSqlContainerOutputWithContext(ctx context.Context) DatabaseAccountSqlContainerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatabaseAccountSqlContainerOutput)
}

type DatabaseAccountSqlContainerOutput struct{ *pulumi.OutputState }

func (DatabaseAccountSqlContainerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DatabaseAccountSqlContainer)(nil)).Elem()
}

func (o DatabaseAccountSqlContainerOutput) ToDatabaseAccountSqlContainerOutput() DatabaseAccountSqlContainerOutput {
	return o
}

func (o DatabaseAccountSqlContainerOutput) ToDatabaseAccountSqlContainerOutputWithContext(ctx context.Context) DatabaseAccountSqlContainerOutput {
	return o
}

func (o DatabaseAccountSqlContainerOutput) ConflictResolutionPolicy() ConflictResolutionPolicyResponsePtrOutput {
	return o.ApplyT(func(v *DatabaseAccountSqlContainer) ConflictResolutionPolicyResponsePtrOutput {
		return v.ConflictResolutionPolicy
	}).(ConflictResolutionPolicyResponsePtrOutput)
}

func (o DatabaseAccountSqlContainerOutput) DefaultTtl() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *DatabaseAccountSqlContainer) pulumi.IntPtrOutput { return v.DefaultTtl }).(pulumi.IntPtrOutput)
}

func (o DatabaseAccountSqlContainerOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DatabaseAccountSqlContainer) pulumi.StringPtrOutput { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o DatabaseAccountSqlContainerOutput) IndexingPolicy() IndexingPolicyResponsePtrOutput {
	return o.ApplyT(func(v *DatabaseAccountSqlContainer) IndexingPolicyResponsePtrOutput { return v.IndexingPolicy }).(IndexingPolicyResponsePtrOutput)
}

func (o DatabaseAccountSqlContainerOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DatabaseAccountSqlContainer) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

func (o DatabaseAccountSqlContainerOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *DatabaseAccountSqlContainer) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o DatabaseAccountSqlContainerOutput) PartitionKey() ContainerPartitionKeyResponsePtrOutput {
	return o.ApplyT(func(v *DatabaseAccountSqlContainer) ContainerPartitionKeyResponsePtrOutput { return v.PartitionKey }).(ContainerPartitionKeyResponsePtrOutput)
}

func (o DatabaseAccountSqlContainerOutput) Rid() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DatabaseAccountSqlContainer) pulumi.StringPtrOutput { return v.Rid }).(pulumi.StringPtrOutput)
}

func (o DatabaseAccountSqlContainerOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *DatabaseAccountSqlContainer) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o DatabaseAccountSqlContainerOutput) Ts() pulumi.AnyOutput {
	return o.ApplyT(func(v *DatabaseAccountSqlContainer) pulumi.AnyOutput { return v.Ts }).(pulumi.AnyOutput)
}

func (o DatabaseAccountSqlContainerOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *DatabaseAccountSqlContainer) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o DatabaseAccountSqlContainerOutput) UniqueKeyPolicy() UniqueKeyPolicyResponsePtrOutput {
	return o.ApplyT(func(v *DatabaseAccountSqlContainer) UniqueKeyPolicyResponsePtrOutput { return v.UniqueKeyPolicy }).(UniqueKeyPolicyResponsePtrOutput)
}

func init() {
	pulumi.RegisterOutputType(DatabaseAccountSqlContainerOutput{})
}
