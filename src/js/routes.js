import MainDashboard from "./components/MainDashboard"
import CostTotals from './components/bank/CostTotals'
import MainTakings from './components/takings/MainTakings'
import ProfitLossMonthly from './components/ProfitLoss/ProfitLossMonthly'


export const routes = [
    { path: '', component: MainDashboard},
    { path: '/costs', component: CostTotals},
    { path: '/totals', component: MainTakings},
    { path: '/profit-loss', component: ProfitLossMonthly}
]