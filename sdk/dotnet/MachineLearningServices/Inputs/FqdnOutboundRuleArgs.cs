// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.MachineLearningServices.Inputs
{

    /// <summary>
    /// FQDN Outbound Rule for the managed network of a machine learning workspace.
    /// </summary>
    public sealed class FqdnOutboundRuleArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Category of a managed network Outbound Rule of a machine learning workspace.
        /// </summary>
        [Input("category")]
        public InputUnion<string, Pulumi.AzureNative.MachineLearningServices.RuleCategory>? Category { get; set; }

        [Input("destination")]
        public Input<string>? Destination { get; set; }

        /// <summary>
        /// Type of a managed network Outbound Rule of a machine learning workspace.
        /// </summary>
        [Input("status")]
        public InputUnion<string, Pulumi.AzureNative.MachineLearningServices.RuleStatus>? Status { get; set; }

        /// <summary>
        /// Type of a managed network Outbound Rule of a machine learning workspace.
        /// Expected value is 'FQDN'.
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        public FqdnOutboundRuleArgs()
        {
        }
        public static new FqdnOutboundRuleArgs Empty => new FqdnOutboundRuleArgs();
    }
}
