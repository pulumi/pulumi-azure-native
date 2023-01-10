


package v20221101preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AFDOrigin struct {
	pulumi.CustomResourceState

	AzureOrigin                 ResourceReferenceResponsePtrOutput                   `pulumi:"azureOrigin"`
	DeploymentStatus            pulumi.StringOutput                                  `pulumi:"deploymentStatus"`
	EnabledState                pulumi.StringPtrOutput                               `pulumi:"enabledState"`
	EnforceCertificateNameCheck pulumi.BoolPtrOutput                                 `pulumi:"enforceCertificateNameCheck"`
	HostName                    pulumi.StringOutput                                  `pulumi:"hostName"`
	HttpPort                    pulumi.IntPtrOutput                                  `pulumi:"httpPort"`
	HttpsPort                   pulumi.IntPtrOutput                                  `pulumi:"httpsPort"`
	Name                        pulumi.StringOutput                                  `pulumi:"name"`
	OriginGroupName             pulumi.StringOutput                                  `pulumi:"originGroupName"`
	OriginHostHeader            pulumi.StringPtrOutput                               `pulumi:"originHostHeader"`
	Priority                    pulumi.IntPtrOutput                                  `pulumi:"priority"`
	ProvisioningState           pulumi.StringOutput                                  `pulumi:"provisioningState"`
	SharedPrivateLinkResource   SharedPrivateLinkResourcePropertiesResponsePtrOutput `pulumi:"sharedPrivateLinkResource"`
	SystemData                  SystemDataResponseOutput                             `pulumi:"systemData"`
	Type                        pulumi.StringOutput                                  `pulumi:"type"`
	Weight                      pulumi.IntPtrOutput                                  `pulumi:"weight"`
}


func NewAFDOrigin(ctx *pulumi.Context,
	name string, args *AFDOriginArgs, opts ...pulumi.ResourceOption) (*AFDOrigin, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.HostName == nil {
		return nil, errors.New("invalid value for required argument 'HostName'")
	}
	if args.OriginGroupName == nil {
		return nil, errors.New("invalid value for required argument 'OriginGroupName'")
	}
	if args.ProfileName == nil {
		return nil, errors.New("invalid value for required argument 'ProfileName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if isZero(args.EnforceCertificateNameCheck) {
		args.EnforceCertificateNameCheck = pulumi.BoolPtr(true)
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:cdn:AFDOrigin"),
		},
		{
			Type: pulumi.String("azure-native:cdn/v20200901:AFDOrigin"),
		},
		{
			Type: pulumi.String("azure-native:cdn/v20210601:AFDOrigin"),
		},
		{
			Type: pulumi.String("azure-native:cdn/v20220501preview:AFDOrigin"),
		},
	})
	opts = append(opts, aliases)
	var resource AFDOrigin
	err := ctx.RegisterResource("azure-native:cdn/v20221101preview:AFDOrigin", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetAFDOrigin(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AFDOriginState, opts ...pulumi.ResourceOption) (*AFDOrigin, error) {
	var resource AFDOrigin
	err := ctx.ReadResource("azure-native:cdn/v20221101preview:AFDOrigin", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type afdoriginState struct {
}

type AFDOriginState struct {
}

func (AFDOriginState) ElementType() reflect.Type {
	return reflect.TypeOf((*afdoriginState)(nil)).Elem()
}

type afdoriginArgs struct {
	AzureOrigin                 *ResourceReference                   `pulumi:"azureOrigin"`
	EnabledState                *string                              `pulumi:"enabledState"`
	EnforceCertificateNameCheck *bool                                `pulumi:"enforceCertificateNameCheck"`
	HostName                    string                               `pulumi:"hostName"`
	HttpPort                    *int                                 `pulumi:"httpPort"`
	HttpsPort                   *int                                 `pulumi:"httpsPort"`
	OriginGroupName             string                               `pulumi:"originGroupName"`
	OriginHostHeader            *string                              `pulumi:"originHostHeader"`
	OriginName                  *string                              `pulumi:"originName"`
	Priority                    *int                                 `pulumi:"priority"`
	ProfileName                 string                               `pulumi:"profileName"`
	ResourceGroupName           string                               `pulumi:"resourceGroupName"`
	SharedPrivateLinkResource   *SharedPrivateLinkResourceProperties `pulumi:"sharedPrivateLinkResource"`
	Weight                      *int                                 `pulumi:"weight"`
}


type AFDOriginArgs struct {
	AzureOrigin                 ResourceReferencePtrInput
	EnabledState                pulumi.StringPtrInput
	EnforceCertificateNameCheck pulumi.BoolPtrInput
	HostName                    pulumi.StringInput
	HttpPort                    pulumi.IntPtrInput
	HttpsPort                   pulumi.IntPtrInput
	OriginGroupName             pulumi.StringInput
	OriginHostHeader            pulumi.StringPtrInput
	OriginName                  pulumi.StringPtrInput
	Priority                    pulumi.IntPtrInput
	ProfileName                 pulumi.StringInput
	ResourceGroupName           pulumi.StringInput
	SharedPrivateLinkResource   SharedPrivateLinkResourcePropertiesPtrInput
	Weight                      pulumi.IntPtrInput
}

func (AFDOriginArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*afdoriginArgs)(nil)).Elem()
}

type AFDOriginInput interface {
	pulumi.Input

	ToAFDOriginOutput() AFDOriginOutput
	ToAFDOriginOutputWithContext(ctx context.Context) AFDOriginOutput
}

func (*AFDOrigin) ElementType() reflect.Type {
	return reflect.TypeOf((**AFDOrigin)(nil)).Elem()
}

func (i *AFDOrigin) ToAFDOriginOutput() AFDOriginOutput {
	return i.ToAFDOriginOutputWithContext(context.Background())
}

func (i *AFDOrigin) ToAFDOriginOutputWithContext(ctx context.Context) AFDOriginOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AFDOriginOutput)
}

