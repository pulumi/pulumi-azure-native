


package v20150801

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


type SiteDeployment struct {
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


func NewSiteDeployment(ctx *pulumi.Context,
	name string, args *SiteDeploymentArgs, opts ...pulumi.ResourceOption) (*SiteDeployment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Name == nil {
		return nil, errors.New("invalid value for required argument 'Name'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:web:SiteDeployment"),
		},
		{
			Type: pulumi.String("azure-native:web/v20160801:SiteDeployment"),
		},
		{
			Type: pulumi.String("azure-native:web/v20180201:SiteDeployment"),
		},
		{
			Type: pulumi.String("azure-native:web/v20181101:SiteDeployment"),
		},
		{
			Type: pulumi.String("azure-native:web/v20190801:SiteDeployment"),
		},
		{
			Type: pulumi.String("azure-native:web/v20200601:SiteDeployment"),
		},
		{
			Type: pulumi.String("azure-native:web/v20200901:SiteDeployment"),
		},
		{
			Type: pulumi.String("azure-native:web/v20201001:SiteDeployment"),
		},
		{
			Type: pulumi.String("azure-native:web/v20201201:SiteDeployment"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210101:SiteDeployment"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210115:SiteDeployment"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210201:SiteDeployment"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210301:SiteDeployment"),
		},
		{
			Type: pulumi.String("azure-native:web/v20220301:SiteDeployment"),
		},
	})
	opts = append(opts, aliases)
	var resource SiteDeployment
	err := ctx.RegisterResource("azure-native:web/v20150801:SiteDeployment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetSiteDeployment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SiteDeploymentState, opts ...pulumi.ResourceOption) (*SiteDeployment, error) {
	var resource SiteDeployment
	err := ctx.ReadResource("azure-native:web/v20150801:SiteDeployment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type siteDeploymentState struct {
}

type SiteDeploymentState struct {
}

func (SiteDeploymentState) ElementType() reflect.Type {
	return reflect.TypeOf((*siteDeploymentState)(nil)).Elem()
}

type siteDeploymentArgs struct {
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
	StartTime         *string           `pulumi:"startTime"`
	Status            *int              `pulumi:"status"`
	Tags              map[string]string `pulumi:"tags"`
	Type              *string           `pulumi:"type"`
}


type SiteDeploymentArgs struct {
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
	StartTime         pulumi.StringPtrInput
	Status            pulumi.IntPtrInput
	Tags              pulumi.StringMapInput
	Type              pulumi.StringPtrInput
}

func (SiteDeploymentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*siteDeploymentArgs)(nil)).Elem()
}

type SiteDeploymentInput interface {
	pulumi.Input

	ToSiteDeploymentOutput() SiteDeploymentOutput
	ToSiteDeploymentOutputWithContext(ctx context.Context) SiteDeploymentOutput
}

func (*SiteDeployment) ElementType() reflect.Type {
	return reflect.TypeOf((**SiteDeployment)(nil)).Elem()
}

func (i *SiteDeployment) ToSiteDeploymentOutput() SiteDeploymentOutput {
	return i.ToSiteDeploymentOutputWithContext(context.Background())
}

func (i *SiteDeployment) ToSiteDeploymentOutputWithContext(ctx context.Context) SiteDeploymentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SiteDeploymentOutput)
}

type SiteDeploymentOutput struct{ *pulumi.OutputState }

func (SiteDeploymentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SiteDeployment)(nil)).Elem()
}

func (o SiteDeploymentOutput) ToSiteDeploymentOutput() SiteDeploymentOutput {
	return o
}

func (o SiteDeploymentOutput) ToSiteDeploymentOutputWithContext(ctx context.Context) SiteDeploymentOutput {
	return o
}

func (o SiteDeploymentOutput) Active() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SiteDeployment) pulumi.BoolPtrOutput { return v.Active }).(pulumi.BoolPtrOutput)
}

func (o SiteDeploymentOutput) Author() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteDeployment) pulumi.StringPtrOutput { return v.Author }).(pulumi.StringPtrOutput)
}

func (o SiteDeploymentOutput) AuthorEmail() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteDeployment) pulumi.StringPtrOutput { return v.AuthorEmail }).(pulumi.StringPtrOutput)
}

func (o SiteDeploymentOutput) Deployer() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteDeployment) pulumi.StringPtrOutput { return v.Deployer }).(pulumi.StringPtrOutput)
}

func (o SiteDeploymentOutput) Details() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteDeployment) pulumi.StringPtrOutput { return v.Details }).(pulumi.StringPtrOutput)
}

func (o SiteDeploymentOutput) EndTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteDeployment) pulumi.StringPtrOutput { return v.EndTime }).(pulumi.StringPtrOutput)
}

func (o SiteDeploymentOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteDeployment) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o SiteDeploymentOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *SiteDeployment) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o SiteDeploymentOutput) Message() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteDeployment) pulumi.StringPtrOutput { return v.Message }).(pulumi.StringPtrOutput)
}

func (o SiteDeploymentOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteDeployment) pulumi.StringPtrOutput { return v.Name }).(pulumi.StringPtrOutput)
}

func (o SiteDeploymentOutput) StartTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteDeployment) pulumi.StringPtrOutput { return v.StartTime }).(pulumi.StringPtrOutput)
}

func (o SiteDeploymentOutput) Status() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SiteDeployment) pulumi.IntPtrOutput { return v.Status }).(pulumi.IntPtrOutput)
}

func (o SiteDeploymentOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *SiteDeployment) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o SiteDeploymentOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteDeployment) pulumi.StringPtrOutput { return v.Type }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(SiteDeploymentOutput{})
}
