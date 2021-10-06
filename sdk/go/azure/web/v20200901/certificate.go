


package v20200901

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Certificate struct {
	pulumi.CustomResourceState

	CanonicalName             pulumi.StringPtrOutput                  `pulumi:"canonicalName"`
	CerBlob                   pulumi.StringOutput                     `pulumi:"cerBlob"`
	ExpirationDate            pulumi.StringOutput                     `pulumi:"expirationDate"`
	FriendlyName              pulumi.StringOutput                     `pulumi:"friendlyName"`
	HostNames                 pulumi.StringArrayOutput                `pulumi:"hostNames"`
	HostingEnvironmentProfile HostingEnvironmentProfileResponseOutput `pulumi:"hostingEnvironmentProfile"`
	IssueDate                 pulumi.StringOutput                     `pulumi:"issueDate"`
	Issuer                    pulumi.StringOutput                     `pulumi:"issuer"`
	KeyVaultId                pulumi.StringPtrOutput                  `pulumi:"keyVaultId"`
	KeyVaultSecretName        pulumi.StringPtrOutput                  `pulumi:"keyVaultSecretName"`
	KeyVaultSecretStatus      pulumi.StringOutput                     `pulumi:"keyVaultSecretStatus"`
	Kind                      pulumi.StringPtrOutput                  `pulumi:"kind"`
	Location                  pulumi.StringOutput                     `pulumi:"location"`
	Name                      pulumi.StringOutput                     `pulumi:"name"`
	PfxBlob                   pulumi.StringPtrOutput                  `pulumi:"pfxBlob"`
	PublicKeyHash             pulumi.StringOutput                     `pulumi:"publicKeyHash"`
	SelfLink                  pulumi.StringOutput                     `pulumi:"selfLink"`
	ServerFarmId              pulumi.StringPtrOutput                  `pulumi:"serverFarmId"`
	SiteName                  pulumi.StringOutput                     `pulumi:"siteName"`
	SubjectName               pulumi.StringOutput                     `pulumi:"subjectName"`
	SystemData                SystemDataResponseOutput                `pulumi:"systemData"`
	Tags                      pulumi.StringMapOutput                  `pulumi:"tags"`
	Thumbprint                pulumi.StringOutput                     `pulumi:"thumbprint"`
	Type                      pulumi.StringOutput                     `pulumi:"type"`
	Valid                     pulumi.BoolOutput                       `pulumi:"valid"`
}


func NewCertificate(ctx *pulumi.Context,
	name string, args *CertificateArgs, opts ...pulumi.ResourceOption) (*Certificate, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Password == nil {
		return nil, errors.New("invalid value for required argument 'Password'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:web/v20200901:Certificate"),
		},
		{
			Type: pulumi.String("azure-native:web:Certificate"),
		},
		{
			Type: pulumi.String("azure-nextgen:web:Certificate"),
		},
		{
			Type: pulumi.String("azure-native:web/v20150801:Certificate"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20150801:Certificate"),
		},
		{
			Type: pulumi.String("azure-native:web/v20160301:Certificate"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20160301:Certificate"),
		},
		{
			Type: pulumi.String("azure-native:web/v20180201:Certificate"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20180201:Certificate"),
		},
		{
			Type: pulumi.String("azure-native:web/v20181101:Certificate"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20181101:Certificate"),
		},
		{
			Type: pulumi.String("azure-native:web/v20190801:Certificate"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20190801:Certificate"),
		},
		{
			Type: pulumi.String("azure-native:web/v20200601:Certificate"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20200601:Certificate"),
		},
		{
			Type: pulumi.String("azure-native:web/v20201001:Certificate"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20201001:Certificate"),
		},
		{
			Type: pulumi.String("azure-native:web/v20201201:Certificate"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20201201:Certificate"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210101:Certificate"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20210101:Certificate"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210115:Certificate"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20210115:Certificate"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210201:Certificate"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20210201:Certificate"),
		},
	})
	opts = append(opts, aliases)
	var resource Certificate
	err := ctx.RegisterResource("azure-native:web/v20200901:Certificate", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetCertificate(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CertificateState, opts ...pulumi.ResourceOption) (*Certificate, error) {
	var resource Certificate
	err := ctx.ReadResource("azure-native:web/v20200901:Certificate", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type certificateState struct {
}

type CertificateState struct {
}

func (CertificateState) ElementType() reflect.Type {
	return reflect.TypeOf((*certificateState)(nil)).Elem()
}

type certificateArgs struct {
	CanonicalName      *string           `pulumi:"canonicalName"`
	HostNames          []string          `pulumi:"hostNames"`
	KeyVaultId         *string           `pulumi:"keyVaultId"`
	KeyVaultSecretName *string           `pulumi:"keyVaultSecretName"`
	Kind               *string           `pulumi:"kind"`
	Location           *string           `pulumi:"location"`
	Name               *string           `pulumi:"name"`
	Password           string            `pulumi:"password"`
	PfxBlob            *string           `pulumi:"pfxBlob"`
	ResourceGroupName  string            `pulumi:"resourceGroupName"`
	ServerFarmId       *string           `pulumi:"serverFarmId"`
	Tags               map[string]string `pulumi:"tags"`
}


type CertificateArgs struct {
	CanonicalName      pulumi.StringPtrInput
	HostNames          pulumi.StringArrayInput
	KeyVaultId         pulumi.StringPtrInput
	KeyVaultSecretName pulumi.StringPtrInput
	Kind               pulumi.StringPtrInput
	Location           pulumi.StringPtrInput
	Name               pulumi.StringPtrInput
	Password           pulumi.StringInput
	PfxBlob            pulumi.StringPtrInput
	ResourceGroupName  pulumi.StringInput
	ServerFarmId       pulumi.StringPtrInput
	Tags               pulumi.StringMapInput
}

func (CertificateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*certificateArgs)(nil)).Elem()
}

type CertificateInput interface {
	pulumi.Input

	ToCertificateOutput() CertificateOutput
	ToCertificateOutputWithContext(ctx context.Context) CertificateOutput
}

func (*Certificate) ElementType() reflect.Type {
	return reflect.TypeOf((*Certificate)(nil))
}

func (i *Certificate) ToCertificateOutput() CertificateOutput {
	return i.ToCertificateOutputWithContext(context.Background())
}

func (i *Certificate) ToCertificateOutputWithContext(ctx context.Context) CertificateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CertificateOutput)
}

type CertificateOutput struct{ *pulumi.OutputState }

func (CertificateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Certificate)(nil))
}

func (o CertificateOutput) ToCertificateOutput() CertificateOutput {
	return o
}

func (o CertificateOutput) ToCertificateOutputWithContext(ctx context.Context) CertificateOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(CertificateOutput{})
}
