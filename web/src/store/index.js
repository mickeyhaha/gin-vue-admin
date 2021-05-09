import Vue from 'vue'
import Vuex from 'vuex'
import VuexPersistence from 'vuex-persist'

import { user } from "@/store/module/user"
import { router } from "@/store/module/router"
import { dictionary } from "@/store/module/dictionary"
Vue.use(Vuex)

import { 
    getPVS_Base_LineList,
    getDeptLineSummary,
    getDCSSMTOutPutList4Chart,
    getPUBMOrderProduce2InfoList4Chart,
    getDCSSMTConsumeAndRejectRate4Chart,
    getDCSSMTMachineEvent4Chart,
    getDCSSMTRunTime4Chart,
} from '@/api/PVS_Base_Line'
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
            // categories: [],
            // series: [
            // ],
            categories: ['4/18', '4/19', '4/20', '4/21', '4/22'],
            series: [
                {
                    name: '不良',
                    data: [320, 420, 210, 110, 620],
                },
            ],
        },
        rejectRate4Chart: {
            categories: ['4/18', '4/19', '4/20', '4/21', '4/22'],
            series: [
                {
                    name: '抛料率',
                    data: [0.3, 0.4, 0.2, 0.1, 0.5],
                },
            ],
        },
        dateOutput4Chart: {
            categories: ['4/18', '4/19', '4/20', '4/21', '4/22'],
            series: [
                {
                    name: '实际产量',
                    data: [400, 400, 200, 100, 600],
                },
                {
                    name: '标准产量',
                    data: [320, 420, 210, 110, 620],
                },
            ],
        },
        dateMachineEvent4Chart: {
            categories: ['4/18', '4/19', '4/20', '4/21', '4/22'],
            series: [
                {
                    name: '卡料',
                    data: [40, 40, 20, 10, 60],
                },
                {
                    name: '报警',
                    data: [32, 42, 21, 11, 62],
                },
            ],
        },
        dateRunTime4Chart: {
            categories: ['4/18', '4/19', '4/20', '4/21', '4/22'],
            series: [
                {
                    name: '异常停机',
                    data: [40, 40, 20, 10, 60],
                },
                {
                    name: '空闲停机',
                    data: [32, 42, 21, 11, 62],
                },
            ],
        },
        stopReason4Chart: {
            categories: ['停机分布'],
            series: [
                {
                    name: '异常停机',
                    data: 40,
                },
                {
                    name: '空闲停机',
                    data: 32,
                },
                {
                    name: '错误停机',
                    data: 32,
                },
                {
                    name: '卡料停机',
                    data: 32,
                },
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
                return state.lines
            }
        },
        async getAoiRate4Chart({ commit, state }, formData) {
            const res = await getTS_AOI_CNTList4Chart({
                page: 1, pageSize: 100,
                lineName: formData.LineName,
                startDate: formData.date[0],
                endDate: formData.date[1],
                shift: formData.Shift,
                workOrderNo: formData.WorkOrderNo
            });
            if (res.code == 0) {
                const aoiRate4Chart = res.data.list[0]
                commit("setAoiRate4Chart", aoiRate4Chart)
                return state.aoiRate4Chart
            }
        },

        // 抛料率
        async getRejectRate4Chart({ commit, state }, formData) {
            const res = await getDCSSMTConsumeAndRejectRate4Chart({ 
                page: 1, pageSize: 100,
                LineName: formData.LineName,
                startDate: formData.date[0],
                endDate: formData.date[1],
                shift: formData.Shift,
                workOrderNo: formData.WorkOrderNo
            });
            if (res.code == 0) {
                const rejectRate = res.data.list[0]
                commit("setRejectRate4Chart", rejectRate)
                return state.rejectRate
            }
        }, 

        async getDCSSMTMachineEvent4Chart({ commit, state }, formData) {
            const res = await getDCSSMTMachineEvent4Chart({
                page: 1, pageSize: 100,
                LineName: formData.LineName,
                startDate: formData.date[0],
                endDate: formData.date[1],
                shift: formData.Shift,
                workOrderNo: formData.WorkOrderNo
            });
            if (res.code == 0) {
                const chartData = res.data.list[0]
                commit("setDateMachineEvent4Chart", chartData)
                return state.dateMachineEvent4Chart
            }
        }, 

        async getDCSSMTRunTime4Chart({ commit, state }, formData) {
            const res = await getDCSSMTRunTime4Chart({
                page: 1, pageSize: 100,
                LineName: formData.LineName,
                startDate: formData.date[0],
                endDate: formData.date[1],
                shift: formData.Shift,
                workOrderNo: formData.WorkOrderNo
            });
            if (res.code == 0) {
                const chartData = res.data.list[0]
                commit("setDateRunTime4Chart", chartData)
                return state.dateRunTime4Chart
            }
        }, 
        
        // 获取产量1
        async getDCSSMTOutPutList4Chart({ commit, state }, formData) {
            const res = await getDCSSMTOutPutList4Chart({
                page: 1, pageSize: 100,
                lineName: formData.LineName,
                startDate: formData.date[0],
                endDate: formData.date[1],
                shift: formData.Shift,
                workOrderNo: formData.WorkOrderNo
            });
            if (res.code == 0) {
                const dateOutput4Chart = res.data.list[0]
                commit("setDateOutput4Chart", dateOutput4Chart)
                return state.dateOutput4Chart
            }
        }, 

        // 获取产量2
        async getPUBMOrderProduce2InfoList4Chart({ commit, state }, formData) {
            const res = await getPUBMOrderProduce2InfoList4Chart({
                page: 1, pageSize: 100,
                LineName: formData.LineName,
                startDate: formData.date[0],
                endDate: formData.date[1],
                shift: formData.Shift,
                workOrderNo: formData.WorkOrderNo
            });
            if (res.code == 0) {
                const dateOutput4Chart = res.data.list[0]
                commit("setDateOutput4Chart", dateOutput4Chart)
                return state.dateOutput4Chart
            }
        }, 

        // 点击deptFilter的submit
        async submitDeptFilter({ commit, dispatch }, formData) {
            if (formData.Shift == 1) {
                formData.date[0] += " 08:00:00"
                formData.date[1] += " 18:59:59"
            } else if (formData.Shift == 2) {
                formData.date[0] += " 19:00:00"
                formData.date[1] += " 07:59:59"
            }
            dispatch('getAoiRate4Chart', formData)
            dispatch('getPUBMOrderProduce2InfoList4Chart', formData)
            dispatch('getRejectRate4Chart', formData)
            dispatch('getDCSSMTMachineEvent4Chart', formData)
            dispatch('getDCSSMTRunTime4Chart', formData)
        },

        // 点击deptLineSummary的一行
        async selectLine({ commit, dispatch}, selectedLine) {
            return
            commit('setSelectedLine', selectedLine)
            dispatch('getAoiRate4Chart', {lineName: selectedLine.LineName})
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
        setRejectRate4Chart(state, rejectRate4Chart) {
            state.rejectRate4Chart = rejectRate4Chart;
        }, 
        setDateMachineEvent4Chart(state, dateMachineEvent4Chart) {
            state.dateMachineEvent4Chart = dateMachineEvent4Chart;
        },
        setDateRunTime4Chart(state, dateRunTime4Chart) {
            state.dateRunTime4Chart = dateRunTime4Chart;
        },
        setDateOutput4Chart(state, dateOutput4Chart) {
            state.dateOutput4Chart = dateOutput4Chart;
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