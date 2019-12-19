import MainDashboard from "./components/MainDashboard";
import CostMain from './components/bank/CostMain'
import CostInput from './components/bank/CostInput'
import CostTotals from './components/bank/CostTotals'
import MainTakings from './components/takings/MainTakings'
import JakTakings from './components/takings/JakTakings'
import PkTakings from './components/takings/PkTakings'
import BaseTakings from './components/takings/BaseTakings'
import ProfitLossMonthlyMonthly from './components/ProfitLoss/ProfitLossMonthly'

import Test from './components/Test'
import ProfitLossMonthly from "./components/ProfitLoss/ProfitLossMonthly";


export const routes = [

    { path: '', component: MainDashboard},
    { path: '', component: CostMain},
    { path: '/test', component: Test},
    { path: '/costs', component: CostTotals},
    { path: '/input', component: CostInput},
    { path: '/total', component: MainTakings},
    { path: '/jakata', component: JakTakings},
    { path: '/pk', component: PkTakings},
    { path: '/base', component: BaseTakings},
    { path: '/profitloss', component: ProfitLossMonthlyMonthly}
]