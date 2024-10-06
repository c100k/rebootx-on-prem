export interface ListQueryParams {
    /**
     * Max number of items to return in the response. Set a reasonable default value (e.g. `50`) in your implementation. Avoid returning too many items at once for client performance reasons.
     * @isInt
     * @minimum 0
     */
    limit?: number;

    /**
     * Cursor from where to start fetching when paginating. The default value should be `0`.
     * @isInt
     * @minimum 0
     */
    offset?: number;
}

export interface ListRes<T extends {}> {
    items: T[];

    /**
     * @isInt
     * @minimum 0
     */
    total: number;
}
