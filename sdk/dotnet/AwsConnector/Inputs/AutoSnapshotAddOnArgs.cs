// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.AwsConnector.Inputs
{

    /// <summary>
    /// Definition of AutoSnapshotAddOn
    /// </summary>
    public sealed class AutoSnapshotAddOnArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The daily time when an automatic snapshot will be created.
        /// </summary>
        [Input("snapshotTimeOfDay")]
        public Input<string>? SnapshotTimeOfDay { get; set; }

        public AutoSnapshotAddOnArgs()
        {
        }
        public static new AutoSnapshotAddOnArgs Empty => new AutoSnapshotAddOnArgs();
    }
}
