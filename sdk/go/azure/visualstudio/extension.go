


package visualstudio

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Extension struct {
	pulumi.CustomResourceState

	Location   pulumi.StringPtrOutput                 `pulumi:"location"`
	Name       pulumi.StringOutput                    `pulumi:"name"`
	Plan       ExtensionResourcePlanResponsePtrOutput `pulumi:"plan"`
	Properties pulumi.StringMapOutput                 `pulumi:"properties"`
	Tags       pulumi.StringMapOutput                 `pulumi:"tags"`
	Type       pulumi.StringOutput                    `pulumi:"type"`
}


func NewExtension(ctx *pulumi.Context,
	name string, args *ExtensionArgs, opts ...pulumi.ResourceOption) (*Extension, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccountResourceName == nil {
		return nil, errors.New("invalid value for required argument 'AccountResourceName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:visualstudio/v20140401preview:Extension"),
		},
		{
			Type: pulumi.String("azure-native:visualstudio/v20171101preview:Extension"),
		},
	})
	opts = append(opts, aliases)
	var resource Extension
	err := ctx.RegisterResource("azure-native:visualstudio:Extension", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetExtension(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ExtensionState, opts ...pulumi.ResourceOption) (*Extension, error) {
	var resource Extension
	err := ctx.ReadResource("azure-native:visualstudio:Extension", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type extensionState struct {
}

type ExtensionState struct {
}

func (ExtensionState) ElementType() reflect.Type {
	return reflect.TypeOf((*extensionState)(nil)).Elem()
}

type extensionArgs struct {
	AccountResourceName   string                 `pulumi:"accountResourceName"`
	ExtensionResourceName *string                `pulumi:"extensionResourceName"`
	Location              *string                `pulumi:"location"`
	Plan                  *ExtensionResourcePlan `pulumi:"plan"`
	Properties            map[string]string      `pulumi:"properties"`
	ResourceGroupName     string                 `pulumi:"resourceGroupName"`
	Tags                  map[string]string      `pulumi:"tags"`
}


type ExtensionArgs struct {
	AccountResourceName   pulumi.StringInput
	ExtensionResourceName pulumi.StringPtrInput
	Location              pulumi.StringPtrInput
	Plan                  ExtensionResourcePlanPtrInput
	Properties            pulumi.StringMapInput
	ResourceGroupName     pulumi.StringInput
	Tags                  pulumi.StringMapInput
}

func (ExtensionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*extensionArgs)(nil)).Elem()
}

type ExtensionInput interface {
	pulumi.Input

	ToExtensionOutput() ExtensionOutput
	ToExtensionOutputWithContext(ctx context.Context) ExtensionOutput
}

func (*Extension) ElementType() reflect.Type {
	return reflect.TypeOf((**Extension)(nil)).Elem()
}

func (i *Extension) ToExtensionOutput() ExtensionOutput {
	return i.ToExtensionOutputWithContext(context.Background())
}

func (i *Extension) ToExtensionOutputWithContext(ctx context.Context) ExtensionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExtensionOutput)
}

type ExtensionOutput struct{ *pulumi.OutputState }

func (ExtensionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Extension)(nil)).Elem()
}

func (o ExtensionOutput) ToExtensionOutput() ExtensionOutput {
	return o
}

func (o ExtensionOutput) ToExtensionOutputWithContext(ctx context.Context) ExtensionOutput {
	return o
}

func (o ExtensionOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Extension) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

func (o ExtensionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Extension) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ExtensionOutput) Plan() ExtensionResourcePlanResponsePtrOutput {
	return o.ApplyT(func(v *Extension) ExtensionResourcePlanResponsePtrOutput { return v.Plan }).(ExtensionResourcePlanResponsePtrOutput)
}

func (o ExtensionOutput) Properties() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Extension) pulumi.StringMapOutput { return v.Properties }).(pulumi.StringMapOutput)
}

func (o ExtensionOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Extension) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o ExtensionOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Extension) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ExtensionOutput{})
}
