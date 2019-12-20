<template>
    <div class="section">
        <h1 class="title is-4">Costs</h1>
        <table class="table">
            <tr>
                <th>Category</th>
                <th>Amount</th>
                <th>Percent</th>
                <th>Average</th>
            </tr>
            <tr>
                <th>Wages (Excluding Izzy)</th>
                <td>{{wagesTotal | toCurrency}}</td>
                <td>{{wagesPercent}}</td>
                <td>{{wagesAverage | toCurrency}}</td>
            </tr>
            <tr>
                <th>Freelance</th>
                <td>{{freelanceTotal | toCurrency}}</td>
                <td>{{freelancePercent}}</td>
                <td>{{freelanceAverage | toCurrency}}</td>
            </tr>
            <tr>
                <th>Drawings (inc Izzy's wage)</th>
                <td>{{drawingsTotal | toCurrency}}</td>
                <td>{{drawingsPercent}}</td>
                <td>{{drawingsAverage | toCurrency}}</td>
            </tr>
            <tr>
                <th>Stock</th>
                <td>{{stockTotal | toCurrency}}</td>
                <td>{{stockPercent}}</td>
                <td>{{stockAverage | toCurrency}}</td>
            </tr>
            <tr>
                <th>VAT</th>
                <td>{{vatTotal | toCurrency}}</td>
                <td>{{vatPercent}}</td>
                <td>{{vatAverage | toCurrency}}</td>
            </tr>
            <tr>
                <th>TAX</th>
                <td>{{taxTotal | toCurrency}}</td>
                <td>{{taxPercent}}</td>
                <td>{{taxAverage | toCurrency}}</td>
            </tr>
            <tr>
                <th>Building</th>
                <td>{{buildingTotal | toCurrency}}</td>
                <td>{{buildingPercent}}</td>
                <td>{{buildingAverage | toCurrency}}</td>
            </tr>
            <tr>
                <th>Marketing</th>
                <td>{{marketingTotal | toCurrency}}</td>
                <td>{{marketingPercent}}</td>
                <td>{{marketingAverage | toCurrency}}</td>
            </tr>
            <tr>
                <th>Condements</th>
                <td>{{condementsTotal | toCurrency}}</td>
                <td>{{condementsPercent}}</td>
                <td>{{condementsAverage | toCurrency}}</td>
            </tr>
            <tr>
                <th>Bank (Inc PDQ costs)</th>
                <td>{{bankTotal | toCurrency}}</td>
                <td>{{bankPercent}}</td>
                <td>{{bankAverage | toCurrency}}</td>
            </tr>
            <tr>
                <th>Utilities</th>
                <td>{{utilitiesTotal | toCurrency}}</td>
                <td>{{utilitiesPercent}}</td>
                <td>{{utilitiesAverage | toCurrency}}</td>
            </tr>
            <tr>
                <th>Loans</th>
                <td>{{loansTotal | toCurrency}}</td>
                <td>{{loansPercent}}</td>
                <td>{{loansAverage | toCurrency}}</td>
            </tr>
            <tr>
                <th>Staff Expenses</th>
                <td>{{staffTotal | toCurrency}}</td>
                <td>{{staffPercent}}</td>
                <td>{{staffAverage | toCurrency}}</td>
            </tr>
            <tr>
                <th>Misc</th>
                <td>{{miscTotal | toCurrency}}</td>
                <td>{{miscPercent}}</td>
                <td>{{miscAverage | toCurrency}}</td>
            </tr>
            <tr>
                <th>Other</th>
                <td>{{otherTotal | toCurrency}}</td>
                <td>{{otherPercent}}</td>
                <td>{{otherAverage | toCurrency}}</td>
            </tr>
            <tr>
                <th>Grand Total</th>
                <td><strong>{{total | toCurrency}}</strong></td>
                <td></td>
                <td>{{totalAverage | toCurrency}}</td>
            </tr>
        </table>
    </div>
</template>

<script>
    export default {
        data() {
            return {
                numMonths: 8,
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
                return this.costs.reduce((sum, val) => sum + val.amount, 0).toFixed(2)
            },
            totalAverage() {
                return (parseInt(this.total) / parseInt(this.numMonths)).toFixed(2)
            },
            wagesTotal() {
                return this.wages.reduce((sum, val) => sum + val.amount, 0).toFixed(2)
            },
            wagesPercent() {
                return (parseInt(this.wagesTotal) / parseInt(this.total) * 100).toFixed(1)
            },
            wagesAverage() {
                return (parseInt(this.wagesTotal) / parseInt(this.numMonths)).toFixed(2)
            },
            freelanceTotal() {
                return this.freelance.reduce((sum, val) => sum + val.amount, 0).toFixed(2)
            },
            freelancePercent() {
                return (parseInt(this.freelanceTotal) / parseInt(this.total) * 100).toFixed(1)
            },
            freelanceAverage() {
                return (parseInt(this.freelanceTotal) / parseInt(this.numMonths)).toFixed(2)
            },
            drawingsTotal() {
                return this.drawings.reduce((sum, val) => sum + val.amount, 0).toFixed(2)
            },
            drawingsPercent() {
                return (parseInt(this.drawingsTotal) / parseInt(this.total) * 100).toFixed(1)
            },
            drawingsAverage() {
                return (parseInt(this.drawingsTotal) / parseInt(this.numMonths)).toFixed(2)
            },
            stockTotal() {
                return this.stock.reduce((sum, val) => sum + val.amount, 0).toFixed(2)
            },
            stockPercent() {
                return (parseInt(this.stockTotal) / parseInt(this.total) * 100).toFixed(1)
            },
            stockAverage() {
                return (parseInt(this.stockTotal) / parseInt(this.numMonths)).toFixed(2)
            },
            vatTotal() {
                return this.vat.reduce((sum, val) => sum + val.amount, 0).toFixed(2)
            },
            vatPercent() {
                return (parseInt(this.vatTotal) / parseInt(this.total) * 100).toFixed(1)
            },
            vatAverage() {
                return (parseInt(this.vatTotal) / parseInt(this.numMonths)).toFixed(2)
            },
            taxTotal() {
                return this.tax.reduce((sum, val) => sum + val.amount, 0).toFixed(2)
            },
            taxPercent() {
                return (parseInt(this.taxTotal) / parseInt(this.total) * 100).toFixed(1)
            },
            taxAverage() {
                return (parseInt(this.taxTotal) / parseInt(this.numMonths)).toFixed(2)
            },
            buildingTotal() {
                return this.building.reduce((sum, val) => sum + val.amount, 0).toFixed(2)
            },
            buildingPercent() {
                return (parseInt(this.buildingTotal) / parseInt(this.total) * 100).toFixed(1)
            },
            buildingAverage() {
                return (parseInt(this.buildingTotal) / parseInt(this.numMonths)).toFixed(2)
            },
            marketingTotal() {
                return this.marketing.reduce((sum, val) => sum + val.amount, 0).toFixed(2)
            },
            marketingPercent() {
                return (parseInt(this.marketingTotal) / parseInt(this.total) * 100).toFixed(1)
            },
            marketingAverage() {
                return (parseInt(this.marketingTotal) / parseInt(this.numMonths)).toFixed(2)
            },
            condementsTotal() {
                return this.condements.reduce((sum, val) => sum + val.amount, 0).toFixed(2)
            },
            condementsPercent() {
                return (parseInt(this.condementsTotal) / parseInt(this.total) * 100).toFixed(1)
            },
            condementsAverage() {
                return (parseInt(this.condementsTotal) / parseInt(this.numMonths)).toFixed(2)
            },
            bankTotal() {
                return this.bank.reduce((sum, val) => sum + val.amount, 0).toFixed(2)
            },
            bankPercent() {
                return (parseInt(this.bankTotal) / parseInt(this.total) * 100).toFixed(1)
            },
            bankAverage() {
                return (parseInt(this.bankTotal) / parseInt(this.numMonths)).toFixed(2)
            },
            utilitiesTotal() {
                return this.utilities.reduce((sum, val) => sum + val.amount, 0).toFixed(2)
            },
            utilitiesPercent() {
                return (parseInt(this.utilitiesTotal) / parseInt(this.total) * 100).toFixed(1)
            },
            utilitiesAverage() {
                return (parseInt(this.utilitiesTotal) / parseInt(this.numMonths)).toFixed(2)
            },
            loansTotal() {
                return this.loans.reduce((sum, val) => sum + val.amount, 0).toFixed(2)
            },
            loansPercent() {
                return (parseInt(this.loansTotal) / parseInt(this.total) * 100).toFixed(1)
            },
            loansAverage() {
                return (parseInt(this.loansTotal) / parseInt(this.numMonths)).toFixed(2)
            },
            staffTotal() {
                return this.staff.reduce((sum, val) => sum + val.amount, 0).toFixed(2)
            },
            staffPercent() {
                return (parseInt(this.staffTotal) / parseInt(this.total) * 100).toFixed(1)
            },
            staffAverage() {
                return (parseInt(this.staffTotal) / parseInt(this.numMonths)).toFixed(2)
            },
            miscTotal() {
                return this.misc.reduce((sum, val) => sum + val.amount, 0).toFixed(2)
            },
            miscPercent() {
                return (parseInt(this.miscTotal) / parseInt(this.total) * 100).toFixed(1)
            },
            miscAverage() {
                return (parseInt(this.miscTotal) / parseInt(this.numMonths)).toFixed(2)
            },
            otherTotal() {
                return this.other.reduce((sum, val) => sum + val.amount, 0).toFixed(2)
            },
            otherPercent() {
                return (parseInt(this.otherTotal) / parseInt(this.total) * 100).toFixed(1)
            },
            otherAverage() {
                return (parseInt(this.otherTotal) / parseInt(this.numMonths)).toFixed(2)
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