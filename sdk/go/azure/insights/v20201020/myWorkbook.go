


package v20201020

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type MyWorkbook struct {
	pulumi.CustomResourceState

	Category       pulumi.StringOutput                        `pulumi:"category"`
	DisplayName    pulumi.StringOutput                        `pulumi:"displayName"`
	Etag           pulumi.StringMapOutput                     `pulumi:"etag"`
	Identity       MyWorkbookManagedIdentityResponsePtrOutput `pulumi:"identity"`
	Kind           pulumi.StringPtrOutput                     `pulumi:"kind"`
	Location       pulumi.StringPtrOutput                     `pulumi:"location"`
	Name           pulumi.StringPtrOutput                     `pulumi:"name"`
	SerializedData pulumi.StringOutput                        `pulumi:"serializedData"`
	SourceId       pulumi.StringPtrOutput                     `pulumi:"sourceId"`
	StorageUri     pulumi.StringPtrOutput                     `pulumi:"storageUri"`
	Tags           pulumi.StringMapOutput                     `pulumi:"tags"`
	TimeModified   pulumi.StringOutput                        `pulumi:"timeModified"`
	Type           pulumi.StringPtrOutput                     `pulumi:"type"`
	UserId         pulumi.StringOutput                        `pulumi:"userId"`
	Version        pulumi.StringPtrOutput                     `pulumi:"version"`
}


func NewMyWorkbook(ctx *pulumi.Context,
	name string, args *MyWorkbookArgs, opts ...pulumi.ResourceOption) (*MyWorkbook, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Category == nil {
		return nil, errors.New("invalid value for required argument 'Category'")
	}
	if args.DisplayName == nil {
		return nil, errors.New("invalid value for required argument 'DisplayName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.SerializedData == nil {
		return nil, errors.New("invalid value for required argument 'SerializedData'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:insights:MyWorkbook"),
		},
		{
			Type: pulumi.String("azure-native:insights/v20150501:MyWorkbook"),
		},
		{
			Type: pulumi.String("azure-native:insights/v20210308:MyWorkbook"),
		},
	})
	opts = append(opts, aliases)
	var resource MyWorkbook
	err := ctx.RegisterResource("azure-native:insights/v20201020:MyWorkbook", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetMyWorkbook(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MyWorkbookState, opts ...pulumi.ResourceOption) (*MyWorkbook, error) {
	var resource MyWorkbook
	err := ctx.ReadResource("azure-native:insights/v20201020:MyWorkbook", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type myWorkbookState struct {
}

type MyWorkbookState struct {
}

func (MyWorkbookState) ElementType() reflect.Type {
	return reflect.TypeOf((*myWorkbookState)(nil)).Elem()
}

type myWorkbookArgs struct {
	Category          string                     `pulumi:"category"`
	DisplayName       string                     `pulumi:"displayName"`
	Etag              map[string]string          `pulumi:"etag"`
	Id                *string                    `pulumi:"id"`
	Identity          *MyWorkbookManagedIdentity `pulumi:"identity"`
	Kind              *string                    `pulumi:"kind"`
	Location          *string                    `pulumi:"location"`
	Name              *string                    `pulumi:"name"`
	ResourceGroupName string                     `pulumi:"resourceGroupName"`
	ResourceName      *string                    `pulumi:"resourceName"`
	SerializedData    string                     `pulumi:"serializedData"`
	SourceId          *string                    `pulumi:"sourceId"`
	StorageUri        *string                    `pulumi:"storageUri"`
	Tags              map[string]string          `pulumi:"tags"`
	Type              *string                    `pulumi:"type"`
	Version           *string                    `pulumi:"version"`
}


type MyWorkbookArgs struct {
	Category          pulumi.StringInput
	DisplayName       pulumi.StringInput
	Etag              pulumi.StringMapInput
	Id                pulumi.StringPtrInput
	Identity          MyWorkbookManagedIdentityPtrInput
	Kind              pulumi.StringPtrInput
	Location          pulumi.StringPtrInput
	Name              pulumi.StringPtrInput
	ResourceGroupName pulumi.StringInput
	ResourceName      pulumi.StringPtrInput
	SerializedData    pulumi.StringInput
	SourceId          pulumi.StringPtrInput
	StorageUri        pulumi.StringPtrInput
	Tags              pulumi.StringMapInput
	Type              pulumi.StringPtrInput
	Version           pulumi.StringPtrInput
}

func (MyWorkbookArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*myWorkbookArgs)(nil)).Elem()
}

type MyWorkbookInput interface {
	pulumi.Input

	ToMyWorkbookOutput() MyWorkbookOutput
	ToMyWorkbookOutputWithContext(ctx context.Context) MyWorkbookOutput
}

func (*MyWorkbook) ElementType() reflect.Type {
	return reflect.TypeOf((*MyWorkbook)(nil))
}

func (i *MyWorkbook) ToMyWorkbookOutput() MyWorkbookOutput {
	return i.ToMyWorkbookOutputWithContext(context.Background())
}

func (i *MyWorkbook) ToMyWorkbookOutputWithContext(ctx context.Context) MyWorkbookOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MyWorkbookOutput)
}

type MyWorkbookOutput struct{ *pulumi.OutputState }

func (MyWorkbookOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MyWorkbook)(nil))
}

func (o MyWorkbookOutput) ToMyWorkbookOutput() MyWorkbookOutput {
	return o
}

func (o MyWorkbookOutput) ToMyWorkbookOutputWithContext(ctx context.Context) MyWorkbookOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(MyWorkbookOutput{})
}
