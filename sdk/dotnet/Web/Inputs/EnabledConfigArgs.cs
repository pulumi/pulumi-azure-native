// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Web.Inputs
{

    /// <summary>
    /// Enabled configuration.
    /// </summary>
    public sealed class EnabledConfigArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// True if configuration is enabled, false if it is disabled and null if configuration is not set.
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        public EnabledConfigArgs()
        {
        }
        public static new EnabledConfigArgs Empty => new EnabledConfigArgs();
    }
}
