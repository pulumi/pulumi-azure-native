// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.EventHub.V20230101Preview
{
    /// <summary>
    /// Single item in a List or Get AuthorizationRule operation
    /// </summary>
    [AzureNativeResourceType("azure-native:eventhub/v20230101preview:EventHubAuthorizationRule")]
    public partial class EventHubAuthorizationRule : global::Pulumi.CustomResource
    {
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
        /// The rights associated with the rule.
        /// </summary>
        [Output("rights")]
        public Output<ImmutableArray<string>> Rights { get; private set; } = null!;

        /// <summary>
        /// The system meta data relating to this resource.
        /// </summary>
        [Output("systemData")]
        public Output<Outputs.SystemDataResponse> SystemData { get; private set; } = null!;

        /// <summary>
        /// The type of the resource. E.g. "Microsoft.EventHub/Namespaces" or "Microsoft.EventHub/Namespaces/EventHubs"
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a EventHubAuthorizationRule resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public EventHubAuthorizationRule(string name, EventHubAuthorizationRuleArgs args, CustomResourceOptions? options = null)
            : base("azure-native:eventhub/v20230101preview:EventHubAuthorizationRule", name, args ?? new EventHubAuthorizationRuleArgs(), MakeResourceOptions(options, ""))
        {
        }

        private EventHubAuthorizationRule(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azure-native:eventhub/v20230101preview:EventHubAuthorizationRule", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new global::Pulumi.Alias { Type = "azure-native:eventhub/v20140901:EventHubAuthorizationRule" },
                    new global::Pulumi.Alias { Type = "azure-native:eventhub/v20150801:EventHubAuthorizationRule" },
                    new global::Pulumi.Alias { Type = "azure-native:eventhub/v20170401:EventHubAuthorizationRule" },
                    new global::Pulumi.Alias { Type = "azure-native:eventhub/v20180101preview:EventHubAuthorizationRule" },
                    new global::Pulumi.Alias { Type = "azure-native:eventhub/v20210101preview:EventHubAuthorizationRule" },
                    new global::Pulumi.Alias { Type = "azure-native:eventhub/v20210601preview:EventHubAuthorizationRule" },
                    new global::Pulumi.Alias { Type = "azure-native:eventhub/v20211101:EventHubAuthorizationRule" },
                    new global::Pulumi.Alias { Type = "azure-native:eventhub/v20220101preview:EventHubAuthorizationRule" },
                    new global::Pulumi.Alias { Type = "azure-native:eventhub/v20221001preview:EventHubAuthorizationRule" },
                    new global::Pulumi.Alias { Type = "azure-native:eventhub/v20240101:EventHubAuthorizationRule" },
                    new global::Pulumi.Alias { Type = "azure-native:eventhub/v20240501preview:EventHubAuthorizationRule" },
                    new global::Pulumi.Alias { Type = "azure-native:eventhub:EventHubAuthorizationRule" },
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing EventHubAuthorizationRule resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static EventHubAuthorizationRule Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new EventHubAuthorizationRule(name, id, options);
        }
    }

    public sealed class EventHubAuthorizationRuleArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The authorization rule name.
        /// </summary>
        [Input("authorizationRuleName")]
        public Input<string>? AuthorizationRuleName { get; set; }

        /// <summary>
        /// The Event Hub name
        /// </summary>
        [Input("eventHubName", required: true)]
        public Input<string> EventHubName { get; set; } = null!;

        /// <summary>
        /// The Namespace name
        /// </summary>
        [Input("namespaceName", required: true)]
        public Input<string> NamespaceName { get; set; } = null!;

        /// <summary>
        /// Name of the resource group within the azure subscription.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        [Input("rights", required: true)]
        private InputList<Union<string, Pulumi.AzureNative.EventHub.V20230101Preview.AccessRights>>? _rights;

        /// <summary>
        /// The rights associated with the rule.
        /// </summary>
        public InputList<Union<string, Pulumi.AzureNative.EventHub.V20230101Preview.AccessRights>> Rights
        {
            get => _rights ?? (_rights = new InputList<Union<string, Pulumi.AzureNative.EventHub.V20230101Preview.AccessRights>>());
            set => _rights = value;
        }

        public EventHubAuthorizationRuleArgs()
        {
        }
        public static new EventHubAuthorizationRuleArgs Empty => new EventHubAuthorizationRuleArgs();
    }
}
