import { createStore } from "vuex";

export default createStore({
  state: {
    transactions: [],
  },
  mutations: {
    setTransactions(state, payload) {
      state.transactions = payload.data;
    },
  },
  actions: {
    pullTransactions(context) {
      return fetch("http://localhost:3000/api/transaction")
        .then((res) => res.json())
        .then((res) => context.commit("setTransactions", res));
    },
    processTransactions(context, payload) {
      return fetch("http://localhost:3000/api/transaction", {
        method: "POST",
        headers: {
          "Content-Type": "text/html",
        },
        body: payload,
      }).then((res) => res.json());
    },
  },
  getters: {
    getTransactions(state) {
      return state.transactions;
    },
  },
  modules: {},
});
