<template>
    <div class="section">
        <h2 class="title is-3">
            PK Turnover
        </h2>
        <table class="table">
            <tr>
                <th>Month</th>
                <th>Services</th>
                <th>Products</th>
                <th>FL Services</th>
                <th>FL Products</th>
                <th>Total</th>
            </tr>
            <tr v-for="(t, index) in t">
                <td>{{t.month_year | moment("MMMM")}}</td>
                <td>{{t.services | toCurrency}}</td>
                <td>{{t.products | toCurrency}}</td>
                <td>{{t.fl_services | toCurrency}}</td>
                <td>{{t.fl_products | toCurrency}}</td>
                <td><strong>{{t.total | toCurrency}}</strong></td>
            </tr>
        </table>
        <p class="is-size-3">Total: {{total | toCurrency}}</p>
    </div>
</template>
<script>
    export default {
        data() {
            return {
                t: []
            }
        },

        computed: {
            total() {
                return this.t.reduce((sum, val) => sum + val.total, 0).toFixed(2);
            }
        },

        created() {
            axios.get('/api/takings/PK').then(response => this.t = response.data)
                .catch(error => {
                    console.log(error)
                })
        }
    }
</script>