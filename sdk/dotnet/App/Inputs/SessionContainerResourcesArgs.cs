// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.App.Inputs
{

    /// <summary>
    /// Container resource requirements for sessions of the session pool.
    /// </summary>
    public sealed class SessionContainerResourcesArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Required CPU in cores, e.g. 0.5
        /// </summary>
        [Input("cpu")]
        public Input<double>? Cpu { get; set; }

        /// <summary>
        /// Required memory, e.g. "250Mb"
        /// </summary>
        [Input("memory")]
        public Input<string>? Memory { get; set; }

        public SessionContainerResourcesArgs()
        {
        }
        public static new SessionContainerResourcesArgs Empty => new SessionContainerResourcesArgs();
    }
}
