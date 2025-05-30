// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.HybridCompute.Inputs
{

    /// <summary>
    /// Private endpoint which the connection belongs to.
    /// </summary>
    public sealed class PrivateEndpointPropertyArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Resource id of the private endpoint.
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        public PrivateEndpointPropertyArgs()
        {
        }
        public static new PrivateEndpointPropertyArgs Empty => new PrivateEndpointPropertyArgs();
    }
}
