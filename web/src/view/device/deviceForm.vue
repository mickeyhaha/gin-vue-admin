<template>
<div>
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
           <el-form-item>
           <el-button @click="save" type="primary">保存</el-button>
           <el-button @click="back" type="primary">返回</el-button>
           </el-form-item>
    </el-form>
</div>
</template>

<script>
import {
    createDevice,
    updateDevice,
    findDevice
} from "@/api/device";  //  此处请自行替换地址
import infoList from "@/mixins/infoList";
export default {
  name: "Device",
  mixins: [infoList],
  data() {
    return {
      type: "",formData: {
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
  methods: {
    async save() {
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
      }
    },
    back(){
        this.$router.go(-1)
    }
  },
  async created() {
   // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if(this.$route.query.id){
    const res = await findDevice({ ID: this.$route.query.id })
    if(res.code == 0){
       this.formData = res.data.redev
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