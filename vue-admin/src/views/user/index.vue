<template>
  <div class="app-container">
    <div class="menu-header">
      <el-form :inline="true" :size="size">
        <el-form-item>
      <el-input
        v-model="listQuery.name"
        placeholder="请输入用户名称"
        style="width: 200px;"
        class="filter-item"
        @keyup.enter.native="handleFilter"
      />
      </el-form-item>
      <el-form-item>
      <el-button class="filter-item" type="primary" icon="el-icon-search" @click="handleFilter">搜索</el-button>
      </el-form-item>
      <el-form-item>
      <el-button
        class="filter-item"
        style="margin-left: 10px;"
        type="primary"
        icon="el-icon-edit"
        @click="handleCreate"
      >新建</el-button>
      </el-form-item>
      </el-form>
    </div>
    <el-table
      v-loading="listLoading"
      :data="list"
      element-loading-text="Loading"
      border
      style="width: 100%"
    >
      <el-table-column label="id" prop="id" sortable align="center" min-width="5%">
        <template slot-scope="scope">
          <span>{{ scope.row.id }}</span>
        </template>
      </el-table-column>
      <el-table-column label="名称" align="center" min-width="10%">
        <template slot-scope="scope">
          <span>{{ scope.row.username }}</span>
        </template>
      </el-table-column>
      <el-table-column label="密码" align="center" min-width="10%">
        <span>*******************</span>
      </el-table-column>
      <el-table-column
        prop="state"
        label="状态"
        align="center"
        :filters="[{ text: '正常', value: '正常' }, { text: '禁用', value: '禁用' }]"
        :filter-method="filterTag"
        min-width="10%"
      >
        <template slot-scope="scope">
          <el-tag :type="scope.row.state | statusFilter">{{ scope.row.state }}</el-tag>
        </template>
      </el-table-column>
      <el-table-column
        prop="user_type"
        label="用户类型"
        align="center"
        :filters="[{ text: '管理员', value: '管理员' }, { text: '研发', value: '研发' }]"
        :filter-method="filterType"
        min-width="10%"
      >
        <template slot-scope="scope">
          <el-tag :type="scope.row.user_type | userTypeFilter">{{ scope.row.user_type }}</el-tag>
        </template>
      </el-table-column>
      <el-table-column prop="date" label="创建日期" sortable align="center" min-width="15%">
        <template slot-scope="scope">
          <i class="el-icon-time" />
          <span style="margin-left: 10px">{{ scope.row.created_at }}</span>
        </template>
      </el-table-column>
      <el-table-column label="操作" align="center" min-width="25%">
        <template slot-scope="scope">
          <el-button size="mini" type="warning" @click="handleResetPwd(scope.$index, scope.row)">重置密码</el-button>
          <el-button size="mini" type="warning" @click="handlePermission(scope.$index, scope.row)">编辑权限</el-button>
          <el-button size="mini" type="danger" @click="handleDelete(scope.$index, scope.row)">删除</el-button>
        </template>
      </el-table-column>
    </el-table>

    <pagination
      v-show="total>0"
      :total="total"
      :page.sync="listQuery.page"
      :limit.sync="listQuery.limit"
      @pagination="fetchData"
    />

    <el-dialog :title="textMap[dialogStatus]" :visible.sync="dialogFormVisible">
      <el-form
        ref="dataForm"
        :rules="rules"
        :model="temp"
        label-position="left"
        label-width="70px"
        style="width: 400px; margin-left:50px;"
      >
        <el-form-item label="用户名" prop="username">
          <el-input v-model="temp.username" />
        </el-form-item>
        <el-form-item v-if="pwdShow" label="密码" prop="password">
          <el-input v-model="temp.password" />
        </el-form-item>
        <el-form-item v-if="permissionShow" label="权限" prop="user_type">
          <el-select v-model="temp.user_type" placeholder="请选择权限">
            <el-option label="管理员" :value="1" />
            <el-option label="研发" :value="2" />
          </el-select>
        </el-form-item>
      </el-form>
      <div slot="footer" class="dialog-footer">
        <el-button @click="dialogFormVisible = false">取消</el-button>
        <el-button type="primary" @click="dialogStatus==='create'?createData():updateData()">确定</el-button>
      </div>
    </el-dialog>
  </div>
