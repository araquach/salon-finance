<template>
  <div class="section">
    <h1 class="title is-3">Breakdown of Costs</h1>
    <p class="is-size-4">Turnover: <span class="has-text-warning is-size-2">{{ takings | toCurrency }}</span></p>
    <br><br>
    <table class="table is-size-3">
      <tr>
        <th></th>
        <th>Category</th>
        <th>Amount</th>
        <th>Percent</th>
        <th>Monthly Average</th>
      </tr>
      <CostBreakDownItem :cost="cost" key="cost.id" v-for="cost in adaptedFigures"/>
      <tr v-if="showTotal" class="is-size-3">
        <td>Yearly</td>
        <td class="has-text-warning">{{ byYear | toCurrency }}</td>
      </tr>
    </table>
  </div>
</template>
<script>
import CostBreakDownItem from "./CostBreakDownItem"
import {mapState, mapMutations} from "vuex"

export default {
  components: {CostBreakDownItem},

  data() {
    return {
      figures: [],
      byYear: null,
      showTotal: false
    }
  },

  filters: {
    showMonth: value => format(parseISO(value), 'LLLL')
  },

  computed: {
    ...mapState({
      takings: state => state.totalTurnover,
    }),

    adaptedFigures() {
      return this.figures.filter((item) => item.category !== 'costs' && item.category !== 'drawings' && item.category !== 'other')
    }
  },

  created() {
      axios
          .get(`/api/costs-by-cat-meeting/all/2021-07-01/2022-02-28`)
          .then(r => r.data)
          .then(data => {
            this.figures = data.figures
            this.byYear = data.by_year
          })
  }
}
</script>
