<template>
  <div>
    <el-table
      :data="tableData"
      @selection-change="handleSelectionChange"
      @row-click="clickData"
      border
      ref="multipleTable"
      stripe
      style="width: 100%; height: 200px"
      tooltip-effect="dark"
    >
    <el-table-column label="日期" width="180">
         <template slot-scope="scope">{{scope.row.CreatedAt|formatDate}}</template>
    </el-table-column>
    
    <el-table-column label="LineID" prop="lineID" width="120"></el-table-column> 
    
    <el-table-column label="LineName" prop="lineName" width="120"></el-table-column> 
    
    <el-table-column label="LineType" prop="lineType" width="120"></el-table-column> 
    
    <el-table-column label="程式读取路径" prop="readFilePath" width="120"></el-table-column> 
    
    <el-table-column label="LineName2" prop="lineName2" width="120"></el-table-column> 
    
    <el-table-column label="制令单号" prop="mOrderNo" width="120"></el-table-column> 
    
    <el-table-column label="Remark" prop="remark" width="120"></el-table-column> 
    
    <el-table-column label="是否报警" prop="warnning" width="120">
         <template slot-scope="scope">{{scope.row.warnning|formatBoolean}}</template>
    </el-table-column>
    
    <el-table-column label="WarnningStr" prop="warnningStr" width="120"></el-table-column> 
    
      <el-table-column label="按钮组">
        <template slot-scope="scope">
          <el-button class="table-button" @click="updatePVS_Base_Line(scope.row)" size="small" type="primary" icon="el-icon-edit">变更</el-button>
        </template>
      </el-table-column>
    </el-table>
  </div>
</template>

<script>
import {
    createPVS_Base_Line,
    deletePVS_Base_Line,
    deletePVS_Base_LineByIds,
    updatePVS_Base_Line,
    findPVS_Base_Line,
    getPVS_Base_LineList
} from "@/api/PVS_Base_Line";  //  此处请自行替换地址
import { formatTimeToStr } from "@/utils/date";
import infoList from "@/mixins/infoList";
import { getLines } from "@/utils/datav";

export default {
  name: "PVS_Base_Line",
  mixins: [infoList],
  data() {
    return {
      listApi: getPVS_Base_LineList,
      dialogFormVisible: false,
      type: "",
      deleteVisible: false,
      multipleSelection: [],formData: {
            lineID:0,
            lineName:"",
            lineType:"",
            readFilePath:"",
            lineName2:"",
            mOrderNo:"",
            remark:"",
            warnning:false,
            warnningStr:"",
            
      }
    };
  },
  filters: {
    formatDate: function(time) {
      if (time != null && time != "") {
        var date = new Date(time);
        return formatTimeToStr(date, "yyyy-MM-dd hh:mm:ss");
      } else {
        return "";
      }
    },
    formatBoolean: function(bool) {
      if (bool != null) {
        return bool ? "是" :"否";
      } else {
        return "";
      }
    }
  },
  methods: {
      //条件搜索前端看此方法
      onSubmit() {
        this.page = 1
        this.pageSize = 100            
        if (this.searchInfo.warnning==""){
          this.searchInfo.warnning=null
        }       
        this.getTableData()
      },
      clickData(val) {
        console.log(val)
      },
      handleSelectionChange(val) {
        this.multipleSelection = val
      },
      deleteRow(row){
        this.$confirm('确定要删除吗?', '提示', {
          confirmButtonText: '确定',
          cancelButtonText: '取消',
          type: 'warning'
        }).then(() => {
           this.deletePVS_Base_Line(row);
        });
      },
      async onDelete() {
        const ids = []
        if(this.multipleSelection.length == 0){
          this.$message({
            type: 'warning',
            message: '请选择要删除的数据'
          })
          return
        }
        this.multipleSelection &&
          this.multipleSelection.map(item => {
            ids.push(item.ID)
          })
        const res = await deletePVS_Base_LineByIds({ ids })
        if (res.code == 0) {
          this.$message({
            type: 'success',
            message: '删除成功'
          })
          if (this.tableData.length == ids.length) {
              this.page--;
          }
          this.deleteVisible = false
          this.getTableData()
        }
      },
    async updatePVS_Base_Line(row) {
      const res = await findPVS_Base_Line({ ID: row.ID });
      this.type = "update";
      if (res.code == 0) {
        this.formData = res.data.rePBL;
        this.dialogFormVisible = true;
      }
    },
    closeDialog() {
      this.dialogFormVisible = false;
      this.formData = {
          lineID:0,
          lineName:"",
          lineType:"",
          readFilePath:"",
          lineName2:"",
          mOrderNo:"",
          remark:"",
          warnning:false,
          warnningStr:"",
          
      };
    },
    async deletePVS_Base_Line(row) {
      const res = await deletePVS_Base_Line({ ID: row.ID });
      if (res.code == 0) {
        this.$message({
          type: "success",
          message: "删除成功"
        });
        if (this.tableData.length == 1) {
            this.page--;
        }
        this.getTableData();
      }
    },
    async enterDialog() {
      let res;
      switch (this.type) {
        case "create":
          res = await createPVS_Base_Line(this.formData);
          break;
        case "update":
          res = await updatePVS_Base_Line(this.formData);
          break;
        default:
          res = await createPVS_Base_Line(this.formData);
          break;
      }
      if (res.code == 0) {
        this.$message({
          type:"success",
          message:"创建/更改成功"
        })
        this.closeDialog();
        this.getTableData();
      }
    },
    openDialog() {
      this.type = "create";
      this.dialogFormVisible = true;
    }
  },
  async created() {
    this.tableData = await getLines();

    // this.$refs["multipleTable"].$nextTick(function () {
    //             this.escrollY(
    //               ".el-table__body-wrapper",
    //               ".el-table__header-wrapper"
    //             );
    // })
  }
};
</script>

<style>
</style>
