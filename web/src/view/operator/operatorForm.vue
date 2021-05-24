<template>
<div>
    <el-form :model="formData" label-position="right" label-width="80px">
             <el-form-item label="操作员ID:">
                <el-input v-model="formData.operator_id" clearable placeholder="请输入" ></el-input>
          </el-form-item>
           
             <el-form-item label="用户名:">
                <el-input v-model="formData.operator_name" clearable placeholder="请输入" ></el-input>
          </el-form-item>
           
             <el-form-item label="状态:">
                 <el-select v-model="formData.status" placeholder="请选择" clearable>
                     <el-option v-for="(item,key) in statusOptions" :key="key" :label="item.label" :value="item.value"></el-option>
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
    createOperator,
    updateOperator,
    findOperator
} from "@/api/operator";  //  此处请自行替换地址
import infoList from "@/mixins/infoList";
export default {
  name: "Operator",
  mixins: [infoList],
  data() {
    return {
      type: "",
      statusOptions:[],
          formData: {
            operator_id:"",
            operator_name:"",
            status:0,
            
      }
    };
  },
  methods: {
    async save() {
      let res;
      switch (this.type) {
        case "create":
          res = await createOperator(this.formData);
          break;
        case "update":
          res = await updateOperator(this.formData);
          break;
        default:
          res = await createOperator(this.formData);
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
    const res = await findOperator({ ID: this.$route.query.id })
    if(res.code == 0){
       this.formData = res.data.reopt
       this.type = "update"
     }
    }else{
     this.type = "create"
   }
  
    await this.getDict("status");
    
}
};
</script>

<style>
</style>