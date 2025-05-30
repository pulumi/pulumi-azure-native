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
    /// NSX Segment
    /// 
    /// Uses Azure REST API version 2023-09-01. In version 2.x of the Azure Native provider, it used API version 2022-05-01.
    /// 
    /// Other available API versions: 2022-05-01, 2023-03-01, 2024-09-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native avs [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
    /// </summary>
    [AzureNativeResourceType("azure-native:avs:WorkloadNetworkSegment")]
    public partial class WorkloadNetworkSegment : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        [Output("azureApiVersion")]
        public Output<string> AzureApiVersion { get; private set; } = null!;

        /// <summary>
        /// Gateway which to connect segment to.
        /// </summary>
        [Output("connectedGateway")]
        public Output<string?> ConnectedGateway { get; private set; } = null!;

        /// <summary>
        /// Display name of the segment.
        /// </summary>
        [Output("displayName")]
        public Output<string?> DisplayName { get; private set; } = null!;

        /// <summary>
        /// The name of the resource
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Port Vif which segment is associated with.
        /// </summary>
        [Output("portVif")]
        public Output<ImmutableArray<Outputs.WorkloadNetworkSegmentPortVifResponse>> PortVif { get; private set; } = null!;

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
        /// Segment status.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// Subnet which to connect segment to.
        /// </summary>
        [Output("subnet")]
        public Output<Outputs.WorkloadNetworkSegmentSubnetResponse?> Subnet { get; private set; } = null!;

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
        /// Create a WorkloadNetworkSegment resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public WorkloadNetworkSegment(string name, WorkloadNetworkSegmentArgs args, CustomResourceOptions? options = null)
            : base("azure-native:avs:WorkloadNetworkSegment", name, args ?? new WorkloadNetworkSegmentArgs(), MakeResourceOptions(options, ""))
        {
        }

        private WorkloadNetworkSegment(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azure-native:avs:WorkloadNetworkSegment", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new global::Pulumi.Alias { Type = "azure-native:avs/v20200717preview:WorkloadNetworkSegment" },
                    new global::Pulumi.Alias { Type = "azure-native:avs/v20210101preview:WorkloadNetworkSegment" },
                    new global::Pulumi.Alias { Type = "azure-native:avs/v20210601:WorkloadNetworkSegment" },
                    new global::Pulumi.Alias { Type = "azure-native:avs/v20211201:WorkloadNetworkSegment" },
                    new global::Pulumi.Alias { Type = "azure-native:avs/v20220501:WorkloadNetworkSegment" },
                    new global::Pulumi.Alias { Type = "azure-native:avs/v20230301:WorkloadNetworkSegment" },
                    new global::Pulumi.Alias { Type = "azure-native:avs/v20230901:WorkloadNetworkSegment" },
                    new global::Pulumi.Alias { Type = "azure-native:avs/v20240901:WorkloadNetworkSegment" },
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing WorkloadNetworkSegment resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static WorkloadNetworkSegment Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new WorkloadNetworkSegment(name, id, options);
        }
    }

    public sealed class WorkloadNetworkSegmentArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Gateway which to connect segment to.
        /// </summary>
        [Input("connectedGateway")]
        public Input<string>? ConnectedGateway { get; set; }

        /// <summary>
        /// Display name of the segment.
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

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
        /// The ID of the NSX Segment
        /// </summary>
        [Input("segmentId")]
        public Input<string>? SegmentId { get; set; }

        /// <summary>
        /// Subnet which to connect segment to.
        /// </summary>
        [Input("subnet")]
        public Input<Inputs.WorkloadNetworkSegmentSubnetArgs>? Subnet { get; set; }

        public WorkloadNetworkSegmentArgs()
        {
        }
        public static new WorkloadNetworkSegmentArgs Empty => new WorkloadNetworkSegmentArgs();
    }
}
