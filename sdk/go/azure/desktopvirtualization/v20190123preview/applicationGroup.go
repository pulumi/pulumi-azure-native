


package v20190123preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ApplicationGroup struct {
	pulumi.CustomResourceState

	ApplicationGroupType pulumi.StringOutput    `pulumi:"applicationGroupType"`
	Description          pulumi.StringPtrOutput `pulumi:"description"`
	FriendlyName         pulumi.StringPtrOutput `pulumi:"friendlyName"`
	HostPoolArmPath      pulumi.StringOutput    `pulumi:"hostPoolArmPath"`
	Location             pulumi.StringOutput    `pulumi:"location"`
	Name                 pulumi.StringOutput    `pulumi:"name"`
	Tags                 pulumi.StringMapOutput `pulumi:"tags"`
	Type                 pulumi.StringOutput    `pulumi:"type"`
	WorkspaceArmPath     pulumi.StringOutput    `pulumi:"workspaceArmPath"`
}


func NewApplicationGroup(ctx *pulumi.Context,
	name string, args *ApplicationGroupArgs, opts ...pulumi.ResourceOption) (*ApplicationGroup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ApplicationGroupType == nil {
		return nil, errors.New("invalid value for required argument 'ApplicationGroupType'")
	}
	if args.HostPoolArmPath == nil {
		return nil, errors.New("invalid value for required argument 'HostPoolArmPath'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:desktopvirtualization:ApplicationGroup"),
		},
		{
			Type: pulumi.String("azure-native:desktopvirtualization/v20190924preview:ApplicationGroup"),
		},
		{
			Type: pulumi.String("azure-native:desktopvirtualization/v20191210preview:ApplicationGroup"),
		},
		{
			Type: pulumi.String("azure-native:desktopvirtualization/v20200921preview:ApplicationGroup"),
		},
		{
			Type: pulumi.String("azure-native:desktopvirtualization/v20201019preview:ApplicationGroup"),
		},
		{
			Type: pulumi.String("azure-native:desktopvirtualization/v20201102preview:ApplicationGroup"),
		},
		{
			Type: pulumi.String("azure-native:desktopvirtualization/v20201110preview:ApplicationGroup"),
		},
		{
			Type: pulumi.String("azure-native:desktopvirtualization/v20210114preview:ApplicationGroup"),
		},
		{
			Type: pulumi.String("azure-native:desktopvirtualization/v20210201preview:ApplicationGroup"),
		},
		{
			Type: pulumi.String("azure-native:desktopvirtualization/v20210309preview:ApplicationGroup"),
		},
		{
			Type: pulumi.String("azure-native:desktopvirtualization/v20210401preview:ApplicationGroup"),
		},
		{
			Type: pulumi.String("azure-native:desktopvirtualization/v20210712:ApplicationGroup"),
		},
		{
			Type: pulumi.String("azure-native:desktopvirtualization/v20210903preview:ApplicationGroup"),
		},
	})
	opts = append(opts, aliases)
	var resource ApplicationGroup
	err := ctx.RegisterResource("azure-native:desktopvirtualization/v20190123preview:ApplicationGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetApplicationGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ApplicationGroupState, opts ...pulumi.ResourceOption) (*ApplicationGroup, error) {
	var resource ApplicationGroup
	err := ctx.ReadResource("azure-native:desktopvirtualization/v20190123preview:ApplicationGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type applicationGroupState struct {
}

type ApplicationGroupState struct {
}

func (ApplicationGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*applicationGroupState)(nil)).Elem()
}

type applicationGroupArgs struct {
	ApplicationGroupName *string           `pulumi:"applicationGroupName"`
	ApplicationGroupType string            `pulumi:"applicationGroupType"`
	Description          *string           `pulumi:"description"`
	FriendlyName         *string           `pulumi:"friendlyName"`
	HostPoolArmPath      string            `pulumi:"hostPoolArmPath"`
	Location             *string           `pulumi:"location"`
	ResourceGroupName    string            `pulumi:"resourceGroupName"`
	Tags                 map[string]string `pulumi:"tags"`
}


type ApplicationGroupArgs struct {
	ApplicationGroupName pulumi.StringPtrInput
	ApplicationGroupType pulumi.StringInput
	Description          pulumi.StringPtrInput
	FriendlyName         pulumi.StringPtrInput
	HostPoolArmPath      pulumi.StringInput
	Location             pulumi.StringPtrInput
	ResourceGroupName    pulumi.StringInput
	Tags                 pulumi.StringMapInput
}

func (ApplicationGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*applicationGroupArgs)(nil)).Elem()
}

type ApplicationGroupInput interface {
	pulumi.Input

	ToApplicationGroupOutput() ApplicationGroupOutput
	ToApplicationGroupOutputWithContext(ctx context.Context) ApplicationGroupOutput
}

func (*ApplicationGroup) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationGroup)(nil)).Elem()
}

func (i *ApplicationGroup) ToApplicationGroupOutput() ApplicationGroupOutput {
	return i.ToApplicationGroupOutputWithContext(context.Background())
}

func (i *ApplicationGroup) ToApplicationGroupOutputWithContext(ctx context.Context) ApplicationGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationGroupOutput)
}

type ApplicationGroupOutput struct{ *pulumi.OutputState }

func (ApplicationGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationGroup)(nil)).Elem()
}

func (o ApplicationGroupOutput) ToApplicationGroupOutput() ApplicationGroupOutput {
	return o
}

func (o ApplicationGroupOutput) ToApplicationGroupOutputWithContext(ctx context.Context) ApplicationGroupOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ApplicationGroupOutput{})
}
