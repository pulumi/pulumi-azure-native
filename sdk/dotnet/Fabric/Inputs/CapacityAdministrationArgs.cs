// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Fabric.Inputs
{

    /// <summary>
    /// The administration properties of the Fabric capacity resource
    /// </summary>
    public sealed class CapacityAdministrationArgs : global::Pulumi.ResourceArgs
    {
        [Input("members", required: true)]
        private InputList<string>? _members;

        /// <summary>
        /// An array of administrator user identities.
        /// </summary>
        public InputList<string> Members
        {
            get => _members ?? (_members = new InputList<string>());
            set => _members = value;
        }

        public CapacityAdministrationArgs()
        {
        }
        public static new CapacityAdministrationArgs Empty => new CapacityAdministrationArgs();
    }
}
