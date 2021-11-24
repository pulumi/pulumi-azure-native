


package v20210501

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type WorkspaceSqlAadAdmin struct {
	pulumi.CustomResourceState

	AdministratorType pulumi.StringPtrOutput `pulumi:"administratorType"`
	Login             pulumi.StringPtrOutput `pulumi:"login"`
	Name              pulumi.StringOutput    `pulumi:"name"`
	Sid               pulumi.StringPtrOutput `pulumi:"sid"`
	TenantId          pulumi.StringPtrOutput `pulumi:"tenantId"`
	Type              pulumi.StringOutput    `pulumi:"type"`
}


func NewWorkspaceSqlAadAdmin(ctx *pulumi.Context,
	name string, args *WorkspaceSqlAadAdminArgs, opts ...pulumi.ResourceOption) (*WorkspaceSqlAadAdmin, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.WorkspaceName == nil {
		return nil, errors.New("invalid value for required argument 'WorkspaceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:synapse:WorkspaceSqlAadAdmin"),
		},
		{
			Type: pulumi.String("azure-native:synapse/v20190601preview:WorkspaceSqlAadAdmin"),
		},
		{
			Type: pulumi.String("azure-native:synapse/v20201201:WorkspaceSqlAadAdmin"),
		},
		{
			Type: pulumi.String("azure-native:synapse/v20210301:WorkspaceSqlAadAdmin"),
		},
		{
			Type: pulumi.String("azure-native:synapse/v20210401preview:WorkspaceSqlAadAdmin"),
		},
		{
			Type: pulumi.String("azure-native:synapse/v20210601:WorkspaceSqlAadAdmin"),
		},
		{
			Type: pulumi.String("azure-native:synapse/v20210601preview:WorkspaceSqlAadAdmin"),
		},
	})
	opts = append(opts, aliases)
	var resource WorkspaceSqlAadAdmin
	err := ctx.RegisterResource("azure-native:synapse/v20210501:WorkspaceSqlAadAdmin", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetWorkspaceSqlAadAdmin(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WorkspaceSqlAadAdminState, opts ...pulumi.ResourceOption) (*WorkspaceSqlAadAdmin, error) {
	var resource WorkspaceSqlAadAdmin
	err := ctx.ReadResource("azure-native:synapse/v20210501:WorkspaceSqlAadAdmin", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type workspaceSqlAadAdminState struct {
}

type WorkspaceSqlAadAdminState struct {
}

func (WorkspaceSqlAadAdminState) ElementType() reflect.Type {
	return reflect.TypeOf((*workspaceSqlAadAdminState)(nil)).Elem()
}

type workspaceSqlAadAdminArgs struct {
	AdministratorType *string `pulumi:"administratorType"`
	Login             *string `pulumi:"login"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
	Sid               *string `pulumi:"sid"`
	TenantId          *string `pulumi:"tenantId"`
	WorkspaceName     string  `pulumi:"workspaceName"`
}


type WorkspaceSqlAadAdminArgs struct {
	AdministratorType pulumi.StringPtrInput
	Login             pulumi.StringPtrInput
	ResourceGroupName pulumi.StringInput
	Sid               pulumi.StringPtrInput
	TenantId          pulumi.StringPtrInput
	WorkspaceName     pulumi.StringInput
}

func (WorkspaceSqlAadAdminArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*workspaceSqlAadAdminArgs)(nil)).Elem()
}

type WorkspaceSqlAadAdminInput interface {
	pulumi.Input

	ToWorkspaceSqlAadAdminOutput() WorkspaceSqlAadAdminOutput
	ToWorkspaceSqlAadAdminOutputWithContext(ctx context.Context) WorkspaceSqlAadAdminOutput
}

func (*WorkspaceSqlAadAdmin) ElementType() reflect.Type {
	return reflect.TypeOf((*WorkspaceSqlAadAdmin)(nil))
}

func (i *WorkspaceSqlAadAdmin) ToWorkspaceSqlAadAdminOutput() WorkspaceSqlAadAdminOutput {
	return i.ToWorkspaceSqlAadAdminOutputWithContext(context.Background())
}

func (i *WorkspaceSqlAadAdmin) ToWorkspaceSqlAadAdminOutputWithContext(ctx context.Context) WorkspaceSqlAadAdminOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkspaceSqlAadAdminOutput)
}

type WorkspaceSqlAadAdminOutput struct{ *pulumi.OutputState }

func (WorkspaceSqlAadAdminOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WorkspaceSqlAadAdmin)(nil))
}

func (o WorkspaceSqlAadAdminOutput) ToWorkspaceSqlAadAdminOutput() WorkspaceSqlAadAdminOutput {
	return o
}

func (o WorkspaceSqlAadAdminOutput) ToWorkspaceSqlAadAdminOutputWithContext(ctx context.Context) WorkspaceSqlAadAdminOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(WorkspaceSqlAadAdminOutput{})
}
