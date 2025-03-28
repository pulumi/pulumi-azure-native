// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.DeviceRegistry
{
    /// <summary>
    /// Asset definition.
    /// 
    /// Uses Azure REST API version 2023-11-01-preview.
    /// 
    /// Other available API versions: 2024-09-01-preview, 2024-11-01.
    /// </summary>
    [AzureNativeResourceType("azure-native:deviceregistry:Asset")]
    public partial class Asset : global::Pulumi.CustomResource
    {
        /// <summary>
        /// A reference to the asset endpoint profile (connection information) used by brokers to connect to an endpoint that provides data points for this asset. Must have the format &lt;ModuleCR.metadata.namespace&gt;/&lt;ModuleCR.metadata.name&gt;.
        /// </summary>
        [Output("assetEndpointProfileUri")]
        public Output<string> AssetEndpointProfileUri { get; private set; } = null!;

        /// <summary>
        /// Resource path to asset type (model) definition.
        /// </summary>
        [Output("assetType")]
        public Output<string?> AssetType { get; private set; } = null!;

        /// <summary>
        /// A set of key-value pairs that contain custom attributes set by the customer.
        /// </summary>
        [Output("attributes")]
        public Output<object?> Attributes { get; private set; } = null!;

        /// <summary>
        /// Array of data points that are part of the asset. Each data point can reference an asset type capability and have per-data point configuration.
        /// </summary>
        [Output("dataPoints")]
        public Output<ImmutableArray<Outputs.DataPointResponse>> DataPoints { get; private set; } = null!;

        /// <summary>
        /// Stringified JSON that contains protocol-specific default configuration for all data points. Each data point can have its own configuration that overrides the default settings here.
        /// </summary>
        [Output("defaultDataPointsConfiguration")]
        public Output<string?> DefaultDataPointsConfiguration { get; private set; } = null!;

        /// <summary>
        /// Stringified JSON that contains connector-specific default configuration for all events. Each event can have its own configuration that overrides the default settings here.
        /// </summary>
        [Output("defaultEventsConfiguration")]
        public Output<string?> DefaultEventsConfiguration { get; private set; } = null!;

        /// <summary>
        /// Human-readable description of the asset.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// Human-readable display name.
        /// </summary>
        [Output("displayName")]
        public Output<string?> DisplayName { get; private set; } = null!;

        /// <summary>
        /// Reference to the documentation.
        /// </summary>
        [Output("documentationUri")]
        public Output<string?> DocumentationUri { get; private set; } = null!;

        /// <summary>
        /// Enabled/Disabled status of the asset.
        /// </summary>
        [Output("enabled")]
        public Output<bool?> Enabled { get; private set; } = null!;

        /// <summary>
        /// Array of events that are part of the asset. Each event can have per-event configuration.
        /// </summary>
        [Output("events")]
        public Output<ImmutableArray<Outputs.EventResponse>> Events { get; private set; } = null!;

        /// <summary>
        /// The extended location.
        /// </summary>
        [Output("extendedLocation")]
        public Output<Outputs.ExtendedLocationResponse> ExtendedLocation { get; private set; } = null!;

        /// <summary>
        /// Asset id provided by the customer.
        /// </summary>
        [Output("externalAssetId")]
        public Output<string?> ExternalAssetId { get; private set; } = null!;

        /// <summary>
        /// Revision number of the hardware.
        /// </summary>
        [Output("hardwareRevision")]
        public Output<string?> HardwareRevision { get; private set; } = null!;

        /// <summary>
        /// The geo-location where the resource lives
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// Asset manufacturer name.
        /// </summary>
        [Output("manufacturer")]
        public Output<string?> Manufacturer { get; private set; } = null!;

        /// <summary>
        /// Asset manufacturer URI.
        /// </summary>
        [Output("manufacturerUri")]
        public Output<string?> ManufacturerUri { get; private set; } = null!;

        /// <summary>
        /// Asset model name.
        /// </summary>
        [Output("model")]
        public Output<string?> Model { get; private set; } = null!;

        /// <summary>
        /// The name of the resource
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Asset product code.
        /// </summary>
        [Output("productCode")]
        public Output<string?> ProductCode { get; private set; } = null!;

        /// <summary>
        /// Provisioning state of the resource.
        /// </summary>
        [Output("provisioningState")]
        public Output<string> ProvisioningState { get; private set; } = null!;

        /// <summary>
        /// Asset serial number.
        /// </summary>
        [Output("serialNumber")]
        public Output<string?> SerialNumber { get; private set; } = null!;

        /// <summary>
        /// Revision number of the software.
        /// </summary>
        [Output("softwareRevision")]
        public Output<string?> SoftwareRevision { get; private set; } = null!;

        /// <summary>
        /// Read only object to reflect changes that have occurred on the Edge. Similar to Kubernetes status property for custom resources.
        /// </summary>
        [Output("status")]
        public Output<Outputs.AssetStatusResponse> Status { get; private set; } = null!;

        /// <summary>
        /// Azure Resource Manager metadata containing createdBy and modifiedBy information.
        /// </summary>
        [Output("systemData")]
        public Output<Outputs.SystemDataResponse> SystemData { get; private set; } = null!;

        /// <summary>
        /// Resource tags.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;

        /// <summary>
        /// Globally unique, immutable, non-reusable id.
        /// </summary>
        [Output("uuid")]
        public Output<string> Uuid { get; private set; } = null!;

        /// <summary>
        /// An integer that is incremented each time the resource is modified.
        /// </summary>
        [Output("version")]
        public Output<int> Version { get; private set; } = null!;


        /// <summary>
        /// Create a Asset resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Asset(string name, AssetArgs args, CustomResourceOptions? options = null)
            : base("azure-native:deviceregistry:Asset", name, args ?? new AssetArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Asset(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azure-native:deviceregistry:Asset", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new global::Pulumi.Alias { Type = "azure-native:deviceregistry/v20231101preview:Asset" },
                    new global::Pulumi.Alias { Type = "azure-native:deviceregistry/v20240901preview:Asset" },
                    new global::Pulumi.Alias { Type = "azure-native:deviceregistry/v20241101:Asset" },
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Asset resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Asset Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Asset(name, id, options);
        }
    }

    public sealed class AssetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// A reference to the asset endpoint profile (connection information) used by brokers to connect to an endpoint that provides data points for this asset. Must have the format &lt;ModuleCR.metadata.namespace&gt;/&lt;ModuleCR.metadata.name&gt;.
        /// </summary>
        [Input("assetEndpointProfileUri", required: true)]
        public Input<string> AssetEndpointProfileUri { get; set; } = null!;

        /// <summary>
        /// Asset name parameter.
        /// </summary>
        [Input("assetName")]
        public Input<string>? AssetName { get; set; }

        /// <summary>
        /// Resource path to asset type (model) definition.
        /// </summary>
        [Input("assetType")]
        public Input<string>? AssetType { get; set; }

        /// <summary>
        /// A set of key-value pairs that contain custom attributes set by the customer.
        /// </summary>
        [Input("attributes")]
        public Input<object>? Attributes { get; set; }

        [Input("dataPoints")]
        private InputList<Inputs.DataPointArgs>? _dataPoints;

        /// <summary>
        /// Array of data points that are part of the asset. Each data point can reference an asset type capability and have per-data point configuration.
        /// </summary>
        public InputList<Inputs.DataPointArgs> DataPoints
        {
            get => _dataPoints ?? (_dataPoints = new InputList<Inputs.DataPointArgs>());
            set => _dataPoints = value;
        }

        /// <summary>
        /// Stringified JSON that contains protocol-specific default configuration for all data points. Each data point can have its own configuration that overrides the default settings here.
        /// </summary>
        [Input("defaultDataPointsConfiguration")]
        public Input<string>? DefaultDataPointsConfiguration { get; set; }

        /// <summary>
        /// Stringified JSON that contains connector-specific default configuration for all events. Each event can have its own configuration that overrides the default settings here.
        /// </summary>
        [Input("defaultEventsConfiguration")]
        public Input<string>? DefaultEventsConfiguration { get; set; }

        /// <summary>
        /// Human-readable description of the asset.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Human-readable display name.
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        /// <summary>
        /// Reference to the documentation.
        /// </summary>
        [Input("documentationUri")]
        public Input<string>? DocumentationUri { get; set; }

        /// <summary>
        /// Enabled/Disabled status of the asset.
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        [Input("events")]
        private InputList<Inputs.EventArgs>? _events;

        /// <summary>
        /// Array of events that are part of the asset. Each event can have per-event configuration.
        /// </summary>
        public InputList<Inputs.EventArgs> Events
        {
            get => _events ?? (_events = new InputList<Inputs.EventArgs>());
            set => _events = value;
        }

        /// <summary>
        /// The extended location.
        /// </summary>
        [Input("extendedLocation", required: true)]
        public Input<Inputs.ExtendedLocationArgs> ExtendedLocation { get; set; } = null!;

        /// <summary>
        /// Asset id provided by the customer.
        /// </summary>
        [Input("externalAssetId")]
        public Input<string>? ExternalAssetId { get; set; }

        /// <summary>
        /// Revision number of the hardware.
        /// </summary>
        [Input("hardwareRevision")]
        public Input<string>? HardwareRevision { get; set; }

        /// <summary>
        /// The geo-location where the resource lives
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// Asset manufacturer name.
        /// </summary>
        [Input("manufacturer")]
        public Input<string>? Manufacturer { get; set; }

        /// <summary>
        /// Asset manufacturer URI.
        /// </summary>
        [Input("manufacturerUri")]
        public Input<string>? ManufacturerUri { get; set; }

        /// <summary>
        /// Asset model name.
        /// </summary>
        [Input("model")]
        public Input<string>? Model { get; set; }

        /// <summary>
        /// Asset product code.
        /// </summary>
        [Input("productCode")]
        public Input<string>? ProductCode { get; set; }

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// Asset serial number.
        /// </summary>
        [Input("serialNumber")]
        public Input<string>? SerialNumber { get; set; }

        /// <summary>
        /// Revision number of the software.
        /// </summary>
        [Input("softwareRevision")]
        public Input<string>? SoftwareRevision { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// Resource tags.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        public AssetArgs()
        {
        }
        public static new AssetArgs Empty => new AssetArgs();
    }
}
