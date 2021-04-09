<template>
  <div id="order-info">
      <el-row>
        <el-col :span="col_span.col1" class="block-top-bottom-content">
          线体:
        </el-col>
        <el-col :span="col_span.col2"  class="block-top-bottom-content">
          <el-select v-model="selectValue" placeholder="请选择" @change="selectChanged" size="small">
            <el-option v-for="(item, i) in lineOptions" :key="i" :label="item.label" :value="item.value"></el-option>
          </el-select>
        </el-col>
      </el-row>
      <el-row>
        <el-col :span="col_span.col1" class="block-top-bottom-content">
          工单:
        </el-col>
        <el-col :span="col_span.col2"  class="block-top-bottom-content">
          {{selectItem.WorkOrderNo}}
        </el-col>
        <el-col :span="col_span.col3"  class="block-top-bottom-content">
          计划产能(H):
        </el-col>
        <el-col :span="col_span.col4"  class="block-top-bottom-content">
          {{selectItem.Qty}}
        </el-col>
      </el-row>
      <el-row>
        <el-col :span="col_span.col1" class="block-top-bottom-content">
          产品:
        </el-col>
        <el-col :span="col_span.col2"  class="block-top-bottom-content">
           {{selectItem.Product}}
        </el-col>
        <el-col :span="col_span.col3"  class="block-top-bottom-content">
          产品名称:
        </el-col>
        <el-col :span="col_span.col4"  class="block-top-bottom-content">
           {{selectItem.Product}}
        </el-col>
      </el-row>
      <el-row>
        <el-col :span="col_span.col1" class="block-top-bottom-content">
          制令单:
        </el-col>
        <el-col :span="col_span.col2"  class="block-top-bottom-content">
          {{selectItem.MOrderNo}}
        </el-col>
        <el-col :span="col_span.col3"  class="block-top-bottom-content">
          面别:
        </el-col>
        <el-col :span="col_span.col4"  class="block-top-bottom-content">
          {{selectItem.PasteSide}}
        </el-col>
      </el-row>
  </div>
</template>

<script>
import {
  getLineCurrOrderList
} from "@/api/PUBMOrderProduce";
import infoList from "@/mixins/infoList";
import { store } from '@/store/index';

export default {
  name: 'OrderInfo',
  mixins: [infoList],
  data () {
    return {
      col_span: {
        col1: 4,
        col2: 8,
        col3: 6,
        col4: 6,
      },
      selectValue: '',
      listApi: getLineCurrOrderList,
      lineOptions: [],
      selectItem: {},
    }
  },
  methods: {
    createData () {
      this.getTableData().then(() => {
        // debugger
        const options = [];
        for (let index = 0; index < this.tableData.length; index++) {
          const element = this.tableData[index]
          options.push({
            'label': element.LineName,
            'value': element.ID,
            'LineName': element.LineName,
            'WorkOrderNo': element.WorkOrderNo,
            'MOrderNo': element.MOrderNo,
            'Product': element.Product,
            'Qty': element.Qty,
            'QtyCompleted': element.QtyCompleted,
            'PasteSide': element.PasteSide,
          });
        }
        this.lineOptions = options;
        this.selectValue = options[0].label;
        this.selectItem = options[0];
        this.$store.commit('update',['selectedLine',this.selectItem]);
      });
      
    },

    selectChanged(value) {
      for (let index = 0; index < this.lineOptions.length; index++) {
        const element = this.lineOptions[index];
        if (element.value == value) {
          this.selectItem = element
        }
      }
      this.$store.commit('update',['selectedLine',this.selectItem]);
    },
  },

  mounted () {
    const { createData } = this

    createData()

    //setInterval(createData, 3000)
  }
}
</script>

<style lang="less">
#order-info {
  width: 100%;
  height: 30%;
  box-shadow: 0 0 3px blue;
  display: flex;
  flex-direction: column;
  background-color: rgba(6, 30, 93, 0.5);
  border-top: 2px solid rgba(1, 153, 209, .5);
  box-sizing: border-box;
  padding: 0px 0px;
  margin-bottom: 10px;

  .dv-conical-column-chart {
    flex: 1;
  }
}
</style>
