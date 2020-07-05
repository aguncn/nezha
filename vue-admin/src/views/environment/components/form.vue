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
      <el-form-item label="描述" prop="description">
        <el-input v-model="dataForm.description" :readonly="readonly" auto-complete="off"></el-input>
      </el-form-item>
    </el-form>
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
import { fetchEnvironment, createEnvironment, updateEnvironment } from '@/api/environment'

	export default {
		name: 'EnvironmentForm',
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
      }
		},
		data() {
			return {
        dataForm: {
          ID: 0,
          name: '环境名称',
          description: '描述'
        },
        dataFormRules: {
          name: [
            {
              required: true,
              message: '请输入环境名称',
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
      if (this.opType === "update") {
        this.getDetail(this.id)
      }
    },
		methods: {
      getDetail(params) {
        this.listLoading = true
        fetchEnvironment(params).then(response => {
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
            createEnvironment(this.dataForm).then(response => {
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
              this.$router.push({ path: this.redirect || '/environment/list' })
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
            updateEnvironment(this.dataForm).then(response => {
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
              this.$router.push({ path: this.redirect || '/environment/list' })
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
    position: relative;

    .createPost-main-container {
      padding: 40px 45px 20px 50px;

      .postInfo-container {
        position: relative;
        @include clearfix;
        margin-bottom: 10px;

        .postInfo-container-item {
          float: left;
        }
      }
    }

    .word-counter {
      width: 40px;
      position: absolute;
      right: 10px;
      top: 0px;
    }
  }

  .article-textarea /deep/ {
    textarea {
      padding-right: 40px;
      resize: none;
      border: none;
      border-radius: 0px;
      border-bottom: 1px solid #bfcbd9;
    }
  }
</style>
