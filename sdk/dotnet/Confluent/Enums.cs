// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.ComponentModel;
using Pulumi;

namespace Pulumi.AzureNative.Confluent
{
    /// <summary>
    /// Kafka Auth Type
    /// </summary>
    [EnumType]
    public readonly struct AuthType : IEquatable<AuthType>
    {
        private readonly string _value;

        private AuthType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static AuthType SERVICE_ACCOUNT { get; } = new AuthType("SERVICE_ACCOUNT");
        public static AuthType KAFKA_API_KEY { get; } = new AuthType("KAFKA_API_KEY");

        public static bool operator ==(AuthType left, AuthType right) => left.Equals(right);
        public static bool operator !=(AuthType left, AuthType right) => !left.Equals(right);

        public static explicit operator string(AuthType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is AuthType other && Equals(other);
        public bool Equals(AuthType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Connector Class
    /// </summary>
    [EnumType]
    public readonly struct ConnectorClass : IEquatable<ConnectorClass>
    {
        private readonly string _value;

        private ConnectorClass(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static ConnectorClass AZUREBLOBSOURCE { get; } = new ConnectorClass("AZUREBLOBSOURCE");
        public static ConnectorClass AZUREBLOBSINK { get; } = new ConnectorClass("AZUREBLOBSINK");

        public static bool operator ==(ConnectorClass left, ConnectorClass right) => left.Equals(right);
        public static bool operator !=(ConnectorClass left, ConnectorClass right) => !left.Equals(right);

        public static explicit operator string(ConnectorClass value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is ConnectorClass other && Equals(other);
        public bool Equals(ConnectorClass other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The connector service type.
    /// </summary>
    [EnumType]
    public readonly struct ConnectorServiceType : IEquatable<ConnectorServiceType>
    {
        private readonly string _value;

        private ConnectorServiceType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static ConnectorServiceType AzureBlobStorageSinkConnector { get; } = new ConnectorServiceType("AzureBlobStorageSinkConnector");
        public static ConnectorServiceType AzureBlobStorageSourceConnector { get; } = new ConnectorServiceType("AzureBlobStorageSourceConnector");
        public static ConnectorServiceType AzureCosmosDBSinkConnector { get; } = new ConnectorServiceType("AzureCosmosDBSinkConnector");
        public static ConnectorServiceType AzureCosmosDBSourceConnector { get; } = new ConnectorServiceType("AzureCosmosDBSourceConnector");
        public static ConnectorServiceType AzureSynapseAnalyticsSinkConnector { get; } = new ConnectorServiceType("AzureSynapseAnalyticsSinkConnector");

        public static bool operator ==(ConnectorServiceType left, ConnectorServiceType right) => left.Equals(right);
        public static bool operator !=(ConnectorServiceType left, ConnectorServiceType right) => !left.Equals(right);

        public static explicit operator string(ConnectorServiceType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is ConnectorServiceType other && Equals(other);
        public bool Equals(ConnectorServiceType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Connector Status
    /// </summary>
    [EnumType]
    public readonly struct ConnectorStatus : IEquatable<ConnectorStatus>
    {
        private readonly string _value;

        private ConnectorStatus(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static ConnectorStatus PROVISIONING { get; } = new ConnectorStatus("PROVISIONING");
        public static ConnectorStatus RUNNING { get; } = new ConnectorStatus("RUNNING");
        public static ConnectorStatus PAUSED { get; } = new ConnectorStatus("PAUSED");
        public static ConnectorStatus FAILED { get; } = new ConnectorStatus("FAILED");

        public static bool operator ==(ConnectorStatus left, ConnectorStatus right) => left.Equals(right);
        public static bool operator !=(ConnectorStatus left, ConnectorStatus right) => !left.Equals(right);

        public static explicit operator string(ConnectorStatus value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is ConnectorStatus other && Equals(other);
        public bool Equals(ConnectorStatus other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Connector Type
    /// </summary>
    [EnumType]
    public readonly struct ConnectorType : IEquatable<ConnectorType>
    {
        private readonly string _value;

        private ConnectorType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static ConnectorType SINK { get; } = new ConnectorType("SINK");
        public static ConnectorType SOURCE { get; } = new ConnectorType("SOURCE");

        public static bool operator ==(ConnectorType left, ConnectorType right) => left.Equals(right);
        public static bool operator !=(ConnectorType left, ConnectorType right) => !left.Equals(right);

        public static explicit operator string(ConnectorType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is ConnectorType other && Equals(other);
        public bool Equals(ConnectorType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Kafka Output Data Format Type
    /// </summary>
    [EnumType]
    public readonly struct DataFormatType : IEquatable<DataFormatType>
    {
        private readonly string _value;

        private DataFormatType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static DataFormatType AVRO { get; } = new DataFormatType("AVRO");
        public static DataFormatType JSON { get; } = new DataFormatType("JSON");
        public static DataFormatType STRING { get; } = new DataFormatType("STRING");
        public static DataFormatType BYTES { get; } = new DataFormatType("BYTES");
        public static DataFormatType PROTOBUF { get; } = new DataFormatType("PROTOBUF");

        public static bool operator ==(DataFormatType left, DataFormatType right) => left.Equals(right);
        public static bool operator !=(DataFormatType left, DataFormatType right) => !left.Equals(right);

        public static explicit operator string(DataFormatType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is DataFormatType other && Equals(other);
        public bool Equals(DataFormatType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Stream governance configuration
    /// </summary>
    [EnumType]
    public readonly struct Package : IEquatable<Package>
    {
        private readonly string _value;

        private Package(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static Package ESSENTIALS { get; } = new Package("ESSENTIALS");
        public static Package ADVANCED { get; } = new Package("ADVANCED");

        public static bool operator ==(Package left, Package right) => left.Equals(right);
        public static bool operator !=(Package left, Package right) => !left.Equals(right);

        public static explicit operator string(Package value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is Package other && Equals(other);
        public bool Equals(Package other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The partner connector type.
    /// </summary>
    [EnumType]
    public readonly struct PartnerConnectorType : IEquatable<PartnerConnectorType>
    {
        private readonly string _value;

        private PartnerConnectorType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static PartnerConnectorType KafkaAzureBlobStorageSource { get; } = new PartnerConnectorType("KafkaAzureBlobStorageSource");
        public static PartnerConnectorType KafkaAzureBlobStorageSink { get; } = new PartnerConnectorType("KafkaAzureBlobStorageSink");
        public static PartnerConnectorType KafkaAzureCosmosDBSource { get; } = new PartnerConnectorType("KafkaAzureCosmosDBSource");
        public static PartnerConnectorType KafkaAzureCosmosDBSink { get; } = new PartnerConnectorType("KafkaAzureCosmosDBSink");
        public static PartnerConnectorType KafkaAzureSynapseAnalyticsSink { get; } = new PartnerConnectorType("KafkaAzureSynapseAnalyticsSink");

        public static bool operator ==(PartnerConnectorType left, PartnerConnectorType right) => left.Equals(right);
        public static bool operator !=(PartnerConnectorType left, PartnerConnectorType right) => !left.Equals(right);

        public static explicit operator string(PartnerConnectorType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is PartnerConnectorType other && Equals(other);
        public bool Equals(PartnerConnectorType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// SaaS Offer Status
    /// </summary>
    [EnumType]
    public readonly struct SaaSOfferStatus : IEquatable<SaaSOfferStatus>
    {
        private readonly string _value;

        private SaaSOfferStatus(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static SaaSOfferStatus Started { get; } = new SaaSOfferStatus("Started");
        public static SaaSOfferStatus PendingFulfillmentStart { get; } = new SaaSOfferStatus("PendingFulfillmentStart");
        public static SaaSOfferStatus InProgress { get; } = new SaaSOfferStatus("InProgress");
        public static SaaSOfferStatus Subscribed { get; } = new SaaSOfferStatus("Subscribed");
        public static SaaSOfferStatus Suspended { get; } = new SaaSOfferStatus("Suspended");
        public static SaaSOfferStatus Reinstated { get; } = new SaaSOfferStatus("Reinstated");
        public static SaaSOfferStatus Succeeded { get; } = new SaaSOfferStatus("Succeeded");
        public static SaaSOfferStatus Failed { get; } = new SaaSOfferStatus("Failed");
        public static SaaSOfferStatus Unsubscribed { get; } = new SaaSOfferStatus("Unsubscribed");
        public static SaaSOfferStatus Updating { get; } = new SaaSOfferStatus("Updating");

        public static bool operator ==(SaaSOfferStatus left, SaaSOfferStatus right) => left.Equals(right);
        public static bool operator !=(SaaSOfferStatus left, SaaSOfferStatus right) => !left.Equals(right);

        public static explicit operator string(SaaSOfferStatus value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is SaaSOfferStatus other && Equals(other);
        public bool Equals(SaaSOfferStatus other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }
}
