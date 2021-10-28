


package v20210701

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type WorkspaceConnection struct {
	pulumi.CustomResourceState

	AuthType    pulumi.StringPtrOutput `pulumi:"authType"`
	Category    pulumi.StringPtrOutput `pulumi:"category"`
	Name        pulumi.StringOutput    `pulumi:"name"`
	Target      pulumi.StringPtrOutput `pulumi:"target"`
	Type        pulumi.StringOutput    `pulumi:"type"`
	Value       pulumi.StringPtrOutput `pulumi:"value"`
	ValueFormat pulumi.StringPtrOutput `pulumi:"valueFormat"`
}


func NewWorkspaceConnection(ctx *pulumi.Context,
	name string, args *WorkspaceConnectionArgs, opts ...pulumi.ResourceOption) (*WorkspaceConnection, error) {
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
			Type: pulumi.String("azure-nextgen:machinelearningservices/v20210701:WorkspaceConnection"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices:WorkspaceConnection"),
		},
		{
			Type: pulumi.String("azure-nextgen:machinelearningservices:WorkspaceConnection"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20200601:WorkspaceConnection"),
		},
		{
			Type: pulumi.String("azure-nextgen:machinelearningservices/v20200601:WorkspaceConnection"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20200801:WorkspaceConnection"),
		},
		{
			Type: pulumi.String("azure-nextgen:machinelearningservices/v20200801:WorkspaceConnection"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20200901preview:WorkspaceConnection"),
		},
		{
			Type: pulumi.String("azure-nextgen:machinelearningservices/v20200901preview:WorkspaceConnection"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20210101:WorkspaceConnection"),
		},
		{
			Type: pulumi.String("azure-nextgen:machinelearningservices/v20210101:WorkspaceConnection"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20210301preview:WorkspaceConnection"),
		},
		{
			Type: pulumi.String("azure-nextgen:machinelearningservices/v20210301preview:WorkspaceConnection"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20210401:WorkspaceConnection"),
		},
		{
			Type: pulumi.String("azure-nextgen:machinelearningservices/v20210401:WorkspaceConnection"),
		},
	})
	opts = append(opts, aliases)
	var resource WorkspaceConnection
	err := ctx.RegisterResource("azure-native:machinelearningservices/v20210701:WorkspaceConnection", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetWorkspaceConnection(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WorkspaceConnectionState, opts ...pulumi.ResourceOption) (*WorkspaceConnection, error) {
	var resource WorkspaceConnection
	err := ctx.ReadResource("azure-native:machinelearningservices/v20210701:WorkspaceConnection", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type workspaceConnectionState struct {
}

type WorkspaceConnectionState struct {
}

func (WorkspaceConnectionState) ElementType() reflect.Type {
	return reflect.TypeOf((*workspaceConnectionState)(nil)).Elem()
}

type workspaceConnectionArgs struct {
	AuthType          *string `pulumi:"authType"`
	Category          *string `pulumi:"category"`
	ConnectionName    *string `pulumi:"connectionName"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
	Target            *string `pulumi:"target"`
	Value             *string `pulumi:"value"`
	ValueFormat       *string `pulumi:"valueFormat"`
	WorkspaceName     string  `pulumi:"workspaceName"`
}


type WorkspaceConnectionArgs struct {
	AuthType          pulumi.StringPtrInput
	Category          pulumi.StringPtrInput
	ConnectionName    pulumi.StringPtrInput
	ResourceGroupName pulumi.StringInput
	Target            pulumi.StringPtrInput
	Value             pulumi.StringPtrInput
	ValueFormat       pulumi.StringPtrInput
	WorkspaceName     pulumi.StringInput
}

func (WorkspaceConnectionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*workspaceConnectionArgs)(nil)).Elem()
}

type WorkspaceConnectionInput interface {
	pulumi.Input

	ToWorkspaceConnectionOutput() WorkspaceConnectionOutput
	ToWorkspaceConnectionOutputWithContext(ctx context.Context) WorkspaceConnectionOutput
}

func (*WorkspaceConnection) ElementType() reflect.Type {
	return reflect.TypeOf((*WorkspaceConnection)(nil))
}

func (i *WorkspaceConnection) ToWorkspaceConnectionOutput() WorkspaceConnectionOutput {
	return i.ToWorkspaceConnectionOutputWithContext(context.Background())
}

func (i *WorkspaceConnection) ToWorkspaceConnectionOutputWithContext(ctx context.Context) WorkspaceConnectionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkspaceConnectionOutput)
}

type WorkspaceConnectionOutput struct{ *pulumi.OutputState }

func (WorkspaceConnectionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WorkspaceConnection)(nil))
}

func (o WorkspaceConnectionOutput) ToWorkspaceConnectionOutput() WorkspaceConnectionOutput {
	return o
}

func (o WorkspaceConnectionOutput) ToWorkspaceConnectionOutputWithContext(ctx context.Context) WorkspaceConnectionOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*WorkspaceConnectionInput)(nil)).Elem(), &WorkspaceConnection{})
	pulumi.RegisterOutputType(WorkspaceConnectionOutput{})
}
