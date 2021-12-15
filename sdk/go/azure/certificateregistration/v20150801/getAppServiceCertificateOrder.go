


package v20150801

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupAppServiceCertificateOrder(ctx *pulumi.Context, args *LookupAppServiceCertificateOrderArgs, opts ...pulumi.InvokeOption) (*LookupAppServiceCertificateOrderResult, error) {
	var rv LookupAppServiceCertificateOrderResult
	err := ctx.Invoke("azure-native:certificateregistration/v20150801:getAppServiceCertificateOrder", args, &rv, opts...)
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
