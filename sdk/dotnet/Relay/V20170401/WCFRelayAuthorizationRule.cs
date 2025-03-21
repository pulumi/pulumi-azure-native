// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Relay.V20170401
{
    /// <summary>
    /// Description of a namespace authorization rule.
    /// </summary>
    [AzureNativeResourceType("azure-native:relay/v20170401:WCFRelayAuthorizationRule")]
    public partial class WCFRelayAuthorizationRule : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Resource name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The rights associated with the rule.
        /// </summary>
        [Output("rights")]
        public Output<ImmutableArray<string>> Rights { get; private set; } = null!;

        /// <summary>
        /// Resource type.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a WCFRelayAuthorizationRule resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public WCFRelayAuthorizationRule(string name, WCFRelayAuthorizationRuleArgs args, CustomResourceOptions? options = null)
            : base("azure-native:relay/v20170401:WCFRelayAuthorizationRule", name, args ?? new WCFRelayAuthorizationRuleArgs(), MakeResourceOptions(options, ""))
        {
        }

        private WCFRelayAuthorizationRule(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azure-native:relay/v20170401:WCFRelayAuthorizationRule", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new global::Pulumi.Alias { Type = "azure-native:relay/v20160701:WCFRelayAuthorizationRule" },
                    new global::Pulumi.Alias { Type = "azure-native:relay/v20211101:WCFRelayAuthorizationRule" },
                    new global::Pulumi.Alias { Type = "azure-native:relay/v20240101:WCFRelayAuthorizationRule" },
                    new global::Pulumi.Alias { Type = "azure-native:relay:WCFRelayAuthorizationRule" },
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing WCFRelayAuthorizationRule resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static WCFRelayAuthorizationRule Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new WCFRelayAuthorizationRule(name, id, options);
        }
    }

    public sealed class WCFRelayAuthorizationRuleArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The authorization rule name.
        /// </summary>
        [Input("authorizationRuleName")]
        public Input<string>? AuthorizationRuleName { get; set; }

        /// <summary>
        /// The namespace name
        /// </summary>
        [Input("namespaceName", required: true)]
        public Input<string> NamespaceName { get; set; } = null!;

        /// <summary>
        /// The relay name.
        /// </summary>
        [Input("relayName", required: true)]
        public Input<string> RelayName { get; set; } = null!;

        /// <summary>
        /// Name of the Resource group within the Azure subscription.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        [Input("rights", required: true)]
        private InputList<Pulumi.AzureNative.Relay.V20170401.AccessRights>? _rights;

        /// <summary>
        /// The rights associated with the rule.
        /// </summary>
        public InputList<Pulumi.AzureNative.Relay.V20170401.AccessRights> Rights
        {
            get => _rights ?? (_rights = new InputList<Pulumi.AzureNative.Relay.V20170401.AccessRights>());
            set => _rights = value;
        }

        public WCFRelayAuthorizationRuleArgs()
        {
        }
        public static new WCFRelayAuthorizationRuleArgs Empty => new WCFRelayAuthorizationRuleArgs();
    }
}
