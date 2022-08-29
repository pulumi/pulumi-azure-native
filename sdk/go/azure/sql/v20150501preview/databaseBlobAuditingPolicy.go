


package v20150501preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type DatabaseBlobAuditingPolicy struct {
	pulumi.CustomResourceState

	AuditActionsAndGroups        pulumi.StringArrayOutput `pulumi:"auditActionsAndGroups"`
	IsStorageSecondaryKeyInUse   pulumi.BoolPtrOutput     `pulumi:"isStorageSecondaryKeyInUse"`
	Kind                         pulumi.StringOutput      `pulumi:"kind"`
	Name                         pulumi.StringOutput      `pulumi:"name"`
	RetentionDays                pulumi.IntPtrOutput      `pulumi:"retentionDays"`
	State                        pulumi.StringOutput      `pulumi:"state"`
	StorageAccountSubscriptionId pulumi.StringPtrOutput   `pulumi:"storageAccountSubscriptionId"`
	StorageEndpoint              pulumi.StringPtrOutput   `pulumi:"storageEndpoint"`
	Type                         pulumi.StringOutput      `pulumi:"type"`
}


func NewDatabaseBlobAuditingPolicy(ctx *pulumi.Context,
	name string, args *DatabaseBlobAuditingPolicyArgs, opts ...pulumi.ResourceOption) (*DatabaseBlobAuditingPolicy, error) {
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
			Type: pulumi.String("azure-native:sql:DatabaseBlobAuditingPolicy"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20170301preview:DatabaseBlobAuditingPolicy"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20200202preview:DatabaseBlobAuditingPolicy"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20200801preview:DatabaseBlobAuditingPolicy"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20201101preview:DatabaseBlobAuditingPolicy"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20210201preview:DatabaseBlobAuditingPolicy"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20210501preview:DatabaseBlobAuditingPolicy"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20210801preview:DatabaseBlobAuditingPolicy"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20211101:DatabaseBlobAuditingPolicy"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20211101preview:DatabaseBlobAuditingPolicy"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20220201preview:DatabaseBlobAuditingPolicy"),
		},
	})
	opts = append(opts, aliases)
	var resource DatabaseBlobAuditingPolicy
	err := ctx.RegisterResource("azure-native:sql/v20150501preview:DatabaseBlobAuditingPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetDatabaseBlobAuditingPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DatabaseBlobAuditingPolicyState, opts ...pulumi.ResourceOption) (*DatabaseBlobAuditingPolicy, error) {
	var resource DatabaseBlobAuditingPolicy
	err := ctx.ReadResource("azure-native:sql/v20150501preview:DatabaseBlobAuditingPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type databaseBlobAuditingPolicyState struct {
}

type DatabaseBlobAuditingPolicyState struct {
}

func (DatabaseBlobAuditingPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*databaseBlobAuditingPolicyState)(nil)).Elem()
}

type databaseBlobAuditingPolicyArgs struct {
	AuditActionsAndGroups        []string                `pulumi:"auditActionsAndGroups"`
	BlobAuditingPolicyName       *string                 `pulumi:"blobAuditingPolicyName"`
	DatabaseName                 string                  `pulumi:"databaseName"`
	IsStorageSecondaryKeyInUse   *bool                   `pulumi:"isStorageSecondaryKeyInUse"`
	ResourceGroupName            string                  `pulumi:"resourceGroupName"`
	RetentionDays                *int                    `pulumi:"retentionDays"`
	ServerName                   string                  `pulumi:"serverName"`
	State                        BlobAuditingPolicyState `pulumi:"state"`
	StorageAccountAccessKey      *string                 `pulumi:"storageAccountAccessKey"`
	StorageAccountSubscriptionId *string                 `pulumi:"storageAccountSubscriptionId"`
	StorageEndpoint              *string                 `pulumi:"storageEndpoint"`
}


