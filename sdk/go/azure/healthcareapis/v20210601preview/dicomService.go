


package v20210601preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type DicomService struct {
	pulumi.CustomResourceState

	AuthenticationConfiguration DicomServiceAuthenticationConfigurationResponsePtrOutput `pulumi:"authenticationConfiguration"`
	Etag                        pulumi.StringPtrOutput                                   `pulumi:"etag"`
	Location                    pulumi.StringPtrOutput                                   `pulumi:"location"`
	Name                        pulumi.StringOutput                                      `pulumi:"name"`
	ProvisioningState           pulumi.StringOutput                                      `pulumi:"provisioningState"`
	ServiceUrl                  pulumi.StringOutput                                      `pulumi:"serviceUrl"`
	SystemData                  SystemDataResponseOutput                                 `pulumi:"systemData"`
	Tags                        pulumi.StringMapOutput                                   `pulumi:"tags"`
	Type                        pulumi.StringOutput                                      `pulumi:"type"`
}


func NewDicomService(ctx *pulumi.Context,
	name string, args *DicomServiceArgs, opts ...pulumi.ResourceOption) (*DicomService, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.WorkspaceName == nil {
		return nil, errors.New("invalid value for required argument 'WorkspaceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:healthcareapis:DicomService"),
		},
		{
			Type: pulumi.String("azure-native:healthcareapis/v20211101:DicomService"),
		},
		{
			Type: pulumi.String("azure-native:healthcareapis/v20220131preview:DicomService"),
		},
		{
			Type: pulumi.String("azure-native:healthcareapis/v20220515:DicomService"),
		},
		{
			Type: pulumi.String("azure-native:healthcareapis/v20220601:DicomService"),
		},
		{
			Type: pulumi.String("azure-native:healthcareapis/v20221001preview:DicomService"),
		},
	})
	opts = append(opts, aliases)
	var resource DicomService
	err := ctx.RegisterResource("azure-native:healthcareapis/v20210601preview:DicomService", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetDicomService(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DicomServiceState, opts ...pulumi.ResourceOption) (*DicomService, error) {
	var resource DicomService
	err := ctx.ReadResource("azure-native:healthcareapis/v20210601preview:DicomService", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type dicomServiceState struct {
}

type DicomServiceState struct {
}

func (DicomServiceState) ElementType() reflect.Type {
	return reflect.TypeOf((*dicomServiceState)(nil)).Elem()
}

type dicomServiceArgs struct {
	DicomServiceName  *string           `pulumi:"dicomServiceName"`
	Location          *string           `pulumi:"location"`
	ResourceGroupName string            `pulumi:"resourceGroupName"`
	Tags              map[string]string `pulumi:"tags"`
	WorkspaceName     string            `pulumi:"workspaceName"`
}


type DicomServiceArgs struct {
	DicomServiceName  pulumi.StringPtrInput
	Location          pulumi.StringPtrInput
	ResourceGroupName pulumi.StringInput
	Tags              pulumi.StringMapInput
	WorkspaceName     pulumi.StringInput
}

func (DicomServiceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dicomServiceArgs)(nil)).Elem()
}

type DicomServiceInput interface {
	pulumi.Input

	ToDicomServiceOutput() DicomServiceOutput
	ToDicomServiceOutputWithContext(ctx context.Context) DicomServiceOutput
}

func (*DicomService) ElementType() reflect.Type {
	return reflect.TypeOf((**DicomService)(nil)).Elem()
}

func (i *DicomService) ToDicomServiceOutput() DicomServiceOutput {
	return i.ToDicomServiceOutputWithContext(context.Background())
}

func (i *DicomService) ToDicomServiceOutputWithContext(ctx context.Context) DicomServiceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DicomServiceOutput)
}

type DicomServiceOutput struct{ *pulumi.OutputState }

func (DicomServiceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DicomService)(nil)).Elem()
}

func (o DicomServiceOutput) ToDicomServiceOutput() DicomServiceOutput {
	return o
}

func (o DicomServiceOutput) ToDicomServiceOutputWithContext(ctx context.Context) DicomServiceOutput {
	return o
}

func (o DicomServiceOutput) AuthenticationConfiguration() DicomServiceAuthenticationConfigurationResponsePtrOutput {
	return o.ApplyT(func(v *DicomService) DicomServiceAuthenticationConfigurationResponsePtrOutput {
		return v.AuthenticationConfiguration
	}).(DicomServiceAuthenticationConfigurationResponsePtrOutput)
}

func (o DicomServiceOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DicomService) pulumi.StringPtrOutput { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o DicomServiceOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DicomService) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

func (o DicomServiceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *DicomService) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o DicomServiceOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *DicomService) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o DicomServiceOutput) ServiceUrl() pulumi.StringOutput {
	return o.ApplyT(func(v *DicomService) pulumi.StringOutput { return v.ServiceUrl }).(pulumi.StringOutput)
}

func (o DicomServiceOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *DicomService) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o DicomServiceOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *DicomService) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o DicomServiceOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *DicomService) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(DicomServiceOutput{})
}
