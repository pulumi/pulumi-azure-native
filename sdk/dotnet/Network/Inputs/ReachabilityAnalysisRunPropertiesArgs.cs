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
    /// Represents the Reachability Analysis Run properties.
    /// </summary>
    public sealed class ReachabilityAnalysisRunPropertiesArgs : global::Pulumi.ResourceArgs
    {
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Id of the intent resource to run analysis on.
        /// </summary>
        [Input("intentId", required: true)]
        public Input<string> IntentId { get; set; } = null!;

        public ReachabilityAnalysisRunPropertiesArgs()
        {
        }
        public static new ReachabilityAnalysisRunPropertiesArgs Empty => new ReachabilityAnalysisRunPropertiesArgs();
    }
}
