


package v20220501

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type WorkspaceConnection struct {
	pulumi.CustomResourceState

	Name       pulumi.StringOutput      `pulumi:"name"`
	Properties pulumi.AnyOutput         `pulumi:"properties"`
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	Type       pulumi.StringOutput      `pulumi:"type"`
}


func NewWorkspaceConnection(ctx *pulumi.Context,
	name string, args *WorkspaceConnectionArgs, opts ...pulumi.ResourceOption) (*WorkspaceConnection, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Properties == nil {
		return nil, errors.New("invalid value for required argument 'Properties'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.WorkspaceName == nil {
		return nil, errors.New("invalid value for required argument 'WorkspaceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:machinelearningservices:WorkspaceConnection"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20200601:WorkspaceConnection"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20200801:WorkspaceConnection"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20200901preview:WorkspaceConnection"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20210101:WorkspaceConnection"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20210301preview:WorkspaceConnection"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20210401:WorkspaceConnection"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20210701:WorkspaceConnection"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20220101preview:WorkspaceConnection"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20220201preview:WorkspaceConnection"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20220601preview:WorkspaceConnection"),
		},
	})
	opts = append(opts, aliases)
	var resource WorkspaceConnection
	err := ctx.RegisterResource("azure-native:machinelearningservices/v20220501:WorkspaceConnection", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetWorkspaceConnection(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WorkspaceConnectionState, opts ...pulumi.ResourceOption) (*WorkspaceConnection, error) {
	var resource WorkspaceConnection
	err := ctx.ReadResource("azure-native:machinelearningservices/v20220501:WorkspaceConnection", name, id, state, &resource, opts...)
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
	ConnectionName    *string     `pulumi:"connectionName"`
	Properties        interface{} `pulumi:"properties"`
	ResourceGroupName string      `pulumi:"resourceGroupName"`
	WorkspaceName     string      `pulumi:"workspaceName"`
}


type WorkspaceConnectionArgs struct {
	ConnectionName    pulumi.StringPtrInput
	Properties        pulumi.Input
	ResourceGroupName pulumi.StringInput
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
	return reflect.TypeOf((**WorkspaceConnection)(nil)).Elem()
}

func (i *WorkspaceConnection) ToWorkspaceConnectionOutput() WorkspaceConnectionOutput {
	return i.ToWorkspaceConnectionOutputWithContext(context.Background())
}

func (i *WorkspaceConnection) ToWorkspaceConnectionOutputWithContext(ctx context.Context) WorkspaceConnectionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkspaceConnectionOutput)
}

type WorkspaceConnectionOutput struct{ *pulumi.OutputState }

func (WorkspaceConnectionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkspaceConnection)(nil)).Elem()
}

func (o WorkspaceConnectionOutput) ToWorkspaceConnectionOutput() WorkspaceConnectionOutput {
	return o
}

func (o WorkspaceConnectionOutput) ToWorkspaceConnectionOutputWithContext(ctx context.Context) WorkspaceConnectionOutput {
	return o
}

func (o WorkspaceConnectionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkspaceConnection) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o WorkspaceConnectionOutput) Properties() pulumi.AnyOutput {
	return o.ApplyT(func(v *WorkspaceConnection) pulumi.AnyOutput { return v.Properties }).(pulumi.AnyOutput)
}

func (o WorkspaceConnectionOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *WorkspaceConnection) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o WorkspaceConnectionOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkspaceConnection) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(WorkspaceConnectionOutput{})
}
