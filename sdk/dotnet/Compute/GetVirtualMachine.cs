// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Compute
{
    public static class GetVirtualMachine
    {
        /// <summary>
        /// Retrieves information about the model view or the instance view of a virtual machine.
        /// 
        /// Uses Azure REST API version 2024-11-01.
        /// 
        /// Other available API versions: 2022-08-01, 2022-11-01, 2023-03-01, 2023-07-01, 2023-09-01, 2024-03-01, 2024-07-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native compute [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Task<GetVirtualMachineResult> InvokeAsync(GetVirtualMachineArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetVirtualMachineResult>("azure-native:compute:getVirtualMachine", args ?? new GetVirtualMachineArgs(), options.WithDefaults());

        /// <summary>
        /// Retrieves information about the model view or the instance view of a virtual machine.
        /// 
        /// Uses Azure REST API version 2024-11-01.
        /// 
        /// Other available API versions: 2022-08-01, 2022-11-01, 2023-03-01, 2023-07-01, 2023-09-01, 2024-03-01, 2024-07-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native compute [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetVirtualMachineResult> Invoke(GetVirtualMachineInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetVirtualMachineResult>("azure-native:compute:getVirtualMachine", args ?? new GetVirtualMachineInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Retrieves information about the model view or the instance view of a virtual machine.
        /// 
        /// Uses Azure REST API version 2024-11-01.
        /// 
        /// Other available API versions: 2022-08-01, 2022-11-01, 2023-03-01, 2023-07-01, 2023-09-01, 2024-03-01, 2024-07-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native compute [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetVirtualMachineResult> Invoke(GetVirtualMachineInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetVirtualMachineResult>("azure-native:compute:getVirtualMachine", args ?? new GetVirtualMachineInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetVirtualMachineArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The expand expression to apply on the operation. 'InstanceView' retrieves a snapshot of the runtime properties of the virtual machine that is managed by the platform and can change outside of control plane operations. 'UserData' retrieves the UserData property as part of the VM model view that was provided by the user during the VM Create/Update operation.
        /// </summary>
        [Input("expand")]
        public string? Expand { get; set; }

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the virtual machine.
        /// </summary>
        [Input("vmName", required: true)]
        public string VmName { get; set; } = null!;

        public GetVirtualMachineArgs()
        {
        }
        public static new GetVirtualMachineArgs Empty => new GetVirtualMachineArgs();
    }

    public sealed class GetVirtualMachineInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The expand expression to apply on the operation. 'InstanceView' retrieves a snapshot of the runtime properties of the virtual machine that is managed by the platform and can change outside of control plane operations. 'UserData' retrieves the UserData property as part of the VM model view that was provided by the user during the VM Create/Update operation.
        /// </summary>
        [Input("expand")]
        public Input<string>? Expand { get; set; }

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the virtual machine.
        /// </summary>
        [Input("vmName", required: true)]
        public Input<string> VmName { get; set; } = null!;

        public GetVirtualMachineInvokeArgs()
        {
        }
        public static new GetVirtualMachineInvokeArgs Empty => new GetVirtualMachineInvokeArgs();
    }


    [OutputType]
    public sealed class GetVirtualMachineResult
    {
        /// <summary>
        /// Specifies additional capabilities enabled or disabled on the virtual machine.
        /// </summary>
        public readonly Outputs.AdditionalCapabilitiesResponse? AdditionalCapabilities;
        /// <summary>
        /// Specifies the gallery applications that should be made available to the VM/VMSS.
        /// </summary>
        public readonly Outputs.ApplicationProfileResponse? ApplicationProfile;
        /// <summary>
        /// Specifies information about the availability set that the virtual machine should be assigned to. Virtual machines specified in the same availability set are allocated to different nodes to maximize availability. For more information about availability sets, see [Availability sets overview](https://docs.microsoft.com/azure/virtual-machines/availability-set-overview). For more information on Azure planned maintenance, see [Maintenance and updates for Virtual Machines in Azure](https://docs.microsoft.com/azure/virtual-machines/maintenance-and-updates). Currently, a VM can only be added to availability set at creation time. The availability set to which the VM is being added should be under the same resource group as the availability set resource. An existing VM cannot be added to an availability set. This property cannot exist along with a non-null properties.virtualMachineScaleSet reference.
        /// </summary>
        public readonly Outputs.SubResourceResponse? AvailabilitySet;
        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        public readonly string AzureApiVersion;
        /// <summary>
        /// Specifies the billing related details of a Azure Spot virtual machine. Minimum api-version: 2019-03-01.
        /// </summary>
        public readonly Outputs.BillingProfileResponse? BillingProfile;
        /// <summary>
        /// Specifies information about the capacity reservation that is used to allocate virtual machine. Minimum api-version: 2021-04-01.
        /// </summary>
        public readonly Outputs.CapacityReservationProfileResponse? CapacityReservation;
        /// <summary>
        /// Specifies the boot diagnostic settings state. Minimum api-version: 2015-06-15.
        /// </summary>
        public readonly Outputs.DiagnosticsProfileResponse? DiagnosticsProfile;
        /// <summary>
        /// Etag is property returned in Create/Update/Get response of the VM, so that customer can supply it in the header to ensure optimistic updates.
        /// </summary>
        public readonly string Etag;
        /// <summary>
        /// Specifies the eviction policy for the Azure Spot virtual machine and Azure Spot scale set. For Azure Spot virtual machines, both 'Deallocate' and 'Delete' are supported and the minimum api-version is 2019-03-01. For Azure Spot scale sets, both 'Deallocate' and 'Delete' are supported and the minimum api-version is 2017-10-30-preview.
        /// </summary>
        public readonly string? EvictionPolicy;
        /// <summary>
        /// The extended location of the Virtual Machine.
        /// </summary>
        public readonly Outputs.ExtendedLocationResponse? ExtendedLocation;
        /// <summary>
        /// Specifies the time alloted for all extensions to start. The time duration should be between 15 minutes and 120 minutes (inclusive) and should be specified in ISO 8601 format. The default value is 90 minutes (PT1H30M). Minimum api-version: 2020-06-01.
        /// </summary>
        public readonly string? ExtensionsTimeBudget;
        /// <summary>
        /// Specifies the hardware settings for the virtual machine.
        /// </summary>
        public readonly Outputs.HardwareProfileResponse? HardwareProfile;
        /// <summary>
        /// Specifies information about the dedicated host that the virtual machine resides in. Minimum api-version: 2018-10-01.
        /// </summary>
        public readonly Outputs.SubResourceResponse? Host;
        /// <summary>
        /// Specifies information about the dedicated host group that the virtual machine resides in. **Note:** User cannot specify both host and hostGroup properties. Minimum api-version: 2020-06-01.
        /// </summary>
        public readonly Outputs.SubResourceResponse? HostGroup;
        /// <summary>
        /// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The identity of the virtual machine, if configured.
        /// </summary>
        public readonly Outputs.VirtualMachineIdentityResponse? Identity;
        /// <summary>
        /// The virtual machine instance view.
        /// </summary>
        public readonly Outputs.VirtualMachineInstanceViewResponse InstanceView;
        /// <summary>
        /// Specifies that the image or disk that is being used was licensed on-premises. &lt;br&gt;&lt;br&gt; Possible values for Windows Server operating system are: &lt;br&gt;&lt;br&gt; Windows_Client &lt;br&gt;&lt;br&gt; Windows_Server &lt;br&gt;&lt;br&gt; Possible values for Linux Server operating system are: &lt;br&gt;&lt;br&gt; RHEL_BYOS (for RHEL) &lt;br&gt;&lt;br&gt; SLES_BYOS (for SUSE) &lt;br&gt;&lt;br&gt; For more information, see [Azure Hybrid Use Benefit for Windows Server](https://docs.microsoft.com/azure/virtual-machines/windows/hybrid-use-benefit-licensing) &lt;br&gt;&lt;br&gt; [Azure Hybrid Use Benefit for Linux Server](https://docs.microsoft.com/azure/virtual-machines/linux/azure-hybrid-benefit-linux) &lt;br&gt;&lt;br&gt; Minimum api-version: 2015-06-15
        /// </summary>
        public readonly string? LicenseType;
        /// <summary>
        /// The geo-location where the resource lives
        /// </summary>
        public readonly string Location;
        /// <summary>
        /// ManagedBy is set to Virtual Machine Scale Set(VMSS) flex ARM resourceID, if the VM is part of the VMSS. This property is used by platform for internal resource group delete optimization.
        /// </summary>
        public readonly string ManagedBy;
        /// <summary>
        /// The name of the resource
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Specifies the network interfaces of the virtual machine.
        /// </summary>
        public readonly Outputs.NetworkProfileResponse? NetworkProfile;
        /// <summary>
        /// Specifies the operating system settings used while creating the virtual machine. Some of the settings cannot be changed once VM is provisioned.
        /// </summary>
        public readonly Outputs.OSProfileResponse? OsProfile;
        /// <summary>
        /// Placement section specifies the user-defined constraints for virtual machine hardware placement. This property cannot be changed once VM is provisioned. Minimum api-version: 2024-11-01.
        /// </summary>
        public readonly Outputs.PlacementResponse? Placement;
        /// <summary>
        /// Specifies information about the marketplace image used to create the virtual machine. This element is only used for marketplace images. Before you can use a marketplace image from an API, you must enable the image for programmatic use.  In the Azure portal, find the marketplace image that you want to use and then click **Want to deploy programmatically, Get Started -&gt;**. Enter any required information and then click **Save**.
        /// </summary>
        public readonly Outputs.PlanResponse? Plan;
        /// <summary>
        /// Specifies the scale set logical fault domain into which the Virtual Machine will be created. By default, the Virtual Machine will by automatically assigned to a fault domain that best maintains balance across available fault domains. This is applicable only if the 'virtualMachineScaleSet' property of this Virtual Machine is set. The Virtual Machine Scale Set that is referenced, must have 'platformFaultDomainCount' greater than 1. This property cannot be updated once the Virtual Machine is created. Fault domain assignment can be viewed in the Virtual Machine Instance View. Minimum api‐version: 2020‐12‐01.
        /// </summary>
        public readonly int? PlatformFaultDomain;
        /// <summary>
        /// Specifies the priority for the virtual machine. Minimum api-version: 2019-03-01
        /// </summary>
        public readonly string? Priority;
        /// <summary>
        /// The provisioning state, which only appears in the response.
        /// </summary>
        public readonly string ProvisioningState;
        /// <summary>
        /// Specifies information about the proximity placement group that the virtual machine should be assigned to. Minimum api-version: 2018-04-01.
        /// </summary>
        public readonly Outputs.SubResourceResponse? ProximityPlacementGroup;
        /// <summary>
        /// The virtual machine child extension resources.
        /// </summary>
        public readonly ImmutableArray<Outputs.VirtualMachineExtensionResponse> Resources;
        /// <summary>
        /// Specifies Redeploy, Reboot and ScheduledEventsAdditionalPublishingTargets Scheduled Event related configurations for the virtual machine.
        /// </summary>
        public readonly Outputs.ScheduledEventsPolicyResponse? ScheduledEventsPolicy;
        /// <summary>
        /// Specifies Scheduled Event related configurations.
        /// </summary>
        public readonly Outputs.ScheduledEventsProfileResponse? ScheduledEventsProfile;
        /// <summary>
        /// Specifies the Security related profile settings for the virtual machine.
        /// </summary>
        public readonly Outputs.SecurityProfileResponse? SecurityProfile;
        /// <summary>
        /// Specifies the storage settings for the virtual machine disks.
        /// </summary>
        public readonly Outputs.StorageProfileResponse? StorageProfile;
        /// <summary>
        /// Azure Resource Manager metadata containing createdBy and modifiedBy information.
        /// </summary>
        public readonly Outputs.SystemDataResponse SystemData;
        /// <summary>
        /// Resource tags.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Tags;
        /// <summary>
        /// Specifies the time at which the Virtual Machine resource was created. Minimum api-version: 2021-11-01.
        /// </summary>
        public readonly string TimeCreated;
        /// <summary>
        /// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// UserData for the VM, which must be base-64 encoded. Customer should not pass any secrets in here. Minimum api-version: 2021-03-01.
        /// </summary>
        public readonly string? UserData;
        /// <summary>
        /// Specifies information about the virtual machine scale set that the virtual machine should be assigned to. Virtual machines specified in the same virtual machine scale set are allocated to different nodes to maximize availability. Currently, a VM can only be added to virtual machine scale set at creation time. An existing VM cannot be added to a virtual machine scale set. This property cannot exist along with a non-null properties.availabilitySet reference. Minimum api‐version: 2019‐03‐01.
        /// </summary>
        public readonly Outputs.SubResourceResponse? VirtualMachineScaleSet;
        /// <summary>
        /// Specifies the VM unique ID which is a 128-bits identifier that is encoded and stored in all Azure IaaS VMs SMBIOS and can be read using platform BIOS commands.
        /// </summary>
        public readonly string VmId;
        /// <summary>
        /// The availability zones.
        /// </summary>
        public readonly ImmutableArray<string> Zones;

        [OutputConstructor]
        private GetVirtualMachineResult(
            Outputs.AdditionalCapabilitiesResponse? additionalCapabilities,

            Outputs.ApplicationProfileResponse? applicationProfile,

            Outputs.SubResourceResponse? availabilitySet,

            string azureApiVersion,

            Outputs.BillingProfileResponse? billingProfile,

            Outputs.CapacityReservationProfileResponse? capacityReservation,

            Outputs.DiagnosticsProfileResponse? diagnosticsProfile,

            string etag,

            string? evictionPolicy,

            Outputs.ExtendedLocationResponse? extendedLocation,

            string? extensionsTimeBudget,

            Outputs.HardwareProfileResponse? hardwareProfile,

            Outputs.SubResourceResponse? host,

            Outputs.SubResourceResponse? hostGroup,

            string id,

            Outputs.VirtualMachineIdentityResponse? identity,

            Outputs.VirtualMachineInstanceViewResponse instanceView,

            string? licenseType,

            string location,

            string managedBy,

            string name,

            Outputs.NetworkProfileResponse? networkProfile,

            Outputs.OSProfileResponse? osProfile,

            Outputs.PlacementResponse? placement,

            Outputs.PlanResponse? plan,

            int? platformFaultDomain,

            string? priority,

            string provisioningState,

            Outputs.SubResourceResponse? proximityPlacementGroup,

            ImmutableArray<Outputs.VirtualMachineExtensionResponse> resources,

            Outputs.ScheduledEventsPolicyResponse? scheduledEventsPolicy,

            Outputs.ScheduledEventsProfileResponse? scheduledEventsProfile,

            Outputs.SecurityProfileResponse? securityProfile,

            Outputs.StorageProfileResponse? storageProfile,

            Outputs.SystemDataResponse systemData,

            ImmutableDictionary<string, string>? tags,

            string timeCreated,

            string type,

            string? userData,

            Outputs.SubResourceResponse? virtualMachineScaleSet,

            string vmId,

            ImmutableArray<string> zones)
        {
            AdditionalCapabilities = additionalCapabilities;
            ApplicationProfile = applicationProfile;
            AvailabilitySet = availabilitySet;
            AzureApiVersion = azureApiVersion;
            BillingProfile = billingProfile;
            CapacityReservation = capacityReservation;
            DiagnosticsProfile = diagnosticsProfile;
            Etag = etag;
            EvictionPolicy = evictionPolicy;
            ExtendedLocation = extendedLocation;
            ExtensionsTimeBudget = extensionsTimeBudget;
            HardwareProfile = hardwareProfile;
            Host = host;
            HostGroup = hostGroup;
            Id = id;
            Identity = identity;
            InstanceView = instanceView;
            LicenseType = licenseType;
            Location = location;
            ManagedBy = managedBy;
            Name = name;
            NetworkProfile = networkProfile;
            OsProfile = osProfile;
            Placement = placement;
            Plan = plan;
            PlatformFaultDomain = platformFaultDomain;
            Priority = priority;
            ProvisioningState = provisioningState;
            ProximityPlacementGroup = proximityPlacementGroup;
            Resources = resources;
            ScheduledEventsPolicy = scheduledEventsPolicy;
            ScheduledEventsProfile = scheduledEventsProfile;
            SecurityProfile = securityProfile;
            StorageProfile = storageProfile;
            SystemData = systemData;
            Tags = tags;
            TimeCreated = timeCreated;
            Type = type;
            UserData = userData;
            VirtualMachineScaleSet = virtualMachineScaleSet;
            VmId = vmId;
            Zones = zones;
        }
    }
}
