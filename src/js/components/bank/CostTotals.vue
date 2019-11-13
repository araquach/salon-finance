<template>
    <div class="section">
        <h1 class="title is-4 has-text-primary">Costs</h1>
        <table class="table">
            <tr>
                <th>Category</th>
                <th>Amount</th>
                <th>Percent</th>
            </tr>
            <tr>
                <th>Wages</th>
                <td>{{wagesTotal}}</td>
                <td>%</td>
            </tr>
            <tr>
                <th>Freelance</th>
                <td>{{freelanceTotal}}</td>
                <td>%</td>
            </tr>
            <tr>
                <th>Drawings</th>
                <td>{{drawingsTotal}}</td>
                <td>%</td>
            </tr>
            <tr>
                <th>Stock</th>
                <td>{{stockTotal}}</td>
                <td>%</td>
            </tr>
            <tr>
                <th>VAT</th>
                <td>{{vatTotal}}</td>
                <td>%</td>
            </tr>
            <tr>
                <th>TAX</th>
                <td>{{taxTotal}}</td>
                <td>%</td>
            </tr>
            <tr>
                <th>Building</th>
                <td>{{buildingTotal}}</td>
                <td>%</td>
            </tr>
            <tr>
                <th>Marketing</th>
                <td>{{marketingTotal}}</td>
                <td>%</td>
            </tr>
            <tr>
                <th>Condements</th>
                <td>{{condementsTotal}}</td>
                <td>%</td>
            </tr>
            <tr>
                <th>Bank</th>
                <td>{{bankTotal}}</td>
                <td>%</td>
            </tr>
            <tr>
                <th>Utilities</th>
                <td>{{utilitiesTotal}}</td>
                <td>%</td>
            </tr>
            <tr>
                <th>Loans</th>
                <td>{{loansTotal}}</td>
                <td>%</td>
            </tr>
            <tr>
                <th>Staff Expenses</th>
                <td>{{staffExpensesTotal}}</td>
                <td>%</td>
            </tr>
            <tr>
                <th>Misc</th>
                <td>{{miscTotal}}</td>
                <td>%</td>
            </tr>
            <tr>
                <th>Other</th>
                <td>{{otherTotal}}</td>
                <td>%</td>
            </tr>
        </table>
    </div>
</template>

<script>
    export default {
        data() {
            return {
                costs: [],
                wages: [],
                freelance: [],
                drawings: [],
                stock: [],
                vat: [],
                tax: [],
                building: [],
                marketing: [],
                condements: [],
                bank: [],
                utilities: [],
                loans: [],
                staff: [],
                misc: [],
                other: []
            }
        },

        computed: {
            total() {
                return this.costs.reduce((sum, val) => sum + val.debit_amount, 0).toFixed(2)
            },
            wagesTotal() {
                return this.wages.reduce((sum, val) => sum + val.debit_amount, 0).toFixed(2)
            },
            freelanceTotal() {
                return this.freelance.reduce((sum, val) => sum + val.debit_amount, 0).toFixed(2)
            },
            drawingsTotal() {
                return this.drawings.reduce((sum, val) => sum + val.debit_amount, 0).toFixed(2)
            },
            stockTotal() {
                return this.stock.reduce((sum, val) => sum + val.debit_amount, 0).toFixed(2)
            },
            vatTotal() {
                return this.vat.reduce((sum, val) => sum + val.debit_amount, 0).toFixed(2)
            },
            taxTotal() {
                return this.tax.reduce((sum, val) => sum + val.debit_amount, 0).toFixed(2)
            },
            buildingTotal() {
                return this.building.reduce((sum, val) => sum + val.debit_amount, 0).toFixed(2)
            },
            marketingTotal() {
                return this.marketing.reduce((sum, val) => sum + val.debit_amount, 0).toFixed(2)
            },
            condementsTotal() {
                return this.condements.reduce((sum, val) => sum + val.debit_amount, 0).toFixed(2)
            },
            bankTotal() {
                return this.bank.reduce((sum, val) => sum + val.debit_amount, 0).toFixed(2)
            },
            utilitiesTotal() {
                return this.utilities.reduce((sum, val) => sum + val.debit_amount, 0).toFixed(2)
            },
            loansTotal() {
                return this.loans.reduce((sum, val) => sum + val.debit_amount, 0).toFixed(2)
            },
            staffExpensesTotal() {
                return this.staff.reduce((sum, val) => sum + val.debit_amount, 0).toFixed(2)
            },
            miscTotal() {
                return this.misc.reduce((sum, val) => sum + val.debit_amount, 0).toFixed(2)
            },
            otherTotal() {
                return this.other.reduce((sum, val) => sum + val.debit_amount, 0).toFixed(2)
            }
        },

        created() {
            axios.get('/api/bankdata').then(response => this.costs = response.data)
                .catch(error => {
                    console.log(error)
                })
            axios.get('/api/costscategory/Stock').then(response => this.stock = response.data)
                .catch(error => {
                    console.log(error)
                })
            axios.get('/api/costscategory/Wages').then(response => this.wages = response.data)
                .catch(error => {
                    console.log(error)
                })
            axios.get('/api/costscategory/Freelance').then(response => this.freelance = response.data)
                .catch(error => {
                    console.log(error)
                })
            axios.get('/api/costscategory/Drawings').then(response => this.drawings = response.data)
                .catch(error => {
                    console.log(error)
                })
            axios.get('/api/costscategory/VAT').then(response => this.vat = response.data)
                .catch(error => {
                    console.log(error)
                })
            axios.get('/api/costscategory/Tax').then(response => this.tax = response.data)
                .catch(error => {
                    console.log(error)
                })
            axios.get('/api/costscategory/Building').then(response => this.building = response.data)
                .catch(error => {
                    console.log(error)
                })
            axios.get('/api/costscategory/Marketing').then(response => this.marketing = response.data)
                .catch(error => {
                    console.log(error)
                })
            axios.get('/api/costscategory/Condements').then(response => this.condements = response.data)
                .catch(error => {
                    console.log(error)
                })
            axios.get('/api/costscategory/Bank').then(response => this.bank = response.data)
                .catch(error => {
                    console.log(error)
                })
            axios.get('/api/costscategory/Utilities').then(response => this.utilities = response.data)
                .catch(error => {
                    console.log(error)
                })
            axios.get('/api/costscategory/Loans').then(response => this.loans = response.data)
                .catch(error => {
                    console.log(error)
                })
            axios.get('/api/costscategory/Staff').then(response => this.staff = response.data)
                .catch(error => {
                    console.log(error)
                })
            axios.get('/api/costscategory/Misc').then(response => this.misc = response.data)
                .catch(error => {
                    console.log(error)
                })
            axios.get('/api/costscategory/Other').then(response => this.other = response.data)
                .catch(error => {
                    console.log(error)
                })
        }
    }
</script>