// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.DnsResolver
{
    /// <summary>
    /// Describes a virtual network link.
    /// 
    /// Uses Azure REST API version 2023-07-01-preview.
    /// 
    /// Other available API versions: 2020-04-01-preview, 2022-07-01, 2025-05-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native dnsresolver [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
    /// </summary>
    [AzureNativeResourceType("azure-native:dnsresolver:PrivateResolverVirtualNetworkLink")]
    public partial class PrivateResolverVirtualNetworkLink : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        [Output("azureApiVersion")]
        public Output<string> AzureApiVersion { get; private set; } = null!;

        /// <summary>
        /// ETag of the virtual network link.
        /// </summary>
        [Output("etag")]
        public Output<string> Etag { get; private set; } = null!;

        /// <summary>
        /// Metadata attached to the virtual network link.
        /// </summary>
        [Output("metadata")]
        public Output<ImmutableDictionary<string, string>?> Metadata { get; private set; } = null!;

        /// <summary>
        /// The name of the resource
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The current provisioning state of the virtual network link. This is a read-only property and any attempt to set this value will be ignored.
        /// </summary>
        [Output("provisioningState")]
        public Output<string> ProvisioningState { get; private set; } = null!;

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
        /// The reference to the virtual network. This cannot be changed after creation.
        /// </summary>
        [Output("virtualNetwork")]
        public Output<Outputs.SubResourceResponse> VirtualNetwork { get; private set; } = null!;


        /// <summary>
        /// Create a PrivateResolverVirtualNetworkLink resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public PrivateResolverVirtualNetworkLink(string name, PrivateResolverVirtualNetworkLinkArgs args, CustomResourceOptions? options = null)
            : base("azure-native:dnsresolver:PrivateResolverVirtualNetworkLink", name, args ?? new PrivateResolverVirtualNetworkLinkArgs(), MakeResourceOptions(options, ""))
        {
        }

        private PrivateResolverVirtualNetworkLink(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azure-native:dnsresolver:PrivateResolverVirtualNetworkLink", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new global::Pulumi.Alias { Type = "azure-native:dnsresolver/v20200401preview:PrivateResolverVirtualNetworkLink" },
                    new global::Pulumi.Alias { Type = "azure-native:dnsresolver/v20220701:PrivateResolverVirtualNetworkLink" },
                    new global::Pulumi.Alias { Type = "azure-native:dnsresolver/v20230701preview:PrivateResolverVirtualNetworkLink" },
                    new global::Pulumi.Alias { Type = "azure-native:dnsresolver/v20250501:PrivateResolverVirtualNetworkLink" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20200401preview:PrivateResolverVirtualNetworkLink" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20220701:PrivateResolverVirtualNetworkLink" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20230701preview:PrivateResolverVirtualNetworkLink" },
                    new global::Pulumi.Alias { Type = "azure-native:network:PrivateResolverVirtualNetworkLink" },
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing PrivateResolverVirtualNetworkLink resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static PrivateResolverVirtualNetworkLink Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new PrivateResolverVirtualNetworkLink(name, id, options);
        }
    }

    public sealed class PrivateResolverVirtualNetworkLinkArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the DNS forwarding ruleset.
        /// </summary>
        [Input("dnsForwardingRulesetName", required: true)]
        public Input<string> DnsForwardingRulesetName { get; set; } = null!;

        [Input("metadata")]
        private InputMap<string>? _metadata;

        /// <summary>
        /// Metadata attached to the virtual network link.
        /// </summary>
        public InputMap<string> Metadata
        {
            get => _metadata ?? (_metadata = new InputMap<string>());
            set => _metadata = value;
        }

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The reference to the virtual network. This cannot be changed after creation.
        /// </summary>
        [Input("virtualNetwork", required: true)]
        public Input<Inputs.SubResourceArgs> VirtualNetwork { get; set; } = null!;

        /// <summary>
        /// The name of the virtual network link.
        /// </summary>
        [Input("virtualNetworkLinkName")]
        public Input<string>? VirtualNetworkLinkName { get; set; }

        public PrivateResolverVirtualNetworkLinkArgs()
        {
        }
        public static new PrivateResolverVirtualNetworkLinkArgs Empty => new PrivateResolverVirtualNetworkLinkArgs();
    }
}
