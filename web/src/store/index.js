import Vue from 'vue'
import Vuex from 'vuex'
import VuexPersistence from 'vuex-persist'

import { user } from "@/store/module/user"
import { router } from "@/store/module/router"
import { dictionary } from "@/store/module/dictionary"
Vue.use(Vuex)

import { getPVS_Base_LineList } from '@/api/PVS_Base_Line'
import {
    getLackWarnings,
    getRejectRateList,
} from "@/api/MoniWholeView";
import {
    getTS_AOI_CNTList4Chart,
} from "@/api/TS_AOI_CNT";


const vuexLocal = new VuexPersistence({
    storage: window.localStorage,
    modules: ['user']
})
export const store = new Vuex.Store({
    state: {
        count: 0,
        lines: [],
        selectedLine: {},
        rejectRate: [],
        aoiRate4Chart: {
            categories: ['Line1', 'Line2', 'Line3', 'Line4', 'Line5'],
            series: [
                {
                    name: '模糊',
                    data: [0.3, 0.4, 0.2, 0.1, 0.5],
                    stackGroup: 'AOI',
                },
                {
                    name: '少料',
                    data: [0.2, 0.3, 0.3, 0.2, 0.5],
                    stackGroup: 'AOI',
                },
                {
                    name: '错点',
                    data: [0.4, 0.2, 0.1, 0.3, 0.1],
                    stackGroup: 'AOI',
                },
            ],
        },
    },
    actions: {
        async getLines({ commit, state }) {
            const res = await getPVS_Base_LineList({ page: 1, pageSize: 100 });
            if (res.code == 0) {
                const lines = res.data.list
                commit("setLines", lines)
                return state.lines
            }
        },
        async getAoiRate4Chart({ commit, state }, lineName) {
            const res = await getTS_AOI_CNTList4Chart({ page: 1, pageSize: 100 });
            if (res.code == 0) {
                const aoiRate4Chart = res.data.list[0]
                commit("setAoiRate4Chart", aoiRate4Chart)
                console.log(aoiRate4Chart)
                return state.aoiRate4Chart
            }
        },
        async getRejectRate({ commit, state }) {
            const res = await getRejectRateList({ page: 1, pageSize: 100 });
            if (res.code == 0) {
                const rejectRate = res.data.list
                commit("setRejectRate", rejectRate)
                return state.rejectRate
            }
        }, 

        async setSelectedLine({ commit, dispatch}, selectedLine) {
            commit('setSelectedLine', selectedLine)
            dispatch('getAoiRate4Chart')
        },
    },
    // 无法执行异步
    mutations: {
        update(state, [key, value]) {
            state[key] = value;
        },
        setLines(state, lines) {
            state.lines = lines;
        },
        setSelectedLine(state, selectedLine) {
            state.selectedLine = selectedLine;
        },
        setRejectRate(state, rejectRate) {
            state.rejectRate = rejectRate;
        },
        setAoiRate4Chart(state, aoiRate4Chart) {
            state.aoiRate4Chart = aoiRate4Chart;
        },
    },
    getters: {
        getLines(state) {
            return state.lines
        },
        getAoiRate4Chart(state) {
            return state.aoiRate4Chart
        },
        getSelectedLine(state) {
            return state.selectedLine
        },
    },
    modules: {
        user,
        router,
        dictionary,
    },
    plugins: [vuexLocal.plugin]
})