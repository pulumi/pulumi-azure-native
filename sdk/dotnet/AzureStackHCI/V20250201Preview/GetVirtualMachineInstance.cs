// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.AzureStackHCI.V20250201Preview
{
    public static class GetVirtualMachineInstance
    {
        /// <summary>
        /// Gets a virtual machine instance
        /// </summary>
        public static Task<GetVirtualMachineInstanceResult> InvokeAsync(GetVirtualMachineInstanceArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetVirtualMachineInstanceResult>("azure-native:azurestackhci/v20250201preview:getVirtualMachineInstance", args ?? new GetVirtualMachineInstanceArgs(), options.WithDefaults());

        /// <summary>
        /// Gets a virtual machine instance
        /// </summary>
        public static Output<GetVirtualMachineInstanceResult> Invoke(GetVirtualMachineInstanceInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetVirtualMachineInstanceResult>("azure-native:azurestackhci/v20250201preview:getVirtualMachineInstance", args ?? new GetVirtualMachineInstanceInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Gets a virtual machine instance
        /// </summary>
        public static Output<GetVirtualMachineInstanceResult> Invoke(GetVirtualMachineInstanceInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetVirtualMachineInstanceResult>("azure-native:azurestackhci/v20250201preview:getVirtualMachineInstance", args ?? new GetVirtualMachineInstanceInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetVirtualMachineInstanceArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The fully qualified Azure Resource manager identifier of the resource.
        /// </summary>
        [Input("resourceUri", required: true)]
        public string ResourceUri { get; set; } = null!;

        public GetVirtualMachineInstanceArgs()
        {
        }
        public static new GetVirtualMachineInstanceArgs Empty => new GetVirtualMachineInstanceArgs();
    }

    public sealed class GetVirtualMachineInstanceInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The fully qualified Azure Resource manager identifier of the resource.
        /// </summary>
        [Input("resourceUri", required: true)]
        public Input<string> ResourceUri { get; set; } = null!;

        public GetVirtualMachineInstanceInvokeArgs()
        {
        }
        public static new GetVirtualMachineInstanceInvokeArgs Empty => new GetVirtualMachineInstanceInvokeArgs();
    }


    [OutputType]
    public sealed class GetVirtualMachineInstanceResult
    {
        /// <summary>
        /// Boolean indicating whether this is an existing local virtual machine or if one should be created.
        /// </summary>
        public readonly bool? CreateFromLocal;
        /// <summary>
        /// The extendedLocation of the resource.
        /// </summary>
        public readonly Outputs.ExtendedLocationResponse? ExtendedLocation;
        /// <summary>
        /// Guest agent install status.
        /// </summary>
        public readonly Outputs.GuestAgentInstallStatusResponse? GuestAgentInstallStatus;
        /// <summary>
        /// HardwareProfile - Specifies the hardware settings for the virtual machine instance.
        /// </summary>
        public readonly Outputs.VirtualMachineInstancePropertiesHardwareProfileResponse? HardwareProfile;
        /// <summary>
        /// HTTP Proxy configuration for the VM.
        /// </summary>
        public readonly Outputs.HttpProxyConfigurationResponse? HttpProxyConfig;
        /// <summary>
        /// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The managed service identities assigned to this resource.
        /// </summary>
        public readonly Outputs.ManagedServiceIdentityResponse? Identity;
        /// <summary>
        /// The virtual machine instance view.
        /// </summary>
        public readonly Outputs.VirtualMachineInstanceViewResponse InstanceView;
        /// <summary>
        /// The name of the resource
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// NetworkProfile - describes the network configuration the virtual machine instance
        /// </summary>
        public readonly Outputs.VirtualMachineInstancePropertiesNetworkProfileResponse? NetworkProfile;
        /// <summary>
        /// OsProfile - describes the configuration of the operating system and sets login data
        /// </summary>
        public readonly Outputs.VirtualMachineInstancePropertiesOsProfileResponse? OsProfile;
        /// <summary>
        /// Provisioning state of the virtual machine instance.
        /// </summary>
        public readonly string ProvisioningState;
        /// <summary>
        /// Unique identifier defined by ARC to identify the guest of the VM.
        /// </summary>
        public readonly string? ResourceUid;
        /// <summary>
        /// SecurityProfile - Specifies the security settings for the virtual machine instance.
        /// </summary>
        public readonly Outputs.VirtualMachineInstancePropertiesSecurityProfileResponse? SecurityProfile;
        /// <summary>
        /// The observed state of virtual machine instances
        /// </summary>
        public readonly Outputs.VirtualMachineInstanceStatusResponse Status;
        /// <summary>
        /// StorageProfile - contains information about the disks and storage information for the virtual machine instance
        /// </summary>
        public readonly Outputs.VirtualMachineInstancePropertiesStorageProfileResponse? StorageProfile;
        /// <summary>
        /// Azure Resource Manager metadata containing createdBy and modifiedBy information.
        /// </summary>
        public readonly Outputs.SystemDataResponse SystemData;
        /// <summary>
        /// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// Unique identifier for the vm resource.
        /// </summary>
        public readonly string VmId;

        [OutputConstructor]
        private GetVirtualMachineInstanceResult(
            bool? createFromLocal,

            Outputs.ExtendedLocationResponse? extendedLocation,

            Outputs.GuestAgentInstallStatusResponse? guestAgentInstallStatus,

            Outputs.VirtualMachineInstancePropertiesHardwareProfileResponse? hardwareProfile,

            Outputs.HttpProxyConfigurationResponse? httpProxyConfig,

            string id,

            Outputs.ManagedServiceIdentityResponse? identity,

            Outputs.VirtualMachineInstanceViewResponse instanceView,

            string name,

            Outputs.VirtualMachineInstancePropertiesNetworkProfileResponse? networkProfile,

            Outputs.VirtualMachineInstancePropertiesOsProfileResponse? osProfile,

            string provisioningState,

            string? resourceUid,

            Outputs.VirtualMachineInstancePropertiesSecurityProfileResponse? securityProfile,

            Outputs.VirtualMachineInstanceStatusResponse status,

            Outputs.VirtualMachineInstancePropertiesStorageProfileResponse? storageProfile,

            Outputs.SystemDataResponse systemData,

            string type,

            string vmId)
        {
            CreateFromLocal = createFromLocal;
            ExtendedLocation = extendedLocation;
            GuestAgentInstallStatus = guestAgentInstallStatus;
            HardwareProfile = hardwareProfile;
            HttpProxyConfig = httpProxyConfig;
            Id = id;
            Identity = identity;
            InstanceView = instanceView;
            Name = name;
            NetworkProfile = networkProfile;
            OsProfile = osProfile;
            ProvisioningState = provisioningState;
            ResourceUid = resourceUid;
            SecurityProfile = securityProfile;
            Status = status;
            StorageProfile = storageProfile;
            SystemData = systemData;
            Type = type;
            VmId = vmId;
        }
    }
}
