


package v20220801preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type WebPubSubCustomDomain struct {
	pulumi.CustomResourceState

	CustomCertificate ResourceReferenceResponseOutput `pulumi:"customCertificate"`
	DomainName        pulumi.StringOutput             `pulumi:"domainName"`
	Name              pulumi.StringOutput             `pulumi:"name"`
	ProvisioningState pulumi.StringOutput             `pulumi:"provisioningState"`
	SystemData        SystemDataResponseOutput        `pulumi:"systemData"`
	Type              pulumi.StringOutput             `pulumi:"type"`
}


func NewWebPubSubCustomDomain(ctx *pulumi.Context,
	name string, args *WebPubSubCustomDomainArgs, opts ...pulumi.ResourceOption) (*WebPubSubCustomDomain, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.CustomCertificate == nil {
		return nil, errors.New("invalid value for required argument 'CustomCertificate'")
	}
	if args.DomainName == nil {
		return nil, errors.New("invalid value for required argument 'DomainName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ResourceName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceName'")
	}
	var resource WebPubSubCustomDomain
	err := ctx.RegisterResource("azure-native:webpubsub/v20220801preview:WebPubSubCustomDomain", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetWebPubSubCustomDomain(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WebPubSubCustomDomainState, opts ...pulumi.ResourceOption) (*WebPubSubCustomDomain, error) {
	var resource WebPubSubCustomDomain
	err := ctx.ReadResource("azure-native:webpubsub/v20220801preview:WebPubSubCustomDomain", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type webPubSubCustomDomainState struct {
}

type WebPubSubCustomDomainState struct {
}

func (WebPubSubCustomDomainState) ElementType() reflect.Type {
	return reflect.TypeOf((*webPubSubCustomDomainState)(nil)).Elem()
}

type webPubSubCustomDomainArgs struct {
	CustomCertificate ResourceReference `pulumi:"customCertificate"`
	DomainName        string            `pulumi:"domainName"`
	Name              *string           `pulumi:"name"`
	ResourceGroupName string            `pulumi:"resourceGroupName"`
	ResourceName      string            `pulumi:"resourceName"`
}


type WebPubSubCustomDomainArgs struct {
	CustomCertificate ResourceReferenceInput
	DomainName        pulumi.StringInput
	Name              pulumi.StringPtrInput
	ResourceGroupName pulumi.StringInput
	ResourceName      pulumi.StringInput
}

func (WebPubSubCustomDomainArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*webPubSubCustomDomainArgs)(nil)).Elem()
}

type WebPubSubCustomDomainInput interface {
	pulumi.Input

	ToWebPubSubCustomDomainOutput() WebPubSubCustomDomainOutput
	ToWebPubSubCustomDomainOutputWithContext(ctx context.Context) WebPubSubCustomDomainOutput
}

func (*WebPubSubCustomDomain) ElementType() reflect.Type {
	return reflect.TypeOf((**WebPubSubCustomDomain)(nil)).Elem()
}

func (i *WebPubSubCustomDomain) ToWebPubSubCustomDomainOutput() WebPubSubCustomDomainOutput {
	return i.ToWebPubSubCustomDomainOutputWithContext(context.Background())
}

func (i *WebPubSubCustomDomain) ToWebPubSubCustomDomainOutputWithContext(ctx context.Context) WebPubSubCustomDomainOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebPubSubCustomDomainOutput)
}

type WebPubSubCustomDomainOutput struct{ *pulumi.OutputState }

func (WebPubSubCustomDomainOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WebPubSubCustomDomain)(nil)).Elem()
}

func (o WebPubSubCustomDomainOutput) ToWebPubSubCustomDomainOutput() WebPubSubCustomDomainOutput {
	return o
}

func (o WebPubSubCustomDomainOutput) ToWebPubSubCustomDomainOutputWithContext(ctx context.Context) WebPubSubCustomDomainOutput {
	return o
}

func (o WebPubSubCustomDomainOutput) CustomCertificate() ResourceReferenceResponseOutput {
	return o.ApplyT(func(v *WebPubSubCustomDomain) ResourceReferenceResponseOutput { return v.CustomCertificate }).(ResourceReferenceResponseOutput)
}

func (o WebPubSubCustomDomainOutput) DomainName() pulumi.StringOutput {
	return o.ApplyT(func(v *WebPubSubCustomDomain) pulumi.StringOutput { return v.DomainName }).(pulumi.StringOutput)
}

func (o WebPubSubCustomDomainOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *WebPubSubCustomDomain) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o WebPubSubCustomDomainOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *WebPubSubCustomDomain) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o WebPubSubCustomDomainOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *WebPubSubCustomDomain) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o WebPubSubCustomDomainOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *WebPubSubCustomDomain) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(WebPubSubCustomDomainOutput{})
}
