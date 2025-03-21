// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Web.V20150801Preview.Inputs
{

    /// <summary>
    /// Back end service per ASE
    /// </summary>
    public sealed class HostingEnvironmentServiceDescriptionsArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Host Id
        /// </summary>
        [Input("hostId")]
        public Input<string>? HostId { get; set; }

        /// <summary>
        /// Hosting environment Id
        /// </summary>
        [Input("hostingEnvironmentId")]
        public Input<string>? HostingEnvironmentId { get; set; }

        /// <summary>
        /// service url to use
        /// </summary>
        [Input("serviceUrl")]
        public Input<string>? ServiceUrl { get; set; }

        /// <summary>
        /// When the backend url is in same ASE, for performance reason this flag can be set to true
        ///             If WebApp.DisableHostNames is also set it improves the security by making the back end accessible only 
        ///             via API calls
        ///             Note: calls will fail if this option is used but back end is not on the same ASE
        /// </summary>
        [Input("useInternalRouting")]
        public Input<bool>? UseInternalRouting { get; set; }

        public HostingEnvironmentServiceDescriptionsArgs()
        {
        }
        public static new HostingEnvironmentServiceDescriptionsArgs Empty => new HostingEnvironmentServiceDescriptionsArgs();
    }
}
