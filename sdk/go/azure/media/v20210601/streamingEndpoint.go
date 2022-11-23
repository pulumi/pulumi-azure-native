


package v20210601

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type StreamingEndpoint struct {
	pulumi.CustomResourceState

	AccessControl           StreamingEndpointAccessControlResponsePtrOutput `pulumi:"accessControl"`
	AvailabilitySetName     pulumi.StringPtrOutput                          `pulumi:"availabilitySetName"`
	CdnEnabled              pulumi.BoolPtrOutput                            `pulumi:"cdnEnabled"`
	CdnProfile              pulumi.StringPtrOutput                          `pulumi:"cdnProfile"`
	CdnProvider             pulumi.StringPtrOutput                          `pulumi:"cdnProvider"`
	Created                 pulumi.StringOutput                             `pulumi:"created"`
	CrossSiteAccessPolicies CrossSiteAccessPoliciesResponsePtrOutput        `pulumi:"crossSiteAccessPolicies"`
	CustomHostNames         pulumi.StringArrayOutput                        `pulumi:"customHostNames"`
	Description             pulumi.StringPtrOutput                          `pulumi:"description"`
	FreeTrialEndTime        pulumi.StringOutput                             `pulumi:"freeTrialEndTime"`
	HostName                pulumi.StringOutput                             `pulumi:"hostName"`
	LastModified            pulumi.StringOutput                             `pulumi:"lastModified"`
	Location                pulumi.StringOutput                             `pulumi:"location"`
	MaxCacheAge             pulumi.Float64PtrOutput                         `pulumi:"maxCacheAge"`
	Name                    pulumi.StringOutput                             `pulumi:"name"`
	ProvisioningState       pulumi.StringOutput                             `pulumi:"provisioningState"`
	ResourceState           pulumi.StringOutput                             `pulumi:"resourceState"`
	ScaleUnits              pulumi.IntOutput                                `pulumi:"scaleUnits"`
	SystemData              SystemDataResponseOutput                        `pulumi:"systemData"`
	Tags                    pulumi.StringMapOutput                          `pulumi:"tags"`
	Type                    pulumi.StringOutput                             `pulumi:"type"`
}


func NewStreamingEndpoint(ctx *pulumi.Context,
	name string, args *StreamingEndpointArgs, opts ...pulumi.ResourceOption) (*StreamingEndpoint, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccountName == nil {
		return nil, errors.New("invalid value for required argument 'AccountName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ScaleUnits == nil {
		return nil, errors.New("invalid value for required argument 'ScaleUnits'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:media:StreamingEndpoint"),
		},
		{
			Type: pulumi.String("azure-native:media/v20180330preview:StreamingEndpoint"),
		},
		{
			Type: pulumi.String("azure-native:media/v20180601preview:StreamingEndpoint"),
		},
		{
			Type: pulumi.String("azure-native:media/v20180701:StreamingEndpoint"),
		},
		{
			Type: pulumi.String("azure-native:media/v20190501preview:StreamingEndpoint"),
		},
		{
			Type: pulumi.String("azure-native:media/v20200501:StreamingEndpoint"),
		},
		{
			Type: pulumi.String("azure-native:media/v20211101:StreamingEndpoint"),
		},
		{
			Type: pulumi.String("azure-native:media/v20220801:StreamingEndpoint"),
		},
	})
	opts = append(opts, aliases)
	var resource StreamingEndpoint
	err := ctx.RegisterResource("azure-native:media/v20210601:StreamingEndpoint", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetStreamingEndpoint(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *StreamingEndpointState, opts ...pulumi.ResourceOption) (*StreamingEndpoint, error) {
	var resource StreamingEndpoint
	err := ctx.ReadResource("azure-native:media/v20210601:StreamingEndpoint", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type streamingEndpointState struct {
}

type StreamingEndpointState struct {
}

func (StreamingEndpointState) ElementType() reflect.Type {
	return reflect.TypeOf((*streamingEndpointState)(nil)).Elem()
}

type streamingEndpointArgs struct {
	AccessControl           *StreamingEndpointAccessControl `pulumi:"accessControl"`
	AccountName             string                          `pulumi:"accountName"`
	AutoStart               *bool                           `pulumi:"autoStart"`
	AvailabilitySetName     *string                         `pulumi:"availabilitySetName"`
	CdnEnabled              *bool                           `pulumi:"cdnEnabled"`
	CdnProfile              *string                         `pulumi:"cdnProfile"`
	CdnProvider             *string                         `pulumi:"cdnProvider"`
	CrossSiteAccessPolicies *CrossSiteAccessPolicies        `pulumi:"crossSiteAccessPolicies"`
	CustomHostNames         []string                        `pulumi:"customHostNames"`
	Description             *string                         `pulumi:"description"`
	Location                *string                         `pulumi:"location"`
	MaxCacheAge             *float64                        `pulumi:"maxCacheAge"`
	ResourceGroupName       string                          `pulumi:"resourceGroupName"`
	ScaleUnits              int                             `pulumi:"scaleUnits"`
	StreamingEndpointName   *string                         `pulumi:"streamingEndpointName"`
	Tags                    map[string]string               `pulumi:"tags"`
}


type StreamingEndpointArgs struct {
	AccessControl           StreamingEndpointAccessControlPtrInput
	AccountName             pulumi.StringInput
	AutoStart               pulumi.BoolPtrInput
	AvailabilitySetName     pulumi.StringPtrInput
	CdnEnabled              pulumi.BoolPtrInput
	CdnProfile              pulumi.StringPtrInput
	CdnProvider             pulumi.StringPtrInput
	CrossSiteAccessPolicies CrossSiteAccessPoliciesPtrInput
	CustomHostNames         pulumi.StringArrayInput
	Description             pulumi.StringPtrInput
	Location                pulumi.StringPtrInput
	MaxCacheAge             pulumi.Float64PtrInput
	ResourceGroupName       pulumi.StringInput
	ScaleUnits              pulumi.IntInput
	StreamingEndpointName   pulumi.StringPtrInput
	Tags                    pulumi.StringMapInput
}

func (StreamingEndpointArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*streamingEndpointArgs)(nil)).Elem()
}

type StreamingEndpointInput interface {
	pulumi.Input

	ToStreamingEndpointOutput() StreamingEndpointOutput
	ToStreamingEndpointOutputWithContext(ctx context.Context) StreamingEndpointOutput
}

func (*StreamingEndpoint) ElementType() reflect.Type {
	return reflect.TypeOf((**StreamingEndpoint)(nil)).Elem()
}

func (i *StreamingEndpoint) ToStreamingEndpointOutput() StreamingEndpointOutput {
	return i.ToStreamingEndpointOutputWithContext(context.Background())
}

func (i *StreamingEndpoint) ToStreamingEndpointOutputWithContext(ctx context.Context) StreamingEndpointOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StreamingEndpointOutput)
}

