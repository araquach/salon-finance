import MainDashboard from "./components/MainDashboard"
import CostTotals from './components/bank/CostTotals'
import ProfitLossMonthly from './components/ProfitLoss/ProfitLossMonthly'


export const routes = [
    { path: '', component: MainDashboard},
    { path: '/costs', component: CostTotals},
    { path: '/profit-loss', component: ProfitLossMonthly}
]