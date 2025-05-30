// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.AwsConnector.Outputs
{

    /// <summary>
    /// Definition of RuleVariables
    /// </summary>
    [OutputType]
    public sealed class RuleVariablesResponse
    {
        /// <summary>
        /// Property ipSets
        /// </summary>
        public readonly ImmutableDictionary<string, Outputs.IPSetResponse>? IpSets;
        /// <summary>
        /// Property portSets
        /// </summary>
        public readonly ImmutableDictionary<string, Outputs.PortSetResponse>? PortSets;

        [OutputConstructor]
        private RuleVariablesResponse(
            ImmutableDictionary<string, Outputs.IPSetResponse>? ipSets,

            ImmutableDictionary<string, Outputs.PortSetResponse>? portSets)
        {
            IpSets = ipSets;
            PortSets = portSets;
        }
    }
}
