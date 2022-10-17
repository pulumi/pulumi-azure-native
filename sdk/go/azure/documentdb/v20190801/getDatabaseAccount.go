


package v20190801

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


func LookupDatabaseAccount(ctx *pulumi.Context, args *LookupDatabaseAccountArgs, opts ...pulumi.InvokeOption) (*LookupDatabaseAccountResult, error) {
	var rv LookupDatabaseAccountResult
	err := ctx.Invoke("azure-native:documentdb/v20190801:getDatabaseAccount", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupDatabaseAccountArgs struct {
	AccountName       string `pulumi:"accountName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupDatabaseAccountResult struct {
	Capabilities                       []CapabilityResponse         `pulumi:"capabilities"`
	ConnectorOffer                     *string                      `pulumi:"connectorOffer"`
	ConsistencyPolicy                  *ConsistencyPolicyResponse   `pulumi:"consistencyPolicy"`
	DatabaseAccountOfferType           string                       `pulumi:"databaseAccountOfferType"`
	DisableKeyBasedMetadataWriteAccess *bool                        `pulumi:"disableKeyBasedMetadataWriteAccess"`
	DocumentEndpoint                   string                       `pulumi:"documentEndpoint"`
	EnableAutomaticFailover            *bool                        `pulumi:"enableAutomaticFailover"`
	EnableCassandraConnector           *bool                        `pulumi:"enableCassandraConnector"`
	EnableMultipleWriteLocations       *bool                        `pulumi:"enableMultipleWriteLocations"`
	FailoverPolicies                   []FailoverPolicyResponse     `pulumi:"failoverPolicies"`
	Id                                 string                       `pulumi:"id"`
	IpRangeFilter                      *string                      `pulumi:"ipRangeFilter"`
	IsVirtualNetworkFilterEnabled      *bool                        `pulumi:"isVirtualNetworkFilterEnabled"`
	Kind                               *string                      `pulumi:"kind"`
	Location                           *string                      `pulumi:"location"`
	Locations                          []LocationResponse           `pulumi:"locations"`
	Name                               string                       `pulumi:"name"`
	ProvisioningState                  string                       `pulumi:"provisioningState"`
	ReadLocations                      []LocationResponse           `pulumi:"readLocations"`
	Tags                               map[string]string            `pulumi:"tags"`
	Type                               string                       `pulumi:"type"`
	VirtualNetworkRules                []VirtualNetworkRuleResponse `pulumi:"virtualNetworkRules"`
	WriteLocations                     []LocationResponse           `pulumi:"writeLocations"`
}


func (val *LookupDatabaseAccountResult) Defaults() *LookupDatabaseAccountResult {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Kind) {
		kind_ := "GlobalDocumentDB"
		tmp.Kind = &kind_
	}
	return &tmp
}

func LookupDatabaseAccountOutput(ctx *pulumi.Context, args LookupDatabaseAccountOutputArgs, opts ...pulumi.InvokeOption) LookupDatabaseAccountResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDatabaseAccountResult, error) {
			args := v.(LookupDatabaseAccountArgs)
			r, err := LookupDatabaseAccount(ctx, &args, opts...)
			var s LookupDatabaseAccountResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupDatabaseAccountResultOutput)
}

type LookupDatabaseAccountOutputArgs struct {
	AccountName       pulumi.StringInput `pulumi:"accountName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupDatabaseAccountOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDatabaseAccountArgs)(nil)).Elem()
}


type LookupDatabaseAccountResultOutput struct{ *pulumi.OutputState }

func (LookupDatabaseAccountResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDatabaseAccountResult)(nil)).Elem()
}

func (o LookupDatabaseAccountResultOutput) ToLookupDatabaseAccountResultOutput() LookupDatabaseAccountResultOutput {
	return o
}

func (o LookupDatabaseAccountResultOutput) ToLookupDatabaseAccountResultOutputWithContext(ctx context.Context) LookupDatabaseAccountResultOutput {
	return o
}

func (o LookupDatabaseAccountResultOutput) Capabilities() CapabilityResponseArrayOutput {
	return o.ApplyT(func(v LookupDatabaseAccountResult) []CapabilityResponse { return v.Capabilities }).(CapabilityResponseArrayOutput)
}

