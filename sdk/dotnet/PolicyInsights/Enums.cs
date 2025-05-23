// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.ComponentModel;
using Pulumi;

namespace Pulumi.AzureNative.PolicyInsights
{
    /// <summary>
    /// The compliance state that should be set on the resource.
    /// </summary>
    [EnumType]
    public readonly struct ComplianceState : IEquatable<ComplianceState>
    {
        private readonly string _value;

        private ComplianceState(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        /// <summary>
        /// The resource is in compliance with the policy.
        /// </summary>
        public static ComplianceState Compliant { get; } = new ComplianceState("Compliant");
        /// <summary>
        /// The resource is not in compliance with the policy.
        /// </summary>
        public static ComplianceState NonCompliant { get; } = new ComplianceState("NonCompliant");
        /// <summary>
        /// The compliance state of the resource is not known.
        /// </summary>
        public static ComplianceState Unknown { get; } = new ComplianceState("Unknown");

        public static bool operator ==(ComplianceState left, ComplianceState right) => left.Equals(right);
        public static bool operator !=(ComplianceState left, ComplianceState right) => !left.Equals(right);

        public static explicit operator string(ComplianceState value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is ComplianceState other && Equals(other);
        public bool Equals(ComplianceState other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The way resources to remediate are discovered. Defaults to ExistingNonCompliant if not specified.
    /// </summary>
    [EnumType]
    public readonly struct ResourceDiscoveryMode : IEquatable<ResourceDiscoveryMode>
    {
        private readonly string _value;

        private ResourceDiscoveryMode(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        /// <summary>
        /// Remediate resources that are already known to be non-compliant.
        /// </summary>
        public static ResourceDiscoveryMode ExistingNonCompliant { get; } = new ResourceDiscoveryMode("ExistingNonCompliant");
        /// <summary>
        /// Re-evaluate the compliance state of resources and then remediate the resources found to be non-compliant. The resourceIds filter cannot be used in this mode.
        /// </summary>
        public static ResourceDiscoveryMode ReEvaluateCompliance { get; } = new ResourceDiscoveryMode("ReEvaluateCompliance");

        public static bool operator ==(ResourceDiscoveryMode left, ResourceDiscoveryMode right) => left.Equals(right);
        public static bool operator !=(ResourceDiscoveryMode left, ResourceDiscoveryMode right) => !left.Equals(right);

        public static explicit operator string(ResourceDiscoveryMode value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is ResourceDiscoveryMode other && Equals(other);
        public bool Equals(ResourceDiscoveryMode other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }
}
