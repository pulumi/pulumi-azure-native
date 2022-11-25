


package v20220308preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Administrator struct {
	pulumi.CustomResourceState

	Name          pulumi.StringOutput      `pulumi:"name"`
	ObjectId      pulumi.StringPtrOutput   `pulumi:"objectId"`
	PrincipalName pulumi.StringPtrOutput   `pulumi:"principalName"`
	PrincipalType pulumi.StringPtrOutput   `pulumi:"principalType"`
	SystemData    SystemDataResponseOutput `pulumi:"systemData"`
	TenantId      pulumi.StringPtrOutput   `pulumi:"tenantId"`
	Type          pulumi.StringOutput      `pulumi:"type"`
}


func NewAdministrator(ctx *pulumi.Context,
	name string, args *AdministratorArgs, opts ...pulumi.ResourceOption) (*Administrator, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ServerName == nil {
		return nil, errors.New("invalid value for required argument 'ServerName'")
	}
	var resource Administrator
	err := ctx.RegisterResource("azure-native:dbforpostgresql/v20220308preview:Administrator", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetAdministrator(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AdministratorState, opts ...pulumi.ResourceOption) (*Administrator, error) {
	var resource Administrator
	err := ctx.ReadResource("azure-native:dbforpostgresql/v20220308preview:Administrator", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type administratorState struct {
}

type AdministratorState struct {
}

func (AdministratorState) ElementType() reflect.Type {
	return reflect.TypeOf((*administratorState)(nil)).Elem()
}

type administratorArgs struct {
	ObjectId          *string `pulumi:"objectId"`
	PrincipalName     *string `pulumi:"principalName"`
	PrincipalType     *string `pulumi:"principalType"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
	ServerName        string  `pulumi:"serverName"`
	TenantId          *string `pulumi:"tenantId"`
}


type AdministratorArgs struct {
	ObjectId          pulumi.StringPtrInput
	PrincipalName     pulumi.StringPtrInput
	PrincipalType     pulumi.StringPtrInput
	ResourceGroupName pulumi.StringInput
	ServerName        pulumi.StringInput
	TenantId          pulumi.StringPtrInput
}

func (AdministratorArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*administratorArgs)(nil)).Elem()
}

type AdministratorInput interface {
	pulumi.Input

	ToAdministratorOutput() AdministratorOutput
	ToAdministratorOutputWithContext(ctx context.Context) AdministratorOutput
}

func (*Administrator) ElementType() reflect.Type {
	return reflect.TypeOf((**Administrator)(nil)).Elem()
}

func (i *Administrator) ToAdministratorOutput() AdministratorOutput {
	return i.ToAdministratorOutputWithContext(context.Background())
}

func (i *Administrator) ToAdministratorOutputWithContext(ctx context.Context) AdministratorOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AdministratorOutput)
}

type AdministratorOutput struct{ *pulumi.OutputState }

func (AdministratorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Administrator)(nil)).Elem()
}

func (o AdministratorOutput) ToAdministratorOutput() AdministratorOutput {
	return o
}

func (o AdministratorOutput) ToAdministratorOutputWithContext(ctx context.Context) AdministratorOutput {
	return o
}

func (o AdministratorOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Administrator) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o AdministratorOutput) ObjectId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Administrator) pulumi.StringPtrOutput { return v.ObjectId }).(pulumi.StringPtrOutput)
}

func (o AdministratorOutput) PrincipalName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Administrator) pulumi.StringPtrOutput { return v.PrincipalName }).(pulumi.StringPtrOutput)
}

func (o AdministratorOutput) PrincipalType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Administrator) pulumi.StringPtrOutput { return v.PrincipalType }).(pulumi.StringPtrOutput)
}

func (o AdministratorOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *Administrator) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o AdministratorOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Administrator) pulumi.StringPtrOutput { return v.TenantId }).(pulumi.StringPtrOutput)
}

func (o AdministratorOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Administrator) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(AdministratorOutput{})
}
