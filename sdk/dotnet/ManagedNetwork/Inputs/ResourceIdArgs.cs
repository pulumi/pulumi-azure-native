// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.ManagedNetwork.Inputs
{

    /// <summary>
    /// Generic pointer to a resource
    /// </summary>
    public sealed class ResourceIdArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Resource Id
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        public ResourceIdArgs()
        {
        }
        public static new ResourceIdArgs Empty => new ResourceIdArgs();
    }
}
