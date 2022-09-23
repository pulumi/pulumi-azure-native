


package v20200202preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ManagedInstanceAdministrator struct {
	pulumi.CustomResourceState

	AdministratorType pulumi.StringOutput    `pulumi:"administratorType"`
	Login             pulumi.StringOutput    `pulumi:"login"`
	Name              pulumi.StringOutput    `pulumi:"name"`
	Sid               pulumi.StringOutput    `pulumi:"sid"`
	TenantId          pulumi.StringPtrOutput `pulumi:"tenantId"`
	Type              pulumi.StringOutput    `pulumi:"type"`
}


func NewManagedInstanceAdministrator(ctx *pulumi.Context,
	name string, args *ManagedInstanceAdministratorArgs, opts ...pulumi.ResourceOption) (*ManagedInstanceAdministrator, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AdministratorType == nil {
		return nil, errors.New("invalid value for required argument 'AdministratorType'")
	}
	if args.Login == nil {
		return nil, errors.New("invalid value for required argument 'Login'")
	}
	if args.ManagedInstanceName == nil {
		return nil, errors.New("invalid value for required argument 'ManagedInstanceName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.Sid == nil {
		return nil, errors.New("invalid value for required argument 'Sid'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:sql:ManagedInstanceAdministrator"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20170301preview:ManagedInstanceAdministrator"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20200801preview:ManagedInstanceAdministrator"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20201101preview:ManagedInstanceAdministrator"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20210201preview:ManagedInstanceAdministrator"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20210501preview:ManagedInstanceAdministrator"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20210801preview:ManagedInstanceAdministrator"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20211101:ManagedInstanceAdministrator"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20211101preview:ManagedInstanceAdministrator"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20220201preview:ManagedInstanceAdministrator"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20220501preview:ManagedInstanceAdministrator"),
		},
	})
	opts = append(opts, aliases)
	var resource ManagedInstanceAdministrator
	err := ctx.RegisterResource("azure-native:sql/v20200202preview:ManagedInstanceAdministrator", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetManagedInstanceAdministrator(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ManagedInstanceAdministratorState, opts ...pulumi.ResourceOption) (*ManagedInstanceAdministrator, error) {
	var resource ManagedInstanceAdministrator
	err := ctx.ReadResource("azure-native:sql/v20200202preview:ManagedInstanceAdministrator", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type managedInstanceAdministratorState struct {
}

type ManagedInstanceAdministratorState struct {
}

func (ManagedInstanceAdministratorState) ElementType() reflect.Type {
	return reflect.TypeOf((*managedInstanceAdministratorState)(nil)).Elem()
}

type managedInstanceAdministratorArgs struct {
	AdministratorName   *string `pulumi:"administratorName"`
	AdministratorType   string  `pulumi:"administratorType"`
	Login               string  `pulumi:"login"`
	ManagedInstanceName string  `pulumi:"managedInstanceName"`
	ResourceGroupName   string  `pulumi:"resourceGroupName"`
	Sid                 string  `pulumi:"sid"`
	TenantId            *string `pulumi:"tenantId"`
}


type ManagedInstanceAdministratorArgs struct {
	AdministratorName   pulumi.StringPtrInput
	AdministratorType   pulumi.StringInput
	Login               pulumi.StringInput
	ManagedInstanceName pulumi.StringInput
	ResourceGroupName   pulumi.StringInput
	Sid                 pulumi.StringInput
	TenantId            pulumi.StringPtrInput
}

func (ManagedInstanceAdministratorArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*managedInstanceAdministratorArgs)(nil)).Elem()
}

type ManagedInstanceAdministratorInput interface {
	pulumi.Input

	ToManagedInstanceAdministratorOutput() ManagedInstanceAdministratorOutput
	ToManagedInstanceAdministratorOutputWithContext(ctx context.Context) ManagedInstanceAdministratorOutput
}

func (*ManagedInstanceAdministrator) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedInstanceAdministrator)(nil)).Elem()
}

func (i *ManagedInstanceAdministrator) ToManagedInstanceAdministratorOutput() ManagedInstanceAdministratorOutput {
	return i.ToManagedInstanceAdministratorOutputWithContext(context.Background())
}

func (i *ManagedInstanceAdministrator) ToManagedInstanceAdministratorOutputWithContext(ctx context.Context) ManagedInstanceAdministratorOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedInstanceAdministratorOutput)
}

type ManagedInstanceAdministratorOutput struct{ *pulumi.OutputState }

func (ManagedInstanceAdministratorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedInstanceAdministrator)(nil)).Elem()
}

func (o ManagedInstanceAdministratorOutput) ToManagedInstanceAdministratorOutput() ManagedInstanceAdministratorOutput {
	return o
}

func (o ManagedInstanceAdministratorOutput) ToManagedInstanceAdministratorOutputWithContext(ctx context.Context) ManagedInstanceAdministratorOutput {
	return o
}

func (o ManagedInstanceAdministratorOutput) AdministratorType() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagedInstanceAdministrator) pulumi.StringOutput { return v.AdministratorType }).(pulumi.StringOutput)
}

func (o ManagedInstanceAdministratorOutput) Login() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagedInstanceAdministrator) pulumi.StringOutput { return v.Login }).(pulumi.StringOutput)
}

func (o ManagedInstanceAdministratorOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagedInstanceAdministrator) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ManagedInstanceAdministratorOutput) Sid() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagedInstanceAdministrator) pulumi.StringOutput { return v.Sid }).(pulumi.StringOutput)
}

func (o ManagedInstanceAdministratorOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedInstanceAdministrator) pulumi.StringPtrOutput { return v.TenantId }).(pulumi.StringPtrOutput)
}

func (o ManagedInstanceAdministratorOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagedInstanceAdministrator) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ManagedInstanceAdministratorOutput{})
}
