<template>
  <div :id="chartId" :style="{width: '300px', height: '300px',display:'inline-block'}"></div>
</template>

<script>
// 引入基本模板
let echarts = require("echarts");

export default {
  name: "lineChart",
  props: ["chartId", "chartData"],
  data() {
    return {};
  },
  mounted() {
    this.drawLine();
  },
  methods: {
    drawLine() {
      // 基于准备好的dom，初始化echarts实例
      let myChart = echarts.init(document.getElementById(this.chartId));
      var option = {
        xAxis: {
          type: "category",
          data: Object.keys(this.chartData)
        },
        yAxis: {
          type: "value"
        },
        series: [
          {
            type: "line",
            data: Object.values(this.chartData)
          }
        ]
      };

      // 绘制图表
      myChart.setOption(option);
    }
  }
};
</script>
