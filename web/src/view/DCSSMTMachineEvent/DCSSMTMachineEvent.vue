<template>
  <div>
    <div class="search-term">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
        <el-form-item label="LineName字段">
          <el-input placeholder="搜索条件" v-model="searchInfo.LineName"></el-input>
        </el-form-item>        
        <el-form-item label="EventName字段">
          <el-input placeholder="搜索条件" v-model="searchInfo.EventName"></el-input>
        </el-form-item>    
        <el-form-item label="EventRemark字段">
          <el-input placeholder="搜索条件" v-model="searchInfo.EventRemark"></el-input>
        </el-form-item>      
        <el-form-item label="TableNo字段">
          <el-input placeholder="搜索条件" v-model="searchInfo.TableNo"></el-input>
        </el-form-item>        
        <el-form-item>
          <el-button @click="onSubmit" type="primary">查询</el-button>
        </el-form-item>
        <el-form-item>
          <el-button @click="openDialog" type="primary">新增DCSSMTMachineEvent表</el-button>
        </el-form-item>
        <el-form-item>
          <el-popover placement="top" v-model="deleteVisible" width="160">
            <p>确定要删除吗？</p>
              <div style="text-align: right; margin: 0">
                <el-button @click="deleteVisible = false" size="mini" type="text">取消</el-button>
                <el-button @click="onDelete" size="mini" type="primary">确定</el-button>
              </div>
            <el-button icon="el-icon-delete" size="mini" slot="reference" type="danger">批量删除</el-button>
          </el-popover>
        </el-form-item>
      </el-form>
    </div>
    <el-table
      :data="tableData"
      @selection-change="handleSelectionChange"
      border
      ref="multipleTable"
      stripe
      style="width: 100%"
      tooltip-effect="dark"
    >
    <el-table-column type="selection" width="55"></el-table-column>
    <el-table-column label="日期" width="180">
         <template slot-scope="scope">{{scope.row.CreatedAt|formatDate}}</template>
    </el-table-column>
    
    <el-table-column label="LineName字段" prop="LineName" width="120"></el-table-column> 
    
    <el-table-column label="MachineCode字段" prop="MachineCode" width="120"></el-table-column> 
    
    <el-table-column label="MachineName字段" prop="MachineName" width="120"></el-table-column> 
    
    <el-table-column label="EventName字段" prop="EventName" width="120"></el-table-column> 
    
    <el-table-column label="EventRemark字段" prop="EventRemark" width="120"></el-table-column> 
    
    <el-table-column label="CreateTime字段" prop="CreateTime" width="120"></el-table-column> 
    
    <el-table-column label="TableNo字段" prop="TableNo" width="120"></el-table-column> 
    
    <el-table-column label="Lane字段" prop="Lane" width="120"></el-table-column> 
    
    <el-table-column label="AddTime字段" prop="AddTime" width="120"></el-table-column> 
    
      <el-table-column label="按钮组">
        <template slot-scope="scope">
          <el-button class="table-button" @click="updateDCSSMTMachineEvent(scope.row)" size="small" type="primary" icon="el-icon-edit">变更</el-button>
          <el-button type="danger" icon="el-icon-delete" size="mini" @click="deleteRow(scope.row)">删除</el-button>
        </template>
      </el-table-column>
    </el-table>

    <el-pagination
      :current-page="page"
      :page-size="pageSize"
      :page-sizes="[10, 30, 50, 100]"
      :style="{float:'right',padding:'20px'}"
      :total="total"
      @current-change="handleCurrentChange"
      @size-change="handleSizeChange"
      layout="total, sizes, prev, pager, next, jumper"
    ></el-pagination>

    <el-dialog :before-close="closeDialog" :visible.sync="dialogFormVisible" title="弹窗操作">
      <el-form :model="formData" label-position="right" label-width="80px">
         <el-form-item label="LineName字段:">
            <el-input v-model="formData.LineName" clearable placeholder="请输入" ></el-input>
      </el-form-item>
       
         <el-form-item label="MachineCode字段:"><el-input v-model.number="formData.MachineCode" clearable placeholder="请输入"></el-input>
      </el-form-item>
       
         <el-form-item label="MachineName字段:">
            <el-input v-model="formData.MachineName" clearable placeholder="请输入" ></el-input>
      </el-form-item>
       
         <el-form-item label="EventName字段:">
            <el-input v-model="formData.EventName" clearable placeholder="请输入" ></el-input>
      </el-form-item>
       
         <el-form-item label="EventRemark字段:">
            <el-input v-model="formData.EventRemark" clearable placeholder="请输入" ></el-input>
      </el-form-item>
       
         <el-form-item label="CreateTime字段:">
              <el-date-picker type="date" placeholder="选择日期" v-model="formData.CreateTime" clearable></el-date-picker>
       </el-form-item>
       
         <el-form-item label="TableNo字段:">
            <el-input v-model="formData.TableNo" clearable placeholder="请输入" ></el-input>
      </el-form-item>
       
         <el-form-item label="Lane字段:"><el-input v-model.number="formData.Lane" clearable placeholder="请输入"></el-input>
      </el-form-item>
       
         <el-form-item label="AddTime字段:">
              <el-date-picker type="date" placeholder="选择日期" v-model="formData.AddTime" clearable></el-date-picker>
       </el-form-item>
       </el-form>
      <div class="dialog-footer" slot="footer">
        <el-button @click="closeDialog">取 消</el-button>
        <el-button @click="enterDialog" type="primary">确 定</el-button>
      </div>
    </el-dialog>
  </div>
