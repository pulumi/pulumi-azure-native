


package v20150801

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type SiteInstanceDeployment struct {
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


func NewSiteInstanceDeployment(ctx *pulumi.Context,
	name string, args *SiteInstanceDeploymentArgs, opts ...pulumi.ResourceOption) (*SiteInstanceDeployment, error) {
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
	var resource SiteInstanceDeployment
	err := ctx.RegisterResource("azure-native:web/v20150801:SiteInstanceDeployment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetSiteInstanceDeployment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SiteInstanceDeploymentState, opts ...pulumi.ResourceOption) (*SiteInstanceDeployment, error) {
	var resource SiteInstanceDeployment
	err := ctx.ReadResource("azure-native:web/v20150801:SiteInstanceDeployment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type siteInstanceDeploymentState struct {
}

type SiteInstanceDeploymentState struct {
}

func (SiteInstanceDeploymentState) ElementType() reflect.Type {
	return reflect.TypeOf((*siteInstanceDeploymentState)(nil)).Elem()
}

type siteInstanceDeploymentArgs struct {
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
	StartTime         *string           `pulumi:"startTime"`
	Status            *int              `pulumi:"status"`
	Tags              map[string]string `pulumi:"tags"`
	Type              *string           `pulumi:"type"`
}


type SiteInstanceDeploymentArgs struct {
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
	StartTime         pulumi.StringPtrInput
	Status            pulumi.IntPtrInput
	Tags              pulumi.StringMapInput
	Type              pulumi.StringPtrInput
}

func (SiteInstanceDeploymentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*siteInstanceDeploymentArgs)(nil)).Elem()
}

type SiteInstanceDeploymentInput interface {
	pulumi.Input

	ToSiteInstanceDeploymentOutput() SiteInstanceDeploymentOutput
	ToSiteInstanceDeploymentOutputWithContext(ctx context.Context) SiteInstanceDeploymentOutput
}

func (*SiteInstanceDeployment) ElementType() reflect.Type {
	return reflect.TypeOf((*SiteInstanceDeployment)(nil))
}

func (i *SiteInstanceDeployment) ToSiteInstanceDeploymentOutput() SiteInstanceDeploymentOutput {
	return i.ToSiteInstanceDeploymentOutputWithContext(context.Background())
}

func (i *SiteInstanceDeployment) ToSiteInstanceDeploymentOutputWithContext(ctx context.Context) SiteInstanceDeploymentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SiteInstanceDeploymentOutput)
}

type SiteInstanceDeploymentOutput struct{ *pulumi.OutputState }

func (SiteInstanceDeploymentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SiteInstanceDeployment)(nil))
}

func (o SiteInstanceDeploymentOutput) ToSiteInstanceDeploymentOutput() SiteInstanceDeploymentOutput {
	return o
}

func (o SiteInstanceDeploymentOutput) ToSiteInstanceDeploymentOutputWithContext(ctx context.Context) SiteInstanceDeploymentOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(SiteInstanceDeploymentOutput{})
}
