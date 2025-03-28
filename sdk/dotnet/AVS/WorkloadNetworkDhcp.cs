// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.AVS
{
    /// <summary>
    /// NSX DHCP
    /// 
    /// Uses Azure REST API version 2022-05-01. In version 1.x of the Azure Native provider, it used API version 2020-07-17-preview.
    /// 
    /// Other available API versions: 2021-01-01-preview, 2023-03-01, 2023-09-01.
    /// </summary>
    [AzureNativeResourceType("azure-native:avs:WorkloadNetworkDhcp")]
    public partial class WorkloadNetworkDhcp : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Resource name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// DHCP properties.
        /// </summary>
        [Output("properties")]
        public Output<Union<Outputs.WorkloadNetworkDhcpRelayResponse, Outputs.WorkloadNetworkDhcpServerResponse>> Properties { get; private set; } = null!;

        /// <summary>
        /// Resource type.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a WorkloadNetworkDhcp resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public WorkloadNetworkDhcp(string name, WorkloadNetworkDhcpArgs args, CustomResourceOptions? options = null)
            : base("azure-native:avs:WorkloadNetworkDhcp", name, args ?? new WorkloadNetworkDhcpArgs(), MakeResourceOptions(options, ""))
        {
        }

        private WorkloadNetworkDhcp(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azure-native:avs:WorkloadNetworkDhcp", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new global::Pulumi.Alias { Type = "azure-native:avs/v20200717preview:WorkloadNetworkDhcp" },
                    new global::Pulumi.Alias { Type = "azure-native:avs/v20210101preview:WorkloadNetworkDhcp" },
                    new global::Pulumi.Alias { Type = "azure-native:avs/v20210601:WorkloadNetworkDhcp" },
                    new global::Pulumi.Alias { Type = "azure-native:avs/v20211201:WorkloadNetworkDhcp" },
                    new global::Pulumi.Alias { Type = "azure-native:avs/v20220501:WorkloadNetworkDhcp" },
                    new global::Pulumi.Alias { Type = "azure-native:avs/v20230301:WorkloadNetworkDhcp" },
                    new global::Pulumi.Alias { Type = "azure-native:avs/v20230901:WorkloadNetworkDhcp" },
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing WorkloadNetworkDhcp resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static WorkloadNetworkDhcp Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new WorkloadNetworkDhcp(name, id, options);
        }
    }

    public sealed class WorkloadNetworkDhcpArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// NSX DHCP identifier. Generally the same as the DHCP display name
        /// </summary>
        [Input("dhcpId")]
        public Input<string>? DhcpId { get; set; }

        /// <summary>
        /// Name of the private cloud
        /// </summary>
        [Input("privateCloudName", required: true)]
        public Input<string> PrivateCloudName { get; set; } = null!;

        /// <summary>
        /// DHCP properties.
        /// </summary>
        [Input("properties")]
        public InputUnion<Inputs.WorkloadNetworkDhcpRelayArgs, Inputs.WorkloadNetworkDhcpServerArgs>? Properties { get; set; }

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        public WorkloadNetworkDhcpArgs()
        {
        }
        public static new WorkloadNetworkDhcpArgs Empty => new WorkloadNetworkDhcpArgs();
    }
}
