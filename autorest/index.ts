import { glob } from "fast-glob";
import { writeFile } from "fs/promises";
import * as path from "path";
import { fileURLToPath } from "url";
import { ConfigurationLoader } from "@autorest/configuration";
import {
  ConsoleLogger,
  FilterLogger,
  AutorestSyncLogger,
} from "@autorest/common";
import { createFolderUri, createFileOrFolderUri } from "@azure-tools/uri";

(async () => {
  const logger = new AutorestSyncLogger({
    sinks: [new ConsoleLogger()],
    processors: [new FilterLogger({ level: "error" })],
  });
  const defaultConfigUri = createFileOrFolderUri(
    require.resolve(
      "@autorest/configuration/resources/default-configuration.md"
    )
  );

  const specificationRootDir = path.resolve(
    "..",
    "azure-rest-api-specs",
    "specification"
  );
  const readmeFiles = await glob("*/resource-manager/readme.md", {
    cwd: specificationRootDir,
  });

  const result: string[] = [];
  for (const readmeFile of readmeFiles) {
    const configFolderUri = createFolderUri(
      path.dirname(path.resolve(specificationRootDir, readmeFile))
    );
    const configLoader = new ConfigurationLoader(
      logger,
      defaultConfigUri,
      configFolderUri
    );
    const config = await configLoader.load([], true);
    const inputFiles = config.config.inputFileUris.map((uri) =>
      path.relative(specificationRootDir, fileURLToPath(uri))
    );
    result.push(...inputFiles);
  }
  writeFile(
    "../reports/autorest-input-files.json",
    JSON.stringify(result, null, 2)
  );
  console.log(`found ${result.length} input files`);
})().catch(console.error);
