<template>
<div>
    <el-form :model="formData" label-position="right" label-width="80px">
             <el-form-item label="deviceId:">
                <el-input v-model="formData.deviceId" clearable placeholder="请输入" ></el-input>
          </el-form-item>
           
             <el-form-item label="温度:">
                  <el-input-number v-model="formData.temp" :precision="2" clearable></el-input-number>
           </el-form-item>
           
             <el-form-item label="上报数据:">
                <el-input v-model="formData.data" clearable placeholder="请输入" ></el-input>
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
    createDeviceLog,
    updateDeviceLog,
    findDeviceLog
} from "@/api/device_log";  //  此处请自行替换地址
import infoList from "@/mixins/infoList";
export default {
  name: "DeviceLog",
  mixins: [infoList],
  data() {
    return {
      type: "",formData: {
            deviceId:"",
            temp:0,
            data:"",
            
      }
    };
  },
  methods: {
    async save() {
      let res;
      switch (this.type) {
        case "create":
          res = await createDeviceLog(this.formData);
          break;
        case "update":
          res = await updateDeviceLog(this.formData);
          break;
        default:
          res = await createDeviceLog(this.formData);
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
    const res = await findDeviceLog({ ID: this.$route.query.id })
    if(res.code == 0){
       this.formData = res.data.redeviceLog
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