import {
    Dashboard,
    ListDashboardMetricsRes,
    ListDashboardsQueryParams,
    ListDashboardsRes,
} from '../model';

export interface DashboardService {
    list(params: ListDashboardsQueryParams): Promise<ListDashboardsRes>;
    listMetrics(dashboardId: Dashboard['id']): Promise<ListDashboardMetricsRes>;
}
