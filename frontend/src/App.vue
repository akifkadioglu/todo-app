<template>
  <v-app>
    <v-main>
      <router-view />
    </v-main>
  </v-app>
</template>

<script>
export default {
  name: "App",
  mounted() {
    var locale = localStorage.getItem(this.$store.state.lang) ?? "";
    if (locale == "") {
      localStorage.setItem(this.$store.state.lang, "en");
    }
  },
};
</script>
