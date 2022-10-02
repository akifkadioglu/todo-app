<template>
  <v-simple-table fixed-header height="520px">
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
          <td class="tableRow">{{ item.title }}</td>
          <td>
            <ContentShowTask :id="item.ID" />
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
import ContentShowTask from "../Task/ShowTask.vue";
export default {
  mounted() {
    this.$store.state.isLoading = true;
    this.getTodoList();
  },
  components: {
    ContentShowTask,
  },
  data() {
    return {
      TodoList: [],
      api: process.env.VUE_APP_BASE_URL,
    };
  },
  computed: {
    computedList: function () {
      var vm = this;
      return this.TodoList.filter(function (item) {
        return item.msg.toLowerCase().indexOf(vm.query.toLowerCase()) !== -1;
      });
    },
  },
  methods: {
    getTodoList() {
      this.axios.get(this.api + "/getTasks").then((response) => {
        this.TodoList = response.data;
        this.$store.state.isLoading = false;
      });
    },
    deleteTask(id, index) {
      this.$store.state.overlay = true;
      this.axios
        .delete(this.api + "/deleteTask", { params: { id } })
        .then((response) => {
          if (response.data) {
            this.TodoList.splice(index, 1);
          }
          this.$store.state.overlay = false;
        });
    },
  },
};
</script>
<style scoped>
.tableRow {
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
  max-width: 150px;
}
</style>
