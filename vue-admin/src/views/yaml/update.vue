<template>
  <yaml-form  :id=parseInt(this.$route.params.id) :listK8s=listK8s :listApplication=listApplication :readonly="false" :opType="'update'" />
</template>

<script>
import YamlForm from './components/form'
import { fetchList as fetchK8sList } from '@/api/k8s'
import { fetchList as fetchApplicationList } from '@/api/application'

export default {
  name: 'EditYaml',
  components: { YamlForm },
  data() {
    return {
      listK8s: null,
      listK8sQuery: {
        page: 1,
        limit: 100,
        name: undefined
      },
      listApplication: null,
      listApplicationQuery: {
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
      fetchK8sList(this.listK8sQuery).then(response => {
        this.listK8s = response.data.list
      })
      fetchApplicationList(this.listApplicationQuery).then(response => {
        this.listApplication = response.data.list
      })
      this.listLoading = false
    }
  }
}
</script>
