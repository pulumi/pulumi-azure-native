


package v20210201preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ExtendedServerBlobAuditingPolicy struct {
	pulumi.CustomResourceState

	AuditActionsAndGroups        pulumi.StringArrayOutput `pulumi:"auditActionsAndGroups"`
	IsAzureMonitorTargetEnabled  pulumi.BoolPtrOutput     `pulumi:"isAzureMonitorTargetEnabled"`
	IsDevopsAuditEnabled         pulumi.BoolPtrOutput     `pulumi:"isDevopsAuditEnabled"`
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


func NewExtendedServerBlobAuditingPolicy(ctx *pulumi.Context,
	name string, args *ExtendedServerBlobAuditingPolicyArgs, opts ...pulumi.ResourceOption) (*ExtendedServerBlobAuditingPolicy, error) {
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
			Type: pulumi.String("azure-native:sql:ExtendedServerBlobAuditingPolicy"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20170301preview:ExtendedServerBlobAuditingPolicy"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20200202preview:ExtendedServerBlobAuditingPolicy"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20200801preview:ExtendedServerBlobAuditingPolicy"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20201101preview:ExtendedServerBlobAuditingPolicy"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20210501preview:ExtendedServerBlobAuditingPolicy"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20210801preview:ExtendedServerBlobAuditingPolicy"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20211101:ExtendedServerBlobAuditingPolicy"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20211101preview:ExtendedServerBlobAuditingPolicy"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20220201preview:ExtendedServerBlobAuditingPolicy"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20220501preview:ExtendedServerBlobAuditingPolicy"),
		},
	})
	opts = append(opts, aliases)
	var resource ExtendedServerBlobAuditingPolicy
	err := ctx.RegisterResource("azure-native:sql/v20210201preview:ExtendedServerBlobAuditingPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetExtendedServerBlobAuditingPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ExtendedServerBlobAuditingPolicyState, opts ...pulumi.ResourceOption) (*ExtendedServerBlobAuditingPolicy, error) {
	var resource ExtendedServerBlobAuditingPolicy
	err := ctx.ReadResource("azure-native:sql/v20210201preview:ExtendedServerBlobAuditingPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type extendedServerBlobAuditingPolicyState struct {
}

type ExtendedServerBlobAuditingPolicyState struct {
}

func (ExtendedServerBlobAuditingPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*extendedServerBlobAuditingPolicyState)(nil)).Elem()
}

type extendedServerBlobAuditingPolicyArgs struct {
	AuditActionsAndGroups        []string                `pulumi:"auditActionsAndGroups"`
	BlobAuditingPolicyName       *string                 `pulumi:"blobAuditingPolicyName"`
	IsAzureMonitorTargetEnabled  *bool                   `pulumi:"isAzureMonitorTargetEnabled"`
	IsDevopsAuditEnabled         *bool                   `pulumi:"isDevopsAuditEnabled"`
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


type ExtendedServerBlobAuditingPolicyArgs struct {
	AuditActionsAndGroups        pulumi.StringArrayInput
	BlobAuditingPolicyName       pulumi.StringPtrInput
	IsAzureMonitorTargetEnabled  pulumi.BoolPtrInput
	IsDevopsAuditEnabled         pulumi.BoolPtrInput
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

func (ExtendedServerBlobAuditingPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*extendedServerBlobAuditingPolicyArgs)(nil)).Elem()
}

type ExtendedServerBlobAuditingPolicyInput interface {
	pulumi.Input

	ToExtendedServerBlobAuditingPolicyOutput() ExtendedServerBlobAuditingPolicyOutput
	ToExtendedServerBlobAuditingPolicyOutputWithContext(ctx context.Context) ExtendedServerBlobAuditingPolicyOutput
}

func (*ExtendedServerBlobAuditingPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**ExtendedServerBlobAuditingPolicy)(nil)).Elem()
}

