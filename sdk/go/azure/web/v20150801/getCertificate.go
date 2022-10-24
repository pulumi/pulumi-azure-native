


package v20150801

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


func LookupCertificate(ctx *pulumi.Context, args *LookupCertificateArgs, opts ...pulumi.InvokeOption) (*LookupCertificateResult, error) {
	var rv LookupCertificateResult
	err := ctx.Invoke("azure-native:web/v20150801:getCertificate", args, &rv, opts...)
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
	CerBlob                   *string                            `pulumi:"cerBlob"`
	ExpirationDate            *string                            `pulumi:"expirationDate"`
	FriendlyName              *string                            `pulumi:"friendlyName"`
	HostNames                 []string                           `pulumi:"hostNames"`
	HostingEnvironmentProfile *HostingEnvironmentProfileResponse `pulumi:"hostingEnvironmentProfile"`
	Id                        *string                            `pulumi:"id"`
	IssueDate                 *string                            `pulumi:"issueDate"`
	Issuer                    *string                            `pulumi:"issuer"`
	Kind                      *string                            `pulumi:"kind"`
	Location                  string                             `pulumi:"location"`
	Name                      *string                            `pulumi:"name"`
	Password                  *string                            `pulumi:"password"`
	PfxBlob                   *string                            `pulumi:"pfxBlob"`
	PublicKeyHash             *string                            `pulumi:"publicKeyHash"`
	SelfLink                  *string                            `pulumi:"selfLink"`
	SiteName                  *string                            `pulumi:"siteName"`
	SubjectName               *string                            `pulumi:"subjectName"`
	Tags                      map[string]string                  `pulumi:"tags"`
	Thumbprint                *string                            `pulumi:"thumbprint"`
	Type                      *string                            `pulumi:"type"`
	Valid                     *bool                              `pulumi:"valid"`
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

func (o LookupCertificateResultOutput) CerBlob() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCertificateResult) *string { return v.CerBlob }).(pulumi.StringPtrOutput)
}

func (o LookupCertificateResultOutput) ExpirationDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCertificateResult) *string { return v.ExpirationDate }).(pulumi.StringPtrOutput)
}

func (o LookupCertificateResultOutput) FriendlyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCertificateResult) *string { return v.FriendlyName }).(pulumi.StringPtrOutput)
}

func (o LookupCertificateResultOutput) HostNames() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupCertificateResult) []string { return v.HostNames }).(pulumi.StringArrayOutput)
}

func (o LookupCertificateResultOutput) HostingEnvironmentProfile() HostingEnvironmentProfileResponsePtrOutput {
	return o.ApplyT(func(v LookupCertificateResult) *HostingEnvironmentProfileResponse { return v.HostingEnvironmentProfile }).(HostingEnvironmentProfileResponsePtrOutput)
}

func (o LookupCertificateResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCertificateResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupCertificateResultOutput) IssueDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCertificateResult) *string { return v.IssueDate }).(pulumi.StringPtrOutput)
}

func (o LookupCertificateResultOutput) Issuer() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCertificateResult) *string { return v.Issuer }).(pulumi.StringPtrOutput)
}

func (o LookupCertificateResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCertificateResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o LookupCertificateResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCertificateResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupCertificateResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCertificateResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o LookupCertificateResultOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCertificateResult) *string { return v.Password }).(pulumi.StringPtrOutput)
}

func (o LookupCertificateResultOutput) PfxBlob() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCertificateResult) *string { return v.PfxBlob }).(pulumi.StringPtrOutput)
}

func (o LookupCertificateResultOutput) PublicKeyHash() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCertificateResult) *string { return v.PublicKeyHash }).(pulumi.StringPtrOutput)
}

func (o LookupCertificateResultOutput) SelfLink() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCertificateResult) *string { return v.SelfLink }).(pulumi.StringPtrOutput)
}

func (o LookupCertificateResultOutput) SiteName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCertificateResult) *string { return v.SiteName }).(pulumi.StringPtrOutput)
}

func (o LookupCertificateResultOutput) SubjectName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCertificateResult) *string { return v.SubjectName }).(pulumi.StringPtrOutput)
}

func (o LookupCertificateResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupCertificateResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupCertificateResultOutput) Thumbprint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCertificateResult) *string { return v.Thumbprint }).(pulumi.StringPtrOutput)
}

func (o LookupCertificateResultOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCertificateResult) *string { return v.Type }).(pulumi.StringPtrOutput)
}

func (o LookupCertificateResultOutput) Valid() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupCertificateResult) *bool { return v.Valid }).(pulumi.BoolPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupCertificateResultOutput{})
}
