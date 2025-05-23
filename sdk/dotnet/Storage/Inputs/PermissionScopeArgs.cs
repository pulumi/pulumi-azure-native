// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Storage.Inputs
{

    public sealed class PermissionScopeArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The permissions for the local user. Possible values include: Read (r), Write (w), Delete (d), List (l), Create (c), Modify Ownership (o), and Modify Permissions (p).
        /// </summary>
        [Input("permissions", required: true)]
        public Input<string> Permissions { get; set; } = null!;

        /// <summary>
        /// The name of resource, normally the container name or the file share name, used by the local user.
        /// </summary>
        [Input("resourceName", required: true)]
        public Input<string> ResourceName { get; set; } = null!;

        /// <summary>
        /// The service used by the local user, e.g. blob, file.
        /// </summary>
        [Input("service", required: true)]
        public Input<string> Service { get; set; } = null!;

        public PermissionScopeArgs()
        {
        }
        public static new PermissionScopeArgs Empty => new PermissionScopeArgs();
    }
}
