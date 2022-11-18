


package v20220101preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Linker struct {
	pulumi.CustomResourceState

	AuthInfo          pulumi.AnyOutput              `pulumi:"authInfo"`
	ClientType        pulumi.StringPtrOutput        `pulumi:"clientType"`
	Name              pulumi.StringOutput           `pulumi:"name"`
	ProvisioningState pulumi.StringOutput           `pulumi:"provisioningState"`
	Scope             pulumi.StringPtrOutput        `pulumi:"scope"`
	SecretStore       SecretStoreResponsePtrOutput  `pulumi:"secretStore"`
	SystemData        SystemDataResponseOutput      `pulumi:"systemData"`
	TargetService     pulumi.AnyOutput              `pulumi:"targetService"`
	Type              pulumi.StringOutput           `pulumi:"type"`
	VNetSolution      VNetSolutionResponsePtrOutput `pulumi:"vNetSolution"`
}


func NewLinker(ctx *pulumi.Context,
	name string, args *LinkerArgs, opts ...pulumi.ResourceOption) (*Linker, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceUri == nil {
		return nil, errors.New("invalid value for required argument 'ResourceUri'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:servicelinker:Linker"),
		},
		{
			Type: pulumi.String("azure-native:servicelinker/v20211101preview:Linker"),
		},
		{
			Type: pulumi.String("azure-native:servicelinker/v20220501:Linker"),
		},
		{
			Type: pulumi.String("azure-native:servicelinker/v20221101preview:Linker"),
		},
	})
	opts = append(opts, aliases)
	var resource Linker
	err := ctx.RegisterResource("azure-native:servicelinker/v20220101preview:Linker", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetLinker(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LinkerState, opts ...pulumi.ResourceOption) (*Linker, error) {
	var resource Linker
	err := ctx.ReadResource("azure-native:servicelinker/v20220101preview:Linker", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type linkerState struct {
}

type LinkerState struct {
}

func (LinkerState) ElementType() reflect.Type {
	return reflect.TypeOf((*linkerState)(nil)).Elem()
}

type linkerArgs struct {
	AuthInfo      interface{}   `pulumi:"authInfo"`
	ClientType    *string       `pulumi:"clientType"`
	LinkerName    *string       `pulumi:"linkerName"`
	ResourceUri   string        `pulumi:"resourceUri"`
	Scope         *string       `pulumi:"scope"`
	SecretStore   *SecretStore  `pulumi:"secretStore"`
	TargetService interface{}   `pulumi:"targetService"`
	VNetSolution  *VNetSolution `pulumi:"vNetSolution"`
}


type LinkerArgs struct {
	AuthInfo      pulumi.Input
	ClientType    pulumi.StringPtrInput
	LinkerName    pulumi.StringPtrInput
	ResourceUri   pulumi.StringInput
	Scope         pulumi.StringPtrInput
	SecretStore   SecretStorePtrInput
	TargetService pulumi.Input
	VNetSolution  VNetSolutionPtrInput
}

func (LinkerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*linkerArgs)(nil)).Elem()
}

type LinkerInput interface {
	pulumi.Input

	ToLinkerOutput() LinkerOutput
	ToLinkerOutputWithContext(ctx context.Context) LinkerOutput
}

func (*Linker) ElementType() reflect.Type {
	return reflect.TypeOf((**Linker)(nil)).Elem()
}

func (i *Linker) ToLinkerOutput() LinkerOutput {
	return i.ToLinkerOutputWithContext(context.Background())
}

func (i *Linker) ToLinkerOutputWithContext(ctx context.Context) LinkerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LinkerOutput)
}

type LinkerOutput struct{ *pulumi.OutputState }

func (LinkerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Linker)(nil)).Elem()
}

func (o LinkerOutput) ToLinkerOutput() LinkerOutput {
	return o
}

func (o LinkerOutput) ToLinkerOutputWithContext(ctx context.Context) LinkerOutput {
	return o
}

func (o LinkerOutput) AuthInfo() pulumi.AnyOutput {
	return o.ApplyT(func(v *Linker) pulumi.AnyOutput { return v.AuthInfo }).(pulumi.AnyOutput)
}

func (o LinkerOutput) ClientType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Linker) pulumi.StringPtrOutput { return v.ClientType }).(pulumi.StringPtrOutput)
}

func (o LinkerOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Linker) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o LinkerOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *Linker) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LinkerOutput) Scope() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Linker) pulumi.StringPtrOutput { return v.Scope }).(pulumi.StringPtrOutput)
}

func (o LinkerOutput) SecretStore() SecretStoreResponsePtrOutput {
	return o.ApplyT(func(v *Linker) SecretStoreResponsePtrOutput { return v.SecretStore }).(SecretStoreResponsePtrOutput)
}

func (o LinkerOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *Linker) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LinkerOutput) TargetService() pulumi.AnyOutput {
	return o.ApplyT(func(v *Linker) pulumi.AnyOutput { return v.TargetService }).(pulumi.AnyOutput)
}

func (o LinkerOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Linker) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o LinkerOutput) VNetSolution() VNetSolutionResponsePtrOutput {
	return o.ApplyT(func(v *Linker) VNetSolutionResponsePtrOutput { return v.VNetSolution }).(VNetSolutionResponsePtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LinkerOutput{})
}
