


package v20201201

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupDomain(ctx *pulumi.Context, args *LookupDomainArgs, opts ...pulumi.InvokeOption) (*LookupDomainResult, error) {
	var rv LookupDomainResult
	err := ctx.Invoke("azure-native:domainregistration/v20201201:getDomain", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupDomainArgs struct {
	DomainName        string `pulumi:"domainName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupDomainResult struct {
	AuthCode                    *string            `pulumi:"authCode"`
	AutoRenew                   *bool              `pulumi:"autoRenew"`
	CreatedTime                 string             `pulumi:"createdTime"`
	DnsType                     *string            `pulumi:"dnsType"`
	DnsZoneId                   *string            `pulumi:"dnsZoneId"`
	DomainNotRenewableReasons   []string           `pulumi:"domainNotRenewableReasons"`
	ExpirationTime              string             `pulumi:"expirationTime"`
	Id                          string             `pulumi:"id"`
	Kind                        *string            `pulumi:"kind"`
	LastRenewedTime             string             `pulumi:"lastRenewedTime"`
	Location                    string             `pulumi:"location"`
	ManagedHostNames            []HostNameResponse `pulumi:"managedHostNames"`
	Name                        string             `pulumi:"name"`
	NameServers                 []string           `pulumi:"nameServers"`
	Privacy                     *bool              `pulumi:"privacy"`
	ProvisioningState           string             `pulumi:"provisioningState"`
	ReadyForDnsRecordManagement bool               `pulumi:"readyForDnsRecordManagement"`
	RegistrationStatus          string             `pulumi:"registrationStatus"`
	Tags                        map[string]string  `pulumi:"tags"`
	TargetDnsType               *string            `pulumi:"targetDnsType"`
	Type                        string             `pulumi:"type"`
}


func (val *LookupDomainResult) Defaults() *LookupDomainResult {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.AutoRenew) {
		autoRenew_ := true
		tmp.AutoRenew = &autoRenew_
	}
	return &tmp
}
