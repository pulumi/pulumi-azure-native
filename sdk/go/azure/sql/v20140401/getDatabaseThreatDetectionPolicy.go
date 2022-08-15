


package v20140401

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupDatabaseThreatDetectionPolicy(ctx *pulumi.Context, args *LookupDatabaseThreatDetectionPolicyArgs, opts ...pulumi.InvokeOption) (*LookupDatabaseThreatDetectionPolicyResult, error) {
	var rv LookupDatabaseThreatDetectionPolicyResult
	err := ctx.Invoke("azure-native:sql/v20140401:getDatabaseThreatDetectionPolicy", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDatabaseThreatDetectionPolicyArgs struct {
	DatabaseName            string `pulumi:"databaseName"`
	ResourceGroupName       string `pulumi:"resourceGroupName"`
	SecurityAlertPolicyName string `pulumi:"securityAlertPolicyName"`
	ServerName              string `pulumi:"serverName"`
}


type LookupDatabaseThreatDetectionPolicyResult struct {
	DisabledAlerts     *string `pulumi:"disabledAlerts"`
	EmailAccountAdmins *string `pulumi:"emailAccountAdmins"`
	EmailAddresses     *string `pulumi:"emailAddresses"`
	Id                 string  `pulumi:"id"`
	Kind               string  `pulumi:"kind"`
	Location           *string `pulumi:"location"`
	Name               string  `pulumi:"name"`
	RetentionDays      *int    `pulumi:"retentionDays"`
	State              string  `pulumi:"state"`
	StorageEndpoint    *string `pulumi:"storageEndpoint"`
	Type               string  `pulumi:"type"`
	UseServerDefault   *string `pulumi:"useServerDefault"`
}

func LookupDatabaseThreatDetectionPolicyOutput(ctx *pulumi.Context, args LookupDatabaseThreatDetectionPolicyOutputArgs, opts ...pulumi.InvokeOption) LookupDatabaseThreatDetectionPolicyResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDatabaseThreatDetectionPolicyResult, error) {
			args := v.(LookupDatabaseThreatDetectionPolicyArgs)
			r, err := LookupDatabaseThreatDetectionPolicy(ctx, &args, opts...)
			var s LookupDatabaseThreatDetectionPolicyResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupDatabaseThreatDetectionPolicyResultOutput)
}

type LookupDatabaseThreatDetectionPolicyOutputArgs struct {
	DatabaseName            pulumi.StringInput `pulumi:"databaseName"`
	ResourceGroupName       pulumi.StringInput `pulumi:"resourceGroupName"`
	SecurityAlertPolicyName pulumi.StringInput `pulumi:"securityAlertPolicyName"`
	ServerName              pulumi.StringInput `pulumi:"serverName"`
}

func (LookupDatabaseThreatDetectionPolicyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDatabaseThreatDetectionPolicyArgs)(nil)).Elem()
}


type LookupDatabaseThreatDetectionPolicyResultOutput struct{ *pulumi.OutputState }

func (LookupDatabaseThreatDetectionPolicyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDatabaseThreatDetectionPolicyResult)(nil)).Elem()
}

func (o LookupDatabaseThreatDetectionPolicyResultOutput) ToLookupDatabaseThreatDetectionPolicyResultOutput() LookupDatabaseThreatDetectionPolicyResultOutput {
	return o
}

func (o LookupDatabaseThreatDetectionPolicyResultOutput) ToLookupDatabaseThreatDetectionPolicyResultOutputWithContext(ctx context.Context) LookupDatabaseThreatDetectionPolicyResultOutput {
	return o
}

func (o LookupDatabaseThreatDetectionPolicyResultOutput) DisabledAlerts() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDatabaseThreatDetectionPolicyResult) *string { return v.DisabledAlerts }).(pulumi.StringPtrOutput)
}

func (o LookupDatabaseThreatDetectionPolicyResultOutput) EmailAccountAdmins() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDatabaseThreatDetectionPolicyResult) *string { return v.EmailAccountAdmins }).(pulumi.StringPtrOutput)
}

func (o LookupDatabaseThreatDetectionPolicyResultOutput) EmailAddresses() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDatabaseThreatDetectionPolicyResult) *string { return v.EmailAddresses }).(pulumi.StringPtrOutput)
}

func (o LookupDatabaseThreatDetectionPolicyResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatabaseThreatDetectionPolicyResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupDatabaseThreatDetectionPolicyResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatabaseThreatDetectionPolicyResult) string { return v.Kind }).(pulumi.StringOutput)
}

func (o LookupDatabaseThreatDetectionPolicyResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDatabaseThreatDetectionPolicyResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupDatabaseThreatDetectionPolicyResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatabaseThreatDetectionPolicyResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupDatabaseThreatDetectionPolicyResultOutput) RetentionDays() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupDatabaseThreatDetectionPolicyResult) *int { return v.RetentionDays }).(pulumi.IntPtrOutput)
}

func (o LookupDatabaseThreatDetectionPolicyResultOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatabaseThreatDetectionPolicyResult) string { return v.State }).(pulumi.StringOutput)
}

func (o LookupDatabaseThreatDetectionPolicyResultOutput) StorageEndpoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDatabaseThreatDetectionPolicyResult) *string { return v.StorageEndpoint }).(pulumi.StringPtrOutput)
}

func (o LookupDatabaseThreatDetectionPolicyResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatabaseThreatDetectionPolicyResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupDatabaseThreatDetectionPolicyResultOutput) UseServerDefault() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDatabaseThreatDetectionPolicyResult) *string { return v.UseServerDefault }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDatabaseThreatDetectionPolicyResultOutput{})
}
