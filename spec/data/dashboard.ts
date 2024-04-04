import { ListDashboardMetricsRes, ListDashboardsRes } from '../model';

export const DASHBOARDS: ListDashboardsRes = {
    items: [
        {
            id: '123',
            name: 'Infra',
        },
        {
            id: '456',
            name: 'Business',
        },
    ],
    total: 2,
};

export const DASHBOARD_METRICS_FOR_BUSINESS: ListDashboardMetricsRes = {
    items: [
        {
            id: '123',
            label: 'Clients #',
            unit: null,
            value: 612,
        },
        {
            id: '456',
            label: 'Turnover',
            unit: '€',
            value: null,
        },
        {
            id: '789',
            label: 'NPS',
            unit: null,
            value: 95,
        },
    ],
    total: 3,
};
