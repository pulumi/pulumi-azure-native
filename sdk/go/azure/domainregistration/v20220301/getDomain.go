


package v20220301

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupDomain(ctx *pulumi.Context, args *LookupDomainArgs, opts ...pulumi.InvokeOption) (*LookupDomainResult, error) {
	var rv LookupDomainResult
	err := ctx.Invoke("azure-native:domainregistration/v20220301:getDomain", args, &rv, opts...)
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

func LookupDomainOutput(ctx *pulumi.Context, args LookupDomainOutputArgs, opts ...pulumi.InvokeOption) LookupDomainResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDomainResult, error) {
			args := v.(LookupDomainArgs)
			r, err := LookupDomain(ctx, &args, opts...)
			var s LookupDomainResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupDomainResultOutput)
}

type LookupDomainOutputArgs struct {
	DomainName        pulumi.StringInput `pulumi:"domainName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupDomainOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDomainArgs)(nil)).Elem()
}


type LookupDomainResultOutput struct{ *pulumi.OutputState }

func (LookupDomainResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDomainResult)(nil)).Elem()
}

func (o LookupDomainResultOutput) ToLookupDomainResultOutput() LookupDomainResultOutput {
	return o
}

func (o LookupDomainResultOutput) ToLookupDomainResultOutputWithContext(ctx context.Context) LookupDomainResultOutput {
	return o
}

func (o LookupDomainResultOutput) AuthCode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDomainResult) *string { return v.AuthCode }).(pulumi.StringPtrOutput)
}

func (o LookupDomainResultOutput) AutoRenew() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupDomainResult) *bool { return v.AutoRenew }).(pulumi.BoolPtrOutput)
}

func (o LookupDomainResultOutput) CreatedTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDomainResult) string { return v.CreatedTime }).(pulumi.StringOutput)
}

func (o LookupDomainResultOutput) DnsType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDomainResult) *string { return v.DnsType }).(pulumi.StringPtrOutput)
}

func (o LookupDomainResultOutput) DnsZoneId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDomainResult) *string { return v.DnsZoneId }).(pulumi.StringPtrOutput)
}

func (o LookupDomainResultOutput) DomainNotRenewableReasons() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupDomainResult) []string { return v.DomainNotRenewableReasons }).(pulumi.StringArrayOutput)
}

func (o LookupDomainResultOutput) ExpirationTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDomainResult) string { return v.ExpirationTime }).(pulumi.StringOutput)
}

func (o LookupDomainResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDomainResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupDomainResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDomainResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o LookupDomainResultOutput) LastRenewedTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDomainResult) string { return v.LastRenewedTime }).(pulumi.StringOutput)
}

func (o LookupDomainResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDomainResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupDomainResultOutput) ManagedHostNames() HostNameResponseArrayOutput {
	return o.ApplyT(func(v LookupDomainResult) []HostNameResponse { return v.ManagedHostNames }).(HostNameResponseArrayOutput)
}

func (o LookupDomainResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDomainResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupDomainResultOutput) NameServers() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupDomainResult) []string { return v.NameServers }).(pulumi.StringArrayOutput)
}

func (o LookupDomainResultOutput) Privacy() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupDomainResult) *bool { return v.Privacy }).(pulumi.BoolPtrOutput)
}

func (o LookupDomainResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDomainResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupDomainResultOutput) ReadyForDnsRecordManagement() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupDomainResult) bool { return v.ReadyForDnsRecordManagement }).(pulumi.BoolOutput)
}

func (o LookupDomainResultOutput) RegistrationStatus() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDomainResult) string { return v.RegistrationStatus }).(pulumi.StringOutput)
}

func (o LookupDomainResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupDomainResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupDomainResultOutput) TargetDnsType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDomainResult) *string { return v.TargetDnsType }).(pulumi.StringPtrOutput)
}

func (o LookupDomainResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDomainResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDomainResultOutput{})
}
