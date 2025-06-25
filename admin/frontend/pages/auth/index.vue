<template>
  <AuthForm
      :type="type"
      @switch="value => type = value"
      @action="(login, password) => action(login, password)"
  />
</template>
<script setup lang="ts">

import AuthForm from "~/components/organisms/AuthForm.vue";

const type: Ref<"sign_in" | "sign_up"> = ref("sign_in");


const action = (login: string, password: string) => {
  console.warn('ACTION:', type.value, login, password);
}

</script>