


package v20210701

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupVirtualMachine(ctx *pulumi.Context, args *LookupVirtualMachineArgs, opts ...pulumi.InvokeOption) (*LookupVirtualMachineResult, error) {
	var rv LookupVirtualMachineResult
	err := ctx.Invoke("azure-native:compute/v20210701:getVirtualMachine", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupVirtualMachineArgs struct {
	Expand            *string `pulumi:"expand"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
	VmName            string  `pulumi:"vmName"`
}


type LookupVirtualMachineResult struct {
	AdditionalCapabilities  *AdditionalCapabilitiesResponse     `pulumi:"additionalCapabilities"`
	ApplicationProfile      *ApplicationProfileResponse         `pulumi:"applicationProfile"`
	AvailabilitySet         *SubResourceResponse                `pulumi:"availabilitySet"`
	BillingProfile          *BillingProfileResponse             `pulumi:"billingProfile"`
	CapacityReservation     *CapacityReservationProfileResponse `pulumi:"capacityReservation"`
	DiagnosticsProfile      *DiagnosticsProfileResponse         `pulumi:"diagnosticsProfile"`
	EvictionPolicy          *string                             `pulumi:"evictionPolicy"`
	ExtendedLocation        *ExtendedLocationResponse           `pulumi:"extendedLocation"`
	ExtensionsTimeBudget    *string                             `pulumi:"extensionsTimeBudget"`
	HardwareProfile         *HardwareProfileResponse            `pulumi:"hardwareProfile"`
	Host                    *SubResourceResponse                `pulumi:"host"`
	HostGroup               *SubResourceResponse                `pulumi:"hostGroup"`
	Id                      string                              `pulumi:"id"`
	Identity                *VirtualMachineIdentityResponse     `pulumi:"identity"`
	InstanceView            VirtualMachineInstanceViewResponse  `pulumi:"instanceView"`
	LicenseType             *string                             `pulumi:"licenseType"`
	Location                string                              `pulumi:"location"`
	Name                    string                              `pulumi:"name"`
	NetworkProfile          *NetworkProfileResponse             `pulumi:"networkProfile"`
	OsProfile               *OSProfileResponse                  `pulumi:"osProfile"`
	Plan                    *PlanResponse                       `pulumi:"plan"`
	PlatformFaultDomain     *int                                `pulumi:"platformFaultDomain"`
	Priority                *string                             `pulumi:"priority"`
	ProvisioningState       string                              `pulumi:"provisioningState"`
	ProximityPlacementGroup *SubResourceResponse                `pulumi:"proximityPlacementGroup"`
	Resources               []VirtualMachineExtensionResponse   `pulumi:"resources"`
	ScheduledEventsProfile  *ScheduledEventsProfileResponse     `pulumi:"scheduledEventsProfile"`
	SecurityProfile         *SecurityProfileResponse            `pulumi:"securityProfile"`
	StorageProfile          *StorageProfileResponse             `pulumi:"storageProfile"`
	Tags                    map[string]string                   `pulumi:"tags"`
	Type                    string                              `pulumi:"type"`
	UserData                *string                             `pulumi:"userData"`
	VirtualMachineScaleSet  *SubResourceResponse                `pulumi:"virtualMachineScaleSet"`
	VmId                    string                              `pulumi:"vmId"`
	Zones                   []string                            `pulumi:"zones"`
}
