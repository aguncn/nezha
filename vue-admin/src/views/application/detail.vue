<template>
  <application-form  :id=parseInt(this.$route.params.id) :listProject=listProject :readonly="true" :opType="'detail'" />
</template>

<script>
import { fetchList as fetchProjectList } from '@/api/project'
import ApplicationForm from './components/form'

export default {
  name: 'DetailApplication',
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
