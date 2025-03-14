// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.RedHatOpenShift.V20231122.Inputs
{

    /// <summary>
    /// APIServerProfile represents an API server profile.
    /// </summary>
    public sealed class APIServerProfileArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// API server visibility.
        /// </summary>
        [Input("visibility")]
        public InputUnion<string, Pulumi.AzureNative.RedHatOpenShift.V20231122.Visibility>? Visibility { get; set; }

        public APIServerProfileArgs()
        {
        }
        public static new APIServerProfileArgs Empty => new APIServerProfileArgs();
    }
}