</template>

<script>
import {
    createDCSSMTMachineEvent,
    deleteDCSSMTMachineEvent,
    deleteDCSSMTMachineEventByIds,
    updateDCSSMTMachineEvent,
    findDCSSMTMachineEvent,
    getDCSSMTMachineEventList
} from "@/api/DCSSMTMachineEvent";  //  此处请自行替换地址
import { formatTimeToStr } from "@/utils/date";
import infoList from "@/mixins/infoList";
export default {
  name: "DCSSMTMachineEvent",
  mixins: [infoList],
  data() {
    return {
      listApi: getDCSSMTMachineEventList,
      dialogFormVisible: false,
      type: "",
      deleteVisible: false,
      multipleSelection: [],formData: {
            LineName:"",
            MachineCode:0,
            MachineName:"",
            EventName:"",
            EventRemark:"",
            CreateTime:new Date(),
            TableNo:"",
            Lane:0,
            AddTime:new Date(),
            
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
        this.pageSize = 10             
        this.getTableData()
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
           this.deleteDCSSMTMachineEvent(row);
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
        const res = await deleteDCSSMTMachineEventByIds({ ids })
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
    async updateDCSSMTMachineEvent(row) {
      const res = await findDCSSMTMachineEvent({ ID: row.ID });
      this.type = "update";
      if (res.code == 0) {
        this.formData = res.data.reDSME;
        this.dialogFormVisible = true;
      }
    },
    closeDialog() {
      this.dialogFormVisible = false;
      this.formData = {
          LineName:"",
          MachineCode:0,
          MachineName:"",
          EventName:"",
          EventRemark:"",
          CreateTime:new Date(),
          TableNo:"",
          Lane:0,
          AddTime:new Date(),
          
      };
    },
    async deleteDCSSMTMachineEvent(row) {
      const res = await deleteDCSSMTMachineEvent({ ID: row.ID });
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
          res = await createDCSSMTMachineEvent(this.formData);
          break;
        case "update":
          res = await updateDCSSMTMachineEvent(this.formData);
          break;
        default:
          res = await createDCSSMTMachineEvent(this.formData);
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
    await this.getTableData();
  
}
};
</script>

<style>
</style>
