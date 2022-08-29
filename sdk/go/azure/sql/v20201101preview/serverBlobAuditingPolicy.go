


package v20201101preview

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
			Type: pulumi.String("azure-native:sql:ServerBlobAuditingPolicy"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20170301preview:ServerBlobAuditingPolicy"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20200202preview:ServerBlobAuditingPolicy"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20200801preview:ServerBlobAuditingPolicy"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20210201preview:ServerBlobAuditingPolicy"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20210501preview:ServerBlobAuditingPolicy"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20210801preview:ServerBlobAuditingPolicy"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20211101:ServerBlobAuditingPolicy"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20211101preview:ServerBlobAuditingPolicy"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20220201preview:ServerBlobAuditingPolicy"),
		},
	})
	opts = append(opts, aliases)
	var resource ServerBlobAuditingPolicy
	err := ctx.RegisterResource("azure-native:sql/v20201101preview:ServerBlobAuditingPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetServerBlobAuditingPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServerBlobAuditingPolicyState, opts ...pulumi.ResourceOption) (*ServerBlobAuditingPolicy, error) {
	var resource ServerBlobAuditingPolicy
	err := ctx.ReadResource("azure-native:sql/v20201101preview:ServerBlobAuditingPolicy", name, id, state, &resource, opts...)
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
	return reflect.TypeOf((**ServerBlobAuditingPolicy)(nil)).Elem()
}

func (i *ServerBlobAuditingPolicy) ToServerBlobAuditingPolicyOutput() ServerBlobAuditingPolicyOutput {
	return i.ToServerBlobAuditingPolicyOutputWithContext(context.Background())
}

func (i *ServerBlobAuditingPolicy) ToServerBlobAuditingPolicyOutputWithContext(ctx context.Context) ServerBlobAuditingPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServerBlobAuditingPolicyOutput)
}

type ServerBlobAuditingPolicyOutput struct{ *pulumi.OutputState }

func (ServerBlobAuditingPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServerBlobAuditingPolicy)(nil)).Elem()
}

func (o ServerBlobAuditingPolicyOutput) ToServerBlobAuditingPolicyOutput() ServerBlobAuditingPolicyOutput {
	return o
}

func (o ServerBlobAuditingPolicyOutput) ToServerBlobAuditingPolicyOutputWithContext(ctx context.Context) ServerBlobAuditingPolicyOutput {
	return o
}

func (o ServerBlobAuditingPolicyOutput) AuditActionsAndGroups() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ServerBlobAuditingPolicy) pulumi.StringArrayOutput { return v.AuditActionsAndGroups }).(pulumi.StringArrayOutput)
}

func (o ServerBlobAuditingPolicyOutput) IsAzureMonitorTargetEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ServerBlobAuditingPolicy) pulumi.BoolPtrOutput { return v.IsAzureMonitorTargetEnabled }).(pulumi.BoolPtrOutput)
}

func (o ServerBlobAuditingPolicyOutput) IsDevopsAuditEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ServerBlobAuditingPolicy) pulumi.BoolPtrOutput { return v.IsDevopsAuditEnabled }).(pulumi.BoolPtrOutput)
}

func (o ServerBlobAuditingPolicyOutput) IsStorageSecondaryKeyInUse() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ServerBlobAuditingPolicy) pulumi.BoolPtrOutput { return v.IsStorageSecondaryKeyInUse }).(pulumi.BoolPtrOutput)
}

func (o ServerBlobAuditingPolicyOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ServerBlobAuditingPolicy) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ServerBlobAuditingPolicyOutput) QueueDelayMs() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ServerBlobAuditingPolicy) pulumi.IntPtrOutput { return v.QueueDelayMs }).(pulumi.IntPtrOutput)
}

func (o ServerBlobAuditingPolicyOutput) RetentionDays() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ServerBlobAuditingPolicy) pulumi.IntPtrOutput { return v.RetentionDays }).(pulumi.IntPtrOutput)
}

func (o ServerBlobAuditingPolicyOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v *ServerBlobAuditingPolicy) pulumi.StringOutput { return v.State }).(pulumi.StringOutput)
}

func (o ServerBlobAuditingPolicyOutput) StorageAccountSubscriptionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServerBlobAuditingPolicy) pulumi.StringPtrOutput { return v.StorageAccountSubscriptionId }).(pulumi.StringPtrOutput)
}

func (o ServerBlobAuditingPolicyOutput) StorageEndpoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServerBlobAuditingPolicy) pulumi.StringPtrOutput { return v.StorageEndpoint }).(pulumi.StringPtrOutput)
}

func (o ServerBlobAuditingPolicyOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ServerBlobAuditingPolicy) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ServerBlobAuditingPolicyOutput{})
}
