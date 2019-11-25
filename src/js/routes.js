import CostMain from './components/bank/CostMain'
import CostInput from './components/bank/CostInput'
import CostTotals from './components/bank/CostTotals'
import TotalTakings from './components/takings/TotalTakings'
import JakTakings from './components/takings/JakTakings'
import PkTakings from './components/takings/PkTakings'
import BaseTakings from './components/takings/BaseTakings'

import Test from './components/Test'

export const routes = [
    { path: '', component: CostMain},
    { path: '/test', component: Test},
    { path: '/costs', component: CostTotals},
    { path: '/input', component: CostInput},
    { path: '/total', component: TotalTakings},
    { path: '/jakata', component: JakTakings},
    { path: '/pk', component: PkTakings},
    { path: '/base', component: BaseTakings},
]