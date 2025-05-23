// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Compute.Inputs
{

    /// <summary>
    /// Specifies eventGridAndResourceGraph related Scheduled Event related configurations.
    /// </summary>
    public sealed class EventGridAndResourceGraphArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies if event grid and resource graph is enabled for Scheduled event related configurations.
        /// </summary>
        [Input("enable")]
        public Input<bool>? Enable { get; set; }

        public EventGridAndResourceGraphArgs()
        {
        }
        public static new EventGridAndResourceGraphArgs Empty => new EventGridAndResourceGraphArgs();
    }
}
