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
    /// The configuration settings of the Azure Active Directory allowed principals.
    /// </summary>
    public sealed class AllowedPrincipalsArgs : global::Pulumi.ResourceArgs
    {
        [Input("groups")]
        private InputList<string>? _groups;

        /// <summary>
        /// The list of the allowed groups.
        /// </summary>
        public InputList<string> Groups
        {
            get => _groups ?? (_groups = new InputList<string>());
            set => _groups = value;
        }

        [Input("identities")]
        private InputList<string>? _identities;

        /// <summary>
        /// The list of the allowed identities.
        /// </summary>
        public InputList<string> Identities
        {
            get => _identities ?? (_identities = new InputList<string>());
            set => _identities = value;
        }

        public AllowedPrincipalsArgs()
        {
        }
        public static new AllowedPrincipalsArgs Empty => new AllowedPrincipalsArgs();
    }
}
