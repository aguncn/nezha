<template>
  <project-form :id=parseInt(this.$route.params.id) :readonly="false" :opType="'create'" />
</template>

<script>
import ProjectForm from './components/form'

export default {
  name: 'CreateProject',
  components: { ProjectForm }

}
</script>
