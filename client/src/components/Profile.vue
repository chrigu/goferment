<template>
  <div class="profile">
    <h1 class="is-size-1">Edit Profile</h1>
    <!-- <form> -->
      <ul id="steps">
        <li v-for="step in steps" :key="step.name">
          <Step v-bind="step" />
        </li>
        <Step @add-step="addStep" :editable="true" />
      </ul>
    <!-- </form> -->
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
    ...mapGetters(['steps'])
  },
  methods: {
    ...mapActions(['addStep'])
  }
}
</script>
