<template>
  <div class="home">
    <HomeAppbar />

    <v-card class="mx-auto base-card">
      <v-card-text class="card-text">
        <div class="display-1 p-3 navbar">
          <div>
            <v-btn @click="changeTheme()" depressed color="primary">
              <v-icon v-if="$vuetify.theme.dark">
                mdi-white-balance-sunny
              </v-icon>
              <v-icon v-if="!$vuetify.theme.dark"> mdi-weather-night </v-icon>
            </v-btn>
          </div>
          <div>{{ $t("header") }}</div>
          <div>
            <v-btn depressed color="primary">
              <v-icon> mdi-pencil </v-icon>
            </v-btn>
          </div>
        </div>
        <v-divider></v-divider>
        <HomeContent />
      </v-card-text>
    </v-card>
    <Snackbar />
    <HomeBottombar />
  </div>
</template>

<script>
// @ is an alias to /src
import HomeAppbar from "@/components/Home/Appbar.vue";
import HomeBottombar from "@/components/Home/Bottombar.vue";
import HomeContent from "@/components/Home/Content.vue";
import Snackbar from "@/components/Snackbar.vue";

export default {
  name: "Home",
  components: {
    HomeAppbar,
    HomeBottombar,
    HomeContent,
    Snackbar,
  },
  data() {
    return {
      snackbar: false,
    };
  },
  methods: {
    changeTheme() {
      this.$vuetify.theme.dark = !this.$vuetify.theme.dark;
      this.$store.state.snackbar = true;
      this.$store.state.snackbarMessage = this.$vuetify.theme.dark
        ? this.$t("dark_mode_on")
        : this.$t("bright_mode_on");
    },
  },
};
</script>
<style scoped>
.base-card {
  height: 700px;
  min-width: 350px;
  max-width: 700px;
  margin-top: -50px;
}
.card-text {
  text-align: center;
  color: black;
  size: 90;
}
</style>
