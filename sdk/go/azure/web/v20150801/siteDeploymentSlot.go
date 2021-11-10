


package v20150801

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type SiteDeploymentSlot struct {
	pulumi.CustomResourceState

	Active      pulumi.BoolPtrOutput   `pulumi:"active"`
	Author      pulumi.StringPtrOutput `pulumi:"author"`
	AuthorEmail pulumi.StringPtrOutput `pulumi:"authorEmail"`
	Deployer    pulumi.StringPtrOutput `pulumi:"deployer"`
	Details     pulumi.StringPtrOutput `pulumi:"details"`
	EndTime     pulumi.StringPtrOutput `pulumi:"endTime"`
	Kind        pulumi.StringPtrOutput `pulumi:"kind"`
	Location    pulumi.StringOutput    `pulumi:"location"`
	Message     pulumi.StringPtrOutput `pulumi:"message"`
	Name        pulumi.StringPtrOutput `pulumi:"name"`
	StartTime   pulumi.StringPtrOutput `pulumi:"startTime"`
	Status      pulumi.IntPtrOutput    `pulumi:"status"`
	Tags        pulumi.StringMapOutput `pulumi:"tags"`
	Type        pulumi.StringPtrOutput `pulumi:"type"`
}


func NewSiteDeploymentSlot(ctx *pulumi.Context,
	name string, args *SiteDeploymentSlotArgs, opts ...pulumi.ResourceOption) (*SiteDeploymentSlot, error) {
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
			Type: pulumi.String("azure-native:web:SiteDeploymentSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20160801:SiteDeploymentSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20180201:SiteDeploymentSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20181101:SiteDeploymentSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20190801:SiteDeploymentSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20200601:SiteDeploymentSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20200901:SiteDeploymentSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20201001:SiteDeploymentSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20201201:SiteDeploymentSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210101:SiteDeploymentSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210115:SiteDeploymentSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210201:SiteDeploymentSlot"),
		},
	})
	opts = append(opts, aliases)
	var resource SiteDeploymentSlot
	err := ctx.RegisterResource("azure-native:web/v20150801:SiteDeploymentSlot", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetSiteDeploymentSlot(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SiteDeploymentSlotState, opts ...pulumi.ResourceOption) (*SiteDeploymentSlot, error) {
	var resource SiteDeploymentSlot
	err := ctx.ReadResource("azure-native:web/v20150801:SiteDeploymentSlot", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type siteDeploymentSlotState struct {
}

type SiteDeploymentSlotState struct {
}

func (SiteDeploymentSlotState) ElementType() reflect.Type {
	return reflect.TypeOf((*siteDeploymentSlotState)(nil)).Elem()
}

type siteDeploymentSlotArgs struct {
	Active            *bool             `pulumi:"active"`
	Author            *string           `pulumi:"author"`
	AuthorEmail       *string           `pulumi:"authorEmail"`
	Deployer          *string           `pulumi:"deployer"`
	Details           *string           `pulumi:"details"`
	EndTime           *string           `pulumi:"endTime"`
	Id                *string           `pulumi:"id"`
	Kind              *string           `pulumi:"kind"`
	Location          *string           `pulumi:"location"`
	Message           *string           `pulumi:"message"`
	Name              string            `pulumi:"name"`
	ResourceGroupName string            `pulumi:"resourceGroupName"`
	Slot              string            `pulumi:"slot"`
	StartTime         *string           `pulumi:"startTime"`
	Status            *int              `pulumi:"status"`
	Tags              map[string]string `pulumi:"tags"`
	Type              *string           `pulumi:"type"`
}


type SiteDeploymentSlotArgs struct {
	Active            pulumi.BoolPtrInput
	Author            pulumi.StringPtrInput
	AuthorEmail       pulumi.StringPtrInput
	Deployer          pulumi.StringPtrInput
	Details           pulumi.StringPtrInput
	EndTime           pulumi.StringPtrInput
	Id                pulumi.StringPtrInput
	Kind              pulumi.StringPtrInput
	Location          pulumi.StringPtrInput
	Message           pulumi.StringPtrInput
	Name              pulumi.StringInput
	ResourceGroupName pulumi.StringInput
	Slot              pulumi.StringInput
	StartTime         pulumi.StringPtrInput
	Status            pulumi.IntPtrInput
	Tags              pulumi.StringMapInput
	Type              pulumi.StringPtrInput
}

func (SiteDeploymentSlotArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*siteDeploymentSlotArgs)(nil)).Elem()
}

type SiteDeploymentSlotInput interface {
	pulumi.Input

	ToSiteDeploymentSlotOutput() SiteDeploymentSlotOutput
	ToSiteDeploymentSlotOutputWithContext(ctx context.Context) SiteDeploymentSlotOutput
}

func (*SiteDeploymentSlot) ElementType() reflect.Type {
	return reflect.TypeOf((*SiteDeploymentSlot)(nil))
}

func (i *SiteDeploymentSlot) ToSiteDeploymentSlotOutput() SiteDeploymentSlotOutput {
	return i.ToSiteDeploymentSlotOutputWithContext(context.Background())
}

func (i *SiteDeploymentSlot) ToSiteDeploymentSlotOutputWithContext(ctx context.Context) SiteDeploymentSlotOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SiteDeploymentSlotOutput)
}

type SiteDeploymentSlotOutput struct{ *pulumi.OutputState }

func (SiteDeploymentSlotOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SiteDeploymentSlot)(nil))
}

func (o SiteDeploymentSlotOutput) ToSiteDeploymentSlotOutput() SiteDeploymentSlotOutput {
	return o
}

func (o SiteDeploymentSlotOutput) ToSiteDeploymentSlotOutputWithContext(ctx context.Context) SiteDeploymentSlotOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(SiteDeploymentSlotOutput{})
}
