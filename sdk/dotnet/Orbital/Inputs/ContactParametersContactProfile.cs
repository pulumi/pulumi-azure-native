// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Orbital.Inputs
{

    /// <summary>
    /// The reference to the contact profile resource.
    /// </summary>
    public sealed class ContactParametersContactProfile : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Resource ID.
        /// </summary>
        [Input("id", required: true)]
        public string Id { get; set; } = null!;

        public ContactParametersContactProfile()
        {
        }
        public static new ContactParametersContactProfile Empty => new ContactParametersContactProfile();
    }
}
