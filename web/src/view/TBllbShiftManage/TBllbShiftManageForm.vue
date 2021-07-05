<template>
<div>
    <el-form :model="formData" label-position="right" label-width="80px">
             <el-form-item label="ShiftManageCode字段:">
                <el-input v-model="formData.ShiftManageCode" clearable placeholder="请输入" ></el-input>
          </el-form-item>
           
             <el-form-item label="ShiftManageName字段:">
                <el-input v-model="formData.ShiftManageName" clearable placeholder="请输入" ></el-input>
          </el-form-item>
           
             <el-form-item label="WorkShopCode字段:">
                <el-input v-model="formData.WorkShopCode" clearable placeholder="请输入" ></el-input>
          </el-form-item>
           
             <el-form-item label="StartTime字段:">
                <el-input v-model="formData.StartTime" clearable placeholder="请输入" ></el-input>
          </el-form-item>
           
             <el-form-item label="EndTime字段:">
                <el-input v-model="formData.EndTime" clearable placeholder="请输入" ></el-input>
          </el-form-item>
           
             <el-form-item label="Creator字段:">
                <el-input v-model="formData.Creator" clearable placeholder="请输入" ></el-input>
          </el-form-item>
           
             <el-form-item label="Createtime字段:">
                  <el-date-picker type="date" placeholder="选择日期" v-model="formData.Createtime" clearable></el-date-picker>
           </el-form-item>
           
             <el-form-item label="Updator字段:">
                <el-input v-model="formData.Updator" clearable placeholder="请输入" ></el-input>
          </el-form-item>
           
             <el-form-item label="Updatetime字段:">
                  <el-date-picker type="date" placeholder="选择日期" v-model="formData.Updatetime" clearable></el-date-picker>
           </el-form-item>
           
             <el-form-item label="TotalMente字段:">
                  <el-input-number v-model="formData.TotalMente" :precision="2" clearable></el-input-number>
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
    createTBllbShiftManage,
    updateTBllbShiftManage,
    findTBllbShiftManage
} from "@/api/TBllbShiftManage";  //  此处请自行替换地址
import infoList from "@/mixins/infoList";
export default {
  name: "TBllbShiftManage",
  mixins: [infoList],
  data() {
    return {
      type: "",formData: {
            ShiftManageCode:"",
            ShiftManageName:"",
            WorkShopCode:"",
            StartTime:"",
            EndTime:"",
            Creator:"",
            Createtime:new Date(),
            Updator:"",
            Updatetime:new Date(),
            TotalMente:0,
            
      }
    };
  },
  methods: {
    async save() {
      let res;
      switch (this.type) {
        case "create":
          res = await createTBllbShiftManage(this.formData);
          break;
        case "update":
          res = await updateTBllbShiftManage(this.formData);
          break;
        default:
          res = await createTBllbShiftManage(this.formData);
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
    const res = await findTBllbShiftManage({ ID: this.$route.query.id })
    if(res.code == 0){
       this.formData = res.data.reTSM
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