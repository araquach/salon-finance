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
            <tr v-for="(jf, index) in jf">
                <td>{{jf.month_year | moment("MMMM")}}</td>
                <td>{{jf.services}}</td>
                <td>{{jf.products}}</td>
                <td>{{jf.fl_services}}</td>
                <td>{{jf.fl_products}}</td>
                <td></td>
            </tr>
        </table>
    </div>
</template>

<script>
    export default {
        props: ['jf']
    }
</script>