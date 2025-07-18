// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.ProviderHub.Inputs
{

    /// <summary>
    /// The expedited rollout definition.
    /// </summary>
    public sealed class DefaultRolloutSpecificationExpeditedRolloutArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Indicates whether expedited rollout is enabled/disabled
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        public DefaultRolloutSpecificationExpeditedRolloutArgs()
        {
        }
        public static new DefaultRolloutSpecificationExpeditedRolloutArgs Empty => new DefaultRolloutSpecificationExpeditedRolloutArgs();
    }
}
