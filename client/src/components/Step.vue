<template>
  <div class="step container">
    <div class="columns">
      <div class="column">
        <div class="field">
          <label class="label">Step name</label>
          <div class="control">
          <input class="input" type="text" placeholder="Text input" v-model="step.name">
          </div>
        </div>
      </div>
      <div class="column">
        <div class="field">
          <label class="label">Duration</label>
          <div class="control">
          <input class="input" type="text" placeholder="Text input" v-model="step.duration">
          </div>
        </div>
      </div>
      <div class="column">
        <div class="field">
          <label class="label">Temperature</label>
          <div class="control">
          <input class="input" type="text" placeholder="Text input" v-model="step.temperature">
          </div>
        </div>
      </div>
      <div class="column">
        <div v-if="editable">
          <button class="button is-primary" @click="addStep">Add step</button>
        </div>
        <div v-else>
          <button class="button is-primary" @click="addStep">Remove step</button>
        </div>
      </div>
    </div>

  </div>
</template>

<script>
// @ is an alias to /src

export default {
  name: 'Step',
  props: {
    temperature: {
      type: Number,
      default: 0
    },
    duration: {
      type: Number,
      default: 0
    },
    name: {
      type: String,
      default: ''
    },
    editable: {
      type: Boolean,
      default: false
    }
  },
  data () {
    return {
      step: {
        temperature: this.temperature,
        duration: this.duration,
        name: this.name
      }
    }
  },
  methods: {
    addStep () {
      this.$emit('add-step', {
        ...this.step
      })

      this.step = {
        temperature: 0,
        duration: 0,
        name: ''
      }
    }
  }
}
</script>
