import { ListDashboardsQueryParams, ListDashboardsRes } from '../model';

export interface DashboardService {
    list(params: ListDashboardsQueryParams): Promise<ListDashboardsRes>;
}
