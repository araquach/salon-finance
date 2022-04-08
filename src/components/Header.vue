<template>
  <b-navbar spaced>
    <template slot="brand">
      <b-navbar-item :to="{ name: 'dashboard' }" tag="router-link">
        <h1 class="title is-3">Salon Finance</h1>
      </b-navbar-item>
    </template>
    <template slot="start">
      <b-navbar-item :to="{ name: 'last-years-takings' }" tag="router-link">
        Last Years Takings
      </b-navbar-item>
      <b-navbar-item :to="{ name: 'last-eight-months' }" tag="router-link">
        Last Eight Months
      </b-navbar-item>
      <b-navbar-item :to="{ name: 'cost-breakdown' }" tag="router-link">
        Cost Breakdown
      </b-navbar-item>
    </template>
    <template slot="end">

    </template>
  </b-navbar>
</template>