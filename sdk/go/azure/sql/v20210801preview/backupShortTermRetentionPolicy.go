


package v20210801preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type BackupShortTermRetentionPolicy struct {
	pulumi.CustomResourceState

	DiffBackupIntervalInHours pulumi.IntPtrOutput `pulumi:"diffBackupIntervalInHours"`
	Name                      pulumi.StringOutput `pulumi:"name"`
	RetentionDays             pulumi.IntPtrOutput `pulumi:"retentionDays"`
	Type                      pulumi.StringOutput `pulumi:"type"`
}


func NewBackupShortTermRetentionPolicy(ctx *pulumi.Context,
	name string, args *BackupShortTermRetentionPolicyArgs, opts ...pulumi.ResourceOption) (*BackupShortTermRetentionPolicy, error) {
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
			Type: pulumi.String("azure-native:sql:BackupShortTermRetentionPolicy"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20171001preview:BackupShortTermRetentionPolicy"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20200202preview:BackupShortTermRetentionPolicy"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20200801preview:BackupShortTermRetentionPolicy"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20201101preview:BackupShortTermRetentionPolicy"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20210201preview:BackupShortTermRetentionPolicy"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20210501preview:BackupShortTermRetentionPolicy"),
		},
	})
	opts = append(opts, aliases)
	var resource BackupShortTermRetentionPolicy
	err := ctx.RegisterResource("azure-native:sql/v20210801preview:BackupShortTermRetentionPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetBackupShortTermRetentionPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BackupShortTermRetentionPolicyState, opts ...pulumi.ResourceOption) (*BackupShortTermRetentionPolicy, error) {
	var resource BackupShortTermRetentionPolicy
	err := ctx.ReadResource("azure-native:sql/v20210801preview:BackupShortTermRetentionPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type backupShortTermRetentionPolicyState struct {
}

type BackupShortTermRetentionPolicyState struct {
}

func (BackupShortTermRetentionPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*backupShortTermRetentionPolicyState)(nil)).Elem()
}

type backupShortTermRetentionPolicyArgs struct {
	DatabaseName              string  `pulumi:"databaseName"`
	DiffBackupIntervalInHours *int    `pulumi:"diffBackupIntervalInHours"`
	PolicyName                *string `pulumi:"policyName"`
	ResourceGroupName         string  `pulumi:"resourceGroupName"`
	RetentionDays             *int    `pulumi:"retentionDays"`
	ServerName                string  `pulumi:"serverName"`
}


type BackupShortTermRetentionPolicyArgs struct {
	DatabaseName              pulumi.StringInput
	DiffBackupIntervalInHours pulumi.IntPtrInput
	PolicyName                pulumi.StringPtrInput
	ResourceGroupName         pulumi.StringInput
	RetentionDays             pulumi.IntPtrInput
	ServerName                pulumi.StringInput
}

func (BackupShortTermRetentionPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*backupShortTermRetentionPolicyArgs)(nil)).Elem()
}

type BackupShortTermRetentionPolicyInput interface {
	pulumi.Input

	ToBackupShortTermRetentionPolicyOutput() BackupShortTermRetentionPolicyOutput
	ToBackupShortTermRetentionPolicyOutputWithContext(ctx context.Context) BackupShortTermRetentionPolicyOutput
}

func (*BackupShortTermRetentionPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**BackupShortTermRetentionPolicy)(nil)).Elem()
}

func (i *BackupShortTermRetentionPolicy) ToBackupShortTermRetentionPolicyOutput() BackupShortTermRetentionPolicyOutput {
	return i.ToBackupShortTermRetentionPolicyOutputWithContext(context.Background())
}

func (i *BackupShortTermRetentionPolicy) ToBackupShortTermRetentionPolicyOutputWithContext(ctx context.Context) BackupShortTermRetentionPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackupShortTermRetentionPolicyOutput)
}

type BackupShortTermRetentionPolicyOutput struct{ *pulumi.OutputState }

func (BackupShortTermRetentionPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BackupShortTermRetentionPolicy)(nil)).Elem()
}

func (o BackupShortTermRetentionPolicyOutput) ToBackupShortTermRetentionPolicyOutput() BackupShortTermRetentionPolicyOutput {
	return o
}

func (o BackupShortTermRetentionPolicyOutput) ToBackupShortTermRetentionPolicyOutputWithContext(ctx context.Context) BackupShortTermRetentionPolicyOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(BackupShortTermRetentionPolicyOutput{})
}
