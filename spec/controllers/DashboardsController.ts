import {
    Controller,
    Example,
    Get,
    OperationId,
    Path,
    Produces,
    Queries,
    Response,
    Route,
    Security,
    SuccessResponse,
    Tags,
} from 'tsoa';

import {
    DASHBOARDS,
    DASHBOARD_METRICS_FOR_BUSINESS,
    DASHBOARD_METRICS_FOR_BUSINESS_1,
    ERR_401,
    ERR_403,
    ERR_404,
} from '../data';
import {
    Dashboard,
    DashboardMetric,
    GetDashboardMetricRes,
    ListDashboardMetricsRes,
    ListDashboardsQueryParams,
    ListDashboardsRes,
} from '../model';
import { ErrorRes } from '../schema';
import { DashboardService } from '../services';

@Route('dashboards')
@Produces('application/json')
@Security('authorizationHeader')
@Tags('Dashboard')
export class DashboardsController extends Controller {
    constructor(private dashboardsService: DashboardService) {
        super();
    }

    /**
     * List the dashboards with their id, name, etc.
     * @summary List the dashboards
     * @param queryParams
     * @returns
     */
    @Get()
    @OperationId('ListDashboards')
    @SuccessResponse(200)
    @Response<ErrorRes>(401, ERR_401, { message: ERR_401 })
    @Response<ErrorRes>(403, ERR_403, { message: ERR_403 })
    @Example<ListDashboardsRes>(
        DASHBOARDS,
        'A list of items with their id, name, etc.',
    )
    @Example<ListDashboardsRes>(
        {
            items: [],
            total: 0,
        },
        'An empty list',
    )
    public async list(
        @Queries() queryParams: ListDashboardsQueryParams,
    ): Promise<ListDashboardsRes> {
        return this.dashboardsService.list(queryParams);
    }

    /**
     * List a dashboard metrics
     *
     * For each metric, you can send the value or `null`. Typically, if the value is "long" to get, set `null` to return early from this call.
     * This will offer a better UX in the app by loading the metrics fast. The app will take care of fetching the actual value asynchronously.
     *
     * @summary List a dashboard metrics
     * @param queryParams
     * @returns
     */
    @Get('{id}/metrics')
    @SuccessResponse(200)
    @Response<ErrorRes>(401, ERR_401, { message: ERR_401 })
    @Response<ErrorRes>(403, ERR_403, { message: ERR_403 })
    @Response<ErrorRes>(404, ERR_404, { message: ERR_404 })
    @Example<ListDashboardMetricsRes>(
        DASHBOARD_METRICS_FOR_BUSINESS,
        'Business metrics',
    )
    @Example<ListDashboardsRes>(
        {
            items: [],
            total: 0,
        },
        'An empty list',
    )
    public async listMetrics(
        @Path() id: Dashboard['id'],
    ): Promise<ListDashboardMetricsRes> {
        return this.dashboardsService.listMetrics(id);
    }

    /**
     * Get a dashboard metric
     * @summary Get a dashboard metric
     * @param id
     * @param metricId
     * @returns
     */
    @Get('{id}/metrics/{metricId}')
    @SuccessResponse(200)
    @Response<ErrorRes>(401, ERR_401, { message: ERR_401 })
    @Response<ErrorRes>(403, ERR_403, { message: ERR_403 })
    @Response<ErrorRes>(404, ERR_404, { message: ERR_404 })
    @Example<GetDashboardMetricRes>(
        DASHBOARD_METRICS_FOR_BUSINESS_1,
        'A business metric',
    )
    public async getMetric(
        @Path() id: Dashboard['id'],
        @Path() metricId: DashboardMetric['id'],
    ): Promise<DashboardMetric> {
        return this.dashboardsService.getMetric(id, metricId);
    }
}
