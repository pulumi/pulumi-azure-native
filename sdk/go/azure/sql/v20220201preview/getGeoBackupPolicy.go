


package v20220201preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupGeoBackupPolicy(ctx *pulumi.Context, args *LookupGeoBackupPolicyArgs, opts ...pulumi.InvokeOption) (*LookupGeoBackupPolicyResult, error) {
	var rv LookupGeoBackupPolicyResult
	err := ctx.Invoke("azure-native:sql/v20220201preview:getGeoBackupPolicy", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupGeoBackupPolicyArgs struct {
	DatabaseName        string `pulumi:"databaseName"`
	GeoBackupPolicyName string `pulumi:"geoBackupPolicyName"`
	ResourceGroupName   string `pulumi:"resourceGroupName"`
	ServerName          string `pulumi:"serverName"`
}


type LookupGeoBackupPolicyResult struct {
	Id          string `pulumi:"id"`
	Kind        string `pulumi:"kind"`
	Location    string `pulumi:"location"`
	Name        string `pulumi:"name"`
	State       string `pulumi:"state"`
	StorageType string `pulumi:"storageType"`
	Type        string `pulumi:"type"`
}

func LookupGeoBackupPolicyOutput(ctx *pulumi.Context, args LookupGeoBackupPolicyOutputArgs, opts ...pulumi.InvokeOption) LookupGeoBackupPolicyResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupGeoBackupPolicyResult, error) {
			args := v.(LookupGeoBackupPolicyArgs)
			r, err := LookupGeoBackupPolicy(ctx, &args, opts...)
			var s LookupGeoBackupPolicyResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupGeoBackupPolicyResultOutput)
}

type LookupGeoBackupPolicyOutputArgs struct {
	DatabaseName        pulumi.StringInput `pulumi:"databaseName"`
	GeoBackupPolicyName pulumi.StringInput `pulumi:"geoBackupPolicyName"`
	ResourceGroupName   pulumi.StringInput `pulumi:"resourceGroupName"`
	ServerName          pulumi.StringInput `pulumi:"serverName"`
}

func (LookupGeoBackupPolicyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupGeoBackupPolicyArgs)(nil)).Elem()
}


type LookupGeoBackupPolicyResultOutput struct{ *pulumi.OutputState }

func (LookupGeoBackupPolicyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupGeoBackupPolicyResult)(nil)).Elem()
}

func (o LookupGeoBackupPolicyResultOutput) ToLookupGeoBackupPolicyResultOutput() LookupGeoBackupPolicyResultOutput {
	return o
}

func (o LookupGeoBackupPolicyResultOutput) ToLookupGeoBackupPolicyResultOutputWithContext(ctx context.Context) LookupGeoBackupPolicyResultOutput {
	return o
}

func (o LookupGeoBackupPolicyResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGeoBackupPolicyResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupGeoBackupPolicyResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGeoBackupPolicyResult) string { return v.Kind }).(pulumi.StringOutput)
}

func (o LookupGeoBackupPolicyResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGeoBackupPolicyResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupGeoBackupPolicyResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGeoBackupPolicyResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupGeoBackupPolicyResultOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGeoBackupPolicyResult) string { return v.State }).(pulumi.StringOutput)
}

func (o LookupGeoBackupPolicyResultOutput) StorageType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGeoBackupPolicyResult) string { return v.StorageType }).(pulumi.StringOutput)
}

func (o LookupGeoBackupPolicyResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGeoBackupPolicyResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupGeoBackupPolicyResultOutput{})
}
