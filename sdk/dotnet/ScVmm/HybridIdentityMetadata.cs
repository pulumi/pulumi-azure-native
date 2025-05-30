// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.ScVmm
{
    /// <summary>
    /// Defines the HybridIdentityMetadata.
    /// 
    /// Uses Azure REST API version 2023-04-01-preview. In version 2.x of the Azure Native provider, it used API version 2022-05-21-preview.
    /// 
    /// Other available API versions: 2022-05-21-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native scvmm [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
    /// </summary>
    [AzureNativeResourceType("azure-native:scvmm:HybridIdentityMetadata")]
    public partial class HybridIdentityMetadata : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        [Output("azureApiVersion")]
        public Output<string> AzureApiVersion { get; private set; } = null!;

        /// <summary>
        /// The identity of the resource.
        /// </summary>
        [Output("identity")]
        public Output<Outputs.IdentityResponse> Identity { get; private set; } = null!;

        /// <summary>
        /// The name of the resource
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Gets or sets the provisioning state.
        /// </summary>
        [Output("provisioningState")]
        public Output<string> ProvisioningState { get; private set; } = null!;

        /// <summary>
        /// Gets or sets the Public Key.
        /// </summary>
        [Output("publicKey")]
        public Output<string?> PublicKey { get; private set; } = null!;

        /// <summary>
        /// Gets or sets the Vm Id.
        /// </summary>
        [Output("resourceUid")]
        public Output<string?> ResourceUid { get; private set; } = null!;

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
        /// Create a HybridIdentityMetadata resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public HybridIdentityMetadata(string name, HybridIdentityMetadataArgs args, CustomResourceOptions? options = null)
            : base("azure-native:scvmm:HybridIdentityMetadata", name, args ?? new HybridIdentityMetadataArgs(), MakeResourceOptions(options, ""))
        {
        }

        private HybridIdentityMetadata(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azure-native:scvmm:HybridIdentityMetadata", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new global::Pulumi.Alias { Type = "azure-native:scvmm/v20220521preview:HybridIdentityMetadata" },
                    new global::Pulumi.Alias { Type = "azure-native:scvmm/v20230401preview:HybridIdentityMetadata" },
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing HybridIdentityMetadata resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static HybridIdentityMetadata Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new HybridIdentityMetadata(name, id, options);
        }
    }

    public sealed class HybridIdentityMetadataArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Name of the hybridIdentityMetadata.
        /// </summary>
        [Input("metadataName")]
        public Input<string>? MetadataName { get; set; }

        /// <summary>
        /// Gets or sets the Public Key.
        /// </summary>
        [Input("publicKey")]
        public Input<string>? PublicKey { get; set; }

        /// <summary>
        /// The name of the resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// Gets or sets the Vm Id.
        /// </summary>
        [Input("resourceUid")]
        public Input<string>? ResourceUid { get; set; }

        /// <summary>
        /// Name of the vm.
        /// </summary>
        [Input("virtualMachineName", required: true)]
        public Input<string> VirtualMachineName { get; set; } = null!;

        public HybridIdentityMetadataArgs()
        {
        }
        public static new HybridIdentityMetadataArgs Empty => new HybridIdentityMetadataArgs();
    }
}
