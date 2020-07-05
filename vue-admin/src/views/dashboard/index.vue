<template>
  <div class="dashboard-editor-container">
    <template>
        <div id="banner">
          <!--动态将图片轮播图的容器高度设置成与图片一致-->
          <el-carousel :height="bannerHeight + 'px'" ref="carousel" @click.native="linkTo">
            <!--遍历图片地址,动态生成轮播图-->
            <el-carousel-item v-for="item in img_list" :key="item">
              <img :src="item" alt="">
            </el-carousel-item>
          </el-carousel>
        </div>
    </template>
    <el-divider></el-divider>
    <el-row :gutter="32">
      <el-col :xs="24" :sm="24" :lg="8">

        <div class="chart-wrapper">
          <pie-chart :successErrorCount="successErrorCount" />
        </div>
      </el-col>
      <el-col :xs="24" :sm="24" :lg="8">

        <div class="chart-wrapper">
          <bar-chart :cycleCount="cycleCount" />
        </div>
      </el-col>
      <el-col :xs="24" :sm="24" :lg="8">
        <el-table v-loading="listLoading" :data="list" border fit highlight-current-row
          style="width: 100%">
          <el-table-column align="center" label="最近发布" min-width="20%">
            <template slot-scope="scope">
              <span>{{ scope.row.name }} </span>
            </template>
          </el-table-column>
          <el-table-column  align="center" label="操作" min-width="10%">
            <template slot-scope="scope">
              <router-link :to="'/deploy/index/'+scope.row.ID">
                <el-button type="primary" :size="size" icon="el-icon-caret-right">
                  部署
                </el-button>
              </router-link>
            </template>
          </el-table-column>
        </el-table>
      </el-col>
    </el-row>

    <el-dialog
      title="提示"
      :visible.sync="dialogVisible"
      width="100%"
      :before-close="handleClose">
      <span><img :src="imgSrc"></span>
    </el-dialog>
  </div>
</template>

<script>
  import { recentList } from '@/api/application'
  import { getOkErrorCount, getCycleCount } from '@/api/deploy'
  import PieChart from './components/PieChart'
  import BarChart from './components/BarChart'

export default {
  name: 'DashboardAdmin',
  components: {
    PieChart,
    BarChart
  },
  data() {
    return {
      size: 'mini',
      listLoading: true,
      list: null,
      successErrorCount: null,
      cycleCount: null,
      dialogVisible: false,
      imgSrc: '',
      img_list:[
        require("@/assets/image/111.png"),
        require("@/assets/image/222.png")
      ],
      // 图片父容器高度
      bannerHeight :1000,
      // 浏览器宽度
      screenWidth :0,
      recentListQuery: {
        limit: 5
      },
      // cycle表示取值周期，比如，以周，月，年为取值周期
      // count表示取值个数，比如，取4周，4月，4年
      cycleCountListQuery: {
        cycle: 7,
        count: 4
      },
    }
  },
  created() {
    this.getRecentList(),
    this.getOkErrorList(),
    this.getCycleCountList()
  },
  methods: {
    setSize:function () {
      // 通过浏览器宽度(图片宽度)计算高度
      this.bannerHeight = 215 / 1920 * this.screenWidth;
    },
    linkTo () {
      let activeIndex = this.$refs.carousel.activeIndex
      this.imgSrc = this.img_list[activeIndex]
      this.dialogVisible = true
    },
     handleClose(done) {
      this.$confirm('确认关闭？')
        .then(_ => {
          done();
        })
        .catch(_ => {});
    },
    getRecentList() {
      this.listLoading = true
        recentList(this.recentListQuery).then(response => {
          this.list = response.data
          this.listLoading = false
        })
    },
    getRecentList() {
      this.listLoading = true
        recentList(this.recentListQuery).then(response => {
          this.list = response.data
          this.listLoading = false
        })
    },
    getOkErrorList() {
      this.listLoading = true
        getOkErrorCount().then(response => {
          //console.log(response)
          this.successErrorCount = response.data
          this.listLoading = false
        })
    },
    getCycleCountList() {
      this.listLoading = true
        getCycleCount(this.cycleCountListQuery).then(response => {
          this.cycleCount = response.data
        })
    },
    handleSetLineChartData(type) {
      this.lineChartData = lineChartData[type]
    }
  },
  mounted() {
      // 首次加载时,需要调用一次
      this.screenWidth =  window.innerWidth;
      this.setSize();
      // 窗口大小发生改变时,调用一次
      window.onresize = () =>{
      this.screenWidth =  window.innerWidth;
      this.setSize();
    }
  }
}
</script>

<style lang="scss" scoped>

  .el-carousel__item h3 {
    color: #475669;
    font-size: 14px;
    opacity: 0.75;
    line-height: 200px;
    margin: 0;
  }

  .el-carousel__item:nth-child(2n) {
     background-color: #99a9bf;
  }

  .el-carousel__item:nth-child(2n+1) {
     background-color: #d3dce6;
  }
  img{
    /*设置图片宽度和浏览器宽度一致*/
    width: 100%;
    height: inherit;
  }
  .dashboard-editor-container {
    padding: 22px;
    background-color: rgb(240, 242, 245);
    position: relative;

    .github-corner {
      position: absolute;
      top: 0px;
      border: 0;
      right: 0;
    }

    .chart-wrapper {
      background: #fff;
      padding: 16px 16px 0;
      margin-bottom: 32px;
    }
  }

@media (max-width:1024px) {
  .chart-wrapper {
    padding: 8px;
  }
}
</style>
