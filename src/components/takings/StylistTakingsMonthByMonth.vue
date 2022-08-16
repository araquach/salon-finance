<template>
  <div>
    <div>
      <div>
        <br>
        <b-dropdown aria-role="list">
          <template #trigger="{ active }">
            <b-button
                label="Select Salon"
                type="is-primary"
                :icon-right="active ? 'menu-up' : 'menu-down'" />
          </template>
          <b-dropdown-item @click="selectSalon('jakata')" aria-role="listitem">Jakata</b-dropdown-item>
          <b-dropdown-item @click="selectSalon('pk')" aria-role="listitem">PK</b-dropdown-item>
          <b-dropdown-item @click="selectSalon('base')" aria-role="listitem">Base</b-dropdown-item>
        </b-dropdown>

        <b-dropdown v-if="salon" aria-role="list">
          <template #trigger="{ active }">
            <b-button
                label="Select Stylist"
                type="is-primary"
                :icon-right="active ? 'menu-up' : 'menu-down'" />
          </template>
          <b-dropdown-item @click="selectStylist(stylist)" v-if="salon === 'jakata'" v-for="stylist in getJakataStylists" :key="stylist.id" aria-role="listitem">{{ stylist.first_name }}</b-dropdown-item>
          <b-dropdown-item @click="selectStylist(stylist)" v-if="salon === 'pk'" v-for="stylist in getPKStylists" :key="stylist.id" aria-role="listitem">{{ stylist.first_name }}</b-dropdown-item>
          <b-dropdown-item @click="selectStylist(stylist)" v-if="salon === 'base'" v-for="stylist in getBaseStylists" :key="stylist.id" aria-role="listitem">{{ stylist.first_name }}</b-dropdown-item>
        </b-dropdown>

        <b-dropdown v-if="stylistSelected" aria-role="list">
          <template #trigger="{ active }">
            <b-button
                label="Select Dates"
                type="is-primary"
                :icon-right="active ? 'menu-up' : 'menu-down'" />
          </template>
          <b-dropdown-item @click="selectDate('quarter')" aria-role="listitem">Last Quarter</b-dropdown-item>
          <b-dropdown-item @click="selectDate('year')" aria-role="listitem">Year</b-dropdown-item>
          <b-dropdown-item @click="selectDate('all')" aria-role="listitem">All Time</b-dropdown-item>
          <b-dropdown-item @click="showDatePicker" aria-role="listitem">Date Range</b-dropdown-item>
        </b-dropdown>
      </div>
    </div>
    <br>
    <LineChartGenerator v-if="loaded"
                        class="chart"
                        :chart-options="chartOptions"
                        :chart-data="getStylistTakingsMonthByMonth"
                        :chart-id="chartId"
                        :dataset-id-key="datasetIdKey"
                        :plugins="plugins"
                        :css-classes="cssClasses"
                        :styles="styles"
                        :width="width"
                        :height="height"
    />
  </div>
</template>

<script>
import {mapGetters, mapState} from "vuex"
import { Line as LineChartGenerator } from 'vue-chartjs/legacy'

import {
  Chart as ChartJS,
  Title,
  Tooltip,
  Legend,
  LineElement,
  LinearScale,
  CategoryScale,
  PointElement
} from 'chart.js'

ChartJS.register(
    Title,
    Tooltip,
    Legend,
    LineElement,
    LinearScale,
    CategoryScale,
    PointElement
)

export default {
  name: 'LineChart',
  components: {
    LineChartGenerator
  },
  props: {
    chartId: {
      type: String,
      default: 'line-chart'
    },
    datasetIdKey: {
      type: String,
      default: 'label'
    },
    width: {
      type: Number,
      default: 400
    },
    height: {
      type: Number,
      default: 400
    },
    cssClasses: {
      default: '',
      type: String
    },
    styles: {
      type: Object,
      default: () => {}
    },
    plugins: {
      type: Array,
      default: () => []
    }
  },
  data() {
    return {
      chartOptions: {
        responsive: true,
        maintainAspectRatio: false
      },
      stylistSelected: false
    }
  },

  methods: {
    selectSalon(i) {
      this.$store.commit('UPDATE_SALON', i)
    },

    selectStylist(stylist) {
      this.$store.commit('UPDATE_STYLIST', stylist)
      this.$store.dispatch('loadStylistTakingsMonthByMonth')
      this.stylistSelected = true
    },

    selectDate(d) {
      this.$store.commit('UPDATE_DATE_RANGE_TYPE', d)
      this.$store.dispatch('loadStylistTakingsMonthByMonth')
    },

    showDatePicker() {

    }
  },

  computed: {
    ...mapState({
      salon: state => state.salon,
      loaded: state => state.loaded
    }),

    ...mapGetters([
        'getStylistTakingsMonthByMonth',
        'getJakataStylists',
        'getPKStylists',
        'getBaseStylists',
        'getChosenStylist'
    ])
  },

  created() {
    this.$store.dispatch('loadStylists')
  }
}
</script>
