// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.CognitiveServices
{
    /// <summary>
    /// The Private Endpoint Connection resource.
    /// 
    /// Uses Azure REST API version 2024-10-01. In version 2.x of the Azure Native provider, it used API version 2023-05-01.
    /// 
    /// Other available API versions: 2023-05-01, 2023-10-01-preview, 2024-04-01-preview, 2024-06-01-preview, 2025-04-01-preview, 2025-06-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native cognitiveservices [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
    /// </summary>
    [AzureNativeResourceType("azure-native:cognitiveservices:PrivateEndpointConnection")]
    public partial class PrivateEndpointConnection : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        [Output("azureApiVersion")]
        public Output<string> AzureApiVersion { get; private set; } = null!;

        /// <summary>
        /// Resource Etag.
        /// </summary>
        [Output("etag")]
        public Output<string> Etag { get; private set; } = null!;

        /// <summary>
        /// The location of the private endpoint connection
        /// </summary>
        [Output("location")]
        public Output<string?> Location { get; private set; } = null!;

        /// <summary>
        /// The name of the resource
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Resource properties.
        /// </summary>
        [Output("properties")]
        public Output<Outputs.PrivateEndpointConnectionPropertiesResponse> Properties { get; private set; } = null!;

        /// <summary>
        /// Metadata pertaining to creation and last modification of the resource.
        /// </summary>
        [Output("systemData")]
        public Output<Outputs.SystemDataResponse> SystemData { get; private set; } = null!;

        /// <summary>
        /// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
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
            : base("azure-native:cognitiveservices:PrivateEndpointConnection", name, args ?? new PrivateEndpointConnectionArgs(), MakeResourceOptions(options, ""))
        {
        }

        private PrivateEndpointConnection(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azure-native:cognitiveservices:PrivateEndpointConnection", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new global::Pulumi.Alias { Type = "azure-native:cognitiveservices/v20170418:PrivateEndpointConnection" },
                    new global::Pulumi.Alias { Type = "azure-native:cognitiveservices/v20210430:PrivateEndpointConnection" },
                    new global::Pulumi.Alias { Type = "azure-native:cognitiveservices/v20211001:PrivateEndpointConnection" },
                    new global::Pulumi.Alias { Type = "azure-native:cognitiveservices/v20220301:PrivateEndpointConnection" },
                    new global::Pulumi.Alias { Type = "azure-native:cognitiveservices/v20221001:PrivateEndpointConnection" },
                    new global::Pulumi.Alias { Type = "azure-native:cognitiveservices/v20221201:PrivateEndpointConnection" },
                    new global::Pulumi.Alias { Type = "azure-native:cognitiveservices/v20230501:PrivateEndpointConnection" },
                    new global::Pulumi.Alias { Type = "azure-native:cognitiveservices/v20231001preview:PrivateEndpointConnection" },
                    new global::Pulumi.Alias { Type = "azure-native:cognitiveservices/v20240401preview:PrivateEndpointConnection" },
                    new global::Pulumi.Alias { Type = "azure-native:cognitiveservices/v20240601preview:PrivateEndpointConnection" },
                    new global::Pulumi.Alias { Type = "azure-native:cognitiveservices/v20241001:PrivateEndpointConnection" },
                    new global::Pulumi.Alias { Type = "azure-native:cognitiveservices/v20250401preview:PrivateEndpointConnection" },
                    new global::Pulumi.Alias { Type = "azure-native:cognitiveservices/v20250601:PrivateEndpointConnection" },
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
        /// The name of Cognitive Services account.
        /// </summary>
        [Input("accountName", required: true)]
        public Input<string> AccountName { get; set; } = null!;

        /// <summary>
        /// The location of the private endpoint connection
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// The name of the private endpoint connection associated with the Cognitive Services Account
        /// </summary>
        [Input("privateEndpointConnectionName")]
        public Input<string>? PrivateEndpointConnectionName { get; set; }

        /// <summary>
        /// Resource properties.
        /// </summary>
        [Input("properties")]
        public Input<Inputs.PrivateEndpointConnectionPropertiesArgs>? Properties { get; set; }

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        public PrivateEndpointConnectionArgs()
        {
        }
        public static new PrivateEndpointConnectionArgs Empty => new PrivateEndpointConnectionArgs();
    }
}
