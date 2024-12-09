import {
    Controller,
    Example,
    Get,
    OperationId,
    Path,
    Post,
    Produces,
    Queries,
    Response,
    Route,
    Security,
    SuccessResponse,
    Tags,
} from 'tsoa';

import { ERR_401, ERR_403, ERR_404 } from '../data/error.js';
import {
    RUNNABLES,
    RUNNABLE_OP_ASYNC_DESC,
    RUNNABLE_OP_ASYNC_RES,
    RUNNABLE_OP_SYNC_DESC,
    RUNNABLE_OP_SYNC_RES,
} from '../data/runnable.js';
import type {
    ListRunnablesQueryParams,
    ListRunnablesRes,
    Runnable,
    RunnableOperationRes,
} from '../model/runnable.js';
import type { ErrorRes } from '../schema/error.js';
import type { RunnableService } from '../services/RunnableService.js';

@Route('runnables')
@Produces('application/json')
@Security('authorizationHeader')
@Tags('Runnable')
export class RunnablesController extends Controller {
    constructor(private runnablesService: RunnableService) {
        super();
    }

    /**
     * List the runnables with their name, status, etc.
     * @summary List the runnables
     * @param queryParams
     * @returns The list of runnables
     */
    @Get()
    @OperationId('ListRunnables')
    @SuccessResponse(200)
    @Response<ErrorRes>(401, ERR_401, { message: ERR_401 })
    @Response<ErrorRes>(403, ERR_403, { message: ERR_403 })
    @Example<ListRunnablesRes>(
        RUNNABLES,
        'A list of items with their name, status, etc.',
    )
    @Example<ListRunnablesRes>(
        {
            items: [],
            total: 0,
        },
        'An empty list',
    )
    public async list(
        @Queries() queryParams: ListRunnablesQueryParams,
    ): Promise<ListRunnablesRes> {
        return this.runnablesService.list(queryParams);
    }

    /**
     * Reboot a runnable
     * @summary Reboot a runnable
     * @param id
     * @returns The result of the operation
     */
    @Post('{id}/reboot')
    @SuccessResponse(201)
    @Response<ErrorRes>(401, ERR_401, { message: ERR_401 })
    @Response<ErrorRes>(403, ERR_403, { message: ERR_403 })
    @Response<ErrorRes>(404, ERR_404, { message: ERR_404 })
    @Example<RunnableOperationRes>(
        RUNNABLE_OP_ASYNC_RES,
        RUNNABLE_OP_ASYNC_DESC,
    )
    @Example<RunnableOperationRes>(RUNNABLE_OP_SYNC_RES, RUNNABLE_OP_SYNC_DESC)
    public async reboot(
        @Path() id: Runnable['id'],
    ): Promise<RunnableOperationRes> {
        return this.runnablesService.reboot(id);
    }

    /**
     * Stop a runnable
     * @summary Stop a runnable
     * @param id
     * @returns The result of the operation
     */
    @Post('{id}/stop')
    @SuccessResponse(201)
    @Response<ErrorRes>(401, ERR_401, { message: ERR_401 })
    @Response<ErrorRes>(403, ERR_403, { message: ERR_403 })
    @Response<ErrorRes>(404, ERR_404, { message: ERR_404 })
    @Example<RunnableOperationRes>(
        RUNNABLE_OP_ASYNC_RES,
        RUNNABLE_OP_ASYNC_DESC,
    )
    @Example<RunnableOperationRes>(RUNNABLE_OP_SYNC_RES, RUNNABLE_OP_SYNC_DESC)
    public async stop(
        @Path() id: Runnable['id'],
    ): Promise<RunnableOperationRes> {
        return this.runnablesService.stop(id);
    }
}
