


package v20210501preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type LongTermRetentionPolicy struct {
	pulumi.CustomResourceState

	MonthlyRetention pulumi.StringPtrOutput `pulumi:"monthlyRetention"`
	Name             pulumi.StringOutput    `pulumi:"name"`
	Type             pulumi.StringOutput    `pulumi:"type"`
	WeekOfYear       pulumi.IntPtrOutput    `pulumi:"weekOfYear"`
	WeeklyRetention  pulumi.StringPtrOutput `pulumi:"weeklyRetention"`
	YearlyRetention  pulumi.StringPtrOutput `pulumi:"yearlyRetention"`
}


func NewLongTermRetentionPolicy(ctx *pulumi.Context,
	name string, args *LongTermRetentionPolicyArgs, opts ...pulumi.ResourceOption) (*LongTermRetentionPolicy, error) {
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
			Type: pulumi.String("azure-native:sql:LongTermRetentionPolicy"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20170301preview:LongTermRetentionPolicy"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20200202preview:LongTermRetentionPolicy"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20200801preview:LongTermRetentionPolicy"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20201101preview:LongTermRetentionPolicy"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20210201preview:LongTermRetentionPolicy"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20210801preview:LongTermRetentionPolicy"),
		},
	})
	opts = append(opts, aliases)
	var resource LongTermRetentionPolicy
	err := ctx.RegisterResource("azure-native:sql/v20210501preview:LongTermRetentionPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetLongTermRetentionPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LongTermRetentionPolicyState, opts ...pulumi.ResourceOption) (*LongTermRetentionPolicy, error) {
	var resource LongTermRetentionPolicy
	err := ctx.ReadResource("azure-native:sql/v20210501preview:LongTermRetentionPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type longTermRetentionPolicyState struct {
}

type LongTermRetentionPolicyState struct {
}

func (LongTermRetentionPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*longTermRetentionPolicyState)(nil)).Elem()
}

type longTermRetentionPolicyArgs struct {
	DatabaseName      string  `pulumi:"databaseName"`
	MonthlyRetention  *string `pulumi:"monthlyRetention"`
	PolicyName        *string `pulumi:"policyName"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
	ServerName        string  `pulumi:"serverName"`
	WeekOfYear        *int    `pulumi:"weekOfYear"`
	WeeklyRetention   *string `pulumi:"weeklyRetention"`
	YearlyRetention   *string `pulumi:"yearlyRetention"`
}


type LongTermRetentionPolicyArgs struct {
	DatabaseName      pulumi.StringInput
	MonthlyRetention  pulumi.StringPtrInput
	PolicyName        pulumi.StringPtrInput
	ResourceGroupName pulumi.StringInput
	ServerName        pulumi.StringInput
	WeekOfYear        pulumi.IntPtrInput
	WeeklyRetention   pulumi.StringPtrInput
	YearlyRetention   pulumi.StringPtrInput
}

func (LongTermRetentionPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*longTermRetentionPolicyArgs)(nil)).Elem()
}

type LongTermRetentionPolicyInput interface {
	pulumi.Input

	ToLongTermRetentionPolicyOutput() LongTermRetentionPolicyOutput
	ToLongTermRetentionPolicyOutputWithContext(ctx context.Context) LongTermRetentionPolicyOutput
}

func (*LongTermRetentionPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**LongTermRetentionPolicy)(nil)).Elem()
}

func (i *LongTermRetentionPolicy) ToLongTermRetentionPolicyOutput() LongTermRetentionPolicyOutput {
	return i.ToLongTermRetentionPolicyOutputWithContext(context.Background())
}

func (i *LongTermRetentionPolicy) ToLongTermRetentionPolicyOutputWithContext(ctx context.Context) LongTermRetentionPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LongTermRetentionPolicyOutput)
}

type LongTermRetentionPolicyOutput struct{ *pulumi.OutputState }

func (LongTermRetentionPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LongTermRetentionPolicy)(nil)).Elem()
}

func (o LongTermRetentionPolicyOutput) ToLongTermRetentionPolicyOutput() LongTermRetentionPolicyOutput {
	return o
}

func (o LongTermRetentionPolicyOutput) ToLongTermRetentionPolicyOutputWithContext(ctx context.Context) LongTermRetentionPolicyOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(LongTermRetentionPolicyOutput{})
}
