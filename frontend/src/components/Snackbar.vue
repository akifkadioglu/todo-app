<template>
  <v-snackbar
    :timeout="-1"
    :value="true"
    elevation="24"
    left
    v-model="this.$store.state.snackbar"
  >
    {{ this.$store.state.snackbarMessage }}
    <template v-slot:action="{ attrs }">
      <v-btn color="red" text v-bind="attrs" @click="closeSnackbar()">
        {{ $t("close") }}
      </v-btn>
    </template>
  </v-snackbar>
</template>
<script>
export default {
  methods: {
    closeSnackbar() {
      this.$store.state.snackbar = false;
    },
  },
};
</script>
