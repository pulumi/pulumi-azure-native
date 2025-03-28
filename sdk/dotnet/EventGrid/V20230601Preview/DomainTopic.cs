// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.EventGrid.V20230601Preview
{
    /// <summary>
    /// Domain Topic.
    /// </summary>
    [AzureNativeResourceType("azure-native:eventgrid/v20230601preview:DomainTopic")]
    public partial class DomainTopic : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Name of the resource.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Provisioning state of the domain topic.
        /// </summary>
        [Output("provisioningState")]
        public Output<string> ProvisioningState { get; private set; } = null!;

        /// <summary>
        /// The system metadata relating to Domain Topic resource.
        /// </summary>
        [Output("systemData")]
        public Output<Outputs.SystemDataResponse> SystemData { get; private set; } = null!;

        /// <summary>
        /// Type of the resource.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a DomainTopic resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public DomainTopic(string name, DomainTopicArgs args, CustomResourceOptions? options = null)
            : base("azure-native:eventgrid/v20230601preview:DomainTopic", name, args ?? new DomainTopicArgs(), MakeResourceOptions(options, ""))
        {
        }

        private DomainTopic(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azure-native:eventgrid/v20230601preview:DomainTopic", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new global::Pulumi.Alias { Type = "azure-native:eventgrid/v20190201preview:DomainTopic" },
                    new global::Pulumi.Alias { Type = "azure-native:eventgrid/v20190601:DomainTopic" },
                    new global::Pulumi.Alias { Type = "azure-native:eventgrid/v20200101preview:DomainTopic" },
                    new global::Pulumi.Alias { Type = "azure-native:eventgrid/v20200401preview:DomainTopic" },
                    new global::Pulumi.Alias { Type = "azure-native:eventgrid/v20200601:DomainTopic" },
                    new global::Pulumi.Alias { Type = "azure-native:eventgrid/v20201015preview:DomainTopic" },
                    new global::Pulumi.Alias { Type = "azure-native:eventgrid/v20210601preview:DomainTopic" },
                    new global::Pulumi.Alias { Type = "azure-native:eventgrid/v20211015preview:DomainTopic" },
                    new global::Pulumi.Alias { Type = "azure-native:eventgrid/v20211201:DomainTopic" },
                    new global::Pulumi.Alias { Type = "azure-native:eventgrid/v20220615:DomainTopic" },
                    new global::Pulumi.Alias { Type = "azure-native:eventgrid/v20231215preview:DomainTopic" },
                    new global::Pulumi.Alias { Type = "azure-native:eventgrid/v20240601preview:DomainTopic" },
                    new global::Pulumi.Alias { Type = "azure-native:eventgrid/v20241215preview:DomainTopic" },
                    new global::Pulumi.Alias { Type = "azure-native:eventgrid/v20250215:DomainTopic" },
                    new global::Pulumi.Alias { Type = "azure-native:eventgrid:DomainTopic" },
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing DomainTopic resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static DomainTopic Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new DomainTopic(name, id, options);
        }
    }

    public sealed class DomainTopicArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Name of the domain.
        /// </summary>
        [Input("domainName", required: true)]
        public Input<string> DomainName { get; set; } = null!;

        /// <summary>
        /// Name of the domain topic.
        /// </summary>
        [Input("domainTopicName")]
        public Input<string>? DomainTopicName { get; set; }

        /// <summary>
        /// The name of the resource group within the user's subscription.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        public DomainTopicArgs()
        {
        }
        public static new DomainTopicArgs Empty => new DomainTopicArgs();
    }
}
