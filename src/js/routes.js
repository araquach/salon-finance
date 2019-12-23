import MainDashboard from "./components/MainDashboard"
import CostMain from './components/bank/CostMain'
import CostInput from './components/bank/CostInput'
import CostTotals from './components/bank/CostTotals'
import MainTakings from './components/takings/MainTakings'
import ProfitLossMonthly from './components/ProfitLoss/ProfitLossMonthly'


export const routes = [
    { path: '', component: MainDashboard},
    { path: '/costs', component: CostTotals},
    { path: '/input', component: CostInput},
    { path: '/totals', component: MainTakings},
    { path: '/profitloss', component: ProfitLossMonthly}
]