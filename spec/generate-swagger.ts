import { readFile } from 'fs/promises';
import { join } from 'path';

import { ExtendedSpecConfig, generateSpec } from 'tsoa';

const basePath = join('/app');

const pkgJsonPath = join(basePath, 'package.json');
const pkgJsonBuff = await readFile(pkgJsonPath);
const pkgJson = pkgJsonBuff.toString();
const { version }: { version: string } = JSON.parse(pkgJson);

const specBasePath = join(basePath, 'spec');
const controllersPath = join(
    specBasePath,
    'controllers',
    '**',
    '*Controller.ts',
);
const outputDirectory = join(specBasePath, '_generated');

const specOptions: ExtendedSpecConfig = {
    basePath: '/',
    controllerPathGlobs: [controllersPath],
    description: `Find all the details about this specification on the [GitHub repository](https://github.com/c100k/rebootx-on-prem).`,
    entryFile: 'index.ts',
    host: 'localhost:9001/cd5331ba',
    name: 'RebootX On-Prem Specification',
    noImplicitAdditionalProperties: 'throw-on-extras',
    outputDirectory,
    schemes: ['http'],
    securityDefinitions: {
        authorizationHeader: {
            in: 'header',
            name: 'authorization',
            type: 'apiKey',
        },
    },
    specVersion: 3,
    version,
};

await generateSpec(specOptions);
