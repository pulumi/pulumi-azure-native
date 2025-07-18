// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Sql
{
    /// <summary>
    /// A long term retention policy.
    /// 
    /// Uses Azure REST API version 2023-08-01. In version 2.x of the Azure Native provider, it used API version 2021-11-01.
    /// 
    /// Other available API versions: 2020-02-02-preview, 2020-08-01-preview, 2020-11-01-preview, 2021-02-01-preview, 2021-05-01-preview, 2021-08-01-preview, 2021-11-01, 2021-11-01-preview, 2022-02-01-preview, 2022-05-01-preview, 2022-08-01-preview, 2022-11-01-preview, 2023-02-01-preview, 2023-05-01-preview, 2023-08-01-preview, 2024-05-01-preview, 2024-11-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native sql [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
    /// </summary>
    [AzureNativeResourceType("azure-native:sql:LongTermRetentionPolicy")]
    public partial class LongTermRetentionPolicy : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        [Output("azureApiVersion")]
        public Output<string> AzureApiVersion { get; private set; } = null!;

        /// <summary>
        /// The monthly retention policy for an LTR backup in an ISO 8601 format.
        /// </summary>
        [Output("monthlyRetention")]
        public Output<string?> MonthlyRetention { get; private set; } = null!;

        /// <summary>
        /// Resource name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Resource type.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;

        /// <summary>
        /// The week of year to take the yearly backup in an ISO 8601 format.
        /// </summary>
        [Output("weekOfYear")]
        public Output<int?> WeekOfYear { get; private set; } = null!;

        /// <summary>
        /// The weekly retention policy for an LTR backup in an ISO 8601 format.
        /// </summary>
        [Output("weeklyRetention")]
        public Output<string?> WeeklyRetention { get; private set; } = null!;

        /// <summary>
        /// The yearly retention policy for an LTR backup in an ISO 8601 format.
        /// </summary>
        [Output("yearlyRetention")]
        public Output<string?> YearlyRetention { get; private set; } = null!;


        /// <summary>
        /// Create a LongTermRetentionPolicy resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public LongTermRetentionPolicy(string name, LongTermRetentionPolicyArgs args, CustomResourceOptions? options = null)
            : base("azure-native:sql:LongTermRetentionPolicy", name, args ?? new LongTermRetentionPolicyArgs(), MakeResourceOptions(options, ""))
        {
        }

        private LongTermRetentionPolicy(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azure-native:sql:LongTermRetentionPolicy", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new global::Pulumi.Alias { Type = "azure-native:sql/v20170301preview:BackupLongTermRetentionPolicy" },
                    new global::Pulumi.Alias { Type = "azure-native:sql/v20170301preview:LongTermRetentionPolicy" },
                    new global::Pulumi.Alias { Type = "azure-native:sql/v20200202preview:LongTermRetentionPolicy" },
                    new global::Pulumi.Alias { Type = "azure-native:sql/v20200801preview:LongTermRetentionPolicy" },
                    new global::Pulumi.Alias { Type = "azure-native:sql/v20201101preview:LongTermRetentionPolicy" },
                    new global::Pulumi.Alias { Type = "azure-native:sql/v20210201preview:LongTermRetentionPolicy" },
                    new global::Pulumi.Alias { Type = "azure-native:sql/v20210501preview:LongTermRetentionPolicy" },
                    new global::Pulumi.Alias { Type = "azure-native:sql/v20210801preview:LongTermRetentionPolicy" },
                    new global::Pulumi.Alias { Type = "azure-native:sql/v20211101:LongTermRetentionPolicy" },
                    new global::Pulumi.Alias { Type = "azure-native:sql/v20211101preview:LongTermRetentionPolicy" },
                    new global::Pulumi.Alias { Type = "azure-native:sql/v20220201preview:LongTermRetentionPolicy" },
                    new global::Pulumi.Alias { Type = "azure-native:sql/v20220501preview:LongTermRetentionPolicy" },
                    new global::Pulumi.Alias { Type = "azure-native:sql/v20220801preview:LongTermRetentionPolicy" },
                    new global::Pulumi.Alias { Type = "azure-native:sql/v20221101preview:LongTermRetentionPolicy" },
                    new global::Pulumi.Alias { Type = "azure-native:sql/v20230201preview:LongTermRetentionPolicy" },
                    new global::Pulumi.Alias { Type = "azure-native:sql/v20230501preview:LongTermRetentionPolicy" },
                    new global::Pulumi.Alias { Type = "azure-native:sql/v20230801:LongTermRetentionPolicy" },
                    new global::Pulumi.Alias { Type = "azure-native:sql/v20230801preview:LongTermRetentionPolicy" },
                    new global::Pulumi.Alias { Type = "azure-native:sql/v20240501preview:LongTermRetentionPolicy" },
                    new global::Pulumi.Alias { Type = "azure-native:sql/v20241101preview:LongTermRetentionPolicy" },
                    new global::Pulumi.Alias { Type = "azure-native:sql:BackupLongTermRetentionPolicy" },
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing LongTermRetentionPolicy resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static LongTermRetentionPolicy Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new LongTermRetentionPolicy(name, id, options);
        }
    }

    public sealed class LongTermRetentionPolicyArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the database.
        /// </summary>
        [Input("databaseName", required: true)]
        public Input<string> DatabaseName { get; set; } = null!;

        /// <summary>
        /// The monthly retention policy for an LTR backup in an ISO 8601 format.
        /// </summary>
        [Input("monthlyRetention")]
        public Input<string>? MonthlyRetention { get; set; }

        /// <summary>
        /// The policy name. Should always be Default.
        /// </summary>
        [Input("policyName")]
        public Input<string>? PolicyName { get; set; }

        /// <summary>
        /// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the server.
        /// </summary>
        [Input("serverName", required: true)]
        public Input<string> ServerName { get; set; } = null!;

        /// <summary>
        /// The week of year to take the yearly backup in an ISO 8601 format.
        /// </summary>
        [Input("weekOfYear")]
        public Input<int>? WeekOfYear { get; set; }

        /// <summary>
        /// The weekly retention policy for an LTR backup in an ISO 8601 format.
        /// </summary>
        [Input("weeklyRetention")]
        public Input<string>? WeeklyRetention { get; set; }

        /// <summary>
        /// The yearly retention policy for an LTR backup in an ISO 8601 format.
        /// </summary>
        [Input("yearlyRetention")]
        public Input<string>? YearlyRetention { get; set; }

        public LongTermRetentionPolicyArgs()
        {
        }
        public static new LongTermRetentionPolicyArgs Empty => new LongTermRetentionPolicyArgs();
    }
}
