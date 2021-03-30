<template>
<div>
    <el-form :model="formData" label-position="right" label-width="80px">
             <el-form-item label="目标产量:"><el-input v-model.number="formData.dest" clearable placeholder="请输入"></el-input>
          </el-form-item>
           
             <el-form-item label="实际产量逗号隔开:">
                <el-input v-model="formData.real" clearable placeholder="请输入" ></el-input>
          </el-form-item>
           
             <el-form-item label="小时数:"><el-input v-model.number="formData.hour" clearable placeholder="请输入"></el-input>
          </el-form-item>
           
             <el-form-item label="开始时间:">
                  <el-date-picker type="date" placeholder="选择日期" v-model="formData.beginTime" clearable></el-date-picker>
           </el-form-item>
           
             <el-form-item label="LineID:"><el-input v-model.number="formData.lineID" clearable placeholder="请输入"></el-input>
          </el-form-item>
           
             <el-form-item label="制令单目标数量:">
                <el-switch active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" v-model="formData.manuDest" clearable ></el-switch>
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
    createVM_TimeOutput,
    updateVM_TimeOutput,
    findVM_TimeOutput
} from "@/api/VM_TimeOutput";  //  此处请自行替换地址
import infoList from "@/mixins/infoList";
export default {
  name: "VM_TimeOutput",
  mixins: [infoList],
  data() {
    return {
      type: "",formData: {
            dest:0,
            real:"",
            hour:0,
            beginTime:new Date(),
            lineID:0,
            manuDest:false,
            
      }
    };
  },
  methods: {
    async save() {
      let res;
      switch (this.type) {
        case "create":
          res = await createVM_TimeOutput(this.formData);
          break;
        case "update":
          res = await updateVM_TimeOutput(this.formData);
          break;
        default:
          res = await createVM_TimeOutput(this.formData);
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
    const res = await findVM_TimeOutput({ ID: this.$route.query.id })
    if(res.code == 0){
       this.formData = res.data.reVTO
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