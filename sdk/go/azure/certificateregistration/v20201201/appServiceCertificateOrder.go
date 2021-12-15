


package v20201201

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
	Contact                                  CertificateOrderContactResponseOutput  `pulumi:"contact"`
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
			Type: pulumi.String("azure-native:certificateregistration/v20190801:AppServiceCertificateOrder"),
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
			Type: pulumi.String("azure-native:certificateregistration/v20210101:AppServiceCertificateOrder"),
		},
		{
			Type: pulumi.String("azure-native:certificateregistration/v20210115:AppServiceCertificateOrder"),
		},
		{
			Type: pulumi.String("azure-native:certificateregistration/v20210201:AppServiceCertificateOrder"),
		},
	})
	opts = append(opts, aliases)
	var resource AppServiceCertificateOrder
	err := ctx.RegisterResource("azure-native:certificateregistration/v20201201:AppServiceCertificateOrder", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetAppServiceCertificateOrder(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AppServiceCertificateOrderState, opts ...pulumi.ResourceOption) (*AppServiceCertificateOrder, error) {
	var resource AppServiceCertificateOrder
	err := ctx.ReadResource("azure-native:certificateregistration/v20201201:AppServiceCertificateOrder", name, id, state, &resource, opts...)
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

func init() {
	pulumi.RegisterOutputType(AppServiceCertificateOrderOutput{})
}
