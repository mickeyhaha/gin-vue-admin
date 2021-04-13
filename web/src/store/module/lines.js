import { getPVS_Base_LineList } from '@/api/PVS_Base_Line'

export const lines = {
    namespaced: true,
    state: {
        lines: [],
    },
    mutations: {
        setLines(state, lines) {
            // state.lines = { ...state.lines, ...lines }
            state.lines = lines;
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
    },
    getters:{
        getLines(state){
            return state.lines
        }
    }
}