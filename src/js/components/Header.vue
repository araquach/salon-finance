<template>
    <b-navbar spaced>
        <template slot="brand">
            <b-navbar-item tag="router-link" :to="{ path: '' }">
                <h1 class="title is-3">Salon Finance</h1>
            </b-navbar-item>
        </template>
        <template slot="start">
            <b-navbar-item tag="router-link" to="/">
                Dashboard
            </b-navbar-item>
            <b-navbar-dropdown label="Menu">
                <b-navbar-item tag="router-link" to="/total">
                    Total
                </b-navbar-item>
                <b-navbar-item tag="router-link" to="/jakata">
                    Jakata
                </b-navbar-item>
                <b-navbar-item tag="router-link" to="/pk">
                    PK
                </b-navbar-item>
                <b-navbar-item tag="router-link" to="/base">
                    Base
                </b-navbar-item>
                <b-navbar-item tag="router-link" to="/profitloss">
                    Monthly Profit/loss
                </b-navbar-item>
                <b-navbar-item tag="router-link" to="/input">
                    Category Input
                </b-navbar-item>
            </b-navbar-dropdown>
        </template>
    </b-navbar>
</template>