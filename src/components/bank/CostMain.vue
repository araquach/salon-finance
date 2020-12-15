<template>
    <div class="section">
        <table class="table is-size-4">
            <tr>
                <th>Profit/Loss</th>
                <th>Monthly Profit/Loss</th>
            </tr>
            <tr>
                <td class="has-text-danger"></td>
                <td class="has-text-danger"></td>
            </tr>
        </table>
    </div>
</template>

<script>
  import { mapState } from "vuex"

  export default {
    computed: {
      ...mapState([])
    },

    created() {

    }
  }
</script>