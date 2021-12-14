


package v20190601preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type SyncMember struct {
	pulumi.CustomResourceState

	DatabaseName                      pulumi.StringPtrOutput `pulumi:"databaseName"`
	DatabaseType                      pulumi.StringPtrOutput `pulumi:"databaseType"`
	Name                              pulumi.StringOutput    `pulumi:"name"`
	PrivateEndpointName               pulumi.StringOutput    `pulumi:"privateEndpointName"`
	ServerName                        pulumi.StringPtrOutput `pulumi:"serverName"`
	SqlServerDatabaseId               pulumi.StringPtrOutput `pulumi:"sqlServerDatabaseId"`
	SyncAgentId                       pulumi.StringPtrOutput `pulumi:"syncAgentId"`
	SyncDirection                     pulumi.StringPtrOutput `pulumi:"syncDirection"`
	SyncMemberAzureDatabaseResourceId pulumi.StringPtrOutput `pulumi:"syncMemberAzureDatabaseResourceId"`
	SyncState                         pulumi.StringOutput    `pulumi:"syncState"`
	Type                              pulumi.StringOutput    `pulumi:"type"`
	UsePrivateLinkConnection          pulumi.BoolPtrOutput   `pulumi:"usePrivateLinkConnection"`
	UserName                          pulumi.StringPtrOutput `pulumi:"userName"`
}


func NewSyncMember(ctx *pulumi.Context,
	name string, args *SyncMemberArgs, opts ...pulumi.ResourceOption) (*SyncMember, error) {
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
	if args.SyncGroupName == nil {
		return nil, errors.New("invalid value for required argument 'SyncGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:sql:SyncMember"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20150501preview:SyncMember"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20200202preview:SyncMember"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20200801preview:SyncMember"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20201101preview:SyncMember"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20210201preview:SyncMember"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20210501preview:SyncMember"),
		},
	})
	opts = append(opts, aliases)
	var resource SyncMember
	err := ctx.RegisterResource("azure-native:sql/v20190601preview:SyncMember", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetSyncMember(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SyncMemberState, opts ...pulumi.ResourceOption) (*SyncMember, error) {
	var resource SyncMember
	err := ctx.ReadResource("azure-native:sql/v20190601preview:SyncMember", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type syncMemberState struct {
}

type SyncMemberState struct {
}

func (SyncMemberState) ElementType() reflect.Type {
	return reflect.TypeOf((*syncMemberState)(nil)).Elem()
}

type syncMemberArgs struct {
	DatabaseName                      string  `pulumi:"databaseName"`
	DatabaseType                      *string `pulumi:"databaseType"`
	Password                          *string `pulumi:"password"`
	ResourceGroupName                 string  `pulumi:"resourceGroupName"`
	ServerName                        string  `pulumi:"serverName"`
	SqlServerDatabaseId               *string `pulumi:"sqlServerDatabaseId"`
	SyncAgentId                       *string `pulumi:"syncAgentId"`
	SyncDirection                     *string `pulumi:"syncDirection"`
	SyncGroupName                     string  `pulumi:"syncGroupName"`
	SyncMemberAzureDatabaseResourceId *string `pulumi:"syncMemberAzureDatabaseResourceId"`
	SyncMemberName                    *string `pulumi:"syncMemberName"`
	UsePrivateLinkConnection          *bool   `pulumi:"usePrivateLinkConnection"`
	UserName                          *string `pulumi:"userName"`
}


type SyncMemberArgs struct {
	DatabaseName                      pulumi.StringInput
	DatabaseType                      pulumi.StringPtrInput
	Password                          pulumi.StringPtrInput
	ResourceGroupName                 pulumi.StringInput
	ServerName                        pulumi.StringInput
	SqlServerDatabaseId               pulumi.StringPtrInput
	SyncAgentId                       pulumi.StringPtrInput
	SyncDirection                     pulumi.StringPtrInput
	SyncGroupName                     pulumi.StringInput
	SyncMemberAzureDatabaseResourceId pulumi.StringPtrInput
	SyncMemberName                    pulumi.StringPtrInput
	UsePrivateLinkConnection          pulumi.BoolPtrInput
	UserName                          pulumi.StringPtrInput
}

func (SyncMemberArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*syncMemberArgs)(nil)).Elem()
}

type SyncMemberInput interface {
	pulumi.Input

	ToSyncMemberOutput() SyncMemberOutput
	ToSyncMemberOutputWithContext(ctx context.Context) SyncMemberOutput
}

func (*SyncMember) ElementType() reflect.Type {
	return reflect.TypeOf((**SyncMember)(nil)).Elem()
}

func (i *SyncMember) ToSyncMemberOutput() SyncMemberOutput {
	return i.ToSyncMemberOutputWithContext(context.Background())
}

func (i *SyncMember) ToSyncMemberOutputWithContext(ctx context.Context) SyncMemberOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SyncMemberOutput)
}

type SyncMemberOutput struct{ *pulumi.OutputState }

func (SyncMemberOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SyncMember)(nil)).Elem()
}

func (o SyncMemberOutput) ToSyncMemberOutput() SyncMemberOutput {
	return o
}

func (o SyncMemberOutput) ToSyncMemberOutputWithContext(ctx context.Context) SyncMemberOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(SyncMemberOutput{})
}
