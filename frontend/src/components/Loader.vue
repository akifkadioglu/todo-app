<template>
  <div class="text-center">
    <v-overlay :value="this.$store.state.overlay">
      <v-progress-circular indeterminate size="64"></v-progress-circular>
    </v-overlay>
  </div>
</template>
