<template>
    <div class="section">
        <h2 class="title is-3">
            Jakata Turnover
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
                <td>{{t.services | addVat | toCurrency}}</td>
                <td>{{t.products | addVat | toCurrency}}</td>
                <td>{{t.fl_services | addVat | toCurrency}}</td>
                <td>{{t.fl_products | addVat |toCurrency}}</td>
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

        filters: {
            addVat(t) {
                return parseFloat(t) + (parseFloat(t) * .25)
            }
        },

        computed: {
            total() {
                return this.t.reduce((sum, val) => sum + val.total, 0).toFixed(2);
            }
        },

        created() {
            axios.get('/api/takings/Jakata').then(response => this.t = response.data)
                .catch(error => {
                    console.log(error)
                })
        }
    }
</script>