


package v20170301preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type BackupLongTermRetentionPolicy struct {
	pulumi.CustomResourceState

	MonthlyRetention pulumi.StringPtrOutput `pulumi:"monthlyRetention"`
	Name             pulumi.StringOutput    `pulumi:"name"`
	Type             pulumi.StringOutput    `pulumi:"type"`
	WeekOfYear       pulumi.IntPtrOutput    `pulumi:"weekOfYear"`
	WeeklyRetention  pulumi.StringPtrOutput `pulumi:"weeklyRetention"`
	YearlyRetention  pulumi.StringPtrOutput `pulumi:"yearlyRetention"`
}


func NewBackupLongTermRetentionPolicy(ctx *pulumi.Context,
	name string, args *BackupLongTermRetentionPolicyArgs, opts ...pulumi.ResourceOption) (*BackupLongTermRetentionPolicy, error) {
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
			Type: pulumi.String("azure-nextgen:sql/v20170301preview:BackupLongTermRetentionPolicy"),
		},
		{
			Type: pulumi.String("azure-native:sql:BackupLongTermRetentionPolicy"),
		},
		{
			Type: pulumi.String("azure-nextgen:sql:BackupLongTermRetentionPolicy"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20200202preview:BackupLongTermRetentionPolicy"),
		},
		{
			Type: pulumi.String("azure-nextgen:sql/v20200202preview:BackupLongTermRetentionPolicy"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20200801preview:BackupLongTermRetentionPolicy"),
		},
		{
			Type: pulumi.String("azure-nextgen:sql/v20200801preview:BackupLongTermRetentionPolicy"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20201101preview:BackupLongTermRetentionPolicy"),
		},
		{
			Type: pulumi.String("azure-nextgen:sql/v20201101preview:BackupLongTermRetentionPolicy"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20210201preview:BackupLongTermRetentionPolicy"),
		},
		{
			Type: pulumi.String("azure-nextgen:sql/v20210201preview:BackupLongTermRetentionPolicy"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20210501preview:BackupLongTermRetentionPolicy"),
		},
		{
			Type: pulumi.String("azure-nextgen:sql/v20210501preview:BackupLongTermRetentionPolicy"),
		},
	})
	opts = append(opts, aliases)
	var resource BackupLongTermRetentionPolicy
	err := ctx.RegisterResource("azure-native:sql/v20170301preview:BackupLongTermRetentionPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetBackupLongTermRetentionPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BackupLongTermRetentionPolicyState, opts ...pulumi.ResourceOption) (*BackupLongTermRetentionPolicy, error) {
	var resource BackupLongTermRetentionPolicy
	err := ctx.ReadResource("azure-native:sql/v20170301preview:BackupLongTermRetentionPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type backupLongTermRetentionPolicyState struct {
}

type BackupLongTermRetentionPolicyState struct {
}

func (BackupLongTermRetentionPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*backupLongTermRetentionPolicyState)(nil)).Elem()
}

type backupLongTermRetentionPolicyArgs struct {
	DatabaseName      string  `pulumi:"databaseName"`
	MonthlyRetention  *string `pulumi:"monthlyRetention"`
	PolicyName        *string `pulumi:"policyName"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
	ServerName        string  `pulumi:"serverName"`
	WeekOfYear        *int    `pulumi:"weekOfYear"`
	WeeklyRetention   *string `pulumi:"weeklyRetention"`
	YearlyRetention   *string `pulumi:"yearlyRetention"`
}


type BackupLongTermRetentionPolicyArgs struct {
	DatabaseName      pulumi.StringInput
	MonthlyRetention  pulumi.StringPtrInput
	PolicyName        pulumi.StringPtrInput
	ResourceGroupName pulumi.StringInput
	ServerName        pulumi.StringInput
	WeekOfYear        pulumi.IntPtrInput
	WeeklyRetention   pulumi.StringPtrInput
	YearlyRetention   pulumi.StringPtrInput
}

func (BackupLongTermRetentionPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*backupLongTermRetentionPolicyArgs)(nil)).Elem()
}

type BackupLongTermRetentionPolicyInput interface {
	pulumi.Input

	ToBackupLongTermRetentionPolicyOutput() BackupLongTermRetentionPolicyOutput
	ToBackupLongTermRetentionPolicyOutputWithContext(ctx context.Context) BackupLongTermRetentionPolicyOutput
}

func (*BackupLongTermRetentionPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((*BackupLongTermRetentionPolicy)(nil))
}

func (i *BackupLongTermRetentionPolicy) ToBackupLongTermRetentionPolicyOutput() BackupLongTermRetentionPolicyOutput {
	return i.ToBackupLongTermRetentionPolicyOutputWithContext(context.Background())
}

func (i *BackupLongTermRetentionPolicy) ToBackupLongTermRetentionPolicyOutputWithContext(ctx context.Context) BackupLongTermRetentionPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackupLongTermRetentionPolicyOutput)
}

type BackupLongTermRetentionPolicyOutput struct{ *pulumi.OutputState }

func (BackupLongTermRetentionPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BackupLongTermRetentionPolicy)(nil))
}

func (o BackupLongTermRetentionPolicyOutput) ToBackupLongTermRetentionPolicyOutput() BackupLongTermRetentionPolicyOutput {
	return o
}

func (o BackupLongTermRetentionPolicyOutput) ToBackupLongTermRetentionPolicyOutputWithContext(ctx context.Context) BackupLongTermRetentionPolicyOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(BackupLongTermRetentionPolicyOutput{})
}
