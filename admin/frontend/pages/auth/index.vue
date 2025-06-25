<template>
  <AuthForm
      :type="type"
      @switch="value => type = value"
  />
</template>
<script setup lang="ts">

import AuthForm from "~/components/organisms/AuthForm.vue";

const type: Ref<"sign_in" | "sign_up"> = ref("sign_in");

</script>