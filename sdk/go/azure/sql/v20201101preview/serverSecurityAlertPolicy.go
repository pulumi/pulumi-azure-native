


package v20201101preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ServerSecurityAlertPolicy struct {
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
			Type: pulumi.String("azure-nextgen:sql/v20201101preview:ServerSecurityAlertPolicy"),
		},
		{
			Type: pulumi.String("azure-native:sql:ServerSecurityAlertPolicy"),
		},
		{
			Type: pulumi.String("azure-nextgen:sql:ServerSecurityAlertPolicy"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20170301preview:ServerSecurityAlertPolicy"),
		},
		{
			Type: pulumi.String("azure-nextgen:sql/v20170301preview:ServerSecurityAlertPolicy"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20200202preview:ServerSecurityAlertPolicy"),
		},
		{
			Type: pulumi.String("azure-nextgen:sql/v20200202preview:ServerSecurityAlertPolicy"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20200801preview:ServerSecurityAlertPolicy"),
		},
		{
			Type: pulumi.String("azure-nextgen:sql/v20200801preview:ServerSecurityAlertPolicy"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20210201preview:ServerSecurityAlertPolicy"),
		},
		{
			Type: pulumi.String("azure-nextgen:sql/v20210201preview:ServerSecurityAlertPolicy"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20210501preview:ServerSecurityAlertPolicy"),
		},
		{
			Type: pulumi.String("azure-nextgen:sql/v20210501preview:ServerSecurityAlertPolicy"),
		},
	})
	opts = append(opts, aliases)
	var resource ServerSecurityAlertPolicy
	err := ctx.RegisterResource("azure-native:sql/v20201101preview:ServerSecurityAlertPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetServerSecurityAlertPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServerSecurityAlertPolicyState, opts ...pulumi.ResourceOption) (*ServerSecurityAlertPolicy, error) {
	var resource ServerSecurityAlertPolicy
	err := ctx.ReadResource("azure-native:sql/v20201101preview:ServerSecurityAlertPolicy", name, id, state, &resource, opts...)
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


type ServerSecurityAlertPolicyArgs struct {
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

func (ServerSecurityAlertPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*serverSecurityAlertPolicyArgs)(nil)).Elem()
}

type ServerSecurityAlertPolicyInput interface {
	pulumi.Input

	ToServerSecurityAlertPolicyOutput() ServerSecurityAlertPolicyOutput
	ToServerSecurityAlertPolicyOutputWithContext(ctx context.Context) ServerSecurityAlertPolicyOutput
}

func (*ServerSecurityAlertPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((*ServerSecurityAlertPolicy)(nil))
}

func (i *ServerSecurityAlertPolicy) ToServerSecurityAlertPolicyOutput() ServerSecurityAlertPolicyOutput {
	return i.ToServerSecurityAlertPolicyOutputWithContext(context.Background())
}

func (i *ServerSecurityAlertPolicy) ToServerSecurityAlertPolicyOutputWithContext(ctx context.Context) ServerSecurityAlertPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServerSecurityAlertPolicyOutput)
}

type ServerSecurityAlertPolicyOutput struct{ *pulumi.OutputState }

func (ServerSecurityAlertPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServerSecurityAlertPolicy)(nil))
}

func (o ServerSecurityAlertPolicyOutput) ToServerSecurityAlertPolicyOutput() ServerSecurityAlertPolicyOutput {
	return o
}

func (o ServerSecurityAlertPolicyOutput) ToServerSecurityAlertPolicyOutputWithContext(ctx context.Context) ServerSecurityAlertPolicyOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ServerSecurityAlertPolicyInput)(nil)).Elem(), &ServerSecurityAlertPolicy{})
	pulumi.RegisterOutputType(ServerSecurityAlertPolicyOutput{})
}
