// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.ServiceFabric
{
    /// <summary>
    /// An application type version resource for the specified application type name resource.
    /// 
    /// Uses Azure REST API version 2024-04-01. In version 2.x of the Azure Native provider, it used API version 2023-03-01-preview.
    /// 
    /// Other available API versions: 2023-03-01-preview, 2023-07-01-preview, 2023-09-01-preview, 2023-11-01-preview, 2023-12-01-preview, 2024-02-01-preview, 2024-06-01-preview, 2024-09-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native servicefabric [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
    /// </summary>
    [AzureNativeResourceType("azure-native:servicefabric:ManagedClusterApplicationTypeVersion")]
    public partial class ManagedClusterApplicationTypeVersion : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The URL to the application package
        /// </summary>
        [Output("appPackageUrl")]
        public Output<string> AppPackageUrl { get; private set; } = null!;

        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        [Output("azureApiVersion")]
        public Output<string> AzureApiVersion { get; private set; } = null!;

        /// <summary>
        /// Resource location depends on the parent resource.
        /// </summary>
        [Output("location")]
        public Output<string?> Location { get; private set; } = null!;

        /// <summary>
        /// Azure resource name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The current deployment or provisioning state, which only appears in the response
        /// </summary>
        [Output("provisioningState")]
        public Output<string> ProvisioningState { get; private set; } = null!;

        /// <summary>
        /// Metadata pertaining to creation and last modification of the resource.
        /// </summary>
        [Output("systemData")]
        public Output<Outputs.SystemDataResponse> SystemData { get; private set; } = null!;

        /// <summary>
        /// Azure resource tags.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// Azure resource type.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a ManagedClusterApplicationTypeVersion resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ManagedClusterApplicationTypeVersion(string name, ManagedClusterApplicationTypeVersionArgs args, CustomResourceOptions? options = null)
            : base("azure-native:servicefabric:ManagedClusterApplicationTypeVersion", name, args ?? new ManagedClusterApplicationTypeVersionArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ManagedClusterApplicationTypeVersion(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azure-native:servicefabric:ManagedClusterApplicationTypeVersion", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new global::Pulumi.Alias { Type = "azure-native:servicefabric/v20210101preview:ManagedClusterApplicationTypeVersion" },
                    new global::Pulumi.Alias { Type = "azure-native:servicefabric/v20210501:ManagedClusterApplicationTypeVersion" },
                    new global::Pulumi.Alias { Type = "azure-native:servicefabric/v20210701preview:ManagedClusterApplicationTypeVersion" },
                    new global::Pulumi.Alias { Type = "azure-native:servicefabric/v20210901privatepreview:ManagedClusterApplicationTypeVersion" },
                    new global::Pulumi.Alias { Type = "azure-native:servicefabric/v20211101preview:ManagedClusterApplicationTypeVersion" },
                    new global::Pulumi.Alias { Type = "azure-native:servicefabric/v20220101:ManagedClusterApplicationTypeVersion" },
                    new global::Pulumi.Alias { Type = "azure-native:servicefabric/v20220201preview:ManagedClusterApplicationTypeVersion" },
                    new global::Pulumi.Alias { Type = "azure-native:servicefabric/v20220601preview:ManagedClusterApplicationTypeVersion" },
                    new global::Pulumi.Alias { Type = "azure-native:servicefabric/v20220801preview:ManagedClusterApplicationTypeVersion" },
                    new global::Pulumi.Alias { Type = "azure-native:servicefabric/v20221001preview:ManagedClusterApplicationTypeVersion" },
                    new global::Pulumi.Alias { Type = "azure-native:servicefabric/v20230201preview:ManagedClusterApplicationTypeVersion" },
                    new global::Pulumi.Alias { Type = "azure-native:servicefabric/v20230301preview:ManagedClusterApplicationTypeVersion" },
                    new global::Pulumi.Alias { Type = "azure-native:servicefabric/v20230701preview:ManagedClusterApplicationTypeVersion" },
                    new global::Pulumi.Alias { Type = "azure-native:servicefabric/v20230901preview:ManagedClusterApplicationTypeVersion" },
                    new global::Pulumi.Alias { Type = "azure-native:servicefabric/v20231101preview:ManagedClusterApplicationTypeVersion" },
                    new global::Pulumi.Alias { Type = "azure-native:servicefabric/v20231201preview:ManagedClusterApplicationTypeVersion" },
                    new global::Pulumi.Alias { Type = "azure-native:servicefabric/v20240201preview:ManagedClusterApplicationTypeVersion" },
                    new global::Pulumi.Alias { Type = "azure-native:servicefabric/v20240401:ManagedClusterApplicationTypeVersion" },
                    new global::Pulumi.Alias { Type = "azure-native:servicefabric/v20240601preview:ManagedClusterApplicationTypeVersion" },
                    new global::Pulumi.Alias { Type = "azure-native:servicefabric/v20240901preview:ManagedClusterApplicationTypeVersion" },
                    new global::Pulumi.Alias { Type = "azure-native:servicefabric/v20241101preview:ManagedClusterApplicationTypeVersion" },
                    new global::Pulumi.Alias { Type = "azure-native:servicefabric/v20250301preview:ManagedClusterApplicationTypeVersion" },
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing ManagedClusterApplicationTypeVersion resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ManagedClusterApplicationTypeVersion Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new ManagedClusterApplicationTypeVersion(name, id, options);
        }
    }

    public sealed class ManagedClusterApplicationTypeVersionArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The URL to the application package
        /// </summary>
        [Input("appPackageUrl", required: true)]
        public Input<string> AppPackageUrl { get; set; } = null!;

        /// <summary>
        /// The name of the application type name resource.
        /// </summary>
        [Input("applicationTypeName", required: true)]
        public Input<string> ApplicationTypeName { get; set; } = null!;

        /// <summary>
        /// The name of the cluster resource.
        /// </summary>
        [Input("clusterName", required: true)]
        public Input<string> ClusterName { get; set; } = null!;

        /// <summary>
        /// Resource location depends on the parent resource.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// The name of the resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// Azure resource tags.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        /// <summary>
        /// The application type version.
        /// </summary>
        [Input("version")]
        public Input<string>? Version { get; set; }

        public ManagedClusterApplicationTypeVersionArgs()
        {
        }
        public static new ManagedClusterApplicationTypeVersionArgs Empty => new ManagedClusterApplicationTypeVersionArgs();
    }
}
