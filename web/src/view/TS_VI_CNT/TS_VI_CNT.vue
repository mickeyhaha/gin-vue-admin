<template>
  <div>
    <div class="search-term">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
        <el-form-item label="Count">
          <el-input placeholder="搜索条件" v-model="searchInfo.count"></el-input>
        </el-form-item>    
        <el-form-item label="ErrCount">
          <el-input placeholder="搜索条件" v-model="searchInfo.errCount"></el-input>
        </el-form-item>      
        <el-form-item label="LineID">
          <el-input placeholder="搜索条件" v-model="searchInfo.lineID"></el-input>
        </el-form-item>    
        <el-form-item label="OrderNo">
          <el-input placeholder="搜索条件" v-model="searchInfo.orderNo"></el-input>
        </el-form-item>    
        <el-form-item label="IssueName">
          <el-input placeholder="搜索条件" v-model="searchInfo.issueName"></el-input>
        </el-form-item>    
        <el-form-item label="Result">
          <el-input placeholder="搜索条件" v-model="searchInfo.result"></el-input>
        </el-form-item>    
        <el-form-item label="VIID">
          <el-input placeholder="搜索条件" v-model="searchInfo.vIID"></el-input>
        </el-form-item>    
        <el-form-item>
          <el-button @click="onSubmit" type="primary">查询</el-button>
        </el-form-item>
        <el-form-item>
          <el-button @click="openDialog" type="primary">新增VI不良信息</el-button>
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
    
    <el-table-column label="Count" prop="count" width="120"></el-table-column> 
    
    <el-table-column label="ErrCount" prop="errCount" width="120"></el-table-column> 
    
    <el-table-column label="CreateTime" prop="createTime" width="120"></el-table-column> 
    
    <el-table-column label="LineID" prop="lineID" width="120"></el-table-column> 
    
    <el-table-column label="OrderNo" prop="orderNo" width="120"></el-table-column> 
    
    <el-table-column label="IssueName" prop="issueName" width="120"></el-table-column> 
    
    <el-table-column label="Result" prop="result" width="120"></el-table-column> 
    
    <el-table-column label="VIID" prop="vIID" width="120"></el-table-column> 
    
      <el-table-column label="按钮组">
        <template slot-scope="scope">
          <el-button class="table-button" @click="updateTS_VI_CNT(scope.row)" size="small" type="primary" icon="el-icon-edit">变更</el-button>
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
         <el-form-item label="Count:"><el-input v-model.number="formData.count" clearable placeholder="请输入"></el-input>
      </el-form-item>
       
         <el-form-item label="ErrCount:"><el-input v-model.number="formData.errCount" clearable placeholder="请输入"></el-input>
      </el-form-item>
       
         <el-form-item label="CreateTime:">
              <el-date-picker type="date" placeholder="选择日期" v-model="formData.createTime" clearable></el-date-picker>
       </el-form-item>
       
         <el-form-item label="LineID:"><el-input v-model.number="formData.lineID" clearable placeholder="请输入"></el-input>
      </el-form-item>
       
         <el-form-item label="OrderNo:">
            <el-input v-model="formData.orderNo" clearable placeholder="请输入" ></el-input>
      </el-form-item>
       
         <el-form-item label="IssueName:">
            <el-input v-model="formData.issueName" clearable placeholder="请输入" ></el-input>
      </el-form-item>
       
         <el-form-item label="Result:"><el-input v-model.number="formData.result" clearable placeholder="请输入"></el-input>
      </el-form-item>
       
         <el-form-item label="VIID:"><el-input v-model.number="formData.vIID" clearable placeholder="请输入"></el-input>
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
    createTS_VI_CNT,
    deleteTS_VI_CNT,
    deleteTS_VI_CNTByIds,
    updateTS_VI_CNT,
    findTS_VI_CNT,
    getTS_VI_CNTList
} from "@/api/TS_VI_CNT";  //  此处请自行替换地址
import { formatTimeToStr } from "@/utils/date";
import infoList from "@/mixins/infoList";
export default {
  name: "TS_VI_CNT",
  mixins: [infoList],
  data() {
    return {
      listApi: getTS_VI_CNTList,
      dialogFormVisible: false,
      type: "",
      deleteVisible: false,
      multipleSelection: [],formData: {
            count:0,
            errCount:0,
            createTime:new Date(),
            lineID:0,
            orderNo:"",
            issueName:"",
            result:0,
            vIID:0,
            
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
           this.deleteTS_VI_CNT(row);
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
        const res = await deleteTS_VI_CNTByIds({ ids })
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
    async updateTS_VI_CNT(row) {
      const res = await findTS_VI_CNT({ ID: row.ID });
      this.type = "update";
      if (res.code == 0) {
        this.formData = res.data.reTVC;
        this.dialogFormVisible = true;
      }
    },
    closeDialog() {
      this.dialogFormVisible = false;
      this.formData = {
          count:0,
          errCount:0,
          createTime:new Date(),
          lineID:0,
          orderNo:"",
          issueName:"",
          result:0,
          vIID:0,
          
      };
    },
    async deleteTS_VI_CNT(row) {
      const res = await deleteTS_VI_CNT({ ID: row.ID });
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
          res = await createTS_VI_CNT(this.formData);
          break;
        case "update":
          res = await updateTS_VI_CNT(this.formData);
          break;
        default:
          res = await createTS_VI_CNT(this.formData);
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
