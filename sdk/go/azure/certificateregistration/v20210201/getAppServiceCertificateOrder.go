


package v20210201

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupAppServiceCertificateOrder(ctx *pulumi.Context, args *LookupAppServiceCertificateOrderArgs, opts ...pulumi.InvokeOption) (*LookupAppServiceCertificateOrderResult, error) {
	var rv LookupAppServiceCertificateOrderResult
	err := ctx.Invoke("azure-native:certificateregistration/v20210201:getAppServiceCertificateOrder", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupAppServiceCertificateOrderArgs struct {
	CertificateOrderName string `pulumi:"certificateOrderName"`
	ResourceGroupName    string `pulumi:"resourceGroupName"`
}


type LookupAppServiceCertificateOrderResult struct {
	AppServiceCertificateNotRenewableReasons []string                                 `pulumi:"appServiceCertificateNotRenewableReasons"`
	AutoRenew                                *bool                                    `pulumi:"autoRenew"`
	Certificates                             map[string]AppServiceCertificateResponse `pulumi:"certificates"`
	Contact                                  CertificateOrderContactResponse          `pulumi:"contact"`
	Csr                                      *string                                  `pulumi:"csr"`
	DistinguishedName                        *string                                  `pulumi:"distinguishedName"`
	DomainVerificationToken                  string                                   `pulumi:"domainVerificationToken"`
	ExpirationTime                           string                                   `pulumi:"expirationTime"`
	Id                                       string                                   `pulumi:"id"`
	Intermediate                             CertificateDetailsResponse               `pulumi:"intermediate"`
	IsPrivateKeyExternal                     bool                                     `pulumi:"isPrivateKeyExternal"`
	KeySize                                  *int                                     `pulumi:"keySize"`
	Kind                                     *string                                  `pulumi:"kind"`
	LastCertificateIssuanceTime              string                                   `pulumi:"lastCertificateIssuanceTime"`
	Location                                 string                                   `pulumi:"location"`
	Name                                     string                                   `pulumi:"name"`
	NextAutoRenewalTimeStamp                 string                                   `pulumi:"nextAutoRenewalTimeStamp"`
	ProductType                              string                                   `pulumi:"productType"`
	ProvisioningState                        string                                   `pulumi:"provisioningState"`
	Root                                     CertificateDetailsResponse               `pulumi:"root"`
	SerialNumber                             string                                   `pulumi:"serialNumber"`
	SignedCertificate                        CertificateDetailsResponse               `pulumi:"signedCertificate"`
	Status                                   string                                   `pulumi:"status"`
	Tags                                     map[string]string                        `pulumi:"tags"`
	Type                                     string                                   `pulumi:"type"`
	ValidityInYears                          *int                                     `pulumi:"validityInYears"`
}


func (val *LookupAppServiceCertificateOrderResult) Defaults() *LookupAppServiceCertificateOrderResult {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.AutoRenew) {
		autoRenew_ := true
		tmp.AutoRenew = &autoRenew_
	}
	if isZero(tmp.KeySize) {
		keySize_ := 2048
		tmp.KeySize = &keySize_
	}
	if isZero(tmp.ValidityInYears) {
		validityInYears_ := 1
		tmp.ValidityInYears = &validityInYears_
	}
	return &tmp
}

func LookupAppServiceCertificateOrderOutput(ctx *pulumi.Context, args LookupAppServiceCertificateOrderOutputArgs, opts ...pulumi.InvokeOption) LookupAppServiceCertificateOrderResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupAppServiceCertificateOrderResult, error) {
			args := v.(LookupAppServiceCertificateOrderArgs)
			r, err := LookupAppServiceCertificateOrder(ctx, &args, opts...)
			var s LookupAppServiceCertificateOrderResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupAppServiceCertificateOrderResultOutput)
}

type LookupAppServiceCertificateOrderOutputArgs struct {
	CertificateOrderName pulumi.StringInput `pulumi:"certificateOrderName"`
	ResourceGroupName    pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupAppServiceCertificateOrderOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAppServiceCertificateOrderArgs)(nil)).Elem()
}


type LookupAppServiceCertificateOrderResultOutput struct{ *pulumi.OutputState }

func (LookupAppServiceCertificateOrderResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAppServiceCertificateOrderResult)(nil)).Elem()
}

func (o LookupAppServiceCertificateOrderResultOutput) ToLookupAppServiceCertificateOrderResultOutput() LookupAppServiceCertificateOrderResultOutput {
	return o
}

func (o LookupAppServiceCertificateOrderResultOutput) ToLookupAppServiceCertificateOrderResultOutputWithContext(ctx context.Context) LookupAppServiceCertificateOrderResultOutput {
	return o
}

