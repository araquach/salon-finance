<template>
    <div class="section">
        <h2 class="title is-4">
            Total Salon Turnover
        </h2>
        <table class="table is-size-5">
            <tr>
                <th>Services</th>
                <th>Products</th>
                <th>FL Services</th>
                <th>FL Products</th>
                <th><strong>Grand Total</strong></th>
            </tr>
            <tr>
                <td>{{serviceTotal | toCurrency}}</td>
                <td>{{productTotal | toCurrency}}</td>
                <td>{{flServiceTotal | toCurrency}}</td>
                <td>{{flProductTotal | toCurrency}}</td>
                <td><strong>{{grandTotal | toCurrency}}</strong></td>
            </tr>
        </table>
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
            grandTotal() {
                return this.t.reduce((sum, val) => sum + val.total, 0).toFixed(2);
            },
            serviceTotal() {
                return this.t.reduce((sum, val) => sum + val.services, 0).toFixed(2);
            },
            productTotal() {
                return this.t.reduce((sum, val) => sum + val.products, 0).toFixed(2);
            },
            flServiceTotal() {
                return this.t.reduce((sum, val) => sum + val.fl_services, 0).toFixed(2);
            },
            flProductTotal() {
                return this.t.reduce((sum, val) => sum + val.fl_products, 0).toFixed(2);
            }
        },

        created() {
            axios.get('/api/takings/All').then(response => this.t = response.data)
                .catch(error => {
                    console.log(error)
                })
        }
    }
</script>