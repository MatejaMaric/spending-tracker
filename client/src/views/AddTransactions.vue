<template>
  <div class="align">
    <textarea
      placeholder="Copy transaction table HTML here"
      cols="100"
      rows="20"
      v-model="tableHtml"
    ></textarea>
    <button @click="submitTable">Submit</button>
  </div>
</template>

<script lang="ts">
import { defineComponent, ref } from "vue";
import { useStore } from "vuex";

export default defineComponent({
  name: "AddTransactions",
  setup() {
    const tableHtml = ref("");
    const store = useStore();

    const submitTable = () => {
      store
        .dispatch("processTransactions", tableHtml.value)
        .then((res) => {
          tableHtml.value = "";
          alert(`Added transactions: ${res.added_rows}`);
        })
        .catch(() => alert("Bad request!"));
    };

    return {
      tableHtml,
      submitTable,
    };
  },
});
</script>

<style lang="scss" scoped>
.align {
  display: flex;
  flex-direction: column;
  align-items: center;

  margin: 1rem 0;
}

button {
  border: none;
  color: white;
  background-color: black;
  padding: 1rem;
  font-size: 1rem;
  margin: 0.5rem 0;
  cursor: pointer;
}
</style>
