<template>
<div>
    <el-form :model="formData" label-position="right" label-width="80px">
             <el-form-item label="LineID:"><el-input v-model.number="formData.lineID" clearable placeholder="请输入"></el-input>
          </el-form-item>
           
             <el-form-item label="LineName:">
                <el-input v-model="formData.lineName" clearable placeholder="请输入" ></el-input>
          </el-form-item>
           
             <el-form-item label="LineType:">
                <el-input v-model="formData.lineType" clearable placeholder="请输入" ></el-input>
          </el-form-item>
           
             <el-form-item label="程式读取路径:">
                <el-input v-model="formData.readFilePath" clearable placeholder="请输入" ></el-input>
          </el-form-item>
           
             <el-form-item label="LineName2:">
                <el-input v-model="formData.lineName2" clearable placeholder="请输入" ></el-input>
          </el-form-item>
           
             <el-form-item label="制令单号:">
                <el-input v-model="formData.mOrderNo" clearable placeholder="请输入" ></el-input>
          </el-form-item>
           
             <el-form-item label="Remark:">
                <el-input v-model="formData.remark" clearable placeholder="请输入" ></el-input>
          </el-form-item>
           
             <el-form-item label="是否报警:">
                <el-switch active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" v-model="formData.warnning" clearable ></el-switch>
          </el-form-item>
           
             <el-form-item label="WarnningStr:">
                <el-input v-model="formData.warnningStr" clearable placeholder="请输入" ></el-input>
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
    createPVS_Base_Line,
    updatePVS_Base_Line,
    findPVS_Base_Line
} from "@/api/PVS_Base_Line";  //  此处请自行替换地址
import infoList from "@/mixins/infoList";
export default {
  name: "PVS_Base_Line",
  mixins: [infoList],
  data() {
    return {
      type: "",formData: {
            lineID:0,
            lineName:"",
            lineType:"",
            readFilePath:"",
            lineName2:"",
            mOrderNo:"",
            remark:"",
            warnning:false,
            warnningStr:"",
            
      }
    };
  },
  methods: {
    async save() {
      let res;
      switch (this.type) {
        case "create":
          res = await createPVS_Base_Line(this.formData);
          break;
        case "update":
          res = await updatePVS_Base_Line(this.formData);
          break;
        default:
          res = await createPVS_Base_Line(this.formData);
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
    const res = await findPVS_Base_Line({ ID: this.$route.query.id })
    if(res.code == 0){
       this.formData = res.data.rePBL
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