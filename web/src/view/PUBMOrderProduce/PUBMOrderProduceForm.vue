<template>
<div>
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
           <el-form-item>
           <el-button @click="save" type="primary">保存</el-button>
           <el-button @click="back" type="primary">返回</el-button>
           </el-form-item>
    </el-form>
</div>
</template>

<script>
import {
    createPUBMOrderProduce2,
    updatePUBMOrderProduce2,
    findPUBMOrderProduce2
} from "@/api/PUBMOrderProduce";  //  此处请自行替换地址
import infoList from "@/mixins/infoList";
export default {
  name: "PUBMOrderProduce2",
  mixins: [infoList],
  data() {
    return {
      type: "",
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
  methods: {
    async save() {
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
      }
    },
    back(){
        this.$router.go(-1)
    }
  },
  async created() {
   // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if(this.$route.query.id){
    const res = await findPUBMOrderProduce2({ ID: this.$route.query.id })
    if(res.code == 0){
       this.formData = res.data.rePUBMOrderProduce
       this.type = "update"
     }
    }else{
     this.type = "create"
   }
  
    await this.getDict("OrderStatus");
    
    await this.getDict("PanelType");
    
}
};
</script>

<style>
</style>