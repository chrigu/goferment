<template>
  <div class="steps">
    <ion-grid>
        <ion-row>
            <ion-col>
                <ion-item>
                    <ion-label position="stacked">Step name</ion-label>
                    <ion-input 
                      name="name"
                      :value="step.name"
                      :disabled="!editable"
                      @ionChange="handleChange"></ion-input>
                </ion-item>
            </ion-col>
            <ion-col>
                <ion-item>
                    <ion-label position="stacked">Duration</ion-label>
                    <ion-input
                      name="duration"
                      :value="step.duration"
                      :disabled="!editable"
                      @ionChange="handleChange" type="number"></ion-input>
                </ion-item>
            </ion-col>
            <ion-col>
                <ion-item>
                    <ion-label position="stacked">Temperature</ion-label>
                    <ion-input
                      name="temperature"
                      :value="step.temperature"
                      :disabled="!editable"
                      @ionChange="handleChange" type="number"></ion-input>
                </ion-item>
            </ion-col>
            <ion-col>
                <div v-if="editable">
                    <ion-button color="primary" @click="addStep">Add Step</ion-button>
                </div>
                <div v-else>
                    <ion-button color="warning" @click="removeStep">Remove</ion-button>
                </div>
            </ion-col>
        </ion-row>
    </ion-grid>
  </div>
</template>

<script lang="js">
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
import { defineComponent, reactive, toRefs, ref } from 'vue';


export default defineComponent({
  name: 'ProfileStep',
    setup (props, {emit}) {
    const { temperature, duration, name, editable } = toRefs(props)

    const step = ref({
      temperature: temperature.value,
      duration: duration.value,
      name: name.value
    })

    const addStep = () => {
      emit('add-step', {
        temperature: step.value.temperature,
        duration: step.value.duration,
        name: step.value.name,
      })

      step.value.temperature = 0
      step.value.duration = 0
      step.value.name = ''
    }

    const removeStep = () => {
      emit('remove-step')
    }

    const handleChange = (e) => {
      console.log(e, step.value);
      const propName = e.target.name;
      step.value[propName] = e.detail.value;
    };

    return {
        step,
        addStep,
        handleChange,
        removeStep
        }
    },
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
   emits: [ 'add-step', 'remove-step' ]
});
</script>
<style scoped lang="scss">
.step-field {
  &__label {
    display: none;
  }
}

</style>

