


package automation

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type SourceControl struct {
	pulumi.CustomResourceState

	AutoSync         pulumi.BoolPtrOutput   `pulumi:"autoSync"`
	Branch           pulumi.StringPtrOutput `pulumi:"branch"`
	CreationTime     pulumi.StringPtrOutput `pulumi:"creationTime"`
	Description      pulumi.StringPtrOutput `pulumi:"description"`
	FolderPath       pulumi.StringPtrOutput `pulumi:"folderPath"`
	LastModifiedTime pulumi.StringPtrOutput `pulumi:"lastModifiedTime"`
	Name             pulumi.StringOutput    `pulumi:"name"`
	PublishRunbook   pulumi.BoolPtrOutput   `pulumi:"publishRunbook"`
	RepoUrl          pulumi.StringPtrOutput `pulumi:"repoUrl"`
	SourceType       pulumi.StringPtrOutput `pulumi:"sourceType"`
	Type             pulumi.StringOutput    `pulumi:"type"`
}


func NewSourceControl(ctx *pulumi.Context,
	name string, args *SourceControlArgs, opts ...pulumi.ResourceOption) (*SourceControl, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AutomationAccountName == nil {
		return nil, errors.New("invalid value for required argument 'AutomationAccountName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:automation/v20170515preview:SourceControl"),
		},
		{
			Type: pulumi.String("azure-native:automation/v20190601:SourceControl"),
		},
		{
			Type: pulumi.String("azure-native:automation/v20200113preview:SourceControl"),
		},
	})
	opts = append(opts, aliases)
	var resource SourceControl
	err := ctx.RegisterResource("azure-native:automation:SourceControl", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetSourceControl(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SourceControlState, opts ...pulumi.ResourceOption) (*SourceControl, error) {
	var resource SourceControl
	err := ctx.ReadResource("azure-native:automation:SourceControl", name, id, state, &resource, opts...)
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
	AutoSync              *bool                                 `pulumi:"autoSync"`
	AutomationAccountName string                                `pulumi:"automationAccountName"`
	Branch                *string                               `pulumi:"branch"`
	Description           *string                               `pulumi:"description"`
	FolderPath            *string                               `pulumi:"folderPath"`
	PublishRunbook        *bool                                 `pulumi:"publishRunbook"`
	RepoUrl               *string                               `pulumi:"repoUrl"`
	ResourceGroupName     string                                `pulumi:"resourceGroupName"`
	SecurityToken         *SourceControlSecurityTokenProperties `pulumi:"securityToken"`
	SourceControlName     *string                               `pulumi:"sourceControlName"`
	SourceType            *string                               `pulumi:"sourceType"`
}


type SourceControlArgs struct {
	AutoSync              pulumi.BoolPtrInput
	AutomationAccountName pulumi.StringInput
	Branch                pulumi.StringPtrInput
	Description           pulumi.StringPtrInput
	FolderPath            pulumi.StringPtrInput
	PublishRunbook        pulumi.BoolPtrInput
	RepoUrl               pulumi.StringPtrInput
	ResourceGroupName     pulumi.StringInput
	SecurityToken         SourceControlSecurityTokenPropertiesPtrInput
	SourceControlName     pulumi.StringPtrInput
	SourceType            pulumi.StringPtrInput
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
	return reflect.TypeOf((**SourceControl)(nil)).Elem()
}

func (i *SourceControl) ToSourceControlOutput() SourceControlOutput {
	return i.ToSourceControlOutputWithContext(context.Background())
}

func (i *SourceControl) ToSourceControlOutputWithContext(ctx context.Context) SourceControlOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceControlOutput)
}

type SourceControlOutput struct{ *pulumi.OutputState }

func (SourceControlOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SourceControl)(nil)).Elem()
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
