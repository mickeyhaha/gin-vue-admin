import Vue from 'vue'
import Vuex from 'vuex'
import VuexPersistence from 'vuex-persist'

import { user } from "@/store/module/user"
import { router } from "@/store/module/router"
import { dictionary } from "@/store/module/dictionary"
Vue.use(Vuex)

import { getPVS_Base_LineList,
    getDeptLineSummary } from '@/api/PVS_Base_Line'
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
            categories: [],
            series: [
            ],
        },
        deptLineSummary: [],
    },
    actions: {
        async getLines({ commit, state }) {
            const res = await getPVS_Base_LineList({ page: 1, pageSize: 100 });
            if (res.code == 0) {
                const lines = res.data.list
                commit("setLines", lines)
                console.log(lines)
                return state.lines
            }
        },
        async getAoiRate4Chart({ commit, state }, lineName) {
            const res = await getTS_AOI_CNTList4Chart({ page: 1, pageSize: 100, lineName: lineName });
            if (res.code == 0) {
                const aoiRate4Chart = res.data.list[0]
                commit("setAoiRate4Chart", aoiRate4Chart)
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

        // 点击dept看板的线体
        async selectLine({ commit, dispatch}, selectedLine) {
            commit('setSelectedLine', selectedLine)
            dispatch('getAoiRate4Chart', selectedLine.LineName)
        },

        async getDeptLineSummary({ commit, state }) {
            const res = await getDeptLineSummary({ page: 1, pageSize: 100 });
            if (res.code == 0) {
                const deptLineSummary = res.data.list
                commit("setDeptLineSummary", deptLineSummary)
                return state.deptLineSummary
            }
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
        setDeptLineSummary(state, deptLineSummary) {
            state.deptLineSummary = deptLineSummary;
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
        getDeptLineSummary(state) {
            return state.deptLineSummary
        },
    },
    modules: {
        user,
        router,
        dictionary,
    },
    plugins: [vuexLocal.plugin]
})