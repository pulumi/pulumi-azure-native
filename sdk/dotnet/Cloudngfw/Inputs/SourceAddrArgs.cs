// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Cloudngfw.Inputs
{

    /// <summary>
    /// Address properties
    /// </summary>
    public sealed class SourceAddrArgs : global::Pulumi.ResourceArgs
    {
        [Input("cidrs")]
        private InputList<string>? _cidrs;

        /// <summary>
        /// special value 'any'
        /// </summary>
        public InputList<string> Cidrs
        {
            get => _cidrs ?? (_cidrs = new InputList<string>());
            set => _cidrs = value;
        }

        [Input("countries")]
        private InputList<string>? _countries;

        /// <summary>
        /// list of countries
        /// </summary>
        public InputList<string> Countries
        {
            get => _countries ?? (_countries = new InputList<string>());
            set => _countries = value;
        }

        [Input("feeds")]
        private InputList<string>? _feeds;

        /// <summary>
        /// list of feeds
        /// </summary>
        public InputList<string> Feeds
        {
            get => _feeds ?? (_feeds = new InputList<string>());
            set => _feeds = value;
        }

        [Input("prefixLists")]
        private InputList<string>? _prefixLists;

        /// <summary>
        /// prefix list
        /// </summary>
        public InputList<string> PrefixLists
        {
            get => _prefixLists ?? (_prefixLists = new InputList<string>());
            set => _prefixLists = value;
        }

        public SourceAddrArgs()
        {
        }
        public static new SourceAddrArgs Empty => new SourceAddrArgs();
    }
}
