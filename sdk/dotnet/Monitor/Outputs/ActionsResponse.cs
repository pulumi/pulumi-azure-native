// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Monitor.Outputs
{

    /// <summary>
    /// Actions to invoke when the alert fires.
    /// </summary>
    [OutputType]
    public sealed class ActionsResponse
    {
        /// <summary>
        /// Action Group resource Ids to invoke when the alert fires.
        /// </summary>
        public readonly ImmutableArray<string> ActionGroups;
        /// <summary>
        /// The properties of an action properties.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? ActionProperties;
        /// <summary>
        /// The properties of an alert payload.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? CustomProperties;

        [OutputConstructor]
        private ActionsResponse(
            ImmutableArray<string> actionGroups,

            ImmutableDictionary<string, string>? actionProperties,

            ImmutableDictionary<string, string>? customProperties)
        {
            ActionGroups = actionGroups;
            ActionProperties = actionProperties;
            CustomProperties = customProperties;
        }
    }
}
