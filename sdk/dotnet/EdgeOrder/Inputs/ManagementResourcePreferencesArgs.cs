// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.EdgeOrder.Inputs
{

    /// <summary>
    /// Management resource preference to link device.
    /// </summary>
    public sealed class ManagementResourcePreferencesArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Customer preferred Management resource ARM ID.
        /// </summary>
        [Input("preferredManagementResourceId")]
        public Input<string>? PreferredManagementResourceId { get; set; }

        public ManagementResourcePreferencesArgs()
        {
        }
        public static new ManagementResourcePreferencesArgs Empty => new ManagementResourcePreferencesArgs();
    }
}
