import type {
    ListDashboardsQueryParams,
    ListDashboardsRes,
} from '../model/dashboard.js';

export interface DashboardService {
    list(params: ListDashboardsQueryParams): Promise<ListDashboardsRes>;
}
