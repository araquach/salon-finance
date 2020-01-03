<template>
    <div class="section">
        <table class="table is-size-4">
            <tr>
                <th>Profit/Loss</th>
                <th>Monthly Profit/Loss</th>
            </tr>
            <tr>
                <td class="has-text-danger">{{pl | toCurrency}}</td>
                <td class="has-text-danger">{{plAverage | toCurrency}}</td>
            </tr>
        </table>
    </div>
</template>

<script>
    export default {
        data() {
            return {
                takings: [],
                costs: [],
                numMonths: 12
            }
        },

        computed: {
            grandTotal() {
                return this.takings.reduce((sum, val) => sum + val.total, 0)
            },
            totalCosts() {
                return this.costs.reduce((sum, val) => sum + val.amount, 0)
            },
            pl() {
                return (this.grandTotal - this.totalCosts).toFixed(2)
            },
            plAverage() {
                return ((this.grandTotal - this.totalCosts) / this.numMonths).toFixed(2)
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