


package sql

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupManagedServerDnsAlias(ctx *pulumi.Context, args *LookupManagedServerDnsAliasArgs, opts ...pulumi.InvokeOption) (*LookupManagedServerDnsAliasResult, error) {
	var rv LookupManagedServerDnsAliasResult
	err := ctx.Invoke("azure-native:sql:getManagedServerDnsAlias", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupManagedServerDnsAliasArgs struct {
	DnsAliasName        string `pulumi:"dnsAliasName"`
	ManagedInstanceName string `pulumi:"managedInstanceName"`
	ResourceGroupName   string `pulumi:"resourceGroupName"`
}


type LookupManagedServerDnsAliasResult struct {
	AzureDnsRecord       string `pulumi:"azureDnsRecord"`
	Id                   string `pulumi:"id"`
	Name                 string `pulumi:"name"`
	PublicAzureDnsRecord string `pulumi:"publicAzureDnsRecord"`
	Type                 string `pulumi:"type"`
}

func LookupManagedServerDnsAliasOutput(ctx *pulumi.Context, args LookupManagedServerDnsAliasOutputArgs, opts ...pulumi.InvokeOption) LookupManagedServerDnsAliasResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupManagedServerDnsAliasResult, error) {
			args := v.(LookupManagedServerDnsAliasArgs)
			r, err := LookupManagedServerDnsAlias(ctx, &args, opts...)
			var s LookupManagedServerDnsAliasResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupManagedServerDnsAliasResultOutput)
}

type LookupManagedServerDnsAliasOutputArgs struct {
	DnsAliasName        pulumi.StringInput `pulumi:"dnsAliasName"`
	ManagedInstanceName pulumi.StringInput `pulumi:"managedInstanceName"`
	ResourceGroupName   pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupManagedServerDnsAliasOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupManagedServerDnsAliasArgs)(nil)).Elem()
}


type LookupManagedServerDnsAliasResultOutput struct{ *pulumi.OutputState }

func (LookupManagedServerDnsAliasResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupManagedServerDnsAliasResult)(nil)).Elem()
}

func (o LookupManagedServerDnsAliasResultOutput) ToLookupManagedServerDnsAliasResultOutput() LookupManagedServerDnsAliasResultOutput {
	return o
}

func (o LookupManagedServerDnsAliasResultOutput) ToLookupManagedServerDnsAliasResultOutputWithContext(ctx context.Context) LookupManagedServerDnsAliasResultOutput {
	return o
}

func (o LookupManagedServerDnsAliasResultOutput) AzureDnsRecord() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagedServerDnsAliasResult) string { return v.AzureDnsRecord }).(pulumi.StringOutput)
}

func (o LookupManagedServerDnsAliasResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagedServerDnsAliasResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupManagedServerDnsAliasResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagedServerDnsAliasResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupManagedServerDnsAliasResultOutput) PublicAzureDnsRecord() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagedServerDnsAliasResult) string { return v.PublicAzureDnsRecord }).(pulumi.StringOutput)
}

func (o LookupManagedServerDnsAliasResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagedServerDnsAliasResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupManagedServerDnsAliasResultOutput{})
}
