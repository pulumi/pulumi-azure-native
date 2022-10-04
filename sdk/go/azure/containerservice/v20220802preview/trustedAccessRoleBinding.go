


package v20220802preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type TrustedAccessRoleBinding struct {
	pulumi.CustomResourceState

	Name              pulumi.StringOutput      `pulumi:"name"`
	ProvisioningState pulumi.StringOutput      `pulumi:"provisioningState"`
	Roles             pulumi.StringArrayOutput `pulumi:"roles"`
	SourceResourceId  pulumi.StringOutput      `pulumi:"sourceResourceId"`
	SystemData        SystemDataResponseOutput `pulumi:"systemData"`
	Type              pulumi.StringOutput      `pulumi:"type"`
}


func NewTrustedAccessRoleBinding(ctx *pulumi.Context,
	name string, args *TrustedAccessRoleBindingArgs, opts ...pulumi.ResourceOption) (*TrustedAccessRoleBinding, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ResourceName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceName'")
	}
	if args.Roles == nil {
		return nil, errors.New("invalid value for required argument 'Roles'")
	}
	if args.SourceResourceId == nil {
		return nil, errors.New("invalid value for required argument 'SourceResourceId'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:containerservice:TrustedAccessRoleBinding"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20220402preview:TrustedAccessRoleBinding"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20220502preview:TrustedAccessRoleBinding"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20220602preview:TrustedAccessRoleBinding"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20220702preview:TrustedAccessRoleBinding"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20220803preview:TrustedAccessRoleBinding"),
		},
	})
	opts = append(opts, aliases)
	var resource TrustedAccessRoleBinding
	err := ctx.RegisterResource("azure-native:containerservice/v20220802preview:TrustedAccessRoleBinding", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetTrustedAccessRoleBinding(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TrustedAccessRoleBindingState, opts ...pulumi.ResourceOption) (*TrustedAccessRoleBinding, error) {
	var resource TrustedAccessRoleBinding
	err := ctx.ReadResource("azure-native:containerservice/v20220802preview:TrustedAccessRoleBinding", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type trustedAccessRoleBindingState struct {
}

type TrustedAccessRoleBindingState struct {
}

func (TrustedAccessRoleBindingState) ElementType() reflect.Type {
	return reflect.TypeOf((*trustedAccessRoleBindingState)(nil)).Elem()
}

type trustedAccessRoleBindingArgs struct {
	ResourceGroupName            string   `pulumi:"resourceGroupName"`
	ResourceName                 string   `pulumi:"resourceName"`
	Roles                        []string `pulumi:"roles"`
	SourceResourceId             string   `pulumi:"sourceResourceId"`
	TrustedAccessRoleBindingName *string  `pulumi:"trustedAccessRoleBindingName"`
}


type TrustedAccessRoleBindingArgs struct {
	ResourceGroupName            pulumi.StringInput
	ResourceName                 pulumi.StringInput
	Roles                        pulumi.StringArrayInput
	SourceResourceId             pulumi.StringInput
	TrustedAccessRoleBindingName pulumi.StringPtrInput
}

func (TrustedAccessRoleBindingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*trustedAccessRoleBindingArgs)(nil)).Elem()
}

type TrustedAccessRoleBindingInput interface {
	pulumi.Input

	ToTrustedAccessRoleBindingOutput() TrustedAccessRoleBindingOutput
	ToTrustedAccessRoleBindingOutputWithContext(ctx context.Context) TrustedAccessRoleBindingOutput
}

func (*TrustedAccessRoleBinding) ElementType() reflect.Type {
	return reflect.TypeOf((**TrustedAccessRoleBinding)(nil)).Elem()
}

func (i *TrustedAccessRoleBinding) ToTrustedAccessRoleBindingOutput() TrustedAccessRoleBindingOutput {
	return i.ToTrustedAccessRoleBindingOutputWithContext(context.Background())
}

func (i *TrustedAccessRoleBinding) ToTrustedAccessRoleBindingOutputWithContext(ctx context.Context) TrustedAccessRoleBindingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TrustedAccessRoleBindingOutput)
}

type TrustedAccessRoleBindingOutput struct{ *pulumi.OutputState }

func (TrustedAccessRoleBindingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TrustedAccessRoleBinding)(nil)).Elem()
}

func (o TrustedAccessRoleBindingOutput) ToTrustedAccessRoleBindingOutput() TrustedAccessRoleBindingOutput {
	return o
}

func (o TrustedAccessRoleBindingOutput) ToTrustedAccessRoleBindingOutputWithContext(ctx context.Context) TrustedAccessRoleBindingOutput {
	return o
}

func (o TrustedAccessRoleBindingOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *TrustedAccessRoleBinding) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o TrustedAccessRoleBindingOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *TrustedAccessRoleBinding) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o TrustedAccessRoleBindingOutput) Roles() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *TrustedAccessRoleBinding) pulumi.StringArrayOutput { return v.Roles }).(pulumi.StringArrayOutput)
}

func (o TrustedAccessRoleBindingOutput) SourceResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v *TrustedAccessRoleBinding) pulumi.StringOutput { return v.SourceResourceId }).(pulumi.StringOutput)
}

func (o TrustedAccessRoleBindingOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *TrustedAccessRoleBinding) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o TrustedAccessRoleBindingOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *TrustedAccessRoleBinding) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(TrustedAccessRoleBindingOutput{})
}
