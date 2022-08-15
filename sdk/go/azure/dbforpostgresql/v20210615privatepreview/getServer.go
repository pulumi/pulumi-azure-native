


package v20210615privatepreview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupServer(ctx *pulumi.Context, args *LookupServerArgs, opts ...pulumi.InvokeOption) (*LookupServerResult, error) {
	var rv LookupServerResult
	err := ctx.Invoke("azure-native:dbforpostgresql/v20210615privatepreview:getServer", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupServerArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ServerName        string `pulumi:"serverName"`
}


type LookupServerResult struct {
	AdministratorLogin       *string                    `pulumi:"administratorLogin"`
	AvailabilityZone         *string                    `pulumi:"availabilityZone"`
	Backup                   *BackupResponse            `pulumi:"backup"`
	FullyQualifiedDomainName string                     `pulumi:"fullyQualifiedDomainName"`
	HighAvailability         *HighAvailabilityResponse  `pulumi:"highAvailability"`
	Id                       string                     `pulumi:"id"`
	Identity                 *IdentityResponse          `pulumi:"identity"`
	Location                 string                     `pulumi:"location"`
	MaintenanceWindow        *MaintenanceWindowResponse `pulumi:"maintenanceWindow"`
	MinorVersion             string                     `pulumi:"minorVersion"`
	Name                     string                     `pulumi:"name"`
	Network                  *NetworkResponse           `pulumi:"network"`
	Sku                      *SkuResponse               `pulumi:"sku"`
	State                    string                     `pulumi:"state"`
	Storage                  *StorageResponse           `pulumi:"storage"`
	SystemData               SystemDataResponse         `pulumi:"systemData"`
	Tags                     map[string]string          `pulumi:"tags"`
	Type                     string                     `pulumi:"type"`
	Version                  *string                    `pulumi:"version"`
}

func LookupServerOutput(ctx *pulumi.Context, args LookupServerOutputArgs, opts ...pulumi.InvokeOption) LookupServerResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupServerResult, error) {
			args := v.(LookupServerArgs)
			r, err := LookupServer(ctx, &args, opts...)
			var s LookupServerResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupServerResultOutput)
}

type LookupServerOutputArgs struct {
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	ServerName        pulumi.StringInput `pulumi:"serverName"`
}

func (LookupServerOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupServerArgs)(nil)).Elem()
}


type LookupServerResultOutput struct{ *pulumi.OutputState }

func (LookupServerResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupServerResult)(nil)).Elem()
}

func (o LookupServerResultOutput) ToLookupServerResultOutput() LookupServerResultOutput {
	return o
}

func (o LookupServerResultOutput) ToLookupServerResultOutputWithContext(ctx context.Context) LookupServerResultOutput {
	return o
}

func (o LookupServerResultOutput) AdministratorLogin() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupServerResult) *string { return v.AdministratorLogin }).(pulumi.StringPtrOutput)
}

func (o LookupServerResultOutput) AvailabilityZone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupServerResult) *string { return v.AvailabilityZone }).(pulumi.StringPtrOutput)
}

func (o LookupServerResultOutput) Backup() BackupResponsePtrOutput {
	return o.ApplyT(func(v LookupServerResult) *BackupResponse { return v.Backup }).(BackupResponsePtrOutput)
}

func (o LookupServerResultOutput) FullyQualifiedDomainName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerResult) string { return v.FullyQualifiedDomainName }).(pulumi.StringOutput)
}

func (o LookupServerResultOutput) HighAvailability() HighAvailabilityResponsePtrOutput {
	return o.ApplyT(func(v LookupServerResult) *HighAvailabilityResponse { return v.HighAvailability }).(HighAvailabilityResponsePtrOutput)
}

func (o LookupServerResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupServerResultOutput) Identity() IdentityResponsePtrOutput {
	return o.ApplyT(func(v LookupServerResult) *IdentityResponse { return v.Identity }).(IdentityResponsePtrOutput)
}

func (o LookupServerResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupServerResultOutput) MaintenanceWindow() MaintenanceWindowResponsePtrOutput {
	return o.ApplyT(func(v LookupServerResult) *MaintenanceWindowResponse { return v.MaintenanceWindow }).(MaintenanceWindowResponsePtrOutput)
}

func (o LookupServerResultOutput) MinorVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerResult) string { return v.MinorVersion }).(pulumi.StringOutput)
}

func (o LookupServerResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupServerResultOutput) Network() NetworkResponsePtrOutput {
	return o.ApplyT(func(v LookupServerResult) *NetworkResponse { return v.Network }).(NetworkResponsePtrOutput)
}

func (o LookupServerResultOutput) Sku() SkuResponsePtrOutput {
	return o.ApplyT(func(v LookupServerResult) *SkuResponse { return v.Sku }).(SkuResponsePtrOutput)
}

func (o LookupServerResultOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerResult) string { return v.State }).(pulumi.StringOutput)
}

func (o LookupServerResultOutput) Storage() StorageResponsePtrOutput {
	return o.ApplyT(func(v LookupServerResult) *StorageResponse { return v.Storage }).(StorageResponsePtrOutput)
}

func (o LookupServerResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupServerResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupServerResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupServerResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupServerResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupServerResultOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupServerResult) *string { return v.Version }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupServerResultOutput{})
}
