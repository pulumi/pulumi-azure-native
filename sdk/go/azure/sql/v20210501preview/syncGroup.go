


package v20210501preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type SyncGroup struct {
	pulumi.CustomResourceState

	ConflictLoggingRetentionInDays pulumi.IntPtrOutput              `pulumi:"conflictLoggingRetentionInDays"`
	ConflictResolutionPolicy       pulumi.StringPtrOutput           `pulumi:"conflictResolutionPolicy"`
	EnableConflictLogging          pulumi.BoolPtrOutput             `pulumi:"enableConflictLogging"`
	HubDatabaseUserName            pulumi.StringPtrOutput           `pulumi:"hubDatabaseUserName"`
	Interval                       pulumi.IntPtrOutput              `pulumi:"interval"`
	LastSyncTime                   pulumi.StringOutput              `pulumi:"lastSyncTime"`
	Name                           pulumi.StringOutput              `pulumi:"name"`
	PrivateEndpointName            pulumi.StringOutput              `pulumi:"privateEndpointName"`
	Schema                         SyncGroupSchemaResponsePtrOutput `pulumi:"schema"`
	Sku                            SkuResponsePtrOutput             `pulumi:"sku"`
	SyncDatabaseId                 pulumi.StringPtrOutput           `pulumi:"syncDatabaseId"`
	SyncState                      pulumi.StringOutput              `pulumi:"syncState"`
	Type                           pulumi.StringOutput              `pulumi:"type"`
	UsePrivateLinkConnection       pulumi.BoolPtrOutput             `pulumi:"usePrivateLinkConnection"`
}


func NewSyncGroup(ctx *pulumi.Context,
	name string, args *SyncGroupArgs, opts ...pulumi.ResourceOption) (*SyncGroup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DatabaseName == nil {
		return nil, errors.New("invalid value for required argument 'DatabaseName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ServerName == nil {
		return nil, errors.New("invalid value for required argument 'ServerName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:sql:SyncGroup"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20150501preview:SyncGroup"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20190601preview:SyncGroup"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20200202preview:SyncGroup"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20200801preview:SyncGroup"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20201101preview:SyncGroup"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20210201preview:SyncGroup"),
		},
	})
	opts = append(opts, aliases)
	var resource SyncGroup
	err := ctx.RegisterResource("azure-native:sql/v20210501preview:SyncGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetSyncGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SyncGroupState, opts ...pulumi.ResourceOption) (*SyncGroup, error) {
	var resource SyncGroup
	err := ctx.ReadResource("azure-native:sql/v20210501preview:SyncGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type syncGroupState struct {
}

type SyncGroupState struct {
}

func (SyncGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*syncGroupState)(nil)).Elem()
}

type syncGroupArgs struct {
	ConflictLoggingRetentionInDays *int             `pulumi:"conflictLoggingRetentionInDays"`
	ConflictResolutionPolicy       *string          `pulumi:"conflictResolutionPolicy"`
	DatabaseName                   string           `pulumi:"databaseName"`
	EnableConflictLogging          *bool            `pulumi:"enableConflictLogging"`
	HubDatabasePassword            *string          `pulumi:"hubDatabasePassword"`
	HubDatabaseUserName            *string          `pulumi:"hubDatabaseUserName"`
	Interval                       *int             `pulumi:"interval"`
	ResourceGroupName              string           `pulumi:"resourceGroupName"`
	Schema                         *SyncGroupSchema `pulumi:"schema"`
	ServerName                     string           `pulumi:"serverName"`
	Sku                            *Sku             `pulumi:"sku"`
	SyncDatabaseId                 *string          `pulumi:"syncDatabaseId"`
	SyncGroupName                  *string          `pulumi:"syncGroupName"`
	UsePrivateLinkConnection       *bool            `pulumi:"usePrivateLinkConnection"`
}


type SyncGroupArgs struct {
	ConflictLoggingRetentionInDays pulumi.IntPtrInput
	ConflictResolutionPolicy       pulumi.StringPtrInput
	DatabaseName                   pulumi.StringInput
	EnableConflictLogging          pulumi.BoolPtrInput
	HubDatabasePassword            pulumi.StringPtrInput
	HubDatabaseUserName            pulumi.StringPtrInput
	Interval                       pulumi.IntPtrInput
	ResourceGroupName              pulumi.StringInput
	Schema                         SyncGroupSchemaPtrInput
	ServerName                     pulumi.StringInput
	Sku                            SkuPtrInput
	SyncDatabaseId                 pulumi.StringPtrInput
	SyncGroupName                  pulumi.StringPtrInput
	UsePrivateLinkConnection       pulumi.BoolPtrInput
}

func (SyncGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*syncGroupArgs)(nil)).Elem()
}

type SyncGroupInput interface {
	pulumi.Input

	ToSyncGroupOutput() SyncGroupOutput
	ToSyncGroupOutputWithContext(ctx context.Context) SyncGroupOutput
}

func (*SyncGroup) ElementType() reflect.Type {
	return reflect.TypeOf((**SyncGroup)(nil)).Elem()
}

func (i *SyncGroup) ToSyncGroupOutput() SyncGroupOutput {
	return i.ToSyncGroupOutputWithContext(context.Background())
}

func (i *SyncGroup) ToSyncGroupOutputWithContext(ctx context.Context) SyncGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SyncGroupOutput)
}

type SyncGroupOutput struct{ *pulumi.OutputState }

func (SyncGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SyncGroup)(nil)).Elem()
}

func (o SyncGroupOutput) ToSyncGroupOutput() SyncGroupOutput {
	return o
}

func (o SyncGroupOutput) ToSyncGroupOutputWithContext(ctx context.Context) SyncGroupOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(SyncGroupOutput{})
}
