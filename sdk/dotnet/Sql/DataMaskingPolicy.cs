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
    /// A database data masking policy.
    /// 
    /// Uses Azure REST API version 2023-08-01. In version 2.x of the Azure Native provider, it used API version 2021-11-01.
    /// 
    /// Other available API versions: 2014-04-01, 2021-11-01, 2022-02-01-preview, 2022-05-01-preview, 2022-08-01-preview, 2022-11-01-preview, 2023-02-01-preview, 2023-05-01-preview, 2023-08-01-preview, 2024-05-01-preview, 2024-11-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native sql [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
    /// </summary>
    [AzureNativeResourceType("azure-native:sql:DataMaskingPolicy")]
    public partial class DataMaskingPolicy : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The list of the application principals. This is a legacy parameter and is no longer used.
        /// </summary>
        [Output("applicationPrincipals")]
        public Output<string> ApplicationPrincipals { get; private set; } = null!;

        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        [Output("azureApiVersion")]
        public Output<string> AzureApiVersion { get; private set; } = null!;

        /// <summary>
        /// The state of the data masking policy.
        /// </summary>
        [Output("dataMaskingState")]
        public Output<string> DataMaskingState { get; private set; } = null!;

        /// <summary>
        /// The list of the exempt principals. Specifies the semicolon-separated list of database users for which the data masking policy does not apply. The specified users receive data results without masking for all of the database queries.
        /// </summary>
        [Output("exemptPrincipals")]
        public Output<string?> ExemptPrincipals { get; private set; } = null!;

        /// <summary>
        /// The kind of Data Masking Policy. Metadata, used for Azure portal.
        /// </summary>
        [Output("kind")]
        public Output<string> Kind { get; private set; } = null!;

        /// <summary>
        /// The location of the data masking policy.
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// The masking level. This is a legacy parameter and is no longer used.
        /// </summary>
        [Output("maskingLevel")]
        public Output<string> MaskingLevel { get; private set; } = null!;

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
        /// Create a DataMaskingPolicy resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public DataMaskingPolicy(string name, DataMaskingPolicyArgs args, CustomResourceOptions? options = null)
            : base("azure-native:sql:DataMaskingPolicy", name, args ?? new DataMaskingPolicyArgs(), MakeResourceOptions(options, ""))
        {
        }

        private DataMaskingPolicy(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azure-native:sql:DataMaskingPolicy", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new global::Pulumi.Alias { Type = "azure-native:sql/v20140401:DataMaskingPolicy" },
                    new global::Pulumi.Alias { Type = "azure-native:sql/v20211101:DataMaskingPolicy" },
                    new global::Pulumi.Alias { Type = "azure-native:sql/v20220201preview:DataMaskingPolicy" },
                    new global::Pulumi.Alias { Type = "azure-native:sql/v20220501preview:DataMaskingPolicy" },
                    new global::Pulumi.Alias { Type = "azure-native:sql/v20220801preview:DataMaskingPolicy" },
                    new global::Pulumi.Alias { Type = "azure-native:sql/v20221101preview:DataMaskingPolicy" },
                    new global::Pulumi.Alias { Type = "azure-native:sql/v20230201preview:DataMaskingPolicy" },
                    new global::Pulumi.Alias { Type = "azure-native:sql/v20230501preview:DataMaskingPolicy" },
                    new global::Pulumi.Alias { Type = "azure-native:sql/v20230801:DataMaskingPolicy" },
                    new global::Pulumi.Alias { Type = "azure-native:sql/v20230801preview:DataMaskingPolicy" },
                    new global::Pulumi.Alias { Type = "azure-native:sql/v20240501preview:DataMaskingPolicy" },
                    new global::Pulumi.Alias { Type = "azure-native:sql/v20241101preview:DataMaskingPolicy" },
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing DataMaskingPolicy resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static DataMaskingPolicy Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new DataMaskingPolicy(name, id, options);
        }
    }

    public sealed class DataMaskingPolicyArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the database for which the data masking policy applies.
        /// </summary>
        [Input("dataMaskingPolicyName")]
        public Input<string>? DataMaskingPolicyName { get; set; }

        /// <summary>
        /// The state of the data masking policy.
        /// </summary>
        [Input("dataMaskingState", required: true)]
        public Input<Pulumi.AzureNative.Sql.DataMaskingState> DataMaskingState { get; set; } = null!;

        /// <summary>
        /// The name of the database.
        /// </summary>
        [Input("databaseName", required: true)]
        public Input<string> DatabaseName { get; set; } = null!;

        /// <summary>
        /// The list of the exempt principals. Specifies the semicolon-separated list of database users for which the data masking policy does not apply. The specified users receive data results without masking for all of the database queries.
        /// </summary>
        [Input("exemptPrincipals")]
        public Input<string>? ExemptPrincipals { get; set; }

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

        public DataMaskingPolicyArgs()
        {
        }
        public static new DataMaskingPolicyArgs Empty => new DataMaskingPolicyArgs();
    }
}
