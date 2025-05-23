// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Kusto.Inputs
{

    /// <summary>
    /// The list of language extension objects.
    /// </summary>
    public sealed class LanguageExtensionsListArgs : global::Pulumi.ResourceArgs
    {
        [Input("value")]
        private InputList<Inputs.LanguageExtensionArgs>? _value;

        /// <summary>
        /// The list of language extensions.
        /// </summary>
        public InputList<Inputs.LanguageExtensionArgs> Value
        {
            get => _value ?? (_value = new InputList<Inputs.LanguageExtensionArgs>());
            set => _value = value;
        }

        public LanguageExtensionsListArgs()
        {
        }
        public static new LanguageExtensionsListArgs Empty => new LanguageExtensionsListArgs();
    }
}
