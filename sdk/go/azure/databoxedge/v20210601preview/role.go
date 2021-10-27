


package v20210601preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


type Role struct {
	pulumi.CustomResourceState

	Kind       pulumi.StringOutput      `pulumi:"kind"`
	Name       pulumi.StringOutput      `pulumi:"name"`
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	Type       pulumi.StringOutput      `pulumi:"type"`
}


func NewRole(ctx *pulumi.Context,
	name string, args *RoleArgs, opts ...pulumi.ResourceOption) (*Role, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DeviceName == nil {
		return nil, errors.New("invalid value for required argument 'DeviceName'")
	}
	if args.Kind == nil {
		return nil, errors.New("invalid value for required argument 'Kind'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:databoxedge/v20210601preview:Role"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge:Role"),
		},
		{
			Type: pulumi.String("azure-nextgen:databoxedge:Role"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20190301:Role"),
		},
		{
			Type: pulumi.String("azure-nextgen:databoxedge/v20190301:Role"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20190701:Role"),
		},
		{
			Type: pulumi.String("azure-nextgen:databoxedge/v20190701:Role"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20190801:Role"),
		},
		{
			Type: pulumi.String("azure-nextgen:databoxedge/v20190801:Role"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20200501preview:Role"),
		},
		{
			Type: pulumi.String("azure-nextgen:databoxedge/v20200501preview:Role"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20200901:Role"),
		},
		{
			Type: pulumi.String("azure-nextgen:databoxedge/v20200901:Role"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20200901preview:Role"),
		},
		{
			Type: pulumi.String("azure-nextgen:databoxedge/v20200901preview:Role"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20201201:Role"),
		},
		{
			Type: pulumi.String("azure-nextgen:databoxedge/v20201201:Role"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20210201:Role"),
		},
		{
			Type: pulumi.String("azure-nextgen:databoxedge/v20210201:Role"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20210201preview:Role"),
		},
		{
			Type: pulumi.String("azure-nextgen:databoxedge/v20210201preview:Role"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20210601:Role"),
		},
		{
			Type: pulumi.String("azure-nextgen:databoxedge/v20210601:Role"),
		},
	})
	opts = append(opts, aliases)
	var resource Role
	err := ctx.RegisterResource("azure-native:databoxedge/v20210601preview:Role", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetRole(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RoleState, opts ...pulumi.ResourceOption) (*Role, error) {
	var resource Role
	err := ctx.ReadResource("azure-native:databoxedge/v20210601preview:Role", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type roleState struct {
}

type RoleState struct {
}

func (RoleState) ElementType() reflect.Type {
	return reflect.TypeOf((*roleState)(nil)).Elem()
}

type roleArgs struct {
	DeviceName        string  `pulumi:"deviceName"`
	Kind              string  `pulumi:"kind"`
	Name              *string `pulumi:"name"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
}


type RoleArgs struct {
	DeviceName        pulumi.StringInput
	Kind              pulumi.StringInput
	Name              pulumi.StringPtrInput
	ResourceGroupName pulumi.StringInput
}

func (RoleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*roleArgs)(nil)).Elem()
}

type RoleInput interface {
	pulumi.Input

	ToRoleOutput() RoleOutput
	ToRoleOutputWithContext(ctx context.Context) RoleOutput
}

func (*Role) ElementType() reflect.Type {
	return reflect.TypeOf((*Role)(nil))
}

func (i *Role) ToRoleOutput() RoleOutput {
	return i.ToRoleOutputWithContext(context.Background())
}

func (i *Role) ToRoleOutputWithContext(ctx context.Context) RoleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RoleOutput)
}

type RoleOutput struct{ *pulumi.OutputState }

func (RoleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Role)(nil))
}

func (o RoleOutput) ToRoleOutput() RoleOutput {
	return o
}

func (o RoleOutput) ToRoleOutputWithContext(ctx context.Context) RoleOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RoleInput)(nil)).Elem(), &Role{})
	pulumi.RegisterOutputType(RoleOutput{})
}
