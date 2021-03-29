import service from '@/utils/request'

// @Tags Machine
// @Summary 创建Machine
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Machine true "创建Machine"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /machine/createMachine [post]
export const createMachine = (data) => {
     return service({
         url: "/machine/createMachine",
         method: 'post',
         data
     })
 }


// @Tags Machine
// @Summary 删除Machine
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Machine true "删除Machine"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /machine/deleteMachine [delete]
 export const deleteMachine = (data) => {
     return service({
         url: "/machine/deleteMachine",
         method: 'delete',
         data
     })
 }

// @Tags Machine
// @Summary 删除Machine
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Machine"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /machine/deleteMachine [delete]
 export const deleteMachineByIds = (data) => {
     return service({
         url: "/machine/deleteMachineByIds",
         method: 'delete',
         data
     })
 }

// @Tags Machine
// @Summary 更新Machine
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Machine true "更新Machine"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /machine/updateMachine [put]
 export const updateMachine = (data) => {
     return service({
         url: "/machine/updateMachine",
         method: 'put',
         data
     })
 }


// @Tags Machine
// @Summary 用id查询Machine
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Machine true "用id查询Machine"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /machine/findMachine [get]
 export const findMachine = (params) => {
     return service({
         url: "/machine/findMachine",
         method: 'get',
         params
     })
 }


// @Tags Machine
// @Summary 分页获取Machine列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取Machine列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /machine/getMachineList [get]
 export const getMachineList = (params) => {
     return service({
         url: "/machine/getMachineList",
         method: 'get',
         params
     })
 }

export function randomExtend(minNum, maxNum) {
    return parseInt(Math.random() * (maxNum - minNum + 1) + minNum, 10)
}

export const getFlopData = () => {
     return [
         {
             title: '管养里程',
             number: {
                 number: [randomExtend(20000, 30000)],
                 content: '{nt}',
                 textAlign: 'right',
                 style: {
                     fill: '#4d99fc',
                     fontWeight: 'bold'
                 }
             },
             unit: '公里'
         },
         {
             title: '桥梁',
             number: {
                 number: [randomExtend(20, 30)],
                 content: '{nt}',
                 textAlign: 'right',
                 style: {
                     fill: '#f46827',
                     fontWeight: 'bold'
                 }
             },
             unit: '座'
         },
         {
             title: '涵洞隧道',
             number: {
                 number: [randomExtend(20, 30)],
                 content: '{nt}',
                 textAlign: 'right',
                 style: {
                     fill: '#40faee',
                     fontWeight: 'bold'
                 }
             },
             unit: '个'
         },
         {
             title: '匝道',
             number: {
                 number: [randomExtend(10, 20)],
                 content: '{nt}',
                 textAlign: 'right',
                 style: {
                     fill: '#4d99fc',
                     fontWeight: 'bold'
                 }
             },
             unit: '个'
         },
         {
             title: '隧道',
             number: {
                 number: [randomExtend(5, 10)],
                 content: '{nt}',
                 textAlign: 'right',
                 style: {
                     fill: '#f46827',
                     fontWeight: 'bold'
                 }
             },
             unit: '个'
         },
         {
             title: '服务区',
             number: {
                 number: [randomExtend(5, 10)],
                 content: '{nt}',
                 textAlign: 'right',
                 style: {
                     fill: '#40faee',
                     fontWeight: 'bold'
                 }
             },
             unit: '个'
         },
         {
             title: '收费站',
             number: {
                 number: [randomExtend(5, 10)],
                 content: '{nt}',
                 textAlign: 'right',
                 style: {
                     fill: '#4d99fc',
                     fontWeight: 'bold'
                 }
             },
             unit: '个'
         },
         {
             title: '超限站',
             number: {
                 number: [randomExtend(5, 10)],
                 content: '{nt}',
                 textAlign: 'right',
                 style: {
                     fill: '#f46827',
                     fontWeight: 'bold'
                 }
             },
             unit: '个'
         },
         {
             title: '停车区',
             number: {
                 number: [randomExtend(5, 10)],
                 content: '{nt}',
                 textAlign: 'right',
                 style: {
                     fill: '#40faee',
                     fontWeight: 'bold'
                 }
             },
             unit: '个'
         }
     ]
 }