// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.ComponentModel;
using Pulumi;

namespace Pulumi.AzureNative.Insights.V20200202
{
    /// <summary>
    /// Type of application being monitored.
    /// </summary>
    [EnumType]
    public readonly struct ApplicationType : IEquatable<ApplicationType>
    {
        private readonly string _value;

        private ApplicationType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static ApplicationType Web { get; } = new ApplicationType("web");
        public static ApplicationType Other { get; } = new ApplicationType("other");

        public static bool operator ==(ApplicationType left, ApplicationType right) => left.Equals(right);
        public static bool operator !=(ApplicationType left, ApplicationType right) => !left.Equals(right);

        public static explicit operator string(ApplicationType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is ApplicationType other && Equals(other);
        public bool Equals(ApplicationType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Used by the Application Insights system to determine what kind of flow this component was created by. This is to be set to 'Bluefield' when creating/updating a component via the REST API.
    /// </summary>
    [EnumType]
    public readonly struct FlowType : IEquatable<FlowType>
    {
        private readonly string _value;

        private FlowType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static FlowType Bluefield { get; } = new FlowType("Bluefield");

        public static bool operator ==(FlowType left, FlowType right) => left.Equals(right);
        public static bool operator !=(FlowType left, FlowType right) => !left.Equals(right);

        public static explicit operator string(FlowType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is FlowType other && Equals(other);
        public bool Equals(FlowType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Indicates the flow of the ingestion.
    /// </summary>
    [EnumType]
    public readonly struct IngestionMode : IEquatable<IngestionMode>
    {
        private readonly string _value;

        private IngestionMode(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static IngestionMode ApplicationInsights { get; } = new IngestionMode("ApplicationInsights");
        public static IngestionMode ApplicationInsightsWithDiagnosticSettings { get; } = new IngestionMode("ApplicationInsightsWithDiagnosticSettings");
        public static IngestionMode LogAnalytics { get; } = new IngestionMode("LogAnalytics");

        public static bool operator ==(IngestionMode left, IngestionMode right) => left.Equals(right);
        public static bool operator !=(IngestionMode left, IngestionMode right) => !left.Equals(right);

        public static explicit operator string(IngestionMode value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is IngestionMode other && Equals(other);
        public bool Equals(IngestionMode other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The network access type for accessing Application Insights query.
    /// </summary>
    [EnumType]
    public readonly struct PublicNetworkAccessType : IEquatable<PublicNetworkAccessType>
    {
        private readonly string _value;

        private PublicNetworkAccessType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        /// <summary>
        /// Enables connectivity to Application Insights through public DNS.
        /// </summary>
        public static PublicNetworkAccessType Enabled { get; } = new PublicNetworkAccessType("Enabled");
        /// <summary>
        /// Disables public connectivity to Application Insights through public DNS.
        /// </summary>
        public static PublicNetworkAccessType Disabled { get; } = new PublicNetworkAccessType("Disabled");

        public static bool operator ==(PublicNetworkAccessType left, PublicNetworkAccessType right) => left.Equals(right);
        public static bool operator !=(PublicNetworkAccessType left, PublicNetworkAccessType right) => !left.Equals(right);

        public static explicit operator string(PublicNetworkAccessType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is PublicNetworkAccessType other && Equals(other);
        public bool Equals(PublicNetworkAccessType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Describes what tool created this Application Insights component. Customers using this API should set this to the default 'rest'.
    /// </summary>
    [EnumType]
    public readonly struct RequestSource : IEquatable<RequestSource>
    {
        private readonly string _value;

        private RequestSource(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static RequestSource Rest { get; } = new RequestSource("rest");

        public static bool operator ==(RequestSource left, RequestSource right) => left.Equals(right);
        public static bool operator !=(RequestSource left, RequestSource right) => !left.Equals(right);

        public static explicit operator string(RequestSource value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is RequestSource other && Equals(other);
        public bool Equals(RequestSource other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }
}
