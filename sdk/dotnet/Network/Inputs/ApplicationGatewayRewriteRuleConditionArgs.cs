// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Network.Inputs
{

    /// <summary>
    /// Set of conditions in the Rewrite Rule in Application Gateway.
    /// </summary>
    public sealed class ApplicationGatewayRewriteRuleConditionArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Setting this parameter to truth value with force the pattern to do a case in-sensitive comparison.
        /// </summary>
        [Input("ignoreCase")]
        public Input<bool>? IgnoreCase { get; set; }

        /// <summary>
        /// Setting this value as truth will force to check the negation of the condition given by the user.
        /// </summary>
        [Input("negate")]
        public Input<bool>? Negate { get; set; }

        /// <summary>
        /// The pattern, either fixed string or regular expression, that evaluates the truthfulness of the condition.
        /// </summary>
        [Input("pattern")]
        public Input<string>? Pattern { get; set; }

        /// <summary>
        /// The condition parameter of the RewriteRuleCondition.
        /// </summary>
        [Input("variable")]
        public Input<string>? Variable { get; set; }

        public ApplicationGatewayRewriteRuleConditionArgs()
        {
        }
        public static new ApplicationGatewayRewriteRuleConditionArgs Empty => new ApplicationGatewayRewriteRuleConditionArgs();
    }
}
