import MainDashboard from "./views/MainDashboard"
import CostTotals from "./components/bank/CostTotals"
import TotalTakings from "./components/takings/TotalTakings"

export const routes = [
    {
        path: '',
        name: 'dashboard',
        component: MainDashboard,
    },
    {
        path: '/cost-totals',
        name: 'cost-totals',
        component: CostTotals
    },
    {
        path: '/total-takings',
        name: 'total-takings',
        component: TotalTakings
    }
]