<template>
  <div>
    <div class="search-term">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
        <el-form-item label="LineID">
          <el-input placeholder="搜索条件" v-model="searchInfo.lineID"></el-input>
        </el-form-item>    
        <el-form-item label="LineName">
          <el-input placeholder="搜索条件" v-model="searchInfo.lineName"></el-input>
        </el-form-item>    
        <el-form-item label="LineType">
          <el-input placeholder="搜索条件" v-model="searchInfo.lineType"></el-input>
        </el-form-item>    
        <el-form-item label="程式读取路径">
          <el-input placeholder="搜索条件" v-model="searchInfo.readFilePath"></el-input>
        </el-form-item>    
        <el-form-item label="LineName2">
          <el-input placeholder="搜索条件" v-model="searchInfo.lineName2"></el-input>
        </el-form-item>    
        <el-form-item label="制令单号">
          <el-input placeholder="搜索条件" v-model="searchInfo.mOrderNo"></el-input>
        </el-form-item>    
        <el-form-item label="Remark">
          <el-input placeholder="搜索条件" v-model="searchInfo.remark"></el-input>
        </el-form-item>    
            <el-form-item label="是否报警" prop="warnning">
            <el-select v-model="searchInfo.warnning" clear placeholder="请选择">
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
        <el-form-item label="WarnningStr">
          <el-input placeholder="搜索条件" v-model="searchInfo.warnningStr"></el-input>
        </el-form-item>    
        <el-form-item>
          <el-button @click="onSubmit" type="primary">查询</el-button>
        </el-form-item>
        <el-form-item>
          <el-button @click="openDialog" type="primary">新增PVS生产线表</el-button>
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
         <el-form-item label="LineID:"><el-input v-model.number="formData.lineID" clearable placeholder="请输入"></el-input>
      </el-form-item>
       
         <el-form-item label="LineName:">
            <el-input v-model="formData.lineName" clearable placeholder="请输入" ></el-input>
      </el-form-item>
       
         <el-form-item label="LineType:">
            <el-input v-model="formData.lineType" clearable placeholder="请输入" ></el-input>
      </el-form-item>
       
         <el-form-item label="程式读取路径:">
            <el-input v-model="formData.readFilePath" clearable placeholder="请输入" ></el-input>
      </el-form-item>
       
         <el-form-item label="LineName2:">
            <el-input v-model="formData.lineName2" clearable placeholder="请输入" ></el-input>
      </el-form-item>
       
         <el-form-item label="制令单号:">
            <el-input v-model="formData.mOrderNo" clearable placeholder="请输入" ></el-input>
      </el-form-item>
       
         <el-form-item label="Remark:">
            <el-input v-model="formData.remark" clearable placeholder="请输入" ></el-input>
      </el-form-item>
       
         <el-form-item label="是否报警:">
            <el-switch active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" v-model="formData.warnning" clearable ></el-switch>
      </el-form-item>
       
         <el-form-item label="WarnningStr:">
            <el-input v-model="formData.warnningStr" clearable placeholder="请输入" ></el-input>
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
    createPVS_Base_Line,
    deletePVS_Base_Line,
    deletePVS_Base_LineByIds,
    updatePVS_Base_Line,
    findPVS_Base_Line,
    getPVS_Base_LineList
} from "@/api/PVS_Base_Line";  //  此处请自行替换地址
import { formatTimeToStr } from "@/utils/date";
import infoList from "@/mixins/infoList";
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
        this.pageSize = 10             
        if (this.searchInfo.warnning==""){
          this.searchInfo.warnning=null
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
    await this.getTableData();
  
}
};
</script>

<style>
</style>
