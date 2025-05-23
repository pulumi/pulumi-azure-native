// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.VoiceServices.Inputs
{

    /// <summary>
    /// Details of DNS Domains delegated to the Communications Gateway.
    /// </summary>
    public sealed class DnsDelegationsPropertiesArgs : global::Pulumi.ResourceArgs
    {
        [Input("delegations")]
        private InputList<Inputs.DnsDelegationPropertiesArgs>? _delegations;

        /// <summary>
        /// DNS Domains to delegate for the creation of DNS Zones by the Azure Communications Gateway
        /// </summary>
        public InputList<Inputs.DnsDelegationPropertiesArgs> Delegations
        {
            get => _delegations ?? (_delegations = new InputList<Inputs.DnsDelegationPropertiesArgs>());
            set => _delegations = value;
        }

        public DnsDelegationsPropertiesArgs()
        {
        }
        public static new DnsDelegationsPropertiesArgs Empty => new DnsDelegationsPropertiesArgs();
    }
}
