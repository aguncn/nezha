<template>
  <div class="page-container">
    <el-form
      ref="dataForm"
      :model="dataForm"
      label-width="100px"
      :size="size"
      label-position="right"
    >
      <el-row>
        <el-col :span="24">
          <el-form-item label="应用" prop="application">
            <el-input :value="application" :readonly="true" />
          </el-form-item>
        </el-col>
      </el-row>
      <el-row>
        <el-col :span="6">
          <el-form-item label="版本" prop="version">
            <el-input v-model="dataForm.version" auto-complete="off" />
          </el-form-item>
        </el-col>
        <el-col :span="9">
          <el-form-item label="镜像" prop="harbor">
            <el-select v-model="dataForm.harbor" placeholder="请选择" style="width:100%" @change="selectHarborTrigger(dataForm.harbor)">
              <el-option
                v-for="(item, index) in listHarbor"
                :key="index"
                :label="item"
                :value="item"
              />
            </el-select>
          </el-form-item>
        </el-col>
        <el-col :span="9">
          <el-form-item label="上线描述" prop="reason">
            <el-input v-model="dataForm.reason" auto-complete="off" />
          </el-form-item>
        </el-col>
      </el-row>
      <el-row>
        <el-col :span="6">
          <el-form-item label="环境" prop="environment_id">
            <el-select v-model="dataForm.environment_id" style="width:100%" @change="selectEnvTrigger(dataForm.environment_id)">
              <el-option
                v-for="(item,index) in listEnvironment"
                :key="index"
                :label="item.name"
                :value="item.ID"
              />
            </el-select>
          </el-form-item>
        </el-col>
        <el-col :span="9">
          <el-form-item v-show="isEnv" label="K8s集群" prop="k8s">
            <el-select v-model="dataForm.k8s_id" style="width:100%" @change="selectK8sTrigger(dataForm.k8s_id)">
              <el-option
                v-for="(item,index) in listK8s"
                :key="index"
                :label="item.name"
                :value="item.ID"
              />
            </el-select>
          </el-form-item>
        </el-col>
        <el-col :span="9">
          <el-form-item v-show="isK8s" label="Yaml" prop="yaml">
            <el-select v-model="dataForm.yaml_id" style="width:100%" @change="selectYamlTrigger(dataForm.yaml_id)">
              <el-option
                v-for="(item,index) in listYaml"
                :key="index"
                :label="item.name"
                :value="item.ID"
              />
            </el-select>
          </el-form-item>
        </el-col>
      </el-row>
    </el-form>
    <div style="text-align: center">
      <template>
        <el-popconfirm title="确定执行？" @onConfirm="onConfirm">
          <el-button slot="reference" :disabled="deployIsDisabled" type="danger">发布上线</el-button>
        </el-popconfirm>
      </template>
    </div>
    <el-divider />
    <el-row>
      <el-col :span="20" :offset="2">
        <el-steps style="text-align: left;" :active="active" finish-status="success">
          <el-step v-for="(item,index) in stepItems" :key="index" :title="item.title" :description="item.description" />
        </el-steps>
        <p />
        <el-input
          v-model="log"
          type="textarea"
          :rows="5"
          placeholder="发布日志输出"
          :disabled="true"
        />
      </el-col>
    </el-row>
  </div>
</template>

<script>
import { mapGetters } from 'vuex'
import { fetchApplication } from '@/api/application'
import { fetchList as fetchEnvironmentList } from '@/api/environment'
import { fetchList as fetchK8sList } from '@/api/k8s'
import { fetchList as fetchYamlList } from '@/api/yaml'
import { createDeploy } from '@/api/deploy'
import { submitDeploy } from '@/api/deploy'
import { statusDeploy } from '@/api/deploy'
import { getDockerTag } from '@/api/deploy'
import { strTime } from '@/utils/datetime'

