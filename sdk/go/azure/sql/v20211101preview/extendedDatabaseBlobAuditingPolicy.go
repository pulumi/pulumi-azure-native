


package v20211101preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ExtendedDatabaseBlobAuditingPolicy struct {
	pulumi.CustomResourceState

	AuditActionsAndGroups        pulumi.StringArrayOutput `pulumi:"auditActionsAndGroups"`
	IsAzureMonitorTargetEnabled  pulumi.BoolPtrOutput     `pulumi:"isAzureMonitorTargetEnabled"`
	IsManagedIdentityInUse       pulumi.BoolPtrOutput     `pulumi:"isManagedIdentityInUse"`
	IsStorageSecondaryKeyInUse   pulumi.BoolPtrOutput     `pulumi:"isStorageSecondaryKeyInUse"`
	Name                         pulumi.StringOutput      `pulumi:"name"`
	PredicateExpression          pulumi.StringPtrOutput   `pulumi:"predicateExpression"`
	QueueDelayMs                 pulumi.IntPtrOutput      `pulumi:"queueDelayMs"`
	RetentionDays                pulumi.IntPtrOutput      `pulumi:"retentionDays"`
	State                        pulumi.StringOutput      `pulumi:"state"`
	StorageAccountSubscriptionId pulumi.StringPtrOutput   `pulumi:"storageAccountSubscriptionId"`
	StorageEndpoint              pulumi.StringPtrOutput   `pulumi:"storageEndpoint"`
	Type                         pulumi.StringOutput      `pulumi:"type"`
}


func NewExtendedDatabaseBlobAuditingPolicy(ctx *pulumi.Context,
	name string, args *ExtendedDatabaseBlobAuditingPolicyArgs, opts ...pulumi.ResourceOption) (*ExtendedDatabaseBlobAuditingPolicy, error) {
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
	if args.State == nil {
		return nil, errors.New("invalid value for required argument 'State'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:sql:ExtendedDatabaseBlobAuditingPolicy"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20170301preview:ExtendedDatabaseBlobAuditingPolicy"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20200202preview:ExtendedDatabaseBlobAuditingPolicy"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20200801preview:ExtendedDatabaseBlobAuditingPolicy"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20201101preview:ExtendedDatabaseBlobAuditingPolicy"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20210201preview:ExtendedDatabaseBlobAuditingPolicy"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20210501preview:ExtendedDatabaseBlobAuditingPolicy"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20210801preview:ExtendedDatabaseBlobAuditingPolicy"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20211101:ExtendedDatabaseBlobAuditingPolicy"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20220201preview:ExtendedDatabaseBlobAuditingPolicy"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20220501preview:ExtendedDatabaseBlobAuditingPolicy"),
		},
	})
	opts = append(opts, aliases)
	var resource ExtendedDatabaseBlobAuditingPolicy
	err := ctx.RegisterResource("azure-native:sql/v20211101preview:ExtendedDatabaseBlobAuditingPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetExtendedDatabaseBlobAuditingPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ExtendedDatabaseBlobAuditingPolicyState, opts ...pulumi.ResourceOption) (*ExtendedDatabaseBlobAuditingPolicy, error) {
	var resource ExtendedDatabaseBlobAuditingPolicy
	err := ctx.ReadResource("azure-native:sql/v20211101preview:ExtendedDatabaseBlobAuditingPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type extendedDatabaseBlobAuditingPolicyState struct {
}

type ExtendedDatabaseBlobAuditingPolicyState struct {
}

func (ExtendedDatabaseBlobAuditingPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*extendedDatabaseBlobAuditingPolicyState)(nil)).Elem()
}

type extendedDatabaseBlobAuditingPolicyArgs struct {
	AuditActionsAndGroups        []string                `pulumi:"auditActionsAndGroups"`
	BlobAuditingPolicyName       *string                 `pulumi:"blobAuditingPolicyName"`
	DatabaseName                 string                  `pulumi:"databaseName"`
	IsAzureMonitorTargetEnabled  *bool                   `pulumi:"isAzureMonitorTargetEnabled"`
	IsManagedIdentityInUse       *bool                   `pulumi:"isManagedIdentityInUse"`
	IsStorageSecondaryKeyInUse   *bool                   `pulumi:"isStorageSecondaryKeyInUse"`
	PredicateExpression          *string                 `pulumi:"predicateExpression"`
	QueueDelayMs                 *int                    `pulumi:"queueDelayMs"`
	ResourceGroupName            string                  `pulumi:"resourceGroupName"`
	RetentionDays                *int                    `pulumi:"retentionDays"`
	ServerName                   string                  `pulumi:"serverName"`
	State                        BlobAuditingPolicyState `pulumi:"state"`
	StorageAccountAccessKey      *string                 `pulumi:"storageAccountAccessKey"`
	StorageAccountSubscriptionId *string                 `pulumi:"storageAccountSubscriptionId"`
	StorageEndpoint              *string                 `pulumi:"storageEndpoint"`
}


