<template>
  <div>
    <div class="search-term">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
        <el-form-item label="目标产量">
          <el-input placeholder="搜索条件" v-model="searchInfo.dest"></el-input>
        </el-form-item>      
        <el-form-item label="小时数">
          <el-input placeholder="搜索条件" v-model="searchInfo.hour"></el-input>
        </el-form-item>    
        <el-form-item label="开始时间">
          <el-input placeholder="搜索条件" v-model="searchInfo.beginTime"></el-input>
        </el-form-item>    
        <el-form-item label="LineID">
          <el-input placeholder="搜索条件" v-model="searchInfo.lineID"></el-input>
        </el-form-item>    
            <el-form-item label="制令单目标数量" prop="manuDest">
            <el-select v-model="searchInfo.manuDest" clear placeholder="请选择">
                <el-option
                    key="true"
                    label="是"
                    value="true">
                </el-option>
                <el-option
                    key="false"
                    label="否"
                    value="false">
                </el-option>
            </el-select>
            </el-form-item>   
        <el-form-item>
          <el-button @click="onSubmit" type="primary">查询</el-button>
        </el-form-item>
        <el-form-item>
          <el-button @click="openDialog" type="primary">新增时段产量</el-button>
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
    
    <el-table-column label="目标产量" prop="dest" width="120"></el-table-column> 
    
    <el-table-column label="实际产量逗号隔开" prop="real" width="120"></el-table-column> 
    
    <el-table-column label="小时数" prop="hour" width="120"></el-table-column> 
    
    <el-table-column label="开始时间" prop="beginTime" width="120"></el-table-column> 
    
    <el-table-column label="LineID" prop="lineID" width="120"></el-table-column> 
    
    <el-table-column label="制令单目标数量" prop="manuDest" width="120">
         <template slot-scope="scope">{{scope.row.manuDest|formatBoolean}}</template>
    </el-table-column>
    
      <el-table-column label="按钮组">
        <template slot-scope="scope">
          <el-button class="table-button" @click="updateVM_TimeOutput(scope.row)" size="small" type="primary" icon="el-icon-edit">变更</el-button>
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
         <el-form-item label="目标产量:"><el-input v-model.number="formData.dest" clearable placeholder="请输入"></el-input>
      </el-form-item>
       
         <el-form-item label="实际产量逗号隔开:">
            <el-input v-model="formData.real" clearable placeholder="请输入" ></el-input>
      </el-form-item>
       
         <el-form-item label="小时数:"><el-input v-model.number="formData.hour" clearable placeholder="请输入"></el-input>
      </el-form-item>
       
         <el-form-item label="开始时间:">
              <el-date-picker type="date" placeholder="选择日期" v-model="formData.beginTime" clearable></el-date-picker>
       </el-form-item>
       
         <el-form-item label="LineID:"><el-input v-model.number="formData.lineID" clearable placeholder="请输入"></el-input>
      </el-form-item>
       
         <el-form-item label="制令单目标数量:">
            <el-switch active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" v-model="formData.manuDest" clearable ></el-switch>
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
    createVM_TimeOutput,
    deleteVM_TimeOutput,
    deleteVM_TimeOutputByIds,
    updateVM_TimeOutput,
    findVM_TimeOutput,
    getVM_TimeOutputList
} from "@/api/VM_TimeOutput";  //  此处请自行替换地址
import { formatTimeToStr } from "@/utils/date";
import infoList from "@/mixins/infoList";
export default {
  name: "VM_TimeOutput",
  mixins: [infoList],
  data() {
    return {
      listApi: getVM_TimeOutputList,
      dialogFormVisible: false,
      type: "",
      deleteVisible: false,
      multipleSelection: [],formData: {
            dest:0,
            real:"",
            hour:0,
            beginTime:new Date(),
            lineID:0,
            manuDest:false,
            
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
        if (this.searchInfo.manuDest==""){
          this.searchInfo.manuDest=null
        }      
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
           this.deleteVM_TimeOutput(row);
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
        const res = await deleteVM_TimeOutputByIds({ ids })
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
    async updateVM_TimeOutput(row) {
      const res = await findVM_TimeOutput({ ID: row.ID });
      this.type = "update";
      if (res.code == 0) {
        this.formData = res.data.reVTO;
        this.dialogFormVisible = true;
      }
    },
    closeDialog() {
      this.dialogFormVisible = false;
      this.formData = {
          dest:0,
          real:"",
          hour:0,
          beginTime:new Date(),
          lineID:0,
          manuDest:false,
          
      };
    },
    async deleteVM_TimeOutput(row) {
      const res = await deleteVM_TimeOutput({ ID: row.ID });
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
          res = await createVM_TimeOutput(this.formData);
          break;
        case "update":
          res = await updateVM_TimeOutput(this.formData);
          break;
        default:
          res = await createVM_TimeOutput(this.formData);
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
