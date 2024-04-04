import {
    Dashboard,
    DashboardMetric,
    GetDashboardMetricRes,
    ListDashboardMetricsRes,
    ListDashboardsQueryParams,
    ListDashboardsRes,
} from '../model';

export interface DashboardService {
    getMetric(
        dashboardId: DashboardMetric['id'],
        metricId: DashboardMetric['id'],
    ): Promise<GetDashboardMetricRes>;
    list(params: ListDashboardsQueryParams): Promise<ListDashboardsRes>;
    listMetrics(dashboardId: Dashboard['id']): Promise<ListDashboardMetricsRes>;
}
