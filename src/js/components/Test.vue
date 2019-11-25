<template>
    <div class="section">
        <h1 class="title">Test</h1>
        <b-dropdown aria-role="list" v-model="selectedSalon">
            <button class="button is-primary" slot="trigger">
                <span>Select Salon</span>
                <b-icon icon="menu-down"></b-icon>
            </button>
            <b-dropdown-item aria-role="listitem" value="Jakata">Jakata</b-dropdown-item>
            <b-dropdown-item aria-role="listitem" value="PK">PK</b-dropdown-item>
            <b-dropdown-item aria-role="listitem" value="Base">Base</b-dropdown-item>
            <b-dropdown-item aria-role="listitem" value="All">All</b-dropdown-item>
        </b-dropdown>
        <br><br>
        <table class="table">
            <tr>
                <th>Salon</th>
                <th>Services</th>
                <th>Products</th>
                <th>Total</th>
            </tr>
            <tr>
                <td>{{selectedSalon}}</td>
                <td>{{totalServices | toCurrency}}</td>
                <td>{{totalProducts | toCurrency}}</td>
                <td>{{grandTotal | toCurrency}}</td>

            </tr>
        </table>
    </div>
</template>

<script>
    export default {
        data() {
            return {
                takings: [],

                selectedSalon: 'All'
            }
        },

        computed: {
            filteredItem() {
                if (this.selectedSalon == 'All') {
                    return this.takings
                } else {
                    return this.takings.filter(item => {
                        return item.salon.indexOf(this.selectedSalon) > -1
                    })
                }
            },

            totalProducts() {
                return this.filteredItem.reduce((sum, val) => sum + val.products, 0).toFixed(2);
            },

            totalServices() {
                return this.filteredItem.reduce((sum, val) => sum + val.services, 0).toFixed(2);
            },

            grandTotal() {
                return Number(this.totalProducts) + Number(this.totalServices)
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