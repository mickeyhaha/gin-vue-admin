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
           
             <el-form-item label="Lane字段:"><el-input v-model.number="formData.Lane" clearable placeholder="请输入"></el-input>
          </el-form-item>
           
             <el-form-item label="SetupName字段:">
                <el-input v-model="formData.SetupName" clearable placeholder="请输入" ></el-input>
          </el-form-item>
           
             <el-form-item label="TimeCode字段:">
                <el-input v-model="formData.TimeCode" clearable placeholder="请输入" ></el-input>
          </el-form-item>
           
             <el-form-item label="TimeValue字段:">
                  <el-input-number v-model="formData.TimeValue" :precision="2" clearable></el-input-number>
           </el-form-item>
           
             <el-form-item label="CreateTime字段:">
                  <el-date-picker type="date" placeholder="选择日期" v-model="formData.CreateTime" clearable></el-date-picker>
           </el-form-item>
           
             <el-form-item label="Remark字段:">
                <el-input v-model="formData.Remark" clearable placeholder="请输入" ></el-input>
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
    createDCSSMTRunTime,
    updateDCSSMTRunTime,
    findDCSSMTRunTime
} from "@/api/DCSSMTRunTime";  //  此处请自行替换地址
import infoList from "@/mixins/infoList";
export default {
  name: "DCSSMTRunTime",
  mixins: [infoList],
  data() {
    return {
      type: "",formData: {
            LineName:"",
            MachineCode:0,
            MachineName:"",
            Lane:0,
            SetupName:"",
            TimeCode:"",
            TimeValue:0,
            CreateTime:new Date(),
            Remark:"",
            AddTime:new Date(),
            
      }
    };
  },
  methods: {
    async save() {
      let res;
      switch (this.type) {
        case "create":
          res = await createDCSSMTRunTime(this.formData);
          break;
        case "update":
          res = await updateDCSSMTRunTime(this.formData);
          break;
        default:
          res = await createDCSSMTRunTime(this.formData);
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
    const res = await findDCSSMTRunTime({ ID: this.$route.query.id })
    if(res.code == 0){
       this.formData = res.data.reDSRT
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