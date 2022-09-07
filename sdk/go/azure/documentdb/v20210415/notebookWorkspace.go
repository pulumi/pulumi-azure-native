


package v20210415

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type NotebookWorkspace struct {
	pulumi.CustomResourceState

	Name                   pulumi.StringOutput `pulumi:"name"`
	NotebookServerEndpoint pulumi.StringOutput `pulumi:"notebookServerEndpoint"`
	Status                 pulumi.StringOutput `pulumi:"status"`
	Type                   pulumi.StringOutput `pulumi:"type"`
}


func NewNotebookWorkspace(ctx *pulumi.Context,
	name string, args *NotebookWorkspaceArgs, opts ...pulumi.ResourceOption) (*NotebookWorkspace, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccountName == nil {
		return nil, errors.New("invalid value for required argument 'AccountName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:documentdb:NotebookWorkspace"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20190801:NotebookWorkspace"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20191212:NotebookWorkspace"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20200301:NotebookWorkspace"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20200401:NotebookWorkspace"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20200601preview:NotebookWorkspace"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20200901:NotebookWorkspace"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210115:NotebookWorkspace"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210301preview:NotebookWorkspace"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210315:NotebookWorkspace"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210401preview:NotebookWorkspace"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210515:NotebookWorkspace"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210615:NotebookWorkspace"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210701preview:NotebookWorkspace"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20211015:NotebookWorkspace"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20211015preview:NotebookWorkspace"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20211115preview:NotebookWorkspace"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20220215preview:NotebookWorkspace"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20220515:NotebookWorkspace"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20220515preview:NotebookWorkspace"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20220815:NotebookWorkspace"),
		},
	})
	opts = append(opts, aliases)
	var resource NotebookWorkspace
	err := ctx.RegisterResource("azure-native:documentdb/v20210415:NotebookWorkspace", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetNotebookWorkspace(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NotebookWorkspaceState, opts ...pulumi.ResourceOption) (*NotebookWorkspace, error) {
	var resource NotebookWorkspace
	err := ctx.ReadResource("azure-native:documentdb/v20210415:NotebookWorkspace", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type notebookWorkspaceState struct {
}

type NotebookWorkspaceState struct {
}

func (NotebookWorkspaceState) ElementType() reflect.Type {
	return reflect.TypeOf((*notebookWorkspaceState)(nil)).Elem()
}

type notebookWorkspaceArgs struct {
	AccountName           string  `pulumi:"accountName"`
	NotebookWorkspaceName *string `pulumi:"notebookWorkspaceName"`
	ResourceGroupName     string  `pulumi:"resourceGroupName"`
}


type NotebookWorkspaceArgs struct {
	AccountName           pulumi.StringInput
	NotebookWorkspaceName pulumi.StringPtrInput
	ResourceGroupName     pulumi.StringInput
}

func (NotebookWorkspaceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*notebookWorkspaceArgs)(nil)).Elem()
}

type NotebookWorkspaceInput interface {
	pulumi.Input

	ToNotebookWorkspaceOutput() NotebookWorkspaceOutput
	ToNotebookWorkspaceOutputWithContext(ctx context.Context) NotebookWorkspaceOutput
}

func (*NotebookWorkspace) ElementType() reflect.Type {
	return reflect.TypeOf((**NotebookWorkspace)(nil)).Elem()
}

func (i *NotebookWorkspace) ToNotebookWorkspaceOutput() NotebookWorkspaceOutput {
	return i.ToNotebookWorkspaceOutputWithContext(context.Background())
}

func (i *NotebookWorkspace) ToNotebookWorkspaceOutputWithContext(ctx context.Context) NotebookWorkspaceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NotebookWorkspaceOutput)
}

type NotebookWorkspaceOutput struct{ *pulumi.OutputState }

func (NotebookWorkspaceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NotebookWorkspace)(nil)).Elem()
}

func (o NotebookWorkspaceOutput) ToNotebookWorkspaceOutput() NotebookWorkspaceOutput {
	return o
}

func (o NotebookWorkspaceOutput) ToNotebookWorkspaceOutputWithContext(ctx context.Context) NotebookWorkspaceOutput {
	return o
}

func (o NotebookWorkspaceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *NotebookWorkspace) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o NotebookWorkspaceOutput) NotebookServerEndpoint() pulumi.StringOutput {
	return o.ApplyT(func(v *NotebookWorkspace) pulumi.StringOutput { return v.NotebookServerEndpoint }).(pulumi.StringOutput)
}

func (o NotebookWorkspaceOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *NotebookWorkspace) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

func (o NotebookWorkspaceOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *NotebookWorkspace) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(NotebookWorkspaceOutput{})
}
