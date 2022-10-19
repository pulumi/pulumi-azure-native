


package automation

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Watcher struct {
	pulumi.CustomResourceState

	CreationTime                pulumi.StringOutput     `pulumi:"creationTime"`
	Description                 pulumi.StringPtrOutput  `pulumi:"description"`
	Etag                        pulumi.StringPtrOutput  `pulumi:"etag"`
	ExecutionFrequencyInSeconds pulumi.Float64PtrOutput `pulumi:"executionFrequencyInSeconds"`
	LastModifiedBy              pulumi.StringOutput     `pulumi:"lastModifiedBy"`
	LastModifiedTime            pulumi.StringOutput     `pulumi:"lastModifiedTime"`
	Location                    pulumi.StringPtrOutput  `pulumi:"location"`
	Name                        pulumi.StringOutput     `pulumi:"name"`
	ScriptName                  pulumi.StringPtrOutput  `pulumi:"scriptName"`
	ScriptParameters            pulumi.StringMapOutput  `pulumi:"scriptParameters"`
	ScriptRunOn                 pulumi.StringPtrOutput  `pulumi:"scriptRunOn"`
	Status                      pulumi.StringOutput     `pulumi:"status"`
	Tags                        pulumi.StringMapOutput  `pulumi:"tags"`
	Type                        pulumi.StringOutput     `pulumi:"type"`
}


func NewWatcher(ctx *pulumi.Context,
	name string, args *WatcherArgs, opts ...pulumi.ResourceOption) (*Watcher, error) {
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
			Type: pulumi.String("azure-native:automation/v20151031:Watcher"),
		},
		{
			Type: pulumi.String("azure-native:automation/v20190601:Watcher"),
		},
		{
			Type: pulumi.String("azure-native:automation/v20200113preview:Watcher"),
		},
	})
	opts = append(opts, aliases)
	var resource Watcher
	err := ctx.RegisterResource("azure-native:automation:Watcher", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetWatcher(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WatcherState, opts ...pulumi.ResourceOption) (*Watcher, error) {
	var resource Watcher
	err := ctx.ReadResource("azure-native:automation:Watcher", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type watcherState struct {
}

type WatcherState struct {
}

func (WatcherState) ElementType() reflect.Type {
	return reflect.TypeOf((*watcherState)(nil)).Elem()
}

type watcherArgs struct {
	AutomationAccountName       string            `pulumi:"automationAccountName"`
	Description                 *string           `pulumi:"description"`
	ExecutionFrequencyInSeconds *float64          `pulumi:"executionFrequencyInSeconds"`
	Location                    *string           `pulumi:"location"`
	ResourceGroupName           string            `pulumi:"resourceGroupName"`
	ScriptName                  *string           `pulumi:"scriptName"`
	ScriptParameters            map[string]string `pulumi:"scriptParameters"`
	ScriptRunOn                 *string           `pulumi:"scriptRunOn"`
	Tags                        map[string]string `pulumi:"tags"`
	WatcherName                 *string           `pulumi:"watcherName"`
}


type WatcherArgs struct {
	AutomationAccountName       pulumi.StringInput
	Description                 pulumi.StringPtrInput
	ExecutionFrequencyInSeconds pulumi.Float64PtrInput
	Location                    pulumi.StringPtrInput
	ResourceGroupName           pulumi.StringInput
	ScriptName                  pulumi.StringPtrInput
	ScriptParameters            pulumi.StringMapInput
	ScriptRunOn                 pulumi.StringPtrInput
	Tags                        pulumi.StringMapInput
	WatcherName                 pulumi.StringPtrInput
}

func (WatcherArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*watcherArgs)(nil)).Elem()
}

type WatcherInput interface {
	pulumi.Input

	ToWatcherOutput() WatcherOutput
	ToWatcherOutputWithContext(ctx context.Context) WatcherOutput
}

func (*Watcher) ElementType() reflect.Type {
	return reflect.TypeOf((**Watcher)(nil)).Elem()
}

func (i *Watcher) ToWatcherOutput() WatcherOutput {
	return i.ToWatcherOutputWithContext(context.Background())
}

func (i *Watcher) ToWatcherOutputWithContext(ctx context.Context) WatcherOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WatcherOutput)
}

type WatcherOutput struct{ *pulumi.OutputState }

func (WatcherOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Watcher)(nil)).Elem()
}

func (o WatcherOutput) ToWatcherOutput() WatcherOutput {
	return o
}

func (o WatcherOutput) ToWatcherOutputWithContext(ctx context.Context) WatcherOutput {
	return o
}

func (o WatcherOutput) CreationTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Watcher) pulumi.StringOutput { return v.CreationTime }).(pulumi.StringOutput)
}

func (o WatcherOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Watcher) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

func (o WatcherOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Watcher) pulumi.StringPtrOutput { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o WatcherOutput) ExecutionFrequencyInSeconds() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *Watcher) pulumi.Float64PtrOutput { return v.ExecutionFrequencyInSeconds }).(pulumi.Float64PtrOutput)
}

func (o WatcherOutput) LastModifiedBy() pulumi.StringOutput {
	return o.ApplyT(func(v *Watcher) pulumi.StringOutput { return v.LastModifiedBy }).(pulumi.StringOutput)
}

func (o WatcherOutput) LastModifiedTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Watcher) pulumi.StringOutput { return v.LastModifiedTime }).(pulumi.StringOutput)
}

func (o WatcherOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Watcher) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

func (o WatcherOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Watcher) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o WatcherOutput) ScriptName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Watcher) pulumi.StringPtrOutput { return v.ScriptName }).(pulumi.StringPtrOutput)
}

func (o WatcherOutput) ScriptParameters() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Watcher) pulumi.StringMapOutput { return v.ScriptParameters }).(pulumi.StringMapOutput)
}

func (o WatcherOutput) ScriptRunOn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Watcher) pulumi.StringPtrOutput { return v.ScriptRunOn }).(pulumi.StringPtrOutput)
}

func (o WatcherOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *Watcher) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

func (o WatcherOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Watcher) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o WatcherOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Watcher) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(WatcherOutput{})
}
