// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Cdn.Outputs
{

    /// <summary>
    /// Defines the url rewrite action for the delivery rule.
    /// </summary>
    [OutputType]
    public sealed class UrlRewriteActionResponse
    {
        /// <summary>
        /// The name of the action for the delivery rule.
        /// Expected value is 'UrlRewrite'.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Defines the parameters for the action.
        /// </summary>
        public readonly Outputs.UrlRewriteActionParametersResponse Parameters;

        [OutputConstructor]
        private UrlRewriteActionResponse(
            string name,

            Outputs.UrlRewriteActionParametersResponse parameters)
        {
            Name = name;
            Parameters = parameters;
        }
    }
}
