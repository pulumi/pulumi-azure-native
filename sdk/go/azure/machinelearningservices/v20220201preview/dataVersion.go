


package v20220201preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type DataVersion struct {
	pulumi.CustomResourceState

	DataVersionBaseDetails pulumi.AnyOutput         `pulumi:"dataVersionBaseDetails"`
	Name                   pulumi.StringOutput      `pulumi:"name"`
	SystemData             SystemDataResponseOutput `pulumi:"systemData"`
	Type                   pulumi.StringOutput      `pulumi:"type"`
}


func NewDataVersion(ctx *pulumi.Context,
	name string, args *DataVersionArgs, opts ...pulumi.ResourceOption) (*DataVersion, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DataVersionBaseDetails == nil {
		return nil, errors.New("invalid value for required argument 'DataVersionBaseDetails'")
	}
	if args.Name == nil {
		return nil, errors.New("invalid value for required argument 'Name'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.WorkspaceName == nil {
		return nil, errors.New("invalid value for required argument 'WorkspaceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:machinelearningservices:DataVersion"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20210301preview:DataVersion"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20220501:DataVersion"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20220601preview:DataVersion"),
		},
	})
	opts = append(opts, aliases)
	var resource DataVersion
	err := ctx.RegisterResource("azure-native:machinelearningservices/v20220201preview:DataVersion", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetDataVersion(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DataVersionState, opts ...pulumi.ResourceOption) (*DataVersion, error) {
	var resource DataVersion
	err := ctx.ReadResource("azure-native:machinelearningservices/v20220201preview:DataVersion", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type dataVersionState struct {
}

type DataVersionState struct {
}

func (DataVersionState) ElementType() reflect.Type {
	return reflect.TypeOf((*dataVersionState)(nil)).Elem()
}

type dataVersionArgs struct {
	DataVersionBaseDetails interface{} `pulumi:"dataVersionBaseDetails"`
	Name                   string      `pulumi:"name"`
	ResourceGroupName      string      `pulumi:"resourceGroupName"`
	Version                *string     `pulumi:"version"`
	WorkspaceName          string      `pulumi:"workspaceName"`
}


type DataVersionArgs struct {
	DataVersionBaseDetails pulumi.Input
	Name                   pulumi.StringInput
	ResourceGroupName      pulumi.StringInput
	Version                pulumi.StringPtrInput
	WorkspaceName          pulumi.StringInput
}

func (DataVersionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dataVersionArgs)(nil)).Elem()
}

type DataVersionInput interface {
	pulumi.Input

	ToDataVersionOutput() DataVersionOutput
	ToDataVersionOutputWithContext(ctx context.Context) DataVersionOutput
}

func (*DataVersion) ElementType() reflect.Type {
	return reflect.TypeOf((**DataVersion)(nil)).Elem()
}

func (i *DataVersion) ToDataVersionOutput() DataVersionOutput {
	return i.ToDataVersionOutputWithContext(context.Background())
}

func (i *DataVersion) ToDataVersionOutputWithContext(ctx context.Context) DataVersionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataVersionOutput)
}

type DataVersionOutput struct{ *pulumi.OutputState }

func (DataVersionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DataVersion)(nil)).Elem()
}

func (o DataVersionOutput) ToDataVersionOutput() DataVersionOutput {
	return o
}

func (o DataVersionOutput) ToDataVersionOutputWithContext(ctx context.Context) DataVersionOutput {
	return o
}

func (o DataVersionOutput) DataVersionBaseDetails() pulumi.AnyOutput {
	return o.ApplyT(func(v *DataVersion) pulumi.AnyOutput { return v.DataVersionBaseDetails }).(pulumi.AnyOutput)
}

func (o DataVersionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *DataVersion) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o DataVersionOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *DataVersion) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o DataVersionOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *DataVersion) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(DataVersionOutput{})
}
