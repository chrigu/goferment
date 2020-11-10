<template>
  <ion-page>
    <page-header title="Go Profile!" />
    
    <ion-content>
  
      <div id="container">
        <ul id="steps">
          <li v-for="(step, i) in steps" :key="step.name">
            <ProfileStep  v-bind="step" @remove-step="removeStep(i)" />
          </li>
          <ProfileStep @add-step="addStep" :editable="true" />
        </ul>
      </div>
    </ion-content>
  </ion-page>
</template>

<script lang="js">
import { IonContent, IonPage } from '@ionic/vue';
import PageHeader from '../components/PageHeader.vue'
import ProfileStep from '../components/ProfileStep.vue'
import Step from '../components/ProfileStep.vue'
import { defineComponent, ref } from 'vue';


export default defineComponent({
  name: 'Profile',
  components: {
    IonContent,
    IonPage,
    PageHeader,
    ProfileStep
  },
  setup () {
    const steps = ref([])
    const addStep = (step) => {
      steps.value = [...steps.value, step]
    }

    const removeStep = (index) => {
      steps.value = [...steps.value.slice(0, index), ...steps.value.slice(index + 1)]
    }

    return {
      steps,
      addStep,
      removeStep
    }
  }
});
</script>

<style scoped>
.steps li {
  list-style: none;
}
</style>
