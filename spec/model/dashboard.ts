import { ListQueryParams, ListRes } from '../schema';

// NOTE : Voluntarily duplicating ListRunnablesQueryParams because tsoa does not handle well another level of inheritance or union type
// @Queries('queryParams') only support 'refObject' or 'nestedObjectLiteral' types. If you want only one query parameter, please use the '@Query' decorator.
export interface ListDashboardsQueryParams extends ListQueryParams {
    /**
     * Filter on one or multiple properties. It's up to you to implement the filtering that you want. It can be as simple as equality check on one specific field (e.g. `name`). It can also be a partial check (e.g. `ILIKE` pattern) on multiple fields.
     */
    q?: string;
}

export type ListDashboardsRes = ListRes<Dashboard>;

/**
 * A collection of metrics
 */
export interface Dashboard {
    id: string;
    name: string;
}
