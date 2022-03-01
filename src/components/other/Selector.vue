<template>
  <div>
    <p>Salon</p>
    <input v-model="salon" type="text" label="salon">
    <b-field label="Select a date">
      <b-datepicker
          v-model="dateRange.startDate"
          placeholder="Click to select..."
          icon="calendar-today"
          :icon-right="selected ? 'close-circle' : ''"
          icon-right-clickable
          @icon-right-click="clearDate"
          trap-focus>
      </b-datepicker>
    </b-field>
    <b-field label="Select a date">
      <b-datepicker
          v-model="dateRange.endDate"
          placeholder="Click to select..."
          icon="calendar-today"
          :icon-right="selected ? 'close-circle' : ''"
          icon-right-clickable
          @icon-right-click="clearDate"
          trap-focus>
      </b-datepicker>
    </b-field>
    <button @click="updateInfo" class="button">Update</button>
  </div>
</template>
<script>
  import {mapMutations} from "vuex"

  export default {
    data() {
      return {
        salon: '',
        dateRange: {
          startDate: new Date(),
          endDate: new Date()
        }
      }
    },

    methods: {
      updateInfo() {
        this.UPDATE_SALON(this.salon)
        this.UPDATE_DATE_RANGE(this.dateRange)
      },

      ...mapMutations([
          'UPDATE_SALON',
          'UPDATE_DATE_RANGE'
      ])
    }
  }
</script>