import type { ListQueryParams, ListRes } from '../schema/list.js';

export interface ListRunnablesQueryParams extends ListQueryParams {
    /**
     * Filter on one or multiple properties. It's up to you to implement the filtering that you want. It can be as simple as equality check on one specific field (e.g. `name`). It can also be a partial check (e.g. `ILIKE` pattern) on multiple fields.
     */
    q?: string;
}

export type ListRunnablesRes = ListRes<Runnable>;

/**
 * A metric associated to a runnable
 */
export interface RunnableMetric {
    /**
     * Try to keep it short to have a great and more readable display in the app (i.e. "CPU", "RAM", "Proc #")
     */
    label: string | null;

    /**
     * The ratio of the value compared to its maximum.
     * For example, if you have 1024 of RAM and 256 are being used, the ratio should be 256 / 1024 = 0.25
     * @maximum 1.0
     * @minimum 0.0
     */
    ratio: number | null;

    /**
     * If provided, it must be an array of two numbers.
     * They respectively define the limits for "warning" and "danger".
     * To illustrate with CPU usage, these values could be [60, 80].
     * In this case, if value < 60, it will be "success".
     * If value < 80 it will be "warning".
     * Everything else will be "danger".
     * If the first value is greater than the second one (i.e. higher is better), the semantics are reversed.
     * @maxItems 2
     * @minItems 2
     */
    thresholds: number[] | null;

    /**
     * Like for the label, try to keep it short to have a great and more readable display (i.e. "MB", "%", "GB/s")
     */
    unit: string | null;

    /**
     * Format it so it's displayed correctly in the app.
     * If it's a percentage, unlike ratio, put directly the actual value (i.e. 25 and not 0.25)
     */
    value: number | null;
}

/**
 * The status of a runnable
 * Any intermediary status that you have on your side must be mapped to the `pending` status.
 */
export const RunnableStatus = {
    OFF: 'off',
    ON: 'on',
    PENDING: 'pending',
    UNKNOWN: 'unknown',
} as const;

export type RunnableStatus = (typeof RunnableStatus)[keyof typeof RunnableStatus];

/**
 * The context in which a runnable is
 * It can be `geo`, defining the geographical zone where the runnable is (e.g. AWS regions code).
 * It can also be `logical`, defining an abstract structure where the runnable is (e.g. GCP project).
 */
export interface RunnableScope {
    label: string;
    value: string;
}

/**
 * The configuration to define how to SSH into the runnable
 */
export interface RunnableSSH {
    keyName: string | null;

    /**
     * @isInt
     * @minimum 0
     */
    port: number;

    username: string;
}

/**
 * The scopes in which the runnable is
 */
export interface RunnableScopes {
    geo: RunnableScope | null;
    logical: RunnableScope | null;
}

/**
 * Anything that runs, can be stopped and rebooted
 * Typical examples are cloud VMs, containers, PaaS applications, etc.
 */
export interface Runnable {
    flavor: string | null;
    fqdn: string | null;
    id: string;
    ipv4: string | null;
    metrics: RunnableMetric[] | null;
    name: string;
    scopes: RunnableScopes;
    ssh: RunnableSSH | null;
    stack: string | null;
    status: RunnableStatus;
}

/**
 * The response of a reboot, stop operation
 */
export interface RunnableOperationRes {
    /**
     * If the operation is triggered via an asynchronous queue and will eventually succeed, the id can be provided here for information
     */
    jobId: string | null;
}
