// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.CloudHealth.Inputs
{

    /// <summary>
    /// Discovery rule properties
    /// </summary>
    public sealed class DiscoveryRulePropertiesArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Whether to add all recommended signals to the discovered entities.
        /// </summary>
        [Input("addRecommendedSignals", required: true)]
        public InputUnion<string, Pulumi.AzureNative.CloudHealth.DiscoveryRuleRecommendedSignalsBehavior> AddRecommendedSignals { get; set; } = null!;

        /// <summary>
        /// Reference to the name of the authentication setting which is used for querying Azure Resource Graph. The same authentication setting will also be assigned to any discovered entities.
        /// </summary>
        [Input("authenticationSetting", required: true)]
        public Input<string> AuthenticationSetting { get; set; } = null!;

        /// <summary>
        /// Whether to create relationships between the discovered entities based on a set of built-in rules. These relationships cannot be manually deleted.
        /// </summary>
        [Input("discoverRelationships", required: true)]
        public InputUnion<string, Pulumi.AzureNative.CloudHealth.DiscoveryRuleRelationshipDiscoveryBehavior> DiscoverRelationships { get; set; } = null!;

        /// <summary>
        /// Display name
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        /// <summary>
        /// Azure Resource Graph query text in KQL syntax. The query must return at least a column named 'id' which contains the resource ID of the discovered resources.
        /// </summary>
        [Input("resourceGraphQuery", required: true)]
        public Input<string> ResourceGraphQuery { get; set; } = null!;

        public DiscoveryRulePropertiesArgs()
        {
        }
        public static new DiscoveryRulePropertiesArgs Empty => new DiscoveryRulePropertiesArgs();
    }
}
