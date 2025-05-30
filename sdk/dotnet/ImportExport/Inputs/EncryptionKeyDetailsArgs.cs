// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.ImportExport.Inputs
{

    /// <summary>
    /// Specifies the encryption key properties
    /// </summary>
    public sealed class EncryptionKeyDetailsArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The type of kek encryption key
        /// </summary>
        [Input("kekType")]
        public InputUnion<string, Pulumi.AzureNative.ImportExport.EncryptionKekType>? KekType { get; set; }

        /// <summary>
        /// Specifies the url for kek encryption key. 
        /// </summary>
        [Input("kekUrl")]
        public Input<string>? KekUrl { get; set; }

        /// <summary>
        /// Specifies the keyvault resource id for kek encryption key. 
        /// </summary>
        [Input("kekVaultResourceID")]
        public Input<string>? KekVaultResourceID { get; set; }

        public EncryptionKeyDetailsArgs()
        {
            KekType = "MicrosoftManaged";
        }
        public static new EncryptionKeyDetailsArgs Empty => new EncryptionKeyDetailsArgs();
    }
}
