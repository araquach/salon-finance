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
                <th>Monthly Average</th>
            </tr>
            <tr>
                <td>{{serviceTotal | addVat | toCurrency}}</td>
                <td>{{productTotal | addVat | toCurrency}}</td>
                <td>{{flServiceTotal | addVat | toCurrency}}</td>
                <td>{{flProductTotal | addVat | toCurrency}}</td>
                <td><strong>{{grandTotal | toCurrency}}</strong></td>
                <td>{{grandTotalAverage | toCurrency}}</td>
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

        filters: {
            addVat(t) {
                return parseFloat(t) + (parseFloat(t) * .25)
            }
        },

        computed: {
            grandTotal() {
                return this.t.reduce((sum, val) => sum + val.total, 0)
            },
            grandTotalAverage() {
                return (this.grandTotal / 9)
            },
            serviceTotal() {
                return this.t.reduce((sum, val) => sum + val.services, 0)
            },
            productTotal() {
                return this.t.reduce((sum, val) => sum + val.products, 0)
            },
            flServiceTotal() {
                return this.t.reduce((sum, val) => sum + val.fl_services, 0)
            },
            flProductTotal() {
                return this.t.reduce((sum, val) => sum + val.fl_products, 0)
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