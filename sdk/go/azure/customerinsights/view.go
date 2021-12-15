


package customerinsights

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type View struct {
	pulumi.CustomResourceState

	Changed     pulumi.StringOutput    `pulumi:"changed"`
	Created     pulumi.StringOutput    `pulumi:"created"`
	Definition  pulumi.StringOutput    `pulumi:"definition"`
	DisplayName pulumi.StringMapOutput `pulumi:"displayName"`
	Name        pulumi.StringOutput    `pulumi:"name"`
	TenantId    pulumi.StringOutput    `pulumi:"tenantId"`
	Type        pulumi.StringOutput    `pulumi:"type"`
	UserId      pulumi.StringPtrOutput `pulumi:"userId"`
	ViewName    pulumi.StringOutput    `pulumi:"viewName"`
}


func NewView(ctx *pulumi.Context,
	name string, args *ViewArgs, opts ...pulumi.ResourceOption) (*View, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Definition == nil {
		return nil, errors.New("invalid value for required argument 'Definition'")
	}
	if args.HubName == nil {
		return nil, errors.New("invalid value for required argument 'HubName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:customerinsights/v20170101:View"),
		},
		{
			Type: pulumi.String("azure-native:customerinsights/v20170426:View"),
		},
	})
	opts = append(opts, aliases)
	var resource View
	err := ctx.RegisterResource("azure-native:customerinsights:View", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetView(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ViewState, opts ...pulumi.ResourceOption) (*View, error) {
	var resource View
	err := ctx.ReadResource("azure-native:customerinsights:View", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type viewState struct {
}

type ViewState struct {
}

func (ViewState) ElementType() reflect.Type {
	return reflect.TypeOf((*viewState)(nil)).Elem()
}

type viewArgs struct {
	Definition        string            `pulumi:"definition"`
	DisplayName       map[string]string `pulumi:"displayName"`
	HubName           string            `pulumi:"hubName"`
	ResourceGroupName string            `pulumi:"resourceGroupName"`
	UserId            *string           `pulumi:"userId"`
	ViewName          *string           `pulumi:"viewName"`
}


type ViewArgs struct {
	Definition        pulumi.StringInput
	DisplayName       pulumi.StringMapInput
	HubName           pulumi.StringInput
	ResourceGroupName pulumi.StringInput
	UserId            pulumi.StringPtrInput
	ViewName          pulumi.StringPtrInput
}

func (ViewArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*viewArgs)(nil)).Elem()
}

type ViewInput interface {
	pulumi.Input

	ToViewOutput() ViewOutput
	ToViewOutputWithContext(ctx context.Context) ViewOutput
}

func (*View) ElementType() reflect.Type {
	return reflect.TypeOf((**View)(nil)).Elem()
}

func (i *View) ToViewOutput() ViewOutput {
	return i.ToViewOutputWithContext(context.Background())
}

func (i *View) ToViewOutputWithContext(ctx context.Context) ViewOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ViewOutput)
}

type ViewOutput struct{ *pulumi.OutputState }

func (ViewOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**View)(nil)).Elem()
}

func (o ViewOutput) ToViewOutput() ViewOutput {
	return o
}

func (o ViewOutput) ToViewOutputWithContext(ctx context.Context) ViewOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ViewOutput{})
}
