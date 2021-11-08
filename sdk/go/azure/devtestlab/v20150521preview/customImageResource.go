


package v20150521preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type CustomImageResource struct {
	pulumi.CustomResourceState

	Author            pulumi.StringPtrOutput                       `pulumi:"author"`
	CreationDate      pulumi.StringPtrOutput                       `pulumi:"creationDate"`
	Description       pulumi.StringPtrOutput                       `pulumi:"description"`
	Location          pulumi.StringPtrOutput                       `pulumi:"location"`
	Name              pulumi.StringPtrOutput                       `pulumi:"name"`
	OsType            pulumi.StringPtrOutput                       `pulumi:"osType"`
	ProvisioningState pulumi.StringPtrOutput                       `pulumi:"provisioningState"`
	Tags              pulumi.StringMapOutput                       `pulumi:"tags"`
	Type              pulumi.StringPtrOutput                       `pulumi:"type"`
	Vhd               CustomImagePropertiesCustomResponsePtrOutput `pulumi:"vhd"`
	Vm                CustomImagePropertiesFromVmResponsePtrOutput `pulumi:"vm"`
}


func NewCustomImageResource(ctx *pulumi.Context,
	name string, args *CustomImageResourceArgs, opts ...pulumi.ResourceOption) (*CustomImageResource, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.LabName == nil {
		return nil, errors.New("invalid value for required argument 'LabName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:devtestlab/v20150521preview:CustomImageResource"),
		},
		{
			Type: pulumi.String("azure-native:devtestlab:CustomImageResource"),
		},
		{
			Type: pulumi.String("azure-nextgen:devtestlab:CustomImageResource"),
		},
		{
			Type: pulumi.String("azure-native:devtestlab/v20160515:CustomImageResource"),
		},
		{
			Type: pulumi.String("azure-nextgen:devtestlab/v20160515:CustomImageResource"),
		},
		{
			Type: pulumi.String("azure-native:devtestlab/v20180915:CustomImageResource"),
		},
		{
			Type: pulumi.String("azure-nextgen:devtestlab/v20180915:CustomImageResource"),
		},
	})
	opts = append(opts, aliases)
	var resource CustomImageResource
	err := ctx.RegisterResource("azure-native:devtestlab/v20150521preview:CustomImageResource", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetCustomImageResource(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CustomImageResourceState, opts ...pulumi.ResourceOption) (*CustomImageResource, error) {
	var resource CustomImageResource
	err := ctx.ReadResource("azure-native:devtestlab/v20150521preview:CustomImageResource", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type customImageResourceState struct {
}

type CustomImageResourceState struct {
}

func (CustomImageResourceState) ElementType() reflect.Type {
	return reflect.TypeOf((*customImageResourceState)(nil)).Elem()
}

type customImageResourceArgs struct {
	Author            *string                      `pulumi:"author"`
	CreationDate      *string                      `pulumi:"creationDate"`
	Description       *string                      `pulumi:"description"`
	Id                *string                      `pulumi:"id"`
	LabName           string                       `pulumi:"labName"`
	Location          *string                      `pulumi:"location"`
	Name              *string                      `pulumi:"name"`
	OsType            *string                      `pulumi:"osType"`
	ProvisioningState *string                      `pulumi:"provisioningState"`
	ResourceGroupName string                       `pulumi:"resourceGroupName"`
	Tags              map[string]string            `pulumi:"tags"`
	Type              *string                      `pulumi:"type"`
	Vhd               *CustomImagePropertiesCustom `pulumi:"vhd"`
	Vm                *CustomImagePropertiesFromVm `pulumi:"vm"`
}


type CustomImageResourceArgs struct {
	Author            pulumi.StringPtrInput
	CreationDate      pulumi.StringPtrInput
	Description       pulumi.StringPtrInput
	Id                pulumi.StringPtrInput
	LabName           pulumi.StringInput
	Location          pulumi.StringPtrInput
	Name              pulumi.StringPtrInput
	OsType            pulumi.StringPtrInput
	ProvisioningState pulumi.StringPtrInput
	ResourceGroupName pulumi.StringInput
	Tags              pulumi.StringMapInput
	Type              pulumi.StringPtrInput
	Vhd               CustomImagePropertiesCustomPtrInput
	Vm                CustomImagePropertiesFromVmPtrInput
}

func (CustomImageResourceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*customImageResourceArgs)(nil)).Elem()
}

type CustomImageResourceInput interface {
	pulumi.Input

	ToCustomImageResourceOutput() CustomImageResourceOutput
	ToCustomImageResourceOutputWithContext(ctx context.Context) CustomImageResourceOutput
}

func (*CustomImageResource) ElementType() reflect.Type {
	return reflect.TypeOf((*CustomImageResource)(nil))
}

func (i *CustomImageResource) ToCustomImageResourceOutput() CustomImageResourceOutput {
	return i.ToCustomImageResourceOutputWithContext(context.Background())
}

func (i *CustomImageResource) ToCustomImageResourceOutputWithContext(ctx context.Context) CustomImageResourceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomImageResourceOutput)
}

type CustomImageResourceOutput struct{ *pulumi.OutputState }

func (CustomImageResourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CustomImageResource)(nil))
}

func (o CustomImageResourceOutput) ToCustomImageResourceOutput() CustomImageResourceOutput {
	return o
}

func (o CustomImageResourceOutput) ToCustomImageResourceOutputWithContext(ctx context.Context) CustomImageResourceOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(CustomImageResourceOutput{})
}
