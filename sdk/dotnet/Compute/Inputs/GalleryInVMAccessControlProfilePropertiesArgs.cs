// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Compute.Inputs
{

    /// <summary>
    /// Describes the properties of a gallery inVMAccessControlProfile.
    /// </summary>
    public sealed class GalleryInVMAccessControlProfilePropertiesArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// This property allows you to specify the Endpoint type for which this profile is defining the access control for. Possible values are: 'WireServer' or 'IMDS'
        /// </summary>
        [Input("applicableHostEndpoint", required: true)]
        public Input<Pulumi.AzureNative.Compute.EndpointTypes> ApplicableHostEndpoint { get; set; } = null!;

        /// <summary>
        /// The description of this gallery inVMAccessControlProfile resources. This property is updatable.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// This property allows you to specify the OS type of the VMs/VMSS for which this profile can be used against. Possible values are: 'Windows' or 'Linux'
        /// </summary>
        [Input("osType", required: true)]
        public Input<Pulumi.AzureNative.Compute.OperatingSystemTypes> OsType { get; set; } = null!;

        public GalleryInVMAccessControlProfilePropertiesArgs()
        {
        }
        public static new GalleryInVMAccessControlProfilePropertiesArgs Empty => new GalleryInVMAccessControlProfilePropertiesArgs();
    }
}
