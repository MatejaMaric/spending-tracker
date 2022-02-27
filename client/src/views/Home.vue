<template>
  <div>
    <Table :transactions="transactions" />
  </div>
</template>

<script lang="ts">
import { computed, defineComponent, onMounted } from "vue";
import { useStore } from "vuex";
import Table from "@/components/Table.vue";

export default defineComponent({
  name: "Home",
  components: {
    Table,
  },
  setup() {
    const store = useStore();
    const transactions = computed(() => store.getters.getTransactions);

    onMounted(async () => {
      await store.dispatch("pullTransactions");
    });

    return {
      transactions,
    };
  },
});
</script>
