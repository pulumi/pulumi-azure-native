// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Migrate.Inputs
{

    /// <summary>
    /// Defines reference to NSG.
    /// </summary>
    public sealed class NsgReferenceArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Gets the ARM resource ID of the tracked resource being referenced.
        /// </summary>
        [Input("sourceArmResourceId", required: true)]
        public Input<string> SourceArmResourceId { get; set; } = null!;

        public NsgReferenceArgs()
        {
        }
        public static new NsgReferenceArgs Empty => new NsgReferenceArgs();
    }
}
