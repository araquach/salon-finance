<template>
    <div>
        <table class="table is-narrow">
            <tr>
                <th>Months</th>
                <th v-for="(month in months">{{month}}</th>
            </tr>
            <tr>
                <th>Costs</th>
                <td v-for="(data, index) in monthCostTotal">{{data | toCurrency}}</td>
            </tr>
            <tr>
                <th>Takings</th>
                <td v-for="(data, index) in monthTakingsTotal">{{data | toCurrency}}</td>
            </tr>
        </table>
    </div>
</template>
<script>
    export default {
        data()  {
            return {
                bankData: [],
                takings: [],
                months: ['Jan', 'Feb', 'Mar', 'Apr', 'May', 'Jun', 'Jul', 'Aug', 'Sep', '0ct', 'Nov', 'Dec'],
            }
        },

        computed: {
            monthCostFilter() {
                const byMonth = [];
                for (let i = 0; i < 12; i++) {
                    byMonth.push(this.bankData.filter(d => new Date(d.date).getMonth() === i));
                }
                return byMonth;
            },
            monthCostTotal() {
                const total = [];
                for (let i = 0; i < 12; i++) {
                    let initialVal = 0;
                    total.push(this.monthCostFilter[i].reduce((acc, current) => acc + current.amount, initialVal));
                }
                return total;
            },
            monthTakingsFilter() {
                const byMonth =[];
                for (let i = 0; i < 12; i++) {
                    byMonth.push(this.takings.filter(d => new Date(d.month_year).getMonth() === i));
                }
                return byMonth;
            },
            monthTakingsTotal() {
                const total = [];
                for (let i = 0; i < 12; i++) {
                    let initialVal = 0;
                    total.push(this.monthTakingsFilter[i].reduce((acc, current) => acc + current.total, initialVal));
                }
                return total;
            }
        },

        created() {
            axios.get('/api/bankdata').then(response => this.bankData = response.data)
                .catch(error => {
                    console.log(error)
                })
            axios.get('/api/takings/All').then(response => this.takings = response.data)
                .catch(error => {
                    console.log(error)
                })
        }
    }
</script>