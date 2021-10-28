


package v20180630

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Runbook struct {
	pulumi.CustomResourceState

	CreationTime       pulumi.StringPtrOutput            `pulumi:"creationTime"`
	Description        pulumi.StringPtrOutput            `pulumi:"description"`
	Draft              RunbookDraftResponsePtrOutput     `pulumi:"draft"`
	Etag               pulumi.StringPtrOutput            `pulumi:"etag"`
	JobCount           pulumi.IntPtrOutput               `pulumi:"jobCount"`
	LastModifiedBy     pulumi.StringPtrOutput            `pulumi:"lastModifiedBy"`
	LastModifiedTime   pulumi.StringPtrOutput            `pulumi:"lastModifiedTime"`
	Location           pulumi.StringPtrOutput            `pulumi:"location"`
	LogActivityTrace   pulumi.IntPtrOutput               `pulumi:"logActivityTrace"`
	LogProgress        pulumi.BoolPtrOutput              `pulumi:"logProgress"`
	LogVerbose         pulumi.BoolPtrOutput              `pulumi:"logVerbose"`
	Name               pulumi.StringOutput               `pulumi:"name"`
	OutputTypes        pulumi.StringArrayOutput          `pulumi:"outputTypes"`
	Parameters         RunbookParameterResponseMapOutput `pulumi:"parameters"`
	ProvisioningState  pulumi.StringPtrOutput            `pulumi:"provisioningState"`
	PublishContentLink ContentLinkResponsePtrOutput      `pulumi:"publishContentLink"`
	RunbookType        pulumi.StringPtrOutput            `pulumi:"runbookType"`
	State              pulumi.StringPtrOutput            `pulumi:"state"`
	Tags               pulumi.StringMapOutput            `pulumi:"tags"`
	Type               pulumi.StringOutput               `pulumi:"type"`
}


func NewRunbook(ctx *pulumi.Context,
	name string, args *RunbookArgs, opts ...pulumi.ResourceOption) (*Runbook, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AutomationAccountName == nil {
		return nil, errors.New("invalid value for required argument 'AutomationAccountName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.RunbookType == nil {
		return nil, errors.New("invalid value for required argument 'RunbookType'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:automation/v20180630:Runbook"),
		},
		{
			Type: pulumi.String("azure-native:automation:Runbook"),
		},
		{
			Type: pulumi.String("azure-nextgen:automation:Runbook"),
		},
		{
			Type: pulumi.String("azure-native:automation/v20151031:Runbook"),
		},
		{
			Type: pulumi.String("azure-nextgen:automation/v20151031:Runbook"),
		},
		{
			Type: pulumi.String("azure-native:automation/v20190601:Runbook"),
		},
		{
			Type: pulumi.String("azure-nextgen:automation/v20190601:Runbook"),
		},
	})
	opts = append(opts, aliases)
	var resource Runbook
	err := ctx.RegisterResource("azure-native:automation/v20180630:Runbook", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetRunbook(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RunbookState, opts ...pulumi.ResourceOption) (*Runbook, error) {
	var resource Runbook
	err := ctx.ReadResource("azure-native:automation/v20180630:Runbook", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type runbookState struct {
}

type RunbookState struct {
}

func (RunbookState) ElementType() reflect.Type {
	return reflect.TypeOf((*runbookState)(nil)).Elem()
}

type runbookArgs struct {
	AutomationAccountName string            `pulumi:"automationAccountName"`
	Description           *string           `pulumi:"description"`
	Draft                 *RunbookDraft     `pulumi:"draft"`
	Location              *string           `pulumi:"location"`
	LogActivityTrace      *int              `pulumi:"logActivityTrace"`
	LogProgress           *bool             `pulumi:"logProgress"`
	LogVerbose            *bool             `pulumi:"logVerbose"`
	Name                  *string           `pulumi:"name"`
	PublishContentLink    *ContentLink      `pulumi:"publishContentLink"`
	ResourceGroupName     string            `pulumi:"resourceGroupName"`
	RunbookName           *string           `pulumi:"runbookName"`
	RunbookType           string            `pulumi:"runbookType"`
	Tags                  map[string]string `pulumi:"tags"`
}


type RunbookArgs struct {
	AutomationAccountName pulumi.StringInput
	Description           pulumi.StringPtrInput
	Draft                 RunbookDraftPtrInput
	Location              pulumi.StringPtrInput
	LogActivityTrace      pulumi.IntPtrInput
	LogProgress           pulumi.BoolPtrInput
	LogVerbose            pulumi.BoolPtrInput
	Name                  pulumi.StringPtrInput
	PublishContentLink    ContentLinkPtrInput
	ResourceGroupName     pulumi.StringInput
	RunbookName           pulumi.StringPtrInput
	RunbookType           pulumi.StringInput
	Tags                  pulumi.StringMapInput
}

func (RunbookArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*runbookArgs)(nil)).Elem()
}

type RunbookInput interface {
	pulumi.Input

	ToRunbookOutput() RunbookOutput
	ToRunbookOutputWithContext(ctx context.Context) RunbookOutput
}

func (*Runbook) ElementType() reflect.Type {
	return reflect.TypeOf((*Runbook)(nil))
}

func (i *Runbook) ToRunbookOutput() RunbookOutput {
	return i.ToRunbookOutputWithContext(context.Background())
}

func (i *Runbook) ToRunbookOutputWithContext(ctx context.Context) RunbookOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RunbookOutput)
}

type RunbookOutput struct{ *pulumi.OutputState }

func (RunbookOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Runbook)(nil))
}

func (o RunbookOutput) ToRunbookOutput() RunbookOutput {
	return o
}

func (o RunbookOutput) ToRunbookOutputWithContext(ctx context.Context) RunbookOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RunbookInput)(nil)).Elem(), &Runbook{})
	pulumi.RegisterOutputType(RunbookOutput{})
}