</template>

<script>
import { getList, createUser, updateUser, deleteUser } from '@/api/user'
import Pagination from '@/components/Pagination' // secondary package based on el-pagination
export default {
  name: 'ComplexTable',
  components: { Pagination },
  filters: {
    statusFilter(Status) {
      const statusMap = {
        正常: 'success',
        禁用: 'danger'
      }
      return statusMap[Status]
    },
    userTypeFilter(userType) {
      const statusMap = {
        管理员: 'success',
        研发: 'primary'
      }
      return statusMap[userType]
    }
  },
  data() {
    return {
      size: 'mini',
      list: null,
      total: 0,
      listLoading: true,
      listQuery: {
        page: 1,
        limit: 10,
        name: undefined
      },
      pwdShow: true,
      permissionShow: true,
      dialogFormVisible: false,
      dialogStatus: '',
      textMap: {
        permission: '编辑权限',
        resetPwd: '重置密码',
        create: '新建用户'
      },
      rules: {
        username: [
          { required: true, message: '请输入用户名', trigger: 'blur' }
        ],
        password: [{ required: true, message: '请输入密码', trigger: 'blur' }],
        user_type: [{ required: true, message: '请选择权限', trigger: 'change' }]
      },
      temp: {
        id: undefined,
        username: '',
        password: '',
        user_type: undefined,
        op_type: ''
      }
    }
  },
  created() {
    this.fetchData()
  },
  methods: {
    handleResetPwd(index, row) {
      this.temp.id = row.id
      this.temp.username = row.username
      this.temp.password = ''
      this.temp.user_type = row.user_type
      this.temp.op_type = 'adminResetPwd'
      this.dialogStatus = 'resetPwd'
      this.permissionShow = false
      this.pwdShow = true
      this.dialogFormVisible = true
      this.$nextTick(() => {
        this.$refs['dataForm'].clearValidate()
      })
    },
    handlePermission(index, row) {
      this.temp.id = row.id
      this.temp.username = row.username
      this.temp.password = row.password
      this.temp.user_type = row.user_type
      this.temp.op_type = 'adminChangePermission'
      this.dialogStatus = 'permission'
      this.pwdShow = false
      this.permissionShow = true
      this.dialogFormVisible = true
      this.$nextTick(() => {
        this.$refs['dataForm'].clearValidate()
      })
    },
    handleDelete(index, row) {
      deleteUser(row.id).then(response => {
        this.fetchData()
        this.dialogFormVisible = false
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
            message: '删除用户成功!',
            type: 'success',
            duration: 2000
          })
        }
      })
    },
    fetchData() {
      this.listLoading = true
      getList(this.listQuery).then(response => {
        this.list = response.data.list
        this.total = response.data.total
        this.listLoading = false
      })
    },
    handleFilter() {
      this.listQuery.page = 1
      this.fetchData()
    },
    resetTemp() {
      this.temp = {
        id: undefined,
        username: '',
        password: ''
      }
    },
    handleCreate() {
      this.resetTemp()
      this.dialogStatus = 'create'
      this.pwdShow = true
      this.permissionShow = true
      this.dialogFormVisible = true
      this.temp.op_type = 'adminAddUser'
      this.$nextTick(() => {
        this.$refs['dataForm'].clearValidate()
      })
    },
    filterTag(value, row) {
      return row.state === value
    },
    filterType(value, row) {
      return row.user_type === value
    },
    createData() {
      this.$refs['dataForm'].validate(valid => {
        if (valid) {
          createUser(this.temp).then(() => {
            this.fetchData()
            this.dialogFormVisible = false
            this.$notify({
              title: 'Success',
              message: '新建用户成功!',
              type: 'success',
              duration: 2000
            })
          })
        }
      })
    },
    updateData() {
      this.$refs['dataForm'].validate(valid => {
        if (valid) {
          const tempData = Object.assign({}, this.temp)
          if (tempData.user_type === '管理员') {
            tempData.user_type = 1
          } else if (tempData.user_type === '研发') {
            tempData.user_type = 2
          }

          updateUser(tempData).then(() => {
            this.fetchData()
            this.dialogFormVisible = false
            this.$notify({
              title: 'Success',
              message: '更新数据成功',
              type: 'success',
              duration: 2000
            })
          })
        }
      })
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
