import { createStore } from "vuex";

const store = createStore({
  state: {
    /* User */
    userName: null,
    dbName: null,
    tableCount: 0,
    tables: [],
    selectedTable: null,
  },
  getters: {
    getDatabase(state) {
      return state.dbName;
    },
    getTableCount(state) {
      return state.tableCount;
    },
  },
  mutations: {
    setDatabase(state, value) {
      state.dbName = value;
    },
    setTable(state, value) {
      state.selectedTable = value;
    },
    setTableCount(state, value) {
      state.tableCount = value;
    },
    setTables(state, value) {
      state.tables = value;
    },
    /* User */
    user(state, payload) {
      if (payload.name) {
        state.userName = payload.name;
      }
    },
  },
  actions: {
    // asideMobileToggle ({ commit, state }, payload = null) {
    //   const isShow = payload !== null ? payload : !state.isAsideMobileExpanded
    //   document.getElementById('app').classList[isShow ? 'add' : 'remove']('ml-60')
    //   document.documentElement.classList[isShow ? 'add' : 'remove']('m-clipped')
    //   commit('basic', {
    //     key: 'isAsideMobileExpanded',
    //     value: isShow
    //   })
    // },
    // formScreenToggle ({ commit, state }, payload) {
    //   commit('basic', {
    //     key: 'isFormScreen',
    //     value: payload
    //   })
    //   document.documentElement.classList[payload ? 'add' : 'remove']('form-screen')
    // }
  },
  modules: {},
});

export default store;
