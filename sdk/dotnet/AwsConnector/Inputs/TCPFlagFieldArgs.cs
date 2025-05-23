// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.AwsConnector.Inputs
{

    /// <summary>
    /// Definition of TCPFlagField
    /// </summary>
    public sealed class TCPFlagFieldArgs : global::Pulumi.ResourceArgs
    {
        [Input("flags")]
        private InputList<string>? _flags;

        /// <summary>
        /// Property flags
        /// </summary>
        public InputList<string> Flags
        {
            get => _flags ?? (_flags = new InputList<string>());
            set => _flags = value;
        }

        [Input("masks")]
        private InputList<string>? _masks;

        /// <summary>
        /// Property masks
        /// </summary>
        public InputList<string> Masks
        {
            get => _masks ?? (_masks = new InputList<string>());
            set => _masks = value;
        }

        public TCPFlagFieldArgs()
        {
        }
        public static new TCPFlagFieldArgs Empty => new TCPFlagFieldArgs();
    }
}
