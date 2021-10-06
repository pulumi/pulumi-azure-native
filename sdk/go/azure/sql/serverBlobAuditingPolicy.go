


package sql

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ServerBlobAuditingPolicy struct {
	pulumi.CustomResourceState

	AuditActionsAndGroups        pulumi.StringArrayOutput `pulumi:"auditActionsAndGroups"`
	IsAzureMonitorTargetEnabled  pulumi.BoolPtrOutput     `pulumi:"isAzureMonitorTargetEnabled"`
	IsDevopsAuditEnabled         pulumi.BoolPtrOutput     `pulumi:"isDevopsAuditEnabled"`
	IsStorageSecondaryKeyInUse   pulumi.BoolPtrOutput     `pulumi:"isStorageSecondaryKeyInUse"`
	Name                         pulumi.StringOutput      `pulumi:"name"`
	QueueDelayMs                 pulumi.IntPtrOutput      `pulumi:"queueDelayMs"`
	RetentionDays                pulumi.IntPtrOutput      `pulumi:"retentionDays"`
	State                        pulumi.StringOutput      `pulumi:"state"`
	StorageAccountSubscriptionId pulumi.StringPtrOutput   `pulumi:"storageAccountSubscriptionId"`
	StorageEndpoint              pulumi.StringPtrOutput   `pulumi:"storageEndpoint"`
	Type                         pulumi.StringOutput      `pulumi:"type"`
}


func NewServerBlobAuditingPolicy(ctx *pulumi.Context,
	name string, args *ServerBlobAuditingPolicyArgs, opts ...pulumi.ResourceOption) (*ServerBlobAuditingPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ServerName == nil {
		return nil, errors.New("invalid value for required argument 'ServerName'")
	}
	if args.State == nil {
		return nil, errors.New("invalid value for required argument 'State'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:sql:ServerBlobAuditingPolicy"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20170301preview:ServerBlobAuditingPolicy"),
		},
		{
			Type: pulumi.String("azure-nextgen:sql/v20170301preview:ServerBlobAuditingPolicy"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20200202preview:ServerBlobAuditingPolicy"),
		},
		{
			Type: pulumi.String("azure-nextgen:sql/v20200202preview:ServerBlobAuditingPolicy"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20200801preview:ServerBlobAuditingPolicy"),
		},
		{
			Type: pulumi.String("azure-nextgen:sql/v20200801preview:ServerBlobAuditingPolicy"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20201101preview:ServerBlobAuditingPolicy"),
		},
		{
			Type: pulumi.String("azure-nextgen:sql/v20201101preview:ServerBlobAuditingPolicy"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20210201preview:ServerBlobAuditingPolicy"),
		},
		{
			Type: pulumi.String("azure-nextgen:sql/v20210201preview:ServerBlobAuditingPolicy"),
		},
	})
	opts = append(opts, aliases)
	var resource ServerBlobAuditingPolicy
	err := ctx.RegisterResource("azure-native:sql:ServerBlobAuditingPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetServerBlobAuditingPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServerBlobAuditingPolicyState, opts ...pulumi.ResourceOption) (*ServerBlobAuditingPolicy, error) {
	var resource ServerBlobAuditingPolicy
	err := ctx.ReadResource("azure-native:sql:ServerBlobAuditingPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type serverBlobAuditingPolicyState struct {
}

type ServerBlobAuditingPolicyState struct {
}

func (ServerBlobAuditingPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*serverBlobAuditingPolicyState)(nil)).Elem()
}

type serverBlobAuditingPolicyArgs struct {
	AuditActionsAndGroups        []string                `pulumi:"auditActionsAndGroups"`
	BlobAuditingPolicyName       *string                 `pulumi:"blobAuditingPolicyName"`
	IsAzureMonitorTargetEnabled  *bool                   `pulumi:"isAzureMonitorTargetEnabled"`
	IsDevopsAuditEnabled         *bool                   `pulumi:"isDevopsAuditEnabled"`
	IsStorageSecondaryKeyInUse   *bool                   `pulumi:"isStorageSecondaryKeyInUse"`
	QueueDelayMs                 *int                    `pulumi:"queueDelayMs"`
	ResourceGroupName            string                  `pulumi:"resourceGroupName"`
	RetentionDays                *int                    `pulumi:"retentionDays"`
	ServerName                   string                  `pulumi:"serverName"`
	State                        BlobAuditingPolicyState `pulumi:"state"`
	StorageAccountAccessKey      *string                 `pulumi:"storageAccountAccessKey"`
	StorageAccountSubscriptionId *string                 `pulumi:"storageAccountSubscriptionId"`
	StorageEndpoint              *string                 `pulumi:"storageEndpoint"`
}


type ServerBlobAuditingPolicyArgs struct {
	AuditActionsAndGroups        pulumi.StringArrayInput
	BlobAuditingPolicyName       pulumi.StringPtrInput
	IsAzureMonitorTargetEnabled  pulumi.BoolPtrInput
	IsDevopsAuditEnabled         pulumi.BoolPtrInput
	IsStorageSecondaryKeyInUse   pulumi.BoolPtrInput
	QueueDelayMs                 pulumi.IntPtrInput
	ResourceGroupName            pulumi.StringInput
	RetentionDays                pulumi.IntPtrInput
	ServerName                   pulumi.StringInput
	State                        BlobAuditingPolicyStateInput
	StorageAccountAccessKey      pulumi.StringPtrInput
	StorageAccountSubscriptionId pulumi.StringPtrInput
	StorageEndpoint              pulumi.StringPtrInput
}

func (ServerBlobAuditingPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*serverBlobAuditingPolicyArgs)(nil)).Elem()
}

type ServerBlobAuditingPolicyInput interface {
	pulumi.Input

	ToServerBlobAuditingPolicyOutput() ServerBlobAuditingPolicyOutput
	ToServerBlobAuditingPolicyOutputWithContext(ctx context.Context) ServerBlobAuditingPolicyOutput
}

func (*ServerBlobAuditingPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((*ServerBlobAuditingPolicy)(nil))
}

func (i *ServerBlobAuditingPolicy) ToServerBlobAuditingPolicyOutput() ServerBlobAuditingPolicyOutput {
	return i.ToServerBlobAuditingPolicyOutputWithContext(context.Background())
}

func (i *ServerBlobAuditingPolicy) ToServerBlobAuditingPolicyOutputWithContext(ctx context.Context) ServerBlobAuditingPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServerBlobAuditingPolicyOutput)
}

type ServerBlobAuditingPolicyOutput struct{ *pulumi.OutputState }

func (ServerBlobAuditingPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServerBlobAuditingPolicy)(nil))
}

func (o ServerBlobAuditingPolicyOutput) ToServerBlobAuditingPolicyOutput() ServerBlobAuditingPolicyOutput {
	return o
}

func (o ServerBlobAuditingPolicyOutput) ToServerBlobAuditingPolicyOutputWithContext(ctx context.Context) ServerBlobAuditingPolicyOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ServerBlobAuditingPolicyOutput{})
}
