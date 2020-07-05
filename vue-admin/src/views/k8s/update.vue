<template>
  <k8s-form  :id=parseInt(this.$route.params.id) :listEnvironment=listEnvironment :readonly="false" :opType="'update'" />
</template>

<script>
import K8sForm from './components/form'
import { fetchList as fetchEnvironmentList } from '@/api/environment'

export default {
  name: 'UpdateK8s',
  components: { K8sForm },
  data() {
    return {
      listEnvironment: null,
      listEnvironmentQuery: {
        page: 1,
        limit: 100,
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
      fetchEnvironmentList(this.listEnvironmentQuery).then(response => {
        this.listEnvironment = response.data.list
      })
      this.listLoading = false
    }
  }
}
</script>
