import MainDashboard from "./views/MainDashboard"
import CostTotals from "./components/bank/CostTotals"
import ProfitLossMonthly from "./components/ProfitLoss/ProfitLossMonthly"

export const routes = [
    {
        path: "",
        name: "dashboard",
        component: MainDashboard
    },
    {
        path: "/costs",
        name: "cost-totals",
        component: CostTotals
    },
    {
        path: "/profit-loss",
        name: "profit-loss",
        component: ProfitLossMonthly
    }
]