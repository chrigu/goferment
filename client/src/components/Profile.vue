<template>
  <div class="profile">
    <h1 class="is-size-1">Edit Profile</h1>
    <p>Duration {{ profileDuration }}</p>
    <!-- <form> -->
      <div class="steps__title step-title container">
        <div class="columns">
          <div class="column is-one-quarter">
            <h3>Step name</h3>
          </div>
          <div class="column is-one-quarter">
            <h3>Duration</h3>
          </div>
          <div class="column is-one-quarter">
            <h3>Temperature</h3>
          </div>
        </div>
      </div>
      <ul id="steps">
        <li v-for="step in steps" :key="step.name">
          <Step v-bind="step" />
        </li>
        <Step @add-step="addStep" :editable="true" />
      </ul>
      <div v-if="steps.length > 0">
        <button class="button is-primary" >Start profile</button>
      </div>
    <!-- </form> -->
    <pre v-if="steps.length > 0">
      {{ profileJson }}
    </pre>
  </div>
</template>

<script>
import Step from '@/components/Step.vue'

import { mapGetters, mapActions } from 'vuex'

export default {
  name: 'Profile',
  components: {
    Step
  },
  computed: {
    ...mapGetters(['steps']),
    profileDuration () {
      return this.steps.reduce((duration, step) => step.duration + duration, 0)
    },
    profileJson () {
      return JSON.stringify({ steps: this.steps })
    }
  },
  methods: {
    ...mapActions(['addStep'])
  }
}
</script>
