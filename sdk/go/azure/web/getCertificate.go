


package web

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupCertificate(ctx *pulumi.Context, args *LookupCertificateArgs, opts ...pulumi.InvokeOption) (*LookupCertificateResult, error) {
	var rv LookupCertificateResult
	err := ctx.Invoke("azure-native:web:getCertificate", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupCertificateArgs struct {
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupCertificateResult struct {
	CanonicalName             *string                           `pulumi:"canonicalName"`
	CerBlob                   string                            `pulumi:"cerBlob"`
	DomainValidationMethod    *string                           `pulumi:"domainValidationMethod"`
	ExpirationDate            string                            `pulumi:"expirationDate"`
	FriendlyName              string                            `pulumi:"friendlyName"`
	HostNames                 []string                          `pulumi:"hostNames"`
	HostingEnvironmentProfile HostingEnvironmentProfileResponse `pulumi:"hostingEnvironmentProfile"`
	Id                        string                            `pulumi:"id"`
	IssueDate                 string                            `pulumi:"issueDate"`
	Issuer                    string                            `pulumi:"issuer"`
	KeyVaultId                *string                           `pulumi:"keyVaultId"`
	KeyVaultSecretName        *string                           `pulumi:"keyVaultSecretName"`
	KeyVaultSecretStatus      string                            `pulumi:"keyVaultSecretStatus"`
	Kind                      *string                           `pulumi:"kind"`
	Location                  string                            `pulumi:"location"`
	Name                      string                            `pulumi:"name"`
	PfxBlob                   *string                           `pulumi:"pfxBlob"`
	PublicKeyHash             string                            `pulumi:"publicKeyHash"`
	SelfLink                  string                            `pulumi:"selfLink"`
	ServerFarmId              *string                           `pulumi:"serverFarmId"`
	SiteName                  string                            `pulumi:"siteName"`
	SubjectName               string                            `pulumi:"subjectName"`
	Tags                      map[string]string                 `pulumi:"tags"`
	Thumbprint                string                            `pulumi:"thumbprint"`
	Type                      string                            `pulumi:"type"`
	Valid                     bool                              `pulumi:"valid"`
}

func LookupCertificateOutput(ctx *pulumi.Context, args LookupCertificateOutputArgs, opts ...pulumi.InvokeOption) LookupCertificateResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupCertificateResult, error) {
			args := v.(LookupCertificateArgs)
			r, err := LookupCertificate(ctx, &args, opts...)
			var s LookupCertificateResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupCertificateResultOutput)
}

type LookupCertificateOutputArgs struct {
	Name              pulumi.StringInput `pulumi:"name"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupCertificateOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCertificateArgs)(nil)).Elem()
}


type LookupCertificateResultOutput struct{ *pulumi.OutputState }

func (LookupCertificateResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCertificateResult)(nil)).Elem()
}

func (o LookupCertificateResultOutput) ToLookupCertificateResultOutput() LookupCertificateResultOutput {
	return o
}

func (o LookupCertificateResultOutput) ToLookupCertificateResultOutputWithContext(ctx context.Context) LookupCertificateResultOutput {
	return o
}

func (o LookupCertificateResultOutput) CanonicalName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCertificateResult) *string { return v.CanonicalName }).(pulumi.StringPtrOutput)
}

func (o LookupCertificateResultOutput) CerBlob() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCertificateResult) string { return v.CerBlob }).(pulumi.StringOutput)
}

func (o LookupCertificateResultOutput) DomainValidationMethod() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCertificateResult) *string { return v.DomainValidationMethod }).(pulumi.StringPtrOutput)
}

func (o LookupCertificateResultOutput) ExpirationDate() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCertificateResult) string { return v.ExpirationDate }).(pulumi.StringOutput)
}

func (o LookupCertificateResultOutput) FriendlyName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCertificateResult) string { return v.FriendlyName }).(pulumi.StringOutput)
}

func (o LookupCertificateResultOutput) HostNames() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupCertificateResult) []string { return v.HostNames }).(pulumi.StringArrayOutput)
}

func (o LookupCertificateResultOutput) HostingEnvironmentProfile() HostingEnvironmentProfileResponseOutput {
	return o.ApplyT(func(v LookupCertificateResult) HostingEnvironmentProfileResponse { return v.HostingEnvironmentProfile }).(HostingEnvironmentProfileResponseOutput)
}

func (o LookupCertificateResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCertificateResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupCertificateResultOutput) IssueDate() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCertificateResult) string { return v.IssueDate }).(pulumi.StringOutput)
}

func (o LookupCertificateResultOutput) Issuer() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCertificateResult) string { return v.Issuer }).(pulumi.StringOutput)
}

func (o LookupCertificateResultOutput) KeyVaultId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCertificateResult) *string { return v.KeyVaultId }).(pulumi.StringPtrOutput)
}

func (o LookupCertificateResultOutput) KeyVaultSecretName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCertificateResult) *string { return v.KeyVaultSecretName }).(pulumi.StringPtrOutput)
}

func (o LookupCertificateResultOutput) KeyVaultSecretStatus() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCertificateResult) string { return v.KeyVaultSecretStatus }).(pulumi.StringOutput)
}

func (o LookupCertificateResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCertificateResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o LookupCertificateResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCertificateResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupCertificateResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCertificateResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupCertificateResultOutput) PfxBlob() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCertificateResult) *string { return v.PfxBlob }).(pulumi.StringPtrOutput)
}

func (o LookupCertificateResultOutput) PublicKeyHash() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCertificateResult) string { return v.PublicKeyHash }).(pulumi.StringOutput)
}

func (o LookupCertificateResultOutput) SelfLink() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCertificateResult) string { return v.SelfLink }).(pulumi.StringOutput)
}

func (o LookupCertificateResultOutput) ServerFarmId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCertificateResult) *string { return v.ServerFarmId }).(pulumi.StringPtrOutput)
}

func (o LookupCertificateResultOutput) SiteName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCertificateResult) string { return v.SiteName }).(pulumi.StringOutput)
}

func (o LookupCertificateResultOutput) SubjectName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCertificateResult) string { return v.SubjectName }).(pulumi.StringOutput)
}

func (o LookupCertificateResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupCertificateResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupCertificateResultOutput) Thumbprint() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCertificateResult) string { return v.Thumbprint }).(pulumi.StringOutput)
}

func (o LookupCertificateResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCertificateResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupCertificateResultOutput) Valid() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupCertificateResult) bool { return v.Valid }).(pulumi.BoolOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupCertificateResultOutput{})
}
