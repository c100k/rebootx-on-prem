import {
    Controller,
    Example,
    Get,
    OperationId,
    Produces,
    Queries,
    Response,
    Route,
    Security,
    SuccessResponse,
    Tags,
} from 'tsoa';

import { DASHBOARDS } from '../data/dashboard.js';
import { ERR_401, ERR_403 } from '../data/error.js';
import type {
    ListDashboardsQueryParams,
    ListDashboardsRes,
} from '../model/dashboard.js';
import type { ErrorRes } from '../schema/error.js';
import type { DashboardService } from '../services/DashboardService.js';

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
     * @returns The list of dashboards
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
}
