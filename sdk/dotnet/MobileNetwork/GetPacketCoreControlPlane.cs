// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.MobileNetwork
{
    public static class GetPacketCoreControlPlane
    {
        /// <summary>
        /// Gets information about the specified packet core control plane.
        /// 
        /// Uses Azure REST API version 2024-04-01.
        /// 
        /// Other available API versions: 2022-04-01-preview, 2022-11-01, 2023-06-01, 2023-09-01, 2024-02-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native mobilenetwork [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Task<GetPacketCoreControlPlaneResult> InvokeAsync(GetPacketCoreControlPlaneArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetPacketCoreControlPlaneResult>("azure-native:mobilenetwork:getPacketCoreControlPlane", args ?? new GetPacketCoreControlPlaneArgs(), options.WithDefaults());

        /// <summary>
        /// Gets information about the specified packet core control plane.
        /// 
        /// Uses Azure REST API version 2024-04-01.
        /// 
        /// Other available API versions: 2022-04-01-preview, 2022-11-01, 2023-06-01, 2023-09-01, 2024-02-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native mobilenetwork [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetPacketCoreControlPlaneResult> Invoke(GetPacketCoreControlPlaneInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetPacketCoreControlPlaneResult>("azure-native:mobilenetwork:getPacketCoreControlPlane", args ?? new GetPacketCoreControlPlaneInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Gets information about the specified packet core control plane.
        /// 
        /// Uses Azure REST API version 2024-04-01.
        /// 
        /// Other available API versions: 2022-04-01-preview, 2022-11-01, 2023-06-01, 2023-09-01, 2024-02-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native mobilenetwork [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetPacketCoreControlPlaneResult> Invoke(GetPacketCoreControlPlaneInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetPacketCoreControlPlaneResult>("azure-native:mobilenetwork:getPacketCoreControlPlane", args ?? new GetPacketCoreControlPlaneInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetPacketCoreControlPlaneArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the packet core control plane.
        /// </summary>
        [Input("packetCoreControlPlaneName", required: true)]
        public string PacketCoreControlPlaneName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetPacketCoreControlPlaneArgs()
        {
        }
        public static new GetPacketCoreControlPlaneArgs Empty => new GetPacketCoreControlPlaneArgs();
    }

    public sealed class GetPacketCoreControlPlaneInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the packet core control plane.
        /// </summary>
        [Input("packetCoreControlPlaneName", required: true)]
        public Input<string> PacketCoreControlPlaneName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        public GetPacketCoreControlPlaneInvokeArgs()
        {
        }
        public static new GetPacketCoreControlPlaneInvokeArgs Empty => new GetPacketCoreControlPlaneInvokeArgs();
    }


    [OutputType]
    public sealed class GetPacketCoreControlPlaneResult
    {
        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        public readonly string AzureApiVersion;
        /// <summary>
        /// The control plane interface on the access network. For 5G networks, this is the N2 interface. For 4G networks, this is the S1-MME interface.
        /// </summary>
        public readonly Outputs.InterfacePropertiesResponse ControlPlaneAccessInterface;
        /// <summary>
        /// The virtual IP address(es) for the control plane on the access network in a High Availability (HA) system. In an HA deployment the access network router should be configured to anycast traffic for this address to the control plane access interfaces on the active and standby nodes. In non-HA system this list should be omitted or empty.
        /// </summary>
        public readonly ImmutableArray<string> ControlPlaneAccessVirtualIpv4Addresses;
        /// <summary>
        /// The core network technology generation (5G core or EPC / 4G core).
        /// </summary>
        public readonly string? CoreNetworkTechnology;
        /// <summary>
        /// Configuration for uploading packet core diagnostics
        /// </summary>
        public readonly Outputs.DiagnosticsUploadConfigurationResponse? DiagnosticsUpload;
        /// <summary>
        /// Configuration for sending packet core events to an Azure Event Hub.
        /// </summary>
        public readonly Outputs.EventHubConfigurationResponse? EventHub;
        /// <summary>
        /// The provisioning state of the secret containing private keys and keyIds for SUPI concealment.
        /// </summary>
        public readonly Outputs.HomeNetworkPrivateKeysProvisioningResponse HomeNetworkPrivateKeysProvisioning;
        /// <summary>
        /// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The identity used to retrieve the ingress certificate from Azure key vault.
        /// </summary>
        public readonly Outputs.ManagedServiceIdentityResponse? Identity;
        /// <summary>
        /// The installation state of the packet core control plane resource.
        /// </summary>
        public readonly Outputs.InstallationResponse? Installation;
        /// <summary>
        /// The currently installed version of the packet core software.
        /// </summary>
        public readonly string InstalledVersion;
        /// <summary>
        /// Settings to allow interoperability with third party components e.g. RANs and UEs.
        /// </summary>
        public readonly object? InteropSettings;
        /// <summary>
        /// The kubernetes ingress configuration to control access to packet core diagnostics over local APIs.
        /// </summary>
        public readonly Outputs.LocalDiagnosticsAccessConfigurationResponse LocalDiagnosticsAccess;
        /// <summary>
        /// The geo-location where the resource lives
        /// </summary>
        public readonly string Location;
        /// <summary>
        /// The name of the resource
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The platform where the packet core is deployed.
        /// </summary>
        public readonly Outputs.PlatformConfigurationResponse Platform;
        /// <summary>
        /// The provisioning state of the packet core control plane resource.
        /// </summary>
        public readonly string ProvisioningState;
        /// <summary>
        /// The previous version of the packet core software that was deployed. Used when performing the rollback action.
        /// </summary>
        public readonly string RollbackVersion;
        /// <summary>
        /// Signaling configuration for the packet core.
        /// </summary>
        public readonly Outputs.SignalingConfigurationResponse? Signaling;
        /// <summary>
        /// Site(s) under which this packet core control plane should be deployed. The sites must be in the same location as the packet core control plane.
        /// </summary>
        public readonly ImmutableArray<Outputs.SiteResourceIdResponse> Sites;
        /// <summary>
        /// The SKU defining the throughput and SIM allowances for this packet core control plane deployment.
        /// </summary>
        public readonly string Sku;
        /// <summary>
        /// Azure Resource Manager metadata containing createdBy and modifiedBy information.
        /// </summary>
        public readonly Outputs.SystemDataResponse SystemData;
        /// <summary>
        /// Resource tags.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Tags;
        /// <summary>
        /// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// The MTU (in bytes) signaled to the UE. The same MTU is set on the user plane data links for all data networks. The MTU set on the user plane access link is calculated to be 60 bytes greater than this value to allow for GTP encapsulation.
        /// </summary>
        public readonly int? UeMtu;
        /// <summary>
        /// The user consent configuration for the packet core.
        /// </summary>
        public readonly Outputs.UserConsentConfigurationResponse? UserConsent;
        /// <summary>
        /// The desired version of the packet core software.
        /// </summary>
        public readonly string? Version;

        [OutputConstructor]
        private GetPacketCoreControlPlaneResult(
            string azureApiVersion,

            Outputs.InterfacePropertiesResponse controlPlaneAccessInterface,

            ImmutableArray<string> controlPlaneAccessVirtualIpv4Addresses,

            string? coreNetworkTechnology,

            Outputs.DiagnosticsUploadConfigurationResponse? diagnosticsUpload,

            Outputs.EventHubConfigurationResponse? eventHub,

            Outputs.HomeNetworkPrivateKeysProvisioningResponse homeNetworkPrivateKeysProvisioning,

            string id,

            Outputs.ManagedServiceIdentityResponse? identity,

            Outputs.InstallationResponse? installation,

            string installedVersion,

            object? interopSettings,

            Outputs.LocalDiagnosticsAccessConfigurationResponse localDiagnosticsAccess,

            string location,

            string name,

            Outputs.PlatformConfigurationResponse platform,

            string provisioningState,

            string rollbackVersion,

            Outputs.SignalingConfigurationResponse? signaling,

            ImmutableArray<Outputs.SiteResourceIdResponse> sites,

            string sku,

            Outputs.SystemDataResponse systemData,

            ImmutableDictionary<string, string>? tags,

            string type,

            int? ueMtu,

            Outputs.UserConsentConfigurationResponse? userConsent,

            string? version)
        {
            AzureApiVersion = azureApiVersion;
            ControlPlaneAccessInterface = controlPlaneAccessInterface;
            ControlPlaneAccessVirtualIpv4Addresses = controlPlaneAccessVirtualIpv4Addresses;
            CoreNetworkTechnology = coreNetworkTechnology;
            DiagnosticsUpload = diagnosticsUpload;
            EventHub = eventHub;
            HomeNetworkPrivateKeysProvisioning = homeNetworkPrivateKeysProvisioning;
            Id = id;
            Identity = identity;
            Installation = installation;
            InstalledVersion = installedVersion;
            InteropSettings = interopSettings;
            LocalDiagnosticsAccess = localDiagnosticsAccess;
            Location = location;
            Name = name;
            Platform = platform;
            ProvisioningState = provisioningState;
            RollbackVersion = rollbackVersion;
            Signaling = signaling;
            Sites = sites;
            Sku = sku;
            SystemData = systemData;
            Tags = tags;
            Type = type;
            UeMtu = ueMtu;
            UserConsent = userConsent;
            Version = version;
        }
    }
}
