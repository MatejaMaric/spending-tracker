<template>
  <div>
    <table>
      <thead>
        <th>Made At</th>
        <th>Processed At</th>
        <th>Description</th>
        <th>Paid</th>
        <th>Received</th>
        <th>Balance</th>
      </thead>
      <tbody>
        <tr v-for="t in transactions" :key="t.ID">
          <td v-text="t.MadeAt"></td>
          <td v-text="t.ProcessedAt"></td>
          <td v-text="t.Description"></td>
          <td v-text="t.Paid"></td>
          <td v-text="t.Received"></td>
          <td v-text="t.Balance"></td>
        </tr>
      </tbody>
    </table>
  </div>
</template>

<script lang="ts">
import { defineComponent } from "vue";

export default defineComponent({
  name: "Table",
  props: {
    transactions: {
      type: Array,
      required: true,
    },
  },
});
</script>

<style scoped lang="scss">
table,
tr,
th,
td {
  border: 1px solid black;
}

table {
  max-width: 1200px;
  margin: 1rem auto;
}
</style>
