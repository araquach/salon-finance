<template>
  <div>
    <table class="table">
      <tr>
        <th>Month</th>
        <th>Jakata</th>
        <th>PK</th>
        <th>Base</th>
        <th>All</th>
      </tr>
      <tr v-for="data in takingsAll">
        <td>{{ data.month | showMonth }}</td>
      </tr>
    </table>
  </div>
</template>
<script>
import {mapState} from "vuex"
import {format, parseISO} from "date-fns";

export default {
  filters: {
    showMonth: value => format(parseISO(value), 'LLLL')
  },

  computed: {
    ...mapState({
      takingsJakata: state => state.totalsByDateRange.jakata,
      takingsPK: state => state.totalsByDateRange.pk,
      takingsBase: state => state.totalsByDateRange.base,
      takingsAll: state => state.totalsByDateRange.all
    })
  }
}
</script>