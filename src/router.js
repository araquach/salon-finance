import Vue from 'vue'
import Router from 'vue-router'
import MainDashboard from "./views/MainDashboard"
import CostsByCategory from "./components/costs/CostsByCategory"
import CostsByMonth from "./components/costs/CostsByMonth"
import TakingsByDateRange from "./components/takings/TakingsByDateRange"
import TakingsByStylist from "./components/takings/TakingsByStylist"
import ProfitLossMonthly from "./components/ProfitLoss/ProfitLossMonthly"
import LastYearsTakings from "./views/meeting/LastYearsTakings"
import LastEightMonths from "./views/meeting/LastEightMonths"
import CostBreakdown from "./views/meeting/CostBreakdown"
import FiguresChart from "./views/meeting/FiguresChart"
import StylistTakingsMonthByMonth from "./components/takings/StylistTakingsMonthByMonth"

Vue.use(Router)

const router = new Router({
    mode: 'history',
    routes: [
        {
            path: '',
            name: 'dashboard',
            component: MainDashboard,
        },
        {
            path: '/costs-by-category',
            name: 'costs-by-category',
            component: CostsByCategory
        },
        {
            path: '/costs-by-month',
            name: 'costs-by-month',
            component: CostsByMonth
        },
        {
            path: '/takings-by-date-range',
            name: 'takings-by-date-range',
            component: TakingsByDateRange
        },
        {
            path: '/takings-by-stylist',
            name: 'takings-by-stylist',
            component: TakingsByStylist
        },
        {
            path: '/profit-loss-monthly',
            name: 'profit-loss-monthly',
            component: ProfitLossMonthly
        },
        // meeting stuff
        {
            path: '/last-years-takings',
            name: 'last-years-takings',
            component: LastYearsTakings
        },
        {
            path: '/last-eight-months',
            name: 'last-eight-months',
            component: LastEightMonths
        },
        {
            path: '/cost-breakdown',
            name: 'cost-breakdown',
            component: CostBreakdown
        },
        {
            path: '/figures-chart',
            name: 'figures-chart',
            component: FiguresChart
        },
        {
            path: '/stylist-takings-month-by-month',
            name: 'stylist-takings-month-by-month',
            component: StylistTakingsMonthByMonth
        }
    ]
})

export default router