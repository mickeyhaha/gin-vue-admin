<template>
  <div>
    <div class="search-term">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
        <el-form-item label="deviceId字段">
          <el-input placeholder="搜索条件" v-model="searchInfo.deviceId"></el-input>
        </el-form-item>    
        <el-form-item label="deviceName字段">
          <el-input placeholder="搜索条件" v-model="searchInfo.deviceName"></el-input>
        </el-form-item>    
        <el-form-item label="displayId字段">
          <el-input placeholder="搜索条件" v-model="searchInfo.displayId"></el-input>
        </el-form-item>    
        <el-form-item label="displayName字段">
          <el-input placeholder="搜索条件" v-model="searchInfo.displayName"></el-input>
        </el-form-item>    
        <el-form-item label="barCode字段">
          <el-input placeholder="搜索条件" v-model="searchInfo.barCode"></el-input>
        </el-form-item>                          
        <el-form-item label="opName">
          <el-input placeholder="搜索条件" v-model="searchInfo.opName"></el-input>
        </el-form-item>    
        <el-form-item>
          <el-button @click="onSubmit" type="primary">查询</el-button>
        </el-form-item>
        <el-form-item>
          <el-button @click="openDialog" type="primary">新增device表</el-button>
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
    
    <el-table-column label="deviceId字段" prop="deviceId" width="120"></el-table-column> 
    
    <el-table-column label="deviceName字段" prop="deviceName" width="120"></el-table-column> 
    
    <el-table-column label="displayId字段" prop="displayId" width="120"></el-table-column> 
    
    <el-table-column label="displayName字段" prop="displayName" width="120"></el-table-column> 
    
    <el-table-column label="barCode字段" prop="barCode" width="120"></el-table-column> 
    
    <el-table-column label="longitude字段" prop="longitude" width="120"></el-table-column> 
    
    <el-table-column label="latitude字段" prop="latitude" width="120"></el-table-column> 
    
    <el-table-column label="placeId字段" prop="placeId" width="120"></el-table-column> 
    
    <el-table-column label="placeName字段" prop="placeName" width="120"></el-table-column> 
    
    <el-table-column label="status字段" prop="status" width="120"></el-table-column> 
    
    <el-table-column label="temp字段" prop="temp" width="120"></el-table-column> 
    
    <el-table-column label="tempMax字段" prop="tempMax" width="120"></el-table-column> 
    
    <el-table-column label="tempMin字段" prop="tempMin" width="120"></el-table-column> 
    
    <el-table-column label="tempAvg字段" prop="tempAvg" width="120"></el-table-column> 
    
    <el-table-column label="threshold字段" prop="threshold" width="120"></el-table-column> 
    
    <el-table-column label="operatorId字段" prop="opId" width="120"></el-table-column> 
    
    <el-table-column label="opName" prop="opName" width="120"></el-table-column> 
    
      <el-table-column label="按钮组">
        <template slot-scope="scope">
          <el-button class="table-button" @click="updateDevice(scope.row)" size="small" type="primary" icon="el-icon-edit">变更</el-button>
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
         <el-form-item label="deviceId字段:">
            <el-input v-model="formData.deviceId" clearable placeholder="请输入" ></el-input>
      </el-form-item>
       
         <el-form-item label="deviceName字段:">
            <el-input v-model="formData.deviceName" clearable placeholder="请输入" ></el-input>
      </el-form-item>
       
         <el-form-item label="displayId字段:">
            <el-input v-model="formData.displayId" clearable placeholder="请输入" ></el-input>
      </el-form-item>
       
         <el-form-item label="displayName字段:">
            <el-input v-model="formData.displayName" clearable placeholder="请输入" ></el-input>
      </el-form-item>
       
         <el-form-item label="barCode字段:">
            <el-input v-model="formData.barCode" clearable placeholder="请输入" ></el-input>
      </el-form-item>
       
         <el-form-item label="longitude字段:">
              <el-input-number v-model="formData.longitude" :precision="2" clearable></el-input-number>
       </el-form-item>
       
         <el-form-item label="latitude字段:">
              <el-input-number v-model="formData.latitude" :precision="2" clearable></el-input-number>
       </el-form-item>
       
         <el-form-item label="placeId字段:"><el-input v-model="formData.placeId" clearable placeholder="请输入"></el-input>
      </el-form-item>
       
         <el-form-item label="placeName字段:"><el-input v-model.number="formData.placeName" clearable placeholder="请输入"></el-input>
      </el-form-item>
       
         <el-form-item label="status字段:"><el-input v-model.number="formData.status" clearable placeholder="请输入"></el-input>
      </el-form-item>
       
         <el-form-item label="temp字段:">
              <el-input-number v-model="formData.temp" :precision="2" clearable></el-input-number>
       </el-form-item>
       
         <el-form-item label="tempMax字段:">
              <el-input-number v-model="formData.tempMax" :precision="2" clearable></el-input-number>
       </el-form-item>
       
         <el-form-item label="tempMin字段:">
              <el-input-number v-model="formData.tempMin" :precision="2" clearable></el-input-number>
       </el-form-item>
       
         <el-form-item label="tempAvg字段:">
              <el-input-number v-model="formData.tempAvg" :precision="2" clearable></el-input-number>
       </el-form-item>
       
         <el-form-item label="threshold字段:">
              <el-input-number v-model="formData.threshold" :precision="2" clearable></el-input-number>
       </el-form-item>
       
         <el-form-item label="operatorId字段:"><el-input v-model.number="formData.opId" clearable placeholder="请输入"></el-input>
      </el-form-item>
       
         <el-form-item label="opName:">
            <el-input v-model="formData.opName" clearable placeholder="请输入" ></el-input>
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
    createDevice,
    deleteDevice,
    deleteDeviceByIds,
    updateDevice,
    findDevice,
    getDeviceList
} from "@/api/device";  //  此处请自行替换地址
import { formatTimeToStr } from "@/utils/date";
import infoList from "@/mixins/infoList";
export default {
  name: "Device",
  mixins: [infoList],
  data() {
    return {
      listApi: getDeviceList,
      dialogFormVisible: false,
      type: "",
      deleteVisible: false,
      multipleSelection: [],formData: {
            deviceId:"",
            deviceName:"",
            displayId:"",
            displayName:"",
            barCode:"",
            longitude:0,
            latitude:0,
            placeId:0,
            placeName:0,
            status:0,
            temp:0,
            tempMax:0,
            tempMin:0,
            tempAvg:0,
            threshold:0,
            opId:0,
            opName:"",
            
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
           this.deleteDevice(row);
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
        const res = await deleteDeviceByIds({ ids })
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
    async updateDevice(row) {
      const res = await findDevice({ ID: row.ID });
      this.type = "update";
      if (res.code == 0) {
        this.formData = res.data.redev;
        this.dialogFormVisible = true;
      }
    },
    closeDialog() {
      this.dialogFormVisible = false;
      this.formData = {
          deviceId:"",
          deviceName:"",
          displayId:"",
          displayName:"",
          barCode:"",
          longitude:0,
          latitude:0,
          placeId:0,
          placeName:0,
          status:0,
          temp:0,
          tempMax:0,
          tempMin:0,
          tempAvg:0,
          threshold:0,
          opId:0,
          opName:"",
          
      };
    },
    async deleteDevice(row) {
      const res = await deleteDevice({ ID: row.ID });
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
          res = await createDevice(this.formData);
          break;
        case "update":
          res = await updateDevice(this.formData);
          break;
        default:
          res = await createDevice(this.formData);
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
