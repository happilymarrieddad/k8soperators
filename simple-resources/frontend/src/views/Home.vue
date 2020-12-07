<template>
  <div class="home">
    <h1>This is a home page</h1>
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
    const response = await fetch(`${baseURL}`)
    const items = await response.json()

    this.items = items;
  }
}
</script>