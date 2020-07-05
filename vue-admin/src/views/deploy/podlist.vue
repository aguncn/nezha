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

    <el-table v-loading="listLoading" :data="list" border fit highlight-current-row style="width: 100%">
      <el-table-column align="center" label="ID" min-width="5%">
        <template slot-scope="scope">
          <span>{{ scope.row.ID }}</span>
        </template>
      </el-table-column>
      <el-table-column align="center" label="Pod名称" min-width="20%">
        <template slot-scope="scope">
          <span>{{ scope.row.podname }}</span>
        </template>
      </el-table-column>
      <el-table-column align="center" label="namespace" min-width="15%">
        <template slot-scope="scope">
          <span>{{ scope.row.namespace }}</span>
        </template>
      </el-table-column>
      <el-table-column align="center" label="container" min-width="15%">
        <template slot-scope="scope">
          <span>{{ scope.row.container }}</span>
        </template>
      </el-table-column>
      <el-table-column align="center" label="环境" min-width="10%">
        <template slot-scope="scope">
          <span>{{ scope.row.Environment.name }}</span>
        </template>
      </el-table-column>
      <el-table-column align="center" label="K8s" min-width="10%">
        <template slot-scope="scope">
          <span>{{ scope.row.K8s.name }}</span>
        </template>
      </el-table-column>
      <el-table-column prop="date" label="创建日期" sortable align="center" min-width="10%">
        <template slot-scope="scope">
          <span style="margin-left: 10px">{{ scope.row.CreatedAt | parseTime('{m}-{d} {h}:{i}') }}</span>
        </template>
      </el-table-column>
      <el-table-column align="center" label="终端" min-width="10%">
        <template slot-scope="scope">
          <a @click="wtlink(scope.row.K8s.terminal, scope.row.namespace, scope.row.podname, scope.row.container)">
            <el-button type="info" :size="size" icon="el-icon-caret-right">
              terminal
            </el-button>
          </a>
        </template>
      </el-table-column>
    </el-table>

    <pagination v-show="total>0" :total="total" :page.sync="listQuery.page" :limit.sync="listQuery.limit" @pagination="getList" />
  </div>

</template>

<script>
import { fetchList } from '@/api/pod'
import { parseTime } from '@/utils'
import Pagination from '@/components/Pagination'
import NzButton from '@/components/NzButton'

export default {
  name: 'PodList',
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
        appID: this.id,
        page: 1,
        limit: 10,
        application_id: this.id,
        name: undefined
      }
    }
  },
  created() {
    this.getList()
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
    wtlink(url, namespace, podname, container) {
      // 插入数据库，pod里加了-1,-2,-3这种区分，一个pod只要不超过10个container，下面的pod截取就是有效的。
      podname = podname.substring(0, podname.length - 2)
      var wturl = url + '/?namespace=' + namespace + '&pod=' + podname + '&container=' + container
      console.log(wturl)
      window.open(wturl, '_blank')
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
    background-color: #000000;
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
