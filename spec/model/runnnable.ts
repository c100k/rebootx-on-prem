import { ListQueryParams, ListRes } from '../schema';

export interface ListRunnablesQueryParams extends ListQueryParams {
    /**
     * Filter on one or multiple properties. It's up to you to implement the filtering that you want. It can be as simple as equality check on one specific field (e.g. `name`). It can also be a partial check (e.g. `ILIKE` pattern) on multiple fields.
     */
    q?: string;
}

export type ListRunnablesRes = ListRes<Runnable>;

/**
 * Defines the actual status of a runnable.
 * Any intermediary status that you have on your side must be mapped to the `pending` status.
 */
export enum RunnableStatus {
    OFF = 'off',
    ON = 'on',
    PENDING = 'pending',
    UNKNOWN = 'unknown',
}

/**
 * Corresponds to a "context" in which a runnable is.
 * It can be `geo`, defining the geographical zone where the runnable is (e.g. AWS regions code).
 * It can also be `logical`, defining an abstract structure where the runnable is (e.g. GCP project).
 */
export interface RunnableScope {
    label: string;
    value: string;
}

export interface RunnableSSH {
    keyName: string | null;

    /**
     * @isInt
     * @minimum 0
     */
    port: number;

    username: string;
}

export interface RunnableScopes {
    geo: RunnableScope | null;
    logical: RunnableScope | null;
}

/**
 * Corresponds to something that "runs", can be "stopped" and "rebooted".
 * Typical examples are cloud VMs, containers, PaaS applications, etc.
 */
export interface Runnable {
    flavor: string | null;
    fqdn: string | null;
    id: string;
    ipv4: string | null;
    name: string;
    scopes: RunnableScopes;
    ssh: RunnableSSH | null;
    stack: string | null;
    status: RunnableStatus;
}

export interface RunnableOperationRes {
    /**
     * If the process has been triggered on an asynchronous queue and will eventually succeed, you can provide this value here
     */
    jobId: string | null;
}
