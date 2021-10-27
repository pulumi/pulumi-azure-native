


package v20180630

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
			Type: pulumi.String("azure-nextgen:automation/v20180630:Python2Package"),
		},
		{
			Type: pulumi.String("azure-native:automation:Python2Package"),
		},
		{
			Type: pulumi.String("azure-nextgen:automation:Python2Package"),
		},
		{
			Type: pulumi.String("azure-native:automation/v20190601:Python2Package"),
		},
		{
			Type: pulumi.String("azure-nextgen:automation/v20190601:Python2Package"),
		},
		{
			Type: pulumi.String("azure-native:automation/v20200113preview:Python2Package"),
		},
		{
			Type: pulumi.String("azure-nextgen:automation/v20200113preview:Python2Package"),
		},
	})
	opts = append(opts, aliases)
	var resource Python2Package
	err := ctx.RegisterResource("azure-native:automation/v20180630:Python2Package", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetPython2Package(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *Python2PackageState, opts ...pulumi.ResourceOption) (*Python2Package, error) {
	var resource Python2Package
	err := ctx.ReadResource("azure-native:automation/v20180630:Python2Package", name, id, state, &resource, opts...)
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
	return reflect.TypeOf((*Python2Package)(nil))
}

func (i *Python2Package) ToPython2PackageOutput() Python2PackageOutput {
	return i.ToPython2PackageOutputWithContext(context.Background())
}

func (i *Python2Package) ToPython2PackageOutputWithContext(ctx context.Context) Python2PackageOutput {
	return pulumi.ToOutputWithContext(ctx, i).(Python2PackageOutput)
}

type Python2PackageOutput struct{ *pulumi.OutputState }

func (Python2PackageOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Python2Package)(nil))
}

func (o Python2PackageOutput) ToPython2PackageOutput() Python2PackageOutput {
	return o
}

func (o Python2PackageOutput) ToPython2PackageOutputWithContext(ctx context.Context) Python2PackageOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*Python2PackageInput)(nil)).Elem(), &Python2Package{})
	pulumi.RegisterOutputType(Python2PackageOutput{})
}
