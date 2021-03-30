<template>
  <div>
    <div class="search-term">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
        <el-form-item label="LineID">
          <el-input placeholder="搜索条件" v-model="searchInfo.lineID"></el-input>
        </el-form-item>    
        <el-form-item label="机器">
          <el-input placeholder="搜索条件" v-model="searchInfo.machineCode"></el-input>
        </el-form-item>    
        <el-form-item label="台车">
          <el-input placeholder="搜索条件" v-model="searchInfo.tableNo"></el-input>
        </el-form-item>    
        <el-form-item label="BCT">
          <el-input placeholder="搜索条件" v-model="searchInfo.bCT"></el-input>
        </el-form-item>    
        <el-form-item label="TrackNo">
          <el-input placeholder="搜索条件" v-model="searchInfo.trackNo"></el-input>
        </el-form-item>    
        <el-form-item label="Division">
          <el-input placeholder="搜索条件" v-model="searchInfo.division"></el-input>
        </el-form-item>    
        <el-form-item label="吸取位置">
          <el-input placeholder="搜索条件" v-model="searchInfo.pickPos"></el-input>
        </el-form-item>    
        <el-form-item label="FeederType">
          <el-input placeholder="搜索条件" v-model="searchInfo.feederType"></el-input>
        </el-form-item>    
        <el-form-item label="FeederNo">
          <el-input placeholder="搜索条件" v-model="searchInfo.feederNo"></el-input>
        </el-form-item>    
        <el-form-item label="料号">
          <el-input placeholder="搜索条件" v-model="searchInfo.matrCode"></el-input>
        </el-form-item>    
        <el-form-item label="PartStatus">
          <el-input placeholder="搜索条件" v-model="searchInfo.partStatus"></el-input>
        </el-form-item>        
        <el-form-item label="PlacedQty">
          <el-input placeholder="搜索条件" v-model="searchInfo.placedQty"></el-input>
        </el-form-item>    
        <el-form-item label="吸取错误">
          <el-input placeholder="搜索条件" v-model="searchInfo.pickError"></el-input>
        </el-form-item>                    
            <el-form-item label="StopTrail" prop="stopTrail">
            <el-select v-model="searchInfo.stopTrail" clear placeholder="请选择">
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
        <el-form-item label="Scanner">
          <el-input placeholder="搜索条件" v-model="searchInfo.scanner"></el-input>
        </el-form-item>          
        <el-form-item>
          <el-button @click="onSubmit" type="primary">查询</el-button>
        </el-form-item>
        <el-form-item>
          <el-button @click="openDialog" type="primary">新增监控视图</el-button>
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
    
    <el-table-column label="机器" prop="machineCode" width="120"></el-table-column> 
    
    <el-table-column label="台车" prop="tableNo" width="120"></el-table-column> 
    
    <el-table-column label="BCT" prop="bCT" width="120"></el-table-column> 
    
    <el-table-column label="TrackNo" prop="trackNo" width="120"></el-table-column> 
    
    <el-table-column label="Division" prop="division" width="120"></el-table-column> 
    
    <el-table-column label="吸取位置" prop="pickPos" width="120"></el-table-column> 
    
    <el-table-column label="FeederType" prop="feederType" width="120"></el-table-column> 
    
    <el-table-column label="FeederNo" prop="feederNo" width="120"></el-table-column> 
    
    <el-table-column label="料号" prop="matrCode" width="120"></el-table-column> 
    
    <el-table-column label="PartStatus" prop="partStatus" width="120"></el-table-column> 
    
    <el-table-column label="CycleTime" prop="cycleTime" width="120"></el-table-column> 
    
    <el-table-column label="PartPos" prop="partPos" width="120"></el-table-column> 
    
    <el-table-column label="PlacedQty" prop="placedQty" width="120"></el-table-column> 
    
      <el-table-column label="吸取错误" prop="pickError" width="120">
        <template slot-scope="scope">
          {{filterDict(scope.row.pickError,"int")}}
        </template>
      </el-table-column>
    
    <el-table-column label="IdentError" prop="identError" width="120"></el-table-column> 
    
    <el-table-column label="OtherError" prop="otherError" width="120"></el-table-column> 
    
    <el-table-column label="TotalConsume" prop="totalConsume" width="120"></el-table-column> 
    
    <el-table-column label="UniqueCode" prop="uniqueCode" width="120"></el-table-column> 
    
    <el-table-column label="LeftQty" prop="leftQty" width="120"></el-table-column> 
    
    <el-table-column label="StandbyCode" prop="standbyCode" width="120"></el-table-column> 
    
    <el-table-column label="StandbyQty" prop="standbyQty" width="120"></el-table-column> 
    
    <el-table-column label="TotalLeftQty" prop="totalLeftQty" width="120"></el-table-column> 
    
    <el-table-column label="StopTrail" prop="stopTrail" width="120">
         <template slot-scope="scope">{{scope.row.stopTrail|formatBoolean}}</template>
    </el-table-column>
    
    <el-table-column label="TotalError" prop="totalError" width="120"></el-table-column> 
    
    <el-table-column label="ChangeTime" prop="changeTime" width="120"></el-table-column> 
    
    <el-table-column label="RejectRate" prop="rejectRate" width="120"></el-table-column> 
    
    <el-table-column label="LeftTime" prop="leftTime" width="120"></el-table-column> 
    
    <el-table-column label="UseLevel" prop="useLevel" width="120"></el-table-column> 
    
    <el-table-column label="Scanner" prop="scanner" width="120"></el-table-column> 
    
    <el-table-column label="ScanTime" prop="scanTime" width="120"></el-table-column> 
    
    <el-table-column label="Remark" prop="remark" width="120"></el-table-column> 
    
    <el-table-column label="IsTray" prop="isTray" width="120">
         <template slot-scope="scope">{{scope.row.isTray|formatBoolean}}</template>
    </el-table-column>
    
      <el-table-column label="按钮组">
        <template slot-scope="scope">
          <el-button class="table-button" @click="updateMoniWholeView(scope.row)" size="small" type="primary" icon="el-icon-edit">变更</el-button>
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
       
         <el-form-item label="机器:"><el-input v-model.number="formData.machineCode" clearable placeholder="请输入"></el-input>
      </el-form-item>
       
         <el-form-item label="台车:">
            <el-input v-model="formData.tableNo" clearable placeholder="请输入" ></el-input>
      </el-form-item>
       
         <el-form-item label="BCT:">
            <el-input v-model="formData.bCT" clearable placeholder="请输入" ></el-input>
      </el-form-item>
       
         <el-form-item label="TrackNo:"><el-input v-model.number="formData.trackNo" clearable placeholder="请输入"></el-input>
      </el-form-item>
       
         <el-form-item label="Division:"><el-input v-model.number="formData.division" clearable placeholder="请输入"></el-input>
      </el-form-item>
       
         <el-form-item label="吸取位置:"><el-input v-model.number="formData.pickPos" clearable placeholder="请输入"></el-input>
      </el-form-item>
       
         <el-form-item label="FeederType:">
            <el-input v-model="formData.feederType" clearable placeholder="请输入" ></el-input>
      </el-form-item>
       
         <el-form-item label="FeederNo:">
            <el-input v-model="formData.feederNo" clearable placeholder="请输入" ></el-input>
      </el-form-item>
       
         <el-form-item label="料号:">
            <el-input v-model="formData.matrCode" clearable placeholder="请输入" ></el-input>
      </el-form-item>
       
         <el-form-item label="PartStatus:"><el-input v-model.number="formData.partStatus" clearable placeholder="请输入"></el-input>
      </el-form-item>
       
         <el-form-item label="CycleTime:">
              <el-input-number v-model="formData.cycleTime" :precision="2" clearable></el-input-number>
       </el-form-item>
       
         <el-form-item label="PartPos:">
            <el-input v-model="formData.partPos" clearable placeholder="请输入" ></el-input>
      </el-form-item>
       
         <el-form-item label="PlacedQty:"><el-input v-model.number="formData.placedQty" clearable placeholder="请输入"></el-input>
      </el-form-item>
       
         <el-form-item label="吸取错误:">
             <el-select v-model="formData.pickError" placeholder="请选择" clearable>
                 <el-option v-for="(item,key) in intOptions" :key="key" :label="item.label" :value="item.value"></el-option>
             </el-select>
      </el-form-item>
       
         <el-form-item label="IdentError:"><el-input v-model.number="formData.identError" clearable placeholder="请输入"></el-input>
      </el-form-item>
       
         <el-form-item label="OtherError:"><el-input v-model.number="formData.otherError" clearable placeholder="请输入"></el-input>
      </el-form-item>
       
         <el-form-item label="TotalConsume:"><el-input v-model.number="formData.totalConsume" clearable placeholder="请输入"></el-input>
      </el-form-item>
       
         <el-form-item label="UniqueCode:">
            <el-input v-model="formData.uniqueCode" clearable placeholder="请输入" ></el-input>
      </el-form-item>
       
         <el-form-item label="LeftQty:"><el-input v-model.number="formData.leftQty" clearable placeholder="请输入"></el-input>
      </el-form-item>
       
         <el-form-item label="StandbyCode:">
            <el-input v-model="formData.standbyCode" clearable placeholder="请输入" ></el-input>
      </el-form-item>
       
         <el-form-item label="StandbyQty:"><el-input v-model.number="formData.standbyQty" clearable placeholder="请输入"></el-input>
      </el-form-item>
       
         <el-form-item label="TotalLeftQty:"><el-input v-model.number="formData.totalLeftQty" clearable placeholder="请输入"></el-input>
      </el-form-item>
       
         <el-form-item label="StopTrail:">
            <el-switch active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" v-model="formData.stopTrail" clearable ></el-switch>
      </el-form-item>
       
         <el-form-item label="TotalError:"><el-input v-model.number="formData.totalError" clearable placeholder="请输入"></el-input>
      </el-form-item>
       
         <el-form-item label="ChangeTime:">
              <el-date-picker type="date" placeholder="选择日期" v-model="formData.changeTime" clearable></el-date-picker>
       </el-form-item>
       
         <el-form-item label="RejectRate:">
              <el-input-number v-model="formData.rejectRate" :precision="2" clearable></el-input-number>
       </el-form-item>
       
         <el-form-item label="LeftTime:">
              <el-input-number v-model="formData.leftTime" :precision="2" clearable></el-input-number>
       </el-form-item>
       
         <el-form-item label="UseLevel:"><el-input v-model.number="formData.useLevel" clearable placeholder="请输入"></el-input>
      </el-form-item>
       
         <el-form-item label="Scanner:">
            <el-input v-model="formData.scanner" clearable placeholder="请输入" ></el-input>
      </el-form-item>
       
         <el-form-item label="ScanTime:">
              <el-date-picker type="date" placeholder="选择日期" v-model="formData.scanTime" clearable></el-date-picker>
       </el-form-item>
       
         <el-form-item label="Remark:">
            <el-input v-model="formData.remark" clearable placeholder="请输入" ></el-input>
      </el-form-item>
       
         <el-form-item label="IsTray:">
            <el-switch active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" v-model="formData.isTray" clearable ></el-switch>
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
    createMoniWholeView,
    deleteMoniWholeView,
    deleteMoniWholeViewByIds,
    updateMoniWholeView,
    findMoniWholeView,
    getMoniWholeViewList
} from "@/api/MoniWholeView";  //  此处请自行替换地址
import { formatTimeToStr } from "@/utils/date";
import infoList from "@/mixins/infoList";
export default {
  name: "MoniWholeView",
  mixins: [infoList],
  data() {
    return {
      listApi: getMoniWholeViewList,
      dialogFormVisible: false,
      type: "",
      deleteVisible: false,
      multipleSelection: [],
      intOptions:[],
          formData: {
            lineID:0,
            machineCode:0,
            tableNo:"",
            bCT:"",
            trackNo:0,
            division:0,
            pickPos:0,
            feederType:"",
            feederNo:"",
            matrCode:"",
            partStatus:0,
            cycleTime:0,
            partPos:"",
            placedQty:0,
            pickError:0,
            identError:0,
            otherError:0,
            totalConsume:0,
            uniqueCode:"",
            leftQty:0,
            standbyCode:"",
            standbyQty:0,
            totalLeftQty:0,
            stopTrail:false,
            totalError:0,
            changeTime:new Date(),
            rejectRate:0,
            leftTime:0,
            useLevel:0,
            scanner:"",
            scanTime:new Date(),
            remark:"",
            isTray:false,
            
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
        if (this.searchInfo.stopTrail==""){
          this.searchInfo.stopTrail=null
        }                
        if (this.searchInfo.isTray==""){
          this.searchInfo.isTray=null
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
           this.deleteMoniWholeView(row);
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
        const res = await deleteMoniWholeViewByIds({ ids })
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
    async updateMoniWholeView(row) {
      const res = await findMoniWholeView({ ID: row.ID });
      this.type = "update";
      if (res.code == 0) {
        this.formData = res.data.reMWV;
        this.dialogFormVisible = true;
      }
    },
    closeDialog() {
      this.dialogFormVisible = false;
      this.formData = {
          lineID:0,
          machineCode:0,
          tableNo:"",
          bCT:"",
          trackNo:0,
          division:0,
          pickPos:0,
          feederType:"",
          feederNo:"",
          matrCode:"",
          partStatus:0,
          cycleTime:0,
          partPos:"",
          placedQty:0,
          pickError:0,
          identError:0,
          otherError:0,
          totalConsume:0,
          uniqueCode:"",
          leftQty:0,
          standbyCode:"",
          standbyQty:0,
          totalLeftQty:0,
          stopTrail:false,
          totalError:0,
          changeTime:new Date(),
          rejectRate:0,
          leftTime:0,
          useLevel:0,
          scanner:"",
          scanTime:new Date(),
          remark:"",
          isTray:false,
          
      };
    },
    async deleteMoniWholeView(row) {
      const res = await deleteMoniWholeView({ ID: row.ID });
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
          res = await createMoniWholeView(this.formData);
          break;
        case "update":
          res = await updateMoniWholeView(this.formData);
          break;
        default:
          res = await createMoniWholeView(this.formData);
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
  
    await this.getDict("int");
    
}
};
</script>

<style>
</style>
