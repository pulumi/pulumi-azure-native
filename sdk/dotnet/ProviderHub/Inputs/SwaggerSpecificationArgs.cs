// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.ProviderHub.Inputs
{

    public sealed class SwaggerSpecificationArgs : global::Pulumi.ResourceArgs
    {
        [Input("apiVersions")]
        private InputList<string>? _apiVersions;
        public InputList<string> ApiVersions
        {
            get => _apiVersions ?? (_apiVersions = new InputList<string>());
            set => _apiVersions = value;
        }

        [Input("swaggerSpecFolderUri")]
        public Input<string>? SwaggerSpecFolderUri { get; set; }

        public SwaggerSpecificationArgs()
        {
        }
        public static new SwaggerSpecificationArgs Empty => new SwaggerSpecificationArgs();
    }
}
