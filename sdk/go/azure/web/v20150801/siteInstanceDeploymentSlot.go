


package v20150801

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


type SiteInstanceDeploymentSlot struct {
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


func NewSiteInstanceDeploymentSlot(ctx *pulumi.Context,
	name string, args *SiteInstanceDeploymentSlotArgs, opts ...pulumi.ResourceOption) (*SiteInstanceDeploymentSlot, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.InstanceId == nil {
		return nil, errors.New("invalid value for required argument 'InstanceId'")
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
	var resource SiteInstanceDeploymentSlot
	err := ctx.RegisterResource("azure-native:web/v20150801:SiteInstanceDeploymentSlot", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetSiteInstanceDeploymentSlot(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SiteInstanceDeploymentSlotState, opts ...pulumi.ResourceOption) (*SiteInstanceDeploymentSlot, error) {
	var resource SiteInstanceDeploymentSlot
	err := ctx.ReadResource("azure-native:web/v20150801:SiteInstanceDeploymentSlot", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type siteInstanceDeploymentSlotState struct {
}

type SiteInstanceDeploymentSlotState struct {
}

func (SiteInstanceDeploymentSlotState) ElementType() reflect.Type {
	return reflect.TypeOf((*siteInstanceDeploymentSlotState)(nil)).Elem()
}

type siteInstanceDeploymentSlotArgs struct {
	Active            *bool             `pulumi:"active"`
	Author            *string           `pulumi:"author"`
	AuthorEmail       *string           `pulumi:"authorEmail"`
	Deployer          *string           `pulumi:"deployer"`
	Details           *string           `pulumi:"details"`
	EndTime           *string           `pulumi:"endTime"`
	Id                *string           `pulumi:"id"`
	InstanceId        string            `pulumi:"instanceId"`
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


type SiteInstanceDeploymentSlotArgs struct {
	Active            pulumi.BoolPtrInput
	Author            pulumi.StringPtrInput
	AuthorEmail       pulumi.StringPtrInput
	Deployer          pulumi.StringPtrInput
	Details           pulumi.StringPtrInput
	EndTime           pulumi.StringPtrInput
	Id                pulumi.StringPtrInput
	InstanceId        pulumi.StringInput
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

func (SiteInstanceDeploymentSlotArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*siteInstanceDeploymentSlotArgs)(nil)).Elem()
}

type SiteInstanceDeploymentSlotInput interface {
	pulumi.Input

	ToSiteInstanceDeploymentSlotOutput() SiteInstanceDeploymentSlotOutput
	ToSiteInstanceDeploymentSlotOutputWithContext(ctx context.Context) SiteInstanceDeploymentSlotOutput
}

func (*SiteInstanceDeploymentSlot) ElementType() reflect.Type {
	return reflect.TypeOf((**SiteInstanceDeploymentSlot)(nil)).Elem()
}

func (i *SiteInstanceDeploymentSlot) ToSiteInstanceDeploymentSlotOutput() SiteInstanceDeploymentSlotOutput {
	return i.ToSiteInstanceDeploymentSlotOutputWithContext(context.Background())
}

func (i *SiteInstanceDeploymentSlot) ToSiteInstanceDeploymentSlotOutputWithContext(ctx context.Context) SiteInstanceDeploymentSlotOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SiteInstanceDeploymentSlotOutput)
}

type SiteInstanceDeploymentSlotOutput struct{ *pulumi.OutputState }

func (SiteInstanceDeploymentSlotOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SiteInstanceDeploymentSlot)(nil)).Elem()
}

func (o SiteInstanceDeploymentSlotOutput) ToSiteInstanceDeploymentSlotOutput() SiteInstanceDeploymentSlotOutput {
	return o
}

func (o SiteInstanceDeploymentSlotOutput) ToSiteInstanceDeploymentSlotOutputWithContext(ctx context.Context) SiteInstanceDeploymentSlotOutput {
	return o
}

func (o SiteInstanceDeploymentSlotOutput) Active() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SiteInstanceDeploymentSlot) pulumi.BoolPtrOutput { return v.Active }).(pulumi.BoolPtrOutput)
}

func (o SiteInstanceDeploymentSlotOutput) Author() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteInstanceDeploymentSlot) pulumi.StringPtrOutput { return v.Author }).(pulumi.StringPtrOutput)
}

func (o SiteInstanceDeploymentSlotOutput) AuthorEmail() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteInstanceDeploymentSlot) pulumi.StringPtrOutput { return v.AuthorEmail }).(pulumi.StringPtrOutput)
}

func (o SiteInstanceDeploymentSlotOutput) Deployer() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteInstanceDeploymentSlot) pulumi.StringPtrOutput { return v.Deployer }).(pulumi.StringPtrOutput)
}

func (o SiteInstanceDeploymentSlotOutput) Details() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteInstanceDeploymentSlot) pulumi.StringPtrOutput { return v.Details }).(pulumi.StringPtrOutput)
}

func (o SiteInstanceDeploymentSlotOutput) EndTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteInstanceDeploymentSlot) pulumi.StringPtrOutput { return v.EndTime }).(pulumi.StringPtrOutput)
}

func (o SiteInstanceDeploymentSlotOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteInstanceDeploymentSlot) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o SiteInstanceDeploymentSlotOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *SiteInstanceDeploymentSlot) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o SiteInstanceDeploymentSlotOutput) Message() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteInstanceDeploymentSlot) pulumi.StringPtrOutput { return v.Message }).(pulumi.StringPtrOutput)
}

func (o SiteInstanceDeploymentSlotOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteInstanceDeploymentSlot) pulumi.StringPtrOutput { return v.Name }).(pulumi.StringPtrOutput)
}

func (o SiteInstanceDeploymentSlotOutput) StartTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteInstanceDeploymentSlot) pulumi.StringPtrOutput { return v.StartTime }).(pulumi.StringPtrOutput)
}

func (o SiteInstanceDeploymentSlotOutput) Status() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SiteInstanceDeploymentSlot) pulumi.IntPtrOutput { return v.Status }).(pulumi.IntPtrOutput)
}

func (o SiteInstanceDeploymentSlotOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *SiteInstanceDeploymentSlot) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o SiteInstanceDeploymentSlotOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteInstanceDeploymentSlot) pulumi.StringPtrOutput { return v.Type }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(SiteInstanceDeploymentSlotOutput{})
}
