<template>
    <div class="section">
        <h2 class="title is-3">
            Salon Turnover
        </h2>
        <table class="table">
            <tr>
                <th>Salon</th>
                <th>Month</th>
                <th>Services</th>
                <th>Products</th>
                <th>FL Services</th>
                <th>FL Products</th>
                <th>Total</th>
            </tr>
            <tr v-for="(t, index) in t">
                <td>{{t.salon}}</td>
                <td>{{t.month_year | moment("MMMM")}}</td>
                <td>{{t.services}}</td>
                <td>{{t.products}}</td>
                <td>{{t.fl_services}}</td>
                <td>{{t.fl_products}}</td>
                <td></td>
            </tr>
        </table>
    </div>
</template>
<script>
    export default {
        props: ['t']
    }
</script>