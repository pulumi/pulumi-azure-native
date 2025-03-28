// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.ComponentModel;
using Pulumi;

namespace Pulumi.AzureNative.KubernetesConfiguration.V20211101Preview
{
    /// <summary>
    /// Specify whether to validate the Kubernetes objects referenced in the Kustomization before applying them to the cluster.
    /// </summary>
    [EnumType]
    public readonly struct KustomizationValidationType : IEquatable<KustomizationValidationType>
    {
        private readonly string _value;

        private KustomizationValidationType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static KustomizationValidationType None { get; } = new KustomizationValidationType("none");
        public static KustomizationValidationType Client { get; } = new KustomizationValidationType("client");
        public static KustomizationValidationType Server { get; } = new KustomizationValidationType("server");

        public static bool operator ==(KustomizationValidationType left, KustomizationValidationType right) => left.Equals(right);
        public static bool operator !=(KustomizationValidationType left, KustomizationValidationType right) => !left.Equals(right);

        public static explicit operator string(KustomizationValidationType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is KustomizationValidationType other && Equals(other);
        public bool Equals(KustomizationValidationType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Scope at which the operator will be installed.
    /// </summary>
    [EnumType]
    public readonly struct ScopeType : IEquatable<ScopeType>
    {
        private readonly string _value;

        private ScopeType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static ScopeType Cluster { get; } = new ScopeType("cluster");
        public static ScopeType @Namespace { get; } = new ScopeType("namespace");

        public static bool operator ==(ScopeType left, ScopeType right) => left.Equals(right);
        public static bool operator !=(ScopeType left, ScopeType right) => !left.Equals(right);

        public static explicit operator string(ScopeType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is ScopeType other && Equals(other);
        public bool Equals(ScopeType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Source Kind to pull the configuration data from.
    /// </summary>
    [EnumType]
    public readonly struct SourceKindType : IEquatable<SourceKindType>
    {
        private readonly string _value;

        private SourceKindType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static SourceKindType GitRepository { get; } = new SourceKindType("GitRepository");

        public static bool operator ==(SourceKindType left, SourceKindType right) => left.Equals(right);
        public static bool operator !=(SourceKindType left, SourceKindType right) => !left.Equals(right);

        public static explicit operator string(SourceKindType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is SourceKindType other && Equals(other);
        public bool Equals(SourceKindType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }
}
