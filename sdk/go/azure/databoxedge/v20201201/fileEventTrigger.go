


package v20201201

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type FileEventTrigger struct {
	pulumi.CustomResourceState

	CustomContextTag pulumi.StringPtrOutput       `pulumi:"customContextTag"`
	Kind             pulumi.StringOutput          `pulumi:"kind"`
	Name             pulumi.StringOutput          `pulumi:"name"`
	SinkInfo         RoleSinkInfoResponseOutput   `pulumi:"sinkInfo"`
	SourceInfo       FileSourceInfoResponseOutput `pulumi:"sourceInfo"`
	SystemData       SystemDataResponseOutput     `pulumi:"systemData"`
	Type             pulumi.StringOutput          `pulumi:"type"`
}


func NewFileEventTrigger(ctx *pulumi.Context,
	name string, args *FileEventTriggerArgs, opts ...pulumi.ResourceOption) (*FileEventTrigger, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DeviceName == nil {
		return nil, errors.New("invalid value for required argument 'DeviceName'")
	}
	if args.Kind == nil {
		return nil, errors.New("invalid value for required argument 'Kind'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.SinkInfo == nil {
		return nil, errors.New("invalid value for required argument 'SinkInfo'")
	}
	if args.SourceInfo == nil {
		return nil, errors.New("invalid value for required argument 'SourceInfo'")
	}
	args.Kind = pulumi.String("FileEvent")
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:databoxedge:FileEventTrigger"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20190301:FileEventTrigger"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20190701:FileEventTrigger"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20190801:FileEventTrigger"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20200501preview:FileEventTrigger"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20200901:FileEventTrigger"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20200901preview:FileEventTrigger"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20210201:FileEventTrigger"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20210201preview:FileEventTrigger"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20210601:FileEventTrigger"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20210601preview:FileEventTrigger"),
		},
	})
	opts = append(opts, aliases)
	var resource FileEventTrigger
	err := ctx.RegisterResource("azure-native:databoxedge/v20201201:FileEventTrigger", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetFileEventTrigger(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FileEventTriggerState, opts ...pulumi.ResourceOption) (*FileEventTrigger, error) {
	var resource FileEventTrigger
	err := ctx.ReadResource("azure-native:databoxedge/v20201201:FileEventTrigger", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type fileEventTriggerState struct {
}

type FileEventTriggerState struct {
}

func (FileEventTriggerState) ElementType() reflect.Type {
	return reflect.TypeOf((*fileEventTriggerState)(nil)).Elem()
}

type fileEventTriggerArgs struct {
	CustomContextTag  *string        `pulumi:"customContextTag"`
	DeviceName        string         `pulumi:"deviceName"`
	Kind              string         `pulumi:"kind"`
	Name              *string        `pulumi:"name"`
	ResourceGroupName string         `pulumi:"resourceGroupName"`
	SinkInfo          RoleSinkInfo   `pulumi:"sinkInfo"`
	SourceInfo        FileSourceInfo `pulumi:"sourceInfo"`
}


type FileEventTriggerArgs struct {
	CustomContextTag  pulumi.StringPtrInput
	DeviceName        pulumi.StringInput
	Kind              pulumi.StringInput
	Name              pulumi.StringPtrInput
	ResourceGroupName pulumi.StringInput
	SinkInfo          RoleSinkInfoInput
	SourceInfo        FileSourceInfoInput
}

func (FileEventTriggerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*fileEventTriggerArgs)(nil)).Elem()
}

type FileEventTriggerInput interface {
	pulumi.Input

	ToFileEventTriggerOutput() FileEventTriggerOutput
	ToFileEventTriggerOutputWithContext(ctx context.Context) FileEventTriggerOutput
}

func (*FileEventTrigger) ElementType() reflect.Type {
	return reflect.TypeOf((*FileEventTrigger)(nil))
}

func (i *FileEventTrigger) ToFileEventTriggerOutput() FileEventTriggerOutput {
	return i.ToFileEventTriggerOutputWithContext(context.Background())
}

func (i *FileEventTrigger) ToFileEventTriggerOutputWithContext(ctx context.Context) FileEventTriggerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FileEventTriggerOutput)
}

type FileEventTriggerOutput struct{ *pulumi.OutputState }

func (FileEventTriggerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FileEventTrigger)(nil))
}

func (o FileEventTriggerOutput) ToFileEventTriggerOutput() FileEventTriggerOutput {
	return o
}

func (o FileEventTriggerOutput) ToFileEventTriggerOutputWithContext(ctx context.Context) FileEventTriggerOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(FileEventTriggerOutput{})
}