func (o LookupAppServiceCertificateOrderResultOutput) AppServiceCertificateNotRenewableReasons() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupAppServiceCertificateOrderResult) []string {
		return v.AppServiceCertificateNotRenewableReasons
	}).(pulumi.StringArrayOutput)
}

func (o LookupAppServiceCertificateOrderResultOutput) AutoRenew() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupAppServiceCertificateOrderResult) *bool { return v.AutoRenew }).(pulumi.BoolPtrOutput)
}

func (o LookupAppServiceCertificateOrderResultOutput) Certificates() AppServiceCertificateResponseMapOutput {
	return o.ApplyT(func(v LookupAppServiceCertificateOrderResult) map[string]AppServiceCertificateResponse {
		return v.Certificates
	}).(AppServiceCertificateResponseMapOutput)
}

func (o LookupAppServiceCertificateOrderResultOutput) Contact() CertificateOrderContactResponseOutput {
	return o.ApplyT(func(v LookupAppServiceCertificateOrderResult) CertificateOrderContactResponse { return v.Contact }).(CertificateOrderContactResponseOutput)
}

func (o LookupAppServiceCertificateOrderResultOutput) Csr() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAppServiceCertificateOrderResult) *string { return v.Csr }).(pulumi.StringPtrOutput)
}

func (o LookupAppServiceCertificateOrderResultOutput) DistinguishedName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAppServiceCertificateOrderResult) *string { return v.DistinguishedName }).(pulumi.StringPtrOutput)
}

func (o LookupAppServiceCertificateOrderResultOutput) DomainVerificationToken() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAppServiceCertificateOrderResult) string { return v.DomainVerificationToken }).(pulumi.StringOutput)
}

func (o LookupAppServiceCertificateOrderResultOutput) ExpirationTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAppServiceCertificateOrderResult) string { return v.ExpirationTime }).(pulumi.StringOutput)
}

func (o LookupAppServiceCertificateOrderResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAppServiceCertificateOrderResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupAppServiceCertificateOrderResultOutput) Intermediate() CertificateDetailsResponseOutput {
	return o.ApplyT(func(v LookupAppServiceCertificateOrderResult) CertificateDetailsResponse { return v.Intermediate }).(CertificateDetailsResponseOutput)
}

func (o LookupAppServiceCertificateOrderResultOutput) IsPrivateKeyExternal() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupAppServiceCertificateOrderResult) bool { return v.IsPrivateKeyExternal }).(pulumi.BoolOutput)
}

func (o LookupAppServiceCertificateOrderResultOutput) KeySize() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupAppServiceCertificateOrderResult) *int { return v.KeySize }).(pulumi.IntPtrOutput)
}

func (o LookupAppServiceCertificateOrderResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAppServiceCertificateOrderResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o LookupAppServiceCertificateOrderResultOutput) LastCertificateIssuanceTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAppServiceCertificateOrderResult) string { return v.LastCertificateIssuanceTime }).(pulumi.StringOutput)
}

func (o LookupAppServiceCertificateOrderResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAppServiceCertificateOrderResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupAppServiceCertificateOrderResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAppServiceCertificateOrderResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupAppServiceCertificateOrderResultOutput) NextAutoRenewalTimeStamp() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAppServiceCertificateOrderResult) string { return v.NextAutoRenewalTimeStamp }).(pulumi.StringOutput)
}

func (o LookupAppServiceCertificateOrderResultOutput) ProductType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAppServiceCertificateOrderResult) string { return v.ProductType }).(pulumi.StringOutput)
}

func (o LookupAppServiceCertificateOrderResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAppServiceCertificateOrderResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupAppServiceCertificateOrderResultOutput) Root() CertificateDetailsResponseOutput {
	return o.ApplyT(func(v LookupAppServiceCertificateOrderResult) CertificateDetailsResponse { return v.Root }).(CertificateDetailsResponseOutput)
}

func (o LookupAppServiceCertificateOrderResultOutput) SerialNumber() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAppServiceCertificateOrderResult) string { return v.SerialNumber }).(pulumi.StringOutput)
}

func (o LookupAppServiceCertificateOrderResultOutput) SignedCertificate() CertificateDetailsResponseOutput {
	return o.ApplyT(func(v LookupAppServiceCertificateOrderResult) CertificateDetailsResponse { return v.SignedCertificate }).(CertificateDetailsResponseOutput)
}

func (o LookupAppServiceCertificateOrderResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAppServiceCertificateOrderResult) string { return v.Status }).(pulumi.StringOutput)
}

func (o LookupAppServiceCertificateOrderResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupAppServiceCertificateOrderResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupAppServiceCertificateOrderResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAppServiceCertificateOrderResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupAppServiceCertificateOrderResultOutput) ValidityInYears() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupAppServiceCertificateOrderResult) *int { return v.ValidityInYears }).(pulumi.IntPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupAppServiceCertificateOrderResultOutput{})
}
