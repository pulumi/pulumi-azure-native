


package v20150801

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type SiteSourceControlSlot struct {
	pulumi.CustomResourceState

	Branch                    pulumi.StringPtrOutput `pulumi:"branch"`
	DeploymentRollbackEnabled pulumi.BoolPtrOutput   `pulumi:"deploymentRollbackEnabled"`
	IsManualIntegration       pulumi.BoolPtrOutput   `pulumi:"isManualIntegration"`
	IsMercurial               pulumi.BoolPtrOutput   `pulumi:"isMercurial"`
	Kind                      pulumi.StringPtrOutput `pulumi:"kind"`
	Location                  pulumi.StringOutput    `pulumi:"location"`
	Name                      pulumi.StringPtrOutput `pulumi:"name"`
	RepoUrl                   pulumi.StringPtrOutput `pulumi:"repoUrl"`
	Tags                      pulumi.StringMapOutput `pulumi:"tags"`
	Type                      pulumi.StringPtrOutput `pulumi:"type"`
}


func NewSiteSourceControlSlot(ctx *pulumi.Context,
	name string, args *SiteSourceControlSlotArgs, opts ...pulumi.ResourceOption) (*SiteSourceControlSlot, error) {
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
			Type: pulumi.String("azure-native:web:SiteSourceControlSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20160801:SiteSourceControlSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20180201:SiteSourceControlSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20181101:SiteSourceControlSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20190801:SiteSourceControlSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20200601:SiteSourceControlSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20200901:SiteSourceControlSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20201001:SiteSourceControlSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20201201:SiteSourceControlSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210101:SiteSourceControlSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210115:SiteSourceControlSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210201:SiteSourceControlSlot"),
		},
	})
	opts = append(opts, aliases)
	var resource SiteSourceControlSlot
	err := ctx.RegisterResource("azure-native:web/v20150801:SiteSourceControlSlot", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetSiteSourceControlSlot(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SiteSourceControlSlotState, opts ...pulumi.ResourceOption) (*SiteSourceControlSlot, error) {
	var resource SiteSourceControlSlot
	err := ctx.ReadResource("azure-native:web/v20150801:SiteSourceControlSlot", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type siteSourceControlSlotState struct {
}

type SiteSourceControlSlotState struct {
}

func (SiteSourceControlSlotState) ElementType() reflect.Type {
	return reflect.TypeOf((*siteSourceControlSlotState)(nil)).Elem()
}

type siteSourceControlSlotArgs struct {
	Branch                    *string           `pulumi:"branch"`
	DeploymentRollbackEnabled *bool             `pulumi:"deploymentRollbackEnabled"`
	Id                        *string           `pulumi:"id"`
	IsManualIntegration       *bool             `pulumi:"isManualIntegration"`
	IsMercurial               *bool             `pulumi:"isMercurial"`
	Kind                      *string           `pulumi:"kind"`
	Location                  *string           `pulumi:"location"`
	Name                      string            `pulumi:"name"`
	RepoUrl                   *string           `pulumi:"repoUrl"`
	ResourceGroupName         string            `pulumi:"resourceGroupName"`
	Slot                      string            `pulumi:"slot"`
	Tags                      map[string]string `pulumi:"tags"`
	Type                      *string           `pulumi:"type"`
}


type SiteSourceControlSlotArgs struct {
	Branch                    pulumi.StringPtrInput
	DeploymentRollbackEnabled pulumi.BoolPtrInput
	Id                        pulumi.StringPtrInput
	IsManualIntegration       pulumi.BoolPtrInput
	IsMercurial               pulumi.BoolPtrInput
	Kind                      pulumi.StringPtrInput
	Location                  pulumi.StringPtrInput
	Name                      pulumi.StringInput
	RepoUrl                   pulumi.StringPtrInput
	ResourceGroupName         pulumi.StringInput
	Slot                      pulumi.StringInput
	Tags                      pulumi.StringMapInput
	Type                      pulumi.StringPtrInput
}

func (SiteSourceControlSlotArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*siteSourceControlSlotArgs)(nil)).Elem()
}

type SiteSourceControlSlotInput interface {
	pulumi.Input

	ToSiteSourceControlSlotOutput() SiteSourceControlSlotOutput
	ToSiteSourceControlSlotOutputWithContext(ctx context.Context) SiteSourceControlSlotOutput
}

func (*SiteSourceControlSlot) ElementType() reflect.Type {
	return reflect.TypeOf((**SiteSourceControlSlot)(nil)).Elem()
}

func (i *SiteSourceControlSlot) ToSiteSourceControlSlotOutput() SiteSourceControlSlotOutput {
	return i.ToSiteSourceControlSlotOutputWithContext(context.Background())
}

func (i *SiteSourceControlSlot) ToSiteSourceControlSlotOutputWithContext(ctx context.Context) SiteSourceControlSlotOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SiteSourceControlSlotOutput)
}

type SiteSourceControlSlotOutput struct{ *pulumi.OutputState }

func (SiteSourceControlSlotOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SiteSourceControlSlot)(nil)).Elem()
}

func (o SiteSourceControlSlotOutput) ToSiteSourceControlSlotOutput() SiteSourceControlSlotOutput {
	return o
}

func (o SiteSourceControlSlotOutput) ToSiteSourceControlSlotOutputWithContext(ctx context.Context) SiteSourceControlSlotOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(SiteSourceControlSlotOutput{})
}
