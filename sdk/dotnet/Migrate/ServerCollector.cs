// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Migrate
{
    /// <summary>
    /// Uses Azure REST API version 2019-10-01. In version 2.x of the Azure Native provider, it used API version 2019-10-01.
    /// </summary>
    [AzureNativeResourceType("azure-native:migrate:ServerCollector")]
    public partial class ServerCollector : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        [Output("azureApiVersion")]
        public Output<string> AzureApiVersion { get; private set; } = null!;

        [Output("eTag")]
        public Output<string?> ETag { get; private set; } = null!;

        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        [Output("properties")]
        public Output<Outputs.CollectorPropertiesResponse> Properties { get; private set; } = null!;

        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a ServerCollector resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ServerCollector(string name, ServerCollectorArgs args, CustomResourceOptions? options = null)
            : base("azure-native:migrate:ServerCollector", name, args ?? new ServerCollectorArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ServerCollector(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azure-native:migrate:ServerCollector", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new global::Pulumi.Alias { Type = "azure-native:migrate/v20191001:ServerCollector" },
                    new global::Pulumi.Alias { Type = "azure-native:migrate/v20230315:ServerCollector" },
                    new global::Pulumi.Alias { Type = "azure-native:migrate/v20230315:ServerCollectorsOperation" },
                    new global::Pulumi.Alias { Type = "azure-native:migrate/v20230401preview:ServerCollector" },
                    new global::Pulumi.Alias { Type = "azure-native:migrate/v20230401preview:ServerCollectorsOperation" },
                    new global::Pulumi.Alias { Type = "azure-native:migrate/v20230501preview:ServerCollector" },
                    new global::Pulumi.Alias { Type = "azure-native:migrate/v20230501preview:ServerCollectorsOperation" },
                    new global::Pulumi.Alias { Type = "azure-native:migrate/v20230909preview:ServerCollector" },
                    new global::Pulumi.Alias { Type = "azure-native:migrate/v20230909preview:ServerCollectorsOperation" },
                    new global::Pulumi.Alias { Type = "azure-native:migrate/v20240101preview:ServerCollector" },
                    new global::Pulumi.Alias { Type = "azure-native:migrate/v20240101preview:ServerCollectorsOperation" },
                    new global::Pulumi.Alias { Type = "azure-native:migrate/v20240115:ServerCollector" },
                    new global::Pulumi.Alias { Type = "azure-native:migrate/v20240303preview:ServerCollector" },
                    new global::Pulumi.Alias { Type = "azure-native:migrate:ServerCollectorsOperation" },
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing ServerCollector resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ServerCollector Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new ServerCollector(name, id, options);
        }
    }

    public sealed class ServerCollectorArgs : global::Pulumi.ResourceArgs
    {
        [Input("eTag")]
        public Input<string>? ETag { get; set; }

        /// <summary>
        /// Name of the Azure Migrate project.
        /// </summary>
        [Input("projectName", required: true)]
        public Input<string> ProjectName { get; set; } = null!;

        [Input("properties")]
        public Input<Inputs.CollectorPropertiesArgs>? Properties { get; set; }

        /// <summary>
        /// Name of the Azure Resource Group that project is part of.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// Unique name of a Server collector within a project.
        /// </summary>
        [Input("serverCollectorName")]
        public Input<string>? ServerCollectorName { get; set; }

        public ServerCollectorArgs()
        {
        }
        public static new ServerCollectorArgs Empty => new ServerCollectorArgs();
    }
}
