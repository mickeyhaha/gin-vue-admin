<template>
<div>
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
           <el-form-item>
           <el-button @click="save" type="primary">保存</el-button>
           <el-button @click="back" type="primary">返回</el-button>
           </el-form-item>
    </el-form>
</div>
</template>

<script>
import {
    createPUB_MOrderProduce,
    updatePUB_MOrderProduce,
    findPUB_MOrderProduce
} from "@/api/PUB_MOrderProduce";  //  此处请自行替换地址
import infoList from "@/mixins/infoList";
export default {
  name: "PUB_MOrderProduce",
  mixins: [infoList],
  data() {
    return {
      type: "",formData: {
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
  methods: {
    async save() {
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
      }
    },
    back(){
        this.$router.go(-1)
    }
  },
  async created() {
   // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if(this.$route.query.id){
    const res = await findPUB_MOrderProduce({ ID: this.$route.query.id })
    if(res.code == 0){
       this.formData = res.data.rePMOP
       this.type = "update"
     }
    }else{
     this.type = "create"
   }
  
}
};
</script>

<style>
</style>