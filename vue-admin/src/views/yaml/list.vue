<template>
  <div class="app-container">
    <div class="menu-header">
      <el-form :inline="true" :size="size">
        <el-form-item>
          <el-input
            v-model="listQuery.name"
            style="width: 300px;"
            placeholder="关键字"
            @keyup.enter.native="handleFilter"
          />
        </el-form-item>
        <el-form-item>
          <nz-button icon="el-icon-search" :label="'搜索'" type="primary" @click="handleFilter" />
        </el-form-item>
        <el-form-item>
          <router-link :to="'/yaml/create'">
            <nz-button icon="fa fa-plus" :label="'新增'" type="primary" />
          </router-link>
        </el-form-item>
      </el-form>
    </div>

    <el-table v-loading="listLoading" :data="list" border fit highlight-current-row style="width: 100%">
      <el-table-column align="center" label="ID" min-width="5%">
        <template slot-scope="scope">
          <span>{{ scope.row.ID }}</span>
        </template>
      </el-table-column>

      <el-table-column align="center" label="Yaml配置名称" min-width="10%">
        <template slot-scope="scope">
          <span>{{ scope.row.name }}</span>
        </template>
      </el-table-column>
      <el-table-column align="center" label="描述" min-width="20%">
        <template slot-scope="scope">
          <span>{{ scope.row.description }}</span>
        </template>
      </el-table-column>
      <el-table-column align="center" label="环境" min-width="10%">
        <template slot-scope="scope">
          <span>{{ scope.row.K8s.Environment.name }}</span>
        </template>
      </el-table-column>
      <el-table-column align="center" label="Application" min-width="15%">
        <template slot-scope="scope">
          <span>{{ scope.row.Application.name }}</span>
        </template>
      </el-table-column>
      <el-table-column align="center" label="K8s集群" min-width="15%">
        <template slot-scope="scope">
          <span>{{ scope.row.K8s.name }}</span>
        </template>
      </el-table-column>
      <el-table-column prop="date" label="创建日期" sortable align="center" min-width="15%">
        <template slot-scope="scope">
          <i class="el-icon-time" />
          <span style="margin-left: 10px">{{ scope.row.CreatedAt | parseTime('{y}-{m}-{d}') }}</span>
        </template>
      </el-table-column>
      <el-table-column align="center" min-width="10%">
        <template slot-scope="scope">
          <router-link :to="'/yaml/detail/'+scope.row.ID">
            <nz-button
              icon="el-icon-info"
              :size="size"
              :plain="true"
              :circle="true"
              :is-label="false"
              type="info"
            />
          </router-link>
          <router-link :to="'/yaml/update/'+scope.row.ID">
            <nz-button
              icon="el-icon-edit"
              :plain="true"
              :circle="true"
              :is-label="false"
              :size="size"
              type="info"
            />
          </router-link>
          <nz-button
            icon="el-icon-delete"
            :plain="true"
            :circle="true"
            :is-label="false"
            :size="size"
            type="danger"
            @click="handleDelete(scope.row)"
          />
        </template>
      </el-table-column>
    </el-table>
    <pagination v-show="total>0" :total="total" :page.sync="listQuery.page" :limit.sync="listQuery.limit" @pagination="getList" />
  </div>
</template>

<script>
import { fetchList, deleteYaml } from '@/api/yaml'
import { parseTime } from '@/utils'
import Pagination from '@/components/Pagination'
import NzButton from '@/components/NzButton'

export default {
  name: 'YamlList',
  components: {
    Pagination,
    NzButton
  },
  filters: {
      parseTime: parseTime
  },
  data() {
    return {
      size: 'mini',
      Filters: {
        label: ''
      },
      title: '',
      dialogVisible: false,
      list: null,
      total: 0,
      listLoading: true,
      readonly: false,
      listQuery: {
        page: 1,
        limit: 10,
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
    handleDelete: function(row) {
      this.$confirm('确定删除' + row.name + '吗？', '提示', {
        type: 'warning'
      }).then(() => {
        deleteYaml(row.ID).then(response => {
          this.getList()
          if (response.msg === 'fail') {
            this.$notify({
              title: 'Fail',
              message: response.detail,
              type: 'error',
              duration: 2000
            })
          } else {
            this.$notify({
              title: 'Success',
              message: '删除Yaml成功!',
              type: 'success',
              duration: 2000
            })
          }
        })
      }).catch(() => {})
    }
  }
}
</script>

<style scoped>
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
