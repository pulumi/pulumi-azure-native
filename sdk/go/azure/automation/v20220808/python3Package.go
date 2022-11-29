


package v20220808

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Python3Package struct {
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


func NewPython3Package(ctx *pulumi.Context,
	name string, args *Python3PackageArgs, opts ...pulumi.ResourceOption) (*Python3Package, error) {
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
	var resource Python3Package
	err := ctx.RegisterResource("azure-native:automation/v20220808:Python3Package", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetPython3Package(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *Python3PackageState, opts ...pulumi.ResourceOption) (*Python3Package, error) {
	var resource Python3Package
	err := ctx.ReadResource("azure-native:automation/v20220808:Python3Package", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type python3PackageState struct {
}

type Python3PackageState struct {
}

func (Python3PackageState) ElementType() reflect.Type {
	return reflect.TypeOf((*python3PackageState)(nil)).Elem()
}

type python3PackageArgs struct {
	AutomationAccountName string            `pulumi:"automationAccountName"`
	ContentLink           ContentLink       `pulumi:"contentLink"`
	PackageName           *string           `pulumi:"packageName"`
	ResourceGroupName     string            `pulumi:"resourceGroupName"`
	Tags                  map[string]string `pulumi:"tags"`
}


type Python3PackageArgs struct {
	AutomationAccountName pulumi.StringInput
	ContentLink           ContentLinkInput
	PackageName           pulumi.StringPtrInput
	ResourceGroupName     pulumi.StringInput
	Tags                  pulumi.StringMapInput
}

func (Python3PackageArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*python3PackageArgs)(nil)).Elem()
}

type Python3PackageInput interface {
	pulumi.Input

	ToPython3PackageOutput() Python3PackageOutput
	ToPython3PackageOutputWithContext(ctx context.Context) Python3PackageOutput
}

func (*Python3Package) ElementType() reflect.Type {
	return reflect.TypeOf((**Python3Package)(nil)).Elem()
}

func (i *Python3Package) ToPython3PackageOutput() Python3PackageOutput {
	return i.ToPython3PackageOutputWithContext(context.Background())
}

func (i *Python3Package) ToPython3PackageOutputWithContext(ctx context.Context) Python3PackageOutput {
	return pulumi.ToOutputWithContext(ctx, i).(Python3PackageOutput)
}

type Python3PackageOutput struct{ *pulumi.OutputState }

func (Python3PackageOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Python3Package)(nil)).Elem()
}

func (o Python3PackageOutput) ToPython3PackageOutput() Python3PackageOutput {
	return o
}

func (o Python3PackageOutput) ToPython3PackageOutputWithContext(ctx context.Context) Python3PackageOutput {
	return o
}

func (o Python3PackageOutput) ActivityCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Python3Package) pulumi.IntPtrOutput { return v.ActivityCount }).(pulumi.IntPtrOutput)
}

func (o Python3PackageOutput) ContentLink() ContentLinkResponsePtrOutput {
	return o.ApplyT(func(v *Python3Package) ContentLinkResponsePtrOutput { return v.ContentLink }).(ContentLinkResponsePtrOutput)
}

func (o Python3PackageOutput) CreationTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Python3Package) pulumi.StringPtrOutput { return v.CreationTime }).(pulumi.StringPtrOutput)
}

func (o Python3PackageOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Python3Package) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

func (o Python3PackageOutput) Error() ModuleErrorInfoResponsePtrOutput {
	return o.ApplyT(func(v *Python3Package) ModuleErrorInfoResponsePtrOutput { return v.Error }).(ModuleErrorInfoResponsePtrOutput)
}

func (o Python3PackageOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Python3Package) pulumi.StringPtrOutput { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o Python3PackageOutput) IsComposite() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Python3Package) pulumi.BoolPtrOutput { return v.IsComposite }).(pulumi.BoolPtrOutput)
}

func (o Python3PackageOutput) IsGlobal() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Python3Package) pulumi.BoolPtrOutput { return v.IsGlobal }).(pulumi.BoolPtrOutput)
}

func (o Python3PackageOutput) LastModifiedTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Python3Package) pulumi.StringPtrOutput { return v.LastModifiedTime }).(pulumi.StringPtrOutput)
}

func (o Python3PackageOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Python3Package) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

func (o Python3PackageOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Python3Package) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o Python3PackageOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Python3Package) pulumi.StringPtrOutput { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

func (o Python3PackageOutput) SizeInBytes() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *Python3Package) pulumi.Float64PtrOutput { return v.SizeInBytes }).(pulumi.Float64PtrOutput)
}

func (o Python3PackageOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Python3Package) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o Python3PackageOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Python3Package) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o Python3PackageOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Python3Package) pulumi.StringPtrOutput { return v.Version }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(Python3PackageOutput{})
}
