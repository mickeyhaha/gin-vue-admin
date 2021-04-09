import Vue from 'vue'
import Vuex from 'vuex'
import VuexPersistence from 'vuex-persist'

import { user } from "@/store/module/user"
import { router } from "@/store/module/router"
import { dictionary } from "@/store/module/dictionary"
Vue.use(Vuex)



const vuexLocal = new VuexPersistence({
    storage: window.localStorage,
    modules: ['user']
})
export const store = new Vuex.Store({
    state: {
        count: 0,
        selectedLine: {},
    },
    mutations: {
        update(state, [key, value]) {
            state[key] = value;
        },
    },
    modules: {
        user,
        router,
        dictionary
    },
    plugins: [vuexLocal.plugin]
})