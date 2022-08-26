


package v20200113preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Module struct {
	pulumi.CustomResourceState

	ActivityCount     pulumi.IntPtrOutput              `pulumi:"activityCount"`
	ContentLink       ContentLinkResponsePtrOutput     `pulumi:"contentLink"`
	CreationTime      pulumi.StringPtrOutput           `pulumi:"creationTime"`
	Description       pulumi.StringPtrOutput           `pulumi:"description"`
	Error             ModuleErrorInfoResponsePtrOutput `pulumi:"error"`
	Etag              pulumi.StringPtrOutput           `pulumi:"etag"`
	IsComposite       pulumi.BoolPtrOutput             `pulumi:"isComposite"`
	IsGlobal          pulumi.BoolPtrOutput             `pulumi:"isGlobal"`
	LastModifiedTime  pulumi.StringPtrOutput           `pulumi:"lastModifiedTime"`
	Location          pulumi.StringPtrOutput           `pulumi:"location"`
	Name              pulumi.StringOutput              `pulumi:"name"`
	ProvisioningState pulumi.StringPtrOutput           `pulumi:"provisioningState"`
	SizeInBytes       pulumi.Float64PtrOutput          `pulumi:"sizeInBytes"`
	Tags              pulumi.StringMapOutput           `pulumi:"tags"`
	Type              pulumi.StringOutput              `pulumi:"type"`
	Version           pulumi.StringPtrOutput           `pulumi:"version"`
}


func NewModule(ctx *pulumi.Context,
	name string, args *ModuleArgs, opts ...pulumi.ResourceOption) (*Module, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AutomationAccountName == nil {
		return nil, errors.New("invalid value for required argument 'AutomationAccountName'")
	}
	if args.ContentLink == nil {
		return nil, errors.New("invalid value for required argument 'ContentLink'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:automation:Module"),
		},
		{
			Type: pulumi.String("azure-native:automation/v20151031:Module"),
		},
		{
			Type: pulumi.String("azure-native:automation/v20190601:Module"),
		},
	})
	opts = append(opts, aliases)
	var resource Module
	err := ctx.RegisterResource("azure-native:automation/v20200113preview:Module", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetModule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ModuleState, opts ...pulumi.ResourceOption) (*Module, error) {
	var resource Module
	err := ctx.ReadResource("azure-native:automation/v20200113preview:Module", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type moduleState struct {
}

type ModuleState struct {
}

func (ModuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*moduleState)(nil)).Elem()
}

type moduleArgs struct {
	AutomationAccountName string            `pulumi:"automationAccountName"`
	ContentLink           ContentLink       `pulumi:"contentLink"`
	Location              *string           `pulumi:"location"`
	ModuleName            *string           `pulumi:"moduleName"`
	Name                  *string           `pulumi:"name"`
	ResourceGroupName     string            `pulumi:"resourceGroupName"`
	Tags                  map[string]string `pulumi:"tags"`
}


type ModuleArgs struct {
	AutomationAccountName pulumi.StringInput
	ContentLink           ContentLinkInput
	Location              pulumi.StringPtrInput
	ModuleName            pulumi.StringPtrInput
	Name                  pulumi.StringPtrInput
	ResourceGroupName     pulumi.StringInput
	Tags                  pulumi.StringMapInput
}

func (ModuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*moduleArgs)(nil)).Elem()
}

type ModuleInput interface {
	pulumi.Input

	ToModuleOutput() ModuleOutput
	ToModuleOutputWithContext(ctx context.Context) ModuleOutput
}

func (*Module) ElementType() reflect.Type {
	return reflect.TypeOf((**Module)(nil)).Elem()
}

func (i *Module) ToModuleOutput() ModuleOutput {
	return i.ToModuleOutputWithContext(context.Background())
}

func (i *Module) ToModuleOutputWithContext(ctx context.Context) ModuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ModuleOutput)
}

type ModuleOutput struct{ *pulumi.OutputState }

func (ModuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Module)(nil)).Elem()
}

func (o ModuleOutput) ToModuleOutput() ModuleOutput {
	return o
}

func (o ModuleOutput) ToModuleOutputWithContext(ctx context.Context) ModuleOutput {
	return o
}

func (o ModuleOutput) ActivityCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Module) pulumi.IntPtrOutput { return v.ActivityCount }).(pulumi.IntPtrOutput)
}

func (o ModuleOutput) ContentLink() ContentLinkResponsePtrOutput {
	return o.ApplyT(func(v *Module) ContentLinkResponsePtrOutput { return v.ContentLink }).(ContentLinkResponsePtrOutput)
}

func (o ModuleOutput) CreationTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Module) pulumi.StringPtrOutput { return v.CreationTime }).(pulumi.StringPtrOutput)
}

func (o ModuleOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Module) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

func (o ModuleOutput) Error() ModuleErrorInfoResponsePtrOutput {
	return o.ApplyT(func(v *Module) ModuleErrorInfoResponsePtrOutput { return v.Error }).(ModuleErrorInfoResponsePtrOutput)
}

func (o ModuleOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Module) pulumi.StringPtrOutput { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o ModuleOutput) IsComposite() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Module) pulumi.BoolPtrOutput { return v.IsComposite }).(pulumi.BoolPtrOutput)
}

func (o ModuleOutput) IsGlobal() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Module) pulumi.BoolPtrOutput { return v.IsGlobal }).(pulumi.BoolPtrOutput)
}

func (o ModuleOutput) LastModifiedTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Module) pulumi.StringPtrOutput { return v.LastModifiedTime }).(pulumi.StringPtrOutput)
}

func (o ModuleOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Module) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

func (o ModuleOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Module) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ModuleOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Module) pulumi.StringPtrOutput { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

func (o ModuleOutput) SizeInBytes() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *Module) pulumi.Float64PtrOutput { return v.SizeInBytes }).(pulumi.Float64PtrOutput)
}

func (o ModuleOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Module) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o ModuleOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Module) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o ModuleOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Module) pulumi.StringPtrOutput { return v.Version }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ModuleOutput{})
}
