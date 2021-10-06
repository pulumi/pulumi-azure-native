


package agfoodplatform

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Extension struct {
	pulumi.CustomResourceState

	ETag                      pulumi.StringOutput      `pulumi:"eTag"`
	ExtensionApiDocsLink      pulumi.StringOutput      `pulumi:"extensionApiDocsLink"`
	ExtensionAuthLink         pulumi.StringOutput      `pulumi:"extensionAuthLink"`
	ExtensionCategory         pulumi.StringOutput      `pulumi:"extensionCategory"`
	ExtensionId               pulumi.StringOutput      `pulumi:"extensionId"`
	InstalledExtensionVersion pulumi.StringOutput      `pulumi:"installedExtensionVersion"`
	Name                      pulumi.StringOutput      `pulumi:"name"`
	SystemData                SystemDataResponseOutput `pulumi:"systemData"`
	Type                      pulumi.StringOutput      `pulumi:"type"`
}


func NewExtension(ctx *pulumi.Context,
	name string, args *ExtensionArgs, opts ...pulumi.ResourceOption) (*Extension, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.FarmBeatsResourceName == nil {
		return nil, errors.New("invalid value for required argument 'FarmBeatsResourceName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:agfoodplatform:Extension"),
		},
		{
			Type: pulumi.String("azure-native:agfoodplatform/v20200512preview:Extension"),
		},
		{
			Type: pulumi.String("azure-nextgen:agfoodplatform/v20200512preview:Extension"),
		},
	})
	opts = append(opts, aliases)
	var resource Extension
	err := ctx.RegisterResource("azure-native:agfoodplatform:Extension", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetExtension(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ExtensionState, opts ...pulumi.ResourceOption) (*Extension, error) {
	var resource Extension
	err := ctx.ReadResource("azure-native:agfoodplatform:Extension", name, id, state, &resource, opts...)
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
	ExtensionId           *string `pulumi:"extensionId"`
	FarmBeatsResourceName string  `pulumi:"farmBeatsResourceName"`
	ResourceGroupName     string  `pulumi:"resourceGroupName"`
}


type ExtensionArgs struct {
	ExtensionId           pulumi.StringPtrInput
	FarmBeatsResourceName pulumi.StringInput
	ResourceGroupName     pulumi.StringInput
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
	return reflect.TypeOf((*Extension)(nil))
}

func (i *Extension) ToExtensionOutput() ExtensionOutput {
	return i.ToExtensionOutputWithContext(context.Background())
}

func (i *Extension) ToExtensionOutputWithContext(ctx context.Context) ExtensionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExtensionOutput)
}

type ExtensionOutput struct{ *pulumi.OutputState }

func (ExtensionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Extension)(nil))
}

func (o ExtensionOutput) ToExtensionOutput() ExtensionOutput {
	return o
}

func (o ExtensionOutput) ToExtensionOutputWithContext(ctx context.Context) ExtensionOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ExtensionOutput{})
}
