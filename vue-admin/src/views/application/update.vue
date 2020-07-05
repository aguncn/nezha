<template>
  <application-form :id=parseInt(this.$route.params.id) :listProject=listProject :readonly="false" :opType="'update'" />
</template>

<script>
import ApplicationForm from './components/form'
import { fetchList as fetchProjectList } from '@/api/project'

export default {
  name: 'EditApplication',
  components: { ApplicationForm },
  data() {
    return {
      listProject: null,
      listProjectQuery: {
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
      fetchProjectList(this.listProjectQuery).then(response => {
        this.listProject = response.data.list
      })
      this.listLoading = false
    }
  }
}
</script>
