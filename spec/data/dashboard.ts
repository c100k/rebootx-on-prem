import { ListDashboardsRes } from '../model/index.js';

export const DASHBOARDS: ListDashboardsRes = {
    items: [
        {
            id: '123',
            metrics: [],
            name: 'Infra',
        },
        {
            id: '456',
            metrics: [
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
            name: 'Business',
        },
    ],
    total: 2,
};
