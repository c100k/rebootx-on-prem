import type {
    ListRunnablesQueryParams,
    ListRunnablesRes,
    Runnable,
    RunnableOperationRes,
} from '../model/runnable.js';

export interface RunnableService {
    list(params: ListRunnablesQueryParams): Promise<ListRunnablesRes>;
    reboot(id: Runnable['id']): Promise<RunnableOperationRes>;
    stop(id: Runnable['id']): Promise<RunnableOperationRes>;
}
