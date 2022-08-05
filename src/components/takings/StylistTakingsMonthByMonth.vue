<template>
  <div>
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
import { mapGetters } from "vuex"
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
      loaded: true,
      chartOptions: {
        responsive: true,
        maintainAspectRatio: false
      }
    }
  },

  computed: {
    ...mapGetters([
        'getStylistTakingsMonthByMonth',
    ])
  },

  created() {
    this.$store.dispatch('loadStylistTakingsMonthByMonth')
  }
}
</script>
