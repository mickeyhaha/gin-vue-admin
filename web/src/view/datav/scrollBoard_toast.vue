<template>
  <div id="scroll-board">
    <div class="title">停机原因</div>
    <el-table
      :data="config.data"
      style="width: 100%;"
    >
      <el-table-column :label="config.header[0]" prop="machine"></el-table-column> 
      <el-table-column :label="config.header[1]" prop="type"></el-table-column> 
      <el-table-column :label="config.header[2]" prop="position"></el-table-column> 
      <el-table-column :label="config.header[3]" prop="reason"></el-table-column> 
      <el-table-column :label="config.header[4]" prop="port"></el-table-column> 
    </el-table>
  </div>
</template>

<script>
export default {
  name: 'ScrollBoard',
  data () {
    return {
      config: {
        header: ['机器', '停机类别', '位置', '停机原因', '停机端口'],
        data: [
          {machine: '机器1', type: '卡纸', position: '1-1', reason: '太厚', port: '1'},
          {machine: '机器1', type: '卡纸', position: '1-1', reason: '太厚', port: '1'},
          {machine: '机器1', type: '卡纸', position: '1-1', reason: '太厚', port: '1'},
          {machine: '机器1', type: '卡纸', position: '1-1', reason: '太厚', port: '1'},
          {machine: '机器1', type: '卡纸', position: '1-1', reason: '太厚', port: '1'},
          {machine: '机器1', type: '卡纸', position: '1-1', reason: '太厚', port: '1'},
        ],
        index: true,
        columnWidth: [50, 170, 300],
        align: ['center'],
        rowNum: 7,
        headerBGC: '#1981f6',
        headerHeight: 45,
        oddRowBGC: 'rgba(0, 44, 81, 0.8)',
        evenRowBGC: 'rgba(10, 29, 50, 0.8)'
      }
    }
  }
}
</script>

<style lang="less" scoped>
/deep/ .el-table{
/* 表格字体颜色 */
color:white;
/* 表格边框颜色 */
border: 0.5px solid #758a99; 
height: 100%;
}
/* 表格内背景颜色 */
/deep/ .el-table th {
border: 1;
color: white;
}
/deep/ .el-table tr{
border: 1;
color: white;
background-color: #5ca8f2;
}
/deep/ .el-table td{
border: 1;
color: white;
background-color: #08183d;
}
/* 双数行背景颜色 */
/deep/ .el-table--striped .el-table__body tr.el-table__row--striped td {
 
background-color:#fff;
background-color: rgba(148, 144, 144, 0.3)
}
/* 使表格背景透明 */
/deep/ .el-table th, .el-table tr {
border: 1;
background-color: transparent;
}
/* 删除表格下横线 */
/deep/ .el-table::before {
left: 0;
bottom: 0;
width: 100%;
height: 0px;
}
/* 表格表头字体颜色 */
/deep/ .el-table thead {
color: white;
font-weight: 500;
background-color: rgba(148, 144, 144, 0.3)
}
</style>

<style lang="less">
#scroll-board {
  width: 100%;
  height: 200px;
  box-shadow: 0 0 3px blue;
  display: flex;
  flex-direction: column;
  background-color: rgba(6, 30, 93, 0.5);
  border-top: 2px solid rgba(1, 153, 209, .5);
  box-sizing: border-box;
  padding: 0px 10px;
  margin-bottom: 10px;
}
</style>
