<template>
  <div class="about">
    <h1>This is a about page</h1>
    <span>{{ items }}</span>
  </div>
</template>

<script>
const baseURL = `${process.env.VUE_APP_apiurl}`

export default {
  data() {
    return {
      items: []
    };
  },
  async mounted() {
    const response = await fetch(`${baseURL}/about`)
    const items = await response.json()

    this.items = items;
  }
}
</script>