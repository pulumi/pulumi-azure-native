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
    /// The service resource.
    /// 
    /// Uses Azure REST API version 2024-11-01-preview. In version 2.x of the Azure Native provider, it used API version 2024-11-01-preview.
    /// 
    /// Other available API versions: 2025-03-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native servicefabric [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
    /// </summary>
    [AzureNativeResourceType("azure-native:servicefabric:Service")]
    public partial class Service : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        [Output("azureApiVersion")]
        public Output<string> AzureApiVersion { get; private set; } = null!;

        /// <summary>
        /// The geo-location where the resource lives
        /// </summary>
        [Output("location")]
        public Output<string?> Location { get; private set; } = null!;

        /// <summary>
        /// The name of the resource
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The service resource properties.
        /// </summary>
        [Output("properties")]
        public Output<Union<Outputs.StatefulServicePropertiesResponse, Outputs.StatelessServicePropertiesResponse>> Properties { get; private set; } = null!;

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
        /// Create a Service resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Service(string name, ServiceArgs args, CustomResourceOptions? options = null)
            : base("azure-native:servicefabric:Service", name, args ?? new ServiceArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Service(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azure-native:servicefabric:Service", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new global::Pulumi.Alias { Type = "azure-native:servicefabric/v20210101preview:Service" },
                    new global::Pulumi.Alias { Type = "azure-native:servicefabric/v20210501:Service" },
                    new global::Pulumi.Alias { Type = "azure-native:servicefabric/v20210601:Service" },
                    new global::Pulumi.Alias { Type = "azure-native:servicefabric/v20210701preview:Service" },
                    new global::Pulumi.Alias { Type = "azure-native:servicefabric/v20210901privatepreview:Service" },
                    new global::Pulumi.Alias { Type = "azure-native:servicefabric/v20211101preview:Service" },
                    new global::Pulumi.Alias { Type = "azure-native:servicefabric/v20220101:Service" },
                    new global::Pulumi.Alias { Type = "azure-native:servicefabric/v20220201preview:Service" },
                    new global::Pulumi.Alias { Type = "azure-native:servicefabric/v20220601preview:Service" },
                    new global::Pulumi.Alias { Type = "azure-native:servicefabric/v20220801preview:Service" },
                    new global::Pulumi.Alias { Type = "azure-native:servicefabric/v20221001preview:Service" },
                    new global::Pulumi.Alias { Type = "azure-native:servicefabric/v20230201preview:Service" },
                    new global::Pulumi.Alias { Type = "azure-native:servicefabric/v20230301preview:ManagedClusterService" },
                    new global::Pulumi.Alias { Type = "azure-native:servicefabric/v20230301preview:Service" },
                    new global::Pulumi.Alias { Type = "azure-native:servicefabric/v20230701preview:ManagedClusterService" },
                    new global::Pulumi.Alias { Type = "azure-native:servicefabric/v20230701preview:Service" },
                    new global::Pulumi.Alias { Type = "azure-native:servicefabric/v20230901preview:ManagedClusterService" },
                    new global::Pulumi.Alias { Type = "azure-native:servicefabric/v20230901preview:Service" },
                    new global::Pulumi.Alias { Type = "azure-native:servicefabric/v20231101preview:ManagedClusterService" },
                    new global::Pulumi.Alias { Type = "azure-native:servicefabric/v20231101preview:Service" },
                    new global::Pulumi.Alias { Type = "azure-native:servicefabric/v20231201preview:ManagedClusterService" },
                    new global::Pulumi.Alias { Type = "azure-native:servicefabric/v20231201preview:Service" },
                    new global::Pulumi.Alias { Type = "azure-native:servicefabric/v20240201preview:ManagedClusterService" },
                    new global::Pulumi.Alias { Type = "azure-native:servicefabric/v20240201preview:Service" },
                    new global::Pulumi.Alias { Type = "azure-native:servicefabric/v20240401:ManagedClusterService" },
                    new global::Pulumi.Alias { Type = "azure-native:servicefabric/v20240401:Service" },
                    new global::Pulumi.Alias { Type = "azure-native:servicefabric/v20240601preview:ManagedClusterService" },
                    new global::Pulumi.Alias { Type = "azure-native:servicefabric/v20240601preview:Service" },
                    new global::Pulumi.Alias { Type = "azure-native:servicefabric/v20240901preview:ManagedClusterService" },
                    new global::Pulumi.Alias { Type = "azure-native:servicefabric/v20240901preview:Service" },
                    new global::Pulumi.Alias { Type = "azure-native:servicefabric/v20241101preview:Service" },
                    new global::Pulumi.Alias { Type = "azure-native:servicefabric/v20250301preview:Service" },
                    new global::Pulumi.Alias { Type = "azure-native:servicefabric:ManagedClusterService" },
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Service resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Service Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Service(name, id, options);
        }
    }

    public sealed class ServiceArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the application resource.
        /// </summary>
        [Input("applicationName", required: true)]
        public Input<string> ApplicationName { get; set; } = null!;

        /// <summary>
        /// The name of the cluster resource.
        /// </summary>
        [Input("clusterName", required: true)]
        public Input<string> ClusterName { get; set; } = null!;

        /// <summary>
        /// The geo-location where the resource lives
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// The service resource properties.
        /// </summary>
        [Input("properties")]
        public InputUnion<Inputs.StatefulServicePropertiesArgs, Inputs.StatelessServicePropertiesArgs>? Properties { get; set; }

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the service resource in the format of {applicationName}~{serviceName}.
        /// </summary>
        [Input("serviceName")]
        public Input<string>? ServiceName { get; set; }

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

        public ServiceArgs()
        {
        }
        public static new ServiceArgs Empty => new ServiceArgs();
    }
}
