<template>
  <ion-page>
    <page-header title="Go Profile!" />
    
    <ion-content>
  
      <div id="container">
        <Profile-Step :editable="true" />
      </div>
    </ion-content>
  </ion-page>
</template>

<script lang="ts">
import { IonContent, IonPage } from '@ionic/vue';
import PageHeader from '../components/PageHeader.vue'
import ProfileStep from '../components/ProfileStep.vue'
import { defineComponent } from 'vue';

export default defineComponent({
  name: 'Profile',
  components: {
    IonContent,
    IonPage,
    PageHeader,
    ProfileStep
  }
});
</script>
