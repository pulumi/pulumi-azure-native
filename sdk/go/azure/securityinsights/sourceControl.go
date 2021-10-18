


package securityinsights

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type SourceControl struct {
	pulumi.CustomResourceState

	ContentTypes       pulumi.StringArrayOutput `pulumi:"contentTypes"`
	CreatedAt          pulumi.StringPtrOutput   `pulumi:"createdAt"`
	CreatedBy          pulumi.StringPtrOutput   `pulumi:"createdBy"`
	CreatedByType      pulumi.StringPtrOutput   `pulumi:"createdByType"`
	Description        pulumi.StringPtrOutput   `pulumi:"description"`
	DisplayName        pulumi.StringOutput      `pulumi:"displayName"`
	Etag               pulumi.StringPtrOutput   `pulumi:"etag"`
	LastModifiedAt     pulumi.StringPtrOutput   `pulumi:"lastModifiedAt"`
	LastModifiedBy     pulumi.StringPtrOutput   `pulumi:"lastModifiedBy"`
	LastModifiedByType pulumi.StringPtrOutput   `pulumi:"lastModifiedByType"`
	Name               pulumi.StringOutput      `pulumi:"name"`
	RepoType           pulumi.StringOutput      `pulumi:"repoType"`
	Repository         RepositoryResponseOutput `pulumi:"repository"`
	SystemData         SystemDataResponseOutput `pulumi:"systemData"`
	Type               pulumi.StringOutput      `pulumi:"type"`
}


func NewSourceControl(ctx *pulumi.Context,
	name string, args *SourceControlArgs, opts ...pulumi.ResourceOption) (*SourceControl, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ContentTypes == nil {
		return nil, errors.New("invalid value for required argument 'ContentTypes'")
	}
	if args.DisplayName == nil {
		return nil, errors.New("invalid value for required argument 'DisplayName'")
	}
	if args.OperationalInsightsResourceProvider == nil {
		return nil, errors.New("invalid value for required argument 'OperationalInsightsResourceProvider'")
	}
	if args.RepoType == nil {
		return nil, errors.New("invalid value for required argument 'RepoType'")
	}
	if args.Repository == nil {
		return nil, errors.New("invalid value for required argument 'Repository'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.WorkspaceName == nil {
		return nil, errors.New("invalid value for required argument 'WorkspaceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:securityinsights:SourceControl"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20210301preview:SourceControl"),
		},
		{
			Type: pulumi.String("azure-nextgen:securityinsights/v20210301preview:SourceControl"),
		},
	})
	opts = append(opts, aliases)
	var resource SourceControl
	err := ctx.RegisterResource("azure-native:securityinsights:SourceControl", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetSourceControl(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SourceControlState, opts ...pulumi.ResourceOption) (*SourceControl, error) {
	var resource SourceControl
	err := ctx.ReadResource("azure-native:securityinsights:SourceControl", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type sourceControlState struct {
}

type SourceControlState struct {
}

func (SourceControlState) ElementType() reflect.Type {
	return reflect.TypeOf((*sourceControlState)(nil)).Elem()
}

type sourceControlArgs struct {
	ContentTypes                        []string   `pulumi:"contentTypes"`
	CreatedAt                           *string    `pulumi:"createdAt"`
	CreatedBy                           *string    `pulumi:"createdBy"`
	CreatedByType                       *string    `pulumi:"createdByType"`
	Description                         *string    `pulumi:"description"`
	DisplayName                         string     `pulumi:"displayName"`
	Etag                                *string    `pulumi:"etag"`
	Id                                  *string    `pulumi:"id"`
	LastModifiedAt                      *string    `pulumi:"lastModifiedAt"`
	LastModifiedBy                      *string    `pulumi:"lastModifiedBy"`
	LastModifiedByType                  *string    `pulumi:"lastModifiedByType"`
	OperationalInsightsResourceProvider string     `pulumi:"operationalInsightsResourceProvider"`
	RepoType                            string     `pulumi:"repoType"`
	Repository                          Repository `pulumi:"repository"`
	ResourceGroupName                   string     `pulumi:"resourceGroupName"`
	SourceControlId                     *string    `pulumi:"sourceControlId"`
	WorkspaceName                       string     `pulumi:"workspaceName"`
}


type SourceControlArgs struct {
	ContentTypes                        pulumi.StringArrayInput
	CreatedAt                           pulumi.StringPtrInput
	CreatedBy                           pulumi.StringPtrInput
	CreatedByType                       pulumi.StringPtrInput
	Description                         pulumi.StringPtrInput
	DisplayName                         pulumi.StringInput
	Etag                                pulumi.StringPtrInput
	Id                                  pulumi.StringPtrInput
	LastModifiedAt                      pulumi.StringPtrInput
	LastModifiedBy                      pulumi.StringPtrInput
	LastModifiedByType                  pulumi.StringPtrInput
	OperationalInsightsResourceProvider pulumi.StringInput
	RepoType                            pulumi.StringInput
	Repository                          RepositoryInput
	ResourceGroupName                   pulumi.StringInput
	SourceControlId                     pulumi.StringPtrInput
	WorkspaceName                       pulumi.StringInput
}

func (SourceControlArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sourceControlArgs)(nil)).Elem()
}

type SourceControlInput interface {
	pulumi.Input

	ToSourceControlOutput() SourceControlOutput
	ToSourceControlOutputWithContext(ctx context.Context) SourceControlOutput
}

func (*SourceControl) ElementType() reflect.Type {
	return reflect.TypeOf((*SourceControl)(nil))
}

func (i *SourceControl) ToSourceControlOutput() SourceControlOutput {
	return i.ToSourceControlOutputWithContext(context.Background())
}

func (i *SourceControl) ToSourceControlOutputWithContext(ctx context.Context) SourceControlOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceControlOutput)
}

type SourceControlOutput struct{ *pulumi.OutputState }

func (SourceControlOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SourceControl)(nil))
}

func (o SourceControlOutput) ToSourceControlOutput() SourceControlOutput {
	return o
}

func (o SourceControlOutput) ToSourceControlOutputWithContext(ctx context.Context) SourceControlOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(SourceControlOutput{})
}
