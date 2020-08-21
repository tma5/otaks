<template>
  <v-app id="inspire">
    <v-app-bar
      app
      clipped-left
    >
      <v-toolbar-title>otaks</v-toolbar-title>
    </v-app-bar>

    <v-main>
      <v-container
        class="fill-height"
        fluid
      >
        <Map></Map>
      </v-container>
    </v-main>

  </v-app>
</template>

<script>
  import Map from "@/components/Map"
  export default {
    props: {
      source: String,
    },
    components: {
      Map,
    },
    data: () => ({
      drawer: null,
    }),
    created () {
      this.$vuetify.theme.dark = true
    },
  }
</script>