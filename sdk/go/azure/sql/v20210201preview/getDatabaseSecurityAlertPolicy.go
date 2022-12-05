


package v20210201preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupDatabaseSecurityAlertPolicy(ctx *pulumi.Context, args *LookupDatabaseSecurityAlertPolicyArgs, opts ...pulumi.InvokeOption) (*LookupDatabaseSecurityAlertPolicyResult, error) {
	var rv LookupDatabaseSecurityAlertPolicyResult
	err := ctx.Invoke("azure-native:sql/v20210201preview:getDatabaseSecurityAlertPolicy", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDatabaseSecurityAlertPolicyArgs struct {
	DatabaseName            string `pulumi:"databaseName"`
	ResourceGroupName       string `pulumi:"resourceGroupName"`
	SecurityAlertPolicyName string `pulumi:"securityAlertPolicyName"`
	ServerName              string `pulumi:"serverName"`
}


type LookupDatabaseSecurityAlertPolicyResult struct {
	CreationTime            string             `pulumi:"creationTime"`
	DisabledAlerts          []string           `pulumi:"disabledAlerts"`
	EmailAccountAdmins      *bool              `pulumi:"emailAccountAdmins"`
	EmailAddresses          []string           `pulumi:"emailAddresses"`
	Id                      string             `pulumi:"id"`
	Name                    string             `pulumi:"name"`
	RetentionDays           *int               `pulumi:"retentionDays"`
	State                   string             `pulumi:"state"`
	StorageAccountAccessKey *string            `pulumi:"storageAccountAccessKey"`
	StorageEndpoint         *string            `pulumi:"storageEndpoint"`
	SystemData              SystemDataResponse `pulumi:"systemData"`
	Type                    string             `pulumi:"type"`
}

func LookupDatabaseSecurityAlertPolicyOutput(ctx *pulumi.Context, args LookupDatabaseSecurityAlertPolicyOutputArgs, opts ...pulumi.InvokeOption) LookupDatabaseSecurityAlertPolicyResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDatabaseSecurityAlertPolicyResult, error) {
			args := v.(LookupDatabaseSecurityAlertPolicyArgs)
			r, err := LookupDatabaseSecurityAlertPolicy(ctx, &args, opts...)
			var s LookupDatabaseSecurityAlertPolicyResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupDatabaseSecurityAlertPolicyResultOutput)
}

type LookupDatabaseSecurityAlertPolicyOutputArgs struct {
	DatabaseName            pulumi.StringInput `pulumi:"databaseName"`
	ResourceGroupName       pulumi.StringInput `pulumi:"resourceGroupName"`
	SecurityAlertPolicyName pulumi.StringInput `pulumi:"securityAlertPolicyName"`
	ServerName              pulumi.StringInput `pulumi:"serverName"`
}

func (LookupDatabaseSecurityAlertPolicyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDatabaseSecurityAlertPolicyArgs)(nil)).Elem()
}


type LookupDatabaseSecurityAlertPolicyResultOutput struct{ *pulumi.OutputState }

func (LookupDatabaseSecurityAlertPolicyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDatabaseSecurityAlertPolicyResult)(nil)).Elem()
}

func (o LookupDatabaseSecurityAlertPolicyResultOutput) ToLookupDatabaseSecurityAlertPolicyResultOutput() LookupDatabaseSecurityAlertPolicyResultOutput {
	return o
}

func (o LookupDatabaseSecurityAlertPolicyResultOutput) ToLookupDatabaseSecurityAlertPolicyResultOutputWithContext(ctx context.Context) LookupDatabaseSecurityAlertPolicyResultOutput {
	return o
}

func (o LookupDatabaseSecurityAlertPolicyResultOutput) CreationTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatabaseSecurityAlertPolicyResult) string { return v.CreationTime }).(pulumi.StringOutput)
}

func (o LookupDatabaseSecurityAlertPolicyResultOutput) DisabledAlerts() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupDatabaseSecurityAlertPolicyResult) []string { return v.DisabledAlerts }).(pulumi.StringArrayOutput)
}

func (o LookupDatabaseSecurityAlertPolicyResultOutput) EmailAccountAdmins() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupDatabaseSecurityAlertPolicyResult) *bool { return v.EmailAccountAdmins }).(pulumi.BoolPtrOutput)
}

func (o LookupDatabaseSecurityAlertPolicyResultOutput) EmailAddresses() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupDatabaseSecurityAlertPolicyResult) []string { return v.EmailAddresses }).(pulumi.StringArrayOutput)
}

func (o LookupDatabaseSecurityAlertPolicyResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatabaseSecurityAlertPolicyResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupDatabaseSecurityAlertPolicyResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatabaseSecurityAlertPolicyResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupDatabaseSecurityAlertPolicyResultOutput) RetentionDays() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupDatabaseSecurityAlertPolicyResult) *int { return v.RetentionDays }).(pulumi.IntPtrOutput)
}

func (o LookupDatabaseSecurityAlertPolicyResultOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatabaseSecurityAlertPolicyResult) string { return v.State }).(pulumi.StringOutput)
}

func (o LookupDatabaseSecurityAlertPolicyResultOutput) StorageAccountAccessKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDatabaseSecurityAlertPolicyResult) *string { return v.StorageAccountAccessKey }).(pulumi.StringPtrOutput)
}

func (o LookupDatabaseSecurityAlertPolicyResultOutput) StorageEndpoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDatabaseSecurityAlertPolicyResult) *string { return v.StorageEndpoint }).(pulumi.StringPtrOutput)
}

func (o LookupDatabaseSecurityAlertPolicyResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupDatabaseSecurityAlertPolicyResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupDatabaseSecurityAlertPolicyResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatabaseSecurityAlertPolicyResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDatabaseSecurityAlertPolicyResultOutput{})
}
