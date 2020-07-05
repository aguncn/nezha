<template>
  <div class="application-container">
    <el-form
      :model="dataForm"
      label-width="100px"
      :rules="dataFormRules"
      ref="dataForm"
      :size="size"
      label-position="right">
      <el-form-item label="ID" prop="ID" v-if="false">
        <el-input v-model="dataForm.ID" :disabled="true" auto-complete="off"></el-input>
      </el-form-item>
      <el-form-item label="名称" prop="name">
        <el-input v-model="dataForm.name" :readonly="readonly" auto-complete="off"></el-input>
      </el-form-item>
      <el-form-item label="中文名称" prop="cn_name">
        <el-input v-model="dataForm.cn_name" :readonly="readonly" auto-complete="off"></el-input>
      </el-form-item>
      <el-form-item label="描述" prop="description">
        <el-input v-model="dataForm.description" :readonly="readonly" auto-complete="off"></el-input>
      </el-form-item>
      <el-form-item label="所属项目" prop="project_id">
        <el-select v-model="dataForm.project_id" :disabled="readonly" style="width:100%">
            <el-option
                v-for="(item,index) in listProject"
                :key="index"
                :label="item.name"
                :value="item.ID" />
        </el-select>
      </el-form-item>

      <el-form-item label="Harbor地址" prop="harbor">
        <el-input v-model="dataForm.harbor" :readonly="readonly" auto-complete="off"></el-input>
      </el-form-item>
    </el-form>
    <div v-if="this.opType==='detail'">
      <router-link :to="'/application/list'">
      <el-button :size="size">返回</el-button>
      </router-link>
    </div>
    <div v-if="this.opType==='update'">
      <el-button :size="size" type="primary" @click.native="submitUpdateForm" >保存</el-button>
    </div>
    <div v-if="this.opType==='create'">
      <el-button :size="size" @click.native="resetForm">Cancel</el-button>
      <el-button :size="size" type="primary" @click.native="submitCreateForm" >保存</el-button>
    </div>
  </div>
</template>

<script>
import { mapGetters } from 'vuex'
import { fetchApplication, createApplication, updateApplication } from '@/api/application'

	export default {
		name: 'ApplicationForm',
		props: {
      id: {
      	type: Number,
      	default: 0
      },
      opType: {
      	type: String,
      	default: ''
      },
			size: {
				type: String,
				default: 'mini'
			},
      readonly: {
      	type: Boolean,
      	default: false
      },
      listProject: {
        type: Array,
        default: null
      }
		},
		data() {
			return {
        dataForm: {
          ID: 0,
          name: '应用名称',
          description: '描述',
          project_id: '',
          git: 'git',
          jenkins: 'jenkins'
        },
        dataFormRules: {
          name: [
            {
              required: true,
              message: '请输入应用名称',
              trigger: 'blur',
            }
          ]
        }

			}
		},
    computed: {
      ...mapGetters([
        'userId',
        'userName'
      ])
    },
    created() {
      if (this.opType === "detail" || this.opType === "update") {
        this.getDetail(this.id)
      }
    },
    methods: {
      getDetail(params) {
        this.listLoading = true
        fetchApplication(params).then(response => {
          this.dataForm = response.data
          this.listLoading = false
        })
      },
      submitCreateForm() {
        this.$refs.dataForm.validate(valid => {
          if (valid) {
            this.dataForm.state = 1
            this.dataForm.user_id = this.userId
            this.dataForm.created_by = this.userName
            console.log(this.dataForm)
            createApplication(this.dataForm).then(response => {
              this.loading = true
              if (response.code !== 200) {
                this.$notify({
                  title: '错误',
                  message: '可能存在相同应用，请仔细核查！',
                  type: 'error',
                  duration: 2000
                })
                this.loading = false
                return
              }
              this.$notify({
                title: '成功',
                message: '新建应用成功',
                type: 'success',
                duration: 2000
              })
              this.$router.push({ path: this.redirect || '/application/list' })
              this.loading = false
            })
          } else {
            console.log('error submit!!')
            return false
          }
        })
      },
      submitUpdateForm() {
        this.$refs.dataForm.validate(valid => {
          if (valid) {
            console.log(this.dataForm, "####555555###")
            updateApplication(this.dataForm).then(response => {
              this.loading = true
              if (response.code !== 200) {
                this.$notify({
                  title: '错误',
                  message: '更新失败，可能存在相同应用,请仔细核查！',
                  type: 'error',
                  duration: 2000
                })
                this.loading = false
                return
              }
              this.$notify({
                title: '成功',
                message: '更新应用成功',
                type: 'success',
                duration: 2000
              })
              this.$router.push({ path: this.redirect || '/application/list' })
              this.loading = false
            })
          } else {
            console.log('error submit!!')
            return false
          }
        })
      },
      resetForm: function() {
        this.dialogVisible = false
        this.readonly = false
        this.$refs['dataForm'].resetFields()
      }
		}
	}
</script>

<style>
  .application-container {
    padding-top: 20px;
    padding-bottom: 20px;
    width: 80%;
    margin:0 auto;
    position: relative;
    border:1px solid #ccc
  }
</style>
