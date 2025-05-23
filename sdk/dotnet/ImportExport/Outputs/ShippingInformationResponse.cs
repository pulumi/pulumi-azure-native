// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.ImportExport.Outputs
{

    /// <summary>
    /// Contains information about the Microsoft datacenter to which the drives should be shipped.
    /// </summary>
    [OutputType]
    public sealed class ShippingInformationResponse
    {
        /// <summary>
        /// Additional shipping information for customer, specific to datacenter to which customer should send their disks.
        /// </summary>
        public readonly string AdditionalInformation;
        /// <summary>
        /// The city name to use when returning the drives.
        /// </summary>
        public readonly string? City;
        /// <summary>
        /// The country or region to use when returning the drives. 
        /// </summary>
        public readonly string? CountryOrRegion;
        /// <summary>
        /// Phone number of the recipient of the returned drives.
        /// </summary>
        public readonly string? Phone;
        /// <summary>
        /// The postal code to use when returning the drives.
        /// </summary>
        public readonly string? PostalCode;
        /// <summary>
        /// The name of the recipient who will receive the hard drives when they are returned. 
        /// </summary>
        public readonly string? RecipientName;
        /// <summary>
        /// The state or province to use when returning the drives.
        /// </summary>
        public readonly string? StateOrProvince;
        /// <summary>
        /// The first line of the street address to use when returning the drives. 
        /// </summary>
        public readonly string? StreetAddress1;
        /// <summary>
        /// The second line of the street address to use when returning the drives. 
        /// </summary>
        public readonly string? StreetAddress2;

        [OutputConstructor]
        private ShippingInformationResponse(
            string additionalInformation,

            string? city,

            string? countryOrRegion,

            string? phone,

            string? postalCode,

            string? recipientName,

            string? stateOrProvince,

            string? streetAddress1,

            string? streetAddress2)
        {
            AdditionalInformation = additionalInformation;
            City = city;
            CountryOrRegion = countryOrRegion;
            Phone = phone;
            PostalCode = postalCode;
            RecipientName = recipientName;
            StateOrProvince = stateOrProvince;
            StreetAddress1 = streetAddress1;
            StreetAddress2 = streetAddress2;
        }
    }
}