export default {
  name: 'DeployList',
  props: {
    id: {
      type: Number,
      default: 0
    }
  },
  data() {
    return {
      size: 'small',
      active: 0,
      deployIsDisabled: true,
      dataForm: {
        name: '',
        reason: '',
        version: 'v0.1',
        harbor: '',
        description: '描述',
        application_id: '',
        k8s_id: '',
        environment_id: '',
        yaml_id: '',
        yaml: ''
      },
      stepItems: [
        { title: '合成yaml文件', description: '' },
        { title: '生成发布批次', description: '' },
        { title: '应用到k8s集群', description: '' },
        { title: '新版本上线', description: '' }
      ],
      isEnv: false,
      isK8s: false,
      application: '',
      listHarbor: [],
      listEnvironment: null,
      listYaml: null,
      listK8s: null,
      listQuery: {
        page: 1,
        limit: 100,
        name: undefined
      },
      listK8sQuery: {
        page: 1,
        limit: 100,
        name: undefined,
        environment_id: undefined
      },
      listYamlQuery: {
        page: 1,
        limit: 100,
        name: undefined,
        k8s_id: undefined,
        application_id: undefined
      },
      listDockerTagQuery: {
        harbor: undefined
      },
      timer: null,
      log: ''
    }
  },
  computed: {
    ...mapGetters([
      'userId',
      'userName'
    ])
  },
  created() {
    this.getEnvironmentList()
    this.getDetail(this.id)
  },
  methods: {
    next() {
      if (this.active++ > 3) this.active = 0
    },
    getDetail(params) {
      this.listLoading = true
      fetchApplication(params).then(response => {
        this.dataForm.application_id = response.data.ID
        this.application = response.data.name
        // 只传递真正的项目应用到后端，因为后端是API，所以后端组合成真正的API地址。
        var harborAddr = response.data.harbor
        var tmpArr = harborAddr.split('/')
        this.listDockerTagQuery.harbor = tmpArr.slice(1, tmpArr.length).join('/')
        getDockerTag(this.listDockerTagQuery).then(response => {
          for (var item of response.data) {
            this.listHarbor.unshift(harborAddr + ':' + item.name)
          }
        })
      })
      this.listLoading = false
    },
    getEnvironmentList() {
      this.listLoading = true
      fetchEnvironmentList(this.listQuery).then(response => {
        this.listEnvironment = response.data.list
      })
      this.listLoading = false
    },
    selectHarborTrigger(value) {
      console.log(value)
      if (this.dataForm.yaml_id !== '') {
        this.deployIsDisabled = false
      }
    },
    selectEnvTrigger(value) {
      this.$set(this.dataForm, 'k8s_id', '')
      this.isEnv = true
      this.listK8sQuery.environment_id = value
      this.listLoading = true
      fetchK8sList(this.listK8sQuery).then(response => {
        this.listK8s = response.data.list
      })
      this.listLoading = false
    },
    selectK8sTrigger(value) {
      this.$set(this.dataForm, 'yaml_id', '')
      this.isK8s = true
      this.listYamlQuery.k8s_id = value
      this.listYamlQuery.application_id = this.id
      this.listLoading = true
      fetchYamlList(this.listYamlQuery).then(response => {
        this.listYaml = response.data.list
      })
      this.listLoading = false
    },
    selectYamlTrigger(value) {
      if (this.dataForm.harbor !== '') {
        this.deployIsDisabled = false
      }
    },
    onConfirm() {
      this.deployIsDisabled = true
      // 先更新发布批次及yaml内容
      this.dataForm.name = strTime()
      for (const i of this.listYaml) {
        if (i.ID === this.dataForm.yaml_id) {
          this.dataForm.yaml = i.yaml
          break
        }
      }
      let noEmpty = true
      for (const key in this.dataForm) {
        if (this.dataForm[key].length === 0) {
          noEmpty = false
        }
      }
      if (!noEmpty) {
        this.$message({ type: 'warning', message: '所有可选项不能为空' })
        return
      }
      this.dataForm.yaml = this.dataForm.yaml.replace('demo-docker-image', this.dataForm.harbor)
      this.active = 1
      this.stepItems[0].description = 'yaml文件替换成功！'

      const self = this
      self.dataForm.state = 1
      self.dataForm.user_id = self.userId
      self.dataForm.created_by = self.userName
      createDeploy(self.dataForm).then(res => {
        this.log = '生成发布批次\n'
        this.log += res['msg']
        this.log += '\n'
        if (res['code'] === 200) {
          self.active = 2
          self.stepItems[1].description = self.dataForm.name
        } else {
          console.log('生成发布批次出错，退出')
          return
        }
        submitDeploy(self.dataForm).then(res => {
          this.log += '应用到K8s集群\n'
          this.log += res['msg']
          this.log += '\n'
          this.log += res['data']
          this.log += '\n'
          if (res['code'] === 200) {
            self.active = 3
          } else {
            console.log('应用到K8s集群出错，退出')
            return
          }

          var timer = setInterval(() => {
            statusDeploy(self.dataForm).then(res => {
              if (res['code'] === 200 && res['msg'] === 'ok') {
                self.active = 4
                this.log += '新版本上线\n'
                this.log += res['msg']
                this.log += '\n'
                this.log += res['data']
                this.log += '\n'
                clearInterval(timer)
                return
              } else if (res['data'].indexOf('ImagePullBackOff') !== -1 ||
                res['data'].indexOf('ErrImagePull') !== -1) {
                this.log += 'error\n'
                this.log += res['data']
                this.log += '\n'
                clearInterval(timer)
                return
              } else {
                this.log += res['data']
                this.log += '\n'
              }
            }).catch((err) => {
              console.log(err)
              return
            })
          }, 5000)
        }).catch((err) => {
          console.log(err)
          return
        })
      }).catch((err) => {
        console.log(err)
        return
      })
    }
  }
}
</script>
<style>
  .el-row {
    margin-bottom: 20px;
    &:last-child {
      margin-bottom: 0;
    }
  }
  .el-col {
    border-radius: 4px;
  }
  .bg-purple-dark {
    background: #99a9bf;
  }
  .bg-purple {
    background: #d3dce6;
  }
  .bg-purple-light {
    background: #e5e9f2;
  }
  .grid-content {
    border-radius: 4px;
    min-height: 36px;
  }
  .row-bg {
    padding: 10px 0;
    background-color: #f9fafc;
  }
</style>
<style scoped>

</style>