type AFDOriginOutput struct{ *pulumi.OutputState }

func (AFDOriginOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AFDOrigin)(nil)).Elem()
}

func (o AFDOriginOutput) ToAFDOriginOutput() AFDOriginOutput {
	return o
}

func (o AFDOriginOutput) ToAFDOriginOutputWithContext(ctx context.Context) AFDOriginOutput {
	return o
}

func (o AFDOriginOutput) AzureOrigin() ResourceReferenceResponsePtrOutput {
	return o.ApplyT(func(v *AFDOrigin) ResourceReferenceResponsePtrOutput { return v.AzureOrigin }).(ResourceReferenceResponsePtrOutput)
}

func (o AFDOriginOutput) DeploymentStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *AFDOrigin) pulumi.StringOutput { return v.DeploymentStatus }).(pulumi.StringOutput)
}

func (o AFDOriginOutput) EnabledState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AFDOrigin) pulumi.StringPtrOutput { return v.EnabledState }).(pulumi.StringPtrOutput)
}

func (o AFDOriginOutput) EnforceCertificateNameCheck() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AFDOrigin) pulumi.BoolPtrOutput { return v.EnforceCertificateNameCheck }).(pulumi.BoolPtrOutput)
}

func (o AFDOriginOutput) HostName() pulumi.StringOutput {
	return o.ApplyT(func(v *AFDOrigin) pulumi.StringOutput { return v.HostName }).(pulumi.StringOutput)
}

func (o AFDOriginOutput) HttpPort() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *AFDOrigin) pulumi.IntPtrOutput { return v.HttpPort }).(pulumi.IntPtrOutput)
}

func (o AFDOriginOutput) HttpsPort() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *AFDOrigin) pulumi.IntPtrOutput { return v.HttpsPort }).(pulumi.IntPtrOutput)
}

func (o AFDOriginOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *AFDOrigin) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o AFDOriginOutput) OriginGroupName() pulumi.StringOutput {
	return o.ApplyT(func(v *AFDOrigin) pulumi.StringOutput { return v.OriginGroupName }).(pulumi.StringOutput)
}

func (o AFDOriginOutput) OriginHostHeader() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AFDOrigin) pulumi.StringPtrOutput { return v.OriginHostHeader }).(pulumi.StringPtrOutput)
}

func (o AFDOriginOutput) Priority() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *AFDOrigin) pulumi.IntPtrOutput { return v.Priority }).(pulumi.IntPtrOutput)
}

func (o AFDOriginOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *AFDOrigin) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o AFDOriginOutput) SharedPrivateLinkResource() SharedPrivateLinkResourcePropertiesResponsePtrOutput {
	return o.ApplyT(func(v *AFDOrigin) SharedPrivateLinkResourcePropertiesResponsePtrOutput {
		return v.SharedPrivateLinkResource
	}).(SharedPrivateLinkResourcePropertiesResponsePtrOutput)
}

func (o AFDOriginOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *AFDOrigin) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o AFDOriginOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *AFDOrigin) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o AFDOriginOutput) Weight() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *AFDOrigin) pulumi.IntPtrOutput { return v.Weight }).(pulumi.IntPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(AFDOriginOutput{})
}
