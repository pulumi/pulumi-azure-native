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
    /// Describes an inbound endpoint for a DNS resolver.
    /// 
    /// Uses Azure REST API version 2023-07-01-preview.
    /// 
    /// Other available API versions: 2020-04-01-preview, 2022-07-01, 2025-05-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native dnsresolver [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
    /// </summary>
    [AzureNativeResourceType("azure-native:dnsresolver:InboundEndpoint")]
    public partial class InboundEndpoint : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        [Output("azureApiVersion")]
        public Output<string> AzureApiVersion { get; private set; } = null!;

        /// <summary>
        /// ETag of the inbound endpoint.
        /// </summary>
        [Output("etag")]
        public Output<string> Etag { get; private set; } = null!;

        /// <summary>
        /// IP configurations for the inbound endpoint.
        /// </summary>
        [Output("ipConfigurations")]
        public Output<ImmutableArray<Outputs.IpConfigurationResponse>> IpConfigurations { get; private set; } = null!;

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
        /// The current provisioning state of the inbound endpoint. This is a read-only property and any attempt to set this value will be ignored.
        /// </summary>
        [Output("provisioningState")]
        public Output<string> ProvisioningState { get; private set; } = null!;

        /// <summary>
        /// The resourceGuid property of the inbound endpoint resource.
        /// </summary>
        [Output("resourceGuid")]
        public Output<string> ResourceGuid { get; private set; } = null!;

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
        /// Create a InboundEndpoint resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public InboundEndpoint(string name, InboundEndpointArgs args, CustomResourceOptions? options = null)
            : base("azure-native:dnsresolver:InboundEndpoint", name, args ?? new InboundEndpointArgs(), MakeResourceOptions(options, ""))
        {
        }

        private InboundEndpoint(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azure-native:dnsresolver:InboundEndpoint", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new global::Pulumi.Alias { Type = "azure-native:dnsresolver/v20200401preview:InboundEndpoint" },
                    new global::Pulumi.Alias { Type = "azure-native:dnsresolver/v20220701:InboundEndpoint" },
                    new global::Pulumi.Alias { Type = "azure-native:dnsresolver/v20230701preview:InboundEndpoint" },
                    new global::Pulumi.Alias { Type = "azure-native:dnsresolver/v20250501:InboundEndpoint" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20200401preview:InboundEndpoint" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20220701:InboundEndpoint" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20230701preview:InboundEndpoint" },
                    new global::Pulumi.Alias { Type = "azure-native:network:InboundEndpoint" },
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing InboundEndpoint resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static InboundEndpoint Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new InboundEndpoint(name, id, options);
        }
    }

    public sealed class InboundEndpointArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the DNS resolver.
        /// </summary>
        [Input("dnsResolverName", required: true)]
        public Input<string> DnsResolverName { get; set; } = null!;

        /// <summary>
        /// The name of the inbound endpoint for the DNS resolver.
        /// </summary>
        [Input("inboundEndpointName")]
        public Input<string>? InboundEndpointName { get; set; }

        [Input("ipConfigurations", required: true)]
        private InputList<Inputs.IpConfigurationArgs>? _ipConfigurations;

        /// <summary>
        /// IP configurations for the inbound endpoint.
        /// </summary>
        public InputList<Inputs.IpConfigurationArgs> IpConfigurations
        {
            get => _ipConfigurations ?? (_ipConfigurations = new InputList<Inputs.IpConfigurationArgs>());
            set => _ipConfigurations = value;
        }

        /// <summary>
        /// The geo-location where the resource lives
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
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

        public InboundEndpointArgs()
        {
        }
        public static new InboundEndpointArgs Empty => new InboundEndpointArgs();
    }
}
