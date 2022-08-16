


package m365securityandcompliance

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type PrivateLinkServicesForEDMUpload struct {
	pulumi.CustomResourceState

	Etag       pulumi.StringPtrOutput                    `pulumi:"etag"`
	Identity   ServicesResourceResponseIdentityPtrOutput `pulumi:"identity"`
	Kind       pulumi.StringOutput                       `pulumi:"kind"`
	Location   pulumi.StringOutput                       `pulumi:"location"`
	Name       pulumi.StringOutput                       `pulumi:"name"`
	Properties ServicesPropertiesResponseOutput          `pulumi:"properties"`
	SystemData SystemDataResponseOutput                  `pulumi:"systemData"`
	Tags       pulumi.StringMapOutput                    `pulumi:"tags"`
	Type       pulumi.StringOutput                       `pulumi:"type"`
}


func NewPrivateLinkServicesForEDMUpload(ctx *pulumi.Context,
	name string, args *PrivateLinkServicesForEDMUploadArgs, opts ...pulumi.ResourceOption) (*PrivateLinkServicesForEDMUpload, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Kind == nil {
		return nil, errors.New("invalid value for required argument 'Kind'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:m365securityandcompliance/v20210325preview:privateLinkServicesForEDMUpload"),
		},
	})
	opts = append(opts, aliases)
	var resource PrivateLinkServicesForEDMUpload
	err := ctx.RegisterResource("azure-native:m365securityandcompliance:privateLinkServicesForEDMUpload", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetPrivateLinkServicesForEDMUpload(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PrivateLinkServicesForEDMUploadState, opts ...pulumi.ResourceOption) (*PrivateLinkServicesForEDMUpload, error) {
	var resource PrivateLinkServicesForEDMUpload
	err := ctx.ReadResource("azure-native:m365securityandcompliance:privateLinkServicesForEDMUpload", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type privateLinkServicesForEDMUploadState struct {
}

type PrivateLinkServicesForEDMUploadState struct {
}

func (PrivateLinkServicesForEDMUploadState) ElementType() reflect.Type {
	return reflect.TypeOf((*privateLinkServicesForEDMUploadState)(nil)).Elem()
}

type privateLinkServicesForEDMUploadArgs struct {
	Identity          *ServicesResourceIdentity `pulumi:"identity"`
	Kind              Kind                      `pulumi:"kind"`
	Location          *string                   `pulumi:"location"`
	Properties        *ServicesProperties       `pulumi:"properties"`
	ResourceGroupName string                    `pulumi:"resourceGroupName"`
	ResourceName      *string                   `pulumi:"resourceName"`
	Tags              map[string]string         `pulumi:"tags"`
}


type PrivateLinkServicesForEDMUploadArgs struct {
	Identity          ServicesResourceIdentityPtrInput
	Kind              KindInput
	Location          pulumi.StringPtrInput
	Properties        ServicesPropertiesPtrInput
	ResourceGroupName pulumi.StringInput
	ResourceName      pulumi.StringPtrInput
	Tags              pulumi.StringMapInput
}

func (PrivateLinkServicesForEDMUploadArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*privateLinkServicesForEDMUploadArgs)(nil)).Elem()
}

type PrivateLinkServicesForEDMUploadInput interface {
	pulumi.Input

	ToPrivateLinkServicesForEDMUploadOutput() PrivateLinkServicesForEDMUploadOutput
	ToPrivateLinkServicesForEDMUploadOutputWithContext(ctx context.Context) PrivateLinkServicesForEDMUploadOutput
}

func (*PrivateLinkServicesForEDMUpload) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateLinkServicesForEDMUpload)(nil)).Elem()
}

func (i *PrivateLinkServicesForEDMUpload) ToPrivateLinkServicesForEDMUploadOutput() PrivateLinkServicesForEDMUploadOutput {
	return i.ToPrivateLinkServicesForEDMUploadOutputWithContext(context.Background())
}

func (i *PrivateLinkServicesForEDMUpload) ToPrivateLinkServicesForEDMUploadOutputWithContext(ctx context.Context) PrivateLinkServicesForEDMUploadOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateLinkServicesForEDMUploadOutput)
}

type PrivateLinkServicesForEDMUploadOutput struct{ *pulumi.OutputState }

func (PrivateLinkServicesForEDMUploadOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateLinkServicesForEDMUpload)(nil)).Elem()
}

func (o PrivateLinkServicesForEDMUploadOutput) ToPrivateLinkServicesForEDMUploadOutput() PrivateLinkServicesForEDMUploadOutput {
	return o
}

func (o PrivateLinkServicesForEDMUploadOutput) ToPrivateLinkServicesForEDMUploadOutputWithContext(ctx context.Context) PrivateLinkServicesForEDMUploadOutput {
	return o
}

func (o PrivateLinkServicesForEDMUploadOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateLinkServicesForEDMUpload) pulumi.StringPtrOutput { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o PrivateLinkServicesForEDMUploadOutput) Identity() ServicesResourceResponseIdentityPtrOutput {
	return o.ApplyT(func(v *PrivateLinkServicesForEDMUpload) ServicesResourceResponseIdentityPtrOutput { return v.Identity }).(ServicesResourceResponseIdentityPtrOutput)
}

func (o PrivateLinkServicesForEDMUploadOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateLinkServicesForEDMUpload) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

func (o PrivateLinkServicesForEDMUploadOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateLinkServicesForEDMUpload) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o PrivateLinkServicesForEDMUploadOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateLinkServicesForEDMUpload) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o PrivateLinkServicesForEDMUploadOutput) Properties() ServicesPropertiesResponseOutput {
	return o.ApplyT(func(v *PrivateLinkServicesForEDMUpload) ServicesPropertiesResponseOutput { return v.Properties }).(ServicesPropertiesResponseOutput)
}

func (o PrivateLinkServicesForEDMUploadOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *PrivateLinkServicesForEDMUpload) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o PrivateLinkServicesForEDMUploadOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *PrivateLinkServicesForEDMUpload) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o PrivateLinkServicesForEDMUploadOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateLinkServicesForEDMUpload) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(PrivateLinkServicesForEDMUploadOutput{})
}
