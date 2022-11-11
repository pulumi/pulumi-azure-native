


package easm

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type LabelByWorkspace struct {
	pulumi.CustomResourceState

	Color             pulumi.StringPtrOutput   `pulumi:"color"`
	DisplayName       pulumi.StringPtrOutput   `pulumi:"displayName"`
	Name              pulumi.StringOutput      `pulumi:"name"`
	ProvisioningState pulumi.StringOutput      `pulumi:"provisioningState"`
	SystemData        SystemDataResponseOutput `pulumi:"systemData"`
	Type              pulumi.StringOutput      `pulumi:"type"`
}


func NewLabelByWorkspace(ctx *pulumi.Context,
	name string, args *LabelByWorkspaceArgs, opts ...pulumi.ResourceOption) (*LabelByWorkspace, error) {
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
			Type: pulumi.String("azure-native:easm/v20220401preview:LabelByWorkspace"),
		},
	})
	opts = append(opts, aliases)
	var resource LabelByWorkspace
	err := ctx.RegisterResource("azure-native:easm:LabelByWorkspace", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetLabelByWorkspace(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LabelByWorkspaceState, opts ...pulumi.ResourceOption) (*LabelByWorkspace, error) {
	var resource LabelByWorkspace
	err := ctx.ReadResource("azure-native:easm:LabelByWorkspace", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type labelByWorkspaceState struct {
}

type LabelByWorkspaceState struct {
}

func (LabelByWorkspaceState) ElementType() reflect.Type {
	return reflect.TypeOf((*labelByWorkspaceState)(nil)).Elem()
}

type labelByWorkspaceArgs struct {
	Color             *string `pulumi:"color"`
	DisplayName       *string `pulumi:"displayName"`
	LabelName         *string `pulumi:"labelName"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
	WorkspaceName     string  `pulumi:"workspaceName"`
}


type LabelByWorkspaceArgs struct {
	Color             pulumi.StringPtrInput
	DisplayName       pulumi.StringPtrInput
	LabelName         pulumi.StringPtrInput
	ResourceGroupName pulumi.StringInput
	WorkspaceName     pulumi.StringInput
}

func (LabelByWorkspaceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*labelByWorkspaceArgs)(nil)).Elem()
}

type LabelByWorkspaceInput interface {
	pulumi.Input

	ToLabelByWorkspaceOutput() LabelByWorkspaceOutput
	ToLabelByWorkspaceOutputWithContext(ctx context.Context) LabelByWorkspaceOutput
}

func (*LabelByWorkspace) ElementType() reflect.Type {
	return reflect.TypeOf((**LabelByWorkspace)(nil)).Elem()
}

func (i *LabelByWorkspace) ToLabelByWorkspaceOutput() LabelByWorkspaceOutput {
	return i.ToLabelByWorkspaceOutputWithContext(context.Background())
}

func (i *LabelByWorkspace) ToLabelByWorkspaceOutputWithContext(ctx context.Context) LabelByWorkspaceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LabelByWorkspaceOutput)
}

type LabelByWorkspaceOutput struct{ *pulumi.OutputState }

func (LabelByWorkspaceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LabelByWorkspace)(nil)).Elem()
}

func (o LabelByWorkspaceOutput) ToLabelByWorkspaceOutput() LabelByWorkspaceOutput {
	return o
}

func (o LabelByWorkspaceOutput) ToLabelByWorkspaceOutputWithContext(ctx context.Context) LabelByWorkspaceOutput {
	return o
}

func (o LabelByWorkspaceOutput) Color() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LabelByWorkspace) pulumi.StringPtrOutput { return v.Color }).(pulumi.StringPtrOutput)
}

func (o LabelByWorkspaceOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LabelByWorkspace) pulumi.StringPtrOutput { return v.DisplayName }).(pulumi.StringPtrOutput)
}

func (o LabelByWorkspaceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *LabelByWorkspace) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o LabelByWorkspaceOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *LabelByWorkspace) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LabelByWorkspaceOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *LabelByWorkspace) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LabelByWorkspaceOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *LabelByWorkspace) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LabelByWorkspaceOutput{})
}
