


package v20201201

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type MECRole struct {
	pulumi.CustomResourceState

	ConnectionString AsymmetricEncryptedSecretResponsePtrOutput `pulumi:"connectionString"`
	Kind             pulumi.StringOutput                        `pulumi:"kind"`
	Name             pulumi.StringOutput                        `pulumi:"name"`
	RoleStatus       pulumi.StringOutput                        `pulumi:"roleStatus"`
	SystemData       SystemDataResponseOutput                   `pulumi:"systemData"`
	Type             pulumi.StringOutput                        `pulumi:"type"`
}


func NewMECRole(ctx *pulumi.Context,
	name string, args *MECRoleArgs, opts ...pulumi.ResourceOption) (*MECRole, error) {
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
	if args.RoleStatus == nil {
		return nil, errors.New("invalid value for required argument 'RoleStatus'")
	}
	args.Kind = pulumi.String("MEC")
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:databoxedge:MECRole"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20190301:MECRole"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20190701:MECRole"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20190801:MECRole"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20200501preview:MECRole"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20200901:MECRole"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20200901preview:MECRole"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20210201:MECRole"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20210201preview:MECRole"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20210601:MECRole"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20210601preview:MECRole"),
		},
	})
	opts = append(opts, aliases)
	var resource MECRole
	err := ctx.RegisterResource("azure-native:databoxedge/v20201201:MECRole", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetMECRole(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MECRoleState, opts ...pulumi.ResourceOption) (*MECRole, error) {
	var resource MECRole
	err := ctx.ReadResource("azure-native:databoxedge/v20201201:MECRole", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type mecroleState struct {
}

type MECRoleState struct {
}

func (MECRoleState) ElementType() reflect.Type {
	return reflect.TypeOf((*mecroleState)(nil)).Elem()
}

type mecroleArgs struct {
	ConnectionString  *AsymmetricEncryptedSecret `pulumi:"connectionString"`
	DeviceName        string                     `pulumi:"deviceName"`
	Kind              string                     `pulumi:"kind"`
	Name              *string                    `pulumi:"name"`
	ResourceGroupName string                     `pulumi:"resourceGroupName"`
	RoleStatus        string                     `pulumi:"roleStatus"`
}


type MECRoleArgs struct {
	ConnectionString  AsymmetricEncryptedSecretPtrInput
	DeviceName        pulumi.StringInput
	Kind              pulumi.StringInput
	Name              pulumi.StringPtrInput
	ResourceGroupName pulumi.StringInput
	RoleStatus        pulumi.StringInput
}

func (MECRoleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*mecroleArgs)(nil)).Elem()
}

type MECRoleInput interface {
	pulumi.Input

	ToMECRoleOutput() MECRoleOutput
	ToMECRoleOutputWithContext(ctx context.Context) MECRoleOutput
}

func (*MECRole) ElementType() reflect.Type {
	return reflect.TypeOf((*MECRole)(nil))
}

func (i *MECRole) ToMECRoleOutput() MECRoleOutput {
	return i.ToMECRoleOutputWithContext(context.Background())
}

func (i *MECRole) ToMECRoleOutputWithContext(ctx context.Context) MECRoleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MECRoleOutput)
}

type MECRoleOutput struct{ *pulumi.OutputState }

func (MECRoleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MECRole)(nil))
}

func (o MECRoleOutput) ToMECRoleOutput() MECRoleOutput {
	return o
}

func (o MECRoleOutput) ToMECRoleOutputWithContext(ctx context.Context) MECRoleOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(MECRoleOutput{})
}
