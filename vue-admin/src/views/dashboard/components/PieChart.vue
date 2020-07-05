<template>
  <div :class="className" :style="{height:height,width:width}" />
</template>

<script>
import echarts from 'echarts'
require('echarts/theme/macarons') // echarts theme
import resize from './mixins/resize'

export default {
  mixins: [resize],
  props: {
    className: {
      type: String,
      default: 'chart'
    },
    width: {
      type: String,
      default: '100%'
    },
    height: {
      type: String,
      default: '300px'
    },
    successErrorCount: {
      type: Array,
      default: null
    }
  },
  data() {
    return {
      chart: null
    }
  },
  watch:{
    successErrorCount(newValue,oldValue){
      this.successErrorCount = newValue
      this.initChart()
    }
  },
  mounted() {
  },
  beforeDestroy() {
    if (!this.chart) {
      return
    }
    this.chart.dispose()
    this.chart = null
  },
  methods: {
    initChart() {
      this.chart = echarts.init(this.$el, 'macarons')
      this.chart.setOption({
        title:{
          text:'部署通过率',
          left:'center',
          top:'40%',
          padding:[24,0],
          textStyle:{
            color:'#fff',
            align:'center'
          }
        },
        tooltip: {
          trigger: 'item',
          formatter: '{a} <br/>{b} : {c} ({d}%)'
        },
        series: [
          {
            name:'成功失败',
            type:'pie',
            radius:'80%',
            data: this.successErrorCount,
            animationEasing: 'cubicInOut',
            animationDuration: 2600
          }
        ]
      })
    }
  }
}
</script>