func (o LookupDatabaseAccountResultOutput) ConnectorOffer() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDatabaseAccountResult) *string { return v.ConnectorOffer }).(pulumi.StringPtrOutput)
}

func (o LookupDatabaseAccountResultOutput) ConsistencyPolicy() ConsistencyPolicyResponsePtrOutput {
	return o.ApplyT(func(v LookupDatabaseAccountResult) *ConsistencyPolicyResponse { return v.ConsistencyPolicy }).(ConsistencyPolicyResponsePtrOutput)
}

func (o LookupDatabaseAccountResultOutput) DatabaseAccountOfferType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatabaseAccountResult) string { return v.DatabaseAccountOfferType }).(pulumi.StringOutput)
}

func (o LookupDatabaseAccountResultOutput) DisableKeyBasedMetadataWriteAccess() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupDatabaseAccountResult) *bool { return v.DisableKeyBasedMetadataWriteAccess }).(pulumi.BoolPtrOutput)
}

func (o LookupDatabaseAccountResultOutput) DocumentEndpoint() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatabaseAccountResult) string { return v.DocumentEndpoint }).(pulumi.StringOutput)
}

func (o LookupDatabaseAccountResultOutput) EnableAutomaticFailover() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupDatabaseAccountResult) *bool { return v.EnableAutomaticFailover }).(pulumi.BoolPtrOutput)
}

func (o LookupDatabaseAccountResultOutput) EnableCassandraConnector() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupDatabaseAccountResult) *bool { return v.EnableCassandraConnector }).(pulumi.BoolPtrOutput)
}

func (o LookupDatabaseAccountResultOutput) EnableMultipleWriteLocations() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupDatabaseAccountResult) *bool { return v.EnableMultipleWriteLocations }).(pulumi.BoolPtrOutput)
}

func (o LookupDatabaseAccountResultOutput) FailoverPolicies() FailoverPolicyResponseArrayOutput {
	return o.ApplyT(func(v LookupDatabaseAccountResult) []FailoverPolicyResponse { return v.FailoverPolicies }).(FailoverPolicyResponseArrayOutput)
}

func (o LookupDatabaseAccountResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatabaseAccountResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupDatabaseAccountResultOutput) IpRangeFilter() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDatabaseAccountResult) *string { return v.IpRangeFilter }).(pulumi.StringPtrOutput)
}

func (o LookupDatabaseAccountResultOutput) IsVirtualNetworkFilterEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupDatabaseAccountResult) *bool { return v.IsVirtualNetworkFilterEnabled }).(pulumi.BoolPtrOutput)
}

func (o LookupDatabaseAccountResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDatabaseAccountResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o LookupDatabaseAccountResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDatabaseAccountResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupDatabaseAccountResultOutput) Locations() LocationResponseArrayOutput {
	return o.ApplyT(func(v LookupDatabaseAccountResult) []LocationResponse { return v.Locations }).(LocationResponseArrayOutput)
}

func (o LookupDatabaseAccountResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatabaseAccountResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupDatabaseAccountResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatabaseAccountResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupDatabaseAccountResultOutput) ReadLocations() LocationResponseArrayOutput {
	return o.ApplyT(func(v LookupDatabaseAccountResult) []LocationResponse { return v.ReadLocations }).(LocationResponseArrayOutput)
}

func (o LookupDatabaseAccountResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupDatabaseAccountResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupDatabaseAccountResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatabaseAccountResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupDatabaseAccountResultOutput) VirtualNetworkRules() VirtualNetworkRuleResponseArrayOutput {
	return o.ApplyT(func(v LookupDatabaseAccountResult) []VirtualNetworkRuleResponse { return v.VirtualNetworkRules }).(VirtualNetworkRuleResponseArrayOutput)
}

func (o LookupDatabaseAccountResultOutput) WriteLocations() LocationResponseArrayOutput {
	return o.ApplyT(func(v LookupDatabaseAccountResult) []LocationResponse { return v.WriteLocations }).(LocationResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDatabaseAccountResultOutput{})
}
