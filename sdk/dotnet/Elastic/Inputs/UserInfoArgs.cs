// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Elastic.Inputs
{

    /// <summary>
    /// User Information to be passed to partners.
    /// </summary>
    public sealed class UserInfoArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Company information of the user to be passed to partners.
        /// </summary>
        [Input("companyInfo")]
        public Input<Inputs.CompanyInfoArgs>? CompanyInfo { get; set; }

        /// <summary>
        /// Company name of the user
        /// </summary>
        [Input("companyName")]
        public Input<string>? CompanyName { get; set; }

        /// <summary>
        /// Email of the user used by Elastic for contacting them if needed
        /// </summary>
        [Input("emailAddress")]
        public Input<string>? EmailAddress { get; set; }

        /// <summary>
        /// First name of the user
        /// </summary>
        [Input("firstName")]
        public Input<string>? FirstName { get; set; }

        /// <summary>
        /// Last name of the user
        /// </summary>
        [Input("lastName")]
        public Input<string>? LastName { get; set; }

        public UserInfoArgs()
        {
        }
        public static new UserInfoArgs Empty => new UserInfoArgs();
    }
}
