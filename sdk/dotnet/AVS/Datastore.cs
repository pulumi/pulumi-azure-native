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
    /// A datastore resource
    /// 
    /// Uses Azure REST API version 2023-09-01. In version 2.x of the Azure Native provider, it used API version 2022-05-01.
    /// 
    /// Other available API versions: 2022-05-01, 2023-03-01, 2024-09-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native avs [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
    /// </summary>
    [AzureNativeResourceType("azure-native:avs:Datastore")]
    public partial class Datastore : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        [Output("azureApiVersion")]
        public Output<string> AzureApiVersion { get; private set; } = null!;

        /// <summary>
        /// An iSCSI volume
        /// </summary>
        [Output("diskPoolVolume")]
        public Output<Outputs.DiskPoolVolumeResponse?> DiskPoolVolume { get; private set; } = null!;

        /// <summary>
        /// An Elastic SAN volume
        /// </summary>
        [Output("elasticSanVolume")]
        public Output<Outputs.ElasticSanVolumeResponse?> ElasticSanVolume { get; private set; } = null!;

        /// <summary>
        /// The name of the resource
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// An Azure NetApp Files volume
        /// </summary>
        [Output("netAppVolume")]
        public Output<Outputs.NetAppVolumeResponse?> NetAppVolume { get; private set; } = null!;

        /// <summary>
        /// The state of the datastore provisioning
        /// </summary>
        [Output("provisioningState")]
        public Output<string> ProvisioningState { get; private set; } = null!;

        /// <summary>
        /// The operational status of the datastore
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
        /// Create a Datastore resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Datastore(string name, DatastoreArgs args, CustomResourceOptions? options = null)
            : base("azure-native:avs:Datastore", name, args ?? new DatastoreArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Datastore(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azure-native:avs:Datastore", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new global::Pulumi.Alias { Type = "azure-native:avs/v20210101preview:Datastore" },
                    new global::Pulumi.Alias { Type = "azure-native:avs/v20210601:Datastore" },
                    new global::Pulumi.Alias { Type = "azure-native:avs/v20211201:Datastore" },
                    new global::Pulumi.Alias { Type = "azure-native:avs/v20220501:Datastore" },
                    new global::Pulumi.Alias { Type = "azure-native:avs/v20230301:Datastore" },
                    new global::Pulumi.Alias { Type = "azure-native:avs/v20230901:Datastore" },
                    new global::Pulumi.Alias { Type = "azure-native:avs/v20240901:Datastore" },
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Datastore resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Datastore Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Datastore(name, id, options);
        }
    }

    public sealed class DatastoreArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Name of the cluster
        /// </summary>
        [Input("clusterName", required: true)]
        public Input<string> ClusterName { get; set; } = null!;

        /// <summary>
        /// Name of the datastore
        /// </summary>
        [Input("datastoreName")]
        public Input<string>? DatastoreName { get; set; }

        /// <summary>
        /// An iSCSI volume
        /// </summary>
        [Input("diskPoolVolume")]
        public Input<Inputs.DiskPoolVolumeArgs>? DiskPoolVolume { get; set; }

        /// <summary>
        /// An Elastic SAN volume
        /// </summary>
        [Input("elasticSanVolume")]
        public Input<Inputs.ElasticSanVolumeArgs>? ElasticSanVolume { get; set; }

        /// <summary>
        /// An Azure NetApp Files volume
        /// </summary>
        [Input("netAppVolume")]
        public Input<Inputs.NetAppVolumeArgs>? NetAppVolume { get; set; }

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

        public DatastoreArgs()
        {
        }
        public static new DatastoreArgs Empty => new DatastoreArgs();
    }
}
