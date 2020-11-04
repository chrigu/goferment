<template>
  <div class="steps">
    <ion-grid>
        <ion-row>
            <ion-col>
                <ion-item>
                    <ion-label position="stacked">Step name</ion-label>
                    <ion-input v-model="step.name"></ion-input>
                </ion-item>
            </ion-col>
            <ion-col>
                <ion-item>
                    <ion-label position="stacked">Duration</ion-label>
                    <ion-input v-model="step.duration" type="number"></ion-input>
                </ion-item>
            </ion-col>
            <ion-col>
                <ion-item>
                    <ion-label position="stacked">Temperature</ion-label>
                    <ion-input v-model="step.temperature" type="number"></ion-input>
                </ion-item>
            </ion-col>
            <ion-col>
                <div v-if="editable">
                    <ion-button color="primary" @click="addStep">Add Step</ion-button>
                </div>
                <div v-else>
                    <ion-button color="warning" @click="addStep">Add Step</ion-button>
                </div>
            </ion-col>
        </ion-row>
    </ion-grid>
  </div>
</template>

<script lang="ts">
// @ is an alias to /src
import {
    IonGrid,
    IonRow,
    IonItem,
    IonCol,
    IonLabel,
    IonInput
    } from '@ionic/vue';
import PageHeader from '../components/PageHeader.vue'
import { defineComponent } from 'vue';

export default defineComponent({
  name: 'ProfileStep',
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
        temperature: Number(this.step.temperature),
        duration: Number(this.step.duration),
        name: this.step.name
      })

      this.step = {
        temperature: 0,
        duration: 0,
        name: ''
      }
    }
  }
});
</script>
<style scoped lang="scss">
.step-field {
  &__label {
    display: none;
  }
}

</style>

