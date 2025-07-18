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
    /// Database, Server or Elastic Pool Advisor.
    /// 
    /// Uses Azure REST API version 2023-08-01. In version 2.x of the Azure Native provider, it used API version 2021-11-01.
    /// 
    /// Other available API versions: 2014-04-01, 2015-05-01-preview, 2020-02-02-preview, 2020-08-01-preview, 2020-11-01-preview, 2021-02-01-preview, 2021-05-01-preview, 2021-08-01-preview, 2021-11-01, 2021-11-01-preview, 2022-02-01-preview, 2022-05-01-preview, 2022-08-01-preview, 2022-11-01-preview, 2023-02-01-preview, 2023-05-01-preview, 2023-08-01-preview, 2024-05-01-preview, 2024-11-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native sql [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
    /// </summary>
    [AzureNativeResourceType("azure-native:sql:ServerAdvisor")]
    public partial class ServerAdvisor : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Gets the status of availability of this advisor to customers. Possible values are 'GA', 'PublicPreview', 'LimitedPublicPreview' and 'PrivatePreview'.
        /// </summary>
        [Output("advisorStatus")]
        public Output<string> AdvisorStatus { get; private set; } = null!;

        /// <summary>
        /// Gets the auto-execute status (whether to let the system execute the recommendations) of this advisor. Possible values are 'Enabled' and 'Disabled'
        /// </summary>
        [Output("autoExecuteStatus")]
        public Output<string> AutoExecuteStatus { get; private set; } = null!;

        /// <summary>
        /// Gets the resource from which current value of auto-execute status is inherited. Auto-execute status can be set on (and inherited from) different levels in the resource hierarchy. Possible values are 'Subscription', 'Server', 'ElasticPool', 'Database' and 'Default' (when status is not explicitly set on any level).
        /// </summary>
        [Output("autoExecuteStatusInheritedFrom")]
        public Output<string> AutoExecuteStatusInheritedFrom { get; private set; } = null!;

        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        [Output("azureApiVersion")]
        public Output<string> AzureApiVersion { get; private set; } = null!;

        /// <summary>
        /// Resource kind.
        /// </summary>
        [Output("kind")]
        public Output<string> Kind { get; private set; } = null!;

        /// <summary>
        /// Gets the time when the current resource was analyzed for recommendations by this advisor.
        /// </summary>
        [Output("lastChecked")]
        public Output<string> LastChecked { get; private set; } = null!;

        /// <summary>
        /// Resource location.
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// Resource name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Gets that status of recommendations for this advisor and reason for not having any recommendations. Possible values include, but are not limited to, 'Ok' (Recommendations available),LowActivity (not enough workload to analyze), 'DbSeemsTuned' (Database is doing well), etc.
        /// </summary>
        [Output("recommendationsStatus")]
        public Output<string> RecommendationsStatus { get; private set; } = null!;

        /// <summary>
        /// Gets the recommended actions for this advisor.
        /// </summary>
        [Output("recommendedActions")]
        public Output<ImmutableArray<Outputs.RecommendedActionResponse>> RecommendedActions { get; private set; } = null!;

        /// <summary>
        /// Resource type.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a ServerAdvisor resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ServerAdvisor(string name, ServerAdvisorArgs args, CustomResourceOptions? options = null)
            : base("azure-native:sql:ServerAdvisor", name, args ?? new ServerAdvisorArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ServerAdvisor(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azure-native:sql:ServerAdvisor", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new global::Pulumi.Alias { Type = "azure-native:sql/v20140401:ServerAdvisor" },
                    new global::Pulumi.Alias { Type = "azure-native:sql/v20150501preview:ServerAdvisor" },
                    new global::Pulumi.Alias { Type = "azure-native:sql/v20200202preview:ServerAdvisor" },
                    new global::Pulumi.Alias { Type = "azure-native:sql/v20200801preview:ServerAdvisor" },
                    new global::Pulumi.Alias { Type = "azure-native:sql/v20201101preview:ServerAdvisor" },
                    new global::Pulumi.Alias { Type = "azure-native:sql/v20210201preview:ServerAdvisor" },
                    new global::Pulumi.Alias { Type = "azure-native:sql/v20210501preview:ServerAdvisor" },
                    new global::Pulumi.Alias { Type = "azure-native:sql/v20210801preview:ServerAdvisor" },
                    new global::Pulumi.Alias { Type = "azure-native:sql/v20211101:ServerAdvisor" },
                    new global::Pulumi.Alias { Type = "azure-native:sql/v20211101preview:ServerAdvisor" },
                    new global::Pulumi.Alias { Type = "azure-native:sql/v20220201preview:ServerAdvisor" },
                    new global::Pulumi.Alias { Type = "azure-native:sql/v20220501preview:ServerAdvisor" },
                    new global::Pulumi.Alias { Type = "azure-native:sql/v20220801preview:ServerAdvisor" },
                    new global::Pulumi.Alias { Type = "azure-native:sql/v20221101preview:ServerAdvisor" },
                    new global::Pulumi.Alias { Type = "azure-native:sql/v20230201preview:ServerAdvisor" },
                    new global::Pulumi.Alias { Type = "azure-native:sql/v20230501preview:ServerAdvisor" },
                    new global::Pulumi.Alias { Type = "azure-native:sql/v20230801:ServerAdvisor" },
                    new global::Pulumi.Alias { Type = "azure-native:sql/v20230801preview:ServerAdvisor" },
                    new global::Pulumi.Alias { Type = "azure-native:sql/v20240501preview:ServerAdvisor" },
                    new global::Pulumi.Alias { Type = "azure-native:sql/v20241101preview:ServerAdvisor" },
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing ServerAdvisor resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ServerAdvisor Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new ServerAdvisor(name, id, options);
        }
    }

    public sealed class ServerAdvisorArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the Server Advisor.
        /// </summary>
        [Input("advisorName")]
        public Input<string>? AdvisorName { get; set; }

        /// <summary>
        /// Gets the auto-execute status (whether to let the system execute the recommendations) of this advisor. Possible values are 'Enabled' and 'Disabled'
        /// </summary>
        [Input("autoExecuteStatus", required: true)]
        public Input<Pulumi.AzureNative.Sql.AutoExecuteStatus> AutoExecuteStatus { get; set; } = null!;

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

        public ServerAdvisorArgs()
        {
        }
        public static new ServerAdvisorArgs Empty => new ServerAdvisorArgs();
    }
}
