<template>
    <div class="section">
        <h2 class="title is-3">
            Base Turnover
        </h2>
        <table class="table">
            <tr>
                <th>Month</th>
                <th>Services</th>
                <th>Products</th>
                <th>Total</th>
            </tr>
            <tr v-for="(bf, index) in bf">
                <td>{{bf.month_year | moment("MMMM")}}</td>
                <td>{{bf.products}}</td>
                <td>{{bf.services}}</td>
                <td>{{total}}</td>
            </tr>
        </table>
    </div>
</template>

<script>
    export default {
        props: ['bf'],

        data() {
            return {

            }
        },

        computed: {
            total() {
                return this.bf.products + this.bf.services
            }
        }
    }
</script>