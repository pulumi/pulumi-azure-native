


package v20211101preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Video struct {
	pulumi.CustomResourceState

	Archival    VideoArchivalResponsePtrOutput  `pulumi:"archival"`
	ContentUrls VideoContentUrlsResponseOutput  `pulumi:"contentUrls"`
	Description pulumi.StringPtrOutput          `pulumi:"description"`
	Flags       VideoFlagsResponseOutput        `pulumi:"flags"`
	MediaInfo   VideoMediaInfoResponsePtrOutput `pulumi:"mediaInfo"`
	Name        pulumi.StringOutput             `pulumi:"name"`
	SystemData  SystemDataResponseOutput        `pulumi:"systemData"`
	Title       pulumi.StringPtrOutput          `pulumi:"title"`
	Type        pulumi.StringOutput             `pulumi:"type"`
}


func NewVideo(ctx *pulumi.Context,
	name string, args *VideoArgs, opts ...pulumi.ResourceOption) (*Video, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccountName == nil {
		return nil, errors.New("invalid value for required argument 'AccountName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:videoanalyzer:Video"),
		},
		{
			Type: pulumi.String("azure-native:videoanalyzer/v20210501preview:Video"),
		},
	})
	opts = append(opts, aliases)
	var resource Video
	err := ctx.RegisterResource("azure-native:videoanalyzer/v20211101preview:Video", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetVideo(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VideoState, opts ...pulumi.ResourceOption) (*Video, error) {
	var resource Video
	err := ctx.ReadResource("azure-native:videoanalyzer/v20211101preview:Video", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type videoState struct {
}

type VideoState struct {
}

func (VideoState) ElementType() reflect.Type {
	return reflect.TypeOf((*videoState)(nil)).Elem()
}

type videoArgs struct {
	AccountName       string          `pulumi:"accountName"`
	Archival          *VideoArchival  `pulumi:"archival"`
	Description       *string         `pulumi:"description"`
	MediaInfo         *VideoMediaInfo `pulumi:"mediaInfo"`
	ResourceGroupName string          `pulumi:"resourceGroupName"`
	Title             *string         `pulumi:"title"`
	VideoName         *string         `pulumi:"videoName"`
}


type VideoArgs struct {
	AccountName       pulumi.StringInput
	Archival          VideoArchivalPtrInput
	Description       pulumi.StringPtrInput
	MediaInfo         VideoMediaInfoPtrInput
	ResourceGroupName pulumi.StringInput
	Title             pulumi.StringPtrInput
	VideoName         pulumi.StringPtrInput
}

func (VideoArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*videoArgs)(nil)).Elem()
}

type VideoInput interface {
	pulumi.Input

	ToVideoOutput() VideoOutput
	ToVideoOutputWithContext(ctx context.Context) VideoOutput
}

func (*Video) ElementType() reflect.Type {
	return reflect.TypeOf((**Video)(nil)).Elem()
}

func (i *Video) ToVideoOutput() VideoOutput {
	return i.ToVideoOutputWithContext(context.Background())
}

func (i *Video) ToVideoOutputWithContext(ctx context.Context) VideoOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VideoOutput)
}

type VideoOutput struct{ *pulumi.OutputState }

func (VideoOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Video)(nil)).Elem()
}

func (o VideoOutput) ToVideoOutput() VideoOutput {
	return o
}

func (o VideoOutput) ToVideoOutputWithContext(ctx context.Context) VideoOutput {
	return o
}

func (o VideoOutput) Archival() VideoArchivalResponsePtrOutput {
	return o.ApplyT(func(v *Video) VideoArchivalResponsePtrOutput { return v.Archival }).(VideoArchivalResponsePtrOutput)
}

func (o VideoOutput) ContentUrls() VideoContentUrlsResponseOutput {
	return o.ApplyT(func(v *Video) VideoContentUrlsResponseOutput { return v.ContentUrls }).(VideoContentUrlsResponseOutput)
}

func (o VideoOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Video) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

func (o VideoOutput) Flags() VideoFlagsResponseOutput {
	return o.ApplyT(func(v *Video) VideoFlagsResponseOutput { return v.Flags }).(VideoFlagsResponseOutput)
}

func (o VideoOutput) MediaInfo() VideoMediaInfoResponsePtrOutput {
	return o.ApplyT(func(v *Video) VideoMediaInfoResponsePtrOutput { return v.MediaInfo }).(VideoMediaInfoResponsePtrOutput)
}

func (o VideoOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Video) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o VideoOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *Video) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o VideoOutput) Title() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Video) pulumi.StringPtrOutput { return v.Title }).(pulumi.StringPtrOutput)
}

func (o VideoOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Video) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(VideoOutput{})
}
