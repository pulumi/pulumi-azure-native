// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.StorageCache.Outputs
{

    /// <summary>
    /// Rule to place restrictions on portions of the cache namespace being presented to clients.
    /// </summary>
    [OutputType]
    public sealed class NfsAccessRuleResponse
    {
        /// <summary>
        /// Access allowed by this rule.
        /// </summary>
        public readonly string Access;
        /// <summary>
        /// GID value that replaces 0 when rootSquash is true. This will use the value of anonymousUID if not provided.
        /// </summary>
        public readonly string? AnonymousGID;
        /// <summary>
        /// UID value that replaces 0 when rootSquash is true. 65534 will be used if not provided.
        /// </summary>
        public readonly string? AnonymousUID;
        /// <summary>
        /// Filter applied to the scope for this rule. The filter's format depends on its scope. 'default' scope matches all clients and has no filter value. 'network' scope takes a filter in CIDR format (for example, 10.99.1.0/24). 'host' takes an IP address or fully qualified domain name as filter. If a client does not match any filter rule and there is no default rule, access is denied.
        /// </summary>
        public readonly string? Filter;
        /// <summary>
        /// Map root accesses to anonymousUID and anonymousGID.
        /// </summary>
        public readonly bool? RootSquash;
        /// <summary>
        /// Scope for this rule. The scope and filter determine which clients match the rule.
        /// </summary>
        public readonly string Scope;
        /// <summary>
        /// For the default policy, allow access to subdirectories under the root export. If this is set to no, clients can only mount the path '/'. If set to yes, clients can mount a deeper path, like '/a/b'.
        /// </summary>
        public readonly bool? SubmountAccess;
        /// <summary>
        /// Allow SUID semantics.
        /// </summary>
        public readonly bool? Suid;

        [OutputConstructor]
        private NfsAccessRuleResponse(
            string access,

            string? anonymousGID,

            string? anonymousUID,

            string? filter,

            bool? rootSquash,

            string scope,

            bool? submountAccess,

            bool? suid)
        {
            Access = access;
            AnonymousGID = anonymousGID;
            AnonymousUID = anonymousUID;
            Filter = filter;
            RootSquash = rootSquash;
            Scope = scope;
            SubmountAccess = submountAccess;
            Suid = suid;
        }
    }
}
