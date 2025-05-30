// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.VirtualMachineImages.Inputs
{

    /// <summary>
    /// Describes an image source from [Azure Gallery Images](https://docs.microsoft.com/en-us/rest/api/compute/virtualmachineimages).
    /// </summary>
    public sealed class ImageTemplatePlatformImageSourceArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Image offer from the [Azure Gallery Images](https://docs.microsoft.com/en-us/rest/api/compute/virtualmachineimages).
        /// </summary>
        [Input("offer")]
        public Input<string>? Offer { get; set; }

        /// <summary>
        /// Optional configuration of purchase plan for platform image.
        /// </summary>
        [Input("planInfo")]
        public Input<Inputs.PlatformImagePurchasePlanArgs>? PlanInfo { get; set; }

        /// <summary>
        /// Image Publisher in [Azure Gallery Images](https://docs.microsoft.com/en-us/rest/api/compute/virtualmachineimages).
        /// </summary>
        [Input("publisher")]
        public Input<string>? Publisher { get; set; }

        /// <summary>
        /// Image sku from the [Azure Gallery Images](https://docs.microsoft.com/en-us/rest/api/compute/virtualmachineimages).
        /// </summary>
        [Input("sku")]
        public Input<string>? Sku { get; set; }

        /// <summary>
        /// Specifies the type of source image you want to start with.
        /// Expected value is 'PlatformImage'.
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        /// <summary>
        /// Image version from the [Azure Gallery Images](https://docs.microsoft.com/en-us/rest/api/compute/virtualmachineimages). If 'latest' is specified here, the version is evaluated when the image build takes place, not when the template is submitted.
        /// </summary>
        [Input("version")]
        public Input<string>? Version { get; set; }

        public ImageTemplatePlatformImageSourceArgs()
        {
        }
        public static new ImageTemplatePlatformImageSourceArgs Empty => new ImageTemplatePlatformImageSourceArgs();
    }
}
