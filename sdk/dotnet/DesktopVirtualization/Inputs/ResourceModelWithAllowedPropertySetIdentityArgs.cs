// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.DesktopVirtualization.Inputs
{

    public sealed class ResourceModelWithAllowedPropertySetIdentityArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The identity type.
        /// </summary>
        [Input("type")]
        public Input<Pulumi.AzureNative.DesktopVirtualization.ResourceIdentityType>? Type { get; set; }

        public ResourceModelWithAllowedPropertySetIdentityArgs()
        {
        }
        public static new ResourceModelWithAllowedPropertySetIdentityArgs Empty => new ResourceModelWithAllowedPropertySetIdentityArgs();
    }
}
