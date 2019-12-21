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
            <tr>

            </tr>
        </table>
    </div>
</template>

<script>
    export default {
        data() {
            return {
                numMonths: 8,
                costs: [],
                categories: ['Wages', 'Freelance', 'Drawings', 'Stock', 'Vat', 'Tax', 'Building', 'Marketing', 'Condements', 'Bank', 'Utilities', 'Loans', 'Staff', 'Misc', 'Other']
            }
        },

        computed: {
            totalAverage() {
                return (parseInt(this.total) / parseInt(this.numMonths)).toFixed(2)
            },
            categoryTotal_backup() {
                return this.wages.reduce((sum, val) => sum + val.amount, 0).toFixed(2)
            },
            categoryTotal() {
                return this.categories.forEach(element => element.reduce((sum, val) => sum + val.amount, 0).toFixed(2));
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