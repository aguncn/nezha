<template>
  <div class="app-container">
    <div class="menu-header">
      <el-form :inline="true" :size="size">
        <el-form-item>
          <el-input
            v-model="listQuery.name"
            style="width: 300px;"
            :placeholder="application"
            @keyup.enter.native="handleFilter"
          />
        </el-form-item>
        <el-form-item>
          <nz-button icon="el-icon-search" :label="'搜索'" type="primary" @click="handleFilter" />
        </el-form-item>
      </el-form>
    </div>

    <el-drawer
      :title="title"
      size="50%"
      :visible.sync="drawer"
      :before-close="handleDrawerClose"
    >
      <table>
        <tr>
          <td>
            <!--这种取值，只取到第一行的值，而不会每行都更新-->
            <!--{{ scope.row.name }}-->
            <pre>
              <code>
{{ content }}
              </code>

            </pre>
          </td>
        </tr>
      </table>
    </el-drawer>

    <el-table v-loading="listLoading" :data="list" border fit highlight-current-row style="width: 100%">
      <el-table-column align="center" label="ID" min-width="5%">
        <template slot-scope="scope">
          <span>{{ scope.row.ID }}</span>
        </template>
      </el-table-column>
      <el-table-column align="center" label="发布批次" min-width="20%">
        <template slot-scope="scope">
          <el-tooltip class="item" effect="dark" :content="scope.row.reason" placement="top-start">
            <span>{{ scope.row.name }}</span>
          </el-tooltip>
        </template>
      </el-table-column>
      <el-table-column align="center" label="版本" min-width="10%">
        <template slot-scope="scope">
          <span>{{ scope.row.version }}</span>
        </template>
      </el-table-column>
      <el-table-column align="center" label="环境" min-width="10%">
        <template slot-scope="scope">
          <span>{{ scope.row.Environment.name }}</span>
        </template>
      </el-table-column>
      <el-table-column align="center" label="K8s群集" min-width="10%">
        <template slot-scope="scope">
          <span>{{ scope.row.K8s.name }}</span>
        </template>
      </el-table-column>
      <el-table-column align="center" label="Yaml" min-width="10%">
        <template slot-scope="scope">
          <el-button plain @click="handleDrawerOpen('Yaml内容', scope.row.yaml)">
            查看
          </el-button>
        </template>
      </el-table-column>
      <el-table-column prop="date" label="创建日期" sortable align="center" min-width="10%">
        <template slot-scope="scope">
          <span style="margin-left: 10px">{{ scope.row.CreatedAt | parseTime('{m}-{d} {h}:{i}') }}</span>
        </template>
      </el-table-column>
      <el-table-column align="center" label="状态" min-width="15%">
        <template slot-scope="scope">
          <span>{{ scope.row.result }}</span>
        </template>
      </el-table-column>
      <el-table-column align="center" label="日志" min-width="10%">
        <template slot-scope="scope">
          <el-button plain type="primary" @click="handleDrawerOpen('Log输出', scope.row.log)">
            查看
          </el-button>
        </template>
      </el-table-column>
    </el-table>

    <pagination v-show="total>0" :total="total" :page.sync="listQuery.page" :limit.sync="listQuery.limit" @pagination="getList" />
  </div>

</template>

<script>
import { fetchList } from '@/api/deploy'
import { parseTime } from '@/utils'
import Pagination from '@/components/Pagination'
import NzButton from '@/components/NzButton'
import { fetchApplication } from '@/api/application'

export default {
  name: 'HistoryList',
  components: {
    Pagination,
    NzButton
  },
  filters: {
    parseTime: parseTime
  },
  props: {
    id: {
      type: Number,
      default: 0
    }
  },
  data() {
    return {
      size: 'mini',
      Filters: {
        label: ''
      },
      application: '',
      drawer: false,
      title: '',
      content: '',
      list: null,
      total: 0,
      listLoading: true,
      readonly: false,
      listQuery: {
        page: 1,
        limit: 10,
        application_id: this.id,
        name: undefined
      }
    }
  },
  created() {
    this.getList()
    this.getDetail(this.id)
  },
  methods: {
    getList() {
      this.listLoading = true
      fetchList(this.listQuery).then(response => {
        this.list = response.data.list
        this.total = response.data.total
        this.listLoading = false
      })
    },
    handleFilter() {
      this.listQuery.page = 1
      this.getList()
    },
    getDetail(params) {
      this.listLoading = true
      fetchApplication(params).then(response => {
        this.application = response.data.name
        this.listLoading = false
      })
    },
    handleDrawerOpen(title, param) {
      this.title = title
      this.drawer = true
      this.content = param
    },
    handleDrawerClose(done) {
      this.content = ''
      done()
    }
  }
}
</script>
<style>
  .el-drawer.rtl{
    overflow: scroll;
}
</style>
<style scoped>
  pre {
    width: auto;
    background-color: #eeeeee;
    overflow-x: auto;
  }
  .menu-header {
    text-align: left;
    font-size: 16px;
    color: rgb(20, 89, 121);
  }

  .edit-input {
    padding-right: 100px;
  }
  .cancel-btn {
    position: absolute;
    right: 15px;
    top: 10px;
  }
</style>
