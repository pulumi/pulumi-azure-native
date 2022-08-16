


package v20210501preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupServerSecurityAlertPolicy(ctx *pulumi.Context, args *LookupServerSecurityAlertPolicyArgs, opts ...pulumi.InvokeOption) (*LookupServerSecurityAlertPolicyResult, error) {
	var rv LookupServerSecurityAlertPolicyResult
	err := ctx.Invoke("azure-native:sql/v20210501preview:getServerSecurityAlertPolicy", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupServerSecurityAlertPolicyArgs struct {
	ResourceGroupName       string `pulumi:"resourceGroupName"`
	SecurityAlertPolicyName string `pulumi:"securityAlertPolicyName"`
	ServerName              string `pulumi:"serverName"`
}


type LookupServerSecurityAlertPolicyResult struct {
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

func LookupServerSecurityAlertPolicyOutput(ctx *pulumi.Context, args LookupServerSecurityAlertPolicyOutputArgs, opts ...pulumi.InvokeOption) LookupServerSecurityAlertPolicyResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupServerSecurityAlertPolicyResult, error) {
			args := v.(LookupServerSecurityAlertPolicyArgs)
			r, err := LookupServerSecurityAlertPolicy(ctx, &args, opts...)
			var s LookupServerSecurityAlertPolicyResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupServerSecurityAlertPolicyResultOutput)
}

type LookupServerSecurityAlertPolicyOutputArgs struct {
	ResourceGroupName       pulumi.StringInput `pulumi:"resourceGroupName"`
	SecurityAlertPolicyName pulumi.StringInput `pulumi:"securityAlertPolicyName"`
	ServerName              pulumi.StringInput `pulumi:"serverName"`
}

func (LookupServerSecurityAlertPolicyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupServerSecurityAlertPolicyArgs)(nil)).Elem()
}


type LookupServerSecurityAlertPolicyResultOutput struct{ *pulumi.OutputState }

func (LookupServerSecurityAlertPolicyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupServerSecurityAlertPolicyResult)(nil)).Elem()
}

func (o LookupServerSecurityAlertPolicyResultOutput) ToLookupServerSecurityAlertPolicyResultOutput() LookupServerSecurityAlertPolicyResultOutput {
	return o
}

func (o LookupServerSecurityAlertPolicyResultOutput) ToLookupServerSecurityAlertPolicyResultOutputWithContext(ctx context.Context) LookupServerSecurityAlertPolicyResultOutput {
	return o
}

func (o LookupServerSecurityAlertPolicyResultOutput) CreationTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerSecurityAlertPolicyResult) string { return v.CreationTime }).(pulumi.StringOutput)
}

func (o LookupServerSecurityAlertPolicyResultOutput) DisabledAlerts() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupServerSecurityAlertPolicyResult) []string { return v.DisabledAlerts }).(pulumi.StringArrayOutput)
}

func (o LookupServerSecurityAlertPolicyResultOutput) EmailAccountAdmins() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupServerSecurityAlertPolicyResult) *bool { return v.EmailAccountAdmins }).(pulumi.BoolPtrOutput)
}

func (o LookupServerSecurityAlertPolicyResultOutput) EmailAddresses() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupServerSecurityAlertPolicyResult) []string { return v.EmailAddresses }).(pulumi.StringArrayOutput)
}

func (o LookupServerSecurityAlertPolicyResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerSecurityAlertPolicyResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupServerSecurityAlertPolicyResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerSecurityAlertPolicyResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupServerSecurityAlertPolicyResultOutput) RetentionDays() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupServerSecurityAlertPolicyResult) *int { return v.RetentionDays }).(pulumi.IntPtrOutput)
}

func (o LookupServerSecurityAlertPolicyResultOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerSecurityAlertPolicyResult) string { return v.State }).(pulumi.StringOutput)
}

func (o LookupServerSecurityAlertPolicyResultOutput) StorageAccountAccessKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupServerSecurityAlertPolicyResult) *string { return v.StorageAccountAccessKey }).(pulumi.StringPtrOutput)
}

func (o LookupServerSecurityAlertPolicyResultOutput) StorageEndpoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupServerSecurityAlertPolicyResult) *string { return v.StorageEndpoint }).(pulumi.StringPtrOutput)
}

func (o LookupServerSecurityAlertPolicyResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupServerSecurityAlertPolicyResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupServerSecurityAlertPolicyResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerSecurityAlertPolicyResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupServerSecurityAlertPolicyResultOutput{})
}
