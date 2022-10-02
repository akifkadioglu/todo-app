<template>
  <v-app>
    <v-main>
      <router-view />
    </v-main>
    <Snackbar />
    <Loader />
  </v-app>
</template>

<script>
import Snackbar from "@/components/Snackbar.vue";
import Loader from "@/components/Loader.vue";

export default {
  name: "App",
  components: {
    Snackbar,
    Loader,
  },
  mounted() {
    var locale = localStorage.getItem(this.$store.state.lang) ?? "";
    var darkmode = localStorage.getItem(this.$store.state.darkmode) ?? "0";
    if (locale == "") {
      localStorage.setItem(this.$store.state.lang, "en");
    }
    if (darkmode == "1") {
      localStorage.setItem(this.$store.state.darkmode, "1");
      this.$vuetify.theme.dark = true;
    } else {
      localStorage.setItem(this.$store.state.darkmode, "0");
      this.$vuetify.theme.dark = false;
    }
  },
};
</script>
<style>
/* width */
::-webkit-scrollbar {
  width: 5px;
  border-radius: 20;
}

/* Track */
::-webkit-scrollbar-track {
  background: #f1f1f1;
}

/* Handle */
::-webkit-scrollbar-thumb {
  background: #888;
}

/* Handle on hover */
::-webkit-scrollbar-thumb:hover {
  background: #555;
}
</style>
