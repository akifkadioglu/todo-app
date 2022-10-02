<template>
  <div class="bottombar">
    <v-card>
      <div class="container">
        <div class="content">
          <div class="text-center">
            <v-menu top :offset-y="true">
              <template v-slot:activator="{ on, attrs }">
                <v-btn
                  color="primary"
                  dark
                  v-bind="attrs"
                  v-on="on"
                  width="131px"
                >
                  {{ $t($i18n.locale) }}
                </v-btn>
              </template>

              <v-list>
                <v-list-item v-for="(lang, i) in langs" :key="`Lang${i}`">
                  <v-list-item-title style="text-align: center">
                    <v-btn @click="changeLang(lang)" :value="lang">
                      {{ $t(lang) }}
                    </v-btn>
                  </v-list-item-title>
                </v-list-item>
              </v-list>
            </v-menu>
          </div>
        </div>
      </div>
    </v-card>
  </div>
</template>

<script>
export default {
  data() {
    return { langs: ["tr", "en"] };
  },
  methods: {
    changeLang(lang) {
      var isChanged = false;
      this.$i18n.locale = lang;
      isChanged = lang != localStorage.getItem(this.$store.state.lang);
      localStorage.setItem(this.$store.state.lang, lang);
      if (isChanged) {
        this.$store.state.snackbar = true;
        this.$store.state.snackbarMessage =
          this.$t("new_lang") + ": " + this.$t(lang);
      }
    },
  },
};
</script>

<style scoped>
.bottombar {
  position: fixed;
  width: 100%;
  bottom: 0px;
}
.container {
  display: table;
  position: relative;
  overflow: hidden;
  text-align: center;
  height: 70px;
  width: 100%;
}
.content {
  display: table-cell;
  vertical-align: middle;
  text-align: center;
}
</style>
