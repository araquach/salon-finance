<template>
    <div class="section">
        <h1 class="title is-4 has-text-primary">Cost Input</h1>

        <div class="field">
            <label class="label">Search</label>
            <div class="control">
                <input v-model="search" class="input" type="text" placeholder="Search">
            </div>
        </div>

        <table class="table">
            <tr>
                <th>Date</th>
                <th>Account</th>
                <th>Description</th>
                <th>Amount</th>
                <th>Select</th>
                <th>Category</th>
            </tr>
            <IndCost v-for="(cost, index) in filteredItem" v-bind:key="cost.id" :c="cost"/>
        </table>
    </div>
</template>

<script>
    import IndCost from "./IndCost"

    export default {
        components: {IndCost},

        data() {
            return {
                search: '',
                costs: []
            }
        },

        computed: {
            total() {
                return this.costs.reduce((sum, val) => sum + val.debit_amount, 0).toFixed(2)
            },

            filteredItem() {
                return this.costs.filter(item => {
                    return item.trans_description.toLowerCase().indexOf(this.search.toLowerCase()) > -1
                })
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