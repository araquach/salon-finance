<template>
    <div class="section">
        <h1 class="title is-3">Costs</h1>
        <p class="is-size-3">Total Costs: {{total}}</p>

        <table class="table">
            <tr>
                <th>Date</th>
                <th>Account</th>
                <th>Description</th>
                <th>Amount</th>
                <th>Category</th>
            </tr>

            <IndCost v-for="(cost, index) in costs" v-bind:key="cost.id" :c="cost"/>

        </table>
    </div>
</template>

<script>
    export default {
        data() {
            return {
                costs: []
            }
        },

        computed: {
            total() {
                return this.costs.reduce((sum, val) => sum + val.debit_amount, 0).toFixed(2);
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