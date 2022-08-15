


package v20190801

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


type AppServiceCertificateOrder struct {
	pulumi.CustomResourceState

	AppServiceCertificateNotRenewableReasons pulumi.StringArrayOutput               `pulumi:"appServiceCertificateNotRenewableReasons"`
	AutoRenew                                pulumi.BoolPtrOutput                   `pulumi:"autoRenew"`
	Certificates                             AppServiceCertificateResponseMapOutput `pulumi:"certificates"`
	Csr                                      pulumi.StringPtrOutput                 `pulumi:"csr"`
	DistinguishedName                        pulumi.StringPtrOutput                 `pulumi:"distinguishedName"`
	DomainVerificationToken                  pulumi.StringOutput                    `pulumi:"domainVerificationToken"`
	ExpirationTime                           pulumi.StringOutput                    `pulumi:"expirationTime"`
	Intermediate                             CertificateDetailsResponseOutput       `pulumi:"intermediate"`
	IsPrivateKeyExternal                     pulumi.BoolOutput                      `pulumi:"isPrivateKeyExternal"`
	KeySize                                  pulumi.IntPtrOutput                    `pulumi:"keySize"`
	Kind                                     pulumi.StringPtrOutput                 `pulumi:"kind"`
	LastCertificateIssuanceTime              pulumi.StringOutput                    `pulumi:"lastCertificateIssuanceTime"`
	Location                                 pulumi.StringOutput                    `pulumi:"location"`
	Name                                     pulumi.StringOutput                    `pulumi:"name"`
	NextAutoRenewalTimeStamp                 pulumi.StringOutput                    `pulumi:"nextAutoRenewalTimeStamp"`
	ProductType                              pulumi.StringOutput                    `pulumi:"productType"`
	ProvisioningState                        pulumi.StringOutput                    `pulumi:"provisioningState"`
	Root                                     CertificateDetailsResponseOutput       `pulumi:"root"`
	SerialNumber                             pulumi.StringOutput                    `pulumi:"serialNumber"`
	SignedCertificate                        CertificateDetailsResponseOutput       `pulumi:"signedCertificate"`
	Status                                   pulumi.StringOutput                    `pulumi:"status"`
	Tags                                     pulumi.StringMapOutput                 `pulumi:"tags"`
	Type                                     pulumi.StringOutput                    `pulumi:"type"`
	ValidityInYears                          pulumi.IntPtrOutput                    `pulumi:"validityInYears"`
}


