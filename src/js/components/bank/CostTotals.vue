<template>
    <div class="section">
        <h1 class="title is-4">Costs</h1>
        <table class="table">
            <tr>
                <th>Category</th>
                <th>Amount</th>
                <th>Percent</th>
                <th>Average</th>
            </tr>
            <tr v-for="category in categoryTotal">
                <td>{{category.category}}</td>
                <td>{{category.amount | toCurrency}}</td>
                <td>{{category.percent.toFixed(1)}}</td>
                <td>{{category.average | toCurrency}}</td>
            <tr>
                <th>Grand Total</th>
                <td><strong>{{grandTotal | toCurrency}}</strong></td>
                <td></td>
                <td>{{totalAverage | toCurrency}}</td>
            </tr>
            </tr>
        </table>
    </div>
</template>

<script>
    export default {
        data() {
            return {
                numMonths: 9,
                costs: []
            }
        },

        computed: {
            grandTotal() {
                return this.costs.reduce((sum, val) => sum + val.amount, 0).toFixed(2)
            },
            totalAverage() {
                return this.grandTotal / this.numMonths
            },
            categoryTotal() {
                let numMonths = this.numMonths
                const result = []
                const mainTotal = this.grandTotal
                this.costs.reduce(function(res, value) {
                    if (!res[value.category]) {
                        res[value.category] = { category: value.category, amount: 0, percent: 0, average: 0 }
                        result.push(res[value.category])
                    }
                    let total = res[value.category].amount += value.amount
                    res[value.category].percent = (total / mainTotal) * 100
                    res[value.category].average = total / numMonths
                    return res
                }, {})
                return result
            }
        },

        created() {
            axios.get('/api/bankdata').then(response => this.costs = response.data)
                .catch(error => {
                    console.log(error)
                })
        }
    }
</script>