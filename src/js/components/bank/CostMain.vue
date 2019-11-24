<template>
    <div class="section">

        <TotalTakings/>
        <CostTotals/>

        <p class="is-size-3 has-text-danger">Profit/loss: {{pl}}</p>
    </div>
</template>

<script>
    import CostTotals from "./CostTotals";
    import TotalTakings from "../takings/TotalTakings";

    export default {
        components: {TotalTakings, CostTotals},

        data() {
            return {
                takings: [],
                costs: []
            }
        },

        computed: {
            grandTotal() {
                return this.takings.reduce((sum, val) => sum + val.total, 0)
            },
            totalCosts() {
                return this.costs.reduce((sum, val) => sum + val.debit_amount, 0)
            },
            pl() {
                return "£" + (this.grandTotal - this.totalCosts).toFixed(2)
            }
        },

        created() {
            axios.get('/api/takings/All').then(response => this.takings = response.data)
                .catch(error => {
                    console.log(error)
                })
            axios.get('/api/bankdata').then(response => this.costs = response.data)
                .catch(error => {
                    console.log(error)
                })
        }
    }
</script>