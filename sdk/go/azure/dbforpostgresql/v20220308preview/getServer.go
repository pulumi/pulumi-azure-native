


package v20220308preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupServer(ctx *pulumi.Context, args *LookupServerArgs, opts ...pulumi.InvokeOption) (*LookupServerResult, error) {
	var rv LookupServerResult
	err := ctx.Invoke("azure-native:dbforpostgresql/v20220308preview:getServer", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupServerArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ServerName        string `pulumi:"serverName"`
}


type LookupServerResult struct {
	AdministratorLogin       *string                       `pulumi:"administratorLogin"`
	AuthConfig               *AuthConfigResponse           `pulumi:"authConfig"`
	AvailabilityZone         *string                       `pulumi:"availabilityZone"`
	Backup                   *BackupResponse               `pulumi:"backup"`
	DataEncryption           *DataEncryptionResponse       `pulumi:"dataEncryption"`
	FullyQualifiedDomainName string                        `pulumi:"fullyQualifiedDomainName"`
	HighAvailability         *HighAvailabilityResponse     `pulumi:"highAvailability"`
	Id                       string                        `pulumi:"id"`
	Identity                 *UserAssignedIdentityResponse `pulumi:"identity"`
	Location                 string                        `pulumi:"location"`
	MaintenanceWindow        *MaintenanceWindowResponse    `pulumi:"maintenanceWindow"`
	MinorVersion             string                        `pulumi:"minorVersion"`
	Name                     string                        `pulumi:"name"`
	Network                  *NetworkResponse              `pulumi:"network"`
	ReplicaCapacity          *int                          `pulumi:"replicaCapacity"`
	ReplicationRole          *string                       `pulumi:"replicationRole"`
	Sku                      *PostgreSqlSkuResponse        `pulumi:"sku"`
	State                    string                        `pulumi:"state"`
	Storage                  *StorageResponse              `pulumi:"storage"`
	SystemData               SystemDataResponse            `pulumi:"systemData"`
	Tags                     map[string]string             `pulumi:"tags"`
	Type                     string                        `pulumi:"type"`
	Version                  *string                       `pulumi:"version"`
}


func (val *LookupServerResult) Defaults() *LookupServerResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.AuthConfig = tmp.AuthConfig.Defaults()

	if isZero(tmp.AvailabilityZone) {
		availabilityZone_ := ""
		tmp.AvailabilityZone = &availabilityZone_
	}
	tmp.Backup = tmp.Backup.Defaults()

	tmp.HighAvailability = tmp.HighAvailability.Defaults()

	tmp.MaintenanceWindow = tmp.MaintenanceWindow.Defaults()

	tmp.Network = tmp.Network.Defaults()

	return &tmp
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

func (o LookupServerResultOutput) AuthConfig() AuthConfigResponsePtrOutput {
	return o.ApplyT(func(v LookupServerResult) *AuthConfigResponse { return v.AuthConfig }).(AuthConfigResponsePtrOutput)
}

func (o LookupServerResultOutput) AvailabilityZone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupServerResult) *string { return v.AvailabilityZone }).(pulumi.StringPtrOutput)
}

func (o LookupServerResultOutput) Backup() BackupResponsePtrOutput {
	return o.ApplyT(func(v LookupServerResult) *BackupResponse { return v.Backup }).(BackupResponsePtrOutput)
}

func (o LookupServerResultOutput) DataEncryption() DataEncryptionResponsePtrOutput {
	return o.ApplyT(func(v LookupServerResult) *DataEncryptionResponse { return v.DataEncryption }).(DataEncryptionResponsePtrOutput)
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

func (o LookupServerResultOutput) Identity() UserAssignedIdentityResponsePtrOutput {
	return o.ApplyT(func(v LookupServerResult) *UserAssignedIdentityResponse { return v.Identity }).(UserAssignedIdentityResponsePtrOutput)
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

func (o LookupServerResultOutput) ReplicaCapacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupServerResult) *int { return v.ReplicaCapacity }).(pulumi.IntPtrOutput)
}

func (o LookupServerResultOutput) ReplicationRole() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupServerResult) *string { return v.ReplicationRole }).(pulumi.StringPtrOutput)
}

func (o LookupServerResultOutput) Sku() PostgreSqlSkuResponsePtrOutput {
	return o.ApplyT(func(v LookupServerResult) *PostgreSqlSkuResponse { return v.Sku }).(PostgreSqlSkuResponsePtrOutput)
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
