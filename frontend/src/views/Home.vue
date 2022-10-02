<template>
  <div class="home">
    <HomeAppbar />
    <transition name="bounce">
      <v-card class="mx-auto base-card" v-if="show">
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

            <HomeCreateTask />
          </div>
          

          <v-divider></v-divider>

          <HomeContent />
        </v-card-text>
      </v-card>
    </transition>

    <HomeBottombar />
  </div>
</template>

<script>
// @ is an alias to /src
import HomeAppbar from "@/components/Home/Appbar.vue";
import HomeBottombar from "@/components/Home/Bottombar.vue";
import HomeContent from "@/components/Home/Content.vue";
import HomeCreateTask from "@/components/Task/CreateTask.vue";

export default {
  name: "Home",
  components: {
    HomeAppbar,
    HomeBottombar,
    HomeContent,
    HomeCreateTask,
  },
  mounted() {
    this.show = !this.show;
  },
  data() {
    return {
      snackbar: false,
      show: false,
    };
  },
  methods: {
    changeTheme() {
      this.$vuetify.theme.dark = !this.$vuetify.theme.dark;
      localStorage.setItem(
        this.$store.state.darkmode,
        !this.$vuetify.theme.dark ? "0" : "1"
      );
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
.bounce-enter-active {
  animation: bounce-in 0.7s;
}
.bounce-leave-active {
  animation: bounce-in 0.7s reverse;
}
@keyframes bounce-in {
  0% {
    transform: scale(0);
  }
  50% {
    transform: scale(1.5);
  }
  100% {
    transform: scale(1);
  }
}
</style>
