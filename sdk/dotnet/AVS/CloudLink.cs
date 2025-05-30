// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.AVS
{
    /// <summary>
    /// A cloud link resource
    /// 
    /// Uses Azure REST API version 2023-09-01. In version 2.x of the Azure Native provider, it used API version 2022-05-01.
    /// 
    /// Other available API versions: 2022-05-01, 2023-03-01, 2024-09-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native avs [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
    /// </summary>
    [AzureNativeResourceType("azure-native:avs:CloudLink")]
    public partial class CloudLink : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        [Output("azureApiVersion")]
        public Output<string> AzureApiVersion { get; private set; } = null!;

        /// <summary>
        /// Identifier of the other private cloud participating in the link.
        /// </summary>
        [Output("linkedCloud")]
        public Output<string?> LinkedCloud { get; private set; } = null!;

        /// <summary>
        /// The name of the resource
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The provisioning state of the resource.
        /// </summary>
        [Output("provisioningState")]
        public Output<string> ProvisioningState { get; private set; } = null!;

        /// <summary>
        /// The state of the cloud link.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// Azure Resource Manager metadata containing createdBy and modifiedBy information.
        /// </summary>
        [Output("systemData")]
        public Output<Outputs.SystemDataResponse> SystemData { get; private set; } = null!;

        /// <summary>
        /// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a CloudLink resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public CloudLink(string name, CloudLinkArgs args, CustomResourceOptions? options = null)
            : base("azure-native:avs:CloudLink", name, args ?? new CloudLinkArgs(), MakeResourceOptions(options, ""))
        {
        }

        private CloudLink(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azure-native:avs:CloudLink", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new global::Pulumi.Alias { Type = "azure-native:avs/v20210601:CloudLink" },
                    new global::Pulumi.Alias { Type = "azure-native:avs/v20211201:CloudLink" },
                    new global::Pulumi.Alias { Type = "azure-native:avs/v20220501:CloudLink" },
                    new global::Pulumi.Alias { Type = "azure-native:avs/v20230301:CloudLink" },
                    new global::Pulumi.Alias { Type = "azure-native:avs/v20230901:CloudLink" },
                    new global::Pulumi.Alias { Type = "azure-native:avs/v20240901:CloudLink" },
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing CloudLink resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static CloudLink Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new CloudLink(name, id, options);
        }
    }

    public sealed class CloudLinkArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Name of the cloud link.
        /// </summary>
        [Input("cloudLinkName")]
        public Input<string>? CloudLinkName { get; set; }

        /// <summary>
        /// Identifier of the other private cloud participating in the link.
        /// </summary>
        [Input("linkedCloud")]
        public Input<string>? LinkedCloud { get; set; }

        /// <summary>
        /// Name of the private cloud
        /// </summary>
        [Input("privateCloudName", required: true)]
        public Input<string> PrivateCloudName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        public CloudLinkArgs()
        {
        }
        public static new CloudLinkArgs Empty => new CloudLinkArgs();
    }
}
