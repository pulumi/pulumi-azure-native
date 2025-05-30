// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.MobileNetwork.Inputs
{

    /// <summary>
    /// Reference to a service resource.
    /// </summary>
    public sealed class ServiceResourceIdArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Service resource ID.
        /// </summary>
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        public ServiceResourceIdArgs()
        {
        }
        public static new ServiceResourceIdArgs Empty => new ServiceResourceIdArgs();
    }
}
