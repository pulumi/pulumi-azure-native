


package v20170801preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type SecurityContact struct {
	pulumi.CustomResourceState

	AlertNotifications pulumi.StringOutput    `pulumi:"alertNotifications"`
	AlertsToAdmins     pulumi.StringOutput    `pulumi:"alertsToAdmins"`
	Email              pulumi.StringOutput    `pulumi:"email"`
	Name               pulumi.StringOutput    `pulumi:"name"`
	Phone              pulumi.StringPtrOutput `pulumi:"phone"`
	Type               pulumi.StringOutput    `pulumi:"type"`
}


func NewSecurityContact(ctx *pulumi.Context,
	name string, args *SecurityContactArgs, opts ...pulumi.ResourceOption) (*SecurityContact, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AlertNotifications == nil {
		return nil, errors.New("invalid value for required argument 'AlertNotifications'")
	}
	if args.AlertsToAdmins == nil {
		return nil, errors.New("invalid value for required argument 'AlertsToAdmins'")
	}
	if args.Email == nil {
		return nil, errors.New("invalid value for required argument 'Email'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:security/v20170801preview:SecurityContact"),
		},
		{
			Type: pulumi.String("azure-native:security:SecurityContact"),
		},
		{
			Type: pulumi.String("azure-nextgen:security:SecurityContact"),
		},
		{
			Type: pulumi.String("azure-native:security/v20200101preview:SecurityContact"),
		},
		{
			Type: pulumi.String("azure-nextgen:security/v20200101preview:SecurityContact"),
		},
	})
	opts = append(opts, aliases)
	var resource SecurityContact
	err := ctx.RegisterResource("azure-native:security/v20170801preview:SecurityContact", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetSecurityContact(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SecurityContactState, opts ...pulumi.ResourceOption) (*SecurityContact, error) {
	var resource SecurityContact
	err := ctx.ReadResource("azure-native:security/v20170801preview:SecurityContact", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type securityContactState struct {
}

type SecurityContactState struct {
}

func (SecurityContactState) ElementType() reflect.Type {
	return reflect.TypeOf((*securityContactState)(nil)).Elem()
}

type securityContactArgs struct {
	AlertNotifications  string  `pulumi:"alertNotifications"`
	AlertsToAdmins      string  `pulumi:"alertsToAdmins"`
	Email               string  `pulumi:"email"`
	Phone               *string `pulumi:"phone"`
	SecurityContactName *string `pulumi:"securityContactName"`
}


type SecurityContactArgs struct {
	AlertNotifications  pulumi.StringInput
	AlertsToAdmins      pulumi.StringInput
	Email               pulumi.StringInput
	Phone               pulumi.StringPtrInput
	SecurityContactName pulumi.StringPtrInput
}

func (SecurityContactArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*securityContactArgs)(nil)).Elem()
}

type SecurityContactInput interface {
	pulumi.Input

	ToSecurityContactOutput() SecurityContactOutput
	ToSecurityContactOutputWithContext(ctx context.Context) SecurityContactOutput
}

func (*SecurityContact) ElementType() reflect.Type {
	return reflect.TypeOf((*SecurityContact)(nil))
}

func (i *SecurityContact) ToSecurityContactOutput() SecurityContactOutput {
	return i.ToSecurityContactOutputWithContext(context.Background())
}

func (i *SecurityContact) ToSecurityContactOutputWithContext(ctx context.Context) SecurityContactOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecurityContactOutput)
}

type SecurityContactOutput struct{ *pulumi.OutputState }

func (SecurityContactOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SecurityContact)(nil))
}

func (o SecurityContactOutput) ToSecurityContactOutput() SecurityContactOutput {
	return o
}

func (o SecurityContactOutput) ToSecurityContactOutputWithContext(ctx context.Context) SecurityContactOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(SecurityContactOutput{})
}
