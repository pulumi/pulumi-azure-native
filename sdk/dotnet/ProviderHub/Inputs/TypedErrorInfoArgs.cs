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
    /// Error information.
    /// </summary>
    public sealed class TypedErrorInfoArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The type of the error.
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        public TypedErrorInfoArgs()
        {
        }
        public static new TypedErrorInfoArgs Empty => new TypedErrorInfoArgs();
    }
}
