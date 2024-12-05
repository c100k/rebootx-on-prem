/**
 * The status of a runnable
 * Any intermediary status that you have on your side must be mapped to the `pending` status.
 */
export var RunnableStatus;
((RunnableStatus) => {
    RunnableStatus['OFF'] = 'off';
    RunnableStatus['ON'] = 'on';
    RunnableStatus['PENDING'] = 'pending';
    RunnableStatus['UNKNOWN'] = 'unknown';
})(RunnableStatus || (RunnableStatus = {}));
//# sourceMappingURL=runnable.js.map
