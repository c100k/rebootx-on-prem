import type { ListQueryParams, ListRes } from '../schema/index.js';

/**
 * A collection of metrics
 */
export interface Dashboard {
    id: string;
    metrics: DashboardMetric[] | null;
    name: string;
}

export interface DashboardMetric {
    id: string;

    /**
     * Unlike `RunnableMetric`, this label can be longer to fit your needs
     */
    label: string | null;

    /**
     * Try to keep it short to have a great and more readable display (i.e. "MB", "%", "GB/s")
     */
    unit: string | null;

    /**
     * Format it so it's displayed correctly in the app.
     * If it's a percentage, unlike ratio, put directly the actual value (i.e. 25 and not 0.25)
     */
    value: number | null;
}

export type GetDashboardMetricRes = DashboardMetric;

// NOTE : Voluntarily duplicating ListRunnablesQueryParams because tsoa does not handle well another level of inheritance or union type
// @Queries('queryParams') only support 'refObject' or 'nestedObjectLiteral' types. If you want only one query parameter, please use the '@Query' decorator.
export interface ListDashboardsQueryParams extends ListQueryParams {
    /**
     * Filter on one or multiple properties. It's up to you to implement the filtering that you want. It can be as simple as equality check on one specific field (e.g. `name`). It can also be a partial check (e.g. `ILIKE` pattern) on multiple fields.
     */
    q?: string;
}

export type ListDashboardsRes = ListRes<Dashboard>;

export type ListDashboardMetricsRes = ListRes<DashboardMetric>;
