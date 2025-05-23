// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.ElasticSan
{
    /// <summary>
    /// Response for ElasticSan request.
    /// 
    /// Uses Azure REST API version 2024-05-01. In version 2.x of the Azure Native provider, it used API version 2021-11-20-preview.
    /// 
    /// Other available API versions: 2021-11-20-preview, 2022-12-01-preview, 2023-01-01, 2024-06-01-preview, 2024-07-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native elasticsan [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
    /// </summary>
    [AzureNativeResourceType("azure-native:elasticsan:ElasticSan")]
    public partial class ElasticSan : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Logical zone for Elastic San resource; example: ["1"].
        /// </summary>
        [Output("availabilityZones")]
        public Output<ImmutableArray<string>> AvailabilityZones { get; private set; } = null!;

        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        [Output("azureApiVersion")]
        public Output<string> AzureApiVersion { get; private set; } = null!;

        /// <summary>
        /// Base size of the Elastic San appliance in TiB.
        /// </summary>
        [Output("baseSizeTiB")]
        public Output<double> BaseSizeTiB { get; private set; } = null!;

        /// <summary>
        /// Extended size of the Elastic San appliance in TiB.
        /// </summary>
        [Output("extendedCapacitySizeTiB")]
        public Output<double> ExtendedCapacitySizeTiB { get; private set; } = null!;

        /// <summary>
        /// The geo-location where the resource lives
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// The name of the resource
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The list of Private Endpoint Connections.
        /// </summary>
        [Output("privateEndpointConnections")]
        public Output<ImmutableArray<Outputs.PrivateEndpointConnectionResponse>> PrivateEndpointConnections { get; private set; } = null!;

        /// <summary>
        /// State of the operation on the resource.
        /// </summary>
        [Output("provisioningState")]
        public Output<string> ProvisioningState { get; private set; } = null!;

        /// <summary>
        /// Allow or disallow public network access to ElasticSan. Value is optional but if passed in, must be 'Enabled' or 'Disabled'.
        /// </summary>
        [Output("publicNetworkAccess")]
        public Output<string?> PublicNetworkAccess { get; private set; } = null!;

        /// <summary>
        /// resource sku
        /// </summary>
        [Output("sku")]
        public Output<Outputs.SkuResponse> Sku { get; private set; } = null!;

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
        /// Total Provisioned IOPS of the Elastic San appliance.
        /// </summary>
        [Output("totalIops")]
        public Output<double> TotalIops { get; private set; } = null!;

        /// <summary>
        /// Total Provisioned MBps Elastic San appliance.
        /// </summary>
        [Output("totalMBps")]
        public Output<double> TotalMBps { get; private set; } = null!;

        /// <summary>
        /// Total size of the Elastic San appliance in TB.
        /// </summary>
        [Output("totalSizeTiB")]
        public Output<double> TotalSizeTiB { get; private set; } = null!;

        /// <summary>
        /// Total size of the provisioned Volumes in GiB.
        /// </summary>
        [Output("totalVolumeSizeGiB")]
        public Output<double> TotalVolumeSizeGiB { get; private set; } = null!;

        /// <summary>
        /// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;

        /// <summary>
        /// Total number of volume groups in this Elastic San appliance.
        /// </summary>
        [Output("volumeGroupCount")]
        public Output<double> VolumeGroupCount { get; private set; } = null!;


        /// <summary>
        /// Create a ElasticSan resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ElasticSan(string name, ElasticSanArgs args, CustomResourceOptions? options = null)
            : base("azure-native:elasticsan:ElasticSan", name, args ?? new ElasticSanArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ElasticSan(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azure-native:elasticsan:ElasticSan", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new global::Pulumi.Alias { Type = "azure-native:elasticsan/v20211120preview:ElasticSan" },
                    new global::Pulumi.Alias { Type = "azure-native:elasticsan/v20221201preview:ElasticSan" },
                    new global::Pulumi.Alias { Type = "azure-native:elasticsan/v20230101:ElasticSan" },
                    new global::Pulumi.Alias { Type = "azure-native:elasticsan/v20240501:ElasticSan" },
                    new global::Pulumi.Alias { Type = "azure-native:elasticsan/v20240601preview:ElasticSan" },
                    new global::Pulumi.Alias { Type = "azure-native:elasticsan/v20240701preview:ElasticSan" },
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing ElasticSan resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ElasticSan Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new ElasticSan(name, id, options);
        }
    }

    public sealed class ElasticSanArgs : global::Pulumi.ResourceArgs
    {
        [Input("availabilityZones")]
        private InputList<string>? _availabilityZones;

        /// <summary>
        /// Logical zone for Elastic San resource; example: ["1"].
        /// </summary>
        public InputList<string> AvailabilityZones
        {
            get => _availabilityZones ?? (_availabilityZones = new InputList<string>());
            set => _availabilityZones = value;
        }

        /// <summary>
        /// Base size of the Elastic San appliance in TiB.
        /// </summary>
        [Input("baseSizeTiB", required: true)]
        public Input<double> BaseSizeTiB { get; set; } = null!;

        /// <summary>
        /// The name of the ElasticSan.
        /// </summary>
        [Input("elasticSanName")]
        public Input<string>? ElasticSanName { get; set; }

        /// <summary>
        /// Extended size of the Elastic San appliance in TiB.
        /// </summary>
        [Input("extendedCapacitySizeTiB", required: true)]
        public Input<double> ExtendedCapacitySizeTiB { get; set; } = null!;

        /// <summary>
        /// The geo-location where the resource lives
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// Allow or disallow public network access to ElasticSan. Value is optional but if passed in, must be 'Enabled' or 'Disabled'.
        /// </summary>
        [Input("publicNetworkAccess")]
        public InputUnion<string, Pulumi.AzureNative.ElasticSan.PublicNetworkAccess>? PublicNetworkAccess { get; set; }

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// resource sku
        /// </summary>
        [Input("sku", required: true)]
        public Input<Inputs.SkuArgs> Sku { get; set; } = null!;

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

        public ElasticSanArgs()
        {
        }
        public static new ElasticSanArgs Empty => new ElasticSanArgs();
    }
}
