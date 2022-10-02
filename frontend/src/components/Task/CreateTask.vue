<template>
  <div>
    <v-col cols="auto">
      <v-dialog transition="dialog-bottom-transition" max-width="600">
        <template v-slot:activator="{ on, attrs }">
          <div>
            <v-btn depressed color="primary" v-bind="attrs" v-on="on">
              <v-icon> mdi-pencil </v-icon>
            </v-btn>
          </div>
        </template>
        <template v-slot:default="dialog">
          <v-card>
            <v-toolbar color="primary" dark elevation="0">
              {{ $t("createTask") }}
            </v-toolbar>
            <v-card-text>
              <v-text-field
                :label="$t('title')"
                counter
                maxlength="60"
                :rules="rules"
                hide-details="auto"
                v-model="title"
              ></v-text-field>
              <v-text-field
                :label="$t('description')"
                hide-details="auto"
                v-model="description"
              ></v-text-field>
            </v-card-text>
            <v-card-actions class="justify-end">
              <v-btn text @click="dialog.value = false">
                {{ $t("close") }}
              </v-btn>
              <v-btn
                text
                @click="(dialog.value = false), addTask()"
                class="text-primary"
              >
                {{ $t("save") }}
              </v-btn>
            </v-card-actions>
          </v-card>
        </template>
      </v-dialog>
    </v-col>
  </div>
</template>

<script>
export default {
  name: "CreateTask",
  data() {
    return {
      api: process.env.VUE_APP_BASE_URL,
      title: null,
      description: null,
      rules: [(value) => !!value || this.$t("required")],
    };
  },
  methods: {
    addTask() {
      this.$store.state.overlay = true;

      this.axios
        .post(this.api + "/addTask", null, {
          params: {
            title: this.title,
            description: this.description,
          },
        })
        .then((response) => {
          this.$store.state.newTask = response.data;
          this.title = null;
          this.description = null;
          this.$store.state.overlay = false;
          this.$store.state.snackbar = true;
          this.$store.state.snackbarMessage = this.$t("taskCreated");
        })
        .catch((err) => {
          this.$store.state.overlay = false;
          this.$store.state.snackbar = true;
          this.$store.state.snackbarMessage = err.response.data["error"];
        });
    },
  },
};
</script>

<style></style>
