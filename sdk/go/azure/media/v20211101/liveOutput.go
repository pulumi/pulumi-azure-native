


package v20211101

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type LiveOutput struct {
	pulumi.CustomResourceState

	ArchiveWindowLength pulumi.StringOutput      `pulumi:"archiveWindowLength"`
	AssetName           pulumi.StringOutput      `pulumi:"assetName"`
	Created             pulumi.StringOutput      `pulumi:"created"`
	Description         pulumi.StringPtrOutput   `pulumi:"description"`
	Hls                 HlsResponsePtrOutput     `pulumi:"hls"`
	LastModified        pulumi.StringOutput      `pulumi:"lastModified"`
	ManifestName        pulumi.StringPtrOutput   `pulumi:"manifestName"`
	Name                pulumi.StringOutput      `pulumi:"name"`
	OutputSnapTime      pulumi.Float64PtrOutput  `pulumi:"outputSnapTime"`
	ProvisioningState   pulumi.StringOutput      `pulumi:"provisioningState"`
	ResourceState       pulumi.StringOutput      `pulumi:"resourceState"`
	SystemData          SystemDataResponseOutput `pulumi:"systemData"`
	Type                pulumi.StringOutput      `pulumi:"type"`
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
			Type: pulumi.String("azure-native:media:LiveOutput"),
		},
		{
			Type: pulumi.String("azure-native:media/v20180330preview:LiveOutput"),
		},
		{
			Type: pulumi.String("azure-native:media/v20180601preview:LiveOutput"),
		},
		{
			Type: pulumi.String("azure-native:media/v20180701:LiveOutput"),
		},
		{
			Type: pulumi.String("azure-native:media/v20190501preview:LiveOutput"),
		},
		{
			Type: pulumi.String("azure-native:media/v20200501:LiveOutput"),
		},
		{
			Type: pulumi.String("azure-native:media/v20210601:LiveOutput"),
		},
		{
			Type: pulumi.String("azure-native:media/v20220801:LiveOutput"),
		},
	})
	opts = append(opts, aliases)
	var resource LiveOutput
	err := ctx.RegisterResource("azure-native:media/v20211101:LiveOutput", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetLiveOutput(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LiveOutputState, opts ...pulumi.ResourceOption) (*LiveOutput, error) {
	var resource LiveOutput
	err := ctx.ReadResource("azure-native:media/v20211101:LiveOutput", name, id, state, &resource, opts...)
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
	return reflect.TypeOf((**LiveOutput)(nil)).Elem()
}

func (i *LiveOutput) ToLiveOutputOutput() LiveOutputOutput {
	return i.ToLiveOutputOutputWithContext(context.Background())
}

func (i *LiveOutput) ToLiveOutputOutputWithContext(ctx context.Context) LiveOutputOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LiveOutputOutput)
}

type LiveOutputOutput struct{ *pulumi.OutputState }

func (LiveOutputOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LiveOutput)(nil)).Elem()
}

func (o LiveOutputOutput) ToLiveOutputOutput() LiveOutputOutput {
	return o
}

func (o LiveOutputOutput) ToLiveOutputOutputWithContext(ctx context.Context) LiveOutputOutput {
	return o
}

func (o LiveOutputOutput) ArchiveWindowLength() pulumi.StringOutput {
	return o.ApplyT(func(v *LiveOutput) pulumi.StringOutput { return v.ArchiveWindowLength }).(pulumi.StringOutput)
}

func (o LiveOutputOutput) AssetName() pulumi.StringOutput {
	return o.ApplyT(func(v *LiveOutput) pulumi.StringOutput { return v.AssetName }).(pulumi.StringOutput)
}

func (o LiveOutputOutput) Created() pulumi.StringOutput {
	return o.ApplyT(func(v *LiveOutput) pulumi.StringOutput { return v.Created }).(pulumi.StringOutput)
}

func (o LiveOutputOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LiveOutput) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LiveOutputOutput) Hls() HlsResponsePtrOutput {
	return o.ApplyT(func(v *LiveOutput) HlsResponsePtrOutput { return v.Hls }).(HlsResponsePtrOutput)
}

func (o LiveOutputOutput) LastModified() pulumi.StringOutput {
	return o.ApplyT(func(v *LiveOutput) pulumi.StringOutput { return v.LastModified }).(pulumi.StringOutput)
}

func (o LiveOutputOutput) ManifestName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LiveOutput) pulumi.StringPtrOutput { return v.ManifestName }).(pulumi.StringPtrOutput)
}

func (o LiveOutputOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *LiveOutput) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o LiveOutputOutput) OutputSnapTime() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *LiveOutput) pulumi.Float64PtrOutput { return v.OutputSnapTime }).(pulumi.Float64PtrOutput)
}

func (o LiveOutputOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *LiveOutput) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LiveOutputOutput) ResourceState() pulumi.StringOutput {
	return o.ApplyT(func(v *LiveOutput) pulumi.StringOutput { return v.ResourceState }).(pulumi.StringOutput)
}

func (o LiveOutputOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *LiveOutput) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LiveOutputOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *LiveOutput) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LiveOutputOutput{})
}