type DatabaseBlobAuditingPolicyArgs struct {
	AuditActionsAndGroups        pulumi.StringArrayInput
	BlobAuditingPolicyName       pulumi.StringPtrInput
	DatabaseName                 pulumi.StringInput
	IsStorageSecondaryKeyInUse   pulumi.BoolPtrInput
	ResourceGroupName            pulumi.StringInput
	RetentionDays                pulumi.IntPtrInput
	ServerName                   pulumi.StringInput
	State                        BlobAuditingPolicyStateInput
	StorageAccountAccessKey      pulumi.StringPtrInput
	StorageAccountSubscriptionId pulumi.StringPtrInput
	StorageEndpoint              pulumi.StringPtrInput
}

func (DatabaseBlobAuditingPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*databaseBlobAuditingPolicyArgs)(nil)).Elem()
}

type DatabaseBlobAuditingPolicyInput interface {
	pulumi.Input

	ToDatabaseBlobAuditingPolicyOutput() DatabaseBlobAuditingPolicyOutput
	ToDatabaseBlobAuditingPolicyOutputWithContext(ctx context.Context) DatabaseBlobAuditingPolicyOutput
}

func (*DatabaseBlobAuditingPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**DatabaseBlobAuditingPolicy)(nil)).Elem()
}

func (i *DatabaseBlobAuditingPolicy) ToDatabaseBlobAuditingPolicyOutput() DatabaseBlobAuditingPolicyOutput {
	return i.ToDatabaseBlobAuditingPolicyOutputWithContext(context.Background())
}

func (i *DatabaseBlobAuditingPolicy) ToDatabaseBlobAuditingPolicyOutputWithContext(ctx context.Context) DatabaseBlobAuditingPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatabaseBlobAuditingPolicyOutput)
}

type DatabaseBlobAuditingPolicyOutput struct{ *pulumi.OutputState }

func (DatabaseBlobAuditingPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DatabaseBlobAuditingPolicy)(nil)).Elem()
}

func (o DatabaseBlobAuditingPolicyOutput) ToDatabaseBlobAuditingPolicyOutput() DatabaseBlobAuditingPolicyOutput {
	return o
}

func (o DatabaseBlobAuditingPolicyOutput) ToDatabaseBlobAuditingPolicyOutputWithContext(ctx context.Context) DatabaseBlobAuditingPolicyOutput {
	return o
}

func (o DatabaseBlobAuditingPolicyOutput) AuditActionsAndGroups() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *DatabaseBlobAuditingPolicy) pulumi.StringArrayOutput { return v.AuditActionsAndGroups }).(pulumi.StringArrayOutput)
}

func (o DatabaseBlobAuditingPolicyOutput) IsStorageSecondaryKeyInUse() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *DatabaseBlobAuditingPolicy) pulumi.BoolPtrOutput { return v.IsStorageSecondaryKeyInUse }).(pulumi.BoolPtrOutput)
}

func (o DatabaseBlobAuditingPolicyOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *DatabaseBlobAuditingPolicy) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

func (o DatabaseBlobAuditingPolicyOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *DatabaseBlobAuditingPolicy) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o DatabaseBlobAuditingPolicyOutput) RetentionDays() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *DatabaseBlobAuditingPolicy) pulumi.IntPtrOutput { return v.RetentionDays }).(pulumi.IntPtrOutput)
}

func (o DatabaseBlobAuditingPolicyOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v *DatabaseBlobAuditingPolicy) pulumi.StringOutput { return v.State }).(pulumi.StringOutput)
}

func (o DatabaseBlobAuditingPolicyOutput) StorageAccountSubscriptionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DatabaseBlobAuditingPolicy) pulumi.StringPtrOutput { return v.StorageAccountSubscriptionId }).(pulumi.StringPtrOutput)
}

func (o DatabaseBlobAuditingPolicyOutput) StorageEndpoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DatabaseBlobAuditingPolicy) pulumi.StringPtrOutput { return v.StorageEndpoint }).(pulumi.StringPtrOutput)
}

func (o DatabaseBlobAuditingPolicyOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *DatabaseBlobAuditingPolicy) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(DatabaseBlobAuditingPolicyOutput{})
}
