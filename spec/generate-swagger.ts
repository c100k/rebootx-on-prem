import { join } from 'path';

import { ExtendedSpecConfig, generateSpec } from 'tsoa';

const basePath = join('/spec');
const controllersPath = join(basePath, 'controllers', '**', '*Controller.ts');
const outputDirectory = join(basePath, '_generated');

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
    version: '0.1.0',
};

await generateSpec(specOptions);
