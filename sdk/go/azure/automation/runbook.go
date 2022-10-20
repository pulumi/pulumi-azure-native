


package automation

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
			Type: pulumi.String("azure-native:automation/v20151031:Runbook"),
		},
		{
			Type: pulumi.String("azure-native:automation/v20180630:Runbook"),
		},
		{
			Type: pulumi.String("azure-native:automation/v20190601:Runbook"),
		},
	})
	opts = append(opts, aliases)
	var resource Runbook
	err := ctx.RegisterResource("azure-native:automation:Runbook", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetRunbook(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RunbookState, opts ...pulumi.ResourceOption) (*Runbook, error) {
	var resource Runbook
	err := ctx.ReadResource("azure-native:automation:Runbook", name, id, state, &resource, opts...)
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
	return reflect.TypeOf((**Runbook)(nil)).Elem()
}

func (i *Runbook) ToRunbookOutput() RunbookOutput {
	return i.ToRunbookOutputWithContext(context.Background())
}

func (i *Runbook) ToRunbookOutputWithContext(ctx context.Context) RunbookOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RunbookOutput)
}

type RunbookOutput struct{ *pulumi.OutputState }

func (RunbookOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Runbook)(nil)).Elem()
}

func (o RunbookOutput) ToRunbookOutput() RunbookOutput {
	return o
}

func (o RunbookOutput) ToRunbookOutputWithContext(ctx context.Context) RunbookOutput {
	return o
}

func (o RunbookOutput) CreationTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Runbook) pulumi.StringPtrOutput { return v.CreationTime }).(pulumi.StringPtrOutput)
}

func (o RunbookOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Runbook) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

func (o RunbookOutput) Draft() RunbookDraftResponsePtrOutput {
	return o.ApplyT(func(v *Runbook) RunbookDraftResponsePtrOutput { return v.Draft }).(RunbookDraftResponsePtrOutput)
}

func (o RunbookOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Runbook) pulumi.StringPtrOutput { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o RunbookOutput) JobCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Runbook) pulumi.IntPtrOutput { return v.JobCount }).(pulumi.IntPtrOutput)
}

func (o RunbookOutput) LastModifiedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Runbook) pulumi.StringPtrOutput { return v.LastModifiedBy }).(pulumi.StringPtrOutput)
}

func (o RunbookOutput) LastModifiedTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Runbook) pulumi.StringPtrOutput { return v.LastModifiedTime }).(pulumi.StringPtrOutput)
}

func (o RunbookOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Runbook) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

func (o RunbookOutput) LogActivityTrace() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Runbook) pulumi.IntPtrOutput { return v.LogActivityTrace }).(pulumi.IntPtrOutput)
}

func (o RunbookOutput) LogProgress() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Runbook) pulumi.BoolPtrOutput { return v.LogProgress }).(pulumi.BoolPtrOutput)
}

func (o RunbookOutput) LogVerbose() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Runbook) pulumi.BoolPtrOutput { return v.LogVerbose }).(pulumi.BoolPtrOutput)
}

func (o RunbookOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Runbook) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o RunbookOutput) OutputTypes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Runbook) pulumi.StringArrayOutput { return v.OutputTypes }).(pulumi.StringArrayOutput)
}

func (o RunbookOutput) Parameters() RunbookParameterResponseMapOutput {
	return o.ApplyT(func(v *Runbook) RunbookParameterResponseMapOutput { return v.Parameters }).(RunbookParameterResponseMapOutput)
}

func (o RunbookOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Runbook) pulumi.StringPtrOutput { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

func (o RunbookOutput) PublishContentLink() ContentLinkResponsePtrOutput {
	return o.ApplyT(func(v *Runbook) ContentLinkResponsePtrOutput { return v.PublishContentLink }).(ContentLinkResponsePtrOutput)
}

func (o RunbookOutput) RunbookType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Runbook) pulumi.StringPtrOutput { return v.RunbookType }).(pulumi.StringPtrOutput)
}

func (o RunbookOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Runbook) pulumi.StringPtrOutput { return v.State }).(pulumi.StringPtrOutput)
}

func (o RunbookOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Runbook) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o RunbookOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Runbook) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(RunbookOutput{})
}
