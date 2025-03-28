// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Sql.V20240501Preview
{
    /// <summary>
    /// Distributed availability group between box and Sql Managed Instance.
    /// </summary>
    [AzureNativeResourceType("azure-native:sql/v20240501preview:DistributedAvailabilityGroup")]
    public partial class DistributedAvailabilityGroup : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Databases in the distributed availability group
        /// </summary>
        [Output("databases")]
        public Output<ImmutableArray<Outputs.DistributedAvailabilityGroupDatabaseResponse>> Databases { get; private set; } = null!;

        /// <summary>
        /// ID of the distributed availability group
        /// </summary>
        [Output("distributedAvailabilityGroupId")]
        public Output<string> DistributedAvailabilityGroupId { get; private set; } = null!;

        /// <summary>
        /// Name of the distributed availability group
        /// </summary>
        [Output("distributedAvailabilityGroupName")]
        public Output<string> DistributedAvailabilityGroupName { get; private set; } = null!;

        /// <summary>
        /// The link failover mode - can be Manual if intended to be used for two-way failover with a supported SQL Server, or None for one-way failover to Azure.
        /// </summary>
        [Output("failoverMode")]
        public Output<string?> FailoverMode { get; private set; } = null!;

        /// <summary>
        /// Managed instance side availability group name
        /// </summary>
        [Output("instanceAvailabilityGroupName")]
        public Output<string?> InstanceAvailabilityGroupName { get; private set; } = null!;

        /// <summary>
        /// Managed instance side link role
        /// </summary>
        [Output("instanceLinkRole")]
        public Output<string?> InstanceLinkRole { get; private set; } = null!;

        /// <summary>
        /// Resource name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// SQL server side availability group name
        /// </summary>
        [Output("partnerAvailabilityGroupName")]
        public Output<string?> PartnerAvailabilityGroupName { get; private set; } = null!;

        /// <summary>
        /// SQL server side endpoint - IP or DNS resolvable name
        /// </summary>
        [Output("partnerEndpoint")]
        public Output<string?> PartnerEndpoint { get; private set; } = null!;

        /// <summary>
        /// SQL server side link role
        /// </summary>
        [Output("partnerLinkRole")]
        public Output<string> PartnerLinkRole { get; private set; } = null!;

        /// <summary>
        /// Replication mode of the link
        /// </summary>
        [Output("replicationMode")]
        public Output<string?> ReplicationMode { get; private set; } = null!;

        /// <summary>
        /// Database seeding mode – can be Automatic (default), or Manual for supported scenarios.
        /// </summary>
        [Output("seedingMode")]
        public Output<string?> SeedingMode { get; private set; } = null!;

        /// <summary>
        /// Resource type.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a DistributedAvailabilityGroup resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public DistributedAvailabilityGroup(string name, DistributedAvailabilityGroupArgs args, CustomResourceOptions? options = null)
            : base("azure-native:sql/v20240501preview:DistributedAvailabilityGroup", name, args ?? new DistributedAvailabilityGroupArgs(), MakeResourceOptions(options, ""))
        {
        }

        private DistributedAvailabilityGroup(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azure-native:sql/v20240501preview:DistributedAvailabilityGroup", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new global::Pulumi.Alias { Type = "azure-native:sql/v20210501preview:DistributedAvailabilityGroup" },
                    new global::Pulumi.Alias { Type = "azure-native:sql/v20210801preview:DistributedAvailabilityGroup" },
                    new global::Pulumi.Alias { Type = "azure-native:sql/v20211101:DistributedAvailabilityGroup" },
                    new global::Pulumi.Alias { Type = "azure-native:sql/v20211101preview:DistributedAvailabilityGroup" },
                    new global::Pulumi.Alias { Type = "azure-native:sql/v20220201preview:DistributedAvailabilityGroup" },
                    new global::Pulumi.Alias { Type = "azure-native:sql/v20220501preview:DistributedAvailabilityGroup" },
                    new global::Pulumi.Alias { Type = "azure-native:sql/v20220801preview:DistributedAvailabilityGroup" },
                    new global::Pulumi.Alias { Type = "azure-native:sql/v20221101preview:DistributedAvailabilityGroup" },
                    new global::Pulumi.Alias { Type = "azure-native:sql/v20230201preview:DistributedAvailabilityGroup" },
                    new global::Pulumi.Alias { Type = "azure-native:sql/v20230501preview:DistributedAvailabilityGroup" },
                    new global::Pulumi.Alias { Type = "azure-native:sql/v20230801:DistributedAvailabilityGroup" },
                    new global::Pulumi.Alias { Type = "azure-native:sql/v20230801preview:DistributedAvailabilityGroup" },
                    new global::Pulumi.Alias { Type = "azure-native:sql:DistributedAvailabilityGroup" },
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing DistributedAvailabilityGroup resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static DistributedAvailabilityGroup Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new DistributedAvailabilityGroup(name, id, options);
        }
    }

    public sealed class DistributedAvailabilityGroupArgs : global::Pulumi.ResourceArgs
    {
        [Input("databases")]
        private InputList<Inputs.DistributedAvailabilityGroupDatabaseArgs>? _databases;

        /// <summary>
        /// Databases in the distributed availability group
        /// </summary>
        public InputList<Inputs.DistributedAvailabilityGroupDatabaseArgs> Databases
        {
            get => _databases ?? (_databases = new InputList<Inputs.DistributedAvailabilityGroupDatabaseArgs>());
            set => _databases = value;
        }

        /// <summary>
        /// The distributed availability group name.
        /// </summary>
        [Input("distributedAvailabilityGroupName")]
        public Input<string>? DistributedAvailabilityGroupName { get; set; }

        /// <summary>
        /// The link failover mode - can be Manual if intended to be used for two-way failover with a supported SQL Server, or None for one-way failover to Azure.
        /// </summary>
        [Input("failoverMode")]
        public InputUnion<string, Pulumi.AzureNative.Sql.V20240501Preview.FailoverModeType>? FailoverMode { get; set; }

        /// <summary>
        /// Managed instance side availability group name
        /// </summary>
        [Input("instanceAvailabilityGroupName")]
        public Input<string>? InstanceAvailabilityGroupName { get; set; }

        /// <summary>
        /// Managed instance side link role
        /// </summary>
        [Input("instanceLinkRole")]
        public InputUnion<string, Pulumi.AzureNative.Sql.V20240501Preview.LinkRole>? InstanceLinkRole { get; set; }

        /// <summary>
        /// The name of the managed instance.
        /// </summary>
        [Input("managedInstanceName", required: true)]
        public Input<string> ManagedInstanceName { get; set; } = null!;

        /// <summary>
        /// SQL server side availability group name
        /// </summary>
        [Input("partnerAvailabilityGroupName")]
        public Input<string>? PartnerAvailabilityGroupName { get; set; }

        /// <summary>
        /// SQL server side endpoint - IP or DNS resolvable name
        /// </summary>
        [Input("partnerEndpoint")]
        public Input<string>? PartnerEndpoint { get; set; }

        /// <summary>
        /// Replication mode of the link
        /// </summary>
        [Input("replicationMode")]
        public InputUnion<string, Pulumi.AzureNative.Sql.V20240501Preview.ReplicationModeType>? ReplicationMode { get; set; }

        /// <summary>
        /// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// Database seeding mode – can be Automatic (default), or Manual for supported scenarios.
        /// </summary>
        [Input("seedingMode")]
        public InputUnion<string, Pulumi.AzureNative.Sql.V20240501Preview.SeedingModeType>? SeedingMode { get; set; }

        public DistributedAvailabilityGroupArgs()
        {
        }
        public static new DistributedAvailabilityGroupArgs Empty => new DistributedAvailabilityGroupArgs();
    }
}
