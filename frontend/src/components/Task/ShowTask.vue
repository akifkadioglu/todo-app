<template>
  <div>
    <v-col cols="auto">
      <v-btn depressed color="primary" @click="getTask(id)">
        <v-icon> mdi-eye </v-icon>
      </v-btn>
      <v-dialog
        v-model="dialog"
        transition="dialog-bottom-transition"
        max-width="600"
      >
        <template>
          <v-card>
            <v-toolbar color="primary" dark elevation="0">
              <div
                class="subtitle-1 text-capitalize text-center"
                style="width: 100%"
              >
                {{ data.title }}
              </div>
            </v-toolbar>
            <v-card-text class="body-2 p-5">
              {{ data.description }}
            </v-card-text>
            <v-card-text class="body-2 p-5 text-end">
              {{ data.CreatedAt |  moment("Do MMMM YYYY, dddd")}}
            </v-card-text>

            <v-card-actions class="justify-end">
              <v-btn text @click="dialog = false">
                {{ $t("close") }}
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
  props: {
    id: { type: Number, default: 0 },
  },
  data() {
    return {
      dialog: false,
      api: process.env.VUE_APP_BASE_URL,
      rules: [(value) => !!value || this.$t("required")],
      data: {
        CreatedAt: "2022-10-02T19:19:15.52+03:00",
        DeletedAt: null,
        ID: 15,
        UpdatedAt: "2022-10-02T19:19:15.52+03:00",
        description: "tamam",
        title: "dünzenli spor yap",
      },
    };
  },
  methods: {
    getTask(id) {
      this.$store.state.overlay = true;
      this.axios
        .post(this.api + "/getTask", null, {
          params: {
            id,
          },
        })
        .then((response) => {
          this.data = response.data;
          this.dialog = true;
          this.$store.state.overlay = false;
        })
        .catch((err) => {
          console.log(err);
          this.$store.state.overlay = false;
        });
    },
  },
};
</script>
