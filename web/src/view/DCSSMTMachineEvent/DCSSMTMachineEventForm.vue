<template>
<div>
    <el-form :model="formData" label-position="right" label-width="80px">
             <el-form-item label="LineName字段:">
                <el-input v-model="formData.LineName" clearable placeholder="请输入" ></el-input>
          </el-form-item>
           
             <el-form-item label="MachineCode字段:"><el-input v-model.number="formData.MachineCode" clearable placeholder="请输入"></el-input>
          </el-form-item>
           
             <el-form-item label="MachineName字段:">
                <el-input v-model="formData.MachineName" clearable placeholder="请输入" ></el-input>
          </el-form-item>
           
             <el-form-item label="EventName字段:">
                <el-input v-model="formData.EventName" clearable placeholder="请输入" ></el-input>
          </el-form-item>
           
             <el-form-item label="EventRemark字段:">
                <el-input v-model="formData.EventRemark" clearable placeholder="请输入" ></el-input>
          </el-form-item>
           
             <el-form-item label="CreateTime字段:">
                  <el-date-picker type="date" placeholder="选择日期" v-model="formData.CreateTime" clearable></el-date-picker>
           </el-form-item>
           
             <el-form-item label="TableNo字段:">
                <el-input v-model="formData.TableNo" clearable placeholder="请输入" ></el-input>
          </el-form-item>
           
             <el-form-item label="Lane字段:"><el-input v-model.number="formData.Lane" clearable placeholder="请输入"></el-input>
          </el-form-item>
           
             <el-form-item label="AddTime字段:">
                  <el-date-picker type="date" placeholder="选择日期" v-model="formData.AddTime" clearable></el-date-picker>
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
    createDCSSMTMachineEvent,
    updateDCSSMTMachineEvent,
    findDCSSMTMachineEvent
} from "@/api/DCSSMTMachineEvent";  //  此处请自行替换地址
import infoList from "@/mixins/infoList";
export default {
  name: "DCSSMTMachineEvent",
  mixins: [infoList],
  data() {
    return {
      type: "",formData: {
            LineName:"",
            MachineCode:0,
            MachineName:"",
            EventName:"",
            EventRemark:"",
            CreateTime:new Date(),
            TableNo:"",
            Lane:0,
            AddTime:new Date(),
            
      }
    };
  },
  methods: {
    async save() {
      let res;
      switch (this.type) {
        case "create":
          res = await createDCSSMTMachineEvent(this.formData);
          break;
        case "update":
          res = await updateDCSSMTMachineEvent(this.formData);
          break;
        default:
          res = await createDCSSMTMachineEvent(this.formData);
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
    const res = await findDCSSMTMachineEvent({ ID: this.$route.query.id })
    if(res.code == 0){
       this.formData = res.data.reDSME
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