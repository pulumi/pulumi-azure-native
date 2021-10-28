


package v20190501preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type LiveOutput struct {
	pulumi.CustomResourceState

	ArchiveWindowLength pulumi.StringOutput     `pulumi:"archiveWindowLength"`
	AssetName           pulumi.StringOutput     `pulumi:"assetName"`
	Created             pulumi.StringOutput     `pulumi:"created"`
	Description         pulumi.StringPtrOutput  `pulumi:"description"`
	Hls                 HlsResponsePtrOutput    `pulumi:"hls"`
	LastModified        pulumi.StringOutput     `pulumi:"lastModified"`
	ManifestName        pulumi.StringPtrOutput  `pulumi:"manifestName"`
	Name                pulumi.StringOutput     `pulumi:"name"`
	OutputSnapTime      pulumi.Float64PtrOutput `pulumi:"outputSnapTime"`
	ProvisioningState   pulumi.StringOutput     `pulumi:"provisioningState"`
	ResourceState       pulumi.StringOutput     `pulumi:"resourceState"`
	Type                pulumi.StringOutput     `pulumi:"type"`
}


func NewLiveOutput(ctx *pulumi.Context,
	name string, args *LiveOutputArgs, opts ...pulumi.ResourceOption) (*LiveOutput, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccountName == nil {
		return nil, errors.New("invalid value for required argument 'AccountName'")
	}
	if args.ArchiveWindowLength == nil {
		return nil, errors.New("invalid value for required argument 'ArchiveWindowLength'")
	}
	if args.AssetName == nil {
		return nil, errors.New("invalid value for required argument 'AssetName'")
	}
	if args.LiveEventName == nil {
		return nil, errors.New("invalid value for required argument 'LiveEventName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:media/v20190501preview:LiveOutput"),
		},
		{
			Type: pulumi.String("azure-native:media:LiveOutput"),
		},
		{
			Type: pulumi.String("azure-nextgen:media:LiveOutput"),
		},
		{
			Type: pulumi.String("azure-native:media/v20180330preview:LiveOutput"),
		},
		{
			Type: pulumi.String("azure-nextgen:media/v20180330preview:LiveOutput"),
		},
		{
			Type: pulumi.String("azure-native:media/v20180601preview:LiveOutput"),
		},
		{
			Type: pulumi.String("azure-nextgen:media/v20180601preview:LiveOutput"),
		},
		{
			Type: pulumi.String("azure-native:media/v20180701:LiveOutput"),
		},
		{
			Type: pulumi.String("azure-nextgen:media/v20180701:LiveOutput"),
		},
		{
			Type: pulumi.String("azure-native:media/v20200501:LiveOutput"),
		},
		{
			Type: pulumi.String("azure-nextgen:media/v20200501:LiveOutput"),
		},
		{
			Type: pulumi.String("azure-native:media/v20210601:LiveOutput"),
		},
		{
			Type: pulumi.String("azure-nextgen:media/v20210601:LiveOutput"),
		},
	})
	opts = append(opts, aliases)
	var resource LiveOutput
	err := ctx.RegisterResource("azure-native:media/v20190501preview:LiveOutput", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetLiveOutput(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LiveOutputState, opts ...pulumi.ResourceOption) (*LiveOutput, error) {
	var resource LiveOutput
	err := ctx.ReadResource("azure-native:media/v20190501preview:LiveOutput", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type liveOutputState struct {
}

type LiveOutputState struct {
}

func (LiveOutputState) ElementType() reflect.Type {
	return reflect.TypeOf((*liveOutputState)(nil)).Elem()
}

type liveOutputArgs struct {
	AccountName         string   `pulumi:"accountName"`
	ArchiveWindowLength string   `pulumi:"archiveWindowLength"`
	AssetName           string   `pulumi:"assetName"`
	Description         *string  `pulumi:"description"`
	Hls                 *Hls     `pulumi:"hls"`
	LiveEventName       string   `pulumi:"liveEventName"`
	LiveOutputName      *string  `pulumi:"liveOutputName"`
	ManifestName        *string  `pulumi:"manifestName"`
	OutputSnapTime      *float64 `pulumi:"outputSnapTime"`
	ResourceGroupName   string   `pulumi:"resourceGroupName"`
}


type LiveOutputArgs struct {
	AccountName         pulumi.StringInput
	ArchiveWindowLength pulumi.StringInput
	AssetName           pulumi.StringInput
	Description         pulumi.StringPtrInput
	Hls                 HlsPtrInput
	LiveEventName       pulumi.StringInput
	LiveOutputName      pulumi.StringPtrInput
	ManifestName        pulumi.StringPtrInput
	OutputSnapTime      pulumi.Float64PtrInput
	ResourceGroupName   pulumi.StringInput
}

func (LiveOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*liveOutputArgs)(nil)).Elem()
}

type LiveOutputInput interface {
	pulumi.Input

	ToLiveOutputOutput() LiveOutputOutput
	ToLiveOutputOutputWithContext(ctx context.Context) LiveOutputOutput
}

func (*LiveOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LiveOutput)(nil))
}

func (i *LiveOutput) ToLiveOutputOutput() LiveOutputOutput {
	return i.ToLiveOutputOutputWithContext(context.Background())
}

func (i *LiveOutput) ToLiveOutputOutputWithContext(ctx context.Context) LiveOutputOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LiveOutputOutput)
}

type LiveOutputOutput struct{ *pulumi.OutputState }

func (LiveOutputOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LiveOutput)(nil))
}

func (o LiveOutputOutput) ToLiveOutputOutput() LiveOutputOutput {
	return o
}

func (o LiveOutputOutput) ToLiveOutputOutputWithContext(ctx context.Context) LiveOutputOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*LiveOutputInput)(nil)).Elem(), &LiveOutput{})
	pulumi.RegisterOutputType(LiveOutputOutput{})
}
