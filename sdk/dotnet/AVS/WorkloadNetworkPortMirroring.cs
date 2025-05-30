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
    /// NSX Port Mirroring
    /// 
    /// Uses Azure REST API version 2023-09-01. In version 2.x of the Azure Native provider, it used API version 2022-05-01.
    /// 
    /// Other available API versions: 2022-05-01, 2023-03-01, 2024-09-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native avs [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
    /// </summary>
    [AzureNativeResourceType("azure-native:avs:WorkloadNetworkPortMirroring")]
    public partial class WorkloadNetworkPortMirroring : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        [Output("azureApiVersion")]
        public Output<string> AzureApiVersion { get; private set; } = null!;

        /// <summary>
        /// Destination VM Group.
        /// </summary>
        [Output("destination")]
        public Output<string?> Destination { get; private set; } = null!;

        /// <summary>
        /// Direction of port mirroring profile.
        /// </summary>
        [Output("direction")]
        public Output<string?> Direction { get; private set; } = null!;

        /// <summary>
        /// Display name of the port mirroring profile.
        /// </summary>
        [Output("displayName")]
        public Output<string?> DisplayName { get; private set; } = null!;

        /// <summary>
        /// The name of the resource
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The provisioning state
        /// </summary>
        [Output("provisioningState")]
        public Output<string> ProvisioningState { get; private set; } = null!;

        /// <summary>
        /// NSX revision number.
        /// </summary>
        [Output("revision")]
        public Output<double?> Revision { get; private set; } = null!;

        /// <summary>
        /// Source VM Group.
        /// </summary>
        [Output("source")]
        public Output<string?> Source { get; private set; } = null!;

        /// <summary>
        /// Port Mirroring Status.
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
        /// Create a WorkloadNetworkPortMirroring resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public WorkloadNetworkPortMirroring(string name, WorkloadNetworkPortMirroringArgs args, CustomResourceOptions? options = null)
            : base("azure-native:avs:WorkloadNetworkPortMirroring", name, args ?? new WorkloadNetworkPortMirroringArgs(), MakeResourceOptions(options, ""))
        {
        }

        private WorkloadNetworkPortMirroring(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azure-native:avs:WorkloadNetworkPortMirroring", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new global::Pulumi.Alias { Type = "azure-native:avs/v20200717preview:WorkloadNetworkPortMirroring" },
                    new global::Pulumi.Alias { Type = "azure-native:avs/v20210101preview:WorkloadNetworkPortMirroring" },
                    new global::Pulumi.Alias { Type = "azure-native:avs/v20210601:WorkloadNetworkPortMirroring" },
                    new global::Pulumi.Alias { Type = "azure-native:avs/v20211201:WorkloadNetworkPortMirroring" },
                    new global::Pulumi.Alias { Type = "azure-native:avs/v20220501:WorkloadNetworkPortMirroring" },
                    new global::Pulumi.Alias { Type = "azure-native:avs/v20230301:WorkloadNetworkPortMirroring" },
                    new global::Pulumi.Alias { Type = "azure-native:avs/v20230901:WorkloadNetworkPortMirroring" },
                    new global::Pulumi.Alias { Type = "azure-native:avs/v20240901:WorkloadNetworkPortMirroring" },
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing WorkloadNetworkPortMirroring resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static WorkloadNetworkPortMirroring Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new WorkloadNetworkPortMirroring(name, id, options);
        }
    }

    public sealed class WorkloadNetworkPortMirroringArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Destination VM Group.
        /// </summary>
        [Input("destination")]
        public Input<string>? Destination { get; set; }

        /// <summary>
        /// Direction of port mirroring profile.
        /// </summary>
        [Input("direction")]
        public InputUnion<string, Pulumi.AzureNative.AVS.PortMirroringDirectionEnum>? Direction { get; set; }

        /// <summary>
        /// Display name of the port mirroring profile.
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        /// <summary>
        /// ID of the NSX port mirroring profile.
        /// </summary>
        [Input("portMirroringId")]
        public Input<string>? PortMirroringId { get; set; }

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

        /// <summary>
        /// NSX revision number.
        /// </summary>
        [Input("revision")]
        public Input<double>? Revision { get; set; }

        /// <summary>
        /// Source VM Group.
        /// </summary>
        [Input("source")]
        public Input<string>? Source { get; set; }

        public WorkloadNetworkPortMirroringArgs()
        {
        }
        public static new WorkloadNetworkPortMirroringArgs Empty => new WorkloadNetworkPortMirroringArgs();
    }
}
