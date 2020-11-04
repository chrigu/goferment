<template>
<ion-app>
  <ion-split-pane when="sm" contentId="main-content">
    <ion-menu contentId="main-content">
      <ion-header>
        <ion-toolbar color="primary">
          <ion-title>Menu</ion-title>
        </ion-toolbar>
      </ion-header>

      <ion-content>
        <ion-list>
          <ion-list-header>
            Navigate
          </ion-list-header>
          <ion-menu-toggle auto-hide="false">
            <ion-item button router-link="/">
              <ion-label>
                Home
              </ion-label>
            </ion-item>
            <ion-item button router-link="/profile">
              <ion-label>
                Profile
              </ion-label>
            </ion-item>
          </ion-menu-toggle>
        </ion-list>
      </ion-content>
    </ion-menu>

    <ion-content class="ion-padding" id="main-content">
      <ion-router-outlet ></ion-router-outlet>
    </ion-content>

  </ion-split-pane>
</ion-app>

</template>
<style>
.my-custom-menu {
  --width: 500px;
}
</style>


<script lang="ts">
import { 
  IonContent, 
  IonHeader, 
  IonItem, 
  IonList, 
  IonListHeader, 
  IonMenu, 
  IonLabel,
  IonTitle, 
  IonRouterOutlet,
  IonIcon,
  IonSplitPane,
  IonApp,
  IonToolbar
} from '@ionic/vue';
import { defineComponent } from 'vue';

export default defineComponent({
  name: 'App',
  components: {
    IonContent, 
    IonHeader, 
    IonItem,
    IonLabel,
    IonList,
    IonListHeader,
    IonMenu, 
    IonRouterOutlet,
    IonTitle, 
    IonSplitPane,
    IonApp,
    IonToolbar
  }
});
</script>
