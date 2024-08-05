import {
    ListRunnablesQueryParams,
    ListRunnablesRes,
    Runnable,
    RunnableOperationRes,
} from '../model/index.js';

export interface RunnableService {
    list(params: ListRunnablesQueryParams): Promise<ListRunnablesRes>;
    reboot(id: Runnable['id']): Promise<RunnableOperationRes>;
    stop(id: Runnable['id']): Promise<RunnableOperationRes>;
}
