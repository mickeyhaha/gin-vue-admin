<template>
<div>
    <el-form :model="formData" label-position="right" label-width="80px">
             <el-form-item label="Count:"><el-input v-model.number="formData.count" clearable placeholder="请输入"></el-input>
          </el-form-item>
           
             <el-form-item label="ErrCount:"><el-input v-model.number="formData.errCount" clearable placeholder="请输入"></el-input>
          </el-form-item>
           
             <el-form-item label="CreateTime:">
                  <el-date-picker type="date" placeholder="选择日期" v-model="formData.createTime" clearable></el-date-picker>
           </el-form-item>
           
             <el-form-item label="LineID:"><el-input v-model.number="formData.lineID" clearable placeholder="请输入"></el-input>
          </el-form-item>
           
             <el-form-item label="OrderNo:">
                <el-input v-model="formData.orderNo" clearable placeholder="请输入" ></el-input>
          </el-form-item>
           
             <el-form-item label="IssueName:">
                <el-input v-model="formData.issueName" clearable placeholder="请输入" ></el-input>
          </el-form-item>
           
             <el-form-item label="Result:"><el-input v-model.number="formData.result" clearable placeholder="请输入"></el-input>
          </el-form-item>
           
             <el-form-item label="SPIID:"><el-input v-model.number="formData.sPIID" clearable placeholder="请输入"></el-input>
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
    createTS_SPI_CNT,
    updateTS_SPI_CNT,
    findTS_SPI_CNT
} from "@/api/TS_SPI_CNT";  //  此处请自行替换地址
import infoList from "@/mixins/infoList";
export default {
  name: "TS_SPI_CNT",
  mixins: [infoList],
  data() {
    return {
      type: "",formData: {
            count:0,
            errCount:0,
            createTime:new Date(),
            lineID:0,
            orderNo:"",
            issueName:"",
            result:0,
            sPIID:0,
            
      }
    };
  },
  methods: {
    async save() {
      let res;
      switch (this.type) {
        case "create":
          res = await createTS_SPI_CNT(this.formData);
          break;
        case "update":
          res = await updateTS_SPI_CNT(this.formData);
          break;
        default:
          res = await createTS_SPI_CNT(this.formData);
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
    const res = await findTS_SPI_CNT({ ID: this.$route.query.id })
    if(res.code == 0){
       this.formData = res.data.reTSC
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