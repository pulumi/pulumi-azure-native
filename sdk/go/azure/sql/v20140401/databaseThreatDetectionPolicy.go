


package v20140401

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type DatabaseThreatDetectionPolicy struct {
	pulumi.CustomResourceState

	DisabledAlerts     pulumi.StringPtrOutput `pulumi:"disabledAlerts"`
	EmailAccountAdmins pulumi.StringPtrOutput `pulumi:"emailAccountAdmins"`
	EmailAddresses     pulumi.StringPtrOutput `pulumi:"emailAddresses"`
	Kind               pulumi.StringOutput    `pulumi:"kind"`
	Location           pulumi.StringPtrOutput `pulumi:"location"`
	Name               pulumi.StringOutput    `pulumi:"name"`
	RetentionDays      pulumi.IntPtrOutput    `pulumi:"retentionDays"`
	State              pulumi.StringOutput    `pulumi:"state"`
	StorageEndpoint    pulumi.StringPtrOutput `pulumi:"storageEndpoint"`
	Type               pulumi.StringOutput    `pulumi:"type"`
	UseServerDefault   pulumi.StringPtrOutput `pulumi:"useServerDefault"`
}


func NewDatabaseThreatDetectionPolicy(ctx *pulumi.Context,
	name string, args *DatabaseThreatDetectionPolicyArgs, opts ...pulumi.ResourceOption) (*DatabaseThreatDetectionPolicy, error) {
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
			Type: pulumi.String("azure-nextgen:sql/v20140401:DatabaseThreatDetectionPolicy"),
		},
		{
			Type: pulumi.String("azure-native:sql:DatabaseThreatDetectionPolicy"),
		},
		{
			Type: pulumi.String("azure-nextgen:sql:DatabaseThreatDetectionPolicy"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20180601preview:DatabaseThreatDetectionPolicy"),
		},
		{
			Type: pulumi.String("azure-nextgen:sql/v20180601preview:DatabaseThreatDetectionPolicy"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20200202preview:DatabaseThreatDetectionPolicy"),
		},
		{
			Type: pulumi.String("azure-nextgen:sql/v20200202preview:DatabaseThreatDetectionPolicy"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20200801preview:DatabaseThreatDetectionPolicy"),
		},
		{
			Type: pulumi.String("azure-nextgen:sql/v20200801preview:DatabaseThreatDetectionPolicy"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20201101preview:DatabaseThreatDetectionPolicy"),
		},
		{
			Type: pulumi.String("azure-nextgen:sql/v20201101preview:DatabaseThreatDetectionPolicy"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20210201preview:DatabaseThreatDetectionPolicy"),
		},
		{
			Type: pulumi.String("azure-nextgen:sql/v20210201preview:DatabaseThreatDetectionPolicy"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20210501preview:DatabaseThreatDetectionPolicy"),
		},
		{
			Type: pulumi.String("azure-nextgen:sql/v20210501preview:DatabaseThreatDetectionPolicy"),
		},
	})
	opts = append(opts, aliases)
	var resource DatabaseThreatDetectionPolicy
	err := ctx.RegisterResource("azure-native:sql/v20140401:DatabaseThreatDetectionPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetDatabaseThreatDetectionPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DatabaseThreatDetectionPolicyState, opts ...pulumi.ResourceOption) (*DatabaseThreatDetectionPolicy, error) {
	var resource DatabaseThreatDetectionPolicy
	err := ctx.ReadResource("azure-native:sql/v20140401:DatabaseThreatDetectionPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type databaseThreatDetectionPolicyState struct {
}

type DatabaseThreatDetectionPolicyState struct {
}

func (DatabaseThreatDetectionPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*databaseThreatDetectionPolicyState)(nil)).Elem()
}

type databaseThreatDetectionPolicyArgs struct {
	DatabaseName            string  `pulumi:"databaseName"`
	DisabledAlerts          *string `pulumi:"disabledAlerts"`
	EmailAccountAdmins      *string `pulumi:"emailAccountAdmins"`
	EmailAddresses          *string `pulumi:"emailAddresses"`
	Location                *string `pulumi:"location"`
	ResourceGroupName       string  `pulumi:"resourceGroupName"`
	RetentionDays           *int    `pulumi:"retentionDays"`
	SecurityAlertPolicyName *string `pulumi:"securityAlertPolicyName"`
	ServerName              string  `pulumi:"serverName"`
	State                   string  `pulumi:"state"`
	StorageAccountAccessKey *string `pulumi:"storageAccountAccessKey"`
	StorageEndpoint         *string `pulumi:"storageEndpoint"`
	UseServerDefault        *string `pulumi:"useServerDefault"`
}


type DatabaseThreatDetectionPolicyArgs struct {
	DatabaseName            pulumi.StringInput
	DisabledAlerts          pulumi.StringPtrInput
	EmailAccountAdmins      pulumi.StringPtrInput
	EmailAddresses          pulumi.StringPtrInput
	Location                pulumi.StringPtrInput
	ResourceGroupName       pulumi.StringInput
	RetentionDays           pulumi.IntPtrInput
	SecurityAlertPolicyName pulumi.StringPtrInput
	ServerName              pulumi.StringInput
	State                   pulumi.StringInput
	StorageAccountAccessKey pulumi.StringPtrInput
	StorageEndpoint         pulumi.StringPtrInput
	UseServerDefault        pulumi.StringPtrInput
}

func (DatabaseThreatDetectionPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*databaseThreatDetectionPolicyArgs)(nil)).Elem()
}

type DatabaseThreatDetectionPolicyInput interface {
	pulumi.Input

	ToDatabaseThreatDetectionPolicyOutput() DatabaseThreatDetectionPolicyOutput
	ToDatabaseThreatDetectionPolicyOutputWithContext(ctx context.Context) DatabaseThreatDetectionPolicyOutput
}

func (*DatabaseThreatDetectionPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((*DatabaseThreatDetectionPolicy)(nil))
}

func (i *DatabaseThreatDetectionPolicy) ToDatabaseThreatDetectionPolicyOutput() DatabaseThreatDetectionPolicyOutput {
	return i.ToDatabaseThreatDetectionPolicyOutputWithContext(context.Background())
}

func (i *DatabaseThreatDetectionPolicy) ToDatabaseThreatDetectionPolicyOutputWithContext(ctx context.Context) DatabaseThreatDetectionPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatabaseThreatDetectionPolicyOutput)
}

type DatabaseThreatDetectionPolicyOutput struct{ *pulumi.OutputState }

func (DatabaseThreatDetectionPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DatabaseThreatDetectionPolicy)(nil))
}

func (o DatabaseThreatDetectionPolicyOutput) ToDatabaseThreatDetectionPolicyOutput() DatabaseThreatDetectionPolicyOutput {
	return o
}

func (o DatabaseThreatDetectionPolicyOutput) ToDatabaseThreatDetectionPolicyOutputWithContext(ctx context.Context) DatabaseThreatDetectionPolicyOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(DatabaseThreatDetectionPolicyOutput{})
}