func (i *ExtendedServerBlobAuditingPolicy) ToExtendedServerBlobAuditingPolicyOutput() ExtendedServerBlobAuditingPolicyOutput {
	return i.ToExtendedServerBlobAuditingPolicyOutputWithContext(context.Background())
}

func (i *ExtendedServerBlobAuditingPolicy) ToExtendedServerBlobAuditingPolicyOutputWithContext(ctx context.Context) ExtendedServerBlobAuditingPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExtendedServerBlobAuditingPolicyOutput)
}

type ExtendedServerBlobAuditingPolicyOutput struct{ *pulumi.OutputState }

func (ExtendedServerBlobAuditingPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ExtendedServerBlobAuditingPolicy)(nil)).Elem()
}

func (o ExtendedServerBlobAuditingPolicyOutput) ToExtendedServerBlobAuditingPolicyOutput() ExtendedServerBlobAuditingPolicyOutput {
	return o
}

func (o ExtendedServerBlobAuditingPolicyOutput) ToExtendedServerBlobAuditingPolicyOutputWithContext(ctx context.Context) ExtendedServerBlobAuditingPolicyOutput {
	return o
}

func (o ExtendedServerBlobAuditingPolicyOutput) AuditActionsAndGroups() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ExtendedServerBlobAuditingPolicy) pulumi.StringArrayOutput { return v.AuditActionsAndGroups }).(pulumi.StringArrayOutput)
}

func (o ExtendedServerBlobAuditingPolicyOutput) IsAzureMonitorTargetEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ExtendedServerBlobAuditingPolicy) pulumi.BoolPtrOutput { return v.IsAzureMonitorTargetEnabled }).(pulumi.BoolPtrOutput)
}

func (o ExtendedServerBlobAuditingPolicyOutput) IsDevopsAuditEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ExtendedServerBlobAuditingPolicy) pulumi.BoolPtrOutput { return v.IsDevopsAuditEnabled }).(pulumi.BoolPtrOutput)
}

func (o ExtendedServerBlobAuditingPolicyOutput) IsStorageSecondaryKeyInUse() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ExtendedServerBlobAuditingPolicy) pulumi.BoolPtrOutput { return v.IsStorageSecondaryKeyInUse }).(pulumi.BoolPtrOutput)
}

func (o ExtendedServerBlobAuditingPolicyOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ExtendedServerBlobAuditingPolicy) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ExtendedServerBlobAuditingPolicyOutput) PredicateExpression() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ExtendedServerBlobAuditingPolicy) pulumi.StringPtrOutput { return v.PredicateExpression }).(pulumi.StringPtrOutput)
}

func (o ExtendedServerBlobAuditingPolicyOutput) QueueDelayMs() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ExtendedServerBlobAuditingPolicy) pulumi.IntPtrOutput { return v.QueueDelayMs }).(pulumi.IntPtrOutput)
}

func (o ExtendedServerBlobAuditingPolicyOutput) RetentionDays() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ExtendedServerBlobAuditingPolicy) pulumi.IntPtrOutput { return v.RetentionDays }).(pulumi.IntPtrOutput)
}

func (o ExtendedServerBlobAuditingPolicyOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v *ExtendedServerBlobAuditingPolicy) pulumi.StringOutput { return v.State }).(pulumi.StringOutput)
}

func (o ExtendedServerBlobAuditingPolicyOutput) StorageAccountSubscriptionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ExtendedServerBlobAuditingPolicy) pulumi.StringPtrOutput {
		return v.StorageAccountSubscriptionId
	}).(pulumi.StringPtrOutput)
}

func (o ExtendedServerBlobAuditingPolicyOutput) StorageEndpoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ExtendedServerBlobAuditingPolicy) pulumi.StringPtrOutput { return v.StorageEndpoint }).(pulumi.StringPtrOutput)
}

func (o ExtendedServerBlobAuditingPolicyOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ExtendedServerBlobAuditingPolicy) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ExtendedServerBlobAuditingPolicyOutput{})
}