func NewAppServiceCertificateOrder(ctx *pulumi.Context,
	name string, args *AppServiceCertificateOrderArgs, opts ...pulumi.ResourceOption) (*AppServiceCertificateOrder, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ProductType == nil {
		return nil, errors.New("invalid value for required argument 'ProductType'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if isZero(args.AutoRenew) {
		args.AutoRenew = pulumi.BoolPtr(true)
	}
	if isZero(args.KeySize) {
		args.KeySize = pulumi.IntPtr(2048)
	}
	if isZero(args.ValidityInYears) {
		args.ValidityInYears = pulumi.IntPtr(1)
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:certificateregistration:AppServiceCertificateOrder"),
		},
		{
			Type: pulumi.String("azure-native:certificateregistration/v20150801:AppServiceCertificateOrder"),
		},
		{
			Type: pulumi.String("azure-native:certificateregistration/v20180201:AppServiceCertificateOrder"),
		},
		{
			Type: pulumi.String("azure-native:certificateregistration/v20200601:AppServiceCertificateOrder"),
		},
		{
			Type: pulumi.String("azure-native:certificateregistration/v20200901:AppServiceCertificateOrder"),
		},
		{
			Type: pulumi.String("azure-native:certificateregistration/v20201001:AppServiceCertificateOrder"),
		},
		{
			Type: pulumi.String("azure-native:certificateregistration/v20201201:AppServiceCertificateOrder"),
		},
		{
			Type: pulumi.String("azure-native:certificateregistration/v20210101:AppServiceCertificateOrder"),
		},
		{
			Type: pulumi.String("azure-native:certificateregistration/v20210115:AppServiceCertificateOrder"),
		},
		{
			Type: pulumi.String("azure-native:certificateregistration/v20210201:AppServiceCertificateOrder"),
		},
		{
			Type: pulumi.String("azure-native:certificateregistration/v20210301:AppServiceCertificateOrder"),
		},
		{
			Type: pulumi.String("azure-native:certificateregistration/v20220301:AppServiceCertificateOrder"),
		},
	})
	opts = append(opts, aliases)
	var resource AppServiceCertificateOrder
	err := ctx.RegisterResource("azure-native:certificateregistration/v20190801:AppServiceCertificateOrder", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetAppServiceCertificateOrder(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AppServiceCertificateOrderState, opts ...pulumi.ResourceOption) (*AppServiceCertificateOrder, error) {
	var resource AppServiceCertificateOrder
	err := ctx.ReadResource("azure-native:certificateregistration/v20190801:AppServiceCertificateOrder", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type appServiceCertificateOrderState struct {
}

type AppServiceCertificateOrderState struct {
}

func (AppServiceCertificateOrderState) ElementType() reflect.Type {
	return reflect.TypeOf((*appServiceCertificateOrderState)(nil)).Elem()
}

type appServiceCertificateOrderArgs struct {
	AutoRenew            *bool                            `pulumi:"autoRenew"`
	CertificateOrderName *string                          `pulumi:"certificateOrderName"`
	Certificates         map[string]AppServiceCertificate `pulumi:"certificates"`
	Csr                  *string                          `pulumi:"csr"`
	DistinguishedName    *string                          `pulumi:"distinguishedName"`
	KeySize              *int                             `pulumi:"keySize"`
	Kind                 *string                          `pulumi:"kind"`
	Location             *string                          `pulumi:"location"`
	ProductType          CertificateProductType           `pulumi:"productType"`
	ResourceGroupName    string                           `pulumi:"resourceGroupName"`
	Tags                 map[string]string                `pulumi:"tags"`
	ValidityInYears      *int                             `pulumi:"validityInYears"`
}


type AppServiceCertificateOrderArgs struct {
	AutoRenew            pulumi.BoolPtrInput
	CertificateOrderName pulumi.StringPtrInput
	Certificates         AppServiceCertificateMapInput
	Csr                  pulumi.StringPtrInput
	DistinguishedName    pulumi.StringPtrInput
	KeySize              pulumi.IntPtrInput
	Kind                 pulumi.StringPtrInput
	Location             pulumi.StringPtrInput
	ProductType          CertificateProductTypeInput
	ResourceGroupName    pulumi.StringInput
	Tags                 pulumi.StringMapInput
	ValidityInYears      pulumi.IntPtrInput
}

func (AppServiceCertificateOrderArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*appServiceCertificateOrderArgs)(nil)).Elem()
}

type AppServiceCertificateOrderInput interface {
	pulumi.Input

	ToAppServiceCertificateOrderOutput() AppServiceCertificateOrderOutput
	ToAppServiceCertificateOrderOutputWithContext(ctx context.Context) AppServiceCertificateOrderOutput
}

func (*AppServiceCertificateOrder) ElementType() reflect.Type {
	return reflect.TypeOf((**AppServiceCertificateOrder)(nil)).Elem()
}

func (i *AppServiceCertificateOrder) ToAppServiceCertificateOrderOutput() AppServiceCertificateOrderOutput {
	return i.ToAppServiceCertificateOrderOutputWithContext(context.Background())
}

func (i *AppServiceCertificateOrder) ToAppServiceCertificateOrderOutputWithContext(ctx context.Context) AppServiceCertificateOrderOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppServiceCertificateOrderOutput)
}

type AppServiceCertificateOrderOutput struct{ *pulumi.OutputState }

func (AppServiceCertificateOrderOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AppServiceCertificateOrder)(nil)).Elem()
}

func (o AppServiceCertificateOrderOutput) ToAppServiceCertificateOrderOutput() AppServiceCertificateOrderOutput {
	return o
}

func (o AppServiceCertificateOrderOutput) ToAppServiceCertificateOrderOutputWithContext(ctx context.Context) AppServiceCertificateOrderOutput {
	return o
}

func (o AppServiceCertificateOrderOutput) AppServiceCertificateNotRenewableReasons() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *AppServiceCertificateOrder) pulumi.StringArrayOutput {
		return v.AppServiceCertificateNotRenewableReasons
	}).(pulumi.StringArrayOutput)
}

func (o AppServiceCertificateOrderOutput) AutoRenew() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AppServiceCertificateOrder) pulumi.BoolPtrOutput { return v.AutoRenew }).(pulumi.BoolPtrOutput)
}

