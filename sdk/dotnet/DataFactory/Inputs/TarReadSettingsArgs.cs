// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.DataFactory.Inputs
{

    /// <summary>
    /// The Tar compression read settings.
    /// </summary>
    public sealed class TarReadSettingsArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Preserve the compression file name as folder path. Type: boolean (or Expression with resultType boolean).
        /// </summary>
        [Input("preserveCompressionFileNameAsFolder")]
        public Input<object>? PreserveCompressionFileNameAsFolder { get; set; }

        /// <summary>
        /// The Compression setting type.
        /// Expected value is 'TarReadSettings'.
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        public TarReadSettingsArgs()
        {
        }
        public static new TarReadSettingsArgs Empty => new TarReadSettingsArgs();
    }
}
