


package v20170301preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type SqlServerRegistration struct {
	pulumi.CustomResourceState

	Location       pulumi.StringOutput    `pulumi:"location"`
	Name           pulumi.StringOutput    `pulumi:"name"`
	PropertyBag    pulumi.StringPtrOutput `pulumi:"propertyBag"`
	ResourceGroup  pulumi.StringPtrOutput `pulumi:"resourceGroup"`
	SubscriptionId pulumi.StringPtrOutput `pulumi:"subscriptionId"`
	Tags           pulumi.StringMapOutput `pulumi:"tags"`
	Type           pulumi.StringOutput    `pulumi:"type"`
}


func NewSqlServerRegistration(ctx *pulumi.Context,
	name string, args *SqlServerRegistrationArgs, opts ...pulumi.ResourceOption) (*SqlServerRegistration, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:azuredata:SqlServerRegistration"),
		},
		{
			Type: pulumi.String("azure-native:azuredata/v20190724preview:SqlServerRegistration"),
		},
	})
	opts = append(opts, aliases)
	var resource SqlServerRegistration
	err := ctx.RegisterResource("azure-native:azuredata/v20170301preview:SqlServerRegistration", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetSqlServerRegistration(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SqlServerRegistrationState, opts ...pulumi.ResourceOption) (*SqlServerRegistration, error) {
	var resource SqlServerRegistration
	err := ctx.ReadResource("azure-native:azuredata/v20170301preview:SqlServerRegistration", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type sqlServerRegistrationState struct {
}

type SqlServerRegistrationState struct {
}

func (SqlServerRegistrationState) ElementType() reflect.Type {
	return reflect.TypeOf((*sqlServerRegistrationState)(nil)).Elem()
}

type sqlServerRegistrationArgs struct {
	Location                  *string           `pulumi:"location"`
	PropertyBag               *string           `pulumi:"propertyBag"`
	ResourceGroup             *string           `pulumi:"resourceGroup"`
	ResourceGroupName         string            `pulumi:"resourceGroupName"`
	SqlServerRegistrationName *string           `pulumi:"sqlServerRegistrationName"`
	SubscriptionId            *string           `pulumi:"subscriptionId"`
	Tags                      map[string]string `pulumi:"tags"`
}


type SqlServerRegistrationArgs struct {
	Location                  pulumi.StringPtrInput
	PropertyBag               pulumi.StringPtrInput
	ResourceGroup             pulumi.StringPtrInput
	ResourceGroupName         pulumi.StringInput
	SqlServerRegistrationName pulumi.StringPtrInput
	SubscriptionId            pulumi.StringPtrInput
	Tags                      pulumi.StringMapInput
}

func (SqlServerRegistrationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sqlServerRegistrationArgs)(nil)).Elem()
}

type SqlServerRegistrationInput interface {
	pulumi.Input

	ToSqlServerRegistrationOutput() SqlServerRegistrationOutput
	ToSqlServerRegistrationOutputWithContext(ctx context.Context) SqlServerRegistrationOutput
}

func (*SqlServerRegistration) ElementType() reflect.Type {
	return reflect.TypeOf((**SqlServerRegistration)(nil)).Elem()
}

func (i *SqlServerRegistration) ToSqlServerRegistrationOutput() SqlServerRegistrationOutput {
	return i.ToSqlServerRegistrationOutputWithContext(context.Background())
}

func (i *SqlServerRegistration) ToSqlServerRegistrationOutputWithContext(ctx context.Context) SqlServerRegistrationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SqlServerRegistrationOutput)
}

type SqlServerRegistrationOutput struct{ *pulumi.OutputState }

func (SqlServerRegistrationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SqlServerRegistration)(nil)).Elem()
}

func (o SqlServerRegistrationOutput) ToSqlServerRegistrationOutput() SqlServerRegistrationOutput {
	return o
}

func (o SqlServerRegistrationOutput) ToSqlServerRegistrationOutputWithContext(ctx context.Context) SqlServerRegistrationOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(SqlServerRegistrationOutput{})
}
