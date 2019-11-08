<template>
    <div class="section">
        <h1 class="title is-3">Costs</h1>

        <table class="table">
            <tr>
                <th>Date</th>
                <th>Account</th>
                <th>Description</th>
                <th>Amount</th>
                <th>Category</th>
            </tr>

            <IndCost v-for="(cost, index) in costs" :c="cost"/>

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

        created() {
            axios.get('/api/bankdata').then(response => this.costs = response.data)
                .catch(error => {
                    console.log(error)
                })
        }
    }
</script>