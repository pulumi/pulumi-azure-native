


package dbforpostgresql

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ServerSecurityAlertPolicy struct {
	pulumi.CustomResourceState

	DisabledAlerts          pulumi.StringArrayOutput `pulumi:"disabledAlerts"`
	EmailAccountAdmins      pulumi.BoolPtrOutput     `pulumi:"emailAccountAdmins"`
	EmailAddresses          pulumi.StringArrayOutput `pulumi:"emailAddresses"`
	Name                    pulumi.StringOutput      `pulumi:"name"`
	RetentionDays           pulumi.IntPtrOutput      `pulumi:"retentionDays"`
	State                   pulumi.StringOutput      `pulumi:"state"`
	StorageAccountAccessKey pulumi.StringPtrOutput   `pulumi:"storageAccountAccessKey"`
	StorageEndpoint         pulumi.StringPtrOutput   `pulumi:"storageEndpoint"`
	Type                    pulumi.StringOutput      `pulumi:"type"`
}


func NewServerSecurityAlertPolicy(ctx *pulumi.Context,
	name string, args *ServerSecurityAlertPolicyArgs, opts ...pulumi.ResourceOption) (*ServerSecurityAlertPolicy, error) {
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
			Type: pulumi.String("azure-native:dbforpostgresql/v20171201:ServerSecurityAlertPolicy"),
		},
		{
			Type: pulumi.String("azure-native:dbforpostgresql/v20171201preview:ServerSecurityAlertPolicy"),
		},
	})
	opts = append(opts, aliases)
	var resource ServerSecurityAlertPolicy
	err := ctx.RegisterResource("azure-native:dbforpostgresql:ServerSecurityAlertPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetServerSecurityAlertPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServerSecurityAlertPolicyState, opts ...pulumi.ResourceOption) (*ServerSecurityAlertPolicy, error) {
	var resource ServerSecurityAlertPolicy
	err := ctx.ReadResource("azure-native:dbforpostgresql:ServerSecurityAlertPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type serverSecurityAlertPolicyState struct {
}

type ServerSecurityAlertPolicyState struct {
}

func (ServerSecurityAlertPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*serverSecurityAlertPolicyState)(nil)).Elem()
}

type serverSecurityAlertPolicyArgs struct {
	DisabledAlerts          []string                           `pulumi:"disabledAlerts"`
	EmailAccountAdmins      *bool                              `pulumi:"emailAccountAdmins"`
	EmailAddresses          []string                           `pulumi:"emailAddresses"`
	ResourceGroupName       string                             `pulumi:"resourceGroupName"`
	RetentionDays           *int                               `pulumi:"retentionDays"`
	SecurityAlertPolicyName *string                            `pulumi:"securityAlertPolicyName"`
	ServerName              string                             `pulumi:"serverName"`
	State                   ServerSecurityAlertPolicyStateEnum `pulumi:"state"`
	StorageAccountAccessKey *string                            `pulumi:"storageAccountAccessKey"`
	StorageEndpoint         *string                            `pulumi:"storageEndpoint"`
}


type ServerSecurityAlertPolicyArgs struct {
	DisabledAlerts          pulumi.StringArrayInput
	EmailAccountAdmins      pulumi.BoolPtrInput
	EmailAddresses          pulumi.StringArrayInput
	ResourceGroupName       pulumi.StringInput
	RetentionDays           pulumi.IntPtrInput
	SecurityAlertPolicyName pulumi.StringPtrInput
	ServerName              pulumi.StringInput
	State                   ServerSecurityAlertPolicyStateEnumInput
	StorageAccountAccessKey pulumi.StringPtrInput
	StorageEndpoint         pulumi.StringPtrInput
}

func (ServerSecurityAlertPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*serverSecurityAlertPolicyArgs)(nil)).Elem()
}

type ServerSecurityAlertPolicyInput interface {
	pulumi.Input

	ToServerSecurityAlertPolicyOutput() ServerSecurityAlertPolicyOutput
	ToServerSecurityAlertPolicyOutputWithContext(ctx context.Context) ServerSecurityAlertPolicyOutput
}

func (*ServerSecurityAlertPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**ServerSecurityAlertPolicy)(nil)).Elem()
}

func (i *ServerSecurityAlertPolicy) ToServerSecurityAlertPolicyOutput() ServerSecurityAlertPolicyOutput {
	return i.ToServerSecurityAlertPolicyOutputWithContext(context.Background())
}

func (i *ServerSecurityAlertPolicy) ToServerSecurityAlertPolicyOutputWithContext(ctx context.Context) ServerSecurityAlertPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServerSecurityAlertPolicyOutput)
}

type ServerSecurityAlertPolicyOutput struct{ *pulumi.OutputState }

func (ServerSecurityAlertPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServerSecurityAlertPolicy)(nil)).Elem()
}

func (o ServerSecurityAlertPolicyOutput) ToServerSecurityAlertPolicyOutput() ServerSecurityAlertPolicyOutput {
	return o
}

func (o ServerSecurityAlertPolicyOutput) ToServerSecurityAlertPolicyOutputWithContext(ctx context.Context) ServerSecurityAlertPolicyOutput {
	return o
}

func (o ServerSecurityAlertPolicyOutput) DisabledAlerts() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ServerSecurityAlertPolicy) pulumi.StringArrayOutput { return v.DisabledAlerts }).(pulumi.StringArrayOutput)
}

func (o ServerSecurityAlertPolicyOutput) EmailAccountAdmins() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ServerSecurityAlertPolicy) pulumi.BoolPtrOutput { return v.EmailAccountAdmins }).(pulumi.BoolPtrOutput)
}

func (o ServerSecurityAlertPolicyOutput) EmailAddresses() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ServerSecurityAlertPolicy) pulumi.StringArrayOutput { return v.EmailAddresses }).(pulumi.StringArrayOutput)
}

func (o ServerSecurityAlertPolicyOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ServerSecurityAlertPolicy) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ServerSecurityAlertPolicyOutput) RetentionDays() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ServerSecurityAlertPolicy) pulumi.IntPtrOutput { return v.RetentionDays }).(pulumi.IntPtrOutput)
}

func (o ServerSecurityAlertPolicyOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v *ServerSecurityAlertPolicy) pulumi.StringOutput { return v.State }).(pulumi.StringOutput)
}

func (o ServerSecurityAlertPolicyOutput) StorageAccountAccessKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServerSecurityAlertPolicy) pulumi.StringPtrOutput { return v.StorageAccountAccessKey }).(pulumi.StringPtrOutput)
}

func (o ServerSecurityAlertPolicyOutput) StorageEndpoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServerSecurityAlertPolicy) pulumi.StringPtrOutput { return v.StorageEndpoint }).(pulumi.StringPtrOutput)
}

func (o ServerSecurityAlertPolicyOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ServerSecurityAlertPolicy) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ServerSecurityAlertPolicyOutput{})
}
