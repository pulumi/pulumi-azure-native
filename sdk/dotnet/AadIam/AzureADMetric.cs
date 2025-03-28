// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.AadIam
{
    /// <summary>
    /// AzureADMetrics resource.
    /// 
    /// Uses Azure REST API version 2020-07-01-preview. In version 1.x of the Azure Native provider, it used API version 2020-07-01-preview.
    /// </summary>
    [AzureNativeResourceType("azure-native:aadiam:AzureADMetric")]
    public partial class AzureADMetric : global::Pulumi.CustomResource
    {
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

        [Output("properties")]
        public Output<Outputs.AzureADMetricsPropertiesFormatResponse> Properties { get; private set; } = null!;

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
        /// Create a AzureADMetric resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AzureADMetric(string name, AzureADMetricArgs args, CustomResourceOptions? options = null)
            : base("azure-native:aadiam:AzureADMetric", name, args ?? new AzureADMetricArgs(), MakeResourceOptions(options, ""))
        {
        }

        private AzureADMetric(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azure-native:aadiam:AzureADMetric", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new global::Pulumi.Alias { Type = "azure-native:aadiam/v20200701preview:AzureADMetric" },
                    new global::Pulumi.Alias { Type = "azure-native:aadiam/v20200701preview:azureADMetric" },
                    new global::Pulumi.Alias { Type = "azure-native:aadiam:azureADMetric" },
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing AzureADMetric resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AzureADMetric Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new AzureADMetric(name, id, options);
        }
    }

    public sealed class AzureADMetricArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Name of the azureADMetrics instance.
        /// </summary>
        [Input("azureADMetricsName")]
        public Input<string>? AzureADMetricsName { get; set; }

        /// <summary>
        /// The geo-location where the resource lives
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// Name of an Azure resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

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

        public AzureADMetricArgs()
        {
        }
        public static new AzureADMetricArgs Empty => new AzureADMetricArgs();
    }
}
