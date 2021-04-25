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
           
             <el-form-item label="TableNo字段:">
                <el-input v-model="formData.TableNo" clearable placeholder="请输入" ></el-input>
          </el-form-item>
           
             <el-form-item label="PickPos字段:"><el-input v-model.number="formData.PickPos" clearable placeholder="请输入"></el-input>
          </el-form-item>
           
             <el-form-item label="SlotNo字段:">
                <el-input v-model="formData.SlotNo" clearable placeholder="请输入" ></el-input>
          </el-form-item>
           
             <el-form-item label="PickError字段:">
                 <el-select v-model="formData.PickError" placeholder="请选择" clearable>
                     <el-option v-for="(item,key) in PickErrorOptions" :key="key" :label="item.label" :value="item.value"></el-option>
                 </el-select>
          </el-form-item>
           
             <el-form-item label="IdentError字段:">
                 <el-select v-model="formData.IdentError" placeholder="请选择" clearable>
                     <el-option v-for="(item,key) in IdentErrorOptions" :key="key" :label="item.label" :value="item.value"></el-option>
                 </el-select>
          </el-form-item>
           
             <el-form-item label="OtherError字段:">
                 <el-select v-model="formData.OtherError" placeholder="请选择" clearable>
                     <el-option v-for="(item,key) in OtherErrorOptions" :key="key" :label="item.label" :value="item.value"></el-option>
                 </el-select>
          </el-form-item>
           
             <el-form-item label="CreateTime字段:">
                  <el-date-picker type="date" placeholder="选择日期" v-model="formData.CreateTime" clearable></el-date-picker>
           </el-form-item>
           
             <el-form-item label="Remark字段:">
                <el-input v-model="formData.Remark" clearable placeholder="请输入" ></el-input>
          </el-form-item>
           
             <el-form-item label="FeederNo字段:">
                <el-input v-model="formData.FeederNo" clearable placeholder="请输入" ></el-input>
          </el-form-item>
           
             <el-form-item label="MatrCode字段:">
                <el-input v-model="formData.MatrCode" clearable placeholder="请输入" ></el-input>
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
    createDCSSMTReject,
    updateDCSSMTReject,
    findDCSSMTReject
} from "@/api/DCSSMTReject";  //  此处请自行替换地址
import infoList from "@/mixins/infoList";
export default {
  name: "DCSSMTReject",
  mixins: [infoList],
  data() {
    return {
      type: "",
      PickErrorOptions:[],
          
      IdentErrorOptions:[],
          
      OtherErrorOptions:[],
          formData: {
            LineName:"",
            MachineCode:0,
            MachineName:"",
            Lane:0,
            TableNo:"",
            PickPos:0,
            SlotNo:"",
            PickError:0,
            IdentError:0,
            OtherError:0,
            CreateTime:new Date(),
            Remark:"",
            FeederNo:"",
            MatrCode:"",
            AddTime:new Date(),
            
      }
    };
  },
  methods: {
    async save() {
      let res;
      switch (this.type) {
        case "create":
          res = await createDCSSMTReject(this.formData);
          break;
        case "update":
          res = await updateDCSSMTReject(this.formData);
          break;
        default:
          res = await createDCSSMTReject(this.formData);
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
    const res = await findDCSSMTReject({ ID: this.$route.query.id })
    if(res.code == 0){
       this.formData = res.data.reDSR
       this.type = "update"
     }
    }else{
     this.type = "create"
   }
  
    await this.getDict("PickError");
    
    await this.getDict("IdentError");
    
    await this.getDict("OtherError");
    
}
};
</script>

<style>
</style>