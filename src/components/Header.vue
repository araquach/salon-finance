<template>
  <b-navbar spaced>
    <template slot="brand">
      <b-navbar-item :to="{ name: 'dashboard' }" tag="router-link">
        <h1 class="title is-3">Salon Finance</h1>
      </b-navbar-item>
    </template>
    <template slot="start">
      <b-navbar-item :to="{ name: 'takings-by-date-range' }" tag="router-link">
        Takings (date range)
      </b-navbar-item>
      <b-navbar-item :to="{ name: 'costs-by-category' }" tag="router-link">
        Costs by Category
      </b-navbar-item>
      <b-navbar-item :to="{ name: 'takings-by-stylist' }" tag="router-link">
        Takings by Stylist
      </b-navbar-item>
      <b-navbar-item :to="{ name: 'costs-by-month' }" tag="router-link">
        Costs by Month
      </b-navbar-item>
      <b-navbar-dropdown label="Menu">

      </b-navbar-dropdown>
    </template>
    <template slot="end">

    </template>
  </b-navbar>
</template>