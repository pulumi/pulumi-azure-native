


package v20210201

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type CloudEdgeManagementRole struct {
	pulumi.CustomResourceState

	EdgeProfile           EdgeProfileResponseOutput `pulumi:"edgeProfile"`
	Kind                  pulumi.StringOutput       `pulumi:"kind"`
	LocalManagementStatus pulumi.StringOutput       `pulumi:"localManagementStatus"`
	Name                  pulumi.StringOutput       `pulumi:"name"`
	RoleStatus            pulumi.StringOutput       `pulumi:"roleStatus"`
	SystemData            SystemDataResponseOutput  `pulumi:"systemData"`
	Type                  pulumi.StringOutput       `pulumi:"type"`
}


func NewCloudEdgeManagementRole(ctx *pulumi.Context,
	name string, args *CloudEdgeManagementRoleArgs, opts ...pulumi.ResourceOption) (*CloudEdgeManagementRole, error) {
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
	args.Kind = pulumi.String("CloudEdgeManagement")
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:databoxedge:CloudEdgeManagementRole"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20190301:CloudEdgeManagementRole"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20190701:CloudEdgeManagementRole"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20190801:CloudEdgeManagementRole"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20200501preview:CloudEdgeManagementRole"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20200901:CloudEdgeManagementRole"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20200901preview:CloudEdgeManagementRole"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20201201:CloudEdgeManagementRole"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20210201preview:CloudEdgeManagementRole"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20210601:CloudEdgeManagementRole"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20210601preview:CloudEdgeManagementRole"),
		},
	})
	opts = append(opts, aliases)
	var resource CloudEdgeManagementRole
	err := ctx.RegisterResource("azure-native:databoxedge/v20210201:CloudEdgeManagementRole", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetCloudEdgeManagementRole(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CloudEdgeManagementRoleState, opts ...pulumi.ResourceOption) (*CloudEdgeManagementRole, error) {
	var resource CloudEdgeManagementRole
	err := ctx.ReadResource("azure-native:databoxedge/v20210201:CloudEdgeManagementRole", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type cloudEdgeManagementRoleState struct {
}

type CloudEdgeManagementRoleState struct {
}

func (CloudEdgeManagementRoleState) ElementType() reflect.Type {
	return reflect.TypeOf((*cloudEdgeManagementRoleState)(nil)).Elem()
}

type cloudEdgeManagementRoleArgs struct {
	DeviceName        string  `pulumi:"deviceName"`
	Kind              string  `pulumi:"kind"`
	Name              *string `pulumi:"name"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
	RoleStatus        string  `pulumi:"roleStatus"`
}


type CloudEdgeManagementRoleArgs struct {
	DeviceName        pulumi.StringInput
	Kind              pulumi.StringInput
	Name              pulumi.StringPtrInput
	ResourceGroupName pulumi.StringInput
	RoleStatus        pulumi.StringInput
}

func (CloudEdgeManagementRoleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*cloudEdgeManagementRoleArgs)(nil)).Elem()
}

type CloudEdgeManagementRoleInput interface {
	pulumi.Input

	ToCloudEdgeManagementRoleOutput() CloudEdgeManagementRoleOutput
	ToCloudEdgeManagementRoleOutputWithContext(ctx context.Context) CloudEdgeManagementRoleOutput
}

func (*CloudEdgeManagementRole) ElementType() reflect.Type {
	return reflect.TypeOf((*CloudEdgeManagementRole)(nil))
}

func (i *CloudEdgeManagementRole) ToCloudEdgeManagementRoleOutput() CloudEdgeManagementRoleOutput {
	return i.ToCloudEdgeManagementRoleOutputWithContext(context.Background())
}

func (i *CloudEdgeManagementRole) ToCloudEdgeManagementRoleOutputWithContext(ctx context.Context) CloudEdgeManagementRoleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CloudEdgeManagementRoleOutput)
}

type CloudEdgeManagementRoleOutput struct{ *pulumi.OutputState }

func (CloudEdgeManagementRoleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CloudEdgeManagementRole)(nil))
}

func (o CloudEdgeManagementRoleOutput) ToCloudEdgeManagementRoleOutput() CloudEdgeManagementRoleOutput {
	return o
}

func (o CloudEdgeManagementRoleOutput) ToCloudEdgeManagementRoleOutputWithContext(ctx context.Context) CloudEdgeManagementRoleOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(CloudEdgeManagementRoleOutput{})
}