type ExtendedDatabaseBlobAuditingPolicyArgs struct {
	AuditActionsAndGroups        pulumi.StringArrayInput
	BlobAuditingPolicyName       pulumi.StringPtrInput
	DatabaseName                 pulumi.StringInput
	IsAzureMonitorTargetEnabled  pulumi.BoolPtrInput
	IsManagedIdentityInUse       pulumi.BoolPtrInput
	IsStorageSecondaryKeyInUse   pulumi.BoolPtrInput
	PredicateExpression          pulumi.StringPtrInput
	QueueDelayMs                 pulumi.IntPtrInput
	ResourceGroupName            pulumi.StringInput
	RetentionDays                pulumi.IntPtrInput
	ServerName                   pulumi.StringInput
	State                        BlobAuditingPolicyStateInput
	StorageAccountAccessKey      pulumi.StringPtrInput
	StorageAccountSubscriptionId pulumi.StringPtrInput
	StorageEndpoint              pulumi.StringPtrInput
}

func (ExtendedDatabaseBlobAuditingPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*extendedDatabaseBlobAuditingPolicyArgs)(nil)).Elem()
}

type ExtendedDatabaseBlobAuditingPolicyInput interface {
	pulumi.Input

	ToExtendedDatabaseBlobAuditingPolicyOutput() ExtendedDatabaseBlobAuditingPolicyOutput
	ToExtendedDatabaseBlobAuditingPolicyOutputWithContext(ctx context.Context) ExtendedDatabaseBlobAuditingPolicyOutput
}

func (*ExtendedDatabaseBlobAuditingPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**ExtendedDatabaseBlobAuditingPolicy)(nil)).Elem()
}

func (i *ExtendedDatabaseBlobAuditingPolicy) ToExtendedDatabaseBlobAuditingPolicyOutput() ExtendedDatabaseBlobAuditingPolicyOutput {
	return i.ToExtendedDatabaseBlobAuditingPolicyOutputWithContext(context.Background())
}

func (i *ExtendedDatabaseBlobAuditingPolicy) ToExtendedDatabaseBlobAuditingPolicyOutputWithContext(ctx context.Context) ExtendedDatabaseBlobAuditingPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExtendedDatabaseBlobAuditingPolicyOutput)
}

type ExtendedDatabaseBlobAuditingPolicyOutput struct{ *pulumi.OutputState }

func (ExtendedDatabaseBlobAuditingPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ExtendedDatabaseBlobAuditingPolicy)(nil)).Elem()
}

func (o ExtendedDatabaseBlobAuditingPolicyOutput) ToExtendedDatabaseBlobAuditingPolicyOutput() ExtendedDatabaseBlobAuditingPolicyOutput {
	return o
}

func (o ExtendedDatabaseBlobAuditingPolicyOutput) ToExtendedDatabaseBlobAuditingPolicyOutputWithContext(ctx context.Context) ExtendedDatabaseBlobAuditingPolicyOutput {
	return o
}

func (o ExtendedDatabaseBlobAuditingPolicyOutput) AuditActionsAndGroups() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ExtendedDatabaseBlobAuditingPolicy) pulumi.StringArrayOutput { return v.AuditActionsAndGroups }).(pulumi.StringArrayOutput)
}

func (o ExtendedDatabaseBlobAuditingPolicyOutput) IsAzureMonitorTargetEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ExtendedDatabaseBlobAuditingPolicy) pulumi.BoolPtrOutput { return v.IsAzureMonitorTargetEnabled }).(pulumi.BoolPtrOutput)
}

func (o ExtendedDatabaseBlobAuditingPolicyOutput) IsManagedIdentityInUse() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ExtendedDatabaseBlobAuditingPolicy) pulumi.BoolPtrOutput { return v.IsManagedIdentityInUse }).(pulumi.BoolPtrOutput)
}

func (o ExtendedDatabaseBlobAuditingPolicyOutput) IsStorageSecondaryKeyInUse() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ExtendedDatabaseBlobAuditingPolicy) pulumi.BoolPtrOutput { return v.IsStorageSecondaryKeyInUse }).(pulumi.BoolPtrOutput)
}

func (o ExtendedDatabaseBlobAuditingPolicyOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ExtendedDatabaseBlobAuditingPolicy) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ExtendedDatabaseBlobAuditingPolicyOutput) PredicateExpression() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ExtendedDatabaseBlobAuditingPolicy) pulumi.StringPtrOutput { return v.PredicateExpression }).(pulumi.StringPtrOutput)
}

func (o ExtendedDatabaseBlobAuditingPolicyOutput) QueueDelayMs() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ExtendedDatabaseBlobAuditingPolicy) pulumi.IntPtrOutput { return v.QueueDelayMs }).(pulumi.IntPtrOutput)
}

func (o ExtendedDatabaseBlobAuditingPolicyOutput) RetentionDays() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ExtendedDatabaseBlobAuditingPolicy) pulumi.IntPtrOutput { return v.RetentionDays }).(pulumi.IntPtrOutput)
}

func (o ExtendedDatabaseBlobAuditingPolicyOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v *ExtendedDatabaseBlobAuditingPolicy) pulumi.StringOutput { return v.State }).(pulumi.StringOutput)
}

func (o ExtendedDatabaseBlobAuditingPolicyOutput) StorageAccountSubscriptionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ExtendedDatabaseBlobAuditingPolicy) pulumi.StringPtrOutput {
		return v.StorageAccountSubscriptionId
	}).(pulumi.StringPtrOutput)
}

func (o ExtendedDatabaseBlobAuditingPolicyOutput) StorageEndpoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ExtendedDatabaseBlobAuditingPolicy) pulumi.StringPtrOutput { return v.StorageEndpoint }).(pulumi.StringPtrOutput)
}

func (o ExtendedDatabaseBlobAuditingPolicyOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ExtendedDatabaseBlobAuditingPolicy) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ExtendedDatabaseBlobAuditingPolicyOutput{})
}
