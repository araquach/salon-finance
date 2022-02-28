import Vue from 'vue'
import Router from 'vue-router'
import MainDashboard from "./views/MainDashboard"
import CostsByCategory from "./components/costs/CostsByCategory"
import CostsByMonth from "./components/costs/CostsByMonth"
import TakingsByDateRange from "./components/takings/TakingsByDateRange"
import TakingsByStylist from "./components/takings/TakingsByStylist"

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
    ]
})

export default router