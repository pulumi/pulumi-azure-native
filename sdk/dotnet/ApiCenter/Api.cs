// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.ApiCenter
{
    /// <summary>
    /// API entity.
    /// 
    /// Uses Azure REST API version 2024-03-15-preview. In version 2.x of the Azure Native provider, it used API version 2024-03-01.
    /// 
    /// Other available API versions: 2024-03-01, 2024-06-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native apicenter [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
    /// </summary>
    [AzureNativeResourceType("azure-native:apicenter:Api")]
    public partial class Api : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        [Output("azureApiVersion")]
        public Output<string> AzureApiVersion { get; private set; } = null!;

        /// <summary>
        /// The set of contacts
        /// </summary>
        [Output("contacts")]
        public Output<ImmutableArray<Outputs.ContactResponse>> Contacts { get; private set; } = null!;

        /// <summary>
        /// The custom metadata defined for API catalog entities.
        /// </summary>
        [Output("customProperties")]
        public Output<object?> CustomProperties { get; private set; } = null!;

        /// <summary>
        /// Description of the API.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The set of external documentation
        /// </summary>
        [Output("externalDocumentation")]
        public Output<ImmutableArray<Outputs.ExternalDocumentationResponse>> ExternalDocumentation { get; private set; } = null!;

        /// <summary>
        /// Kind of API. For example, REST or GraphQL.
        /// </summary>
        [Output("kind")]
        public Output<string> Kind { get; private set; } = null!;

        /// <summary>
        /// The license information for the API.
        /// </summary>
        [Output("license")]
        public Output<Outputs.LicenseResponse?> License { get; private set; } = null!;

        /// <summary>
        /// Current lifecycle stage of the API.
        /// </summary>
        [Output("lifecycleStage")]
        public Output<string> LifecycleStage { get; private set; } = null!;

        /// <summary>
        /// The name of the resource
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Short description of the API.
        /// </summary>
        [Output("summary")]
        public Output<string?> Summary { get; private set; } = null!;

        /// <summary>
        /// Azure Resource Manager metadata containing createdBy and modifiedBy information.
        /// </summary>
        [Output("systemData")]
        public Output<Outputs.SystemDataResponse> SystemData { get; private set; } = null!;

        /// <summary>
        /// Terms of service for the API.
        /// </summary>
        [Output("termsOfService")]
        public Output<Outputs.TermsOfServiceResponse?> TermsOfService { get; private set; } = null!;

        /// <summary>
        /// API title.
        /// </summary>
        [Output("title")]
        public Output<string> Title { get; private set; } = null!;

        /// <summary>
        /// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a Api resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Api(string name, ApiArgs args, CustomResourceOptions? options = null)
            : base("azure-native:apicenter:Api", name, args ?? new ApiArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Api(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azure-native:apicenter:Api", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new global::Pulumi.Alias { Type = "azure-native:apicenter/v20240301:Api" },
                    new global::Pulumi.Alias { Type = "azure-native:apicenter/v20240315preview:Api" },
                    new global::Pulumi.Alias { Type = "azure-native:apicenter/v20240601preview:Api" },
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Api resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Api Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Api(name, id, options);
        }
    }

    public sealed class ApiArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the API.
        /// </summary>
        [Input("apiName")]
        public Input<string>? ApiName { get; set; }

        [Input("contacts")]
        private InputList<Inputs.ContactArgs>? _contacts;

        /// <summary>
        /// The set of contacts
        /// </summary>
        public InputList<Inputs.ContactArgs> Contacts
        {
            get => _contacts ?? (_contacts = new InputList<Inputs.ContactArgs>());
            set => _contacts = value;
        }

        /// <summary>
        /// The custom metadata defined for API catalog entities.
        /// </summary>
        [Input("customProperties")]
        public Input<object>? CustomProperties { get; set; }

        /// <summary>
        /// Description of the API.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("externalDocumentation")]
        private InputList<Inputs.ExternalDocumentationArgs>? _externalDocumentation;

        /// <summary>
        /// The set of external documentation
        /// </summary>
        public InputList<Inputs.ExternalDocumentationArgs> ExternalDocumentation
        {
            get => _externalDocumentation ?? (_externalDocumentation = new InputList<Inputs.ExternalDocumentationArgs>());
            set => _externalDocumentation = value;
        }

        /// <summary>
        /// Kind of API. For example, REST or GraphQL.
        /// </summary>
        [Input("kind", required: true)]
        public InputUnion<string, Pulumi.AzureNative.ApiCenter.ApiKind> Kind { get; set; } = null!;

        /// <summary>
        /// The license information for the API.
        /// </summary>
        [Input("license")]
        public Input<Inputs.LicenseArgs>? License { get; set; }

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of Azure API Center service.
        /// </summary>
        [Input("serviceName", required: true)]
        public Input<string> ServiceName { get; set; } = null!;

        /// <summary>
        /// Short description of the API.
        /// </summary>
        [Input("summary")]
        public Input<string>? Summary { get; set; }

        /// <summary>
        /// Terms of service for the API.
        /// </summary>
        [Input("termsOfService")]
        public Input<Inputs.TermsOfServiceArgs>? TermsOfService { get; set; }

        /// <summary>
        /// API title.
        /// </summary>
        [Input("title", required: true)]
        public Input<string> Title { get; set; } = null!;

        /// <summary>
        /// The name of the workspace.
        /// </summary>
        [Input("workspaceName", required: true)]
        public Input<string> WorkspaceName { get; set; } = null!;

        public ApiArgs()
        {
        }
        public static new ApiArgs Empty => new ApiArgs();
    }
}
