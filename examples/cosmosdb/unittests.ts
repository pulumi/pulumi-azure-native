// Copyright 2016-2020, Pulumi Corporation.  All rights reserved.

import * as azurerm from "../../sdk/nodejs";
import * as pulumi from "../../sdk/nodejs/node_modules/@pulumi/pulumi";
import "mocha";
import { distanceBetweenRegions } from "./distance";

let cosmosdbPrimaryLocation = "";
let cosmosdbAccountLocations: string[] = [];

pulumi.runtime.setMocks({
    newResource: function(type: string, name: string, inputs: any): {id: string, state: any} {
        console.log(type);
        if (type === "azurerm:documentdb:DatabaseAccount") {
            cosmosdbAccountLocations = inputs.properties.locations.map((l: any) => l.locationName);
            cosmosdbPrimaryLocation = inputs.location;
        }

        switch (type) {
            default:
                return {
                    id: inputs.name + "_id",
                    state: {
                        ...inputs,
                    },
                };
        }
    },
    call: function(token: string, args: any, provider?: string) {
        return args;
    },
});

// It's important to import the program _after_ the mocks are defined.
import * as sut from "./index";

describe("Cosmos App component", function() {

    describe("Locations", function() {

        it("Cosmos DB account has two locations", function(done) {
            sut.cosmosdbAccount.id.apply(id => {
                if (cosmosdbAccountLocations.length === 2) {
                    done();
                } else {
                    done(new Error(`Wrong number of regions: expected 2 got ${cosmosdbAccountLocations.length}`));
                }
            });
        });

        it("Primary location is set to the first region", function(done) {
            sut.cosmosdbAccount.id.apply(id => {
                if (cosmosdbPrimaryLocation === cosmosdbAccountLocations[0]) {
                    done();
                } else {
                    done(new Error(`Wrong primary regions: expected ${cosmosdbAccountLocations[0]} got ${cosmosdbPrimaryLocation}`));
                }
            });
        });

        it("Max distance between regions is at least 500 km", function(done) {
            sut.cosmosdbAccount.id.apply(id => {
                let max = 0;
                // Iterate throu all pairs of regions and calculate locations.
                for (const regionA of cosmosdbAccountLocations) {
                    for (const regionB of cosmosdbAccountLocations) {
                        const distance = distanceBetweenRegions(regionA, regionB);
                        if (distance > 500) {
                            done();
                            return;
                        }
                        max = Math.max(max, distance);
                    }
                }
                done(new Error(`No regions are at least 500 km apart: max is ${max} km`));
            });
        });
    });
});
