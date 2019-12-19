<template>
    <tr>
        <td>{{c.date | moment("MMM")}}</td>
        <td>{{c.account}}</td>
        <td>{{c.description}}</td>
        <td>{{c.amount.toFixed(2)}}</td>
        <td>
            <b-dropdown v-model="category" aria-role="list">
                <button class="button is-primary" slot="trigger">
                    <span>Category</span>
                    <b-icon icon="menu-down"></b-icon>
                </button>

                <b-dropdown-item @click="addCategory" aria-role="listitem" value="Wages">Wages</b-dropdown-item>
                <b-dropdown-item @click="addCategory" aria-role="listitem" value="Freelance">Freelance</b-dropdown-item>
                <b-dropdown-item @click="addCategory" aria-role="listitem" value="Drawings">Drawings</b-dropdown-item>
                <b-dropdown-item @click="addCategory" aria-role="listitem" value="Stock">Stock</b-dropdown-item>
                <b-dropdown-item @click="addCategory" aria-role="listitem" value="VAT">VAT</b-dropdown-item>
                <b-dropdown-item @click="addCategory" aria-role="listitem" value="Tax">TAX</b-dropdown-item>
                <b-dropdown-item @click="addCategory" aria-role="listitem" value="Building">Building</b-dropdown-item>
                <b-dropdown-item @click="addCategory" aria-role="listitem" value="Marketing">Marketing</b-dropdown-item>
                <b-dropdown-item @click="addCategory" aria-role="listitem" value="Condements">Condements</b-dropdown-item>
                <b-dropdown-item @click="addCategory" aria-role="listitem" value="Bank">Bank</b-dropdown-item>
                <b-dropdown-item @click="addCategory" aria-role="listitem" value="Utilities">Utilities</b-dropdown-item>
                <b-dropdown-item @click="addCategory" aria-role="listitem" value="Loans">Loans</b-dropdown-item>
                <b-dropdown-item @click="addCategory" aria-role="listitem" value="Staff">Staff Expenses</b-dropdown-item>
                <b-dropdown-item @click="addCategory" aria-role="listitem" value="Misc">Misc Purchases</b-dropdown-item>
                <b-dropdown-item @click="addCategory" aria-role="listitem" value="Other">Other</b-dropdown-item>

            </b-dropdown>
        </td>
        <td v-if="!category">{{c.category}}</td>
        <td v-else>{{category}}</td>
    </tr>
</template>

<script>
    export default {
        props: ['c'],

        data() {
            return {
                category: ''
            }
        },

        methods: {
            addCategory() {

                axios.put('/api/bankdata/' + this.c.id, {
                    category: this.category
                })
                    .then(response => {
                        this.resetCategory
                    })
                    .catch(error => {
                        console.log(err)
                    })
            },

            resetCategory() {
                this.category = ''
            }
        },

        created() {

        }

    }
</script>