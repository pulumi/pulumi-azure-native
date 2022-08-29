


package sql

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupDisasterRecoveryConfiguration(ctx *pulumi.Context, args *LookupDisasterRecoveryConfigurationArgs, opts ...pulumi.InvokeOption) (*LookupDisasterRecoveryConfigurationResult, error) {
	var rv LookupDisasterRecoveryConfigurationResult
	err := ctx.Invoke("azure-native:sql:getDisasterRecoveryConfiguration", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDisasterRecoveryConfigurationArgs struct {
	DisasterRecoveryConfigurationName string `pulumi:"disasterRecoveryConfigurationName"`
	ResourceGroupName                 string `pulumi:"resourceGroupName"`
	ServerName                        string `pulumi:"serverName"`
}


type LookupDisasterRecoveryConfigurationResult struct {
	AutoFailover             string `pulumi:"autoFailover"`
	FailoverPolicy           string `pulumi:"failoverPolicy"`
	Id                       string `pulumi:"id"`
	Location                 string `pulumi:"location"`
	LogicalServerName        string `pulumi:"logicalServerName"`
	Name                     string `pulumi:"name"`
	PartnerLogicalServerName string `pulumi:"partnerLogicalServerName"`
	PartnerServerId          string `pulumi:"partnerServerId"`
	Role                     string `pulumi:"role"`
	Status                   string `pulumi:"status"`
	Type                     string `pulumi:"type"`
}

func LookupDisasterRecoveryConfigurationOutput(ctx *pulumi.Context, args LookupDisasterRecoveryConfigurationOutputArgs, opts ...pulumi.InvokeOption) LookupDisasterRecoveryConfigurationResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDisasterRecoveryConfigurationResult, error) {
			args := v.(LookupDisasterRecoveryConfigurationArgs)
			r, err := LookupDisasterRecoveryConfiguration(ctx, &args, opts...)
			var s LookupDisasterRecoveryConfigurationResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupDisasterRecoveryConfigurationResultOutput)
}

type LookupDisasterRecoveryConfigurationOutputArgs struct {
	DisasterRecoveryConfigurationName pulumi.StringInput `pulumi:"disasterRecoveryConfigurationName"`
	ResourceGroupName                 pulumi.StringInput `pulumi:"resourceGroupName"`
	ServerName                        pulumi.StringInput `pulumi:"serverName"`
}

func (LookupDisasterRecoveryConfigurationOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDisasterRecoveryConfigurationArgs)(nil)).Elem()
}


type LookupDisasterRecoveryConfigurationResultOutput struct{ *pulumi.OutputState }

func (LookupDisasterRecoveryConfigurationResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDisasterRecoveryConfigurationResult)(nil)).Elem()
}

func (o LookupDisasterRecoveryConfigurationResultOutput) ToLookupDisasterRecoveryConfigurationResultOutput() LookupDisasterRecoveryConfigurationResultOutput {
	return o
}

func (o LookupDisasterRecoveryConfigurationResultOutput) ToLookupDisasterRecoveryConfigurationResultOutputWithContext(ctx context.Context) LookupDisasterRecoveryConfigurationResultOutput {
	return o
}

func (o LookupDisasterRecoveryConfigurationResultOutput) AutoFailover() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDisasterRecoveryConfigurationResult) string { return v.AutoFailover }).(pulumi.StringOutput)
}

func (o LookupDisasterRecoveryConfigurationResultOutput) FailoverPolicy() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDisasterRecoveryConfigurationResult) string { return v.FailoverPolicy }).(pulumi.StringOutput)
}

func (o LookupDisasterRecoveryConfigurationResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDisasterRecoveryConfigurationResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupDisasterRecoveryConfigurationResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDisasterRecoveryConfigurationResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupDisasterRecoveryConfigurationResultOutput) LogicalServerName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDisasterRecoveryConfigurationResult) string { return v.LogicalServerName }).(pulumi.StringOutput)
}

func (o LookupDisasterRecoveryConfigurationResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDisasterRecoveryConfigurationResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupDisasterRecoveryConfigurationResultOutput) PartnerLogicalServerName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDisasterRecoveryConfigurationResult) string { return v.PartnerLogicalServerName }).(pulumi.StringOutput)
}

func (o LookupDisasterRecoveryConfigurationResultOutput) PartnerServerId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDisasterRecoveryConfigurationResult) string { return v.PartnerServerId }).(pulumi.StringOutput)
}

func (o LookupDisasterRecoveryConfigurationResultOutput) Role() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDisasterRecoveryConfigurationResult) string { return v.Role }).(pulumi.StringOutput)
}

func (o LookupDisasterRecoveryConfigurationResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDisasterRecoveryConfigurationResult) string { return v.Status }).(pulumi.StringOutput)
}

func (o LookupDisasterRecoveryConfigurationResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDisasterRecoveryConfigurationResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDisasterRecoveryConfigurationResultOutput{})
}
