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
                takings: []
            }
        },


        created() {
            axios.get('/api/takings/All').then(response => this.takings = response.data)
                .catch(error => {
                    console.log(error)
                })
        }
    }
</script>