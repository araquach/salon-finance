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
            <tr v-for="(pf, index) in pf">
                <td>{{pf.month_year | moment("MMMM")}}</td>
                <td>{{pf.services}}</td>
                <td>{{pf.products}}</td>
                <td>{{pf.fl_services}}</td>
                <td>{{pf.fl_products}}</td>
                <td></td>
            </tr>
        </table>
    </div>
</template>
<script>
    export default {
        props: ['pf']
    }
</script>