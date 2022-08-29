


package v20200801preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type DatabaseSecurityAlertPolicy struct {
	pulumi.CustomResourceState

	CreationTime            pulumi.StringOutput      `pulumi:"creationTime"`
	DisabledAlerts          pulumi.StringArrayOutput `pulumi:"disabledAlerts"`
	EmailAccountAdmins      pulumi.BoolPtrOutput     `pulumi:"emailAccountAdmins"`
	EmailAddresses          pulumi.StringArrayOutput `pulumi:"emailAddresses"`
	Name                    pulumi.StringOutput      `pulumi:"name"`
	RetentionDays           pulumi.IntPtrOutput      `pulumi:"retentionDays"`
	State                   pulumi.StringOutput      `pulumi:"state"`
	StorageAccountAccessKey pulumi.StringPtrOutput   `pulumi:"storageAccountAccessKey"`
	StorageEndpoint         pulumi.StringPtrOutput   `pulumi:"storageEndpoint"`
	SystemData              SystemDataResponseOutput `pulumi:"systemData"`
	Type                    pulumi.StringOutput      `pulumi:"type"`
}


func NewDatabaseSecurityAlertPolicy(ctx *pulumi.Context,
	name string, args *DatabaseSecurityAlertPolicyArgs, opts ...pulumi.ResourceOption) (*DatabaseSecurityAlertPolicy, error) {
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
			Type: pulumi.String("azure-native:sql:DatabaseSecurityAlertPolicy"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20140401:DatabaseSecurityAlertPolicy"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20180601preview:DatabaseSecurityAlertPolicy"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20200202preview:DatabaseSecurityAlertPolicy"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20201101preview:DatabaseSecurityAlertPolicy"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20210201preview:DatabaseSecurityAlertPolicy"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20210501preview:DatabaseSecurityAlertPolicy"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20210801preview:DatabaseSecurityAlertPolicy"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20211101:DatabaseSecurityAlertPolicy"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20211101preview:DatabaseSecurityAlertPolicy"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20220201preview:DatabaseSecurityAlertPolicy"),
		},
	})
	opts = append(opts, aliases)
	var resource DatabaseSecurityAlertPolicy
	err := ctx.RegisterResource("azure-native:sql/v20200801preview:DatabaseSecurityAlertPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetDatabaseSecurityAlertPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DatabaseSecurityAlertPolicyState, opts ...pulumi.ResourceOption) (*DatabaseSecurityAlertPolicy, error) {
	var resource DatabaseSecurityAlertPolicy
	err := ctx.ReadResource("azure-native:sql/v20200801preview:DatabaseSecurityAlertPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type databaseSecurityAlertPolicyState struct {
}

type DatabaseSecurityAlertPolicyState struct {
}

func (DatabaseSecurityAlertPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*databaseSecurityAlertPolicyState)(nil)).Elem()
}

type databaseSecurityAlertPolicyArgs struct {
	DatabaseName            string                    `pulumi:"databaseName"`
	DisabledAlerts          []string                  `pulumi:"disabledAlerts"`
	EmailAccountAdmins      *bool                     `pulumi:"emailAccountAdmins"`
	EmailAddresses          []string                  `pulumi:"emailAddresses"`
	ResourceGroupName       string                    `pulumi:"resourceGroupName"`
	RetentionDays           *int                      `pulumi:"retentionDays"`
	SecurityAlertPolicyName *string                   `pulumi:"securityAlertPolicyName"`
	ServerName              string                    `pulumi:"serverName"`
	State                   SecurityAlertsPolicyState `pulumi:"state"`
	StorageAccountAccessKey *string                   `pulumi:"storageAccountAccessKey"`
	StorageEndpoint         *string                   `pulumi:"storageEndpoint"`
}


type DatabaseSecurityAlertPolicyArgs struct {
	DatabaseName            pulumi.StringInput
	DisabledAlerts          pulumi.StringArrayInput
	EmailAccountAdmins      pulumi.BoolPtrInput
	EmailAddresses          pulumi.StringArrayInput
	ResourceGroupName       pulumi.StringInput
	RetentionDays           pulumi.IntPtrInput
	SecurityAlertPolicyName pulumi.StringPtrInput
	ServerName              pulumi.StringInput
	State                   SecurityAlertsPolicyStateInput
	StorageAccountAccessKey pulumi.StringPtrInput
	StorageEndpoint         pulumi.StringPtrInput
}

func (DatabaseSecurityAlertPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*databaseSecurityAlertPolicyArgs)(nil)).Elem()
}

type DatabaseSecurityAlertPolicyInput interface {
	pulumi.Input

	ToDatabaseSecurityAlertPolicyOutput() DatabaseSecurityAlertPolicyOutput
	ToDatabaseSecurityAlertPolicyOutputWithContext(ctx context.Context) DatabaseSecurityAlertPolicyOutput
}

func (*DatabaseSecurityAlertPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**DatabaseSecurityAlertPolicy)(nil)).Elem()
}

func (i *DatabaseSecurityAlertPolicy) ToDatabaseSecurityAlertPolicyOutput() DatabaseSecurityAlertPolicyOutput {
	return i.ToDatabaseSecurityAlertPolicyOutputWithContext(context.Background())
}

func (i *DatabaseSecurityAlertPolicy) ToDatabaseSecurityAlertPolicyOutputWithContext(ctx context.Context) DatabaseSecurityAlertPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatabaseSecurityAlertPolicyOutput)
}

type DatabaseSecurityAlertPolicyOutput struct{ *pulumi.OutputState }

func (DatabaseSecurityAlertPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DatabaseSecurityAlertPolicy)(nil)).Elem()
}

func (o DatabaseSecurityAlertPolicyOutput) ToDatabaseSecurityAlertPolicyOutput() DatabaseSecurityAlertPolicyOutput {
	return o
}

func (o DatabaseSecurityAlertPolicyOutput) ToDatabaseSecurityAlertPolicyOutputWithContext(ctx context.Context) DatabaseSecurityAlertPolicyOutput {
	return o
}

func (o DatabaseSecurityAlertPolicyOutput) CreationTime() pulumi.StringOutput {
	return o.ApplyT(func(v *DatabaseSecurityAlertPolicy) pulumi.StringOutput { return v.CreationTime }).(pulumi.StringOutput)
}

func (o DatabaseSecurityAlertPolicyOutput) DisabledAlerts() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *DatabaseSecurityAlertPolicy) pulumi.StringArrayOutput { return v.DisabledAlerts }).(pulumi.StringArrayOutput)
}

func (o DatabaseSecurityAlertPolicyOutput) EmailAccountAdmins() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *DatabaseSecurityAlertPolicy) pulumi.BoolPtrOutput { return v.EmailAccountAdmins }).(pulumi.BoolPtrOutput)
}

func (o DatabaseSecurityAlertPolicyOutput) EmailAddresses() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *DatabaseSecurityAlertPolicy) pulumi.StringArrayOutput { return v.EmailAddresses }).(pulumi.StringArrayOutput)
}

func (o DatabaseSecurityAlertPolicyOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *DatabaseSecurityAlertPolicy) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o DatabaseSecurityAlertPolicyOutput) RetentionDays() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *DatabaseSecurityAlertPolicy) pulumi.IntPtrOutput { return v.RetentionDays }).(pulumi.IntPtrOutput)
}

func (o DatabaseSecurityAlertPolicyOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v *DatabaseSecurityAlertPolicy) pulumi.StringOutput { return v.State }).(pulumi.StringOutput)
}

func (o DatabaseSecurityAlertPolicyOutput) StorageAccountAccessKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DatabaseSecurityAlertPolicy) pulumi.StringPtrOutput { return v.StorageAccountAccessKey }).(pulumi.StringPtrOutput)
}

func (o DatabaseSecurityAlertPolicyOutput) StorageEndpoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DatabaseSecurityAlertPolicy) pulumi.StringPtrOutput { return v.StorageEndpoint }).(pulumi.StringPtrOutput)
}

func (o DatabaseSecurityAlertPolicyOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *DatabaseSecurityAlertPolicy) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o DatabaseSecurityAlertPolicyOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *DatabaseSecurityAlertPolicy) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(DatabaseSecurityAlertPolicyOutput{})
}
