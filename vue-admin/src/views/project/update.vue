<template>
  <project-form :id=parseInt(this.$route.params.id) :readonly="false" :opType="'update'" />
</template>

<script>
import ProjectForm from './components/form'

export default {
  name: 'EditProject',
  components: { ProjectForm }
}
</script>
