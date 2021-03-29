<template>
  <div>
    <div class="search-term">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
        <el-form-item label="生产线">
          <el-input placeholder="搜索条件" v-model="searchInfo.lineName"></el-input>
        </el-form-item>    
        <el-form-item label="制令单">
          <el-input placeholder="搜索条件" v-model="searchInfo.mOrderNo"></el-input>
        </el-form-item>    
        <el-form-item label="工单">
          <el-input placeholder="搜索条件" v-model="searchInfo.workOrderNo"></el-input>
        </el-form-item>    
        <el-form-item label="客户">
          <el-input placeholder="搜索条件" v-model="searchInfo.customer"></el-input>
        </el-form-item>    
        <el-form-item label="客户订单号">
          <el-input placeholder="搜索条件" v-model="searchInfo.custOrderNo"></el-input>
        </el-form-item>    
        <el-form-item label="机型">
          <el-input placeholder="搜索条件" v-model="searchInfo.machineType"></el-input>
        </el-form-item>    
        <el-form-item label="BOM编号">
          <el-input placeholder="搜索条件" v-model="searchInfo.bOMNo"></el-input>
        </el-form-item>    
        <el-form-item label="BOM版本">
          <el-input placeholder="搜索条件" v-model="searchInfo.bOMVersion"></el-input>
        </el-form-item>    
        <el-form-item label="产品编号">
          <el-input placeholder="搜索条件" v-model="searchInfo.product"></el-input>
        </el-form-item>      
        <el-form-item label="版面">
          <el-input placeholder="搜索条件" v-model="searchInfo.pasteSide"></el-input>
        </el-form-item>          
        <el-form-item label="开始时间">
          <el-input placeholder="搜索条件" v-model="searchInfo.beginTime"></el-input>
        </el-form-item>        
        <el-form-item label="备注">
          <el-input placeholder="搜索条件" v-model="searchInfo.remark"></el-input>
        </el-form-item>                    
        <el-form-item>
          <el-button @click="onSubmit" type="primary">查询</el-button>
        </el-form-item>
        <el-form-item>
          <el-button @click="openDialog" type="primary">新增PVS制令单表</el-button>
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
    
    <el-table-column label="生产线" prop="lineName" width="120"></el-table-column> 
    
    <el-table-column label="制令单" prop="mOrderNo" width="120"></el-table-column> 
    
    <el-table-column label="工单" prop="workOrderNo" width="120"></el-table-column> 
    
    <el-table-column label="客户" prop="customer" width="120"></el-table-column> 
    
    <el-table-column label="客户订单号" prop="custOrderNo" width="120"></el-table-column> 
    
    <el-table-column label="机型" prop="machineType" width="120"></el-table-column> 
    
    <el-table-column label="BOM编号" prop="bOMNo" width="120"></el-table-column> 
    
    <el-table-column label="BOM版本" prop="bOMVersion" width="120"></el-table-column> 
    
    <el-table-column label="产品编号" prop="product" width="120"></el-table-column> 
    
    <el-table-column label="周期" prop="cycleTime" width="120"></el-table-column> 
    
    <el-table-column label="版面" prop="pasteSide" width="120"></el-table-column> 
    
    <el-table-column label="拼版数" prop="panelCount" width="120"></el-table-column> 
    
    <el-table-column label="计划数量" prop="qty" width="120"></el-table-column> 
    
    <el-table-column label="完成数量" prop="qtyCompleted" width="120"></el-table-column> 
    
    <el-table-column label="开始时间" prop="beginTime" width="120"></el-table-column> 
    
    <el-table-column label="结束时间" prop="endTime" width="120"></el-table-column> 
    
    <el-table-column label="状态" prop="status" width="120"></el-table-column> 
    
    <el-table-column label="备注" prop="remark" width="120"></el-table-column> 
    
    <el-table-column label="是否启动JIT拉动" prop="jITRunning" width="120"></el-table-column> 
    
    <el-table-column label="创建时间" prop="createTime" width="120"></el-table-column> 
    
    <el-table-column label="最后时间" prop="lMTime" width="120"></el-table-column> 
    
    <el-table-column label="PCB类型" prop="panelType" width="120"></el-table-column> 
    
    <el-table-column label="样品单" prop="sampleRate" width="120"></el-table-column> 
    
    <el-table-column label="板面" prop="sideIndex" width="120"></el-table-column> 
    
    <el-table-column label="轨道一" prop="lane1" width="120">
         <template slot-scope="scope">{{scope.row.lane1|formatBoolean}}</template>
    </el-table-column>
    
    <el-table-column label="轨道二" prop="lane2" width="120">
         <template slot-scope="scope">{{scope.row.lane2|formatBoolean}}</template>
    </el-table-column>
    
      <el-table-column label="按钮组">
        <template slot-scope="scope">
          <el-button class="table-button" @click="updatePUB_MOrderProduce(scope.row)" size="small" type="primary" icon="el-icon-edit">变更</el-button>
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
         <el-form-item label="生产线:">
            <el-input v-model="formData.lineName" clearable placeholder="请输入" ></el-input>
      </el-form-item>
       
         <el-form-item label="制令单:">
            <el-input v-model="formData.mOrderNo" clearable placeholder="请输入" ></el-input>
      </el-form-item>
       
         <el-form-item label="工单:">
            <el-input v-model="formData.workOrderNo" clearable placeholder="请输入" ></el-input>
      </el-form-item>
       
         <el-form-item label="客户:">
            <el-input v-model="formData.customer" clearable placeholder="请输入" ></el-input>
      </el-form-item>
       
         <el-form-item label="客户订单号:">
            <el-input v-model="formData.custOrderNo" clearable placeholder="请输入" ></el-input>
      </el-form-item>
       
         <el-form-item label="机型:">
            <el-input v-model="formData.machineType" clearable placeholder="请输入" ></el-input>
      </el-form-item>
       
         <el-form-item label="BOM编号:">
            <el-input v-model="formData.bOMNo" clearable placeholder="请输入" ></el-input>
      </el-form-item>
       
         <el-form-item label="BOM版本:">
            <el-input v-model="formData.bOMVersion" clearable placeholder="请输入" ></el-input>
      </el-form-item>
       
         <el-form-item label="产品编号:">
            <el-input v-model="formData.product" clearable placeholder="请输入" ></el-input>
      </el-form-item>
       
         <el-form-item label="周期:"><el-input v-model.number="formData.cycleTime" clearable placeholder="请输入"></el-input>
      </el-form-item>
       
         <el-form-item label="版面:">
            <el-input v-model="formData.pasteSide" clearable placeholder="请输入" ></el-input>
      </el-form-item>
       
         <el-form-item label="拼版数:"><el-input v-model.number="formData.panelCount" clearable placeholder="请输入"></el-input>
      </el-form-item>
       
         <el-form-item label="计划数量:"><el-input v-model.number="formData.qty" clearable placeholder="请输入"></el-input>
      </el-form-item>
       
         <el-form-item label="完成数量:"><el-input v-model.number="formData.qtyCompleted" clearable placeholder="请输入"></el-input>
      </el-form-item>
       
         <el-form-item label="开始时间:">
              <el-date-picker type="date" placeholder="选择日期" v-model="formData.beginTime" clearable></el-date-picker>
       </el-form-item>
       
         <el-form-item label="结束时间:">
              <el-date-picker type="date" placeholder="选择日期" v-model="formData.endTime" clearable></el-date-picker>
       </el-form-item>
       
         <el-form-item label="状态:"><el-input v-model.number="formData.status" clearable placeholder="请输入"></el-input>
      </el-form-item>
       
         <el-form-item label="备注:">
            <el-input v-model="formData.remark" clearable placeholder="请输入" ></el-input>
      </el-form-item>
       
         <el-form-item label="是否启动JIT拉动:"><el-input v-model.number="formData.jITRunning" clearable placeholder="请输入"></el-input>
      </el-form-item>
       
         <el-form-item label="创建时间:">
              <el-date-picker type="date" placeholder="选择日期" v-model="formData.createTime" clearable></el-date-picker>
       </el-form-item>
       
         <el-form-item label="最后时间:">
              <el-date-picker type="date" placeholder="选择日期" v-model="formData.lMTime" clearable></el-date-picker>
       </el-form-item>
       
         <el-form-item label="PCB类型:"><el-input v-model.number="formData.panelType" clearable placeholder="请输入"></el-input>
      </el-form-item>
       
         <el-form-item label="样品单:">
            <el-input v-model="formData.sampleRate" clearable placeholder="请输入" ></el-input>
      </el-form-item>
       
         <el-form-item label="板面:">
            <el-input v-model="formData.sideIndex" clearable placeholder="请输入" ></el-input>
      </el-form-item>
       
         <el-form-item label="轨道一:">
            <el-switch active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" v-model="formData.lane1" clearable ></el-switch>
      </el-form-item>
       
         <el-form-item label="轨道二:">
            <el-switch active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" v-model="formData.lane2" clearable ></el-switch>
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
    createPUB_MOrderProduce,
    deletePUB_MOrderProduce,
    deletePUB_MOrderProduceByIds,
    updatePUB_MOrderProduce,
    findPUB_MOrderProduce,
    getPUB_MOrderProduceList
} from "@/api/PUB_MOrderProduce";  //  此处请自行替换地址
import { formatTimeToStr } from "@/utils/date";
import infoList from "@/mixins/infoList";
export default {
  name: "PUB_MOrderProduce",
  mixins: [infoList],
  data() {
    return {
      listApi: getPUB_MOrderProduceList,
      dialogFormVisible: false,
      type: "",
      deleteVisible: false,
      multipleSelection: [],formData: {
            lineName:"",
            mOrderNo:"",
            workOrderNo:"",
            customer:"",
            custOrderNo:"",
            machineType:"",
            bOMNo:"",
            bOMVersion:"",
            product:"",
            cycleTime:0,
            pasteSide:"",
            panelCount:0,
            qty:0,
            qtyCompleted:0,
            beginTime:new Date(),
            endTime:new Date(),
            status:0,
            remark:"",
            jITRunning:0,
            createTime:new Date(),
            lMTime:new Date(),
            panelType:0,
            sampleRate:"",
            sideIndex:"",
            lane1:false,
            lane2:false,
            
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
        if (this.searchInfo.lane1==""){
          this.searchInfo.lane1=null
        }        
        if (this.searchInfo.lane2==""){
          this.searchInfo.lane2=null
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
           this.deletePUB_MOrderProduce(row);
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
        const res = await deletePUB_MOrderProduceByIds({ ids })
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
    async updatePUB_MOrderProduce(row) {
      const res = await findPUB_MOrderProduce({ ID: row.ID });
      this.type = "update";
      if (res.code == 0) {
        this.formData = res.data.rePMOP;
        this.dialogFormVisible = true;
      }
    },
    closeDialog() {
      this.dialogFormVisible = false;
      this.formData = {
          lineName:"",
          mOrderNo:"",
          workOrderNo:"",
          customer:"",
          custOrderNo:"",
          machineType:"",
          bOMNo:"",
          bOMVersion:"",
          product:"",
          cycleTime:0,
          pasteSide:"",
          panelCount:0,
          qty:0,
          qtyCompleted:0,
          beginTime:new Date(),
          endTime:new Date(),
          status:0,
          remark:"",
          jITRunning:0,
          createTime:new Date(),
          lMTime:new Date(),
          panelType:0,
          sampleRate:"",
          sideIndex:"",
          lane1:false,
          lane2:false,
          
      };
    },
    async deletePUB_MOrderProduce(row) {
      const res = await deletePUB_MOrderProduce({ ID: row.ID });
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
          res = await createPUB_MOrderProduce(this.formData);
          break;
        case "update":
          res = await updatePUB_MOrderProduce(this.formData);
          break;
        default:
          res = await createPUB_MOrderProduce(this.formData);
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
