// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Cloudngfw.Inputs
{

    /// <summary>
    /// This field is only present if Strata Cloud Manager is managing the policy for this firewall
    /// </summary>
    public sealed class StrataCloudManagerConfigArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Strata Cloud Manager name which is intended to manage the policy for this firewall.
        /// </summary>
        [Input("cloudManagerName", required: true)]
        public Input<string> CloudManagerName { get; set; } = null!;

        public StrataCloudManagerConfigArgs()
        {
        }
        public static new StrataCloudManagerConfigArgs Empty => new StrataCloudManagerConfigArgs();
    }
}
