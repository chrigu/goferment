<template>
  <div class="config">
    <div class="columns">
      <div class="column is-one-quarter">
        <aside class="menu">
          <p class="menu-label">
            General
          </p>
          <ul class="menu-list">
            <li><router-link to="/config/profile">Profile</router-link></li>
            <li><a>Customers</a></li>
          </ul>
          <p class="menu-label">
            Administration
          </p>
        </aside>
      </div>
      <div class="column">
        <router-view/>
      </div>
    </div>
  </div>
</template>
