


package v20150801

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type SiteSourceControl struct {
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


func NewSiteSourceControl(ctx *pulumi.Context,
	name string, args *SiteSourceControlArgs, opts ...pulumi.ResourceOption) (*SiteSourceControl, error) {
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
			Type: pulumi.String("azure-nextgen:web/v20150801:SiteSourceControl"),
		},
		{
			Type: pulumi.String("azure-native:web:SiteSourceControl"),
		},
		{
			Type: pulumi.String("azure-nextgen:web:SiteSourceControl"),
		},
		{
			Type: pulumi.String("azure-native:web/v20160801:SiteSourceControl"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20160801:SiteSourceControl"),
		},
		{
			Type: pulumi.String("azure-native:web/v20180201:SiteSourceControl"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20180201:SiteSourceControl"),
		},
		{
			Type: pulumi.String("azure-native:web/v20181101:SiteSourceControl"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20181101:SiteSourceControl"),
		},
		{
			Type: pulumi.String("azure-native:web/v20190801:SiteSourceControl"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20190801:SiteSourceControl"),
		},
		{
			Type: pulumi.String("azure-native:web/v20200601:SiteSourceControl"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20200601:SiteSourceControl"),
		},
		{
			Type: pulumi.String("azure-native:web/v20200901:SiteSourceControl"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20200901:SiteSourceControl"),
		},
		{
			Type: pulumi.String("azure-native:web/v20201001:SiteSourceControl"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20201001:SiteSourceControl"),
		},
		{
			Type: pulumi.String("azure-native:web/v20201201:SiteSourceControl"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20201201:SiteSourceControl"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210101:SiteSourceControl"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20210101:SiteSourceControl"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210115:SiteSourceControl"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20210115:SiteSourceControl"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210201:SiteSourceControl"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20210201:SiteSourceControl"),
		},
	})
	opts = append(opts, aliases)
	var resource SiteSourceControl
	err := ctx.RegisterResource("azure-native:web/v20150801:SiteSourceControl", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetSiteSourceControl(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SiteSourceControlState, opts ...pulumi.ResourceOption) (*SiteSourceControl, error) {
	var resource SiteSourceControl
	err := ctx.ReadResource("azure-native:web/v20150801:SiteSourceControl", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type siteSourceControlState struct {
}

type SiteSourceControlState struct {
}

func (SiteSourceControlState) ElementType() reflect.Type {
	return reflect.TypeOf((*siteSourceControlState)(nil)).Elem()
}

type siteSourceControlArgs struct {
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
	Tags                      map[string]string `pulumi:"tags"`
	Type                      *string           `pulumi:"type"`
}


type SiteSourceControlArgs struct {
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
	Tags                      pulumi.StringMapInput
	Type                      pulumi.StringPtrInput
}

func (SiteSourceControlArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*siteSourceControlArgs)(nil)).Elem()
}

type SiteSourceControlInput interface {
	pulumi.Input

	ToSiteSourceControlOutput() SiteSourceControlOutput
	ToSiteSourceControlOutputWithContext(ctx context.Context) SiteSourceControlOutput
}

func (*SiteSourceControl) ElementType() reflect.Type {
	return reflect.TypeOf((*SiteSourceControl)(nil))
}

func (i *SiteSourceControl) ToSiteSourceControlOutput() SiteSourceControlOutput {
	return i.ToSiteSourceControlOutputWithContext(context.Background())
}

func (i *SiteSourceControl) ToSiteSourceControlOutputWithContext(ctx context.Context) SiteSourceControlOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SiteSourceControlOutput)
}

type SiteSourceControlOutput struct{ *pulumi.OutputState }

func (SiteSourceControlOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SiteSourceControl)(nil))
}

func (o SiteSourceControlOutput) ToSiteSourceControlOutput() SiteSourceControlOutput {
	return o
}

func (o SiteSourceControlOutput) ToSiteSourceControlOutputWithContext(ctx context.Context) SiteSourceControlOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SiteSourceControlInput)(nil)).Elem(), &SiteSourceControl{})
	pulumi.RegisterOutputType(SiteSourceControlOutput{})
}
