// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Web.Inputs
{

    /// <summary>
    /// The configuration settings of the Apple provider.
    /// </summary>
    public sealed class AppleArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// &lt;code&gt;false&lt;/code&gt; if the Apple provider should not be enabled despite the set registration; otherwise, &lt;code&gt;true&lt;/code&gt;.
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        /// <summary>
        /// The configuration settings of the login flow.
        /// </summary>
        [Input("login")]
        public Input<Inputs.LoginScopesArgs>? Login { get; set; }

        /// <summary>
        /// The configuration settings of the Apple registration.
        /// </summary>
        [Input("registration")]
        public Input<Inputs.AppleRegistrationArgs>? Registration { get; set; }

        public AppleArgs()
        {
        }
        public static new AppleArgs Empty => new AppleArgs();
    }
}
