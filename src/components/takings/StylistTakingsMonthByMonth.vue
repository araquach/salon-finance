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

        <b-dropdown aria-role="list">
          <template #trigger="{ active }">
            <b-button
                label="Select Stylist"
                type="is-primary"
                :icon-right="active ? 'menu-up' : 'menu-down'" />
          </template>
          <b-dropdown-item v-for="stylist in getJakataStylists" aria-role="listitem">{{ stylist.first_name }}</b-dropdown-item>

        </b-dropdown>
      </div>
    </div>
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
      loaded: false,
      chartOptions: {
        responsive: true,
        maintainAspectRatio: false
      }
    }
  },

  methods: {
    selectSalon(i) {
      this.$store.commit('UPDATE_SALON', i)
    }
  },

  computed: {
    ...mapState({
      salons: state => state.salons,
      salon: state => state.salon
    }),

    ...mapGetters([
        'getStylistTakingsMonthByMonth',
        'getJakataStylists',
        'getPKStylists',
        'getBaseStylists'
    ])
  },

  created() {
    this.$store.dispatch('loadStylists')
    this.$store.dispatch('loadStylistTakingsMonthByMonth')
  }
}
</script>
