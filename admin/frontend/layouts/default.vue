<script setup lang="ts">

import {Navbar} from "~/components/organisms/Navbar";

</script>

<template>
  <div class="h-screen flex flex-col">
    <Navbar/>

    <div class="p-10 flex-grow">
      <NuxtPage/>
    </div>
  </div>

</template>

<style scoped>

</style>