// Copyright 2016-2024, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

var (
	metadata20220901 = `{"portal":"https://portal.azure.com","authentication":{"loginEndpoint":"https://login.microsoftonline.com","audiences":["https://management.core.windows.net/","https://management.azure.com/"],"tenant":"common","identityProvider":"AAD"},"media":"https://rest.media.azure.net","graphAudience":"https://graph.windows.net/","graph":"https://graph.windows.net/","name":"AzureCloud","suffixes":{"azureDataLakeStoreFileSystem":"azuredatalakestore.net","acrLoginServer":"azurecr.io","sqlServerHostname":"database.windows.net","azureDataLakeAnalyticsCatalogAndJob":"azuredatalakeanalytics.net","keyVaultDns":"vault.azure.net","storage":"core.windows.net","azureFrontDoorEndpointSuffix":"azurefd.net","storageSyncEndpointSuffix":"afs.azure.net","mhsmDns":"managedhsm.azure.net","mysqlServerEndpoint":"mysql.database.azure.com","postgresqlServerEndpoint":"postgres.database.azure.com","mariadbServerEndpoint":"mariadb.database.azure.com","synapseAnalytics":"dev.azuresynapse.net","attestationEndpoint":"attest.azure.net"},"batch":"https://batch.core.windows.net/","resourceManager":"https://management.azure.com/","vmImageAliasDoc":"https://raw.githubusercontent.com/Azure/azure-rest-api-specs/master/arm-compute/quickstart-templates/aliases.json","activeDirectoryDataLake":"https://datalake.azure.net/","sqlManagement":"https://management.core.windows.net:8443/","microsoftGraphResourceId":"https://graph.microsoft.com/","appInsightsResourceId":"https://api.applicationinsights.io","appInsightsTelemetryChannelResourceId":"https://dc.applicationinsights.azure.com/v2/track","attestationResourceId":"https://attest.azure.net","synapseAnalyticsResourceId":"https://dev.azuresynapse.net","logAnalyticsResourceId":"https://api.loganalytics.io","ossrDbmsResourceId":"https://ossrdbms-aad.database.windows.net"}`
)

func main() {
	var (
		addr = flag.String("addr", ":8080", "address to listen on")
		cert = flag.String("cert", "", "path to TLS certificate file (PEM format)")
		key  = flag.String("key", "", "path to TLS private key file (PEM format)")
	)
	flag.Parse()

	http.HandleFunc("/metadata/endpoints", func(w http.ResponseWriter, r *http.Request) {
		apiVersion := r.URL.Query().Get("api-version")
		log.Printf("Metadata request received for API Version: %s", apiVersion)
		if apiVersion != "2022-09-01" {
			http.Error(w, "Unsupported API version", http.StatusBadRequest)
			return
		}
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		fmt.Fprint(w, metadata20220901)
	})

	if *cert != "" && *key != "" {
		log.Printf("Local Azure Metadata Service listening with HTTPS on %s", *addr)
		if err := http.ListenAndServeTLS(*addr, *cert, *key, nil); err != nil {
			log.Fatalf("HTTPS server failed: %v", err)
		}
	} else {
		log.Printf("Local Azure Metadata Service listening with HTTP on %s (no TLS cert/key provided)", *addr)
		if err := http.ListenAndServe(*addr, nil); err != nil {
			log.Fatalf("HTTP server failed: %v", err)
		}
	}
}
