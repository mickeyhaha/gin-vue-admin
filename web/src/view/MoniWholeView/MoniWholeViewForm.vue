<template>
<div>
    <el-form :model="formData" label-position="right" label-width="80px">
             <el-form-item label="LineID:"><el-input v-model.number="formData.lineID" clearable placeholder="请输入"></el-input>
          </el-form-item>
           
             <el-form-item label="机器:"><el-input v-model.number="formData.machineCode" clearable placeholder="请输入"></el-input>
          </el-form-item>
           
             <el-form-item label="台车:">
                <el-input v-model="formData.tableNo" clearable placeholder="请输入" ></el-input>
          </el-form-item>
           
             <el-form-item label="BCT:">
                <el-input v-model="formData.bCT" clearable placeholder="请输入" ></el-input>
          </el-form-item>
           
             <el-form-item label="TrackNo:"><el-input v-model.number="formData.trackNo" clearable placeholder="请输入"></el-input>
          </el-form-item>
           
             <el-form-item label="Division:"><el-input v-model.number="formData.division" clearable placeholder="请输入"></el-input>
          </el-form-item>
           
             <el-form-item label="吸取位置:"><el-input v-model.number="formData.pickPos" clearable placeholder="请输入"></el-input>
          </el-form-item>
           
             <el-form-item label="FeederType:">
                <el-input v-model="formData.feederType" clearable placeholder="请输入" ></el-input>
          </el-form-item>
           
             <el-form-item label="FeederNo:">
                <el-input v-model="formData.feederNo" clearable placeholder="请输入" ></el-input>
          </el-form-item>
           
             <el-form-item label="料号:">
                <el-input v-model="formData.matrCode" clearable placeholder="请输入" ></el-input>
          </el-form-item>
           
             <el-form-item label="PartStatus:"><el-input v-model.number="formData.partStatus" clearable placeholder="请输入"></el-input>
          </el-form-item>
           
             <el-form-item label="CycleTime:">
                  <el-input-number v-model="formData.cycleTime" :precision="2" clearable></el-input-number>
           </el-form-item>
           
             <el-form-item label="PartPos:">
                <el-input v-model="formData.partPos" clearable placeholder="请输入" ></el-input>
          </el-form-item>
           
             <el-form-item label="PlacedQty:"><el-input v-model.number="formData.placedQty" clearable placeholder="请输入"></el-input>
          </el-form-item>
           
             <el-form-item label="吸取错误:">
                 <el-select v-model="formData.pickError" placeholder="请选择" clearable>
                     <el-option v-for="(item,key) in intOptions" :key="key" :label="item.label" :value="item.value"></el-option>
                 </el-select>
          </el-form-item>
           
             <el-form-item label="IdentError:"><el-input v-model.number="formData.identError" clearable placeholder="请输入"></el-input>
          </el-form-item>
           
             <el-form-item label="OtherError:"><el-input v-model.number="formData.otherError" clearable placeholder="请输入"></el-input>
          </el-form-item>
           
             <el-form-item label="TotalConsume:"><el-input v-model.number="formData.totalConsume" clearable placeholder="请输入"></el-input>
          </el-form-item>
           
             <el-form-item label="UniqueCode:">
                <el-input v-model="formData.uniqueCode" clearable placeholder="请输入" ></el-input>
          </el-form-item>
           
             <el-form-item label="LeftQty:"><el-input v-model.number="formData.leftQty" clearable placeholder="请输入"></el-input>
          </el-form-item>
           
             <el-form-item label="StandbyCode:">
                <el-input v-model="formData.standbyCode" clearable placeholder="请输入" ></el-input>
          </el-form-item>
           
             <el-form-item label="StandbyQty:"><el-input v-model.number="formData.standbyQty" clearable placeholder="请输入"></el-input>
          </el-form-item>
           
             <el-form-item label="TotalLeftQty:"><el-input v-model.number="formData.totalLeftQty" clearable placeholder="请输入"></el-input>
          </el-form-item>
           
             <el-form-item label="StopTrail:">
                <el-switch active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" v-model="formData.stopTrail" clearable ></el-switch>
          </el-form-item>
           
             <el-form-item label="TotalError:"><el-input v-model.number="formData.totalError" clearable placeholder="请输入"></el-input>
          </el-form-item>
           
             <el-form-item label="ChangeTime:">
                  <el-date-picker type="date" placeholder="选择日期" v-model="formData.changeTime" clearable></el-date-picker>
           </el-form-item>
           
             <el-form-item label="RejectRate:">
                  <el-input-number v-model="formData.rejectRate" :precision="2" clearable></el-input-number>
           </el-form-item>
           
             <el-form-item label="LeftTime:">
                  <el-input-number v-model="formData.leftTime" :precision="2" clearable></el-input-number>
           </el-form-item>
           
             <el-form-item label="UseLevel:"><el-input v-model.number="formData.useLevel" clearable placeholder="请输入"></el-input>
          </el-form-item>
           
             <el-form-item label="Scanner:">
                <el-input v-model="formData.scanner" clearable placeholder="请输入" ></el-input>
          </el-form-item>
           
             <el-form-item label="ScanTime:">
                  <el-date-picker type="date" placeholder="选择日期" v-model="formData.scanTime" clearable></el-date-picker>
           </el-form-item>
           
             <el-form-item label="Remark:">
                <el-input v-model="formData.remark" clearable placeholder="请输入" ></el-input>
          </el-form-item>
           
             <el-form-item label="IsTray:">
                <el-switch active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" v-model="formData.isTray" clearable ></el-switch>
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
    createMoniWholeView,
    updateMoniWholeView,
    findMoniWholeView
} from "@/api/MoniWholeView";  //  此处请自行替换地址
import infoList from "@/mixins/infoList";
export default {
  name: "MoniWholeView",
  mixins: [infoList],
  data() {
    return {
      type: "",
      intOptions:[],
          formData: {
            lineID:0,
            machineCode:0,
            tableNo:"",
            bCT:"",
            trackNo:0,
            division:0,
            pickPos:0,
            feederType:"",
            feederNo:"",
            matrCode:"",
            partStatus:0,
            cycleTime:0,
            partPos:"",
            placedQty:0,
            pickError:0,
            identError:0,
            otherError:0,
            totalConsume:0,
            uniqueCode:"",
            leftQty:0,
            standbyCode:"",
            standbyQty:0,
            totalLeftQty:0,
            stopTrail:false,
            totalError:0,
            changeTime:new Date(),
            rejectRate:0,
            leftTime:0,
            useLevel:0,
            scanner:"",
            scanTime:new Date(),
            remark:"",
            isTray:false,
            
      }
    };
  },
  methods: {
    async save() {
      let res;
      switch (this.type) {
        case "create":
          res = await createMoniWholeView(this.formData);
          break;
        case "update":
          res = await updateMoniWholeView(this.formData);
          break;
        default:
          res = await createMoniWholeView(this.formData);
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
    const res = await findMoniWholeView({ ID: this.$route.query.id })
    if(res.code == 0){
       this.formData = res.data.reMWV
       this.type = "update"
     }
    }else{
     this.type = "create"
   }
  
    await this.getDict("int");
    
}
};
</script>

<style>
</style>