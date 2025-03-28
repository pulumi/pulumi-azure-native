// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.ComponentModel;
using Pulumi;

namespace Pulumi.AzureNative.SecurityInsights.V20210401
{
    /// <summary>
    /// The source of the watchlist
    /// </summary>
    [EnumType]
    public readonly struct Source : IEquatable<Source>
    {
        private readonly string _value;

        private Source(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static Source Local_file { get; } = new Source("Local file");
        public static Source Remote_storage { get; } = new Source("Remote storage");

        public static bool operator ==(Source left, Source right) => left.Equals(right);
        public static bool operator !=(Source left, Source right) => !left.Equals(right);

        public static explicit operator string(Source value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is Source other && Equals(other);
        public bool Equals(Source other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The kind of the entity.
    /// </summary>
    [EnumType]
    public readonly struct ThreatIntelligenceResourceInnerKind : IEquatable<ThreatIntelligenceResourceInnerKind>
    {
        private readonly string _value;

        private ThreatIntelligenceResourceInnerKind(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        /// <summary>
        /// Entity represents threat intelligence indicator in the system.
        /// </summary>
        public static ThreatIntelligenceResourceInnerKind Indicator { get; } = new ThreatIntelligenceResourceInnerKind("indicator");

        public static bool operator ==(ThreatIntelligenceResourceInnerKind left, ThreatIntelligenceResourceInnerKind right) => left.Equals(right);
        public static bool operator !=(ThreatIntelligenceResourceInnerKind left, ThreatIntelligenceResourceInnerKind right) => !left.Equals(right);

        public static explicit operator string(ThreatIntelligenceResourceInnerKind value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is ThreatIntelligenceResourceInnerKind other && Equals(other);
        public bool Equals(ThreatIntelligenceResourceInnerKind other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }
}
