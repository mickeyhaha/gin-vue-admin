<template>
<div>
    <el-form :model="formData" label-position="right" label-width="80px">
             <el-form-item label="ID字段:"><el-input v-model.number="formData.ID" clearable placeholder="请输入"></el-input>
          </el-form-item>
           
             <el-form-item label="LineID字段:"><el-input v-model.number="formData.LineID" clearable placeholder="请输入"></el-input>
          </el-form-item>
           
             <el-form-item label="MachineType字段:">
                <el-input v-model="formData.MachineType" clearable placeholder="请输入" ></el-input>
          </el-form-item>
           
             <el-form-item label="OrderNo字段:">
                <el-input v-model="formData.OrderNo" clearable placeholder="请输入" ></el-input>
          </el-form-item>
           
             <el-form-item label="Program字段:">
                <el-input v-model="formData.Program" clearable placeholder="请输入" ></el-input>
          </el-form-item>
           
             <el-form-item label="PCBCode字段:">
                <el-input v-model="formData.PCBCode" clearable placeholder="请输入" ></el-input>
          </el-form-item>
           
             <el-form-item label="Result字段:"></el-form-item>
           
             <el-form-item label="CreateTime字段:">
                  <el-date-picker type="date" placeholder="选择日期" v-model="formData.CreateTime" clearable></el-date-picker>
           </el-form-item>
           
             <el-form-item label="SysOwner字段:">
                <el-input v-model="formData.SysOwner" clearable placeholder="请输入" ></el-input>
          </el-form-item>
           
             <el-form-item label="Ext1字段:">
                <el-input v-model="formData.Ext1" clearable placeholder="请输入" ></el-input>
          </el-form-item>
           
             <el-form-item label="Ext2字段:">
                <el-input v-model="formData.Ext2" clearable placeholder="请输入" ></el-input>
          </el-form-item>
           
             <el-form-item label="Ext3字段:">
                <el-input v-model="formData.Ext3" clearable placeholder="请输入" ></el-input>
          </el-form-item>
           
             <el-form-item label="Ext4字段:">
                <el-input v-model="formData.Ext4" clearable placeholder="请输入" ></el-input>
          </el-form-item>
           
             <el-form-item label="Ext5字段:">
                <el-input v-model="formData.Ext5" clearable placeholder="请输入" ></el-input>
          </el-form-item>
           
             <el-form-item label="Ext6字段:">
                <el-input v-model="formData.Ext6" clearable placeholder="请输入" ></el-input>
          </el-form-item>
           
             <el-form-item label="Ext7字段:">
                <el-input v-model="formData.Ext7" clearable placeholder="请输入" ></el-input>
          </el-form-item>
           
             <el-form-item label="Ext8字段:">
                <el-input v-model="formData.Ext8" clearable placeholder="请输入" ></el-input>
          </el-form-item>
           
             <el-form-item label="Ext9字段:">
                <el-input v-model="formData.Ext9" clearable placeholder="请输入" ></el-input>
          </el-form-item>
           
             <el-form-item label="Ext10字段:">
                <el-input v-model="formData.Ext10" clearable placeholder="请输入" ></el-input>
          </el-form-item>
           
             <el-form-item label="InsType字段:"><el-input v-model.number="formData.InsType" clearable placeholder="请输入"></el-input>
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
    createTBllbAOITba,
    updateTBllbAOITba,
    findTBllbAOITba
} from "@/api/TBllbAOITba";  //  此处请自行替换地址
import infoList from "@/mixins/infoList";
export default {
  name: "TBllbAOITba",
  mixins: [infoList],
  data() {
    return {
      type: "",formData: {
            ID:0,
            LineID:0,
            MachineType:"",
            OrderNo:"",
            Program:"",
            PCBCode:"",
            CreateTime:new Date(),
            SysOwner:"",
            Ext1:"",
            Ext2:"",
            Ext3:"",
            Ext4:"",
            Ext5:"",
            Ext6:"",
            Ext7:"",
            Ext8:"",
            Ext9:"",
            Ext10:"",
            InsType:0,
            
      }
    };
  },
  methods: {
    async save() {
      let res;
      switch (this.type) {
        case "create":
          res = await createTBllbAOITba(this.formData);
          break;
        case "update":
          res = await updateTBllbAOITba(this.formData);
          break;
        default:
          res = await createTBllbAOITba(this.formData);
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
    const res = await findTBllbAOITba({ ID: this.$route.query.id })
    if(res.code == 0){
       this.formData = res.data.reTBAT
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