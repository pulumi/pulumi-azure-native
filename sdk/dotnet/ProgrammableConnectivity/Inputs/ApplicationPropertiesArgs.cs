// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.ProgrammableConnectivity.Inputs
{

    /// <summary>
    /// Details about the Application that would use the Operator's Network APIs.
    /// </summary>
    public sealed class ApplicationPropertiesArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Description of the application.
        /// </summary>
        [Input("applicationDescription")]
        public Input<string>? ApplicationDescription { get; set; }

        /// <summary>
        /// The category that describes the application.
        /// </summary>
        [Input("applicationType")]
        public Input<string>? ApplicationType { get; set; }

        /// <summary>
        /// Legal name of the organization owning the application.
        /// </summary>
        [Input("legalName")]
        public Input<string>? LegalName { get; set; }

        /// <summary>
        /// Name of the application.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// A description of the organization owning the application.
        /// </summary>
        [Input("organizationDescription")]
        public Input<string>? OrganizationDescription { get; set; }

        /// <summary>
        /// Email address of the Privacy contact or Data Protection officer of the organization.
        /// </summary>
        [Input("privacyContactEmailAddress")]
        public Input<string>? PrivacyContactEmailAddress { get; set; }

        /// <summary>
        /// Unique Tax Number for the user's organization in the country/region the APC Gateway is being purchased.
        /// </summary>
        [Input("taxNumber")]
        public Input<string>? TaxNumber { get; set; }

        public ApplicationPropertiesArgs()
        {
        }
        public static new ApplicationPropertiesArgs Empty => new ApplicationPropertiesArgs();
    }
}
