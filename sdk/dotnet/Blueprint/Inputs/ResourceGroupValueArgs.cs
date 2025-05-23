// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Blueprint.Inputs
{

    /// <summary>
    /// Represents an Azure resource group.
    /// </summary>
    public sealed class ResourceGroupValueArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Location of the resource group.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// Name of the resource group.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public ResourceGroupValueArgs()
        {
        }
        public static new ResourceGroupValueArgs Empty => new ResourceGroupValueArgs();
    }
}
