<template>
  <div>
    <div class="search-term">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">                                
        <el-form-item label="Status字段">
          <el-input placeholder="搜索条件" v-model="searchInfo.Status"></el-input>
        </el-form-item>            
        <el-form-item label="PCB类型">
          <el-input placeholder="搜索条件" v-model="searchInfo.PanelType"></el-input>
        </el-form-item>    
        <el-form-item>
          <el-button @click="onSubmit" type="primary">查询</el-button>
        </el-form-item>
        <el-form-item>
          <el-button @click="openDialog" type="primary">新增PUBMOrderProduce表</el-button>
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
    
    <el-table-column label="MOrderNo字段" prop="MOrderNo" width="120"></el-table-column> 
    
    <el-table-column label="WorkOrderNo字段" prop="WorkOrderNo" width="120"></el-table-column> 
    
    <el-table-column label="Customer字段" prop="Customer" width="120"></el-table-column> 
    
    <el-table-column label="CustOrderNo字段" prop="CustOrderNo" width="120"></el-table-column> 
    
    <el-table-column label="MachineType字段" prop="MachineType" width="120"></el-table-column> 
    
    <el-table-column label="BOM编号" prop="BOMNo" width="120"></el-table-column> 
    
    <el-table-column label="BOM版本" prop="BOMVersion" width="120"></el-table-column> 
    
    <el-table-column label="产品编号" prop="Product" width="120"></el-table-column> 
    
    <el-table-column label="周期" prop="CycleTime" width="120"></el-table-column> 
    
    <el-table-column label="版面" prop="PasteSide" width="120"></el-table-column> 
    
    <el-table-column label="拼版数" prop="PanelCount" width="120"></el-table-column> 
    
    <el-table-column label="计划数量" prop="Qty" width="120"></el-table-column> 
    
    <el-table-column label="QtyCompleted字段" prop="QtyCompleted" width="120"></el-table-column> 
    
    <el-table-column label="开始时间" prop="BeginTime" width="120"></el-table-column> 
    
    <el-table-column label="结束时间" prop="EndTime" width="120"></el-table-column> 
    
      <el-table-column label="Status字段" prop="Status" width="120">
        <template slot-scope="scope">
          {{filterDict(scope.row.Status,"OrderStatus")}}
        </template>
      </el-table-column>
    
    <el-table-column label="备注" prop="Remark" width="120"></el-table-column> 
    
    <el-table-column label="是否启动JIT拉动" prop="JITRunning" width="120"></el-table-column> 
    
    <el-table-column label="CreateTime字段" prop="CreateTime" width="120"></el-table-column> 
    
    <el-table-column label="最后时间" prop="LMTime" width="120"></el-table-column> 
    
      <el-table-column label="PCB类型" prop="PanelType" width="120">
        <template slot-scope="scope">
          {{filterDict(scope.row.PanelType,"PanelType")}}
        </template>
      </el-table-column>
    
      <el-table-column label="按钮组">
        <template slot-scope="scope">
          <el-button class="table-button" @click="updatePUBMOrderProduce2(scope.row)" size="small" type="primary" icon="el-icon-edit">变更</el-button>
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
       
         <el-form-item label="MOrderNo字段:">
            <el-input v-model="formData.MOrderNo" clearable placeholder="请输入" ></el-input>
      </el-form-item>
       
         <el-form-item label="WorkOrderNo字段:">
            <el-input v-model="formData.WorkOrderNo" clearable placeholder="请输入" ></el-input>
      </el-form-item>
       
         <el-form-item label="Customer字段:">
            <el-input v-model="formData.Customer" clearable placeholder="请输入" ></el-input>
      </el-form-item>
       
         <el-form-item label="CustOrderNo字段:">
            <el-input v-model="formData.CustOrderNo" clearable placeholder="请输入" ></el-input>
      </el-form-item>
       
         <el-form-item label="MachineType字段:">
            <el-input v-model="formData.MachineType" clearable placeholder="请输入" ></el-input>
      </el-form-item>
       
         <el-form-item label="BOM编号:">
            <el-input v-model="formData.BOMNo" clearable placeholder="请输入" ></el-input>
      </el-form-item>
       
         <el-form-item label="BOM版本:">
            <el-input v-model="formData.BOMVersion" clearable placeholder="请输入" ></el-input>
      </el-form-item>
       
         <el-form-item label="产品编号:">
            <el-input v-model="formData.Product" clearable placeholder="请输入" ></el-input>
      </el-form-item>
       
         <el-form-item label="周期:"><el-input v-model.number="formData.CycleTime" clearable placeholder="请输入"></el-input>
      </el-form-item>
       
         <el-form-item label="版面:">
            <el-input v-model="formData.PasteSide" clearable placeholder="请输入" ></el-input>
      </el-form-item>
       
         <el-form-item label="拼版数:"><el-input v-model.number="formData.PanelCount" clearable placeholder="请输入"></el-input>
      </el-form-item>
       
         <el-form-item label="计划数量:"><el-input v-model.number="formData.Qty" clearable placeholder="请输入"></el-input>
      </el-form-item>
       
         <el-form-item label="QtyCompleted字段:"><el-input v-model.number="formData.QtyCompleted" clearable placeholder="请输入"></el-input>
      </el-form-item>
       
         <el-form-item label="开始时间:">
              <el-date-picker type="date" placeholder="选择日期" v-model="formData.BeginTime" clearable></el-date-picker>
       </el-form-item>
       
         <el-form-item label="结束时间:">
              <el-date-picker type="date" placeholder="选择日期" v-model="formData.EndTime" clearable></el-date-picker>
       </el-form-item>
       
         <el-form-item label="Status字段:">
             <el-select v-model="formData.Status" placeholder="请选择" clearable>
                 <el-option v-for="(item,key) in OrderStatusOptions" :key="key" :label="item.label" :value="item.value"></el-option>
             </el-select>
      </el-form-item>
       
         <el-form-item label="备注:">
            <el-input v-model="formData.Remark" clearable placeholder="请输入" ></el-input>
      </el-form-item>
       
         <el-form-item label="是否启动JIT拉动:"><el-input v-model.number="formData.JITRunning" clearable placeholder="请输入"></el-input>
      </el-form-item>
       
         <el-form-item label="CreateTime字段:">
              <el-date-picker type="date" placeholder="选择日期" v-model="formData.CreateTime" clearable></el-date-picker>
       </el-form-item>
       
         <el-form-item label="最后时间:">
              <el-date-picker type="date" placeholder="选择日期" v-model="formData.LMTime" clearable></el-date-picker>
       </el-form-item>
       
         <el-form-item label="PCB类型:">
             <el-select v-model="formData.PanelType" placeholder="请选择" clearable>
                 <el-option v-for="(item,key) in PanelTypeOptions" :key="key" :label="item.label" :value="item.value"></el-option>
             </el-select>
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
    createPUBMOrderProduce2,
    deletePUBMOrderProduce2,
    deletePUBMOrderProduce2ByIds,
    updatePUBMOrderProduce2,
    findPUBMOrderProduce2,
    getPUBMOrderProduce2List
} from "@/api/PUBMOrderProduce";  //  此处请自行替换地址
import { formatTimeToStr } from "@/utils/date";
import infoList from "@/mixins/infoList";
export default {
  name: "PUBMOrderProduce2",
  mixins: [infoList],
  data() {
    return {
      listApi: getPUBMOrderProduce2List,
      dialogFormVisible: false,
      type: "",
      deleteVisible: false,
      multipleSelection: [],
      OrderStatusOptions:[],
          
      PanelTypeOptions:[],
          formData: {
            LineName:"",
            MOrderNo:"",
            WorkOrderNo:"",
            Customer:"",
            CustOrderNo:"",
            MachineType:"",
            BOMNo:"",
            BOMVersion:"",
            Product:"",
            CycleTime:0,
            PasteSide:"",
            PanelCount:0,
            Qty:0,
            QtyCompleted:0,
            BeginTime:new Date(),
            EndTime:new Date(),
            Status:0,
            Remark:"",
            JITRunning:0,
            CreateTime:new Date(),
            LMTime:new Date(),
            PanelType:0,
            
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
           this.deletePUBMOrderProduce2(row);
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
        const res = await deletePUBMOrderProduce2ByIds({ ids })
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
    async updatePUBMOrderProduce2(row) {
      const res = await findPUBMOrderProduce2({ ID: row.ID });
      this.type = "update";
      if (res.code == 0) {
        this.formData = res.data.rePUBMOrderProduce;
        this.dialogFormVisible = true;
      }
    },
    closeDialog() {
      this.dialogFormVisible = false;
      this.formData = {
          LineName:"",
          MOrderNo:"",
          WorkOrderNo:"",
          Customer:"",
          CustOrderNo:"",
          MachineType:"",
          BOMNo:"",
          BOMVersion:"",
          Product:"",
          CycleTime:0,
          PasteSide:"",
          PanelCount:0,
          Qty:0,
          QtyCompleted:0,
          BeginTime:new Date(),
          EndTime:new Date(),
          Status:0,
          Remark:"",
          JITRunning:0,
          CreateTime:new Date(),
          LMTime:new Date(),
          PanelType:0,
          
      };
    },
    async deletePUBMOrderProduce2(row) {
      const res = await deletePUBMOrderProduce2({ ID: row.ID });
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
          res = await createPUBMOrderProduce2(this.formData);
          break;
        case "update":
          res = await updatePUBMOrderProduce2(this.formData);
          break;
        default:
          res = await createPUBMOrderProduce2(this.formData);
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
  
    await this.getDict("OrderStatus");
    
    await this.getDict("PanelType");
    
}
};
</script>

<style>
</style>
