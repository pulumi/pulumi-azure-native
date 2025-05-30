// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.HybridNetwork.Inputs
{

    /// <summary>
    /// Network function definition group properties.
    /// </summary>
    public sealed class NetworkFunctionDefinitionGroupPropertiesFormatArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The network function definition group description.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        public NetworkFunctionDefinitionGroupPropertiesFormatArgs()
        {
        }
        public static new NetworkFunctionDefinitionGroupPropertiesFormatArgs Empty => new NetworkFunctionDefinitionGroupPropertiesFormatArgs();
    }
}
