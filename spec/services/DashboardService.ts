import {
    ListDashboardsQueryParams,
    ListDashboardsRes,
} from '../model/index.js';

export interface DashboardService {
    list(params: ListDashboardsQueryParams): Promise<ListDashboardsRes>;
}
