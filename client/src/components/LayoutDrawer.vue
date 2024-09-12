<script setup lang="ts">
   import { faBars } from '@fortawesome/free-solid-svg-icons';
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome';
   
</script>

<template>
<div class="drawer lg:drawer-open">
  <input id="my-drawer-2" type="checkbox" class="drawer-toggle" />
  <section class="drawer-content p-4 flex flex-col gap-2">
    <label for="my-drawer-2" class="lg:hidden cursor-pointer">
      <FontAwesomeIcon :icon="faBars" />
    </label>
    <div>
      <slot></slot>
    </div>
  </section>
  <div class="drawer-side">
    <label for="my-drawer-2" aria-label="close sidebar" class="drawer-overlay"></label>
    <ul class="menu bg-base-200 text-base-content min-h-full w-64 lg:w-32 p-4">
      <!-- Sidebar content here -->
      <li><RouterLink to="/dashboard">Dashboard</RouterLink></li>
      <li><a class="pointer-events-none">Settings</a></li>
    </ul>
  </div>
</div>
</template>