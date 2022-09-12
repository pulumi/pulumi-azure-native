


package v20220901preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type EyesOn struct {
	pulumi.CustomResourceState

	Etag       pulumi.StringPtrOutput   `pulumi:"etag"`
	IsEnabled  pulumi.BoolOutput        `pulumi:"isEnabled"`
	Kind       pulumi.StringOutput      `pulumi:"kind"`
	Name       pulumi.StringOutput      `pulumi:"name"`
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	Type       pulumi.StringOutput      `pulumi:"type"`
}


func NewEyesOn(ctx *pulumi.Context,
	name string, args *EyesOnArgs, opts ...pulumi.ResourceOption) (*EyesOn, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Kind == nil {
		return nil, errors.New("invalid value for required argument 'Kind'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.WorkspaceName == nil {
		return nil, errors.New("invalid value for required argument 'WorkspaceName'")
	}
	args.Kind = pulumi.String("EyesOn")
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:securityinsights:EyesOn"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20190101preview:EyesOn"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20210301preview:EyesOn"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20210901preview:EyesOn"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20211001preview:EyesOn"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220101preview:EyesOn"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220401preview:EyesOn"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220501preview:EyesOn"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220601preview:EyesOn"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220701preview:EyesOn"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220801preview:EyesOn"),
		},
	})
	opts = append(opts, aliases)
	var resource EyesOn
	err := ctx.RegisterResource("azure-native:securityinsights/v20220901preview:EyesOn", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetEyesOn(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EyesOnState, opts ...pulumi.ResourceOption) (*EyesOn, error) {
	var resource EyesOn
	err := ctx.ReadResource("azure-native:securityinsights/v20220901preview:EyesOn", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type eyesOnState struct {
}

type EyesOnState struct {
}

func (EyesOnState) ElementType() reflect.Type {
	return reflect.TypeOf((*eyesOnState)(nil)).Elem()
}

type eyesOnArgs struct {
	Kind              string  `pulumi:"kind"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
	SettingsName      *string `pulumi:"settingsName"`
	WorkspaceName     string  `pulumi:"workspaceName"`
}


type EyesOnArgs struct {
	Kind              pulumi.StringInput
	ResourceGroupName pulumi.StringInput
	SettingsName      pulumi.StringPtrInput
	WorkspaceName     pulumi.StringInput
}

func (EyesOnArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*eyesOnArgs)(nil)).Elem()
}

type EyesOnInput interface {
	pulumi.Input

	ToEyesOnOutput() EyesOnOutput
	ToEyesOnOutputWithContext(ctx context.Context) EyesOnOutput
}

func (*EyesOn) ElementType() reflect.Type {
	return reflect.TypeOf((**EyesOn)(nil)).Elem()
}

func (i *EyesOn) ToEyesOnOutput() EyesOnOutput {
	return i.ToEyesOnOutputWithContext(context.Background())
}

func (i *EyesOn) ToEyesOnOutputWithContext(ctx context.Context) EyesOnOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EyesOnOutput)
}

type EyesOnOutput struct{ *pulumi.OutputState }

func (EyesOnOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EyesOn)(nil)).Elem()
}

func (o EyesOnOutput) ToEyesOnOutput() EyesOnOutput {
	return o
}

func (o EyesOnOutput) ToEyesOnOutputWithContext(ctx context.Context) EyesOnOutput {
	return o
}

func (o EyesOnOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EyesOn) pulumi.StringPtrOutput { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o EyesOnOutput) IsEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v *EyesOn) pulumi.BoolOutput { return v.IsEnabled }).(pulumi.BoolOutput)
}

func (o EyesOnOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *EyesOn) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

func (o EyesOnOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *EyesOn) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o EyesOnOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *EyesOn) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o EyesOnOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *EyesOn) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(EyesOnOutput{})
}