func (o AppServiceCertificateOrderOutput) Certificates() AppServiceCertificateResponseMapOutput {
	return o.ApplyT(func(v *AppServiceCertificateOrder) AppServiceCertificateResponseMapOutput { return v.Certificates }).(AppServiceCertificateResponseMapOutput)
}

func (o AppServiceCertificateOrderOutput) Csr() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AppServiceCertificateOrder) pulumi.StringPtrOutput { return v.Csr }).(pulumi.StringPtrOutput)
}

func (o AppServiceCertificateOrderOutput) DistinguishedName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AppServiceCertificateOrder) pulumi.StringPtrOutput { return v.DistinguishedName }).(pulumi.StringPtrOutput)
}

func (o AppServiceCertificateOrderOutput) DomainVerificationToken() pulumi.StringOutput {
	return o.ApplyT(func(v *AppServiceCertificateOrder) pulumi.StringOutput { return v.DomainVerificationToken }).(pulumi.StringOutput)
}

func (o AppServiceCertificateOrderOutput) ExpirationTime() pulumi.StringOutput {
	return o.ApplyT(func(v *AppServiceCertificateOrder) pulumi.StringOutput { return v.ExpirationTime }).(pulumi.StringOutput)
}

func (o AppServiceCertificateOrderOutput) Intermediate() CertificateDetailsResponseOutput {
	return o.ApplyT(func(v *AppServiceCertificateOrder) CertificateDetailsResponseOutput { return v.Intermediate }).(CertificateDetailsResponseOutput)
}

func (o AppServiceCertificateOrderOutput) IsPrivateKeyExternal() pulumi.BoolOutput {
	return o.ApplyT(func(v *AppServiceCertificateOrder) pulumi.BoolOutput { return v.IsPrivateKeyExternal }).(pulumi.BoolOutput)
}

func (o AppServiceCertificateOrderOutput) KeySize() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *AppServiceCertificateOrder) pulumi.IntPtrOutput { return v.KeySize }).(pulumi.IntPtrOutput)
}

func (o AppServiceCertificateOrderOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AppServiceCertificateOrder) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o AppServiceCertificateOrderOutput) LastCertificateIssuanceTime() pulumi.StringOutput {
	return o.ApplyT(func(v *AppServiceCertificateOrder) pulumi.StringOutput { return v.LastCertificateIssuanceTime }).(pulumi.StringOutput)
}

func (o AppServiceCertificateOrderOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *AppServiceCertificateOrder) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o AppServiceCertificateOrderOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *AppServiceCertificateOrder) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o AppServiceCertificateOrderOutput) NextAutoRenewalTimeStamp() pulumi.StringOutput {
	return o.ApplyT(func(v *AppServiceCertificateOrder) pulumi.StringOutput { return v.NextAutoRenewalTimeStamp }).(pulumi.StringOutput)
}

func (o AppServiceCertificateOrderOutput) ProductType() pulumi.StringOutput {
	return o.ApplyT(func(v *AppServiceCertificateOrder) pulumi.StringOutput { return v.ProductType }).(pulumi.StringOutput)
}

func (o AppServiceCertificateOrderOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *AppServiceCertificateOrder) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o AppServiceCertificateOrderOutput) Root() CertificateDetailsResponseOutput {
	return o.ApplyT(func(v *AppServiceCertificateOrder) CertificateDetailsResponseOutput { return v.Root }).(CertificateDetailsResponseOutput)
}

func (o AppServiceCertificateOrderOutput) SerialNumber() pulumi.StringOutput {
	return o.ApplyT(func(v *AppServiceCertificateOrder) pulumi.StringOutput { return v.SerialNumber }).(pulumi.StringOutput)
}

func (o AppServiceCertificateOrderOutput) SignedCertificate() CertificateDetailsResponseOutput {
	return o.ApplyT(func(v *AppServiceCertificateOrder) CertificateDetailsResponseOutput { return v.SignedCertificate }).(CertificateDetailsResponseOutput)
}

func (o AppServiceCertificateOrderOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *AppServiceCertificateOrder) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

func (o AppServiceCertificateOrderOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *AppServiceCertificateOrder) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o AppServiceCertificateOrderOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *AppServiceCertificateOrder) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o AppServiceCertificateOrderOutput) ValidityInYears() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *AppServiceCertificateOrder) pulumi.IntPtrOutput { return v.ValidityInYears }).(pulumi.IntPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(AppServiceCertificateOrderOutput{})
}
