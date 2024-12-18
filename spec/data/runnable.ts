import {
    type ListRunnablesRes,
    type RunnableOperationRes,
    RunnableStatus,
} from '../model/runnable.js';

export const RUNNABLE_OP_ASYNC_DESC: string =
    'The operation has been executed asynchronously and will eventually succeed';
export const RUNNABLE_OP_ASYNC_RES: RunnableOperationRes = {
    jobId: 'e22df54c-09b2-47cc-be7f-128b15e609c4',
};

export const RUNNABLE_OP_SYNC_DESC: string =
    'The operation has been successfully executed synchronously';
export const RUNNABLE_OP_SYNC_RES: RunnableOperationRes = {
    jobId: null,
};

export const RUNNABLES: ListRunnablesRes = {
    items: [
        {
            flavor: 'medium',
            fqdn: 'server01.mycompany.com',
            id: '123',
            ipv4: '192.168.0.26',
            metrics: [
                {
                    label: 'CPU',
                    ratio: null,
                    thresholds: [65, 85],
                    unit: '%',
                    value: 28,
                },
                {
                    label: 'RAM',
                    ratio: 0.125,
                    thresholds: [3000, 3800],
                    unit: 'MB',
                    value: 512,
                },
            ],
            name: 'server01',
            scopes: {
                geo: {
                    label: 'Paris 01',
                    value: 'par-01',
                },
                logical: {
                    label: 'Project 1',
                    value: 'project-1',
                },
            },
            ssh: {
                keyName: 'keypair-01',
                port: 22,
                username: 'admin',
            },
            stack: 'nodejs',
            status: RunnableStatus.OFF,
        },
        {
            flavor: 'medium',
            fqdn: 'server02.mycompany.com',
            id: '456',
            ipv4: '192.168.0.27',
            metrics: [
                {
                    label: 'CPU',
                    ratio: null,
                    thresholds: [65, 85],
                    unit: '%',
                    value: 82,
                },
                {
                    label: 'RAM',
                    ratio: 0.25,
                    thresholds: [3000, 3800],
                    unit: 'GB',
                    value: 1,
                },
            ],
            name: 'server02',
            scopes: {
                geo: {
                    label: 'Paris 01',
                    value: 'par-01',
                },
                logical: {
                    label: 'Project 1',
                    value: 'project-1',
                },
            },
            ssh: {
                keyName: 'keypair-01',
                port: 22,
                username: 'admin',
            },
            stack: 'go',
            status: RunnableStatus.OFF,
        },
    ],
    total: 2,
};
