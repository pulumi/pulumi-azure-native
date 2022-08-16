


package v20200701preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupServer(ctx *pulumi.Context, args *LookupServerArgs, opts ...pulumi.InvokeOption) (*LookupServerResult, error) {
	var rv LookupServerResult
	err := ctx.Invoke("azure-native:dbformysql/v20200701preview:getServer", args, &rv, opts...)
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
	AdministratorLogin       *string                           `pulumi:"administratorLogin"`
	AvailabilityZone         *string                           `pulumi:"availabilityZone"`
	ByokEnforcement          string                            `pulumi:"byokEnforcement"`
	DelegatedSubnetArguments *DelegatedSubnetArgumentsResponse `pulumi:"delegatedSubnetArguments"`
	EarliestRestoreDate      string                            `pulumi:"earliestRestoreDate"`
	FullyQualifiedDomainName string                            `pulumi:"fullyQualifiedDomainName"`
	HaEnabled                *string                           `pulumi:"haEnabled"`
	HaState                  string                            `pulumi:"haState"`
	Id                       string                            `pulumi:"id"`
	Identity                 *IdentityResponse                 `pulumi:"identity"`
	Location                 string                            `pulumi:"location"`
	MaintenanceWindow        *MaintenanceWindowResponse        `pulumi:"maintenanceWindow"`
	Name                     string                            `pulumi:"name"`
	PrivateDnsZoneArguments  *PrivateDnsZoneArgumentsResponse  `pulumi:"privateDnsZoneArguments"`
	PublicNetworkAccess      string                            `pulumi:"publicNetworkAccess"`
	ReplicaCapacity          int                               `pulumi:"replicaCapacity"`
	ReplicationRole          *string                           `pulumi:"replicationRole"`
	Sku                      *SkuResponse                      `pulumi:"sku"`
	SourceServerId           *string                           `pulumi:"sourceServerId"`
	SslEnforcement           *string                           `pulumi:"sslEnforcement"`
	StandbyAvailabilityZone  string                            `pulumi:"standbyAvailabilityZone"`
	State                    string                            `pulumi:"state"`
	StorageProfile           *StorageProfileResponse           `pulumi:"storageProfile"`
	Tags                     map[string]string                 `pulumi:"tags"`
	Type                     string                            `pulumi:"type"`
	Version                  *string                           `pulumi:"version"`
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

func (o LookupServerResultOutput) ByokEnforcement() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerResult) string { return v.ByokEnforcement }).(pulumi.StringOutput)
}

func (o LookupServerResultOutput) DelegatedSubnetArguments() DelegatedSubnetArgumentsResponsePtrOutput {
	return o.ApplyT(func(v LookupServerResult) *DelegatedSubnetArgumentsResponse { return v.DelegatedSubnetArguments }).(DelegatedSubnetArgumentsResponsePtrOutput)
}

func (o LookupServerResultOutput) EarliestRestoreDate() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerResult) string { return v.EarliestRestoreDate }).(pulumi.StringOutput)
}

func (o LookupServerResultOutput) FullyQualifiedDomainName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerResult) string { return v.FullyQualifiedDomainName }).(pulumi.StringOutput)
}

func (o LookupServerResultOutput) HaEnabled() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupServerResult) *string { return v.HaEnabled }).(pulumi.StringPtrOutput)
}

func (o LookupServerResultOutput) HaState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerResult) string { return v.HaState }).(pulumi.StringOutput)
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

func (o LookupServerResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupServerResultOutput) PrivateDnsZoneArguments() PrivateDnsZoneArgumentsResponsePtrOutput {
	return o.ApplyT(func(v LookupServerResult) *PrivateDnsZoneArgumentsResponse { return v.PrivateDnsZoneArguments }).(PrivateDnsZoneArgumentsResponsePtrOutput)
}

func (o LookupServerResultOutput) PublicNetworkAccess() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerResult) string { return v.PublicNetworkAccess }).(pulumi.StringOutput)
}

func (o LookupServerResultOutput) ReplicaCapacity() pulumi.IntOutput {
	return o.ApplyT(func(v LookupServerResult) int { return v.ReplicaCapacity }).(pulumi.IntOutput)
}

func (o LookupServerResultOutput) ReplicationRole() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupServerResult) *string { return v.ReplicationRole }).(pulumi.StringPtrOutput)
}

func (o LookupServerResultOutput) Sku() SkuResponsePtrOutput {
	return o.ApplyT(func(v LookupServerResult) *SkuResponse { return v.Sku }).(SkuResponsePtrOutput)
}

func (o LookupServerResultOutput) SourceServerId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupServerResult) *string { return v.SourceServerId }).(pulumi.StringPtrOutput)
}

func (o LookupServerResultOutput) SslEnforcement() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupServerResult) *string { return v.SslEnforcement }).(pulumi.StringPtrOutput)
}

func (o LookupServerResultOutput) StandbyAvailabilityZone() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerResult) string { return v.StandbyAvailabilityZone }).(pulumi.StringOutput)
}

func (o LookupServerResultOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerResult) string { return v.State }).(pulumi.StringOutput)
}

func (o LookupServerResultOutput) StorageProfile() StorageProfileResponsePtrOutput {
	return o.ApplyT(func(v LookupServerResult) *StorageProfileResponse { return v.StorageProfile }).(StorageProfileResponsePtrOutput)
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
