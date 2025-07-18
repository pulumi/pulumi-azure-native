// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.DeviceRegistry
{
    public static class GetAsset
    {
        /// <summary>
        /// Get a Asset
        /// 
        /// Uses Azure REST API version 2024-11-01.
        /// 
        /// Other available API versions: 2023-11-01-preview, 2024-09-01-preview, 2025-07-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native deviceregistry [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Task<GetAssetResult> InvokeAsync(GetAssetArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetAssetResult>("azure-native:deviceregistry:getAsset", args ?? new GetAssetArgs(), options.WithDefaults());

        /// <summary>
        /// Get a Asset
        /// 
        /// Uses Azure REST API version 2024-11-01.
        /// 
        /// Other available API versions: 2023-11-01-preview, 2024-09-01-preview, 2025-07-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native deviceregistry [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetAssetResult> Invoke(GetAssetInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetAssetResult>("azure-native:deviceregistry:getAsset", args ?? new GetAssetInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Get a Asset
        /// 
        /// Uses Azure REST API version 2024-11-01.
        /// 
        /// Other available API versions: 2023-11-01-preview, 2024-09-01-preview, 2025-07-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native deviceregistry [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetAssetResult> Invoke(GetAssetInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetAssetResult>("azure-native:deviceregistry:getAsset", args ?? new GetAssetInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetAssetArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Asset name parameter.
        /// </summary>
        [Input("assetName", required: true)]
        public string AssetName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetAssetArgs()
        {
        }
        public static new GetAssetArgs Empty => new GetAssetArgs();
    }

    public sealed class GetAssetInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Asset name parameter.
        /// </summary>
        [Input("assetName", required: true)]
        public Input<string> AssetName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        public GetAssetInvokeArgs()
        {
        }
        public static new GetAssetInvokeArgs Empty => new GetAssetInvokeArgs();
    }


    [OutputType]
    public sealed class GetAssetResult
    {
        /// <summary>
        /// A reference to the asset endpoint profile (connection information) used by brokers to connect to an endpoint that provides data points for this asset. Must provide asset endpoint profile name.
        /// </summary>
        public readonly string AssetEndpointProfileRef;
        /// <summary>
        /// A set of key-value pairs that contain custom attributes set by the customer.
        /// </summary>
        public readonly object? Attributes;
        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        public readonly string AzureApiVersion;
        /// <summary>
        /// Array of datasets that are part of the asset. Each dataset describes the data points that make up the set.
        /// </summary>
        public readonly ImmutableArray<Outputs.DatasetResponse> Datasets;
        /// <summary>
        /// Stringified JSON that contains connector-specific default configuration for all datasets. Each dataset can have its own configuration that overrides the default settings here.
        /// </summary>
        public readonly string? DefaultDatasetsConfiguration;
        /// <summary>
        /// Stringified JSON that contains connector-specific default configuration for all events. Each event can have its own configuration that overrides the default settings here.
        /// </summary>
        public readonly string? DefaultEventsConfiguration;
        /// <summary>
        /// Object that describes the default topic information for the asset.
        /// </summary>
        public readonly Outputs.TopicResponse? DefaultTopic;
        /// <summary>
        /// Human-readable description of the asset.
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// Reference to a list of discovered assets. Populated only if the asset has been created from discovery flow. Discovered asset names must be provided.
        /// </summary>
        public readonly ImmutableArray<string> DiscoveredAssetRefs;
        /// <summary>
        /// Human-readable display name.
        /// </summary>
        public readonly string? DisplayName;
        /// <summary>
        /// Reference to the documentation.
        /// </summary>
        public readonly string? DocumentationUri;
        /// <summary>
        /// Enabled/Disabled status of the asset.
        /// </summary>
        public readonly bool? Enabled;
        /// <summary>
        /// Array of events that are part of the asset. Each event can have per-event configuration.
        /// </summary>
        public readonly ImmutableArray<Outputs.EventResponse> Events;
        /// <summary>
        /// The extended location.
        /// </summary>
        public readonly Outputs.ExtendedLocationResponse ExtendedLocation;
        /// <summary>
        /// Asset id provided by the customer.
        /// </summary>
        public readonly string? ExternalAssetId;
        /// <summary>
        /// Revision number of the hardware.
        /// </summary>
        public readonly string? HardwareRevision;
        /// <summary>
        /// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The geo-location where the resource lives
        /// </summary>
        public readonly string Location;
        /// <summary>
        /// Asset manufacturer name.
        /// </summary>
        public readonly string? Manufacturer;
        /// <summary>
        /// Asset manufacturer URI.
        /// </summary>
        public readonly string? ManufacturerUri;
        /// <summary>
        /// Asset model name.
        /// </summary>
        public readonly string? Model;
        /// <summary>
        /// The name of the resource
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Asset product code.
        /// </summary>
        public readonly string? ProductCode;
        /// <summary>
        /// Provisioning state of the resource.
        /// </summary>
        public readonly string ProvisioningState;
        /// <summary>
        /// Asset serial number.
        /// </summary>
        public readonly string? SerialNumber;
        /// <summary>
        /// Revision number of the software.
        /// </summary>
        public readonly string? SoftwareRevision;
        /// <summary>
        /// Read only object to reflect changes that have occurred on the Edge. Similar to Kubernetes status property for custom resources.
        /// </summary>
        public readonly Outputs.AssetStatusResponse Status;
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
        /// Globally unique, immutable, non-reusable id.
        /// </summary>
        public readonly string Uuid;
        /// <summary>
        /// An integer that is incremented each time the resource is modified.
        /// </summary>
        public readonly double Version;

        [OutputConstructor]
        private GetAssetResult(
            string assetEndpointProfileRef,

            object? attributes,

            string azureApiVersion,

            ImmutableArray<Outputs.DatasetResponse> datasets,

            string? defaultDatasetsConfiguration,

            string? defaultEventsConfiguration,

            Outputs.TopicResponse? defaultTopic,

            string? description,

            ImmutableArray<string> discoveredAssetRefs,

            string? displayName,

            string? documentationUri,

            bool? enabled,

            ImmutableArray<Outputs.EventResponse> events,

            Outputs.ExtendedLocationResponse extendedLocation,

            string? externalAssetId,

            string? hardwareRevision,

            string id,

            string location,

            string? manufacturer,

            string? manufacturerUri,

            string? model,

            string name,

            string? productCode,

            string provisioningState,

            string? serialNumber,

            string? softwareRevision,

            Outputs.AssetStatusResponse status,

            Outputs.SystemDataResponse systemData,

            ImmutableDictionary<string, string>? tags,

            string type,

            string uuid,

            double version)
        {
            AssetEndpointProfileRef = assetEndpointProfileRef;
            Attributes = attributes;
            AzureApiVersion = azureApiVersion;
            Datasets = datasets;
            DefaultDatasetsConfiguration = defaultDatasetsConfiguration;
            DefaultEventsConfiguration = defaultEventsConfiguration;
            DefaultTopic = defaultTopic;
            Description = description;
            DiscoveredAssetRefs = discoveredAssetRefs;
            DisplayName = displayName;
            DocumentationUri = documentationUri;
            Enabled = enabled;
            Events = events;
            ExtendedLocation = extendedLocation;
            ExternalAssetId = externalAssetId;
            HardwareRevision = hardwareRevision;
            Id = id;
            Location = location;
            Manufacturer = manufacturer;
            ManufacturerUri = manufacturerUri;
            Model = model;
            Name = name;
            ProductCode = productCode;
            ProvisioningState = provisioningState;
            SerialNumber = serialNumber;
            SoftwareRevision = softwareRevision;
            Status = status;
            SystemData = systemData;
            Tags = tags;
            Type = type;
            Uuid = uuid;
            Version = version;
        }
    }
}
