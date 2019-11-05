<template>
    <div>
        <h1 class="title is-1">
            Base Turnover
        </h1>
        <table class="table">
            <tr>
                <th>Month</th>
                <th>Services</th>
                <th>Products</th>
            </tr>
            <tr v-for="(bf, index) in baseFigures">
                <td>{{bf.month_year}}</td>
                <td>{{bf.products}}</td>
                <td>{{bf.services}}</td>
            </tr>
        </table>
    </div>
</template>

<script>
    export default {
        data() {
            return {
                baseFigures: []
            }
        },

        created() {
            axios.get('/api/base').then(response => this.baseFigures = response.data)
                .catch(error => {
                    console.log(error)
                })
        }
    }
</script>