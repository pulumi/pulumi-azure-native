// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.LoadTestService
{
    /// <summary>
    /// LoadTest profile mapping resource details
    /// 
    /// Uses Azure REST API version 2023-12-01-preview. In version 2.x of the Azure Native provider, it used API version 2023-12-01-preview.
    /// 
    /// Other available API versions: 2024-12-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native loadtestservice [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
    /// </summary>
    [AzureNativeResourceType("azure-native:loadtestservice:LoadTestProfileMapping")]
    public partial class LoadTestProfileMapping : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        [Output("azureApiVersion")]
        public Output<string> AzureApiVersion { get; private set; } = null!;

        /// <summary>
        /// Mapped Azure Load Test resource Id.
        /// </summary>
        [Output("azureLoadTestingResourceId")]
        public Output<string?> AzureLoadTestingResourceId { get; private set; } = null!;

        /// <summary>
        /// The name of the resource
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Mapped source resource Id.
        /// </summary>
        [Output("sourceResourceId")]
        public Output<string?> SourceResourceId { get; private set; } = null!;

        /// <summary>
        /// Azure Resource Manager metadata containing createdBy and modifiedBy information.
        /// </summary>
        [Output("systemData")]
        public Output<Outputs.SystemDataResponse> SystemData { get; private set; } = null!;

        /// <summary>
        /// Mapped Azure Load Test resource test-profile-id.
        /// </summary>
        [Output("testProfileId")]
        public Output<string?> TestProfileId { get; private set; } = null!;

        /// <summary>
        /// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a LoadTestProfileMapping resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public LoadTestProfileMapping(string name, LoadTestProfileMappingArgs args, CustomResourceOptions? options = null)
            : base("azure-native:loadtestservice:LoadTestProfileMapping", name, args ?? new LoadTestProfileMappingArgs(), MakeResourceOptions(options, ""))
        {
        }

        private LoadTestProfileMapping(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azure-native:loadtestservice:LoadTestProfileMapping", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new global::Pulumi.Alias { Type = "azure-native:loadtestservice/v20231201preview:LoadTestProfileMapping" },
                    new global::Pulumi.Alias { Type = "azure-native:loadtestservice/v20241201preview:LoadTestProfileMapping" },
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing LoadTestProfileMapping resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static LoadTestProfileMapping Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new LoadTestProfileMapping(name, id, options);
        }
    }

    public sealed class LoadTestProfileMappingArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Mapped Azure Load Test resource Id.
        /// </summary>
        [Input("azureLoadTestingResourceId")]
        public Input<string>? AzureLoadTestingResourceId { get; set; }

        /// <summary>
        /// Load Test Profile Mapping name
        /// </summary>
        [Input("loadTestProfileMappingName")]
        public Input<string>? LoadTestProfileMappingName { get; set; }

        /// <summary>
        /// The fully qualified Azure Resource manager identifier of the resource.
        /// </summary>
        [Input("resourceUri", required: true)]
        public Input<string> ResourceUri { get; set; } = null!;

        /// <summary>
        /// Mapped source resource Id.
        /// </summary>
        [Input("sourceResourceId")]
        public Input<string>? SourceResourceId { get; set; }

        /// <summary>
        /// Mapped Azure Load Test resource test-profile-id.
        /// </summary>
        [Input("testProfileId")]
        public Input<string>? TestProfileId { get; set; }

        public LoadTestProfileMappingArgs()
        {
        }
        public static new LoadTestProfileMappingArgs Empty => new LoadTestProfileMappingArgs();
    }
}
