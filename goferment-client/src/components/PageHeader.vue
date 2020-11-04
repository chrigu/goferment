<template>
    <ion-header :translucent="true">
        <ion-toolbar>
            <ion-buttons slot="start">
            <ion-menu-toggle>
                <ion-button>
                <ion-icon slot="icon-only" name="menu"></ion-icon>
                </ion-button>
            </ion-menu-toggle>
            </ion-buttons>
            <ion-title>{{title}}</ion-title>
        </ion-toolbar>
    </ion-header>
</template>


<script lang="ts">
import { 
  IonHeader, 
  IonTitle, 
  IonToolbar,
  IonIcon,
  IonMenuToggle,
} from '@ionic/vue';
import { defineComponent } from 'vue';

export default defineComponent({
  name: 'PageHeader',
  components: {
    IonTitle, 
    IonIcon,
    IonToolbar,
    IonMenuToggle
  },
  props: ['title']
});
</script>
