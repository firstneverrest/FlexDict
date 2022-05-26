<template>
  <main
    class="container flex flex-col justify-center items-center w-full h-screen"
  >
    <img src="~/assets/banner.png" alt="banner" />
    <h4 class="text-secondary">Never stop learning together</h4>
    <p class="mt-2">Find and collect new words every day.</p>
    <form class="flex flex-col justify-center w-4/6 mt-6">
      <input
        type="text"
        class="input-border border-gray-600 rounded-2xl py-2 px-3 text-xs"
        placeholder="Username"
      />
      <input
        type="password"
        class="input-border border-gray-600 rounded-2xl py-2 px-3 text-xs mt-2 mb-4"
        placeholder="Password"
      />
      <button
        class="bg-primary rounded-2xl py-2 px-3 outline-none cursor-pointer"
      >
        Login
      </button>
    </form>
    <p class="mt-4">
      No FlexDict account? <a href="#" class="underline">sign up</a>
    </p>
  </main>
</template>

<script lang="ts">
import Vue from 'vue'

export default Vue.extend({
  name: 'LoginPage',
})
</script>
