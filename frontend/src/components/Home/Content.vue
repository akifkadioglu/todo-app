<template>
  <v-simple-table>
    <template v-slot:default>
      <thead>
        <tr>
          <th class="text-center">{{ $t("id") }}</th>
          <th class="text-center">{{ $t("name") }}</th>
          <th class="text-center">{{ $t("content") }}</th>
          <th class="text-center">{{ $t("delete") }}</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="(item, i) in TodoList" :key="item.name">
          <td>{{ item.ID }}</td>
          <td>{{ item.Title }}</td>
          <td>
            <v-btn depressed color="primary">
              <v-icon> mdi-eye </v-icon>
            </v-btn>
          </td>
          <td>
            <v-btn depressed color="error" @click="deleteTask(item.ID, i)">
              <v-icon> mdi-delete-outline </v-icon>
            </v-btn>
          </td>
        </tr>
      </tbody>
    </template>
  </v-simple-table>
</template>

<script>
export default {
  mounted() {
    this.getTodoList();
  },
  data() {
    return {
      TodoList: [
        {
          CreatedAt: "2022-10-02T06:21:05.278+03:00",
          ID: 1,
          Title: "not 1",
          Description: "dsaf",
        },
      ],
      api: process.env.VUE_APP_BASE_URL,
    };
  },
  methods: {
    getTodoList() {
      this.axios.get(this.api + "/getNotes").then((response) => {
        this.TodoList = response.data;
      });
    },
    deleteTask(id, index) {
      this.axios
        .delete(this.api + "/deleteNote", { params: { id } })
        .then((response) => {
          if (response.data) {
            this.TodoList.splice(index, 1);
          }
        });
    },
  },
};
</script>

<style></style>
