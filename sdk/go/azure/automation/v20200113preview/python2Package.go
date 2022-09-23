


package v20200113preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Python2Package struct {
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


func NewPython2Package(ctx *pulumi.Context,
	name string, args *Python2PackageArgs, opts ...pulumi.ResourceOption) (*Python2Package, error) {
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
			Type: pulumi.String("azure-native:automation:Python2Package"),
		},
		{
			Type: pulumi.String("azure-native:automation/v20180630:Python2Package"),
		},
		{
			Type: pulumi.String("azure-native:automation/v20190601:Python2Package"),
		},
		{
			Type: pulumi.String("azure-native:automation/v20220808:Python2Package"),
		},
	})
	opts = append(opts, aliases)
	var resource Python2Package
	err := ctx.RegisterResource("azure-native:automation/v20200113preview:Python2Package", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetPython2Package(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *Python2PackageState, opts ...pulumi.ResourceOption) (*Python2Package, error) {
	var resource Python2Package
	err := ctx.ReadResource("azure-native:automation/v20200113preview:Python2Package", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type python2PackageState struct {
}

type Python2PackageState struct {
}

func (Python2PackageState) ElementType() reflect.Type {
	return reflect.TypeOf((*python2PackageState)(nil)).Elem()
}

type python2PackageArgs struct {
	AutomationAccountName string            `pulumi:"automationAccountName"`
	ContentLink           ContentLink       `pulumi:"contentLink"`
	PackageName           *string           `pulumi:"packageName"`
	ResourceGroupName     string            `pulumi:"resourceGroupName"`
	Tags                  map[string]string `pulumi:"tags"`
}


type Python2PackageArgs struct {
	AutomationAccountName pulumi.StringInput
	ContentLink           ContentLinkInput
	PackageName           pulumi.StringPtrInput
	ResourceGroupName     pulumi.StringInput
	Tags                  pulumi.StringMapInput
}

func (Python2PackageArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*python2PackageArgs)(nil)).Elem()
}

type Python2PackageInput interface {
	pulumi.Input

	ToPython2PackageOutput() Python2PackageOutput
	ToPython2PackageOutputWithContext(ctx context.Context) Python2PackageOutput
}

func (*Python2Package) ElementType() reflect.Type {
	return reflect.TypeOf((**Python2Package)(nil)).Elem()
}

func (i *Python2Package) ToPython2PackageOutput() Python2PackageOutput {
	return i.ToPython2PackageOutputWithContext(context.Background())
}

func (i *Python2Package) ToPython2PackageOutputWithContext(ctx context.Context) Python2PackageOutput {
	return pulumi.ToOutputWithContext(ctx, i).(Python2PackageOutput)
}

type Python2PackageOutput struct{ *pulumi.OutputState }

func (Python2PackageOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Python2Package)(nil)).Elem()
}

func (o Python2PackageOutput) ToPython2PackageOutput() Python2PackageOutput {
	return o
}

func (o Python2PackageOutput) ToPython2PackageOutputWithContext(ctx context.Context) Python2PackageOutput {
	return o
}

func (o Python2PackageOutput) ActivityCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Python2Package) pulumi.IntPtrOutput { return v.ActivityCount }).(pulumi.IntPtrOutput)
}

func (o Python2PackageOutput) ContentLink() ContentLinkResponsePtrOutput {
	return o.ApplyT(func(v *Python2Package) ContentLinkResponsePtrOutput { return v.ContentLink }).(ContentLinkResponsePtrOutput)
}

func (o Python2PackageOutput) CreationTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Python2Package) pulumi.StringPtrOutput { return v.CreationTime }).(pulumi.StringPtrOutput)
}

func (o Python2PackageOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Python2Package) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

func (o Python2PackageOutput) Error() ModuleErrorInfoResponsePtrOutput {
	return o.ApplyT(func(v *Python2Package) ModuleErrorInfoResponsePtrOutput { return v.Error }).(ModuleErrorInfoResponsePtrOutput)
}

func (o Python2PackageOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Python2Package) pulumi.StringPtrOutput { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o Python2PackageOutput) IsComposite() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Python2Package) pulumi.BoolPtrOutput { return v.IsComposite }).(pulumi.BoolPtrOutput)
}

func (o Python2PackageOutput) IsGlobal() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Python2Package) pulumi.BoolPtrOutput { return v.IsGlobal }).(pulumi.BoolPtrOutput)
}

func (o Python2PackageOutput) LastModifiedTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Python2Package) pulumi.StringPtrOutput { return v.LastModifiedTime }).(pulumi.StringPtrOutput)
}

func (o Python2PackageOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Python2Package) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

func (o Python2PackageOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Python2Package) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o Python2PackageOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Python2Package) pulumi.StringPtrOutput { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

func (o Python2PackageOutput) SizeInBytes() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *Python2Package) pulumi.Float64PtrOutput { return v.SizeInBytes }).(pulumi.Float64PtrOutput)
}

func (o Python2PackageOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Python2Package) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o Python2PackageOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Python2Package) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o Python2PackageOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Python2Package) pulumi.StringPtrOutput { return v.Version }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(Python2PackageOutput{})
}
