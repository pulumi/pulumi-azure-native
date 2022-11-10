


package v20210410privatepreview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupServer(ctx *pulumi.Context, args *LookupServerArgs, opts ...pulumi.InvokeOption) (*LookupServerResult, error) {
	var rv LookupServerResult
	err := ctx.Invoke("azure-native:dbforpostgresql/v20210410privatepreview:getServer", args, &rv, opts...)
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
	AdministratorLogin       *string                                           `pulumi:"administratorLogin"`
	AvailabilityZone         *string                                           `pulumi:"availabilityZone"`
	ByokEnforcement          string                                            `pulumi:"byokEnforcement"`
	DelegatedSubnetArguments *ServerPropertiesResponseDelegatedSubnetArguments `pulumi:"delegatedSubnetArguments"`
	DisplayName              *string                                           `pulumi:"displayName"`
	FullyQualifiedDomainName string                                            `pulumi:"fullyQualifiedDomainName"`
	HaEnabled                *string                                           `pulumi:"haEnabled"`
	HaState                  string                                            `pulumi:"haState"`
	Id                       string                                            `pulumi:"id"`
	Identity                 *IdentityResponse                                 `pulumi:"identity"`
	Location                 string                                            `pulumi:"location"`
	MaintenanceWindow        *MaintenanceWindowResponse                        `pulumi:"maintenanceWindow"`
	Name                     string                                            `pulumi:"name"`
	PointInTimeUTC           *string                                           `pulumi:"pointInTimeUTC"`
	PrivateDnsZoneArguments  *ServerPropertiesResponsePrivateDnsZoneArguments  `pulumi:"privateDnsZoneArguments"`
	PublicNetworkAccess      string                                            `pulumi:"publicNetworkAccess"`
	Sku                      *SkuResponse                                      `pulumi:"sku"`
	SourceResourceGroupName  *string                                           `pulumi:"sourceResourceGroupName"`
	SourceServerName         *string                                           `pulumi:"sourceServerName"`
	SourceSubscriptionId     *string                                           `pulumi:"sourceSubscriptionId"`
	StandbyAvailabilityZone  string                                            `pulumi:"standbyAvailabilityZone"`
	State                    string                                            `pulumi:"state"`
	StorageProfile           *StorageProfileResponse                           `pulumi:"storageProfile"`
	Tags                     map[string]string                                 `pulumi:"tags"`
	Type                     string                                            `pulumi:"type"`
	Version                  *string                                           `pulumi:"version"`
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

func (o LookupServerResultOutput) DelegatedSubnetArguments() ServerPropertiesResponseDelegatedSubnetArgumentsPtrOutput {
	return o.ApplyT(func(v LookupServerResult) *ServerPropertiesResponseDelegatedSubnetArguments {
		return v.DelegatedSubnetArguments
	}).(ServerPropertiesResponseDelegatedSubnetArgumentsPtrOutput)
}

func (o LookupServerResultOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupServerResult) *string { return v.DisplayName }).(pulumi.StringPtrOutput)
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

func (o LookupServerResultOutput) PointInTimeUTC() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupServerResult) *string { return v.PointInTimeUTC }).(pulumi.StringPtrOutput)
}

func (o LookupServerResultOutput) PrivateDnsZoneArguments() ServerPropertiesResponsePrivateDnsZoneArgumentsPtrOutput {
	return o.ApplyT(func(v LookupServerResult) *ServerPropertiesResponsePrivateDnsZoneArguments {
		return v.PrivateDnsZoneArguments
	}).(ServerPropertiesResponsePrivateDnsZoneArgumentsPtrOutput)
}

func (o LookupServerResultOutput) PublicNetworkAccess() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerResult) string { return v.PublicNetworkAccess }).(pulumi.StringOutput)
}

func (o LookupServerResultOutput) Sku() SkuResponsePtrOutput {
	return o.ApplyT(func(v LookupServerResult) *SkuResponse { return v.Sku }).(SkuResponsePtrOutput)
}

func (o LookupServerResultOutput) SourceResourceGroupName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupServerResult) *string { return v.SourceResourceGroupName }).(pulumi.StringPtrOutput)
}

func (o LookupServerResultOutput) SourceServerName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupServerResult) *string { return v.SourceServerName }).(pulumi.StringPtrOutput)
}

func (o LookupServerResultOutput) SourceSubscriptionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupServerResult) *string { return v.SourceSubscriptionId }).(pulumi.StringPtrOutput)
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
