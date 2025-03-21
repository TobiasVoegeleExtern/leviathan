<template>
  <div>
    <h1>Welcome to Vue with TypeScript</h1>
    <p>Go Backend Response: {{ goMessage }}</p>
  </div>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import BackendService from "./services/BackendService"; // Import the service

export default defineComponent({
  name: "App",
  data() {
    return {
      goMessage: "Loading..." as string,
    };
  },
  async mounted() {
    try {
      this.goMessage = await BackendService.getPing();
    } catch (error) {
      this.goMessage = "Error fetching from Go backend";
      console.error(error);
    }
  },
});
</script>

<style scoped>
h1 {
  color: #42b983;
}
</style>