


package v20201201

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type WebAppDeploymentSlot struct {
	pulumi.CustomResourceState

	Active      pulumi.BoolPtrOutput   `pulumi:"active"`
	Author      pulumi.StringPtrOutput `pulumi:"author"`
	AuthorEmail pulumi.StringPtrOutput `pulumi:"authorEmail"`
	Deployer    pulumi.StringPtrOutput `pulumi:"deployer"`
	Details     pulumi.StringPtrOutput `pulumi:"details"`
	EndTime     pulumi.StringPtrOutput `pulumi:"endTime"`
	Kind        pulumi.StringPtrOutput `pulumi:"kind"`
	Message     pulumi.StringPtrOutput `pulumi:"message"`
	Name        pulumi.StringOutput    `pulumi:"name"`
	StartTime   pulumi.StringPtrOutput `pulumi:"startTime"`
	Status      pulumi.IntPtrOutput    `pulumi:"status"`
	Type        pulumi.StringOutput    `pulumi:"type"`
}


func NewWebAppDeploymentSlot(ctx *pulumi.Context,
	name string, args *WebAppDeploymentSlotArgs, opts ...pulumi.ResourceOption) (*WebAppDeploymentSlot, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Name == nil {
		return nil, errors.New("invalid value for required argument 'Name'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.Slot == nil {
		return nil, errors.New("invalid value for required argument 'Slot'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:web:WebAppDeploymentSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20150801:WebAppDeploymentSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20160801:WebAppDeploymentSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20180201:WebAppDeploymentSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20181101:WebAppDeploymentSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20190801:WebAppDeploymentSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20200601:WebAppDeploymentSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20200901:WebAppDeploymentSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20201001:WebAppDeploymentSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210101:WebAppDeploymentSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210115:WebAppDeploymentSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210201:WebAppDeploymentSlot"),
		},
	})
	opts = append(opts, aliases)
	var resource WebAppDeploymentSlot
	err := ctx.RegisterResource("azure-native:web/v20201201:WebAppDeploymentSlot", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetWebAppDeploymentSlot(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WebAppDeploymentSlotState, opts ...pulumi.ResourceOption) (*WebAppDeploymentSlot, error) {
	var resource WebAppDeploymentSlot
	err := ctx.ReadResource("azure-native:web/v20201201:WebAppDeploymentSlot", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type webAppDeploymentSlotState struct {
}

type WebAppDeploymentSlotState struct {
}

func (WebAppDeploymentSlotState) ElementType() reflect.Type {
	return reflect.TypeOf((*webAppDeploymentSlotState)(nil)).Elem()
}

type webAppDeploymentSlotArgs struct {
	Active            *bool   `pulumi:"active"`
	Author            *string `pulumi:"author"`
	AuthorEmail       *string `pulumi:"authorEmail"`
	Deployer          *string `pulumi:"deployer"`
	Details           *string `pulumi:"details"`
	EndTime           *string `pulumi:"endTime"`
	Id                *string `pulumi:"id"`
	Kind              *string `pulumi:"kind"`
	Message           *string `pulumi:"message"`
	Name              string  `pulumi:"name"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
	Slot              string  `pulumi:"slot"`
	StartTime         *string `pulumi:"startTime"`
	Status            *int    `pulumi:"status"`
}


type WebAppDeploymentSlotArgs struct {
	Active            pulumi.BoolPtrInput
	Author            pulumi.StringPtrInput
	AuthorEmail       pulumi.StringPtrInput
	Deployer          pulumi.StringPtrInput
	Details           pulumi.StringPtrInput
	EndTime           pulumi.StringPtrInput
	Id                pulumi.StringPtrInput
	Kind              pulumi.StringPtrInput
	Message           pulumi.StringPtrInput
	Name              pulumi.StringInput
	ResourceGroupName pulumi.StringInput
	Slot              pulumi.StringInput
	StartTime         pulumi.StringPtrInput
	Status            pulumi.IntPtrInput
}

func (WebAppDeploymentSlotArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*webAppDeploymentSlotArgs)(nil)).Elem()
}

type WebAppDeploymentSlotInput interface {
	pulumi.Input

	ToWebAppDeploymentSlotOutput() WebAppDeploymentSlotOutput
	ToWebAppDeploymentSlotOutputWithContext(ctx context.Context) WebAppDeploymentSlotOutput
}

func (*WebAppDeploymentSlot) ElementType() reflect.Type {
	return reflect.TypeOf((*WebAppDeploymentSlot)(nil))
}

func (i *WebAppDeploymentSlot) ToWebAppDeploymentSlotOutput() WebAppDeploymentSlotOutput {
	return i.ToWebAppDeploymentSlotOutputWithContext(context.Background())
}

func (i *WebAppDeploymentSlot) ToWebAppDeploymentSlotOutputWithContext(ctx context.Context) WebAppDeploymentSlotOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebAppDeploymentSlotOutput)
}

type WebAppDeploymentSlotOutput struct{ *pulumi.OutputState }

func (WebAppDeploymentSlotOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WebAppDeploymentSlot)(nil))
}

func (o WebAppDeploymentSlotOutput) ToWebAppDeploymentSlotOutput() WebAppDeploymentSlotOutput {
	return o
}

func (o WebAppDeploymentSlotOutput) ToWebAppDeploymentSlotOutputWithContext(ctx context.Context) WebAppDeploymentSlotOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(WebAppDeploymentSlotOutput{})
}
