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
            </tr>
        </table>
    </div>
</template>

<script>
    export default {
        data() {
            return {
                numMonths: 8,
                costs: []
            }
        },

        computed: {
            totalAverage() {
                return (parseInt(this.total) / parseInt(this.numMonths)).toFixed(2)
            },
            categoryTotal() {
                const result = [];
                this.costs.reduce(function(res, value) {
                    if (!res[value.category]) {
                        res[value.category] = { category: value.category, amount: 0 }
                        result.push(res[value.category])
                    }
                    res[value.category].amount += value.amount
                    return res
                }, {})
                return result
            },
            categoryPercent() {
                return (parseInt(this.wagesTotal) / parseInt(this.total) * 100).toFixed(1)
            },
            categoryAverage() {
                return (parseInt(this.wagesTotal) / parseInt(this.numMonths)).toFixed(2)
            },
            total() {
                return this.costs.reduce((sum, val) => sum + val.amount, 0).toFixed(2)
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