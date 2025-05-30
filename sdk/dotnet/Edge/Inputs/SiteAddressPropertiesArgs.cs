// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Edge.Inputs
{

    /// <summary>
    /// Site address properties
    /// </summary>
    public sealed class SiteAddressPropertiesArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// City of the address
        /// </summary>
        [Input("city")]
        public Input<string>? City { get; set; }

        /// <summary>
        /// Country of the address
        /// </summary>
        [Input("country")]
        public Input<string>? Country { get; set; }

        /// <summary>
        /// Postal or ZIP code of the address
        /// </summary>
        [Input("postalCode")]
        public Input<string>? PostalCode { get; set; }

        /// <summary>
        /// State or province of the address
        /// </summary>
        [Input("stateOrProvince")]
        public Input<string>? StateOrProvince { get; set; }

        /// <summary>
        /// First line of the street address
        /// </summary>
        [Input("streetAddress1")]
        public Input<string>? StreetAddress1 { get; set; }

        /// <summary>
        /// Second line of the street address
        /// </summary>
        [Input("streetAddress2")]
        public Input<string>? StreetAddress2 { get; set; }

        public SiteAddressPropertiesArgs()
        {
        }
        public static new SiteAddressPropertiesArgs Empty => new SiteAddressPropertiesArgs();
    }
}
