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
    /// A private endpoint connection for a project.
    /// 
    /// Uses Azure REST API version 2019-10-01. In version 2.x of the Azure Native provider, it used API version 2019-10-01.
    /// </summary>
    [AzureNativeResourceType("azure-native:migrate:PrivateEndpointConnection")]
    public partial class PrivateEndpointConnection : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        [Output("azureApiVersion")]
        public Output<string> AzureApiVersion { get; private set; } = null!;

        /// <summary>
        /// For optimistic concurrency control.
        /// </summary>
        [Output("eTag")]
        public Output<string?> ETag { get; private set; } = null!;

        /// <summary>
        /// Name of the private endpoint endpoint connection.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Properties of the private endpoint endpoint connection.
        /// </summary>
        [Output("properties")]
        public Output<Outputs.PrivateEndpointConnectionPropertiesResponse> Properties { get; private set; } = null!;

        /// <summary>
        /// Type of the object = [Microsoft.Migrate/assessmentProjects/privateEndpointConnections].
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a PrivateEndpointConnection resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public PrivateEndpointConnection(string name, PrivateEndpointConnectionArgs args, CustomResourceOptions? options = null)
            : base("azure-native:migrate:PrivateEndpointConnection", name, args ?? new PrivateEndpointConnectionArgs(), MakeResourceOptions(options, ""))
        {
        }

        private PrivateEndpointConnection(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azure-native:migrate:PrivateEndpointConnection", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new global::Pulumi.Alias { Type = "azure-native:migrate/v20191001:PrivateEndpointConnection" },
                    new global::Pulumi.Alias { Type = "azure-native:migrate/v20230315:PrivateEndpointConnection" },
                    new global::Pulumi.Alias { Type = "azure-native:migrate/v20230315:PrivateEndpointConnectionOperation" },
                    new global::Pulumi.Alias { Type = "azure-native:migrate/v20230401preview:PrivateEndpointConnection" },
                    new global::Pulumi.Alias { Type = "azure-native:migrate/v20230401preview:PrivateEndpointConnectionOperation" },
                    new global::Pulumi.Alias { Type = "azure-native:migrate/v20230501preview:PrivateEndpointConnection" },
                    new global::Pulumi.Alias { Type = "azure-native:migrate/v20230501preview:PrivateEndpointConnectionOperation" },
                    new global::Pulumi.Alias { Type = "azure-native:migrate/v20230909preview:PrivateEndpointConnection" },
                    new global::Pulumi.Alias { Type = "azure-native:migrate/v20230909preview:PrivateEndpointConnectionOperation" },
                    new global::Pulumi.Alias { Type = "azure-native:migrate/v20240101preview:PrivateEndpointConnection" },
                    new global::Pulumi.Alias { Type = "azure-native:migrate/v20240101preview:PrivateEndpointConnectionOperation" },
                    new global::Pulumi.Alias { Type = "azure-native:migrate/v20240115:PrivateEndpointConnection" },
                    new global::Pulumi.Alias { Type = "azure-native:migrate/v20240303preview:PrivateEndpointConnection" },
                    new global::Pulumi.Alias { Type = "azure-native:migrate:PrivateEndpointConnectionOperation" },
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing PrivateEndpointConnection resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static PrivateEndpointConnection Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new PrivateEndpointConnection(name, id, options);
        }
    }

    public sealed class PrivateEndpointConnectionArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// For optimistic concurrency control.
        /// </summary>
        [Input("eTag")]
        public Input<string>? ETag { get; set; }

        /// <summary>
        /// Unique name of a private endpoint connection within a project.
        /// </summary>
        [Input("privateEndpointConnectionName")]
        public Input<string>? PrivateEndpointConnectionName { get; set; }

        /// <summary>
        /// Name of the Azure Migrate project.
        /// </summary>
        [Input("projectName", required: true)]
        public Input<string> ProjectName { get; set; } = null!;

        /// <summary>
        /// Properties of the private endpoint endpoint connection.
        /// </summary>
        [Input("properties", required: true)]
        public Input<Inputs.PrivateEndpointConnectionPropertiesArgs> Properties { get; set; } = null!;

        /// <summary>
        /// Name of the Azure Resource Group that project is part of.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        public PrivateEndpointConnectionArgs()
        {
        }
        public static new PrivateEndpointConnectionArgs Empty => new PrivateEndpointConnectionArgs();
    }
}
