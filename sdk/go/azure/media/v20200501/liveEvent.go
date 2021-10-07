


package v20200501

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type LiveEvent struct {
	pulumi.CustomResourceState

	Created                 pulumi.StringOutput                       `pulumi:"created"`
	CrossSiteAccessPolicies CrossSiteAccessPoliciesResponsePtrOutput  `pulumi:"crossSiteAccessPolicies"`
	Description             pulumi.StringPtrOutput                    `pulumi:"description"`
	Encoding                LiveEventEncodingResponsePtrOutput        `pulumi:"encoding"`
	HostnamePrefix          pulumi.StringPtrOutput                    `pulumi:"hostnamePrefix"`
	Input                   LiveEventInputResponseOutput              `pulumi:"input"`
	LastModified            pulumi.StringOutput                       `pulumi:"lastModified"`
	Location                pulumi.StringOutput                       `pulumi:"location"`
	Name                    pulumi.StringOutput                       `pulumi:"name"`
	Preview                 LiveEventPreviewResponsePtrOutput         `pulumi:"preview"`
	ProvisioningState       pulumi.StringOutput                       `pulumi:"provisioningState"`
	ResourceState           pulumi.StringOutput                       `pulumi:"resourceState"`
	StreamOptions           pulumi.StringArrayOutput                  `pulumi:"streamOptions"`
	SystemData              SystemDataResponseOutput                  `pulumi:"systemData"`
	Tags                    pulumi.StringMapOutput                    `pulumi:"tags"`
	Transcriptions          LiveEventTranscriptionResponseArrayOutput `pulumi:"transcriptions"`
	Type                    pulumi.StringOutput                       `pulumi:"type"`
	UseStaticHostname       pulumi.BoolPtrOutput                      `pulumi:"useStaticHostname"`
}


func NewLiveEvent(ctx *pulumi.Context,
	name string, args *LiveEventArgs, opts ...pulumi.ResourceOption) (*LiveEvent, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccountName == nil {
		return nil, errors.New("invalid value for required argument 'AccountName'")
	}
	if args.Input == nil {
		return nil, errors.New("invalid value for required argument 'Input'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:media/v20200501:LiveEvent"),
		},
		{
			Type: pulumi.String("azure-native:media:LiveEvent"),
		},
		{
			Type: pulumi.String("azure-nextgen:media:LiveEvent"),
		},
		{
			Type: pulumi.String("azure-native:media/v20180330preview:LiveEvent"),
		},
		{
			Type: pulumi.String("azure-nextgen:media/v20180330preview:LiveEvent"),
		},
		{
			Type: pulumi.String("azure-native:media/v20180601preview:LiveEvent"),
		},
		{
			Type: pulumi.String("azure-nextgen:media/v20180601preview:LiveEvent"),
		},
		{
			Type: pulumi.String("azure-native:media/v20180701:LiveEvent"),
		},
		{
			Type: pulumi.String("azure-nextgen:media/v20180701:LiveEvent"),
		},
		{
			Type: pulumi.String("azure-native:media/v20190501preview:LiveEvent"),
		},
		{
			Type: pulumi.String("azure-nextgen:media/v20190501preview:LiveEvent"),
		},
		{
			Type: pulumi.String("azure-native:media/v20210601:LiveEvent"),
		},
		{
			Type: pulumi.String("azure-nextgen:media/v20210601:LiveEvent"),
		},
	})
	opts = append(opts, aliases)
	var resource LiveEvent
	err := ctx.RegisterResource("azure-native:media/v20200501:LiveEvent", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetLiveEvent(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LiveEventState, opts ...pulumi.ResourceOption) (*LiveEvent, error) {
	var resource LiveEvent
	err := ctx.ReadResource("azure-native:media/v20200501:LiveEvent", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type liveEventState struct {
}

type LiveEventState struct {
}

func (LiveEventState) ElementType() reflect.Type {
	return reflect.TypeOf((*liveEventState)(nil)).Elem()
}

type liveEventArgs struct {
	AccountName             string                   `pulumi:"accountName"`
	AutoStart               *bool                    `pulumi:"autoStart"`
	CrossSiteAccessPolicies *CrossSiteAccessPolicies `pulumi:"crossSiteAccessPolicies"`
	Description             *string                  `pulumi:"description"`
	Encoding                *LiveEventEncoding       `pulumi:"encoding"`
	HostnamePrefix          *string                  `pulumi:"hostnamePrefix"`
	Input                   LiveEventInputType       `pulumi:"input"`
	LiveEventName           *string                  `pulumi:"liveEventName"`
	Location                *string                  `pulumi:"location"`
	Preview                 *LiveEventPreview        `pulumi:"preview"`
	ResourceGroupName       string                   `pulumi:"resourceGroupName"`
	StreamOptions           []string                 `pulumi:"streamOptions"`
	Tags                    map[string]string        `pulumi:"tags"`
	Transcriptions          []LiveEventTranscription `pulumi:"transcriptions"`
	UseStaticHostname       *bool                    `pulumi:"useStaticHostname"`
}


type LiveEventArgs struct {
	AccountName             pulumi.StringInput
	AutoStart               pulumi.BoolPtrInput
	CrossSiteAccessPolicies CrossSiteAccessPoliciesPtrInput
	Description             pulumi.StringPtrInput
	Encoding                LiveEventEncodingPtrInput
	HostnamePrefix          pulumi.StringPtrInput
	Input                   LiveEventInputTypeInput
	LiveEventName           pulumi.StringPtrInput
	Location                pulumi.StringPtrInput
	Preview                 LiveEventPreviewPtrInput
	ResourceGroupName       pulumi.StringInput
	StreamOptions           pulumi.StringArrayInput
	Tags                    pulumi.StringMapInput
	Transcriptions          LiveEventTranscriptionArrayInput
	UseStaticHostname       pulumi.BoolPtrInput
}

func (LiveEventArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*liveEventArgs)(nil)).Elem()
}

type LiveEventInput interface {
	pulumi.Input

	ToLiveEventOutput() LiveEventOutput
	ToLiveEventOutputWithContext(ctx context.Context) LiveEventOutput
}

func (*LiveEvent) ElementType() reflect.Type {
	return reflect.TypeOf((*LiveEvent)(nil))
}

func (i *LiveEvent) ToLiveEventOutput() LiveEventOutput {
	return i.ToLiveEventOutputWithContext(context.Background())
}

func (i *LiveEvent) ToLiveEventOutputWithContext(ctx context.Context) LiveEventOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LiveEventOutput)
}

type LiveEventOutput struct{ *pulumi.OutputState }

func (LiveEventOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LiveEvent)(nil))
}

func (o LiveEventOutput) ToLiveEventOutput() LiveEventOutput {
	return o
}

func (o LiveEventOutput) ToLiveEventOutputWithContext(ctx context.Context) LiveEventOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(LiveEventOutput{})
}
