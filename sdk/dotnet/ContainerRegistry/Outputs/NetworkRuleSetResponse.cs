// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.ContainerRegistry.Outputs
{

    /// <summary>
    /// The network rule set for a container registry.
    /// </summary>
    [OutputType]
    public sealed class NetworkRuleSetResponse
    {
        /// <summary>
        /// The default action of allow or deny when no other rules match.
        /// </summary>
        public readonly string DefaultAction;
        /// <summary>
        /// The IP ACL rules.
        /// </summary>
        public readonly ImmutableArray<Outputs.IPRuleResponse> IpRules;

        [OutputConstructor]
        private NetworkRuleSetResponse(
            string defaultAction,

            ImmutableArray<Outputs.IPRuleResponse> ipRules)
        {
            DefaultAction = defaultAction;
            IpRules = ipRules;
        }
    }
}