type StreamingEndpointOutput struct{ *pulumi.OutputState }

func (StreamingEndpointOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StreamingEndpoint)(nil)).Elem()
}

func (o StreamingEndpointOutput) ToStreamingEndpointOutput() StreamingEndpointOutput {
	return o
}

func (o StreamingEndpointOutput) ToStreamingEndpointOutputWithContext(ctx context.Context) StreamingEndpointOutput {
	return o
}

func (o StreamingEndpointOutput) AccessControl() StreamingEndpointAccessControlResponsePtrOutput {
	return o.ApplyT(func(v *StreamingEndpoint) StreamingEndpointAccessControlResponsePtrOutput { return v.AccessControl }).(StreamingEndpointAccessControlResponsePtrOutput)
}

func (o StreamingEndpointOutput) AvailabilitySetName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StreamingEndpoint) pulumi.StringPtrOutput { return v.AvailabilitySetName }).(pulumi.StringPtrOutput)
}

func (o StreamingEndpointOutput) CdnEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *StreamingEndpoint) pulumi.BoolPtrOutput { return v.CdnEnabled }).(pulumi.BoolPtrOutput)
}

func (o StreamingEndpointOutput) CdnProfile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StreamingEndpoint) pulumi.StringPtrOutput { return v.CdnProfile }).(pulumi.StringPtrOutput)
}

func (o StreamingEndpointOutput) CdnProvider() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StreamingEndpoint) pulumi.StringPtrOutput { return v.CdnProvider }).(pulumi.StringPtrOutput)
}

func (o StreamingEndpointOutput) Created() pulumi.StringOutput {
	return o.ApplyT(func(v *StreamingEndpoint) pulumi.StringOutput { return v.Created }).(pulumi.StringOutput)
}

func (o StreamingEndpointOutput) CrossSiteAccessPolicies() CrossSiteAccessPoliciesResponsePtrOutput {
	return o.ApplyT(func(v *StreamingEndpoint) CrossSiteAccessPoliciesResponsePtrOutput { return v.CrossSiteAccessPolicies }).(CrossSiteAccessPoliciesResponsePtrOutput)
}

func (o StreamingEndpointOutput) CustomHostNames() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *StreamingEndpoint) pulumi.StringArrayOutput { return v.CustomHostNames }).(pulumi.StringArrayOutput)
}

func (o StreamingEndpointOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StreamingEndpoint) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

func (o StreamingEndpointOutput) FreeTrialEndTime() pulumi.StringOutput {
	return o.ApplyT(func(v *StreamingEndpoint) pulumi.StringOutput { return v.FreeTrialEndTime }).(pulumi.StringOutput)
}

func (o StreamingEndpointOutput) HostName() pulumi.StringOutput {
	return o.ApplyT(func(v *StreamingEndpoint) pulumi.StringOutput { return v.HostName }).(pulumi.StringOutput)
}

func (o StreamingEndpointOutput) LastModified() pulumi.StringOutput {
	return o.ApplyT(func(v *StreamingEndpoint) pulumi.StringOutput { return v.LastModified }).(pulumi.StringOutput)
}

func (o StreamingEndpointOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *StreamingEndpoint) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o StreamingEndpointOutput) MaxCacheAge() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *StreamingEndpoint) pulumi.Float64PtrOutput { return v.MaxCacheAge }).(pulumi.Float64PtrOutput)
}

func (o StreamingEndpointOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *StreamingEndpoint) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o StreamingEndpointOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *StreamingEndpoint) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o StreamingEndpointOutput) ResourceState() pulumi.StringOutput {
	return o.ApplyT(func(v *StreamingEndpoint) pulumi.StringOutput { return v.ResourceState }).(pulumi.StringOutput)
}

func (o StreamingEndpointOutput) ScaleUnits() pulumi.IntOutput {
	return o.ApplyT(func(v *StreamingEndpoint) pulumi.IntOutput { return v.ScaleUnits }).(pulumi.IntOutput)
}

func (o StreamingEndpointOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *StreamingEndpoint) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o StreamingEndpointOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *StreamingEndpoint) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o StreamingEndpointOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *StreamingEndpoint) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(StreamingEndpointOutput{})
}
