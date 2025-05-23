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
    /// Definition of ImageScanStatus
    /// </summary>
    public sealed class ImageScanStatusArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// &lt;p&gt;The description of the image scan status.&lt;/p&gt;
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// &lt;p&gt;The current state of an image scan.&lt;/p&gt;
        /// </summary>
        [Input("status")]
        public Input<Inputs.ScanStatusEnumValueArgs>? Status { get; set; }

        public ImageScanStatusArgs()
        {
        }
        public static new ImageScanStatusArgs Empty => new ImageScanStatusArgs();
    }
}
