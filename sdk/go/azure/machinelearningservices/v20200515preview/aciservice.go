


package v20200515preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ACIService struct {
	pulumi.CustomResourceState

	Identity   IdentityResponsePtrOutput `pulumi:"identity"`
	Location   pulumi.StringPtrOutput    `pulumi:"location"`
	Name       pulumi.StringOutput       `pulumi:"name"`
	Properties pulumi.AnyOutput          `pulumi:"properties"`
	Sku        SkuResponsePtrOutput      `pulumi:"sku"`
	Tags       pulumi.StringMapOutput    `pulumi:"tags"`
	Type       pulumi.StringOutput       `pulumi:"type"`
}


func NewACIService(ctx *pulumi.Context,
	name string, args *ACIServiceArgs, opts ...pulumi.ResourceOption) (*ACIService, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ComputeType == nil {
		return nil, errors.New("invalid value for required argument 'ComputeType'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.WorkspaceName == nil {
		return nil, errors.New("invalid value for required argument 'WorkspaceName'")
	}
	if isZero(args.AppInsightsEnabled) {
		args.AppInsightsEnabled = pulumi.BoolPtr(false)
	}
	if isZero(args.AuthEnabled) {
		args.AuthEnabled = pulumi.BoolPtr(false)
	}
	args.ComputeType = pulumi.String("ACI")
	if isZero(args.SslEnabled) {
		args.SslEnabled = pulumi.BoolPtr(false)
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:machinelearningservices:ACIService"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20200501preview:ACIService"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20200901preview:ACIService"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20210101:ACIService"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20210401:ACIService"),
		},
	})
	opts = append(opts, aliases)
	var resource ACIService
	err := ctx.RegisterResource("azure-native:machinelearningservices/v20200515preview:ACIService", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetACIService(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ACIServiceState, opts ...pulumi.ResourceOption) (*ACIService, error) {
	var resource ACIService
	err := ctx.ReadResource("azure-native:machinelearningservices/v20200515preview:ACIService", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type aciserviceState struct {
}

type ACIServiceState struct {
}

func (ACIServiceState) ElementType() reflect.Type {
	return reflect.TypeOf((*aciserviceState)(nil)).Elem()
}

type aciserviceArgs struct {
	AppInsightsEnabled            *bool                                        `pulumi:"appInsightsEnabled"`
	AuthEnabled                   *bool                                        `pulumi:"authEnabled"`
	Cname                         *string                                      `pulumi:"cname"`
	ComputeType                   string                                       `pulumi:"computeType"`
	ContainerResourceRequirements *ContainerResourceRequirements               `pulumi:"containerResourceRequirements"`
	DataCollection                *ACIServiceCreateRequestDataCollection       `pulumi:"dataCollection"`
	Description                   *string                                      `pulumi:"description"`
	DnsNameLabel                  *string                                      `pulumi:"dnsNameLabel"`
	EncryptionProperties          *ACIServiceCreateRequestEncryptionProperties `pulumi:"encryptionProperties"`
	EnvironmentImageRequest       *CreateServiceRequestEnvironmentImageRequest `pulumi:"environmentImageRequest"`
	Keys                          *CreateServiceRequestKeys                    `pulumi:"keys"`
	KvTags                        map[string]string                            `pulumi:"kvTags"`
	Location                      *string                                      `pulumi:"location"`
	Properties                    map[string]string                            `pulumi:"properties"`
	ResourceGroupName             string                                       `pulumi:"resourceGroupName"`
	ServiceName                   *string                                      `pulumi:"serviceName"`
	SslCertificate                *string                                      `pulumi:"sslCertificate"`
	SslEnabled                    *bool                                        `pulumi:"sslEnabled"`
	SslKey                        *string                                      `pulumi:"sslKey"`
	VnetConfiguration             *ACIServiceCreateRequestVnetConfiguration    `pulumi:"vnetConfiguration"`
	WorkspaceName                 string                                       `pulumi:"workspaceName"`
}


type ACIServiceArgs struct {
	AppInsightsEnabled            pulumi.BoolPtrInput
	AuthEnabled                   pulumi.BoolPtrInput
	Cname                         pulumi.StringPtrInput
	ComputeType                   pulumi.StringInput
	ContainerResourceRequirements ContainerResourceRequirementsPtrInput
	DataCollection                ACIServiceCreateRequestDataCollectionPtrInput
	Description                   pulumi.StringPtrInput
	DnsNameLabel                  pulumi.StringPtrInput
	EncryptionProperties          ACIServiceCreateRequestEncryptionPropertiesPtrInput
	EnvironmentImageRequest       CreateServiceRequestEnvironmentImageRequestPtrInput
	Keys                          CreateServiceRequestKeysPtrInput
	KvTags                        pulumi.StringMapInput
	Location                      pulumi.StringPtrInput
	Properties                    pulumi.StringMapInput
	ResourceGroupName             pulumi.StringInput
	ServiceName                   pulumi.StringPtrInput
	SslCertificate                pulumi.StringPtrInput
	SslEnabled                    pulumi.BoolPtrInput
	SslKey                        pulumi.StringPtrInput
	VnetConfiguration             ACIServiceCreateRequestVnetConfigurationPtrInput
	WorkspaceName                 pulumi.StringInput
}

func (ACIServiceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*aciserviceArgs)(nil)).Elem()
}

type ACIServiceInput interface {
	pulumi.Input

	ToACIServiceOutput() ACIServiceOutput
	ToACIServiceOutputWithContext(ctx context.Context) ACIServiceOutput
}

func (*ACIService) ElementType() reflect.Type {
	return reflect.TypeOf((**ACIService)(nil)).Elem()
}

func (i *ACIService) ToACIServiceOutput() ACIServiceOutput {
	return i.ToACIServiceOutputWithContext(context.Background())
}

func (i *ACIService) ToACIServiceOutputWithContext(ctx context.Context) ACIServiceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ACIServiceOutput)
}

type ACIServiceOutput struct{ *pulumi.OutputState }

func (ACIServiceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ACIService)(nil)).Elem()
}

func (o ACIServiceOutput) ToACIServiceOutput() ACIServiceOutput {
	return o
}

func (o ACIServiceOutput) ToACIServiceOutputWithContext(ctx context.Context) ACIServiceOutput {
	return o
}

func (o ACIServiceOutput) Identity() IdentityResponsePtrOutput {
	return o.ApplyT(func(v *ACIService) IdentityResponsePtrOutput { return v.Identity }).(IdentityResponsePtrOutput)
}

func (o ACIServiceOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ACIService) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

func (o ACIServiceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ACIService) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ACIServiceOutput) Properties() pulumi.AnyOutput {
	return o.ApplyT(func(v *ACIService) pulumi.AnyOutput { return v.Properties }).(pulumi.AnyOutput)
}

func (o ACIServiceOutput) Sku() SkuResponsePtrOutput {
	return o.ApplyT(func(v *ACIService) SkuResponsePtrOutput { return v.Sku }).(SkuResponsePtrOutput)
}

func (o ACIServiceOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ACIService) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o ACIServiceOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ACIService) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ACIServiceOutput{})
}
